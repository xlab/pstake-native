// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lscosmos/v1beta1/governance_proposal.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// RegisterCosmosChainProposal defines the details needed to register cosmos chain for
// liquid staking transactions through lscosmos module
type RegisterCosmosChainProposal struct {
	Title                string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description          string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	IBCConnection        string `protobuf:"bytes,3,opt,name=i_b_c_connection,json=iBCConnection,proto3" json:"i_b_c_connection,omitempty"`
	TokenTransferChannel string `protobuf:"bytes,4,opt,name=token_transfer_channel,json=tokenTransferChannel,proto3" json:"token_transfer_channel,omitempty"`
	TokenTransferPort    string `protobuf:"bytes,5,opt,name=token_transfer_port,json=tokenTransferPort,proto3" json:"token_transfer_port,omitempty"`
	BaseDenom            string `protobuf:"bytes,6,opt,name=base_denom,json=baseDenom,proto3" json:"base_denom,omitempty"`
	MintDenom            string `protobuf:"bytes,7,opt,name=mint_denom,json=mintDenom,proto3" json:"mint_denom,omitempty"`
	MinDeposit           string `protobuf:"bytes,8,opt,name=min_deposit,json=minDeposit,proto3" json:"min_deposit,omitempty"`
	PStakeDepositFee     string `protobuf:"bytes,9,opt,name=p_stake_deposit_fee,json=pStakeDepositFee,proto3" json:"p_stake_deposit_fee,omitempty"`
}

func (m *RegisterCosmosChainProposal) Reset()      { *m = RegisterCosmosChainProposal{} }
func (*RegisterCosmosChainProposal) ProtoMessage() {}
func (*RegisterCosmosChainProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_abbb79eadcf33bd7, []int{0}
}
func (m *RegisterCosmosChainProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterCosmosChainProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterCosmosChainProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterCosmosChainProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterCosmosChainProposal.Merge(m, src)
}
func (m *RegisterCosmosChainProposal) XXX_Size() int {
	return m.Size()
}
func (m *RegisterCosmosChainProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterCosmosChainProposal.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterCosmosChainProposal proto.InternalMessageInfo

type CosmosIBCParams struct {
	IBCConnection        string                                 `protobuf:"bytes,1,opt,name=i_b_c_connection,json=iBCConnection,proto3" json:"i_b_c_connection,omitempty"`
	TokenTransferChannel string                                 `protobuf:"bytes,2,opt,name=token_transfer_channel,json=tokenTransferChannel,proto3" json:"token_transfer_channel,omitempty"`
	TokenTransferPort    string                                 `protobuf:"bytes,3,opt,name=token_transfer_port,json=tokenTransferPort,proto3" json:"token_transfer_port,omitempty"`
	BaseDenom            string                                 `protobuf:"bytes,4,opt,name=base_denom,json=baseDenom,proto3" json:"base_denom,omitempty"`
	MintDenom            string                                 `protobuf:"bytes,5,opt,name=mint_denom,json=mintDenom,proto3" json:"mint_denom,omitempty"`
	MinDeposit           github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=min_deposit,json=minDeposit,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_deposit"`
	PStakeDepositFee     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,7,opt,name=p_stake_deposit_fee,json=pStakeDepositFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"p_stake_deposit_fee"`
}

func (m *CosmosIBCParams) Reset()         { *m = CosmosIBCParams{} }
func (m *CosmosIBCParams) String() string { return proto.CompactTextString(m) }
func (*CosmosIBCParams) ProtoMessage()    {}
func (*CosmosIBCParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_abbb79eadcf33bd7, []int{1}
}
func (m *CosmosIBCParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CosmosIBCParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CosmosIBCParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CosmosIBCParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CosmosIBCParams.Merge(m, src)
}
func (m *CosmosIBCParams) XXX_Size() int {
	return m.Size()
}
func (m *CosmosIBCParams) XXX_DiscardUnknown() {
	xxx_messageInfo_CosmosIBCParams.DiscardUnknown(m)
}

var xxx_messageInfo_CosmosIBCParams proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RegisterCosmosChainProposal)(nil), "lscosmos.v1beta1.RegisterCosmosChainProposal")
	proto.RegisterType((*CosmosIBCParams)(nil), "lscosmos.v1beta1.CosmosIBCParams")
}

func init() {
	proto.RegisterFile("lscosmos/v1beta1/governance_proposal.proto", fileDescriptor_abbb79eadcf33bd7)
}

var fileDescriptor_abbb79eadcf33bd7 = []byte{
	// 508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x31, 0x6f, 0xd3, 0x40,
	0x18, 0x86, 0xed, 0x26, 0x69, 0x9b, 0xab, 0x10, 0xc1, 0x8d, 0x90, 0x55, 0xc4, 0xb9, 0xea, 0x00,
	0x08, 0x29, 0xb6, 0x2a, 0x18, 0x10, 0x63, 0x1c, 0x21, 0x75, 0x6a, 0x54, 0x98, 0x90, 0xe0, 0x74,
	0xb9, 0x7c, 0x4d, 0x4e, 0x89, 0xef, 0x4e, 0x77, 0x47, 0x04, 0x1b, 0x3f, 0x81, 0x81, 0xa1, 0x63,
	0x46, 0x7e, 0x4a, 0xc6, 0x8e, 0x88, 0xa1, 0x82, 0x64, 0xe1, 0x47, 0x30, 0x20, 0x9f, 0xdd, 0xa8,
	0xa9, 0x22, 0x44, 0x26, 0xdb, 0xef, 0xf3, 0x7e, 0x77, 0xd6, 0x63, 0x1f, 0x7a, 0x3a, 0x36, 0x4c,
	0x9a, 0x4c, 0x9a, 0x64, 0x72, 0xdc, 0x03, 0x4b, 0x8f, 0x93, 0x81, 0x9c, 0x80, 0x16, 0x54, 0x30,
	0x20, 0x4a, 0x4b, 0x25, 0x0d, 0x1d, 0xc7, 0x4a, 0x4b, 0x2b, 0x83, 0xc6, 0x75, 0x37, 0x2e, 0xbb,
	0x07, 0xcd, 0x81, 0x1c, 0x48, 0x07, 0x93, 0xfc, 0xae, 0xe8, 0x1d, 0xe0, 0x72, 0xc5, 0x1e, 0x35,
	0xb0, 0x5c, 0x96, 0x49, 0x2e, 0x0a, 0x7e, 0xf4, 0x67, 0x0b, 0x3d, 0x38, 0x83, 0x01, 0x37, 0x16,
	0x74, 0xea, 0xaa, 0xe9, 0x90, 0x72, 0xd1, 0x2d, 0x77, 0x0b, 0x9a, 0xa8, 0x66, 0xb9, 0x1d, 0x43,
	0xe8, 0x1f, 0xfa, 0x4f, 0xea, 0x67, 0xc5, 0x43, 0x70, 0x88, 0xf6, 0xfa, 0x60, 0x98, 0xe6, 0xca,
	0x72, 0x29, 0xc2, 0x2d, 0xc7, 0x6e, 0x46, 0xc1, 0x63, 0xd4, 0xe0, 0xa4, 0x47, 0x18, 0x61, 0x52,
	0x08, 0x60, 0xae, 0x56, 0x71, 0xb5, 0x3b, 0xbc, 0x9d, 0xa6, 0xcb, 0x30, 0x78, 0x8e, 0xee, 0x5b,
	0x39, 0x02, 0x41, 0xac, 0xa6, 0xc2, 0x9c, 0x83, 0x26, 0x6c, 0x48, 0x85, 0x80, 0x71, 0x58, 0x75,
	0xf5, 0xa6, 0xa3, 0x6f, 0x4a, 0x98, 0x16, 0x2c, 0x88, 0xd1, 0xfe, 0xad, 0x29, 0x25, 0xb5, 0x0d,
	0x6b, 0x6e, 0xe4, 0xde, 0xca, 0x48, 0x57, 0x6a, 0x1b, 0x3c, 0x44, 0x28, 0x37, 0x40, 0xfa, 0x20,
	0x64, 0x16, 0x6e, 0xbb, 0x5a, 0x3d, 0x4f, 0x3a, 0x79, 0x90, 0xe3, 0x8c, 0x0b, 0x5b, 0xe2, 0x9d,
	0x02, 0xe7, 0x49, 0x81, 0x23, 0xb4, 0x97, 0x71, 0x41, 0xfa, 0xa0, 0xa4, 0xe1, 0x36, 0xdc, 0x75,
	0x3c, 0x9f, 0xe8, 0x14, 0x49, 0xd0, 0x42, 0xfb, 0x8a, 0x18, 0x4b, 0x47, 0x70, 0x5d, 0x22, 0xe7,
	0x00, 0x61, 0xdd, 0x15, 0x1b, 0xea, 0x75, 0x4e, 0xca, 0xee, 0x2b, 0x80, 0x97, 0xbb, 0x17, 0xd3,
	0xc8, 0xfb, 0x3d, 0x8d, 0xbc, 0xa3, 0xaf, 0x15, 0x74, 0xb7, 0xd0, 0x7e, 0xd2, 0x4e, 0xbb, 0x54,
	0xd3, 0xcc, 0xac, 0x55, 0xe7, 0x6f, 0xa6, 0x6e, 0x6b, 0x73, 0x75, 0x95, 0xff, 0x53, 0x57, 0xfd,
	0xb7, 0xba, 0xda, 0x6d, 0x75, 0xa7, 0xab, 0xea, 0x9c, 0xf9, 0x76, 0x3c, 0xbb, 0x8a, 0xbc, 0x1f,
	0x57, 0xd1, 0xa3, 0x01, 0xb7, 0xc3, 0x0f, 0xbd, 0x98, 0xc9, 0x2c, 0x29, 0xff, 0xd3, 0xe2, 0xd2,
	0x32, 0xfd, 0x51, 0x62, 0x3f, 0x29, 0x30, 0xf1, 0x89, 0xb0, 0x2b, 0xaa, 0xdf, 0xad, 0x57, 0xbd,
	0xb3, 0xf1, 0xc2, 0x1d, 0x60, 0x6b, 0x3e, 0x4d, 0xf5, 0x62, 0x1a, 0xf9, 0xed, 0xf7, 0xb3, 0x5f,
	0xd8, 0xfb, 0x3c, 0xc7, 0xde, 0xb7, 0x39, 0xf6, 0x67, 0x73, 0xec, 0x5f, 0xce, 0xb1, 0xff, 0x73,
	0x8e, 0xfd, 0x2f, 0x0b, 0xec, 0x5d, 0x2e, 0xb0, 0xf7, 0x7d, 0x81, 0xbd, 0xb7, 0x2f, 0x6e, 0xec,
	0xa2, 0x40, 0x9b, 0xfc, 0x10, 0x09, 0x06, 0xa7, 0x02, 0x12, 0xe5, 0xde, 0xb0, 0x25, 0xa8, 0xe5,
	0x13, 0x48, 0x3e, 0x26, 0xcb, 0x93, 0xed, 0xf6, 0xee, 0x6d, 0xbb, 0xc3, 0xf7, 0xec, 0x6f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x09, 0x90, 0xc0, 0x2b, 0xf2, 0x03, 0x00, 0x00,
}

func (this *CosmosIBCParams) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CosmosIBCParams)
	if !ok {
		that2, ok := that.(CosmosIBCParams)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.IBCConnection != that1.IBCConnection {
		return false
	}
	if this.TokenTransferChannel != that1.TokenTransferChannel {
		return false
	}
	if this.TokenTransferPort != that1.TokenTransferPort {
		return false
	}
	if this.BaseDenom != that1.BaseDenom {
		return false
	}
	if this.MintDenom != that1.MintDenom {
		return false
	}
	if !this.MinDeposit.Equal(that1.MinDeposit) {
		return false
	}
	if !this.PStakeDepositFee.Equal(that1.PStakeDepositFee) {
		return false
	}
	return true
}
func (m *RegisterCosmosChainProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterCosmosChainProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterCosmosChainProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PStakeDepositFee) > 0 {
		i -= len(m.PStakeDepositFee)
		copy(dAtA[i:], m.PStakeDepositFee)
		i = encodeVarintGovernanceProposal(dAtA, i, uint64(len(m.PStakeDepositFee)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.MinDeposit) > 0 {
		i -= len(m.MinDeposit)
		copy(dAtA[i:], m.MinDeposit)
		i = encodeVarintGovernanceProposal(dAtA, i, uint64(len(m.MinDeposit)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.MintDenom) > 0 {
		i -= len(m.MintDenom)
		copy(dAtA[i:], m.MintDenom)
		i = encodeVarintGovernanceProposal(dAtA, i, uint64(len(m.MintDenom)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.BaseDenom) > 0 {
		i -= len(m.BaseDenom)
		copy(dAtA[i:], m.BaseDenom)
		i = encodeVarintGovernanceProposal(dAtA, i, uint64(len(m.BaseDenom)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.TokenTransferPort) > 0 {
		i -= len(m.TokenTransferPort)
		copy(dAtA[i:], m.TokenTransferPort)
		i = encodeVarintGovernanceProposal(dAtA, i, uint64(len(m.TokenTransferPort)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.TokenTransferChannel) > 0 {
		i -= len(m.TokenTransferChannel)
		copy(dAtA[i:], m.TokenTransferChannel)
		i = encodeVarintGovernanceProposal(dAtA, i, uint64(len(m.TokenTransferChannel)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.IBCConnection) > 0 {
		i -= len(m.IBCConnection)
		copy(dAtA[i:], m.IBCConnection)
		i = encodeVarintGovernanceProposal(dAtA, i, uint64(len(m.IBCConnection)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGovernanceProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGovernanceProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CosmosIBCParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CosmosIBCParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CosmosIBCParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.PStakeDepositFee.Size()
		i -= size
		if _, err := m.PStakeDepositFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGovernanceProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.MinDeposit.Size()
		i -= size
		if _, err := m.MinDeposit.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGovernanceProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.MintDenom) > 0 {
		i -= len(m.MintDenom)
		copy(dAtA[i:], m.MintDenom)
		i = encodeVarintGovernanceProposal(dAtA, i, uint64(len(m.MintDenom)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.BaseDenom) > 0 {
		i -= len(m.BaseDenom)
		copy(dAtA[i:], m.BaseDenom)
		i = encodeVarintGovernanceProposal(dAtA, i, uint64(len(m.BaseDenom)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.TokenTransferPort) > 0 {
		i -= len(m.TokenTransferPort)
		copy(dAtA[i:], m.TokenTransferPort)
		i = encodeVarintGovernanceProposal(dAtA, i, uint64(len(m.TokenTransferPort)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.TokenTransferChannel) > 0 {
		i -= len(m.TokenTransferChannel)
		copy(dAtA[i:], m.TokenTransferChannel)
		i = encodeVarintGovernanceProposal(dAtA, i, uint64(len(m.TokenTransferChannel)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.IBCConnection) > 0 {
		i -= len(m.IBCConnection)
		copy(dAtA[i:], m.IBCConnection)
		i = encodeVarintGovernanceProposal(dAtA, i, uint64(len(m.IBCConnection)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGovernanceProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovGovernanceProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RegisterCosmosChainProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGovernanceProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGovernanceProposal(uint64(l))
	}
	l = len(m.IBCConnection)
	if l > 0 {
		n += 1 + l + sovGovernanceProposal(uint64(l))
	}
	l = len(m.TokenTransferChannel)
	if l > 0 {
		n += 1 + l + sovGovernanceProposal(uint64(l))
	}
	l = len(m.TokenTransferPort)
	if l > 0 {
		n += 1 + l + sovGovernanceProposal(uint64(l))
	}
	l = len(m.BaseDenom)
	if l > 0 {
		n += 1 + l + sovGovernanceProposal(uint64(l))
	}
	l = len(m.MintDenom)
	if l > 0 {
		n += 1 + l + sovGovernanceProposal(uint64(l))
	}
	l = len(m.MinDeposit)
	if l > 0 {
		n += 1 + l + sovGovernanceProposal(uint64(l))
	}
	l = len(m.PStakeDepositFee)
	if l > 0 {
		n += 1 + l + sovGovernanceProposal(uint64(l))
	}
	return n
}

func (m *CosmosIBCParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.IBCConnection)
	if l > 0 {
		n += 1 + l + sovGovernanceProposal(uint64(l))
	}
	l = len(m.TokenTransferChannel)
	if l > 0 {
		n += 1 + l + sovGovernanceProposal(uint64(l))
	}
	l = len(m.TokenTransferPort)
	if l > 0 {
		n += 1 + l + sovGovernanceProposal(uint64(l))
	}
	l = len(m.BaseDenom)
	if l > 0 {
		n += 1 + l + sovGovernanceProposal(uint64(l))
	}
	l = len(m.MintDenom)
	if l > 0 {
		n += 1 + l + sovGovernanceProposal(uint64(l))
	}
	l = m.MinDeposit.Size()
	n += 1 + l + sovGovernanceProposal(uint64(l))
	l = m.PStakeDepositFee.Size()
	n += 1 + l + sovGovernanceProposal(uint64(l))
	return n
}

func sovGovernanceProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGovernanceProposal(x uint64) (n int) {
	return sovGovernanceProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RegisterCosmosChainProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGovernanceProposal
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
			return fmt.Errorf("proto: RegisterCosmosChainProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterCosmosChainProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovernanceProposal
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
				return ErrInvalidLengthGovernanceProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGovernanceProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovernanceProposal
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
				return ErrInvalidLengthGovernanceProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGovernanceProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IBCConnection", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovernanceProposal
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
				return ErrInvalidLengthGovernanceProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGovernanceProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IBCConnection = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenTransferChannel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovernanceProposal
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
				return ErrInvalidLengthGovernanceProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGovernanceProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenTransferChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenTransferPort", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovernanceProposal
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
				return ErrInvalidLengthGovernanceProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGovernanceProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenTransferPort = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovernanceProposal
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
				return ErrInvalidLengthGovernanceProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGovernanceProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovernanceProposal
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
				return ErrInvalidLengthGovernanceProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGovernanceProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinDeposit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovernanceProposal
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
				return ErrInvalidLengthGovernanceProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGovernanceProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinDeposit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PStakeDepositFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovernanceProposal
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
				return ErrInvalidLengthGovernanceProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGovernanceProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PStakeDepositFee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGovernanceProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGovernanceProposal
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
func (m *CosmosIBCParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGovernanceProposal
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
			return fmt.Errorf("proto: CosmosIBCParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CosmosIBCParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IBCConnection", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovernanceProposal
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
				return ErrInvalidLengthGovernanceProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGovernanceProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IBCConnection = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenTransferChannel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovernanceProposal
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
				return ErrInvalidLengthGovernanceProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGovernanceProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenTransferChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenTransferPort", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovernanceProposal
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
				return ErrInvalidLengthGovernanceProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGovernanceProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenTransferPort = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovernanceProposal
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
				return ErrInvalidLengthGovernanceProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGovernanceProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovernanceProposal
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
				return ErrInvalidLengthGovernanceProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGovernanceProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinDeposit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovernanceProposal
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
				return ErrInvalidLengthGovernanceProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGovernanceProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinDeposit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PStakeDepositFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovernanceProposal
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
				return ErrInvalidLengthGovernanceProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGovernanceProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PStakeDepositFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGovernanceProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGovernanceProposal
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
func skipGovernanceProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGovernanceProposal
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
					return 0, ErrIntOverflowGovernanceProposal
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
					return 0, ErrIntOverflowGovernanceProposal
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
				return 0, ErrInvalidLengthGovernanceProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGovernanceProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGovernanceProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGovernanceProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGovernanceProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGovernanceProposal = fmt.Errorf("proto: unexpected end of group")
)
