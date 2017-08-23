// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	messages.proto

It has these top-level messages:
	Token
	TokenResponse
*/
package main

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

type Token struct {
	AccessToken  string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	TokenType    string `protobuf:"bytes,2,opt,name=token_type,json=tokenType" json:"token_type,omitempty"`
	RefreshToken string `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken" json:"refresh_token,omitempty"`
	ExpiresAt    int64  `protobuf:"varint,4,opt,name=expires_at,json=expiresAt" json:"expires_at,omitempty"`
	Counter      int64  `protobuf:"varint,5,opt,name=counter" json:"counter,omitempty"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Token) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *Token) GetTokenType() string {
	if m != nil {
		return m.TokenType
	}
	return ""
}

func (m *Token) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *Token) GetExpiresAt() int64 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

func (m *Token) GetCounter() int64 {
	if m != nil {
		return m.Counter
	}
	return 0
}

type TokenResponse struct {
	MyReply string `protobuf:"bytes,1,opt,name=my_reply,json=myReply" json:"my_reply,omitempty"`
}

func (m *TokenResponse) Reset()                    { *m = TokenResponse{} }
func (m *TokenResponse) String() string            { return proto.CompactTextString(m) }
func (*TokenResponse) ProtoMessage()               {}
func (*TokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TokenResponse) GetMyReply() string {
	if m != nil {
		return m.MyReply
	}
	return ""
}

func init() {
	proto.RegisterType((*Token)(nil), "main.Token")
	proto.RegisterType((*TokenResponse)(nil), "main.TokenResponse")
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0xcd, 0x4e, 0x85, 0x30,
	0x10, 0x85, 0x53, 0x01, 0x91, 0x11, 0x5c, 0x74, 0x55, 0x17, 0x26, 0x88, 0x1b, 0xe2, 0xc2, 0x8d,
	0x4f, 0xe0, 0x2b, 0x34, 0xec, 0x9b, 0x4a, 0x46, 0x25, 0xda, 0x9f, 0x74, 0x6a, 0x62, 0x9f, 0xc8,
	0xd7, 0xbc, 0xa1, 0xf4, 0xde, 0xe5, 0xf9, 0xce, 0xc9, 0x97, 0x19, 0xb8, 0x33, 0x48, 0xa4, 0x3f,
	0x91, 0x5e, 0x7c, 0x70, 0xd1, 0xf1, 0xda, 0xe8, 0xcd, 0x4e, 0xff, 0x0c, 0x9a, 0xc5, 0x7d, 0xa3,
	0xe5, 0x8f, 0xd0, 0xeb, 0x75, 0x45, 0x22, 0x15, 0xf7, 0x2c, 0xd8, 0xc8, 0xe6, 0x4e, 0xde, 0x1e,
	0xec, 0x98, 0x3c, 0x00, 0xe4, 0x4e, 0xc5, 0xe4, 0x51, 0x5c, 0xe5, 0x41, 0x97, 0xc9, 0x92, 0x3c,
	0xf2, 0x27, 0x18, 0x02, 0x7e, 0x04, 0xa4, 0xaf, 0xa2, 0xa8, 0xf2, 0xa2, 0x2f, 0xf0, 0xe2, 0xc0,
	0x3f, 0xbf, 0x05, 0x24, 0xa5, 0xa3, 0xa8, 0x47, 0x36, 0x57, 0xb2, 0x2b, 0xe4, 0x2d, 0x72, 0x01,
	0xed, 0xea, 0x7e, 0x6d, 0xc4, 0x20, 0x9a, 0xdc, 0x9d, 0xe3, 0xf4, 0x0c, 0x43, 0x36, 0x48, 0x24,
	0xef, 0x2c, 0x21, 0xbf, 0x87, 0x1b, 0x93, 0x54, 0x40, 0xff, 0x93, 0xca, 0xb1, 0xad, 0x49, 0x72,
	0x8f, 0xef, 0xd7, 0xf9, 0xc5, 0xd7, 0x53, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4a, 0xaf, 0xfa, 0xfe,
	0xf4, 0x00, 0x00, 0x00,
}
