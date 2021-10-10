package evaluators

import (
	"flag"
	"fmt"
	ra "golang.conradwood.net/apis/rpcaclapi"
	"golang.conradwood.net/go-easyops/cache"
	"golang.conradwood.net/go-easyops/tokens"
	"strconv"
	"time"
)

var (
	service_rules_cache = cache.NewResolvingCache("service_rules", time.Duration(60), 2000)
	refresh_after       = flag.Int("refresh_service_rules", 600, "age in service_rules cache in `seconds` after which entries will be refreshed")
)

type rc struct {
	sr *ra.ServiceRules
}

func (r *ruleEvaluator) getRulesForServiceID(serviceid uint64) (*ra.ServiceRules, error) {
	service_rules_cache.SetRefreshAfter(time.Duration(*refresh_after) * time.Second)
	key := fmt.Sprintf("%d", serviceid)
	o, err := service_rules_cache.Retrieve(key, func(k string) (interface{}, error) {
		return r.get_service_rules_by_key(key)
	})
	if err != nil {
		return nil, err
	}
	if o != nil {
		return o.(*rc).sr, nil
	}
	return nil, nil
}

func (r *ruleEvaluator) get_service_rules_by_key(key string) (interface{}, error) {
	serviceid, err := strconv.ParseUint(key, 10, 64)
	if err != nil {
		return nil, err
	}
	ctx := tokens.ContextWithToken()
	res, err := r.dbrules.GetRulesForService(ctx, &ra.ServiceID{ID: serviceid})
	if err != nil {
		return nil, err
	}
	return &rc{sr: res}, nil
}
