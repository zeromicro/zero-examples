// Code generated by protoc-gen-go.
// source: portal.proto
// DO NOT EDIT!

/*
Package portal is a generated protocol buffer package.

It is generated from these files:
	portal.proto

It has these top-level messages:
	PortalRequest
	PortalResponse
*/
package portal

import (
	"fmt"
	"math"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PortalRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *PortalRequest) Reset()                    { *m = PortalRequest{} }
func (m *PortalRequest) String() string            { return proto.CompactTextString(m) }
func (*PortalRequest) ProtoMessage()               {}
func (*PortalRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PortalRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PortalResponse struct {
	Response string `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
}

func (m *PortalResponse) Reset()                    { *m = PortalResponse{} }
func (m *PortalResponse) String() string            { return proto.CompactTextString(m) }
func (*PortalResponse) ProtoMessage()               {}
func (*PortalResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PortalResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func init() {
	proto.RegisterType((*PortalRequest)(nil), "portal.PortalRequest")
	proto.RegisterType((*PortalResponse)(nil), "portal.PortalResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ context.Context
	_ grpc.ClientConn
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Portal service

type PortalClient interface {
	Portal(ctx context.Context, in *PortalRequest, opts ...grpc.CallOption) (*PortalResponse, error)
}

type portalClient struct {
	cc *grpc.ClientConn
}

func NewPortalClient(cc *grpc.ClientConn) PortalClient {
	return &portalClient{cc}
}

func (c *portalClient) Portal(ctx context.Context, in *PortalRequest, opts ...grpc.CallOption) (*PortalResponse, error) {
	out := new(PortalResponse)
	err := grpc.Invoke(ctx, "/portal.Portal/Portal", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Portal service

type PortalServer interface {
	Portal(context.Context, *PortalRequest) (*PortalResponse, error)
}

func RegisterPortalServer(s *grpc.Server, srv PortalServer) {
	s.RegisterService(&_Portal_serviceDesc, srv)
}

func _Portal_Portal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PortalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalServer).Portal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/portal.Portal/Portal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalServer).Portal(ctx, req.(*PortalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Portal_serviceDesc = grpc.ServiceDesc{
	ServiceName: "portal.Portal",
	HandlerType: (*PortalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Portal",
			Handler:    _Portal_Portal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "portal.proto",
}

func init() { proto.RegisterFile("portal.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 122 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0xc8, 0x2f, 0x2a,
	0x49, 0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0x94, 0xb9, 0x78,
	0x03, 0xc0, 0xac, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4,
	0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x49, 0x87, 0x8b, 0x0f, 0xa6,
	0xa8, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x8a, 0x8b, 0xa3, 0x08, 0xca, 0x86, 0xaa, 0x84,
	0xf3, 0x8d, 0x1c, 0xb9, 0xd8, 0x20, 0xaa, 0x85, 0xcc, 0xe1, 0x2c, 0x51, 0x3d, 0xa8, 0xed, 0x28,
	0x96, 0x49, 0x89, 0xa1, 0x0b, 0x43, 0x8c, 0x48, 0x62, 0x03, 0x3b, 0xd2, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0xce, 0x05, 0x16, 0xd0, 0xb4, 0x00, 0x00, 0x00,
}
