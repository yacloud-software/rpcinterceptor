package resolvers

import (
	"context"
	"fmt"
	apb "golang.conradwood.net/apis/auth"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/cache"
	"golang.conradwood.net/go-easyops/common"
	"golang.conradwood.net/go-easyops/tokens"
	"time"
)

type ResolvedUser interface {
	User() *apb.User
	SignedUser() *apb.SignedUser
}

var (
	idCache    *cache.Cache
	tokenCache *cache.Cache
)

type cacheEntry struct {
	det *apb.SignedUser
	dec *apb.User
}

func (c *cacheEntry) SignedUser() *apb.SignedUser {
	return c.det
}
func (c *cacheEntry) User() *apb.User {
	return c.dec
}
func init() {
	tokenCache = cache.New("tokencache", time.Duration(5*60)*time.Second, 100)
	idCache = cache.New("idcache", time.Duration(5*60)*time.Second, 100)
}

// get a user by id
// this does not distinguish between "error" and "invalid token".
// we treat an error as "invalid token" (as it's most presumed to be most common)
// missing tokens return no error, but nil response
func (r *Resolvers) AsRootGetUserByID(ctx context.Context, id string) (ResolvedUser, error) {
	if id == "" {
		return nil, nil
	}
	axs := idCache.Get(id)
	if axs != nil {
		return axs.(*cacheEntry), nil
	}

	vr := apb.ByIDRequest{UserID: id}
	ctx = tokens.ContextWithToken() // we need to call it with our context. only we have permission to do this
	det, err := authManager.SignedGetUserByID(ctx, &vr)
	if err != nil {
		return nil, err
	}
	dec := common.VerifySignedUser(det)
	if dec == nil {
		return nil, fmt.Errorf("invalid signature")
	}
	if !dec.Active {
		return nil, fmt.Errorf("invalid user")
	}
	ce := &cacheEntry{det: det, dec: dec}
	idCache.Put(id, ce)
	return ce, nil
}

// get a user by token
// this does not distinguish between "error" and "invalid token".
// we treat an error as "invalid token" (as it's most presumed to be most common)
// missing tokens return no error, but nil response
func (r *Resolvers) SignedGetUserByToken(ctx context.Context, token string) (ResolvedUser, error) {
	if token == "" {
		return nil, nil
	}
	axs := tokenCache.Get(token)
	if axs != nil {
		return axs.(*cacheEntry), nil
	}

	vr := apb.AuthenticateTokenRequest{Token: token}
	det, err := authremote.GetAuthClient().SignedGetByToken(ctx, &vr)
	if err != nil {
		return nil, err
	}
	dec := common.VerifySignedUser(det.User)
	if dec == nil {
		return nil, fmt.Errorf("invalid signature")
	}
	if !det.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	if !dec.Active {
		return nil, fmt.Errorf("invalid user")
	}
	ce := &cacheEntry{det: det.User, dec: dec}
	tokenCache.Put(token, ce)
	idCache.Put(ce.dec.ID, ce)
	return ce, nil
}
