// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: traits/booking.proto

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

// BookingApiClient is the client API for BookingApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookingApiClient interface {
	// list bookings for a given bookable
	ListBookings(ctx context.Context, in *ListBookingsRequest, opts ...grpc.CallOption) (*ListBookingsResponse, error)
	// check in an existing booking
	CheckInBooking(ctx context.Context, in *CheckInBookingRequest, opts ...grpc.CallOption) (*CheckInBookingResponse, error)
	// check out an existing booking
	CheckOutBooking(ctx context.Context, in *CheckOutBookingRequest, opts ...grpc.CallOption) (*CheckOutBookingResponse, error)
	// create a new booking
	CreateBooking(ctx context.Context, in *CreateBookingRequest, opts ...grpc.CallOption) (*CreateBookingResponse, error)
	// update an existing booking
	UpdateBooking(ctx context.Context, in *UpdateBookingRequest, opts ...grpc.CallOption) (*UpdateBookingResponse, error)
	// request updates to booking changes for a given bookable
	PullBookings(ctx context.Context, in *ListBookingsRequest, opts ...grpc.CallOption) (BookingApi_PullBookingsClient, error)
}

type bookingApiClient struct {
	cc grpc.ClientConnInterface
}

func NewBookingApiClient(cc grpc.ClientConnInterface) BookingApiClient {
	return &bookingApiClient{cc}
}

func (c *bookingApiClient) ListBookings(ctx context.Context, in *ListBookingsRequest, opts ...grpc.CallOption) (*ListBookingsResponse, error) {
	out := new(ListBookingsResponse)
	err := c.cc.Invoke(ctx, "/smartcore.traits.BookingApi/ListBookings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingApiClient) CheckInBooking(ctx context.Context, in *CheckInBookingRequest, opts ...grpc.CallOption) (*CheckInBookingResponse, error) {
	out := new(CheckInBookingResponse)
	err := c.cc.Invoke(ctx, "/smartcore.traits.BookingApi/CheckInBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingApiClient) CheckOutBooking(ctx context.Context, in *CheckOutBookingRequest, opts ...grpc.CallOption) (*CheckOutBookingResponse, error) {
	out := new(CheckOutBookingResponse)
	err := c.cc.Invoke(ctx, "/smartcore.traits.BookingApi/CheckOutBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingApiClient) CreateBooking(ctx context.Context, in *CreateBookingRequest, opts ...grpc.CallOption) (*CreateBookingResponse, error) {
	out := new(CreateBookingResponse)
	err := c.cc.Invoke(ctx, "/smartcore.traits.BookingApi/CreateBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingApiClient) UpdateBooking(ctx context.Context, in *UpdateBookingRequest, opts ...grpc.CallOption) (*UpdateBookingResponse, error) {
	out := new(UpdateBookingResponse)
	err := c.cc.Invoke(ctx, "/smartcore.traits.BookingApi/UpdateBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingApiClient) PullBookings(ctx context.Context, in *ListBookingsRequest, opts ...grpc.CallOption) (BookingApi_PullBookingsClient, error) {
	stream, err := c.cc.NewStream(ctx, &BookingApi_ServiceDesc.Streams[0], "/smartcore.traits.BookingApi/PullBookings", opts...)
	if err != nil {
		return nil, err
	}
	x := &bookingApiPullBookingsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BookingApi_PullBookingsClient interface {
	Recv() (*PullBookingsResponse, error)
	grpc.ClientStream
}

type bookingApiPullBookingsClient struct {
	grpc.ClientStream
}

func (x *bookingApiPullBookingsClient) Recv() (*PullBookingsResponse, error) {
	m := new(PullBookingsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BookingApiServer is the server API for BookingApi service.
// All implementations must embed UnimplementedBookingApiServer
// for forward compatibility
type BookingApiServer interface {
	// list bookings for a given bookable
	ListBookings(context.Context, *ListBookingsRequest) (*ListBookingsResponse, error)
	// check in an existing booking
	CheckInBooking(context.Context, *CheckInBookingRequest) (*CheckInBookingResponse, error)
	// check out an existing booking
	CheckOutBooking(context.Context, *CheckOutBookingRequest) (*CheckOutBookingResponse, error)
	// create a new booking
	CreateBooking(context.Context, *CreateBookingRequest) (*CreateBookingResponse, error)
	// update an existing booking
	UpdateBooking(context.Context, *UpdateBookingRequest) (*UpdateBookingResponse, error)
	// request updates to booking changes for a given bookable
	PullBookings(*ListBookingsRequest, BookingApi_PullBookingsServer) error
	mustEmbedUnimplementedBookingApiServer()
}

// UnimplementedBookingApiServer must be embedded to have forward compatible implementations.
type UnimplementedBookingApiServer struct {
}

func (UnimplementedBookingApiServer) ListBookings(context.Context, *ListBookingsRequest) (*ListBookingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBookings not implemented")
}
func (UnimplementedBookingApiServer) CheckInBooking(context.Context, *CheckInBookingRequest) (*CheckInBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckInBooking not implemented")
}
func (UnimplementedBookingApiServer) CheckOutBooking(context.Context, *CheckOutBookingRequest) (*CheckOutBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckOutBooking not implemented")
}
func (UnimplementedBookingApiServer) CreateBooking(context.Context, *CreateBookingRequest) (*CreateBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBooking not implemented")
}
func (UnimplementedBookingApiServer) UpdateBooking(context.Context, *UpdateBookingRequest) (*UpdateBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBooking not implemented")
}
func (UnimplementedBookingApiServer) PullBookings(*ListBookingsRequest, BookingApi_PullBookingsServer) error {
	return status.Errorf(codes.Unimplemented, "method PullBookings not implemented")
}
func (UnimplementedBookingApiServer) mustEmbedUnimplementedBookingApiServer() {}

// UnsafeBookingApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookingApiServer will
// result in compilation errors.
type UnsafeBookingApiServer interface {
	mustEmbedUnimplementedBookingApiServer()
}

func RegisterBookingApiServer(s grpc.ServiceRegistrar, srv BookingApiServer) {
	s.RegisterService(&BookingApi_ServiceDesc, srv)
}

func _BookingApi_ListBookings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBookingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingApiServer).ListBookings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.traits.BookingApi/ListBookings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingApiServer).ListBookings(ctx, req.(*ListBookingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingApi_CheckInBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckInBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingApiServer).CheckInBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.traits.BookingApi/CheckInBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingApiServer).CheckInBooking(ctx, req.(*CheckInBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingApi_CheckOutBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckOutBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingApiServer).CheckOutBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.traits.BookingApi/CheckOutBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingApiServer).CheckOutBooking(ctx, req.(*CheckOutBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingApi_CreateBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingApiServer).CreateBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.traits.BookingApi/CreateBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingApiServer).CreateBooking(ctx, req.(*CreateBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingApi_UpdateBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingApiServer).UpdateBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.traits.BookingApi/UpdateBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingApiServer).UpdateBooking(ctx, req.(*UpdateBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingApi_PullBookings_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListBookingsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BookingApiServer).PullBookings(m, &bookingApiPullBookingsServer{stream})
}

type BookingApi_PullBookingsServer interface {
	Send(*PullBookingsResponse) error
	grpc.ServerStream
}

type bookingApiPullBookingsServer struct {
	grpc.ServerStream
}

func (x *bookingApiPullBookingsServer) Send(m *PullBookingsResponse) error {
	return x.ServerStream.SendMsg(m)
}

// BookingApi_ServiceDesc is the grpc.ServiceDesc for BookingApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookingApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.traits.BookingApi",
	HandlerType: (*BookingApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListBookings",
			Handler:    _BookingApi_ListBookings_Handler,
		},
		{
			MethodName: "CheckInBooking",
			Handler:    _BookingApi_CheckInBooking_Handler,
		},
		{
			MethodName: "CheckOutBooking",
			Handler:    _BookingApi_CheckOutBooking_Handler,
		},
		{
			MethodName: "CreateBooking",
			Handler:    _BookingApi_CreateBooking_Handler,
		},
		{
			MethodName: "UpdateBooking",
			Handler:    _BookingApi_UpdateBooking_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PullBookings",
			Handler:       _BookingApi_PullBookings_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "traits/booking.proto",
}

// BookingInfoClient is the client API for BookingInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookingInfoClient interface {
	// Get information about how a named device implements Booking features
	DescribeBooking(ctx context.Context, in *DescribeBookingRequest, opts ...grpc.CallOption) (*BookingSupport, error)
}

type bookingInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewBookingInfoClient(cc grpc.ClientConnInterface) BookingInfoClient {
	return &bookingInfoClient{cc}
}

func (c *bookingInfoClient) DescribeBooking(ctx context.Context, in *DescribeBookingRequest, opts ...grpc.CallOption) (*BookingSupport, error) {
	out := new(BookingSupport)
	err := c.cc.Invoke(ctx, "/smartcore.traits.BookingInfo/DescribeBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookingInfoServer is the server API for BookingInfo service.
// All implementations must embed UnimplementedBookingInfoServer
// for forward compatibility
type BookingInfoServer interface {
	// Get information about how a named device implements Booking features
	DescribeBooking(context.Context, *DescribeBookingRequest) (*BookingSupport, error)
	mustEmbedUnimplementedBookingInfoServer()
}

// UnimplementedBookingInfoServer must be embedded to have forward compatible implementations.
type UnimplementedBookingInfoServer struct {
}

func (UnimplementedBookingInfoServer) DescribeBooking(context.Context, *DescribeBookingRequest) (*BookingSupport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeBooking not implemented")
}
func (UnimplementedBookingInfoServer) mustEmbedUnimplementedBookingInfoServer() {}

// UnsafeBookingInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookingInfoServer will
// result in compilation errors.
type UnsafeBookingInfoServer interface {
	mustEmbedUnimplementedBookingInfoServer()
}

func RegisterBookingInfoServer(s grpc.ServiceRegistrar, srv BookingInfoServer) {
	s.RegisterService(&BookingInfo_ServiceDesc, srv)
}

func _BookingInfo_DescribeBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingInfoServer).DescribeBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.traits.BookingInfo/DescribeBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingInfoServer).DescribeBooking(ctx, req.(*DescribeBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookingInfo_ServiceDesc is the grpc.ServiceDesc for BookingInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookingInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.traits.BookingInfo",
	HandlerType: (*BookingInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeBooking",
			Handler:    _BookingInfo_DescribeBooking_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "traits/booking.proto",
}
