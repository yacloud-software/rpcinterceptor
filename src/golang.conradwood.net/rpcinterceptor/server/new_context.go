package main

import (
	"context"
	"flag"
	"fmt"
	apb "golang.conradwood.net/apis/auth"
	ic "golang.conradwood.net/apis/rpcinterceptor"
	"golang.conradwood.net/go-easyops/auth"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/cache"
	"golang.conradwood.net/go-easyops/common"
	"golang.conradwood.net/go-easyops/ctx"
	"time"
)

var (
	print_callers = flag.Bool("print_callers", false, "if true print callers from request(payload)")
	token_cache   = cache.New("token_cache", time.Duration(60)*time.Minute, 10000)
)

func deprecate_rpc_interceptor(ictx context.Context, req *ic.InterceptRPCRequest) (*ic.InterceptRPCResponse, error) {
	ls := ctx.GetLocalState(ictx)
	if *print_callers || *debug {
		fmt.Printf("Intercept from %s/%s\n", req.Service, req.Method)
	}
	if *debug {
		fmt.Printf("Request: %#v\n", req)
		m := req.InMetadata
		if m != nil {
			u := auth.UserIDString(m.User)
			s := auth.UserIDString(m.Service)
			fmt.Printf("Request.Meta: ServiceToken: %s, (user:%s,service:%s)\n", m.ServiceToken, u, s)
		}
		//fmt.Printf("Context: %s\n", ictx)
		fmt.Printf("Creating response from context %s\n", ctx.Context2String(ictx))
	}
	res := &ic.InterceptRPCResponse{
		RequestID:           "rpcinterceptor_obsolete",
		Reject:              false,
		CallerMethodID:      1,
		Source:              "nosource parsing in rpcinterceptor",
		CalleeServiceID:     1,
		CallerService:       common.VerifySignedUser(ls.CallingService()),
		CallerUser:          common.VerifySignedUser(ls.User()),
		SignedCallerService: ls.CallingService(),
		SignedCallerUser:    ls.User(),
	}

	// if we do not have the information, try from context.metadata
	mt := req.InMetadata
	if mt != nil {
		if res.SignedCallerService == nil {
			res.SignedCallerService = mt.SignedService
			res.CallerService = common.VerifySignedUser(res.SignedCallerService)
		}
		if res.SignedCallerUser == nil {
			res.SignedCallerUser = mt.SignedUser
			res.CallerUser = common.VerifySignedUser(res.SignedCallerUser)
		}

		// if we still do not have the information, try from context.metadata.servicetoken
		if res.SignedCallerService == nil && mt.ServiceToken != "" {
			csu := token_cache.Get(mt.ServiceToken)
			var su *apb.SignedUser
			if csu != nil {
				su = csu.(*apb.SignedUser)
			} else {
				su = authremote.SignedGetByToken(ictx, mt.ServiceToken)
				token_cache.Put(mt.ServiceToken, su)
			}
			res.SignedCallerService = su
			res.CallerService = common.VerifySignedUser(res.SignedCallerService)
		}
	}

	if *debug {
		fmt.Printf("Response with user %s and service %s\n", auth.UserIDString(res.CallerUser), auth.UserIDString(res.CallerService))
	}
	return res, nil
}
