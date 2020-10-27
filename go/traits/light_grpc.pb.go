// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package traits

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LightApiClient is the client API for LightApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LightApiClient interface {
	// request that the brightness be changed
	UpdateBrightness(ctx context.Context, in *UpdateBrightnessRequest, opts ...grpc.CallOption) (*Brightness, error)
	// get the current value of the range
	GetBrightness(ctx context.Context, in *GetBrightnessRequest, opts ...grpc.CallOption) (*Brightness, error)
	// request updates to changes in the range value
	PullBrightness(ctx context.Context, in *PullBrightnessRequest, opts ...grpc.CallOption) (LightApi_PullBrightnessClient, error)
}

type lightApiClient struct {
	cc grpc.ClientConnInterface
}

func NewLightApiClient(cc grpc.ClientConnInterface) LightApiClient {
	return &lightApiClient{cc}
}

func (c *lightApiClient) UpdateBrightness(ctx context.Context, in *UpdateBrightnessRequest, opts ...grpc.CallOption) (*Brightness, error) {
	out := new(Brightness)
	err := c.cc.Invoke(ctx, "/smartcore.traits.LightApi/UpdateBrightness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lightApiClient) GetBrightness(ctx context.Context, in *GetBrightnessRequest, opts ...grpc.CallOption) (*Brightness, error) {
	out := new(Brightness)
	err := c.cc.Invoke(ctx, "/smartcore.traits.LightApi/GetBrightness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lightApiClient) PullBrightness(ctx context.Context, in *PullBrightnessRequest, opts ...grpc.CallOption) (LightApi_PullBrightnessClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LightApi_serviceDesc.Streams[0], "/smartcore.traits.LightApi/PullBrightness", opts...)
	if err != nil {
		return nil, err
	}
	x := &lightApiPullBrightnessClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LightApi_PullBrightnessClient interface {
	Recv() (*PullBrightnessResponse, error)
	grpc.ClientStream
}

type lightApiPullBrightnessClient struct {
	grpc.ClientStream
}

func (x *lightApiPullBrightnessClient) Recv() (*PullBrightnessResponse, error) {
	m := new(PullBrightnessResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LightApiServer is the server API for LightApi service.
// All implementations must embed UnimplementedLightApiServer
// for forward compatibility
type LightApiServer interface {
	// request that the brightness be changed
	UpdateBrightness(context.Context, *UpdateBrightnessRequest) (*Brightness, error)
	// get the current value of the range
	GetBrightness(context.Context, *GetBrightnessRequest) (*Brightness, error)
	// request updates to changes in the range value
	PullBrightness(*PullBrightnessRequest, LightApi_PullBrightnessServer) error
	mustEmbedUnimplementedLightApiServer()
}

// UnimplementedLightApiServer must be embedded to have forward compatible implementations.
type UnimplementedLightApiServer struct {
}

func (*UnimplementedLightApiServer) UpdateBrightness(context.Context, *UpdateBrightnessRequest) (*Brightness, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBrightness not implemented")
}
func (*UnimplementedLightApiServer) GetBrightness(context.Context, *GetBrightnessRequest) (*Brightness, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBrightness not implemented")
}
func (*UnimplementedLightApiServer) PullBrightness(*PullBrightnessRequest, LightApi_PullBrightnessServer) error {
	return status.Errorf(codes.Unimplemented, "method PullBrightness not implemented")
}
func (*UnimplementedLightApiServer) mustEmbedUnimplementedLightApiServer() {}

func RegisterLightApiServer(s *grpc.Server, srv LightApiServer) {
	s.RegisterService(&_LightApi_serviceDesc, srv)
}

func _LightApi_UpdateBrightness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBrightnessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LightApiServer).UpdateBrightness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.traits.LightApi/UpdateBrightness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LightApiServer).UpdateBrightness(ctx, req.(*UpdateBrightnessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LightApi_GetBrightness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBrightnessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LightApiServer).GetBrightness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.traits.LightApi/GetBrightness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LightApiServer).GetBrightness(ctx, req.(*GetBrightnessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LightApi_PullBrightness_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullBrightnessRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LightApiServer).PullBrightness(m, &lightApiPullBrightnessServer{stream})
}

type LightApi_PullBrightnessServer interface {
	Send(*PullBrightnessResponse) error
	grpc.ServerStream
}

type lightApiPullBrightnessServer struct {
	grpc.ServerStream
}

func (x *lightApiPullBrightnessServer) Send(m *PullBrightnessResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _LightApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.traits.LightApi",
	HandlerType: (*LightApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateBrightness",
			Handler:    _LightApi_UpdateBrightness_Handler,
		},
		{
			MethodName: "GetBrightness",
			Handler:    _LightApi_GetBrightness_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PullBrightness",
			Handler:       _LightApi_PullBrightness_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "traits/light.proto",
}

// LightInfoClient is the client API for LightInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LightInfoClient interface {
	// Get information about how a named device implements Brightness features
	DescribeBrightness(ctx context.Context, in *DescribeBrightnessRequest, opts ...grpc.CallOption) (*BrightnessSupport, error)
}

type lightInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewLightInfoClient(cc grpc.ClientConnInterface) LightInfoClient {
	return &lightInfoClient{cc}
}

func (c *lightInfoClient) DescribeBrightness(ctx context.Context, in *DescribeBrightnessRequest, opts ...grpc.CallOption) (*BrightnessSupport, error) {
	out := new(BrightnessSupport)
	err := c.cc.Invoke(ctx, "/smartcore.traits.LightInfo/DescribeBrightness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LightInfoServer is the server API for LightInfo service.
// All implementations must embed UnimplementedLightInfoServer
// for forward compatibility
type LightInfoServer interface {
	// Get information about how a named device implements Brightness features
	DescribeBrightness(context.Context, *DescribeBrightnessRequest) (*BrightnessSupport, error)
	mustEmbedUnimplementedLightInfoServer()
}

// UnimplementedLightInfoServer must be embedded to have forward compatible implementations.
type UnimplementedLightInfoServer struct {
}

func (*UnimplementedLightInfoServer) DescribeBrightness(context.Context, *DescribeBrightnessRequest) (*BrightnessSupport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeBrightness not implemented")
}
func (*UnimplementedLightInfoServer) mustEmbedUnimplementedLightInfoServer() {}

func RegisterLightInfoServer(s *grpc.Server, srv LightInfoServer) {
	s.RegisterService(&_LightInfo_serviceDesc, srv)
}

func _LightInfo_DescribeBrightness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeBrightnessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LightInfoServer).DescribeBrightness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.traits.LightInfo/DescribeBrightness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LightInfoServer).DescribeBrightness(ctx, req.(*DescribeBrightnessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LightInfo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.traits.LightInfo",
	HandlerType: (*LightInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeBrightness",
			Handler:    _LightInfo_DescribeBrightness_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "traits/light.proto",
}
