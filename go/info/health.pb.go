// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.3
// source: info/health.proto

package info

import (
	types "github.com/smart-core-os/sc-api/go/types"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HealthState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Health of the connection to the physical device
	Connection *ConnectionHealth `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
	// Health of communication with the device
	Comm *CommHealth `protobuf:"bytes,2,opt,name=comm,proto3" json:"comm,omitempty"`
}

func (x *HealthState) Reset() {
	*x = HealthState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_health_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthState) ProtoMessage() {}

func (x *HealthState) ProtoReflect() protoreflect.Message {
	mi := &file_info_health_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthState.ProtoReflect.Descriptor instead.
func (*HealthState) Descriptor() ([]byte, []int) {
	return file_info_health_proto_rawDescGZIP(), []int{0}
}

func (x *HealthState) GetConnection() *ConnectionHealth {
	if x != nil {
		return x.Connection
	}
	return nil
}

func (x *HealthState) GetComm() *CommHealth {
	if x != nil {
		return x.Comm
	}
	return nil
}

// Describes the health of a connection
type ConnectionHealth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// are we currently connected
	Status types.Connectivity `protobuf:"varint,1,opt,name=status,proto3,enum=smartcore.types.Connectivity" json:"status,omitempty"`
	// When was the last time a successful connection was established
	ConnectTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=connect_time,json=connectTime,proto3" json:"connect_time,omitempty"`
	// When was the last time a connection was closed/ended
	DisconnectTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=disconnect_time,json=disconnectTime,proto3" json:"disconnect_time,omitempty"`
}

func (x *ConnectionHealth) Reset() {
	*x = ConnectionHealth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_health_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionHealth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionHealth) ProtoMessage() {}

func (x *ConnectionHealth) ProtoReflect() protoreflect.Message {
	mi := &file_info_health_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionHealth.ProtoReflect.Descriptor instead.
func (*ConnectionHealth) Descriptor() ([]byte, []int) {
	return file_info_health_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectionHealth) GetStatus() types.Connectivity {
	if x != nil {
		return x.Status
	}
	return types.Connectivity(0)
}

func (x *ConnectionHealth) GetConnectTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ConnectTime
	}
	return nil
}

func (x *ConnectionHealth) GetDisconnectTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DisconnectTime
	}
	return nil
}

// Describes the health of communication.
type CommHealth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status types.CommStatus `protobuf:"varint,1,opt,name=status,proto3,enum=smartcore.types.CommStatus" json:"status,omitempty"`
	// The most recent time the status changed to COMM_FAILURE
	FailureTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=failure_time,json=failureTime,proto3" json:"failure_time,omitempty"`
	// The most recent time the status changed to COMM_SUCCESS
	SuccessTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=success_time,json=successTime,proto3" json:"success_time,omitempty"`
}

func (x *CommHealth) Reset() {
	*x = CommHealth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_health_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommHealth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommHealth) ProtoMessage() {}

func (x *CommHealth) ProtoReflect() protoreflect.Message {
	mi := &file_info_health_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommHealth.ProtoReflect.Descriptor instead.
func (*CommHealth) Descriptor() ([]byte, []int) {
	return file_info_health_proto_rawDescGZIP(), []int{2}
}

func (x *CommHealth) GetStatus() types.CommStatus {
	if x != nil {
		return x.Status
	}
	return types.CommStatus(0)
}

func (x *CommHealth) GetFailureTime() *timestamppb.Timestamp {
	if x != nil {
		return x.FailureTime
	}
	return nil
}

func (x *CommHealth) GetSuccessTime() *timestamppb.Timestamp {
	if x != nil {
		return x.SuccessTime
	}
	return nil
}

type GetHealthStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the device to query the health state for
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetHealthStateRequest) Reset() {
	*x = GetHealthStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_health_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHealthStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHealthStateRequest) ProtoMessage() {}

func (x *GetHealthStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_info_health_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHealthStateRequest.ProtoReflect.Descriptor instead.
func (*GetHealthStateRequest) Descriptor() ([]byte, []int) {
	return file_info_health_proto_rawDescGZIP(), []int{3}
}

func (x *GetHealthStateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PullHealthStatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the device to subscribe to the health state for
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PullHealthStatesRequest) Reset() {
	*x = PullHealthStatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_health_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullHealthStatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullHealthStatesRequest) ProtoMessage() {}

func (x *PullHealthStatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_info_health_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullHealthStatesRequest.ProtoReflect.Descriptor instead.
func (*PullHealthStatesRequest) Descriptor() ([]byte, []int) {
	return file_info_health_proto_rawDescGZIP(), []int{4}
}

func (x *PullHealthStatesRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PullHealthStatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Changes []*HealthStateChange `protobuf:"bytes,1,rep,name=changes,proto3" json:"changes,omitempty"`
}

func (x *PullHealthStatesResponse) Reset() {
	*x = PullHealthStatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_health_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullHealthStatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullHealthStatesResponse) ProtoMessage() {}

func (x *PullHealthStatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_info_health_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullHealthStatesResponse.ProtoReflect.Descriptor instead.
func (*PullHealthStatesResponse) Descriptor() ([]byte, []int) {
	return file_info_health_proto_rawDescGZIP(), []int{5}
}

func (x *PullHealthStatesResponse) GetChanges() []*HealthStateChange {
	if x != nil {
		return x.Changes
	}
	return nil
}

type HealthStateChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the device that is the source of the change
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// When the change was applied
	ChangeTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=change_time,json=changeTime,proto3" json:"change_time,omitempty"`
	// The new state for the change
	Health *HealthState `protobuf:"bytes,3,opt,name=health,proto3" json:"health,omitempty"`
}

func (x *HealthStateChange) Reset() {
	*x = HealthStateChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_health_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthStateChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthStateChange) ProtoMessage() {}

func (x *HealthStateChange) ProtoReflect() protoreflect.Message {
	mi := &file_info_health_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthStateChange.ProtoReflect.Descriptor instead.
func (*HealthStateChange) Descriptor() ([]byte, []int) {
	return file_info_health_proto_rawDescGZIP(), []int{6}
}

func (x *HealthStateChange) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HealthStateChange) GetChangeTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ChangeTime
	}
	return nil
}

func (x *HealthStateChange) GetHealth() *HealthState {
	if x != nil {
		return x.Health
	}
	return nil
}

var File_info_health_proto protoreflect.FileDescriptor

var file_info_health_proto_rawDesc = []byte{
	0x0a, 0x11, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x69,
	0x6e, 0x66, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x0b,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x40, 0x0a, 0x0a, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x6f,
	0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a,
	0x04, 0x63, 0x6f, 0x6d, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x6d,
	0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x04, 0x63, 0x6f, 0x6d, 0x6d, 0x22, 0xcd, 0x01,
	0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x12, 0x35, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x43, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xbf, 0x01,
	0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x33, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x73,
	0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x3d, 0x0a, 0x0c, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0b, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x3d, 0x0a, 0x0c, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0b, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0x2b, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2d, 0x0a, 0x17,
	0x50, 0x75, 0x6c, 0x6c, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x57, 0x0a, 0x18, 0x50,
	0x75, 0x6c, 0x6c, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x07, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x73, 0x22, 0x99, 0x01, 0x0a, 0x11, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3b,
	0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x6d,
	0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x32, 0xcb, 0x01, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x56, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x25, 0x2e,
	0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x22, 0x00, 0x12, 0x69, 0x0a, 0x10, 0x50, 0x75, 0x6c, 0x6c, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x27, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x28, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x69, 0x6e, 0x66,
	0x6f, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x6e,
	0x0a, 0x12, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x69, 0x6e, 0x66, 0x6f, 0x42, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6d, 0x61, 0x72, 0x74, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2d, 0x6f, 0x73, 0x2f, 0x73, 0x63,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0xaa, 0x02, 0x0e, 0x53,
	0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0xca, 0x02, 0x0e,
	0x53, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x5c, 0x49, 0x6e, 0x66, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_info_health_proto_rawDescOnce sync.Once
	file_info_health_proto_rawDescData = file_info_health_proto_rawDesc
)

func file_info_health_proto_rawDescGZIP() []byte {
	file_info_health_proto_rawDescOnce.Do(func() {
		file_info_health_proto_rawDescData = protoimpl.X.CompressGZIP(file_info_health_proto_rawDescData)
	})
	return file_info_health_proto_rawDescData
}

var file_info_health_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_info_health_proto_goTypes = []any{
	(*HealthState)(nil),              // 0: smartcore.info.HealthState
	(*ConnectionHealth)(nil),         // 1: smartcore.info.ConnectionHealth
	(*CommHealth)(nil),               // 2: smartcore.info.CommHealth
	(*GetHealthStateRequest)(nil),    // 3: smartcore.info.GetHealthStateRequest
	(*PullHealthStatesRequest)(nil),  // 4: smartcore.info.PullHealthStatesRequest
	(*PullHealthStatesResponse)(nil), // 5: smartcore.info.PullHealthStatesResponse
	(*HealthStateChange)(nil),        // 6: smartcore.info.HealthStateChange
	(types.Connectivity)(0),          // 7: smartcore.types.Connectivity
	(*timestamppb.Timestamp)(nil),    // 8: google.protobuf.Timestamp
	(types.CommStatus)(0),            // 9: smartcore.types.CommStatus
}
var file_info_health_proto_depIdxs = []int32{
	1,  // 0: smartcore.info.HealthState.connection:type_name -> smartcore.info.ConnectionHealth
	2,  // 1: smartcore.info.HealthState.comm:type_name -> smartcore.info.CommHealth
	7,  // 2: smartcore.info.ConnectionHealth.status:type_name -> smartcore.types.Connectivity
	8,  // 3: smartcore.info.ConnectionHealth.connect_time:type_name -> google.protobuf.Timestamp
	8,  // 4: smartcore.info.ConnectionHealth.disconnect_time:type_name -> google.protobuf.Timestamp
	9,  // 5: smartcore.info.CommHealth.status:type_name -> smartcore.types.CommStatus
	8,  // 6: smartcore.info.CommHealth.failure_time:type_name -> google.protobuf.Timestamp
	8,  // 7: smartcore.info.CommHealth.success_time:type_name -> google.protobuf.Timestamp
	6,  // 8: smartcore.info.PullHealthStatesResponse.changes:type_name -> smartcore.info.HealthStateChange
	8,  // 9: smartcore.info.HealthStateChange.change_time:type_name -> google.protobuf.Timestamp
	0,  // 10: smartcore.info.HealthStateChange.health:type_name -> smartcore.info.HealthState
	3,  // 11: smartcore.info.Health.GetHealthState:input_type -> smartcore.info.GetHealthStateRequest
	4,  // 12: smartcore.info.Health.PullHealthStates:input_type -> smartcore.info.PullHealthStatesRequest
	0,  // 13: smartcore.info.Health.GetHealthState:output_type -> smartcore.info.HealthState
	5,  // 14: smartcore.info.Health.PullHealthStates:output_type -> smartcore.info.PullHealthStatesResponse
	13, // [13:15] is the sub-list for method output_type
	11, // [11:13] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_info_health_proto_init() }
func file_info_health_proto_init() {
	if File_info_health_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_info_health_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*HealthState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_info_health_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ConnectionHealth); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_info_health_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CommHealth); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_info_health_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetHealthStateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_info_health_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*PullHealthStatesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_info_health_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*PullHealthStatesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_info_health_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*HealthStateChange); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_info_health_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_info_health_proto_goTypes,
		DependencyIndexes: file_info_health_proto_depIdxs,
		MessageInfos:      file_info_health_proto_msgTypes,
	}.Build()
	File_info_health_proto = out.File
	file_info_health_proto_rawDesc = nil
	file_info_health_proto_goTypes = nil
	file_info_health_proto_depIdxs = nil
}
