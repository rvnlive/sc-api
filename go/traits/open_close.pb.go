// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: traits/open_close.proto

package traits

import (
	types "git.vanti.co.uk/smartcore/sc-api/go/types"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Possible directions the device can open/close
type OpenCloseDirection int32

const (
	OpenCloseDirection_UNSPECIFIED OpenCloseDirection = 0
	OpenCloseDirection_UP          OpenCloseDirection = 1
	OpenCloseDirection_DOWN        OpenCloseDirection = 2
	OpenCloseDirection_LEFT        OpenCloseDirection = 3
	OpenCloseDirection_RIGHT       OpenCloseDirection = 4
	OpenCloseDirection_IN          OpenCloseDirection = 5
	OpenCloseDirection_OUT         OpenCloseDirection = 6
)

// Enum value maps for OpenCloseDirection.
var (
	OpenCloseDirection_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UP",
		2: "DOWN",
		3: "LEFT",
		4: "RIGHT",
		5: "IN",
		6: "OUT",
	}
	OpenCloseDirection_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UP":          1,
		"DOWN":        2,
		"LEFT":        3,
		"RIGHT":       4,
		"IN":          5,
		"OUT":         6,
	}
)

func (x OpenCloseDirection) Enum() *OpenCloseDirection {
	p := new(OpenCloseDirection)
	*p = x
	return p
}

func (x OpenCloseDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OpenCloseDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_traits_open_close_proto_enumTypes[0].Descriptor()
}

func (OpenCloseDirection) Type() protoreflect.EnumType {
	return &file_traits_open_close_proto_enumTypes[0]
}

func (x OpenCloseDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OpenCloseDirection.Descriptor instead.
func (OpenCloseDirection) EnumDescriptor() ([]byte, []int) {
	return file_traits_open_close_proto_rawDescGZIP(), []int{0}
}

// Trait.attribute for OpenClose devices
type OpenCloseAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Define how the value for positions behave. The bounds for these attributes will be absent and always in the range
	// 0-100 to represent a percentage
	PositionAttributes *types.FloatAttributes `protobuf:"bytes,1,opt,name=position_attributes,json=positionAttributes,proto3" json:"position_attributes,omitempty"`
	// The supported directions for the device. If none then assume only a single direction is supported
	Directions []OpenCloseDirection `protobuf:"varint,2,rep,packed,name=directions,proto3,enum=smartcore.traits.OpenCloseDirection" json:"directions,omitempty"`
	// Does the device support the Stop method
	SupportsStop bool `protobuf:"varint,3,opt,name=supports_stop,json=supportsStop,proto3" json:"supports_stop,omitempty"`
}

func (x *OpenCloseAttributes) Reset() {
	*x = OpenCloseAttributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_open_close_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenCloseAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenCloseAttributes) ProtoMessage() {}

func (x *OpenCloseAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_traits_open_close_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenCloseAttributes.ProtoReflect.Descriptor instead.
func (*OpenCloseAttributes) Descriptor() ([]byte, []int) {
	return file_traits_open_close_proto_rawDescGZIP(), []int{0}
}

func (x *OpenCloseAttributes) GetPositionAttributes() *types.FloatAttributes {
	if x != nil {
		return x.PositionAttributes
	}
	return nil
}

func (x *OpenCloseAttributes) GetDirections() []OpenCloseDirection {
	if x != nil {
		return x.Directions
	}
	return nil
}

func (x *OpenCloseAttributes) GetSupportsStop() bool {
	if x != nil {
		return x.SupportsStop
	}
	return false
}

// All open/closable elements for the device
type OpenClosePositions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The state of each of the devices open-able elements. Most devices will only have a single value here as they can
	// only be opened from a single direction.
	States []*OpenClosePosition `protobuf:"bytes,1,rep,name=states,proto3" json:"states,omitempty"`
}

func (x *OpenClosePositions) Reset() {
	*x = OpenClosePositions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_open_close_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenClosePositions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenClosePositions) ProtoMessage() {}

func (x *OpenClosePositions) ProtoReflect() protoreflect.Message {
	mi := &file_traits_open_close_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenClosePositions.ProtoReflect.Descriptor instead.
func (*OpenClosePositions) Descriptor() ([]byte, []int) {
	return file_traits_open_close_proto_rawDescGZIP(), []int{1}
}

func (x *OpenClosePositions) GetStates() []*OpenClosePosition {
	if x != nil {
		return x.States
	}
	return nil
}

// Defines the position of one open close element for the device
type OpenClosePosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 0-100, the current value for the open close position of the device in this direction
	PositionPercent *types.FloatVar `protobuf:"bytes,1,opt,name=position_percent,json=positionPercent,proto3" json:"position_percent,omitempty"`
	// Optional direction for devices that support multiple open/close directions
	Direction OpenCloseDirection `protobuf:"varint,2,opt,name=direction,proto3,enum=smartcore.traits.OpenCloseDirection" json:"direction,omitempty"`
}

func (x *OpenClosePosition) Reset() {
	*x = OpenClosePosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_open_close_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenClosePosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenClosePosition) ProtoMessage() {}

func (x *OpenClosePosition) ProtoReflect() protoreflect.Message {
	mi := &file_traits_open_close_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenClosePosition.ProtoReflect.Descriptor instead.
func (*OpenClosePosition) Descriptor() ([]byte, []int) {
	return file_traits_open_close_proto_rawDescGZIP(), []int{2}
}

func (x *OpenClosePosition) GetPositionPercent() *types.FloatVar {
	if x != nil {
		return x.PositionPercent
	}
	return nil
}

func (x *OpenClosePosition) GetDirection() OpenCloseDirection {
	if x != nil {
		return x.Direction
	}
	return OpenCloseDirection_UNSPECIFIED
}

type GetOpenClosePositionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the device to get the state for
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetOpenClosePositionsRequest) Reset() {
	*x = GetOpenClosePositionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_open_close_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOpenClosePositionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOpenClosePositionsRequest) ProtoMessage() {}

func (x *GetOpenClosePositionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_open_close_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOpenClosePositionsRequest.ProtoReflect.Descriptor instead.
func (*GetOpenClosePositionsRequest) Descriptor() ([]byte, []int) {
	return file_traits_open_close_proto_rawDescGZIP(), []int{3}
}

func (x *GetOpenClosePositionsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateOpenClosePositionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the name of the device to update the states for
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// target states for each directional element of the device. If you don't want to change an open direction don't
	// include it in the states list.
	States *OpenClosePositions `protobuf:"bytes,2,opt,name=states,proto3" json:"states,omitempty"`
	// indicate whether the state change is a delta or absolute value
	Delta bool `protobuf:"varint,3,opt,name=delta,proto3" json:"delta,omitempty"`
}

func (x *UpdateOpenClosePositionsRequest) Reset() {
	*x = UpdateOpenClosePositionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_open_close_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOpenClosePositionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOpenClosePositionsRequest) ProtoMessage() {}

func (x *UpdateOpenClosePositionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_open_close_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOpenClosePositionsRequest.ProtoReflect.Descriptor instead.
func (*UpdateOpenClosePositionsRequest) Descriptor() ([]byte, []int) {
	return file_traits_open_close_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateOpenClosePositionsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateOpenClosePositionsRequest) GetStates() *OpenClosePositions {
	if x != nil {
		return x.States
	}
	return nil
}

func (x *UpdateOpenClosePositionsRequest) GetDelta() bool {
	if x != nil {
		return x.Delta
	}
	return false
}

type StopOpenCloseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the name of the device to stop
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *StopOpenCloseRequest) Reset() {
	*x = StopOpenCloseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_open_close_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopOpenCloseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopOpenCloseRequest) ProtoMessage() {}

func (x *StopOpenCloseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_open_close_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopOpenCloseRequest.ProtoReflect.Descriptor instead.
func (*StopOpenCloseRequest) Descriptor() ([]byte, []int) {
	return file_traits_open_close_proto_rawDescGZIP(), []int{5}
}

func (x *StopOpenCloseRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request to begin a subscription for open close state changes
type PullOpenClosePositionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the device we want events from
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Indicate whether we want to be notified of values while tweening or not. The device may ignore this depending on
	// capabilities
	ExcludeTweening bool `protobuf:"varint,2,opt,name=exclude_tweening,json=excludeTweening,proto3" json:"exclude_tweening,omitempty"`
}

func (x *PullOpenClosePositionsRequest) Reset() {
	*x = PullOpenClosePositionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_open_close_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullOpenClosePositionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullOpenClosePositionsRequest) ProtoMessage() {}

func (x *PullOpenClosePositionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_open_close_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullOpenClosePositionsRequest.ProtoReflect.Descriptor instead.
func (*PullOpenClosePositionsRequest) Descriptor() ([]byte, []int) {
	return file_traits_open_close_proto_rawDescGZIP(), []int{6}
}

func (x *PullOpenClosePositionsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PullOpenClosePositionsRequest) GetExcludeTweening() bool {
	if x != nil {
		return x.ExcludeTweening
	}
	return false
}

// A response as part of the stream of changes to the range value
type PullOpenClosePositionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of changes
	Changes []*PullOpenClosePositionsResponse_Change `protobuf:"bytes,1,rep,name=changes,proto3" json:"changes,omitempty"`
}

func (x *PullOpenClosePositionsResponse) Reset() {
	*x = PullOpenClosePositionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_open_close_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullOpenClosePositionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullOpenClosePositionsResponse) ProtoMessage() {}

func (x *PullOpenClosePositionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_traits_open_close_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullOpenClosePositionsResponse.ProtoReflect.Descriptor instead.
func (*PullOpenClosePositionsResponse) Descriptor() ([]byte, []int) {
	return file_traits_open_close_proto_rawDescGZIP(), []int{7}
}

func (x *PullOpenClosePositionsResponse) GetChanges() []*PullOpenClosePositionsResponse_Change {
	if x != nil {
		return x.Changes
	}
	return nil
}

// A change to a single devices open close state value
type PullOpenClosePositionsResponse_Change struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name for the device that issued the change
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// when the change occurred
	ChangeTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=change_time,json=changeTime,proto3" json:"change_time,omitempty"`
	// The new value for the state
	State *OpenClosePositions `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *PullOpenClosePositionsResponse_Change) Reset() {
	*x = PullOpenClosePositionsResponse_Change{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_open_close_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullOpenClosePositionsResponse_Change) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullOpenClosePositionsResponse_Change) ProtoMessage() {}

func (x *PullOpenClosePositionsResponse_Change) ProtoReflect() protoreflect.Message {
	mi := &file_traits_open_close_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullOpenClosePositionsResponse_Change.ProtoReflect.Descriptor instead.
func (*PullOpenClosePositionsResponse_Change) Descriptor() ([]byte, []int) {
	return file_traits_open_close_proto_rawDescGZIP(), []int{7, 0}
}

func (x *PullOpenClosePositionsResponse_Change) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PullOpenClosePositionsResponse_Change) GetChangeTime() *timestamp.Timestamp {
	if x != nil {
		return x.ChangeTime
	}
	return nil
}

func (x *PullOpenClosePositionsResponse_Change) GetState() *OpenClosePositions {
	if x != nil {
		return x.State
	}
	return nil
}

var File_traits_open_close_proto protoreflect.FileDescriptor

var file_traits_open_close_proto_rawDesc = []byte{
	0x0a, 0x17, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x63, 0x6c,
	0x6f, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73, 0x6d, 0x61, 0x72, 0x74,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xd7, 0x01, 0x0a, 0x13, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x55, 0x0a, 0x13, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x46, 0x6c, 0x6f, 0x61,
	0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x12, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12,
	0x44, 0x0a, 0x0a, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x5f, 0x73, 0x74, 0x6f, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x73, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x53, 0x74, 0x6f, 0x70, 0x22, 0x51, 0x0a, 0x12, 0x4f, 0x70,
	0x65, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x3b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61,
	0x69, 0x74, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x22, 0xa1, 0x01,
	0x0a, 0x11, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x48, 0x0a, 0x10, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x72, 0x52, 0x0f, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x42, 0x0a,
	0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x24, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61,
	0x69, 0x74, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x32, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6c, 0x6f, 0x73,
	0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x89, 0x01, 0x0a, 0x1f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73,
	0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x64,
	0x65, 0x6c, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x74,
	0x61, 0x22, 0x2a, 0x0a, 0x14, 0x53, 0x74, 0x6f, 0x70, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5e, 0x0a,
	0x1d, 0x50, 0x75, 0x6c, 0x6c, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x74, 0x77,
	0x65, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x65, 0x78,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x54, 0x77, 0x65, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x22, 0x8b, 0x02,
	0x0a, 0x1e, 0x50, 0x75, 0x6c, 0x6c, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x51, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x37, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72,
	0x61, 0x69, 0x74, 0x73, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x73, 0x1a, 0x95, 0x01, 0x0a, 0x06, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x3a, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74,
	0x73, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2a, 0x5d, 0x0a, 0x12, 0x4f,
	0x70, 0x65, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x55, 0x50, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f,
	0x57, 0x4e, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x45, 0x46, 0x54, 0x10, 0x03, 0x12, 0x09,
	0x0a, 0x05, 0x52, 0x49, 0x47, 0x48, 0x54, 0x10, 0x04, 0x12, 0x06, 0x0a, 0x02, 0x49, 0x4e, 0x10,
	0x05, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x55, 0x54, 0x10, 0x06, 0x32, 0xac, 0x03, 0x0a, 0x0c, 0x4f,
	0x70, 0x65, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x41, 0x70, 0x69, 0x12, 0x64, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2e, 0x2e, 0x73, 0x6d,
	0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x6d,
	0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x4f,
	0x70, 0x65, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x6a, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x31, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x70,
	0x65, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x43,
	0x6c, 0x6f, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x54, 0x0a,
	0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x26, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x4f, 0x70, 0x65,
	0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73,
	0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x74, 0x0a, 0x0d, 0x50, 0x75, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2f, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x4f, 0x70, 0x65, 0x6e,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x4f, 0x70, 0x65,
	0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x7a, 0x0a, 0x14, 0x64, 0x65, 0x76,
	0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74,
	0x73, 0x42, 0x0e, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x2e, 0x76, 0x61, 0x6e, 0x74, 0x69, 0x2e, 0x63,
	0x6f, 0x2e, 0x75, 0x6b, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73,
	0x63, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0xaa,
	0x02, 0x10, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x69,
	0x74, 0x73, 0xca, 0x02, 0x10, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x5c, 0x54,
	0x72, 0x61, 0x69, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_traits_open_close_proto_rawDescOnce sync.Once
	file_traits_open_close_proto_rawDescData = file_traits_open_close_proto_rawDesc
)

func file_traits_open_close_proto_rawDescGZIP() []byte {
	file_traits_open_close_proto_rawDescOnce.Do(func() {
		file_traits_open_close_proto_rawDescData = protoimpl.X.CompressGZIP(file_traits_open_close_proto_rawDescData)
	})
	return file_traits_open_close_proto_rawDescData
}

var file_traits_open_close_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_traits_open_close_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_traits_open_close_proto_goTypes = []interface{}{
	(OpenCloseDirection)(0),                       // 0: smartcore.traits.OpenCloseDirection
	(*OpenCloseAttributes)(nil),                   // 1: smartcore.traits.OpenCloseAttributes
	(*OpenClosePositions)(nil),                    // 2: smartcore.traits.OpenClosePositions
	(*OpenClosePosition)(nil),                     // 3: smartcore.traits.OpenClosePosition
	(*GetOpenClosePositionsRequest)(nil),          // 4: smartcore.traits.GetOpenClosePositionsRequest
	(*UpdateOpenClosePositionsRequest)(nil),       // 5: smartcore.traits.UpdateOpenClosePositionsRequest
	(*StopOpenCloseRequest)(nil),                  // 6: smartcore.traits.StopOpenCloseRequest
	(*PullOpenClosePositionsRequest)(nil),         // 7: smartcore.traits.PullOpenClosePositionsRequest
	(*PullOpenClosePositionsResponse)(nil),        // 8: smartcore.traits.PullOpenClosePositionsResponse
	(*PullOpenClosePositionsResponse_Change)(nil), // 9: smartcore.traits.PullOpenClosePositionsResponse.Change
	(*types.FloatAttributes)(nil),                 // 10: smartcore.api.types.FloatAttributes
	(*types.FloatVar)(nil),                        // 11: smartcore.api.types.FloatVar
	(*timestamp.Timestamp)(nil),                   // 12: google.protobuf.Timestamp
}
var file_traits_open_close_proto_depIdxs = []int32{
	10, // 0: smartcore.traits.OpenCloseAttributes.position_attributes:type_name -> smartcore.api.types.FloatAttributes
	0,  // 1: smartcore.traits.OpenCloseAttributes.directions:type_name -> smartcore.traits.OpenCloseDirection
	3,  // 2: smartcore.traits.OpenClosePositions.states:type_name -> smartcore.traits.OpenClosePosition
	11, // 3: smartcore.traits.OpenClosePosition.position_percent:type_name -> smartcore.api.types.FloatVar
	0,  // 4: smartcore.traits.OpenClosePosition.direction:type_name -> smartcore.traits.OpenCloseDirection
	2,  // 5: smartcore.traits.UpdateOpenClosePositionsRequest.states:type_name -> smartcore.traits.OpenClosePositions
	9,  // 6: smartcore.traits.PullOpenClosePositionsResponse.changes:type_name -> smartcore.traits.PullOpenClosePositionsResponse.Change
	12, // 7: smartcore.traits.PullOpenClosePositionsResponse.Change.change_time:type_name -> google.protobuf.Timestamp
	2,  // 8: smartcore.traits.PullOpenClosePositionsResponse.Change.state:type_name -> smartcore.traits.OpenClosePositions
	4,  // 9: smartcore.traits.OpenCloseApi.GetPositions:input_type -> smartcore.traits.GetOpenClosePositionsRequest
	5,  // 10: smartcore.traits.OpenCloseApi.UpdatePositions:input_type -> smartcore.traits.UpdateOpenClosePositionsRequest
	6,  // 11: smartcore.traits.OpenCloseApi.Stop:input_type -> smartcore.traits.StopOpenCloseRequest
	7,  // 12: smartcore.traits.OpenCloseApi.PullPositions:input_type -> smartcore.traits.PullOpenClosePositionsRequest
	2,  // 13: smartcore.traits.OpenCloseApi.GetPositions:output_type -> smartcore.traits.OpenClosePositions
	2,  // 14: smartcore.traits.OpenCloseApi.UpdatePositions:output_type -> smartcore.traits.OpenClosePositions
	2,  // 15: smartcore.traits.OpenCloseApi.Stop:output_type -> smartcore.traits.OpenClosePositions
	8,  // 16: smartcore.traits.OpenCloseApi.PullPositions:output_type -> smartcore.traits.PullOpenClosePositionsResponse
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_traits_open_close_proto_init() }
func file_traits_open_close_proto_init() {
	if File_traits_open_close_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_traits_open_close_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenCloseAttributes); i {
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
		file_traits_open_close_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenClosePositions); i {
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
		file_traits_open_close_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenClosePosition); i {
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
		file_traits_open_close_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOpenClosePositionsRequest); i {
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
		file_traits_open_close_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOpenClosePositionsRequest); i {
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
		file_traits_open_close_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopOpenCloseRequest); i {
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
		file_traits_open_close_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullOpenClosePositionsRequest); i {
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
		file_traits_open_close_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullOpenClosePositionsResponse); i {
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
		file_traits_open_close_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullOpenClosePositionsResponse_Change); i {
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
			RawDescriptor: file_traits_open_close_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_traits_open_close_proto_goTypes,
		DependencyIndexes: file_traits_open_close_proto_depIdxs,
		EnumInfos:         file_traits_open_close_proto_enumTypes,
		MessageInfos:      file_traits_open_close_proto_msgTypes,
	}.Build()
	File_traits_open_close_proto = out.File
	file_traits_open_close_proto_rawDesc = nil
	file_traits_open_close_proto_goTypes = nil
	file_traits_open_close_proto_depIdxs = nil
}
