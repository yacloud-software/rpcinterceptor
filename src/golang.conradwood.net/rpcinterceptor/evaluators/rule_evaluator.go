package evaluators

import (
	"flag"
	"fmt"
	apb "golang.conradwood.net/apis/auth"
	ra "golang.conradwood.net/apis/rpcaclapi"
	"golang.conradwood.net/go-easyops/auth"
)

var (
	debug = flag.Bool("debug_evaluate", false, "debug the rule evaluators")
)

type RuleEvaluator interface {
	EvaluateCall(service *ra.Service, CallingService *apb.User, CallingUser *apb.User) (bool, error)
	EvaluateRules(rules *ra.ServiceRules, CallingService *apb.User, CallingUser *apb.User) bool
}

func NewRuleEvaluator() RuleEvaluator {
	StartDB()
	res := &ruleEvaluator{}
	service_rules_cache.SetAsyncRetriever(res.get_service_rules_by_key)
	return res
}

type ruleEvaluator struct {
	dbrules *DBRules
}

// given a set of service rules and a callingservice/user -> evaluate. return true if grant
func (r *ruleEvaluator) EvaluateRules(rules *ra.ServiceRules, CallingService *apb.User, CallingUser *apb.User) bool {
	logger := newLogger()
	if rules == nil || len(rules.Rules) == 0 {
		if *debug {
			logger.Printf("no rules\n")
		}
		return false
	}
	actualMatch := ra.RuleMatch_NEITHER
	if CallingService != nil && CallingUser != nil {
		actualMatch = ra.RuleMatch_SERVICE_AND_USER
	} else if CallingService != nil {
		actualMatch = ra.RuleMatch_SERVICE_ONLY
	} else if CallingUser != nil {
		actualMatch = ra.RuleMatch_USER_ONLY
	}

	srcmatches := 0
	for _, r := range rules.Rules {
		if !ruleMatchesSource(r, actualMatch) {
			continue
		}
		srcmatches++
		if !serviceMatch(r, CallingService) {
			continue
		}
		if !userMatch(r, CallingUser) {
			continue
		}
		if *debug {
			logger.Printf("Rule #%d matched\n", r.ID)
		}
		return r.Grant
	}
	logger.Printf("No rules matched (%d) evaluated. SourceMatches=%d\n", len(rules.Rules), srcmatches)
	return false
}

// check if what we actually have (combo user/service) matches this rules. (true if it does)
func ruleMatchesSource(rule *ra.Rule, actual ra.RuleMatch) bool {
	for _, m := range rule.Match {
		if m == ra.RuleMatch_INVALID {
			fmt.Printf("WARNING: Rule #%d has an invalid rulematcher", rule.ID)
			continue
		}
		if m == actual {
			return true
		}
	}
	return false
}
func serviceMatch(rule *ra.Rule, svc *apb.User) bool {
	if svc == nil || len(rule.FromServiceID) == 0 {
		return true
	}
	for _, f := range rule.FromServiceID {
		if svc.ID == f {
			return true
		}
	}
	return false
}
func userMatch(rule *ra.Rule, user *apb.User) bool {
	if user == nil || len(rule.GroupID) == 0 {
		return true
	}
	for _, g := range rule.GroupID {
		for _, ug := range user.Groups {
			if ug.ID == g {
				return true
			}
		}
	}
	return false
}

// get services rules and evaluate rules
func (r *ruleEvaluator) EvaluateCall(service *ra.Service, CallingService *apb.User, CallingUser *apb.User) (bool, error) {
	if service == nil {
		panic("No service to check rules on")
	}
	rules, err := r.getRulesForServiceID(service.ID)
	if err != nil {
		fmt.Printf("Error getting rules for service %s: %s\n", service.Name, err)
		return false, err
	}
	result := r.EvaluateRules(rules, CallingService, CallingUser)
	if *debug {
		fmt.Printf("Result for rules for access by user %s (coming from service %s) to service %s (#%d): %v\n",
			auth.UserIDString(CallingUser),
			auth.UserIDString(CallingService),
			service.Name, service.ID, result)
	}
	return result, nil
}
