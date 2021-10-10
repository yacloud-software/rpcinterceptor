package main

import (
	"flag"
	"fmt"
	apb "golang.conradwood.net/apis/auth"
	"golang.conradwood.net/apis/common"
	el "golang.conradwood.net/apis/errorlogger"
	rc "golang.conradwood.net/apis/rpcaclapi"
	ic "golang.conradwood.net/apis/rpcinterceptor"
	"golang.conradwood.net/evaluators"
	"golang.conradwood.net/go-easyops/auth"
	"golang.conradwood.net/go-easyops/authremote"
	gc "golang.conradwood.net/go-easyops/common"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/prometheus"
	"golang.conradwood.net/go-easyops/server"
	"golang.conradwood.net/go-easyops/sql"
	"golang.conradwood.net/go-easyops/tokens"
	"golang.conradwood.net/go-easyops/utils"
	"golang.conradwood.net/learn"
	"golang.conradwood.net/resolvers"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"os"
	"time"
)

var (
	use_errorlogger  = flag.Bool("use_error_logger", true, "use error logger")
	use_rules        = flag.Bool("use_rules", true, "feature flag: use rules to assign permissions")
	resolve_by_token = flag.Bool("resolve_by_token", true, "feature flag: resolve by service token if it is not provided in metadata")
	maintain_user_id = flag.Bool("maintain_userids", true, "maintain user id mapping between servicetokens and services")
	mock             = flag.Bool("mock", false, "'mock' - do not access database or dependencies, just say yes. Intended for local testing")
	allow_all        = flag.Bool("allow_all", false, "allow all access")
	debug            = flag.Bool("debug", false, "enable debug mode")
	trace            = flag.Bool("trace", false, "enable trace mode")
	port             = flag.Int("port", 10002, "The grpc server port")
	rev              *resolvers.Resolvers
	mockid           = uint64(2)
	rejReqs          = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "rpcinterceptor_rejected_requests",
			Help: "V=1 UNIT=ops DESC=number of requests rejected",
		},
		[]string{"rejectreason"},
	)
	authManager apb.AuthManagerServiceClient
	eval        evaluators.RuleEvaluator
)

type rpcInterceptorServer struct {
}

func main() {
	flag.Parse()
	prometheus.MustRegister(rejReqs)
	fmt.Printf("Starting RPCInterceptor service...\n")
	eval = evaluators.NewRuleEvaluator()
	if !*mock {
		start_requestid_creator()
		fmt.Printf("Starting dbtools()...\n")
		startdbtools()
	}
	sd := server.NewServerDef()
	sd.Port = *port
	sd.NoAuth = true // don't intercept ourselves or we loop like crazy ;)
	sd.Register = server.Register(
		func(server *grpc.Server) error {
			e := new(rpcInterceptorServer)
			ic.RegisterRPCInterceptorServiceServer(server, e)
			return nil
		},
	)
	err := server.ServerStartup(sd)
	utils.Bail("Unable to start server", err)
	os.Exit(0)
}

/************************************
* grpc functions
************************************/
type requestinfo struct {
	target_service *rc.Service
	target_method  *rc.Method
}

func (e *rpcInterceptorServer) GetServiceByUserID(ctx context.Context, req *ic.ServiceByUserIDRequest) (*ic.Service, error) {
	db, err := sql.Open()
	if err != nil {
		fmt.Printf("Failed to connect to db: %s\n", err)
		return nil, err
	}
	rows, err := db.QueryContext(ctx, "list_servicenames", "select id,servicename,userid from service where userid = $1", req.UserID)
	if err != nil {
		return nil, fmt.Errorf("Unable to select services from db: %s", err)
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, errors.NotFound(ctx, `service "%s" not found`, req.UserID)
	}
	res := &ic.Service{}
	err = rows.Scan(&res.ID, &res.Name, &res.UserID)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (e *rpcInterceptorServer) InterceptRPC(ctx context.Context, req *ic.InterceptRPCRequest) (*ic.InterceptRPCResponse, error) {
	if *mock {
		fmt.Printf("RPC response mocked (%s/%s()).\n", req.Service, req.Method)
		return &ic.InterceptRPCResponse{}, nil
	}
	ri := &requestinfo{}
	resp, err := ri.DoInterceptRPC(ctx, req)
	if err != nil {
		return resp, err
	}
	if resp != nil && resp.Reject {
		rejReqs.With(prometheus.Labels{"rejectreason": fmt.Sprintf("%s", resp.RejectReason)}).Inc()
		logError(ctx, req, resp, ri.target_service)
	}
	/*
		if resp != nil && !resp.Reject {
			if resp.CallerUser != nil && resp.SignedCallerUser == nil {
				panic("must not return v2 only (v3 required)")
			}
		}
	*/
	return resp, err
}
func (ri *requestinfo) DoInterceptRPC(ctx context.Context, req *ic.InterceptRPCRequest) (*ic.InterceptRPCResponse, error) {
	var err error
	if !*mock && rev == nil {
		authManager = authremote.GetAuthManagerClient()
		rev, err = resolvers.NewResolver()
		if err != nil {
			return nil, err
		}
		fmt.Printf("Got resolver...\n")
	}

	if req.InMetadata == nil {
		if !*allow_all {
			req.InMetadata = &ic.InMetadata{}
		} else {
			return nil, errors.Error(ctx, codes.PermissionDenied, "access denied", "a context was submitted to rpcinterceptor without any metadata from %s.%s", req.Service, req.Method)
		}
	}
	rid := req.InMetadata.RequestID
	if rid == "" {
		rid = req_get_requestid()
	}
	csvc := rev.GetServiceByID(ctx, req.InMetadata.CallerServiceID)
	var csvcname string
	if csvc != nil {
		csvcname = csvc.Name
	} else {
		csvcname = fmt.Sprintf("%d", req.InMetadata.CallerServiceID)
	}
	debugmsg := fmt.Sprintf("[Request \"%s\" from %s to %s.%s] ", rid, csvcname, req.Service, req.Method)

	var submeth *rc.Method
	submid := uint64(0)
	// methodid from req.InMetadata:
	submid = req.InMetadata.CallerMethodID
	if submid != 0 {
		submeth = rev.GetMethodByID(ctx, submid)
		//fmt.Printf("SubMeth: %#v\n", submeth)
	}

	if *debug {
		fmt.Printf("%sIntercepted %s.%s\n", debugmsg, req.Service, req.Method)
		// they do have access, print debug stuff and return:
		fmt.Printf("%sInterceptRequest: %#v\n", debugmsg, req)
		fmt.Printf("%sIC: %#v\n", debugmsg, req.InMetadata)
		if submeth == nil {
			fmt.Printf("%sSubmitted calling method: [NONE] (id=%d)\n", debugmsg, submid)
		} else {
			fmt.Printf("%sSubmitted calling method: %2d service=%2d (%s)\n", debugmsg, submeth.ID, submeth.ServiceID, submeth.Name)
		}
	}
	resp := &ic.InterceptRPCResponse{
		RequestID:    rid,
		Reject:       true,
		RejectReason: ic.RejectReason_NonSpecific,
	}
	if *mock {
		resp.Reject = false
		resp.CallerMethodID = newMockID()
		if *debug {
			fmt.Printf("Returning MOCKED: %#v\n", resp)
		}
		return resp, nil
	}

	// resolve the tokens to users/services:
	//	resp.SignedCallerUser, err = rev.SignedGetUserByToken(ctx, req.InMetadata.UserToken) // returns nil (no error) if token==""
	ru, err := rev.SignedGetUserByToken(ctx, req.InMetadata.UserToken) // returns nil (no error) if token==""
	if err != nil {
		fmt.Printf("%sfailed to authenticate user: %s\n", debugmsg, err)
		resp.RejectReason = ic.RejectReason_UserRejected
		learn.Learn(rev, req, resp)
		return resp, nil
	}
	if ru != nil {
		resp.CallerUser = ru.User()
		resp.SignedCallerUser = ru.SignedUser()
	}
	// if we cannot resolve user via token, check for a signed userobject
	if resp.CallerUser == nil {
		su := gc.VerifySignedUser(req.InMetadata.SignedUser)
		if su != nil {
			// we got a v2signed user
			resp.SignedCallerUser = req.InMetadata.SignedUser
			resp.CallerUser = su
		} else if gc.VerifySignature(req.InMetadata.User) {
			// we only got a v1 signed user
			ru, err := rev.AsRootGetUserByID(ctx, req.InMetadata.User.ID)
			if err != nil {
				fmt.Printf("Failed to get user: %s\n", err)
				return nil, err
			}
			if ru != nil {
				resp.CallerUser = ru.User()
				resp.SignedCallerUser = ru.SignedUser()
			}

		}
	}
	if resp.CallerUser == nil {
		if *trace {
			fmt.Printf("%sthis call has no user\n", debugmsg)
		}
	}
	rs, err := rev.SignedGetUserByToken(ctx, req.InMetadata.ServiceToken)
	if err != nil {
		fmt.Printf("%sfailed to authenticate service: %s\n", debugmsg, err)
		resp.RejectReason = ic.RejectReason_ServiceRejected
		learn.Learn(rev, req, resp)
		return resp, nil
	}
	if rs != nil {
		resp.CallerService = rs.User()
		resp.SignedCallerService = rs.SignedUser()
	}
	// if we cannot resolve service via token, check for a signed userobject
	if resp.CallerService == nil {
		s := gc.VerifySignedUser(req.InMetadata.SignedService)
		if s != nil {
			resp.SignedCallerService = req.InMetadata.SignedService
			resp.CallerService = s
		} else {
			if gc.VerifySignature(req.InMetadata.Service) {
				// accept it
				resp.CallerService = req.InMetadata.Service
			}
		}
	}
	/*
		from := udShort(resp.CallerService) + "/" + udShort(resp.CallerUser)
		debugmsg = fmt.Sprintf("[Request \"%s\" from %s to %s.%s] ", rid, from, req.Service, req.Method)
	*/
	// sometimes we don't have a usertoken, but we have a service and a userid. if so,
	// resolve it
	if req.InMetadata.UserID != "" && resp.CallerUser == nil && resp.CallerService != nil {
		// TODO: investigate if this is right.
		ru, err := rev.AsRootGetUserByID(ctx, req.InMetadata.UserID)
		if err != nil {
			if *debug {
				fmt.Printf("Failed to lookup userid \"%s\":%s\n", req.InMetadata.UserID, err)
			}
			return nil, err
		}
		if ru != nil {
			resp.CallerUser = ru.User()
			resp.SignedCallerUser = ru.SignedUser()
		}

	}

	// do we have a user? if so stick into variable
	uid := ""
	if resp.CallerUser != nil {
		uid = resp.CallerUser.ID
	}

	// get serviceuser of caller (if any)
	sid := ""
	from := "[NOSERVICE]" + "/" + udShort(resp.CallerUser)
	if resp.CallerService != nil {
		from = udShort(resp.CallerService) + "/" + udShort(resp.CallerUser)
		sid = resp.CallerService.ID // would prefer the serviceID here... hm..
	}
	if resp.CallerService == nil && resp.CallerUser == nil {
		from = "[" + req.Source + "]"
	}
	debugmsg = fmt.Sprintf("[Request \"%s\" from %s to %s.%s] ", rid, from, req.Service, req.Method)

	// mid == ID of method being invoked
	mid := findMethodID(ctx, req.Service, req.Method)
	// resolved...
	if *debug {
		fmt.Printf("%sidentified user=%s, service=%s (method #%d) [%s.%s]\n", debugmsg, udShort(resp.CallerUser), udShort(resp.CallerService), mid, req.Service, req.Method)
	}
	resp.CallerMethodID = mid

	// which service is calling the intercepted rpc (the CALLEE)
	resp.CalleeServiceID = req.InMetadata.CallerServiceID
	if resp.CalleeServiceID == 0 && *resolve_by_token {
		sv := rev.GetServiceByToken(ctx, req.InMetadata.ServiceToken)
		if *debug {
			fmt.Printf("%sServicetoken user: %s\n", debugmsg, sv)
		}
		if sv != nil {
			resp.CalleeServiceID = sv.ID
		}
	}

	// now check if user/services/org have access to this service/method
	err = ri.setAccess(ctx, req, resp, mid) // will set resp.Reject
	if err != nil {
		fmt.Printf("%sWhilst checking access, error occured: %s\n", debugmsg, err)
		return nil, err
	}
	if *allow_all {
		resp.Reject = false
	}

	if !resp.Reject {
		// "audit"
		err = Audit(ctx, mid, sid, uid, submeth)
		if err != nil {
			if *debug {
				fmt.Printf("Error during audit: %s\n", err)
			}
			return nil, err
		}
	}
	err = storeLog(ctx, req, resp, mid, submid)
	if err != nil {
		fmt.Printf("Failed to store log: %s\n", err)
	}

	if resp.Reject {
		// setAccess did not allow access
		fmt.Printf("%sAccess denied for user=%s, service=%s (%s)\n", debugmsg, auth.Description(resp.CallerUser), auth.Description(resp.CallerService), resp.RejectReason)
		return resp, nil
	}
	if *trace {
		fmt.Printf("%sAccess granted\n", debugmsg)
	}
	go learn.Learn(rev, req, resp)
	return resp, nil
}

// userdetail to short string
func udShort(det *apb.User) string {
	if det == nil {
		return "[NOUSER]"
	}
	if det.Abbrev != "" {
		return fmt.Sprintf("%s(%s)", det.Abbrev, det.ID)
	}
	return fmt.Sprintf("%s(%s)", det.Email, det.ID)

}

// given a request and response will set the "Reject" flag and reason appropriately.
// this means specifically, if "Reject" flag is set after this function returns, the
// call is not authorised
func (ri *requestinfo) setAccess(ctx context.Context, req *ic.InterceptRPCRequest, resp *ic.InterceptRPCResponse, methodID uint64) error {
	ri.target_method = rev.GetMethodByID(ctx, methodID)
	if ri.target_method == nil {
		return fmt.Errorf("No method %d", methodID)
	}
	ri.target_service = rev.GetServiceByID(ctx, ri.target_method.ServiceID)
	if *debug {
		fmt.Printf("ServiceID: %d\n", ri.target_service.ID)
	}
	// TODO - remove this
	if auth.IsRootUser(resp.CallerUser) {
		resp.Reject = false
		return nil
	}
	if resp.CallerService != nil {
		resp.Reject = false
		return nil
	}
	// TOOD - use rules only

	b, err := eval.EvaluateCall(ri.target_service, resp.CallerService, resp.CallerUser)
	if err != nil {
		fmt.Printf("Error evaluating call: %s\n", err)
	} else {
		if *use_rules {
			// eventually we can accept this as the whole truth (thus removing "if b" above)
			// also note the reverse here. evaluate returns true if call is to be accepted
			resp.Reject = !b
			return nil
		}
	}

	if resp.CallerUser != nil {
		// anyone may call auth.whoami()
		if req.Service == "auth.AuthManagerService" && req.Method == "WhoAmI" {
			resp.Reject = false
			return nil
		}
		if *debug {
			fmt.Printf("Not accepting user %#v because he's not a root group member\n", auth.Description(resp.CallerUser))
			for i, g := range resp.CallerUser.Groups {
				fmt.Printf("%2d. Group: %s (%s)\n", (i + 1), g.Name, g.ID)
			}
		}

	} else {
		if *debug {
			fmt.Printf("No user provided\n")
		}
	}
	resp.Reject = true
	resp.RejectReason = ic.RejectReason_ServiceMissing
	return nil

}

func findMethodID(ctx context.Context, sn string, mn string) uint64 {
	m := rev.GetMethodByName(ctx, sn, mn)
	if m != nil {
		return m.ID
	}
	// method is new - create it
	err := logService(ctx, sn, mn)
	if err != nil {
		fmt.Printf("Failed to create service %s and method %s: %s\n", sn, mn, err)
	}
	// retry getting it
	rev.Clear()
	m = rev.GetMethodByName(ctx, sn, mn)
	if m != nil {
		return m.ID
	}
	fmt.Printf("Retry from cache failed too for %s.%s\n", sn, mn)
	return 0
}

func newMockID() uint64 {
	mockid++
	return mockid
}

/**********************************************************************
* start up new servers - they need their serviceid to report them to us
**********************************************************************/
func (r *rpcInterceptorServer) GetMyServiceID(ctx context.Context, req *ic.ServiceIDRequest) (*ic.ServiceIDResponse, error) {
	fmt.Printf("Request to resolve token %s from %s to a serviceid \n", req.Token, req.MyName)
	ru, err := rev.SignedGetUserByToken(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	u := ru.User()
	if u == nil {
		return nil, fmt.Errorf("Your token does not resolve to a valid user")
	}
	if !u.ServiceAccount {
		return nil, fmt.Errorf("User %s is not a service", auth.Description(u))
	}
	svc, err := getOrCreateService(ctx, req.MyName, u.ID)
	if err != nil {
		return nil, fmt.Errorf("unable to resolve service: \"%s\" (%s)", req.MyName, err)
	}
	if svc == 0 {
		return nil, fmt.Errorf("no such service: \"%s\"", req.MyName)
	}
	if !*mock && *maintain_user_id {
		sv := rev.GetServiceByID(ctx, svc)
		MaintainUserID(ctx, sv, u)
		if err != nil {
			fmt.Printf("Error maintaining userid: %s\n", err)
			// purposefully ignore error here
		}
		if sv != nil {
			fmt.Printf("Service userid: %s\n", sv.UserID)
		}
	}
	fmt.Printf("Service %s resolved to serviceID: %d\n", req.MyName, svc)
	return &ic.ServiceIDResponse{ServiceID: svc}, nil
}

func MaintainUserID(ctx context.Context, svc *rc.Service, user *apb.User) error {
	if svc.UserID != "" {
		return nil // fast-path, normally, do nothing
	}
	if user == nil {
		fmt.Printf("WARNING: MaintainUserID() called with no user\n")
		return nil
	}
	svc.UserID = user.ID
	db, err := sql.Open()
	if err != nil {
		return err
	}
	_, err = db.ExecContext(ctx, "update_servicename", "update service set userid = $1 where id = $2", svc.UserID, svc.ID)
	if err != nil {
		return err
	}
	return nil
}

/**********************************************************************
* expose learnings
**********************************************************************/

func (e *rpcInterceptorServer) ClearLearnings(ctx context.Context, req *common.Void) (*common.Void, error) {
	learn.Reset()
	return &common.Void{}, nil
}
func (e *rpcInterceptorServer) GetLearnings(ctx context.Context, req *common.Void) (*ic.Learnings, error) {
	// cannot protect this, since it is not intercepted
	/*
		if !auth.IsRoot(ctx) {
			return nil, errors.AccessDenied(ctx, "GetLearnings() for root only")
		}
	*/
	res := &ic.Learnings{}
	res.Learnings = learn.Learnings()
	rev.Clear()
	return res, nil
}

func logError(ctx context.Context, req *ic.InterceptRPCRequest, resp *ic.InterceptRPCResponse, svc *rc.Service) {
	if !*use_errorlogger {
		return
	}
	if req == nil {
		return
	}
	service_id := "undefined"
	if svc != nil {
		service_id = fmt.Sprintf("%d", svc.ID)
	}

	user := resp.CallerUser
	uid := ""
	userdesc := "[nouser]"
	if user != nil {
		uid = user.ID
		userdesc = fmt.Sprintf("[user #%s(%s)]", user.ID, user.Email)
	}
	servicedesc := "[noservice]"
	reqid := ""
	if req.InMetadata != nil {
		reqid = req.InMetadata.RequestID
		cs := req.InMetadata.Service
		if cs != nil {
			servicedesc = fmt.Sprintf("[service #%s(%s)]", cs.ID, cs.Email)
		}
	}

	e := &el.ErrorLogRequest{
		UserID:       uid,
		ServiceName:  req.Service,
		MethodName:   req.Method,
		Timestamp:    uint32(time.Now().Unix()),
		ErrorCode:    uint32(codes.PermissionDenied),
		ErrorMessage: "access denied",
		LogMessage:   fmt.Sprintf("rpcinterceptor denied access to service (rpc_service_id=%s) from %s %s", service_id, userdesc, servicedesc),
		RequestID:    reqid,
	}
	ctx = tokens.ContextWithToken()
	_, err := el.GetErrorLoggerClient().Log(ctx, e)
	if err != nil {
		fmt.Printf("unable to log error: %s\n", err)
	}
}
