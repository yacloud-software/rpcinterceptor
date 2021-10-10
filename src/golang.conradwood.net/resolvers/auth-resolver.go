package resolvers

import (
	"fmt"
	"golang.conradwood.net/apis/auth"
	"golang.conradwood.net/apis/common"
	"golang.conradwood.net/go-easyops/cache"
	"golang.conradwood.net/go-easyops/utils"
	"golang.org/x/net/context"
	"runtime/debug"
	"time"
)

// cache for auth server

var (
	groupCache = cache.New("groupcache", time.Duration(5)*time.Minute, 100)
	userCache  = cache.New("usercache", time.Duration(5)*time.Minute, 100)
)

type groupCacheEntry struct {
	ag *auth.Group
}
type userCacheEntry struct {
	ag *auth.User
}

func (r *Resolvers) GetGroupByID(ctx context.Context, id string, nonLdap bool) *auth.Group {
	acs := groupCache.Get(id)
	if acs != nil {
		return acs.(*groupCacheEntry).ag
	}
	gl, err := authManager.ListGroups(ctx, &common.Void{})
	if err != nil {
		fmt.Printf("Failed to get groups from auth server: %s\n", err)
		return nil
	}
	var res *groupCacheEntry
	for _, g := range gl.Groups {
		gce := &groupCacheEntry{ag: g}
		groupCache.Put(g.ID, gce)
		if g.ID == id {
			res = gce
		}
	}

	if res == nil {
		// negative cache
		groupCache.Put(id, &groupCacheEntry{})
		return nil
	}
	return res.ag

}

func (r *Resolvers) GetUserByID(ctx context.Context, id string) (*auth.User, error) {
	acs := userCache.Get(id)
	if acs != nil {
		return acs.(*userCacheEntry).ag, nil
	}
	gd, err := authManager.GetUserByID(ctx, &auth.ByIDRequest{UserID: id})
	if err != nil {
		userCache.Put(id, &userCacheEntry{}) // negative cache
		fmt.Printf("Failed to get user %s from auth server: %s\n", id, utils.ErrorString(err))
		debug.PrintStack()
		return nil, err
	}
	gce := &userCacheEntry{ag: gd}
	userCache.Put(gd.ID, gce)
	return gce.ag, nil

}
