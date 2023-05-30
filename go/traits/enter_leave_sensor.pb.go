// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: traits/enter_leave_sensor.proto

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

// Possible directions.
type EnterLeaveEvent_Direction int32

const (
	EnterLeaveEvent_DIRECTION_UNSPECIFIED EnterLeaveEvent_Direction = 0
	// Something entered.
	EnterLeaveEvent_ENTER EnterLeaveEvent_Direction = 1
	// Something left.
	EnterLeaveEvent_LEAVE EnterLeaveEvent_Direction = 2
)

// Enum value maps for EnterLeaveEvent_Direction.
var (
	EnterLeaveEvent_Direction_name = map[int32]string{
		0: "DIRECTION_UNSPECIFIED",
		1: "ENTER",
		2: "LEAVE",
	}
	EnterLeaveEvent_Direction_value = map[string]int32{
		"DIRECTION_UNSPECIFIED": 0,
		"ENTER":                 1,
		"LEAVE":                 2,
	}
)

func (x EnterLeaveEvent_Direction) Enum() *EnterLeaveEvent_Direction {
	p := new(EnterLeaveEvent_Direction)
	*p = x
	return p
}

func (x EnterLeaveEvent_Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnterLeaveEvent_Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_traits_enter_leave_sensor_proto_enumTypes[0].Descriptor()
}

func (EnterLeaveEvent_Direction) Type() protoreflect.EnumType {
	return &file_traits_enter_leave_sensor_proto_enumTypes[0]
}

func (x EnterLeaveEvent_Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnterLeaveEvent_Direction.Descriptor instead.
func (EnterLeaveEvent_Direction) EnumDescriptor() ([]byte, []int) {
	return file_traits_enter_leave_sensor_proto_rawDescGZIP(), []int{0, 0}
}

// EnterLeaveEvent describes an enter or leave event, and optionally the entity that entered or left.
type EnterLeaveEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Direction describes in which direction movement was detected.
	Direction EnterLeaveEvent_Direction `protobuf:"varint,1,opt,name=direction,proto3,enum=smartcore.traits.EnterLeaveEvent_Direction" json:"direction,omitempty"`
	// Occupant describes the entity either entering or leaving.
	// Optional.
	Occupant *EnterLeaveEvent_Occupant `protobuf:"bytes,2,opt,name=occupant,proto3" json:"occupant,omitempty"`
	// Enter total, if present, indicates how many ENTER events have been recognised by the device.
	EnterTotal *int32 `protobuf:"varint,3,opt,name=enter_total,json=enterTotal,proto3,oneof" json:"enter_total,omitempty"`
	// Leave total, if present, indicates how many LEAVE events have been recognised by the device.
	LeaveTotal *int32 `protobuf:"varint,4,opt,name=leave_total,json=leaveTotal,proto3,oneof" json:"leave_total,omitempty"`
}

func (x *EnterLeaveEvent) Reset() {
	*x = EnterLeaveEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_enter_leave_sensor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnterLeaveEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnterLeaveEvent) ProtoMessage() {}

func (x *EnterLeaveEvent) ProtoReflect() protoreflect.Message {
	mi := &file_traits_enter_leave_sensor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnterLeaveEvent.ProtoReflect.Descriptor instead.
func (*EnterLeaveEvent) Descriptor() ([]byte, []int) {
	return file_traits_enter_leave_sensor_proto_rawDescGZIP(), []int{0}
}

func (x *EnterLeaveEvent) GetDirection() EnterLeaveEvent_Direction {
	if x != nil {
		return x.Direction
	}
	return EnterLeaveEvent_DIRECTION_UNSPECIFIED
}

func (x *EnterLeaveEvent) GetOccupant() *EnterLeaveEvent_Occupant {
	if x != nil {
		return x.Occupant
	}
	return nil
}

func (x *EnterLeaveEvent) GetEnterTotal() int32 {
	if x != nil && x.EnterTotal != nil {
		return *x.EnterTotal
	}
	return 0
}

func (x *EnterLeaveEvent) GetLeaveTotal() int32 {
	if x != nil && x.LeaveTotal != nil {
		return *x.LeaveTotal
	}
	return 0
}

type PullEnterLeaveEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the device to fetch the state for
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Fields to fetch relative to the EnterLeaveEvent type
	ReadMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=read_mask,json=readMask,proto3" json:"read_mask,omitempty"`
	// If true, only updates to the state will be returned.
	// The first event sent when false will be equivalent to GetEnterLeaveEvent, direction will be unspecified.
	UpdatesOnly bool `protobuf:"varint,3,opt,name=updates_only,json=updatesOnly,proto3" json:"updates_only,omitempty"`
}

func (x *PullEnterLeaveEventsRequest) Reset() {
	*x = PullEnterLeaveEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_enter_leave_sensor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullEnterLeaveEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullEnterLeaveEventsRequest) ProtoMessage() {}

func (x *PullEnterLeaveEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_enter_leave_sensor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullEnterLeaveEventsRequest.ProtoReflect.Descriptor instead.
func (*PullEnterLeaveEventsRequest) Descriptor() ([]byte, []int) {
	return file_traits_enter_leave_sensor_proto_rawDescGZIP(), []int{1}
}

func (x *PullEnterLeaveEventsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PullEnterLeaveEventsRequest) GetReadMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.ReadMask
	}
	return nil
}

func (x *PullEnterLeaveEventsRequest) GetUpdatesOnly() bool {
	if x != nil {
		return x.UpdatesOnly
	}
	return false
}

type PullEnterLeaveEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of changes
	Changes []*PullEnterLeaveEventsResponse_Change `protobuf:"bytes,1,rep,name=changes,proto3" json:"changes,omitempty"`
}

func (x *PullEnterLeaveEventsResponse) Reset() {
	*x = PullEnterLeaveEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_enter_leave_sensor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullEnterLeaveEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullEnterLeaveEventsResponse) ProtoMessage() {}

func (x *PullEnterLeaveEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_traits_enter_leave_sensor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullEnterLeaveEventsResponse.ProtoReflect.Descriptor instead.
func (*PullEnterLeaveEventsResponse) Descriptor() ([]byte, []int) {
	return file_traits_enter_leave_sensor_proto_rawDescGZIP(), []int{2}
}

func (x *PullEnterLeaveEventsResponse) GetChanges() []*PullEnterLeaveEventsResponse_Change {
	if x != nil {
		return x.Changes
	}
	return nil
}

type GetEnterLeaveEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the device to fetch the state for
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Fields to fetch relative to the EnterLeaveEvent type
	ReadMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=read_mask,json=readMask,proto3" json:"read_mask,omitempty"`
}

func (x *GetEnterLeaveEventRequest) Reset() {
	*x = GetEnterLeaveEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_enter_leave_sensor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEnterLeaveEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEnterLeaveEventRequest) ProtoMessage() {}

func (x *GetEnterLeaveEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_enter_leave_sensor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEnterLeaveEventRequest.ProtoReflect.Descriptor instead.
func (*GetEnterLeaveEventRequest) Descriptor() ([]byte, []int) {
	return file_traits_enter_leave_sensor_proto_rawDescGZIP(), []int{3}
}

func (x *GetEnterLeaveEventRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetEnterLeaveEventRequest) GetReadMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.ReadMask
	}
	return nil
}

type ResetEnterLeaveTotalsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the device to fetch the state for
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ResetEnterLeaveTotalsRequest) Reset() {
	*x = ResetEnterLeaveTotalsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_enter_leave_sensor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetEnterLeaveTotalsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetEnterLeaveTotalsRequest) ProtoMessage() {}

func (x *ResetEnterLeaveTotalsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_enter_leave_sensor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetEnterLeaveTotalsRequest.ProtoReflect.Descriptor instead.
func (*ResetEnterLeaveTotalsRequest) Descriptor() ([]byte, []int) {
	return file_traits_enter_leave_sensor_proto_rawDescGZIP(), []int{4}
}

func (x *ResetEnterLeaveTotalsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ResetEnterLeaveTotalsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResetEnterLeaveTotalsResponse) Reset() {
	*x = ResetEnterLeaveTotalsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_enter_leave_sensor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetEnterLeaveTotalsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetEnterLeaveTotalsResponse) ProtoMessage() {}

func (x *ResetEnterLeaveTotalsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_traits_enter_leave_sensor_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetEnterLeaveTotalsResponse.ProtoReflect.Descriptor instead.
func (*ResetEnterLeaveTotalsResponse) Descriptor() ([]byte, []int) {
	return file_traits_enter_leave_sensor_proto_rawDescGZIP(), []int{5}
}

// An Occupant is some entity that triggers the enter or leave event.
// It could be an object, like a car, or a person or animal.
type EnterLeaveEvent_Occupant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A unique name for this occupant as measured by the device.
	// Name allows correlation between enter and leave events.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// An official title for this occupant.
	Title string `protobuf:"bytes,10,opt,name=title,proto3" json:"title,omitempty"`
	// A recognisable display name for this occupant.
	DisplayName string `protobuf:"bytes,11,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// A picture of the occupant.
	Picture *types.Image `protobuf:"bytes,12,opt,name=picture,proto3" json:"picture,omitempty"`
	// A url that resolves to more information on this occupant.
	Url string `protobuf:"bytes,13,opt,name=url,proto3" json:"url,omitempty"`
	// An email address for the occupant.
	Email string `protobuf:"bytes,14,opt,name=email,proto3" json:"email,omitempty"`
	// IDs holds external occupant ids.
	// For example this might hold an id representing a person in an access control system.
	// The map key should uniquely represent the domain for the id, for example "my-access-system/user-id", the use of uris is not
	// required, values should be the id of this consumable in that domain, for example "sma81r6t1c5o3r58e1-3r8u16l1es".
	Ids map[string]string `protobuf:"bytes,100,rep,name=ids,proto3" json:"ids,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Additional properties associated with the occupant.
	// Typically displayed to a user to help them to distinguish between similar occupants.
	More map[string]string `protobuf:"bytes,101,rep,name=more,proto3" json:"more,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *EnterLeaveEvent_Occupant) Reset() {
	*x = EnterLeaveEvent_Occupant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_enter_leave_sensor_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnterLeaveEvent_Occupant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnterLeaveEvent_Occupant) ProtoMessage() {}

func (x *EnterLeaveEvent_Occupant) ProtoReflect() protoreflect.Message {
	mi := &file_traits_enter_leave_sensor_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnterLeaveEvent_Occupant.ProtoReflect.Descriptor instead.
func (*EnterLeaveEvent_Occupant) Descriptor() ([]byte, []int) {
	return file_traits_enter_leave_sensor_proto_rawDescGZIP(), []int{0, 0}
}

func (x *EnterLeaveEvent_Occupant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EnterLeaveEvent_Occupant) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *EnterLeaveEvent_Occupant) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *EnterLeaveEvent_Occupant) GetPicture() *types.Image {
	if x != nil {
		return x.Picture
	}
	return nil
}

func (x *EnterLeaveEvent_Occupant) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *EnterLeaveEvent_Occupant) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *EnterLeaveEvent_Occupant) GetIds() map[string]string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *EnterLeaveEvent_Occupant) GetMore() map[string]string {
	if x != nil {
		return x.More
	}
	return nil
}

// A change, i.e. a new event, has been recorded.
type PullEnterLeaveEventsResponse_Change struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name for the device that issued the change
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// when the change occurred
	ChangeTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=change_time,json=changeTime,proto3" json:"change_time,omitempty"`
	// The new value for the state
	EnterLeaveEvent *EnterLeaveEvent `protobuf:"bytes,3,opt,name=enter_leave_event,json=enterLeaveEvent,proto3" json:"enter_leave_event,omitempty"`
}

func (x *PullEnterLeaveEventsResponse_Change) Reset() {
	*x = PullEnterLeaveEventsResponse_Change{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_enter_leave_sensor_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullEnterLeaveEventsResponse_Change) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullEnterLeaveEventsResponse_Change) ProtoMessage() {}

func (x *PullEnterLeaveEventsResponse_Change) ProtoReflect() protoreflect.Message {
	mi := &file_traits_enter_leave_sensor_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullEnterLeaveEventsResponse_Change.ProtoReflect.Descriptor instead.
func (*PullEnterLeaveEventsResponse_Change) Descriptor() ([]byte, []int) {
	return file_traits_enter_leave_sensor_proto_rawDescGZIP(), []int{2, 0}
}

func (x *PullEnterLeaveEventsResponse_Change) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PullEnterLeaveEventsResponse_Change) GetChangeTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ChangeTime
	}
	return nil
}

func (x *PullEnterLeaveEventsResponse_Change) GetEnterLeaveEvent() *EnterLeaveEvent {
	if x != nil {
		return x.EnterLeaveEvent
	}
	return nil
}

var File_traits_enter_leave_sensor_proto protoreflect.FileDescriptor

var file_traits_enter_leave_sensor_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x6c,
	0x65, 0x61, 0x76, 0x65, 0x5f, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61,
	0x69, 0x74, 0x73, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x06, 0x0a, 0x0f, 0x45, 0x6e,
	0x74, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x49, 0x0a,
	0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x2b, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61,
	0x69, 0x74, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x08, 0x6f, 0x63, 0x63, 0x75,
	0x70, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x6d, 0x61,
	0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x45, 0x6e,
	0x74, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x4f, 0x63,
	0x63, 0x75, 0x70, 0x61, 0x6e, 0x74, 0x52, 0x08, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e, 0x74,
	0x12, 0x24, 0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x6c, 0x65, 0x61, 0x76, 0x65, 0x5f,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x0a, 0x6c,
	0x65, 0x61, 0x76, 0x65, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x88, 0x01, 0x01, 0x1a, 0xb3, 0x03, 0x0a,
	0x08, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52,
	0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x45, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x64, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e,
	0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73,
	0x2e, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e, 0x74, 0x2e, 0x49, 0x64, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x48, 0x0a, 0x04, 0x6d, 0x6f, 0x72, 0x65, 0x18,
	0x65, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x4c, 0x65,
	0x61, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e,
	0x74, 0x2e, 0x4d, 0x6f, 0x72, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6d, 0x6f, 0x72,
	0x65, 0x1a, 0x36, 0x0a, 0x08, 0x49, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x37, 0x0a, 0x09, 0x4d, 0x6f, 0x72,
	0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x3c, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x19, 0x0a, 0x15, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4e,
	0x54, 0x45, 0x52, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x45, 0x41, 0x56, 0x45, 0x10, 0x02,
	0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6c, 0x65, 0x61, 0x76, 0x65, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x22, 0x8d, 0x01, 0x0a, 0x1b, 0x50, 0x75, 0x6c, 0x6c, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x4c, 0x65,
	0x61, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d,
	0x61, 0x73, 0x6b, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x21, 0x0a,
	0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x4f, 0x6e, 0x6c, 0x79,
	0x22, 0x9a, 0x02, 0x0a, 0x1c, 0x50, 0x75, 0x6c, 0x6c, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x4c, 0x65,
	0x61, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4f, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x35, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74,
	0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x4c,
	0x65, 0x61, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x73, 0x1a, 0xa8, 0x01, 0x0a, 0x06, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4d,
	0x0a, 0x11, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x6c, 0x65, 0x61, 0x76, 0x65, 0x5f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x6d, 0x61, 0x72,
	0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x45, 0x6e, 0x74,
	0x65, 0x72, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x0f, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x68, 0x0a,
	0x19, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37,
	0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x08, 0x72,
	0x65, 0x61, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x32, 0x0a, 0x1c, 0x52, 0x65, 0x73, 0x65, 0x74,
	0x45, 0x6e, 0x74, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1f, 0x0a, 0x1d, 0x52,
	0x65, 0x73, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xee, 0x02, 0x0a,
	0x13, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x53, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x41, 0x70, 0x69, 0x12, 0x77, 0x0a, 0x14, 0x50, 0x75, 0x6c, 0x6c, 0x45, 0x6e, 0x74, 0x65,
	0x72, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2d, 0x2e, 0x73,
	0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e,
	0x50, 0x75, 0x6c, 0x6c, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x73, 0x6d,
	0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x50,
	0x75, 0x6c, 0x6c, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x64, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x2b, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x4c,
	0x65, 0x61, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61,
	0x69, 0x74, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x78, 0x0a, 0x15, 0x52, 0x65, 0x73, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x65,
	0x72, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x73, 0x12, 0x2e, 0x2e, 0x73,
	0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e,
	0x52, 0x65, 0x73, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x73,
	0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e,
	0x52, 0x65, 0x73, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x16, 0x0a,
	0x14, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x53, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x80, 0x01, 0x0a, 0x14, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x6d,
	0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x42, 0x15,
	0x45, 0x6e, 0x74, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2d, 0x6f,
	0x73, 0x2f, 0x73, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x69,
	0x74, 0x73, 0xaa, 0x02, 0x10, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54,
	0x72, 0x61, 0x69, 0x74, 0x73, 0xca, 0x02, 0x10, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72,
	0x65, 0x5c, 0x54, 0x72, 0x61, 0x69, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_traits_enter_leave_sensor_proto_rawDescOnce sync.Once
	file_traits_enter_leave_sensor_proto_rawDescData = file_traits_enter_leave_sensor_proto_rawDesc
)

func file_traits_enter_leave_sensor_proto_rawDescGZIP() []byte {
	file_traits_enter_leave_sensor_proto_rawDescOnce.Do(func() {
		file_traits_enter_leave_sensor_proto_rawDescData = protoimpl.X.CompressGZIP(file_traits_enter_leave_sensor_proto_rawDescData)
	})
	return file_traits_enter_leave_sensor_proto_rawDescData
}

var file_traits_enter_leave_sensor_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_traits_enter_leave_sensor_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_traits_enter_leave_sensor_proto_goTypes = []interface{}{
	(EnterLeaveEvent_Direction)(0),              // 0: smartcore.traits.EnterLeaveEvent.Direction
	(*EnterLeaveEvent)(nil),                     // 1: smartcore.traits.EnterLeaveEvent
	(*PullEnterLeaveEventsRequest)(nil),         // 2: smartcore.traits.PullEnterLeaveEventsRequest
	(*PullEnterLeaveEventsResponse)(nil),        // 3: smartcore.traits.PullEnterLeaveEventsResponse
	(*GetEnterLeaveEventRequest)(nil),           // 4: smartcore.traits.GetEnterLeaveEventRequest
	(*ResetEnterLeaveTotalsRequest)(nil),        // 5: smartcore.traits.ResetEnterLeaveTotalsRequest
	(*ResetEnterLeaveTotalsResponse)(nil),       // 6: smartcore.traits.ResetEnterLeaveTotalsResponse
	(*EnterLeaveEvent_Occupant)(nil),            // 7: smartcore.traits.EnterLeaveEvent.Occupant
	nil,                                         // 8: smartcore.traits.EnterLeaveEvent.Occupant.IdsEntry
	nil,                                         // 9: smartcore.traits.EnterLeaveEvent.Occupant.MoreEntry
	(*PullEnterLeaveEventsResponse_Change)(nil), // 10: smartcore.traits.PullEnterLeaveEventsResponse.Change
	(*fieldmaskpb.FieldMask)(nil),               // 11: google.protobuf.FieldMask
	(*types.Image)(nil),                         // 12: smartcore.types.Image
	(*timestamppb.Timestamp)(nil),               // 13: google.protobuf.Timestamp
}
var file_traits_enter_leave_sensor_proto_depIdxs = []int32{
	0,  // 0: smartcore.traits.EnterLeaveEvent.direction:type_name -> smartcore.traits.EnterLeaveEvent.Direction
	7,  // 1: smartcore.traits.EnterLeaveEvent.occupant:type_name -> smartcore.traits.EnterLeaveEvent.Occupant
	11, // 2: smartcore.traits.PullEnterLeaveEventsRequest.read_mask:type_name -> google.protobuf.FieldMask
	10, // 3: smartcore.traits.PullEnterLeaveEventsResponse.changes:type_name -> smartcore.traits.PullEnterLeaveEventsResponse.Change
	11, // 4: smartcore.traits.GetEnterLeaveEventRequest.read_mask:type_name -> google.protobuf.FieldMask
	12, // 5: smartcore.traits.EnterLeaveEvent.Occupant.picture:type_name -> smartcore.types.Image
	8,  // 6: smartcore.traits.EnterLeaveEvent.Occupant.ids:type_name -> smartcore.traits.EnterLeaveEvent.Occupant.IdsEntry
	9,  // 7: smartcore.traits.EnterLeaveEvent.Occupant.more:type_name -> smartcore.traits.EnterLeaveEvent.Occupant.MoreEntry
	13, // 8: smartcore.traits.PullEnterLeaveEventsResponse.Change.change_time:type_name -> google.protobuf.Timestamp
	1,  // 9: smartcore.traits.PullEnterLeaveEventsResponse.Change.enter_leave_event:type_name -> smartcore.traits.EnterLeaveEvent
	2,  // 10: smartcore.traits.EnterLeaveSensorApi.PullEnterLeaveEvents:input_type -> smartcore.traits.PullEnterLeaveEventsRequest
	4,  // 11: smartcore.traits.EnterLeaveSensorApi.GetEnterLeaveEvent:input_type -> smartcore.traits.GetEnterLeaveEventRequest
	5,  // 12: smartcore.traits.EnterLeaveSensorApi.ResetEnterLeaveTotals:input_type -> smartcore.traits.ResetEnterLeaveTotalsRequest
	3,  // 13: smartcore.traits.EnterLeaveSensorApi.PullEnterLeaveEvents:output_type -> smartcore.traits.PullEnterLeaveEventsResponse
	1,  // 14: smartcore.traits.EnterLeaveSensorApi.GetEnterLeaveEvent:output_type -> smartcore.traits.EnterLeaveEvent
	6,  // 15: smartcore.traits.EnterLeaveSensorApi.ResetEnterLeaveTotals:output_type -> smartcore.traits.ResetEnterLeaveTotalsResponse
	13, // [13:16] is the sub-list for method output_type
	10, // [10:13] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_traits_enter_leave_sensor_proto_init() }
func file_traits_enter_leave_sensor_proto_init() {
	if File_traits_enter_leave_sensor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_traits_enter_leave_sensor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnterLeaveEvent); i {
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
		file_traits_enter_leave_sensor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullEnterLeaveEventsRequest); i {
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
		file_traits_enter_leave_sensor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullEnterLeaveEventsResponse); i {
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
		file_traits_enter_leave_sensor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEnterLeaveEventRequest); i {
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
		file_traits_enter_leave_sensor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetEnterLeaveTotalsRequest); i {
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
		file_traits_enter_leave_sensor_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetEnterLeaveTotalsResponse); i {
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
		file_traits_enter_leave_sensor_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnterLeaveEvent_Occupant); i {
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
		file_traits_enter_leave_sensor_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullEnterLeaveEventsResponse_Change); i {
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
	file_traits_enter_leave_sensor_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_traits_enter_leave_sensor_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_traits_enter_leave_sensor_proto_goTypes,
		DependencyIndexes: file_traits_enter_leave_sensor_proto_depIdxs,
		EnumInfos:         file_traits_enter_leave_sensor_proto_enumTypes,
		MessageInfos:      file_traits_enter_leave_sensor_proto_msgTypes,
	}.Build()
	File_traits_enter_leave_sensor_proto = out.File
	file_traits_enter_leave_sensor_proto_rawDesc = nil
	file_traits_enter_leave_sensor_proto_goTypes = nil
	file_traits_enter_leave_sensor_proto_depIdxs = nil
}
