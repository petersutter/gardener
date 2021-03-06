// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/v3/token_bucket.proto

package envoy_type_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type TokenBucket struct {
	MaxTokens            uint32                `protobuf:"varint,1,opt,name=max_tokens,json=maxTokens,proto3" json:"max_tokens,omitempty"`
	TokensPerFill        *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=tokens_per_fill,json=tokensPerFill,proto3" json:"tokens_per_fill,omitempty"`
	FillInterval         *duration.Duration    `protobuf:"bytes,3,opt,name=fill_interval,json=fillInterval,proto3" json:"fill_interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TokenBucket) Reset()         { *m = TokenBucket{} }
func (m *TokenBucket) String() string { return proto.CompactTextString(m) }
func (*TokenBucket) ProtoMessage()    {}
func (*TokenBucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b88f84e3f7d4627, []int{0}
}

func (m *TokenBucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenBucket.Unmarshal(m, b)
}
func (m *TokenBucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenBucket.Marshal(b, m, deterministic)
}
func (m *TokenBucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenBucket.Merge(m, src)
}
func (m *TokenBucket) XXX_Size() int {
	return xxx_messageInfo_TokenBucket.Size(m)
}
func (m *TokenBucket) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenBucket.DiscardUnknown(m)
}

var xxx_messageInfo_TokenBucket proto.InternalMessageInfo

func (m *TokenBucket) GetMaxTokens() uint32 {
	if m != nil {
		return m.MaxTokens
	}
	return 0
}

func (m *TokenBucket) GetTokensPerFill() *wrappers.UInt32Value {
	if m != nil {
		return m.TokensPerFill
	}
	return nil
}

func (m *TokenBucket) GetFillInterval() *duration.Duration {
	if m != nil {
		return m.FillInterval
	}
	return nil
}

func init() {
	proto.RegisterType((*TokenBucket)(nil), "envoy.type.v3.TokenBucket")
}

func init() { proto.RegisterFile("envoy/type/v3/token_bucket.proto", fileDescriptor_5b88f84e3f7d4627) }

var fileDescriptor_5b88f84e3f7d4627 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x4f, 0x4a, 0x03, 0x31,
	0x18, 0xc5, 0x9b, 0x5a, 0xac, 0xa6, 0x0e, 0xca, 0x2c, 0x74, 0xfc, 0x57, 0x46, 0x17, 0x52, 0xba,
	0x48, 0xb0, 0xb3, 0x73, 0x19, 0x44, 0xa8, 0x20, 0x94, 0xa2, 0x6e, 0x87, 0xd4, 0xa6, 0x25, 0x34,
	0x4d, 0x42, 0x26, 0x13, 0xdb, 0x1b, 0x78, 0x06, 0x8f, 0xe0, 0x3d, 0x3c, 0x8f, 0xfb, 0xae, 0x64,
	0x32, 0x53, 0xac, 0x74, 0x17, 0x78, 0xbf, 0xf7, 0xbe, 0x97, 0xef, 0x83, 0x31, 0x93, 0x4e, 0x2d,
	0xb1, 0x5d, 0x6a, 0x86, 0x5d, 0x82, 0xad, 0x9a, 0x31, 0x99, 0x8e, 0xf2, 0xb7, 0x19, 0xb3, 0x48,
	0x1b, 0x65, 0x55, 0x18, 0x78, 0x02, 0x15, 0x04, 0x72, 0xc9, 0x59, 0x7b, 0xaa, 0xd4, 0x54, 0x30,
	0xec, 0xc5, 0x51, 0x3e, 0xc1, 0xe3, 0xdc, 0x50, 0xcb, 0x95, 0x2c, 0xf1, 0x6d, 0xfd, 0xdd, 0x50,
	0xad, 0x99, 0xc9, 0x2a, 0xfd, 0x2a, 0x1f, 0x6b, 0x8a, 0xa9, 0x94, 0xca, 0x7a, 0x5b, 0x86, 0x1d,
	0x33, 0x19, 0x57, 0x92, 0xcb, 0x69, 0x85, 0x9c, 0x38, 0x2a, 0xf8, 0x98, 0x5a, 0x86, 0xd7, 0x8f,
	0x52, 0xb8, 0xfe, 0x01, 0xb0, 0xf5, 0x5c, 0x34, 0x24, 0xbe, 0x60, 0x78, 0x03, 0xe1, 0x9c, 0x2e,
	0x52, 0x5f, 0x3a, 0x8b, 0x40, 0x0c, 0x3a, 0x01, 0x69, 0xae, 0x48, 0xa3, 0x5b, 0x8f, 0x6b, 0xc3,
	0xfd, 0x39, 0x5d, 0x78, 0x38, 0x0b, 0x9f, 0xe0, 0x61, 0xc9, 0xa4, 0x9a, 0x99, 0x74, 0xc2, 0x85,
	0x88, 0xea, 0x31, 0xe8, 0xb4, 0x7a, 0x17, 0xa8, 0x6c, 0x8b, 0xd6, 0x6d, 0xd1, 0x4b, 0x5f, 0xda,
	0xa4, 0xf7, 0x4a, 0x45, 0xce, 0xfe, 0xa2, 0x82, 0xd2, 0x3d, 0x60, 0xe6, 0x81, 0x0b, 0x11, 0x3e,
	0xc2, 0xa0, 0xc8, 0x48, 0xb9, 0xb4, 0xcc, 0x38, 0x2a, 0xa2, 0x1d, 0x1f, 0x76, 0xba, 0x15, 0x76,
	0x5f, 0xad, 0x86, 0xc0, 0x15, 0x69, 0x7e, 0x81, 0xc6, 0x1e, 0xe8, 0xd6, 0x86, 0x07, 0x85, 0xb7,
	0x5f, 0x59, 0xef, 0x2e, 0x3f, 0xbf, 0x3f, 0xda, 0x11, 0x3c, 0xde, 0x58, 0xf2, 0xc6, 0x0f, 0xc9,
	0x2d, 0x3c, 0xe7, 0x0a, 0x79, 0x51, 0x1b, 0xb5, 0x58, 0xa2, 0x7f, 0xc7, 0x20, 0x47, 0x1b, 0xec,
	0xa0, 0x98, 0x3a, 0x00, 0xa3, 0x5d, 0x3f, 0x3e, 0xf9, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x44, 0x24,
	0x7f, 0x93, 0xda, 0x01, 0x00, 0x00,
}
