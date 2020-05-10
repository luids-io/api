// Code generated by protoc-gen-go. DO NOT EDIT.
// source: archive.proto

package pb

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

type SaveEventRequest struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 EventType            `protobuf:"varint,2,opt,name=type,proto3,enum=luids.event.v1.EventType" json:"type,omitempty"`
	Code                 int32                `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
	Level                EventLevel           `protobuf:"varint,4,opt,name=level,proto3,enum=luids.event.v1.EventLevel" json:"level,omitempty"`
	CreatedTs            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_ts,json=createdTs,proto3" json:"created_ts,omitempty"`
	Source               *EventSource         `protobuf:"bytes,6,opt,name=source,proto3" json:"source,omitempty"`
	Processors           []*ProcessInfo       `protobuf:"bytes,7,rep,name=processors,proto3" json:"processors,omitempty"`
	Data                 *EventData           `protobuf:"bytes,8,opt,name=data,proto3" json:"data,omitempty"`
	Codename             string               `protobuf:"bytes,9,opt,name=codename,proto3" json:"codename,omitempty"`
	Description          string               `protobuf:"bytes,10,opt,name=description,proto3" json:"description,omitempty"`
	Tags                 []string             `protobuf:"bytes,11,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SaveEventRequest) Reset()         { *m = SaveEventRequest{} }
func (m *SaveEventRequest) String() string { return proto.CompactTextString(m) }
func (*SaveEventRequest) ProtoMessage()    {}
func (*SaveEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_04f37ff213ec9fca, []int{0}
}

func (m *SaveEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveEventRequest.Unmarshal(m, b)
}
func (m *SaveEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveEventRequest.Marshal(b, m, deterministic)
}
func (m *SaveEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveEventRequest.Merge(m, src)
}
func (m *SaveEventRequest) XXX_Size() int {
	return xxx_messageInfo_SaveEventRequest.Size(m)
}
func (m *SaveEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveEventRequest proto.InternalMessageInfo

func (m *SaveEventRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SaveEventRequest) GetType() EventType {
	if m != nil {
		return m.Type
	}
	return EventType_UNDEFINED
}

func (m *SaveEventRequest) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SaveEventRequest) GetLevel() EventLevel {
	if m != nil {
		return m.Level
	}
	return EventLevel_INFO
}

func (m *SaveEventRequest) GetCreatedTs() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedTs
	}
	return nil
}

func (m *SaveEventRequest) GetSource() *EventSource {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *SaveEventRequest) GetProcessors() []*ProcessInfo {
	if m != nil {
		return m.Processors
	}
	return nil
}

func (m *SaveEventRequest) GetData() *EventData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *SaveEventRequest) GetCodename() string {
	if m != nil {
		return m.Codename
	}
	return ""
}

func (m *SaveEventRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SaveEventRequest) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type SaveEventResponse struct {
	StorageID            string   `protobuf:"bytes,1,opt,name=storageID,proto3" json:"storageID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveEventResponse) Reset()         { *m = SaveEventResponse{} }
func (m *SaveEventResponse) String() string { return proto.CompactTextString(m) }
func (*SaveEventResponse) ProtoMessage()    {}
func (*SaveEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_04f37ff213ec9fca, []int{1}
}

func (m *SaveEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveEventResponse.Unmarshal(m, b)
}
func (m *SaveEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveEventResponse.Marshal(b, m, deterministic)
}
func (m *SaveEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveEventResponse.Merge(m, src)
}
func (m *SaveEventResponse) XXX_Size() int {
	return xxx_messageInfo_SaveEventResponse.Size(m)
}
func (m *SaveEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SaveEventResponse proto.InternalMessageInfo

func (m *SaveEventResponse) GetStorageID() string {
	if m != nil {
		return m.StorageID
	}
	return ""
}

func init() {
	proto.RegisterType((*SaveEventRequest)(nil), "luids.event.v1.SaveEventRequest")
	proto.RegisterType((*SaveEventResponse)(nil), "luids.event.v1.SaveEventResponse")
}

func init() { proto.RegisterFile("archive.proto", fileDescriptor_04f37ff213ec9fca) }

var fileDescriptor_04f37ff213ec9fca = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4f, 0x6f, 0xd4, 0x3c,
	0x10, 0xc6, 0xdf, 0xec, 0xbf, 0x36, 0xb3, 0x7a, 0x57, 0xe0, 0x93, 0x09, 0x1c, 0xc2, 0x4a, 0x88,
	0x5c, 0x1a, 0xd3, 0xed, 0x09, 0xf5, 0x04, 0x2a, 0x87, 0x4a, 0x1c, 0x90, 0xbb, 0x27, 0x24, 0x84,
	0x1c, 0x67, 0x36, 0x6b, 0x29, 0x89, 0x8d, 0xed, 0x44, 0xea, 0x07, 0xe4, 0x7b, 0xa1, 0x38, 0x4b,
	0xbb, 0x54, 0x0b, 0xb7, 0xb1, 0xfd, 0x7b, 0xe6, 0xb1, 0x9e, 0x19, 0xf8, 0x5f, 0x58, 0xb9, 0x57,
	0x3d, 0xe6, 0xc6, 0x6a, 0xaf, 0xc9, 0xaa, 0xee, 0x54, 0xe9, 0x72, 0xec, 0xb1, 0xf5, 0x79, 0x7f,
	0x99, 0x5c, 0x57, 0xca, 0xef, 0xbb, 0x22, 0x97, 0xba, 0x61, 0x95, 0xae, 0x45, 0x5b, 0xb1, 0x00,
	0x16, 0xdd, 0x8e, 0x19, 0x7f, 0x6f, 0xd0, 0x31, 0xaf, 0x1a, 0x74, 0x5e, 0x34, 0xe6, 0xb1, 0x1a,
	0x9b, 0x25, 0x9b, 0x23, 0x71, 0xe8, 0x7b, 0xa1, 0x34, 0x13, 0x46, 0x31, 0x27, 0xf7, 0xd8, 0x08,
	0xc7, 0x82, 0x0d, 0x93, 0xba, 0x69, 0x74, 0x3b, 0x6a, 0xd6, 0x3f, 0xa7, 0xf0, 0xec, 0x4e, 0xf4,
	0xf8, 0x69, 0x78, 0xe2, 0xf8, 0xa3, 0x43, 0xe7, 0xc9, 0x0a, 0x26, 0xaa, 0xa4, 0x51, 0x1a, 0x65,
	0x31, 0x9f, 0xa8, 0x92, 0x5c, 0xc0, 0x6c, 0xf0, 0xa6, 0x93, 0x34, 0xca, 0x56, 0x9b, 0x17, 0xf9,
	0x9f, 0x9f, 0xce, 0x83, 0x76, 0x7b, 0x6f, 0x90, 0x07, 0x8c, 0x10, 0x98, 0x49, 0x5d, 0x22, 0x9d,
	0xa6, 0x51, 0x36, 0xe7, 0xa1, 0x26, 0xef, 0x60, 0x5e, 0x63, 0x8f, 0x35, 0x9d, 0x85, 0x1e, 0xc9,
	0xc9, 0x1e, 0x9f, 0x07, 0x82, 0x8f, 0x20, 0x79, 0x0f, 0x20, 0x2d, 0x0a, 0x8f, 0xe5, 0x77, 0xef,
	0xe8, 0x3c, 0x8d, 0xb2, 0xe5, 0x26, 0xc9, 0x2b, 0xad, 0xab, 0xfa, 0x90, 0x5e, 0xd1, 0xed, 0xf2,
	0xed, 0xef, 0x0c, 0x78, 0x7c, 0xa0, 0xb7, 0x8e, 0x5c, 0xc1, 0xc2, 0xe9, 0xce, 0x4a, 0xa4, 0x8b,
	0x20, 0x7b, 0x79, 0xd2, 0xed, 0x2e, 0x20, 0xfc, 0x80, 0x92, 0x6b, 0x00, 0x63, 0xb5, 0x44, 0xe7,
	0xb4, 0x75, 0xf4, 0x2c, 0x9d, 0x9e, 0x12, 0x7e, 0x19, 0x89, 0xdb, 0x76, 0xa7, 0xf9, 0x11, 0x3e,
	0x24, 0x54, 0x0a, 0x2f, 0xe8, 0x79, 0xf0, 0x3b, 0x9d, 0xd0, 0x8d, 0xf0, 0x82, 0x07, 0x8c, 0x24,
	0x70, 0x3e, 0xa4, 0xd2, 0x8a, 0x06, 0x69, 0x1c, 0x62, 0x7e, 0x38, 0x93, 0x14, 0x96, 0x25, 0x3a,
	0x69, 0x95, 0xf1, 0x4a, 0xb7, 0x14, 0xc2, 0xf3, 0xf1, 0xd5, 0x90, 0xaf, 0x17, 0x95, 0xa3, 0xcb,
	0x74, 0x9a, 0xc5, 0x3c, 0xd4, 0xeb, 0x4b, 0x78, 0x7e, 0x34, 0x46, 0x67, 0x74, 0xeb, 0x90, 0xbc,
	0x82, 0xd8, 0x79, 0x6d, 0x45, 0x85, 0xb7, 0x37, 0x87, 0x71, 0x3e, 0x5e, 0x6c, 0xbe, 0xc1, 0xd9,
	0x87, 0x71, 0x19, 0x09, 0x87, 0xf8, 0x41, 0x4d, 0xd2, 0xa7, 0xbf, 0x7f, 0xba, 0x1f, 0xc9, 0xeb,
	0x7f, 0x10, 0xa3, 0xf5, 0xfa, 0xbf, 0x8f, 0x6f, 0xbf, 0xbe, 0xf9, 0xdb, 0x3e, 0x8e, 0x7b, 0x58,
	0x59, 0x23, 0x99, 0x29, 0x8a, 0x45, 0x18, 0xe6, 0xd5, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd0,
	0xdb, 0xa0, 0xa5, 0x1b, 0x03, 0x00, 0x00,
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
	SaveEvent(ctx context.Context, in *SaveEventRequest, opts ...grpc.CallOption) (*SaveEventResponse, error)
}

type archiveClient struct {
	cc *grpc.ClientConn
}

func NewArchiveClient(cc *grpc.ClientConn) ArchiveClient {
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

func (*UnimplementedArchiveServer) SaveEvent(ctx context.Context, req *SaveEventRequest) (*SaveEventResponse, error) {
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
	Metadata: "archive.proto",
}