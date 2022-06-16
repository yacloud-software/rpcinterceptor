package main

import (
	"flag"
	"fmt"
	apb "golang.conradwood.net/apis/auth"
	"golang.conradwood.net/apis/common"
	ic "golang.conradwood.net/apis/rpcaclapi"
	rc "golang.conradwood.net/apis/rpcinterceptor"
	"golang.conradwood.net/rpcinterceptor/evaluators"
	"golang.conradwood.net/go-easyops/auth"
	"golang.conradwood.net/go-easyops/cache"
	"golang.conradwood.net/go-easyops/client"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/server"
	"golang.conradwood.net/go-easyops/sql"
	"golang.conradwood.net/go-easyops/utils"
	"golang.conradwood.net/rpcinterceptor/resolvers"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"os"
	"time"
)

var (
	authServer  apb.AuthenticationServiceClient
	authManager apb.AuthManagerServiceClient
	debug       = flag.Bool("debug", false, "enable debug mode")
	port        = flag.Int("port", 10005, "The grpc server port")
	tokenCache  *cache.Cache
	rev         *resolvers.Resolvers
	dbrules     *evaluators.DBRules
)

type rpcACLServer struct {
}

func main() {
	flag.Parse()
	evaluators.StartDB()
	dbrules = &evaluators.DBRules{}
	var err error
	fmt.Printf("Starting RPCInterceptor service...\n")
	fmt.Printf("Starting auth client...\n")

	// we really have a hard dependency on the auth client.
	// we can't start if we don't have auth.
	authServer = apb.NewAuthenticationServiceClient(client.Connect("auth.AuthenticationService"))
	authManager = apb.NewAuthManagerServiceClient(client.Connect("auth.AuthManagerService"))
	rev, err = resolvers.NewResolver() // should NEVER return an error
	// (because the only possible error is "auth*service unavailable", but if it were the previous two lines would blocK)
	utils.Bail("failed to create resolver", err)
	tokenCache = cache.New("rpcacl_tokencache", time.Duration(5*60)*time.Second, 100)
	sd := server.NewServerDef()
	sd.Port = *port
	sd.Register = server.Register(
		func(server *grpc.Server) error {
			e := new(rpcACLServer)
			ic.RegisterRPCACLServiceServer(server, e)
			return nil
		},
	)
	err = server.ServerStartup(sd)
	utils.Bail("Unable to start server", err)
	os.Exit(0)
}

/************************************
* grpc functions
************************************/

func (e *rpcACLServer) GetMostRecentLogs(ctx context.Context, req *ic.LogEntryRequest) (*ic.LogEntryList, error) {
	db, err := sql.Open()
	if err != nil {
		fmt.Printf("Failed to connect to db: %s\n", err)
		return nil, err
	}
	me := req.MaxEntries
	if me == 0 {
		me = 10 // default
	}
	rows, err := db.QueryContext(ctx, "list_recent_logs", "select occured,oreq,method_id,callerservice,calleruserid,reject,rejectreason from logentry order by id desc limit $1", me)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	res := &ic.LogEntryList{}
	for rows.Next() {
		var mid uint64
		var callerservice string
		var calleruserid string
		le := &ic.LogEntry{Response: &rc.InterceptRPCResponse{}}
		err = rows.Scan(&le.Timestamp, &le.Response.RequestID, &mid, &callerservice, &calleruserid, &le.Response.Reject, &le.Response.RejectReason)
		if err != nil {
			return nil, fmt.Errorf("Failed to scan logentries: %s", err)
		}
		m := rev.GetMethodByID(ctx, mid)
		if m == nil {
			fmt.Printf("Skippig logentry - No such method: %d\n", mid)
			continue
		}
		le.Method = m
		le.Service = rev.GetServiceByID(ctx, m.ServiceID)
		if callerservice != "" {
			le.Response.CallerService, _ = rev.GetUserByID(ctx, callerservice)
		}
		if calleruserid != "" {
			le.Response.CallerUser, _ = rev.GetUserByID(ctx, calleruserid)
		}
		res.Entries = append(res.Entries, le)
	}
	return res, nil
}

// get logs for calling user
func (e *rpcACLServer) GetMyMostRecentLogs(ctx context.Context, req *ic.LogEntryRequest) (*ic.LogEntryList, error) {
	db, err := sql.Open()
	if err != nil {
		fmt.Printf("Failed to connect to db: %s\n", err)
		return nil, err
	}
	user := auth.GetUser(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}
	rows, err := db.QueryContext(ctx, "select_recent_logs_for_self", "select occured,oreq,method_id,callerservice,calleruserid,reject,rejectreason from logentry where calleruserid = $1 order by id desc limit $2", user.ID, req.MaxEntries)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	res := &ic.LogEntryList{}
	for rows.Next() {
		var mid uint64
		var callerservice string
		var calleruserid string
		le := &ic.LogEntry{Response: &rc.InterceptRPCResponse{}}
		err = rows.Scan(&le.Timestamp, &le.Response.RequestID, &mid, &callerservice, &calleruserid, &le.Response.Reject, &le.Response.RejectReason)
		if err != nil {
			return nil, fmt.Errorf("Failed to scan logentries: %s", err)
		}
		m := rev.GetMethodByID(ctx, mid)
		if m == nil {
			fmt.Printf("Skippig logentry - No such method: %d\n", mid)
			continue
		}
		le.Method = m
		le.Service = rev.GetServiceByID(ctx, m.ServiceID)
		if callerservice != "" {
			le.Response.CallerService, _ = rev.GetUserByID(ctx, callerservice)
		}
		if calleruserid != "" {
			le.Response.CallerUser, _ = rev.GetUserByID(ctx, calleruserid)
		}
		res.Entries = append(res.Entries, le)
	}
	return res, nil
}
func (e *rpcACLServer) ListCalls(ctx context.Context, req *common.Void) (*ic.CallList, error) {
	db, err := sql.Open()
	if err != nil {
		fmt.Printf("Failed to connect to db: %s\n", err)
		return nil, err
	}
	rows, err := db.QueryContext(ctx, "select_calls_to_method", "select methodid,serviceuserid,userid from methodaudit order by methodid")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	res := &ic.CallList{}
	for rows.Next() {
		c := &ic.Call{}
		var mid uint64
		var suid string
		var userid string
		err = rows.Scan(&mid, &suid, &userid)
		if err != nil {
			return nil, err
		}
		m := rev.GetMethodByID(ctx, mid)
		if m != nil {
			fm := &ic.FullMethod{ID: mid, Name: m.Name}
			fm.Service = rev.GetServiceByID(ctx, m.ServiceID)
			c.CalledMethod = fm
			res.Calls = append(res.Calls, c)
			c.CallingService, _ = rev.GetUserByID(ctx, suid)
			c.CallingUser, _ = rev.GetUserByID(ctx, userid)
		}
	}
	return res, nil
}

func (e *rpcACLServer) ServiceNameToID(ctx context.Context, req *ic.ServiceNameRequest) (*ic.ServiceIDResponse, error) {
	if !auth.IsRoot(ctx) {
		return nil, errors.AccessDenied(ctx, "access denied to resolve service names")
	}
	db, err := sql.Open()
	if err != nil {
		fmt.Printf("Failed to connect to db: %s\n", err)
		return nil, err
	}
	user := auth.GetUser(ctx)
	if user == nil {
		return nil, errors.Unauthenticated(ctx, "access denied")
	}
	rows, err := db.QueryContext(ctx, "select_serviceid", "select id from service where servicename = $1", req.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, errors.InvalidArgs(ctx, "no such service", "no such service: %s", req.Name)
	}
	res := &ic.ServiceIDResponse{}
	err = rows.Scan(&res.ID)
	if err != nil {
		return nil, err
	}
	return res, nil

}
