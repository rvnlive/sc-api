// Code generated by protoc-gen-go. DO NOT EDIT.
// source: device/traits/microphone.proto

package traits

import (
	context "context"
	fmt "fmt"
	types "git.vanti.co.uk/smartcore/sc-api/go/types"
	proto "github.com/golang/protobuf/proto"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Trait.attribute describing capabilities of the microphone device
type MicrophoneAttributes struct {
	// Attributes associated with the gain property of the microphone. Note that a step value that is the same as the bounds
	// implies that the microphone only supports mute
	GainAttributes *types.FloatAttributes `protobuf:"bytes,1,opt,name=gain_attributes,json=gainAttributes,proto3" json:"gain_attributes,omitempty"`
	// How is mute implemented by the device. Can help to customise behaviour of interfaces to the device, e.g. by
	// disallowing gain change when muted
	MuteSupport          types.MuteSupport `protobuf:"varint,3,opt,name=mute_support,json=muteSupport,proto3,enum=smartcore.api.types.MuteSupport" json:"mute_support,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MicrophoneAttributes) Reset()         { *m = MicrophoneAttributes{} }
func (m *MicrophoneAttributes) String() string { return proto.CompactTextString(m) }
func (*MicrophoneAttributes) ProtoMessage()    {}
func (*MicrophoneAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_87e3316f3d02b252, []int{0}
}

func (m *MicrophoneAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MicrophoneAttributes.Unmarshal(m, b)
}
func (m *MicrophoneAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MicrophoneAttributes.Marshal(b, m, deterministic)
}
func (m *MicrophoneAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MicrophoneAttributes.Merge(m, src)
}
func (m *MicrophoneAttributes) XXX_Size() int {
	return xxx_messageInfo_MicrophoneAttributes.Size(m)
}
func (m *MicrophoneAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_MicrophoneAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_MicrophoneAttributes proto.InternalMessageInfo

func (m *MicrophoneAttributes) GetGainAttributes() *types.FloatAttributes {
	if m != nil {
		return m.GainAttributes
	}
	return nil
}

func (m *MicrophoneAttributes) GetMuteSupport() types.MuteSupport {
	if m != nil {
		return m.MuteSupport
	}
	return types.MuteSupport_MUTE_NATIVE
}

type GetMicrophoneGainRequest struct {
	// Name of the device to fetch the state for
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// fields to fetch relative the the Volume type
	Fields               *field_mask.FieldMask `protobuf:"bytes,2,opt,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetMicrophoneGainRequest) Reset()         { *m = GetMicrophoneGainRequest{} }
func (m *GetMicrophoneGainRequest) String() string { return proto.CompactTextString(m) }
func (*GetMicrophoneGainRequest) ProtoMessage()    {}
func (*GetMicrophoneGainRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_87e3316f3d02b252, []int{1}
}

func (m *GetMicrophoneGainRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMicrophoneGainRequest.Unmarshal(m, b)
}
func (m *GetMicrophoneGainRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMicrophoneGainRequest.Marshal(b, m, deterministic)
}
func (m *GetMicrophoneGainRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMicrophoneGainRequest.Merge(m, src)
}
func (m *GetMicrophoneGainRequest) XXX_Size() int {
	return xxx_messageInfo_GetMicrophoneGainRequest.Size(m)
}
func (m *GetMicrophoneGainRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMicrophoneGainRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMicrophoneGainRequest proto.InternalMessageInfo

func (m *GetMicrophoneGainRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetMicrophoneGainRequest) GetFields() *field_mask.FieldMask {
	if m != nil {
		return m.Fields
	}
	return nil
}

type UpdateMicrophoneGainRequest struct {
	// Name of the device to update the state for
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The updated gain value
	Gain *types.Volume `protobuf:"bytes,2,opt,name=gain,proto3" json:"gain,omitempty"`
	// Only supported for Volume.level. Update the value relative to the current value.
	Delta bool `protobuf:"varint,3,opt,name=delta,proto3" json:"delta,omitempty"`
	// The fields we intend to update relative to the Volume type
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateMicrophoneGainRequest) Reset()         { *m = UpdateMicrophoneGainRequest{} }
func (m *UpdateMicrophoneGainRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateMicrophoneGainRequest) ProtoMessage()    {}
func (*UpdateMicrophoneGainRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_87e3316f3d02b252, []int{2}
}

func (m *UpdateMicrophoneGainRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateMicrophoneGainRequest.Unmarshal(m, b)
}
func (m *UpdateMicrophoneGainRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateMicrophoneGainRequest.Marshal(b, m, deterministic)
}
func (m *UpdateMicrophoneGainRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateMicrophoneGainRequest.Merge(m, src)
}
func (m *UpdateMicrophoneGainRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateMicrophoneGainRequest.Size(m)
}
func (m *UpdateMicrophoneGainRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateMicrophoneGainRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateMicrophoneGainRequest proto.InternalMessageInfo

func (m *UpdateMicrophoneGainRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateMicrophoneGainRequest) GetGain() *types.Volume {
	if m != nil {
		return m.Gain
	}
	return nil
}

func (m *UpdateMicrophoneGainRequest) GetDelta() bool {
	if m != nil {
		return m.Delta
	}
	return false
}

func (m *UpdateMicrophoneGainRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type PullMicrophoneGainRequest struct {
	// Name of the device to fetch the state for
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// fields to fetch relative the the Volume type
	Fields               *field_mask.FieldMask `protobuf:"bytes,2,opt,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PullMicrophoneGainRequest) Reset()         { *m = PullMicrophoneGainRequest{} }
func (m *PullMicrophoneGainRequest) String() string { return proto.CompactTextString(m) }
func (*PullMicrophoneGainRequest) ProtoMessage()    {}
func (*PullMicrophoneGainRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_87e3316f3d02b252, []int{3}
}

func (m *PullMicrophoneGainRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PullMicrophoneGainRequest.Unmarshal(m, b)
}
func (m *PullMicrophoneGainRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PullMicrophoneGainRequest.Marshal(b, m, deterministic)
}
func (m *PullMicrophoneGainRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullMicrophoneGainRequest.Merge(m, src)
}
func (m *PullMicrophoneGainRequest) XXX_Size() int {
	return xxx_messageInfo_PullMicrophoneGainRequest.Size(m)
}
func (m *PullMicrophoneGainRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PullMicrophoneGainRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PullMicrophoneGainRequest proto.InternalMessageInfo

func (m *PullMicrophoneGainRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PullMicrophoneGainRequest) GetFields() *field_mask.FieldMask {
	if m != nil {
		return m.Fields
	}
	return nil
}

type PullMicrophoneGainResponse struct {
	// Changes since the last message
	Changes              []*types.VolumeChange `protobuf:"bytes,1,rep,name=changes,proto3" json:"changes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PullMicrophoneGainResponse) Reset()         { *m = PullMicrophoneGainResponse{} }
func (m *PullMicrophoneGainResponse) String() string { return proto.CompactTextString(m) }
func (*PullMicrophoneGainResponse) ProtoMessage()    {}
func (*PullMicrophoneGainResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_87e3316f3d02b252, []int{4}
}

func (m *PullMicrophoneGainResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PullMicrophoneGainResponse.Unmarshal(m, b)
}
func (m *PullMicrophoneGainResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PullMicrophoneGainResponse.Marshal(b, m, deterministic)
}
func (m *PullMicrophoneGainResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullMicrophoneGainResponse.Merge(m, src)
}
func (m *PullMicrophoneGainResponse) XXX_Size() int {
	return xxx_messageInfo_PullMicrophoneGainResponse.Size(m)
}
func (m *PullMicrophoneGainResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PullMicrophoneGainResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PullMicrophoneGainResponse proto.InternalMessageInfo

func (m *PullMicrophoneGainResponse) GetChanges() []*types.VolumeChange {
	if m != nil {
		return m.Changes
	}
	return nil
}

func init() {
	proto.RegisterType((*MicrophoneAttributes)(nil), "smartcore.api.device.traits.MicrophoneAttributes")
	proto.RegisterType((*GetMicrophoneGainRequest)(nil), "smartcore.api.device.traits.GetMicrophoneGainRequest")
	proto.RegisterType((*UpdateMicrophoneGainRequest)(nil), "smartcore.api.device.traits.UpdateMicrophoneGainRequest")
	proto.RegisterType((*PullMicrophoneGainRequest)(nil), "smartcore.api.device.traits.PullMicrophoneGainRequest")
	proto.RegisterType((*PullMicrophoneGainResponse)(nil), "smartcore.api.device.traits.PullMicrophoneGainResponse")
}

func init() { proto.RegisterFile("device/traits/microphone.proto", fileDescriptor_87e3316f3d02b252) }

var fileDescriptor_87e3316f3d02b252 = []byte{
	// 525 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0xd5, 0xa6, 0xa1, 0x2d, 0x13, 0xd4, 0x22, 0xab, 0x87, 0xb0, 0x91, 0x20, 0x44, 0x1c, 0x72,
	0xc1, 0x86, 0x20, 0x3e, 0xa4, 0x9e, 0xd2, 0xa2, 0xe6, 0x14, 0x29, 0xda, 0x02, 0x12, 0x28, 0x28,
	0x72, 0x36, 0xee, 0xd6, 0xca, 0xee, 0xda, 0xac, 0xc7, 0x91, 0x38, 0xf0, 0x6b, 0xb8, 0xf5, 0xce,
	0x9f, 0xe0, 0xc2, 0x5f, 0x42, 0x6b, 0xa7, 0x49, 0x23, 0xa5, 0x0b, 0x1c, 0xb8, 0xd9, 0x33, 0xef,
	0xcd, 0x1b, 0xbf, 0x99, 0x5d, 0x78, 0x38, 0x13, 0x0b, 0x19, 0x0b, 0x86, 0x05, 0x97, 0x68, 0x58,
	0x26, 0xe3, 0x42, 0xe9, 0x4b, 0x95, 0x0b, 0xaa, 0x0b, 0x85, 0x8a, 0xb4, 0x4c, 0xc6, 0x0b, 0x8c,
	0x55, 0x21, 0x28, 0xd7, 0x92, 0x7a, 0x34, 0xf5, 0xe8, 0xb0, 0x9d, 0x28, 0x95, 0xa4, 0x82, 0x39,
	0xe8, 0xd4, 0x5e, 0xb0, 0x0b, 0x29, 0xd2, 0xd9, 0x24, 0xe3, 0x66, 0xee, 0xe9, 0x21, 0xc1, 0xaf,
	0x5a, 0x18, 0x96, 0xdb, 0x6c, 0x2a, 0x8a, 0x65, 0xec, 0xbe, 0x8f, 0xd9, 0x5c, 0xe2, 0x26, 0x6a,
	0xa1, 0x52, 0x9b, 0x2d, 0x85, 0x3b, 0x57, 0x01, 0x1c, 0x0d, 0x57, 0xdd, 0xf4, 0x11, 0x0b, 0x39,
	0xb5, 0x28, 0x0c, 0x19, 0xc2, 0x61, 0xc2, 0x65, 0x3e, 0xe1, 0xab, 0x50, 0x33, 0x68, 0x07, 0xdd,
	0x46, 0xef, 0x09, 0xdd, 0xec, 0xd5, 0x15, 0xa5, 0x67, 0xa9, 0xe2, 0xb8, 0xa6, 0x47, 0x07, 0x25,
	0xf9, 0x46, 0xb9, 0x53, 0xb8, 0x97, 0x59, 0x14, 0x13, 0x63, 0xb5, 0x56, 0x05, 0x36, 0x77, 0xda,
	0x41, 0xf7, 0xa0, 0xd7, 0xde, 0x5a, 0x6b, 0x68, 0x51, 0x9c, 0x7b, 0x5c, 0xd4, 0xc8, 0xd6, 0x97,
	0xce, 0x14, 0x9a, 0x03, 0x81, 0xeb, 0x76, 0x07, 0x5c, 0xe6, 0x91, 0xf8, 0x62, 0x85, 0x41, 0x42,
	0xa0, 0x9e, 0xf3, 0x4c, 0xb8, 0x26, 0xef, 0x46, 0xee, 0x4c, 0x7a, 0xb0, 0xeb, 0xac, 0x32, 0xcd,
	0x9a, 0x6b, 0x3d, 0xa4, 0xde, 0x49, 0x7a, 0xed, 0x24, 0x3d, 0x2b, 0xd3, 0x43, 0x6e, 0xe6, 0xd1,
	0x12, 0xd9, 0xf9, 0x11, 0x40, 0xeb, 0xbd, 0x9e, 0x71, 0x14, 0x7f, 0xaf, 0xc3, 0xa0, 0x5e, 0x3e,
	0x77, 0xa9, 0xd2, 0xda, 0xfa, 0xa8, 0x0f, 0xce, 0xf5, 0xc8, 0x01, 0xc9, 0x11, 0xdc, 0x99, 0x89,
	0x14, 0xb9, 0xb3, 0x61, 0x3f, 0xf2, 0x17, 0x72, 0x0c, 0x0d, 0xeb, 0x94, 0xdd, 0x68, 0x9b, 0xf5,
	0x3f, 0xf6, 0x0c, 0x1e, 0x5e, 0x9e, 0x3b, 0x31, 0x3c, 0x18, 0xd9, 0x34, 0xfd, 0xbf, 0xe6, 0x7c,
	0x84, 0x70, 0x9b, 0x88, 0xd1, 0x2a, 0x37, 0x82, 0x1c, 0xc3, 0x5e, 0x7c, 0xc9, 0xf3, 0xc4, 0xad,
	0xca, 0x4e, 0xb7, 0xd1, 0x7b, 0x5c, 0xe1, 0xc4, 0xa9, 0x43, 0x46, 0xd7, 0x8c, 0xde, 0xaf, 0x1a,
	0xc0, 0xba, 0x2e, 0xf9, 0x0c, 0x7b, 0x03, 0x81, 0x65, 0x79, 0xf2, 0x92, 0x56, 0x7c, 0x1c, 0xf4,
	0xb6, 0x85, 0x08, 0xab, 0xc6, 0x40, 0x62, 0x00, 0x3f, 0x64, 0xa7, 0xf0, 0xa6, 0x52, 0xa1, 0x62,
	0x1b, 0xaa, 0x45, 0xbe, 0xc1, 0x7e, 0xe9, 0x96, 0x93, 0x78, 0x55, 0x29, 0x71, 0xeb, 0xe4, 0xc2,
	0xd7, 0xff, 0xcc, 0xf3, 0xc3, 0x78, 0x16, 0x9c, 0x7c, 0x0f, 0xe0, 0xd1, 0x4c, 0x2c, 0xaa, 0x0a,
	0x9c, 0x1c, 0xae, 0xd9, 0xa3, 0x72, 0xec, 0xa3, 0xe0, 0xd3, 0xf3, 0x44, 0x22, 0x5d, 0xf0, 0x1c,
	0x25, 0x8d, 0x15, 0xb5, 0x73, 0xb6, 0xe2, 0x33, 0x13, 0x3f, 0xe5, 0x5a, 0xb2, 0x44, 0xb1, 0x8d,
	0xdf, 0xd9, 0x55, 0xad, 0x75, 0xbe, 0xd2, 0xe8, 0x6b, 0x49, 0xdf, 0x7a, 0x8d, 0x77, 0x2e, 0xfb,
	0xf3, 0x46, 0x76, 0xdc, 0xd7, 0x72, 0xec, 0xb3, 0x63, 0x9f, 0x9d, 0xee, 0xba, 0x75, 0x7b, 0xf1,
	0x3b, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x2b, 0x75, 0x1a, 0x22, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MicrophoneClient is the client API for Microphone service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MicrophoneClient interface {
	// Get the current state of the gain for the device
	GetGain(ctx context.Context, in *GetMicrophoneGainRequest, opts ...grpc.CallOption) (*types.Volume, error)
	// update the gain state for the device
	UpdateGain(ctx context.Context, in *UpdateMicrophoneGainRequest, opts ...grpc.CallOption) (*types.Volume, error)
	PullGain(ctx context.Context, in *PullMicrophoneGainRequest, opts ...grpc.CallOption) (Microphone_PullGainClient, error)
}

type microphoneClient struct {
	cc *grpc.ClientConn
}

func NewMicrophoneClient(cc *grpc.ClientConn) MicrophoneClient {
	return &microphoneClient{cc}
}

func (c *microphoneClient) GetGain(ctx context.Context, in *GetMicrophoneGainRequest, opts ...grpc.CallOption) (*types.Volume, error) {
	out := new(types.Volume)
	err := c.cc.Invoke(ctx, "/smartcore.api.device.traits.Microphone/GetGain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microphoneClient) UpdateGain(ctx context.Context, in *UpdateMicrophoneGainRequest, opts ...grpc.CallOption) (*types.Volume, error) {
	out := new(types.Volume)
	err := c.cc.Invoke(ctx, "/smartcore.api.device.traits.Microphone/UpdateGain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microphoneClient) PullGain(ctx context.Context, in *PullMicrophoneGainRequest, opts ...grpc.CallOption) (Microphone_PullGainClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Microphone_serviceDesc.Streams[0], "/smartcore.api.device.traits.Microphone/PullGain", opts...)
	if err != nil {
		return nil, err
	}
	x := &microphonePullGainClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Microphone_PullGainClient interface {
	Recv() (*PullMicrophoneGainResponse, error)
	grpc.ClientStream
}

type microphonePullGainClient struct {
	grpc.ClientStream
}

func (x *microphonePullGainClient) Recv() (*PullMicrophoneGainResponse, error) {
	m := new(PullMicrophoneGainResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MicrophoneServer is the server API for Microphone service.
type MicrophoneServer interface {
	// Get the current state of the gain for the device
	GetGain(context.Context, *GetMicrophoneGainRequest) (*types.Volume, error)
	// update the gain state for the device
	UpdateGain(context.Context, *UpdateMicrophoneGainRequest) (*types.Volume, error)
	PullGain(*PullMicrophoneGainRequest, Microphone_PullGainServer) error
}

// UnimplementedMicrophoneServer can be embedded to have forward compatible implementations.
type UnimplementedMicrophoneServer struct {
}

func (*UnimplementedMicrophoneServer) GetGain(ctx context.Context, req *GetMicrophoneGainRequest) (*types.Volume, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGain not implemented")
}
func (*UnimplementedMicrophoneServer) UpdateGain(ctx context.Context, req *UpdateMicrophoneGainRequest) (*types.Volume, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGain not implemented")
}
func (*UnimplementedMicrophoneServer) PullGain(req *PullMicrophoneGainRequest, srv Microphone_PullGainServer) error {
	return status.Errorf(codes.Unimplemented, "method PullGain not implemented")
}

func RegisterMicrophoneServer(s *grpc.Server, srv MicrophoneServer) {
	s.RegisterService(&_Microphone_serviceDesc, srv)
}

func _Microphone_GetGain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMicrophoneGainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicrophoneServer).GetGain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.api.device.traits.Microphone/GetGain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicrophoneServer).GetGain(ctx, req.(*GetMicrophoneGainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Microphone_UpdateGain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMicrophoneGainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicrophoneServer).UpdateGain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.api.device.traits.Microphone/UpdateGain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicrophoneServer).UpdateGain(ctx, req.(*UpdateMicrophoneGainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Microphone_PullGain_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullMicrophoneGainRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MicrophoneServer).PullGain(m, &microphonePullGainServer{stream})
}

type Microphone_PullGainServer interface {
	Send(*PullMicrophoneGainResponse) error
	grpc.ServerStream
}

type microphonePullGainServer struct {
	grpc.ServerStream
}

func (x *microphonePullGainServer) Send(m *PullMicrophoneGainResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Microphone_serviceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.api.device.traits.Microphone",
	HandlerType: (*MicrophoneServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGain",
			Handler:    _Microphone_GetGain_Handler,
		},
		{
			MethodName: "UpdateGain",
			Handler:    _Microphone_UpdateGain_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PullGain",
			Handler:       _Microphone_PullGain_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "device/traits/microphone.proto",
}