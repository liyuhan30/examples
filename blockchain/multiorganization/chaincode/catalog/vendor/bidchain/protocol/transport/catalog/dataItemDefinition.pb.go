// Code generated by protoc-gen-go. DO NOT EDIT.
// source: catalog-dataItemDefinition.proto

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

type DataItemDefinition struct {
	DataItemDefinitionId string                     `protobuf:"bytes,1,opt,name=dataItemDefinitionId,proto3" json:"dataItemDefinitionId,omitempty"`
	DataItemName         string                     `protobuf:"bytes,2,opt,name=dataItemName,proto3" json:"dataItemName,omitempty"`
	DataItemVersion      int32                      `protobuf:"varint,3,opt,name=dataItemVersion,proto3" json:"dataItemVersion,omitempty"`
	PublicFieldList      []*DataFieldDefinition     `protobuf:"bytes,4,rep,name=publicFieldList,proto3" json:"publicFieldList,omitempty"`
	Weight               int32                      `protobuf:"varint,5,opt,name=weight,proto3" json:"weight,omitempty"`
	MergeFieldList       []*DataItemMergedFieldData `protobuf:"bytes,6,rep,name=mergeFieldList,proto3" json:"mergeFieldList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *DataItemDefinition) Reset()         { *m = DataItemDefinition{} }
func (m *DataItemDefinition) String() string { return proto.CompactTextString(m) }
func (*DataItemDefinition) ProtoMessage()    {}
func (*DataItemDefinition) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc63dacc5bb62984, []int{0}
}

func (m *DataItemDefinition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataItemDefinition.Unmarshal(m, b)
}
func (m *DataItemDefinition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataItemDefinition.Marshal(b, m, deterministic)
}
func (m *DataItemDefinition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataItemDefinition.Merge(m, src)
}
func (m *DataItemDefinition) XXX_Size() int {
	return xxx_messageInfo_DataItemDefinition.Size(m)
}
func (m *DataItemDefinition) XXX_DiscardUnknown() {
	xxx_messageInfo_DataItemDefinition.DiscardUnknown(m)
}

var xxx_messageInfo_DataItemDefinition proto.InternalMessageInfo

func (m *DataItemDefinition) GetDataItemDefinitionId() string {
	if m != nil {
		return m.DataItemDefinitionId
	}
	return ""
}

func (m *DataItemDefinition) GetDataItemName() string {
	if m != nil {
		return m.DataItemName
	}
	return ""
}

func (m *DataItemDefinition) GetDataItemVersion() int32 {
	if m != nil {
		return m.DataItemVersion
	}
	return 0
}

func (m *DataItemDefinition) GetPublicFieldList() []*DataFieldDefinition {
	if m != nil {
		return m.PublicFieldList
	}
	return nil
}

func (m *DataItemDefinition) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *DataItemDefinition) GetMergeFieldList() []*DataItemMergedFieldData {
	if m != nil {
		return m.MergeFieldList
	}
	return nil
}

func init() {
	proto.RegisterType((*DataItemDefinition)(nil), "DataItemDefinition")
}

func init() { proto.RegisterFile("catalog-dataItemDefinition.proto", fileDescriptor_bc63dacc5bb62984) }

var fileDescriptor_bc63dacc5bb62984 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x49, 0x6b, 0x03, 0xae, 0x62, 0x61, 0x29, 0xb2, 0x78, 0x8a, 0x95, 0x42, 0x2e, 0x26,
	0x50, 0xef, 0x22, 0x12, 0x84, 0x82, 0x7a, 0xc8, 0xc1, 0x83, 0xb7, 0x49, 0x76, 0x4c, 0x07, 0x92,
	0xdd, 0xb0, 0x19, 0xf1, 0x6f, 0xf9, 0x13, 0xc5, 0x6d, 0x8b, 0x36, 0xc9, 0x71, 0xde, 0xfb, 0x78,
	0x6f, 0xf7, 0x89, 0xa8, 0x04, 0x86, 0xda, 0x56, 0xb7, 0x1a, 0x18, 0x36, 0x8c, 0x4d, 0x86, 0x1f,
	0x64, 0x88, 0xc9, 0x9a, 0xa4, 0x75, 0x96, 0xed, 0xd5, 0xf5, 0x7f, 0xe2, 0x89, 0xb0, 0xd6, 0x03,
	0x64, 0xd5, 0x0f, 0x79, 0x41, 0x57, 0xa1, 0xde, 0xc1, 0xc0, 0xb0, 0xc3, 0x96, 0xdf, 0x13, 0x21,
	0xb3, 0x41, 0x8d, 0x5c, 0x8b, 0xc5, 0xb0, 0x7c, 0xa3, 0x55, 0x10, 0x05, 0xf1, 0x69, 0x3e, 0xea,
	0xc9, 0xa5, 0x38, 0x3f, 0xe8, 0xaf, 0xd0, 0xa0, 0x9a, 0x78, 0xf6, 0x48, 0x93, 0xb1, 0x98, 0x1f,
	0xee, 0x37, 0x74, 0x1d, 0x59, 0xa3, 0xa6, 0x51, 0x10, 0xcf, 0xf2, 0xbe, 0x2c, 0xef, 0xc5, 0xbc,
	0xfd, 0x2c, 0x6a, 0x2a, 0xfd, 0x8b, 0x9f, 0xa9, 0x63, 0x75, 0x12, 0x4d, 0xe3, 0xb3, 0xf5, 0x22,
	0xc9, 0x86, 0x9f, 0xce, 0xfb, 0xb0, 0xbc, 0x14, 0xe1, 0x17, 0x52, 0xb5, 0x65, 0x35, 0xf3, 0x05,
	0xfb, 0x4b, 0x3e, 0x88, 0x8b, 0xe6, 0x77, 0x89, 0xbf, 0xd8, 0xd0, 0xc7, 0xaa, 0x24, 0x1b, 0x1f,
	0x2a, 0xef, 0xf1, 0x8f, 0xab, 0xf7, 0x9b, 0x82, 0x74, 0xb9, 0x05, 0x32, 0xa9, 0x1f, 0xb1, 0xb4,
	0x75, 0xca, 0x0e, 0x4c, 0xd7, 0x5a, 0xc7, 0xe9, 0x7e, 0xf7, 0x22, 0xf4, 0xde, 0xdd, 0x4f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x0e, 0x1b, 0x54, 0x97, 0xce, 0x01, 0x00, 0x00,
}