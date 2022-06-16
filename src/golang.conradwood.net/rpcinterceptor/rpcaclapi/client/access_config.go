package main

// work out what access is required from an audit

import (
	"golang.conradwood.net/apis/auth"
	"golang.conradwood.net/apis/rpcaclapi"
)

type AccessConfig struct {
	Methods []*AccessMethod
	calls   []*rpcaclapi.Call
}
type AccessMethod struct {
	accessConfig    *AccessConfig
	TargetMethod    *rpcaclapi.FullMethod
	CallingServices []*auth.User
	CallingUsers    []*auth.User
}
type AccessConfigSummary struct {
	Methods []*AccessMethodSummary
}
type AccessMethodSummary struct {
	TargetMethodID            uint64
	CallingServicesWithNoUser []string `yaml:"callingserviceswithnouser,omitempty"`
	CallingGroupIDs           []string `yaml:"callinggroups,omitempty"`
}

func CreateAccessConfig(calls *rpcaclapi.CallList) (*AccessConfig, error) {
	res := &AccessConfig{}
	for _, c := range calls.Calls {
		res.Add(c)
	}
	return res, nil
}

func (ac *AccessConfig) Add(c *rpcaclapi.Call) {
	ac.calls = append(ac.calls, c)
	am := ac.Method(c.CalledMethod)
	if am == nil {
		am = &AccessMethod{TargetMethod: c.CalledMethod}
		am.accessConfig = ac
		ac.Methods = append(ac.Methods, am)
	}
	if c.CallingService != nil {
		am.CallingServices = append(am.CallingServices, c.CallingService)
	}
	if c.CallingUser != nil {
		am.CallingUsers = append(am.CallingUsers, c.CallingUser)
	}
}

func (ac *AccessConfig) Method(m *rpcaclapi.FullMethod) *AccessMethod {
	for _, xm := range ac.Methods {
		if m.ID == xm.TargetMethod.ID {
			return xm
		}
	}
	return nil
}
func (am *AccessMethod) CallingServicesWithNoUser() []*auth.User {
	var res []*auth.User
	for _, c := range am.accessConfig.calls {
		if c.CalledMethod.ID != am.TargetMethod.ID {
			continue
		}
		if c.CallingUser != nil {
			continue
		}
		if c.CallingService == nil {
			continue
		}
		res = append(res, c.CallingService)
	}
	return res
}
func (am *AccessMethod) CallingUserGroups() []*auth.Group {
	g := &GroupOptimizer{Users: am.CallingUsers}
	return g.Groups()

}

func (am *AccessMethod) Summary() *AccessMethodSummary {
	res := &AccessMethodSummary{TargetMethodID: am.TargetMethod.ID}
	for _, csnu := range am.CallingServicesWithNoUser() {
		res.CallingServicesWithNoUser = append(res.CallingServicesWithNoUser, csnu.ID)
	}
	for _, cg := range am.CallingUserGroups() {
		res.CallingGroupIDs = append(res.CallingGroupIDs, cg.ID)
	}
	return res
}

func (ac *AccessConfig) Summary() *AccessConfigSummary {
	res := &AccessConfigSummary{}
	for _, xm := range ac.Methods {
		ams := xm.Summary()
		res.Methods = append(res.Methods, ams)
	}
	return res
}
