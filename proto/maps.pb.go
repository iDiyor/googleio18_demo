// Code generated by protoc-gen-go. DO NOT EDIT.
// source: maps.proto

/*
Package maps is a generated protocol buffer package.

It is generated from these files:
	maps.proto

It has these top-level messages:
	GetDistanceRequest
	GetDistanceResponse
*/
package maps

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

type GetDistanceRequest struct {
	Origin string `protobuf:"bytes,1,opt,name=origin" json:"origin,omitempty"`
	Dest   string `protobuf:"bytes,2,opt,name=dest" json:"dest,omitempty"`
}

func (m *GetDistanceRequest) Reset()                    { *m = GetDistanceRequest{} }
func (m *GetDistanceRequest) String() string            { return proto.CompactTextString(m) }
func (*GetDistanceRequest) ProtoMessage()               {}
func (*GetDistanceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetDistanceRequest) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *GetDistanceRequest) GetDest() string {
	if m != nil {
		return m.Dest
	}
	return ""
}

type GetDistanceResponse struct {
	Distance int64 `protobuf:"varint,1,opt,name=distance" json:"distance,omitempty"`
}

func (m *GetDistanceResponse) Reset()                    { *m = GetDistanceResponse{} }
func (m *GetDistanceResponse) String() string            { return proto.CompactTextString(m) }
func (*GetDistanceResponse) ProtoMessage()               {}
func (*GetDistanceResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetDistanceResponse) GetDistance() int64 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func init() {
	proto.RegisterType((*GetDistanceRequest)(nil), "GetDistanceRequest")
	proto.RegisterType((*GetDistanceResponse)(nil), "GetDistanceResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Maps service

type MapsClient interface {
	GetDistance(ctx context.Context, in *GetDistanceRequest, opts ...grpc.CallOption) (*GetDistanceResponse, error)
}

type mapsClient struct {
	cc *grpc.ClientConn
}

func NewMapsClient(cc *grpc.ClientConn) MapsClient {
	return &mapsClient{cc}
}

func (c *mapsClient) GetDistance(ctx context.Context, in *GetDistanceRequest, opts ...grpc.CallOption) (*GetDistanceResponse, error) {
	out := new(GetDistanceResponse)
	err := grpc.Invoke(ctx, "/Maps/GetDistance", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Maps service

type MapsServer interface {
	GetDistance(context.Context, *GetDistanceRequest) (*GetDistanceResponse, error)
}

func RegisterMapsServer(s *grpc.Server, srv MapsServer) {
	s.RegisterService(&_Maps_serviceDesc, srv)
}

func _Maps_GetDistance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDistanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapsServer).GetDistance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Maps/GetDistance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapsServer).GetDistance(ctx, req.(*GetDistanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Maps_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Maps",
	HandlerType: (*MapsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDistance",
			Handler:    _Maps_GetDistance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "maps.proto",
}

func init() { proto.RegisterFile("maps.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 146 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x4d, 0x2c, 0x28,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x72, 0xe0, 0x12, 0x72, 0x4f, 0x2d, 0x71, 0xc9, 0x2c,
	0x2e, 0x49, 0xcc, 0x4b, 0x4e, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe3, 0x62,
	0xcb, 0x2f, 0xca, 0x4c, 0xcf, 0xcc, 0x93, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf2, 0x84,
	0x84, 0xb8, 0x58, 0x52, 0x52, 0x8b, 0x4b, 0x24, 0x98, 0xc0, 0xa2, 0x60, 0xb6, 0x92, 0x21, 0x97,
	0x30, 0x8a, 0x09, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x52, 0x5c, 0x1c, 0x29, 0x50, 0x31,
	0xb0, 0x21, 0xcc, 0x41, 0x70, 0xbe, 0x91, 0x03, 0x17, 0x8b, 0x6f, 0x62, 0x41, 0xb1, 0x90, 0x05,
	0x17, 0x37, 0x92, 0x56, 0x21, 0x61, 0x3d, 0x4c, 0xa7, 0x48, 0x89, 0xe8, 0x61, 0x31, 0x3d, 0x89,
	0x0d, 0xec, 0x7a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x39, 0x6f, 0x17, 0xcb, 0x00,
	0x00, 0x00,
}