// Code generated by protoc-gen-go. DO NOT EDIT.
// source: device/traits/health.proto

package traits

import (
	context "context"
	fmt "fmt"
	types "git.vanti.co.uk/smartcore/sc-api/go/types"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Trait.attribute for describing the capabilities of the device wrt this trait
type HealthAttributes struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthAttributes) Reset()         { *m = HealthAttributes{} }
func (m *HealthAttributes) String() string { return proto.CompactTextString(m) }
func (*HealthAttributes) ProtoMessage()    {}
func (*HealthAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_575fc0ae32bebffa, []int{0}
}

func (m *HealthAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthAttributes.Unmarshal(m, b)
}
func (m *HealthAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthAttributes.Marshal(b, m, deterministic)
}
func (m *HealthAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthAttributes.Merge(m, src)
}
func (m *HealthAttributes) XXX_Size() int {
	return xxx_messageInfo_HealthAttributes.Size(m)
}
func (m *HealthAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_HealthAttributes proto.InternalMessageInfo

type HealthState struct {
	// Health of the connection to the physical device
	Connection *ConnectionHealth `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
	// Health of communication with the device
	Comm                 *CommHealth `protobuf:"bytes,2,opt,name=comm,proto3" json:"comm,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *HealthState) Reset()         { *m = HealthState{} }
func (m *HealthState) String() string { return proto.CompactTextString(m) }
func (*HealthState) ProtoMessage()    {}
func (*HealthState) Descriptor() ([]byte, []int) {
	return fileDescriptor_575fc0ae32bebffa, []int{1}
}

func (m *HealthState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthState.Unmarshal(m, b)
}
func (m *HealthState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthState.Marshal(b, m, deterministic)
}
func (m *HealthState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthState.Merge(m, src)
}
func (m *HealthState) XXX_Size() int {
	return xxx_messageInfo_HealthState.Size(m)
}
func (m *HealthState) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthState.DiscardUnknown(m)
}

var xxx_messageInfo_HealthState proto.InternalMessageInfo

func (m *HealthState) GetConnection() *ConnectionHealth {
	if m != nil {
		return m.Connection
	}
	return nil
}

func (m *HealthState) GetComm() *CommHealth {
	if m != nil {
		return m.Comm
	}
	return nil
}

// Describes the health of a connection
type ConnectionHealth struct {
	// are we currently connected
	Status types.Connectivity `protobuf:"varint,1,opt,name=status,proto3,enum=smartcore.api.types.Connectivity" json:"status,omitempty"`
	// When was the last time a successful connection was established
	ConnectTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=connect_time,json=connectTime,proto3" json:"connect_time,omitempty"`
	// When was the last time a connection was closed/ended
	DisconnectTime       *timestamp.Timestamp `protobuf:"bytes,3,opt,name=disconnect_time,json=disconnectTime,proto3" json:"disconnect_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ConnectionHealth) Reset()         { *m = ConnectionHealth{} }
func (m *ConnectionHealth) String() string { return proto.CompactTextString(m) }
func (*ConnectionHealth) ProtoMessage()    {}
func (*ConnectionHealth) Descriptor() ([]byte, []int) {
	return fileDescriptor_575fc0ae32bebffa, []int{2}
}

func (m *ConnectionHealth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionHealth.Unmarshal(m, b)
}
func (m *ConnectionHealth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionHealth.Marshal(b, m, deterministic)
}
func (m *ConnectionHealth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionHealth.Merge(m, src)
}
func (m *ConnectionHealth) XXX_Size() int {
	return xxx_messageInfo_ConnectionHealth.Size(m)
}
func (m *ConnectionHealth) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionHealth.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionHealth proto.InternalMessageInfo

func (m *ConnectionHealth) GetStatus() types.Connectivity {
	if m != nil {
		return m.Status
	}
	return types.Connectivity_NOT_APPLICABLE
}

func (m *ConnectionHealth) GetConnectTime() *timestamp.Timestamp {
	if m != nil {
		return m.ConnectTime
	}
	return nil
}

func (m *ConnectionHealth) GetDisconnectTime() *timestamp.Timestamp {
	if m != nil {
		return m.DisconnectTime
	}
	return nil
}

// Describes the health of communication.
type CommHealth struct {
	Status types.CommStatus `protobuf:"varint,1,opt,name=status,proto3,enum=smartcore.api.types.CommStatus" json:"status,omitempty"`
	// The most recent time the status changed to COMM_FAILURE
	FailureTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=failure_time,json=failureTime,proto3" json:"failure_time,omitempty"`
	// The most recent time the status changed to COMM_SUCCESS
	SuccessTime          *timestamp.Timestamp `protobuf:"bytes,3,opt,name=success_time,json=successTime,proto3" json:"success_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CommHealth) Reset()         { *m = CommHealth{} }
func (m *CommHealth) String() string { return proto.CompactTextString(m) }
func (*CommHealth) ProtoMessage()    {}
func (*CommHealth) Descriptor() ([]byte, []int) {
	return fileDescriptor_575fc0ae32bebffa, []int{3}
}

func (m *CommHealth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommHealth.Unmarshal(m, b)
}
func (m *CommHealth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommHealth.Marshal(b, m, deterministic)
}
func (m *CommHealth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommHealth.Merge(m, src)
}
func (m *CommHealth) XXX_Size() int {
	return xxx_messageInfo_CommHealth.Size(m)
}
func (m *CommHealth) XXX_DiscardUnknown() {
	xxx_messageInfo_CommHealth.DiscardUnknown(m)
}

var xxx_messageInfo_CommHealth proto.InternalMessageInfo

func (m *CommHealth) GetStatus() types.CommStatus {
	if m != nil {
		return m.Status
	}
	return types.CommStatus_COMM_UNKNOWN
}

func (m *CommHealth) GetFailureTime() *timestamp.Timestamp {
	if m != nil {
		return m.FailureTime
	}
	return nil
}

func (m *CommHealth) GetSuccessTime() *timestamp.Timestamp {
	if m != nil {
		return m.SuccessTime
	}
	return nil
}

type GetHealthStateRequest struct {
	// Name of the device to query the health state for
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHealthStateRequest) Reset()         { *m = GetHealthStateRequest{} }
func (m *GetHealthStateRequest) String() string { return proto.CompactTextString(m) }
func (*GetHealthStateRequest) ProtoMessage()    {}
func (*GetHealthStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_575fc0ae32bebffa, []int{4}
}

func (m *GetHealthStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHealthStateRequest.Unmarshal(m, b)
}
func (m *GetHealthStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHealthStateRequest.Marshal(b, m, deterministic)
}
func (m *GetHealthStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHealthStateRequest.Merge(m, src)
}
func (m *GetHealthStateRequest) XXX_Size() int {
	return xxx_messageInfo_GetHealthStateRequest.Size(m)
}
func (m *GetHealthStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHealthStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetHealthStateRequest proto.InternalMessageInfo

func (m *GetHealthStateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PullHealthStatesRequest struct {
	// Name of the device to subscribe to the health state for
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PullHealthStatesRequest) Reset()         { *m = PullHealthStatesRequest{} }
func (m *PullHealthStatesRequest) String() string { return proto.CompactTextString(m) }
func (*PullHealthStatesRequest) ProtoMessage()    {}
func (*PullHealthStatesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_575fc0ae32bebffa, []int{5}
}

func (m *PullHealthStatesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PullHealthStatesRequest.Unmarshal(m, b)
}
func (m *PullHealthStatesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PullHealthStatesRequest.Marshal(b, m, deterministic)
}
func (m *PullHealthStatesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullHealthStatesRequest.Merge(m, src)
}
func (m *PullHealthStatesRequest) XXX_Size() int {
	return xxx_messageInfo_PullHealthStatesRequest.Size(m)
}
func (m *PullHealthStatesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PullHealthStatesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PullHealthStatesRequest proto.InternalMessageInfo

func (m *PullHealthStatesRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PullHealthStatesResponse struct {
	Changes              []*HealthStateChange `protobuf:"bytes,1,rep,name=changes,proto3" json:"changes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PullHealthStatesResponse) Reset()         { *m = PullHealthStatesResponse{} }
func (m *PullHealthStatesResponse) String() string { return proto.CompactTextString(m) }
func (*PullHealthStatesResponse) ProtoMessage()    {}
func (*PullHealthStatesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_575fc0ae32bebffa, []int{6}
}

func (m *PullHealthStatesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PullHealthStatesResponse.Unmarshal(m, b)
}
func (m *PullHealthStatesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PullHealthStatesResponse.Marshal(b, m, deterministic)
}
func (m *PullHealthStatesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullHealthStatesResponse.Merge(m, src)
}
func (m *PullHealthStatesResponse) XXX_Size() int {
	return xxx_messageInfo_PullHealthStatesResponse.Size(m)
}
func (m *PullHealthStatesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PullHealthStatesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PullHealthStatesResponse proto.InternalMessageInfo

func (m *PullHealthStatesResponse) GetChanges() []*HealthStateChange {
	if m != nil {
		return m.Changes
	}
	return nil
}

type HealthStateChange struct {
	// The name of the device that is the source of the change
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// When the change was applied
	ChangeTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=change_time,json=changeTime,proto3" json:"change_time,omitempty"`
	// The new state for the change
	Health               *HealthState `protobuf:"bytes,3,opt,name=health,proto3" json:"health,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *HealthStateChange) Reset()         { *m = HealthStateChange{} }
func (m *HealthStateChange) String() string { return proto.CompactTextString(m) }
func (*HealthStateChange) ProtoMessage()    {}
func (*HealthStateChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_575fc0ae32bebffa, []int{7}
}

func (m *HealthStateChange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthStateChange.Unmarshal(m, b)
}
func (m *HealthStateChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthStateChange.Marshal(b, m, deterministic)
}
func (m *HealthStateChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthStateChange.Merge(m, src)
}
func (m *HealthStateChange) XXX_Size() int {
	return xxx_messageInfo_HealthStateChange.Size(m)
}
func (m *HealthStateChange) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthStateChange.DiscardUnknown(m)
}

var xxx_messageInfo_HealthStateChange proto.InternalMessageInfo

func (m *HealthStateChange) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HealthStateChange) GetChangeTime() *timestamp.Timestamp {
	if m != nil {
		return m.ChangeTime
	}
	return nil
}

func (m *HealthStateChange) GetHealth() *HealthState {
	if m != nil {
		return m.Health
	}
	return nil
}

func init() {
	proto.RegisterType((*HealthAttributes)(nil), "smartcore.api.device.traits.HealthAttributes")
	proto.RegisterType((*HealthState)(nil), "smartcore.api.device.traits.HealthState")
	proto.RegisterType((*ConnectionHealth)(nil), "smartcore.api.device.traits.ConnectionHealth")
	proto.RegisterType((*CommHealth)(nil), "smartcore.api.device.traits.CommHealth")
	proto.RegisterType((*GetHealthStateRequest)(nil), "smartcore.api.device.traits.GetHealthStateRequest")
	proto.RegisterType((*PullHealthStatesRequest)(nil), "smartcore.api.device.traits.PullHealthStatesRequest")
	proto.RegisterType((*PullHealthStatesResponse)(nil), "smartcore.api.device.traits.PullHealthStatesResponse")
	proto.RegisterType((*HealthStateChange)(nil), "smartcore.api.device.traits.HealthStateChange")
}

func init() { proto.RegisterFile("device/traits/health.proto", fileDescriptor_575fc0ae32bebffa) }

var fileDescriptor_575fc0ae32bebffa = []byte{
	// 556 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xc1, 0x6b, 0x13, 0x4f,
	0x18, 0xed, 0xb4, 0x25, 0x3f, 0x7e, 0xdf, 0x96, 0x58, 0x07, 0xd4, 0xb2, 0x3d, 0x54, 0xf7, 0x62,
	0x40, 0x32, 0xab, 0x51, 0x11, 0x29, 0x05, 0xd3, 0x08, 0xf6, 0x22, 0x84, 0x4d, 0x4f, 0x12, 0x90,
	0xc9, 0x66, 0x9a, 0x0c, 0x66, 0x77, 0xd6, 0x9d, 0x6f, 0x03, 0x3d, 0xfa, 0x67, 0x78, 0xf2, 0xe6,
	0xc1, 0x7f, 0xc3, 0x9b, 0x37, 0xff, 0x21, 0x91, 0xcc, 0x4c, 0x92, 0x4d, 0xda, 0x6e, 0xe3, 0x2d,
	0xcc, 0xf7, 0xde, 0xcb, 0x7b, 0xef, 0x9b, 0x1d, 0xf0, 0x87, 0x62, 0x2a, 0x63, 0x11, 0x62, 0xce,
	0x25, 0xea, 0x70, 0x2c, 0xf8, 0x04, 0xc7, 0x2c, 0xcb, 0x15, 0x2a, 0x7a, 0xa8, 0x13, 0x9e, 0x63,
	0xac, 0x72, 0xc1, 0x78, 0x26, 0x99, 0x45, 0x32, 0x8b, 0xf4, 0x8f, 0x46, 0x4a, 0x8d, 0x26, 0x22,
	0x34, 0xd0, 0x41, 0x71, 0x11, 0xa2, 0x4c, 0x84, 0x46, 0x9e, 0x64, 0x96, 0xed, 0xdf, 0xc7, 0xcb,
	0x4c, 0xe8, 0x30, 0x56, 0x69, 0x2a, 0x62, 0x94, 0x2a, 0xb5, 0xe7, 0x01, 0x85, 0xfd, 0x33, 0xf3,
	0x2f, 0x6d, 0xc4, 0x5c, 0x0e, 0x0a, 0x14, 0x3a, 0xf8, 0x4a, 0xc0, 0xb3, 0x87, 0x3d, 0xe4, 0x28,
	0xe8, 0x7b, 0x80, 0x25, 0xef, 0x80, 0x3c, 0x24, 0x0d, 0xaf, 0xd5, 0x64, 0x15, 0x76, 0x58, 0x67,
	0x01, 0xb7, 0x3a, 0x51, 0x49, 0x80, 0x1e, 0xc3, 0x6e, 0xac, 0x92, 0xe4, 0x60, 0xdb, 0x08, 0x3d,
	0xbe, 0x45, 0x28, 0x49, 0x9c, 0x84, 0x21, 0x05, 0xbf, 0x09, 0xec, 0xaf, 0xab, 0xd3, 0xd7, 0x50,
	0xd3, 0xc8, 0xb1, 0xd0, 0xc6, 0x5c, 0xbd, 0xf5, 0x68, 0x4d, 0xd3, 0x64, 0x5f, 0x98, 0x9a, 0x4a,
	0xbc, 0x8c, 0x1c, 0x81, 0x9e, 0xc0, 0x9e, 0xb3, 0xf6, 0x71, 0x56, 0x99, 0x33, 0xe5, 0x33, 0xdb,
	0x27, 0x9b, 0xf7, 0xc9, 0xce, 0xe7, 0x7d, 0x46, 0x9e, 0xc3, 0xcf, 0x4e, 0x68, 0x07, 0xee, 0x0c,
	0xa5, 0x5e, 0x51, 0xd8, 0xb9, 0x55, 0xa1, 0xbe, 0xa4, 0xcc, 0x0e, 0x83, 0x9f, 0x04, 0x60, 0x19,
	0x94, 0xbe, 0x5a, 0x4b, 0x73, 0x74, 0x43, 0x9a, 0x24, 0xe9, 0x19, 0x58, 0x39, 0xcb, 0x05, 0x97,
	0x93, 0x22, 0x17, 0x1b, 0x67, 0x71, 0x78, 0x93, 0xe5, 0x04, 0xf6, 0x74, 0x11, 0xc7, 0x42, 0xeb,
	0x4d, 0x83, 0x78, 0x0e, 0x6f, 0x52, 0x3c, 0x81, 0x7b, 0xef, 0x04, 0x96, 0xee, 0x4d, 0x24, 0x3e,
	0x17, 0x42, 0x23, 0xa5, 0xb0, 0x9b, 0xf2, 0x44, 0x98, 0x34, 0xff, 0x47, 0xe6, 0x77, 0xd0, 0x84,
	0x07, 0xdd, 0x62, 0x32, 0x29, 0xa1, 0x75, 0x15, 0x7c, 0x08, 0x07, 0x57, 0xe1, 0x3a, 0x53, 0xa9,
	0x16, 0xf4, 0x0c, 0xfe, 0x8b, 0xc7, 0x3c, 0x1d, 0x89, 0x59, 0x5f, 0x3b, 0x0d, 0xaf, 0xc5, 0x2a,
	0x6f, 0x54, 0x49, 0xa3, 0x63, 0x68, 0xd1, 0x9c, 0x1e, 0x7c, 0x27, 0x70, 0xf7, 0xca, 0xf8, 0x3a,
	0x3f, 0xf4, 0x18, 0x3c, 0x4b, 0xda, 0xb4, 0x68, 0xb0, 0x70, 0xd3, 0xf3, 0x1b, 0xa8, 0xd9, 0x0f,
	0xdb, 0x35, 0xdc, 0xd8, 0xd4, 0x6f, 0xe4, 0x78, 0xad, 0x3f, 0x04, 0x6a, 0xee, 0xb2, 0x64, 0x50,
	0x5f, 0x6d, 0x9d, 0xb6, 0x2a, 0xe5, 0xae, 0x5d, 0x91, 0xbf, 0xb1, 0x85, 0x60, 0x8b, 0x7e, 0x21,
	0x50, 0x9f, 0x2f, 0xc3, 0xae, 0x82, 0xbe, 0xa8, 0xa4, 0xdf, 0xb0, 0x68, 0xff, 0xe5, 0x3f, 0xb2,
	0xec, 0xbe, 0x83, 0xad, 0xa7, 0xe4, 0xf4, 0x1b, 0x81, 0xa3, 0xa1, 0x98, 0x56, 0x09, 0x9c, 0xba,
	0x27, 0xac, 0x3b, 0x5b, 0x46, 0x97, 0x7c, 0x78, 0x36, 0x92, 0xc8, 0xa6, 0x3c, 0x45, 0xc9, 0x62,
	0xc5, 0x8a, 0x4f, 0xe1, 0x82, 0x1b, 0xea, 0xb8, 0xc9, 0x33, 0x19, 0x8e, 0x54, 0xb8, 0xf2, 0xfc,
	0xfe, 0xd8, 0x3e, 0xec, 0x2d, 0xf4, 0xdb, 0x99, 0x64, 0x6f, 0xad, 0xfe, 0xb9, 0x99, 0xfe, 0x2a,
	0x4d, 0xfb, 0xed, 0x4c, 0xf6, 0xed, 0xb4, 0x6f, 0xa7, 0x83, 0x9a, 0xb9, 0x04, 0xcf, 0xff, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x4f, 0x16, 0x87, 0x75, 0xd2, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HealthClient is the client API for Health service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HealthClient interface {
	GetHealthState(ctx context.Context, in *GetHealthStateRequest, opts ...grpc.CallOption) (*HealthState, error)
	PullHealStates(ctx context.Context, in *PullHealthStatesRequest, opts ...grpc.CallOption) (Health_PullHealStatesClient, error)
}

type healthClient struct {
	cc *grpc.ClientConn
}

func NewHealthClient(cc *grpc.ClientConn) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) GetHealthState(ctx context.Context, in *GetHealthStateRequest, opts ...grpc.CallOption) (*HealthState, error) {
	out := new(HealthState)
	err := c.cc.Invoke(ctx, "/smartcore.api.device.traits.Health/GetHealthState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthClient) PullHealStates(ctx context.Context, in *PullHealthStatesRequest, opts ...grpc.CallOption) (Health_PullHealStatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Health_serviceDesc.Streams[0], "/smartcore.api.device.traits.Health/PullHealStates", opts...)
	if err != nil {
		return nil, err
	}
	x := &healthPullHealStatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Health_PullHealStatesClient interface {
	Recv() (*PullHealthStatesResponse, error)
	grpc.ClientStream
}

type healthPullHealStatesClient struct {
	grpc.ClientStream
}

func (x *healthPullHealStatesClient) Recv() (*PullHealthStatesResponse, error) {
	m := new(PullHealthStatesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HealthServer is the server API for Health service.
type HealthServer interface {
	GetHealthState(context.Context, *GetHealthStateRequest) (*HealthState, error)
	PullHealStates(*PullHealthStatesRequest, Health_PullHealStatesServer) error
}

// UnimplementedHealthServer can be embedded to have forward compatible implementations.
type UnimplementedHealthServer struct {
}

func (*UnimplementedHealthServer) GetHealthState(ctx context.Context, req *GetHealthStateRequest) (*HealthState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHealthState not implemented")
}
func (*UnimplementedHealthServer) PullHealStates(req *PullHealthStatesRequest, srv Health_PullHealStatesServer) error {
	return status.Errorf(codes.Unimplemented, "method PullHealStates not implemented")
}

func RegisterHealthServer(s *grpc.Server, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
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
		FullMethod: "/smartcore.api.device.traits.Health/GetHealthState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).GetHealthState(ctx, req.(*GetHealthStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Health_PullHealStates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullHealthStatesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HealthServer).PullHealStates(m, &healthPullHealStatesServer{stream})
}

type Health_PullHealStatesServer interface {
	Send(*PullHealthStatesResponse) error
	grpc.ServerStream
}

type healthPullHealStatesServer struct {
	grpc.ServerStream
}

func (x *healthPullHealStatesServer) Send(m *PullHealthStatesResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Health_serviceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.api.device.traits.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHealthState",
			Handler:    _Health_GetHealthState_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PullHealStates",
			Handler:       _Health_PullHealStates_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "device/traits/health.proto",
}