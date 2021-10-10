package main

import (
	"flag"
	"fmt"
	"golang.conradwood.net/apis/common"
	ic "golang.conradwood.net/apis/rpcinterceptor"
	"golang.conradwood.net/go-easyops/client"
	"golang.conradwood.net/go-easyops/tokens"
	"golang.conradwood.net/go-easyops/utils"
	"golang.conradwood.net/learn"
	"os"
)

var (
	rs    = flag.Bool("reset", false, "reset learnings")
	rpcic ic.RPCInterceptorServiceClient
)

func main() {
	flag.Parse()
	rpcic = ic.NewRPCInterceptorServiceClient(client.Connect("rpcinterceptor.RPCInterceptorService"))
	if *rs {
		_, err := rpcic.ClearLearnings(tokens.ContextWithToken(), &common.Void{})
		utils.Bail("Failed to clear learnings", err)
	} else {
		learnings()
	}
}

func learnings() {
	ctx := tokens.ContextWithToken()
	ls, err := rpcic.GetLearnings(ctx, &common.Void{})
	utils.Bail("failed to get learnings", err)
	fmt.Printf("Got %d learnings\n", len(ls.Learnings))
	a := learn.FromProto(ctx, ls)
	if a == nil {
		fmt.Printf("Failed to parse learnings\n")
		os.Exit(10)
	}
	for _, s := range a.Services() {
		fmt.Printf("Service: %s\n", s)
		for _, c := range s.Callers() {
			fmt.Printf("   Called by: %s\n", c)
		}
	}
}
