package main

import (
	"flag"
	"fmt"
	"golang.conradwood.net/apis/auth"
	"golang.conradwood.net/apis/common"
	pb "golang.conradwood.net/apis/rpcaclapi"
	"golang.conradwood.net/go-easyops/client"
	"golang.conradwood.net/go-easyops/tokens"
	"golang.conradwood.net/go-easyops/utils"
	"golang.org/x/net/context"
	"gopkg.in/yaml.v2"
	"os"
	"strconv"
	"strings"
)

var (
	use_yaml   = flag.Bool("yaml", false, "output lists in yaml")
	listcalls  = flag.Bool("list_calls", false, "show current list of who's calling whom")
	method     = flag.String("method", "", "Methodname")
	group      = flag.String("group", "", "group `ID`")
	user       = flag.String("user", "", "user ID")
	nuser      = flag.String("notuser", "", "user ID to exclude")
	grant      = flag.Bool("grant", false, "grant a user access to a service (requires user|group service params)")
	revoke     = flag.Bool("revoke", false, "revoke a user's access to a method (requires user, method, service params)")
	service    = flag.String("service", "", "Servicename")
	services   = flag.Bool("services", false, "List Servicename and id")
	getperms   = flag.Bool("permissions", false, "get permissions on method,service")
	getgroups  = flag.Bool("groups", false, "get all groups")
	addgroup   = flag.Bool("addgroup", false, "add a group to a method")
	remgroup   = flag.Bool("removegroup", false, "remove a group from a method")
	showacl    = flag.Bool("showaccess", false, "show access to a method or a group")
	showlogs   = flag.Bool("logs", false, "show logs")
	showerrors = flag.Bool("errors", false, "show errors")
	serviceids = flag.String("serviceids", "", "comma delimited list of service ids (e.g. for logs)")
	methodids  = flag.String("methodids", "", "comma delimited list of method ids (e.g. for logs)")
	rpcapi     pb.RPCACLServiceClient
	authClient auth.AuthManagerServiceClient
)

func main() {
	flag.Parse()

	rpcapi = pb.NewRPCACLServiceClient(client.Connect("rpcaclapi.RPCACLService"))
	authClient = auth.NewAuthManagerServiceClient(client.Connect("auth.AuthManagerService"))

	if *listcalls {
		listCalls()
	}
	if *showlogs {
		showLogs()
	}
	if *showerrors {
		showErrors()
	}
	if *getperms {
		getPermissions()
	}
	if *getgroups {
		GetGroups()
	}
	if *addgroup {
		AddGroup()
	}
	if *remgroup {
		RemGroup()
	}
	if *grant {
		GrantAccess()
	}
	if *revoke {
		RevokeAccess()
	}
	if *services {
		ctx := tokens.ContextWithToken()
		svs, err := rpcapi.GetServices(ctx, &common.Void{})
		utils.Bail("could not get services", err)
		for _, s := range svs.Services {
			fmt.Printf("%04d %s\n", s.ID, s.Name)
		}
	}
	os.Exit(0)
}
func showErrors() {
	ctx := tokens.ContextWithToken()
	sl := &pb.ErrorSearchRequest{
		CallerServiceIDs: stringToIntArray(*serviceids),
		CalledMethodIDs:  stringToIntArray(*methodids),
		ExcludeUserIDs:   userids(*nuser),
		UserIDs:          userids(*user),
	}
	les, err := rpcapi.SearchErrors(ctx, sl)
	utils.Bail("failed to get errors", err)
	fmt.Printf("Showing %d errors\n", len(les.Entries))
	for _, rid := range les.Entries {
		uab := pretty(rid.CallerUser)
		fmt.Printf("%d %7s [%s] %s%s: Code: %d \"%s\" LOG == \"%s\"\n", rid.Timestamp, rid.RequestID, uab, rid.Service.Name, rid.Method.Name, rid.ErrorCode, rid.DisplayMessage, rid.LogMessage)
	}
	fmt.Println()

}

func showLogs() {
	ctx := tokens.ContextWithToken()
	sl := &pb.LogSearchRequest{
		CallerServiceIDs: stringToIntArray(*serviceids),
		CalledMethodIDs:  stringToIntArray(*methodids),
	}
	les, err := rpcapi.SearchLogs(ctx, sl)
	utils.Bail("failed to get logs", err)
	fmt.Printf("Showing %d logentries\n", len(les.Entries))
	ls := NewLogSorter(les)
	for _, rid := range ls.RequestIDs() {
		for _, le := range ls.ByRequestID(*rid) {
			uab := "n/a"
			if le.Response.CallerUser != nil {
				uab = le.Response.CallerUser.Email
			}
			fmt.Printf("%d %7s [%s] %s%s\n", le.Timestamp, le.Response.RequestID, uab, le.Service.Name, le.Method.Name)
		}
		fmt.Println()
	}
}

func AddGroup() {
	m := GetUniqueMethod()
	fmt.Printf("Adding group %#v\n", m)
	addGroupToMethod(tokens.ContextWithToken(), *group, m)
}

func addGroupToMethod(ctx context.Context, g string, m *pb.Method) {
	_, err := rpcapi.AddGroupToMethod(ctx, &pb.MethodIDAndGroupID{GroupID: g, MethodID: m.ID})
	utils.Bail(fmt.Sprintf("failed to add group \"%s\" to method %s", g, m.Name), err)
}

func RemGroup() {
	m := GetUniqueMethod()
	fmt.Printf("Removing group %#v\n", m)
	remGroupFromMethod(tokens.ContextWithToken(), *group, m)
}

func remGroupFromMethod(ctx context.Context, g string, m *pb.Method) {
	_, err := rpcapi.RemoveGroupFromMethod(ctx, &pb.MethodIDAndGroupID{GroupID: g, MethodID: m.ID})
	utils.Bail(fmt.Sprintf("failed to remove group \"%s\" to method %s", g, m.Name), err)
}

func GrantAccess() {
	if *group == "" {
		utils.Bail("missing parameter", fmt.Errorf("-group must be set"))
	}
	if *service == "" {
		utils.Bail("missing parameter", fmt.Errorf("-service must be set"))
	}
	gsr := &pb.GroupServiceRequest{ServiceName: *service, GroupID: *group}
	ctx := tokens.ContextWithToken()
	svc, err := rpcapi.AuthoriseGroupService(ctx, gsr)
	utils.Bail("failed to authorise group", err)
	fmt.Printf("Granted access to service %d(%s) to group %s\n", svc.ID, svc.Name, gsr.GroupID)
	fmt.Printf("Done")
}

func RevokeAccess() {
	fmt.Printf("Revoke access to RPC is currently disabled\n")
	// FEATURE DISABLED. DO NOT BRING BACK WITHOUT CONSULTING THE TEAM AND CONFIRMING HOW RPC ACCESS SHOULD BE GRANTED
	/*s := GetService(*service)
	m := GetMethod(s, *method)

	if *user == "" {
		fmt.Printf("User required\n")
		os.Exit(10)
	}

	ctx := tokens.ContextWithToken()

	userGroup := getAuthUserGroup(ctx, s, m, false)

	// remove user from group
	_, err := authClient.RemoveUserFromGroup(ctx, &auth.RemoveFromGroupRequest{UserID: *user, GroupID: userGroup.ID})
	if err != nil {
		utils.Bail("Error removing user from auth group", err)
	}

	fmt.Printf("Revoked user %s access to method\n", *user)*/
}

// return a consistent name for auth user groups created for an rpc interceptor method
func getAuthUserGroupName(s *pb.Service, m *pb.Method) string {
	return fmt.Sprintf("%s.%s", s.Name, m.Name)
}

// find matching user group and create one if requested
func getAuthUserGroup(ctx context.Context, s *pb.Service, m *pb.Method, create bool) *auth.Group {
	groupName := getAuthUserGroupName(s, m)
	groups, err := authClient.ListGroups(ctx, &common.Void{})
	utils.Bail("Error retrieving existing groups", err)

	var userGroup *auth.Group
	for _, g := range groups.Groups {
		if g.Name == groupName {
			userGroup = g
			break
		}
	}

	if userGroup == nil && create {
		panic("not working")
		/*
				res, err := authClient.CreateGroup(ctx, &auth.CreateGroupRequest{CreateGroup: &auth.CreateGroup{
					Name:      groupName,
					ForeignID: fmt.Sprintf("%d", m.ID),
					Origin:    auth.GroupOrigin_RPC_ACL_SERVICE,
				}})

				if err != nil {
					utils.Bail("Error creating an auth group", err)
				}
			userGroup = res.Group
		*/
	}

	return userGroup
}

func GetGroups() {
	ctx := tokens.ContextWithToken()
	gl, err := rpcapi.GetUserGroups(ctx, &common.Void{})
	utils.Bail("could not get groups", err)
	for _, g := range gl.Groups {
		PrintGroup("", g)
	}
}

func getPermissions() {
	if *group != "" {
		getPermissionsForGroup()
	}

	// permissions for method...
	meth := GetUniqueMethod()
	gl, err := rpcapi.GetGroupsForMethod(tokens.ContextWithToken(), meth)
	utils.Bail("failed to get groups for method", err)
	fmt.Printf("Method %s has %d groupacls:\n", meth.Name, len(gl.Groups))
	for _, g := range gl.Groups {
		PrintGroup("  ", g)
	}
}

func getPermissionsForGroup() {
	fmt.Printf("Permissions for group %s:\n", *group)
	ctx := tokens.ContextWithToken()
	gl, err := rpcapi.GetUserGroups(ctx, &common.Void{})
	utils.Bail("could not get groups", err)
	var ag *auth.Group
	for _, g := range gl.Groups {
		if g.ID == *group {
			ag = g
		}
	}
	if ag == nil {
		fmt.Printf("No such group \"%s\"\n", *group)
		os.Exit(10)
	}
	ml, err := rpcapi.GetMethodsForGroups(ctx, ag)
	utils.Bail("Could not get methods for groups", err)
	fmt.Printf("Group %s has access to %d methods.\n", ag.ID, len(ml.Methods))
	for _, m := range ml.Methods {
		fmt.Printf("   %v\n", m)
	}

}

/*************************************************************************
* pretty printers
*************************************************************************/
func PrintGroup(prefix string, g *auth.Group) {
	ctx := tokens.ContextWithToken()
	ul, err := rpcapi.ListUsersInGroup(ctx, &auth.ListGroupRequest{GroupID: g.ID})
	utils.Bail("Failed to get users for group", err)
	fmt.Printf("%sGroup %s %s (%d users)\n", prefix, g.ID, g.Name, len(ul.Users))
	for _, u := range ul.Users {
		fmt.Printf("%s    User: %#v\n", prefix, u)
	}
}

/*************************************************************************
* resolver method/service fuzzy matchers
*************************************************************************/
// get method as specified by command line parameter
func GetUniqueMethod() *pb.Method {
	svc := GetService(*service)
	meth := GetMethod(svc, *method)

	return meth
}

func GetMethod(s *pb.Service, name string) *pb.Method {
	methods := GetMethods(s, name)
	if len(methods) != 1 {
		fmt.Printf("Error need one method matching \"%s\" in service %v, but got %d\n", name, s, len(methods))
		for _, m := range methods {
			fmt.Printf("%s\n", m.Name)
		}
		os.Exit(10)
	}

	fmt.Printf("MethodID: %d Name: %s (%s)\n", methods[0].ID, methods[0].Name, *method)
	return methods[0]
}

func GetMethods(s *pb.Service, name string) []*pb.Method {
	ml, err := rpcapi.GetMethods(tokens.ContextWithToken(),
		&pb.GetMethodsRequest{ServiceID: s.ID})
	utils.Bail("failed to get methods", err)
	var res []*pb.Method
	ln := strings.ToLower(name)
	for _, m := range ml.Methods {
		if strings.Contains(strings.ToLower(m.Name), ln) {
			res = append(res, m)
		}
	}
	return res
}

func GetService(name string) *pb.Service {
	svcs := GetServices(name)
	if len(svcs) != 1 {
		fmt.Printf("Error: Need one service matching \"%s\", but got %d:\n", name, len(svcs))
		for _, s := range svcs {
			fmt.Printf("ID: %d\tName: %s\n", s.ID, s.Name)
		}
		os.Exit(10)
	}

	fmt.Printf("ServiceID: %d Name: %s (%s)\n", svcs[0].ID, svcs[0].Name, *service)
	return svcs[0]
}

func GetServices(name string) []*pb.Service {
	if name == "" {
		fmt.Printf("Servicename is required\n")
		os.Exit(10)
	}
	ln := strings.ToLower(name)
	s, err := rpcapi.GetServices(tokens.ContextWithToken(), &common.Void{})
	utils.Bail("Failed to get services", err)
	var res []*pb.Service
	for _, sn := range s.Services {
		//	fmt.Printf("Service: %s\n", sn.Name)
		if strings.Contains(strings.ToLower(sn.Name), ln) {
			res = append(res, sn)
		}
	}
	return res
}

func stringToIntArray(s string) []uint64 {
	var res []uint64
	for _, sx := range strings.Split(s, ",") {
		if sx == "" {
			continue
		}
		l, err := strconv.ParseInt(sx, 10, 64)
		utils.Bail("not a number", err)
		res = append(res, uint64(l))
	}
	return res
}

func listCalls() {
	ctx := tokens.ContextWithToken()
	calls, err := rpcapi.ListCalls(ctx, &common.Void{})
	utils.Bail("failed to list calls", err)
	if !*use_yaml {
		for _, call := range calls.Calls {
			fmt.Printf("Service %40s | user %20s called %s.%s\n", pretty(call.CallingService), pretty(call.CallingUser), call.CalledMethod.Service.Name, call.CalledMethod.Name)
		}
		return
	}
	// build a yaml format
	ac, err := CreateAccessConfig(calls)
	utils.Bail("failed to build accessconfig", err)
	sum := ac.Summary()
	b, err := yaml.Marshal(sum)
	utils.Bail("Failed to marshal", err)
	fmt.Printf("Config:\n%s\n", string(b))
}
func pretty(user *auth.User) string {
	if user == nil {
		return "[NONE]"
	}
	if user.Abbrev != "" {
		return fmt.Sprintf("%s(%s)", user.Abbrev, user.ID)
	}
	if user.Email != "" {
		return fmt.Sprintf("%s(%s)", user.Email, user.ID)
	}
	return user.ID
}

func userids(para string) []string {
	if para == "" {
		return []string{}
	}
	return []string{para}
}
