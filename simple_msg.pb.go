// Code generated by protoc-gen-go. DO NOT EDIT.
// source: simple_msg.proto

package simple_msg

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0xce, 0xcc, 0x2d,
	0xc8, 0x49, 0x8d, 0xcf, 0x2d, 0x4e, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x28, 0xe5, 0x73, 0xb1, 0x7b, 0xa4, 0x26, 0xa6, 0x04, 0xa5, 0x16, 0x0a, 0x09, 0x70, 0x31, 0x27,
	0xe7, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x81, 0x98, 0x42, 0x62, 0x5c, 0x6c, 0xc5,
	0xa5, 0x49, 0x20, 0x41, 0x26, 0xb0, 0x20, 0x94, 0x07, 0x52, 0x59, 0x9c, 0x5a, 0x28, 0xc1, 0x0c,
	0x51, 0x59, 0x9c, 0x5a, 0x28, 0x24, 0xc1, 0xc5, 0x5e, 0x96, 0x5a, 0x54, 0x9c, 0x99, 0x9f, 0x27,
	0xc1, 0x02, 0x16, 0x85, 0x71, 0x85, 0xf8, 0xb8, 0x98, 0x52, 0x2b, 0x24, 0x58, 0x15, 0x18, 0x35,
	0x78, 0x82, 0x98, 0x52, 0x2b, 0x94, 0x16, 0x33, 0x42, 0x6d, 0x2c, 0x2e, 0xa0, 0x8f, 0x8d, 0x42,
	0x92, 0x5c, 0x1c, 0xa9, 0x45, 0x45, 0xf1, 0xc9, 0xf9, 0x29, 0xa9, 0x12, 0x6c, 0x10, 0xa5, 0xa9,
	0x45, 0x45, 0xce, 0xf9, 0x29, 0xa9, 0x42, 0xe2, 0x5c, 0x20, 0x26, 0x28, 0x20, 0x24, 0xd8, 0x15,
	0x18, 0x35, 0x38, 0x83, 0xd8, 0x52, 0x8b, 0x8a, 0x7c, 0x8b, 0xd3, 0x9d, 0x78, 0xa2, 0x90, 0x02,
	0x29, 0x89, 0x0d, 0x1c, 0x6e, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1b, 0xc0, 0x5d, 0x96,
	0x4b, 0x01, 0x00, 0x00,
}
