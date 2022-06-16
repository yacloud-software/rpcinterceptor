package resolvers

import (
	"context"
	"fmt"
	ic "golang.conradwood.net/apis/rpcaclapi"
	"golang.conradwood.net/go-easyops/auth"
	"golang.conradwood.net/go-easyops/cache"
	"golang.conradwood.net/go-easyops/sql"
	"time"
)

var (
	serviceCache      = cache.New("servicecache", time.Duration(5)*time.Minute, 100)      // by service name
	serviceTokenCache = cache.New("servicetokencache", time.Duration(5)*time.Minute, 100) // by service name
)

type serviceCacheEntry struct {
	m *ic.Service
}

func (r *Resolvers) GetServiceByToken(ctx context.Context, token string) *ic.Service {
	o := serviceTokenCache.Get(token)
	if o != nil {
		return o.(*serviceCacheEntry).m
	}

	user, err := r.SignedGetUserByToken(ctx, token)
	if err != nil {
		fmt.Printf("failed to resolve user by token: %s\n", err)
		return nil
	}
	if user == nil {
		return nil
	}
	db, err := sql.Open()
	if err != nil {
		fmt.Printf("scache1: Failed to connect to db: %s\n", err)
		return nil
	}

	rows, err := db.QueryContext(ctx, "service_by_userid", "select id from service where userid = $1", user.User().ID)
	if err != nil {
		fmt.Printf("Unable to select services by userid from db: %s\n", err)
		return nil
	}
	if !rows.Next() {
		fmt.Printf("Resolved token to User %s (#%s) but no service for it\n", auth.Description(user.User()), user.User().ID)
		rows.Close()
		return nil
	}
	var id uint64
	err = rows.Scan(&id)
	rows.Close()
	icm := r.GetServiceByID(ctx, id)
	//fmt.Printf("Service: %s\n", icm)
	serviceTokenCache.Put(token, &serviceCacheEntry{m: icm})
	return icm
}

// not cached yet...
func (r *Resolvers) GetServiceByName(ctx context.Context, servicename string) *ic.Service {
	key := servicename
	acs := serviceCache.Get(key)
	if acs != nil {
		return acs.(*serviceCacheEntry).m
	}
	db, err := sql.Open()
	if err != nil {
		fmt.Printf("scache1: Failed to connect to db: %s\n", err)
		return nil
	}
	rows, err := db.QueryContext(ctx, "service_by_name", "select id,servicename,userid from service where servicename = $1", servicename)
	if err != nil {
		fmt.Printf("Unable to select services from db: %s\n", err)
		return nil
	}
	defer rows.Close()
	if !rows.Next() {
		serviceCache.Put(key, &serviceCacheEntry{})
		return nil
	}
	icm := ic.Service{}
	err = rows.Scan(&icm.ID, &icm.Name, &icm.UserID)
	if err != nil {
		fmt.Printf("failed to scan db for service: %s\n", err)
		return nil
	}
	gce := &serviceCacheEntry{m: &icm}
	serviceCache.Put(key, gce)
	return &icm

}

func (r *Resolvers) GetServiceByID(ctx context.Context, id uint64) *ic.Service {
	key := fmt.Sprintf("%d", id)
	acs := serviceCache.Get(key)
	if acs != nil {
		return acs.(*serviceCacheEntry).m
	}
	db, err := sql.Open()
	if err != nil {
		fmt.Printf("scache2: Failed to connect to db: %s\n", err)
		return nil
	}
	rows, err := db.QueryContext(ctx, "service_by_id", "select id,servicename,userid from service where id = $1", id)
	if err != nil {
		fmt.Printf("Unable to select services from db: %s\n", err)
		return nil
	}
	if !rows.Next() {
		serviceCache.Put(key, &serviceCacheEntry{})
		rows.Close()
		return nil
	}
	icm := ic.Service{}
	err = rows.Scan(&icm.ID, &icm.Name, &icm.UserID)
	rows.Close()
	if err != nil {
		fmt.Printf("failed to scan db for service: %s\n", err)
		return nil
	}
	gce := &serviceCacheEntry{m: &icm}
	serviceCache.Put(key, gce)
	return &icm

}
