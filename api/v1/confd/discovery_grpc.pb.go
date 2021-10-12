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
	StreamResources(ctx context.Context, in *Discovery, opts ...grpc.CallOption) (ConfdServer_StreamResourcesClient, error)
	Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (*DownloadResponse, error)
}

type confdServerClient struct {
	cc grpc.ClientConnInterface
}

func NewConfdServerClient(cc grpc.ClientConnInterface) ConfdServerClient {
	return &confdServerClient{cc}
}

func (c *confdServerClient) StreamResources(ctx context.Context, in *Discovery, opts ...grpc.CallOption) (ConfdServer_StreamResourcesClient, error) {
	stream, err := c.cc.NewStream(ctx, &ConfdServer_ServiceDesc.Streams[0], "/grape.api.v1.confd.ConfdServer/StreamResources", opts...)
	if err != nil {
		return nil, err
	}
	x := &confdServerStreamResourcesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ConfdServer_StreamResourcesClient interface {
	Recv() (*Configs, error)
	grpc.ClientStream
}

type confdServerStreamResourcesClient struct {
	grpc.ClientStream
}

func (x *confdServerStreamResourcesClient) Recv() (*Configs, error) {
	m := new(Configs)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *confdServerClient) Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (*DownloadResponse, error) {
	out := new(DownloadResponse)
	err := c.cc.Invoke(ctx, "/grape.api.v1.confd.ConfdServer/Download", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfdServerServer is the server API for ConfdServer service.
// All implementations must embed UnimplementedConfdServerServer
// for forward compatibility
type ConfdServerServer interface {
	StreamResources(*Discovery, ConfdServer_StreamResourcesServer) error
	Download(context.Context, *DownloadRequest) (*DownloadResponse, error)
	mustEmbedUnimplementedConfdServerServer()
}

// UnimplementedConfdServerServer must be embedded to have forward compatible implementations.
type UnimplementedConfdServerServer struct {
}

func (UnimplementedConfdServerServer) StreamResources(*Discovery, ConfdServer_StreamResourcesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamResources not implemented")
}
func (UnimplementedConfdServerServer) Download(context.Context, *DownloadRequest) (*DownloadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Download not implemented")
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

func _ConfdServer_StreamResources_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Discovery)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConfdServerServer).StreamResources(m, &confdServerStreamResourcesServer{stream})
}

type ConfdServer_StreamResourcesServer interface {
	Send(*Configs) error
	grpc.ServerStream
}

type confdServerStreamResourcesServer struct {
	grpc.ServerStream
}

func (x *confdServerStreamResourcesServer) Send(m *Configs) error {
	return x.ServerStream.SendMsg(m)
}

func _ConfdServer_Download_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfdServerServer).Download(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grape.api.v1.confd.ConfdServer/Download",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfdServerServer).Download(ctx, req.(*DownloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConfdServer_ServiceDesc is the grpc.ServiceDesc for ConfdServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConfdServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grape.api.v1.confd.ConfdServer",
	HandlerType: (*ConfdServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Download",
			Handler:    _ConfdServer_Download_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamResources",
			Handler:       _ConfdServer_StreamResources_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grape/api/v1/confd/discovery.proto",
}