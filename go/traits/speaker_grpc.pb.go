// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package traits

import (
	context "context"
	types "github.com/smart-core-os/sc-api/go/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SpeakerApiClient is the client API for SpeakerApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpeakerApiClient interface {
	// Get the current state of the volume for the device
	GetVolume(ctx context.Context, in *GetSpeakerVolumeRequest, opts ...grpc.CallOption) (*types.AudioLevel, error)
	// update the volume state for the device
	UpdateVolume(ctx context.Context, in *UpdateSpeakerVolumeRequest, opts ...grpc.CallOption) (*types.AudioLevel, error)
	PullVolume(ctx context.Context, in *PullSpeakerVolumeRequest, opts ...grpc.CallOption) (SpeakerApi_PullVolumeClient, error)
}

type speakerApiClient struct {
	cc grpc.ClientConnInterface
}

func NewSpeakerApiClient(cc grpc.ClientConnInterface) SpeakerApiClient {
	return &speakerApiClient{cc}
}

func (c *speakerApiClient) GetVolume(ctx context.Context, in *GetSpeakerVolumeRequest, opts ...grpc.CallOption) (*types.AudioLevel, error) {
	out := new(types.AudioLevel)
	err := c.cc.Invoke(ctx, "/smartcore.traits.SpeakerApi/GetVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *speakerApiClient) UpdateVolume(ctx context.Context, in *UpdateSpeakerVolumeRequest, opts ...grpc.CallOption) (*types.AudioLevel, error) {
	out := new(types.AudioLevel)
	err := c.cc.Invoke(ctx, "/smartcore.traits.SpeakerApi/UpdateVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *speakerApiClient) PullVolume(ctx context.Context, in *PullSpeakerVolumeRequest, opts ...grpc.CallOption) (SpeakerApi_PullVolumeClient, error) {
	stream, err := c.cc.NewStream(ctx, &SpeakerApi_ServiceDesc.Streams[0], "/smartcore.traits.SpeakerApi/PullVolume", opts...)
	if err != nil {
		return nil, err
	}
	x := &speakerApiPullVolumeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SpeakerApi_PullVolumeClient interface {
	Recv() (*PullSpeakerVolumeResponse, error)
	grpc.ClientStream
}

type speakerApiPullVolumeClient struct {
	grpc.ClientStream
}

func (x *speakerApiPullVolumeClient) Recv() (*PullSpeakerVolumeResponse, error) {
	m := new(PullSpeakerVolumeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SpeakerApiServer is the server API for SpeakerApi service.
// All implementations must embed UnimplementedSpeakerApiServer
// for forward compatibility
type SpeakerApiServer interface {
	// Get the current state of the volume for the device
	GetVolume(context.Context, *GetSpeakerVolumeRequest) (*types.AudioLevel, error)
	// update the volume state for the device
	UpdateVolume(context.Context, *UpdateSpeakerVolumeRequest) (*types.AudioLevel, error)
	PullVolume(*PullSpeakerVolumeRequest, SpeakerApi_PullVolumeServer) error
	mustEmbedUnimplementedSpeakerApiServer()
}

// UnimplementedSpeakerApiServer must be embedded to have forward compatible implementations.
type UnimplementedSpeakerApiServer struct {
}

func (UnimplementedSpeakerApiServer) GetVolume(context.Context, *GetSpeakerVolumeRequest) (*types.AudioLevel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVolume not implemented")
}
func (UnimplementedSpeakerApiServer) UpdateVolume(context.Context, *UpdateSpeakerVolumeRequest) (*types.AudioLevel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVolume not implemented")
}
func (UnimplementedSpeakerApiServer) PullVolume(*PullSpeakerVolumeRequest, SpeakerApi_PullVolumeServer) error {
	return status.Errorf(codes.Unimplemented, "method PullVolume not implemented")
}
func (UnimplementedSpeakerApiServer) mustEmbedUnimplementedSpeakerApiServer() {}

// UnsafeSpeakerApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpeakerApiServer will
// result in compilation errors.
type UnsafeSpeakerApiServer interface {
	mustEmbedUnimplementedSpeakerApiServer()
}

func RegisterSpeakerApiServer(s grpc.ServiceRegistrar, srv SpeakerApiServer) {
	s.RegisterService(&SpeakerApi_ServiceDesc, srv)
}

func _SpeakerApi_GetVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSpeakerVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpeakerApiServer).GetVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.traits.SpeakerApi/GetVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpeakerApiServer).GetVolume(ctx, req.(*GetSpeakerVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpeakerApi_UpdateVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSpeakerVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpeakerApiServer).UpdateVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.traits.SpeakerApi/UpdateVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpeakerApiServer).UpdateVolume(ctx, req.(*UpdateSpeakerVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpeakerApi_PullVolume_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullSpeakerVolumeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SpeakerApiServer).PullVolume(m, &speakerApiPullVolumeServer{stream})
}

type SpeakerApi_PullVolumeServer interface {
	Send(*PullSpeakerVolumeResponse) error
	grpc.ServerStream
}

type speakerApiPullVolumeServer struct {
	grpc.ServerStream
}

func (x *speakerApiPullVolumeServer) Send(m *PullSpeakerVolumeResponse) error {
	return x.ServerStream.SendMsg(m)
}

// SpeakerApi_ServiceDesc is the grpc.ServiceDesc for SpeakerApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SpeakerApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.traits.SpeakerApi",
	HandlerType: (*SpeakerApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVolume",
			Handler:    _SpeakerApi_GetVolume_Handler,
		},
		{
			MethodName: "UpdateVolume",
			Handler:    _SpeakerApi_UpdateVolume_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PullVolume",
			Handler:       _SpeakerApi_PullVolume_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "traits/speaker.proto",
}

// SpeakerInfoClient is the client API for SpeakerInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpeakerInfoClient interface {
	// Get information about how a named device implements volume features
	DescribeVolume(ctx context.Context, in *DescribeVolumeRequest, opts ...grpc.CallOption) (*VolumeSupport, error)
}

type speakerInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewSpeakerInfoClient(cc grpc.ClientConnInterface) SpeakerInfoClient {
	return &speakerInfoClient{cc}
}

func (c *speakerInfoClient) DescribeVolume(ctx context.Context, in *DescribeVolumeRequest, opts ...grpc.CallOption) (*VolumeSupport, error) {
	out := new(VolumeSupport)
	err := c.cc.Invoke(ctx, "/smartcore.traits.SpeakerInfo/DescribeVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpeakerInfoServer is the server API for SpeakerInfo service.
// All implementations must embed UnimplementedSpeakerInfoServer
// for forward compatibility
type SpeakerInfoServer interface {
	// Get information about how a named device implements volume features
	DescribeVolume(context.Context, *DescribeVolumeRequest) (*VolumeSupport, error)
	mustEmbedUnimplementedSpeakerInfoServer()
}

// UnimplementedSpeakerInfoServer must be embedded to have forward compatible implementations.
type UnimplementedSpeakerInfoServer struct {
}

func (UnimplementedSpeakerInfoServer) DescribeVolume(context.Context, *DescribeVolumeRequest) (*VolumeSupport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeVolume not implemented")
}
func (UnimplementedSpeakerInfoServer) mustEmbedUnimplementedSpeakerInfoServer() {}

// UnsafeSpeakerInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpeakerInfoServer will
// result in compilation errors.
type UnsafeSpeakerInfoServer interface {
	mustEmbedUnimplementedSpeakerInfoServer()
}

func RegisterSpeakerInfoServer(s grpc.ServiceRegistrar, srv SpeakerInfoServer) {
	s.RegisterService(&SpeakerInfo_ServiceDesc, srv)
}

func _SpeakerInfo_DescribeVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpeakerInfoServer).DescribeVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.traits.SpeakerInfo/DescribeVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpeakerInfoServer).DescribeVolume(ctx, req.(*DescribeVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SpeakerInfo_ServiceDesc is the grpc.ServiceDesc for SpeakerInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SpeakerInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.traits.SpeakerInfo",
	HandlerType: (*SpeakerInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeVolume",
			Handler:    _SpeakerInfo_DescribeVolume_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "traits/speaker.proto",
}
