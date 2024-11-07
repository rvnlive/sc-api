// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.3
// source: types/number.proto

package types

import (
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

// Possible behaviours for values that are not allowed by the number
type InvalidNumberBehaviour int32

const (
	// A default behaviour will be applied, typically RESTRICT
	InvalidNumberBehaviour_INVALID_NUMBER_BEHAVIOUR_UNSPECIFIED InvalidNumberBehaviour = 0
	// The value will be restricted to the most appropriate value, typically the closest
	InvalidNumberBehaviour_RESTRICT InvalidNumberBehaviour = 1
	// An error will be raised to alert to the invalid value
	InvalidNumberBehaviour_ERROR InvalidNumberBehaviour = 2
	// Ignore the bounds and apply the value anyway
	InvalidNumberBehaviour_ALLOW InvalidNumberBehaviour = 3
)

// Enum value maps for InvalidNumberBehaviour.
var (
	InvalidNumberBehaviour_name = map[int32]string{
		0: "INVALID_NUMBER_BEHAVIOUR_UNSPECIFIED",
		1: "RESTRICT",
		2: "ERROR",
		3: "ALLOW",
	}
	InvalidNumberBehaviour_value = map[string]int32{
		"INVALID_NUMBER_BEHAVIOUR_UNSPECIFIED": 0,
		"RESTRICT":                             1,
		"ERROR":                                2,
		"ALLOW":                                3,
	}
)

func (x InvalidNumberBehaviour) Enum() *InvalidNumberBehaviour {
	p := new(InvalidNumberBehaviour)
	*p = x
	return p
}

func (x InvalidNumberBehaviour) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InvalidNumberBehaviour) Descriptor() protoreflect.EnumDescriptor {
	return file_types_number_proto_enumTypes[0].Descriptor()
}

func (InvalidNumberBehaviour) Type() protoreflect.EnumType {
	return &file_types_number_proto_enumTypes[0]
}

func (x InvalidNumberBehaviour) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InvalidNumberBehaviour.Descriptor instead.
func (InvalidNumberBehaviour) EnumDescriptor() ([]byte, []int) {
	return file_types_number_proto_rawDescGZIP(), []int{0}
}

// Options for how invalid number values will be handled
type NumberCapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If the value is less than the minimum, behave this way. Should default to RESTRICT
	Min InvalidNumberBehaviour `protobuf:"varint,1,opt,name=min,proto3,enum=smartcore.types.InvalidNumberBehaviour" json:"min,omitempty"`
	// If the value does not lie on a step value, behave this way. Should default to ALLOW
	Step InvalidNumberBehaviour `protobuf:"varint,2,opt,name=step,proto3,enum=smartcore.types.InvalidNumberBehaviour" json:"step,omitempty"`
	// If the value is greater than the maximum, behave this way. Should default to RESTRICT
	Max InvalidNumberBehaviour `protobuf:"varint,3,opt,name=max,proto3,enum=smartcore.types.InvalidNumberBehaviour" json:"max,omitempty"`
}

func (x *NumberCapping) Reset() {
	*x = NumberCapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_number_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumberCapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumberCapping) ProtoMessage() {}

func (x *NumberCapping) ProtoReflect() protoreflect.Message {
	mi := &file_types_number_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumberCapping.ProtoReflect.Descriptor instead.
func (*NumberCapping) Descriptor() ([]byte, []int) {
	return file_types_number_proto_rawDescGZIP(), []int{0}
}

func (x *NumberCapping) GetMin() InvalidNumberBehaviour {
	if x != nil {
		return x.Min
	}
	return InvalidNumberBehaviour_INVALID_NUMBER_BEHAVIOUR_UNSPECIFIED
}

func (x *NumberCapping) GetStep() InvalidNumberBehaviour {
	if x != nil {
		return x.Step
	}
	return InvalidNumberBehaviour_INVALID_NUMBER_BEHAVIOUR_UNSPECIFIED
}

func (x *NumberCapping) GetMax() InvalidNumberBehaviour {
	if x != nil {
		return x.Max
	}
	return InvalidNumberBehaviour_INVALID_NUMBER_BEHAVIOUR_UNSPECIFIED
}

// Defines the bounds for a number. Either of the min or max values can be absent which means those aspects are
// unbounded. It makes no sense for both to be unset, if so then this is no longer a bounds but that isn't
// disallowed.
//
// Contrary to most programming practices both min and max are inclusive and denote the max allowed value, just like
// htmls input type=range.
type Int32Bounds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the minimum value (inclusive). If absent then there is no minimum, though this is typically Int32.Min
	Min *int32 `protobuf:"varint,1,opt,name=min,proto3,oneof" json:"min,omitempty"`
	// the maximum value (inclusive). If absent then there is no maximum, though this is typically Int32.Max
	Max *int32 `protobuf:"varint,2,opt,name=max,proto3,oneof" json:"max,omitempty"`
}

func (x *Int32Bounds) Reset() {
	*x = Int32Bounds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_number_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Int32Bounds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int32Bounds) ProtoMessage() {}

func (x *Int32Bounds) ProtoReflect() protoreflect.Message {
	mi := &file_types_number_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int32Bounds.ProtoReflect.Descriptor instead.
func (*Int32Bounds) Descriptor() ([]byte, []int) {
	return file_types_number_proto_rawDescGZIP(), []int{1}
}

func (x *Int32Bounds) GetMin() int32 {
	if x != nil && x.Min != nil {
		return *x.Min
	}
	return 0
}

func (x *Int32Bounds) GetMax() int32 {
	if x != nil && x.Max != nil {
		return *x.Max
	}
	return 0
}

// Describes features and restrictions applied to a number that is typically used as part of an apis value
type Int32Attributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bounds *Int32Bounds `protobuf:"bytes,1,opt,name=bounds,proto3" json:"bounds,omitempty"`
	// The stepping interval, how little the value can change in one go. A default value of 0 is equivalent to a
	// continuous step, i.e. step=1
	Step int32 `protobuf:"varint,3,opt,name=step,proto3" json:"step,omitempty"`
	// Indicate whether the device supports delta adjustments when changing the value
	SupportsDelta bool `protobuf:"varint,4,opt,name=supports_delta,json=supportsDelta,proto3" json:"supports_delta,omitempty"`
	// Indicate the level of ramp support the device has
	RampSupport TweenSupport `protobuf:"varint,5,opt,name=ramp_support,json=rampSupport,proto3,enum=smartcore.types.TweenSupport" json:"ramp_support,omitempty"`
	// Provide information on how invalid values will be handled
	DefaultCapping *NumberCapping `protobuf:"bytes,6,opt,name=default_capping,json=defaultCapping,proto3" json:"default_capping,omitempty"`
}

func (x *Int32Attributes) Reset() {
	*x = Int32Attributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_number_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Int32Attributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int32Attributes) ProtoMessage() {}

func (x *Int32Attributes) ProtoReflect() protoreflect.Message {
	mi := &file_types_number_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int32Attributes.ProtoReflect.Descriptor instead.
func (*Int32Attributes) Descriptor() ([]byte, []int) {
	return file_types_number_proto_rawDescGZIP(), []int{2}
}

func (x *Int32Attributes) GetBounds() *Int32Bounds {
	if x != nil {
		return x.Bounds
	}
	return nil
}

func (x *Int32Attributes) GetStep() int32 {
	if x != nil {
		return x.Step
	}
	return 0
}

func (x *Int32Attributes) GetSupportsDelta() bool {
	if x != nil {
		return x.SupportsDelta
	}
	return false
}

func (x *Int32Attributes) GetRampSupport() TweenSupport {
	if x != nil {
		return x.RampSupport
	}
	return TweenSupport_TWEEN_SUPPORT_UNSPECIFIED
}

func (x *Int32Attributes) GetDefaultCapping() *NumberCapping {
	if x != nil {
		return x.DefaultCapping
	}
	return nil
}

// Defines the bounds for a number. Either of the min or max values can be absent which means those aspects are
// unbounded. It makes no sense for both to be unset, if so then this is no longer a bounds but that isn't
// disallowed.
//
// Contrary to most programming practices both min and max are inclusive and denote the max allowed value, just like
// htmls input type=range.
type FloatBounds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the minimum value (inclusive). If absent then there is no minimum, though this is typically Float.Min
	Min *float32 `protobuf:"fixed32,1,opt,name=min,proto3,oneof" json:"min,omitempty"`
	// the maximum value (inclusive). If absent then there is no maximum, though this is typically Float.Max
	Max *float32 `protobuf:"fixed32,2,opt,name=max,proto3,oneof" json:"max,omitempty"`
}

func (x *FloatBounds) Reset() {
	*x = FloatBounds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_number_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FloatBounds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FloatBounds) ProtoMessage() {}

func (x *FloatBounds) ProtoReflect() protoreflect.Message {
	mi := &file_types_number_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FloatBounds.ProtoReflect.Descriptor instead.
func (*FloatBounds) Descriptor() ([]byte, []int) {
	return file_types_number_proto_rawDescGZIP(), []int{3}
}

func (x *FloatBounds) GetMin() float32 {
	if x != nil && x.Min != nil {
		return *x.Min
	}
	return 0
}

func (x *FloatBounds) GetMax() float32 {
	if x != nil && x.Max != nil {
		return *x.Max
	}
	return 0
}

// Describes features and restrictions applied to a number that is typically used as part of an apis value
type FloatAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bounds *FloatBounds `protobuf:"bytes,1,opt,name=bounds,proto3" json:"bounds,omitempty"`
	// The stepping interval, how little the value can change in one go. A default value of 0 is equivalent to a
	// continuous step, i.e. all possible float values are supported
	Step float32 `protobuf:"fixed32,3,opt,name=step,proto3" json:"step,omitempty"`
	// Indicate whether the device supports delta adjustments when changing the value
	SupportsDelta bool `protobuf:"varint,4,opt,name=supports_delta,json=supportsDelta,proto3" json:"supports_delta,omitempty"`
	// Indicate the level of ramp support the device has
	RampSupport TweenSupport `protobuf:"varint,5,opt,name=ramp_support,json=rampSupport,proto3,enum=smartcore.types.TweenSupport" json:"ramp_support,omitempty"`
	// Provide information on how invalid values will be handled
	DefaultCapping *NumberCapping `protobuf:"bytes,6,opt,name=default_capping,json=defaultCapping,proto3" json:"default_capping,omitempty"`
}

func (x *FloatAttributes) Reset() {
	*x = FloatAttributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_number_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FloatAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FloatAttributes) ProtoMessage() {}

func (x *FloatAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_types_number_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FloatAttributes.ProtoReflect.Descriptor instead.
func (*FloatAttributes) Descriptor() ([]byte, []int) {
	return file_types_number_proto_rawDescGZIP(), []int{4}
}

func (x *FloatAttributes) GetBounds() *FloatBounds {
	if x != nil {
		return x.Bounds
	}
	return nil
}

func (x *FloatAttributes) GetStep() float32 {
	if x != nil {
		return x.Step
	}
	return 0
}

func (x *FloatAttributes) GetSupportsDelta() bool {
	if x != nil {
		return x.SupportsDelta
	}
	return false
}

func (x *FloatAttributes) GetRampSupport() TweenSupport {
	if x != nil {
		return x.RampSupport
	}
	return TweenSupport_TWEEN_SUPPORT_UNSPECIFIED
}

func (x *FloatAttributes) GetDefaultCapping() *NumberCapping {
	if x != nil {
		return x.DefaultCapping
	}
	return nil
}

var File_types_number_proto protoreflect.FileDescriptor

var file_types_number_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x11, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x77, 0x65,
	0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x01, 0x0a, 0x0d, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x43, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x39, 0x0a, 0x03, 0x6d, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x75, 0x72,
	0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x3b, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x75, 0x72, 0x52, 0x04, 0x73, 0x74,
	0x65, 0x70, 0x12, 0x39, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x27, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x75, 0x72, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x22, 0x4b, 0x0a,
	0x0b, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x12, 0x15, 0x0a, 0x03,
	0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x03, 0x6d, 0x69, 0x6e,
	0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x01, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6d,
	0x69, 0x6e, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6d, 0x61, 0x78, 0x22, 0x8d, 0x02, 0x0a, 0x0f, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x34,
	0x0a, 0x06, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x52, 0x06, 0x62, 0x6f,
	0x75, 0x6e, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x75, 0x70, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0d, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x12,
	0x40, 0x0a, 0x0c, 0x72, 0x61, 0x6d, 0x70, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x77, 0x65, 0x65, 0x6e, 0x53, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x0b, 0x72, 0x61, 0x6d, 0x70, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x47, 0x0a, 0x0f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x6d, 0x61,
	0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x43, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x0e, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x43, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x22, 0x4b, 0x0a, 0x0b, 0x46, 0x6c,
	0x6f, 0x61, 0x74, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x12, 0x15, 0x0a, 0x03, 0x6d, 0x69, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x88, 0x01, 0x01,
	0x12, 0x15, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x48, 0x01, 0x52,
	0x03, 0x6d, 0x61, 0x78, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6d, 0x69, 0x6e, 0x42,
	0x06, 0x0a, 0x04, 0x5f, 0x6d, 0x61, 0x78, 0x22, 0x8d, 0x02, 0x0a, 0x0f, 0x46, 0x6c, 0x6f, 0x61,
	0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x06, 0x62,
	0x6f, 0x75, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x6d,
	0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x46, 0x6c,
	0x6f, 0x61, 0x74, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x52, 0x06, 0x62, 0x6f, 0x75, 0x6e, 0x64,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x04, 0x73, 0x74, 0x65, 0x70, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x5f, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x73,
	0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x12, 0x40, 0x0a, 0x0c,
	0x72, 0x61, 0x6d, 0x70, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x77, 0x65, 0x65, 0x6e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x0b, 0x72, 0x61, 0x6d, 0x70, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x47,
	0x0a, 0x0f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x61, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x43, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x0e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x43, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2a, 0x66, 0x0a, 0x16, 0x49, 0x6e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x75,
	0x72, 0x12, 0x28, 0x0a, 0x24, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x4e, 0x55, 0x4d,
	0x42, 0x45, 0x52, 0x5f, 0x42, 0x45, 0x48, 0x41, 0x56, 0x49, 0x4f, 0x55, 0x52, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x52,
	0x45, 0x53, 0x54, 0x52, 0x49, 0x43, 0x54, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x10, 0x03, 0x42,
	0x60, 0x0a, 0x13, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x42, 0x0b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2d, 0x6f, 0x73, 0x2f,
	0x73, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0xaa,
	0x02, 0x0f, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_number_proto_rawDescOnce sync.Once
	file_types_number_proto_rawDescData = file_types_number_proto_rawDesc
)

func file_types_number_proto_rawDescGZIP() []byte {
	file_types_number_proto_rawDescOnce.Do(func() {
		file_types_number_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_number_proto_rawDescData)
	})
	return file_types_number_proto_rawDescData
}

var file_types_number_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_types_number_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_types_number_proto_goTypes = []any{
	(InvalidNumberBehaviour)(0), // 0: smartcore.types.InvalidNumberBehaviour
	(*NumberCapping)(nil),       // 1: smartcore.types.NumberCapping
	(*Int32Bounds)(nil),         // 2: smartcore.types.Int32Bounds
	(*Int32Attributes)(nil),     // 3: smartcore.types.Int32Attributes
	(*FloatBounds)(nil),         // 4: smartcore.types.FloatBounds
	(*FloatAttributes)(nil),     // 5: smartcore.types.FloatAttributes
	(TweenSupport)(0),           // 6: smartcore.types.TweenSupport
}
var file_types_number_proto_depIdxs = []int32{
	0, // 0: smartcore.types.NumberCapping.min:type_name -> smartcore.types.InvalidNumberBehaviour
	0, // 1: smartcore.types.NumberCapping.step:type_name -> smartcore.types.InvalidNumberBehaviour
	0, // 2: smartcore.types.NumberCapping.max:type_name -> smartcore.types.InvalidNumberBehaviour
	2, // 3: smartcore.types.Int32Attributes.bounds:type_name -> smartcore.types.Int32Bounds
	6, // 4: smartcore.types.Int32Attributes.ramp_support:type_name -> smartcore.types.TweenSupport
	1, // 5: smartcore.types.Int32Attributes.default_capping:type_name -> smartcore.types.NumberCapping
	4, // 6: smartcore.types.FloatAttributes.bounds:type_name -> smartcore.types.FloatBounds
	6, // 7: smartcore.types.FloatAttributes.ramp_support:type_name -> smartcore.types.TweenSupport
	1, // 8: smartcore.types.FloatAttributes.default_capping:type_name -> smartcore.types.NumberCapping
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_types_number_proto_init() }
func file_types_number_proto_init() {
	if File_types_number_proto != nil {
		return
	}
	file_types_tween_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_types_number_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*NumberCapping); i {
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
		file_types_number_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Int32Bounds); i {
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
		file_types_number_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Int32Attributes); i {
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
		file_types_number_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*FloatBounds); i {
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
		file_types_number_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*FloatAttributes); i {
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
	file_types_number_proto_msgTypes[1].OneofWrappers = []any{}
	file_types_number_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_types_number_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_number_proto_goTypes,
		DependencyIndexes: file_types_number_proto_depIdxs,
		EnumInfos:         file_types_number_proto_enumTypes,
		MessageInfos:      file_types_number_proto_msgTypes,
	}.Build()
	File_types_number_proto = out.File
	file_types_number_proto_rawDesc = nil
	file_types_number_proto_goTypes = nil
	file_types_number_proto_depIdxs = nil
}
