// Code generated by protoc-gen-go.
// source: golang.conradwood.net/apis/googlecast/googlecast.proto
// DO NOT EDIT!

/*
Package googlecast is a generated protocol buffer package.

It is generated from these files:
	golang.conradwood.net/apis/googlecast/googlecast.proto

It has these top-level messages:
	Device
	DeviceList
	PlaySoundRequest
	GoogleID
	DeviceName
*/
package googlecast

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

type Device struct {
	// local number, volatile
	ID   uint64 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	IP   string `protobuf:"bytes,3,opt,name=IP" json:"IP,omitempty"`
	// stable id, generated by google on the device itself
	GoogleID string `protobuf:"bytes,4,opt,name=GoogleID" json:"GoogleID,omitempty"`
}

func (m *Device) Reset()                    { *m = Device{} }
func (m *Device) String() string            { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()               {}
func (*Device) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Device) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Device) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Device) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *Device) GetGoogleID() string {
	if m != nil {
		return m.GoogleID
	}
	return ""
}

type DeviceList struct {
	// detected devices
	Devices []*Device `protobuf:"bytes,1,rep,name=Devices" json:"Devices,omitempty"`
}

func (m *DeviceList) Reset()                    { *m = DeviceList{} }
func (m *DeviceList) String() string            { return proto.CompactTextString(m) }
func (*DeviceList) ProtoMessage()               {}
func (*DeviceList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DeviceList) GetDevices() []*Device {
	if m != nil {
		return m.Devices
	}
	return nil
}

type PlaySoundRequest struct {
	URL      string `protobuf:"bytes,1,opt,name=URL" json:"URL,omitempty"`
	DeviceID uint64 `protobuf:"varint,2,opt,name=DeviceID" json:"DeviceID,omitempty"`
}

func (m *PlaySoundRequest) Reset()                    { *m = PlaySoundRequest{} }
func (m *PlaySoundRequest) String() string            { return proto.CompactTextString(m) }
func (*PlaySoundRequest) ProtoMessage()               {}
func (*PlaySoundRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PlaySoundRequest) GetURL() string {
	if m != nil {
		return m.URL
	}
	return ""
}

func (m *PlaySoundRequest) GetDeviceID() uint64 {
	if m != nil {
		return m.DeviceID
	}
	return 0
}

type GoogleID struct {
	GoogleID string `protobuf:"bytes,1,opt,name=GoogleID" json:"GoogleID,omitempty"`
}

func (m *GoogleID) Reset()                    { *m = GoogleID{} }
func (m *GoogleID) String() string            { return proto.CompactTextString(m) }
func (*GoogleID) ProtoMessage()               {}
func (*GoogleID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GoogleID) GetGoogleID() string {
	if m != nil {
		return m.GoogleID
	}
	return ""
}

type DeviceName struct {
	Name string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
}

func (m *DeviceName) Reset()                    { *m = DeviceName{} }
func (m *DeviceName) String() string            { return proto.CompactTextString(m) }
func (*DeviceName) ProtoMessage()               {}
func (*DeviceName) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DeviceName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Device)(nil), "googlecast.Device")
	proto.RegisterType((*DeviceList)(nil), "googlecast.DeviceList")
	proto.RegisterType((*PlaySoundRequest)(nil), "googlecast.PlaySoundRequest")
	proto.RegisterType((*GoogleID)(nil), "googlecast.GoogleID")
	proto.RegisterType((*DeviceName)(nil), "googlecast.DeviceName")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GoogleCast service

type GoogleCastClient interface {
	// get a device by name (e.g. "Ducky")
	DeviceByName(ctx context.Context, in *DeviceName, opts ...grpc.CallOption) (*Device, error)
	// get a device by google id
	DeviceByGoogleID(ctx context.Context, in *GoogleID, opts ...grpc.CallOption) (*Device, error)
	ListDevices(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*DeviceList, error)
	PlaySound(ctx context.Context, in *PlaySoundRequest, opts ...grpc.CallOption) (*common.Void, error)
}

type googleCastClient struct {
	cc *grpc.ClientConn
}

func NewGoogleCastClient(cc *grpc.ClientConn) GoogleCastClient {
	return &googleCastClient{cc}
}

func (c *googleCastClient) DeviceByName(ctx context.Context, in *DeviceName, opts ...grpc.CallOption) (*Device, error) {
	out := new(Device)
	err := grpc.Invoke(ctx, "/googlecast.GoogleCast/DeviceByName", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleCastClient) DeviceByGoogleID(ctx context.Context, in *GoogleID, opts ...grpc.CallOption) (*Device, error) {
	out := new(Device)
	err := grpc.Invoke(ctx, "/googlecast.GoogleCast/DeviceByGoogleID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleCastClient) ListDevices(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*DeviceList, error) {
	out := new(DeviceList)
	err := grpc.Invoke(ctx, "/googlecast.GoogleCast/ListDevices", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleCastClient) PlaySound(ctx context.Context, in *PlaySoundRequest, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/googlecast.GoogleCast/PlaySound", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GoogleCast service

type GoogleCastServer interface {
	// get a device by name (e.g. "Ducky")
	DeviceByName(context.Context, *DeviceName) (*Device, error)
	// get a device by google id
	DeviceByGoogleID(context.Context, *GoogleID) (*Device, error)
	ListDevices(context.Context, *common.Void) (*DeviceList, error)
	PlaySound(context.Context, *PlaySoundRequest) (*common.Void, error)
}

func RegisterGoogleCastServer(s *grpc.Server, srv GoogleCastServer) {
	s.RegisterService(&_GoogleCast_serviceDesc, srv)
}

func _GoogleCast_DeviceByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoogleCastServer).DeviceByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/googlecast.GoogleCast/DeviceByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoogleCastServer).DeviceByName(ctx, req.(*DeviceName))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoogleCast_DeviceByGoogleID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoogleID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoogleCastServer).DeviceByGoogleID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/googlecast.GoogleCast/DeviceByGoogleID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoogleCastServer).DeviceByGoogleID(ctx, req.(*GoogleID))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoogleCast_ListDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoogleCastServer).ListDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/googlecast.GoogleCast/ListDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoogleCastServer).ListDevices(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoogleCast_PlaySound_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaySoundRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoogleCastServer).PlaySound(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/googlecast.GoogleCast/PlaySound",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoogleCastServer).PlaySound(ctx, req.(*PlaySoundRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GoogleCast_serviceDesc = grpc.ServiceDesc{
	ServiceName: "googlecast.GoogleCast",
	HandlerType: (*GoogleCastServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeviceByName",
			Handler:    _GoogleCast_DeviceByName_Handler,
		},
		{
			MethodName: "DeviceByGoogleID",
			Handler:    _GoogleCast_DeviceByGoogleID_Handler,
		},
		{
			MethodName: "ListDevices",
			Handler:    _GoogleCast_ListDevices_Handler,
		},
		{
			MethodName: "PlaySound",
			Handler:    _GoogleCast_PlaySound_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golang.conradwood.net/apis/googlecast/googlecast.proto",
}

func init() {
	proto.RegisterFile("golang.conradwood.net/apis/googlecast/googlecast.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x52, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x65, 0xd3, 0x50, 0xed, 0xb4, 0x48, 0x59, 0x44, 0x42, 0x10, 0x09, 0x39, 0x48, 0x0f, 0x92,
	0x42, 0x0b, 0x0a, 0xc5, 0x83, 0xd4, 0x80, 0x04, 0x8a, 0x94, 0x15, 0xc5, 0xeb, 0x9a, 0x2c, 0x21,
	0xd0, 0x66, 0x6a, 0xb3, 0x55, 0xfa, 0xdf, 0x7e, 0x80, 0x64, 0xb7, 0x49, 0xb7, 0x35, 0x78, 0xea,
	0xcc, 0xf4, 0xbd, 0xd9, 0xf7, 0x5e, 0x06, 0x6e, 0x53, 0x5c, 0xf0, 0x3c, 0x0d, 0x62, 0xcc, 0xd7,
	0x3c, 0xf9, 0x46, 0x4c, 0x82, 0x5c, 0xc8, 0x21, 0x5f, 0x65, 0xc5, 0x30, 0x45, 0x4c, 0x17, 0x22,
	0xe6, 0x85, 0x34, 0xca, 0x60, 0xb5, 0x46, 0x89, 0x14, 0xf6, 0x13, 0x37, 0xf8, 0x67, 0x47, 0x8c,
	0xcb, 0x25, 0xe6, 0xbb, 0x1f, 0xcd, 0xf5, 0xdf, 0xa1, 0x1d, 0x8a, 0xaf, 0x2c, 0x16, 0xf4, 0x0c,
	0xac, 0x28, 0x74, 0x88, 0x47, 0x06, 0x36, 0xb3, 0xa2, 0x90, 0x52, 0xb0, 0x9f, 0xf9, 0x52, 0x38,
	0x96, 0x47, 0x06, 0x1d, 0xa6, 0x6a, 0x85, 0x99, 0x3b, 0x2d, 0x35, 0xb1, 0xa2, 0x39, 0x75, 0xe1,
	0xf4, 0x49, 0xbd, 0x1d, 0x85, 0x8e, 0xad, 0xa6, 0x75, 0xef, 0x4f, 0x00, 0xf4, 0xe6, 0x59, 0x56,
	0x48, 0x7a, 0x03, 0x27, 0xba, 0x2b, 0x1c, 0xe2, 0xb5, 0x06, 0xdd, 0x11, 0x0d, 0x0c, 0x1f, 0xfa,
	0x2f, 0x56, 0x41, 0xfc, 0x07, 0xe8, 0xcf, 0x17, 0x7c, 0xfb, 0x82, 0x9b, 0x3c, 0x61, 0xe2, 0x73,
	0x23, 0x0a, 0x49, 0xfb, 0xd0, 0x7a, 0x65, 0x33, 0x25, 0xb0, 0xc3, 0xca, 0xb2, 0x7c, 0x5d, 0x13,
	0xa2, 0x50, 0xa9, 0xb4, 0x59, 0xdd, 0xfb, 0xd7, 0x7b, 0x65, 0x07, 0x2a, 0xc9, 0x91, 0x4a, 0xaf,
	0x52, 0xa9, 0xfc, 0x55, 0x9e, 0xc9, 0xde, 0xf3, 0xe8, 0x87, 0x00, 0x68, 0xf8, 0x23, 0x2f, 0x24,
	0x9d, 0x40, 0x4f, 0x13, 0xa6, 0x5b, 0x45, 0xb9, 0xf8, 0xeb, 0xa3, 0x9c, 0xbb, 0x0d, 0xfe, 0xe8,
	0x3d, 0xf4, 0x2b, 0x6e, 0x2d, 0xee, 0xdc, 0xc4, 0x55, 0xd3, 0x46, 0xf6, 0x18, 0xba, 0x65, 0x94,
	0xbb, 0x8c, 0x68, 0x2f, 0xd8, 0x7d, 0xc8, 0x37, 0xcc, 0x12, 0xb7, 0x41, 0x86, 0xca, 0xfd, 0x0e,
	0x3a, 0x75, 0x92, 0xf4, 0xd2, 0x04, 0x1d, 0x07, 0xec, 0x1e, 0x2c, 0x9c, 0x7a, 0x70, 0x95, 0x0b,
	0x69, 0xde, 0x51, 0x79, 0x43, 0xc6, 0x82, 0x8f, 0xb6, 0xba, 0xa0, 0xf1, 0x6f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x9e, 0x36, 0x03, 0x63, 0xb7, 0x02, 0x00, 0x00,
}