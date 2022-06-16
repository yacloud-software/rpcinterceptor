package main

import (
	"fmt"
	"golang.conradwood.net/apis/auth"
	"golang.conradwood.net/apis/common"
	ic "golang.conradwood.net/apis/rpcaclapi"
	af "golang.conradwood.net/go-easyops/auth"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/sql"
	"golang.org/x/net/context"
	"time"
)

// func (e *rpcACLServer) InterceptRPC(ctx context.Context, req *ic.InterceptRPCRequest) (*ic.InterceptRPCResponse, error) {
func (e *rpcACLServer) GetServiceByUserID(ctx context.Context, req *ic.ServiceByUserIDRequest) (*ic.Service, error) {
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
		return nil, errors.NotFound(ctx, "not found")
	}
	res := &ic.Service{}
	err = rows.Scan(&res.ID, &res.Name, &res.UserID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// func (e *rpcACLServer) InterceptRPC(ctx context.Context, req *ic.InterceptRPCRequest) (*ic.InterceptRPCResponse, error) {

func (e *rpcACLServer) GetServiceByID(ctx context.Context, req *ic.ServiceByIDRequest) (*ic.Service, error) {
	db, err := sql.Open()
	if err != nil {
		fmt.Printf("Failed to connect to db: %s\n", err)
		return nil, err
	}
	rows, err := db.QueryContext(ctx, "list_servicenames", "select id,servicename,userid from service where id = $1", req.ID)
	if err != nil {
		return nil, fmt.Errorf("Unable to select services from db: %s", err)
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, errors.NotFound(ctx, "not found")
	}
	res := &ic.Service{}
	err = rows.Scan(&res.ID, &res.Name, &res.UserID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e *rpcACLServer) GetServices(ctx context.Context, req *common.Void) (*ic.ServiceList, error) {
	db, err := sql.Open()
	if err != nil {
		fmt.Printf("Failed to connect to db: %s\n", err)
		return nil, err
	}
	rows, err := db.QueryContext(ctx, "list_servicenames", "select id,servicename,userid from service order by servicename asc")
	if err != nil {
		return nil, fmt.Errorf("Unable to select services from db: %s", err)
	}
	defer rows.Close()
	res := &ic.ServiceList{}
	for rows.Next() {
		sc := &ic.Service{}
		err = rows.Scan(&sc.ID, &sc.Name, &sc.UserID)
		if err != nil {
			return nil, fmt.Errorf("Unable to scan services from db: %s", err)
		}
		res.Services = append(res.Services, sc)
	}
	return res, nil

}
func (e *rpcACLServer) GetMethods(ctx context.Context, req *ic.GetMethodsRequest) (*ic.MethodList, error) {
	db, err := sql.Open()
	if err != nil {
		fmt.Printf("Failed to connect to db: %s\n", err)
		return nil, err
	}
	rows, err := db.QueryContext(ctx, "list_methods", "select id,methodname,service_id from method where service_id = $1 order by methodname asc", req.ServiceID)
	if err != nil {
		return nil, fmt.Errorf("Unable to select methods from db: %s", err)
	}
	defer rows.Close()
	res := &ic.MethodList{}
	for rows.Next() {
		sc := &ic.Method{}
		err = rows.Scan(&sc.ID, &sc.Name, &sc.ServiceID)
		if err != nil {
			return nil, fmt.Errorf("Unable to scan methods from db: %s", err)
		}
		res.Methods = append(res.Methods, sc)
	}
	return res, nil
}
func (e *rpcACLServer) GetUserGroups(ctx context.Context, req *common.Void) (*auth.GroupList, error) {
	return authManager.ListGroups(ctx, req)
}
func (e *rpcACLServer) AddUserToGroup(ctx context.Context, req *auth.AddToGroupRequest) (*common.Void, error) {
	return &common.Void{}, nil
}
func (e *rpcACLServer) RemoveUserFromGroup(ctx context.Context, req *auth.RemoveFromGroupRequest) (*common.Void, error) {
	return &common.Void{}, nil
}
func (e *rpcACLServer) ListUsersInGroup(ctx context.Context, req *auth.ListGroupRequest) (*auth.UserListResponse, error) {
	//return authManager.ListUsersInGroup(ctx, req)
	return nil, fmt.Errorf("not implemented")
}
func (e *rpcACLServer) GetGroupsForMethod(ctx context.Context, req *ic.Method) (*auth.GroupList, error) {
	db, err := sql.Open()
	if err != nil {
		fmt.Printf("Failed to connect to db: %s\n", err)
		return nil, err
	}
	rows, err := db.QueryContext(ctx, "list_groups_for_method", "select group_id from methodacl where method_id = $1", req.ID)
	if err != nil {
		return nil, fmt.Errorf("Unable to select methodacls from db: %s", err)
	}
	defer rows.Close()
	res := &auth.GroupList{}
	for rows.Next() {
		var gid string
		err = rows.Scan(&gid)
		if err != nil {
			return nil, fmt.Errorf("Unable to scan methodacls from db: %s", err)
		}
		ag := rev.GetGroupByID(ctx, gid, true)
		if ag != nil {
			res.Groups = append(res.Groups, ag)
		}
	}
	return res, nil
}
func (e *rpcACLServer) AddGroupToMethod(ctx context.Context, req *ic.MethodIDAndGroupID) (*auth.GroupList, error) {
	fmt.Printf("Request to add group %s to method #%d\n", req.GroupID, req.MethodID)
	if req.GroupID == "" {
		return nil, fmt.Errorf("GroupID required")
	}
	if req.MethodID == 0 {
		return nil, fmt.Errorf("MethodID required")
	}
	if !IsACLManager(ctx) {
		return nil, fmt.Errorf("This method is restricted to ACLManagers only")
	}

	m := rev.GetMethodByID(ctx, req.MethodID)
	if m == nil {
		return nil, fmt.Errorf("No method %d", req.MethodID)
	}
	res, err := e.GetGroupsForMethod(ctx, m) // previous state
	db, err := sql.Open()
	if err != nil {
		fmt.Printf("Failed to connect to db: %s\n", err)
		return nil, err
	}
	whoami := af.GetUser(ctx)
	_, err = db.QueryContext(ctx, "insert_into_methodacl", "insert into methodacl (created_at,created_by,method_id,group_id) values ($1,$2,$3,$4)",
		time.Now().Unix(),
		whoami.ID,
		req.MethodID,
		req.GroupID,
	)
	if err == nil {
		Audit_GroupAddedToMethod(ctx, req.MethodID, req.GroupID)
	}
	// if the entry is active already, no error. (we did what the user asked us to, so no worry)
	if err != nil && db.CheckDuplicateRowError(err) {
		err = nil
	}
	return res, err
}
func (e *rpcACLServer) RemoveGroupFromMethod(ctx context.Context, req *ic.MethodIDAndGroupID) (*auth.GroupList, error) {
	fmt.Printf("Request to remove group %s to method #%d\n", req.GroupID, req.MethodID)
	if req.GroupID == "" {
		return nil, fmt.Errorf("GroupID required")
	}
	if req.MethodID == 0 {
		return nil, fmt.Errorf("MethodID required")
	}
	if !IsACLManager(ctx) {
		return nil, fmt.Errorf("This method is restricted to ACLManagers only")
	}
	m := rev.GetMethodByID(ctx, req.MethodID)
	if m == nil {
		return nil, fmt.Errorf("No method %d", req.MethodID)
	}
	res, err := e.GetGroupsForMethod(ctx, m) // previous state
	// is there a link group<->method atm?
	isalloc := false
	for _, k := range res.Groups {
		if k.ID == req.GroupID {
			isalloc = true
			break
		}
	}
	if !isalloc {
		// shortcut: no allocation atm, so no change
		return res, nil
	}
	db, err := sql.Open()
	if err != nil {
		fmt.Printf("Failed to connect to db: %s\n", err)
		return nil, err
	}
	_, err = db.QueryContext(ctx, "delete_medhodacl", "delete from methodacl where method_id = $1 and group_id=$2",
		req.MethodID,
		req.GroupID,
	)
	if err == nil {
		Audit_GroupRemovedFromMethod(ctx, req.MethodID, req.GroupID)
	}
	return res, err
}
func (e *rpcACLServer) GetMethodsForGroups(ctx context.Context, req *auth.Group) (*ic.FullMethodList, error) {
	db, err := sql.Open()
	if err != nil {
		fmt.Printf("Failed to connect to db: %s\n", err)
		return nil, err
	}
	rows, err := db.QueryContext(ctx, "select_methods_for_group", "select method_id from methodacl where group_id = $1", req.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	res := &ic.FullMethodList{}
	for rows.Next() {
		var mid uint64
		err = rows.Scan(&mid)
		if err != nil {
			return nil, fmt.Errorf("Failed to scan db for methodacls: %s", err)
		}
		m := rev.GetMethodByID(ctx, mid)
		if m == nil {
			// should not happen (foreign key in db). Cache bug?
			return nil, fmt.Errorf("Inconsistency: link in methodacl to method %d, which does not exist", mid)
		}
		s := rev.GetServiceByID(ctx, m.ServiceID)
		if s == nil {
			// should not happen (foreign key in db). Cache bug?
			return nil, fmt.Errorf("Inconsistency: Method %d refers to service %d, which does not exist", m.ID, m.ServiceID)
		}
		fm := &ic.FullMethod{ID: m.ID, Name: m.Name, Service: s}
		res.Methods = append(res.Methods, fm)

	}
	return res, nil
}

func IsACLManager(ctx context.Context) bool {
	whoami := af.GetUser(ctx)
	if whoami == nil {
		fmt.Printf("This method is restricted to actual users only\n")
		return false
	}
	if !af.IsInGroup(ctx, "fooadmingroup") {
		fmt.Printf("This method is restricted to guru users only\n")
		return false
	}
	return true
}
