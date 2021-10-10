package main

import (
	"flag"
	"fmt"
	"golang.conradwood.net/go-easyops/prometheus"
	rc "golang.conradwood.net/apis/rpcaclapi"
	"golang.conradwood.net/go-easyops/sql"
	"golang.conradwood.net/go-easyops/tokens"
	"golang.conradwood.net/go-easyops/utils"
	"golang.org/x/net/context"
	"sync"
	"time"
)

var (
	db_flush_cycle = flag.Int("acl_flush_delay", 5, "minutes between flushing acls to db")
	enable_audit   = flag.Bool("enable_audit", false, "Enable auditing of method access")
	knownACLS      = map[string]*knownACL{}
	checkLock      sync.Mutex
	aclSum         = prometheus.NewSummary(prometheus.SummaryOpts{
		Name:       "rpcinterceptor_acl_performance",
		Help:       "V=1 UNIT=durationms DESC=summary of duration for lookups of acls",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.95: 0.005, 0.99: 0.001},
	})
	aclSizeGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "rpcinterceptor_size_acls",
			Help: "V=1 UNIT=short DESC=number of acl combinations (method/service/user)",
		},
		[]string{},
	)
)

type knownACL struct {
	id     uint32
	method uint64
	sid    string
	uid    string
	source uint64
	calls  uint32
	indb   bool
}

func init() {
	prometheus.MustRegister(aclSum, aclSizeGauge)
	go audit_db_loop()
}

func aclKey(m uint64, service string, u string, source uint64) string {
	return fmt.Sprintf("%d-%s-%s-%d", m, service, u, source)
}

// purpose of this package is to log access by whom to what.
// we keep a delay so not to update the DB on *every* call

// a method maybe invoked from another service or user or with both
// here we keep a history of how methods were invoked.
// eventually this should build up an entire database of
// current, implicit ACLs
func Audit(ctx context.Context, method uint64, serviceuserid string, userid string, rpc *rc.Method) error {
	if !*enable_audit {
		return nil
	}
	var source uint64
	if rpc != nil {
		source = rpc.ID
	}
	// safety valve:
	l := len(knownACLS)
	if l > 10000 {
		return nil
	}
	aclSizeGauge.With(prometheus.Labels{}).Set(float64(l))
	start := time.Now()
	checkLock.Lock()
	defer checkLock.Unlock()
	if *debug {
		fmt.Printf("Source of rpc call %d\n", source)
	}
	key := aclKey(method, serviceuserid, userid, source)
	if k, exists := knownACLS[key]; exists {
		k.calls++
		aclSum.Observe(time.Since(start).Seconds())
		return nil
	}
	ka := &knownACL{method: method, sid: serviceuserid, uid: userid, source: source, indb: false}
	if *debug {
		fmt.Printf("New entry: %v\n", ka)
	}
	knownACLS[key] = ka
	aclSum.Observe(time.Since(start).Seconds())
	return nil
}

func audit_db_loop() {
	for {
		if *debug {
			fmt.Printf("Storing acl entries found\n")
		}
		audit_db()
		utils.RandomStall(*db_flush_cycle)
	}
}

func audit_db() {
	for _, a := range knownACLS {
		if a.id > 0 || entryExists(a) {
			if *debug {
				fmt.Printf("Updating entry id %d\n", a.id)
			}
			updateEntry(a)
			continue
		}
		if *debug {
			fmt.Printf("Saving new entry %v\n", a)
		}
		saveEntry(a)
	}
}

func entryExists(a *knownACL) bool {
	db, err := sql.Open()
	if err != nil {
		fmt.Printf("audit1: Failed to connect to db: %s\n", err)
		return false
	}
	rows, err := db.QueryContext(tokens.ContextWithToken(), "select_audit_entry", "select id from methodaudit where methodid = $1 and serviceuserid = $2 and userid = $3 and sourceid = $4", a.method, a.sid, a.uid, a.source)
	if err != nil {
		fmt.Printf("could not select methodaudit row: %s\n", err)
		return false
	}
	if rows.Next() {
		err = rows.Scan(&a.id)
		return err == nil
	}
	return false
}

func saveEntry(a *knownACL) {
	db, err := sql.Open()
	if err != nil {
		fmt.Printf("audit2: Failed to connect to db: %s\n", err)
		return
	}
	_, err = db.ExecContext(tokens.ContextWithToken(), "insert_audit_entry", "insert into methodaudit (methodid,serviceuserid,userid,sourceid,calls) values ($1,$2,$3,$4,$5)", a.method, a.sid, a.uid, a.source, a.calls)
	if err != nil {
		fmt.Printf("could not insert methodaudit row: %s\n", err)
	}
}

func updateEntry(a *knownACL) {
	db, err := sql.Open()
	if err != nil {
		fmt.Printf("audit3: Failed to connect to db: %s\n", err)
		return
	}
	_, err = db.ExecContext(tokens.ContextWithToken(), "update_audit_entry", "update methodaudit set calls = calls + $1 where id = $2", a.calls, a.id)
	if err != nil {
		fmt.Printf("could not update methodaudit row: %s\n", err)
	}
}
