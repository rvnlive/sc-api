// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: info/health.proto

package info

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

// HealthClient is the client API for Health service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HealthClient interface {
	GetHealthState(ctx context.Context, in *GetHealthStateRequest, opts ...grpc.CallOption) (*HealthState, error)
	PullHealthStates(ctx context.Context, in *PullHealthStatesRequest, opts ...grpc.CallOption) (Health_PullHealthStatesClient, error)
}

type healthClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthClient(cc grpc.ClientConnInterface) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) GetHealthState(ctx context.Context, in *GetHealthStateRequest, opts ...grpc.CallOption) (*HealthState, error) {
	out := new(HealthState)
	err := c.cc.Invoke(ctx, "/smartcore.info.Health/GetHealthState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthClient) PullHealthStates(ctx context.Context, in *PullHealthStatesRequest, opts ...grpc.CallOption) (Health_PullHealthStatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Health_ServiceDesc.Streams[0], "/smartcore.info.Health/PullHealthStates", opts...)
	if err != nil {
		return nil, err
	}
	x := &healthPullHealthStatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Health_PullHealthStatesClient interface {
	Recv() (*PullHealthStatesResponse, error)
	grpc.ClientStream
}

type healthPullHealthStatesClient struct {
	grpc.ClientStream
}

func (x *healthPullHealthStatesClient) Recv() (*PullHealthStatesResponse, error) {
	m := new(PullHealthStatesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HealthServer is the server API for Health service.
// All implementations must embed UnimplementedHealthServer
// for forward compatibility
type HealthServer interface {
	GetHealthState(context.Context, *GetHealthStateRequest) (*HealthState, error)
	PullHealthStates(*PullHealthStatesRequest, Health_PullHealthStatesServer) error
	mustEmbedUnimplementedHealthServer()
}

// UnimplementedHealthServer must be embedded to have forward compatible implementations.
type UnimplementedHealthServer struct {
}

func (UnimplementedHealthServer) GetHealthState(context.Context, *GetHealthStateRequest) (*HealthState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHealthState not implemented")
}
func (UnimplementedHealthServer) PullHealthStates(*PullHealthStatesRequest, Health_PullHealthStatesServer) error {
	return status.Errorf(codes.Unimplemented, "method PullHealthStates not implemented")
}
func (UnimplementedHealthServer) mustEmbedUnimplementedHealthServer() {}

// UnsafeHealthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthServer will
// result in compilation errors.
type UnsafeHealthServer interface {
	mustEmbedUnimplementedHealthServer()
}

func RegisterHealthServer(s grpc.ServiceRegistrar, srv HealthServer) {
	s.RegisterService(&Health_ServiceDesc, srv)
}

func _Health_GetHealthState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHealthStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).GetHealthState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.info.Health/GetHealthState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).GetHealthState(ctx, req.(*GetHealthStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Health_PullHealthStates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullHealthStatesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HealthServer).PullHealthStates(m, &healthPullHealthStatesServer{stream})
}

type Health_PullHealthStatesServer interface {
	Send(*PullHealthStatesResponse) error
	grpc.ServerStream
}

type healthPullHealthStatesServer struct {
	grpc.ServerStream
}

func (x *healthPullHealthStatesServer) Send(m *PullHealthStatesResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Health_ServiceDesc is the grpc.ServiceDesc for Health service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Health_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.info.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHealthState",
			Handler:    _Health_GetHealthState_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PullHealthStates",
			Handler:       _Health_PullHealthStates_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "info/health.proto",
}
