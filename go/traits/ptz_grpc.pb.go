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
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PtzApiClient is the client API for PtzApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PtzApiClient interface {
	// Set the target state for the device
	GetPtz(ctx context.Context, in *GetPtzRequest, opts ...grpc.CallOption) (*Ptz, error)
	// Set the target state for the device
	UpdatePtz(ctx context.Context, in *UpdatePtzRequest, opts ...grpc.CallOption) (*Ptz, error)
	// Stop the device if it is extending or retracting, returns the current known state after stopping.
	Stop(ctx context.Context, in *StopPtzRequest, opts ...grpc.CallOption) (*Ptz, error)
	// Create a preset for the ptz position. If no preset ptz position is specified then use the current position
	CreatePreset(ctx context.Context, in *CreatePtzPresetRequest, opts ...grpc.CallOption) (*PtzPreset, error)
	// Get notified of changes to the OnOffState of a device
	PullPtz(ctx context.Context, in *PullPtzRequest, opts ...grpc.CallOption) (PtzApi_PullPtzClient, error)
}

type ptzApiClient struct {
	cc grpc.ClientConnInterface
}

func NewPtzApiClient(cc grpc.ClientConnInterface) PtzApiClient {
	return &ptzApiClient{cc}
}

func (c *ptzApiClient) GetPtz(ctx context.Context, in *GetPtzRequest, opts ...grpc.CallOption) (*Ptz, error) {
	out := new(Ptz)
	err := c.cc.Invoke(ctx, "/smartcore.traits.PtzApi/GetPtz", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ptzApiClient) UpdatePtz(ctx context.Context, in *UpdatePtzRequest, opts ...grpc.CallOption) (*Ptz, error) {
	out := new(Ptz)
	err := c.cc.Invoke(ctx, "/smartcore.traits.PtzApi/UpdatePtz", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ptzApiClient) Stop(ctx context.Context, in *StopPtzRequest, opts ...grpc.CallOption) (*Ptz, error) {
	out := new(Ptz)
	err := c.cc.Invoke(ctx, "/smartcore.traits.PtzApi/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ptzApiClient) CreatePreset(ctx context.Context, in *CreatePtzPresetRequest, opts ...grpc.CallOption) (*PtzPreset, error) {
	out := new(PtzPreset)
	err := c.cc.Invoke(ctx, "/smartcore.traits.PtzApi/CreatePreset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ptzApiClient) PullPtz(ctx context.Context, in *PullPtzRequest, opts ...grpc.CallOption) (PtzApi_PullPtzClient, error) {
	stream, err := c.cc.NewStream(ctx, &PtzApi_ServiceDesc.Streams[0], "/smartcore.traits.PtzApi/PullPtz", opts...)
	if err != nil {
		return nil, err
	}
	x := &ptzApiPullPtzClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PtzApi_PullPtzClient interface {
	Recv() (*PullPtzResponse, error)
	grpc.ClientStream
}

type ptzApiPullPtzClient struct {
	grpc.ClientStream
}

func (x *ptzApiPullPtzClient) Recv() (*PullPtzResponse, error) {
	m := new(PullPtzResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PtzApiServer is the server API for PtzApi service.
// All implementations must embed UnimplementedPtzApiServer
// for forward compatibility
type PtzApiServer interface {
	// Set the target state for the device
	GetPtz(context.Context, *GetPtzRequest) (*Ptz, error)
	// Set the target state for the device
	UpdatePtz(context.Context, *UpdatePtzRequest) (*Ptz, error)
	// Stop the device if it is extending or retracting, returns the current known state after stopping.
	Stop(context.Context, *StopPtzRequest) (*Ptz, error)
	// Create a preset for the ptz position. If no preset ptz position is specified then use the current position
	CreatePreset(context.Context, *CreatePtzPresetRequest) (*PtzPreset, error)
	// Get notified of changes to the OnOffState of a device
	PullPtz(*PullPtzRequest, PtzApi_PullPtzServer) error
	mustEmbedUnimplementedPtzApiServer()
}

// UnimplementedPtzApiServer must be embedded to have forward compatible implementations.
type UnimplementedPtzApiServer struct {
}

func (UnimplementedPtzApiServer) GetPtz(context.Context, *GetPtzRequest) (*Ptz, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPtz not implemented")
}
func (UnimplementedPtzApiServer) UpdatePtz(context.Context, *UpdatePtzRequest) (*Ptz, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePtz not implemented")
}
func (UnimplementedPtzApiServer) Stop(context.Context, *StopPtzRequest) (*Ptz, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedPtzApiServer) CreatePreset(context.Context, *CreatePtzPresetRequest) (*PtzPreset, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePreset not implemented")
}
func (UnimplementedPtzApiServer) PullPtz(*PullPtzRequest, PtzApi_PullPtzServer) error {
	return status.Errorf(codes.Unimplemented, "method PullPtz not implemented")
}
func (UnimplementedPtzApiServer) mustEmbedUnimplementedPtzApiServer() {}

// UnsafePtzApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PtzApiServer will
// result in compilation errors.
type UnsafePtzApiServer interface {
	mustEmbedUnimplementedPtzApiServer()
}

func RegisterPtzApiServer(s grpc.ServiceRegistrar, srv PtzApiServer) {
	s.RegisterService(&PtzApi_ServiceDesc, srv)
}

func _PtzApi_GetPtz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPtzRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PtzApiServer).GetPtz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.traits.PtzApi/GetPtz",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PtzApiServer).GetPtz(ctx, req.(*GetPtzRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PtzApi_UpdatePtz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePtzRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PtzApiServer).UpdatePtz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.traits.PtzApi/UpdatePtz",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PtzApiServer).UpdatePtz(ctx, req.(*UpdatePtzRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PtzApi_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopPtzRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PtzApiServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.traits.PtzApi/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PtzApiServer).Stop(ctx, req.(*StopPtzRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PtzApi_CreatePreset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePtzPresetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PtzApiServer).CreatePreset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.traits.PtzApi/CreatePreset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PtzApiServer).CreatePreset(ctx, req.(*CreatePtzPresetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PtzApi_PullPtz_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullPtzRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PtzApiServer).PullPtz(m, &ptzApiPullPtzServer{stream})
}

type PtzApi_PullPtzServer interface {
	Send(*PullPtzResponse) error
	grpc.ServerStream
}

type ptzApiPullPtzServer struct {
	grpc.ServerStream
}

func (x *ptzApiPullPtzServer) Send(m *PullPtzResponse) error {
	return x.ServerStream.SendMsg(m)
}

// PtzApi_ServiceDesc is the grpc.ServiceDesc for PtzApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PtzApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.traits.PtzApi",
	HandlerType: (*PtzApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPtz",
			Handler:    _PtzApi_GetPtz_Handler,
		},
		{
			MethodName: "UpdatePtz",
			Handler:    _PtzApi_UpdatePtz_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _PtzApi_Stop_Handler,
		},
		{
			MethodName: "CreatePreset",
			Handler:    _PtzApi_CreatePreset_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PullPtz",
			Handler:       _PtzApi_PullPtz_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "traits/ptz.proto",
}

// PtzInfoClient is the client API for PtzInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PtzInfoClient interface {
	// Get information about how a named device implements Ptz features
	DescribePtz(ctx context.Context, in *DescribePtzRequest, opts ...grpc.CallOption) (*PtzSupport, error)
}

type ptzInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewPtzInfoClient(cc grpc.ClientConnInterface) PtzInfoClient {
	return &ptzInfoClient{cc}
}

func (c *ptzInfoClient) DescribePtz(ctx context.Context, in *DescribePtzRequest, opts ...grpc.CallOption) (*PtzSupport, error) {
	out := new(PtzSupport)
	err := c.cc.Invoke(ctx, "/smartcore.traits.PtzInfo/DescribePtz", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PtzInfoServer is the server API for PtzInfo service.
// All implementations must embed UnimplementedPtzInfoServer
// for forward compatibility
type PtzInfoServer interface {
	// Get information about how a named device implements Ptz features
	DescribePtz(context.Context, *DescribePtzRequest) (*PtzSupport, error)
	mustEmbedUnimplementedPtzInfoServer()
}

// UnimplementedPtzInfoServer must be embedded to have forward compatible implementations.
type UnimplementedPtzInfoServer struct {
}

func (UnimplementedPtzInfoServer) DescribePtz(context.Context, *DescribePtzRequest) (*PtzSupport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribePtz not implemented")
}
func (UnimplementedPtzInfoServer) mustEmbedUnimplementedPtzInfoServer() {}

// UnsafePtzInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PtzInfoServer will
// result in compilation errors.
type UnsafePtzInfoServer interface {
	mustEmbedUnimplementedPtzInfoServer()
}

func RegisterPtzInfoServer(s grpc.ServiceRegistrar, srv PtzInfoServer) {
	s.RegisterService(&PtzInfo_ServiceDesc, srv)
}

func _PtzInfo_DescribePtz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribePtzRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PtzInfoServer).DescribePtz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.traits.PtzInfo/DescribePtz",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PtzInfoServer).DescribePtz(ctx, req.(*DescribePtzRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PtzInfo_ServiceDesc is the grpc.ServiceDesc for PtzInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PtzInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.traits.PtzInfo",
	HandlerType: (*PtzInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribePtz",
			Handler:    _PtzInfo_DescribePtz_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "traits/ptz.proto",
}
