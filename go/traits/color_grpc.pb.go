// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: traits/color.proto

package traits

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ColorApi_GetColor_FullMethodName    = "/smartcore.traits.ColorApi/GetColor"
	ColorApi_UpdateColor_FullMethodName = "/smartcore.traits.ColorApi/UpdateColor"
	ColorApi_PullColor_FullMethodName   = "/smartcore.traits.ColorApi/PullColor"
)

// ColorApiClient is the client API for ColorApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Control the color of a device. This could be a smart-bulb, or a DMX fitting.
type ColorApiClient interface {
	// get the current color
	GetColor(ctx context.Context, in *GetColorRequest, opts ...grpc.CallOption) (*Color, error)
	// request that the color be changed
	UpdateColor(ctx context.Context, in *UpdateColorRequest, opts ...grpc.CallOption) (*Color, error)
	// request updates to changes to the color value
	PullColor(ctx context.Context, in *PullColorRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PullColorResponse], error)
}

type colorApiClient struct {
	cc grpc.ClientConnInterface
}

func NewColorApiClient(cc grpc.ClientConnInterface) ColorApiClient {
	return &colorApiClient{cc}
}

func (c *colorApiClient) GetColor(ctx context.Context, in *GetColorRequest, opts ...grpc.CallOption) (*Color, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Color)
	err := c.cc.Invoke(ctx, ColorApi_GetColor_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *colorApiClient) UpdateColor(ctx context.Context, in *UpdateColorRequest, opts ...grpc.CallOption) (*Color, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Color)
	err := c.cc.Invoke(ctx, ColorApi_UpdateColor_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *colorApiClient) PullColor(ctx context.Context, in *PullColorRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PullColorResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ColorApi_ServiceDesc.Streams[0], ColorApi_PullColor_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[PullColorRequest, PullColorResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ColorApi_PullColorClient = grpc.ServerStreamingClient[PullColorResponse]

// ColorApiServer is the server API for ColorApi service.
// All implementations must embed UnimplementedColorApiServer
// for forward compatibility.
//
// Control the color of a device. This could be a smart-bulb, or a DMX fitting.
type ColorApiServer interface {
	// get the current color
	GetColor(context.Context, *GetColorRequest) (*Color, error)
	// request that the color be changed
	UpdateColor(context.Context, *UpdateColorRequest) (*Color, error)
	// request updates to changes to the color value
	PullColor(*PullColorRequest, grpc.ServerStreamingServer[PullColorResponse]) error
	mustEmbedUnimplementedColorApiServer()
}

// UnimplementedColorApiServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedColorApiServer struct{}

func (UnimplementedColorApiServer) GetColor(context.Context, *GetColorRequest) (*Color, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetColor not implemented")
}
func (UnimplementedColorApiServer) UpdateColor(context.Context, *UpdateColorRequest) (*Color, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateColor not implemented")
}
func (UnimplementedColorApiServer) PullColor(*PullColorRequest, grpc.ServerStreamingServer[PullColorResponse]) error {
	return status.Errorf(codes.Unimplemented, "method PullColor not implemented")
}
func (UnimplementedColorApiServer) mustEmbedUnimplementedColorApiServer() {}
func (UnimplementedColorApiServer) testEmbeddedByValue()                  {}

// UnsafeColorApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ColorApiServer will
// result in compilation errors.
type UnsafeColorApiServer interface {
	mustEmbedUnimplementedColorApiServer()
}

func RegisterColorApiServer(s grpc.ServiceRegistrar, srv ColorApiServer) {
	// If the following call pancis, it indicates UnimplementedColorApiServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ColorApi_ServiceDesc, srv)
}

func _ColorApi_GetColor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetColorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ColorApiServer).GetColor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ColorApi_GetColor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ColorApiServer).GetColor(ctx, req.(*GetColorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ColorApi_UpdateColor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateColorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ColorApiServer).UpdateColor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ColorApi_UpdateColor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ColorApiServer).UpdateColor(ctx, req.(*UpdateColorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ColorApi_PullColor_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullColorRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ColorApiServer).PullColor(m, &grpc.GenericServerStream[PullColorRequest, PullColorResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ColorApi_PullColorServer = grpc.ServerStreamingServer[PullColorResponse]

// ColorApi_ServiceDesc is the grpc.ServiceDesc for ColorApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ColorApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.traits.ColorApi",
	HandlerType: (*ColorApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetColor",
			Handler:    _ColorApi_GetColor_Handler,
		},
		{
			MethodName: "UpdateColor",
			Handler:    _ColorApi_UpdateColor_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PullColor",
			Handler:       _ColorApi_PullColor_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "traits/color.proto",
}

const (
	ColorInfo_DescribeColor_FullMethodName = "/smartcore.traits.ColorInfo/DescribeColor"
)

// ColorInfoClient is the client API for ColorInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Describes the capabilities of a specific named device with respect to this trait.
type ColorInfoClient interface {
	// Get information about how a named device implements Color features
	DescribeColor(ctx context.Context, in *DescribeColorRequest, opts ...grpc.CallOption) (*ColorSupport, error)
}

type colorInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewColorInfoClient(cc grpc.ClientConnInterface) ColorInfoClient {
	return &colorInfoClient{cc}
}

func (c *colorInfoClient) DescribeColor(ctx context.Context, in *DescribeColorRequest, opts ...grpc.CallOption) (*ColorSupport, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ColorSupport)
	err := c.cc.Invoke(ctx, ColorInfo_DescribeColor_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ColorInfoServer is the server API for ColorInfo service.
// All implementations must embed UnimplementedColorInfoServer
// for forward compatibility.
//
// Describes the capabilities of a specific named device with respect to this trait.
type ColorInfoServer interface {
	// Get information about how a named device implements Color features
	DescribeColor(context.Context, *DescribeColorRequest) (*ColorSupport, error)
	mustEmbedUnimplementedColorInfoServer()
}

// UnimplementedColorInfoServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedColorInfoServer struct{}

func (UnimplementedColorInfoServer) DescribeColor(context.Context, *DescribeColorRequest) (*ColorSupport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeColor not implemented")
}
func (UnimplementedColorInfoServer) mustEmbedUnimplementedColorInfoServer() {}
func (UnimplementedColorInfoServer) testEmbeddedByValue()                   {}

// UnsafeColorInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ColorInfoServer will
// result in compilation errors.
type UnsafeColorInfoServer interface {
	mustEmbedUnimplementedColorInfoServer()
}

func RegisterColorInfoServer(s grpc.ServiceRegistrar, srv ColorInfoServer) {
	// If the following call pancis, it indicates UnimplementedColorInfoServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ColorInfo_ServiceDesc, srv)
}

func _ColorInfo_DescribeColor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeColorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ColorInfoServer).DescribeColor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ColorInfo_DescribeColor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ColorInfoServer).DescribeColor(ctx, req.(*DescribeColorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ColorInfo_ServiceDesc is the grpc.ServiceDesc for ColorInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ColorInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.traits.ColorInfo",
	HandlerType: (*ColorInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeColor",
			Handler:    _ColorInfo_DescribeColor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "traits/color.proto",
}
