// Code generated by protoc-gen-go. DO NOT EDIT.
// source: catalog-catalog.proto

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

type Catalog struct {
	CatalogBasicInfo       *CatalogBasicInfo       `protobuf:"bytes,1,opt,name=catalogBasicInfo,proto3" json:"catalogBasicInfo,omitempty"`
	DataItemDefinitionList []*DataItemDefinition   `protobuf:"bytes,2,rep,name=dataItemDefinitionList,proto3" json:"dataItemDefinitionList,omitempty"`
	GlossaryList           []*DataTypeDescriptor   `protobuf:"bytes,3,rep,name=glossaryList,proto3" json:"glossaryList,omitempty"`
	TagList                []*Tag                  `protobuf:"bytes,4,rep,name=tagList,proto3" json:"tagList,omitempty"`
	ConstraintList         []*ForeignKeyConstraint `protobuf:"bytes,5,rep,name=constraintList,proto3" json:"constraintList,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                `json:"-"`
	XXX_unrecognized       []byte                  `json:"-"`
	XXX_sizecache          int32                   `json:"-"`
}

func (m *Catalog) Reset()         { *m = Catalog{} }
func (m *Catalog) String() string { return proto.CompactTextString(m) }
func (*Catalog) ProtoMessage()    {}
func (*Catalog) Descriptor() ([]byte, []int) {
	return fileDescriptor_02fea0e6db6eee2a, []int{0}
}

func (m *Catalog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Catalog.Unmarshal(m, b)
}
func (m *Catalog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Catalog.Marshal(b, m, deterministic)
}
func (m *Catalog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Catalog.Merge(m, src)
}
func (m *Catalog) XXX_Size() int {
	return xxx_messageInfo_Catalog.Size(m)
}
func (m *Catalog) XXX_DiscardUnknown() {
	xxx_messageInfo_Catalog.DiscardUnknown(m)
}

var xxx_messageInfo_Catalog proto.InternalMessageInfo

func (m *Catalog) GetCatalogBasicInfo() *CatalogBasicInfo {
	if m != nil {
		return m.CatalogBasicInfo
	}
	return nil
}

func (m *Catalog) GetDataItemDefinitionList() []*DataItemDefinition {
	if m != nil {
		return m.DataItemDefinitionList
	}
	return nil
}

func (m *Catalog) GetGlossaryList() []*DataTypeDescriptor {
	if m != nil {
		return m.GlossaryList
	}
	return nil
}

func (m *Catalog) GetTagList() []*Tag {
	if m != nil {
		return m.TagList
	}
	return nil
}

func (m *Catalog) GetConstraintList() []*ForeignKeyConstraint {
	if m != nil {
		return m.ConstraintList
	}
	return nil
}

func init() {
	proto.RegisterType((*Catalog)(nil), "Catalog")
}

func init() { proto.RegisterFile("catalog-catalog.proto", fileDescriptor_02fea0e6db6eee2a) }

var fileDescriptor_02fea0e6db6eee2a = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0xe9, 0x9f, 0xef, 0x2b, 0x4c, 0x45, 0x6c, 0xa4, 0x52, 0x5c, 0x94, 0x50, 0x11, 0xba,
	0x31, 0x81, 0xba, 0x70, 0xd5, 0x4d, 0x13, 0x84, 0x52, 0x57, 0xa1, 0x2b, 0x77, 0xb7, 0xd3, 0x49,
	0xbc, 0x10, 0xe7, 0x86, 0x99, 0xbb, 0xc9, 0x43, 0xf9, 0x8e, 0xc2, 0x34, 0x23, 0x36, 0xa9, 0xab,
	0x81, 0x73, 0x7e, 0xe7, 0x70, 0xe7, 0x88, 0xa9, 0x04, 0x86, 0x92, 0x8a, 0xa7, 0xe6, 0x8d, 0x2a,
	0x43, 0x4c, 0xf7, 0x13, 0x2f, 0x33, 0x78, 0x69, 0xde, 0x22, 0x37, 0x60, 0x51, 0x6e, 0x75, 0x4e,
	0x8d, 0x1f, 0x7a, 0xff, 0x08, 0x0c, 0x5b, 0x56, 0x9f, 0xa9, 0xca, 0x51, 0x23, 0x23, 0xe9, 0x86,
	0x58, 0x78, 0x22, 0x27, 0xa3, 0xb0, 0xd0, 0x3b, 0x55, 0x27, 0xa4, 0x2d, 0x1b, 0x40, 0xcd, 0x97,
	0x5a, 0xf6, 0x75, 0xa5, 0x52, 0x65, 0xa5, 0xc1, 0x8a, 0xc9, 0x9c, 0x88, 0xc5, 0x57, 0x5f, 0x8c,
	0x92, 0x13, 0x14, 0xac, 0xc5, 0x4d, 0xfb, 0x9a, 0x59, 0x2f, 0xec, 0x2d, 0xc7, 0xab, 0x49, 0x94,
	0xb4, 0x8c, 0xac, 0x83, 0x06, 0x3b, 0x71, 0xd7, 0x3d, 0xf6, 0x0d, 0x2d, 0xcf, 0xfa, 0xe1, 0x60,
	0x39, 0x5e, 0xdd, 0x46, 0x69, 0xc7, 0xce, 0xfe, 0x88, 0x04, 0x2f, 0xe2, 0xaa, 0x28, 0xc9, 0x5a,
	0x30, 0xb5, 0xab, 0x18, 0xfc, 0xaa, 0x38, 0xff, 0x48, 0x76, 0x06, 0x06, 0x73, 0x31, 0x62, 0x28,
	0x5c, 0x66, 0xe8, 0x32, 0xc3, 0x68, 0x0f, 0x45, 0xe6, 0xc5, 0x60, 0x2d, 0xae, 0xe5, 0xcf, 0x4c,
	0x0e, 0xfb, 0xe7, 0xb0, 0x69, 0xf4, 0x7a, 0x61, 0xc7, 0xac, 0x05, 0x6f, 0x1e, 0xdf, 0x1f, 0x0e,
	0x78, 0x94, 0x1f, 0x80, 0x3a, 0x76, 0x0b, 0x4a, 0x2a, 0x63, 0x36, 0xa0, 0x6d, 0x45, 0x86, 0xe3,
	0x66, 0x94, 0xc3, 0x7f, 0xe7, 0x3d, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x95, 0xef, 0x89, 0xa9,
	0x11, 0x02, 0x00, 0x00,
}