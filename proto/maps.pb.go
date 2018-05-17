// Code generated by protoc-gen-go. DO NOT EDIT.
// source: maps.proto

/*
Package maps is a generated protocol buffer package.

It is generated from these files:
	maps.proto

It has these top-level messages:
	GetDistanceRequest
	GetDistanceResponse
	Location
	EmptyResponse
	GetVenuesRequest
	GetVenuesResponse
	Venue
*/
package maps

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
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

type Location struct {
	Lon float64 `protobuf:"fixed64,1,opt,name=lon" json:"lon,omitempty"`
	Lat float64 `protobuf:"fixed64,2,opt,name=lat" json:"lat,omitempty"`
}

func (m *Location) Reset()                    { *m = Location{} }
func (m *Location) String() string            { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()               {}
func (*Location) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Location) GetLon() float64 {
	if m != nil {
		return m.Lon
	}
	return 0
}

func (m *Location) GetLat() float64 {
	if m != nil {
		return m.Lat
	}
	return 0
}

type EmptyResponse struct {
}

func (m *EmptyResponse) Reset()                    { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string            { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()               {}
func (*EmptyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type GetVenuesRequest struct {
	Location string `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`
	Radius   int64  `protobuf:"varint,2,opt,name=radius" json:"radius,omitempty"`
}

func (m *GetVenuesRequest) Reset()                    { *m = GetVenuesRequest{} }
func (m *GetVenuesRequest) String() string            { return proto.CompactTextString(m) }
func (*GetVenuesRequest) ProtoMessage()               {}
func (*GetVenuesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetVenuesRequest) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *GetVenuesRequest) GetRadius() int64 {
	if m != nil {
		return m.Radius
	}
	return 0
}

type GetVenuesResponse struct {
	Venues []*Venue `protobuf:"bytes,1,rep,name=venues" json:"venues,omitempty"`
}

func (m *GetVenuesResponse) Reset()                    { *m = GetVenuesResponse{} }
func (m *GetVenuesResponse) String() string            { return proto.CompactTextString(m) }
func (*GetVenuesResponse) ProtoMessage()               {}
func (*GetVenuesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetVenuesResponse) GetVenues() []*Venue {
	if m != nil {
		return m.Venues
	}
	return nil
}

type Venue struct {
	Id      string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name    string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Address string  `protobuf:"bytes,3,opt,name=address" json:"address,omitempty"`
	Rating  float64 `protobuf:"fixed64,4,opt,name=rating" json:"rating,omitempty"`
}

func (m *Venue) Reset()                    { *m = Venue{} }
func (m *Venue) String() string            { return proto.CompactTextString(m) }
func (*Venue) ProtoMessage()               {}
func (*Venue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Venue) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Venue) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Venue) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Venue) GetRating() float64 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func init() {
	proto.RegisterType((*GetDistanceRequest)(nil), "GetDistanceRequest")
	proto.RegisterType((*GetDistanceResponse)(nil), "GetDistanceResponse")
	proto.RegisterType((*Location)(nil), "Location")
	proto.RegisterType((*EmptyResponse)(nil), "EmptyResponse")
	proto.RegisterType((*GetVenuesRequest)(nil), "GetVenuesRequest")
	proto.RegisterType((*GetVenuesResponse)(nil), "GetVenuesResponse")
	proto.RegisterType((*Venue)(nil), "Venue")
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
	StreamLocation(ctx context.Context, opts ...grpc.CallOption) (Maps_StreamLocationClient, error)
	GetVenues(ctx context.Context, in *GetVenuesRequest, opts ...grpc.CallOption) (*GetVenuesResponse, error)
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

func (c *mapsClient) StreamLocation(ctx context.Context, opts ...grpc.CallOption) (Maps_StreamLocationClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Maps_serviceDesc.Streams[0], c.cc, "/Maps/StreamLocation", opts...)
	if err != nil {
		return nil, err
	}
	x := &mapsStreamLocationClient{stream}
	return x, nil
}

type Maps_StreamLocationClient interface {
	Send(*Location) error
	CloseAndRecv() (*EmptyResponse, error)
	grpc.ClientStream
}

type mapsStreamLocationClient struct {
	grpc.ClientStream
}

func (x *mapsStreamLocationClient) Send(m *Location) error {
	return x.ClientStream.SendMsg(m)
}

func (x *mapsStreamLocationClient) CloseAndRecv() (*EmptyResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(EmptyResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mapsClient) GetVenues(ctx context.Context, in *GetVenuesRequest, opts ...grpc.CallOption) (*GetVenuesResponse, error) {
	out := new(GetVenuesResponse)
	err := grpc.Invoke(ctx, "/Maps/GetVenues", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Maps service

type MapsServer interface {
	GetDistance(context.Context, *GetDistanceRequest) (*GetDistanceResponse, error)
	StreamLocation(Maps_StreamLocationServer) error
	GetVenues(context.Context, *GetVenuesRequest) (*GetVenuesResponse, error)
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

func _Maps_StreamLocation_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MapsServer).StreamLocation(&mapsStreamLocationServer{stream})
}

type Maps_StreamLocationServer interface {
	SendAndClose(*EmptyResponse) error
	Recv() (*Location, error)
	grpc.ServerStream
}

type mapsStreamLocationServer struct {
	grpc.ServerStream
}

func (x *mapsStreamLocationServer) SendAndClose(m *EmptyResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *mapsStreamLocationServer) Recv() (*Location, error) {
	m := new(Location)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Maps_GetVenues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVenuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapsServer).GetVenues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Maps/GetVenues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapsServer).GetVenues(ctx, req.(*GetVenuesRequest))
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
		{
			MethodName: "GetVenues",
			Handler:    _Maps_GetVenues_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamLocation",
			Handler:       _Maps_StreamLocation_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "maps.proto",
}

func init() { proto.RegisterFile("maps.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x41, 0x6e, 0xda, 0x40,
	0x14, 0x86, 0x65, 0x9b, 0xba, 0xf0, 0x68, 0x29, 0x3c, 0xaa, 0xca, 0xb2, 0xaa, 0x0a, 0xcd, 0x8a,
	0x4d, 0xc7, 0x02, 0x36, 0x5d, 0x76, 0x91, 0xc0, 0x26, 0xd9, 0x38, 0x52, 0x76, 0x59, 0x4c, 0xf0,
	0xc8, 0x1a, 0x09, 0xcf, 0x38, 0x9e, 0x01, 0x29, 0xdb, 0x5c, 0x21, 0x87, 0xc9, 0x41, 0x72, 0x85,
	0x1c, 0x24, 0xf2, 0x78, 0x6c, 0x41, 0xc8, 0xee, 0xfd, 0xbf, 0xad, 0xff, 0xfd, 0xef, 0xb3, 0x01,
	0x0a, 0x56, 0x6a, 0x5a, 0x56, 0xca, 0xa8, 0xf8, 0x77, 0xae, 0x54, 0xbe, 0xe3, 0x09, 0x2b, 0x45,
	0xc2, 0xa4, 0x54, 0x86, 0x19, 0xa1, 0xa4, 0x7b, 0x4a, 0xfe, 0x03, 0x6e, 0xb8, 0xb9, 0x10, 0xda,
	0x30, 0xb9, 0xe5, 0x29, 0x7f, 0xd8, 0x73, 0x6d, 0xf0, 0x17, 0x84, 0xaa, 0x12, 0xb9, 0x90, 0x91,
	0x37, 0xf3, 0xe6, 0x83, 0xd4, 0x29, 0x44, 0xe8, 0x65, 0x5c, 0x9b, 0xc8, 0xb7, 0xae, 0x9d, 0xc9,
	0x02, 0xa6, 0x27, 0x09, 0xba, 0x54, 0x52, 0x73, 0x8c, 0xa1, 0x9f, 0x39, 0xcf, 0x86, 0x04, 0x69,
	0xa7, 0x09, 0x85, 0xfe, 0x95, 0xda, 0xda, 0x1e, 0x38, 0x86, 0x60, 0xa7, 0x9a, 0x3d, 0x5e, 0x5a,
	0x8f, 0xd6, 0x61, 0xcd, 0x8e, 0xda, 0x61, 0x86, 0xfc, 0x80, 0xef, 0x97, 0x45, 0x69, 0x1e, 0xdb,
	0x70, 0xb2, 0x86, 0xf1, 0x86, 0x9b, 0x5b, 0x2e, 0xf7, 0x5c, 0xb7, 0x9d, 0x63, 0xe8, 0xef, 0x5c,
	0xa8, 0x6b, 0xdd, 0xe9, 0xfa, 0x9e, 0x8a, 0x65, 0x62, 0xaf, 0x6d, 0x6a, 0x90, 0x3a, 0x45, 0x56,
	0x30, 0x39, 0xca, 0x71, 0xcd, 0xff, 0x40, 0x78, 0xb0, 0x4e, 0xe4, 0xcd, 0x82, 0xf9, 0x70, 0x19,
	0x52, 0xfb, 0x42, 0xea, 0x5c, 0x72, 0x07, 0x5f, 0xac, 0x81, 0x23, 0xf0, 0x45, 0xe6, 0x76, 0xf9,
	0x22, 0xab, 0xe9, 0x48, 0x56, 0xf0, 0x96, 0x4e, 0x3d, 0x63, 0x04, 0x5f, 0x59, 0x96, 0x55, 0x5c,
	0xeb, 0x28, 0xb0, 0x76, 0x2b, 0x9b, 0x4e, 0x46, 0xc8, 0x3c, 0xea, 0xd9, 0x4b, 0x9d, 0x5a, 0xbe,
	0x78, 0xd0, 0xbb, 0x66, 0xa5, 0xc6, 0x7f, 0x30, 0x3c, 0x02, 0x8b, 0x53, 0x7a, 0xfe, 0xa1, 0xe2,
	0x9f, 0xf4, 0x33, 0xf6, 0x7f, 0x61, 0x74, 0x63, 0x2a, 0xce, 0x8a, 0x8e, 0xf2, 0x80, 0xb6, 0x63,
	0x3c, 0xa2, 0x27, 0x2c, 0xe7, 0x1e, 0xae, 0x61, 0xd0, 0x51, 0xc0, 0x09, 0xfd, 0x48, 0x36, 0x46,
	0x7a, 0x06, 0x89, 0xe0, 0xd3, 0xeb, 0xdb, 0xb3, 0xff, 0x0d, 0x21, 0x39, 0x2c, 0x92, 0x06, 0xcc,
	0x7d, 0x68, 0x7f, 0xa9, 0xd5, 0x7b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x61, 0xb2, 0x7f, 0xf9, 0x7e,
	0x02, 0x00, 0x00,
}
