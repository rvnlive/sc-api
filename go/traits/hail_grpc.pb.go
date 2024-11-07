// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: traits/hail.proto

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
	HailApi_CreateHail_FullMethodName = "/smartcore.traits.HailApi/CreateHail"
	HailApi_GetHail_FullMethodName    = "/smartcore.traits.HailApi/GetHail"
	HailApi_UpdateHail_FullMethodName = "/smartcore.traits.HailApi/UpdateHail"
	HailApi_DeleteHail_FullMethodName = "/smartcore.traits.HailApi/DeleteHail"
	HailApi_PullHail_FullMethodName   = "/smartcore.traits.HailApi/PullHail"
	HailApi_ListHails_FullMethodName  = "/smartcore.traits.HailApi/ListHails"
	HailApi_PullHails_FullMethodName  = "/smartcore.traits.HailApi/PullHails"
)

// HailApiClient is the client API for HailApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// HailApi describes devices that may be asked to transport passengers from one location to another.
// To interact with these devices you create a hail, a description of the origin and destination for your journey.
// The device keeps you updated on the preparation or progress of your journey by updating the fields in Hail.
//
// HailApi applies to devices like lifts or taxis, where the transportation is initiated by the caller. It does not
// apply to scheduled/regular transportation - like busses or escalators - where the presence of the passenger is not
// required.
type HailApiClient interface {
	// Create a new hail against the device.
	// Creating a hail instructs the device that a passenger would like to be transported from one location to another.
	// The response from this request may include additional location information than that supplied.
	// See Hail.origin for more details.
	//
	// If Hail.id is specified and is used by an existing hail the device will return an error.
	// If Hail.id is specified the device may return an error indicating that client supplied ids are not supported.
	CreateHail(ctx context.Context, in *CreateHailRequest, opts ...grpc.CallOption) (*Hail, error)
	// Retrieve details on a specific hail by id as returned by CreateHail.
	GetHail(ctx context.Context, in *GetHailRequest, opts ...grpc.CallOption) (*Hail, error)
	// Update a hail by id.
	// Hail.id must be present.
	UpdateHail(ctx context.Context, in *UpdateHailRequest, opts ...grpc.CallOption) (*Hail, error)
	// Delete a hail by id.
	DeleteHail(ctx context.Context, in *DeleteHailRequest, opts ...grpc.CallOption) (*DeleteHailResponse, error)
	// Subscribe to changes in a specific hail.
	PullHail(ctx context.Context, in *PullHailRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PullHailResponse], error)
	// List known hails.
	ListHails(ctx context.Context, in *ListHailsRequest, opts ...grpc.CallOption) (*ListHailsResponse, error)
	// Subscribe to changes in the known hails.
	PullHails(ctx context.Context, in *PullHailsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PullHailsResponse], error)
}

type hailApiClient struct {
	cc grpc.ClientConnInterface
}

func NewHailApiClient(cc grpc.ClientConnInterface) HailApiClient {
	return &hailApiClient{cc}
}

func (c *hailApiClient) CreateHail(ctx context.Context, in *CreateHailRequest, opts ...grpc.CallOption) (*Hail, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Hail)
	err := c.cc.Invoke(ctx, HailApi_CreateHail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hailApiClient) GetHail(ctx context.Context, in *GetHailRequest, opts ...grpc.CallOption) (*Hail, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Hail)
	err := c.cc.Invoke(ctx, HailApi_GetHail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hailApiClient) UpdateHail(ctx context.Context, in *UpdateHailRequest, opts ...grpc.CallOption) (*Hail, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Hail)
	err := c.cc.Invoke(ctx, HailApi_UpdateHail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hailApiClient) DeleteHail(ctx context.Context, in *DeleteHailRequest, opts ...grpc.CallOption) (*DeleteHailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteHailResponse)
	err := c.cc.Invoke(ctx, HailApi_DeleteHail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hailApiClient) PullHail(ctx context.Context, in *PullHailRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PullHailResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &HailApi_ServiceDesc.Streams[0], HailApi_PullHail_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[PullHailRequest, PullHailResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type HailApi_PullHailClient = grpc.ServerStreamingClient[PullHailResponse]

func (c *hailApiClient) ListHails(ctx context.Context, in *ListHailsRequest, opts ...grpc.CallOption) (*ListHailsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListHailsResponse)
	err := c.cc.Invoke(ctx, HailApi_ListHails_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hailApiClient) PullHails(ctx context.Context, in *PullHailsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PullHailsResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &HailApi_ServiceDesc.Streams[1], HailApi_PullHails_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[PullHailsRequest, PullHailsResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type HailApi_PullHailsClient = grpc.ServerStreamingClient[PullHailsResponse]

// HailApiServer is the server API for HailApi service.
// All implementations must embed UnimplementedHailApiServer
// for forward compatibility.
//
// HailApi describes devices that may be asked to transport passengers from one location to another.
// To interact with these devices you create a hail, a description of the origin and destination for your journey.
// The device keeps you updated on the preparation or progress of your journey by updating the fields in Hail.
//
// HailApi applies to devices like lifts or taxis, where the transportation is initiated by the caller. It does not
// apply to scheduled/regular transportation - like busses or escalators - where the presence of the passenger is not
// required.
type HailApiServer interface {
	// Create a new hail against the device.
	// Creating a hail instructs the device that a passenger would like to be transported from one location to another.
	// The response from this request may include additional location information than that supplied.
	// See Hail.origin for more details.
	//
	// If Hail.id is specified and is used by an existing hail the device will return an error.
	// If Hail.id is specified the device may return an error indicating that client supplied ids are not supported.
	CreateHail(context.Context, *CreateHailRequest) (*Hail, error)
	// Retrieve details on a specific hail by id as returned by CreateHail.
	GetHail(context.Context, *GetHailRequest) (*Hail, error)
	// Update a hail by id.
	// Hail.id must be present.
	UpdateHail(context.Context, *UpdateHailRequest) (*Hail, error)
	// Delete a hail by id.
	DeleteHail(context.Context, *DeleteHailRequest) (*DeleteHailResponse, error)
	// Subscribe to changes in a specific hail.
	PullHail(*PullHailRequest, grpc.ServerStreamingServer[PullHailResponse]) error
	// List known hails.
	ListHails(context.Context, *ListHailsRequest) (*ListHailsResponse, error)
	// Subscribe to changes in the known hails.
	PullHails(*PullHailsRequest, grpc.ServerStreamingServer[PullHailsResponse]) error
	mustEmbedUnimplementedHailApiServer()
}

// UnimplementedHailApiServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHailApiServer struct{}

func (UnimplementedHailApiServer) CreateHail(context.Context, *CreateHailRequest) (*Hail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHail not implemented")
}
func (UnimplementedHailApiServer) GetHail(context.Context, *GetHailRequest) (*Hail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHail not implemented")
}
func (UnimplementedHailApiServer) UpdateHail(context.Context, *UpdateHailRequest) (*Hail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHail not implemented")
}
func (UnimplementedHailApiServer) DeleteHail(context.Context, *DeleteHailRequest) (*DeleteHailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHail not implemented")
}
func (UnimplementedHailApiServer) PullHail(*PullHailRequest, grpc.ServerStreamingServer[PullHailResponse]) error {
	return status.Errorf(codes.Unimplemented, "method PullHail not implemented")
}
func (UnimplementedHailApiServer) ListHails(context.Context, *ListHailsRequest) (*ListHailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHails not implemented")
}
func (UnimplementedHailApiServer) PullHails(*PullHailsRequest, grpc.ServerStreamingServer[PullHailsResponse]) error {
	return status.Errorf(codes.Unimplemented, "method PullHails not implemented")
}
func (UnimplementedHailApiServer) mustEmbedUnimplementedHailApiServer() {}
func (UnimplementedHailApiServer) testEmbeddedByValue()                 {}

// UnsafeHailApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HailApiServer will
// result in compilation errors.
type UnsafeHailApiServer interface {
	mustEmbedUnimplementedHailApiServer()
}

func RegisterHailApiServer(s grpc.ServiceRegistrar, srv HailApiServer) {
	// If the following call pancis, it indicates UnimplementedHailApiServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&HailApi_ServiceDesc, srv)
}

func _HailApi_CreateHail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HailApiServer).CreateHail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HailApi_CreateHail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HailApiServer).CreateHail(ctx, req.(*CreateHailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HailApi_GetHail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HailApiServer).GetHail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HailApi_GetHail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HailApiServer).GetHail(ctx, req.(*GetHailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HailApi_UpdateHail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HailApiServer).UpdateHail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HailApi_UpdateHail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HailApiServer).UpdateHail(ctx, req.(*UpdateHailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HailApi_DeleteHail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HailApiServer).DeleteHail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HailApi_DeleteHail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HailApiServer).DeleteHail(ctx, req.(*DeleteHailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HailApi_PullHail_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullHailRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HailApiServer).PullHail(m, &grpc.GenericServerStream[PullHailRequest, PullHailResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type HailApi_PullHailServer = grpc.ServerStreamingServer[PullHailResponse]

func _HailApi_ListHails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListHailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HailApiServer).ListHails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HailApi_ListHails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HailApiServer).ListHails(ctx, req.(*ListHailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HailApi_PullHails_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullHailsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HailApiServer).PullHails(m, &grpc.GenericServerStream[PullHailsRequest, PullHailsResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type HailApi_PullHailsServer = grpc.ServerStreamingServer[PullHailsResponse]

// HailApi_ServiceDesc is the grpc.ServiceDesc for HailApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HailApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.traits.HailApi",
	HandlerType: (*HailApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateHail",
			Handler:    _HailApi_CreateHail_Handler,
		},
		{
			MethodName: "GetHail",
			Handler:    _HailApi_GetHail_Handler,
		},
		{
			MethodName: "UpdateHail",
			Handler:    _HailApi_UpdateHail_Handler,
		},
		{
			MethodName: "DeleteHail",
			Handler:    _HailApi_DeleteHail_Handler,
		},
		{
			MethodName: "ListHails",
			Handler:    _HailApi_ListHails_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PullHail",
			Handler:       _HailApi_PullHail_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PullHails",
			Handler:       _HailApi_PullHails_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "traits/hail.proto",
}

// HailInfoClient is the client API for HailInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HailInfoClient interface {
}

type hailInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewHailInfoClient(cc grpc.ClientConnInterface) HailInfoClient {
	return &hailInfoClient{cc}
}

// HailInfoServer is the server API for HailInfo service.
// All implementations must embed UnimplementedHailInfoServer
// for forward compatibility.
type HailInfoServer interface {
	mustEmbedUnimplementedHailInfoServer()
}

// UnimplementedHailInfoServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHailInfoServer struct{}

func (UnimplementedHailInfoServer) mustEmbedUnimplementedHailInfoServer() {}
func (UnimplementedHailInfoServer) testEmbeddedByValue()                  {}

// UnsafeHailInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HailInfoServer will
// result in compilation errors.
type UnsafeHailInfoServer interface {
	mustEmbedUnimplementedHailInfoServer()
}

func RegisterHailInfoServer(s grpc.ServiceRegistrar, srv HailInfoServer) {
	// If the following call pancis, it indicates UnimplementedHailInfoServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&HailInfo_ServiceDesc, srv)
}

// HailInfo_ServiceDesc is the grpc.ServiceDesc for HailInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HailInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.traits.HailInfo",
	HandlerType: (*HailInfoServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "traits/hail.proto",
}
