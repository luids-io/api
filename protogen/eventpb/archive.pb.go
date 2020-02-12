// Code generated by protoc-gen-go. DO NOT EDIT.
// source: archive.proto

package eventpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
	Event                *Event   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
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

func (m *SaveEventRequest) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type SaveEventResponse struct {
	EventID              string   `protobuf:"bytes,1,opt,name=eventID,proto3" json:"eventID,omitempty"`
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

func (m *SaveEventResponse) GetEventID() string {
	if m != nil {
		return m.EventID
	}
	return ""
}

func init() {
	proto.RegisterType((*SaveEventRequest)(nil), "luids.event.v1.SaveEventRequest")
	proto.RegisterType((*SaveEventResponse)(nil), "luids.event.v1.SaveEventResponse")
}

func init() { proto.RegisterFile("archive.proto", fileDescriptor_04f37ff213ec9fca) }

var fileDescriptor_04f37ff213ec9fca = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x2c, 0x4a, 0xce,
	0xc8, 0x2c, 0x4b, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcb, 0x29, 0xcd, 0x4c, 0x29,
	0xd6, 0x4b, 0x2d, 0x4b, 0xcd, 0x2b, 0xd1, 0x2b, 0x33, 0x94, 0x32, 0x4a, 0xcf, 0x2c, 0xc9, 0x28,
	0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x07, 0x4b, 0xe9, 0x66, 0xe6, 0xeb, 0x27, 0x16, 0x64, 0xea,
	0x17, 0x27, 0x67, 0xa4, 0xe6, 0x26, 0x16, 0xeb, 0x83, 0x55, 0xea, 0x27, 0xe7, 0xe7, 0xe6, 0xe6,
	0xe7, 0x41, 0xcc, 0x50, 0xb2, 0xe7, 0x12, 0x08, 0x4e, 0x2c, 0x4b, 0x75, 0x05, 0xc9, 0x04, 0xa5,
	0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x69, 0x73, 0xb1, 0x82, 0x55, 0x4a, 0x30, 0x2a, 0x30, 0x6a,
	0x70, 0x1b, 0x89, 0xea, 0xa1, 0xda, 0xa3, 0x07, 0x51, 0x0c, 0x51, 0xa3, 0xa4, 0xcb, 0x25, 0x88,
	0x64, 0x40, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x04, 0x17, 0x3b, 0x58, 0xd6, 0xd3, 0x05,
	0x6c, 0x06, 0x67, 0x10, 0x8c, 0x6b, 0x14, 0xcb, 0xc5, 0xee, 0x08, 0xf1, 0x84, 0x50, 0x10, 0x17,
	0x27, 0x5c, 0xa7, 0x90, 0x02, 0xba, 0x25, 0xe8, 0xae, 0x92, 0x52, 0xc4, 0xa3, 0x02, 0x62, 0xad,
	0x12, 0x83, 0x93, 0x56, 0x94, 0x06, 0xae, 0x40, 0x00, 0xfb, 0x37, 0x3d, 0x35, 0x0f, 0x12, 0x0a,
	0x05, 0x49, 0x49, 0x6c, 0x60, 0x11, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x53, 0xdd, 0xd3,
	0x0d, 0x56, 0x01, 0x00, 0x00,
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