// Code generated by protoc-gen-go. DO NOT EDIT.
// source: apns/apns.proto

package apns

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

type Header struct {
	// APNsサービスのアドレス
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// PEMエンコードされた秘密鍵
	KeyPEMBlock []byte `protobuf:"bytes,2,opt,name=keyPEMBlock,proto3" json:"keyPEMBlock,omitempty"`
	// PEMエンコードされた証明書
	CertPEMBlock []byte `protobuf:"bytes,3,opt,name=certPEMBlock,proto3" json:"certPEMBlock,omitempty"`
	// Do not set to true in production
	InsecureSkipVerify bool `protobuf:"varint,4,opt,name=insecureSkipVerify,proto3" json:"insecureSkipVerify,omitempty"`
	// JWT認証時のIssuer
	Issuer string `protobuf:"bytes,5,opt,name=issuer,proto3" json:"issuer,omitempty"`
	// JWT認証時のKey ID
	KeyID string `protobuf:"bytes,6,opt,name=keyID,proto3" json:"keyID,omitempty"`
	// JWT認証時のEC P-256秘密鍵
	PrivateKey string `protobuf:"bytes,7,opt,name=privateKey,proto3" json:"privateKey,omitempty"`
	// JWT認証時のトピック(e.g. Bundle ID)
	Topic string `protobuf:"bytes,8,opt,name=topic,proto3" json:"topic,omitempty"`
	// apns-push-type ('alert' / 'background'; default 'alert')
	PushType             string   `protobuf:"bytes,9,opt,name=pushType,proto3" json:"pushType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f97b0cb6762d65b, []int{0}
}

func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Header) GetKeyPEMBlock() []byte {
	if m != nil {
		return m.KeyPEMBlock
	}
	return nil
}

func (m *Header) GetCertPEMBlock() []byte {
	if m != nil {
		return m.CertPEMBlock
	}
	return nil
}

func (m *Header) GetInsecureSkipVerify() bool {
	if m != nil {
		return m.InsecureSkipVerify
	}
	return false
}

func (m *Header) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *Header) GetKeyID() string {
	if m != nil {
		return m.KeyID
	}
	return ""
}

func (m *Header) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

func (m *Header) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Header) GetPushType() string {
	if m != nil {
		return m.PushType
	}
	return ""
}

func init() {
	proto.RegisterType((*Header)(nil), "apns.Header")
}

func init() { proto.RegisterFile("apns/apns.proto", fileDescriptor_9f97b0cb6762d65b) }

var fileDescriptor_9f97b0cb6762d65b = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x6c, 0xd3, 0x74, 0x2c, 0x08, 0x4b, 0x91, 0xc5, 0x83, 0x84, 0xe2, 0x21, 0xa7,
	0xe4, 0xe0, 0xd9, 0x83, 0xc1, 0x82, 0x22, 0x82, 0x44, 0xf1, 0xe0, 0x2d, 0xdd, 0x8c, 0xed, 0x92,
	0x9a, 0x1d, 0x76, 0x37, 0xc2, 0xfa, 0x24, 0x3e, 0x83, 0x4f, 0x29, 0xd9, 0x6a, 0x89, 0xd0, 0xcb,
	0xb2, 0xdf, 0xff, 0xcd, 0xc0, 0xf0, 0xc3, 0x49, 0x45, 0xad, 0xc9, 0xfb, 0x27, 0x23, 0xad, 0xac,
	0x62, 0xa3, 0xfe, 0xbf, 0xf8, 0x0a, 0x21, 0xba, 0xc5, 0xaa, 0x46, 0xcd, 0x38, 0x4c, 0xaa, 0xba,
	0xd6, 0x68, 0x0c, 0x0f, 0x92, 0x20, 0x9d, 0x96, 0x7f, 0xc8, 0x12, 0x38, 0x6e, 0xd0, 0x3d, 0x2e,
	0x1f, 0x8a, 0xad, 0x12, 0x0d, 0x0f, 0x93, 0x20, 0x9d, 0x95, 0xc3, 0x88, 0x2d, 0x60, 0x26, 0x50,
	0xdb, 0xfd, 0xc8, 0x91, 0x1f, 0xf9, 0x97, 0xb1, 0x0c, 0x98, 0x6c, 0x0d, 0x8a, 0x4e, 0xe3, 0x53,
	0x23, 0xe9, 0x05, 0xb5, 0x7c, 0x73, 0x7c, 0x94, 0x04, 0x69, 0x5c, 0x1e, 0x30, 0xec, 0x14, 0x22,
	0x69, 0x4c, 0x87, 0x9a, 0x8f, 0xfd, 0x39, 0xbf, 0xc4, 0xe6, 0x30, 0x6e, 0xd0, 0xdd, 0xdd, 0xf0,
	0xc8, 0xc7, 0x3b, 0x60, 0xe7, 0x00, 0xa4, 0xe5, 0x47, 0x65, 0xf1, 0x1e, 0x1d, 0x9f, 0x78, 0x35,
	0x48, 0xfa, 0x2d, 0xab, 0x48, 0x0a, 0x1e, 0xef, 0xb6, 0x3c, 0xb0, 0x33, 0x88, 0xa9, 0x33, 0x9b,
	0x67, 0x47, 0xc8, 0xa7, 0x5e, 0xec, 0xb9, 0xb8, 0x7a, 0xbd, 0x58, 0x4b, 0xbb, 0xe9, 0x56, 0x99,
	0x50, 0xef, 0x79, 0xa1, 0xb6, 0xf6, 0x73, 0xd9, 0xae, 0x65, 0x8b, 0x79, 0x45, 0xd2, 0xe4, 0x9a,
	0x84, 0xaf, 0xf3, 0x3b, 0x9c, 0x0f, 0x5c, 0x56, 0x92, 0xc8, 0xae, 0xa9, 0x35, 0xab, 0xc8, 0xd7,
	0x7c, 0xf9, 0x13, 0x00, 0x00, 0xff, 0xff, 0xd3, 0xd8, 0x5a, 0x20, 0x79, 0x01, 0x00, 0x00,
}
