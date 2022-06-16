package resolvers

import (
	"fmt"
	ic "golang.conradwood.net/apis/rpcaclapi"
	"golang.conradwood.net/go-easyops/cache"
	"golang.conradwood.net/go-easyops/sql"
	"golang.org/x/net/context"
	"time"
)

var (
	methodCache     = cache.New("methodcache", time.Duration(5)*time.Minute, 100)
	methodNameCache = cache.New("methodnamecache", time.Duration(5)*time.Minute, 100)
)

type methodCacheEntry struct {
	m *ic.Method
}

func (r *Resolvers) GetMethodByName(ctx context.Context, service string, method string) *ic.Method {
	key := fmt.Sprintf("%s-%s", service, method)
	svc := r.GetServiceByName(ctx, service)
	if svc == nil {
		return nil
	}
	acs := methodNameCache.Get(key)
	if acs != nil {
		return acs.(*methodCacheEntry).m
	}
	db, err := sql.Open()
	if err != nil {
		fmt.Printf("mcache1: Failed to connect to db: %s\n", err)
		return nil
	}
	rows, err := db.QueryContext(ctx, "method_by_name_and_serviceid", "select id,methodname,service_id from method where methodname = $1 and service_id = $2", method, svc.ID)
	if err != nil {
		fmt.Printf("Unable to select methods from db: %s\n", err)
		return nil
	}
	defer rows.Close()
	if !rows.Next() {
		methodNameCache.Put(key, &methodCacheEntry{})
		return nil
	}
	icm := ic.Method{}
	err = rows.Scan(&icm.ID, &icm.Name, &icm.ServiceID)
	if err != nil {
		fmt.Printf("failed to scan db for method: %s\n", err)
		return nil
	}
	gce := &methodCacheEntry{m: &icm}
	methodNameCache.Put(key, gce)
	return &icm
}

func (r *Resolvers) GetMethodByID(ctx context.Context, id uint64) *ic.Method {
	key := fmt.Sprintf("%d", id)
	acs := methodCache.Get(key)
	if acs != nil {
		return acs.(*methodCacheEntry).m
	}
	db, err := sql.Open()
	if err != nil {
		fmt.Printf("mache2: Failed to connect to db: %s\n", err)
		return nil
	}
	rows, err := db.QueryContext(ctx, "method_by_id", "select id,methodname,service_id from method where id = $1", id)
	if err != nil {
		fmt.Printf("Unable to select methods from db: %s\n", err)
		return nil
	}
	if !rows.Next() {
		methodCache.Put(key, &methodCacheEntry{})
		rows.Close()
		return nil
	}
	icm := ic.Method{}
	err = rows.Scan(&icm.ID, &icm.Name, &icm.ServiceID)
	rows.Close()
	if err != nil {
		fmt.Printf("failed to scan db for method: %s\n", err)
		return nil
	}
	gce := &methodCacheEntry{m: &icm}
	methodCache.Put(key, gce)
	return &icm

}
