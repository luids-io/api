// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. View LICENSE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: archive.proto

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

type SaveConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Connection *ConnectionData `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
}

func (x *SaveConnectionRequest) Reset() {
	*x = SaveConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_archive_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveConnectionRequest) ProtoMessage() {}

func (x *SaveConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_archive_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveConnectionRequest.ProtoReflect.Descriptor instead.
func (*SaveConnectionRequest) Descriptor() ([]byte, []int) {
	return file_archive_proto_rawDescGZIP(), []int{0}
}

func (x *SaveConnectionRequest) GetConnection() *ConnectionData {
	if x != nil {
		return x.Connection
	}
	return nil
}

type SaveConnectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SaveConnectionResponse) Reset() {
	*x = SaveConnectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_archive_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveConnectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveConnectionResponse) ProtoMessage() {}

func (x *SaveConnectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_archive_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveConnectionResponse.ProtoReflect.Descriptor instead.
func (*SaveConnectionResponse) Descriptor() ([]byte, []int) {
	return file_archive_proto_rawDescGZIP(), []int{1}
}

func (x *SaveConnectionResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SaveCertificateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Certificate *CertificateData `protobuf:"bytes,1,opt,name=certificate,proto3" json:"certificate,omitempty"`
}

func (x *SaveCertificateRequest) Reset() {
	*x = SaveCertificateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_archive_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveCertificateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveCertificateRequest) ProtoMessage() {}

func (x *SaveCertificateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_archive_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveCertificateRequest.ProtoReflect.Descriptor instead.
func (*SaveCertificateRequest) Descriptor() ([]byte, []int) {
	return file_archive_proto_rawDescGZIP(), []int{2}
}

func (x *SaveCertificateRequest) GetCertificate() *CertificateData {
	if x != nil {
		return x.Certificate
	}
	return nil
}

type SaveCertificateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SaveCertificateResponse) Reset() {
	*x = SaveCertificateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_archive_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveCertificateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveCertificateResponse) ProtoMessage() {}

func (x *SaveCertificateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_archive_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveCertificateResponse.ProtoReflect.Descriptor instead.
func (*SaveCertificateResponse) Descriptor() ([]byte, []int) {
	return file_archive_proto_rawDescGZIP(), []int{3}
}

func (x *SaveCertificateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SaveRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Record *RecordData `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
}

func (x *SaveRecordRequest) Reset() {
	*x = SaveRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_archive_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveRecordRequest) ProtoMessage() {}

func (x *SaveRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_archive_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveRecordRequest.ProtoReflect.Descriptor instead.
func (*SaveRecordRequest) Descriptor() ([]byte, []int) {
	return file_archive_proto_rawDescGZIP(), []int{4}
}

func (x *SaveRecordRequest) GetRecord() *RecordData {
	if x != nil {
		return x.Record
	}
	return nil
}

var File_archive_proto protoreflect.FileDescriptor

var file_archive_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x6c, 0x75, 0x69, 0x64, 0x73, 0x2e, 0x74, 0x6c, 0x73, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x76,
	0x31, 0x1a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x75,
	0x69, 0x64, 0x73, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x73, 0x2f, 0x74, 0x6c, 0x73, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x70, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x15,
	0x53, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6c, 0x75, 0x69, 0x64,
	0x73, 0x2e, 0x74, 0x6c, 0x73, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x28, 0x0a, 0x16, 0x53, 0x61, 0x76, 0x65, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x5d, 0x0a, 0x16, 0x53, 0x61, 0x76, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x0b, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x6c, 0x75, 0x69, 0x64, 0x73, 0x2e, 0x74, 0x6c, 0x73, 0x75, 0x74, 0x69, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x22, 0x29, 0x0a, 0x17, 0x53, 0x61, 0x76, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x49, 0x0a, 0x11, 0x53,
	0x61, 0x76, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x34, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x6c, 0x75, 0x69, 0x64, 0x73, 0x2e, 0x74, 0x6c, 0x73, 0x75, 0x74, 0x69, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x06,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x32, 0xac, 0x02, 0x0a, 0x07, 0x41, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x12, 0x65, 0x0a, 0x0e, 0x53, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x6c, 0x75, 0x69, 0x64, 0x73, 0x2e, 0x74, 0x6c, 0x73,
	0x75, 0x74, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e,
	0x6c, 0x75, 0x69, 0x64, 0x73, 0x2e, 0x74, 0x6c, 0x73, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x0f, 0x53, 0x61, 0x76,
	0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x6c,
	0x75, 0x69, 0x64, 0x73, 0x2e, 0x74, 0x6c, 0x73, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x61, 0x76, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6c, 0x75, 0x69, 0x64, 0x73, 0x2e, 0x74,
	0x6c, 0x73, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x12, 0x23, 0x2e, 0x6c, 0x75, 0x69, 0x64, 0x73, 0x2e, 0x74, 0x6c, 0x73,
	0x75, 0x74, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x28, 0x01, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x75, 0x69, 0x64, 0x73, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x74, 0x6c, 0x73, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_archive_proto_rawDescOnce sync.Once
	file_archive_proto_rawDescData = file_archive_proto_rawDesc
)

func file_archive_proto_rawDescGZIP() []byte {
	file_archive_proto_rawDescOnce.Do(func() {
		file_archive_proto_rawDescData = protoimpl.X.CompressGZIP(file_archive_proto_rawDescData)
	})
	return file_archive_proto_rawDescData
}

var file_archive_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_archive_proto_goTypes = []interface{}{
	(*SaveConnectionRequest)(nil),   // 0: luids.tlsutil.v1.SaveConnectionRequest
	(*SaveConnectionResponse)(nil),  // 1: luids.tlsutil.v1.SaveConnectionResponse
	(*SaveCertificateRequest)(nil),  // 2: luids.tlsutil.v1.SaveCertificateRequest
	(*SaveCertificateResponse)(nil), // 3: luids.tlsutil.v1.SaveCertificateResponse
	(*SaveRecordRequest)(nil),       // 4: luids.tlsutil.v1.SaveRecordRequest
	(*ConnectionData)(nil),          // 5: luids.tlsutil.v1.ConnectionData
	(*CertificateData)(nil),         // 6: luids.tlsutil.v1.CertificateData
	(*RecordData)(nil),              // 7: luids.tlsutil.v1.RecordData
	(*empty.Empty)(nil),             // 8: google.protobuf.Empty
}
var file_archive_proto_depIdxs = []int32{
	5, // 0: luids.tlsutil.v1.SaveConnectionRequest.connection:type_name -> luids.tlsutil.v1.ConnectionData
	6, // 1: luids.tlsutil.v1.SaveCertificateRequest.certificate:type_name -> luids.tlsutil.v1.CertificateData
	7, // 2: luids.tlsutil.v1.SaveRecordRequest.record:type_name -> luids.tlsutil.v1.RecordData
	0, // 3: luids.tlsutil.v1.Archive.SaveConnection:input_type -> luids.tlsutil.v1.SaveConnectionRequest
	2, // 4: luids.tlsutil.v1.Archive.SaveCertificate:input_type -> luids.tlsutil.v1.SaveCertificateRequest
	4, // 5: luids.tlsutil.v1.Archive.StreamRecords:input_type -> luids.tlsutil.v1.SaveRecordRequest
	1, // 6: luids.tlsutil.v1.Archive.SaveConnection:output_type -> luids.tlsutil.v1.SaveConnectionResponse
	3, // 7: luids.tlsutil.v1.Archive.SaveCertificate:output_type -> luids.tlsutil.v1.SaveCertificateResponse
	8, // 8: luids.tlsutil.v1.Archive.StreamRecords:output_type -> google.protobuf.Empty
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_archive_proto_init() }
func file_archive_proto_init() {
	if File_archive_proto != nil {
		return
	}
	file_github_com_luids_io_api_schemas_tlsutil_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_archive_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveConnectionRequest); i {
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
		file_archive_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveConnectionResponse); i {
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
		file_archive_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveCertificateRequest); i {
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
		file_archive_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveCertificateResponse); i {
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
		file_archive_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveRecordRequest); i {
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
			RawDescriptor: file_archive_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_archive_proto_goTypes,
		DependencyIndexes: file_archive_proto_depIdxs,
		MessageInfos:      file_archive_proto_msgTypes,
	}.Build()
	File_archive_proto = out.File
	file_archive_proto_rawDesc = nil
	file_archive_proto_goTypes = nil
	file_archive_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ArchiveClient is the client API for Archive service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ArchiveClient interface {
	SaveConnection(ctx context.Context, in *SaveConnectionRequest, opts ...grpc.CallOption) (*SaveConnectionResponse, error)
	SaveCertificate(ctx context.Context, in *SaveCertificateRequest, opts ...grpc.CallOption) (*SaveCertificateResponse, error)
	StreamRecords(ctx context.Context, opts ...grpc.CallOption) (Archive_StreamRecordsClient, error)
}

type archiveClient struct {
	cc grpc.ClientConnInterface
}

func NewArchiveClient(cc grpc.ClientConnInterface) ArchiveClient {
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

func (*UnimplementedArchiveServer) SaveConnection(context.Context, *SaveConnectionRequest) (*SaveConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveConnection not implemented")
}
func (*UnimplementedArchiveServer) SaveCertificate(context.Context, *SaveCertificateRequest) (*SaveCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveCertificate not implemented")
}
func (*UnimplementedArchiveServer) StreamRecords(Archive_StreamRecordsServer) error {
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
