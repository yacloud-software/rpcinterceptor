// Code generated by protoc-gen-go.
// source: golang.conradwood.net/apis/usagestats/usagestats.proto
// DO NOT EDIT!

/*
Package usagestats is a generated protocol buffer package.

It is generated from these files:
	golang.conradwood.net/apis/usagestats/usagestats.proto

It has these top-level messages:
	LogHttpCallRequest
	LogGrpcCallRequest
	QueryCallsRequest
	QueryCallsByUserResponse
	QueryCallsByGroupResponse
	QueryCallsByResultResponse
*/
package usagestats

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
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

// type of call
type CallType int32

const (
	CallType_Unknown CallType = 0
	CallType_Http    CallType = 1
	CallType_Grpc    CallType = 2
)

var CallType_name = map[int32]string{
	0: "Unknown",
	1: "Http",
	2: "Grpc",
}
var CallType_value = map[string]int32{
	"Unknown": 0,
	"Http":    1,
	"Grpc":    2,
}

func (x CallType) String() string {
	return proto.EnumName(CallType_name, int32(x))
}
func (CallType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// log an http call
type LogHttpCallRequest struct {
	Url       string `protobuf:"bytes,1,opt,name=Url" json:"Url,omitempty"`
	Success   bool   `protobuf:"varint,2,opt,name=Success" json:"Success,omitempty"`
	Timestamp uint32 `protobuf:"varint,3,opt,name=Timestamp" json:"Timestamp,omitempty"`
}

func (m *LogHttpCallRequest) Reset()                    { *m = LogHttpCallRequest{} }
func (m *LogHttpCallRequest) String() string            { return proto.CompactTextString(m) }
func (*LogHttpCallRequest) ProtoMessage()               {}
func (*LogHttpCallRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LogHttpCallRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *LogHttpCallRequest) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *LogHttpCallRequest) GetTimestamp() uint32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

// log a grpc call
type LogGrpcCallRequest struct {
	Module    string `protobuf:"bytes,1,opt,name=Module" json:"Module,omitempty"`
	Method    string `protobuf:"bytes,2,opt,name=Method" json:"Method,omitempty"`
	Success   bool   `protobuf:"varint,3,opt,name=Success" json:"Success,omitempty"`
	Timestamp uint32 `protobuf:"varint,4,opt,name=Timestamp" json:"Timestamp,omitempty"`
}

func (m *LogGrpcCallRequest) Reset()                    { *m = LogGrpcCallRequest{} }
func (m *LogGrpcCallRequest) String() string            { return proto.CompactTextString(m) }
func (*LogGrpcCallRequest) ProtoMessage()               {}
func (*LogGrpcCallRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LogGrpcCallRequest) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
}

func (m *LogGrpcCallRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *LogGrpcCallRequest) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *LogGrpcCallRequest) GetTimestamp() uint32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

// if day, week or month is 0, the data returned
// will cover the week, month or year
type QueryCallsRequest struct {
	Month uint32 `protobuf:"varint,1,opt,name=Month" json:"Month,omitempty"`
	Year  uint32 `protobuf:"varint,2,opt,name=Year" json:"Year,omitempty"`
}

func (m *QueryCallsRequest) Reset()                    { *m = QueryCallsRequest{} }
func (m *QueryCallsRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryCallsRequest) ProtoMessage()               {}
func (*QueryCallsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *QueryCallsRequest) GetMonth() uint32 {
	if m != nil {
		return m.Month
	}
	return 0
}

func (m *QueryCallsRequest) GetYear() uint32 {
	if m != nil {
		return m.Year
	}
	return 0
}

type QueryCallsByUserResponse struct {
	CallRecords []*QueryCallsByUserResponse_CallRecord `protobuf:"bytes,1,rep,name=CallRecords" json:"CallRecords,omitempty"`
}

func (m *QueryCallsByUserResponse) Reset()                    { *m = QueryCallsByUserResponse{} }
func (m *QueryCallsByUserResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryCallsByUserResponse) ProtoMessage()               {}
func (*QueryCallsByUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *QueryCallsByUserResponse) GetCallRecords() []*QueryCallsByUserResponse_CallRecord {
	if m != nil {
		return m.CallRecords
	}
	return nil
}

type QueryCallsByUserResponse_CallRecord struct {
	Abbrev    string   `protobuf:"bytes,1,opt,name=Abbrev" json:"Abbrev,omitempty"`
	FirstName string   `protobuf:"bytes,2,opt,name=FirstName" json:"FirstName,omitempty"`
	LastName  string   `protobuf:"bytes,3,opt,name=LastName" json:"LastName,omitempty"`
	CallPath  []string `protobuf:"bytes,4,rep,name=CallPath" json:"CallPath,omitempty"`
	CallType  CallType `protobuf:"varint,5,opt,name=CallType,enum=usagestats.CallType" json:"CallType,omitempty"`
}

func (m *QueryCallsByUserResponse_CallRecord) Reset()         { *m = QueryCallsByUserResponse_CallRecord{} }
func (m *QueryCallsByUserResponse_CallRecord) String() string { return proto.CompactTextString(m) }
func (*QueryCallsByUserResponse_CallRecord) ProtoMessage()    {}
func (*QueryCallsByUserResponse_CallRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

func (m *QueryCallsByUserResponse_CallRecord) GetAbbrev() string {
	if m != nil {
		return m.Abbrev
	}
	return ""
}

func (m *QueryCallsByUserResponse_CallRecord) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *QueryCallsByUserResponse_CallRecord) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *QueryCallsByUserResponse_CallRecord) GetCallPath() []string {
	if m != nil {
		return m.CallPath
	}
	return nil
}

func (m *QueryCallsByUserResponse_CallRecord) GetCallType() CallType {
	if m != nil {
		return m.CallType
	}
	return CallType_Unknown
}

type QueryCallsByGroupResponse struct {
	CallRecords []*QueryCallsByGroupResponse_CallRecord `protobuf:"bytes,1,rep,name=CallRecords" json:"CallRecords,omitempty"`
}

func (m *QueryCallsByGroupResponse) Reset()                    { *m = QueryCallsByGroupResponse{} }
func (m *QueryCallsByGroupResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryCallsByGroupResponse) ProtoMessage()               {}
func (*QueryCallsByGroupResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *QueryCallsByGroupResponse) GetCallRecords() []*QueryCallsByGroupResponse_CallRecord {
	if m != nil {
		return m.CallRecords
	}
	return nil
}

type QueryCallsByGroupResponse_CallRecord struct {
	GroupName string   `protobuf:"bytes,1,opt,name=GroupName" json:"GroupName,omitempty"`
	CallPath  []string `protobuf:"bytes,2,rep,name=CallPath" json:"CallPath,omitempty"`
	CallType  CallType `protobuf:"varint,3,opt,name=CallType,enum=usagestats.CallType" json:"CallType,omitempty"`
}

func (m *QueryCallsByGroupResponse_CallRecord) Reset()         { *m = QueryCallsByGroupResponse_CallRecord{} }
func (m *QueryCallsByGroupResponse_CallRecord) String() string { return proto.CompactTextString(m) }
func (*QueryCallsByGroupResponse_CallRecord) ProtoMessage()    {}
func (*QueryCallsByGroupResponse_CallRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4, 0}
}

func (m *QueryCallsByGroupResponse_CallRecord) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *QueryCallsByGroupResponse_CallRecord) GetCallPath() []string {
	if m != nil {
		return m.CallPath
	}
	return nil
}

func (m *QueryCallsByGroupResponse_CallRecord) GetCallType() CallType {
	if m != nil {
		return m.CallType
	}
	return CallType_Unknown
}

type QueryCallsByResultResponse struct {
	CallRecords []*QueryCallsByResultResponse_CallRecord `protobuf:"bytes,1,rep,name=CallRecords" json:"CallRecords,omitempty"`
}

func (m *QueryCallsByResultResponse) Reset()                    { *m = QueryCallsByResultResponse{} }
func (m *QueryCallsByResultResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryCallsByResultResponse) ProtoMessage()               {}
func (*QueryCallsByResultResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *QueryCallsByResultResponse) GetCallRecords() []*QueryCallsByResultResponse_CallRecord {
	if m != nil {
		return m.CallRecords
	}
	return nil
}

type QueryCallsByResultResponse_CallRecord struct {
	Day             uint32 `protobuf:"varint,1,opt,name=Day" json:"Day,omitempty"`
	GoodAccessCount uint32 `protobuf:"varint,2,opt,name=GoodAccessCount" json:"GoodAccessCount,omitempty"`
	BadAccessCount  uint32 `protobuf:"varint,3,opt,name=BadAccessCount" json:"BadAccessCount,omitempty"`
}

func (m *QueryCallsByResultResponse_CallRecord) Reset()         { *m = QueryCallsByResultResponse_CallRecord{} }
func (m *QueryCallsByResultResponse_CallRecord) String() string { return proto.CompactTextString(m) }
func (*QueryCallsByResultResponse_CallRecord) ProtoMessage()    {}
func (*QueryCallsByResultResponse_CallRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5, 0}
}

func (m *QueryCallsByResultResponse_CallRecord) GetDay() uint32 {
	if m != nil {
		return m.Day
	}
	return 0
}

func (m *QueryCallsByResultResponse_CallRecord) GetGoodAccessCount() uint32 {
	if m != nil {
		return m.GoodAccessCount
	}
	return 0
}

func (m *QueryCallsByResultResponse_CallRecord) GetBadAccessCount() uint32 {
	if m != nil {
		return m.BadAccessCount
	}
	return 0
}

func init() {
	proto.RegisterType((*LogHttpCallRequest)(nil), "usagestats.LogHttpCallRequest")
	proto.RegisterType((*LogGrpcCallRequest)(nil), "usagestats.LogGrpcCallRequest")
	proto.RegisterType((*QueryCallsRequest)(nil), "usagestats.QueryCallsRequest")
	proto.RegisterType((*QueryCallsByUserResponse)(nil), "usagestats.QueryCallsByUserResponse")
	proto.RegisterType((*QueryCallsByUserResponse_CallRecord)(nil), "usagestats.QueryCallsByUserResponse.CallRecord")
	proto.RegisterType((*QueryCallsByGroupResponse)(nil), "usagestats.QueryCallsByGroupResponse")
	proto.RegisterType((*QueryCallsByGroupResponse_CallRecord)(nil), "usagestats.QueryCallsByGroupResponse.CallRecord")
	proto.RegisterType((*QueryCallsByResultResponse)(nil), "usagestats.QueryCallsByResultResponse")
	proto.RegisterType((*QueryCallsByResultResponse_CallRecord)(nil), "usagestats.QueryCallsByResultResponse.CallRecord")
	proto.RegisterEnum("usagestats.CallType", CallType_name, CallType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UsageStatsService service

type UsageStatsServiceClient interface {
	// logging calls
	LogHttpCall(ctx context.Context, in *LogHttpCallRequest, opts ...grpc.CallOption) (*common.Void, error)
	LogGrpcCall(ctx context.Context, in *LogGrpcCallRequest, opts ...grpc.CallOption) (*common.Void, error)
	// query calls
	QueryCallsByUser(ctx context.Context, in *QueryCallsRequest, opts ...grpc.CallOption) (*QueryCallsByUserResponse, error)
	QueryCallsByGroup(ctx context.Context, in *QueryCallsRequest, opts ...grpc.CallOption) (*QueryCallsByGroupResponse, error)
	QueryCallsByResult(ctx context.Context, in *QueryCallsRequest, opts ...grpc.CallOption) (*QueryCallsByResultResponse, error)
}

type usageStatsServiceClient struct {
	cc *grpc.ClientConn
}

func NewUsageStatsServiceClient(cc *grpc.ClientConn) UsageStatsServiceClient {
	return &usageStatsServiceClient{cc}
}

func (c *usageStatsServiceClient) LogHttpCall(ctx context.Context, in *LogHttpCallRequest, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/usagestats.UsageStatsService/LogHttpCall", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usageStatsServiceClient) LogGrpcCall(ctx context.Context, in *LogGrpcCallRequest, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/usagestats.UsageStatsService/LogGrpcCall", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usageStatsServiceClient) QueryCallsByUser(ctx context.Context, in *QueryCallsRequest, opts ...grpc.CallOption) (*QueryCallsByUserResponse, error) {
	out := new(QueryCallsByUserResponse)
	err := grpc.Invoke(ctx, "/usagestats.UsageStatsService/QueryCallsByUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usageStatsServiceClient) QueryCallsByGroup(ctx context.Context, in *QueryCallsRequest, opts ...grpc.CallOption) (*QueryCallsByGroupResponse, error) {
	out := new(QueryCallsByGroupResponse)
	err := grpc.Invoke(ctx, "/usagestats.UsageStatsService/QueryCallsByGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usageStatsServiceClient) QueryCallsByResult(ctx context.Context, in *QueryCallsRequest, opts ...grpc.CallOption) (*QueryCallsByResultResponse, error) {
	out := new(QueryCallsByResultResponse)
	err := grpc.Invoke(ctx, "/usagestats.UsageStatsService/QueryCallsByResult", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UsageStatsService service

type UsageStatsServiceServer interface {
	// logging calls
	LogHttpCall(context.Context, *LogHttpCallRequest) (*common.Void, error)
	LogGrpcCall(context.Context, *LogGrpcCallRequest) (*common.Void, error)
	// query calls
	QueryCallsByUser(context.Context, *QueryCallsRequest) (*QueryCallsByUserResponse, error)
	QueryCallsByGroup(context.Context, *QueryCallsRequest) (*QueryCallsByGroupResponse, error)
	QueryCallsByResult(context.Context, *QueryCallsRequest) (*QueryCallsByResultResponse, error)
}

func RegisterUsageStatsServiceServer(s *grpc.Server, srv UsageStatsServiceServer) {
	s.RegisterService(&_UsageStatsService_serviceDesc, srv)
}

func _UsageStatsService_LogHttpCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogHttpCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsageStatsServiceServer).LogHttpCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usagestats.UsageStatsService/LogHttpCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsageStatsServiceServer).LogHttpCall(ctx, req.(*LogHttpCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsageStatsService_LogGrpcCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogGrpcCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsageStatsServiceServer).LogGrpcCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usagestats.UsageStatsService/LogGrpcCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsageStatsServiceServer).LogGrpcCall(ctx, req.(*LogGrpcCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsageStatsService_QueryCallsByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCallsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsageStatsServiceServer).QueryCallsByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usagestats.UsageStatsService/QueryCallsByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsageStatsServiceServer).QueryCallsByUser(ctx, req.(*QueryCallsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsageStatsService_QueryCallsByGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCallsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsageStatsServiceServer).QueryCallsByGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usagestats.UsageStatsService/QueryCallsByGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsageStatsServiceServer).QueryCallsByGroup(ctx, req.(*QueryCallsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsageStatsService_QueryCallsByResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCallsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsageStatsServiceServer).QueryCallsByResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usagestats.UsageStatsService/QueryCallsByResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsageStatsServiceServer).QueryCallsByResult(ctx, req.(*QueryCallsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UsageStatsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "usagestats.UsageStatsService",
	HandlerType: (*UsageStatsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LogHttpCall",
			Handler:    _UsageStatsService_LogHttpCall_Handler,
		},
		{
			MethodName: "LogGrpcCall",
			Handler:    _UsageStatsService_LogGrpcCall_Handler,
		},
		{
			MethodName: "QueryCallsByUser",
			Handler:    _UsageStatsService_QueryCallsByUser_Handler,
		},
		{
			MethodName: "QueryCallsByGroup",
			Handler:    _UsageStatsService_QueryCallsByGroup_Handler,
		},
		{
			MethodName: "QueryCallsByResult",
			Handler:    _UsageStatsService_QueryCallsByResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golang.conradwood.net/apis/usagestats/usagestats.proto",
}

func init() {
	proto.RegisterFile("golang.conradwood.net/apis/usagestats/usagestats.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 626 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x55, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0xfd, 0x1c, 0xa7, 0x7f, 0x93, 0xaf, 0xc5, 0x5d, 0x55, 0xc8, 0x44, 0x05, 0x59, 0x11, 0x2d,
	0x16, 0x48, 0x6e, 0x29, 0x12, 0x37, 0x88, 0x8b, 0xb6, 0x88, 0x72, 0xd1, 0x22, 0xea, 0x34, 0xa0,
	0x0a, 0x09, 0x69, 0x6b, 0xaf, 0xd2, 0x08, 0xc7, 0x6b, 0x76, 0xd7, 0x2d, 0x91, 0x78, 0x14, 0xc4,
	0x13, 0xf0, 0x56, 0x5c, 0xf1, 0x16, 0x68, 0xd7, 0x76, 0xbc, 0x76, 0x7e, 0x9a, 0xab, 0xec, 0xcc,
	0x78, 0xe6, 0xe4, 0x9c, 0x9d, 0x99, 0x85, 0x97, 0x7d, 0x1a, 0xe1, 0xb8, 0xef, 0x05, 0x34, 0x66,
	0x38, 0xbc, 0xa5, 0x34, 0xf4, 0x62, 0x22, 0xf6, 0x70, 0x32, 0xe0, 0x7b, 0x29, 0xc7, 0x7d, 0xc2,
	0x05, 0x16, 0xfa, 0xd1, 0x4b, 0x18, 0x15, 0x14, 0x41, 0xe9, 0x69, 0x7b, 0x73, 0x6a, 0x04, 0x74,
	0x38, 0xa4, 0x71, 0xfe, 0x93, 0xe5, 0x76, 0xbe, 0x00, 0x3a, 0xa5, 0xfd, 0x77, 0x42, 0x24, 0xc7,
	0x38, 0x8a, 0x7c, 0xf2, 0x2d, 0x25, 0x5c, 0x20, 0x0b, 0xcc, 0x1e, 0x8b, 0x6c, 0xc3, 0x31, 0xdc,
	0x35, 0x5f, 0x1e, 0x91, 0x0d, 0x2b, 0xdd, 0x34, 0x08, 0x08, 0xe7, 0x76, 0xc3, 0x31, 0xdc, 0x55,
	0xbf, 0x30, 0xd1, 0x36, 0xac, 0x5d, 0x0c, 0x86, 0x12, 0x7e, 0x98, 0xd8, 0xa6, 0x63, 0xb8, 0xeb,
	0x7e, 0xe9, 0xe8, 0xfc, 0x50, 0xf5, 0x4f, 0x58, 0x12, 0xe8, 0xf5, 0xef, 0xc3, 0xf2, 0x19, 0x0d,
	0xd3, 0x88, 0xe4, 0x10, 0xb9, 0xa5, 0xfc, 0x44, 0x5c, 0xd3, 0x50, 0x81, 0x48, 0xbf, 0xb2, 0x74,
	0x74, 0x73, 0x0e, 0x7a, 0xb3, 0x8e, 0xfe, 0x1a, 0x36, 0xcf, 0x53, 0xc2, 0x46, 0x12, 0x9b, 0x17,
	0xe0, 0x5b, 0xb0, 0x74, 0x46, 0x63, 0x71, 0xad, 0xb0, 0xd7, 0xfd, 0xcc, 0x40, 0x08, 0x9a, 0x97,
	0x04, 0x33, 0x05, 0xbc, 0xee, 0xab, 0x73, 0xe7, 0x67, 0x03, 0xec, 0x32, 0xff, 0x68, 0xd4, 0xe3,
	0x84, 0xf9, 0x84, 0x27, 0x34, 0xe6, 0x04, 0x9d, 0x43, 0x2b, 0xa3, 0x14, 0x50, 0x16, 0x72, 0xdb,
	0x70, 0x4c, 0xb7, 0x75, 0xb0, 0xe7, 0x69, 0xb7, 0x33, 0x2b, 0xd5, 0x2b, 0xf3, 0x7c, 0xbd, 0x46,
	0xfb, 0xb7, 0x01, 0x50, 0xda, 0x52, 0x8d, 0xc3, 0xab, 0x2b, 0x46, 0x6e, 0x0a, 0x95, 0x32, 0x4b,
	0x72, 0x7e, 0x3b, 0x60, 0x5c, 0xbc, 0xc7, 0x43, 0x92, 0x0b, 0x55, 0x3a, 0x50, 0x1b, 0x56, 0x4f,
	0x71, 0x1e, 0x34, 0x55, 0x70, 0x6c, 0xcb, 0x98, 0xac, 0xff, 0x01, 0x8b, 0x6b, 0xbb, 0xe9, 0x98,
	0x32, 0x56, 0xd8, 0x68, 0x3f, 0x8b, 0x5d, 0x8c, 0x12, 0x62, 0x2f, 0x39, 0x86, 0xbb, 0x71, 0xb0,
	0xa5, 0x93, 0x29, 0x62, 0xfe, 0xf8, 0xab, 0xce, 0x5f, 0x03, 0x1e, 0xe8, 0x1c, 0x4f, 0x18, 0x4d,
	0x93, 0xb1, 0x3e, 0xfe, 0x34, 0x7d, 0xf6, 0x67, 0xe9, 0x53, 0xc9, 0x9d, 0x29, 0xd0, 0xf7, 0x8a,
	0x3e, 0xdb, 0xb0, 0xa6, 0xd2, 0x14, 0xd5, 0x4c, 0xa2, 0xd2, 0x51, 0xe1, 0xda, 0x98, 0xc3, 0xd5,
	0x5c, 0x88, 0xeb, 0x1f, 0x03, 0xda, 0xfa, 0xff, 0xf5, 0x09, 0x4f, 0x23, 0x31, 0x26, 0xdb, 0x9d,
	0x46, 0xf6, 0xf9, 0x2c, 0xb2, 0xd5, 0xe4, 0x99, 0x6c, 0x93, 0x0a, 0x5b, 0x0b, 0xcc, 0x37, 0x78,
	0x94, 0x37, 0xad, 0x3c, 0x22, 0x17, 0xee, 0x9d, 0x50, 0x1a, 0x1e, 0xaa, 0x49, 0x38, 0xa6, 0x69,
	0x2c, 0xf2, 0xee, 0xad, 0xbb, 0xd1, 0x2e, 0x6c, 0x1c, 0xe1, 0xca, 0x87, 0xd9, 0xa0, 0xd6, 0xbc,
	0x4f, 0x9f, 0x95, 0xba, 0xa0, 0x16, 0xac, 0xf4, 0xe2, 0xaf, 0x31, 0xbd, 0x8d, 0xad, 0xff, 0xd0,
	0x2a, 0x34, 0xe5, 0x8e, 0xb0, 0x0c, 0x79, 0x92, 0xd3, 0x6c, 0x35, 0x0e, 0x7e, 0x99, 0xb0, 0xd9,
	0x93, 0x04, 0xbb, 0x92, 0x60, 0x97, 0xb0, 0x9b, 0x41, 0x40, 0xd0, 0x2b, 0x68, 0x69, 0x0b, 0x05,
	0x3d, 0xd2, 0x35, 0x98, 0xdc, 0x34, 0xed, 0xff, 0xbd, 0x7c, 0x1d, 0x7d, 0xa4, 0x83, 0x30, 0x4f,
	0x2e, 0xb6, 0xc5, 0x44, 0x72, 0x6d, 0x8d, 0xd4, 0x92, 0x3f, 0x81, 0x55, 0x9f, 0x38, 0xf4, 0x70,
	0xfa, 0x15, 0x14, 0x05, 0x1e, 0x2f, 0x32, 0xae, 0xe8, 0x52, 0xdf, 0x22, 0x79, 0xab, 0xde, 0x55,
	0x79, 0x67, 0xa1, 0x46, 0x47, 0x9f, 0x01, 0x4d, 0x36, 0xc6, 0x5d, 0xb5, 0x77, 0x17, 0xeb, 0xab,
	0xa3, 0x27, 0xb0, 0x13, 0x13, 0xa1, 0x3f, 0x05, 0xf9, 0xe3, 0x20, 0x5f, 0x03, 0xad, 0xc6, 0xd5,
	0xb2, 0x7a, 0x0b, 0x5e, 0xfc, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x3b, 0x19, 0x3c, 0x81, 0x06,
	0x00, 0x00,
}