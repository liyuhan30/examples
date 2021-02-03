// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base-store_cryptoMessageEx.proto

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

type StoreCryptoMessageEx struct {
	PlatformId           string               `protobuf:"bytes,1,opt,name=platformId,proto3" json:"platformId,omitempty"`
	HashBeforeEncrypt    []byte               `protobuf:"bytes,2,opt,name=hashBeforeEncrypt,proto3" json:"hashBeforeEncrypt,omitempty"`
	HashMethod           string               `protobuf:"bytes,3,opt,name=hashMethod,proto3" json:"hashMethod,omitempty"`
	EncryptMethod        string               `protobuf:"bytes,4,opt,name=encryptMethod,proto3" json:"encryptMethod,omitempty"`
	DataType             string               `protobuf:"bytes,5,opt,name=dataType,proto3" json:"dataType,omitempty"`
	MessageData          []byte               `protobuf:"bytes,6,opt,name=messageData,proto3" json:"messageData,omitempty"`
	EnvelopInfo          []*StoreEnvelopeInfo `protobuf:"bytes,7,rep,name=envelopInfo,proto3" json:"envelopInfo,omitempty"`
	Extension            string               `protobuf:"bytes,8,opt,name=extension,proto3" json:"extension,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *StoreCryptoMessageEx) Reset()         { *m = StoreCryptoMessageEx{} }
func (m *StoreCryptoMessageEx) String() string { return proto.CompactTextString(m) }
func (*StoreCryptoMessageEx) ProtoMessage()    {}
func (*StoreCryptoMessageEx) Descriptor() ([]byte, []int) {
	return fileDescriptor_906da4fe93c43b1e, []int{0}
}

func (m *StoreCryptoMessageEx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreCryptoMessageEx.Unmarshal(m, b)
}
func (m *StoreCryptoMessageEx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreCryptoMessageEx.Marshal(b, m, deterministic)
}
func (m *StoreCryptoMessageEx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreCryptoMessageEx.Merge(m, src)
}
func (m *StoreCryptoMessageEx) XXX_Size() int {
	return xxx_messageInfo_StoreCryptoMessageEx.Size(m)
}
func (m *StoreCryptoMessageEx) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreCryptoMessageEx.DiscardUnknown(m)
}

var xxx_messageInfo_StoreCryptoMessageEx proto.InternalMessageInfo

func (m *StoreCryptoMessageEx) GetPlatformId() string {
	if m != nil {
		return m.PlatformId
	}
	return ""
}

func (m *StoreCryptoMessageEx) GetHashBeforeEncrypt() []byte {
	if m != nil {
		return m.HashBeforeEncrypt
	}
	return nil
}

func (m *StoreCryptoMessageEx) GetHashMethod() string {
	if m != nil {
		return m.HashMethod
	}
	return ""
}

func (m *StoreCryptoMessageEx) GetEncryptMethod() string {
	if m != nil {
		return m.EncryptMethod
	}
	return ""
}

func (m *StoreCryptoMessageEx) GetDataType() string {
	if m != nil {
		return m.DataType
	}
	return ""
}

func (m *StoreCryptoMessageEx) GetMessageData() []byte {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *StoreCryptoMessageEx) GetEnvelopInfo() []*StoreEnvelopeInfo {
	if m != nil {
		return m.EnvelopInfo
	}
	return nil
}

func (m *StoreCryptoMessageEx) GetExtension() string {
	if m != nil {
		return m.Extension
	}
	return ""
}

func init() {
	proto.RegisterType((*StoreCryptoMessageEx)(nil), "StoreCryptoMessageEx")
}

func init() { proto.RegisterFile("base-store_cryptoMessageEx.proto", fileDescriptor_906da4fe93c43b1e) }

var fileDescriptor_906da4fe93c43b1e = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xcd, 0x4e, 0x02, 0x31,
	0x14, 0x85, 0x03, 0x28, 0xc2, 0x1d, 0x5d, 0xd8, 0xb8, 0xa8, 0xf8, 0x93, 0xc6, 0xb8, 0x98, 0x85,
	0x62, 0xa2, 0x3e, 0x01, 0x3a, 0x0b, 0x16, 0x6c, 0x46, 0x57, 0x6e, 0x48, 0x61, 0xee, 0x38, 0x26,
	0x43, 0x6f, 0xd3, 0x36, 0x06, 0xde, 0xcd, 0x87, 0x33, 0x6d, 0x89, 0x94, 0xb8, 0x69, 0xd2, 0x73,
	0xbe, 0xfb, 0x77, 0x40, 0x2c, 0xa4, 0xc5, 0x7b, 0xeb, 0xc8, 0xe0, 0x7c, 0x69, 0x36, 0xda, 0xd1,
	0x0c, 0xad, 0x95, 0x9f, 0x58, 0xac, 0xc7, 0xda, 0x90, 0xa3, 0xd1, 0x55, 0x42, 0xa0, 0xfa, 0xc6,
	0x96, 0x34, 0x4e, 0x55, 0x4d, 0xd1, 0xbe, 0xf9, 0xe9, 0xc2, 0xd9, 0x9b, 0x37, 0x5f, 0xf6, 0xab,
	0xd9, 0x35, 0x80, 0x6e, 0xa5, 0xab, 0xc9, 0xac, 0xa6, 0x15, 0xef, 0x88, 0x4e, 0x3e, 0x2c, 0x13,
	0x85, 0xdd, 0xc1, 0x69, 0x23, 0x6d, 0x33, 0xc1, 0x9a, 0x0c, 0x16, 0x2a, 0x0c, 0xe7, 0x5d, 0xd1,
	0xc9, 0x8f, 0xcb, 0xff, 0x86, 0xef, 0xe6, 0xc5, 0x19, 0xba, 0x86, 0x2a, 0xde, 0x8b, 0xdd, 0x76,
	0x0a, 0xbb, 0x85, 0x13, 0x8c, 0xe8, 0x16, 0x39, 0x08, 0xc8, 0xbe, 0xc8, 0x46, 0x30, 0xa8, 0xa4,
	0x93, 0xef, 0x1b, 0x8d, 0xfc, 0x30, 0x00, 0x7f, 0x7f, 0x26, 0x20, 0x5b, 0xc5, 0xe5, 0x5f, 0xa5,
	0x93, 0xbc, 0x1f, 0x36, 0x49, 0x25, 0xf6, 0x0c, 0xd9, 0x36, 0x00, 0x7f, 0x3f, 0x3f, 0x12, 0xbd,
	0x3c, 0x7b, 0x64, 0xe3, 0x70, 0x7d, 0x91, 0x24, 0x53, 0xa6, 0x18, 0xbb, 0x84, 0x21, 0xae, 0x1d,
	0x2a, 0xfb, 0x45, 0x8a, 0x0f, 0xc2, 0xd0, 0x9d, 0x30, 0xb9, 0xf8, 0x38, 0x0f, 0x39, 0x2e, 0xa9,
	0x7d, 0x08, 0x19, 0xc7, 0x77, 0xee, 0x43, 0x5f, 0xf4, 0x83, 0xf5, 0xf4, 0x1b, 0x00, 0x00, 0xff,
	0xff, 0xfa, 0xfa, 0x44, 0x4e, 0xa5, 0x01, 0x00, 0x00,
}
