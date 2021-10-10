package main

import (
	"context"
	ra "golang.conradwood.net/apis/rpcaclapi"
)

// todo delegate to rule evaluator
func (e *rpcACLServer) GetRulesForService(ctx context.Context, req *ra.ServiceID) (*ra.ServiceRules, error) {
	return dbrules.GetRulesForService(ctx, req)
}

// todo delegate to rule evaluator
func (e *rpcACLServer) GetRulesForServices(ctx context.Context, req *ra.ServiceIDList) (*ra.ServiceRules, error) {
	return dbrules.GetRulesForServices(ctx, req)
}
func (e *rpcACLServer) AuthoriseGroupService(ctx context.Context, req *ra.GroupServiceRequest) (*ra.Service, error) {
	return dbrules.AuthoriseGroupService(ctx, req)
}
