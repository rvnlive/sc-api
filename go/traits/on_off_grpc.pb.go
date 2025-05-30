// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: traits/on_off.proto

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
	OnOffApi_GetOnOff_FullMethodName    = "/smartcore.traits.OnOffApi/GetOnOff"
	OnOffApi_UpdateOnOff_FullMethodName = "/smartcore.traits.OnOffApi/UpdateOnOff"
	OnOffApi_PullOnOff_FullMethodName   = "/smartcore.traits.OnOffApi/PullOnOff"
)

// OnOffApiClient is the client API for OnOffApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Applicable to devices that have a basic binary on and off functionality.
type OnOffApiClient interface {
	// Get the current state of the device
	GetOnOff(ctx context.Context, in *GetOnOffRequest, opts ...grpc.CallOption) (*OnOff, error)
	// Update the device to be on or off
	UpdateOnOff(ctx context.Context, in *UpdateOnOffRequest, opts ...grpc.CallOption) (*OnOff, error)
	// Get notified of changes to the OnOff state of a device
	PullOnOff(ctx context.Context, in *PullOnOffRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PullOnOffResponse], error)
}

type onOffApiClient struct {
	cc grpc.ClientConnInterface
}

func NewOnOffApiClient(cc grpc.ClientConnInterface) OnOffApiClient {
	return &onOffApiClient{cc}
}

func (c *onOffApiClient) GetOnOff(ctx context.Context, in *GetOnOffRequest, opts ...grpc.CallOption) (*OnOff, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OnOff)
	err := c.cc.Invoke(ctx, OnOffApi_GetOnOff_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onOffApiClient) UpdateOnOff(ctx context.Context, in *UpdateOnOffRequest, opts ...grpc.CallOption) (*OnOff, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OnOff)
	err := c.cc.Invoke(ctx, OnOffApi_UpdateOnOff_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onOffApiClient) PullOnOff(ctx context.Context, in *PullOnOffRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PullOnOffResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &OnOffApi_ServiceDesc.Streams[0], OnOffApi_PullOnOff_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[PullOnOffRequest, PullOnOffResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type OnOffApi_PullOnOffClient = grpc.ServerStreamingClient[PullOnOffResponse]

// OnOffApiServer is the server API for OnOffApi service.
// All implementations must embed UnimplementedOnOffApiServer
// for forward compatibility.
//
// Applicable to devices that have a basic binary on and off functionality.
type OnOffApiServer interface {
	// Get the current state of the device
	GetOnOff(context.Context, *GetOnOffRequest) (*OnOff, error)
	// Update the device to be on or off
	UpdateOnOff(context.Context, *UpdateOnOffRequest) (*OnOff, error)
	// Get notified of changes to the OnOff state of a device
	PullOnOff(*PullOnOffRequest, grpc.ServerStreamingServer[PullOnOffResponse]) error
	mustEmbedUnimplementedOnOffApiServer()
}

// UnimplementedOnOffApiServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOnOffApiServer struct{}

func (UnimplementedOnOffApiServer) GetOnOff(context.Context, *GetOnOffRequest) (*OnOff, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOnOff not implemented")
}
func (UnimplementedOnOffApiServer) UpdateOnOff(context.Context, *UpdateOnOffRequest) (*OnOff, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOnOff not implemented")
}
func (UnimplementedOnOffApiServer) PullOnOff(*PullOnOffRequest, grpc.ServerStreamingServer[PullOnOffResponse]) error {
	return status.Errorf(codes.Unimplemented, "method PullOnOff not implemented")
}
func (UnimplementedOnOffApiServer) mustEmbedUnimplementedOnOffApiServer() {}
func (UnimplementedOnOffApiServer) testEmbeddedByValue()                  {}

// UnsafeOnOffApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OnOffApiServer will
// result in compilation errors.
type UnsafeOnOffApiServer interface {
	mustEmbedUnimplementedOnOffApiServer()
}

func RegisterOnOffApiServer(s grpc.ServiceRegistrar, srv OnOffApiServer) {
	// If the following call pancis, it indicates UnimplementedOnOffApiServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OnOffApi_ServiceDesc, srv)
}

func _OnOffApi_GetOnOff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOnOffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnOffApiServer).GetOnOff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OnOffApi_GetOnOff_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnOffApiServer).GetOnOff(ctx, req.(*GetOnOffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnOffApi_UpdateOnOff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOnOffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnOffApiServer).UpdateOnOff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OnOffApi_UpdateOnOff_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnOffApiServer).UpdateOnOff(ctx, req.(*UpdateOnOffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnOffApi_PullOnOff_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullOnOffRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OnOffApiServer).PullOnOff(m, &grpc.GenericServerStream[PullOnOffRequest, PullOnOffResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type OnOffApi_PullOnOffServer = grpc.ServerStreamingServer[PullOnOffResponse]

// OnOffApi_ServiceDesc is the grpc.ServiceDesc for OnOffApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OnOffApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.traits.OnOffApi",
	HandlerType: (*OnOffApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOnOff",
			Handler:    _OnOffApi_GetOnOff_Handler,
		},
		{
			MethodName: "UpdateOnOff",
			Handler:    _OnOffApi_UpdateOnOff_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PullOnOff",
			Handler:       _OnOffApi_PullOnOff_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "traits/on_off.proto",
}

const (
	OnOffInfo_DescribeOnOff_FullMethodName = "/smartcore.traits.OnOffInfo/DescribeOnOff"
)

// OnOffInfoClient is the client API for OnOffInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Describes the capabilities of a specific named device with respect to this trait.
type OnOffInfoClient interface {
	// Get information about how a named device implements OnOff features
	DescribeOnOff(ctx context.Context, in *DescribeOnOffRequest, opts ...grpc.CallOption) (*OnOffSupport, error)
}

type onOffInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewOnOffInfoClient(cc grpc.ClientConnInterface) OnOffInfoClient {
	return &onOffInfoClient{cc}
}

func (c *onOffInfoClient) DescribeOnOff(ctx context.Context, in *DescribeOnOffRequest, opts ...grpc.CallOption) (*OnOffSupport, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OnOffSupport)
	err := c.cc.Invoke(ctx, OnOffInfo_DescribeOnOff_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OnOffInfoServer is the server API for OnOffInfo service.
// All implementations must embed UnimplementedOnOffInfoServer
// for forward compatibility.
//
// Describes the capabilities of a specific named device with respect to this trait.
type OnOffInfoServer interface {
	// Get information about how a named device implements OnOff features
	DescribeOnOff(context.Context, *DescribeOnOffRequest) (*OnOffSupport, error)
	mustEmbedUnimplementedOnOffInfoServer()
}

// UnimplementedOnOffInfoServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOnOffInfoServer struct{}

func (UnimplementedOnOffInfoServer) DescribeOnOff(context.Context, *DescribeOnOffRequest) (*OnOffSupport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeOnOff not implemented")
}
func (UnimplementedOnOffInfoServer) mustEmbedUnimplementedOnOffInfoServer() {}
func (UnimplementedOnOffInfoServer) testEmbeddedByValue()                   {}

// UnsafeOnOffInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OnOffInfoServer will
// result in compilation errors.
type UnsafeOnOffInfoServer interface {
	mustEmbedUnimplementedOnOffInfoServer()
}

func RegisterOnOffInfoServer(s grpc.ServiceRegistrar, srv OnOffInfoServer) {
	// If the following call pancis, it indicates UnimplementedOnOffInfoServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OnOffInfo_ServiceDesc, srv)
}

func _OnOffInfo_DescribeOnOff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeOnOffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnOffInfoServer).DescribeOnOff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OnOffInfo_DescribeOnOff_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnOffInfoServer).DescribeOnOff(ctx, req.(*DescribeOnOffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OnOffInfo_ServiceDesc is the grpc.ServiceDesc for OnOffInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OnOffInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.traits.OnOffInfo",
	HandlerType: (*OnOffInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeOnOff",
			Handler:    _OnOffInfo_DescribeOnOff_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "traits/on_off.proto",
}
