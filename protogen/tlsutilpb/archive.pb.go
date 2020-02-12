// Code generated by protoc-gen-go. DO NOT EDIT.
// source: archive.proto

package tlsutilpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SaveConnectionRequest struct {
	Connection           *ConnectionData `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SaveConnectionRequest) Reset()         { *m = SaveConnectionRequest{} }
func (m *SaveConnectionRequest) String() string { return proto.CompactTextString(m) }
func (*SaveConnectionRequest) ProtoMessage()    {}
func (*SaveConnectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_04f37ff213ec9fca, []int{0}
}

func (m *SaveConnectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveConnectionRequest.Unmarshal(m, b)
}
func (m *SaveConnectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveConnectionRequest.Marshal(b, m, deterministic)
}
func (m *SaveConnectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveConnectionRequest.Merge(m, src)
}
func (m *SaveConnectionRequest) XXX_Size() int {
	return xxx_messageInfo_SaveConnectionRequest.Size(m)
}
func (m *SaveConnectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveConnectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveConnectionRequest proto.InternalMessageInfo

func (m *SaveConnectionRequest) GetConnection() *ConnectionData {
	if m != nil {
		return m.Connection
	}
	return nil
}

type SaveConnectionResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveConnectionResponse) Reset()         { *m = SaveConnectionResponse{} }
func (m *SaveConnectionResponse) String() string { return proto.CompactTextString(m) }
func (*SaveConnectionResponse) ProtoMessage()    {}
func (*SaveConnectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_04f37ff213ec9fca, []int{1}
}

func (m *SaveConnectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveConnectionResponse.Unmarshal(m, b)
}
func (m *SaveConnectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveConnectionResponse.Marshal(b, m, deterministic)
}
func (m *SaveConnectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveConnectionResponse.Merge(m, src)
}
func (m *SaveConnectionResponse) XXX_Size() int {
	return xxx_messageInfo_SaveConnectionResponse.Size(m)
}
func (m *SaveConnectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveConnectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SaveConnectionResponse proto.InternalMessageInfo

func (m *SaveConnectionResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type SaveCertificateRequest struct {
	Certificate          *CertificateData `protobuf:"bytes,1,opt,name=certificate,proto3" json:"certificate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SaveCertificateRequest) Reset()         { *m = SaveCertificateRequest{} }
func (m *SaveCertificateRequest) String() string { return proto.CompactTextString(m) }
func (*SaveCertificateRequest) ProtoMessage()    {}
func (*SaveCertificateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_04f37ff213ec9fca, []int{2}
}

func (m *SaveCertificateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveCertificateRequest.Unmarshal(m, b)
}
func (m *SaveCertificateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveCertificateRequest.Marshal(b, m, deterministic)
}
func (m *SaveCertificateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveCertificateRequest.Merge(m, src)
}
func (m *SaveCertificateRequest) XXX_Size() int {
	return xxx_messageInfo_SaveCertificateRequest.Size(m)
}
func (m *SaveCertificateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveCertificateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveCertificateRequest proto.InternalMessageInfo

func (m *SaveCertificateRequest) GetCertificate() *CertificateData {
	if m != nil {
		return m.Certificate
	}
	return nil
}

type SaveCertificateResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveCertificateResponse) Reset()         { *m = SaveCertificateResponse{} }
func (m *SaveCertificateResponse) String() string { return proto.CompactTextString(m) }
func (*SaveCertificateResponse) ProtoMessage()    {}
func (*SaveCertificateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_04f37ff213ec9fca, []int{3}
}

func (m *SaveCertificateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveCertificateResponse.Unmarshal(m, b)
}
func (m *SaveCertificateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveCertificateResponse.Marshal(b, m, deterministic)
}
func (m *SaveCertificateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveCertificateResponse.Merge(m, src)
}
func (m *SaveCertificateResponse) XXX_Size() int {
	return xxx_messageInfo_SaveCertificateResponse.Size(m)
}
func (m *SaveCertificateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveCertificateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SaveCertificateResponse proto.InternalMessageInfo

func (m *SaveCertificateResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type SaveRecordRequest struct {
	Record               *RecordData `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SaveRecordRequest) Reset()         { *m = SaveRecordRequest{} }
func (m *SaveRecordRequest) String() string { return proto.CompactTextString(m) }
func (*SaveRecordRequest) ProtoMessage()    {}
func (*SaveRecordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_04f37ff213ec9fca, []int{4}
}

func (m *SaveRecordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveRecordRequest.Unmarshal(m, b)
}
func (m *SaveRecordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveRecordRequest.Marshal(b, m, deterministic)
}
func (m *SaveRecordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveRecordRequest.Merge(m, src)
}
func (m *SaveRecordRequest) XXX_Size() int {
	return xxx_messageInfo_SaveRecordRequest.Size(m)
}
func (m *SaveRecordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveRecordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveRecordRequest proto.InternalMessageInfo

func (m *SaveRecordRequest) GetRecord() *RecordData {
	if m != nil {
		return m.Record
	}
	return nil
}

func init() {
	proto.RegisterType((*SaveConnectionRequest)(nil), "luids.tlsutil.v1.SaveConnectionRequest")
	proto.RegisterType((*SaveConnectionResponse)(nil), "luids.tlsutil.v1.SaveConnectionResponse")
	proto.RegisterType((*SaveCertificateRequest)(nil), "luids.tlsutil.v1.SaveCertificateRequest")
	proto.RegisterType((*SaveCertificateResponse)(nil), "luids.tlsutil.v1.SaveCertificateResponse")
	proto.RegisterType((*SaveRecordRequest)(nil), "luids.tlsutil.v1.SaveRecordRequest")
}

func init() { proto.RegisterFile("archive.proto", fileDescriptor_04f37ff213ec9fca) }

var fileDescriptor_04f37ff213ec9fca = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x4e, 0xf3, 0x30,
	0x10, 0x6c, 0x7b, 0xe8, 0xa7, 0xcf, 0x55, 0x0b, 0x58, 0xa2, 0xa0, 0x88, 0x43, 0x09, 0x07, 0x52,
	0x04, 0xb6, 0x68, 0xfb, 0x00, 0x40, 0xe1, 0xc0, 0x0d, 0xa5, 0x27, 0x90, 0x38, 0x38, 0xce, 0x36,
	0xb1, 0x94, 0xc4, 0x21, 0x76, 0x2a, 0xf5, 0x9d, 0x78, 0x48, 0x54, 0x27, 0xfd, 0x4d, 0x2b, 0xb8,
	0xf8, 0xb0, 0x9e, 0x99, 0x9d, 0xdd, 0x1d, 0xd4, 0x66, 0x19, 0x0f, 0xc5, 0x0c, 0x48, 0x9a, 0x49,
	0x2d, 0xf1, 0x71, 0x94, 0x0b, 0x5f, 0x11, 0x1d, 0xa9, 0x5c, 0x8b, 0x88, 0xcc, 0xee, 0xad, 0x51,
	0x20, 0x74, 0x98, 0x7b, 0x84, 0xcb, 0x98, 0x9a, 0xcf, 0x3b, 0x21, 0x29, 0x4b, 0x05, 0x55, 0x3c,
	0x84, 0x98, 0x29, 0x5a, 0x62, 0x29, 0x97, 0x71, 0x2c, 0x93, 0x42, 0xc7, 0x1a, 0x6e, 0xb0, 0x02,
	0x19, 0xb1, 0x24, 0xa0, 0xe6, 0xc3, 0xcb, 0xa7, 0x34, 0xd5, 0xf3, 0x14, 0x14, 0x85, 0x38, 0xd5,
	0xf3, 0xe2, 0x2d, 0x48, 0xf6, 0x3b, 0x3a, 0x9d, 0xb0, 0x19, 0x8c, 0x65, 0x92, 0x00, 0xd7, 0x42,
	0x26, 0x2e, 0x7c, 0xe5, 0xa0, 0x34, 0x7e, 0x40, 0x88, 0xaf, 0x8a, 0xe7, 0xf5, 0x5e, 0xdd, 0x69,
	0x0d, 0x7a, 0x64, 0xd7, 0x2a, 0x59, 0x13, 0x9f, 0x99, 0x66, 0xee, 0x06, 0xc7, 0x76, 0x50, 0x77,
	0x57, 0x5a, 0xa5, 0x32, 0x51, 0x80, 0x3b, 0xa8, 0x21, 0x7c, 0xa3, 0xf9, 0xdf, 0x6d, 0x08, 0xdf,
	0xfe, 0x2c, 0x91, 0x90, 0x69, 0x31, 0x15, 0x9c, 0x69, 0x58, 0xba, 0x18, 0xa3, 0x16, 0x5f, 0x57,
	0x4b, 0x1b, 0x97, 0x7b, 0x6c, 0xac, 0x41, 0xc6, 0xc7, 0x26, 0xcb, 0xee, 0xa3, 0xb3, 0x8a, 0xfc,
	0x01, 0x27, 0xaf, 0xe8, 0x64, 0x01, 0x75, 0x81, 0xcb, 0xcc, 0x5f, 0x9a, 0x18, 0xa1, 0x66, 0x66,
	0x0a, 0x65, 0xff, 0x8b, 0x6a, 0xff, 0x82, 0x60, 0x5a, 0x97, 0xd8, 0xc1, 0x77, 0x03, 0xfd, 0x7b,
	0x2c, 0x0e, 0x8d, 0x01, 0x75, 0xb6, 0x57, 0x81, 0xaf, 0xab, 0x1a, 0x7b, 0xef, 0x60, 0x39, 0xbf,
	0x03, 0x8b, 0x59, 0xec, 0x1a, 0x0e, 0xd1, 0xd1, 0xce, 0xa0, 0xf8, 0x10, 0xbd, 0xb2, 0x6a, 0xab,
	0xff, 0x07, 0xe4, 0xaa, 0xd3, 0x1b, 0x6a, 0x4f, 0x74, 0x06, 0x2c, 0x2e, 0x06, 0x57, 0xf8, 0x6a,
	0x3f, 0x7b, 0x6b, 0x91, 0x56, 0x97, 0x04, 0x52, 0x06, 0x51, 0x19, 0x7c, 0x2f, 0x9f, 0x92, 0x97,
	0x45, 0x14, 0xed, 0x9a, 0x53, 0x7f, 0xba, 0xfd, 0xb8, 0x39, 0x94, 0x7a, 0x83, 0x0e, 0x20, 0x59,
	0xc6, 0x3e, 0xf5, 0xbc, 0xa6, 0xa9, 0x0d, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x77, 0x7c, 0xba,
	0x1b, 0x4b, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ArchiveClient is the client API for Archive service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ArchiveClient interface {
	SaveConnection(ctx context.Context, in *SaveConnectionRequest, opts ...grpc.CallOption) (*SaveConnectionResponse, error)
	SaveCertificate(ctx context.Context, in *SaveCertificateRequest, opts ...grpc.CallOption) (*SaveCertificateResponse, error)
	StreamRecords(ctx context.Context, opts ...grpc.CallOption) (Archive_StreamRecordsClient, error)
}

type archiveClient struct {
	cc *grpc.ClientConn
}

func NewArchiveClient(cc *grpc.ClientConn) ArchiveClient {
	return &archiveClient{cc}
}

func (c *archiveClient) SaveConnection(ctx context.Context, in *SaveConnectionRequest, opts ...grpc.CallOption) (*SaveConnectionResponse, error) {
	out := new(SaveConnectionResponse)
	err := c.cc.Invoke(ctx, "/luids.tlsutil.v1.Archive/SaveConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *archiveClient) SaveCertificate(ctx context.Context, in *SaveCertificateRequest, opts ...grpc.CallOption) (*SaveCertificateResponse, error) {
	out := new(SaveCertificateResponse)
	err := c.cc.Invoke(ctx, "/luids.tlsutil.v1.Archive/SaveCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *archiveClient) StreamRecords(ctx context.Context, opts ...grpc.CallOption) (Archive_StreamRecordsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Archive_serviceDesc.Streams[0], "/luids.tlsutil.v1.Archive/StreamRecords", opts...)
	if err != nil {
		return nil, err
	}
	x := &archiveStreamRecordsClient{stream}
	return x, nil
}

type Archive_StreamRecordsClient interface {
	Send(*SaveRecordRequest) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type archiveStreamRecordsClient struct {
	grpc.ClientStream
}

func (x *archiveStreamRecordsClient) Send(m *SaveRecordRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *archiveStreamRecordsClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ArchiveServer is the server API for Archive service.
type ArchiveServer interface {
	SaveConnection(context.Context, *SaveConnectionRequest) (*SaveConnectionResponse, error)
	SaveCertificate(context.Context, *SaveCertificateRequest) (*SaveCertificateResponse, error)
	StreamRecords(Archive_StreamRecordsServer) error
}

// UnimplementedArchiveServer can be embedded to have forward compatible implementations.
type UnimplementedArchiveServer struct {
}

func (*UnimplementedArchiveServer) SaveConnection(ctx context.Context, req *SaveConnectionRequest) (*SaveConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveConnection not implemented")
}
func (*UnimplementedArchiveServer) SaveCertificate(ctx context.Context, req *SaveCertificateRequest) (*SaveCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveCertificate not implemented")
}
func (*UnimplementedArchiveServer) StreamRecords(srv Archive_StreamRecordsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamRecords not implemented")
}

func RegisterArchiveServer(s *grpc.Server, srv ArchiveServer) {
	s.RegisterService(&_Archive_serviceDesc, srv)
}

func _Archive_SaveConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArchiveServer).SaveConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/luids.tlsutil.v1.Archive/SaveConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArchiveServer).SaveConnection(ctx, req.(*SaveConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Archive_SaveCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArchiveServer).SaveCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/luids.tlsutil.v1.Archive/SaveCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArchiveServer).SaveCertificate(ctx, req.(*SaveCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Archive_StreamRecords_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ArchiveServer).StreamRecords(&archiveStreamRecordsServer{stream})
}

type Archive_StreamRecordsServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*SaveRecordRequest, error)
	grpc.ServerStream
}

type archiveStreamRecordsServer struct {
	grpc.ServerStream
}

func (x *archiveStreamRecordsServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *archiveStreamRecordsServer) Recv() (*SaveRecordRequest, error) {
	m := new(SaveRecordRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Archive_serviceDesc = grpc.ServiceDesc{
	ServiceName: "luids.tlsutil.v1.Archive",
	HandlerType: (*ArchiveServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveConnection",
			Handler:    _Archive_SaveConnection_Handler,
		},
		{
			MethodName: "SaveCertificate",
			Handler:    _Archive_SaveCertificate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamRecords",
			Handler:       _Archive_StreamRecords_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "archive.proto",
}