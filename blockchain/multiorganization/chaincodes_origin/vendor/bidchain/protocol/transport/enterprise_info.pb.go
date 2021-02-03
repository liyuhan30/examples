// Code generated by protoc-gen-go. DO NOT EDIT.
// source: enterprise_info.proto

package transport

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// 副本企业信息
type EnterpriseInfo struct {
	// 基本信息
	PublicInfo *PublicInfo `protobuf:"bytes,1,opt,name=publicInfo,proto3" json:"publicInfo,omitempty"`
	// 副本加密字段信息
	EncryptedMirrorList []*EnterpriseInfo_EncryptedMirrorDataInfoPair `protobuf:"bytes,2,rep,name=encryptedMirrorList,proto3" json:"encryptedMirrorList,omitempty"`
	// 平台信息
	PlatformId string `protobuf:"bytes,3,opt,name=platformId,proto3" json:"platformId,omitempty"`
	// 组织结构代码(作为企业的唯一id)
	OrgCode              string   `protobuf:"bytes,4,opt,name=orgCode,proto3" json:"orgCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnterpriseInfo) Reset()         { *m = EnterpriseInfo{} }
func (m *EnterpriseInfo) String() string { return proto.CompactTextString(m) }
func (*EnterpriseInfo) ProtoMessage()    {}
func (*EnterpriseInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac7e7634d19ce216, []int{0}
}

func (m *EnterpriseInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnterpriseInfo.Unmarshal(m, b)
}
func (m *EnterpriseInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnterpriseInfo.Marshal(b, m, deterministic)
}
func (m *EnterpriseInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnterpriseInfo.Merge(m, src)
}
func (m *EnterpriseInfo) XXX_Size() int {
	return xxx_messageInfo_EnterpriseInfo.Size(m)
}
func (m *EnterpriseInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_EnterpriseInfo.DiscardUnknown(m)
}

var xxx_messageInfo_EnterpriseInfo proto.InternalMessageInfo

func (m *EnterpriseInfo) GetPublicInfo() *PublicInfo {
	if m != nil {
		return m.PublicInfo
	}
	return nil
}

func (m *EnterpriseInfo) GetEncryptedMirrorList() []*EnterpriseInfo_EncryptedMirrorDataInfoPair {
	if m != nil {
		return m.EncryptedMirrorList
	}
	return nil
}

func (m *EnterpriseInfo) GetPlatformId() string {
	if m != nil {
		return m.PlatformId
	}
	return ""
}

func (m *EnterpriseInfo) GetOrgCode() string {
	if m != nil {
		return m.OrgCode
	}
	return ""
}

type EnterpriseInfo_EncryptedMirrorDataInfoPair struct {
	Title                string               `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Data                 *EncryptedMirrorData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *EnterpriseInfo_EncryptedMirrorDataInfoPair) Reset() {
	*m = EnterpriseInfo_EncryptedMirrorDataInfoPair{}
}
func (m *EnterpriseInfo_EncryptedMirrorDataInfoPair) String() string {
	return proto.CompactTextString(m)
}
func (*EnterpriseInfo_EncryptedMirrorDataInfoPair) ProtoMessage() {}
func (*EnterpriseInfo_EncryptedMirrorDataInfoPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac7e7634d19ce216, []int{0, 0}
}

func (m *EnterpriseInfo_EncryptedMirrorDataInfoPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnterpriseInfo_EncryptedMirrorDataInfoPair.Unmarshal(m, b)
}
func (m *EnterpriseInfo_EncryptedMirrorDataInfoPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnterpriseInfo_EncryptedMirrorDataInfoPair.Marshal(b, m, deterministic)
}
func (m *EnterpriseInfo_EncryptedMirrorDataInfoPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnterpriseInfo_EncryptedMirrorDataInfoPair.Merge(m, src)
}
func (m *EnterpriseInfo_EncryptedMirrorDataInfoPair) XXX_Size() int {
	return xxx_messageInfo_EnterpriseInfo_EncryptedMirrorDataInfoPair.Size(m)
}
func (m *EnterpriseInfo_EncryptedMirrorDataInfoPair) XXX_DiscardUnknown() {
	xxx_messageInfo_EnterpriseInfo_EncryptedMirrorDataInfoPair.DiscardUnknown(m)
}

var xxx_messageInfo_EnterpriseInfo_EncryptedMirrorDataInfoPair proto.InternalMessageInfo

func (m *EnterpriseInfo_EncryptedMirrorDataInfoPair) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *EnterpriseInfo_EncryptedMirrorDataInfoPair) GetData() *EncryptedMirrorData {
	if m != nil {
		return m.Data
	}
	return nil
}

// 添加企业信息对应的输入信息
type AddEnterpriseInfoRequest struct {
	InfoList             []*EnterpriseInfo `protobuf:"bytes,1,rep,name=infoList,proto3" json:"infoList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AddEnterpriseInfoRequest) Reset()         { *m = AddEnterpriseInfoRequest{} }
func (m *AddEnterpriseInfoRequest) String() string { return proto.CompactTextString(m) }
func (*AddEnterpriseInfoRequest) ProtoMessage()    {}
func (*AddEnterpriseInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac7e7634d19ce216, []int{1}
}

func (m *AddEnterpriseInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddEnterpriseInfoRequest.Unmarshal(m, b)
}
func (m *AddEnterpriseInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddEnterpriseInfoRequest.Marshal(b, m, deterministic)
}
func (m *AddEnterpriseInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddEnterpriseInfoRequest.Merge(m, src)
}
func (m *AddEnterpriseInfoRequest) XXX_Size() int {
	return xxx_messageInfo_AddEnterpriseInfoRequest.Size(m)
}
func (m *AddEnterpriseInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddEnterpriseInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddEnterpriseInfoRequest proto.InternalMessageInfo

func (m *AddEnterpriseInfoRequest) GetInfoList() []*EnterpriseInfo {
	if m != nil {
		return m.InfoList
	}
	return nil
}

// 添加企业信息的响应信息
type AddEnterpriseInfoResponse struct {
	// enterpriseMirrorId: enterpriseId
	InfoList             []*AddEnterpriseInfoResponseInfo `protobuf:"bytes,1,rep,name=infoList,proto3" json:"infoList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *AddEnterpriseInfoResponse) Reset()         { *m = AddEnterpriseInfoResponse{} }
func (m *AddEnterpriseInfoResponse) String() string { return proto.CompactTextString(m) }
func (*AddEnterpriseInfoResponse) ProtoMessage()    {}
func (*AddEnterpriseInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac7e7634d19ce216, []int{2}
}

func (m *AddEnterpriseInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddEnterpriseInfoResponse.Unmarshal(m, b)
}
func (m *AddEnterpriseInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddEnterpriseInfoResponse.Marshal(b, m, deterministic)
}
func (m *AddEnterpriseInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddEnterpriseInfoResponse.Merge(m, src)
}
func (m *AddEnterpriseInfoResponse) XXX_Size() int {
	return xxx_messageInfo_AddEnterpriseInfoResponse.Size(m)
}
func (m *AddEnterpriseInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddEnterpriseInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddEnterpriseInfoResponse proto.InternalMessageInfo

func (m *AddEnterpriseInfoResponse) GetInfoList() []*AddEnterpriseInfoResponseInfo {
	if m != nil {
		return m.InfoList
	}
	return nil
}

type AddEnterpriseInfoResponseInfo struct {
	EnterpriseMirrorId   string   `protobuf:"bytes,1,opt,name=enterpriseMirrorId,proto3" json:"enterpriseMirrorId,omitempty"`
	EnterpriseId         string   `protobuf:"bytes,2,opt,name=enterpriseId,proto3" json:"enterpriseId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddEnterpriseInfoResponseInfo) Reset()         { *m = AddEnterpriseInfoResponseInfo{} }
func (m *AddEnterpriseInfoResponseInfo) String() string { return proto.CompactTextString(m) }
func (*AddEnterpriseInfoResponseInfo) ProtoMessage()    {}
func (*AddEnterpriseInfoResponseInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac7e7634d19ce216, []int{2, 0}
}

func (m *AddEnterpriseInfoResponseInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddEnterpriseInfoResponseInfo.Unmarshal(m, b)
}
func (m *AddEnterpriseInfoResponseInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddEnterpriseInfoResponseInfo.Marshal(b, m, deterministic)
}
func (m *AddEnterpriseInfoResponseInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddEnterpriseInfoResponseInfo.Merge(m, src)
}
func (m *AddEnterpriseInfoResponseInfo) XXX_Size() int {
	return xxx_messageInfo_AddEnterpriseInfoResponseInfo.Size(m)
}
func (m *AddEnterpriseInfoResponseInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AddEnterpriseInfoResponseInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AddEnterpriseInfoResponseInfo proto.InternalMessageInfo

func (m *AddEnterpriseInfoResponseInfo) GetEnterpriseMirrorId() string {
	if m != nil {
		return m.EnterpriseMirrorId
	}
	return ""
}

func (m *AddEnterpriseInfoResponseInfo) GetEnterpriseId() string {
	if m != nil {
		return m.EnterpriseId
	}
	return ""
}

// 添加企业时发出的事件
// 事件名称: addNewEnterprises
type EnterpriseInfoEvent struct {
	InfoList             []*EnterpriseInfo `protobuf:"bytes,1,rep,name=infoList,proto3" json:"infoList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *EnterpriseInfoEvent) Reset()         { *m = EnterpriseInfoEvent{} }
func (m *EnterpriseInfoEvent) String() string { return proto.CompactTextString(m) }
func (*EnterpriseInfoEvent) ProtoMessage()    {}
func (*EnterpriseInfoEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac7e7634d19ce216, []int{3}
}

func (m *EnterpriseInfoEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnterpriseInfoEvent.Unmarshal(m, b)
}
func (m *EnterpriseInfoEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnterpriseInfoEvent.Marshal(b, m, deterministic)
}
func (m *EnterpriseInfoEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnterpriseInfoEvent.Merge(m, src)
}
func (m *EnterpriseInfoEvent) XXX_Size() int {
	return xxx_messageInfo_EnterpriseInfoEvent.Size(m)
}
func (m *EnterpriseInfoEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_EnterpriseInfoEvent.DiscardUnknown(m)
}

var xxx_messageInfo_EnterpriseInfoEvent proto.InternalMessageInfo

func (m *EnterpriseInfoEvent) GetInfoList() []*EnterpriseInfo {
	if m != nil {
		return m.InfoList
	}
	return nil
}

// 获取某些企业所有最新副本信息 时
// 获取某些企业所有历史副本信息
type GetEnterpriseInfoRequest struct {
	EnterpriseIds        []string `protobuf:"bytes,1,rep,name=enterpriseIds,proto3" json:"enterpriseIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEnterpriseInfoRequest) Reset()         { *m = GetEnterpriseInfoRequest{} }
func (m *GetEnterpriseInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetEnterpriseInfoRequest) ProtoMessage()    {}
func (*GetEnterpriseInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac7e7634d19ce216, []int{4}
}

func (m *GetEnterpriseInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEnterpriseInfoRequest.Unmarshal(m, b)
}
func (m *GetEnterpriseInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEnterpriseInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetEnterpriseInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEnterpriseInfoRequest.Merge(m, src)
}
func (m *GetEnterpriseInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetEnterpriseInfoRequest.Size(m)
}
func (m *GetEnterpriseInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEnterpriseInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetEnterpriseInfoRequest proto.InternalMessageInfo

func (m *GetEnterpriseInfoRequest) GetEnterpriseIds() []string {
	if m != nil {
		return m.EnterpriseIds
	}
	return nil
}

// 获取某些企业所有最新副本信息
// 获取某些企业所有历史部分信息
type GetEnterpriseInfoResponse struct {
	Infos                []*GetEnterpriseInfoResponseInfo `protobuf:"bytes,1,rep,name=infos,proto3" json:"infos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *GetEnterpriseInfoResponse) Reset()         { *m = GetEnterpriseInfoResponse{} }
func (m *GetEnterpriseInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetEnterpriseInfoResponse) ProtoMessage()    {}
func (*GetEnterpriseInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac7e7634d19ce216, []int{5}
}

func (m *GetEnterpriseInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEnterpriseInfoResponse.Unmarshal(m, b)
}
func (m *GetEnterpriseInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEnterpriseInfoResponse.Marshal(b, m, deterministic)
}
func (m *GetEnterpriseInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEnterpriseInfoResponse.Merge(m, src)
}
func (m *GetEnterpriseInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetEnterpriseInfoResponse.Size(m)
}
func (m *GetEnterpriseInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEnterpriseInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetEnterpriseInfoResponse proto.InternalMessageInfo

func (m *GetEnterpriseInfoResponse) GetInfos() []*GetEnterpriseInfoResponseInfo {
	if m != nil {
		return m.Infos
	}
	return nil
}

type GetEnterpriseInfoResponseInfo struct {
	// 企业id
	EnterpriseId string `protobuf:"bytes,1,opt,name=enterpriseId,proto3" json:"enterpriseId,omitempty"`
	// 企业信息
	EnterpriseInfos      []*EnterpriseResponseInfo `protobuf:"bytes,2,rep,name=enterpriseInfos,proto3" json:"enterpriseInfos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *GetEnterpriseInfoResponseInfo) Reset()         { *m = GetEnterpriseInfoResponseInfo{} }
func (m *GetEnterpriseInfoResponseInfo) String() string { return proto.CompactTextString(m) }
func (*GetEnterpriseInfoResponseInfo) ProtoMessage()    {}
func (*GetEnterpriseInfoResponseInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac7e7634d19ce216, []int{5, 0}
}

func (m *GetEnterpriseInfoResponseInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEnterpriseInfoResponseInfo.Unmarshal(m, b)
}
func (m *GetEnterpriseInfoResponseInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEnterpriseInfoResponseInfo.Marshal(b, m, deterministic)
}
func (m *GetEnterpriseInfoResponseInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEnterpriseInfoResponseInfo.Merge(m, src)
}
func (m *GetEnterpriseInfoResponseInfo) XXX_Size() int {
	return xxx_messageInfo_GetEnterpriseInfoResponseInfo.Size(m)
}
func (m *GetEnterpriseInfoResponseInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEnterpriseInfoResponseInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GetEnterpriseInfoResponseInfo proto.InternalMessageInfo

func (m *GetEnterpriseInfoResponseInfo) GetEnterpriseId() string {
	if m != nil {
		return m.EnterpriseId
	}
	return ""
}

func (m *GetEnterpriseInfoResponseInfo) GetEnterpriseInfos() []*EnterpriseResponseInfo {
	if m != nil {
		return m.EnterpriseInfos
	}
	return nil
}

type EnterpriseResponseInfo struct {
	// 企业信息
	EnterpriseInfo *EnterpriseInfo `protobuf:"bytes,1,opt,name=enterpriseInfo,proto3" json:"enterpriseInfo,omitempty"`
	// 更新时间
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	// 交易id
	TxId                 string   `protobuf:"bytes,4,opt,name=txId,proto3" json:"txId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnterpriseResponseInfo) Reset()         { *m = EnterpriseResponseInfo{} }
func (m *EnterpriseResponseInfo) String() string { return proto.CompactTextString(m) }
func (*EnterpriseResponseInfo) ProtoMessage()    {}
func (*EnterpriseResponseInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac7e7634d19ce216, []int{6}
}

func (m *EnterpriseResponseInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnterpriseResponseInfo.Unmarshal(m, b)
}
func (m *EnterpriseResponseInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnterpriseResponseInfo.Marshal(b, m, deterministic)
}
func (m *EnterpriseResponseInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnterpriseResponseInfo.Merge(m, src)
}
func (m *EnterpriseResponseInfo) XXX_Size() int {
	return xxx_messageInfo_EnterpriseResponseInfo.Size(m)
}
func (m *EnterpriseResponseInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_EnterpriseResponseInfo.DiscardUnknown(m)
}

var xxx_messageInfo_EnterpriseResponseInfo proto.InternalMessageInfo

func (m *EnterpriseResponseInfo) GetEnterpriseInfo() *EnterpriseInfo {
	if m != nil {
		return m.EnterpriseInfo
	}
	return nil
}

func (m *EnterpriseResponseInfo) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *EnterpriseResponseInfo) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func init() {
	proto.RegisterType((*EnterpriseInfo)(nil), "EnterpriseInfo")
	proto.RegisterType((*EnterpriseInfo_EncryptedMirrorDataInfoPair)(nil), "EnterpriseInfo.EncryptedMirrorDataInfoPair")
	proto.RegisterType((*AddEnterpriseInfoRequest)(nil), "AddEnterpriseInfoRequest")
	proto.RegisterType((*AddEnterpriseInfoResponse)(nil), "AddEnterpriseInfoResponse")
	proto.RegisterType((*AddEnterpriseInfoResponseInfo)(nil), "AddEnterpriseInfoResponse.info")
	proto.RegisterType((*EnterpriseInfoEvent)(nil), "EnterpriseInfoEvent")
	proto.RegisterType((*GetEnterpriseInfoRequest)(nil), "GetEnterpriseInfoRequest")
	proto.RegisterType((*GetEnterpriseInfoResponse)(nil), "GetEnterpriseInfoResponse")
	proto.RegisterType((*GetEnterpriseInfoResponseInfo)(nil), "GetEnterpriseInfoResponse.info")
	proto.RegisterType((*EnterpriseResponseInfo)(nil), "EnterpriseResponseInfo")
}

func init() { proto.RegisterFile("enterprise_info.proto", fileDescriptor_ac7e7634d19ce216) }

var fileDescriptor_ac7e7634d19ce216 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xd1, 0x6a, 0xdb, 0x30,
	0x14, 0xc5, 0x69, 0xba, 0xcd, 0x37, 0x5b, 0x3b, 0xd4, 0x6c, 0x73, 0x3d, 0x58, 0x83, 0xd9, 0x43,
	0x20, 0xa0, 0x40, 0xc7, 0x18, 0x6c, 0x2f, 0x6b, 0xb7, 0x50, 0x0c, 0x1b, 0x14, 0xd1, 0xa7, 0x42,
	0x29, 0x4a, 0xa4, 0x04, 0x41, 0x2c, 0x79, 0x92, 0x3c, 0xb6, 0xcf, 0xd9, 0x3f, 0x6c, 0x1f, 0xb0,
	0x3f, 0x1b, 0x96, 0x95, 0xc4, 0x4e, 0xdc, 0x41, 0xdf, 0x74, 0x8f, 0x8e, 0xcf, 0xd5, 0x39, 0x96,
	0x2e, 0x3c, 0xe3, 0xd2, 0x72, 0x9d, 0x6b, 0x61, 0xf8, 0xad, 0x90, 0x73, 0x85, 0x73, 0xad, 0xac,
	0x8a, 0x4f, 0x16, 0x4a, 0x2d, 0x96, 0x7c, 0xec, 0xaa, 0x69, 0x31, 0x1f, 0x5b, 0x91, 0x71, 0x63,
	0x69, 0x96, 0x7b, 0xc2, 0xd3, 0xc9, 0xfa, 0x3b, 0x8f, 0xf4, 0x67, 0xfa, 0x67, 0x6e, 0xd5, 0x6d,
	0xc6, 0x8d, 0xa1, 0x0b, 0x8f, 0x26, 0x7f, 0x3a, 0x70, 0xb0, 0xa1, 0xa6, 0x72, 0xae, 0xd0, 0x08,
	0x20, 0x2f, 0xa6, 0x4b, 0x31, 0x2b, 0xab, 0x28, 0x18, 0x04, 0xc3, 0xde, 0x69, 0x0f, 0x5f, 0xae,
	0x21, 0x52, 0xdb, 0x46, 0x37, 0x70, 0xc4, 0xa5, 0x53, 0xe6, 0xec, 0xab, 0xd0, 0x5a, 0xe9, 0x2f,
	0xc2, 0xd8, 0xa8, 0x33, 0xd8, 0x1b, 0xf6, 0x4e, 0x47, 0xb8, 0x29, 0x8d, 0x27, 0x4d, 0xea, 0x67,
	0x6a, 0x69, 0x89, 0x5f, 0x52, 0xa1, 0x49, 0x9b, 0x0e, 0x7a, 0x05, 0x90, 0x2f, 0xa9, 0x9d, 0x2b,
	0x9d, 0xa5, 0x2c, 0xda, 0x1b, 0x04, 0xc3, 0x90, 0xd4, 0x10, 0x14, 0xc1, 0x43, 0xa5, 0x17, 0x9f,
	0x14, 0xe3, 0x51, 0xd7, 0x6d, 0xae, 0xca, 0xf8, 0x06, 0x5e, 0xfe, 0xa7, 0x1b, 0xea, 0xc3, 0xbe,
	0x15, 0x76, 0xc9, 0x9d, 0xbf, 0x90, 0x54, 0x05, 0x1a, 0x42, 0x97, 0x51, 0x4b, 0xa3, 0x8e, 0x33,
	0xdd, 0x6f, 0x3b, 0x2f, 0x71, 0x8c, 0xe4, 0x02, 0xa2, 0x33, 0xc6, 0x9a, 0xf6, 0x08, 0xff, 0x56,
	0x70, 0x63, 0xd1, 0x08, 0x1e, 0x95, 0xbf, 0xca, 0x05, 0x11, 0xb8, 0x20, 0x0e, 0xb7, 0x82, 0x20,
	0x6b, 0x42, 0xf2, 0x3b, 0x80, 0xe3, 0x16, 0x25, 0x93, 0x2b, 0x69, 0x38, 0xfa, 0xb0, 0x23, 0x75,
	0x82, 0xef, 0x64, 0x63, 0xd1, 0x90, 0x8e, 0xaf, 0xa1, 0x5b, 0xae, 0x11, 0x06, 0xb4, 0xb9, 0x45,
	0x95, 0x93, 0x94, 0x79, 0xe3, 0x2d, 0x3b, 0x28, 0x81, 0xc7, 0x1b, 0x34, 0x65, 0x2e, 0x8d, 0x90,
	0x34, 0xb0, 0xe4, 0x1c, 0x8e, 0x9a, 0x87, 0x98, 0x7c, 0xe7, 0xf2, 0x9e, 0xd6, 0x3f, 0x42, 0x74,
	0xc1, 0x6d, 0x7b, 0x86, 0xaf, 0xe1, 0x49, 0xbd, 0x9f, 0x71, 0x6a, 0x21, 0x69, 0x82, 0xc9, 0xdf,
	0x00, 0x8e, 0x5b, 0x24, 0x7c, 0x78, 0x6f, 0x61, 0xbf, 0xec, 0x65, 0xd6, 0xc9, 0xdd, 0x49, 0xad,
	0x92, 0xab, 0xd8, 0x71, 0xe6, 0x63, 0xdb, 0x8e, 0x21, 0xd8, 0x8d, 0x01, 0x9d, 0xc1, 0x21, 0x6f,
	0x28, 0x1a, 0x7f, 0xf5, 0x5f, 0xd4, 0x6c, 0xaf, 0xba, 0xb8, 0x8e, 0xdb, 0xfc, 0xe4, 0x57, 0x00,
	0xcf, 0xdb, 0xb9, 0xe8, 0x1d, 0x1c, 0x34, 0xd9, 0xfe, 0x35, 0xee, 0x64, 0xba, 0x45, 0x43, 0xef,
	0x01, 0x8a, 0x9c, 0x51, 0xcb, 0xaf, 0x44, 0xc6, 0xfd, 0x6d, 0x8e, 0x71, 0x35, 0x33, 0xf0, 0x6a,
	0x66, 0xe0, 0xab, 0xd5, 0xcc, 0x20, 0x35, 0x36, 0x42, 0xd0, 0xb5, 0x3f, 0x52, 0xe6, 0xdf, 0x93,
	0x5b, 0x9f, 0x8f, 0x21, 0x99, 0x49, 0x3c, 0x15, 0xcc, 0x14, 0x12, 0x67, 0xdc, 0xd2, 0xf2, 0x0d,
	0xe0, 0x99, 0xca, 0x32, 0x25, 0x2b, 0x3d, 0x46, 0xd5, 0x75, 0x68, 0x35, 0x95, 0x26, 0x57, 0xda,
	0x4e, 0x1f, 0x38, 0xf0, 0xcd, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x66, 0xce, 0xbe, 0xbf,
	0x04, 0x00, 0x00,
}