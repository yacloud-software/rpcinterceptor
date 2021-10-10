package main

import (
	"fmt"
	"golang.conradwood.net/go-easyops/auth"
	"golang.org/x/net/context"
)

//eventually this will be an auditlog rather than just a printed message...
func Audit_GroupAddedToMethod(ctx context.Context, methodid uint64, groupid string) {
	whoami := auth.GetUser(ctx)
	s := "ANONYMOUS"
	if whoami != nil {
		s = fmt.Sprintf("%s/%s", whoami.ID, whoami.Email)
	}
	fmt.Printf("AUDIT: User %s added permissions to method %d to group %s\n", s, methodid, groupid)
}

//eventually this will be an auditlog rather than just a printed message...
func Audit_GroupRemovedFromMethod(ctx context.Context, methodid uint64, groupid string) {
	whoami := auth.GetUser(ctx)
	s := "ANONYMOUS"
	if whoami != nil {
		s = fmt.Sprintf("%s/%s", whoami.ID, whoami.Email)
	}
	fmt.Printf("AUDIT: User %s removed permissions from method %d for group %s\n", s, methodid, groupid)
}
