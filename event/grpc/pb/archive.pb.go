// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. View LICENSE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: github.com/luids-io/api/schemas/event/archive.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

type SaveEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type        EventType            `protobuf:"varint,2,opt,name=type,proto3,enum=luids.event.v1.EventType" json:"type,omitempty"`
	Code        int32                `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
	Level       EventLevel           `protobuf:"varint,4,opt,name=level,proto3,enum=luids.event.v1.EventLevel" json:"level,omitempty"`
	CreatedTs   *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_ts,json=createdTs,proto3" json:"created_ts,omitempty"`
	ReceivedTs  *timestamp.Timestamp `protobuf:"bytes,6,opt,name=received_ts,json=receivedTs,proto3" json:"received_ts,omitempty"`
	Source      *EventSource         `protobuf:"bytes,7,opt,name=source,proto3" json:"source,omitempty"`
	Processors  []*ProcessInfo       `protobuf:"bytes,8,rep,name=processors,proto3" json:"processors,omitempty"`
	Data        *EventData           `protobuf:"bytes,9,opt,name=data,proto3" json:"data,omitempty"`
	Duplicates  int32                `protobuf:"varint,10,opt,name=duplicates,proto3" json:"duplicates,omitempty"`
	Codename    string               `protobuf:"bytes,11,opt,name=codename,proto3" json:"codename,omitempty"`
	Description string               `protobuf:"bytes,12,opt,name=description,proto3" json:"description,omitempty"`
	Tags        []string             `protobuf:"bytes,13,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *SaveEventRequest) Reset() {
	*x = SaveEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_luids_io_api_schemas_event_archive_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveEventRequest) ProtoMessage() {}

func (x *SaveEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_luids_io_api_schemas_event_archive_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveEventRequest.ProtoReflect.Descriptor instead.
func (*SaveEventRequest) Descriptor() ([]byte, []int) {
	return file_github_com_luids_io_api_schemas_event_archive_proto_rawDescGZIP(), []int{0}
}

func (x *SaveEventRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SaveEventRequest) GetType() EventType {
	if x != nil {
		return x.Type
	}
	return EventType_UNDEFINED
}

func (x *SaveEventRequest) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SaveEventRequest) GetLevel() EventLevel {
	if x != nil {
		return x.Level
	}
	return EventLevel_INFO
}

func (x *SaveEventRequest) GetCreatedTs() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedTs
	}
	return nil
}

func (x *SaveEventRequest) GetReceivedTs() *timestamp.Timestamp {
	if x != nil {
		return x.ReceivedTs
	}
	return nil
}

func (x *SaveEventRequest) GetSource() *EventSource {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *SaveEventRequest) GetProcessors() []*ProcessInfo {
	if x != nil {
		return x.Processors
	}
	return nil
}

func (x *SaveEventRequest) GetData() *EventData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SaveEventRequest) GetDuplicates() int32 {
	if x != nil {
		return x.Duplicates
	}
	return 0
}

func (x *SaveEventRequest) GetCodename() string {
	if x != nil {
		return x.Codename
	}
	return ""
}

func (x *SaveEventRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SaveEventRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type SaveEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StorageID string `protobuf:"bytes,1,opt,name=storageID,proto3" json:"storageID,omitempty"`
}

func (x *SaveEventResponse) Reset() {
	*x = SaveEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_luids_io_api_schemas_event_archive_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveEventResponse) ProtoMessage() {}

func (x *SaveEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_luids_io_api_schemas_event_archive_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveEventResponse.ProtoReflect.Descriptor instead.
func (*SaveEventResponse) Descriptor() ([]byte, []int) {
	return file_github_com_luids_io_api_schemas_event_archive_proto_rawDescGZIP(), []int{1}
}

func (x *SaveEventResponse) GetStorageID() string {
	if x != nil {
		return x.StorageID
	}
	return ""
}

var File_github_com_luids_io_api_schemas_event_archive_proto protoreflect.FileDescriptor

var file_github_com_luids_io_api_schemas_event_archive_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x75, 0x69,
	0x64, 0x73, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6c, 0x75, 0x69, 0x64, 0x73, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6c, 0x75, 0x69, 0x64, 0x73, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x04, 0x0a, 0x10, 0x53,
	0x61, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x2d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e,
	0x6c, 0x75, 0x69, 0x64, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1a, 0x2e, 0x6c, 0x75, 0x69, 0x64, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x73, 0x12,
	0x3b, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x74, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x54, 0x73, 0x12, 0x33, 0x0a, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6c,
	0x75, 0x69, 0x64, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6c, 0x75, 0x69, 0x64, 0x73, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x73, 0x12, 0x2d,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6c,
	0x75, 0x69, 0x64, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a,
	0x0a, 0x64, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x64, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x6f, 0x64, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6f, 0x64, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22,
	0x31, 0x0a, 0x11, 0x53, 0x61, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x49, 0x44, 0x32, 0x5d, 0x0a, 0x07, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x12, 0x52, 0x0a,
	0x09, 0x53, 0x61, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x2e, 0x6c, 0x75, 0x69,
	0x64, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6c,
	0x75, 0x69, 0x64, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61,
	0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x75, 0x69, 0x64, 0x73, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_github_com_luids_io_api_schemas_event_archive_proto_rawDescOnce sync.Once
	file_github_com_luids_io_api_schemas_event_archive_proto_rawDescData = file_github_com_luids_io_api_schemas_event_archive_proto_rawDesc
)

func file_github_com_luids_io_api_schemas_event_archive_proto_rawDescGZIP() []byte {
	file_github_com_luids_io_api_schemas_event_archive_proto_rawDescOnce.Do(func() {
		file_github_com_luids_io_api_schemas_event_archive_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_luids_io_api_schemas_event_archive_proto_rawDescData)
	})
	return file_github_com_luids_io_api_schemas_event_archive_proto_rawDescData
}

var file_github_com_luids_io_api_schemas_event_archive_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_luids_io_api_schemas_event_archive_proto_goTypes = []interface{}{
	(*SaveEventRequest)(nil),    // 0: luids.event.v1.SaveEventRequest
	(*SaveEventResponse)(nil),   // 1: luids.event.v1.SaveEventResponse
	(EventType)(0),              // 2: luids.event.v1.EventType
	(EventLevel)(0),             // 3: luids.event.v1.EventLevel
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*EventSource)(nil),         // 5: luids.event.v1.EventSource
	(*ProcessInfo)(nil),         // 6: luids.event.v1.ProcessInfo
	(*EventData)(nil),           // 7: luids.event.v1.EventData
}
var file_github_com_luids_io_api_schemas_event_archive_proto_depIdxs = []int32{
	2, // 0: luids.event.v1.SaveEventRequest.type:type_name -> luids.event.v1.EventType
	3, // 1: luids.event.v1.SaveEventRequest.level:type_name -> luids.event.v1.EventLevel
	4, // 2: luids.event.v1.SaveEventRequest.created_ts:type_name -> google.protobuf.Timestamp
	4, // 3: luids.event.v1.SaveEventRequest.received_ts:type_name -> google.protobuf.Timestamp
	5, // 4: luids.event.v1.SaveEventRequest.source:type_name -> luids.event.v1.EventSource
	6, // 5: luids.event.v1.SaveEventRequest.processors:type_name -> luids.event.v1.ProcessInfo
	7, // 6: luids.event.v1.SaveEventRequest.data:type_name -> luids.event.v1.EventData
	0, // 7: luids.event.v1.Archive.SaveEvent:input_type -> luids.event.v1.SaveEventRequest
	1, // 8: luids.event.v1.Archive.SaveEvent:output_type -> luids.event.v1.SaveEventResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_github_com_luids_io_api_schemas_event_archive_proto_init() }
func file_github_com_luids_io_api_schemas_event_archive_proto_init() {
	if File_github_com_luids_io_api_schemas_event_archive_proto != nil {
		return
	}
	file_github_com_luids_io_api_schemas_event_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_github_com_luids_io_api_schemas_event_archive_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveEventRequest); i {
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
		file_github_com_luids_io_api_schemas_event_archive_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveEventResponse); i {
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
			RawDescriptor: file_github_com_luids_io_api_schemas_event_archive_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_luids_io_api_schemas_event_archive_proto_goTypes,
		DependencyIndexes: file_github_com_luids_io_api_schemas_event_archive_proto_depIdxs,
		MessageInfos:      file_github_com_luids_io_api_schemas_event_archive_proto_msgTypes,
	}.Build()
	File_github_com_luids_io_api_schemas_event_archive_proto = out.File
	file_github_com_luids_io_api_schemas_event_archive_proto_rawDesc = nil
	file_github_com_luids_io_api_schemas_event_archive_proto_goTypes = nil
	file_github_com_luids_io_api_schemas_event_archive_proto_depIdxs = nil
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
	SaveEvent(ctx context.Context, in *SaveEventRequest, opts ...grpc.CallOption) (*SaveEventResponse, error)
}

type archiveClient struct {
	cc grpc.ClientConnInterface
}

func NewArchiveClient(cc grpc.ClientConnInterface) ArchiveClient {
	return &archiveClient{cc}
}

func (c *archiveClient) SaveEvent(ctx context.Context, in *SaveEventRequest, opts ...grpc.CallOption) (*SaveEventResponse, error) {
	out := new(SaveEventResponse)
	err := c.cc.Invoke(ctx, "/luids.event.v1.Archive/SaveEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArchiveServer is the server API for Archive service.
type ArchiveServer interface {
	SaveEvent(context.Context, *SaveEventRequest) (*SaveEventResponse, error)
}

// UnimplementedArchiveServer can be embedded to have forward compatible implementations.
type UnimplementedArchiveServer struct {
}

func (*UnimplementedArchiveServer) SaveEvent(context.Context, *SaveEventRequest) (*SaveEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveEvent not implemented")
}

func RegisterArchiveServer(s *grpc.Server, srv ArchiveServer) {
	s.RegisterService(&_Archive_serviceDesc, srv)
}

func _Archive_SaveEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArchiveServer).SaveEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/luids.event.v1.Archive/SaveEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArchiveServer).SaveEvent(ctx, req.(*SaveEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Archive_serviceDesc = grpc.ServiceDesc{
	ServiceName: "luids.event.v1.Archive",
	HandlerType: (*ArchiveServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveEvent",
			Handler:    _Archive_SaveEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/luids-io/api/schemas/event/archive.proto",
}
