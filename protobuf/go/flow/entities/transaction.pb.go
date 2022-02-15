// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flow/entities/transaction.proto

package entities

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

type TransactionStatus int32

const (
	TransactionStatus_UNKNOWN   TransactionStatus = 0
	TransactionStatus_PENDING   TransactionStatus = 1
	TransactionStatus_FINALIZED TransactionStatus = 2
	TransactionStatus_EXECUTED  TransactionStatus = 3
	TransactionStatus_SEALED    TransactionStatus = 4
	TransactionStatus_EXPIRED   TransactionStatus = 5
)

var TransactionStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "PENDING",
	2: "FINALIZED",
	3: "EXECUTED",
	4: "SEALED",
	5: "EXPIRED",
}

var TransactionStatus_value = map[string]int32{
	"UNKNOWN":   0,
	"PENDING":   1,
	"FINALIZED": 2,
	"EXECUTED":  3,
	"SEALED":    4,
	"EXPIRED":   5,
}

func (x TransactionStatus) String() string {
	return proto.EnumName(TransactionStatus_name, int32(x))
}

func (TransactionStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_84361b1ad2e0936e, []int{0}
}

type Transaction struct {
	Script               []byte                   `protobuf:"bytes,1,opt,name=script,proto3" json:"script,omitempty"`
	Arguments            [][]byte                 `protobuf:"bytes,2,rep,name=arguments,proto3" json:"arguments,omitempty"`
	ReferenceBlockId     []byte                   `protobuf:"bytes,3,opt,name=reference_block_id,json=referenceBlockId,proto3" json:"reference_block_id,omitempty"`
	GasLimit             uint64                   `protobuf:"varint,4,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	ProposalKey          *Transaction_ProposalKey `protobuf:"bytes,5,opt,name=proposal_key,json=proposalKey,proto3" json:"proposal_key,omitempty"`
	Payer                []byte                   `protobuf:"bytes,6,opt,name=payer,proto3" json:"payer,omitempty"`
	Authorizers          [][]byte                 `protobuf:"bytes,7,rep,name=authorizers,proto3" json:"authorizers,omitempty"`
	PayloadSignatures    []*Transaction_Signature `protobuf:"bytes,8,rep,name=payload_signatures,json=payloadSignatures,proto3" json:"payload_signatures,omitempty"`
	EnvelopeSignatures   []*Transaction_Signature `protobuf:"bytes,9,rep,name=envelope_signatures,json=envelopeSignatures,proto3" json:"envelope_signatures,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_84361b1ad2e0936e, []int{0}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetScript() []byte {
	if m != nil {
		return m.Script
	}
	return nil
}

func (m *Transaction) GetArguments() [][]byte {
	if m != nil {
		return m.Arguments
	}
	return nil
}

func (m *Transaction) GetReferenceBlockId() []byte {
	if m != nil {
		return m.ReferenceBlockId
	}
	return nil
}

func (m *Transaction) GetGasLimit() uint64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *Transaction) GetProposalKey() *Transaction_ProposalKey {
	if m != nil {
		return m.ProposalKey
	}
	return nil
}

func (m *Transaction) GetPayer() []byte {
	if m != nil {
		return m.Payer
	}
	return nil
}

func (m *Transaction) GetAuthorizers() [][]byte {
	if m != nil {
		return m.Authorizers
	}
	return nil
}

func (m *Transaction) GetPayloadSignatures() []*Transaction_Signature {
	if m != nil {
		return m.PayloadSignatures
	}
	return nil
}

func (m *Transaction) GetEnvelopeSignatures() []*Transaction_Signature {
	if m != nil {
		return m.EnvelopeSignatures
	}
	return nil
}

type Transaction_ProposalKey struct {
	Address              []byte   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	KeyId                uint32   `protobuf:"varint,2,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	SequenceNumber       uint64   `protobuf:"varint,3,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction_ProposalKey) Reset()         { *m = Transaction_ProposalKey{} }
func (m *Transaction_ProposalKey) String() string { return proto.CompactTextString(m) }
func (*Transaction_ProposalKey) ProtoMessage()    {}
func (*Transaction_ProposalKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_84361b1ad2e0936e, []int{0, 0}
}

func (m *Transaction_ProposalKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction_ProposalKey.Unmarshal(m, b)
}
func (m *Transaction_ProposalKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction_ProposalKey.Marshal(b, m, deterministic)
}
func (m *Transaction_ProposalKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction_ProposalKey.Merge(m, src)
}
func (m *Transaction_ProposalKey) XXX_Size() int {
	return xxx_messageInfo_Transaction_ProposalKey.Size(m)
}
func (m *Transaction_ProposalKey) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction_ProposalKey.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction_ProposalKey proto.InternalMessageInfo

func (m *Transaction_ProposalKey) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Transaction_ProposalKey) GetKeyId() uint32 {
	if m != nil {
		return m.KeyId
	}
	return 0
}

func (m *Transaction_ProposalKey) GetSequenceNumber() uint64 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

type Transaction_Signature struct {
	Address              []byte   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	KeyId                uint32   `protobuf:"varint,2,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	Signature            []byte   `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction_Signature) Reset()         { *m = Transaction_Signature{} }
func (m *Transaction_Signature) String() string { return proto.CompactTextString(m) }
func (*Transaction_Signature) ProtoMessage()    {}
func (*Transaction_Signature) Descriptor() ([]byte, []int) {
	return fileDescriptor_84361b1ad2e0936e, []int{0, 1}
}

func (m *Transaction_Signature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction_Signature.Unmarshal(m, b)
}
func (m *Transaction_Signature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction_Signature.Marshal(b, m, deterministic)
}
func (m *Transaction_Signature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction_Signature.Merge(m, src)
}
func (m *Transaction_Signature) XXX_Size() int {
	return xxx_messageInfo_Transaction_Signature.Size(m)
}
func (m *Transaction_Signature) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction_Signature.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction_Signature proto.InternalMessageInfo

func (m *Transaction_Signature) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Transaction_Signature) GetKeyId() uint32 {
	if m != nil {
		return m.KeyId
	}
	return 0
}

func (m *Transaction_Signature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterEnum("flow.entities.TransactionStatus", TransactionStatus_name, TransactionStatus_value)
	proto.RegisterType((*Transaction)(nil), "flow.entities.Transaction")
	proto.RegisterType((*Transaction_ProposalKey)(nil), "flow.entities.Transaction.ProposalKey")
	proto.RegisterType((*Transaction_Signature)(nil), "flow.entities.Transaction.Signature")
}

func init() { proto.RegisterFile("flow/entities/transaction.proto", fileDescriptor_84361b1ad2e0936e) }

var fileDescriptor_84361b1ad2e0936e = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x5f, 0x8b, 0x9b, 0x4e,
	0x14, 0xfd, 0x99, 0xff, 0x5e, 0x93, 0x5f, 0xdd, 0xe9, 0x1f, 0x64, 0x1b, 0xa8, 0x94, 0xd2, 0x4a,
	0x29, 0x5a, 0xb6, 0x9f, 0x60, 0xb7, 0xb1, 0x45, 0x36, 0xd8, 0x60, 0x36, 0x74, 0xc9, 0x4b, 0x98,
	0xe8, 0x8d, 0x2b, 0x31, 0x8e, 0x9d, 0x19, 0x5b, 0xec, 0x63, 0x3f, 0x79, 0xd1, 0x8d, 0x49, 0xf6,
	0xa5, 0xb0, 0x2f, 0xc2, 0x39, 0xf7, 0xcc, 0x99, 0x7b, 0xe7, 0x78, 0xe1, 0xd5, 0x26, 0x65, 0xbf,
	0x1c, 0xcc, 0x64, 0x22, 0x13, 0x14, 0x8e, 0xe4, 0x34, 0x13, 0x34, 0x94, 0x09, 0xcb, 0xec, 0x9c,
	0x33, 0xc9, 0xc8, 0xa8, 0x12, 0xd8, 0x8d, 0xe0, 0xf5, 0x9f, 0x2e, 0x68, 0x37, 0x47, 0x11, 0x79,
	0x01, 0x3d, 0x11, 0xf2, 0x24, 0x97, 0x86, 0x62, 0x2a, 0xd6, 0x30, 0xd8, 0x23, 0x32, 0x06, 0x95,
	0xf2, 0xb8, 0xd8, 0x61, 0x26, 0x85, 0xd1, 0x32, 0xdb, 0xd6, 0x30, 0x38, 0x12, 0xe4, 0x03, 0x10,
	0x8e, 0x1b, 0xe4, 0x98, 0x85, 0xb8, 0x5a, 0xa7, 0x2c, 0xdc, 0xae, 0x92, 0xc8, 0x68, 0xd7, 0x0e,
	0xfa, 0xa1, 0x72, 0x55, 0x15, 0xbc, 0x88, 0xbc, 0x04, 0x35, 0xa6, 0x62, 0x95, 0x26, 0xbb, 0x44,
	0x1a, 0x1d, 0x53, 0xb1, 0x3a, 0xc1, 0x20, 0xa6, 0x62, 0x5a, 0x61, 0xe2, 0xc1, 0x30, 0xe7, 0x2c,
	0x67, 0x82, 0xa6, 0xab, 0x2d, 0x96, 0x46, 0xd7, 0x54, 0x2c, 0xed, 0xe2, 0xad, 0xfd, 0xa0, 0x6d,
	0xfb, 0xa4, 0x65, 0x7b, 0xb6, 0x97, 0x5f, 0x63, 0x19, 0x68, 0xf9, 0x11, 0x90, 0x67, 0xd0, 0xcd,
	0x69, 0x89, 0xdc, 0xe8, 0xd5, 0x8d, 0xdc, 0x03, 0x62, 0x82, 0x46, 0x0b, 0x79, 0xc7, 0x78, 0xf2,
	0x1b, 0xb9, 0x30, 0xfa, 0xf5, 0x2c, 0xa7, 0x14, 0x99, 0x03, 0xc9, 0x69, 0x99, 0x32, 0x1a, 0xad,
	0x44, 0x12, 0x67, 0x54, 0x16, 0x1c, 0x85, 0x31, 0x30, 0xdb, 0x96, 0x76, 0xf1, 0xe6, 0x1f, 0x8d,
	0xcc, 0x1b, 0x71, 0x70, 0xb6, 0x3f, 0x7f, 0x60, 0x04, 0x59, 0xc0, 0x53, 0xcc, 0x7e, 0x62, 0xca,
	0x72, 0x3c, 0x75, 0x55, 0x1f, 0xe1, 0x4a, 0x1a, 0x83, 0xa3, 0xed, 0x79, 0x0c, 0xda, 0xc9, 0xfc,
	0xc4, 0x80, 0x3e, 0x8d, 0x22, 0x8e, 0x42, 0xec, 0xf3, 0x6b, 0x20, 0x79, 0x0e, 0xbd, 0x2d, 0x96,
	0x55, 0x2c, 0x2d, 0x53, 0xb1, 0x46, 0x41, 0x77, 0x8b, 0xa5, 0x17, 0x91, 0x77, 0xf0, 0x44, 0xe0,
	0x8f, 0xa2, 0x0e, 0x2e, 0x2b, 0x76, 0x6b, 0xe4, 0x75, 0x6c, 0x9d, 0xe0, 0xff, 0x86, 0xf6, 0x6b,
	0xf6, 0x7c, 0x09, 0xea, 0xe1, 0xda, 0xc7, 0x5f, 0x33, 0x06, 0xf5, 0x30, 0xf4, 0xfe, 0xbf, 0x38,
	0x12, 0xef, 0x43, 0x38, 0x3b, 0x99, 0x78, 0x2e, 0xa9, 0x2c, 0x04, 0xd1, 0xa0, 0xbf, 0xf0, 0xaf,
	0xfd, 0x6f, 0xdf, 0x7d, 0xfd, 0xbf, 0x0a, 0xcc, 0x5c, 0x7f, 0xe2, 0xf9, 0x5f, 0x75, 0x85, 0x8c,
	0x40, 0xfd, 0xe2, 0xf9, 0x97, 0x53, 0x6f, 0xe9, 0x4e, 0xf4, 0x16, 0x19, 0xc2, 0xc0, 0xbd, 0x75,
	0x3f, 0x2f, 0x6e, 0xdc, 0x89, 0xde, 0x26, 0x00, 0xbd, 0xb9, 0x7b, 0x39, 0x75, 0x27, 0x7a, 0xa7,
	0x3a, 0xe5, 0xde, 0xce, 0xbc, 0xc0, 0x9d, 0xe8, 0xdd, 0xab, 0x19, 0x8c, 0x19, 0x8f, 0x6d, 0x96,
	0xd5, 0x4f, 0x5d, 0x2f, 0xc3, 0xba, 0xd8, 0x1c, 0xde, 0x7c, 0xf9, 0x31, 0x4e, 0xe4, 0x5d, 0xb1,
	0xb6, 0x43, 0xb6, 0x73, 0xee, 0x45, 0x4e, 0xfd, 0x69, 0x94, 0x4e, 0xcc, 0x9c, 0x07, 0xcb, 0xb5,
	0xee, 0xd5, 0xa5, 0x4f, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x7b, 0xf2, 0x21, 0x74, 0x03,
	0x00, 0x00,
}
