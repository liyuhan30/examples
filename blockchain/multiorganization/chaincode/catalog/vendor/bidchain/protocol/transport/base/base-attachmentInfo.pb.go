// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base-attachmentInfo.proto

package base

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

type AttachmentInfo struct {
	IpfsFileHashId       string   `protobuf:"bytes,1,opt,name=ipfsFileHashId,proto3" json:"ipfsFileHashId,omitempty"`
	FileHash             string   `protobuf:"bytes,2,opt,name=fileHash,proto3" json:"fileHash,omitempty"`
	HashMethod           string   `protobuf:"bytes,3,opt,name=hashMethod,proto3" json:"hashMethod,omitempty"`
	ChannelName          string   `protobuf:"bytes,4,opt,name=channelName,proto3" json:"channelName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttachmentInfo) Reset()         { *m = AttachmentInfo{} }
func (m *AttachmentInfo) String() string { return proto.CompactTextString(m) }
func (*AttachmentInfo) ProtoMessage()    {}
func (*AttachmentInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ac4789b7c58e02a, []int{0}
}

func (m *AttachmentInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttachmentInfo.Unmarshal(m, b)
}
func (m *AttachmentInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttachmentInfo.Marshal(b, m, deterministic)
}
func (m *AttachmentInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttachmentInfo.Merge(m, src)
}
func (m *AttachmentInfo) XXX_Size() int {
	return xxx_messageInfo_AttachmentInfo.Size(m)
}
func (m *AttachmentInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AttachmentInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AttachmentInfo proto.InternalMessageInfo

func (m *AttachmentInfo) GetIpfsFileHashId() string {
	if m != nil {
		return m.IpfsFileHashId
	}
	return ""
}

func (m *AttachmentInfo) GetFileHash() string {
	if m != nil {
		return m.FileHash
	}
	return ""
}

func (m *AttachmentInfo) GetHashMethod() string {
	if m != nil {
		return m.HashMethod
	}
	return ""
}

func (m *AttachmentInfo) GetChannelName() string {
	if m != nil {
		return m.ChannelName
	}
	return ""
}

func init() {
	proto.RegisterType((*AttachmentInfo)(nil), "AttachmentInfo")
}

func init() { proto.RegisterFile("base-attachmentInfo.proto", fileDescriptor_6ac4789b7c58e02a) }

var fileDescriptor_6ac4789b7c58e02a = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x4a, 0x2c, 0x4e,
	0xd5, 0x4d, 0x2c, 0x29, 0x49, 0x4c, 0xce, 0xc8, 0x4d, 0xcd, 0x2b, 0xf1, 0xcc, 0x4b, 0xcb, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x9a, 0xc6, 0xc8, 0xc5, 0xe7, 0x88, 0x22, 0x21, 0xa4, 0xc6,
	0xc5, 0x97, 0x59, 0x90, 0x56, 0xec, 0x96, 0x99, 0x93, 0xea, 0x91, 0x58, 0x9c, 0xe1, 0x99, 0x22,
	0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x84, 0x26, 0x2a, 0x24, 0xc5, 0xc5, 0x91, 0x06, 0xe5, 0x49,
	0x30, 0x81, 0x55, 0xc0, 0xf9, 0x42, 0x72, 0x5c, 0x5c, 0x19, 0x89, 0xc5, 0x19, 0xbe, 0xa9, 0x25,
	0x19, 0xf9, 0x29, 0x12, 0xcc, 0x60, 0x59, 0x24, 0x11, 0x21, 0x05, 0x2e, 0xee, 0xe4, 0x8c, 0xc4,
	0xbc, 0xbc, 0xd4, 0x1c, 0xbf, 0xc4, 0xdc, 0x54, 0x09, 0x16, 0xb0, 0x02, 0x64, 0x21, 0x27, 0xc9,
	0x28, 0x71, 0xb0, 0x0b, 0x93, 0xf3, 0x73, 0xf4, 0x4b, 0x8a, 0x12, 0xf3, 0x8a, 0x0b, 0xf2, 0x8b,
	0x4a, 0xf4, 0x41, 0x1e, 0x49, 0x62, 0x03, 0x4b, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x3c,
	0xb5, 0xb0, 0x19, 0xd7, 0x00, 0x00, 0x00,
}