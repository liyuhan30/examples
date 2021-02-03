// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base-store_stringWrapper.proto

package store_base

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

type StoreStringWrapper struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreStringWrapper) Reset()         { *m = StoreStringWrapper{} }
func (m *StoreStringWrapper) String() string { return proto.CompactTextString(m) }
func (*StoreStringWrapper) ProtoMessage()    {}
func (*StoreStringWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_887a9ac5cb886fb4, []int{0}
}

func (m *StoreStringWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreStringWrapper.Unmarshal(m, b)
}
func (m *StoreStringWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreStringWrapper.Marshal(b, m, deterministic)
}
func (m *StoreStringWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreStringWrapper.Merge(m, src)
}
func (m *StoreStringWrapper) XXX_Size() int {
	return xxx_messageInfo_StoreStringWrapper.Size(m)
}
func (m *StoreStringWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreStringWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_StoreStringWrapper proto.InternalMessageInfo

func (m *StoreStringWrapper) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*StoreStringWrapper)(nil), "StoreStringWrapper")
}

func init() { proto.RegisterFile("base-store_stringWrapper.proto", fileDescriptor_887a9ac5cb886fb4) }

var fileDescriptor_887a9ac5cb886fb4 = []byte{
	// 105 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0x4a, 0x2c, 0x4e,
	0xd5, 0x2d, 0x2e, 0xc9, 0x2f, 0x4a, 0x8d, 0x2f, 0x2e, 0x29, 0xca, 0xcc, 0x4b, 0x0f, 0x2f, 0x4a,
	0x2c, 0x28, 0x48, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xd2, 0xe2, 0x12, 0x0a, 0x06,
	0x49, 0x06, 0x23, 0xcb, 0x09, 0x89, 0x70, 0xb1, 0x96, 0x25, 0xe6, 0x94, 0xa6, 0x4a, 0x30, 0x2a,
	0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x4e, 0xd2, 0x51, 0x92, 0x60, 0x4d, 0xc9, 0xf9, 0x39, 0xfa,
	0x60, 0x13, 0x21, 0x64, 0x3c, 0xc8, 0x8a, 0x24, 0x36, 0xb0, 0x94, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0x94, 0xbc, 0x78, 0x0b, 0x71, 0x00, 0x00, 0x00,
}
