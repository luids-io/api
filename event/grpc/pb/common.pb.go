// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type EventType int32

const (
	EventType_UNDEFINED EventType = 0
	EventType_SECURITY  EventType = 1
)

var EventType_name = map[int32]string{
	0: "UNDEFINED",
	1: "SECURITY",
}

var EventType_value = map[string]int32{
	"UNDEFINED": 0,
	"SECURITY":  1,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}

func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

type EventLevel int32

const (
	EventLevel_INFO     EventLevel = 0
	EventLevel_LOW      EventLevel = 1
	EventLevel_MEDIUM   EventLevel = 2
	EventLevel_HIGH     EventLevel = 3
	EventLevel_CRITICAL EventLevel = 4
)

var EventLevel_name = map[int32]string{
	0: "INFO",
	1: "LOW",
	2: "MEDIUM",
	3: "HIGH",
	4: "CRITICAL",
}

var EventLevel_value = map[string]int32{
	"INFO":     0,
	"LOW":      1,
	"MEDIUM":   2,
	"HIGH":     3,
	"CRITICAL": 4,
}

func (x EventLevel) String() string {
	return proto.EnumName(EventLevel_name, int32(x))
}

func (EventLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

type EventData_DataEnc int32

const (
	EventData_NODATA EventData_DataEnc = 0
	EventData_JSON   EventData_DataEnc = 1
)

var EventData_DataEnc_name = map[int32]string{
	0: "NODATA",
	1: "JSON",
}

var EventData_DataEnc_value = map[string]int32{
	"NODATA": 0,
	"JSON":   1,
}

func (x EventData_DataEnc) String() string {
	return proto.EnumName(EventData_DataEnc_name, int32(x))
}

func (EventData_DataEnc) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2, 0}
}

type EventSource struct {
	Hostname             string   `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Program              string   `protobuf:"bytes,2,opt,name=program,proto3" json:"program,omitempty"`
	Instance             string   `protobuf:"bytes,3,opt,name=instance,proto3" json:"instance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventSource) Reset()         { *m = EventSource{} }
func (m *EventSource) String() string { return proto.CompactTextString(m) }
func (*EventSource) ProtoMessage()    {}
func (*EventSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *EventSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventSource.Unmarshal(m, b)
}
func (m *EventSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventSource.Marshal(b, m, deterministic)
}
func (m *EventSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSource.Merge(m, src)
}
func (m *EventSource) XXX_Size() int {
	return xxx_messageInfo_EventSource.Size(m)
}
func (m *EventSource) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSource.DiscardUnknown(m)
}

var xxx_messageInfo_EventSource proto.InternalMessageInfo

func (m *EventSource) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *EventSource) GetProgram() string {
	if m != nil {
		return m.Program
	}
	return ""
}

func (m *EventSource) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

type ProcessInfo struct {
	ReceivedTs           *timestamp.Timestamp `protobuf:"bytes,1,opt,name=received_ts,json=receivedTs,proto3" json:"received_ts,omitempty"`
	Processor            *EventSource         `protobuf:"bytes,2,opt,name=processor,proto3" json:"processor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ProcessInfo) Reset()         { *m = ProcessInfo{} }
func (m *ProcessInfo) String() string { return proto.CompactTextString(m) }
func (*ProcessInfo) ProtoMessage()    {}
func (*ProcessInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *ProcessInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessInfo.Unmarshal(m, b)
}
func (m *ProcessInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessInfo.Marshal(b, m, deterministic)
}
func (m *ProcessInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessInfo.Merge(m, src)
}
func (m *ProcessInfo) XXX_Size() int {
	return xxx_messageInfo_ProcessInfo.Size(m)
}
func (m *ProcessInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessInfo proto.InternalMessageInfo

func (m *ProcessInfo) GetReceivedTs() *timestamp.Timestamp {
	if m != nil {
		return m.ReceivedTs
	}
	return nil
}

func (m *ProcessInfo) GetProcessor() *EventSource {
	if m != nil {
		return m.Processor
	}
	return nil
}

type EventData struct {
	DataEnc              EventData_DataEnc `protobuf:"varint,1,opt,name=dataEnc,proto3,enum=luids.event.v1.EventData_DataEnc" json:"dataEnc,omitempty"`
	Data                 []byte            `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *EventData) Reset()         { *m = EventData{} }
func (m *EventData) String() string { return proto.CompactTextString(m) }
func (*EventData) ProtoMessage()    {}
func (*EventData) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

func (m *EventData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventData.Unmarshal(m, b)
}
func (m *EventData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventData.Marshal(b, m, deterministic)
}
func (m *EventData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventData.Merge(m, src)
}
func (m *EventData) XXX_Size() int {
	return xxx_messageInfo_EventData.Size(m)
}
func (m *EventData) XXX_DiscardUnknown() {
	xxx_messageInfo_EventData.DiscardUnknown(m)
}

var xxx_messageInfo_EventData proto.InternalMessageInfo

func (m *EventData) GetDataEnc() EventData_DataEnc {
	if m != nil {
		return m.DataEnc
	}
	return EventData_NODATA
}

func (m *EventData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("luids.event.v1.EventType", EventType_name, EventType_value)
	proto.RegisterEnum("luids.event.v1.EventLevel", EventLevel_name, EventLevel_value)
	proto.RegisterEnum("luids.event.v1.EventData_DataEnc", EventData_DataEnc_name, EventData_DataEnc_value)
	proto.RegisterType((*EventSource)(nil), "luids.event.v1.EventSource")
	proto.RegisterType((*ProcessInfo)(nil), "luids.event.v1.ProcessInfo")
	proto.RegisterType((*EventData)(nil), "luids.event.v1.EventData")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0x6b, 0xdb, 0x30,
	0x14, 0xc6, 0xe3, 0x26, 0x34, 0xc9, 0x73, 0x56, 0x84, 0x4e, 0x21, 0x3b, 0x6c, 0x0b, 0x8c, 0x95,
	0xc2, 0x64, 0x96, 0x9d, 0x46, 0x4e, 0x59, 0xec, 0xae, 0x1e, 0xa9, 0x33, 0x1c, 0x87, 0xb1, 0x5d,
	0x8a, 0xa2, 0xa8, 0xae, 0x21, 0x96, 0x84, 0xa4, 0x04, 0x7a, 0xd8, 0x71, 0xff, 0xf7, 0x90, 0x52,
	0xaf, 0x19, 0xec, 0xf6, 0x1e, 0xef, 0xf7, 0xbe, 0xef, 0x7b, 0x42, 0x30, 0x60, 0xb2, 0xae, 0xa5,
	0x20, 0x4a, 0x4b, 0x2b, 0xf1, 0xc5, 0x6e, 0x5f, 0x6d, 0x0d, 0xe1, 0x07, 0x2e, 0x2c, 0x39, 0x7c,
	0x18, 0x4d, 0xcb, 0xca, 0x3e, 0xec, 0x37, 0x84, 0xc9, 0x3a, 0x2a, 0xe5, 0x8e, 0x8a, 0x32, 0xf2,
	0xe0, 0x66, 0x7f, 0x1f, 0x29, 0xfb, 0xa8, 0xb8, 0x89, 0x6c, 0x55, 0x73, 0x63, 0x69, 0xad, 0x9e,
	0xab, 0xa3, 0xd8, 0xf8, 0x0e, 0xc2, 0xc4, 0x09, 0xad, 0xe4, 0x5e, 0x33, 0x8e, 0x47, 0xd0, 0x7b,
	0x90, 0xc6, 0x0a, 0x5a, 0xf3, 0x61, 0xf0, 0x3a, 0xb8, 0xec, 0xe7, 0x7f, 0x7b, 0x3c, 0x84, 0xae,
	0xd2, 0xb2, 0xd4, 0xb4, 0x1e, 0x9e, 0xf9, 0x51, 0xd3, 0xba, 0xad, 0x4a, 0x18, 0x4b, 0x05, 0xe3,
	0xc3, 0xf6, 0x71, 0xab, 0xe9, 0xc7, 0xbf, 0x03, 0x08, 0xbf, 0x69, 0xc9, 0xb8, 0x31, 0xa9, 0xb8,
	0x97, 0x78, 0x0a, 0xa1, 0xe6, 0x8c, 0x57, 0x07, 0xbe, 0xbd, 0xb3, 0xc6, 0x9b, 0x84, 0x93, 0x11,
	0x29, 0xa5, 0x2c, 0x77, 0x9c, 0x34, 0xc1, 0x49, 0xd1, 0xe4, 0xcc, 0xa1, 0xc1, 0x0b, 0x83, 0x3f,
	0x41, 0x5f, 0x1d, 0xb5, 0xa4, 0xf6, 0x21, 0xc2, 0xc9, 0x4b, 0xf2, 0xef, 0x73, 0x90, 0x93, 0x73,
	0xf2, 0x67, 0x7a, 0xfc, 0x0b, 0xfa, 0x7e, 0x12, 0x53, 0x4b, 0xf1, 0x14, 0xba, 0x5b, 0x6a, 0x69,
	0x22, 0x98, 0x0f, 0x70, 0x31, 0x79, 0xf3, 0x5f, 0x15, 0xc7, 0x92, 0xf8, 0x08, 0xe6, 0xcd, 0x06,
	0xc6, 0xd0, 0x71, 0xa5, 0xf7, 0x1f, 0xe4, 0xbe, 0x1e, 0xbf, 0x82, 0xee, 0x13, 0x87, 0x01, 0xce,
	0xb3, 0x65, 0x3c, 0x2b, 0x66, 0xa8, 0x85, 0x7b, 0xd0, 0xf9, 0xba, 0x5a, 0x66, 0x28, 0xb8, 0xba,
	0x7c, 0xb2, 0x2f, 0x1e, 0x15, 0xc7, 0x2f, 0xa0, 0xbf, 0xce, 0xe2, 0xe4, 0x3a, 0xcd, 0x92, 0x18,
	0xb5, 0xf0, 0x00, 0x7a, 0xab, 0x64, 0xbe, 0xce, 0xd3, 0xe2, 0x07, 0x0a, 0xae, 0xe6, 0x00, 0x9e,
	0x5c, 0xf0, 0x03, 0xdf, 0x39, 0x85, 0x34, 0xbb, 0x5e, 0xa2, 0x16, 0xee, 0x42, 0x7b, 0xb1, 0xfc,
	0x8e, 0x02, 0x67, 0x70, 0x9b, 0xc4, 0xe9, 0xfa, 0x16, 0x9d, 0xb9, 0xf1, 0x4d, 0xfa, 0xe5, 0x06,
	0xb5, 0x9d, 0xc8, 0x3c, 0x4f, 0x8b, 0x74, 0x3e, 0x5b, 0xa0, 0xce, 0xe7, 0x77, 0x3f, 0xdf, 0x9e,
	0xfc, 0x0a, 0x7f, 0xdb, 0xfb, 0x4a, 0x46, 0x54, 0x55, 0x91, 0x3f, 0x31, 0x2a, 0xb5, 0x62, 0x91,
	0xda, 0x6c, 0xce, 0xfd, 0x8b, 0x7f, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x5e, 0x6e, 0x72,
	0x63, 0x02, 0x00, 0x00,
}