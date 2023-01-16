package main

import (
	"context"
	ic "golang.conradwood.net/apis/rpcinterceptor"
	//	"golang.conradwood.net/go-easyops/auth"
	"fmt"
	"golang.conradwood.net/go-easyops/common"
	"golang.conradwood.net/go-easyops/ctx"
)

func deprecate_rpc_interceptor(ictx context.Context, req *ic.InterceptRPCRequest) (*ic.InterceptRPCResponse, error) {
	ls := ctx.GetLocalState(ictx)
	if *debug {
		fmt.Printf("Request: %#v\n", req)
		fmt.Printf("Request.Meta: %#v\n", req.InMetadata)
		fmt.Printf("Context: %s\n", ictx)
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
	return res, nil
}
