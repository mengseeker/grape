// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package confd

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ConfdServerClient is the client API for ConfdServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfdServerClient interface {
	StreamDiscovery(ctx context.Context, in *Discovery, opts ...grpc.CallOption) (ConfdServer_StreamDiscoveryClient, error)
}

type confdServerClient struct {
	cc grpc.ClientConnInterface
}

func NewConfdServerClient(cc grpc.ClientConnInterface) ConfdServerClient {
	return &confdServerClient{cc}
}

func (c *confdServerClient) StreamDiscovery(ctx context.Context, in *Discovery, opts ...grpc.CallOption) (ConfdServer_StreamDiscoveryClient, error) {
	stream, err := c.cc.NewStream(ctx, &ConfdServer_ServiceDesc.Streams[0], "/api.v1.confd.ConfdServer/StreamDiscovery", opts...)
	if err != nil {
		return nil, err
	}
	x := &confdServerStreamDiscoveryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ConfdServer_StreamDiscoveryClient interface {
	Recv() (*Configs, error)
	grpc.ClientStream
}

type confdServerStreamDiscoveryClient struct {
	grpc.ClientStream
}

func (x *confdServerStreamDiscoveryClient) Recv() (*Configs, error) {
	m := new(Configs)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ConfdServerServer is the server API for ConfdServer service.
// All implementations must embed UnimplementedConfdServerServer
// for forward compatibility
type ConfdServerServer interface {
	StreamDiscovery(*Discovery, ConfdServer_StreamDiscoveryServer) error
	mustEmbedUnimplementedConfdServerServer()
}

// UnimplementedConfdServerServer must be embedded to have forward compatible implementations.
type UnimplementedConfdServerServer struct {
}

func (UnimplementedConfdServerServer) StreamDiscovery(*Discovery, ConfdServer_StreamDiscoveryServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamDiscovery not implemented")
}
func (UnimplementedConfdServerServer) mustEmbedUnimplementedConfdServerServer() {}

// UnsafeConfdServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfdServerServer will
// result in compilation errors.
type UnsafeConfdServerServer interface {
	mustEmbedUnimplementedConfdServerServer()
}

func RegisterConfdServerServer(s grpc.ServiceRegistrar, srv ConfdServerServer) {
	s.RegisterService(&ConfdServer_ServiceDesc, srv)
}

func _ConfdServer_StreamDiscovery_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Discovery)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConfdServerServer).StreamDiscovery(m, &confdServerStreamDiscoveryServer{stream})
}

type ConfdServer_StreamDiscoveryServer interface {
	Send(*Configs) error
	grpc.ServerStream
}

type confdServerStreamDiscoveryServer struct {
	grpc.ServerStream
}

func (x *confdServerStreamDiscoveryServer) Send(m *Configs) error {
	return x.ServerStream.SendMsg(m)
}

// ConfdServer_ServiceDesc is the grpc.ServiceDesc for ConfdServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConfdServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.confd.ConfdServer",
	HandlerType: (*ConfdServerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamDiscovery",
			Handler:       _ConfdServer_StreamDiscovery_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/v1/confd/discovery.proto",
}
