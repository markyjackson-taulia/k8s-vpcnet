// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipam.proto

/*
Package vpcnetpb is a generated protocol buffer package.

It is generated from these files:
	ipam.proto

It has these top-level messages:
	AddRequest
	AddResponse
	DelRequest
	DelResponse
*/
package vpcnetpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type AddRequest struct {
	ContainerID  string `protobuf:"bytes,1,opt,name=ContainerID" json:"ContainerID,omitempty"`
	PodName      string `protobuf:"bytes,2,opt,name=PodName" json:"PodName,omitempty"`
	PodNamespace string `protobuf:"bytes,3,opt,name=PodNamespace" json:"PodNamespace,omitempty"`
}

func (m *AddRequest) Reset()                    { *m = AddRequest{} }
func (m *AddRequest) String() string            { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()               {}
func (*AddRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddRequest) GetContainerID() string {
	if m != nil {
		return m.ContainerID
	}
	return ""
}

func (m *AddRequest) GetPodName() string {
	if m != nil {
		return m.PodName
	}
	return ""
}

func (m *AddRequest) GetPodNamespace() string {
	if m != nil {
		return m.PodNamespace
	}
	return ""
}

type AddResponse struct {
	AllocatedIP string `protobuf:"bytes,1,opt,name=AllocatedIP" json:"AllocatedIP,omitempty"`
	ENIIP       string `protobuf:"bytes,2,opt,name=ENIIP" json:"ENIIP,omitempty"`
	SubnetCIDR  string `protobuf:"bytes,3,opt,name=SubnetCIDR" json:"SubnetCIDR,omitempty"`
}

func (m *AddResponse) Reset()                    { *m = AddResponse{} }
func (m *AddResponse) String() string            { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()               {}
func (*AddResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddResponse) GetAllocatedIP() string {
	if m != nil {
		return m.AllocatedIP
	}
	return ""
}

func (m *AddResponse) GetENIIP() string {
	if m != nil {
		return m.ENIIP
	}
	return ""
}

func (m *AddResponse) GetSubnetCIDR() string {
	if m != nil {
		return m.SubnetCIDR
	}
	return ""
}

type DelRequest struct {
	ContainerID string `protobuf:"bytes,1,opt,name=ContainerID" json:"ContainerID,omitempty"`
}

func (m *DelRequest) Reset()                    { *m = DelRequest{} }
func (m *DelRequest) String() string            { return proto.CompactTextString(m) }
func (*DelRequest) ProtoMessage()               {}
func (*DelRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DelRequest) GetContainerID() string {
	if m != nil {
		return m.ContainerID
	}
	return ""
}

type DelResponse struct {
}

func (m *DelResponse) Reset()                    { *m = DelResponse{} }
func (m *DelResponse) String() string            { return proto.CompactTextString(m) }
func (*DelResponse) ProtoMessage()               {}
func (*DelResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*AddRequest)(nil), "vpcnet.AddRequest")
	proto.RegisterType((*AddResponse)(nil), "vpcnet.AddResponse")
	proto.RegisterType((*DelRequest)(nil), "vpcnet.DelRequest")
	proto.RegisterType((*DelResponse)(nil), "vpcnet.DelResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for IPAM service

type IPAMClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	Del(ctx context.Context, in *DelRequest, opts ...grpc.CallOption) (*DelResponse, error)
}

type iPAMClient struct {
	cc *grpc.ClientConn
}

func NewIPAMClient(cc *grpc.ClientConn) IPAMClient {
	return &iPAMClient{cc}
}

func (c *iPAMClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := grpc.Invoke(ctx, "/vpcnet.IPAM/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPAMClient) Del(ctx context.Context, in *DelRequest, opts ...grpc.CallOption) (*DelResponse, error) {
	out := new(DelResponse)
	err := grpc.Invoke(ctx, "/vpcnet.IPAM/Del", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IPAM service

type IPAMServer interface {
	Add(context.Context, *AddRequest) (*AddResponse, error)
	Del(context.Context, *DelRequest) (*DelResponse, error)
}

func RegisterIPAMServer(s *grpc.Server, srv IPAMServer) {
	s.RegisterService(&_IPAM_serviceDesc, srv)
}

func _IPAM_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPAMServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vpcnet.IPAM/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPAMServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPAM_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPAMServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vpcnet.IPAM/Del",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPAMServer).Del(ctx, req.(*DelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IPAM_serviceDesc = grpc.ServiceDesc{
	ServiceName: "vpcnet.IPAM",
	HandlerType: (*IPAMServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _IPAM_Add_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _IPAM_Del_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ipam.proto",
}

func init() { proto.RegisterFile("ipam.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x5b, 0x0a, 0x05, 0xfe, 0xc2, 0x72, 0x30, 0x44, 0x0c, 0xa8, 0xf2, 0xc4, 0x14, 0x21,
	0x78, 0x82, 0xd0, 0x30, 0x78, 0xa0, 0x8a, 0xc2, 0xc6, 0xe6, 0xc4, 0x37, 0x14, 0xb9, 0xb6, 0x69,
	0x5c, 0x9e, 0x1f, 0xd5, 0x4e, 0x95, 0x74, 0x63, 0xf3, 0x7d, 0xe7, 0x5f, 0xdf, 0xf9, 0x0c, 0x6c,
	0xbc, 0xda, 0xe6, 0x7e, 0xe7, 0x82, 0xa3, 0xf9, 0xaf, 0x6f, 0x2d, 0x07, 0x61, 0x80, 0x42, 0xeb,
	0x9a, 0x7f, 0xf6, 0xdc, 0x05, 0x5a, 0x62, 0xb1, 0x72, 0x36, 0xa8, 0x8d, 0xe5, 0x9d, 0x2c, 0xb3,
	0xe9, 0x72, 0xfa, 0x74, 0x5d, 0x8f, 0x11, 0x65, 0xb8, 0xac, 0x9c, 0x5e, 0xab, 0x2d, 0x67, 0x67,
	0xb1, 0x7b, 0x2c, 0x49, 0xe0, 0xa6, 0x3f, 0x76, 0x5e, 0xb5, 0x9c, 0xcd, 0x62, 0xfb, 0x84, 0x09,
	0xc6, 0x22, 0xda, 0x3a, 0xef, 0x6c, 0xc7, 0x07, 0x5d, 0x61, 0x8c, 0x6b, 0x55, 0x60, 0x2d, 0xab,
	0xa3, 0x6e, 0x84, 0xe8, 0x1e, 0x17, 0xef, 0x6b, 0x29, 0xab, 0x5e, 0x96, 0x0a, 0x7a, 0x04, 0x3e,
	0xf7, 0x8d, 0xe5, 0xb0, 0x92, 0x65, 0xdd, 0x8b, 0x46, 0x44, 0xe4, 0x40, 0xc9, 0xe6, 0xdf, 0x8f,
	0x12, 0xb7, 0x58, 0xc4, 0xfb, 0x69, 0xac, 0x97, 0x6f, 0x9c, 0xcb, 0xaa, 0xf8, 0xa0, 0x67, 0xcc,
	0x0a, 0xad, 0x89, 0xf2, 0xb4, 0xab, 0x7c, 0x58, 0xd4, 0xc3, 0xdd, 0x09, 0x4b, 0x39, 0x31, 0x39,
	0x24, 0x4a, 0x36, 0x43, 0x62, 0x98, 0x62, 0x48, 0x8c, 0x4c, 0x62, 0xf2, 0x86, 0xaf, 0xab, 0xc4,
	0x7d, 0xd3, 0xcc, 0xe3, 0xd7, 0xbc, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x37, 0xe0, 0xcb,
	0xa8, 0x01, 0x00, 0x00,
}