// Code generated by protoc-gen-go. DO NOT EDIT.
// source: simple_msg.proto

package simple_msg

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

type HeadReq struct {
	Cmd                  int32    `protobuf:"varint,1,opt,name=cmd,proto3" json:"cmd,omitempty"`
	Subcmd               int32    `protobuf:"varint,2,opt,name=subcmd,proto3" json:"subcmd,omitempty"`
	Seq                  int32    `protobuf:"varint,3,opt,name=seq,proto3" json:"seq,omitempty"`
	Version              int32    `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	Ex                   []byte   `protobuf:"bytes,5,opt,name=ex,proto3" json:"ex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeadReq) Reset()         { *m = HeadReq{} }
func (m *HeadReq) String() string { return proto.CompactTextString(m) }
func (*HeadReq) ProtoMessage()    {}
func (*HeadReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8cac7452d7cb78a, []int{0}
}

func (m *HeadReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeadReq.Unmarshal(m, b)
}
func (m *HeadReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeadReq.Marshal(b, m, deterministic)
}
func (m *HeadReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeadReq.Merge(m, src)
}
func (m *HeadReq) XXX_Size() int {
	return xxx_messageInfo_HeadReq.Size(m)
}
func (m *HeadReq) XXX_DiscardUnknown() {
	xxx_messageInfo_HeadReq.DiscardUnknown(m)
}

var xxx_messageInfo_HeadReq proto.InternalMessageInfo

func (m *HeadReq) GetCmd() int32 {
	if m != nil {
		return m.Cmd
	}
	return 0
}

func (m *HeadReq) GetSubcmd() int32 {
	if m != nil {
		return m.Subcmd
	}
	return 0
}

func (m *HeadReq) GetSeq() int32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *HeadReq) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *HeadReq) GetEx() []byte {
	if m != nil {
		return m.Ex
	}
	return nil
}

type HeadRsp struct {
	Cmd                  int32    `protobuf:"varint,1,opt,name=cmd,proto3" json:"cmd,omitempty"`
	Subcmd               int32    `protobuf:"varint,2,opt,name=subcmd,proto3" json:"subcmd,omitempty"`
	Seq                  int32    `protobuf:"varint,3,opt,name=seq,proto3" json:"seq,omitempty"`
	Version              int32    `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	Ex                   []byte   `protobuf:"bytes,5,opt,name=ex,proto3" json:"ex,omitempty"`
	ErrCode              int32    `protobuf:"varint,6,opt,name=err_code,json=errCode,proto3" json:"err_code,omitempty"`
	ErrMsg               string   `protobuf:"bytes,7,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeadRsp) Reset()         { *m = HeadRsp{} }
func (m *HeadRsp) String() string { return proto.CompactTextString(m) }
func (*HeadRsp) ProtoMessage()    {}
func (*HeadRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8cac7452d7cb78a, []int{1}
}

func (m *HeadRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeadRsp.Unmarshal(m, b)
}
func (m *HeadRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeadRsp.Marshal(b, m, deterministic)
}
func (m *HeadRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeadRsp.Merge(m, src)
}
func (m *HeadRsp) XXX_Size() int {
	return xxx_messageInfo_HeadRsp.Size(m)
}
func (m *HeadRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_HeadRsp.DiscardUnknown(m)
}

var xxx_messageInfo_HeadRsp proto.InternalMessageInfo

func (m *HeadRsp) GetCmd() int32 {
	if m != nil {
		return m.Cmd
	}
	return 0
}

func (m *HeadRsp) GetSubcmd() int32 {
	if m != nil {
		return m.Subcmd
	}
	return 0
}

func (m *HeadRsp) GetSeq() int32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *HeadRsp) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *HeadRsp) GetEx() []byte {
	if m != nil {
		return m.Ex
	}
	return nil
}

func (m *HeadRsp) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *HeadRsp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func init() {
	proto.RegisterType((*HeadReq)(nil), "simple_msg.HeadReq")
	proto.RegisterType((*HeadRsp)(nil), "simple_msg.HeadRsp")
}

func init() { proto.RegisterFile("simple_msg.proto", fileDescriptor_f8cac7452d7cb78a) }

var fileDescriptor_f8cac7452d7cb78a = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x90, 0xb1, 0x4a, 0x03, 0x41,
	0x10, 0x86, 0xdd, 0x4b, 0xb2, 0x67, 0x86, 0x20, 0x61, 0x04, 0x5d, 0xad, 0x8e, 0xab, 0xb6, 0x0a,
	0xa2, 0xe0, 0x03, 0x98, 0xc6, 0x26, 0xcd, 0xda, 0xd9, 0x04, 0x93, 0x1d, 0x8e, 0x80, 0xe7, 0xee,
	0xcd, 0xa8, 0xe4, 0x7d, 0x7c, 0x51, 0xd9, 0xbd, 0x13, 0x0f, 0xec, 0xed, 0xe6, 0xff, 0xf8, 0xe1,
	0x1f, 0x3e, 0x58, 0xca, 0xa1, 0x8d, 0xaf, 0xb4, 0x6d, 0xa5, 0x59, 0x45, 0x0e, 0xef, 0x01, 0xe1,
	0x97, 0xd4, 0x01, 0xca, 0x47, 0x7a, 0xf1, 0x8e, 0x3a, 0x5c, 0xc2, 0x64, 0xdf, 0x7a, 0xa3, 0x2a,
	0x65, 0x67, 0x2e, 0x9d, 0x78, 0x01, 0x5a, 0x3e, 0x76, 0x09, 0x16, 0x19, 0x0e, 0x29, 0x35, 0x85,
	0x3a, 0x33, 0xe9, 0x9b, 0x42, 0x1d, 0x1a, 0x28, 0x3f, 0x89, 0xe5, 0x10, 0xde, 0xcc, 0x34, 0xd3,
	0x9f, 0x88, 0x67, 0x50, 0xd0, 0xd1, 0xcc, 0x2a, 0x65, 0x17, 0xae, 0xa0, 0x63, 0xfd, 0xa5, 0x86,
	0x45, 0x89, 0xff, 0xb3, 0x88, 0x57, 0x70, 0x4a, 0xcc, 0xdb, 0x7d, 0xf0, 0x64, 0x74, 0x5f, 0x25,
	0xe6, 0x75, 0xf0, 0x84, 0x97, 0x90, 0xce, 0x24, 0xc2, 0x94, 0x95, 0xb2, 0x73, 0xa7, 0x89, 0x79,
	0x23, 0xcd, 0xed, 0x1a, 0xe6, 0x4f, 0x59, 0xd2, 0x46, 0x1a, 0xbc, 0x87, 0x69, 0xfa, 0x18, 0xcf,
	0x57, 0x23, 0x95, 0x83, 0xb5, 0xeb, 0xbf, 0x50, 0x62, 0x7d, 0x62, 0xd5, 0x8d, 0x7a, 0x58, 0x3c,
	0x8f, 0x4c, 0xef, 0x74, 0x96, 0x7f, 0xf7, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x49, 0xe7, 0xb0, 0x30,
	0x90, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SimpleMsgClient is the client API for SimpleMsg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SimpleMsgClient interface {
	Head(ctx context.Context, opts ...grpc.CallOption) (SimpleMsg_HeadClient, error)
}

type simpleMsgClient struct {
	cc *grpc.ClientConn
}

func NewSimpleMsgClient(cc *grpc.ClientConn) SimpleMsgClient {
	return &simpleMsgClient{cc}
}

func (c *simpleMsgClient) Head(ctx context.Context, opts ...grpc.CallOption) (SimpleMsg_HeadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SimpleMsg_serviceDesc.Streams[0], "/simple_msg.SimpleMsg/Head", opts...)
	if err != nil {
		return nil, err
	}
	x := &simpleMsgHeadClient{stream}
	return x, nil
}

type SimpleMsg_HeadClient interface {
	Send(*HeadReq) error
	Recv() (*HeadRsp, error)
	grpc.ClientStream
}

type simpleMsgHeadClient struct {
	grpc.ClientStream
}

func (x *simpleMsgHeadClient) Send(m *HeadReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *simpleMsgHeadClient) Recv() (*HeadRsp, error) {
	m := new(HeadRsp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SimpleMsgServer is the server API for SimpleMsg service.
type SimpleMsgServer interface {
	Head(SimpleMsg_HeadServer) error
}

// UnimplementedSimpleMsgServer can be embedded to have forward compatible implementations.
type UnimplementedSimpleMsgServer struct {
}

func (*UnimplementedSimpleMsgServer) Head(srv SimpleMsg_HeadServer) error {
	return status.Errorf(codes.Unimplemented, "method Head not implemented")
}

func RegisterSimpleMsgServer(s *grpc.Server, srv SimpleMsgServer) {
	s.RegisterService(&_SimpleMsg_serviceDesc, srv)
}

func _SimpleMsg_Head_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SimpleMsgServer).Head(&simpleMsgHeadServer{stream})
}

type SimpleMsg_HeadServer interface {
	Send(*HeadRsp) error
	Recv() (*HeadReq, error)
	grpc.ServerStream
}

type simpleMsgHeadServer struct {
	grpc.ServerStream
}

func (x *simpleMsgHeadServer) Send(m *HeadRsp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *simpleMsgHeadServer) Recv() (*HeadReq, error) {
	m := new(HeadReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SimpleMsg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "simple_msg.SimpleMsg",
	HandlerType: (*SimpleMsgServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Head",
			Handler:       _SimpleMsg_Head_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "simple_msg.proto",
}
