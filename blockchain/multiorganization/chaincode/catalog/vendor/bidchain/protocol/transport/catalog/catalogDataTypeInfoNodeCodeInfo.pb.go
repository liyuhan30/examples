// Code generated by protoc-gen-go. DO NOT EDIT.
// source: catalog-catalogDataTypeInfoNodeCodeInfo.proto

package catalog

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

type CatalogDataTypeInfoNodeCodeInfo struct {
	CodeName             string   `protobuf:"bytes,1,opt,name=codeName,proto3" json:"codeName,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CatalogDataTypeInfoNodeCodeInfo) Reset()         { *m = CatalogDataTypeInfoNodeCodeInfo{} }
func (m *CatalogDataTypeInfoNodeCodeInfo) String() string { return proto.CompactTextString(m) }
func (*CatalogDataTypeInfoNodeCodeInfo) ProtoMessage()    {}
func (*CatalogDataTypeInfoNodeCodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc7a785332378887, []int{0}
}

func (m *CatalogDataTypeInfoNodeCodeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CatalogDataTypeInfoNodeCodeInfo.Unmarshal(m, b)
}
func (m *CatalogDataTypeInfoNodeCodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CatalogDataTypeInfoNodeCodeInfo.Marshal(b, m, deterministic)
}
func (m *CatalogDataTypeInfoNodeCodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CatalogDataTypeInfoNodeCodeInfo.Merge(m, src)
}
func (m *CatalogDataTypeInfoNodeCodeInfo) XXX_Size() int {
	return xxx_messageInfo_CatalogDataTypeInfoNodeCodeInfo.Size(m)
}
func (m *CatalogDataTypeInfoNodeCodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CatalogDataTypeInfoNodeCodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CatalogDataTypeInfoNodeCodeInfo proto.InternalMessageInfo

func (m *CatalogDataTypeInfoNodeCodeInfo) GetCodeName() string {
	if m != nil {
		return m.CodeName
	}
	return ""
}

func (m *CatalogDataTypeInfoNodeCodeInfo) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*CatalogDataTypeInfoNodeCodeInfo)(nil), "CatalogDataTypeInfoNodeCodeInfo")
}

func init() {
	proto.RegisterFile("catalog-catalogDataTypeInfoNodeCodeInfo.proto", fileDescriptor_dc7a785332378887)
}

var fileDescriptor_dc7a785332378887 = []byte{
	// 142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4d, 0x4e, 0x2c, 0x49,
	0xcc, 0xc9, 0x4f, 0x87, 0xd1, 0x2e, 0x89, 0x25, 0x89, 0x21, 0x95, 0x05, 0xa9, 0x9e, 0x79, 0x69,
	0xf9, 0x7e, 0xf9, 0x29, 0xa9, 0xce, 0xf9, 0x29, 0x60, 0xb6, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe,
	0x52, 0x38, 0x97, 0xbc, 0x33, 0x7e, 0x85, 0x42, 0x52, 0x5c, 0x1c, 0xc9, 0xf9, 0x29, 0xa9, 0x7e,
	0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x70, 0xbe, 0x90, 0x04, 0x17, 0x7b,
	0x72, 0x7e, 0x5e, 0x49, 0x6a, 0x5e, 0x89, 0x04, 0x13, 0x58, 0x0a, 0xc6, 0x75, 0x52, 0x8d, 0x52,
	0x4e, 0xca, 0x4c, 0x49, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x07, 0x5b, 0x95, 0x9c, 0x9f, 0xa3, 0x5f,
	0x52, 0x94, 0x98, 0x57, 0x5c, 0x90, 0x5f, 0x54, 0xa2, 0x0f, 0x75, 0x5d, 0x12, 0x1b, 0x58, 0xce,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x7f, 0x6e, 0xd9, 0xb7, 0x00, 0x00, 0x00,
}
