// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/feegrant/v1beta1/feegrant.proto

package types

import (
	fmt "fmt"
	types1 "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/regen-network/cosmos-proto"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// BasicFeeAllowance implements FeeAllowance with a one-time grant of tokens
// that optionally expires. The delegatee can use up to SpendLimit to cover fees.
type BasicFeeAllowance struct {
	SpendLimit github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=spend_limit,json=spendLimit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"spend_limit"`
	Expiration ExpiresAt                                `protobuf:"bytes,2,opt,name=expiration,proto3" json:"expiration"`
}

func (m *BasicFeeAllowance) Reset()         { *m = BasicFeeAllowance{} }
func (m *BasicFeeAllowance) String() string { return proto.CompactTextString(m) }
func (*BasicFeeAllowance) ProtoMessage()    {}
func (*BasicFeeAllowance) Descriptor() ([]byte, []int) {
	return fileDescriptor_7279582900c30aea, []int{0}
}
func (m *BasicFeeAllowance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BasicFeeAllowance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BasicFeeAllowance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BasicFeeAllowance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BasicFeeAllowance.Merge(m, src)
}
func (m *BasicFeeAllowance) XXX_Size() int {
	return m.Size()
}
func (m *BasicFeeAllowance) XXX_DiscardUnknown() {
	xxx_messageInfo_BasicFeeAllowance.DiscardUnknown(m)
}

var xxx_messageInfo_BasicFeeAllowance proto.InternalMessageInfo

func (m *BasicFeeAllowance) GetSpendLimit() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.SpendLimit
	}
	return nil
}

func (m *BasicFeeAllowance) GetExpiration() ExpiresAt {
	if m != nil {
		return m.Expiration
	}
	return ExpiresAt{}
}

// PeriodicFeeAllowance extends FeeAllowance to allow for both a maximum cap,
// as well as a limit per time period.
type PeriodicFeeAllowance struct {
	Basic            BasicFeeAllowance                        `protobuf:"bytes,1,opt,name=basic,proto3" json:"basic"`
	Period           Duration                                 `protobuf:"bytes,2,opt,name=period,proto3" json:"period"`
	PeriodSpendLimit github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=period_spend_limit,json=periodSpendLimit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"period_spend_limit"`
	PeriodCanSpend   github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=period_can_spend,json=periodCanSpend,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"period_can_spend"`
	PeriodReset      ExpiresAt                                `protobuf:"bytes,5,opt,name=period_reset,json=periodReset,proto3" json:"period_reset"`
}

func (m *PeriodicFeeAllowance) Reset()         { *m = PeriodicFeeAllowance{} }
func (m *PeriodicFeeAllowance) String() string { return proto.CompactTextString(m) }
func (*PeriodicFeeAllowance) ProtoMessage()    {}
func (*PeriodicFeeAllowance) Descriptor() ([]byte, []int) {
	return fileDescriptor_7279582900c30aea, []int{1}
}
func (m *PeriodicFeeAllowance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PeriodicFeeAllowance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PeriodicFeeAllowance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PeriodicFeeAllowance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeriodicFeeAllowance.Merge(m, src)
}
func (m *PeriodicFeeAllowance) XXX_Size() int {
	return m.Size()
}
func (m *PeriodicFeeAllowance) XXX_DiscardUnknown() {
	xxx_messageInfo_PeriodicFeeAllowance.DiscardUnknown(m)
}

var xxx_messageInfo_PeriodicFeeAllowance proto.InternalMessageInfo

func (m *PeriodicFeeAllowance) GetBasic() BasicFeeAllowance {
	if m != nil {
		return m.Basic
	}
	return BasicFeeAllowance{}
}

func (m *PeriodicFeeAllowance) GetPeriod() Duration {
	if m != nil {
		return m.Period
	}
	return Duration{}
}

func (m *PeriodicFeeAllowance) GetPeriodSpendLimit() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.PeriodSpendLimit
	}
	return nil
}

func (m *PeriodicFeeAllowance) GetPeriodCanSpend() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.PeriodCanSpend
	}
	return nil
}

func (m *PeriodicFeeAllowance) GetPeriodReset() ExpiresAt {
	if m != nil {
		return m.PeriodReset
	}
	return ExpiresAt{}
}

// Duration is a repeating unit of either clock time or number of blocks.
// This is designed to be added to an ExpiresAt struct.
type Duration struct {
	Clock time.Duration `protobuf:"bytes,1,opt,name=clock,proto3,stdduration" json:"clock"`
	Block int64         `protobuf:"varint,2,opt,name=block,proto3" json:"block,omitempty"`
}

func (m *Duration) Reset()         { *m = Duration{} }
func (m *Duration) String() string { return proto.CompactTextString(m) }
func (*Duration) ProtoMessage()    {}
func (*Duration) Descriptor() ([]byte, []int) {
	return fileDescriptor_7279582900c30aea, []int{2}
}
func (m *Duration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Duration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Duration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Duration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Duration.Merge(m, src)
}
func (m *Duration) XXX_Size() int {
	return m.Size()
}
func (m *Duration) XXX_DiscardUnknown() {
	xxx_messageInfo_Duration.DiscardUnknown(m)
}

var xxx_messageInfo_Duration proto.InternalMessageInfo

func (m *Duration) GetClock() time.Duration {
	if m != nil {
		return m.Clock
	}
	return 0
}

func (m *Duration) GetBlock() int64 {
	if m != nil {
		return m.Block
	}
	return 0
}

// ExpiresAt is a point in time where something expires.
// It may be *either* block time or block height
type ExpiresAt struct {
	Time   time.Time `protobuf:"bytes,1,opt,name=time,proto3,stdtime" json:"time"`
	Height int64     `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
}

func (m *ExpiresAt) Reset()         { *m = ExpiresAt{} }
func (m *ExpiresAt) String() string { return proto.CompactTextString(m) }
func (*ExpiresAt) ProtoMessage()    {}
func (*ExpiresAt) Descriptor() ([]byte, []int) {
	return fileDescriptor_7279582900c30aea, []int{3}
}
func (m *ExpiresAt) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExpiresAt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExpiresAt.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExpiresAt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExpiresAt.Merge(m, src)
}
func (m *ExpiresAt) XXX_Size() int {
	return m.Size()
}
func (m *ExpiresAt) XXX_DiscardUnknown() {
	xxx_messageInfo_ExpiresAt.DiscardUnknown(m)
}

var xxx_messageInfo_ExpiresAt proto.InternalMessageInfo

func (m *ExpiresAt) GetTime() time.Time {
	if m != nil {
		return m.Time
	}
	return time.Time{}
}

func (m *ExpiresAt) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

// FeeAllowanceGrant is stored in the KVStore to record a grant with full context
type FeeAllowanceGrant struct {
	Granter   string      `protobuf:"bytes,1,opt,name=granter,proto3" json:"granter,omitempty"`
	Grantee   string      `protobuf:"bytes,2,opt,name=grantee,proto3" json:"grantee,omitempty"`
	Allowance *types1.Any `protobuf:"bytes,3,opt,name=allowance,proto3" json:"allowance,omitempty"`
}

func (m *FeeAllowanceGrant) Reset()         { *m = FeeAllowanceGrant{} }
func (m *FeeAllowanceGrant) String() string { return proto.CompactTextString(m) }
func (*FeeAllowanceGrant) ProtoMessage()    {}
func (*FeeAllowanceGrant) Descriptor() ([]byte, []int) {
	return fileDescriptor_7279582900c30aea, []int{4}
}
func (m *FeeAllowanceGrant) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FeeAllowanceGrant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FeeAllowanceGrant.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FeeAllowanceGrant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeeAllowanceGrant.Merge(m, src)
}
func (m *FeeAllowanceGrant) XXX_Size() int {
	return m.Size()
}
func (m *FeeAllowanceGrant) XXX_DiscardUnknown() {
	xxx_messageInfo_FeeAllowanceGrant.DiscardUnknown(m)
}

var xxx_messageInfo_FeeAllowanceGrant proto.InternalMessageInfo

func (m *FeeAllowanceGrant) GetGranter() string {
	if m != nil {
		return m.Granter
	}
	return ""
}

func (m *FeeAllowanceGrant) GetGrantee() string {
	if m != nil {
		return m.Grantee
	}
	return ""
}

func (m *FeeAllowanceGrant) GetAllowance() *types1.Any {
	if m != nil {
		return m.Allowance
	}
	return nil
}

func init() {
	proto.RegisterType((*BasicFeeAllowance)(nil), "cosmos.feegrant.v1beta1.BasicFeeAllowance")
	proto.RegisterType((*PeriodicFeeAllowance)(nil), "cosmos.feegrant.v1beta1.PeriodicFeeAllowance")
	proto.RegisterType((*Duration)(nil), "cosmos.feegrant.v1beta1.Duration")
	proto.RegisterType((*ExpiresAt)(nil), "cosmos.feegrant.v1beta1.ExpiresAt")
	proto.RegisterType((*FeeAllowanceGrant)(nil), "cosmos.feegrant.v1beta1.FeeAllowanceGrant")
}

func init() {
	proto.RegisterFile("cosmos/feegrant/v1beta1/feegrant.proto", fileDescriptor_7279582900c30aea)
}

var fileDescriptor_7279582900c30aea = []byte{
	// 577 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x6e, 0xd3, 0x40,
	0x14, 0x8e, 0x49, 0x13, 0x9a, 0x17, 0x40, 0x64, 0x14, 0x81, 0x9b, 0x85, 0x53, 0xb2, 0x40, 0x11,
	0x52, 0x6c, 0x5a, 0x36, 0xa8, 0x1b, 0x14, 0x97, 0xb6, 0x20, 0x58, 0x20, 0xc3, 0x0a, 0x81, 0x22,
	0xdb, 0x99, 0x3a, 0xa3, 0x3a, 0x1e, 0xcb, 0x33, 0x81, 0xe6, 0x12, 0xa8, 0x4b, 0xce, 0xc0, 0x9a,
	0x43, 0x54, 0xac, 0x2a, 0x56, 0xac, 0x5a, 0x94, 0x9c, 0x80, 0x1b, 0xa0, 0xf9, 0x71, 0x12, 0xa5,
	0x04, 0x81, 0xd4, 0x55, 0x3c, 0xf3, 0xde, 0xf7, 0xf3, 0xbe, 0x99, 0x09, 0xdc, 0x0f, 0x29, 0x1b,
	0x52, 0xe6, 0x1c, 0x62, 0x1c, 0x65, 0x7e, 0xc2, 0x9d, 0x0f, 0x5b, 0x01, 0xe6, 0xfe, 0xd6, 0x6c,
	0xc3, 0x4e, 0x33, 0xca, 0x29, 0xba, 0xab, 0xfa, 0xec, 0xd9, 0xb6, 0xee, 0x6b, 0xd4, 0x23, 0x1a,
	0x51, 0xd9, 0xe3, 0x88, 0x2f, 0xd5, 0xde, 0xd8, 0x88, 0x28, 0x8d, 0x62, 0xec, 0xc8, 0x55, 0x30,
	0x3a, 0x74, 0xfc, 0x64, 0x9c, 0x97, 0x14, 0x53, 0x4f, 0x61, 0x34, 0xad, 0x2a, 0x59, 0xda, 0x4c,
	0xe0, 0x33, 0x3c, 0x33, 0x12, 0x52, 0x92, 0xe8, 0x7a, 0x73, 0x99, 0x95, 0x93, 0x21, 0x66, 0xdc,
	0x1f, 0xa6, 0xaa, 0xa1, 0x75, 0x6e, 0x40, 0xcd, 0xf5, 0x19, 0x09, 0xf7, 0x31, 0xee, 0xc6, 0x31,
	0xfd, 0xe8, 0x27, 0x21, 0x46, 0x31, 0x54, 0x59, 0x8a, 0x93, 0x7e, 0x2f, 0x26, 0x43, 0xc2, 0x4d,
	0x63, 0xb3, 0xd8, 0xae, 0x6e, 0x6f, 0xd8, 0x5a, 0x5a, 0x88, 0xe5, 0xd3, 0xd8, 0xbb, 0x94, 0x24,
	0xee, 0xc3, 0xd3, 0xf3, 0x66, 0xe1, 0xcb, 0x45, 0xb3, 0x1d, 0x11, 0x3e, 0x18, 0x05, 0x76, 0x48,
	0x87, 0xda, 0xa7, 0xfe, 0xe9, 0xb0, 0xfe, 0x91, 0xc3, 0xc7, 0x29, 0x66, 0x12, 0xc0, 0x3c, 0x90,
	0xfc, 0x2f, 0x05, 0x3d, 0x7a, 0x06, 0x80, 0x8f, 0x53, 0x92, 0xf9, 0x9c, 0xd0, 0xc4, 0xbc, 0xb6,
	0x69, 0xb4, 0xab, 0xdb, 0x2d, 0x7b, 0x45, 0x7c, 0xf6, 0x9e, 0x68, 0xc5, 0xac, 0xcb, 0xdd, 0x35,
	0xa1, 0xea, 0x2d, 0x60, 0x77, 0x6a, 0xdf, 0xbf, 0x76, 0x6e, 0x2e, 0x4e, 0xf2, 0xbc, 0xf5, 0xab,
	0x08, 0xf5, 0x57, 0x38, 0x23, 0xb4, 0xbf, 0x34, 0xe3, 0x3e, 0x94, 0x02, 0x31, 0xb8, 0x69, 0x48,
	0xc1, 0x07, 0x2b, 0x05, 0x2f, 0xc5, 0xa3, 0x85, 0x15, 0x1c, 0x3d, 0x81, 0x72, 0x2a, 0xf9, 0xb5,
	0xf3, 0x7b, 0x2b, 0x89, 0x9e, 0x8e, 0x94, 0x4d, 0x8d, 0xd7, 0x30, 0x34, 0x06, 0xa4, 0xbe, 0x7a,
	0x8b, 0x99, 0x17, 0xaf, 0x3e, 0xf3, 0xdb, 0x4a, 0xe6, 0xf5, 0x3c, 0xf9, 0x11, 0xe8, 0xbd, 0x5e,
	0xe8, 0x27, 0x4a, 0xde, 0x5c, 0xbb, 0x7a, 0xe1, 0x5b, 0x4a, 0x64, 0xd7, 0x4f, 0xa4, 0x36, 0x7a,
	0x01, 0x37, 0xb4, 0x6c, 0x86, 0x19, 0xe6, 0x66, 0xe9, 0x3f, 0x8f, 0xbc, 0xaa, 0xd0, 0x9e, 0x00,
	0xff, 0xe9, 0xcc, 0xdf, 0xc1, 0x7a, 0x9e, 0x35, 0xda, 0x81, 0x52, 0x18, 0xd3, 0xf0, 0x48, 0x1f,
	0x73, 0xc3, 0x56, 0x2f, 0xc2, 0xce, 0x5f, 0x84, 0xfd, 0x26, 0x7f, 0x11, 0xee, 0xba, 0x20, 0xff,
	0x7c, 0xd1, 0x34, 0x3c, 0x05, 0x41, 0x75, 0x28, 0x05, 0x12, 0x2b, 0x4e, 0xb6, 0xe8, 0xa9, 0x45,
	0xeb, 0x3d, 0x54, 0x66, 0x86, 0xd0, 0x63, 0x58, 0x13, 0x4f, 0xea, 0x5f, 0xd9, 0x4f, 0x04, 0xbb,
	0x44, 0xa0, 0x3b, 0x50, 0x1e, 0x60, 0x12, 0x0d, 0xb8, 0x66, 0xd7, 0xab, 0xd6, 0x27, 0x03, 0x6a,
	0x8b, 0xe3, 0x1c, 0x88, 0x28, 0x90, 0x09, 0xd7, 0x65, 0x26, 0x38, 0x93, 0x52, 0x15, 0x2f, 0x5f,
	0xce, 0x2b, 0x58, 0x12, 0xcd, 0x2a, 0x18, 0xed, 0x41, 0xc5, 0xcf, 0x59, 0xcc, 0xa2, 0x34, 0x58,
	0xbf, 0x64, 0xb0, 0x9b, 0x8c, 0xdd, 0xda, 0xb7, 0xe5, 0x08, 0xbd, 0x39, 0xd2, 0x3d, 0x38, 0x9d,
	0x58, 0xc6, 0xd9, 0xc4, 0x32, 0x7e, 0x4e, 0x2c, 0xe3, 0x64, 0x6a, 0x15, 0xce, 0xa6, 0x56, 0xe1,
	0xc7, 0xd4, 0x2a, 0xbc, 0xed, 0xfc, 0xf5, 0x06, 0x1c, 0xcf, 0xff, 0x22, 0xe5, 0x65, 0x08, 0xca,
	0x52, 0xf4, 0xd1, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x07, 0x5c, 0x4a, 0xc1, 0x42, 0x05, 0x00,
	0x00,
}

func (m *BasicFeeAllowance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BasicFeeAllowance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BasicFeeAllowance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Expiration.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintFeegrant(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.SpendLimit) > 0 {
		for iNdEx := len(m.SpendLimit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SpendLimit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFeegrant(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *PeriodicFeeAllowance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PeriodicFeeAllowance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PeriodicFeeAllowance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.PeriodReset.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintFeegrant(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.PeriodCanSpend) > 0 {
		for iNdEx := len(m.PeriodCanSpend) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PeriodCanSpend[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFeegrant(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.PeriodSpendLimit) > 0 {
		for iNdEx := len(m.PeriodSpendLimit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PeriodSpendLimit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFeegrant(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size, err := m.Period.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintFeegrant(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Basic.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintFeegrant(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Duration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Duration) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Duration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Block != 0 {
		i = encodeVarintFeegrant(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x10
	}
	n5, err5 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Clock, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Clock):])
	if err5 != nil {
		return 0, err5
	}
	i -= n5
	i = encodeVarintFeegrant(dAtA, i, uint64(n5))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ExpiresAt) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExpiresAt) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExpiresAt) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		i = encodeVarintFeegrant(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x10
	}
	n6, err6 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Time):])
	if err6 != nil {
		return 0, err6
	}
	i -= n6
	i = encodeVarintFeegrant(dAtA, i, uint64(n6))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *FeeAllowanceGrant) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeeAllowanceGrant) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeeAllowanceGrant) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Allowance != nil {
		{
			size, err := m.Allowance.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFeegrant(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Grantee) > 0 {
		i -= len(m.Grantee)
		copy(dAtA[i:], m.Grantee)
		i = encodeVarintFeegrant(dAtA, i, uint64(len(m.Grantee)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Granter) > 0 {
		i -= len(m.Granter)
		copy(dAtA[i:], m.Granter)
		i = encodeVarintFeegrant(dAtA, i, uint64(len(m.Granter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFeegrant(dAtA []byte, offset int, v uint64) int {
	offset -= sovFeegrant(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BasicFeeAllowance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.SpendLimit) > 0 {
		for _, e := range m.SpendLimit {
			l = e.Size()
			n += 1 + l + sovFeegrant(uint64(l))
		}
	}
	l = m.Expiration.Size()
	n += 1 + l + sovFeegrant(uint64(l))
	return n
}

func (m *PeriodicFeeAllowance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Basic.Size()
	n += 1 + l + sovFeegrant(uint64(l))
	l = m.Period.Size()
	n += 1 + l + sovFeegrant(uint64(l))
	if len(m.PeriodSpendLimit) > 0 {
		for _, e := range m.PeriodSpendLimit {
			l = e.Size()
			n += 1 + l + sovFeegrant(uint64(l))
		}
	}
	if len(m.PeriodCanSpend) > 0 {
		for _, e := range m.PeriodCanSpend {
			l = e.Size()
			n += 1 + l + sovFeegrant(uint64(l))
		}
	}
	l = m.PeriodReset.Size()
	n += 1 + l + sovFeegrant(uint64(l))
	return n
}

func (m *Duration) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Clock)
	n += 1 + l + sovFeegrant(uint64(l))
	if m.Block != 0 {
		n += 1 + sovFeegrant(uint64(m.Block))
	}
	return n
}

func (m *ExpiresAt) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovFeegrant(uint64(l))
	if m.Height != 0 {
		n += 1 + sovFeegrant(uint64(m.Height))
	}
	return n
}

func (m *FeeAllowanceGrant) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Granter)
	if l > 0 {
		n += 1 + l + sovFeegrant(uint64(l))
	}
	l = len(m.Grantee)
	if l > 0 {
		n += 1 + l + sovFeegrant(uint64(l))
	}
	if m.Allowance != nil {
		l = m.Allowance.Size()
		n += 1 + l + sovFeegrant(uint64(l))
	}
	return n
}

func sovFeegrant(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFeegrant(x uint64) (n int) {
	return sovFeegrant(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BasicFeeAllowance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFeegrant
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BasicFeeAllowance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BasicFeeAllowance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpendLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpendLimit = append(m.SpendLimit, types.Coin{})
			if err := m.SpendLimit[len(m.SpendLimit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Expiration.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFeegrant(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFeegrant
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFeegrant
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PeriodicFeeAllowance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFeegrant
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PeriodicFeeAllowance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PeriodicFeeAllowance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Basic", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Basic.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Period", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Period.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeriodSpendLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeriodSpendLimit = append(m.PeriodSpendLimit, types.Coin{})
			if err := m.PeriodSpendLimit[len(m.PeriodSpendLimit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeriodCanSpend", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeriodCanSpend = append(m.PeriodCanSpend, types.Coin{})
			if err := m.PeriodCanSpend[len(m.PeriodCanSpend)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeriodReset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PeriodReset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFeegrant(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFeegrant
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFeegrant
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Duration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFeegrant
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Duration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Duration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Clock", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Clock, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Block |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFeegrant(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFeegrant
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFeegrant
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ExpiresAt) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFeegrant
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExpiresAt: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExpiresAt: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Time, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFeegrant(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFeegrant
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFeegrant
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FeeAllowanceGrant) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFeegrant
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FeeAllowanceGrant: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeeAllowanceGrant: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Granter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Grantee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allowance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Allowance == nil {
				m.Allowance = &types1.Any{}
			}
			if err := m.Allowance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFeegrant(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFeegrant
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFeegrant
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipFeegrant(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFeegrant
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFeegrant
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFeegrant
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthFeegrant
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFeegrant
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFeegrant
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFeegrant        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFeegrant          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFeegrant = fmt.Errorf("proto: unexpected end of group")
)