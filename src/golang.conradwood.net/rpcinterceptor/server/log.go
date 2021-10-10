package main

import (
	"flag"
	"fmt"
	ic "golang.conradwood.net/apis/rpcinterceptor"
	"golang.conradwood.net/go-easyops/cache"
	"golang.conradwood.net/go-easyops/sql"
	"golang.conradwood.net/go-easyops/utils"
	"golang.org/x/net/context"
	"time"
)

var (
	clean_delay  = flag.Int("clean_interval", 10, "interval in `minutes` in which to remove old entries from the database")
	clean_old    = flag.Int64("clean_is_old", 120, "age of an entry in `minutes` until it is removed from the database")
	log_calls    = flag.Bool("log_calls", true, "log all calls to the database (expect high volume (1 query of sum(all_rpc)))")
	smCache      *cache.Cache
	serviceCache *cache.Cache
)

func storeLog(ctx context.Context, req *ic.InterceptRPCRequest, resp *ic.InterceptRPCResponse, methodID uint64, callingmethodid uint64) error {
	if *log_calls {
		im := req.InMetadata
		if im == nil {
			im = &ic.InMetadata{}
		}
		db, err := sql.Open()
		if err != nil {
			fmt.Printf("log1: Failed to connect to db: %s\n", err)
			return err
		}
		cs := ""
		if resp.CallerService != nil {
			cs = resp.CallerService.ID
		}
		cu := ""
		if resp.CallerUser != nil {
			cu = resp.CallerUser.ID
		}
		occ := time.Now().Unix()
		rej := 0
		if resp.Reject {
			rej = 1
		}

		_, err = db.ExecContext(ctx, "insert_logentry", "insert into logentry (occured,ireq,oreq,method_id,callerservice,calleruserid,reject,rejectreason,callingmethod) values ($1,$2,$3,$4,$5,$6,$7,$8,$9)",
			occ,
			im.RequestID,
			resp.RequestID,
			methodID,
			cs,
			cu,
			rej,
			int(resp.RejectReason),
			callingmethodid,
		)
		if err != nil && *debug {
			fmt.Printf("log2: Failed to insert log entry: %s. request: %#v\n,response: %#v\n", err, req, resp)
		}
		if err != nil {
			return fmt.Errorf("Failed to insert log entry: %s", err)
		}
	}
	e := logService(ctx, req.Service, req.Method)
	if e != nil {
		fmt.Printf("log3: Failed to create \"%s.%s\": %s\n", req.Service, req.Method, e)
	}
	return nil
}

func cleandb_loop() {
	for {
		cleandb()
		clean_error_db()
		utils.RandomStall(*clean_delay)
	}
}

func cleandb() {
	now := time.Now().Unix()
	now = now - 60*(*clean_old)
	db, err := sql.Open()
	if err != nil {
		fmt.Printf("log4: Failed to connect to db: %s\n", err)
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	_, err = db.ExecContext(ctx, "deletelogentries", "delete from logentry where occured < $1", now)

	// 	_, err = db.Exec("delete from logentry where occured < $1", now)
	if err != nil {
		fmt.Printf("Cleaner failed: %s\n", err)
	}
}
func startdbtools() {
	smCache = cache.New("servicemethodcache", time.Duration(4)*time.Hour, 100)
	serviceCache = cache.New("servicecache2", time.Duration(4)*time.Hour, 100)
	go cleandb_loop()
}

func logService(ctx context.Context, service string, method string) error {
	hash := service + method
	if smCache.Get(hash) != nil {
		return nil
	}
	id, err := getOrCreateService(ctx, service, "")
	if err != nil {
		return err
	}
	if id == 0 {
		return fmt.Errorf("NO ID for service %s", service)
	}
	mid, err := getOrCreateMethod(ctx, id, method)
	if err != nil {
		return err
	}
	mc := &MethodCacheEntry{ID: mid, ServiceID: id}
	smCache.Put(hash, mc)
	return nil
}

type ServiceCacheEntry struct {
	ID uint64
}

func getOrCreateService(ctx context.Context, service string, userid string) (uint64, error) {
	db, err := sql.Open()
	if err != nil {
		fmt.Printf("log5: Failed to connect to db: %s\n", err)
		return 0, err
	}
	x := serviceCache.Get(service)
	if x != nil {
		xs := x.(*ServiceCacheEntry)
		return xs.ID, nil
	}
	rows, err := db.QueryContext(ctx, "select_service", "select id from service where servicename = $1", service)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	if rows.Next() {
		ns := &ServiceCacheEntry{}
		err = rows.Scan(&ns.ID)
		if err != nil {
			return 0, err
		}
		serviceCache.Put(service, ns)
		return ns.ID, nil
	}
	rows, err = db.QueryContext(ctx, "create_service", "insert into service (lastseen,servicename,userid) values($1,$2,$3) returning ID", time.Now().Unix(), service, userid)
	if err != nil {
		return 0, fmt.Errorf("Error adding service %s to database: %s", service, err)
	}
	defer rows.Close()
	if !rows.Next() {
		return 0, fmt.Errorf("unable to insert service (no rows)")
	}
	var id uint64
	err = rows.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve row for service %s: %s", service, err)
	}
	return id, nil

}

type MethodCacheEntry struct {
	ID        uint64
	ServiceID uint64
}

func getOrCreateMethod(ctx context.Context, serviceid uint64, method string) (uint64, error) {
	db, err := sql.Open()
	if err != nil {
		fmt.Printf("log6: Failed to connect to db: %s\n", err)
		return 0, err
	}
	rows, err := db.QueryContext(ctx, "get_method", "select id from method where methodname =  $1 and service_id = $2", method, serviceid)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	if rows.Next() {
		var id uint64
		err = rows.Scan(&id)
		if err != nil {
			return 0, err
		}
		return id, nil
	}
	rows, err = db.QueryContext(ctx, "insert_method", "insert into method (lastseen,service_id,methodname) values($1,$2,$3) returning ID", time.Now().Unix(), serviceid, method)
	if err != nil {
		return 0, fmt.Errorf("Error adding method %s to database: %s", method, err)
	}
	defer rows.Close()
	if !rows.Next() {
		return 0, fmt.Errorf("unable to insert method (no rows)")
	}
	var id uint64
	err = rows.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve row for method %s: %s", method, err)
	}
	setDefaultPermissions(ctx, id)
	return id, nil
}

func setDefaultPermissions(ctx context.Context, method_id uint64) {
	db, err := sql.Open()
	if err != nil {
		fmt.Printf("log7: Failed to connect to db: %s\n", err)
		return
	}

	groupid := "34"  // developers
	creator := "639" // rpcinterceptor service account
	_, err = db.QueryContext(ctx, "insert_methodacl", "insert into methodacl (created_at,created_by,method_id,group_id) values ($1,$2,$3,$4)",
		time.Now().Unix(),
		creator,
		method_id, groupid)
	if err != nil {
		fmt.Printf("Failed to set default permissions: %s\n", err)
		// print & ignore
	}
}
