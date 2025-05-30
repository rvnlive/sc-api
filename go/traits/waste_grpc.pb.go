// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: traits/waste.proto

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
	WasteApi_ListWasteRecords_FullMethodName = "/smartcore.traits.WasteApi/ListWasteRecords"
	WasteApi_PullWasteRecords_FullMethodName = "/smartcore.traits.WasteApi/PullWasteRecords"
)

// WasteApiClient is the client API for WasteApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// WasteApi exposes details about units of waste that have been produced by the building
type WasteApiClient interface {
	ListWasteRecords(ctx context.Context, in *ListWasteRecordsRequest, opts ...grpc.CallOption) (*ListWasteRecordsResponse, error)
	PullWasteRecords(ctx context.Context, in *PullWasteRecordsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PullWasteRecordsResponse], error)
}

type wasteApiClient struct {
	cc grpc.ClientConnInterface
}

func NewWasteApiClient(cc grpc.ClientConnInterface) WasteApiClient {
	return &wasteApiClient{cc}
}

func (c *wasteApiClient) ListWasteRecords(ctx context.Context, in *ListWasteRecordsRequest, opts ...grpc.CallOption) (*ListWasteRecordsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListWasteRecordsResponse)
	err := c.cc.Invoke(ctx, WasteApi_ListWasteRecords_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wasteApiClient) PullWasteRecords(ctx context.Context, in *PullWasteRecordsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PullWasteRecordsResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &WasteApi_ServiceDesc.Streams[0], WasteApi_PullWasteRecords_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[PullWasteRecordsRequest, PullWasteRecordsResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type WasteApi_PullWasteRecordsClient = grpc.ServerStreamingClient[PullWasteRecordsResponse]

// WasteApiServer is the server API for WasteApi service.
// All implementations must embed UnimplementedWasteApiServer
// for forward compatibility.
//
// WasteApi exposes details about units of waste that have been produced by the building
type WasteApiServer interface {
	ListWasteRecords(context.Context, *ListWasteRecordsRequest) (*ListWasteRecordsResponse, error)
	PullWasteRecords(*PullWasteRecordsRequest, grpc.ServerStreamingServer[PullWasteRecordsResponse]) error
	mustEmbedUnimplementedWasteApiServer()
}

// UnimplementedWasteApiServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWasteApiServer struct{}

func (UnimplementedWasteApiServer) ListWasteRecords(context.Context, *ListWasteRecordsRequest) (*ListWasteRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWasteRecords not implemented")
}
func (UnimplementedWasteApiServer) PullWasteRecords(*PullWasteRecordsRequest, grpc.ServerStreamingServer[PullWasteRecordsResponse]) error {
	return status.Errorf(codes.Unimplemented, "method PullWasteRecords not implemented")
}
func (UnimplementedWasteApiServer) mustEmbedUnimplementedWasteApiServer() {}
func (UnimplementedWasteApiServer) testEmbeddedByValue()                  {}

// UnsafeWasteApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WasteApiServer will
// result in compilation errors.
type UnsafeWasteApiServer interface {
	mustEmbedUnimplementedWasteApiServer()
}

func RegisterWasteApiServer(s grpc.ServiceRegistrar, srv WasteApiServer) {
	// If the following call pancis, it indicates UnimplementedWasteApiServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WasteApi_ServiceDesc, srv)
}

func _WasteApi_ListWasteRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWasteRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WasteApiServer).ListWasteRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WasteApi_ListWasteRecords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WasteApiServer).ListWasteRecords(ctx, req.(*ListWasteRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WasteApi_PullWasteRecords_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullWasteRecordsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WasteApiServer).PullWasteRecords(m, &grpc.GenericServerStream[PullWasteRecordsRequest, PullWasteRecordsResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type WasteApi_PullWasteRecordsServer = grpc.ServerStreamingServer[PullWasteRecordsResponse]

// WasteApi_ServiceDesc is the grpc.ServiceDesc for WasteApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WasteApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.traits.WasteApi",
	HandlerType: (*WasteApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListWasteRecords",
			Handler:    _WasteApi_ListWasteRecords_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PullWasteRecords",
			Handler:       _WasteApi_PullWasteRecords_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "traits/waste.proto",
}

const (
	WasteInfo_DescribeWasteRecord_FullMethodName = "/smartcore.traits.WasteInfo/DescribeWasteRecord"
)

// WasteInfoClient is the client API for WasteInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WasteInfoClient interface {
	DescribeWasteRecord(ctx context.Context, in *DescribeWasteRecordRequest, opts ...grpc.CallOption) (*WasteRecordSupport, error)
}

type wasteInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewWasteInfoClient(cc grpc.ClientConnInterface) WasteInfoClient {
	return &wasteInfoClient{cc}
}

func (c *wasteInfoClient) DescribeWasteRecord(ctx context.Context, in *DescribeWasteRecordRequest, opts ...grpc.CallOption) (*WasteRecordSupport, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WasteRecordSupport)
	err := c.cc.Invoke(ctx, WasteInfo_DescribeWasteRecord_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WasteInfoServer is the server API for WasteInfo service.
// All implementations must embed UnimplementedWasteInfoServer
// for forward compatibility.
type WasteInfoServer interface {
	DescribeWasteRecord(context.Context, *DescribeWasteRecordRequest) (*WasteRecordSupport, error)
	mustEmbedUnimplementedWasteInfoServer()
}

// UnimplementedWasteInfoServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWasteInfoServer struct{}

func (UnimplementedWasteInfoServer) DescribeWasteRecord(context.Context, *DescribeWasteRecordRequest) (*WasteRecordSupport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeWasteRecord not implemented")
}
func (UnimplementedWasteInfoServer) mustEmbedUnimplementedWasteInfoServer() {}
func (UnimplementedWasteInfoServer) testEmbeddedByValue()                   {}

// UnsafeWasteInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WasteInfoServer will
// result in compilation errors.
type UnsafeWasteInfoServer interface {
	mustEmbedUnimplementedWasteInfoServer()
}

func RegisterWasteInfoServer(s grpc.ServiceRegistrar, srv WasteInfoServer) {
	// If the following call pancis, it indicates UnimplementedWasteInfoServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WasteInfo_ServiceDesc, srv)
}

func _WasteInfo_DescribeWasteRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeWasteRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WasteInfoServer).DescribeWasteRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WasteInfo_DescribeWasteRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WasteInfoServer).DescribeWasteRecord(ctx, req.(*DescribeWasteRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WasteInfo_ServiceDesc is the grpc.ServiceDesc for WasteInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WasteInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.traits.WasteInfo",
	HandlerType: (*WasteInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeWasteRecord",
			Handler:    _WasteInfo_DescribeWasteRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "traits/waste.proto",
}
