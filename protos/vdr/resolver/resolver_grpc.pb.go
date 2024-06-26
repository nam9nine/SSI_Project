// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: protos/vdr/resolver/resolver.proto

package resolver

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

// DIDResolverClient is the client API for DIDResolver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DIDResolverClient interface {
	ResolveDID(ctx context.Context, in *ResolveDIDReq, opts ...grpc.CallOption) (*ResolveDIDRes, error)
}

type dIDResolverClient struct {
	cc grpc.ClientConnInterface
}

func NewDIDResolverClient(cc grpc.ClientConnInterface) DIDResolverClient {
	return &dIDResolverClient{cc}
}

func (c *dIDResolverClient) ResolveDID(ctx context.Context, in *ResolveDIDReq, opts ...grpc.CallOption) (*ResolveDIDRes, error) {
	out := new(ResolveDIDRes)
	err := c.cc.Invoke(ctx, "/vdr.resolver.DIDResolver/ResolveDID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DIDResolverServer is the server API for DIDResolver service.
// All implementations must embed UnimplementedDIDResolverServer
// for forward compatibility
type DIDResolverServer interface {
	ResolveDID(context.Context, *ResolveDIDReq) (*ResolveDIDRes, error)
	mustEmbedUnimplementedDIDResolverServer()
}

// UnimplementedDIDResolverServer must be embedded to have forward compatible implementations.
type UnimplementedDIDResolverServer struct {
}

func (UnimplementedDIDResolverServer) ResolveDID(context.Context, *ResolveDIDReq) (*ResolveDIDRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResolveDID not implemented")
}
func (UnimplementedDIDResolverServer) mustEmbedUnimplementedDIDResolverServer() {}

// UnsafeDIDResolverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DIDResolverServer will
// result in compilation errors.
type UnsafeDIDResolverServer interface {
	mustEmbedUnimplementedDIDResolverServer()
}

func RegisterDIDResolverServer(s grpc.ServiceRegistrar, srv DIDResolverServer) {
	s.RegisterService(&DIDResolver_ServiceDesc, srv)
}

func _DIDResolver_ResolveDID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveDIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DIDResolverServer).ResolveDID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vdr.resolver.DIDResolver/ResolveDID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DIDResolverServer).ResolveDID(ctx, req.(*ResolveDIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

// DIDResolver_ServiceDesc is the grpc.ServiceDesc for DIDResolver service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DIDResolver_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vdr.resolver.DIDResolver",
	HandlerType: (*DIDResolverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ResolveDID",
			Handler:    _DIDResolver_ResolveDID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/vdr/resolver/resolver.proto",
}
