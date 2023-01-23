// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: traits/emergency.proto

package traits

import (
	types "github.com/smart-core-os/sc-api/go/types"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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

// Levels of emergency
type Emergency_Level int32

const (
	// The level is unknown, applications may choose to be cautious or care-free with this lack of information.
	Emergency_LEVEL_UNSPECIFIED Emergency_Level = 0
	// There is no emergency.
	Emergency_OK Emergency_Level = 1
	// There is a chance of an emergency soon. This might be reported by a smoke detector that has detected some particles
	// but not enough to trigger the alarm yet.
	Emergency_WARNING Emergency_Level = 2
	// There is an emergency. Please protect yourself and follow emergency procedures.
	Emergency_EMERGENCY Emergency_Level = 3
)

// Enum value maps for Emergency_Level.
var (
	Emergency_Level_name = map[int32]string{
		0: "LEVEL_UNSPECIFIED",
		1: "OK",
		2: "WARNING",
		3: "EMERGENCY",
	}
	Emergency_Level_value = map[string]int32{
		"LEVEL_UNSPECIFIED": 0,
		"OK":                1,
		"WARNING":           2,
		"EMERGENCY":         3,
	}
)

func (x Emergency_Level) Enum() *Emergency_Level {
	p := new(Emergency_Level)
	*p = x
	return p
}

func (x Emergency_Level) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Emergency_Level) Descriptor() protoreflect.EnumDescriptor {
	return file_traits_emergency_proto_enumTypes[0].Descriptor()
}

func (Emergency_Level) Type() protoreflect.EnumType {
	return &file_traits_emergency_proto_enumTypes[0]
}

func (x Emergency_Level) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Emergency_Level.Descriptor instead.
func (Emergency_Level) EnumDescriptor() ([]byte, []int) {
	return file_traits_emergency_proto_rawDescGZIP(), []int{0, 0}
}

// An emergency is a life threatening event.
type Emergency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The level of emergency that is active.
	Level Emergency_Level `protobuf:"varint,1,opt,name=level,proto3,enum=smartcore.traits.Emergency_Level" json:"level,omitempty"`
	// A textual description of the reason for the emergency. For example "Fire alarm triggered", "Excessive pressure"
	Reason string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	// The time the emergency level was changed. When writing and absent, and the level has changed, then server time will be used.
	LevelChangeTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=level_change_time,json=levelChangeTime,proto3" json:"level_change_time,omitempty"`
	// Silent emergencies are ongoing but don't require continued or excessive user feedback.
	Silent bool `protobuf:"varint,4,opt,name=silent,proto3" json:"silent,omitempty"`
	// The emergency being fabricated, typically for testing.
	Drill bool `protobuf:"varint,5,opt,name=drill,proto3" json:"drill,omitempty"`
}

func (x *Emergency) Reset() {
	*x = Emergency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_emergency_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Emergency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Emergency) ProtoMessage() {}

func (x *Emergency) ProtoReflect() protoreflect.Message {
	mi := &file_traits_emergency_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Emergency.ProtoReflect.Descriptor instead.
func (*Emergency) Descriptor() ([]byte, []int) {
	return file_traits_emergency_proto_rawDescGZIP(), []int{0}
}

func (x *Emergency) GetLevel() Emergency_Level {
	if x != nil {
		return x.Level
	}
	return Emergency_LEVEL_UNSPECIFIED
}

func (x *Emergency) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *Emergency) GetLevelChangeTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LevelChangeTime
	}
	return nil
}

func (x *Emergency) GetSilent() bool {
	if x != nil {
		return x.Silent
	}
	return false
}

func (x *Emergency) GetDrill() bool {
	if x != nil {
		return x.Drill
	}
	return false
}

// EmergencySupport describes the capabilities of devices implementing this trait
type EmergencySupport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// How a named device supports read/write/pull apis
	ResourceSupport *types.ResourceSupport `protobuf:"bytes,1,opt,name=resource_support,json=resourceSupport,proto3" json:"resource_support,omitempty"`
}

func (x *EmergencySupport) Reset() {
	*x = EmergencySupport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_emergency_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmergencySupport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmergencySupport) ProtoMessage() {}

func (x *EmergencySupport) ProtoReflect() protoreflect.Message {
	mi := &file_traits_emergency_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmergencySupport.ProtoReflect.Descriptor instead.
func (*EmergencySupport) Descriptor() ([]byte, []int) {
	return file_traits_emergency_proto_rawDescGZIP(), []int{1}
}

func (x *EmergencySupport) GetResourceSupport() *types.ResourceSupport {
	if x != nil {
		return x.ResourceSupport
	}
	return nil
}

type GetEmergencyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The state fields to fetch
	ReadMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=read_mask,json=readMask,proto3" json:"read_mask,omitempty"`
}

func (x *GetEmergencyRequest) Reset() {
	*x = GetEmergencyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_emergency_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmergencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmergencyRequest) ProtoMessage() {}

func (x *GetEmergencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_emergency_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmergencyRequest.ProtoReflect.Descriptor instead.
func (*GetEmergencyRequest) Descriptor() ([]byte, []int) {
	return file_traits_emergency_proto_rawDescGZIP(), []int{2}
}

func (x *GetEmergencyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetEmergencyRequest) GetReadMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.ReadMask
	}
	return nil
}

type UpdateEmergencyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Emergency *Emergency `protobuf:"bytes,2,opt,name=emergency,proto3" json:"emergency,omitempty"`
	// The fields relative to state we intend to update
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateEmergencyRequest) Reset() {
	*x = UpdateEmergencyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_emergency_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEmergencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEmergencyRequest) ProtoMessage() {}

func (x *UpdateEmergencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_emergency_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEmergencyRequest.ProtoReflect.Descriptor instead.
func (*UpdateEmergencyRequest) Descriptor() ([]byte, []int) {
	return file_traits_emergency_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateEmergencyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateEmergencyRequest) GetEmergency() *Emergency {
	if x != nil {
		return x.Emergency
	}
	return nil
}

func (x *UpdateEmergencyRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type PullEmergencyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The state fields to fetch
	ReadMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=read_mask,json=readMask,proto3" json:"read_mask,omitempty"`
	// When true the device will only send changes to the resource value.
	// The default behaviour is to send the current value immediately followed by any updates as they happen.
	UpdatesOnly bool `protobuf:"varint,3,opt,name=updates_only,json=updatesOnly,proto3" json:"updates_only,omitempty"`
}

func (x *PullEmergencyRequest) Reset() {
	*x = PullEmergencyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_emergency_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullEmergencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullEmergencyRequest) ProtoMessage() {}

func (x *PullEmergencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_emergency_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullEmergencyRequest.ProtoReflect.Descriptor instead.
func (*PullEmergencyRequest) Descriptor() ([]byte, []int) {
	return file_traits_emergency_proto_rawDescGZIP(), []int{4}
}

func (x *PullEmergencyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PullEmergencyRequest) GetReadMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.ReadMask
	}
	return nil
}

func (x *PullEmergencyRequest) GetUpdatesOnly() bool {
	if x != nil {
		return x.UpdatesOnly
	}
	return false
}

type PullEmergencyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Changes []*PullEmergencyResponse_Change `protobuf:"bytes,1,rep,name=changes,proto3" json:"changes,omitempty"`
}

func (x *PullEmergencyResponse) Reset() {
	*x = PullEmergencyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_emergency_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullEmergencyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullEmergencyResponse) ProtoMessage() {}

func (x *PullEmergencyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_traits_emergency_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullEmergencyResponse.ProtoReflect.Descriptor instead.
func (*PullEmergencyResponse) Descriptor() ([]byte, []int) {
	return file_traits_emergency_proto_rawDescGZIP(), []int{5}
}

func (x *PullEmergencyResponse) GetChanges() []*PullEmergencyResponse_Change {
	if x != nil {
		return x.Changes
	}
	return nil
}

type DescribeEmergencyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the device
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DescribeEmergencyRequest) Reset() {
	*x = DescribeEmergencyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_emergency_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeEmergencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeEmergencyRequest) ProtoMessage() {}

func (x *DescribeEmergencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_emergency_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeEmergencyRequest.ProtoReflect.Descriptor instead.
func (*DescribeEmergencyRequest) Descriptor() ([]byte, []int) {
	return file_traits_emergency_proto_rawDescGZIP(), []int{6}
}

func (x *DescribeEmergencyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PullEmergencyResponse_Change struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name for the device that issued the change
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// when the change occurred
	ChangeTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=change_time,json=changeTime,proto3" json:"change_time,omitempty"`
	// The new value for the emergency.
	Emergency *Emergency `protobuf:"bytes,3,opt,name=emergency,proto3" json:"emergency,omitempty"`
}

func (x *PullEmergencyResponse_Change) Reset() {
	*x = PullEmergencyResponse_Change{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_emergency_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullEmergencyResponse_Change) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullEmergencyResponse_Change) ProtoMessage() {}

func (x *PullEmergencyResponse_Change) ProtoReflect() protoreflect.Message {
	mi := &file_traits_emergency_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullEmergencyResponse_Change.ProtoReflect.Descriptor instead.
func (*PullEmergencyResponse_Change) Descriptor() ([]byte, []int) {
	return file_traits_emergency_proto_rawDescGZIP(), []int{5, 0}
}

func (x *PullEmergencyResponse_Change) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PullEmergencyResponse_Change) GetChangeTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ChangeTime
	}
	return nil
}

func (x *PullEmergencyResponse_Change) GetEmergency() *Emergency {
	if x != nil {
		return x.Emergency
	}
	return nil
}

var File_traits_emergency_proto protoreflect.FileDescriptor

var file_traits_emergency_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2f, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e,
	0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x96, 0x02, 0x0a, 0x09, 0x45, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x37, 0x0a,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x73,
	0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e,
	0x45, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x46,
	0x0a, 0x11, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x6c, 0x65, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x69, 0x6c, 0x65, 0x6e, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x64, 0x72, 0x69, 0x6c, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x64,
	0x72, 0x69, 0x6c, 0x6c, 0x22, 0x42, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x15, 0x0a,
	0x11, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x4d, 0x45,
	0x52, 0x47, 0x45, 0x4e, 0x43, 0x59, 0x10, 0x03, 0x22, 0x5f, 0x0a, 0x10, 0x45, 0x6d, 0x65, 0x72,
	0x67, 0x65, 0x6e, 0x63, 0x79, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x4b, 0x0a, 0x10,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x62, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x45, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d,
	0x61, 0x73, 0x6b, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0xa4, 0x01,
	0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x09,
	0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69,
	0x74, 0x73, 0x2e, 0x45, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x09, 0x65, 0x6d,
	0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4d, 0x61, 0x73, 0x6b, 0x22, 0x86, 0x01, 0x0a, 0x14, 0x50, 0x75, 0x6c, 0x6c, 0x45, 0x6d, 0x65,
	0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x37, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x73, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x4f, 0x6e, 0x6c, 0x79, 0x22, 0xf8, 0x01,
	0x0a, 0x15, 0x50, 0x75, 0x6c, 0x6c, 0x45, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x50, 0x75, 0x6c, 0x6c,
	0x45, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x73, 0x1a, 0x94, 0x01, 0x0a, 0x06, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x39, 0x0a,
	0x09, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61,
	0x69, 0x74, 0x73, 0x2e, 0x45, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x09, 0x65,
	0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x2e, 0x0a, 0x18, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x45, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xa0, 0x02, 0x0a, 0x0c, 0x45, 0x6d, 0x65,
	0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x41, 0x70, 0x69, 0x12, 0x52, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x45, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x25, 0x2e, 0x73, 0x6d, 0x61, 0x72,
	0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x45, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61,
	0x69, 0x74, 0x73, 0x2e, 0x45, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x58, 0x0a,
	0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x28, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61,
	0x69, 0x74, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x65, 0x72, 0x67, 0x65,
	0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x6d, 0x61,
	0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x45, 0x6d,
	0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x62, 0x0a, 0x0d, 0x50, 0x75, 0x6c, 0x6c, 0x45,
	0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x26, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x50, 0x75, 0x6c, 0x6c,
	0x45, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x27, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61,
	0x69, 0x74, 0x73, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x45, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x32, 0x74, 0x0a, 0x0d, 0x45,
	0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x63, 0x0a, 0x11,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x45, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x2a, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72,
	0x61, 0x69, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x45, 0x6d, 0x65,
	0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73,
	0x2e, 0x45, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x42, 0x79, 0x0a, 0x14, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x42, 0x0e, 0x45, 0x6d, 0x65, 0x72, 0x67,
	0x65, 0x6e, 0x63, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x2d, 0x63, 0x6f,
	0x72, 0x65, 0x2d, 0x6f, 0x73, 0x2f, 0x73, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f,
	0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0xaa, 0x02, 0x10, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x74, 0x73, 0xca, 0x02, 0x10, 0x53, 0x6d, 0x61, 0x72,
	0x74, 0x63, 0x6f, 0x72, 0x65, 0x5c, 0x54, 0x72, 0x61, 0x69, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_traits_emergency_proto_rawDescOnce sync.Once
	file_traits_emergency_proto_rawDescData = file_traits_emergency_proto_rawDesc
)

func file_traits_emergency_proto_rawDescGZIP() []byte {
	file_traits_emergency_proto_rawDescOnce.Do(func() {
		file_traits_emergency_proto_rawDescData = protoimpl.X.CompressGZIP(file_traits_emergency_proto_rawDescData)
	})
	return file_traits_emergency_proto_rawDescData
}

var file_traits_emergency_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_traits_emergency_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_traits_emergency_proto_goTypes = []interface{}{
	(Emergency_Level)(0),                 // 0: smartcore.traits.Emergency.Level
	(*Emergency)(nil),                    // 1: smartcore.traits.Emergency
	(*EmergencySupport)(nil),             // 2: smartcore.traits.EmergencySupport
	(*GetEmergencyRequest)(nil),          // 3: smartcore.traits.GetEmergencyRequest
	(*UpdateEmergencyRequest)(nil),       // 4: smartcore.traits.UpdateEmergencyRequest
	(*PullEmergencyRequest)(nil),         // 5: smartcore.traits.PullEmergencyRequest
	(*PullEmergencyResponse)(nil),        // 6: smartcore.traits.PullEmergencyResponse
	(*DescribeEmergencyRequest)(nil),     // 7: smartcore.traits.DescribeEmergencyRequest
	(*PullEmergencyResponse_Change)(nil), // 8: smartcore.traits.PullEmergencyResponse.Change
	(*timestamppb.Timestamp)(nil),        // 9: google.protobuf.Timestamp
	(*types.ResourceSupport)(nil),        // 10: smartcore.types.ResourceSupport
	(*fieldmaskpb.FieldMask)(nil),        // 11: google.protobuf.FieldMask
}
var file_traits_emergency_proto_depIdxs = []int32{
	0,  // 0: smartcore.traits.Emergency.level:type_name -> smartcore.traits.Emergency.Level
	9,  // 1: smartcore.traits.Emergency.level_change_time:type_name -> google.protobuf.Timestamp
	10, // 2: smartcore.traits.EmergencySupport.resource_support:type_name -> smartcore.types.ResourceSupport
	11, // 3: smartcore.traits.GetEmergencyRequest.read_mask:type_name -> google.protobuf.FieldMask
	1,  // 4: smartcore.traits.UpdateEmergencyRequest.emergency:type_name -> smartcore.traits.Emergency
	11, // 5: smartcore.traits.UpdateEmergencyRequest.update_mask:type_name -> google.protobuf.FieldMask
	11, // 6: smartcore.traits.PullEmergencyRequest.read_mask:type_name -> google.protobuf.FieldMask
	8,  // 7: smartcore.traits.PullEmergencyResponse.changes:type_name -> smartcore.traits.PullEmergencyResponse.Change
	9,  // 8: smartcore.traits.PullEmergencyResponse.Change.change_time:type_name -> google.protobuf.Timestamp
	1,  // 9: smartcore.traits.PullEmergencyResponse.Change.emergency:type_name -> smartcore.traits.Emergency
	3,  // 10: smartcore.traits.EmergencyApi.GetEmergency:input_type -> smartcore.traits.GetEmergencyRequest
	4,  // 11: smartcore.traits.EmergencyApi.UpdateEmergency:input_type -> smartcore.traits.UpdateEmergencyRequest
	5,  // 12: smartcore.traits.EmergencyApi.PullEmergency:input_type -> smartcore.traits.PullEmergencyRequest
	7,  // 13: smartcore.traits.EmergencyInfo.DescribeEmergency:input_type -> smartcore.traits.DescribeEmergencyRequest
	1,  // 14: smartcore.traits.EmergencyApi.GetEmergency:output_type -> smartcore.traits.Emergency
	1,  // 15: smartcore.traits.EmergencyApi.UpdateEmergency:output_type -> smartcore.traits.Emergency
	6,  // 16: smartcore.traits.EmergencyApi.PullEmergency:output_type -> smartcore.traits.PullEmergencyResponse
	2,  // 17: smartcore.traits.EmergencyInfo.DescribeEmergency:output_type -> smartcore.traits.EmergencySupport
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_traits_emergency_proto_init() }
func file_traits_emergency_proto_init() {
	if File_traits_emergency_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_traits_emergency_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Emergency); i {
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
		file_traits_emergency_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmergencySupport); i {
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
		file_traits_emergency_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmergencyRequest); i {
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
		file_traits_emergency_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEmergencyRequest); i {
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
		file_traits_emergency_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullEmergencyRequest); i {
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
		file_traits_emergency_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullEmergencyResponse); i {
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
		file_traits_emergency_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeEmergencyRequest); i {
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
		file_traits_emergency_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullEmergencyResponse_Change); i {
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
			RawDescriptor: file_traits_emergency_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_traits_emergency_proto_goTypes,
		DependencyIndexes: file_traits_emergency_proto_depIdxs,
		EnumInfos:         file_traits_emergency_proto_enumTypes,
		MessageInfos:      file_traits_emergency_proto_msgTypes,
	}.Build()
	File_traits_emergency_proto = out.File
	file_traits_emergency_proto_rawDesc = nil
	file_traits_emergency_proto_goTypes = nil
	file_traits_emergency_proto_depIdxs = nil
}
