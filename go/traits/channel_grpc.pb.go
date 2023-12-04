// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: traits/channel.proto

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

const (
	ChannelApi_GetChosenChannel_FullMethodName  = "/smartcore.traits.ChannelApi/GetChosenChannel"
	ChannelApi_ChooseChannel_FullMethodName     = "/smartcore.traits.ChannelApi/ChooseChannel"
	ChannelApi_AdjustChannel_FullMethodName     = "/smartcore.traits.ChannelApi/AdjustChannel"
	ChannelApi_ReturnChannel_FullMethodName     = "/smartcore.traits.ChannelApi/ReturnChannel"
	ChannelApi_PullChosenChannel_FullMethodName = "/smartcore.traits.ChannelApi/PullChosenChannel"
)

// ChannelApiClient is the client API for ChannelApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChannelApiClient interface {
	// GetChosenChannel returns the currently chosen channel for the device.
	GetChosenChannel(ctx context.Context, in *GetChosenChannelRequest, opts ...grpc.CallOption) (*Channel, error)
	// ChooseChannel will set the currently chosen channel to the one specified. Returns the chosen channel
	ChooseChannel(ctx context.Context, in *ChooseChannelRequest, opts ...grpc.CallOption) (*Channel, error)
	// AdjustChannel will choose a channel based on the current channel, i.e. 'next channel'. Returns the chosen channel
	AdjustChannel(ctx context.Context, in *AdjustChannelRequest, opts ...grpc.CallOption) (*Channel, error)
	// ReturnChannel will choose the previously chosen channel. Returns the chosen channel
	ReturnChannel(ctx context.Context, in *ReturnChannelRequest, opts ...grpc.CallOption) (*Channel, error)
	// PullChosenChannel subscribes to chosen channel changes
	PullChosenChannel(ctx context.Context, in *PullChosenChannelRequest, opts ...grpc.CallOption) (ChannelApi_PullChosenChannelClient, error)
}

type channelApiClient struct {
	cc grpc.ClientConnInterface
}

func NewChannelApiClient(cc grpc.ClientConnInterface) ChannelApiClient {
	return &channelApiClient{cc}
}

func (c *channelApiClient) GetChosenChannel(ctx context.Context, in *GetChosenChannelRequest, opts ...grpc.CallOption) (*Channel, error) {
	out := new(Channel)
	err := c.cc.Invoke(ctx, ChannelApi_GetChosenChannel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelApiClient) ChooseChannel(ctx context.Context, in *ChooseChannelRequest, opts ...grpc.CallOption) (*Channel, error) {
	out := new(Channel)
	err := c.cc.Invoke(ctx, ChannelApi_ChooseChannel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelApiClient) AdjustChannel(ctx context.Context, in *AdjustChannelRequest, opts ...grpc.CallOption) (*Channel, error) {
	out := new(Channel)
	err := c.cc.Invoke(ctx, ChannelApi_AdjustChannel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelApiClient) ReturnChannel(ctx context.Context, in *ReturnChannelRequest, opts ...grpc.CallOption) (*Channel, error) {
	out := new(Channel)
	err := c.cc.Invoke(ctx, ChannelApi_ReturnChannel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelApiClient) PullChosenChannel(ctx context.Context, in *PullChosenChannelRequest, opts ...grpc.CallOption) (ChannelApi_PullChosenChannelClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChannelApi_ServiceDesc.Streams[0], ChannelApi_PullChosenChannel_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &channelApiPullChosenChannelClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChannelApi_PullChosenChannelClient interface {
	Recv() (*PullChosenChannelResponse, error)
	grpc.ClientStream
}

type channelApiPullChosenChannelClient struct {
	grpc.ClientStream
}

func (x *channelApiPullChosenChannelClient) Recv() (*PullChosenChannelResponse, error) {
	m := new(PullChosenChannelResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChannelApiServer is the server API for ChannelApi service.
// All implementations must embed UnimplementedChannelApiServer
// for forward compatibility
type ChannelApiServer interface {
	// GetChosenChannel returns the currently chosen channel for the device.
	GetChosenChannel(context.Context, *GetChosenChannelRequest) (*Channel, error)
	// ChooseChannel will set the currently chosen channel to the one specified. Returns the chosen channel
	ChooseChannel(context.Context, *ChooseChannelRequest) (*Channel, error)
	// AdjustChannel will choose a channel based on the current channel, i.e. 'next channel'. Returns the chosen channel
	AdjustChannel(context.Context, *AdjustChannelRequest) (*Channel, error)
	// ReturnChannel will choose the previously chosen channel. Returns the chosen channel
	ReturnChannel(context.Context, *ReturnChannelRequest) (*Channel, error)
	// PullChosenChannel subscribes to chosen channel changes
	PullChosenChannel(*PullChosenChannelRequest, ChannelApi_PullChosenChannelServer) error
	mustEmbedUnimplementedChannelApiServer()
}

// UnimplementedChannelApiServer must be embedded to have forward compatible implementations.
type UnimplementedChannelApiServer struct {
}

func (UnimplementedChannelApiServer) GetChosenChannel(context.Context, *GetChosenChannelRequest) (*Channel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChosenChannel not implemented")
}
func (UnimplementedChannelApiServer) ChooseChannel(context.Context, *ChooseChannelRequest) (*Channel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChooseChannel not implemented")
}
func (UnimplementedChannelApiServer) AdjustChannel(context.Context, *AdjustChannelRequest) (*Channel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdjustChannel not implemented")
}
func (UnimplementedChannelApiServer) ReturnChannel(context.Context, *ReturnChannelRequest) (*Channel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReturnChannel not implemented")
}
func (UnimplementedChannelApiServer) PullChosenChannel(*PullChosenChannelRequest, ChannelApi_PullChosenChannelServer) error {
	return status.Errorf(codes.Unimplemented, "method PullChosenChannel not implemented")
}
func (UnimplementedChannelApiServer) mustEmbedUnimplementedChannelApiServer() {}

// UnsafeChannelApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChannelApiServer will
// result in compilation errors.
type UnsafeChannelApiServer interface {
	mustEmbedUnimplementedChannelApiServer()
}

func RegisterChannelApiServer(s grpc.ServiceRegistrar, srv ChannelApiServer) {
	s.RegisterService(&ChannelApi_ServiceDesc, srv)
}

func _ChannelApi_GetChosenChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChosenChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelApiServer).GetChosenChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChannelApi_GetChosenChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelApiServer).GetChosenChannel(ctx, req.(*GetChosenChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelApi_ChooseChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChooseChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelApiServer).ChooseChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChannelApi_ChooseChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelApiServer).ChooseChannel(ctx, req.(*ChooseChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelApi_AdjustChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdjustChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelApiServer).AdjustChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChannelApi_AdjustChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelApiServer).AdjustChannel(ctx, req.(*AdjustChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelApi_ReturnChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReturnChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelApiServer).ReturnChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChannelApi_ReturnChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelApiServer).ReturnChannel(ctx, req.(*ReturnChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelApi_PullChosenChannel_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullChosenChannelRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChannelApiServer).PullChosenChannel(m, &channelApiPullChosenChannelServer{stream})
}

type ChannelApi_PullChosenChannelServer interface {
	Send(*PullChosenChannelResponse) error
	grpc.ServerStream
}

type channelApiPullChosenChannelServer struct {
	grpc.ServerStream
}

func (x *channelApiPullChosenChannelServer) Send(m *PullChosenChannelResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ChannelApi_ServiceDesc is the grpc.ServiceDesc for ChannelApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChannelApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.traits.ChannelApi",
	HandlerType: (*ChannelApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChosenChannel",
			Handler:    _ChannelApi_GetChosenChannel_Handler,
		},
		{
			MethodName: "ChooseChannel",
			Handler:    _ChannelApi_ChooseChannel_Handler,
		},
		{
			MethodName: "AdjustChannel",
			Handler:    _ChannelApi_AdjustChannel_Handler,
		},
		{
			MethodName: "ReturnChannel",
			Handler:    _ChannelApi_ReturnChannel_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PullChosenChannel",
			Handler:       _ChannelApi_PullChosenChannel_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "traits/channel.proto",
}

const (
	ChannelInfo_DescribeChosenChannel_FullMethodName = "/smartcore.traits.ChannelInfo/DescribeChosenChannel"
)

// ChannelInfoClient is the client API for ChannelInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChannelInfoClient interface {
	// Get information about how a named device implements channel choosing and adjustment
	DescribeChosenChannel(ctx context.Context, in *DescribeChosenChannelRequest, opts ...grpc.CallOption) (*ChosenChannelSupport, error)
}

type channelInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewChannelInfoClient(cc grpc.ClientConnInterface) ChannelInfoClient {
	return &channelInfoClient{cc}
}

func (c *channelInfoClient) DescribeChosenChannel(ctx context.Context, in *DescribeChosenChannelRequest, opts ...grpc.CallOption) (*ChosenChannelSupport, error) {
	out := new(ChosenChannelSupport)
	err := c.cc.Invoke(ctx, ChannelInfo_DescribeChosenChannel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChannelInfoServer is the server API for ChannelInfo service.
// All implementations must embed UnimplementedChannelInfoServer
// for forward compatibility
type ChannelInfoServer interface {
	// Get information about how a named device implements channel choosing and adjustment
	DescribeChosenChannel(context.Context, *DescribeChosenChannelRequest) (*ChosenChannelSupport, error)
	mustEmbedUnimplementedChannelInfoServer()
}

// UnimplementedChannelInfoServer must be embedded to have forward compatible implementations.
type UnimplementedChannelInfoServer struct {
}

func (UnimplementedChannelInfoServer) DescribeChosenChannel(context.Context, *DescribeChosenChannelRequest) (*ChosenChannelSupport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeChosenChannel not implemented")
}
func (UnimplementedChannelInfoServer) mustEmbedUnimplementedChannelInfoServer() {}

// UnsafeChannelInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChannelInfoServer will
// result in compilation errors.
type UnsafeChannelInfoServer interface {
	mustEmbedUnimplementedChannelInfoServer()
}

func RegisterChannelInfoServer(s grpc.ServiceRegistrar, srv ChannelInfoServer) {
	s.RegisterService(&ChannelInfo_ServiceDesc, srv)
}

func _ChannelInfo_DescribeChosenChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeChosenChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelInfoServer).DescribeChosenChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChannelInfo_DescribeChosenChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelInfoServer).DescribeChosenChannel(ctx, req.(*DescribeChosenChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChannelInfo_ServiceDesc is the grpc.ServiceDesc for ChannelInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChannelInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.traits.ChannelInfo",
	HandlerType: (*ChannelInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeChosenChannel",
			Handler:    _ChannelInfo_DescribeChosenChannel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "traits/channel.proto",
}
