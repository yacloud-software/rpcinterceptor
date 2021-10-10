package evaluators

/*
* read/write rules from/to database
* the API is as-defined in rpcaclapi.proto
 */
import (
	"context"
	gsql "database/sql"
	"fmt"
	ra "golang.conradwood.net/apis/rpcaclapi"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/sql"
	db "golang.conradwood.net/rpcdb"
	"strings"
)

var (
	store_dbrules       *db.DBDBRule
	store_dbrulematch   *db.DBDBRuleMatch
	store_dbruleservice *db.DBDBRuleService
	store_dbrulegroup   *db.DBDBRuleGroup
	psql                *sql.DB
)

// must be called before anything else.
func StartDB() error {
	var err error
	if psql == nil {
		psql, err = sql.Open()
		if err != nil {
			return err
		}
	}
	if store_dbrules == nil {
		store_dbrules = db.NewDBDBRule(psql)
	}
	if store_dbrulematch == nil {
		store_dbrulematch = db.NewDBDBRuleMatch(psql)
	}
	if store_dbruleservice == nil {
		store_dbruleservice = db.NewDBDBRuleService(psql)
	}
	if store_dbrulegroup == nil {
		store_dbrulegroup = db.NewDBDBRuleGroup(psql)
	}
	return nil
}

type DBRules struct {
}

func (d *DBRules) GetRulesForService(ctx context.Context, req *ra.ServiceID) (*ra.ServiceRules, error) {
	rl := &ruleLoader{ServiceID: req.ID, ctx: ctx}
	res, err := rl.Load()
	if err != nil {
		return nil, err
	}
	return res, err
}
func (d *DBRules) GetRulesForServices(ctx context.Context, req *ra.ServiceIDList) (*ra.ServiceRules, error) {
	return nil, errors.NotImplemented(ctx, "getrulesforservices()")

}

func (d *DBRules) AuthoriseGroupService(ctx context.Context, req *ra.GroupServiceRequest) (*ra.Service, error) {
	var err error
	if len(req.GroupID) == 0 {
		return nil, errors.InvalidArgs(ctx, "groupname too short", "missing groupname")
	}
	var svc *ra.Service
	sid := req.ServiceID
	if sid == 0 {
		if len(req.ServiceName) < 3 {
			return nil, errors.InvalidArgs(ctx, "servicename too short", "servicename must be at least 3 chars (or serviceid provided")
		}
		svcs, err := getServiceMatchName(ctx, req.ServiceName)
		if err != nil {
			return nil, err
		}
		var fs []*ra.Service
		for _, s := range svcs {
			if strings.Contains(s.Name, req.ServiceName) {
				fs = append(fs, s)
			}
		}
		if len(fs) == 0 {
			return nil, errors.NotFound(ctx, "no service matching %s found", req.ServiceName)
		}
		if len(fs) > 1 {
			return nil, errors.InvalidArgs(ctx, "servicename multiple matches", "servicename \"%s\" matched %d services", req.ServiceName, len(fs))
		}
		svc = fs[0]
	} else {
		svc, err = getServiceByID(ctx, sid)
		if err != nil {
			return nil, err
		}
	}
	sid = svc.ID
	fmt.Printf("Adding group %s to service #%d\n", req.GroupID, sid)

	// TODO - check if it is authorised already and if group exists

	rule := &ra.DBRule{ServiceID: sid, ResultOnMatch: true}
	id, err := store_dbrules.Save(ctx, rule)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Created rule #%d\n", id)

	/*		Match:   []RuleMatch{ra.RuleMatch_USER_ONLY, ra.RuleMatch_SERVICE_AND_USER},
		GroupID: []string{req.GroupID},
		Grant:   true,
	}
	*/
	_, err = store_dbrulematch.Save(ctx, &ra.DBRuleMatch{RuleID: id, RuleMatch: uint32(ra.RuleMatch_USER_ONLY)})
	if err != nil {
		return nil, err
	}
	_, err = store_dbrulematch.Save(ctx, &ra.DBRuleMatch{RuleID: id, RuleMatch: uint32(ra.RuleMatch_SERVICE_AND_USER)})
	if err != nil {
		return nil, err
	}
	_, err = store_dbrulegroup.Save(ctx, &ra.DBRuleGroup{RuleID: id, GroupID: req.GroupID})
	if err != nil {
		return nil, err
	}
	return svc, nil
}

/********************************************************************************************
database access
********************************************************************************************/

type ruleLoader struct {
	ServiceID uint64
	ctx       context.Context
	err       error
	res       *ra.ServiceRules
}

func (r *ruleLoader) Load() (*ra.ServiceRules, error) {
	r.res = &ra.ServiceRules{ServiceID: r.ServiceID}
	dbrules, err := store_dbrules.ByServiceID(r.ctx, r.ServiceID)
	if err != nil {
		return nil, err
	}
	for _, dbr := range dbrules {
		r.res.Rules = append(r.res.Rules, &ra.Rule{
			ID:    dbr.ID,
			Grant: dbr.ResultOnMatch,
		})
	}
	for _, rule := range r.res.Rules {
		r.LoadMatchers(rule)
		r.LoadServices(rule)
		r.LoadGroups(rule)
	}
	return r.res, r.err
}

func (r *ruleLoader) LoadMatchers(rule *ra.Rule) {
	ms, err := store_dbrulematch.ByRuleID(r.ctx, rule.ID)
	if err != nil {
		r.err = err
		return
	}
	for _, m := range ms {
		rule.Match = append(rule.Match, ra.RuleMatch(m.RuleMatch))
	}
}
func (r *ruleLoader) LoadServices(rule *ra.Rule) {
	gs, err := store_dbruleservice.ByRuleID(r.ctx, rule.ID)
	if err != nil {
		r.err = err
		return
	}
	for _, g := range gs {
		rule.FromServiceID = append(rule.FromServiceID, g.ServiceUserID)
	}
}
func (r *ruleLoader) LoadGroups(rule *ra.Rule) {
	gs, err := store_dbrulegroup.ByRuleID(r.ctx, rule.ID)
	if err != nil {
		r.err = err
		return
	}
	for _, g := range gs {
		rule.GroupID = append(rule.GroupID, g.GroupID)
	}
}

func getServiceMatchName(ctx context.Context, name string) ([]*ra.Service, error) {
	rows, err := psql.QueryContext(ctx, "servicenamematch", "select id,servicename,userid from service where servicename like $1", "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return servicerows(rows)
}
func getServiceByID(ctx context.Context, svcid uint64) (*ra.Service, error) {
	rows, err := psql.QueryContext(ctx, "servicebyid", "select id,servicename,userid from service where id = $1", svcid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	svcs, err := servicerows(rows)
	if err != nil {
		return nil, err
	}
	if len(svcs) < 1 {
		return nil, errors.NotFound(ctx, "no service with id %d", svcid)
	}
	return svcs[0], nil
}
func servicerows(rows *gsql.Rows) ([]*ra.Service, error) {
	var res []*ra.Service
	for rows.Next() {
		n := &ra.Service{}
		err := rows.Scan(&n.ID, &n.Name, &n.UserID)
		if err != nil {
			return nil, err
		}
		res = append(res, n)
	}
	return res, nil
}
