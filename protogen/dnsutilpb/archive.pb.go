// Code generated by protoc-gen-go. DO NOT EDIT.
// source: archive.proto

package dnsutilpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type SaveResolvRequest struct {
	ServerIp             string               `protobuf:"bytes,1,opt,name=server_ip,json=serverIp,proto3" json:"server_ip,omitempty"`
	ClientIp             string               `protobuf:"bytes,2,opt,name=client_ip,json=clientIp,proto3" json:"client_ip,omitempty"`
	Name                 string               `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ResolvedIps          []string             `protobuf:"bytes,4,rep,name=resolved_ips,json=resolvedIps,proto3" json:"resolved_ips,omitempty"`
	Ts                   *timestamp.Timestamp `protobuf:"bytes,5,opt,name=ts,proto3" json:"ts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SaveResolvRequest) Reset()         { *m = SaveResolvRequest{} }
func (m *SaveResolvRequest) String() string { return proto.CompactTextString(m) }
func (*SaveResolvRequest) ProtoMessage()    {}
func (*SaveResolvRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_04f37ff213ec9fca, []int{0}
}

func (m *SaveResolvRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveResolvRequest.Unmarshal(m, b)
}
func (m *SaveResolvRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveResolvRequest.Marshal(b, m, deterministic)
}
func (m *SaveResolvRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveResolvRequest.Merge(m, src)
}
func (m *SaveResolvRequest) XXX_Size() int {
	return xxx_messageInfo_SaveResolvRequest.Size(m)
}
func (m *SaveResolvRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveResolvRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveResolvRequest proto.InternalMessageInfo

func (m *SaveResolvRequest) GetServerIp() string {
	if m != nil {
		return m.ServerIp
	}
	return ""
}

func (m *SaveResolvRequest) GetClientIp() string {
	if m != nil {
		return m.ClientIp
	}
	return ""
}

func (m *SaveResolvRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SaveResolvRequest) GetResolvedIps() []string {
	if m != nil {
		return m.ResolvedIps
	}
	return nil
}

func (m *SaveResolvRequest) GetTs() *timestamp.Timestamp {
	if m != nil {
		return m.Ts
	}
	return nil
}

type SaveResolvResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveResolvResponse) Reset()         { *m = SaveResolvResponse{} }
func (m *SaveResolvResponse) String() string { return proto.CompactTextString(m) }
func (*SaveResolvResponse) ProtoMessage()    {}
func (*SaveResolvResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_04f37ff213ec9fca, []int{1}
}

func (m *SaveResolvResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveResolvResponse.Unmarshal(m, b)
}
func (m *SaveResolvResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveResolvResponse.Marshal(b, m, deterministic)
}
func (m *SaveResolvResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveResolvResponse.Merge(m, src)
}
func (m *SaveResolvResponse) XXX_Size() int {
	return xxx_messageInfo_SaveResolvResponse.Size(m)
}
func (m *SaveResolvResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveResolvResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SaveResolvResponse proto.InternalMessageInfo

func (m *SaveResolvResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*SaveResolvRequest)(nil), "luids.dnsutil.v1.SaveResolvRequest")
	proto.RegisterType((*SaveResolvResponse)(nil), "luids.dnsutil.v1.SaveResolvResponse")
}

func init() { proto.RegisterFile("archive.proto", fileDescriptor_04f37ff213ec9fca) }

var fileDescriptor_04f37ff213ec9fca = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4f, 0xfa, 0x40,
	0x10, 0xc5, 0xff, 0x2d, 0xfc, 0x55, 0x16, 0x35, 0xba, 0xa7, 0x06, 0x2f, 0x88, 0x1c, 0x08, 0xd1,
	0x6d, 0xc4, 0xa3, 0x27, 0xbd, 0x71, 0xad, 0x5e, 0xf4, 0x42, 0x5a, 0x3a, 0x94, 0x49, 0xda, 0xee,
	0xd8, 0xd9, 0x6d, 0xe2, 0xb7, 0xf2, 0x23, 0x1a, 0x76, 0x21, 0x10, 0x4d, 0xbc, 0x4d, 0x7e, 0x6f,
	0xe6, 0xf5, 0xf5, 0xad, 0x38, 0x4b, 0x9b, 0xe5, 0x1a, 0x5b, 0x50, 0xd4, 0x68, 0xa3, 0xe5, 0x45,
	0x69, 0x31, 0x67, 0x95, 0xd7, 0x6c, 0x0d, 0x96, 0xaa, 0xbd, 0x1f, 0x3c, 0x16, 0x68, 0xd6, 0x36,
	0x53, 0x4b, 0x5d, 0xc5, 0x85, 0x2e, 0xd3, 0xba, 0x88, 0xdd, 0x6a, 0x66, 0x57, 0x31, 0x99, 0x4f,
	0x02, 0x8e, 0x0d, 0x56, 0xc0, 0x26, 0xad, 0x68, 0x3f, 0x79, 0xbb, 0xd1, 0x57, 0x20, 0x2e, 0x5f,
	0xd2, 0x16, 0x12, 0x60, 0x5d, 0xb6, 0x09, 0x7c, 0x58, 0x60, 0x23, 0xaf, 0x44, 0x8f, 0xa1, 0x69,
	0xa1, 0x59, 0x20, 0x45, 0xc1, 0x30, 0x98, 0xf4, 0x92, 0x13, 0x0f, 0xe6, 0xb4, 0x11, 0x97, 0x25,
	0x42, 0x6d, 0x36, 0x62, 0xe8, 0x45, 0x0f, 0xe6, 0x24, 0xa5, 0xe8, 0xd6, 0x69, 0x05, 0x51, 0xc7,
	0x71, 0x37, 0xcb, 0x6b, 0x71, 0xda, 0x38, 0x7b, 0xc8, 0x17, 0x48, 0x1c, 0x75, 0x87, 0x9d, 0x49,
	0x2f, 0xe9, 0xef, 0xd8, 0x9c, 0x58, 0x4e, 0x45, 0x68, 0x38, 0xfa, 0x3f, 0x0c, 0x26, 0xfd, 0xd9,
	0x40, 0x15, 0x5a, 0x17, 0xe5, 0xf6, 0x87, 0x33, 0xbb, 0x52, 0xaf, 0xbb, 0xd0, 0x49, 0x68, 0x78,
	0x34, 0x16, 0xf2, 0x30, 0x31, 0x93, 0xae, 0x19, 0xe4, 0xb9, 0x08, 0x31, 0xdf, 0x66, 0x0d, 0x31,
	0x9f, 0xe5, 0xe2, 0xf8, 0xc9, 0x17, 0x27, 0xdf, 0x84, 0xd8, 0x1f, 0xc8, 0x1b, 0xf5, 0xb3, 0x41,
	0xf5, 0xab, 0x80, 0xc1, 0xf8, 0xef, 0x25, 0xff, 0xcd, 0xd1, 0xbf, 0xe7, 0xdb, 0xf7, 0xe9, 0x41,
	0xfb, 0xee, 0xe6, 0x0e, 0x75, 0x9c, 0x12, 0xfa, 0x37, 0x28, 0xa0, 0x8e, 0xb7, 0x1e, 0x94, 0x65,
	0x47, 0x8e, 0x3d, 0x7c, 0x07, 0x00, 0x00, 0xff, 0xff, 0x28, 0x3b, 0x0c, 0x2c, 0xd3, 0x01, 0x00,
	0x00,
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
	SaveResolv(ctx context.Context, in *SaveResolvRequest, opts ...grpc.CallOption) (*SaveResolvResponse, error)
}

type archiveClient struct {
	cc *grpc.ClientConn
}

func NewArchiveClient(cc *grpc.ClientConn) ArchiveClient {
	return &archiveClient{cc}
}

func (c *archiveClient) SaveResolv(ctx context.Context, in *SaveResolvRequest, opts ...grpc.CallOption) (*SaveResolvResponse, error) {
	out := new(SaveResolvResponse)
	err := c.cc.Invoke(ctx, "/luids.dnsutil.v1.Archive/SaveResolv", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArchiveServer is the server API for Archive service.
type ArchiveServer interface {
	SaveResolv(context.Context, *SaveResolvRequest) (*SaveResolvResponse, error)
}

// UnimplementedArchiveServer can be embedded to have forward compatible implementations.
type UnimplementedArchiveServer struct {
}

func (*UnimplementedArchiveServer) SaveResolv(ctx context.Context, req *SaveResolvRequest) (*SaveResolvResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveResolv not implemented")
}

func RegisterArchiveServer(s *grpc.Server, srv ArchiveServer) {
	s.RegisterService(&_Archive_serviceDesc, srv)
}

func _Archive_SaveResolv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveResolvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArchiveServer).SaveResolv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/luids.dnsutil.v1.Archive/SaveResolv",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArchiveServer).SaveResolv(ctx, req.(*SaveResolvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Archive_serviceDesc = grpc.ServiceDesc{
	ServiceName: "luids.dnsutil.v1.Archive",
	HandlerType: (*ArchiveServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveResolv",
			Handler:    _Archive_SaveResolv_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "archive.proto",
}
