package resolvers

import (
	apb "golang.conradwood.net/apis/auth"
	"golang.conradwood.net/go-easyops/authremote"
)

var (
	authServer  apb.AuthenticationServiceClient
	authManager apb.AuthManagerServiceClient
)

type Resolvers struct {
}

func NewResolver() (*Resolvers, error) {
	r := &Resolvers{}
	if authServer == nil {
		authServer = authremote.GetAuthClient()
	}
	if authManager == nil {
		authManager = authremote.GetAuthManagerClient()

	}
	return r, nil
}

func (r *Resolvers) Clear() {
	serviceCache.Clear()
	methodNameCache.Clear()
	methodCache.Clear()
}
