package learn

import (
	"context"
	"flag"
	"fmt"
	au "golang.conradwood.net/apis/auth"
	aa "golang.conradwood.net/apis/rpcaclapi"
	ic "golang.conradwood.net/apis/rpcinterceptor"
	"golang.conradwood.net/go-easyops/auth"
	"golang.conradwood.net/resolvers"
)

var (
	learnCache = make(map[uint64][]*Category)
	debug      = flag.Bool("debug_learn", false, "debug learning code")
	learnflag  = flag.Bool("learn", false, "learn mode means all access is tracked in-ram by type, e.g. service a accesses service b etc. This can be queried via RPC at any time")
)

type Category struct {
	typ    int
	userid string
	from   uint64 // service id
	to     uint64 // service id
	count  uint64
}
type learnreq struct {
	prefix        string
	resolver      *resolvers.Resolvers
	resp          *ic.InterceptRPCResponse
	req           *ic.InterceptRPCRequest
	ctx           context.Context
	CallerService *aa.Service // service being intercepted
	CallerMethod  *aa.Method  // method being intercepted
	CalleeService *aa.Service // service invoking "Caller"
	User          *au.User
}

// no context here - we run totally asynchronously and independently from the serving path
func Learn(r *resolvers.Resolvers, req *ic.InterceptRPCRequest, resp *ic.InterceptRPCResponse) {
	if !*learnflag {
		return
	}
	lo := learnreq{
		ctx:      context.Background(),
		resolver: r,
		req:      req,
		resp:     resp,
		prefix:   fmt.Sprintf("[%06s]", resp.RequestID),
	}
	//caller
	mid := lo.resp.CallerMethodID
	md := lo.resolver.GetMethodByID(lo.ctx, mid)
	svid := md.ServiceID
	sv := lo.resolver.GetServiceByID(lo.ctx, svid)
	lo.CallerService = sv
	lo.CallerMethod = md
	// callee
	svid = lo.resp.CalleeServiceID
	lo.CalleeService = lo.resolver.GetServiceByID(lo.ctx, svid)

	// user
	lo.User = lo.resp.CallerUser
	// process what we got..
	lo.process()
}
func (lo *learnreq) Printf(msg string, args ...interface{}) {
	m := fmt.Sprintf("%s %s", lo.prefix, msg)
	fmt.Printf(m, args...)
}

// build a unique key for the origin (the caller, the source) of this request
func (lo *learnreq) CallerID() string {
	return fmt.Sprintf("[#%003d %s->#%003d %s()]",
		lo.CallerService.ID,
		lo.CallerService.Name,
		lo.CallerMethod.ID,
		lo.CallerMethod.Name)
}

// build a unique key for the origin (the caller, the source) of this request
func (lo *learnreq) CalleeID() string {
	// we might not HAVE a service that called us (e.g. h2gproxy)
	if lo.CalleeService == nil {
		return "[NOSERVICE]"
	}
	return fmt.Sprintf("[#%003d %s]", lo.CalleeService.ID, lo.CalleeService.Name)
}
func (lo *learnreq) UserID() string {
	if lo.User == nil {
		return "[NOUSER]"
	}
	return fmt.Sprintf("%s (%s)", lo.User.ID, auth.Description(lo.User))

}
func Reset() {
	learnCache = make(map[uint64][]*Category)
}

func Learnings() []*ic.Learning {
	var res []*ic.Learning
	for _, v := range learnCache {
		for _, c := range v {
			l := &ic.Learning{
				FromServiceID: c.from,
				ToServiceID:   c.to,
				UserID:        c.userid,
				Count:         c.count,
			}
			res = append(res, l)
		}
	}
	return res
}

func (lo *learnreq) process() {
	// we break it down into these categories:
	// 1. service calls service with user
	// 2. [noservice] calls service with user
	// 3. service calls service with no user

	var category *Category
	if lo.CalleeService != nil && lo.CallerService != nil && lo.User != nil {
		category = &Category{typ: 1, from: lo.CalleeService.ID, to: lo.CallerService.ID, userid: lo.User.ID}
	} else if lo.CalleeService == nil && lo.CallerService != nil && lo.User != nil {
		category = &Category{typ: 2, to: lo.CallerService.ID, userid: lo.User.ID}
	} else if lo.CalleeService != nil && lo.CallerService != nil && lo.User == nil {
		category = &Category{typ: 3, from: lo.CalleeService.ID, to: lo.CallerService.ID}
	} else {
		lo.Printf("??? : %s ---> %s /// as user %s\n", lo.CalleeID(), lo.CallerID(), lo.UserID())
		return
	}

	if *debug {
		if lo.resp.Reject {
			lo.Printf("REJ %d: %s ---> %s /// as user %s\n", category.typ, lo.CalleeID(), lo.CallerID(), lo.UserID())
		} else {
			lo.Printf("ACC %d: %s ---> %s /// as user %s\n", category.typ, lo.CalleeID(), lo.CallerID(), lo.UserID())
		}
	}
	sc := learnCache[lo.CallerService.ID]
	var mc *Category
	for _, s := range sc {
		if s.equals(category) {
			mc = s
			break
		}
	}
	if mc == nil {
		mc = category
		sc = append(sc, mc)
		learnCache[lo.CallerService.ID] = sc
	}
	mc.count++
}
func (c *Category) equals(c2 *Category) bool {
	if c.typ != c2.typ {
		return false
	}
	if c.from != c2.from {
		return false
	}
	if c.to != c2.to {
		return false
	}
	if c.typ == 1 {
		if c.userid != c2.userid {
			return false
		}
	}
	return true
}
