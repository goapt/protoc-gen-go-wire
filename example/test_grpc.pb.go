// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package example

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

// DummyClient is the client API for Dummy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DummyClient interface {
	Run(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error)
}

type dummyClient struct {
	cc grpc.ClientConnInterface
}

func NewDummyClient(cc grpc.ClientConnInterface) DummyClient {
	return &dummyClient{cc}
}

func (c *dummyClient) Run(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error) {
	out := new(TestResponse)
	err := c.cc.Invoke(ctx, "/Dummy/Run", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DummyServer is the server API for Dummy service.
// All implementations should embed UnimplementedDummyServer
// for forward compatibility
type DummyServer interface {
	Run(context.Context, *TestRequest) (*TestResponse, error)
}

// UnimplementedDummyServer should be embedded to have forward compatible implementations.
type UnimplementedDummyServer struct {
}

func (UnimplementedDummyServer) Run(context.Context, *TestRequest) (*TestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Run not implemented")
}

// UnsafeDummyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DummyServer will
// result in compilation errors.
type UnsafeDummyServer interface {
	mustEmbedUnimplementedDummyServer()
}

func RegisterDummyServer(s grpc.ServiceRegistrar, srv DummyServer) {
	s.RegisterService(&Dummy_ServiceDesc, srv)
}

func _Dummy_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DummyServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Dummy/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DummyServer).Run(ctx, req.(*TestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Dummy_ServiceDesc is the grpc.ServiceDesc for Dummy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Dummy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Dummy",
	HandlerType: (*DummyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Run",
			Handler:    _Dummy_Run_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/test.proto",
}