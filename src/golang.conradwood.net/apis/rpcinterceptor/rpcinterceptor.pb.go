// Code generated by protoc-gen-go.
// source: protos/golang.conradwood.net/apis/rpcinterceptor/rpcinterceptor.proto
// DO NOT EDIT!

/*
Package rpcinterceptor is a generated protocol buffer package.

It is generated from these files:
	protos/golang.conradwood.net/apis/rpcinterceptor/rpcinterceptor.proto

It has these top-level messages:
	InterceptRPCResponse
	ServiceByUserIDRequest
	Service
	InterceptRPCRequest
	CTXRoutingTags
	InMetadata
	LogErrorRequest
	ServiceIDRequest
	ServiceIDResponse
	Learning
	Learnings
*/
package rpcinterceptor

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import auth "golang.conradwood.net/apis/auth"
import common "golang.conradwood.net/apis/common"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RejectReason int32

const (
	RejectReason_NonSpecific     RejectReason = 0
	RejectReason_UserRejected    RejectReason = 1
	RejectReason_ServiceRejected RejectReason = 2
	RejectReason_OrgRejected     RejectReason = 3
	RejectReason_UserMissing     RejectReason = 4
	RejectReason_ServiceMissing  RejectReason = 5
	RejectReason_OrgMissing      RejectReason = 6
)

var RejectReason_name = map[int32]string{
	0: "NonSpecific",
	1: "UserRejected",
	2: "ServiceRejected",
	3: "OrgRejected",
	4: "UserMissing",
	5: "ServiceMissing",
	6: "OrgMissing",
}
var RejectReason_value = map[string]int32{
	"NonSpecific":     0,
	"UserRejected":    1,
	"ServiceRejected": 2,
	"OrgRejected":     3,
	"UserMissing":     4,
	"ServiceMissing":  5,
	"OrgMissing":      6,
}

func (x RejectReason) String() string {
	return proto.EnumName(RejectReason_name, int32(x))
}
func (RejectReason) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type InterceptRPCResponse struct {
	RequestID            string           `protobuf:"bytes,1,opt,name=RequestID" json:"RequestID,omitempty"`
	CallerService        *auth.User       `protobuf:"bytes,2,opt,name=CallerService" json:"CallerService,omitempty"`
	CallerUser           *auth.User       `protobuf:"bytes,3,opt,name=CallerUser" json:"CallerUser,omitempty"`
	CallerSudoUser       *auth.User       `protobuf:"bytes,4,opt,name=CallerSudoUser" json:"CallerSudoUser,omitempty"`
	Reject               bool             `protobuf:"varint,6,opt,name=Reject" json:"Reject,omitempty"`
	RejectReason         RejectReason     `protobuf:"varint,7,opt,name=RejectReason,enum=rpcinterceptor.RejectReason" json:"RejectReason,omitempty"`
	CallerMethodID       uint64           `protobuf:"varint,8,opt,name=CallerMethodID" json:"CallerMethodID,omitempty"`
	Source               string           `protobuf:"bytes,9,opt,name=Source" json:"Source,omitempty"`
	CalleeServiceID      uint64           `protobuf:"varint,10,opt,name=CalleeServiceID" json:"CalleeServiceID,omitempty"`
	SignedCallerService  *auth.SignedUser `protobuf:"bytes,11,opt,name=SignedCallerService" json:"SignedCallerService,omitempty"`
	SignedCallerUser     *auth.SignedUser `protobuf:"bytes,12,opt,name=SignedCallerUser" json:"SignedCallerUser,omitempty"`
	SignedCallerSudoUser *auth.SignedUser `protobuf:"bytes,13,opt,name=SignedCallerSudoUser" json:"SignedCallerSudoUser,omitempty"`
}

func (m *InterceptRPCResponse) Reset()                    { *m = InterceptRPCResponse{} }
func (m *InterceptRPCResponse) String() string            { return proto.CompactTextString(m) }
func (*InterceptRPCResponse) ProtoMessage()               {}
func (*InterceptRPCResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *InterceptRPCResponse) GetRequestID() string {
	if m != nil {
		return m.RequestID
	}
	return ""
}

func (m *InterceptRPCResponse) GetCallerService() *auth.User {
	if m != nil {
		return m.CallerService
	}
	return nil
}

func (m *InterceptRPCResponse) GetCallerUser() *auth.User {
	if m != nil {
		return m.CallerUser
	}
	return nil
}

func (m *InterceptRPCResponse) GetCallerSudoUser() *auth.User {
	if m != nil {
		return m.CallerSudoUser
	}
	return nil
}

func (m *InterceptRPCResponse) GetReject() bool {
	if m != nil {
		return m.Reject
	}
	return false
}

func (m *InterceptRPCResponse) GetRejectReason() RejectReason {
	if m != nil {
		return m.RejectReason
	}
	return RejectReason_NonSpecific
}

func (m *InterceptRPCResponse) GetCallerMethodID() uint64 {
	if m != nil {
		return m.CallerMethodID
	}
	return 0
}

func (m *InterceptRPCResponse) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *InterceptRPCResponse) GetCalleeServiceID() uint64 {
	if m != nil {
		return m.CalleeServiceID
	}
	return 0
}

func (m *InterceptRPCResponse) GetSignedCallerService() *auth.SignedUser {
	if m != nil {
		return m.SignedCallerService
	}
	return nil
}

func (m *InterceptRPCResponse) GetSignedCallerUser() *auth.SignedUser {
	if m != nil {
		return m.SignedCallerUser
	}
	return nil
}

func (m *InterceptRPCResponse) GetSignedCallerSudoUser() *auth.SignedUser {
	if m != nil {
		return m.SignedCallerSudoUser
	}
	return nil
}

type ServiceByUserIDRequest struct {
	UserID string `protobuf:"bytes,1,opt,name=UserID" json:"UserID,omitempty"`
}

func (m *ServiceByUserIDRequest) Reset()                    { *m = ServiceByUserIDRequest{} }
func (m *ServiceByUserIDRequest) String() string            { return proto.CompactTextString(m) }
func (*ServiceByUserIDRequest) ProtoMessage()               {}
func (*ServiceByUserIDRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ServiceByUserIDRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type Service struct {
	ID     uint64 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	UserID string `protobuf:"bytes,3,opt,name=UserID" json:"UserID,omitempty"`
}

func (m *Service) Reset()                    { *m = Service{} }
func (m *Service) String() string            { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()               {}
func (*Service) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Service) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type InterceptRPCRequest struct {
	InMetadata *InMetadata `protobuf:"bytes,1,opt,name=InMetadata" json:"InMetadata,omitempty"`
	Service    string      `protobuf:"bytes,2,opt,name=Service" json:"Service,omitempty"`
	Method     string      `protobuf:"bytes,3,opt,name=Method" json:"Method,omitempty"`
	Source     string      `protobuf:"bytes,4,opt,name=Source" json:"Source,omitempty"`
}

func (m *InterceptRPCRequest) Reset()                    { *m = InterceptRPCRequest{} }
func (m *InterceptRPCRequest) String() string            { return proto.CompactTextString(m) }
func (*InterceptRPCRequest) ProtoMessage()               {}
func (*InterceptRPCRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *InterceptRPCRequest) GetInMetadata() *InMetadata {
	if m != nil {
		return m.InMetadata
	}
	return nil
}

func (m *InterceptRPCRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *InterceptRPCRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *InterceptRPCRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

type CTXRoutingTags struct {
	Tags            map[string]string `protobuf:"bytes,1,rep,name=Tags" json:"Tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	FallbackToPlain bool              `protobuf:"varint,2,opt,name=FallbackToPlain" json:"FallbackToPlain,omitempty"`
	Propagate       bool              `protobuf:"varint,3,opt,name=Propagate" json:"Propagate,omitempty"`
}

func (m *CTXRoutingTags) Reset()                    { *m = CTXRoutingTags{} }
func (m *CTXRoutingTags) String() string            { return proto.CompactTextString(m) }
func (*CTXRoutingTags) ProtoMessage()               {}
func (*CTXRoutingTags) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CTXRoutingTags) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *CTXRoutingTags) GetFallbackToPlain() bool {
	if m != nil {
		return m.FallbackToPlain
	}
	return false
}

func (m *CTXRoutingTags) GetPropagate() bool {
	if m != nil {
		return m.Propagate
	}
	return false
}

// the stuff we're transporting within a context between services
type InMetadata struct {
	RequestID       string              `protobuf:"bytes,1,opt,name=RequestID" json:"RequestID,omitempty"`
	FooBar          string              `protobuf:"bytes,2,opt,name=FooBar" json:"FooBar,omitempty"`
	UserToken       string              `protobuf:"bytes,3,opt,name=UserToken" json:"UserToken,omitempty"`
	ServiceToken    string              `protobuf:"bytes,4,opt,name=ServiceToken" json:"ServiceToken,omitempty"`
	UserID          string              `protobuf:"bytes,5,opt,name=UserID" json:"UserID,omitempty"`
	CallerMethodID  uint64              `protobuf:"varint,7,opt,name=CallerMethodID" json:"CallerMethodID,omitempty"`
	CallerServiceID uint64              `protobuf:"varint,8,opt,name=CallerServiceID" json:"CallerServiceID,omitempty"`
	RoutingInfo     uint32              `protobuf:"varint,9,opt,name=RoutingInfo" json:"RoutingInfo,omitempty"`
	Version         uint32              `protobuf:"varint,10,opt,name=Version" json:"Version,omitempty"`
	Service         *auth.User          `protobuf:"bytes,11,opt,name=Service" json:"Service,omitempty"`
	User            *auth.User          `protobuf:"bytes,12,opt,name=User" json:"User,omitempty"`
	SignedService   *auth.SignedUser    `protobuf:"bytes,13,opt,name=SignedService" json:"SignedService,omitempty"`
	SignedUser      *auth.SignedUser    `protobuf:"bytes,14,opt,name=SignedUser" json:"SignedUser,omitempty"`
	Trace           bool                `protobuf:"varint,15,opt,name=Trace" json:"Trace,omitempty"`
	Debug           bool                `protobuf:"varint,16,opt,name=Debug" json:"Debug,omitempty"`
	SignedSession   *auth.SignedSession `protobuf:"bytes,17,opt,name=SignedSession" json:"SignedSession,omitempty"`
	RoutingTags     *CTXRoutingTags     `protobuf:"bytes,18,opt,name=RoutingTags" json:"RoutingTags,omitempty"`
}

func (m *InMetadata) Reset()                    { *m = InMetadata{} }
func (m *InMetadata) String() string            { return proto.CompactTextString(m) }
func (*InMetadata) ProtoMessage()               {}
func (*InMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *InMetadata) GetRequestID() string {
	if m != nil {
		return m.RequestID
	}
	return ""
}

func (m *InMetadata) GetFooBar() string {
	if m != nil {
		return m.FooBar
	}
	return ""
}

func (m *InMetadata) GetUserToken() string {
	if m != nil {
		return m.UserToken
	}
	return ""
}

func (m *InMetadata) GetServiceToken() string {
	if m != nil {
		return m.ServiceToken
	}
	return ""
}

func (m *InMetadata) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *InMetadata) GetCallerMethodID() uint64 {
	if m != nil {
		return m.CallerMethodID
	}
	return 0
}

func (m *InMetadata) GetCallerServiceID() uint64 {
	if m != nil {
		return m.CallerServiceID
	}
	return 0
}

func (m *InMetadata) GetRoutingInfo() uint32 {
	if m != nil {
		return m.RoutingInfo
	}
	return 0
}

func (m *InMetadata) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *InMetadata) GetService() *auth.User {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *InMetadata) GetUser() *auth.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *InMetadata) GetSignedService() *auth.SignedUser {
	if m != nil {
		return m.SignedService
	}
	return nil
}

func (m *InMetadata) GetSignedUser() *auth.SignedUser {
	if m != nil {
		return m.SignedUser
	}
	return nil
}

func (m *InMetadata) GetTrace() bool {
	if m != nil {
		return m.Trace
	}
	return false
}

func (m *InMetadata) GetDebug() bool {
	if m != nil {
		return m.Debug
	}
	return false
}

func (m *InMetadata) GetSignedSession() *auth.SignedSession {
	if m != nil {
		return m.SignedSession
	}
	return nil
}

func (m *InMetadata) GetRoutingTags() *CTXRoutingTags {
	if m != nil {
		return m.RoutingTags
	}
	return nil
}

type LogErrorRequest struct {
	InMetadata     *InMetadata `protobuf:"bytes,1,opt,name=InMetadata" json:"InMetadata,omitempty"`
	Service        string      `protobuf:"bytes,2,opt,name=Service" json:"Service,omitempty"`
	Method         string      `protobuf:"bytes,3,opt,name=Method" json:"Method,omitempty"`
	ErrorCode      uint32      `protobuf:"varint,4,opt,name=ErrorCode" json:"ErrorCode,omitempty"`
	DisplayMessage string      `protobuf:"bytes,5,opt,name=DisplayMessage" json:"DisplayMessage,omitempty"`
	LogMessage     string      `protobuf:"bytes,6,opt,name=LogMessage" json:"LogMessage,omitempty"`
}

func (m *LogErrorRequest) Reset()                    { *m = LogErrorRequest{} }
func (m *LogErrorRequest) String() string            { return proto.CompactTextString(m) }
func (*LogErrorRequest) ProtoMessage()               {}
func (*LogErrorRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *LogErrorRequest) GetInMetadata() *InMetadata {
	if m != nil {
		return m.InMetadata
	}
	return nil
}

func (m *LogErrorRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *LogErrorRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *LogErrorRequest) GetErrorCode() uint32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *LogErrorRequest) GetDisplayMessage() string {
	if m != nil {
		return m.DisplayMessage
	}
	return ""
}

func (m *LogErrorRequest) GetLogMessage() string {
	if m != nil {
		return m.LogMessage
	}
	return ""
}

type ServiceIDRequest struct {
	Token  string `protobuf:"bytes,1,opt,name=Token" json:"Token,omitempty"`
	MyName string `protobuf:"bytes,2,opt,name=MyName" json:"MyName,omitempty"`
}

func (m *ServiceIDRequest) Reset()                    { *m = ServiceIDRequest{} }
func (m *ServiceIDRequest) String() string            { return proto.CompactTextString(m) }
func (*ServiceIDRequest) ProtoMessage()               {}
func (*ServiceIDRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ServiceIDRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ServiceIDRequest) GetMyName() string {
	if m != nil {
		return m.MyName
	}
	return ""
}

type ServiceIDResponse struct {
	ServiceID uint64 `protobuf:"varint,2,opt,name=ServiceID" json:"ServiceID,omitempty"`
}

func (m *ServiceIDResponse) Reset()                    { *m = ServiceIDResponse{} }
func (m *ServiceIDResponse) String() string            { return proto.CompactTextString(m) }
func (*ServiceIDResponse) ProtoMessage()               {}
func (*ServiceIDResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ServiceIDResponse) GetServiceID() uint64 {
	if m != nil {
		return m.ServiceID
	}
	return 0
}

type Learning struct {
	ToServiceID   uint64 `protobuf:"varint,1,opt,name=ToServiceID" json:"ToServiceID,omitempty"`
	FromServiceID uint64 `protobuf:"varint,2,opt,name=FromServiceID" json:"FromServiceID,omitempty"`
	UserID        string `protobuf:"bytes,3,opt,name=UserID" json:"UserID,omitempty"`
	Count         uint64 `protobuf:"varint,4,opt,name=Count" json:"Count,omitempty"`
}

func (m *Learning) Reset()                    { *m = Learning{} }
func (m *Learning) String() string            { return proto.CompactTextString(m) }
func (*Learning) ProtoMessage()               {}
func (*Learning) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Learning) GetToServiceID() uint64 {
	if m != nil {
		return m.ToServiceID
	}
	return 0
}

func (m *Learning) GetFromServiceID() uint64 {
	if m != nil {
		return m.FromServiceID
	}
	return 0
}

func (m *Learning) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *Learning) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Learnings struct {
	Learnings []*Learning `protobuf:"bytes,1,rep,name=Learnings" json:"Learnings,omitempty"`
}

func (m *Learnings) Reset()                    { *m = Learnings{} }
func (m *Learnings) String() string            { return proto.CompactTextString(m) }
func (*Learnings) ProtoMessage()               {}
func (*Learnings) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Learnings) GetLearnings() []*Learning {
	if m != nil {
		return m.Learnings
	}
	return nil
}

func init() {
	proto.RegisterType((*InterceptRPCResponse)(nil), "rpcinterceptor.InterceptRPCResponse")
	proto.RegisterType((*ServiceByUserIDRequest)(nil), "rpcinterceptor.ServiceByUserIDRequest")
	proto.RegisterType((*Service)(nil), "rpcinterceptor.Service")
	proto.RegisterType((*InterceptRPCRequest)(nil), "rpcinterceptor.InterceptRPCRequest")
	proto.RegisterType((*CTXRoutingTags)(nil), "rpcinterceptor.CTXRoutingTags")
	proto.RegisterType((*InMetadata)(nil), "rpcinterceptor.InMetadata")
	proto.RegisterType((*LogErrorRequest)(nil), "rpcinterceptor.LogErrorRequest")
	proto.RegisterType((*ServiceIDRequest)(nil), "rpcinterceptor.ServiceIDRequest")
	proto.RegisterType((*ServiceIDResponse)(nil), "rpcinterceptor.ServiceIDResponse")
	proto.RegisterType((*Learning)(nil), "rpcinterceptor.Learning")
	proto.RegisterType((*Learnings)(nil), "rpcinterceptor.Learnings")
	proto.RegisterEnum("rpcinterceptor.RejectReason", RejectReason_name, RejectReason_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RPCInterceptorService service

type RPCInterceptorServiceClient interface {
	// extract useful information from a context
	InterceptRPC(ctx context.Context, in *InterceptRPCRequest, opts ...grpc.CallOption) (*InterceptRPCResponse, error)
	// log an rpc error
	LogError(ctx context.Context, in *LogErrorRequest, opts ...grpc.CallOption) (*common.Void, error)
	// get a serviceID by token (bootstrapping servers)
	GetMyServiceID(ctx context.Context, in *ServiceIDRequest, opts ...grpc.CallOption) (*ServiceIDResponse, error)
	// if in learning mode, we can retrieve what has been learned so far:
	GetLearnings(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*Learnings, error)
	// reset learnings...
	ClearLearnings(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*common.Void, error)
	// given a userid for a service, this will return the rpcinterceptor service
	GetServiceByUserID(ctx context.Context, in *ServiceByUserIDRequest, opts ...grpc.CallOption) (*Service, error)
}

type rPCInterceptorServiceClient struct {
	cc *grpc.ClientConn
}

func NewRPCInterceptorServiceClient(cc *grpc.ClientConn) RPCInterceptorServiceClient {
	return &rPCInterceptorServiceClient{cc}
}

func (c *rPCInterceptorServiceClient) InterceptRPC(ctx context.Context, in *InterceptRPCRequest, opts ...grpc.CallOption) (*InterceptRPCResponse, error) {
	out := new(InterceptRPCResponse)
	err := grpc.Invoke(ctx, "/rpcinterceptor.RPCInterceptorService/InterceptRPC", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCInterceptorServiceClient) LogError(ctx context.Context, in *LogErrorRequest, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/rpcinterceptor.RPCInterceptorService/LogError", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCInterceptorServiceClient) GetMyServiceID(ctx context.Context, in *ServiceIDRequest, opts ...grpc.CallOption) (*ServiceIDResponse, error) {
	out := new(ServiceIDResponse)
	err := grpc.Invoke(ctx, "/rpcinterceptor.RPCInterceptorService/GetMyServiceID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCInterceptorServiceClient) GetLearnings(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*Learnings, error) {
	out := new(Learnings)
	err := grpc.Invoke(ctx, "/rpcinterceptor.RPCInterceptorService/GetLearnings", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCInterceptorServiceClient) ClearLearnings(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/rpcinterceptor.RPCInterceptorService/ClearLearnings", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCInterceptorServiceClient) GetServiceByUserID(ctx context.Context, in *ServiceByUserIDRequest, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := grpc.Invoke(ctx, "/rpcinterceptor.RPCInterceptorService/GetServiceByUserID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RPCInterceptorService service

type RPCInterceptorServiceServer interface {
	// extract useful information from a context
	InterceptRPC(context.Context, *InterceptRPCRequest) (*InterceptRPCResponse, error)
	// log an rpc error
	LogError(context.Context, *LogErrorRequest) (*common.Void, error)
	// get a serviceID by token (bootstrapping servers)
	GetMyServiceID(context.Context, *ServiceIDRequest) (*ServiceIDResponse, error)
	// if in learning mode, we can retrieve what has been learned so far:
	GetLearnings(context.Context, *common.Void) (*Learnings, error)
	// reset learnings...
	ClearLearnings(context.Context, *common.Void) (*common.Void, error)
	// given a userid for a service, this will return the rpcinterceptor service
	GetServiceByUserID(context.Context, *ServiceByUserIDRequest) (*Service, error)
}

func RegisterRPCInterceptorServiceServer(s *grpc.Server, srv RPCInterceptorServiceServer) {
	s.RegisterService(&_RPCInterceptorService_serviceDesc, srv)
}

func _RPCInterceptorService_InterceptRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InterceptRPCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCInterceptorServiceServer).InterceptRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcinterceptor.RPCInterceptorService/InterceptRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCInterceptorServiceServer).InterceptRPC(ctx, req.(*InterceptRPCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCInterceptorService_LogError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogErrorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCInterceptorServiceServer).LogError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcinterceptor.RPCInterceptorService/LogError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCInterceptorServiceServer).LogError(ctx, req.(*LogErrorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCInterceptorService_GetMyServiceID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCInterceptorServiceServer).GetMyServiceID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcinterceptor.RPCInterceptorService/GetMyServiceID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCInterceptorServiceServer).GetMyServiceID(ctx, req.(*ServiceIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCInterceptorService_GetLearnings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCInterceptorServiceServer).GetLearnings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcinterceptor.RPCInterceptorService/GetLearnings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCInterceptorServiceServer).GetLearnings(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCInterceptorService_ClearLearnings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCInterceptorServiceServer).ClearLearnings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcinterceptor.RPCInterceptorService/ClearLearnings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCInterceptorServiceServer).ClearLearnings(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCInterceptorService_GetServiceByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceByUserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCInterceptorServiceServer).GetServiceByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcinterceptor.RPCInterceptorService/GetServiceByUserID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCInterceptorServiceServer).GetServiceByUserID(ctx, req.(*ServiceByUserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RPCInterceptorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpcinterceptor.RPCInterceptorService",
	HandlerType: (*RPCInterceptorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InterceptRPC",
			Handler:    _RPCInterceptorService_InterceptRPC_Handler,
		},
		{
			MethodName: "LogError",
			Handler:    _RPCInterceptorService_LogError_Handler,
		},
		{
			MethodName: "GetMyServiceID",
			Handler:    _RPCInterceptorService_GetMyServiceID_Handler,
		},
		{
			MethodName: "GetLearnings",
			Handler:    _RPCInterceptorService_GetLearnings_Handler,
		},
		{
			MethodName: "ClearLearnings",
			Handler:    _RPCInterceptorService_ClearLearnings_Handler,
		},
		{
			MethodName: "GetServiceByUserID",
			Handler:    _RPCInterceptorService_GetServiceByUserID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/golang.conradwood.net/apis/rpcinterceptor/rpcinterceptor.proto",
}

func init() {
	proto.RegisterFile("protos/golang.conradwood.net/apis/rpcinterceptor/rpcinterceptor.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 1109 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x56, 0xdd, 0x6e, 0xdb, 0x46,
	0x13, 0xfd, 0x28, 0xc9, 0x8a, 0x34, 0xfa, 0xb1, 0xb2, 0xce, 0xe7, 0xb2, 0x82, 0xe1, 0xaa, 0xac,
	0x11, 0xa8, 0x6e, 0x21, 0xbb, 0x2a, 0x90, 0x34, 0x41, 0x2e, 0x0c, 0x4b, 0xb6, 0x21, 0xc0, 0x4a,
	0x8c, 0xb5, 0x12, 0xb4, 0x97, 0x6b, 0x6a, 0xc3, 0xb0, 0x96, 0x77, 0xd5, 0x25, 0x95, 0x42, 0x77,
	0x05, 0x7a, 0xdb, 0x37, 0xe8, 0xdb, 0xf4, 0x05, 0x0a, 0xf4, 0x1d, 0xfa, 0x1e, 0xc5, 0xfe, 0x50,
	0x5c, 0x52, 0x96, 0x7b, 0xd9, 0x1b, 0x89, 0x73, 0x78, 0x66, 0x76, 0x76, 0xe6, 0xcc, 0x72, 0xe1,
	0x6c, 0x2e, 0x78, 0xcc, 0xa3, 0xa3, 0x80, 0xcf, 0x08, 0x0b, 0x7a, 0x3e, 0x67, 0x82, 0x4c, 0x7f,
	0xe6, 0x7c, 0xda, 0x63, 0x34, 0x3e, 0x22, 0xf3, 0x30, 0x3a, 0x12, 0x73, 0x3f, 0x64, 0x31, 0x15,
	0x3e, 0x9d, 0xc7, 0x5c, 0xe4, 0xcc, 0x9e, 0xf2, 0x47, 0xcd, 0x2c, 0xda, 0x3e, 0x7c, 0x20, 0x1e,
	0x59, 0xc4, 0x1f, 0xd4, 0x8f, 0xf6, 0x6d, 0xf7, 0x1e, 0xe0, 0xfa, 0xfc, 0xee, 0x8e, 0x33, 0xf3,
	0xa7, 0xf9, 0xde, 0x5f, 0x25, 0x78, 0x32, 0x4a, 0xd6, 0xc2, 0x57, 0x03, 0x4c, 0xa3, 0x39, 0x67,
	0x11, 0x45, 0x7b, 0x50, 0xc5, 0xf4, 0xa7, 0x05, 0x8d, 0xe2, 0xd1, 0xd0, 0x75, 0x3a, 0x4e, 0xb7,
	0x8a, 0x53, 0x00, 0x1d, 0x43, 0x63, 0x40, 0x66, 0x33, 0x2a, 0xae, 0xa9, 0xf8, 0x18, 0xfa, 0xd4,
	0x2d, 0x74, 0x9c, 0x6e, 0xad, 0x0f, 0x3d, 0x95, 0xca, 0xdb, 0x88, 0x0a, 0x9c, 0x25, 0xa0, 0x43,
	0x00, 0x0d, 0xc8, 0x97, 0x6e, 0x71, 0x8d, 0x6e, 0xbd, 0x45, 0x7d, 0x68, 0x1a, 0xe7, 0xc5, 0x94,
	0x2b, 0x7e, 0x69, 0x8d, 0x9f, 0x63, 0xa0, 0x5d, 0x28, 0x63, 0xfa, 0x23, 0xf5, 0x63, 0xb7, 0xdc,
	0x71, 0xba, 0x15, 0x6c, 0x2c, 0x74, 0x02, 0x75, 0xfd, 0x84, 0x29, 0x89, 0x38, 0x73, 0x1f, 0x75,
	0x9c, 0x6e, 0xb3, 0xbf, 0xd7, 0xcb, 0x55, 0xde, 0xe6, 0xe0, 0x8c, 0x07, 0x7a, 0x9a, 0x64, 0x33,
	0xa6, 0xf1, 0x07, 0x3e, 0x1d, 0x0d, 0xdd, 0x4a, 0xc7, 0xe9, 0x96, 0x70, 0x0e, 0x95, 0x19, 0x5c,
	0xf3, 0x85, 0xf0, 0xa9, 0x5b, 0x55, 0xe5, 0x32, 0x16, 0xea, 0xc2, 0xb6, 0x62, 0x52, 0x53, 0x8a,
	0xd1, 0xd0, 0x05, 0x15, 0x20, 0x0f, 0xa3, 0x53, 0xd8, 0xb9, 0x0e, 0x03, 0x46, 0xa7, 0xd9, 0xda,
	0xd6, 0xd4, 0xe6, 0x5b, 0x7a, 0xf3, 0x9a, 0xa0, 0x4a, 0x70, 0x1f, 0x19, 0xbd, 0x82, 0x96, 0x0d,
	0xab, 0xea, 0xd5, 0x37, 0x04, 0x58, 0x63, 0xa2, 0x21, 0x3c, 0xc9, 0x04, 0x4d, 0xea, 0xdf, 0xd8,
	0x10, 0xe1, 0x5e, 0xb6, 0x77, 0x0c, 0xbb, 0x26, 0x9d, 0xd3, 0xa5, 0x04, 0x46, 0x43, 0xa3, 0x1c,
	0x59, 0x23, 0x0d, 0x18, 0x49, 0x19, 0xcb, 0x3b, 0x83, 0x47, 0xc9, 0x06, 0x9a, 0x50, 0x30, 0xaf,
	0x4b, 0xb8, 0x30, 0x1a, 0x22, 0x04, 0xa5, 0xd7, 0xe4, 0x4e, 0x2b, 0xac, 0x8a, 0xd5, 0xb3, 0x15,
	0xa6, 0x98, 0x09, 0xf3, 0xbb, 0x03, 0x3b, 0x59, 0x35, 0xeb, 0x65, 0x5f, 0x02, 0x8c, 0xd8, 0x98,
	0xc6, 0x64, 0x4a, 0x62, 0xa2, 0x62, 0xd7, 0xfa, 0xed, 0xbc, 0x04, 0x52, 0x06, 0xb6, 0xd8, 0xc8,
	0x5d, 0xa5, 0x66, 0x52, 0x58, 0x65, 0xba, 0x0b, 0x65, 0xdd, 0xfc, 0x24, 0x0b, 0x6d, 0x59, 0x42,
	0x28, 0xd9, 0x42, 0xf0, 0xfe, 0x74, 0xa0, 0x39, 0x98, 0x7c, 0x8f, 0xf9, 0x22, 0x0e, 0x59, 0x30,
	0x21, 0x41, 0x84, 0x5e, 0x41, 0x49, 0xfe, 0xbb, 0x4e, 0xa7, 0xd8, 0xad, 0xf5, 0xbb, 0xf9, 0x94,
	0xb2, 0xec, 0x9e, 0xfc, 0x39, 0x63, 0xb1, 0x58, 0x62, 0xe5, 0x25, 0x95, 0x75, 0x4e, 0x66, 0xb3,
	0x1b, 0xe2, 0xdf, 0x4e, 0xf8, 0xd5, 0x8c, 0x84, 0x4c, 0xa5, 0x58, 0xc1, 0x79, 0x58, 0x4e, 0xf3,
	0x95, 0xe0, 0x73, 0x12, 0x90, 0x98, 0xaa, 0x6c, 0x2b, 0x38, 0x05, 0xda, 0xcf, 0xa1, 0xba, 0x0a,
	0x8d, 0x5a, 0x50, 0xbc, 0xa5, 0x4b, 0xd3, 0x1f, 0xf9, 0x88, 0x9e, 0xc0, 0xd6, 0x47, 0x32, 0x5b,
	0x24, 0xfb, 0xd7, 0xc6, 0xcb, 0xc2, 0x77, 0x8e, 0xf7, 0xeb, 0x96, 0x5d, 0xd8, 0x7f, 0x39, 0x33,
	0x76, 0xa1, 0x7c, 0xce, 0xf9, 0x29, 0x11, 0x26, 0x8e, 0xb1, 0xa4, 0x97, 0x6c, 0xdf, 0x84, 0xdf,
	0x52, 0x66, 0x2a, 0x99, 0x02, 0xc8, 0x83, 0xba, 0xa9, 0xb7, 0x26, 0xe8, 0x92, 0x66, 0x30, 0x4b,
	0x0e, 0x5b, 0xb6, 0x1c, 0xee, 0x99, 0xdc, 0x47, 0xf7, 0x4e, 0x6e, 0x32, 0xa1, 0x22, 0x9d, 0xd0,
	0x8a, 0x35, 0xa1, 0x29, 0x8c, 0x3a, 0x50, 0x33, 0x0d, 0x19, 0xb1, 0xf7, 0x5c, 0x0d, 0x7a, 0x03,
	0xdb, 0x90, 0x94, 0xcb, 0x3b, 0x2a, 0xa2, 0x90, 0x33, 0x35, 0xe5, 0x0d, 0x9c, 0x98, 0xe8, 0x20,
	0x15, 0x52, 0x6d, 0xed, 0x38, 0x5b, 0x89, 0x6a, 0x1f, 0x4a, 0xd6, 0xcc, 0xda, 0x14, 0x85, 0xa3,
	0x67, 0xd0, 0xd0, 0x33, 0x97, 0xc4, 0xda, 0x34, 0x9a, 0x59, 0x1a, 0x3a, 0x06, 0x48, 0x5f, 0xba,
	0xcd, 0x0d, 0x4e, 0x16, 0x47, 0xb6, 0x7d, 0x22, 0x88, 0x4f, 0xdd, 0x6d, 0xa5, 0x17, 0x6d, 0x48,
	0x74, 0x48, 0x6f, 0x16, 0x81, 0xdb, 0xd2, 0xa8, 0x32, 0xd0, 0x8b, 0x34, 0xab, 0x48, 0xed, 0xfd,
	0xb1, 0x5a, 0x60, 0xc7, 0x5e, 0xc0, 0xbc, 0xc2, 0x59, 0x26, 0x3a, 0x59, 0x95, 0x54, 0x4d, 0x02,
	0x52, 0x8e, 0xfb, 0x0f, 0x4f, 0x02, 0xb6, 0x5d, 0xbc, 0xbf, 0x1d, 0xd8, 0xbe, 0xe4, 0xc1, 0x99,
	0x10, 0x5c, 0xfc, 0x37, 0x13, 0xbf, 0x07, 0x55, 0xb5, 0xfa, 0x80, 0x4f, 0xf5, 0xd0, 0x37, 0x70,
	0x0a, 0x48, 0x19, 0x0e, 0xc3, 0x68, 0x3e, 0x23, 0xcb, 0x31, 0x8d, 0x22, 0x12, 0x50, 0x23, 0xd3,
	0x1c, 0x8a, 0xf6, 0x01, 0x2e, 0x79, 0x90, 0x70, 0xca, 0x8a, 0x63, 0x21, 0xde, 0x09, 0xb4, 0x56,
	0x4a, 0x4c, 0xf6, 0x29, 0x9b, 0xa4, 0xe6, 0x42, 0x8f, 0xdb, 0xd6, 0x6a, 0x20, 0xc6, 0x4b, 0xeb,
	0xd4, 0x34, 0x96, 0xf7, 0x0d, 0x3c, 0xb6, 0x22, 0xa4, 0x5f, 0xfa, 0x54, 0xf7, 0x05, 0xa5, 0xfb,
	0x14, 0xf0, 0x7e, 0x71, 0xa0, 0x72, 0x49, 0x89, 0x60, 0x21, 0x0b, 0xa4, 0xfc, 0x27, 0x3c, 0x25,
	0xeb, 0x43, 0xda, 0x86, 0xd0, 0x01, 0x34, 0xce, 0x05, 0xbf, 0xcb, 0x07, 0xcc, 0x82, 0x9b, 0xce,
	0x6f, 0xb9, 0x9b, 0x01, 0x5f, 0xb0, 0x58, 0xd5, 0xb0, 0x84, 0xb5, 0xe1, 0x0d, 0xa0, 0x9a, 0x64,
	0x10, 0xa1, 0x67, 0x96, 0x61, 0x8e, 0x4d, 0x37, 0xdf, 0xd7, 0x84, 0x80, 0x53, 0xea, 0xe1, 0x6f,
	0x4e, 0xf6, 0x22, 0x80, 0xb6, 0xa1, 0xf6, 0x9a, 0xb3, 0xeb, 0x39, 0xf5, 0xc3, 0xf7, 0xa1, 0xdf,
	0xfa, 0x1f, 0x6a, 0x41, 0x5d, 0xcd, 0x80, 0x22, 0xd1, 0x69, 0xcb, 0x41, 0x3b, 0xb0, 0x6d, 0x72,
	0x5e, 0x81, 0x05, 0xe9, 0xf7, 0x46, 0x04, 0x2b, 0xa0, 0x28, 0x01, 0xe9, 0x37, 0x0e, 0xa3, 0x28,
	0x64, 0x41, 0xab, 0x84, 0x10, 0x34, 0x8d, 0x5b, 0x82, 0x6d, 0xa1, 0x26, 0xc0, 0x1b, 0x11, 0x24,
	0x76, 0xb9, 0xff, 0x47, 0x11, 0xfe, 0x8f, 0xaf, 0x06, 0xa3, 0x34, 0xeb, 0x44, 0x63, 0x3f, 0x40,
	0xdd, 0xfe, 0x84, 0xa1, 0x2f, 0xd6, 0x55, 0xbb, 0xf6, 0x81, 0x6b, 0x1f, 0x3c, 0x4c, 0x32, 0x9d,
	0x7e, 0x01, 0x95, 0x64, 0x4e, 0xd0, 0x67, 0x6b, 0x45, 0xcb, 0x4e, 0x50, 0xbb, 0xde, 0x33, 0x17,
	0xc5, 0x77, 0x3c, 0x9c, 0xa2, 0xb7, 0xd0, 0xbc, 0xa0, 0xf1, 0x78, 0x69, 0x1d, 0x85, 0xf9, 0x00,
	0x79, 0x6d, 0xb6, 0x3f, 0x7f, 0x80, 0x61, 0x32, 0x7a, 0x0e, 0xf5, 0x0b, 0x1a, 0xa7, 0xdd, 0xcd,
	0x2c, 0xda, 0xfe, 0x74, 0x53, 0x63, 0x23, 0xf4, 0x35, 0x34, 0x07, 0x33, 0x4a, 0xc4, 0x26, 0xd7,
	0x7c, 0xf6, 0xe8, 0x82, 0xc6, 0xb9, 0x3b, 0x09, 0x7a, 0xba, 0x21, 0xbf, 0xdc, 0xa5, 0xa5, 0xfd,
	0xc9, 0x26, 0xde, 0x57, 0xf0, 0x25, 0xa3, 0xb1, 0x7d, 0xd7, 0x36, 0xb7, 0x6f, 0x79, 0xdd, 0xce,
	0x39, 0xdd, 0x94, 0xd5, 0x85, 0xfb, 0xdb, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x84, 0x2b,
	0xae, 0x25, 0x0c, 0x00, 0x00,
}
