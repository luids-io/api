// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. View LICENSE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: xlist.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

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

// Enum value maps for Resource.
var (
	Resource_name = map[int32]string{
		0: "IPV4",
		1: "IPV6",
		2: "DOMAIN",
		3: "MD5",
		4: "SHA1",
		5: "SHA256",
	}
	Resource_value = map[string]int32{
		"IPV4":   0,
		"IPV6":   1,
		"DOMAIN": 2,
		"MD5":    3,
		"SHA1":   4,
		"SHA256": 5,
	}
)

func (x Resource) Enum() *Resource {
	p := new(Resource)
	*p = x
	return p
}

func (x Resource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Resource) Descriptor() protoreflect.EnumDescriptor {
	return file_xlist_proto_enumTypes[0].Descriptor()
}

func (Resource) Type() protoreflect.EnumType {
	return &file_xlist_proto_enumTypes[0]
}

func (x Resource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Resource.Descriptor instead.
func (Resource) EnumDescriptor() ([]byte, []int) {
	return file_xlist_proto_rawDescGZIP(), []int{0}
}

type CheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Resource Resource `protobuf:"varint,2,opt,name=resource,proto3,enum=luids.xlist.v1.Resource" json:"resource,omitempty"`
}

func (x *CheckRequest) Reset() {
	*x = CheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xlist_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRequest) ProtoMessage() {}

func (x *CheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_xlist_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRequest.ProtoReflect.Descriptor instead.
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return file_xlist_proto_rawDescGZIP(), []int{0}
}

func (x *CheckRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CheckRequest) GetResource() Resource {
	if x != nil {
		return x.Resource
	}
	return Resource_IPV4
}

type CheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool   `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Reason string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	TTL    int32  `protobuf:"varint,3,opt,name=TTL,proto3" json:"TTL,omitempty"`
}

func (x *CheckResponse) Reset() {
	*x = CheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xlist_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResponse) ProtoMessage() {}

func (x *CheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_xlist_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResponse.ProtoReflect.Descriptor instead.
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return file_xlist_proto_rawDescGZIP(), []int{1}
}

func (x *CheckResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *CheckResponse) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *CheckResponse) GetTTL() int32 {
	if x != nil {
		return x.TTL
	}
	return 0
}

type ResourcesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resources []Resource `protobuf:"varint,1,rep,packed,name=resources,proto3,enum=luids.xlist.v1.Resource" json:"resources,omitempty"`
}

func (x *ResourcesResponse) Reset() {
	*x = ResourcesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xlist_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourcesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourcesResponse) ProtoMessage() {}

func (x *ResourcesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_xlist_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourcesResponse.ProtoReflect.Descriptor instead.
func (*ResourcesResponse) Descriptor() ([]byte, []int) {
	return file_xlist_proto_rawDescGZIP(), []int{2}
}

func (x *ResourcesResponse) GetResources() []Resource {
	if x != nil {
		return x.Resources
	}
	return nil
}

var File_xlist_proto protoreflect.FileDescriptor

var file_xlist_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x78, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6c,
	0x75, 0x69, 0x64, 0x73, 0x2e, 0x78, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x0c, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x34,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x18, 0x2e, 0x6c, 0x75, 0x69, 0x64, 0x73, 0x2e, 0x78, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x22, 0x51, 0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x54, 0x54, 0x4c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x54, 0x54, 0x4c, 0x22, 0x4b, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x09,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x18, 0x2e, 0x6c, 0x75, 0x69, 0x64, 0x73, 0x2e, 0x78, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2a, 0x49, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x08, 0x0a, 0x04, 0x49, 0x50, 0x56, 0x34, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x50,
	0x56, 0x36, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x10, 0x02,
	0x12, 0x07, 0x0a, 0x03, 0x4d, 0x44, 0x35, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x48, 0x41,
	0x31, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x48, 0x41, 0x32, 0x35, 0x36, 0x10, 0x05, 0x32,
	0x99, 0x01, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x46, 0x0a, 0x05, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x12, 0x1c, 0x2e, 0x6c, 0x75, 0x69, 0x64, 0x73, 0x2e, 0x78, 0x6c, 0x69, 0x73, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x6c, 0x75, 0x69, 0x64, 0x73, 0x2e, 0x78, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x48, 0x0a, 0x09, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x21, 0x2e, 0x6c, 0x75, 0x69, 0x64, 0x73, 0x2e, 0x78,
	0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x27, 0x5a, 0x25, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x75, 0x69, 0x64, 0x73, 0x2d,
	0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x78, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_xlist_proto_rawDescOnce sync.Once
	file_xlist_proto_rawDescData = file_xlist_proto_rawDesc
)

func file_xlist_proto_rawDescGZIP() []byte {
	file_xlist_proto_rawDescOnce.Do(func() {
		file_xlist_proto_rawDescData = protoimpl.X.CompressGZIP(file_xlist_proto_rawDescData)
	})
	return file_xlist_proto_rawDescData
}

var file_xlist_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_xlist_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_xlist_proto_goTypes = []interface{}{
	(Resource)(0),             // 0: luids.xlist.v1.Resource
	(*CheckRequest)(nil),      // 1: luids.xlist.v1.CheckRequest
	(*CheckResponse)(nil),     // 2: luids.xlist.v1.CheckResponse
	(*ResourcesResponse)(nil), // 3: luids.xlist.v1.ResourcesResponse
	(*empty.Empty)(nil),       // 4: google.protobuf.Empty
}
var file_xlist_proto_depIdxs = []int32{
	0, // 0: luids.xlist.v1.CheckRequest.resource:type_name -> luids.xlist.v1.Resource
	0, // 1: luids.xlist.v1.ResourcesResponse.resources:type_name -> luids.xlist.v1.Resource
	1, // 2: luids.xlist.v1.Check.Check:input_type -> luids.xlist.v1.CheckRequest
	4, // 3: luids.xlist.v1.Check.Resources:input_type -> google.protobuf.Empty
	2, // 4: luids.xlist.v1.Check.Check:output_type -> luids.xlist.v1.CheckResponse
	3, // 5: luids.xlist.v1.Check.Resources:output_type -> luids.xlist.v1.ResourcesResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_xlist_proto_init() }
func file_xlist_proto_init() {
	if File_xlist_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_xlist_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_xlist_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_xlist_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourcesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_xlist_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_xlist_proto_goTypes,
		DependencyIndexes: file_xlist_proto_depIdxs,
		EnumInfos:         file_xlist_proto_enumTypes,
		MessageInfos:      file_xlist_proto_msgTypes,
	}.Build()
	File_xlist_proto = out.File
	file_xlist_proto_rawDesc = nil
	file_xlist_proto_goTypes = nil
	file_xlist_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CheckClient is the client API for Check service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CheckClient interface {
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
	Resources(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ResourcesResponse, error)
}

type checkClient struct {
	cc grpc.ClientConnInterface
}

func NewCheckClient(cc grpc.ClientConnInterface) CheckClient {
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

// CheckServer is the server API for Check service.
type CheckServer interface {
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
	Resources(context.Context, *empty.Empty) (*ResourcesResponse, error)
}

// UnimplementedCheckServer can be embedded to have forward compatible implementations.
type UnimplementedCheckServer struct {
}

func (*UnimplementedCheckServer) Check(context.Context, *CheckRequest) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (*UnimplementedCheckServer) Resources(context.Context, *empty.Empty) (*ResourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Resources not implemented")
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "xlist.proto",
}
