// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base-store_envelopeInfo.proto

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

type StoreEnvelopeInfo struct {
	EnvelopeData         []byte   `protobuf:"bytes,1,opt,name=envelopeData,proto3" json:"envelopeData,omitempty"`
	EncryptMethod        string   `protobuf:"bytes,2,opt,name=encryptMethod,proto3" json:"encryptMethod,omitempty"`
	EncryptPublicKey     string   `protobuf:"bytes,3,opt,name=encryptPublicKey,proto3" json:"encryptPublicKey,omitempty"`
	EnvelopeIdentifier   string   `protobuf:"bytes,4,opt,name=envelopeIdentifier,proto3" json:"envelopeIdentifier,omitempty"`
	Extension            string   `protobuf:"bytes,5,opt,name=extension,proto3" json:"extension,omitempty"`
	Description          string   `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreEnvelopeInfo) Reset()         { *m = StoreEnvelopeInfo{} }
func (m *StoreEnvelopeInfo) String() string { return proto.CompactTextString(m) }
func (*StoreEnvelopeInfo) ProtoMessage()    {}
func (*StoreEnvelopeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c6a4ef9630a6a74, []int{0}
}

func (m *StoreEnvelopeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreEnvelopeInfo.Unmarshal(m, b)
}
func (m *StoreEnvelopeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreEnvelopeInfo.Marshal(b, m, deterministic)
}
func (m *StoreEnvelopeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreEnvelopeInfo.Merge(m, src)
}
func (m *StoreEnvelopeInfo) XXX_Size() int {
	return xxx_messageInfo_StoreEnvelopeInfo.Size(m)
}
func (m *StoreEnvelopeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreEnvelopeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_StoreEnvelopeInfo proto.InternalMessageInfo

func (m *StoreEnvelopeInfo) GetEnvelopeData() []byte {
	if m != nil {
		return m.EnvelopeData
	}
	return nil
}

func (m *StoreEnvelopeInfo) GetEncryptMethod() string {
	if m != nil {
		return m.EncryptMethod
	}
	return ""
}

func (m *StoreEnvelopeInfo) GetEncryptPublicKey() string {
	if m != nil {
		return m.EncryptPublicKey
	}
	return ""
}

func (m *StoreEnvelopeInfo) GetEnvelopeIdentifier() string {
	if m != nil {
		return m.EnvelopeIdentifier
	}
	return ""
}

func (m *StoreEnvelopeInfo) GetExtension() string {
	if m != nil {
		return m.Extension
	}
	return ""
}

func (m *StoreEnvelopeInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*StoreEnvelopeInfo)(nil), "StoreEnvelopeInfo")
}

func init() { proto.RegisterFile("base-store_envelopeInfo.proto", fileDescriptor_1c6a4ef9630a6a74) }

var fileDescriptor_1c6a4ef9630a6a74 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0x4a, 0x2c, 0x4e,
	0xd5, 0x2d, 0x2e, 0xc9, 0x2f, 0x4a, 0x8d, 0x4f, 0xcd, 0x2b, 0x4b, 0xcd, 0xc9, 0x2f, 0x48, 0xf5,
	0xcc, 0x4b, 0xcb, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xfa, 0xc9, 0xc8, 0x25, 0x18, 0x0c,
	0x92, 0x74, 0x45, 0x92, 0x13, 0x52, 0xe2, 0xe2, 0x81, 0xa9, 0x75, 0x49, 0x2c, 0x49, 0x94, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x09, 0x42, 0x11, 0x13, 0x52, 0xe1, 0xe2, 0x4d, 0xcd, 0x4b, 0x2e, 0xaa,
	0x2c, 0x28, 0xf1, 0x4d, 0x2d, 0xc9, 0xc8, 0x4f, 0x91, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x42,
	0x15, 0x14, 0xd2, 0xe2, 0x12, 0x80, 0x0a, 0x04, 0x94, 0x26, 0xe5, 0x64, 0x26, 0x7b, 0xa7, 0x56,
	0x4a, 0x30, 0x83, 0x15, 0x62, 0x88, 0x0b, 0xe9, 0x71, 0x09, 0xc1, 0x5d, 0x98, 0x92, 0x9a, 0x57,
	0x92, 0x99, 0x96, 0x99, 0x5a, 0x24, 0xc1, 0x02, 0x56, 0x8d, 0x45, 0x46, 0x48, 0x86, 0x8b, 0x33,
	0xb5, 0xa2, 0x24, 0x35, 0xaf, 0x38, 0x33, 0x3f, 0x4f, 0x82, 0x15, 0xac, 0x0c, 0x21, 0x20, 0xa4,
	0xc0, 0xc5, 0x9d, 0x92, 0x5a, 0x9c, 0x5c, 0x94, 0x59, 0x50, 0x02, 0x92, 0x67, 0x03, 0xcb, 0x23,
	0x0b, 0x39, 0x49, 0x47, 0x49, 0x82, 0x03, 0x21, 0x39, 0x3f, 0x47, 0x1f, 0x1c, 0x40, 0x10, 0x32,
	0x1e, 0x14, 0x62, 0x49, 0x6c, 0x60, 0x29, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf5, 0xe2,
	0x98, 0xf8, 0x40, 0x01, 0x00, 0x00,
}
