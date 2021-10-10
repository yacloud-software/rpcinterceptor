// Code generated by protoc-gen-go.
// source: golang.conradwood.net/apis/dhcpconfigurator/dhcpconfigurator.proto
// DO NOT EDIT!

/*
Package dhcpconfigurator is a generated protocol buffer package.

It is generated from these files:
	golang.conradwood.net/apis/dhcpconfigurator/dhcpconfigurator.proto

It has these top-level messages:
	SetMappingRequest
*/
package dhcpconfigurator

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

type SetMappingRequest struct {
	Hostname string `protobuf:"bytes,1,opt,name=Hostname" json:"Hostname,omitempty"`
	MAC      string `protobuf:"bytes,2,opt,name=MAC" json:"MAC,omitempty"`
}

func (m *SetMappingRequest) Reset()                    { *m = SetMappingRequest{} }
func (m *SetMappingRequest) String() string            { return proto.CompactTextString(m) }
func (*SetMappingRequest) ProtoMessage()               {}
func (*SetMappingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SetMappingRequest) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *SetMappingRequest) GetMAC() string {
	if m != nil {
		return m.MAC
	}
	return ""
}

func init() {
	proto.RegisterType((*SetMappingRequest)(nil), "dhcpconfigurator.SetMappingRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DHCPConfigurator service

type DHCPConfiguratorClient interface {
	SetMapping(ctx context.Context, in *SetMappingRequest, opts ...grpc.CallOption) (*common.Void, error)
}

type dHCPConfiguratorClient struct {
	cc *grpc.ClientConn
}

func NewDHCPConfiguratorClient(cc *grpc.ClientConn) DHCPConfiguratorClient {
	return &dHCPConfiguratorClient{cc}
}

func (c *dHCPConfiguratorClient) SetMapping(ctx context.Context, in *SetMappingRequest, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/dhcpconfigurator.DHCPConfigurator/SetMapping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DHCPConfigurator service

type DHCPConfiguratorServer interface {
	SetMapping(context.Context, *SetMappingRequest) (*common.Void, error)
}

func RegisterDHCPConfiguratorServer(s *grpc.Server, srv DHCPConfiguratorServer) {
	s.RegisterService(&_DHCPConfigurator_serviceDesc, srv)
}

func _DHCPConfigurator_SetMapping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetMappingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DHCPConfiguratorServer).SetMapping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dhcpconfigurator.DHCPConfigurator/SetMapping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DHCPConfiguratorServer).SetMapping(ctx, req.(*SetMappingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DHCPConfigurator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dhcpconfigurator.DHCPConfigurator",
	HandlerType: (*DHCPConfiguratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetMapping",
			Handler:    _DHCPConfigurator_SetMapping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golang.conradwood.net/apis/dhcpconfigurator/dhcpconfigurator.proto",
}

func init() {
	proto.RegisterFile("golang.conradwood.net/apis/dhcpconfigurator/dhcpconfigurator.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x72, 0x4a, 0xcf, 0xcf, 0x49,
	0xcc, 0x4b, 0xd7, 0x4b, 0xce, 0xcf, 0x2b, 0x4a, 0x4c, 0x29, 0xcf, 0xcf, 0x4f, 0xd1, 0xcb, 0x4b,
	0x2d, 0xd1, 0x4f, 0x2c, 0xc8, 0x2c, 0xd6, 0x4f, 0xc9, 0x48, 0x2e, 0x48, 0xce, 0xcf, 0x4b, 0xcb,
	0x4c, 0x2f, 0x2d, 0x4a, 0x2c, 0xc9, 0x2f, 0xc2, 0x10, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0x40, 0x17, 0x97, 0xd2, 0xc3, 0x63, 0x6a, 0x72, 0x7e, 0x6e, 0x6e, 0x7e, 0x1e, 0x94, 0x82,
	0x98, 0xa0, 0xe4, 0xc8, 0x25, 0x18, 0x9c, 0x5a, 0xe2, 0x9b, 0x58, 0x50, 0x90, 0x99, 0x97, 0x1e,
	0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x24, 0xc5, 0xc5, 0xe1, 0x91, 0x5f, 0x5c, 0x92, 0x97,
	0x98, 0x9b, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe7, 0x0b, 0x09, 0x70, 0x31, 0xfb,
	0x3a, 0x3a, 0x4b, 0x30, 0x81, 0x85, 0x41, 0x4c, 0xa3, 0x60, 0x2e, 0x01, 0x17, 0x0f, 0xe7, 0x00,
	0x67, 0x24, 0x67, 0x08, 0xd9, 0x73, 0x71, 0x21, 0x8c, 0x15, 0x52, 0xd6, 0xc3, 0x70, 0x3f, 0x86,
	0xa5, 0x52, 0x3c, 0x7a, 0x50, 0x87, 0x85, 0xe5, 0x67, 0xa6, 0x38, 0xe9, 0x72, 0x69, 0xe7, 0xa5,
	0x96, 0x20, 0x7b, 0x03, 0xea, 0x31, 0x90, 0x4f, 0x30, 0x8c, 0x4b, 0x62, 0x03, 0xfb, 0xc6, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0x51, 0x4c, 0x00, 0x3e, 0x55, 0x01, 0x00, 0x00,
}