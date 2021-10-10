package learn

import (
	"context"
	"fmt"
	apb "golang.conradwood.net/apis/auth"
	"golang.conradwood.net/apis/common"
	raa "golang.conradwood.net/apis/rpcaclapi"
	ic "golang.conradwood.net/apis/rpcinterceptor"
	"golang.conradwood.net/go-easyops/auth"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/utils"
)

type Analysis struct {
	proto     *ic.Learnings
	learnings []*ic.Learning
	services  []*Service
	users     []*apb.User
}

type Service struct {
	a    *Analysis
	ID   uint64
	Name string
}

func (s *Service) String() string {
	return fmt.Sprintf("#%02d(%s)", s.ID, s.Name)
}

func FromProto(ctx context.Context, p *ic.Learnings) *Analysis {
	res := &Analysis{proto: p, learnings: p.Learnings}
	res.Services()
	ra := raa.GetRPCACLServiceClient()
	rs, err := ra.GetServices(ctx, &common.Void{})
	if err != nil {
		fmt.Printf("Failed to get services: %s\n", err)
		return nil
	}
	for _, s := range rs.Services {
		svc := res.ServiceByID(s.ID)
		svc.Name = s.Name
	}
	for _, l := range res.learnings {
		if l.UserID == "" {
			continue
		}
		u := res.UserByID(l.UserID)
		if u != nil {
			continue
		}
		u, err := authremote.GetUserByID(ctx, l.UserID)
		if err != nil {
			fmt.Printf("failed to get user: %s\n", utils.ErrorString(err))
			return nil
		}
		res.users = append(res.users, u)
	}
	return res
}
func (a *Analysis) UserByID(id string) *apb.User {
	for _, u := range a.users {
		if u.ID == id {
			return u
		}
	}
	return nil
}

func (a *Analysis) ServiceByID(id uint64) *Service {
	if id == 0 {
		return nil
	}
	for _, s := range a.services {
		if s.ID == id {
			return s
		}
	}
	svc := &Service{a: a, ID: id}
	a.services = append(a.services, svc)
	return svc
}
func (a *Analysis) Services() []*Service {
	var res []*Service
	for _, l := range a.learnings {
		contained := false
		for _, rs := range res {
			if rs.ID == l.ToServiceID {
				contained = true
				break
			}
		}
		if !contained {
			svc := a.ServiceByID(l.ToServiceID)
			res = append(res, svc)
		}
	}
	return res
}

type Call struct {
	From *Service
	To   *Service
	User *apb.User
}

func (c *Call) String() string {
	s := "nobody"
	if c.User != nil {
		s = auth.Description(c.User)
	}
	sv := "[noservice]"
	if c.From != nil {
		sv = c.From.String()
	}
	return fmt.Sprintf("%s (as %s)", sv, s)
}

func (s *Service) Callers() []*Call {
	var res []*Call
	for _, l := range s.a.learnings {
		if l.ToServiceID != s.ID {
			continue
		}
		c := &Call{User: s.a.UserByID(l.UserID),
			From: s.a.ServiceByID(l.FromServiceID),
			To:   s.a.ServiceByID(l.ToServiceID),
		}
		res = append(res, c)
	}
	return res
}
