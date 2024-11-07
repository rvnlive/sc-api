// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: info/info.proto

package info

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
	Info_ListDevices_FullMethodName = "/smartcore.info.Info/ListDevices"
	Info_PullDevices_FullMethodName = "/smartcore.info.Info/PullDevices"
)

// InfoClient is the client API for Info service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Allow exposing information about an endpoint related to devices.
//
// Deprecated: use the Parent trait for listing devices, use Metadata trait for display name.
type InfoClient interface {
	// Get devices that this service knows about
	ListDevices(ctx context.Context, in *ListDevicesRequest, opts ...grpc.CallOption) (*ListDevicesResponse, error)
	// Open a server stream that responds with changes to the set of devices described in the request
	PullDevices(ctx context.Context, in *PullDevicesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PullDevicesResponse], error)
}

type infoClient struct {
	cc grpc.ClientConnInterface
}

func NewInfoClient(cc grpc.ClientConnInterface) InfoClient {
	return &infoClient{cc}
}

func (c *infoClient) ListDevices(ctx context.Context, in *ListDevicesRequest, opts ...grpc.CallOption) (*ListDevicesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListDevicesResponse)
	err := c.cc.Invoke(ctx, Info_ListDevices_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoClient) PullDevices(ctx context.Context, in *PullDevicesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PullDevicesResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Info_ServiceDesc.Streams[0], Info_PullDevices_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[PullDevicesRequest, PullDevicesResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Info_PullDevicesClient = grpc.ServerStreamingClient[PullDevicesResponse]

// InfoServer is the server API for Info service.
// All implementations must embed UnimplementedInfoServer
// for forward compatibility.
//
// Allow exposing information about an endpoint related to devices.
//
// Deprecated: use the Parent trait for listing devices, use Metadata trait for display name.
type InfoServer interface {
	// Get devices that this service knows about
	ListDevices(context.Context, *ListDevicesRequest) (*ListDevicesResponse, error)
	// Open a server stream that responds with changes to the set of devices described in the request
	PullDevices(*PullDevicesRequest, grpc.ServerStreamingServer[PullDevicesResponse]) error
	mustEmbedUnimplementedInfoServer()
}

// UnimplementedInfoServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedInfoServer struct{}

func (UnimplementedInfoServer) ListDevices(context.Context, *ListDevicesRequest) (*ListDevicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDevices not implemented")
}
func (UnimplementedInfoServer) PullDevices(*PullDevicesRequest, grpc.ServerStreamingServer[PullDevicesResponse]) error {
	return status.Errorf(codes.Unimplemented, "method PullDevices not implemented")
}
func (UnimplementedInfoServer) mustEmbedUnimplementedInfoServer() {}
func (UnimplementedInfoServer) testEmbeddedByValue()              {}

// UnsafeInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InfoServer will
// result in compilation errors.
type UnsafeInfoServer interface {
	mustEmbedUnimplementedInfoServer()
}

func RegisterInfoServer(s grpc.ServiceRegistrar, srv InfoServer) {
	// If the following call pancis, it indicates UnimplementedInfoServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Info_ServiceDesc, srv)
}

func _Info_ListDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServer).ListDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Info_ListDevices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServer).ListDevices(ctx, req.(*ListDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Info_PullDevices_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullDevicesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InfoServer).PullDevices(m, &grpc.GenericServerStream[PullDevicesRequest, PullDevicesResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Info_PullDevicesServer = grpc.ServerStreamingServer[PullDevicesResponse]

// Info_ServiceDesc is the grpc.ServiceDesc for Info service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Info_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.info.Info",
	HandlerType: (*InfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDevices",
			Handler:    _Info_ListDevices_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PullDevices",
			Handler:       _Info_PullDevices_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "info/info.proto",
}
