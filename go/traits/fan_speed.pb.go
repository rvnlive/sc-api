// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.1
// source: traits/fan_speed.proto

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

type FanSpeed_Direction int32

const (
	FanSpeed_DIRECTION_UNSPECIFIED FanSpeed_Direction = 0
	FanSpeed_FORWARD               FanSpeed_Direction = 1
	FanSpeed_BACKWARD              FanSpeed_Direction = 2
)

// Enum value maps for FanSpeed_Direction.
var (
	FanSpeed_Direction_name = map[int32]string{
		0: "DIRECTION_UNSPECIFIED",
		1: "FORWARD",
		2: "BACKWARD",
	}
	FanSpeed_Direction_value = map[string]int32{
		"DIRECTION_UNSPECIFIED": 0,
		"FORWARD":               1,
		"BACKWARD":              2,
	}
)

func (x FanSpeed_Direction) Enum() *FanSpeed_Direction {
	p := new(FanSpeed_Direction)
	*p = x
	return p
}

func (x FanSpeed_Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FanSpeed_Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_traits_fan_speed_proto_enumTypes[0].Descriptor()
}

func (FanSpeed_Direction) Type() protoreflect.EnumType {
	return &file_traits_fan_speed_proto_enumTypes[0]
}

func (x FanSpeed_Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FanSpeed_Direction.Descriptor instead.
func (FanSpeed_Direction) EnumDescriptor() ([]byte, []int) {
	return file_traits_fan_speed_proto_rawDescGZIP(), []int{0, 0}
}

// FanSpeed describes how fast a fan is spinning.
type FanSpeed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Percentage describes the fan speed relative to it's minimum and maximum fan speed.
	// In the range [0,100].
	Percentage float32 `protobuf:"fixed32,1,opt,name=percentage,proto3" json:"percentage,omitempty"`
	// Preset describes the fan speed in terms of qualitative values, low, high, etc.
	// Available values are reported via FanSpeedSupport.presets.
	Preset string `protobuf:"bytes,2,opt,name=preset,proto3" json:"preset,omitempty"`
	// Preset index is the index in FanSpeedSupport.presets of the currently selected preset.
	// This field is primarily used for relative speed adjustment, a value of 1 when the update is relative, adds 1
	// to the active preset, increasing the fan speed by one adjustment. Negative values would decrease the fan speed
	// by an equivalent number of steps.
	PresetIndex int32 `protobuf:"varint,3,opt,name=preset_index,json=presetIndex,proto3" json:"preset_index,omitempty"`
	// Direction describes the direction of the fan rotation.
	// This is typically described relative to the function of the fan.
	// For example a table top fan might describe FORWARD as blowing air out the front of the fan.
	Direction FanSpeed_Direction `protobuf:"varint,4,opt,name=direction,proto3,enum=smartcore.traits.FanSpeed_Direction" json:"direction,omitempty"`
}

func (x *FanSpeed) Reset() {
	*x = FanSpeed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_fan_speed_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FanSpeed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FanSpeed) ProtoMessage() {}

func (x *FanSpeed) ProtoReflect() protoreflect.Message {
	mi := &file_traits_fan_speed_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FanSpeed.ProtoReflect.Descriptor instead.
func (*FanSpeed) Descriptor() ([]byte, []int) {
	return file_traits_fan_speed_proto_rawDescGZIP(), []int{0}
}

func (x *FanSpeed) GetPercentage() float32 {
	if x != nil {
		return x.Percentage
	}
	return 0
}

func (x *FanSpeed) GetPreset() string {
	if x != nil {
		return x.Preset
	}
	return ""
}

func (x *FanSpeed) GetPresetIndex() int32 {
	if x != nil {
		return x.PresetIndex
	}
	return 0
}

func (x *FanSpeed) GetDirection() FanSpeed_Direction {
	if x != nil {
		return x.Direction
	}
	return FanSpeed_DIRECTION_UNSPECIFIED
}

type FanSpeedSupport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// How a named device supports read/write/pull apis
	ResourceSupport *types.ResourceSupport `protobuf:"bytes,1,opt,name=resource_support,json=resourceSupport,proto3" json:"resource_support,omitempty"`
	// Presets is the list of supported fan speed presets, in ascending fan speed order - i.e. index 0 is the slowest.
	Presets []string `protobuf:"bytes,2,rep,name=presets,proto3" json:"presets,omitempty"`
}

func (x *FanSpeedSupport) Reset() {
	*x = FanSpeedSupport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_fan_speed_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FanSpeedSupport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FanSpeedSupport) ProtoMessage() {}

func (x *FanSpeedSupport) ProtoReflect() protoreflect.Message {
	mi := &file_traits_fan_speed_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FanSpeedSupport.ProtoReflect.Descriptor instead.
func (*FanSpeedSupport) Descriptor() ([]byte, []int) {
	return file_traits_fan_speed_proto_rawDescGZIP(), []int{1}
}

func (x *FanSpeedSupport) GetResourceSupport() *types.ResourceSupport {
	if x != nil {
		return x.ResourceSupport
	}
	return nil
}

func (x *FanSpeedSupport) GetPresets() []string {
	if x != nil {
		return x.Presets
	}
	return nil
}

type GetFanSpeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the device to fetch the state for
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Fields to fetch relative to the FanSpeed type
	ReadMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=read_mask,json=readMask,proto3" json:"read_mask,omitempty"`
}

func (x *GetFanSpeedRequest) Reset() {
	*x = GetFanSpeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_fan_speed_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFanSpeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFanSpeedRequest) ProtoMessage() {}

func (x *GetFanSpeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_fan_speed_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFanSpeedRequest.ProtoReflect.Descriptor instead.
func (*GetFanSpeedRequest) Descriptor() ([]byte, []int) {
	return file_traits_fan_speed_proto_rawDescGZIP(), []int{2}
}

func (x *GetFanSpeedRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetFanSpeedRequest) GetReadMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.ReadMask
	}
	return nil
}

type UpdateFanSpeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the device to fetch the state for
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The new fan speed.
	FanSpeed *FanSpeed `protobuf:"bytes,2,opt,name=fan_speed,json=fanSpeed,proto3" json:"fan_speed,omitempty"`
	// Fields to fetch relative to the FanSpeed type
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// Values in fan_speed are relative to their current values.
	// FanSpeed.preset cannot be directly set relatively, use preset_index instead to adjust the selected preset index.
	Relative bool `protobuf:"varint,4,opt,name=relative,proto3" json:"relative,omitempty"`
}

func (x *UpdateFanSpeedRequest) Reset() {
	*x = UpdateFanSpeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_fan_speed_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFanSpeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFanSpeedRequest) ProtoMessage() {}

func (x *UpdateFanSpeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_fan_speed_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFanSpeedRequest.ProtoReflect.Descriptor instead.
func (*UpdateFanSpeedRequest) Descriptor() ([]byte, []int) {
	return file_traits_fan_speed_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateFanSpeedRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateFanSpeedRequest) GetFanSpeed() *FanSpeed {
	if x != nil {
		return x.FanSpeed
	}
	return nil
}

func (x *UpdateFanSpeedRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (x *UpdateFanSpeedRequest) GetRelative() bool {
	if x != nil {
		return x.Relative
	}
	return false
}

type ReverseFanSpeedDirectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the device to reverse the fan speed for.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ReverseFanSpeedDirectionRequest) Reset() {
	*x = ReverseFanSpeedDirectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_fan_speed_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReverseFanSpeedDirectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReverseFanSpeedDirectionRequest) ProtoMessage() {}

func (x *ReverseFanSpeedDirectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_fan_speed_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReverseFanSpeedDirectionRequest.ProtoReflect.Descriptor instead.
func (*ReverseFanSpeedDirectionRequest) Descriptor() ([]byte, []int) {
	return file_traits_fan_speed_proto_rawDescGZIP(), []int{4}
}

func (x *ReverseFanSpeedDirectionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PullFanSpeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the device to fetch the state for
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Fields to fetch relative to the FanSpeed type
	ReadMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=read_mask,json=readMask,proto3" json:"read_mask,omitempty"`
	// When true the device will only send changes to the resource value.
	// The default behaviour is to send the current value immediately followed by any updates as they happen.
	UpdatesOnly bool `protobuf:"varint,3,opt,name=updates_only,json=updatesOnly,proto3" json:"updates_only,omitempty"`
}

func (x *PullFanSpeedRequest) Reset() {
	*x = PullFanSpeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_fan_speed_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullFanSpeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullFanSpeedRequest) ProtoMessage() {}

func (x *PullFanSpeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_fan_speed_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullFanSpeedRequest.ProtoReflect.Descriptor instead.
func (*PullFanSpeedRequest) Descriptor() ([]byte, []int) {
	return file_traits_fan_speed_proto_rawDescGZIP(), []int{5}
}

func (x *PullFanSpeedRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PullFanSpeedRequest) GetReadMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.ReadMask
	}
	return nil
}

func (x *PullFanSpeedRequest) GetUpdatesOnly() bool {
	if x != nil {
		return x.UpdatesOnly
	}
	return false
}

type PullFanSpeedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Changes since the last message.
	Changes []*PullFanSpeedResponse_Change `protobuf:"bytes,1,rep,name=changes,proto3" json:"changes,omitempty"`
}

func (x *PullFanSpeedResponse) Reset() {
	*x = PullFanSpeedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_fan_speed_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullFanSpeedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullFanSpeedResponse) ProtoMessage() {}

func (x *PullFanSpeedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_traits_fan_speed_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullFanSpeedResponse.ProtoReflect.Descriptor instead.
func (*PullFanSpeedResponse) Descriptor() ([]byte, []int) {
	return file_traits_fan_speed_proto_rawDescGZIP(), []int{6}
}

func (x *PullFanSpeedResponse) GetChanges() []*PullFanSpeedResponse_Change {
	if x != nil {
		return x.Changes
	}
	return nil
}

type DescribeFanSpeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the device
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DescribeFanSpeedRequest) Reset() {
	*x = DescribeFanSpeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_fan_speed_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeFanSpeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeFanSpeedRequest) ProtoMessage() {}

func (x *DescribeFanSpeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_fan_speed_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeFanSpeedRequest.ProtoReflect.Descriptor instead.
func (*DescribeFanSpeedRequest) Descriptor() ([]byte, []int) {
	return file_traits_fan_speed_proto_rawDescGZIP(), []int{7}
}

func (x *DescribeFanSpeedRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PullFanSpeedResponse_Change struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the device that issued the change.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// When the change occurred
	ChangeTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=change_time,json=changeTime,proto3" json:"change_time,omitempty"`
	// The new value for the current fan speed.
	FanSpeed *FanSpeed `protobuf:"bytes,3,opt,name=fan_speed,json=fanSpeed,proto3" json:"fan_speed,omitempty"`
}

func (x *PullFanSpeedResponse_Change) Reset() {
	*x = PullFanSpeedResponse_Change{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_fan_speed_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullFanSpeedResponse_Change) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullFanSpeedResponse_Change) ProtoMessage() {}

func (x *PullFanSpeedResponse_Change) ProtoReflect() protoreflect.Message {
	mi := &file_traits_fan_speed_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullFanSpeedResponse_Change.ProtoReflect.Descriptor instead.
func (*PullFanSpeedResponse_Change) Descriptor() ([]byte, []int) {
	return file_traits_fan_speed_proto_rawDescGZIP(), []int{6, 0}
}

func (x *PullFanSpeedResponse_Change) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PullFanSpeedResponse_Change) GetChangeTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ChangeTime
	}
	return nil
}

func (x *PullFanSpeedResponse_Change) GetFanSpeed() *FanSpeed {
	if x != nil {
		return x.FanSpeed
	}
	return nil
}

var File_traits_fan_speed_proto protoreflect.FileDescriptor

var file_traits_fan_speed_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2f, 0x66, 0x61, 0x6e, 0x5f, 0x73, 0x70, 0x65,
	0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xec, 0x01, 0x0a, 0x08, 0x46, 0x61, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72,
	0x65, 0x73, 0x65, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x72, 0x65, 0x73,
	0x65, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x42, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x73, 0x6d, 0x61,
	0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x46, 0x61,
	0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x41, 0x0a, 0x09, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x49, 0x52, 0x45,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x4f, 0x52, 0x57, 0x41, 0x52, 0x44, 0x10, 0x01,
	0x12, 0x0c, 0x0a, 0x08, 0x42, 0x41, 0x43, 0x4b, 0x57, 0x41, 0x52, 0x44, 0x10, 0x02, 0x22, 0x78,
	0x0a, 0x0f, 0x46, 0x61, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x4b, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x6d,
	0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x0f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x73, 0x22, 0x61, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x46,
	0x61, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73,
	0x6b, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0xbd, 0x01, 0x0a, 0x15,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x61, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x66, 0x61, 0x6e,
	0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73,
	0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e,
	0x46, 0x61, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x08, 0x66, 0x61, 0x6e, 0x53, 0x70, 0x65,
	0x65, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d,
	0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x22, 0x35, 0x0a, 0x1f, 0x52,
	0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x46, 0x61, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x85, 0x01, 0x0a, 0x13, 0x50, 0x75, 0x6c, 0x6c, 0x46, 0x61, 0x6e, 0x53, 0x70,
	0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37,
	0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x08, 0x72,
	0x65, 0x61, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x73, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x4f, 0x6e, 0x6c, 0x79, 0x22, 0xf4, 0x01, 0x0a, 0x14, 0x50,
	0x75, 0x6c, 0x6c, 0x46, 0x61, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x46, 0x61, 0x6e, 0x53,
	0x70, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x1a, 0x92, 0x01, 0x0a,
	0x06, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x66, 0x61, 0x6e, 0x5f,
	0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x6d,
	0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x46,
	0x61, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x08, 0x66, 0x61, 0x6e, 0x53, 0x70, 0x65, 0x65,
	0x64, 0x22, 0x2d, 0x0a, 0x17, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x46, 0x61, 0x6e,
	0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x32, 0x81, 0x03, 0x0a, 0x0b, 0x46, 0x61, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x41, 0x70, 0x69,
	0x12, 0x4f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x46, 0x61, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12,
	0x24, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69,
	0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x46, 0x61, 0x6e, 0x53, 0x70, 0x65, 0x65,
	0x64, 0x12, 0x55, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x61, 0x6e, 0x53, 0x70,
	0x65, 0x65, 0x64, 0x12, 0x27, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x61, 0x6e,
	0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73,
	0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e,
	0x46, 0x61, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x5f, 0x0a, 0x0c, 0x50, 0x75, 0x6c, 0x6c,
	0x46, 0x61, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x25, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x50, 0x75, 0x6c, 0x6c,
	0x46, 0x61, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x26, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69,
	0x74, 0x73, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x46, 0x61, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x69, 0x0a, 0x18, 0x52, 0x65, 0x76,
	0x65, 0x72, 0x73, 0x65, 0x46, 0x61, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65,
	0x46, 0x61, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x46, 0x61, 0x6e, 0x53,
	0x70, 0x65, 0x65, 0x64, 0x32, 0x70, 0x0a, 0x0c, 0x46, 0x61, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x60, 0x0a, 0x10, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x46, 0x61, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x29, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x46, 0x61, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x46, 0x61, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x53,
	0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x78, 0x0a, 0x14, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x6d,
	0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x42, 0x0d,
	0x46, 0x61, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x61, 0x72,
	0x74, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2d, 0x6f, 0x73, 0x2f, 0x73, 0x63, 0x2d, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0xaa, 0x02, 0x10, 0x53, 0x6d, 0x61,
	0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x74, 0x73, 0xca, 0x02, 0x10,
	0x53, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x5c, 0x54, 0x72, 0x61, 0x69, 0x74, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_traits_fan_speed_proto_rawDescOnce sync.Once
	file_traits_fan_speed_proto_rawDescData = file_traits_fan_speed_proto_rawDesc
)

func file_traits_fan_speed_proto_rawDescGZIP() []byte {
	file_traits_fan_speed_proto_rawDescOnce.Do(func() {
		file_traits_fan_speed_proto_rawDescData = protoimpl.X.CompressGZIP(file_traits_fan_speed_proto_rawDescData)
	})
	return file_traits_fan_speed_proto_rawDescData
}

var file_traits_fan_speed_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_traits_fan_speed_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_traits_fan_speed_proto_goTypes = []interface{}{
	(FanSpeed_Direction)(0),                 // 0: smartcore.traits.FanSpeed.Direction
	(*FanSpeed)(nil),                        // 1: smartcore.traits.FanSpeed
	(*FanSpeedSupport)(nil),                 // 2: smartcore.traits.FanSpeedSupport
	(*GetFanSpeedRequest)(nil),              // 3: smartcore.traits.GetFanSpeedRequest
	(*UpdateFanSpeedRequest)(nil),           // 4: smartcore.traits.UpdateFanSpeedRequest
	(*ReverseFanSpeedDirectionRequest)(nil), // 5: smartcore.traits.ReverseFanSpeedDirectionRequest
	(*PullFanSpeedRequest)(nil),             // 6: smartcore.traits.PullFanSpeedRequest
	(*PullFanSpeedResponse)(nil),            // 7: smartcore.traits.PullFanSpeedResponse
	(*DescribeFanSpeedRequest)(nil),         // 8: smartcore.traits.DescribeFanSpeedRequest
	(*PullFanSpeedResponse_Change)(nil),     // 9: smartcore.traits.PullFanSpeedResponse.Change
	(*types.ResourceSupport)(nil),           // 10: smartcore.types.ResourceSupport
	(*fieldmaskpb.FieldMask)(nil),           // 11: google.protobuf.FieldMask
	(*timestamppb.Timestamp)(nil),           // 12: google.protobuf.Timestamp
}
var file_traits_fan_speed_proto_depIdxs = []int32{
	0,  // 0: smartcore.traits.FanSpeed.direction:type_name -> smartcore.traits.FanSpeed.Direction
	10, // 1: smartcore.traits.FanSpeedSupport.resource_support:type_name -> smartcore.types.ResourceSupport
	11, // 2: smartcore.traits.GetFanSpeedRequest.read_mask:type_name -> google.protobuf.FieldMask
	1,  // 3: smartcore.traits.UpdateFanSpeedRequest.fan_speed:type_name -> smartcore.traits.FanSpeed
	11, // 4: smartcore.traits.UpdateFanSpeedRequest.update_mask:type_name -> google.protobuf.FieldMask
	11, // 5: smartcore.traits.PullFanSpeedRequest.read_mask:type_name -> google.protobuf.FieldMask
	9,  // 6: smartcore.traits.PullFanSpeedResponse.changes:type_name -> smartcore.traits.PullFanSpeedResponse.Change
	12, // 7: smartcore.traits.PullFanSpeedResponse.Change.change_time:type_name -> google.protobuf.Timestamp
	1,  // 8: smartcore.traits.PullFanSpeedResponse.Change.fan_speed:type_name -> smartcore.traits.FanSpeed
	3,  // 9: smartcore.traits.FanSpeedApi.GetFanSpeed:input_type -> smartcore.traits.GetFanSpeedRequest
	4,  // 10: smartcore.traits.FanSpeedApi.UpdateFanSpeed:input_type -> smartcore.traits.UpdateFanSpeedRequest
	6,  // 11: smartcore.traits.FanSpeedApi.PullFanSpeed:input_type -> smartcore.traits.PullFanSpeedRequest
	5,  // 12: smartcore.traits.FanSpeedApi.ReverseFanSpeedDirection:input_type -> smartcore.traits.ReverseFanSpeedDirectionRequest
	8,  // 13: smartcore.traits.FanSpeedInfo.DescribeFanSpeed:input_type -> smartcore.traits.DescribeFanSpeedRequest
	1,  // 14: smartcore.traits.FanSpeedApi.GetFanSpeed:output_type -> smartcore.traits.FanSpeed
	1,  // 15: smartcore.traits.FanSpeedApi.UpdateFanSpeed:output_type -> smartcore.traits.FanSpeed
	7,  // 16: smartcore.traits.FanSpeedApi.PullFanSpeed:output_type -> smartcore.traits.PullFanSpeedResponse
	1,  // 17: smartcore.traits.FanSpeedApi.ReverseFanSpeedDirection:output_type -> smartcore.traits.FanSpeed
	2,  // 18: smartcore.traits.FanSpeedInfo.DescribeFanSpeed:output_type -> smartcore.traits.FanSpeedSupport
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_traits_fan_speed_proto_init() }
func file_traits_fan_speed_proto_init() {
	if File_traits_fan_speed_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_traits_fan_speed_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FanSpeed); i {
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
		file_traits_fan_speed_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FanSpeedSupport); i {
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
		file_traits_fan_speed_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFanSpeedRequest); i {
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
		file_traits_fan_speed_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFanSpeedRequest); i {
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
		file_traits_fan_speed_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReverseFanSpeedDirectionRequest); i {
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
		file_traits_fan_speed_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullFanSpeedRequest); i {
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
		file_traits_fan_speed_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullFanSpeedResponse); i {
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
		file_traits_fan_speed_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeFanSpeedRequest); i {
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
		file_traits_fan_speed_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullFanSpeedResponse_Change); i {
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
			RawDescriptor: file_traits_fan_speed_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_traits_fan_speed_proto_goTypes,
		DependencyIndexes: file_traits_fan_speed_proto_depIdxs,
		EnumInfos:         file_traits_fan_speed_proto_enumTypes,
		MessageInfos:      file_traits_fan_speed_proto_msgTypes,
	}.Build()
	File_traits_fan_speed_proto = out.File
	file_traits_fan_speed_proto_rawDesc = nil
	file_traits_fan_speed_proto_goTypes = nil
	file_traits_fan_speed_proto_depIdxs = nil
}
