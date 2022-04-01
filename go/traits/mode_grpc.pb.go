// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: traits/mode.proto

package traits

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

// ModeApiClient is the client API for ModeApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModeApiClient interface {
	// GetModeValues returns the current state for all modes for a device.
	GetModeValues(ctx context.Context, in *GetModeValuesRequest, opts ...grpc.CallOption) (*ModeValues, error)
	// UpdateModeValues sets the current mode for the device.
	// Multiple modes may be set at the same time.
	// The device may error if setting a mode that is not supported, in any case the returned ModeValues contains the
	// state after the update.
	// Relative updates to the mode may either cap or wrap the mode value if adjusting the value relatively would either
	// underflow or overflow the available modes.
	UpdateModeValues(ctx context.Context, in *UpdateModeValuesRequest, opts ...grpc.CallOption) (*ModeValues, error)
	// PullModeValues returns a stream of updates to the modes set on the device.
	PullModeValues(ctx context.Context, in *PullModeValuesRequest, opts ...grpc.CallOption) (ModeApi_PullModeValuesClient, error)
}

type modeApiClient struct {
	cc grpc.ClientConnInterface
}

func NewModeApiClient(cc grpc.ClientConnInterface) ModeApiClient {
	return &modeApiClient{cc}
}

func (c *modeApiClient) GetModeValues(ctx context.Context, in *GetModeValuesRequest, opts ...grpc.CallOption) (*ModeValues, error) {
	out := new(ModeValues)
	err := c.cc.Invoke(ctx, "/smartcore.traits.ModeApi/GetModeValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modeApiClient) UpdateModeValues(ctx context.Context, in *UpdateModeValuesRequest, opts ...grpc.CallOption) (*ModeValues, error) {
	out := new(ModeValues)
	err := c.cc.Invoke(ctx, "/smartcore.traits.ModeApi/UpdateModeValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modeApiClient) PullModeValues(ctx context.Context, in *PullModeValuesRequest, opts ...grpc.CallOption) (ModeApi_PullModeValuesClient, error) {
	stream, err := c.cc.NewStream(ctx, &ModeApi_ServiceDesc.Streams[0], "/smartcore.traits.ModeApi/PullModeValues", opts...)
	if err != nil {
		return nil, err
	}
	x := &modeApiPullModeValuesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ModeApi_PullModeValuesClient interface {
	Recv() (*PullModeValuesResponse, error)
	grpc.ClientStream
}

type modeApiPullModeValuesClient struct {
	grpc.ClientStream
}

func (x *modeApiPullModeValuesClient) Recv() (*PullModeValuesResponse, error) {
	m := new(PullModeValuesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ModeApiServer is the server API for ModeApi service.
// All implementations must embed UnimplementedModeApiServer
// for forward compatibility
type ModeApiServer interface {
	// GetModeValues returns the current state for all modes for a device.
	GetModeValues(context.Context, *GetModeValuesRequest) (*ModeValues, error)
	// UpdateModeValues sets the current mode for the device.
	// Multiple modes may be set at the same time.
	// The device may error if setting a mode that is not supported, in any case the returned ModeValues contains the
	// state after the update.
	// Relative updates to the mode may either cap or wrap the mode value if adjusting the value relatively would either
	// underflow or overflow the available modes.
	UpdateModeValues(context.Context, *UpdateModeValuesRequest) (*ModeValues, error)
	// PullModeValues returns a stream of updates to the modes set on the device.
	PullModeValues(*PullModeValuesRequest, ModeApi_PullModeValuesServer) error
	mustEmbedUnimplementedModeApiServer()
}

// UnimplementedModeApiServer must be embedded to have forward compatible implementations.
type UnimplementedModeApiServer struct {
}

func (UnimplementedModeApiServer) GetModeValues(context.Context, *GetModeValuesRequest) (*ModeValues, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModeValues not implemented")
}
func (UnimplementedModeApiServer) UpdateModeValues(context.Context, *UpdateModeValuesRequest) (*ModeValues, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateModeValues not implemented")
}
func (UnimplementedModeApiServer) PullModeValues(*PullModeValuesRequest, ModeApi_PullModeValuesServer) error {
	return status.Errorf(codes.Unimplemented, "method PullModeValues not implemented")
}
func (UnimplementedModeApiServer) mustEmbedUnimplementedModeApiServer() {}

// UnsafeModeApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModeApiServer will
// result in compilation errors.
type UnsafeModeApiServer interface {
	mustEmbedUnimplementedModeApiServer()
}

func RegisterModeApiServer(s grpc.ServiceRegistrar, srv ModeApiServer) {
	s.RegisterService(&ModeApi_ServiceDesc, srv)
}

func _ModeApi_GetModeValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModeValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModeApiServer).GetModeValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.traits.ModeApi/GetModeValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModeApiServer).GetModeValues(ctx, req.(*GetModeValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModeApi_UpdateModeValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateModeValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModeApiServer).UpdateModeValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.traits.ModeApi/UpdateModeValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModeApiServer).UpdateModeValues(ctx, req.(*UpdateModeValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModeApi_PullModeValues_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullModeValuesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ModeApiServer).PullModeValues(m, &modeApiPullModeValuesServer{stream})
}

type ModeApi_PullModeValuesServer interface {
	Send(*PullModeValuesResponse) error
	grpc.ServerStream
}

type modeApiPullModeValuesServer struct {
	grpc.ServerStream
}

func (x *modeApiPullModeValuesServer) Send(m *PullModeValuesResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ModeApi_ServiceDesc is the grpc.ServiceDesc for ModeApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ModeApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.traits.ModeApi",
	HandlerType: (*ModeApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetModeValues",
			Handler:    _ModeApi_GetModeValues_Handler,
		},
		{
			MethodName: "UpdateModeValues",
			Handler:    _ModeApi_UpdateModeValues_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PullModeValues",
			Handler:       _ModeApi_PullModeValues_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "traits/mode.proto",
}

// ModeInfoClient is the client API for ModeInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModeInfoClient interface {
	// Get information about how a named device implements Modes features
	DescribeModes(ctx context.Context, in *DescribeModesRequest, opts ...grpc.CallOption) (*ModesSupport, error)
}

type modeInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewModeInfoClient(cc grpc.ClientConnInterface) ModeInfoClient {
	return &modeInfoClient{cc}
}

func (c *modeInfoClient) DescribeModes(ctx context.Context, in *DescribeModesRequest, opts ...grpc.CallOption) (*ModesSupport, error) {
	out := new(ModesSupport)
	err := c.cc.Invoke(ctx, "/smartcore.traits.ModeInfo/DescribeModes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModeInfoServer is the server API for ModeInfo service.
// All implementations must embed UnimplementedModeInfoServer
// for forward compatibility
type ModeInfoServer interface {
	// Get information about how a named device implements Modes features
	DescribeModes(context.Context, *DescribeModesRequest) (*ModesSupport, error)
	mustEmbedUnimplementedModeInfoServer()
}

// UnimplementedModeInfoServer must be embedded to have forward compatible implementations.
type UnimplementedModeInfoServer struct {
}

func (UnimplementedModeInfoServer) DescribeModes(context.Context, *DescribeModesRequest) (*ModesSupport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeModes not implemented")
}
func (UnimplementedModeInfoServer) mustEmbedUnimplementedModeInfoServer() {}

// UnsafeModeInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModeInfoServer will
// result in compilation errors.
type UnsafeModeInfoServer interface {
	mustEmbedUnimplementedModeInfoServer()
}

func RegisterModeInfoServer(s grpc.ServiceRegistrar, srv ModeInfoServer) {
	s.RegisterService(&ModeInfo_ServiceDesc, srv)
}

func _ModeInfo_DescribeModes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeModesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModeInfoServer).DescribeModes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.traits.ModeInfo/DescribeModes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModeInfoServer).DescribeModes(ctx, req.(*DescribeModesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ModeInfo_ServiceDesc is the grpc.ServiceDesc for ModeInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ModeInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.traits.ModeInfo",
	HandlerType: (*ModeInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeModes",
			Handler:    _ModeInfo_DescribeModes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "traits/mode.proto",
}
