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

// EmergencyApiClient is the client API for EmergencyApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmergencyApiClient interface {
	GetEmergency(ctx context.Context, in *GetEmergencyRequest, opts ...grpc.CallOption) (*Emergency, error)
	UpdateEmergency(ctx context.Context, in *UpdateEmergencyRequest, opts ...grpc.CallOption) (*Emergency, error)
	PullEmergency(ctx context.Context, in *PullEmergencyRequest, opts ...grpc.CallOption) (EmergencyApi_PullEmergencyClient, error)
}

type emergencyApiClient struct {
	cc grpc.ClientConnInterface
}

func NewEmergencyApiClient(cc grpc.ClientConnInterface) EmergencyApiClient {
	return &emergencyApiClient{cc}
}

func (c *emergencyApiClient) GetEmergency(ctx context.Context, in *GetEmergencyRequest, opts ...grpc.CallOption) (*Emergency, error) {
	out := new(Emergency)
	err := c.cc.Invoke(ctx, "/smartcore.api.device.traits.EmergencyApi/GetEmergency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emergencyApiClient) UpdateEmergency(ctx context.Context, in *UpdateEmergencyRequest, opts ...grpc.CallOption) (*Emergency, error) {
	out := new(Emergency)
	err := c.cc.Invoke(ctx, "/smartcore.api.device.traits.EmergencyApi/UpdateEmergency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emergencyApiClient) PullEmergency(ctx context.Context, in *PullEmergencyRequest, opts ...grpc.CallOption) (EmergencyApi_PullEmergencyClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EmergencyApi_serviceDesc.Streams[0], "/smartcore.api.device.traits.EmergencyApi/PullEmergency", opts...)
	if err != nil {
		return nil, err
	}
	x := &emergencyApiPullEmergencyClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EmergencyApi_PullEmergencyClient interface {
	Recv() (*PullEmergencyResponse, error)
	grpc.ClientStream
}

type emergencyApiPullEmergencyClient struct {
	grpc.ClientStream
}

func (x *emergencyApiPullEmergencyClient) Recv() (*PullEmergencyResponse, error) {
	m := new(PullEmergencyResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EmergencyApiServer is the server API for EmergencyApi service.
// All implementations must embed UnimplementedEmergencyApiServer
// for forward compatibility
type EmergencyApiServer interface {
	GetEmergency(context.Context, *GetEmergencyRequest) (*Emergency, error)
	UpdateEmergency(context.Context, *UpdateEmergencyRequest) (*Emergency, error)
	PullEmergency(*PullEmergencyRequest, EmergencyApi_PullEmergencyServer) error
	mustEmbedUnimplementedEmergencyApiServer()
}

// UnimplementedEmergencyApiServer must be embedded to have forward compatible implementations.
type UnimplementedEmergencyApiServer struct {
}

func (*UnimplementedEmergencyApiServer) GetEmergency(context.Context, *GetEmergencyRequest) (*Emergency, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmergency not implemented")
}
func (*UnimplementedEmergencyApiServer) UpdateEmergency(context.Context, *UpdateEmergencyRequest) (*Emergency, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmergency not implemented")
}
func (*UnimplementedEmergencyApiServer) PullEmergency(*PullEmergencyRequest, EmergencyApi_PullEmergencyServer) error {
	return status.Errorf(codes.Unimplemented, "method PullEmergency not implemented")
}
func (*UnimplementedEmergencyApiServer) mustEmbedUnimplementedEmergencyApiServer() {}

func RegisterEmergencyApiServer(s *grpc.Server, srv EmergencyApiServer) {
	s.RegisterService(&_EmergencyApi_serviceDesc, srv)
}

func _EmergencyApi_GetEmergency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEmergencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmergencyApiServer).GetEmergency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.api.device.traits.EmergencyApi/GetEmergency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmergencyApiServer).GetEmergency(ctx, req.(*GetEmergencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmergencyApi_UpdateEmergency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEmergencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmergencyApiServer).UpdateEmergency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.api.device.traits.EmergencyApi/UpdateEmergency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmergencyApiServer).UpdateEmergency(ctx, req.(*UpdateEmergencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmergencyApi_PullEmergency_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullEmergencyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EmergencyApiServer).PullEmergency(m, &emergencyApiPullEmergencyServer{stream})
}

type EmergencyApi_PullEmergencyServer interface {
	Send(*PullEmergencyResponse) error
	grpc.ServerStream
}

type emergencyApiPullEmergencyServer struct {
	grpc.ServerStream
}

func (x *emergencyApiPullEmergencyServer) Send(m *PullEmergencyResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _EmergencyApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.api.device.traits.EmergencyApi",
	HandlerType: (*EmergencyApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEmergency",
			Handler:    _EmergencyApi_GetEmergency_Handler,
		},
		{
			MethodName: "UpdateEmergency",
			Handler:    _EmergencyApi_UpdateEmergency_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PullEmergency",
			Handler:       _EmergencyApi_PullEmergency_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "device/traits/emergency.proto",
}
