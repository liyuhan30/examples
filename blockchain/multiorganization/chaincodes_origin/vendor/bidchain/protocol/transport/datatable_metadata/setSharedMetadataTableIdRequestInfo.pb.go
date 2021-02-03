// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datatable_metadata-setSharedMetadataTableIdRequestInfo.proto

package datatable_metadata

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

type SetSharedMetadataTableIdRequestInfo struct {
	PlatformId           string   `protobuf:"bytes,1,opt,name=platformId,proto3" json:"platformId,omitempty"`
	TableIdList          []string `protobuf:"bytes,2,rep,name=tableIdList,proto3" json:"tableIdList,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetSharedMetadataTableIdRequestInfo) Reset()         { *m = SetSharedMetadataTableIdRequestInfo{} }
func (m *SetSharedMetadataTableIdRequestInfo) String() string { return proto.CompactTextString(m) }
func (*SetSharedMetadataTableIdRequestInfo) ProtoMessage()    {}
func (*SetSharedMetadataTableIdRequestInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b253e1134f7fba1, []int{0}
}

func (m *SetSharedMetadataTableIdRequestInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetSharedMetadataTableIdRequestInfo.Unmarshal(m, b)
}
func (m *SetSharedMetadataTableIdRequestInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetSharedMetadataTableIdRequestInfo.Marshal(b, m, deterministic)
}
func (m *SetSharedMetadataTableIdRequestInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetSharedMetadataTableIdRequestInfo.Merge(m, src)
}
func (m *SetSharedMetadataTableIdRequestInfo) XXX_Size() int {
	return xxx_messageInfo_SetSharedMetadataTableIdRequestInfo.Size(m)
}
func (m *SetSharedMetadataTableIdRequestInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SetSharedMetadataTableIdRequestInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SetSharedMetadataTableIdRequestInfo proto.InternalMessageInfo

func (m *SetSharedMetadataTableIdRequestInfo) GetPlatformId() string {
	if m != nil {
		return m.PlatformId
	}
	return ""
}

func (m *SetSharedMetadataTableIdRequestInfo) GetTableIdList() []string {
	if m != nil {
		return m.TableIdList
	}
	return nil
}

func init() {
	proto.RegisterType((*SetSharedMetadataTableIdRequestInfo)(nil), "SetSharedMetadataTableIdRequestInfo")
}

func init() {
	proto.RegisterFile("datatable_metadata-setSharedMetadataTableIdRequestInfo.proto", fileDescriptor_8b253e1134f7fba1)
}

var fileDescriptor_8b253e1134f7fba1 = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x49, 0x49, 0x2c, 0x49,
	0x2c, 0x49, 0x4c, 0xca, 0x49, 0x8d, 0xcf, 0x4d, 0x2d, 0x49, 0x04, 0xf1, 0x74, 0x8b, 0x53, 0x4b,
	0x82, 0x33, 0x12, 0x8b, 0x52, 0x53, 0x7c, 0xa1, 0x22, 0x21, 0x20, 0x79, 0xcf, 0x94, 0xa0, 0xd4,
	0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0xcf, 0xbc, 0xb4, 0x7c, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0xa5,
	0x74, 0x2e, 0xe5, 0x60, 0xc2, 0x8a, 0x85, 0xe4, 0xb8, 0xb8, 0x0a, 0x72, 0x12, 0x4b, 0xd2, 0xf2,
	0x8b, 0x72, 0x3d, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x90, 0x44, 0x84, 0x14, 0xb8,
	0xb8, 0x4b, 0x20, 0xba, 0x7c, 0x32, 0x8b, 0x4b, 0x24, 0x98, 0x14, 0x98, 0x35, 0x38, 0x83, 0x90,
	0x85, 0x9c, 0x0c, 0xa2, 0xf4, 0x92, 0x32, 0x53, 0x92, 0x33, 0x12, 0x33, 0xf3, 0xf4, 0xc1, 0x56,
	0x27, 0xe7, 0xe7, 0xe8, 0x97, 0x14, 0x25, 0xe6, 0x15, 0x17, 0xe4, 0x17, 0x95, 0xe8, 0x63, 0x7a,
	0x22, 0x89, 0x0d, 0xac, 0xcc, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x61, 0x4d, 0xe5, 0x60, 0xe1,
	0x00, 0x00, 0x00,
}