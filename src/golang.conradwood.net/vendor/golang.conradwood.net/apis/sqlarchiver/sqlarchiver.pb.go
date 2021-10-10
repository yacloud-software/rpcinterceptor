// Code generated by protoc-gen-go.
// source: golang.conradwood.net/apis/sqlarchiver/sqlarchiver.proto
// DO NOT EDIT!

/*
Package sqlarchiver is a generated protocol buffer package.

It is generated from these files:
	golang.conradwood.net/apis/sqlarchiver/sqlarchiver.proto

It has these top-level messages:
	ArchiveJob
	Archive
	CreateArchiveJobRequest
	CreateArchiveJobResponse
	RunArchiveJobRequest
	DeleteArchiveJobRequest
	ListArchiveJobsResponse
	ListArchivesResponse
	ListArchiveEntriesRequest
	ListArchiveEntriesResponse
*/
package sqlarchiver

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

// request to archive rows from given database table every 24 hrs
type ArchiveJob struct {
	// name of database
	// postgres or mysql
	DBType string `protobuf:"bytes,1,opt,name=DBType" json:"DBType,omitempty"`
	// name of database
	DBDb string `protobuf:"bytes,2,opt,name=DBDb" json:"DBDb,omitempty"`
	// datanamse user name
	DBUser string `protobuf:"bytes,3,opt,name=DBUser" json:"DBUser,omitempty"`
	// database password
	DBPw string `protobuf:"bytes,4,opt,name=DBPw" json:"DBPw,omitempty"`
	// databaser host
	DBHost string `protobuf:"bytes,5,opt,name=DBHost" json:"DBHost,omitempty"`
	// table to archive
	Table string `protobuf:"bytes,6,opt,name=Table" json:"Table,omitempty"`
	// whhere cluase for select & delete
	Where string `protobuf:"bytes,7,opt,name=Where" json:"Where,omitempty"`
	// repeat period for job ( in minutes)
	Period uint32 `protobuf:"varint,8,opt,name=Period" json:"Period,omitempty"`
	// job id used only job queries, ignored for create
	ID uint64 `protobuf:"varint,9,opt,name=ID" json:"ID,omitempty"`
}

func (m *ArchiveJob) Reset()                    { *m = ArchiveJob{} }
func (m *ArchiveJob) String() string            { return proto.CompactTextString(m) }
func (*ArchiveJob) ProtoMessage()               {}
func (*ArchiveJob) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ArchiveJob) GetDBType() string {
	if m != nil {
		return m.DBType
	}
	return ""
}

func (m *ArchiveJob) GetDBDb() string {
	if m != nil {
		return m.DBDb
	}
	return ""
}

func (m *ArchiveJob) GetDBUser() string {
	if m != nil {
		return m.DBUser
	}
	return ""
}

func (m *ArchiveJob) GetDBPw() string {
	if m != nil {
		return m.DBPw
	}
	return ""
}

func (m *ArchiveJob) GetDBHost() string {
	if m != nil {
		return m.DBHost
	}
	return ""
}

func (m *ArchiveJob) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *ArchiveJob) GetWhere() string {
	if m != nil {
		return m.Where
	}
	return ""
}

func (m *ArchiveJob) GetPeriod() uint32 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *ArchiveJob) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

// description of an archive
type Archive struct {
	// job used to create
	JobID uint64 `protobuf:"varint,1,opt,name=JobID" json:"JobID,omitempty"`
	// time created
	Timestamp uint32 `protobuf:"varint,2,opt,name=Timestamp" json:"Timestamp,omitempty"`
	// number of entries
	EntryCount uint32 `protobuf:"varint,3,opt,name=EntryCount" json:"EntryCount,omitempty"`
	// archive id used only archive queries, ignored for create
	ID uint64 `protobuf:"varint,4,opt,name=ID" json:"ID,omitempty"`
}

func (m *Archive) Reset()                    { *m = Archive{} }
func (m *Archive) String() string            { return proto.CompactTextString(m) }
func (*Archive) ProtoMessage()               {}
func (*Archive) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Archive) GetJobID() uint64 {
	if m != nil {
		return m.JobID
	}
	return 0
}

func (m *Archive) GetTimestamp() uint32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Archive) GetEntryCount() uint32 {
	if m != nil {
		return m.EntryCount
	}
	return 0
}

func (m *Archive) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

type CreateArchiveJobRequest struct {
	ArchiveJob *ArchiveJob `protobuf:"bytes,1,opt,name=ArchiveJob" json:"ArchiveJob,omitempty"`
}

func (m *CreateArchiveJobRequest) Reset()                    { *m = CreateArchiveJobRequest{} }
func (m *CreateArchiveJobRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateArchiveJobRequest) ProtoMessage()               {}
func (*CreateArchiveJobRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateArchiveJobRequest) GetArchiveJob() *ArchiveJob {
	if m != nil {
		return m.ArchiveJob
	}
	return nil
}

type CreateArchiveJobResponse struct {
	JobID uint64 `protobuf:"varint,1,opt,name=JobID" json:"JobID,omitempty"`
}

func (m *CreateArchiveJobResponse) Reset()                    { *m = CreateArchiveJobResponse{} }
func (m *CreateArchiveJobResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateArchiveJobResponse) ProtoMessage()               {}
func (*CreateArchiveJobResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CreateArchiveJobResponse) GetJobID() uint64 {
	if m != nil {
		return m.JobID
	}
	return 0
}

type RunArchiveJobRequest struct {
	JobID uint64 `protobuf:"varint,1,opt,name=JobID" json:"JobID,omitempty"`
}

func (m *RunArchiveJobRequest) Reset()                    { *m = RunArchiveJobRequest{} }
func (m *RunArchiveJobRequest) String() string            { return proto.CompactTextString(m) }
func (*RunArchiveJobRequest) ProtoMessage()               {}
func (*RunArchiveJobRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RunArchiveJobRequest) GetJobID() uint64 {
	if m != nil {
		return m.JobID
	}
	return 0
}

type DeleteArchiveJobRequest struct {
	JobID uint64 `protobuf:"varint,1,opt,name=JobID" json:"JobID,omitempty"`
}

func (m *DeleteArchiveJobRequest) Reset()                    { *m = DeleteArchiveJobRequest{} }
func (m *DeleteArchiveJobRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteArchiveJobRequest) ProtoMessage()               {}
func (*DeleteArchiveJobRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DeleteArchiveJobRequest) GetJobID() uint64 {
	if m != nil {
		return m.JobID
	}
	return 0
}

type ListArchiveJobsResponse struct {
	ArchiveJobs []*ArchiveJob `protobuf:"bytes,1,rep,name=ArchiveJobs" json:"ArchiveJobs,omitempty"`
}

func (m *ListArchiveJobsResponse) Reset()                    { *m = ListArchiveJobsResponse{} }
func (m *ListArchiveJobsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListArchiveJobsResponse) ProtoMessage()               {}
func (*ListArchiveJobsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ListArchiveJobsResponse) GetArchiveJobs() []*ArchiveJob {
	if m != nil {
		return m.ArchiveJobs
	}
	return nil
}

type ListArchivesResponse struct {
	Archives []*Archive `protobuf:"bytes,1,rep,name=Archives" json:"Archives,omitempty"`
}

func (m *ListArchivesResponse) Reset()                    { *m = ListArchivesResponse{} }
func (m *ListArchivesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListArchivesResponse) ProtoMessage()               {}
func (*ListArchivesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ListArchivesResponse) GetArchives() []*Archive {
	if m != nil {
		return m.Archives
	}
	return nil
}

type ListArchiveEntriesRequest struct {
	ArchiveID uint64 `protobuf:"varint,1,opt,name=ArchiveID" json:"ArchiveID,omitempty"`
	Limit     uint64 `protobuf:"varint,2,opt,name=Limit" json:"Limit,omitempty"`
}

func (m *ListArchiveEntriesRequest) Reset()                    { *m = ListArchiveEntriesRequest{} }
func (m *ListArchiveEntriesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListArchiveEntriesRequest) ProtoMessage()               {}
func (*ListArchiveEntriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ListArchiveEntriesRequest) GetArchiveID() uint64 {
	if m != nil {
		return m.ArchiveID
	}
	return 0
}

func (m *ListArchiveEntriesRequest) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ListArchiveEntriesResponse struct {
	Entries []string `protobuf:"bytes,1,rep,name=Entries" json:"Entries,omitempty"`
}

func (m *ListArchiveEntriesResponse) Reset()                    { *m = ListArchiveEntriesResponse{} }
func (m *ListArchiveEntriesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListArchiveEntriesResponse) ProtoMessage()               {}
func (*ListArchiveEntriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ListArchiveEntriesResponse) GetEntries() []string {
	if m != nil {
		return m.Entries
	}
	return nil
}

func init() {
	proto.RegisterType((*ArchiveJob)(nil), "sqlarchiver.ArchiveJob")
	proto.RegisterType((*Archive)(nil), "sqlarchiver.Archive")
	proto.RegisterType((*CreateArchiveJobRequest)(nil), "sqlarchiver.CreateArchiveJobRequest")
	proto.RegisterType((*CreateArchiveJobResponse)(nil), "sqlarchiver.CreateArchiveJobResponse")
	proto.RegisterType((*RunArchiveJobRequest)(nil), "sqlarchiver.RunArchiveJobRequest")
	proto.RegisterType((*DeleteArchiveJobRequest)(nil), "sqlarchiver.DeleteArchiveJobRequest")
	proto.RegisterType((*ListArchiveJobsResponse)(nil), "sqlarchiver.ListArchiveJobsResponse")
	proto.RegisterType((*ListArchivesResponse)(nil), "sqlarchiver.ListArchivesResponse")
	proto.RegisterType((*ListArchiveEntriesRequest)(nil), "sqlarchiver.ListArchiveEntriesRequest")
	proto.RegisterType((*ListArchiveEntriesResponse)(nil), "sqlarchiver.ListArchiveEntriesResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SQLArchiverJobService service

type SQLArchiverJobServiceClient interface {
	// create a job
	CreateArchiveJob(ctx context.Context, in *CreateArchiveJobRequest, opts ...grpc.CallOption) (*CreateArchiveJobResponse, error)
	// list jobs
	ListArchiveJobs(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*ListArchiveJobsResponse, error)
	// delete a job
	DeleteArchiveJob(ctx context.Context, in *DeleteArchiveJobRequest, opts ...grpc.CallOption) (*common.Void, error)
	// run a job
	RunArchiveJobNow(ctx context.Context, in *RunArchiveJobRequest, opts ...grpc.CallOption) (*common.Void, error)
	// list archives by job id & time
	ListArchives(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*ListArchivesResponse, error)
	// list the entries (insert statements) in an archive
	ListArchiveEntries(ctx context.Context, in *ListArchiveEntriesRequest, opts ...grpc.CallOption) (*ListArchiveEntriesResponse, error)
}

type sQLArchiverJobServiceClient struct {
	cc *grpc.ClientConn
}

func NewSQLArchiverJobServiceClient(cc *grpc.ClientConn) SQLArchiverJobServiceClient {
	return &sQLArchiverJobServiceClient{cc}
}

func (c *sQLArchiverJobServiceClient) CreateArchiveJob(ctx context.Context, in *CreateArchiveJobRequest, opts ...grpc.CallOption) (*CreateArchiveJobResponse, error) {
	out := new(CreateArchiveJobResponse)
	err := grpc.Invoke(ctx, "/sqlarchiver.SQLArchiverJobService/CreateArchiveJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sQLArchiverJobServiceClient) ListArchiveJobs(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*ListArchiveJobsResponse, error) {
	out := new(ListArchiveJobsResponse)
	err := grpc.Invoke(ctx, "/sqlarchiver.SQLArchiverJobService/ListArchiveJobs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sQLArchiverJobServiceClient) DeleteArchiveJob(ctx context.Context, in *DeleteArchiveJobRequest, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/sqlarchiver.SQLArchiverJobService/DeleteArchiveJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sQLArchiverJobServiceClient) RunArchiveJobNow(ctx context.Context, in *RunArchiveJobRequest, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/sqlarchiver.SQLArchiverJobService/RunArchiveJobNow", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sQLArchiverJobServiceClient) ListArchives(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*ListArchivesResponse, error) {
	out := new(ListArchivesResponse)
	err := grpc.Invoke(ctx, "/sqlarchiver.SQLArchiverJobService/ListArchives", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sQLArchiverJobServiceClient) ListArchiveEntries(ctx context.Context, in *ListArchiveEntriesRequest, opts ...grpc.CallOption) (*ListArchiveEntriesResponse, error) {
	out := new(ListArchiveEntriesResponse)
	err := grpc.Invoke(ctx, "/sqlarchiver.SQLArchiverJobService/ListArchiveEntries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SQLArchiverJobService service

type SQLArchiverJobServiceServer interface {
	// create a job
	CreateArchiveJob(context.Context, *CreateArchiveJobRequest) (*CreateArchiveJobResponse, error)
	// list jobs
	ListArchiveJobs(context.Context, *common.Void) (*ListArchiveJobsResponse, error)
	// delete a job
	DeleteArchiveJob(context.Context, *DeleteArchiveJobRequest) (*common.Void, error)
	// run a job
	RunArchiveJobNow(context.Context, *RunArchiveJobRequest) (*common.Void, error)
	// list archives by job id & time
	ListArchives(context.Context, *common.Void) (*ListArchivesResponse, error)
	// list the entries (insert statements) in an archive
	ListArchiveEntries(context.Context, *ListArchiveEntriesRequest) (*ListArchiveEntriesResponse, error)
}

func RegisterSQLArchiverJobServiceServer(s *grpc.Server, srv SQLArchiverJobServiceServer) {
	s.RegisterService(&_SQLArchiverJobService_serviceDesc, srv)
}

func _SQLArchiverJobService_CreateArchiveJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArchiveJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SQLArchiverJobServiceServer).CreateArchiveJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sqlarchiver.SQLArchiverJobService/CreateArchiveJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SQLArchiverJobServiceServer).CreateArchiveJob(ctx, req.(*CreateArchiveJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SQLArchiverJobService_ListArchiveJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SQLArchiverJobServiceServer).ListArchiveJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sqlarchiver.SQLArchiverJobService/ListArchiveJobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SQLArchiverJobServiceServer).ListArchiveJobs(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _SQLArchiverJobService_DeleteArchiveJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArchiveJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SQLArchiverJobServiceServer).DeleteArchiveJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sqlarchiver.SQLArchiverJobService/DeleteArchiveJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SQLArchiverJobServiceServer).DeleteArchiveJob(ctx, req.(*DeleteArchiveJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SQLArchiverJobService_RunArchiveJobNow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunArchiveJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SQLArchiverJobServiceServer).RunArchiveJobNow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sqlarchiver.SQLArchiverJobService/RunArchiveJobNow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SQLArchiverJobServiceServer).RunArchiveJobNow(ctx, req.(*RunArchiveJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SQLArchiverJobService_ListArchives_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SQLArchiverJobServiceServer).ListArchives(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sqlarchiver.SQLArchiverJobService/ListArchives",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SQLArchiverJobServiceServer).ListArchives(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _SQLArchiverJobService_ListArchiveEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListArchiveEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SQLArchiverJobServiceServer).ListArchiveEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sqlarchiver.SQLArchiverJobService/ListArchiveEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SQLArchiverJobServiceServer).ListArchiveEntries(ctx, req.(*ListArchiveEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SQLArchiverJobService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sqlarchiver.SQLArchiverJobService",
	HandlerType: (*SQLArchiverJobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateArchiveJob",
			Handler:    _SQLArchiverJobService_CreateArchiveJob_Handler,
		},
		{
			MethodName: "ListArchiveJobs",
			Handler:    _SQLArchiverJobService_ListArchiveJobs_Handler,
		},
		{
			MethodName: "DeleteArchiveJob",
			Handler:    _SQLArchiverJobService_DeleteArchiveJob_Handler,
		},
		{
			MethodName: "RunArchiveJobNow",
			Handler:    _SQLArchiverJobService_RunArchiveJobNow_Handler,
		},
		{
			MethodName: "ListArchives",
			Handler:    _SQLArchiverJobService_ListArchives_Handler,
		},
		{
			MethodName: "ListArchiveEntries",
			Handler:    _SQLArchiverJobService_ListArchiveEntries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golang.conradwood.net/apis/sqlarchiver/sqlarchiver.proto",
}

func init() {
	proto.RegisterFile("golang.conradwood.net/apis/sqlarchiver/sqlarchiver.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 574 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x54, 0x5d, 0x6f, 0xda, 0x30,
	0x14, 0x55, 0xda, 0x14, 0xca, 0xa5, 0x6c, 0xc8, 0x62, 0xc3, 0x8b, 0xaa, 0x89, 0x46, 0x5b, 0xc7,
	0xc3, 0x14, 0xaa, 0x4e, 0xda, 0xc7, 0xd3, 0x34, 0x48, 0xa7, 0x82, 0xd0, 0xc6, 0x5c, 0xb6, 0x3d,
	0x4e, 0x09, 0x58, 0x6d, 0x24, 0x12, 0xd3, 0x24, 0x14, 0xf5, 0x5f, 0xee, 0xc7, 0xec, 0x07, 0x4c,
	0xb1, 0x0d, 0x71, 0x42, 0xd2, 0x3e, 0x91, 0x73, 0x38, 0xf7, 0xf8, 0x9e, 0xdc, 0x1b, 0xc3, 0xc7,
	0x6b, 0xb6, 0x70, 0x82, 0x6b, 0x6b, 0xc6, 0x82, 0xd0, 0x99, 0xaf, 0x19, 0x9b, 0x5b, 0x01, 0x8d,
	0x7b, 0xce, 0xd2, 0x8b, 0x7a, 0xd1, 0xed, 0xc2, 0x09, 0x67, 0x37, 0xde, 0x1d, 0x0d, 0xd5, 0x67,
	0x6b, 0x19, 0xb2, 0x98, 0xa1, 0xba, 0x42, 0x19, 0xd6, 0x03, 0x36, 0x33, 0xe6, 0xfb, 0x2c, 0x90,
	0x3f, 0xa2, 0xd8, 0xfc, 0xab, 0x01, 0x7c, 0x11, 0xc5, 0x23, 0xe6, 0xa2, 0xe7, 0x50, 0xb1, 0xfb,
	0xd3, 0xfb, 0x25, 0xc5, 0x5a, 0x47, 0xeb, 0xd6, 0x88, 0x44, 0x08, 0x81, 0x6e, 0xf7, 0x6d, 0x17,
	0xef, 0x71, 0x96, 0x3f, 0x0b, 0xed, 0xcf, 0x88, 0x86, 0x78, 0x7f, 0xa3, 0x4d, 0x90, 0xd0, 0x4e,
	0xd6, 0x58, 0xdf, 0x68, 0x27, 0x6b, 0xa1, 0xbd, 0x64, 0x51, 0x8c, 0x0f, 0x36, 0xda, 0x04, 0xa1,
	0x16, 0x1c, 0x4c, 0x1d, 0x77, 0x41, 0x71, 0x85, 0xd3, 0x02, 0x24, 0xec, 0xef, 0x1b, 0x1a, 0x52,
	0x5c, 0x15, 0x2c, 0x07, 0x89, 0xc7, 0x84, 0x86, 0x1e, 0x9b, 0xe3, 0xc3, 0x8e, 0xd6, 0x6d, 0x10,
	0x89, 0xd0, 0x13, 0xd8, 0x1b, 0xda, 0xb8, 0xd6, 0xd1, 0xba, 0x3a, 0xd9, 0x1b, 0xda, 0xa6, 0x0f,
	0x55, 0x99, 0x28, 0x31, 0x1a, 0x31, 0x77, 0x68, 0xf3, 0x34, 0x3a, 0x11, 0x00, 0x1d, 0x43, 0x6d,
	0xea, 0xf9, 0x34, 0x8a, 0x1d, 0x7f, 0xc9, 0x13, 0x35, 0x48, 0x4a, 0xa0, 0x97, 0x00, 0x17, 0x41,
	0x1c, 0xde, 0x0f, 0xd8, 0x2a, 0x88, 0x79, 0xb4, 0x06, 0x51, 0x18, 0x79, 0x9c, 0xbe, 0x3d, 0x8e,
	0x40, 0x7b, 0x10, 0x52, 0x27, 0xa6, 0xe9, 0x6b, 0x24, 0xf4, 0x76, 0x45, 0xa3, 0x18, 0x7d, 0x50,
	0xdf, 0x2d, 0xef, 0xa1, 0x7e, 0xde, 0xb6, 0xd4, 0x09, 0x2a, 0x35, 0x8a, 0xd4, 0x3c, 0x03, 0xbc,
	0xeb, 0x19, 0x2d, 0x59, 0x10, 0x95, 0x64, 0x32, 0xdf, 0x42, 0x8b, 0xac, 0x82, 0xdd, 0x16, 0x8a,
	0xd5, 0x3d, 0x68, 0xdb, 0x74, 0x41, 0x8b, 0x7a, 0x2e, 0x2e, 0x98, 0x42, 0x7b, 0xec, 0x45, 0x71,
	0x2a, 0x8f, 0xb6, 0xfd, 0x7c, 0x82, 0xba, 0x42, 0x63, 0xad, 0xb3, 0xff, 0x50, 0x4a, 0x55, 0x6b,
	0x5e, 0x42, 0x4b, 0x71, 0x4d, 0x2d, 0xcf, 0xe0, 0x70, 0xc3, 0x49, 0xbf, 0x56, 0x91, 0x1f, 0xd9,
	0xaa, 0xcc, 0xef, 0xf0, 0x42, 0x71, 0x4a, 0xa6, 0xe5, 0x25, 0x7e, 0x22, 0xd2, 0x31, 0xd4, 0xe4,
	0x1f, 0xdb, 0x58, 0x29, 0x91, 0x04, 0x1e, 0x7b, 0xbe, 0x17, 0xf3, 0x4d, 0xd0, 0x89, 0x00, 0xe6,
	0x7b, 0x30, 0x8a, 0x0c, 0x65, 0x83, 0x18, 0xaa, 0x92, 0xe2, 0xfd, 0xd5, 0xc8, 0x06, 0x9e, 0xff,
	0xdb, 0x87, 0x67, 0x57, 0x3f, 0xc6, 0xb2, 0x2e, 0x1c, 0x31, 0xf7, 0x8a, 0x86, 0x77, 0xde, 0x8c,
	0xa2, 0x3f, 0xd0, 0xcc, 0xcf, 0x14, 0xbd, 0xca, 0xc4, 0x2a, 0x59, 0x23, 0xe3, 0xf5, 0x23, 0x2a,
	0xd9, 0xd4, 0x05, 0x3c, 0xcd, 0xcd, 0x08, 0x1d, 0x59, 0xf2, 0x63, 0xff, 0xc5, 0xbc, 0xb9, 0x91,
	0x3d, 0xad, 0x6c, 0x9e, 0x5f, 0xa1, 0x99, 0xdf, 0x8d, 0x5c, 0x9f, 0x25, 0xab, 0x63, 0x64, 0x4e,
	0x43, 0x03, 0x68, 0x66, 0x36, 0xf2, 0x1b, 0x5b, 0xa3, 0x93, 0x8c, 0x4f, 0xd1, 0xc2, 0xe6, 0x4c,
	0x3e, 0xc3, 0x91, 0xba, 0x21, 0xb9, 0x40, 0x27, 0x65, 0x81, 0xd2, 0x34, 0x14, 0xd0, 0xee, 0x1c,
	0xd1, 0x69, 0x59, 0x61, 0x76, 0x73, 0x8c, 0x37, 0x8f, 0xea, 0xc4, 0x31, 0xfd, 0x2e, 0x9c, 0x06,
	0x34, 0x56, 0x6f, 0x5d, 0x79, 0x0f, 0x27, 0x17, 0xaf, 0x6a, 0xe2, 0x56, 0xf8, 0xbd, 0xfb, 0xee,
	0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x56, 0x99, 0xc7, 0xbe, 0xf0, 0x05, 0x00, 0x00,
}