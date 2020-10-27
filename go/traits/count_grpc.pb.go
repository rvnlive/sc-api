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

// CountApiClient is the client API for CountApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CountApiClient interface {
	// Get the current count state
	GetCount(ctx context.Context, in *GetCountRequest, opts ...grpc.CallOption) (*Count, error)
	// Reset the counts to 0 and update the reset time.
	ResetCount(ctx context.Context, in *ResetCountRequest, opts ...grpc.CallOption) (*Count, error)
	// Update one or more properties of the count.
	UpdateCount(ctx context.Context, in *UpdateCountRequest, opts ...grpc.CallOption) (*Count, error)
	// Subscribe to changes to the count.
	PullCounts(ctx context.Context, in *PullCountsRequest, opts ...grpc.CallOption) (CountApi_PullCountsClient, error)
}

type countApiClient struct {
	cc grpc.ClientConnInterface
}

func NewCountApiClient(cc grpc.ClientConnInterface) CountApiClient {
	return &countApiClient{cc}
}

func (c *countApiClient) GetCount(ctx context.Context, in *GetCountRequest, opts ...grpc.CallOption) (*Count, error) {
	out := new(Count)
	err := c.cc.Invoke(ctx, "/smartcore.traits.CountApi/GetCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *countApiClient) ResetCount(ctx context.Context, in *ResetCountRequest, opts ...grpc.CallOption) (*Count, error) {
	out := new(Count)
	err := c.cc.Invoke(ctx, "/smartcore.traits.CountApi/ResetCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *countApiClient) UpdateCount(ctx context.Context, in *UpdateCountRequest, opts ...grpc.CallOption) (*Count, error) {
	out := new(Count)
	err := c.cc.Invoke(ctx, "/smartcore.traits.CountApi/UpdateCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *countApiClient) PullCounts(ctx context.Context, in *PullCountsRequest, opts ...grpc.CallOption) (CountApi_PullCountsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CountApi_serviceDesc.Streams[0], "/smartcore.traits.CountApi/PullCounts", opts...)
	if err != nil {
		return nil, err
	}
	x := &countApiPullCountsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CountApi_PullCountsClient interface {
	Recv() (*PullCountsResponse, error)
	grpc.ClientStream
}

type countApiPullCountsClient struct {
	grpc.ClientStream
}

func (x *countApiPullCountsClient) Recv() (*PullCountsResponse, error) {
	m := new(PullCountsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CountApiServer is the server API for CountApi service.
// All implementations must embed UnimplementedCountApiServer
// for forward compatibility
type CountApiServer interface {
	// Get the current count state
	GetCount(context.Context, *GetCountRequest) (*Count, error)
	// Reset the counts to 0 and update the reset time.
	ResetCount(context.Context, *ResetCountRequest) (*Count, error)
	// Update one or more properties of the count.
	UpdateCount(context.Context, *UpdateCountRequest) (*Count, error)
	// Subscribe to changes to the count.
	PullCounts(*PullCountsRequest, CountApi_PullCountsServer) error
	mustEmbedUnimplementedCountApiServer()
}

// UnimplementedCountApiServer must be embedded to have forward compatible implementations.
type UnimplementedCountApiServer struct {
}

func (*UnimplementedCountApiServer) GetCount(context.Context, *GetCountRequest) (*Count, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCount not implemented")
}
func (*UnimplementedCountApiServer) ResetCount(context.Context, *ResetCountRequest) (*Count, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetCount not implemented")
}
func (*UnimplementedCountApiServer) UpdateCount(context.Context, *UpdateCountRequest) (*Count, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCount not implemented")
}
func (*UnimplementedCountApiServer) PullCounts(*PullCountsRequest, CountApi_PullCountsServer) error {
	return status.Errorf(codes.Unimplemented, "method PullCounts not implemented")
}
func (*UnimplementedCountApiServer) mustEmbedUnimplementedCountApiServer() {}

func RegisterCountApiServer(s *grpc.Server, srv CountApiServer) {
	s.RegisterService(&_CountApi_serviceDesc, srv)
}

func _CountApi_GetCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountApiServer).GetCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.traits.CountApi/GetCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountApiServer).GetCount(ctx, req.(*GetCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CountApi_ResetCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountApiServer).ResetCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.traits.CountApi/ResetCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountApiServer).ResetCount(ctx, req.(*ResetCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CountApi_UpdateCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountApiServer).UpdateCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.traits.CountApi/UpdateCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountApiServer).UpdateCount(ctx, req.(*UpdateCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CountApi_PullCounts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullCountsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CountApiServer).PullCounts(m, &countApiPullCountsServer{stream})
}

type CountApi_PullCountsServer interface {
	Send(*PullCountsResponse) error
	grpc.ServerStream
}

type countApiPullCountsServer struct {
	grpc.ServerStream
}

func (x *countApiPullCountsServer) Send(m *PullCountsResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _CountApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.traits.CountApi",
	HandlerType: (*CountApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCount",
			Handler:    _CountApi_GetCount_Handler,
		},
		{
			MethodName: "ResetCount",
			Handler:    _CountApi_ResetCount_Handler,
		},
		{
			MethodName: "UpdateCount",
			Handler:    _CountApi_UpdateCount_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PullCounts",
			Handler:       _CountApi_PullCounts_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "traits/count.proto",
}

// CountInfoClient is the client API for CountInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CountInfoClient interface {
	// Get information about how a named device implements Count features
	DescribeCount(ctx context.Context, in *DescribeCountRequest, opts ...grpc.CallOption) (*CountSupport, error)
}

type countInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewCountInfoClient(cc grpc.ClientConnInterface) CountInfoClient {
	return &countInfoClient{cc}
}

func (c *countInfoClient) DescribeCount(ctx context.Context, in *DescribeCountRequest, opts ...grpc.CallOption) (*CountSupport, error) {
	out := new(CountSupport)
	err := c.cc.Invoke(ctx, "/smartcore.traits.CountInfo/DescribeCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CountInfoServer is the server API for CountInfo service.
// All implementations must embed UnimplementedCountInfoServer
// for forward compatibility
type CountInfoServer interface {
	// Get information about how a named device implements Count features
	DescribeCount(context.Context, *DescribeCountRequest) (*CountSupport, error)
	mustEmbedUnimplementedCountInfoServer()
}

// UnimplementedCountInfoServer must be embedded to have forward compatible implementations.
type UnimplementedCountInfoServer struct {
}

func (*UnimplementedCountInfoServer) DescribeCount(context.Context, *DescribeCountRequest) (*CountSupport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeCount not implemented")
}
func (*UnimplementedCountInfoServer) mustEmbedUnimplementedCountInfoServer() {}

func RegisterCountInfoServer(s *grpc.Server, srv CountInfoServer) {
	s.RegisterService(&_CountInfo_serviceDesc, srv)
}

func _CountInfo_DescribeCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountInfoServer).DescribeCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.traits.CountInfo/DescribeCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountInfoServer).DescribeCount(ctx, req.(*DescribeCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CountInfo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.traits.CountInfo",
	HandlerType: (*CountInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeCount",
			Handler:    _CountInfo_DescribeCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "traits/count.proto",
}
