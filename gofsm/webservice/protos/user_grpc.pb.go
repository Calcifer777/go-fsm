// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: protos/user.proto

package protos

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

// TutorialClient is the client API for Tutorial service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TutorialClient interface {
	CreateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*User, error)
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
}

type tutorialClient struct {
	cc grpc.ClientConnInterface
}

func NewTutorialClient(cc grpc.ClientConnInterface) TutorialClient {
	return &tutorialClient{cc}
}

func (c *tutorialClient) CreateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/Tutorial/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tutorialClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := c.cc.Invoke(ctx, "/Tutorial/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TutorialServer is the server API for Tutorial service.
// All implementations must embed UnimplementedTutorialServer
// for forward compatibility
type TutorialServer interface {
	CreateUser(context.Context, *UserRequest) (*User, error)
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
	mustEmbedUnimplementedTutorialServer()
}

// UnimplementedTutorialServer must be embedded to have forward compatible implementations.
type UnimplementedTutorialServer struct {
}

func (UnimplementedTutorialServer) CreateUser(context.Context, *UserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedTutorialServer) Greet(context.Context, *GreetRequest) (*GreetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}
func (UnimplementedTutorialServer) mustEmbedUnimplementedTutorialServer() {}

// UnsafeTutorialServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TutorialServer will
// result in compilation errors.
type UnsafeTutorialServer interface {
	mustEmbedUnimplementedTutorialServer()
}

func RegisterTutorialServer(s grpc.ServiceRegistrar, srv TutorialServer) {
	s.RegisterService(&Tutorial_ServiceDesc, srv)
}

func _Tutorial_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TutorialServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Tutorial/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TutorialServer).CreateUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tutorial_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TutorialServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Tutorial/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TutorialServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Tutorial_ServiceDesc is the grpc.ServiceDesc for Tutorial service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tutorial_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Tutorial",
	HandlerType: (*TutorialServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _Tutorial_CreateUser_Handler,
		},
		{
			MethodName: "Greet",
			Handler:    _Tutorial_Greet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/user.proto",
}
