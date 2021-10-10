package main

import (
	"flag"
	"fmt"
	"golang.conradwood.net/apis/common"
	ic "golang.conradwood.net/apis/rpcinterceptor"
	"golang.conradwood.net/go-easyops/sql"
	"golang.org/x/net/context"
	"time"
)

const (
	ERRCOLS = ("occured,ireq,method_id,callerservice,calleruserid,errcode,displaymessage,logmessage")
)

var (
	clean_err_old = flag.Int("clean_errors", 1800, "seconds before cleaning error log entries")
)

func (e *rpcInterceptorServer) LogError(ctx context.Context, req *ic.LogErrorRequest) (*common.Void, error) {
	cv := &common.Void{}
	if req.InMetadata == nil {
		return nil, fmt.Errorf("rpcinterceptor does not log errors without InMetadata")
	}
	im := req.InMetadata
	if im.UserID == "" {
		return nil, fmt.Errorf("rpcinterceptor does not log errors without userid")
	}
	db, err := sql.Open()
	if err != nil {
		fmt.Printf("errorlog1: Failed to connect to db: %s\n", err)
		return nil, err
	}
	_, err = db.ExecContext(ctx, "inserterror", "insert into errentry ("+ERRCOLS+") values ($1,$2,$3,$4,$5,$6,$7,$8)",
		time.Now().Unix(),
		im.RequestID,
		im.CallerMethodID,
		req.Service,
		im.UserID,
		req.ErrorCode,
		req.DisplayMessage,
		req.LogMessage,
	)
	if err != nil {
		fmt.Printf("could not insert errorentry row: %s\n", err)
		return nil, err
	}
	return cv, nil
}

func clean_error_db() {
	now := time.Now().Unix()
	now = now - 60*int64(*clean_err_old)
	db, err := sql.Open()
	if err != nil {
		fmt.Printf("errorlog2: Failed to connect to db: %s\n", err)
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	_, err = db.ExecContext(ctx, "deletelogentries", "delete from errentry where occured < $1", now)

	// 	_, err = db.Exec("delete from logentry where occured < $1", now)
	if err != nil {
		fmt.Printf("Cleaner failed: %s\n", err)
	}
}
