// Code generated by protoc-gen-go.
// source: golang.lakeparime.com/apis/echoservice/echoservice.proto
// DO NOT EDIT!

/*
Package echo is a generated protocol buffer package.

It is generated from these files:
	golang.lakeparime.com/apis/echoservice/echoservice.proto

It has these top-level messages:
	PingResponse
	PingRequest
	BlockchainResponse
*/
package echo

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

// this is returned by "Ping"
type PingResponse struct {
	// this is a string which the echoserver fills in
	Response string `protobuf:"bytes,1,opt,name=Response" json:"Response,omitempty"`
}

func (m *PingResponse) Reset()                    { *m = PingResponse{} }
func (m *PingResponse) String() string            { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()               {}
func (*PingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PingResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

// this is sent to the server. All fields must have a unique number
type PingRequest struct {
	// this is a number which the client fills in
	Number uint64 `protobuf:"varint,1,opt,name=Number" json:"Number,omitempty"`
	// this is a string which the client fills in
	Message string `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PingRequest) GetNumber() uint64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *PingRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type BlockchainResponse struct {
	Hashes uint64 `protobuf:"varint,1,opt,name=Hashes" json:"Hashes,omitempty"`
}

func (m *BlockchainResponse) Reset()                    { *m = BlockchainResponse{} }
func (m *BlockchainResponse) String() string            { return proto.CompactTextString(m) }
func (*BlockchainResponse) ProtoMessage()               {}
func (*BlockchainResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BlockchainResponse) GetHashes() uint64 {
	if m != nil {
		return m.Hashes
	}
	return 0
}

func init() {
	proto.RegisterType((*PingResponse)(nil), "echo.PingResponse")
	proto.RegisterType((*PingRequest)(nil), "echo.PingRequest")
	proto.RegisterType((*BlockchainResponse)(nil), "echo.BlockchainResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Echo service

type EchoClient interface {
	// the simplest of all RPCs. just print out the message and return an error or "OK"
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	// pretend to do a complicated blockchain calculation
	Blockchain(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*BlockchainResponse, error)
}

type echoClient struct {
	cc *grpc.ClientConn
}

func NewEchoClient(cc *grpc.ClientConn) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/echo.Echo/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoClient) Blockchain(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*BlockchainResponse, error) {
	out := new(BlockchainResponse)
	err := grpc.Invoke(ctx, "/echo.Echo/Blockchain", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Echo service

type EchoServer interface {
	// the simplest of all RPCs. just print out the message and return an error or "OK"
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	// pretend to do a complicated blockchain calculation
	Blockchain(context.Context, *common.Void) (*BlockchainResponse, error)
}

func RegisterEchoServer(s *grpc.Server, srv EchoServer) {
	s.RegisterService(&_Echo_serviceDesc, srv)
}

func _Echo_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.Echo/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Echo_Blockchain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).Blockchain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.Echo/Blockchain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).Blockchain(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

var _Echo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "echo.Echo",
	HandlerType: (*EchoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Echo_Ping_Handler,
		},
		{
			MethodName: "Blockchain",
			Handler:    _Echo_Blockchain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golang.lakeparime.com/apis/echoservice/echoservice.proto",
}

func init() {
	proto.RegisterFile("golang.lakeparime.com/apis/echoservice/echoservice.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x50, 0x4d, 0x4b, 0xc4, 0x30,
	0x10, 0x65, 0xa5, 0x54, 0x1d, 0xf7, 0x62, 0x0e, 0x4b, 0xe9, 0x45, 0xed, 0x49, 0x44, 0xb3, 0xa0,
	0x1e, 0xbc, 0x09, 0x0b, 0x82, 0x17, 0x45, 0x7a, 0xf0, 0x9e, 0x4d, 0x87, 0x36, 0xb4, 0xc9, 0xd4,
	0xa4, 0xab, 0x7f, 0x5f, 0xda, 0xa4, 0x35, 0xe0, 0x29, 0x79, 0xc3, 0xfb, 0x98, 0x79, 0xf0, 0x54,
	0x53, 0x27, 0x4c, 0xcd, 0x3b, 0xd1, 0x62, 0x2f, 0xac, 0xd2, 0xc8, 0x25, 0xe9, 0xad, 0xe8, 0x95,
	0xdb, 0xa2, 0x6c, 0xc8, 0xa1, 0xfd, 0x56, 0x12, 0xe3, 0x3f, 0xef, 0x2d, 0x0d, 0xc4, 0x92, 0x71,
	0x94, 0xf3, 0xa0, 0x97, 0x64, 0xac, 0xa8, 0x7e, 0x88, 0x2a, 0x6e, 0x70, 0xf0, 0x7a, 0x49, 0x5a,
	0x93, 0x09, 0x8f, 0x57, 0x15, 0x37, 0xb0, 0xfe, 0x50, 0xa6, 0x2e, 0xd1, 0xf5, 0x64, 0x1c, 0xb2,
	0x1c, 0x4e, 0xe6, 0x7f, 0xb6, 0xba, 0x5c, 0x5d, 0x9f, 0x96, 0x0b, 0x2e, 0x9e, 0xe1, 0xcc, 0x73,
	0xbf, 0x0e, 0xe8, 0x06, 0xb6, 0x81, 0xf4, 0xfd, 0xa0, 0xf7, 0x68, 0x27, 0x62, 0x52, 0x06, 0xc4,
	0x32, 0x38, 0x7e, 0x43, 0xe7, 0x44, 0x8d, 0xd9, 0xd1, 0xe4, 0x30, 0xc3, 0xe2, 0x16, 0xd8, 0xae,
	0x23, 0xd9, 0xca, 0x46, 0x28, 0xb3, 0x44, 0x6e, 0x20, 0x7d, 0x15, 0xae, 0x41, 0x37, 0xfb, 0x78,
	0x74, 0xdf, 0x42, 0xf2, 0x22, 0x1b, 0x62, 0x77, 0x90, 0x8c, 0xb1, 0xec, 0x9c, 0x8f, 0x17, 0xf2,
	0x68, 0x85, 0x9c, 0xc5, 0xa3, 0x60, 0xf7, 0x08, 0xf0, 0x17, 0xc2, 0xd6, 0x3c, 0x9c, 0xfb, 0x49,
	0xaa, 0xca, 0x33, 0xcf, 0xff, 0xbf, 0xc4, 0xee, 0x0a, 0x2e, 0x24, 0xe9, 0xb8, 0xf6, 0xb1, 0x32,
	0x1e, 0xd5, 0xbc, 0x4f, 0xa7, 0xc6, 0x1e, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x27, 0x85, 0xca,
	0x48, 0xa3, 0x01, 0x00, 0x00,
}