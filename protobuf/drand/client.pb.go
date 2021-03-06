// Code generated by protoc-gen-go. DO NOT EDIT.
// source: drand/client.proto

package drand // import "github.com/dedis/drand/protobuf/drand"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import crypto "github.com/dedis/drand/protobuf/crypto"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// PublicRandRequest requests a public random value that has been generated in a
// unbiasable way and verifiable.
type PublicRandRequest struct {
	// round uniquely identifies a beacon. If round == 0, then the response will
	// contain the last.
	// XXX better ways to do that...
	Round                uint64   `protobuf:"varint,1,opt,name=round,proto3" json:"round,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublicRandRequest) Reset()         { *m = PublicRandRequest{} }
func (m *PublicRandRequest) String() string { return proto.CompactTextString(m) }
func (*PublicRandRequest) ProtoMessage()    {}
func (*PublicRandRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_89a0ad9403dde017, []int{0}
}
func (m *PublicRandRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublicRandRequest.Unmarshal(m, b)
}
func (m *PublicRandRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublicRandRequest.Marshal(b, m, deterministic)
}
func (dst *PublicRandRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicRandRequest.Merge(dst, src)
}
func (m *PublicRandRequest) XXX_Size() int {
	return xxx_messageInfo_PublicRandRequest.Size(m)
}
func (m *PublicRandRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicRandRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublicRandRequest proto.InternalMessageInfo

func (m *PublicRandRequest) GetRound() uint64 {
	if m != nil {
		return m.Round
	}
	return 0
}

// PublicRandResponse holds a signature which is the random value. It can be
// verified thanks to the distributed public key of the nodes that have ran the
// DKG protocol and is unbiasable. The randomness can be verified using the BLS
// verification routine with the message "round || previous_rand".
type PublicRandResponse struct {
	Round                uint64        `protobuf:"varint,1,opt,name=round,proto3" json:"round,omitempty"`
	Previous             []byte        `protobuf:"bytes,2,opt,name=previous,proto3" json:"previous,omitempty"`
	Randomness           *crypto.Point `protobuf:"bytes,3,opt,name=randomness,proto3" json:"randomness,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PublicRandResponse) Reset()         { *m = PublicRandResponse{} }
func (m *PublicRandResponse) String() string { return proto.CompactTextString(m) }
func (*PublicRandResponse) ProtoMessage()    {}
func (*PublicRandResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_89a0ad9403dde017, []int{1}
}
func (m *PublicRandResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublicRandResponse.Unmarshal(m, b)
}
func (m *PublicRandResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublicRandResponse.Marshal(b, m, deterministic)
}
func (dst *PublicRandResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicRandResponse.Merge(dst, src)
}
func (m *PublicRandResponse) XXX_Size() int {
	return xxx_messageInfo_PublicRandResponse.Size(m)
}
func (m *PublicRandResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicRandResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PublicRandResponse proto.InternalMessageInfo

func (m *PublicRandResponse) GetRound() uint64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *PublicRandResponse) GetPrevious() []byte {
	if m != nil {
		return m.Previous
	}
	return nil
}

func (m *PublicRandResponse) GetRandomness() *crypto.Point {
	if m != nil {
		return m.Randomness
	}
	return nil
}

// PrivateRandRequest is the message to send when requesting a private random
// value.
type PrivateRandRequest struct {
	// Request must contains a public key towards which to encrypt the private
	// randomness.
	Request              *ECIESObject `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PrivateRandRequest) Reset()         { *m = PrivateRandRequest{} }
func (m *PrivateRandRequest) String() string { return proto.CompactTextString(m) }
func (*PrivateRandRequest) ProtoMessage()    {}
func (*PrivateRandRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_89a0ad9403dde017, []int{2}
}
func (m *PrivateRandRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrivateRandRequest.Unmarshal(m, b)
}
func (m *PrivateRandRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrivateRandRequest.Marshal(b, m, deterministic)
}
func (dst *PrivateRandRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrivateRandRequest.Merge(dst, src)
}
func (m *PrivateRandRequest) XXX_Size() int {
	return xxx_messageInfo_PrivateRandRequest.Size(m)
}
func (m *PrivateRandRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrivateRandRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrivateRandRequest proto.InternalMessageInfo

func (m *PrivateRandRequest) GetRequest() *ECIESObject {
	if m != nil {
		return m.Request
	}
	return nil
}

type PrivateRandResponse struct {
	// Response contains the private randomness encrypted towards the client's
	// request key.
	Response             *ECIESObject `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PrivateRandResponse) Reset()         { *m = PrivateRandResponse{} }
func (m *PrivateRandResponse) String() string { return proto.CompactTextString(m) }
func (*PrivateRandResponse) ProtoMessage()    {}
func (*PrivateRandResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_89a0ad9403dde017, []int{3}
}
func (m *PrivateRandResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrivateRandResponse.Unmarshal(m, b)
}
func (m *PrivateRandResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrivateRandResponse.Marshal(b, m, deterministic)
}
func (dst *PrivateRandResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrivateRandResponse.Merge(dst, src)
}
func (m *PrivateRandResponse) XXX_Size() int {
	return xxx_messageInfo_PrivateRandResponse.Size(m)
}
func (m *PrivateRandResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrivateRandResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrivateRandResponse proto.InternalMessageInfo

func (m *PrivateRandResponse) GetResponse() *ECIESObject {
	if m != nil {
		return m.Response
	}
	return nil
}

type ECIESObject struct {
	Ephemeral            *crypto.Point `protobuf:"bytes,1,opt,name=ephemeral,proto3" json:"ephemeral,omitempty"`
	Ciphertext           []byte        `protobuf:"bytes,2,opt,name=ciphertext,proto3" json:"ciphertext,omitempty"`
	Nonce                []byte        `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ECIESObject) Reset()         { *m = ECIESObject{} }
func (m *ECIESObject) String() string { return proto.CompactTextString(m) }
func (*ECIESObject) ProtoMessage()    {}
func (*ECIESObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_89a0ad9403dde017, []int{4}
}
func (m *ECIESObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ECIESObject.Unmarshal(m, b)
}
func (m *ECIESObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ECIESObject.Marshal(b, m, deterministic)
}
func (dst *ECIESObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ECIESObject.Merge(dst, src)
}
func (m *ECIESObject) XXX_Size() int {
	return xxx_messageInfo_ECIESObject.Size(m)
}
func (m *ECIESObject) XXX_DiscardUnknown() {
	xxx_messageInfo_ECIESObject.DiscardUnknown(m)
}

var xxx_messageInfo_ECIESObject proto.InternalMessageInfo

func (m *ECIESObject) GetEphemeral() *crypto.Point {
	if m != nil {
		return m.Ephemeral
	}
	return nil
}

func (m *ECIESObject) GetCiphertext() []byte {
	if m != nil {
		return m.Ciphertext
	}
	return nil
}

func (m *ECIESObject) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

// DistKeyRequest requests the distributed public key used during the randomness generation process
type DistKeyRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DistKeyRequest) Reset()         { *m = DistKeyRequest{} }
func (m *DistKeyRequest) String() string { return proto.CompactTextString(m) }
func (*DistKeyRequest) ProtoMessage()    {}
func (*DistKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_89a0ad9403dde017, []int{5}
}
func (m *DistKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistKeyRequest.Unmarshal(m, b)
}
func (m *DistKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistKeyRequest.Marshal(b, m, deterministic)
}
func (dst *DistKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistKeyRequest.Merge(dst, src)
}
func (m *DistKeyRequest) XXX_Size() int {
	return xxx_messageInfo_DistKeyRequest.Size(m)
}
func (m *DistKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DistKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DistKeyRequest proto.InternalMessageInfo

type DistKeyResponse struct {
	Key                  *crypto.Point `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DistKeyResponse) Reset()         { *m = DistKeyResponse{} }
func (m *DistKeyResponse) String() string { return proto.CompactTextString(m) }
func (*DistKeyResponse) ProtoMessage()    {}
func (*DistKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_89a0ad9403dde017, []int{6}
}
func (m *DistKeyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistKeyResponse.Unmarshal(m, b)
}
func (m *DistKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistKeyResponse.Marshal(b, m, deterministic)
}
func (dst *DistKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistKeyResponse.Merge(dst, src)
}
func (m *DistKeyResponse) XXX_Size() int {
	return xxx_messageInfo_DistKeyResponse.Size(m)
}
func (m *DistKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DistKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DistKeyResponse proto.InternalMessageInfo

func (m *DistKeyResponse) GetKey() *crypto.Point {
	if m != nil {
		return m.Key
	}
	return nil
}

type HomeRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HomeRequest) Reset()         { *m = HomeRequest{} }
func (m *HomeRequest) String() string { return proto.CompactTextString(m) }
func (*HomeRequest) ProtoMessage()    {}
func (*HomeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_89a0ad9403dde017, []int{7}
}
func (m *HomeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HomeRequest.Unmarshal(m, b)
}
func (m *HomeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HomeRequest.Marshal(b, m, deterministic)
}
func (dst *HomeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HomeRequest.Merge(dst, src)
}
func (m *HomeRequest) XXX_Size() int {
	return xxx_messageInfo_HomeRequest.Size(m)
}
func (m *HomeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HomeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HomeRequest proto.InternalMessageInfo

type HomeResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HomeResponse) Reset()         { *m = HomeResponse{} }
func (m *HomeResponse) String() string { return proto.CompactTextString(m) }
func (*HomeResponse) ProtoMessage()    {}
func (*HomeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_89a0ad9403dde017, []int{8}
}
func (m *HomeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HomeResponse.Unmarshal(m, b)
}
func (m *HomeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HomeResponse.Marshal(b, m, deterministic)
}
func (dst *HomeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HomeResponse.Merge(dst, src)
}
func (m *HomeResponse) XXX_Size() int {
	return xxx_messageInfo_HomeResponse.Size(m)
}
func (m *HomeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HomeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HomeResponse proto.InternalMessageInfo

func (m *HomeResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*PublicRandRequest)(nil), "drand.PublicRandRequest")
	proto.RegisterType((*PublicRandResponse)(nil), "drand.PublicRandResponse")
	proto.RegisterType((*PrivateRandRequest)(nil), "drand.PrivateRandRequest")
	proto.RegisterType((*PrivateRandResponse)(nil), "drand.PrivateRandResponse")
	proto.RegisterType((*ECIESObject)(nil), "drand.ECIESObject")
	proto.RegisterType((*DistKeyRequest)(nil), "drand.DistKeyRequest")
	proto.RegisterType((*DistKeyResponse)(nil), "drand.DistKeyResponse")
	proto.RegisterType((*HomeRequest)(nil), "drand.HomeRequest")
	proto.RegisterType((*HomeResponse)(nil), "drand.HomeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RandomnessClient is the client API for Randomness service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RandomnessClient interface {
	Public(ctx context.Context, in *PublicRandRequest, opts ...grpc.CallOption) (*PublicRandResponse, error)
	Private(ctx context.Context, in *PrivateRandRequest, opts ...grpc.CallOption) (*PrivateRandResponse, error)
}

type randomnessClient struct {
	cc *grpc.ClientConn
}

func NewRandomnessClient(cc *grpc.ClientConn) RandomnessClient {
	return &randomnessClient{cc}
}

func (c *randomnessClient) Public(ctx context.Context, in *PublicRandRequest, opts ...grpc.CallOption) (*PublicRandResponse, error) {
	out := new(PublicRandResponse)
	err := c.cc.Invoke(ctx, "/drand.Randomness/Public", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *randomnessClient) Private(ctx context.Context, in *PrivateRandRequest, opts ...grpc.CallOption) (*PrivateRandResponse, error) {
	out := new(PrivateRandResponse)
	err := c.cc.Invoke(ctx, "/drand.Randomness/Private", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RandomnessServer is the server API for Randomness service.
type RandomnessServer interface {
	Public(context.Context, *PublicRandRequest) (*PublicRandResponse, error)
	Private(context.Context, *PrivateRandRequest) (*PrivateRandResponse, error)
}

func RegisterRandomnessServer(s *grpc.Server, srv RandomnessServer) {
	s.RegisterService(&_Randomness_serviceDesc, srv)
}

func _Randomness_Public_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublicRandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RandomnessServer).Public(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drand.Randomness/Public",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RandomnessServer).Public(ctx, req.(*PublicRandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Randomness_Private_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrivateRandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RandomnessServer).Private(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drand.Randomness/Private",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RandomnessServer).Private(ctx, req.(*PrivateRandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Randomness_serviceDesc = grpc.ServiceDesc{
	ServiceName: "drand.Randomness",
	HandlerType: (*RandomnessServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Public",
			Handler:    _Randomness_Public_Handler,
		},
		{
			MethodName: "Private",
			Handler:    _Randomness_Private_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "drand/client.proto",
}

// InfoClient is the client API for Info service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InfoClient interface {
	DistKey(ctx context.Context, in *DistKeyRequest, opts ...grpc.CallOption) (*DistKeyResponse, error)
	Home(ctx context.Context, in *HomeRequest, opts ...grpc.CallOption) (*HomeResponse, error)
}

type infoClient struct {
	cc *grpc.ClientConn
}

func NewInfoClient(cc *grpc.ClientConn) InfoClient {
	return &infoClient{cc}
}

func (c *infoClient) DistKey(ctx context.Context, in *DistKeyRequest, opts ...grpc.CallOption) (*DistKeyResponse, error) {
	out := new(DistKeyResponse)
	err := c.cc.Invoke(ctx, "/drand.Info/DistKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoClient) Home(ctx context.Context, in *HomeRequest, opts ...grpc.CallOption) (*HomeResponse, error) {
	out := new(HomeResponse)
	err := c.cc.Invoke(ctx, "/drand.Info/Home", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InfoServer is the server API for Info service.
type InfoServer interface {
	DistKey(context.Context, *DistKeyRequest) (*DistKeyResponse, error)
	Home(context.Context, *HomeRequest) (*HomeResponse, error)
}

func RegisterInfoServer(s *grpc.Server, srv InfoServer) {
	s.RegisterService(&_Info_serviceDesc, srv)
}

func _Info_DistKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DistKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServer).DistKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drand.Info/DistKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServer).DistKey(ctx, req.(*DistKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Info_Home_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServer).Home(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drand.Info/Home",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServer).Home(ctx, req.(*HomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Info_serviceDesc = grpc.ServiceDesc{
	ServiceName: "drand.Info",
	HandlerType: (*InfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DistKey",
			Handler:    _Info_DistKey_Handler,
		},
		{
			MethodName: "Home",
			Handler:    _Info_Home_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "drand/client.proto",
}

func init() { proto.RegisterFile("drand/client.proto", fileDescriptor_client_89a0ad9403dde017) }

var fileDescriptor_client_89a0ad9403dde017 = []byte{
	// 528 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xe5, 0x36, 0x4d, 0xda, 0x49, 0x28, 0x64, 0xd2, 0x96, 0xd4, 0x42, 0x28, 0xb2, 0xf8,
	0x13, 0xaa, 0xca, 0x96, 0xd2, 0x1b, 0x12, 0x97, 0x42, 0x24, 0x2a, 0x0e, 0x44, 0xee, 0x89, 0xde,
	0x1c, 0x7b, 0x92, 0x2c, 0x4d, 0x76, 0x5d, 0xef, 0x3a, 0x22, 0x42, 0x5c, 0x78, 0x05, 0x6e, 0x3c,
	0x16, 0xbc, 0x02, 0x0f, 0x82, 0xb2, 0xbb, 0x31, 0x4e, 0x1b, 0x7a, 0xf3, 0x37, 0xf3, 0xf9, 0xb7,
	0xb3, 0x9f, 0xc7, 0x80, 0x49, 0x16, 0xf1, 0x24, 0x88, 0xa7, 0x8c, 0xb8, 0xf2, 0xd3, 0x4c, 0x28,
	0x81, 0x3b, 0xba, 0xe6, 0x1e, 0xc4, 0xd9, 0x22, 0x55, 0x22, 0xa0, 0x29, 0xcd, 0x8a, 0xa6, 0xfb,
	0x64, 0x2c, 0xc4, 0x78, 0x4a, 0x41, 0x94, 0xb2, 0x20, 0xe2, 0x5c, 0xa8, 0x48, 0x31, 0xc1, 0xa5,
	0xe9, 0x7a, 0xaf, 0xa0, 0x39, 0xc8, 0x87, 0x53, 0x16, 0x87, 0x11, 0x4f, 0x42, 0xba, 0xc9, 0x49,
	0x2a, 0x3c, 0x80, 0x9d, 0x4c, 0xe4, 0x3c, 0x69, 0x3b, 0x1d, 0xa7, 0x5b, 0x09, 0x8d, 0xf0, 0xe6,
	0x80, 0x65, 0xab, 0x4c, 0x05, 0x97, 0xb4, 0xd9, 0x8b, 0x2e, 0xec, 0xa6, 0x19, 0xcd, 0x99, 0xc8,
	0x65, 0x7b, 0xab, 0xe3, 0x74, 0x1b, 0x61, 0xa1, 0xd1, 0x07, 0x58, 0x8e, 0x2b, 0x66, 0x9c, 0xa4,
	0x6c, 0x6f, 0x77, 0x9c, 0x6e, 0xbd, 0xb7, 0xef, 0xaf, 0x86, 0x1e, 0x08, 0xc6, 0x55, 0x58, 0x72,
	0x78, 0xe7, 0x80, 0x83, 0x8c, 0xcd, 0x23, 0x45, 0xe5, 0x19, 0x4f, 0xa1, 0x96, 0x99, 0x47, 0x7d,
	0x72, 0xbd, 0x87, 0xbe, 0x4e, 0xc1, 0xef, 0xbf, 0xbd, 0xe8, 0x5f, 0x7e, 0x1c, 0x7e, 0xa6, 0x58,
	0x85, 0x2b, 0x8b, 0xd7, 0x87, 0xd6, 0x1a, 0xc3, 0x0e, 0xef, 0xc3, 0x6e, 0x66, 0x9f, 0xef, 0xa1,
	0x14, 0x1e, 0xef, 0x06, 0xea, 0xa5, 0x06, 0x9e, 0xc2, 0x1e, 0xa5, 0x13, 0x9a, 0x51, 0x16, 0x4d,
	0xed, 0xfb, 0xb7, 0x2f, 0xf2, 0xcf, 0x80, 0x4f, 0x01, 0x62, 0x96, 0x4e, 0x28, 0x53, 0xf4, 0x45,
	0xd9, 0x54, 0x4a, 0x95, 0x65, 0x92, 0x5c, 0xf0, 0x98, 0x74, 0x24, 0x8d, 0xd0, 0x08, 0xef, 0x11,
	0xec, 0xbf, 0x63, 0x52, 0x7d, 0xa0, 0x85, 0xbd, 0xb9, 0x77, 0x06, 0x0f, 0x8b, 0x8a, 0xbd, 0x47,
	0x07, 0xb6, 0xaf, 0x69, 0xf1, 0x9f, 0x11, 0x96, 0x2d, 0xef, 0x01, 0xd4, 0xdf, 0x8b, 0x19, 0xad,
	0x18, 0x2f, 0xa0, 0x61, 0xa4, 0x05, 0x1c, 0x41, 0x55, 0xaa, 0x48, 0xe5, 0x52, 0x33, 0xf6, 0x42,
	0xab, 0x7a, 0xbf, 0x1c, 0x80, 0xb0, 0xf8, 0x14, 0xc8, 0xa0, 0x6a, 0x56, 0x00, 0xdb, 0x36, 0xa7,
	0x3b, 0xcb, 0xe3, 0x1e, 0x6f, 0xe8, 0xd8, 0xf8, 0x4e, 0xbe, 0xff, 0xfe, 0xf3, 0x63, 0xeb, 0x19,
	0xd6, 0xf5, 0x32, 0xa6, 0xda, 0x70, 0x75, 0x88, 0xad, 0x92, 0x0c, 0xbe, 0xea, 0x05, 0xfa, 0x86,
	0x9f, 0xa0, 0x66, 0xbf, 0x18, 0x16, 0xc4, 0x3b, 0x5b, 0xe0, 0xba, 0x9b, 0x5a, 0xf6, 0xb4, 0xc7,
	0xfa, 0xb4, 0xa6, 0xd7, 0x30, 0x78, 0xe3, 0x78, 0xed, 0x9c, 0xf4, 0x7e, 0x3a, 0x50, 0xb9, 0xe0,
	0x23, 0x81, 0x97, 0x50, 0xb3, 0x49, 0xe2, 0xa1, 0x05, 0xad, 0x67, 0xed, 0x1e, 0xdd, 0x2e, 0x5b,
	0xf6, 0xb1, 0x66, 0xb7, 0xb0, 0xa9, 0xd9, 0x8c, 0x8f, 0x44, 0x90, 0x30, 0xa9, 0xae, 0x69, 0x81,
	0x6f, 0xa0, 0xb2, 0x8c, 0x16, 0x57, 0x9b, 0x54, 0x8a, 0xdd, 0x6d, 0xad, 0xd5, 0x2c, 0xab, 0xa1,
	0x59, 0x55, 0xac, 0x2c, 0x59, 0xe7, 0x2f, 0xaf, 0x9e, 0x8f, 0x99, 0x9a, 0xe4, 0x43, 0x3f, 0x16,
	0xb3, 0x20, 0xa1, 0x84, 0xc9, 0xc0, 0xfc, 0xf2, 0xfa, 0x87, 0x1d, 0xe6, 0x23, 0x23, 0x87, 0x55,
	0xad, 0xcf, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xd2, 0xe2, 0x59, 0xee, 0x11, 0x04, 0x00, 0x00,
}
