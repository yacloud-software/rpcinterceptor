// Code generated by protoc-gen-go.
// source: golang.conradwood.net/apis/getestservice/echoservice.proto
// DO NOT EDIT!

/*
Package getestservice is a generated protocol buffer package.

It is generated from these files:
	golang.conradwood.net/apis/getestservice/echoservice.proto

It has these top-level messages:
	PingRequest
	PingResponse
*/
package getestservice

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type PingRequest struct {
	SequenceNumber uint32 `protobuf:"varint,1,opt,name=SequenceNumber" json:"SequenceNumber,omitempty"`
	Payload        string `protobuf:"bytes,2,opt,name=Payload" json:"Payload,omitempty"`
	TTL            uint32 `protobuf:"varint,3,opt,name=TTL" json:"TTL,omitempty"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PingRequest) GetSequenceNumber() uint32 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func (m *PingRequest) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *PingRequest) GetTTL() uint32 {
	if m != nil {
		return m.TTL
	}
	return 0
}

type PingResponse struct {
	Response *PingRequest `protobuf:"bytes,1,opt,name=Response" json:"Response,omitempty"`
}

func (m *PingResponse) Reset()                    { *m = PingResponse{} }
func (m *PingResponse) String() string            { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()               {}
func (*PingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PingResponse) GetResponse() *PingRequest {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "getestservice.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "getestservice.PingResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EchoService service

type EchoServiceClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type echoServiceClient struct {
	cc *grpc.ClientConn
}

func NewEchoServiceClient(cc *grpc.ClientConn) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/getestservice.EchoService/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EchoService service

type EchoServiceServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
}

func RegisterEchoServiceServer(s *grpc.Server, srv EchoServiceServer) {
	s.RegisterService(&_EchoService_serviceDesc, srv)
}

func _EchoService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/getestservice.EchoService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EchoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "getestservice.EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _EchoService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golang.conradwood.net/apis/getestservice/echoservice.proto",
}

// Client API for EchoStreamService service

type EchoStreamServiceClient interface {
	SendToServer(ctx context.Context, opts ...grpc.CallOption) (EchoStreamService_SendToServerClient, error)
}

type echoStreamServiceClient struct {
	cc *grpc.ClientConn
}

func NewEchoStreamServiceClient(cc *grpc.ClientConn) EchoStreamServiceClient {
	return &echoStreamServiceClient{cc}
}

func (c *echoStreamServiceClient) SendToServer(ctx context.Context, opts ...grpc.CallOption) (EchoStreamService_SendToServerClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_EchoStreamService_serviceDesc.Streams[0], c.cc, "/getestservice.EchoStreamService/SendToServer", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoStreamServiceSendToServerClient{stream}
	return x, nil
}

type EchoStreamService_SendToServerClient interface {
	Send(*PingRequest) error
	CloseAndRecv() (*PingResponse, error)
	grpc.ClientStream
}

type echoStreamServiceSendToServerClient struct {
	grpc.ClientStream
}

func (x *echoStreamServiceSendToServerClient) Send(m *PingRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *echoStreamServiceSendToServerClient) CloseAndRecv() (*PingResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for EchoStreamService service

type EchoStreamServiceServer interface {
	SendToServer(EchoStreamService_SendToServerServer) error
}

func RegisterEchoStreamServiceServer(s *grpc.Server, srv EchoStreamServiceServer) {
	s.RegisterService(&_EchoStreamService_serviceDesc, srv)
}

func _EchoStreamService_SendToServer_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EchoStreamServiceServer).SendToServer(&echoStreamServiceSendToServerServer{stream})
}

type EchoStreamService_SendToServerServer interface {
	SendAndClose(*PingResponse) error
	Recv() (*PingRequest, error)
	grpc.ServerStream
}

type echoStreamServiceSendToServerServer struct {
	grpc.ServerStream
}

func (x *echoStreamServiceSendToServerServer) SendAndClose(m *PingResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *echoStreamServiceSendToServerServer) Recv() (*PingRequest, error) {
	m := new(PingRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _EchoStreamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "getestservice.EchoStreamService",
	HandlerType: (*EchoStreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendToServer",
			Handler:       _EchoStreamService_SendToServer_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "golang.conradwood.net/apis/getestservice/echoservice.proto",
}

func init() {
	proto.RegisterFile("golang.conradwood.net/apis/getestservice/echoservice.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x89, 0x15, 0xff, 0x4c, 0x5a, 0xd1, 0x3d, 0x85, 0x7a, 0x29, 0x3d, 0x48, 0x44, 0xd8,
	0x40, 0x05, 0x0f, 0x5e, 0x04, 0x41, 0x41, 0x90, 0x52, 0x92, 0x9c, 0x85, 0xed, 0x66, 0x48, 0x03,
	0xed, 0x4e, 0xdc, 0xdd, 0x2a, 0x7e, 0x7b, 0xd9, 0x6d, 0x22, 0x89, 0x88, 0x87, 0xde, 0xde, 0xee,
	0xbc, 0xf7, 0x9b, 0x81, 0x07, 0xf7, 0x25, 0xad, 0x85, 0x2a, 0xb9, 0x24, 0xa5, 0x45, 0xf1, 0x49,
	0x54, 0x70, 0x85, 0x36, 0x11, 0x75, 0x65, 0x92, 0x12, 0x2d, 0x1a, 0x6b, 0x50, 0x7f, 0x54, 0x12,
	0x13, 0x94, 0x2b, 0x6a, 0x34, 0xaf, 0x35, 0x59, 0x62, 0xa3, 0x9e, 0x61, 0x2a, 0x20, 0x5c, 0x54,
	0xaa, 0x4c, 0xf1, 0x7d, 0x8b, 0xc6, 0xb2, 0x2b, 0x38, 0xcb, 0x9c, 0x54, 0x12, 0xe7, 0xdb, 0xcd,
	0x12, 0x75, 0x14, 0x4c, 0x82, 0x78, 0x94, 0xfe, 0xfa, 0x65, 0x11, 0x1c, 0x2f, 0xc4, 0xd7, 0x9a,
	0x44, 0x11, 0x1d, 0x4c, 0x82, 0xf8, 0x34, 0x6d, 0x9f, 0xec, 0x1c, 0x06, 0x79, 0xfe, 0x1a, 0x0d,
	0x7c, 0xcc, 0xc9, 0xe9, 0x33, 0x0c, 0x77, 0x2b, 0x4c, 0x4d, 0xca, 0x20, 0xbb, 0x83, 0x93, 0x56,
	0x7b, 0x7a, 0x38, 0x1b, 0xf3, 0xde, 0x51, 0xbc, 0x73, 0x51, 0xfa, 0xe3, 0x9d, 0xcd, 0x21, 0x7c,
	0x92, 0x2b, 0xca, 0x76, 0x26, 0xf6, 0x00, 0x87, 0xce, 0xc7, 0xfe, 0x09, 0x8f, 0x2f, 0xff, 0x9c,
	0x35, 0xbc, 0x37, 0xb8, 0xf0, 0x3c, 0xab, 0x51, 0x6c, 0x5a, 0xea, 0x0b, 0x0c, 0x33, 0x54, 0x45,
	0xee, 0xd7, 0xa0, 0xde, 0x9b, 0x1e, 0x07, 0x8f, 0x37, 0x70, 0xad, 0xd0, 0x76, 0x4b, 0x6a, 0x6a,
	0x73, 0x3d, 0xb9, 0x68, 0x27, 0xbb, 0x3c, 0xf2, 0xed, 0xdc, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff,
	0x35, 0xfb, 0xe0, 0xdd, 0xdb, 0x01, 0x00, 0x00,
}