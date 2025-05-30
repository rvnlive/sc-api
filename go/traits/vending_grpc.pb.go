// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: traits/vending.proto

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
	VendingApi_ListConsumables_FullMethodName = "/smartcore.traits.VendingApi/ListConsumables"
	VendingApi_PullConsumables_FullMethodName = "/smartcore.traits.VendingApi/PullConsumables"
	VendingApi_GetStock_FullMethodName        = "/smartcore.traits.VendingApi/GetStock"
	VendingApi_UpdateStock_FullMethodName     = "/smartcore.traits.VendingApi/UpdateStock"
	VendingApi_PullStock_FullMethodName       = "/smartcore.traits.VendingApi/PullStock"
	VendingApi_ListInventory_FullMethodName   = "/smartcore.traits.VendingApi/ListInventory"
	VendingApi_PullInventory_FullMethodName   = "/smartcore.traits.VendingApi/PullInventory"
	VendingApi_Dispense_FullMethodName        = "/smartcore.traits.VendingApi/Dispense"
	VendingApi_StopDispense_FullMethodName    = "/smartcore.traits.VendingApi/StopDispense"
)

// VendingApiClient is the client API for VendingApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// VendingApi describes devices that hold a stock and can dispense their consumables on command.
// Applies to devices like vending machines, water taps, coffee machines, and pet food dispensers.
//
// The VendingApi encodes two different resources, the consumables and the stock.
// Consumables represent information about the products that are available from this device.
// Stock is a measure of how many/much of a consumables is currently available.
// When referring to the stock of all consumables, this trait uses the term inventory.
//
// Quantities, both available and dispensed, can be either discrete items like apples, or continuous like water.
// Some consumables are both discrete and continuous, you can have 0.3L or 1 cup of water.
type VendingApiClient interface {
	// List the consumables available from this device.
	// This has no information on stock, only descriptive information about the physical item.
	ListConsumables(ctx context.Context, in *ListConsumablesRequest, opts ...grpc.CallOption) (*ListConsumablesResponse, error)
	// Subscribe to change in consumables available from this device.
	PullConsumables(ctx context.Context, in *PullConsumablesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PullConsumablesResponse], error)
	// Get the current stock levels for a consumable.
	// If the consumable is unknown will return NotFound.
	// If the consumable is out of stock, returns a Stock with a zero remaining quantity.
	GetStock(ctx context.Context, in *GetStockRequest, opts ...grpc.CallOption) (*Consumable_Stock, error)
	// Update the quantity of a consumable in the device.
	//
	// This is unlikely to actually create physical consumables out of thin air, but is useful if the device doesn't have
	// internal sensors monitoring the physical quantity of its consumables. For example to let a cookie dispenser know
	// that it was refilled with 100 new tasty cookies.
	UpdateStock(ctx context.Context, in *UpdateStockRequest, opts ...grpc.CallOption) (*Consumable_Stock, error)
	// Subscribe to changes in the stock levels for a consumable.
	PullStock(ctx context.Context, in *PullStockRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PullStockResponse], error)
	// List the stock for each consumable in the device.
	// The device should return zero-quantity stock for consumables that are out-of-stock.
	ListInventory(ctx context.Context, in *ListInventoryRequest, opts ...grpc.CallOption) (*ListInventoryResponse, error)
	// Subscribe to changes in consumable stock.
	PullInventory(ctx context.Context, in *PullInventoryRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PullInventoryResponse], error)
	// Trigger the dispensing of some quantity of consumables.
	// Devices should return only stock items that are mentioned in the request.
	Dispense(ctx context.Context, in *DispenseRequest, opts ...grpc.CallOption) (*Consumable_Stock, error)
	// Stop the dispensing of the mentioned consumables.
	// Devices should return only stock items that are mentioned in the request.
	StopDispense(ctx context.Context, in *StopDispenseRequest, opts ...grpc.CallOption) (*Consumable_Stock, error)
}

type vendingApiClient struct {
	cc grpc.ClientConnInterface
}

func NewVendingApiClient(cc grpc.ClientConnInterface) VendingApiClient {
	return &vendingApiClient{cc}
}

func (c *vendingApiClient) ListConsumables(ctx context.Context, in *ListConsumablesRequest, opts ...grpc.CallOption) (*ListConsumablesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListConsumablesResponse)
	err := c.cc.Invoke(ctx, VendingApi_ListConsumables_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendingApiClient) PullConsumables(ctx context.Context, in *PullConsumablesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PullConsumablesResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &VendingApi_ServiceDesc.Streams[0], VendingApi_PullConsumables_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[PullConsumablesRequest, PullConsumablesResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type VendingApi_PullConsumablesClient = grpc.ServerStreamingClient[PullConsumablesResponse]

func (c *vendingApiClient) GetStock(ctx context.Context, in *GetStockRequest, opts ...grpc.CallOption) (*Consumable_Stock, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Consumable_Stock)
	err := c.cc.Invoke(ctx, VendingApi_GetStock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendingApiClient) UpdateStock(ctx context.Context, in *UpdateStockRequest, opts ...grpc.CallOption) (*Consumable_Stock, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Consumable_Stock)
	err := c.cc.Invoke(ctx, VendingApi_UpdateStock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendingApiClient) PullStock(ctx context.Context, in *PullStockRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PullStockResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &VendingApi_ServiceDesc.Streams[1], VendingApi_PullStock_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[PullStockRequest, PullStockResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type VendingApi_PullStockClient = grpc.ServerStreamingClient[PullStockResponse]

func (c *vendingApiClient) ListInventory(ctx context.Context, in *ListInventoryRequest, opts ...grpc.CallOption) (*ListInventoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListInventoryResponse)
	err := c.cc.Invoke(ctx, VendingApi_ListInventory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendingApiClient) PullInventory(ctx context.Context, in *PullInventoryRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PullInventoryResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &VendingApi_ServiceDesc.Streams[2], VendingApi_PullInventory_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[PullInventoryRequest, PullInventoryResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type VendingApi_PullInventoryClient = grpc.ServerStreamingClient[PullInventoryResponse]

func (c *vendingApiClient) Dispense(ctx context.Context, in *DispenseRequest, opts ...grpc.CallOption) (*Consumable_Stock, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Consumable_Stock)
	err := c.cc.Invoke(ctx, VendingApi_Dispense_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendingApiClient) StopDispense(ctx context.Context, in *StopDispenseRequest, opts ...grpc.CallOption) (*Consumable_Stock, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Consumable_Stock)
	err := c.cc.Invoke(ctx, VendingApi_StopDispense_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VendingApiServer is the server API for VendingApi service.
// All implementations must embed UnimplementedVendingApiServer
// for forward compatibility.
//
// VendingApi describes devices that hold a stock and can dispense their consumables on command.
// Applies to devices like vending machines, water taps, coffee machines, and pet food dispensers.
//
// The VendingApi encodes two different resources, the consumables and the stock.
// Consumables represent information about the products that are available from this device.
// Stock is a measure of how many/much of a consumables is currently available.
// When referring to the stock of all consumables, this trait uses the term inventory.
//
// Quantities, both available and dispensed, can be either discrete items like apples, or continuous like water.
// Some consumables are both discrete and continuous, you can have 0.3L or 1 cup of water.
type VendingApiServer interface {
	// List the consumables available from this device.
	// This has no information on stock, only descriptive information about the physical item.
	ListConsumables(context.Context, *ListConsumablesRequest) (*ListConsumablesResponse, error)
	// Subscribe to change in consumables available from this device.
	PullConsumables(*PullConsumablesRequest, grpc.ServerStreamingServer[PullConsumablesResponse]) error
	// Get the current stock levels for a consumable.
	// If the consumable is unknown will return NotFound.
	// If the consumable is out of stock, returns a Stock with a zero remaining quantity.
	GetStock(context.Context, *GetStockRequest) (*Consumable_Stock, error)
	// Update the quantity of a consumable in the device.
	//
	// This is unlikely to actually create physical consumables out of thin air, but is useful if the device doesn't have
	// internal sensors monitoring the physical quantity of its consumables. For example to let a cookie dispenser know
	// that it was refilled with 100 new tasty cookies.
	UpdateStock(context.Context, *UpdateStockRequest) (*Consumable_Stock, error)
	// Subscribe to changes in the stock levels for a consumable.
	PullStock(*PullStockRequest, grpc.ServerStreamingServer[PullStockResponse]) error
	// List the stock for each consumable in the device.
	// The device should return zero-quantity stock for consumables that are out-of-stock.
	ListInventory(context.Context, *ListInventoryRequest) (*ListInventoryResponse, error)
	// Subscribe to changes in consumable stock.
	PullInventory(*PullInventoryRequest, grpc.ServerStreamingServer[PullInventoryResponse]) error
	// Trigger the dispensing of some quantity of consumables.
	// Devices should return only stock items that are mentioned in the request.
	Dispense(context.Context, *DispenseRequest) (*Consumable_Stock, error)
	// Stop the dispensing of the mentioned consumables.
	// Devices should return only stock items that are mentioned in the request.
	StopDispense(context.Context, *StopDispenseRequest) (*Consumable_Stock, error)
	mustEmbedUnimplementedVendingApiServer()
}

// UnimplementedVendingApiServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedVendingApiServer struct{}

func (UnimplementedVendingApiServer) ListConsumables(context.Context, *ListConsumablesRequest) (*ListConsumablesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConsumables not implemented")
}
func (UnimplementedVendingApiServer) PullConsumables(*PullConsumablesRequest, grpc.ServerStreamingServer[PullConsumablesResponse]) error {
	return status.Errorf(codes.Unimplemented, "method PullConsumables not implemented")
}
func (UnimplementedVendingApiServer) GetStock(context.Context, *GetStockRequest) (*Consumable_Stock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStock not implemented")
}
func (UnimplementedVendingApiServer) UpdateStock(context.Context, *UpdateStockRequest) (*Consumable_Stock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStock not implemented")
}
func (UnimplementedVendingApiServer) PullStock(*PullStockRequest, grpc.ServerStreamingServer[PullStockResponse]) error {
	return status.Errorf(codes.Unimplemented, "method PullStock not implemented")
}
func (UnimplementedVendingApiServer) ListInventory(context.Context, *ListInventoryRequest) (*ListInventoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInventory not implemented")
}
func (UnimplementedVendingApiServer) PullInventory(*PullInventoryRequest, grpc.ServerStreamingServer[PullInventoryResponse]) error {
	return status.Errorf(codes.Unimplemented, "method PullInventory not implemented")
}
func (UnimplementedVendingApiServer) Dispense(context.Context, *DispenseRequest) (*Consumable_Stock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dispense not implemented")
}
func (UnimplementedVendingApiServer) StopDispense(context.Context, *StopDispenseRequest) (*Consumable_Stock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopDispense not implemented")
}
func (UnimplementedVendingApiServer) mustEmbedUnimplementedVendingApiServer() {}
func (UnimplementedVendingApiServer) testEmbeddedByValue()                    {}

// UnsafeVendingApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VendingApiServer will
// result in compilation errors.
type UnsafeVendingApiServer interface {
	mustEmbedUnimplementedVendingApiServer()
}

func RegisterVendingApiServer(s grpc.ServiceRegistrar, srv VendingApiServer) {
	// If the following call pancis, it indicates UnimplementedVendingApiServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&VendingApi_ServiceDesc, srv)
}

func _VendingApi_ListConsumables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListConsumablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VendingApiServer).ListConsumables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VendingApi_ListConsumables_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VendingApiServer).ListConsumables(ctx, req.(*ListConsumablesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VendingApi_PullConsumables_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullConsumablesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VendingApiServer).PullConsumables(m, &grpc.GenericServerStream[PullConsumablesRequest, PullConsumablesResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type VendingApi_PullConsumablesServer = grpc.ServerStreamingServer[PullConsumablesResponse]

func _VendingApi_GetStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VendingApiServer).GetStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VendingApi_GetStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VendingApiServer).GetStock(ctx, req.(*GetStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VendingApi_UpdateStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VendingApiServer).UpdateStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VendingApi_UpdateStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VendingApiServer).UpdateStock(ctx, req.(*UpdateStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VendingApi_PullStock_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullStockRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VendingApiServer).PullStock(m, &grpc.GenericServerStream[PullStockRequest, PullStockResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type VendingApi_PullStockServer = grpc.ServerStreamingServer[PullStockResponse]

func _VendingApi_ListInventory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInventoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VendingApiServer).ListInventory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VendingApi_ListInventory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VendingApiServer).ListInventory(ctx, req.(*ListInventoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VendingApi_PullInventory_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullInventoryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VendingApiServer).PullInventory(m, &grpc.GenericServerStream[PullInventoryRequest, PullInventoryResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type VendingApi_PullInventoryServer = grpc.ServerStreamingServer[PullInventoryResponse]

func _VendingApi_Dispense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DispenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VendingApiServer).Dispense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VendingApi_Dispense_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VendingApiServer).Dispense(ctx, req.(*DispenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VendingApi_StopDispense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopDispenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VendingApiServer).StopDispense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VendingApi_StopDispense_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VendingApiServer).StopDispense(ctx, req.(*StopDispenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VendingApi_ServiceDesc is the grpc.ServiceDesc for VendingApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VendingApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.traits.VendingApi",
	HandlerType: (*VendingApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListConsumables",
			Handler:    _VendingApi_ListConsumables_Handler,
		},
		{
			MethodName: "GetStock",
			Handler:    _VendingApi_GetStock_Handler,
		},
		{
			MethodName: "UpdateStock",
			Handler:    _VendingApi_UpdateStock_Handler,
		},
		{
			MethodName: "ListInventory",
			Handler:    _VendingApi_ListInventory_Handler,
		},
		{
			MethodName: "Dispense",
			Handler:    _VendingApi_Dispense_Handler,
		},
		{
			MethodName: "StopDispense",
			Handler:    _VendingApi_StopDispense_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PullConsumables",
			Handler:       _VendingApi_PullConsumables_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PullStock",
			Handler:       _VendingApi_PullStock_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PullInventory",
			Handler:       _VendingApi_PullInventory_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "traits/vending.proto",
}

// VendingInfoClient is the client API for VendingInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VendingInfoClient interface {
}

type vendingInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewVendingInfoClient(cc grpc.ClientConnInterface) VendingInfoClient {
	return &vendingInfoClient{cc}
}

// VendingInfoServer is the server API for VendingInfo service.
// All implementations must embed UnimplementedVendingInfoServer
// for forward compatibility.
type VendingInfoServer interface {
	mustEmbedUnimplementedVendingInfoServer()
}

// UnimplementedVendingInfoServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedVendingInfoServer struct{}

func (UnimplementedVendingInfoServer) mustEmbedUnimplementedVendingInfoServer() {}
func (UnimplementedVendingInfoServer) testEmbeddedByValue()                     {}

// UnsafeVendingInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VendingInfoServer will
// result in compilation errors.
type UnsafeVendingInfoServer interface {
	mustEmbedUnimplementedVendingInfoServer()
}

func RegisterVendingInfoServer(s grpc.ServiceRegistrar, srv VendingInfoServer) {
	// If the following call pancis, it indicates UnimplementedVendingInfoServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&VendingInfo_ServiceDesc, srv)
}

// VendingInfo_ServiceDesc is the grpc.ServiceDesc for VendingInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VendingInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.traits.VendingInfo",
	HandlerType: (*VendingInfoServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "traits/vending.proto",
}
