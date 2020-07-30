// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. View LICENSE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: github.com/luids-io/api/schemas/dnsutil/resolv.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type ResolvCollectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientIp    string   `protobuf:"bytes,1,opt,name=client_ip,json=clientIp,proto3" json:"client_ip,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ResolvedIps []string `protobuf:"bytes,3,rep,name=resolved_ips,json=resolvedIps,proto3" json:"resolved_ips,omitempty"`
}

func (x *ResolvCollectRequest) Reset() {
	*x = ResolvCollectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_luids_io_api_schemas_dnsutil_resolv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolvCollectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolvCollectRequest) ProtoMessage() {}

func (x *ResolvCollectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_luids_io_api_schemas_dnsutil_resolv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolvCollectRequest.ProtoReflect.Descriptor instead.
func (*ResolvCollectRequest) Descriptor() ([]byte, []int) {
	return file_github_com_luids_io_api_schemas_dnsutil_resolv_proto_rawDescGZIP(), []int{0}
}

func (x *ResolvCollectRequest) GetClientIp() string {
	if x != nil {
		return x.ClientIp
	}
	return ""
}

func (x *ResolvCollectRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResolvCollectRequest) GetResolvedIps() []string {
	if x != nil {
		return x.ResolvedIps
	}
	return nil
}

type ResolvCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientIp   string `protobuf:"bytes,1,opt,name=client_ip,json=clientIp,proto3" json:"client_ip,omitempty"`
	ResolvedIp string `protobuf:"bytes,2,opt,name=resolved_ip,json=resolvedIp,proto3" json:"resolved_ip,omitempty"`
	Name       string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ResolvCheckRequest) Reset() {
	*x = ResolvCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_luids_io_api_schemas_dnsutil_resolv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolvCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolvCheckRequest) ProtoMessage() {}

func (x *ResolvCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_luids_io_api_schemas_dnsutil_resolv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolvCheckRequest.ProtoReflect.Descriptor instead.
func (*ResolvCheckRequest) Descriptor() ([]byte, []int) {
	return file_github_com_luids_io_api_schemas_dnsutil_resolv_proto_rawDescGZIP(), []int{1}
}

func (x *ResolvCheckRequest) GetClientIp() string {
	if x != nil {
		return x.ClientIp
	}
	return ""
}

func (x *ResolvCheckRequest) GetResolvedIp() string {
	if x != nil {
		return x.ResolvedIp
	}
	return ""
}

func (x *ResolvCheckRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ResolvCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result  bool                 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	LastTs  *timestamp.Timestamp `protobuf:"bytes,2,opt,name=last_ts,json=lastTs,proto3" json:"last_ts,omitempty"`
	StoreTs *timestamp.Timestamp `protobuf:"bytes,3,opt,name=store_ts,json=storeTs,proto3" json:"store_ts,omitempty"`
}

func (x *ResolvCheckResponse) Reset() {
	*x = ResolvCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_luids_io_api_schemas_dnsutil_resolv_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolvCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolvCheckResponse) ProtoMessage() {}

func (x *ResolvCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_luids_io_api_schemas_dnsutil_resolv_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolvCheckResponse.ProtoReflect.Descriptor instead.
func (*ResolvCheckResponse) Descriptor() ([]byte, []int) {
	return file_github_com_luids_io_api_schemas_dnsutil_resolv_proto_rawDescGZIP(), []int{2}
}

func (x *ResolvCheckResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *ResolvCheckResponse) GetLastTs() *timestamp.Timestamp {
	if x != nil {
		return x.LastTs
	}
	return nil
}

func (x *ResolvCheckResponse) GetStoreTs() *timestamp.Timestamp {
	if x != nil {
		return x.StoreTs
	}
	return nil
}

var File_github_com_luids_io_api_schemas_dnsutil_resolv_proto protoreflect.FileDescriptor

var file_github_com_luids_io_api_schemas_dnsutil_resolv_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x75, 0x69,
	0x64, 0x73, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x73, 0x2f, 0x64, 0x6e, 0x73, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6c, 0x75, 0x69, 0x64, 0x73, 0x2e, 0x64, 0x6e,
	0x73, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x5f, 0x69, 0x70, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x49,
	0x70, 0x73, 0x22, 0x66, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x64, 0x5f, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x64, 0x49, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x99, 0x01, 0x0a, 0x13, 0x52,
	0x65, 0x73, 0x6f, 0x6c, 0x76, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x73, 0x12,
	0x35, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x54, 0x73, 0x32, 0x5c, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x4b, 0x0a, 0x07, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x12, 0x26, 0x2e, 0x6c, 0x75, 0x69, 0x64, 0x73, 0x2e, 0x64, 0x6e, 0x73, 0x75, 0x74,
	0x69, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x32, 0x65, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x12, 0x56, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x24, 0x2e, 0x6c,
	0x75, 0x69, 0x64, 0x73, 0x2e, 0x64, 0x6e, 0x73, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6c, 0x75, 0x69, 0x64, 0x73, 0x2e, 0x64, 0x6e, 0x73, 0x75, 0x74,
	0x69, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x29, 0x5a, 0x27, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x75, 0x69, 0x64, 0x73, 0x2d,
	0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x6e, 0x73, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_luids_io_api_schemas_dnsutil_resolv_proto_rawDescOnce sync.Once
	file_github_com_luids_io_api_schemas_dnsutil_resolv_proto_rawDescData = file_github_com_luids_io_api_schemas_dnsutil_resolv_proto_rawDesc
)

func file_github_com_luids_io_api_schemas_dnsutil_resolv_proto_rawDescGZIP() []byte {
	file_github_com_luids_io_api_schemas_dnsutil_resolv_proto_rawDescOnce.Do(func() {
		file_github_com_luids_io_api_schemas_dnsutil_resolv_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_luids_io_api_schemas_dnsutil_resolv_proto_rawDescData)
	})
	return file_github_com_luids_io_api_schemas_dnsutil_resolv_proto_rawDescData
}

var file_github_com_luids_io_api_schemas_dnsutil_resolv_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_luids_io_api_schemas_dnsutil_resolv_proto_goTypes = []interface{}{
	(*ResolvCollectRequest)(nil), // 0: luids.dnsutil.v1.ResolvCollectRequest
	(*ResolvCheckRequest)(nil),   // 1: luids.dnsutil.v1.ResolvCheckRequest
	(*ResolvCheckResponse)(nil),  // 2: luids.dnsutil.v1.ResolvCheckResponse
	(*timestamp.Timestamp)(nil),  // 3: google.protobuf.Timestamp
	(*empty.Empty)(nil),          // 4: google.protobuf.Empty
}
var file_github_com_luids_io_api_schemas_dnsutil_resolv_proto_depIdxs = []int32{
	3, // 0: luids.dnsutil.v1.ResolvCheckResponse.last_ts:type_name -> google.protobuf.Timestamp
	3, // 1: luids.dnsutil.v1.ResolvCheckResponse.store_ts:type_name -> google.protobuf.Timestamp
	0, // 2: luids.dnsutil.v1.ResolvCollect.Collect:input_type -> luids.dnsutil.v1.ResolvCollectRequest
	1, // 3: luids.dnsutil.v1.ResolvCheck.Check:input_type -> luids.dnsutil.v1.ResolvCheckRequest
	4, // 4: luids.dnsutil.v1.ResolvCollect.Collect:output_type -> google.protobuf.Empty
	2, // 5: luids.dnsutil.v1.ResolvCheck.Check:output_type -> luids.dnsutil.v1.ResolvCheckResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_github_com_luids_io_api_schemas_dnsutil_resolv_proto_init() }
func file_github_com_luids_io_api_schemas_dnsutil_resolv_proto_init() {
	if File_github_com_luids_io_api_schemas_dnsutil_resolv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_luids_io_api_schemas_dnsutil_resolv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolvCollectRequest); i {
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
		file_github_com_luids_io_api_schemas_dnsutil_resolv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolvCheckRequest); i {
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
		file_github_com_luids_io_api_schemas_dnsutil_resolv_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolvCheckResponse); i {
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
			RawDescriptor: file_github_com_luids_io_api_schemas_dnsutil_resolv_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_github_com_luids_io_api_schemas_dnsutil_resolv_proto_goTypes,
		DependencyIndexes: file_github_com_luids_io_api_schemas_dnsutil_resolv_proto_depIdxs,
		MessageInfos:      file_github_com_luids_io_api_schemas_dnsutil_resolv_proto_msgTypes,
	}.Build()
	File_github_com_luids_io_api_schemas_dnsutil_resolv_proto = out.File
	file_github_com_luids_io_api_schemas_dnsutil_resolv_proto_rawDesc = nil
	file_github_com_luids_io_api_schemas_dnsutil_resolv_proto_goTypes = nil
	file_github_com_luids_io_api_schemas_dnsutil_resolv_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ResolvCollectClient is the client API for ResolvCollect service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResolvCollectClient interface {
	Collect(ctx context.Context, in *ResolvCollectRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type resolvCollectClient struct {
	cc grpc.ClientConnInterface
}

func NewResolvCollectClient(cc grpc.ClientConnInterface) ResolvCollectClient {
	return &resolvCollectClient{cc}
}

func (c *resolvCollectClient) Collect(ctx context.Context, in *ResolvCollectRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/luids.dnsutil.v1.ResolvCollect/Collect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResolvCollectServer is the server API for ResolvCollect service.
type ResolvCollectServer interface {
	Collect(context.Context, *ResolvCollectRequest) (*empty.Empty, error)
}

// UnimplementedResolvCollectServer can be embedded to have forward compatible implementations.
type UnimplementedResolvCollectServer struct {
}

func (*UnimplementedResolvCollectServer) Collect(context.Context, *ResolvCollectRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Collect not implemented")
}

func RegisterResolvCollectServer(s *grpc.Server, srv ResolvCollectServer) {
	s.RegisterService(&_ResolvCollect_serviceDesc, srv)
}

func _ResolvCollect_Collect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolvCollectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResolvCollectServer).Collect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/luids.dnsutil.v1.ResolvCollect/Collect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResolvCollectServer).Collect(ctx, req.(*ResolvCollectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResolvCollect_serviceDesc = grpc.ServiceDesc{
	ServiceName: "luids.dnsutil.v1.ResolvCollect",
	HandlerType: (*ResolvCollectServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Collect",
			Handler:    _ResolvCollect_Collect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/luids-io/api/schemas/dnsutil/resolv.proto",
}

// ResolvCheckClient is the client API for ResolvCheck service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResolvCheckClient interface {
	Check(ctx context.Context, in *ResolvCheckRequest, opts ...grpc.CallOption) (*ResolvCheckResponse, error)
}

type resolvCheckClient struct {
	cc grpc.ClientConnInterface
}

func NewResolvCheckClient(cc grpc.ClientConnInterface) ResolvCheckClient {
	return &resolvCheckClient{cc}
}

func (c *resolvCheckClient) Check(ctx context.Context, in *ResolvCheckRequest, opts ...grpc.CallOption) (*ResolvCheckResponse, error) {
	out := new(ResolvCheckResponse)
	err := c.cc.Invoke(ctx, "/luids.dnsutil.v1.ResolvCheck/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResolvCheckServer is the server API for ResolvCheck service.
type ResolvCheckServer interface {
	Check(context.Context, *ResolvCheckRequest) (*ResolvCheckResponse, error)
}

// UnimplementedResolvCheckServer can be embedded to have forward compatible implementations.
type UnimplementedResolvCheckServer struct {
}

func (*UnimplementedResolvCheckServer) Check(context.Context, *ResolvCheckRequest) (*ResolvCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}

func RegisterResolvCheckServer(s *grpc.Server, srv ResolvCheckServer) {
	s.RegisterService(&_ResolvCheck_serviceDesc, srv)
}

func _ResolvCheck_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolvCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResolvCheckServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/luids.dnsutil.v1.ResolvCheck/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResolvCheckServer).Check(ctx, req.(*ResolvCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResolvCheck_serviceDesc = grpc.ServiceDesc{
	ServiceName: "luids.dnsutil.v1.ResolvCheck",
	HandlerType: (*ResolvCheckServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _ResolvCheck_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/luids-io/api/schemas/dnsutil/resolv.proto",
}
