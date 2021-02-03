// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base-store_integerWrapper.proto

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

type StoreIntegerWrapper struct {
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreIntegerWrapper) Reset()         { *m = StoreIntegerWrapper{} }
func (m *StoreIntegerWrapper) String() string { return proto.CompactTextString(m) }
func (*StoreIntegerWrapper) ProtoMessage()    {}
func (*StoreIntegerWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_41440200b12a182f, []int{0}
}

func (m *StoreIntegerWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreIntegerWrapper.Unmarshal(m, b)
}
func (m *StoreIntegerWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreIntegerWrapper.Marshal(b, m, deterministic)
}
func (m *StoreIntegerWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreIntegerWrapper.Merge(m, src)
}
func (m *StoreIntegerWrapper) XXX_Size() int {
	return xxx_messageInfo_StoreIntegerWrapper.Size(m)
}
func (m *StoreIntegerWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreIntegerWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_StoreIntegerWrapper proto.InternalMessageInfo

func (m *StoreIntegerWrapper) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*StoreIntegerWrapper)(nil), "StoreIntegerWrapper")
}

func init() { proto.RegisterFile("base-store_integerWrapper.proto", fileDescriptor_41440200b12a182f) }

var fileDescriptor_41440200b12a182f = []byte{
	// 106 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0x4a, 0x2c, 0x4e,
	0xd5, 0x2d, 0x2e, 0xc9, 0x2f, 0x4a, 0x8d, 0xcf, 0xcc, 0x2b, 0x49, 0x4d, 0x4f, 0x2d, 0x0a, 0x2f,
	0x4a, 0x2c, 0x28, 0x48, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xd2, 0xe6, 0x12, 0x0e,
	0x06, 0xc9, 0x7a, 0xa2, 0x48, 0x0a, 0x89, 0x70, 0xb1, 0x96, 0x25, 0xe6, 0x94, 0xa6, 0x4a, 0x30,
	0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x41, 0x38, 0x4e, 0xd2, 0x51, 0x92, 0x60, 0x5d, 0xc9, 0xf9, 0x39,
	0xfa, 0x60, 0x33, 0x21, 0x64, 0x3c, 0xc8, 0x92, 0x24, 0x36, 0xb0, 0x94, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0xe1, 0x71, 0x6c, 0x51, 0x73, 0x00, 0x00, 0x00,
}