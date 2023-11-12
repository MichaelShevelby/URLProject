// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.0
// source: api/generateshorturl/generateshorturl.proto

package generateshorturl

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

// GenerateShortUrlClient is the client API for GenerateShortUrl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GenerateShortUrlClient interface {
	GenerateShortUrl(ctx context.Context, in *GenerateShortUrlRequest, opts ...grpc.CallOption) (*GenerateShortUrlResponse, error)
}

type generateShortUrlClient struct {
	cc grpc.ClientConnInterface
}

func NewGenerateShortUrlClient(cc grpc.ClientConnInterface) GenerateShortUrlClient {
	return &generateShortUrlClient{cc}
}

func (c *generateShortUrlClient) GenerateShortUrl(ctx context.Context, in *GenerateShortUrlRequest, opts ...grpc.CallOption) (*GenerateShortUrlResponse, error) {
	out := new(GenerateShortUrlResponse)
	err := c.cc.Invoke(ctx, "/generateshorturl.GenerateShortUrl/GenerateShortUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GenerateShortUrlServer is the server API for GenerateShortUrl service.
// All implementations must embed UnimplementedGenerateShortUrlServer
// for forward compatibility
type GenerateShortUrlServer interface {
	GenerateShortUrl(context.Context, *GenerateShortUrlRequest) (*GenerateShortUrlResponse, error)
	mustEmbedUnimplementedGenerateShortUrlServer()
}

// UnimplementedGenerateShortUrlServer must be embedded to have forward compatible implementations.
type UnimplementedGenerateShortUrlServer struct {
}

func (UnimplementedGenerateShortUrlServer) GenerateShortUrl(context.Context, *GenerateShortUrlRequest) (*GenerateShortUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateShortUrl not implemented")
}
func (UnimplementedGenerateShortUrlServer) mustEmbedUnimplementedGenerateShortUrlServer() {}

// UnsafeGenerateShortUrlServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GenerateShortUrlServer will
// result in compilation errors.
type UnsafeGenerateShortUrlServer interface {
	mustEmbedUnimplementedGenerateShortUrlServer()
}

func RegisterGenerateShortUrlServer(s grpc.ServiceRegistrar, srv GenerateShortUrlServer) {
	s.RegisterService(&GenerateShortUrl_ServiceDesc, srv)
}

func _GenerateShortUrl_GenerateShortUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateShortUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenerateShortUrlServer).GenerateShortUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generateshorturl.GenerateShortUrl/GenerateShortUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenerateShortUrlServer).GenerateShortUrl(ctx, req.(*GenerateShortUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GenerateShortUrl_ServiceDesc is the grpc.ServiceDesc for GenerateShortUrl service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GenerateShortUrl_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "generateshorturl.GenerateShortUrl",
	HandlerType: (*GenerateShortUrlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateShortUrl",
			Handler:    _GenerateShortUrl_GenerateShortUrl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/generateshorturl/generateshorturl.proto",
}
