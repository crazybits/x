// Code generated by protoc-gen-go.
// source: crazybits/x/crypto/crypto.proto
// DO NOT EDIT!

/*
Package crypto is a generated protocol buffer package.

It is generated from these files:
	crazybits/x/crypto/crypto.proto

It has these top-level messages:
	PublicKey
	PrivateKey
	Signature
*/
package crypto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CryptoType int32

const (
	CryptoType_ECDSA CryptoType = 0
	CryptoType_RSA   CryptoType = 1
	CryptoType_DSA   CryptoType = 2
)

var CryptoType_name = map[int32]string{
	0: "ECDSA",
	1: "RSA",
	2: "DSA",
}
var CryptoType_value = map[string]int32{
	"ECDSA": 0,
	"RSA":   1,
	"DSA":   2,
}

func (x CryptoType) String() string {
	return proto.EnumName(CryptoType_name, int32(x))
}
func (CryptoType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type PublicKey struct {
	Type CryptoType `protobuf:"varint,1,opt,name=type,enum=crypto.CryptoType" json:"type,omitempty"`
	Key  []byte     `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *PublicKey) Reset()                    { *m = PublicKey{} }
func (m *PublicKey) String() string            { return proto.CompactTextString(m) }
func (*PublicKey) ProtoMessage()               {}
func (*PublicKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type PrivateKey struct {
	Type CryptoType `protobuf:"varint,1,opt,name=type,enum=crypto.CryptoType" json:"type,omitempty"`
	Key  []byte     `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *PrivateKey) Reset()                    { *m = PrivateKey{} }
func (m *PrivateKey) String() string            { return proto.CompactTextString(m) }
func (*PrivateKey) ProtoMessage()               {}
func (*PrivateKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Signature struct {
	Type CryptoType `protobuf:"varint,1,opt,name=type,enum=crypto.CryptoType" json:"type,omitempty"`
	R    []byte     `protobuf:"bytes,2,opt,name=r,proto3" json:"r,omitempty"`
	S    []byte     `protobuf:"bytes,3,opt,name=s,proto3" json:"s,omitempty"`
}

func (m *Signature) Reset()                    { *m = Signature{} }
func (m *Signature) String() string            { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()               {}
func (*Signature) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*PublicKey)(nil), "crypto.PublicKey")
	proto.RegisterType((*PrivateKey)(nil), "crypto.PrivateKey")
	proto.RegisterType((*Signature)(nil), "crypto.Signature")
	proto.RegisterEnum("crypto.CryptoType", CryptoType_name, CryptoType_value)
}

func init() { proto.RegisterFile("crazybits/x/crypto/crypto.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0x2e, 0x4a, 0xac,
	0xaa, 0x4c, 0xca, 0x2c, 0x29, 0xd6, 0xaf, 0xd0, 0x4f, 0x2e, 0xaa, 0x2c, 0x28, 0xc9, 0x87, 0x52,
	0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6c, 0x10, 0x9e, 0x92, 0x2b, 0x17, 0x67, 0x40, 0x69,
	0x52, 0x4e, 0x66, 0xb2, 0x77, 0x6a, 0xa5, 0x90, 0x1a, 0x17, 0x4b, 0x49, 0x65, 0x41, 0xaa, 0x04,
	0xa3, 0x02, 0xa3, 0x06, 0x9f, 0x91, 0x90, 0x1e, 0x54, 0x87, 0x33, 0x98, 0x0a, 0xa9, 0x2c, 0x48,
	0x0d, 0x02, 0xcb, 0x0b, 0x09, 0x70, 0x31, 0x67, 0xa7, 0x56, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0xf0,
	0x04, 0x81, 0x98, 0x4a, 0x6e, 0x5c, 0x5c, 0x01, 0x45, 0x99, 0x65, 0x89, 0x25, 0xa9, 0x94, 0x99,
	0xe3, 0xcf, 0xc5, 0x19, 0x9c, 0x99, 0x9e, 0x97, 0x58, 0x52, 0x5a, 0x94, 0x4a, 0xb4, 0x31, 0x3c,
	0x5c, 0x8c, 0x45, 0x50, 0x43, 0x18, 0x8b, 0x40, 0xbc, 0x62, 0x09, 0x66, 0x08, 0xaf, 0x58, 0x4b,
	0x93, 0x8b, 0x0b, 0xa1, 0x5e, 0x88, 0x93, 0x8b, 0xd5, 0xd5, 0xd9, 0x25, 0xd8, 0x51, 0x80, 0x41,
	0x88, 0x9d, 0x8b, 0x39, 0x28, 0xd8, 0x51, 0x80, 0x11, 0xc4, 0x00, 0x89, 0x30, 0x25, 0xb1, 0x81,
	0x43, 0xc6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xe7, 0xec, 0xa4, 0x0d, 0x3c, 0x01, 0x00, 0x00,
}
