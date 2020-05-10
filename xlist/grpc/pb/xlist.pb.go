// Code generated by protoc-gen-go. DO NOT EDIT.
// source: xlist.proto

package pb

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

//enums
type Resource int32

const (
	Resource_IPV4   Resource = 0
	Resource_IPV6   Resource = 1
	Resource_DOMAIN Resource = 2
	Resource_MD5    Resource = 3
	Resource_SHA1   Resource = 4
	Resource_SHA256 Resource = 5
)

var Resource_name = map[int32]string{
	0: "IPV4",
	1: "IPV6",
	2: "DOMAIN",
	3: "MD5",
	4: "SHA1",
	5: "SHA256",
}

var Resource_value = map[string]int32{
	"IPV4":   0,
	"IPV6":   1,
	"DOMAIN": 2,
	"MD5":    3,
	"SHA1":   4,
	"SHA256": 5,
}

func (x Resource) String() string {
	return proto.EnumName(Resource_name, int32(x))
}

func (Resource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b5c410e285fee728, []int{0}
}

type CheckRequest struct {
	Request              *Request `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckRequest) Reset()         { *m = CheckRequest{} }
func (m *CheckRequest) String() string { return proto.CompactTextString(m) }
func (*CheckRequest) ProtoMessage()    {}
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5c410e285fee728, []int{0}
}

func (m *CheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckRequest.Unmarshal(m, b)
}
func (m *CheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckRequest.Marshal(b, m, deterministic)
}
func (m *CheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckRequest.Merge(m, src)
}
func (m *CheckRequest) XXX_Size() int {
	return xxx_messageInfo_CheckRequest.Size(m)
}
func (m *CheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckRequest proto.InternalMessageInfo

func (m *CheckRequest) GetRequest() *Request {
	if m != nil {
		return m.Request
	}
	return nil
}

type CheckResponse struct {
	Response             *Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CheckResponse) Reset()         { *m = CheckResponse{} }
func (m *CheckResponse) String() string { return proto.CompactTextString(m) }
func (*CheckResponse) ProtoMessage()    {}
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5c410e285fee728, []int{1}
}

func (m *CheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckResponse.Unmarshal(m, b)
}
func (m *CheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckResponse.Marshal(b, m, deterministic)
}
func (m *CheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckResponse.Merge(m, src)
}
func (m *CheckResponse) XXX_Size() int {
	return xxx_messageInfo_CheckResponse.Size(m)
}
func (m *CheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckResponse proto.InternalMessageInfo

func (m *CheckResponse) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

type ResourcesResponse struct {
	Resources            []Resource `protobuf:"varint,1,rep,packed,name=resources,proto3,enum=luids.xlist.v1.Resource" json:"resources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ResourcesResponse) Reset()         { *m = ResourcesResponse{} }
func (m *ResourcesResponse) String() string { return proto.CompactTextString(m) }
func (*ResourcesResponse) ProtoMessage()    {}
func (*ResourcesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5c410e285fee728, []int{2}
}

func (m *ResourcesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourcesResponse.Unmarshal(m, b)
}
func (m *ResourcesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourcesResponse.Marshal(b, m, deterministic)
}
func (m *ResourcesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourcesResponse.Merge(m, src)
}
func (m *ResourcesResponse) XXX_Size() int {
	return xxx_messageInfo_ResourcesResponse.Size(m)
}
func (m *ResourcesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourcesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResourcesResponse proto.InternalMessageInfo

func (m *ResourcesResponse) GetResources() []Resource {
	if m != nil {
		return m.Resources
	}
	return nil
}

// NOTE: request and response are in separate messages because they
// will be useful in a future multichecking api
type Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Resource             Resource `protobuf:"varint,2,opt,name=resource,proto3,enum=luids.xlist.v1.Resource" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5c410e285fee728, []int{3}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Request) GetResource() Resource {
	if m != nil {
		return m.Resource
	}
	return Resource_IPV4
}

type Response struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Reason               string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	TTL                  int32    `protobuf:"varint,3,opt,name=TTL,json=tTL,proto3" json:"TTL,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5c410e285fee728, []int{4}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *Response) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *Response) GetTTL() int32 {
	if m != nil {
		return m.TTL
	}
	return 0
}

func init() {
	proto.RegisterEnum("luids.xlist.v1.Resource", Resource_name, Resource_value)
	proto.RegisterType((*CheckRequest)(nil), "luids.xlist.v1.CheckRequest")
	proto.RegisterType((*CheckResponse)(nil), "luids.xlist.v1.CheckResponse")
	proto.RegisterType((*ResourcesResponse)(nil), "luids.xlist.v1.ResourcesResponse")
	proto.RegisterType((*Request)(nil), "luids.xlist.v1.Request")
	proto.RegisterType((*Response)(nil), "luids.xlist.v1.Response")
}

func init() { proto.RegisterFile("xlist.proto", fileDescriptor_b5c410e285fee728) }

var fileDescriptor_b5c410e285fee728 = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xdf, 0x6f, 0xd3, 0x30,
	0x10, 0xc7, 0x9b, 0xa5, 0xbf, 0x72, 0x83, 0x29, 0xf8, 0x61, 0x44, 0x15, 0x48, 0x21, 0x12, 0xa2,
	0x42, 0xc2, 0x56, 0xbb, 0xad, 0xe2, 0x35, 0xb0, 0xa1, 0x56, 0x74, 0x30, 0xb9, 0x15, 0x0f, 0xbc,
	0x25, 0xc1, 0x64, 0x11, 0x69, 0x6c, 0xe2, 0x04, 0xb1, 0xbf, 0x93, 0x7f, 0x08, 0xd5, 0x76, 0x42,
	0xf9, 0x11, 0x5e, 0xa2, 0xaf, 0xef, 0xbe, 0xf9, 0xdc, 0xd9, 0x77, 0x70, 0xfc, 0x3d, 0xcf, 0x64,
	0x85, 0x45, 0xc9, 0x2b, 0x8e, 0x4e, 0xf2, 0x3a, 0xfb, 0x24, 0xb1, 0x0e, 0x7d, 0x9b, 0x4d, 0xce,
	0xd2, 0xac, 0xba, 0xad, 0x63, 0x9c, 0xf0, 0x1d, 0x49, 0x79, 0x1e, 0x15, 0x29, 0x51, 0xc6, 0xb8,
	0xfe, 0x4c, 0x44, 0x75, 0x27, 0x98, 0x24, 0x6c, 0x27, 0xaa, 0x3b, 0xfd, 0xd5, 0x90, 0x20, 0x84,
	0x7b, 0xaf, 0x6f, 0x59, 0xf2, 0x85, 0xb2, 0xaf, 0x35, 0x93, 0x15, 0x9a, 0xc1, 0xa8, 0xd4, 0xd2,
	0xb3, 0x7c, 0x6b, 0x7a, 0x3c, 0x7f, 0x88, 0x7f, 0x2f, 0x83, 0x8d, 0x93, 0x36, 0xbe, 0xe0, 0x0a,
	0xee, 0x1b, 0x84, 0x14, 0xbc, 0x90, 0x0c, 0x9d, 0xc3, 0xb8, 0x34, 0xda, 0x40, 0xbc, 0xbf, 0x21,
	0x3a, 0x4f, 0x5b, 0x67, 0xf0, 0x16, 0x1e, 0x50, 0x26, 0x79, 0x5d, 0x26, 0x4c, 0xb6, 0xa8, 0x05,
	0x38, 0x65, 0x13, 0xf4, 0x2c, 0xdf, 0x9e, 0x9e, 0xfc, 0x93, 0xa5, 0x0c, 0xf4, 0x97, 0x35, 0xd8,
	0xc0, 0xa8, 0xb9, 0x11, 0x82, 0x7e, 0x11, 0xed, 0x74, 0x27, 0x0e, 0x55, 0xda, 0x74, 0xa8, 0xbc,
	0xde, 0x91, 0x6f, 0xfd, 0x97, 0xda, 0x3a, 0x83, 0x35, 0x8c, 0xdb, 0xc6, 0x4e, 0x61, 0x58, 0x32,
	0x59, 0xe7, 0xfa, 0x99, 0xc6, 0xd4, 0x9c, 0x74, 0x3c, 0x92, 0xbc, 0x50, 0x5c, 0x87, 0x9a, 0x13,
	0x72, 0xc1, 0xde, 0x6e, 0xd7, 0x9e, 0xed, 0x5b, 0xd3, 0x01, 0xb5, 0xab, 0xed, 0xfa, 0xf9, 0x4a,
	0xd1, 0x14, 0x19, 0x8d, 0xa1, 0xbf, 0xba, 0xf9, 0x70, 0xee, 0xf6, 0x8c, 0x5a, 0xb8, 0x16, 0x02,
	0x18, 0x5e, 0xbe, 0xbf, 0x0e, 0x57, 0xef, 0xdc, 0x23, 0x34, 0x02, 0xfb, 0xfa, 0xf2, 0xc2, 0xb5,
	0xf7, 0xe9, 0xcd, 0x32, 0x9c, 0xb9, 0xfd, 0x7d, 0x7a, 0xb3, 0x0c, 0xe7, 0x17, 0x0b, 0x77, 0x30,
	0xff, 0x61, 0xc1, 0x40, 0x8d, 0x00, 0xbd, 0x69, 0xc4, 0xa3, 0x3f, 0xef, 0x73, 0x38, 0xe5, 0xc9,
	0xe3, 0x8e, 0xac, 0x19, 0x45, 0x0f, 0x2d, 0xc1, 0x69, 0x87, 0x81, 0x4e, 0x71, 0xca, 0x79, 0x9a,
	0x33, 0xdc, 0xac, 0x13, 0xbe, 0xda, 0x6f, 0xd0, 0xe4, 0x49, 0xd7, 0x9b, 0xc9, 0x03, 0xd2, 0x4b,
	0xe8, 0xdf, 0x64, 0x45, 0xda, 0x09, 0xe9, 0x88, 0x07, 0xbd, 0x57, 0xcf, 0x3e, 0x3e, 0x3d, 0xd8,
	0x68, 0x55, 0xea, 0x45, 0xc6, 0x49, 0x24, 0x32, 0xa2, 0x2a, 0x92, 0xb4, 0x14, 0x09, 0x11, 0x71,
	0x3c, 0x54, 0xbf, 0x9e, 0xfd, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x61, 0xb9, 0x1f, 0xd9, 0x1e, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CheckClient is the client API for Check service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CheckClient interface {
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
	Resources(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ResourcesResponse, error)
	Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
}

type checkClient struct {
	cc *grpc.ClientConn
}

func NewCheckClient(cc *grpc.ClientConn) CheckClient {
	return &checkClient{cc}
}

func (c *checkClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/luids.xlist.v1.Check/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkClient) Resources(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ResourcesResponse, error) {
	out := new(ResourcesResponse)
	err := c.cc.Invoke(ctx, "/luids.xlist.v1.Check/Resources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkClient) Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/luids.xlist.v1.Check/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckServer is the server API for Check service.
type CheckServer interface {
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
	Resources(context.Context, *empty.Empty) (*ResourcesResponse, error)
	Ping(context.Context, *empty.Empty) (*empty.Empty, error)
}

// UnimplementedCheckServer can be embedded to have forward compatible implementations.
type UnimplementedCheckServer struct {
}

func (*UnimplementedCheckServer) Check(ctx context.Context, req *CheckRequest) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (*UnimplementedCheckServer) Resources(ctx context.Context, req *empty.Empty) (*ResourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Resources not implemented")
}
func (*UnimplementedCheckServer) Ping(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterCheckServer(s *grpc.Server, srv CheckServer) {
	s.RegisterService(&_Check_serviceDesc, srv)
}

func _Check_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/luids.xlist.v1.Check/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Check_Resources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckServer).Resources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/luids.xlist.v1.Check/Resources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckServer).Resources(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Check_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/luids.xlist.v1.Check/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckServer).Ping(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Check_serviceDesc = grpc.ServiceDesc{
	ServiceName: "luids.xlist.v1.Check",
	HandlerType: (*CheckServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Check_Check_Handler,
		},
		{
			MethodName: "Resources",
			Handler:    _Check_Resources_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Check_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "xlist.proto",
}