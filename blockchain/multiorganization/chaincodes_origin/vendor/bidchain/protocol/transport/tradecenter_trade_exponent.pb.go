// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tradecenter_trade_exponent.proto

package transport

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

// 交易中心交易指数
type TradeCenterTradeExponent struct {
	// 当日交易指数
	TradeExp string `protobuf:"bytes,1,opt,name=tradeExp,proto3" json:"tradeExp,omitempty"`
	// 时间戳
	TimeStamp string `protobuf:"bytes,2,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
	// 交易金额-金额
	TradeAmount string `protobuf:"bytes,3,opt,name=tradeAmount,proto3" json:"tradeAmount,omitempty"`
	// 交易金额-指数
	TradeAmountExp string `protobuf:"bytes,4,opt,name=tradeAmountExp,proto3" json:"tradeAmountExp,omitempty"`
	// 交易金额-占比
	TradeAmountProportion string `protobuf:"bytes,5,opt,name=tradeAmountProportion,proto3" json:"tradeAmountProportion,omitempty"`
	// 交易项目-数量
	TradeProProCount string `protobuf:"bytes,6,opt,name=tradeProProCount,proto3" json:"tradeProProCount,omitempty"`
	// 交易项目-指数
	TradeProExp string `protobuf:"bytes,7,opt,name=tradeProExp,proto3" json:"tradeProExp,omitempty"`
	// 交易项目-占比
	TradeProProportion string `protobuf:"bytes,8,opt,name=tradeProProportion,proto3" json:"tradeProProportion,omitempty"`
	// 节约资金-金额
	TradeSavingAmount string `protobuf:"bytes,9,opt,name=tradeSavingAmount,proto3" json:"tradeSavingAmount,omitempty"`
	// 节约资金-指数
	TradeSavingExp string `protobuf:"bytes,10,opt,name=tradeSavingExp,proto3" json:"tradeSavingExp,omitempty"`
	// 节约资金-占比
	TradeSavingProportion string `protobuf:"bytes,11,opt,name=tradeSavingProportion,proto3" json:"tradeSavingProportion,omitempty"`
	// 投标对象-数量
	TradeBidderCount string `protobuf:"bytes,12,opt,name=tradeBidderCount,proto3" json:"tradeBidderCount,omitempty"`
	// 投标对象-指数
	TradeBidderExp string `protobuf:"bytes,13,opt,name=tradeBidderExp,proto3" json:"tradeBidderExp,omitempty"`
	// 投标对象-占比
	TradeBidderProportion string `protobuf:"bytes,14,opt,name=tradeBidderProportion,proto3" json:"tradeBidderProportion,omitempty"`
	// 平台id
	PlatformId           string   `protobuf:"bytes,15,opt,name=platformId,proto3" json:"platformId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TradeCenterTradeExponent) Reset()         { *m = TradeCenterTradeExponent{} }
func (m *TradeCenterTradeExponent) String() string { return proto.CompactTextString(m) }
func (*TradeCenterTradeExponent) ProtoMessage()    {}
func (*TradeCenterTradeExponent) Descriptor() ([]byte, []int) {
	return fileDescriptor_51b0796c4aeef7e1, []int{0}
}

func (m *TradeCenterTradeExponent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradeCenterTradeExponent.Unmarshal(m, b)
}
func (m *TradeCenterTradeExponent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradeCenterTradeExponent.Marshal(b, m, deterministic)
}
func (m *TradeCenterTradeExponent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradeCenterTradeExponent.Merge(m, src)
}
func (m *TradeCenterTradeExponent) XXX_Size() int {
	return xxx_messageInfo_TradeCenterTradeExponent.Size(m)
}
func (m *TradeCenterTradeExponent) XXX_DiscardUnknown() {
	xxx_messageInfo_TradeCenterTradeExponent.DiscardUnknown(m)
}

var xxx_messageInfo_TradeCenterTradeExponent proto.InternalMessageInfo

func (m *TradeCenterTradeExponent) GetTradeExp() string {
	if m != nil {
		return m.TradeExp
	}
	return ""
}

func (m *TradeCenterTradeExponent) GetTimeStamp() string {
	if m != nil {
		return m.TimeStamp
	}
	return ""
}

func (m *TradeCenterTradeExponent) GetTradeAmount() string {
	if m != nil {
		return m.TradeAmount
	}
	return ""
}

func (m *TradeCenterTradeExponent) GetTradeAmountExp() string {
	if m != nil {
		return m.TradeAmountExp
	}
	return ""
}

func (m *TradeCenterTradeExponent) GetTradeAmountProportion() string {
	if m != nil {
		return m.TradeAmountProportion
	}
	return ""
}

func (m *TradeCenterTradeExponent) GetTradeProProCount() string {
	if m != nil {
		return m.TradeProProCount
	}
	return ""
}

func (m *TradeCenterTradeExponent) GetTradeProExp() string {
	if m != nil {
		return m.TradeProExp
	}
	return ""
}

func (m *TradeCenterTradeExponent) GetTradeProProportion() string {
	if m != nil {
		return m.TradeProProportion
	}
	return ""
}

func (m *TradeCenterTradeExponent) GetTradeSavingAmount() string {
	if m != nil {
		return m.TradeSavingAmount
	}
	return ""
}

func (m *TradeCenterTradeExponent) GetTradeSavingExp() string {
	if m != nil {
		return m.TradeSavingExp
	}
	return ""
}

func (m *TradeCenterTradeExponent) GetTradeSavingProportion() string {
	if m != nil {
		return m.TradeSavingProportion
	}
	return ""
}

func (m *TradeCenterTradeExponent) GetTradeBidderCount() string {
	if m != nil {
		return m.TradeBidderCount
	}
	return ""
}

func (m *TradeCenterTradeExponent) GetTradeBidderExp() string {
	if m != nil {
		return m.TradeBidderExp
	}
	return ""
}

func (m *TradeCenterTradeExponent) GetTradeBidderProportion() string {
	if m != nil {
		return m.TradeBidderProportion
	}
	return ""
}

func (m *TradeCenterTradeExponent) GetPlatformId() string {
	if m != nil {
		return m.PlatformId
	}
	return ""
}

type AddTradeCenterTradeExponentRequest struct {
	TradeCenterTradeExponentList []*TradeCenterTradeExponent `protobuf:"bytes,1,rep,name=tradeCenterTradeExponentList,proto3" json:"tradeCenterTradeExponentList,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}                    `json:"-"`
	XXX_unrecognized             []byte                      `json:"-"`
	XXX_sizecache                int32                       `json:"-"`
}

func (m *AddTradeCenterTradeExponentRequest) Reset()         { *m = AddTradeCenterTradeExponentRequest{} }
func (m *AddTradeCenterTradeExponentRequest) String() string { return proto.CompactTextString(m) }
func (*AddTradeCenterTradeExponentRequest) ProtoMessage()    {}
func (*AddTradeCenterTradeExponentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_51b0796c4aeef7e1, []int{1}
}

func (m *AddTradeCenterTradeExponentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddTradeCenterTradeExponentRequest.Unmarshal(m, b)
}
func (m *AddTradeCenterTradeExponentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddTradeCenterTradeExponentRequest.Marshal(b, m, deterministic)
}
func (m *AddTradeCenterTradeExponentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddTradeCenterTradeExponentRequest.Merge(m, src)
}
func (m *AddTradeCenterTradeExponentRequest) XXX_Size() int {
	return xxx_messageInfo_AddTradeCenterTradeExponentRequest.Size(m)
}
func (m *AddTradeCenterTradeExponentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddTradeCenterTradeExponentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddTradeCenterTradeExponentRequest proto.InternalMessageInfo

func (m *AddTradeCenterTradeExponentRequest) GetTradeCenterTradeExponentList() []*TradeCenterTradeExponent {
	if m != nil {
		return m.TradeCenterTradeExponentList
	}
	return nil
}

func init() {
	proto.RegisterType((*TradeCenterTradeExponent)(nil), "TradeCenterTradeExponent")
	proto.RegisterType((*AddTradeCenterTradeExponentRequest)(nil), "AddTradeCenterTradeExponentRequest")
}

func init() { proto.RegisterFile("tradecenter_trade_exponent.proto", fileDescriptor_51b0796c4aeef7e1) }

var fileDescriptor_51b0796c4aeef7e1 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x4d, 0x4b, 0xf3, 0x40,
	0x14, 0x85, 0xc9, 0xdb, 0xb7, 0xb5, 0xb9, 0xd5, 0xaa, 0x03, 0x42, 0x94, 0x22, 0x21, 0x0b, 0x29,
	0x22, 0x11, 0xd4, 0x3f, 0xd0, 0x16, 0x17, 0x82, 0x8b, 0xd0, 0xba, 0x12, 0xa4, 0x4c, 0x3b, 0xa3,
	0x04, 0xcc, 0x4c, 0x9c, 0xdc, 0x8a, 0x7b, 0x37, 0xfe, 0x6c, 0xc9, 0x9d, 0x3a, 0x0e, 0xfd, 0x70,
	0x97, 0x79, 0xce, 0x99, 0xf6, 0xe9, 0xa1, 0x81, 0x18, 0x0d, 0x17, 0x72, 0x2e, 0x15, 0x4a, 0x33,
	0xa5, 0xe7, 0xa9, 0xfc, 0x28, 0xb5, 0x92, 0x0a, 0xd3, 0xd2, 0x68, 0xd4, 0xc9, 0x57, 0x13, 0xa2,
	0x87, 0x3a, 0x18, 0x51, 0x89, 0x1e, 0x6f, 0x97, 0x15, 0x76, 0x02, 0x6d, 0x5c, 0x82, 0x28, 0x88,
	0x83, 0x7e, 0x38, 0x76, 0x67, 0xd6, 0x83, 0x10, 0xf3, 0x42, 0x4e, 0x90, 0x17, 0x65, 0xf4, 0x8f,
	0xc2, 0x5f, 0xc0, 0x62, 0xe8, 0x50, 0x73, 0x50, 0xe8, 0x85, 0xc2, 0xa8, 0x41, 0xb9, 0x8f, 0xd8,
	0x19, 0x74, 0xbd, 0x63, 0xfd, 0x0d, 0xff, 0xa9, 0xb4, 0x42, 0xd9, 0x0d, 0x1c, 0x79, 0x24, 0x33,
	0xba, 0xd4, 0x06, 0x73, 0xad, 0xa2, 0x26, 0xd5, 0x37, 0x87, 0xec, 0x1c, 0x0e, 0x28, 0xc8, 0x8c,
	0xce, 0x8c, 0x1e, 0x91, 0x44, 0x8b, 0x2e, 0xac, 0x71, 0xe7, 0x9a, 0x19, 0x5d, 0x6b, 0xec, 0x78,
	0xae, 0x16, 0xb1, 0x14, 0x98, 0x77, 0xeb, 0x47, 0xa0, 0x4d, 0xc5, 0x0d, 0x09, 0xbb, 0x80, 0x43,
	0xa2, 0x13, 0xfe, 0x9e, 0xab, 0x97, 0xe5, 0x06, 0x21, 0xd5, 0xd7, 0x03, 0xb7, 0x84, 0x85, 0xb5,
	0x02, 0x78, 0x4b, 0x38, 0xea, 0x96, 0xb0, 0xc4, 0x13, 0xe9, 0x78, 0x4b, 0xac, 0x86, 0x6e, 0x89,
	0x61, 0x2e, 0x84, 0x34, 0x76, 0x89, 0x5d, 0x6f, 0x09, 0x8f, 0x3b, 0x13, 0xcb, 0x6a, 0x93, 0x3d,
	0xcf, 0xc4, 0x51, 0x67, 0x62, 0x89, 0x67, 0xd2, 0xf5, 0x4c, 0x56, 0x43, 0x76, 0x0a, 0x50, 0xbe,
	0x72, 0x7c, 0xd6, 0xa6, 0xb8, 0x13, 0xd1, 0x3e, 0x55, 0x3d, 0x92, 0x7c, 0x06, 0x90, 0x0c, 0x84,
	0xd8, 0xf6, 0x6f, 0x1c, 0xcb, 0xb7, 0x85, 0xac, 0x90, 0x3d, 0x41, 0x0f, 0xb7, 0x54, 0xee, 0xf3,
	0x0a, 0xa3, 0x20, 0x6e, 0xf4, 0x3b, 0x57, 0xc7, 0xe9, 0xd6, 0xcf, 0xf9, 0xf3, 0xfa, 0xf0, 0x12,
	0x92, 0xb9, 0x4a, 0x67, 0xb9, 0xa8, 0x16, 0x2a, 0x2d, 0x24, 0x72, 0xc1, 0x91, 0xa7, 0x73, 0x5d,
	0x14, 0x5a, 0xd9, 0x77, 0x46, 0x70, 0xfd, 0x18, 0xa2, 0xe1, 0xaa, 0xaa, 0x7f, 0xd9, 0xac, 0x45,
	0xf0, 0xfa, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xba, 0xdc, 0xeb, 0x80, 0x6c, 0x03, 0x00, 0x00,
}