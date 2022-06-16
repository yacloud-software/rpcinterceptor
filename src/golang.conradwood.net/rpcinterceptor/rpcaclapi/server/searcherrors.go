package main

import (
	"fmt"
	ic "golang.conradwood.net/apis/rpcaclapi"
	"golang.conradwood.net/go-easyops/sql"
	"golang.org/x/net/context"
)

/*
the proto is somewhat complex, here it is as a refresher
message ErrorSearchRequest {
  // only include entries which these user(s) made (empty list=all)
  repeated string UserIDs = 1;
  // only include entries which users in these groups made (empty list=all)
  repeated string GroupIDs = 2;
  // automatically include entries which do not match the SearchRequest but have the same requestid as one that matched
  bool FullRequests = 3;
  // unix epoch starttime (only include entries younger than this). 0 = all
  uint32 StartTime = 4;
  // unix epoch endtime (only include entries older than this). 0 = all
  uint32 EndTime = 5;
  // only include entries where the caller service is in this list (empty list=all)
  repeated uint64 CallerServiceIDs = 6;
  // only include entries where the called service is in this list (empty list=all)
  repeated uint64 CalledServiceIDs = 7;
  // only include entries where the caller method is in this list (empty list=all)
  repeated uint64 CallerMethodIDs = 8;
  // only include entries where the called method is in this list (empty list=all)
  repeated uint64 CalledMethodIDs = 9;
}
*/
func (e *rpcACLServer) SearchErrors(ctx context.Context, req *ic.ErrorSearchRequest) (*ic.ErrorEntryList, error) {
	// verify strings...
	for _, u := range req.UserIDs {
		if !isValidUserID(u) {
			if *debug {
				fmt.Printf("\"%s\" is not a valid userid\n", u)
			}
			return nil, fmt.Errorf("\"%s\" is not a valid userid", u)
		}
	}
	for _, u := range req.GroupIDs {
		if !isValidGroupID(u) {
			if *debug {
				fmt.Printf("\"%s\" is not a valid userid\n", u)
			}
			return nil, fmt.Errorf("\"%s\" is not a valid userid", u)
		}
	}
	if *debug {
		fmt.Printf("search: %d RequestIDs\n", len(req.RequestIDs))
		fmt.Printf("search: %d UserIDs\n", len(req.UserIDs))
		fmt.Printf("search: %d ExcludeUserIDs\n", len(req.ExcludeUserIDs))
		fmt.Printf("search: %d GroupIDs\n", len(req.GroupIDs))
		fmt.Printf("search: %d CallerServiceIDs\n", len(req.CallerServiceIDs))
		fmt.Printf("search: %d CalledServiceIDs\n", len(req.CalledServiceIDs))
	}

	db, err := sql.Open()
	if err != nil {
		fmt.Printf("Failed to connect to db: %s\n", err)
		return nil, err
	}

	// translate serviceids into methodids:
	mids := req.CalledMethodIDs
	for _, sid := range req.CallerServiceIDs {
		smids := e.serviceToMethodIDs(sid)
		mids = append(mids, smids...)
	}
	for _, sid := range req.CalledServiceIDs {
		smids := e.serviceToMethodIDs(sid)
		mids = append(mids, smids...)
	}
	var eclauses []string
	eclauses = append(eclauses, addStringWhere("calleruserid", req.UserIDs))
	eclauses = append(eclauses, addStringNotWhere("calleruserid", req.ExcludeUserIDs))
	eclauses = append(eclauses, addStringWhere("ireq", req.RequestIDs))
	eclauses = append(eclauses, addIntWhere("method_id", mids))
	var clauses []string
	for _, c := range eclauses {
		if c != "" {
			clauses = append(clauses, c)
		}
	}
	where := ""
	if len(clauses) != 0 {
		where = "where "
	}
	delim := ""
	for _, c := range clauses {
		if c == "" {
			continue
		}
		where = where + delim + c
		delim = " AND "
	}
	fmt.Printf("WHERE: <%s>\n", where)
	res := &ic.ErrorEntryList{}
	rows, err := db.QueryContext(ctx, "list_errors", "select occured,ireq,method_id,callerservice,calleruserid,errcode,displaymessage,logmessage from errentry  "+where+" order by id desc limit $1", 200)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var mid uint64
		var callerservice string
		var calleruserid string
		le := &ic.ErrorEntry{}
		err = rows.Scan(&le.Timestamp, &le.RequestID, &mid, &callerservice, &calleruserid, &le.ErrorCode, &le.DisplayMessage, &le.LogMessage)
		if err != nil {
			return nil, fmt.Errorf("Failed to scan errorentries: %s", err)
		}
		m := rev.GetMethodByID(ctx, mid)
		if m == nil {
			fmt.Printf("Skippig errorentry - No such method: %d\n", mid)
			continue
		}
		le.Method = m
		le.Service = rev.GetServiceByID(ctx, m.ServiceID)
		if calleruserid != "" {
			le.CallerUser, _ = rev.GetUserByID(ctx, calleruserid)
		}

		res.Entries = append(res.Entries, le)
	}
	return res, nil
}
