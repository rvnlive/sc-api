// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: traits/temperature_sensor.proto

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

// Trait.attribute for temperature sensors
type TemperatureSensorAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// How the device thinks, all communications will be in celsius but the device's native unit is this.
	NativeUnit types.TemperatureUnit `protobuf:"varint,1,opt,name=native_unit,json=nativeUnit,proto3,enum=smartcore.api.types.TemperatureUnit" json:"native_unit,omitempty"`
}

func (x *TemperatureSensorAttributes) Reset() {
	*x = TemperatureSensorAttributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_temperature_sensor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemperatureSensorAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemperatureSensorAttributes) ProtoMessage() {}

func (x *TemperatureSensorAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_traits_temperature_sensor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemperatureSensorAttributes.ProtoReflect.Descriptor instead.
func (*TemperatureSensorAttributes) Descriptor() ([]byte, []int) {
	return file_traits_temperature_sensor_proto_rawDescGZIP(), []int{0}
}

func (x *TemperatureSensorAttributes) GetNativeUnit() types.TemperatureUnit {
	if x != nil {
		return x.NativeUnit
	}
	return types.TemperatureUnit_CELSIUS
}

type GetTemperatureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the name of the device to get the state for
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetTemperatureRequest) Reset() {
	*x = GetTemperatureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_temperature_sensor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTemperatureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTemperatureRequest) ProtoMessage() {}

func (x *GetTemperatureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_temperature_sensor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTemperatureRequest.ProtoReflect.Descriptor instead.
func (*GetTemperatureRequest) Descriptor() ([]byte, []int) {
	return file_traits_temperature_sensor_proto_rawDescGZIP(), []int{1}
}

func (x *GetTemperatureRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PullTemperatureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the name of the device to get the state for
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PullTemperatureRequest) Reset() {
	*x = PullTemperatureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_temperature_sensor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullTemperatureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullTemperatureRequest) ProtoMessage() {}

func (x *PullTemperatureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_temperature_sensor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullTemperatureRequest.ProtoReflect.Descriptor instead.
func (*PullTemperatureRequest) Descriptor() ([]byte, []int) {
	return file_traits_temperature_sensor_proto_rawDescGZIP(), []int{2}
}

func (x *PullTemperatureRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// response containing temperature changes
type PullTemperatureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the list of changes that have occurred since the last event
	Changes []*TemperatureChange `protobuf:"bytes,1,rep,name=changes,proto3" json:"changes,omitempty"`
}

func (x *PullTemperatureResponse) Reset() {
	*x = PullTemperatureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_temperature_sensor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullTemperatureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullTemperatureResponse) ProtoMessage() {}

func (x *PullTemperatureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_traits_temperature_sensor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullTemperatureResponse.ProtoReflect.Descriptor instead.
func (*PullTemperatureResponse) Descriptor() ([]byte, []int) {
	return file_traits_temperature_sensor_proto_rawDescGZIP(), []int{3}
}

func (x *PullTemperatureResponse) GetChanges() []*TemperatureChange {
	if x != nil {
		return x.Changes
	}
	return nil
}

// A change to the temperature reading of the device
type TemperatureChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name for the device that issued the change
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// when the change occurred
	CreateTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// the new state for the device
	Temperature *types.Temperature `protobuf:"bytes,3,opt,name=temperature,proto3" json:"temperature,omitempty"`
}

func (x *TemperatureChange) Reset() {
	*x = TemperatureChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_temperature_sensor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemperatureChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemperatureChange) ProtoMessage() {}

func (x *TemperatureChange) ProtoReflect() protoreflect.Message {
	mi := &file_traits_temperature_sensor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemperatureChange.ProtoReflect.Descriptor instead.
func (*TemperatureChange) Descriptor() ([]byte, []int) {
	return file_traits_temperature_sensor_proto_rawDescGZIP(), []int{4}
}

func (x *TemperatureChange) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TemperatureChange) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *TemperatureChange) GetTemperature() *types.Temperature {
	if x != nil {
		return x.Temperature
	}
	return nil
}

var File_traits_temperature_sensor_proto protoreflect.FileDescriptor

var file_traits_temperature_sensor_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61,
	0x69, 0x74, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x75, 0x6e, 0x69, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a, 0x1b, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x45, 0x0a, 0x0b, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f,
	0x75, 0x6e, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x73, 0x6d, 0x61,
	0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x55, 0x6e, 0x69, 0x74,
	0x52, 0x0a, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x22, 0x2b, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2c, 0x0a, 0x16, 0x50, 0x75, 0x6c,
	0x6c, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x58, 0x0a, 0x17, 0x50, 0x75, 0x6c, 0x6c, 0x54,
	0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3d, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x73, 0x22, 0xa8, 0x01, 0x0a, 0x11, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x42, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52,
	0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x32, 0xdc, 0x01, 0x0a,
	0x11, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x12, 0x5b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x27, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x6a, 0x0a, 0x0f, 0x50, 0x75, 0x6c, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x28, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74,
	0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x73,
	0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e,
	0x50, 0x75, 0x6c, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x82, 0x01, 0x0a, 0x14,
	0x64, 0x65, 0x76, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72,
	0x61, 0x69, 0x74, 0x73, 0x42, 0x16, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2a,
	0x67, 0x69, 0x74, 0x2e, 0x76, 0x61, 0x6e, 0x74, 0x69, 0x2e, 0x63, 0x6f, 0x2e, 0x75, 0x6b, 0x2f,
	0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x63, 0x2d, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0xaa, 0x02, 0x10, 0x53, 0x6d, 0x61,
	0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x74, 0x73, 0xca, 0x02, 0x10,
	0x53, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x5c, 0x54, 0x72, 0x61, 0x69, 0x74, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_traits_temperature_sensor_proto_rawDescOnce sync.Once
	file_traits_temperature_sensor_proto_rawDescData = file_traits_temperature_sensor_proto_rawDesc
)

func file_traits_temperature_sensor_proto_rawDescGZIP() []byte {
	file_traits_temperature_sensor_proto_rawDescOnce.Do(func() {
		file_traits_temperature_sensor_proto_rawDescData = protoimpl.X.CompressGZIP(file_traits_temperature_sensor_proto_rawDescData)
	})
	return file_traits_temperature_sensor_proto_rawDescData
}

var file_traits_temperature_sensor_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_traits_temperature_sensor_proto_goTypes = []interface{}{
	(*TemperatureSensorAttributes)(nil), // 0: smartcore.traits.TemperatureSensorAttributes
	(*GetTemperatureRequest)(nil),       // 1: smartcore.traits.GetTemperatureRequest
	(*PullTemperatureRequest)(nil),      // 2: smartcore.traits.PullTemperatureRequest
	(*PullTemperatureResponse)(nil),     // 3: smartcore.traits.PullTemperatureResponse
	(*TemperatureChange)(nil),           // 4: smartcore.traits.TemperatureChange
	(types.TemperatureUnit)(0),          // 5: smartcore.api.types.TemperatureUnit
	(*timestamp.Timestamp)(nil),         // 6: google.protobuf.Timestamp
	(*types.Temperature)(nil),           // 7: smartcore.api.types.Temperature
}
var file_traits_temperature_sensor_proto_depIdxs = []int32{
	5, // 0: smartcore.traits.TemperatureSensorAttributes.native_unit:type_name -> smartcore.api.types.TemperatureUnit
	4, // 1: smartcore.traits.PullTemperatureResponse.changes:type_name -> smartcore.traits.TemperatureChange
	6, // 2: smartcore.traits.TemperatureChange.create_time:type_name -> google.protobuf.Timestamp
	7, // 3: smartcore.traits.TemperatureChange.temperature:type_name -> smartcore.api.types.Temperature
	1, // 4: smartcore.traits.TemperatureSensor.GetTemperature:input_type -> smartcore.traits.GetTemperatureRequest
	2, // 5: smartcore.traits.TemperatureSensor.PullTemperature:input_type -> smartcore.traits.PullTemperatureRequest
	7, // 6: smartcore.traits.TemperatureSensor.GetTemperature:output_type -> smartcore.api.types.Temperature
	3, // 7: smartcore.traits.TemperatureSensor.PullTemperature:output_type -> smartcore.traits.PullTemperatureResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_traits_temperature_sensor_proto_init() }
func file_traits_temperature_sensor_proto_init() {
	if File_traits_temperature_sensor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_traits_temperature_sensor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemperatureSensorAttributes); i {
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
		file_traits_temperature_sensor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTemperatureRequest); i {
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
		file_traits_temperature_sensor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullTemperatureRequest); i {
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
		file_traits_temperature_sensor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullTemperatureResponse); i {
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
		file_traits_temperature_sensor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemperatureChange); i {
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
			RawDescriptor: file_traits_temperature_sensor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_traits_temperature_sensor_proto_goTypes,
		DependencyIndexes: file_traits_temperature_sensor_proto_depIdxs,
		MessageInfos:      file_traits_temperature_sensor_proto_msgTypes,
	}.Build()
	File_traits_temperature_sensor_proto = out.File
	file_traits_temperature_sensor_proto_rawDesc = nil
	file_traits_temperature_sensor_proto_goTypes = nil
	file_traits_temperature_sensor_proto_depIdxs = nil
}