// Code generated by protoc-gen-go.
// source: users.singingcat.net/1/clock/clock.proto
// DO NOT EDIT!

/*
Package clock is a generated protocol buffer package.

It is generated from these files:
	users.singingcat.net/1/clock/clock.proto

It has these top-level messages:
	Config
*/
package clock

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// an example
type Config struct {
	Debug bool `protobuf:"varint,1,opt,name=Debug" json:"Debug,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Config) GetDebug() bool {
	if m != nil {
		return m.Debug
	}
	return false
}

func init() {
	proto.RegisterType((*Config)(nil), "clock.Config")
}

func init() { proto.RegisterFile("users.singingcat.net/1/clock/clock.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x24, 0xcd, 0xb1, 0xca, 0xc2, 0x30,
	0x10, 0xc0, 0x71, 0x3a, 0xb4, 0x7c, 0x64, 0x2c, 0xdf, 0xe0, 0x24, 0x87, 0x53, 0xa7, 0x14, 0x71,
	0x14, 0x17, 0xf5, 0x09, 0x7c, 0x83, 0x6b, 0x7a, 0xc6, 0xd8, 0x90, 0x0b, 0xb9, 0x8b, 0xe0, 0xdb,
	0x0b, 0xed, 0xf2, 0x87, 0xdf, 0xf4, 0x37, 0x43, 0x15, 0x2a, 0x62, 0x25, 0x24, 0x1f, 0x92, 0x77,
	0xa8, 0x36, 0x91, 0x8e, 0xc7, 0xd1, 0x45, 0x76, 0xcb, 0x56, 0x9b, 0x0b, 0x2b, 0xf7, 0xed, 0x8a,
	0xc3, 0xde, 0x74, 0x37, 0x4e, 0xcf, 0xe0, 0xfb, 0x7f, 0xd3, 0xde, 0x69, 0xaa, 0x7e, 0xd7, 0x40,
	0x33, 0xfc, 0x3d, 0x36, 0x5c, 0x2f, 0xe6, 0x5c, 0x28, 0xf3, 0x54, 0x43, 0x9c, 0xa9, 0xd8, 0x88,
	0x4a, 0x19, 0xd5, 0xbd, 0x42, 0xf2, 0x30, 0x33, 0x09, 0x24, 0x56, 0xf8, 0x92, 0x82, 0xd4, 0x9c,
	0xb9, 0x28, 0xbc, 0xf1, 0x83, 0x90, 0xd1, 0x2d, 0xe8, 0x49, 0xa6, 0x6e, 0x9d, 0x9d, 0x7e, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x52, 0xae, 0x87, 0xc7, 0x98, 0x00, 0x00, 0x00,
}