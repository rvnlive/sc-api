// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: traits/air_quality_sensor.proto

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

// Comfort encodes levels of comfort for an area.
type AirQuality_Comfort int32

const (
	// The comfort level is unknown
	AirQuality_COMFORT_UNSPECIFIED AirQuality_Comfort = 0
	// The area should be comfortable for occupants
	AirQuality_COMFORTABLE AirQuality_Comfort = 1
	// The area might be uncomfortable for occupants
	AirQuality_UNCOMFORTABLE AirQuality_Comfort = 2
)

// Enum value maps for AirQuality_Comfort.
var (
	AirQuality_Comfort_name = map[int32]string{
		0: "COMFORT_UNSPECIFIED",
		1: "COMFORTABLE",
		2: "UNCOMFORTABLE",
	}
	AirQuality_Comfort_value = map[string]int32{
		"COMFORT_UNSPECIFIED": 0,
		"COMFORTABLE":         1,
		"UNCOMFORTABLE":       2,
	}
)

func (x AirQuality_Comfort) Enum() *AirQuality_Comfort {
	p := new(AirQuality_Comfort)
	*p = x
	return p
}

func (x AirQuality_Comfort) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AirQuality_Comfort) Descriptor() protoreflect.EnumDescriptor {
	return file_traits_air_quality_sensor_proto_enumTypes[0].Descriptor()
}

func (AirQuality_Comfort) Type() protoreflect.EnumType {
	return &file_traits_air_quality_sensor_proto_enumTypes[0]
}

func (x AirQuality_Comfort) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AirQuality_Comfort.Descriptor instead.
func (AirQuality_Comfort) EnumDescriptor() ([]byte, []int) {
	return file_traits_air_quality_sensor_proto_rawDescGZIP(), []int{0, 0}
}

// AirQuality holds the value of all supported sensors.
type AirQuality struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The CO2 level in parts per million
	CarbonDioxideLevel *float32 `protobuf:"fixed32,1,opt,name=carbon_dioxide_level,json=carbonDioxideLevel,proto3,oneof" json:"carbon_dioxide_level,omitempty"`
	// The VOC level in parts per million
	VolatileOrganicCompounds *float32 `protobuf:"fixed32,2,opt,name=volatile_organic_compounds,json=volatileOrganicCompounds,proto3,oneof" json:"volatile_organic_compounds,omitempty"`
	// The air pressure in hPa
	AirPressure *float32 `protobuf:"fixed32,3,opt,name=air_pressure,json=airPressure,proto3,oneof" json:"air_pressure,omitempty"`
	// General comfort of the area
	Comfort AirQuality_Comfort `protobuf:"varint,4,opt,name=comfort,proto3,enum=smartcore.traits.AirQuality_Comfort" json:"comfort,omitempty"`
	// A percentage [0,100] reading for how infectious the air is.
	// Typically a combination of other sensor readings, the combination algorithm is undefined.
	InfectionRisk *float32 `protobuf:"fixed32,5,opt,name=infection_risk,json=infectionRisk,proto3,oneof" json:"infection_risk,omitempty"`
	// An air quality score for the area mapped to a percentage [0,100], with 100 being the best quality.
	Score *float32 `protobuf:"fixed32,6,opt,name=score,proto3,oneof" json:"score,omitempty"`
}

func (x *AirQuality) Reset() {
	*x = AirQuality{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_air_quality_sensor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AirQuality) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AirQuality) ProtoMessage() {}

func (x *AirQuality) ProtoReflect() protoreflect.Message {
	mi := &file_traits_air_quality_sensor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AirQuality.ProtoReflect.Descriptor instead.
func (*AirQuality) Descriptor() ([]byte, []int) {
	return file_traits_air_quality_sensor_proto_rawDescGZIP(), []int{0}
}

func (x *AirQuality) GetCarbonDioxideLevel() float32 {
	if x != nil && x.CarbonDioxideLevel != nil {
		return *x.CarbonDioxideLevel
	}
	return 0
}

func (x *AirQuality) GetVolatileOrganicCompounds() float32 {
	if x != nil && x.VolatileOrganicCompounds != nil {
		return *x.VolatileOrganicCompounds
	}
	return 0
}

func (x *AirQuality) GetAirPressure() float32 {
	if x != nil && x.AirPressure != nil {
		return *x.AirPressure
	}
	return 0
}

func (x *AirQuality) GetComfort() AirQuality_Comfort {
	if x != nil {
		return x.Comfort
	}
	return AirQuality_COMFORT_UNSPECIFIED
}

func (x *AirQuality) GetInfectionRisk() float32 {
	if x != nil && x.InfectionRisk != nil {
		return *x.InfectionRisk
	}
	return 0
}

func (x *AirQuality) GetScore() float32 {
	if x != nil && x.Score != nil {
		return *x.Score
	}
	return 0
}

// AirQualitySupport describes the capabilities of devices implementing this trait
type AirQualitySupport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// How a named device supports read/pull apis
	ResourceSupport *types.ResourceSupport `protobuf:"bytes,1,opt,name=resource_support,json=resourceSupport,proto3" json:"resource_support,omitempty"`
	// If present then the sensor supports reporting CO2 levels, in parts-per-million.
	CarbonDioxideLevel *types.FloatBounds `protobuf:"bytes,2,opt,name=carbon_dioxide_level,json=carbonDioxideLevel,proto3" json:"carbon_dioxide_level,omitempty"`
	// If present then the sensor supports reporting volatile organic compounds, in parts-per-million.
	VolatileOrganicCompounds *types.FloatBounds `protobuf:"bytes,3,opt,name=volatile_organic_compounds,json=volatileOrganicCompounds,proto3" json:"volatile_organic_compounds,omitempty"`
	// If present then the sensor supports reporting air pressure, in hPa.
	AirPressure *types.FloatBounds `protobuf:"bytes,4,opt,name=air_pressure,json=airPressure,proto3" json:"air_pressure,omitempty"`
	// If non-empty then the sensor supports reporting a general comfort reading of any of the provided types.
	// Unknown should be assumed to be present if supported.
	Comfort []AirQuality_Comfort `protobuf:"varint,5,rep,packed,name=comfort,proto3,enum=smartcore.traits.AirQuality_Comfort" json:"comfort,omitempty"`
}

func (x *AirQualitySupport) Reset() {
	*x = AirQualitySupport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_air_quality_sensor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AirQualitySupport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AirQualitySupport) ProtoMessage() {}

func (x *AirQualitySupport) ProtoReflect() protoreflect.Message {
	mi := &file_traits_air_quality_sensor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AirQualitySupport.ProtoReflect.Descriptor instead.
func (*AirQualitySupport) Descriptor() ([]byte, []int) {
	return file_traits_air_quality_sensor_proto_rawDescGZIP(), []int{1}
}

func (x *AirQualitySupport) GetResourceSupport() *types.ResourceSupport {
	if x != nil {
		return x.ResourceSupport
	}
	return nil
}

func (x *AirQualitySupport) GetCarbonDioxideLevel() *types.FloatBounds {
	if x != nil {
		return x.CarbonDioxideLevel
	}
	return nil
}

func (x *AirQualitySupport) GetVolatileOrganicCompounds() *types.FloatBounds {
	if x != nil {
		return x.VolatileOrganicCompounds
	}
	return nil
}

func (x *AirQualitySupport) GetAirPressure() *types.FloatBounds {
	if x != nil {
		return x.AirPressure
	}
	return nil
}

func (x *AirQualitySupport) GetComfort() []AirQuality_Comfort {
	if x != nil {
		return x.Comfort
	}
	return nil
}

type GetAirQualityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the device to request state from
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The state fields to fetch
	ReadMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=read_mask,json=readMask,proto3" json:"read_mask,omitempty"`
}

func (x *GetAirQualityRequest) Reset() {
	*x = GetAirQualityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_air_quality_sensor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAirQualityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAirQualityRequest) ProtoMessage() {}

func (x *GetAirQualityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_air_quality_sensor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAirQualityRequest.ProtoReflect.Descriptor instead.
func (*GetAirQualityRequest) Descriptor() ([]byte, []int) {
	return file_traits_air_quality_sensor_proto_rawDescGZIP(), []int{2}
}

func (x *GetAirQualityRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetAirQualityRequest) GetReadMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.ReadMask
	}
	return nil
}

type PullAirQualityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the device to request state from
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The state fields to pull
	ReadMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=read_mask,json=readMask,proto3" json:"read_mask,omitempty"`
	// When true the device will only send changes to the resource value.
	// The default behaviour is to send the current value immediately followed by any updates as they happen.
	UpdatesOnly bool `protobuf:"varint,3,opt,name=updates_only,json=updatesOnly,proto3" json:"updates_only,omitempty"`
}

func (x *PullAirQualityRequest) Reset() {
	*x = PullAirQualityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_air_quality_sensor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullAirQualityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullAirQualityRequest) ProtoMessage() {}

func (x *PullAirQualityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_air_quality_sensor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullAirQualityRequest.ProtoReflect.Descriptor instead.
func (*PullAirQualityRequest) Descriptor() ([]byte, []int) {
	return file_traits_air_quality_sensor_proto_rawDescGZIP(), []int{3}
}

func (x *PullAirQualityRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PullAirQualityRequest) GetReadMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.ReadMask
	}
	return nil
}

func (x *PullAirQualityRequest) GetUpdatesOnly() bool {
	if x != nil {
		return x.UpdatesOnly
	}
	return false
}

type PullAirQualityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Changes []*PullAirQualityResponse_Change `protobuf:"bytes,1,rep,name=changes,proto3" json:"changes,omitempty"`
}

func (x *PullAirQualityResponse) Reset() {
	*x = PullAirQualityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_air_quality_sensor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullAirQualityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullAirQualityResponse) ProtoMessage() {}

func (x *PullAirQualityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_traits_air_quality_sensor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullAirQualityResponse.ProtoReflect.Descriptor instead.
func (*PullAirQualityResponse) Descriptor() ([]byte, []int) {
	return file_traits_air_quality_sensor_proto_rawDescGZIP(), []int{4}
}

func (x *PullAirQualityResponse) GetChanges() []*PullAirQualityResponse_Change {
	if x != nil {
		return x.Changes
	}
	return nil
}

type DescribeAirQualityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the device
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DescribeAirQualityRequest) Reset() {
	*x = DescribeAirQualityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_air_quality_sensor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeAirQualityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeAirQualityRequest) ProtoMessage() {}

func (x *DescribeAirQualityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_air_quality_sensor_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeAirQualityRequest.ProtoReflect.Descriptor instead.
func (*DescribeAirQualityRequest) Descriptor() ([]byte, []int) {
	return file_traits_air_quality_sensor_proto_rawDescGZIP(), []int{5}
}

func (x *DescribeAirQualityRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PullAirQualityResponse_Change struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the device that has changed.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The time the change happened.
	ChangeTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=change_time,json=changeTime,proto3" json:"change_time,omitempty"`
	// The new state of the sensor. Changed fields should be reported via the update_mask property.
	AirQuality *AirQuality `protobuf:"bytes,3,opt,name=air_quality,json=airQuality,proto3" json:"air_quality,omitempty"`
	// The state fields that have changed.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *PullAirQualityResponse_Change) Reset() {
	*x = PullAirQualityResponse_Change{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_air_quality_sensor_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullAirQualityResponse_Change) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullAirQualityResponse_Change) ProtoMessage() {}

func (x *PullAirQualityResponse_Change) ProtoReflect() protoreflect.Message {
	mi := &file_traits_air_quality_sensor_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullAirQualityResponse_Change.ProtoReflect.Descriptor instead.
func (*PullAirQualityResponse_Change) Descriptor() ([]byte, []int) {
	return file_traits_air_quality_sensor_proto_rawDescGZIP(), []int{4, 0}
}

func (x *PullAirQualityResponse_Change) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PullAirQualityResponse_Change) GetChangeTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ChangeTime
	}
	return nil
}

func (x *PullAirQualityResponse_Change) GetAirQuality() *AirQuality {
	if x != nil {
		return x.AirQuality
	}
	return nil
}

func (x *PullAirQualityResponse_Change) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

var File_traits_air_quality_sensor_proto protoreflect.FileDescriptor

var file_traits_air_quality_sensor_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2f, 0x61, 0x69, 0x72, 0x5f, 0x71, 0x75, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61,
	0x69, 0x74, 0x73, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x69, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x03, 0x0a,
	0x0a, 0x41, 0x69, 0x72, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x35, 0x0a, 0x14, 0x63,
	0x61, 0x72, 0x62, 0x6f, 0x6e, 0x5f, 0x64, 0x69, 0x6f, 0x78, 0x69, 0x64, 0x65, 0x5f, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x12, 0x63, 0x61, 0x72,
	0x62, 0x6f, 0x6e, 0x44, 0x69, 0x6f, 0x78, 0x69, 0x64, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x88,
	0x01, 0x01, 0x12, 0x41, 0x0a, 0x1a, 0x76, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6c, 0x65, 0x5f, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x75, 0x6e, 0x64, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x48, 0x01, 0x52, 0x18, 0x76, 0x6f, 0x6c, 0x61, 0x74, 0x69,
	0x6c, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x63, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x75, 0x6e,
	0x64, 0x73, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x61, 0x69, 0x72, 0x5f, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x48, 0x02, 0x52, 0x0b, 0x61,
	0x69, 0x72, 0x50, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x88, 0x01, 0x01, 0x12, 0x3e, 0x0a,
	0x07, 0x63, 0x6f, 0x6d, 0x66, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24,
	0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74,
	0x73, 0x2e, 0x41, 0x69, 0x72, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x6f, 0x6d,
	0x66, 0x6f, 0x72, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x66, 0x6f, 0x72, 0x74, 0x12, 0x2a, 0x0a,
	0x0e, 0x69, 0x6e, 0x66, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x69, 0x73, 0x6b, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x02, 0x48, 0x03, 0x52, 0x0d, 0x69, 0x6e, 0x66, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x69, 0x73, 0x6b, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x48, 0x04, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x88, 0x01, 0x01, 0x22, 0x46, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x66, 0x6f, 0x72, 0x74, 0x12,
	0x17, 0x0a, 0x13, 0x43, 0x4f, 0x4d, 0x46, 0x4f, 0x52, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x4f, 0x4d, 0x46,
	0x4f, 0x52, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x43,
	0x4f, 0x4d, 0x46, 0x4f, 0x52, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x42, 0x17, 0x0a, 0x15,
	0x5f, 0x63, 0x61, 0x72, 0x62, 0x6f, 0x6e, 0x5f, 0x64, 0x69, 0x6f, 0x78, 0x69, 0x64, 0x65, 0x5f,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x1d, 0x0a, 0x1b, 0x5f, 0x76, 0x6f, 0x6c, 0x61, 0x74, 0x69,
	0x6c, 0x65, 0x5f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x75, 0x6e, 0x64, 0x73, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x61, 0x69, 0x72, 0x5f, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x75, 0x72, 0x65, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x69, 0x6e, 0x66, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x69, 0x73, 0x6b, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x22, 0x8d, 0x03, 0x0a, 0x11, 0x41, 0x69, 0x72, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74,
	0x79, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x4b, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x4e, 0x0a, 0x14, 0x63, 0x61, 0x72, 0x62, 0x6f, 0x6e, 0x5f,
	0x64, 0x69, 0x6f, 0x78, 0x69, 0x64, 0x65, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x42, 0x6f, 0x75, 0x6e, 0x64,
	0x73, 0x52, 0x12, 0x63, 0x61, 0x72, 0x62, 0x6f, 0x6e, 0x44, 0x69, 0x6f, 0x78, 0x69, 0x64, 0x65,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x5a, 0x0a, 0x1a, 0x76, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6c,
	0x65, 0x5f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x75,
	0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x6d, 0x61, 0x72,
	0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x46, 0x6c, 0x6f, 0x61,
	0x74, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x52, 0x18, 0x76, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6c,
	0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x63, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x75, 0x6e, 0x64,
	0x73, 0x12, 0x3f, 0x0a, 0x0c, 0x61, 0x69, 0x72, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x42,
	0x6f, 0x75, 0x6e, 0x64, 0x73, 0x52, 0x0b, 0x61, 0x69, 0x72, 0x50, 0x72, 0x65, 0x73, 0x73, 0x75,
	0x72, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x66, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x41, 0x69, 0x72, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74,
	0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x66, 0x6f, 0x72, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x66, 0x6f,
	0x72, 0x74, 0x22, 0x63, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x69, 0x72, 0x51, 0x75, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37,
	0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x08, 0x72,
	0x65, 0x61, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x87, 0x01, 0x0a, 0x15, 0x50, 0x75, 0x6c, 0x6c,
	0x41, 0x69, 0x72, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6d, 0x61,
	0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4d, 0x61, 0x73, 0x6b, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x21,
	0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x4f, 0x6e, 0x6c,
	0x79, 0x22, 0xbb, 0x02, 0x0a, 0x16, 0x50, 0x75, 0x6c, 0x6c, 0x41, 0x69, 0x72, 0x51, 0x75, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x07,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e,
	0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73,
	0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x41, 0x69, 0x72, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x07,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x1a, 0xd5, 0x01, 0x0a, 0x06, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x61, 0x69, 0x72, 0x5f, 0x71, 0x75, 0x61, 0x6c, 0x69,
	0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x41, 0x69, 0x72, 0x51,
	0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x61, 0x69, 0x72, 0x51, 0x75, 0x61, 0x6c, 0x69,
	0x74, 0x79, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d,
	0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22,
	0x2f, 0x0a, 0x19, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x69, 0x72, 0x51, 0x75,
	0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x32, 0xd3, 0x01, 0x0a, 0x13, 0x41, 0x69, 0x72, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x53,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x41, 0x70, 0x69, 0x12, 0x55, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41,
	0x69, 0x72, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x26, 0x2e, 0x73, 0x6d, 0x61, 0x72,
	0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x69, 0x72, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72,
	0x61, 0x69, 0x74, 0x73, 0x2e, 0x41, 0x69, 0x72, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12,
	0x65, 0x0a, 0x0e, 0x50, 0x75, 0x6c, 0x6c, 0x41, 0x69, 0x72, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74,
	0x79, 0x12, 0x27, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72,
	0x61, 0x69, 0x74, 0x73, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x41, 0x69, 0x72, 0x51, 0x75, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x73, 0x6d, 0x61,
	0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x50, 0x75,
	0x6c, 0x6c, 0x41, 0x69, 0x72, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x32, 0x7e, 0x0a, 0x14, 0x41, 0x69, 0x72, 0x51, 0x75, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x66,
	0x0a, 0x12, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x69, 0x72, 0x51, 0x75, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x12, 0x2b, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x41, 0x69, 0x72, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72,
	0x61, 0x69, 0x74, 0x73, 0x2e, 0x41, 0x69, 0x72, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x53,
	0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x80, 0x01, 0x0a, 0x14, 0x64, 0x65, 0x76, 0x2e, 0x73,
	0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x42,
	0x15, 0x41, 0x69, 0x72, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2d,
	0x6f, 0x73, 0x2f, 0x73, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x72, 0x61,
	0x69, 0x74, 0x73, 0xaa, 0x02, 0x10, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x54, 0x72, 0x61, 0x69, 0x74, 0x73, 0xca, 0x02, 0x10, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f,
	0x72, 0x65, 0x5c, 0x54, 0x72, 0x61, 0x69, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_traits_air_quality_sensor_proto_rawDescOnce sync.Once
	file_traits_air_quality_sensor_proto_rawDescData = file_traits_air_quality_sensor_proto_rawDesc
)

func file_traits_air_quality_sensor_proto_rawDescGZIP() []byte {
	file_traits_air_quality_sensor_proto_rawDescOnce.Do(func() {
		file_traits_air_quality_sensor_proto_rawDescData = protoimpl.X.CompressGZIP(file_traits_air_quality_sensor_proto_rawDescData)
	})
	return file_traits_air_quality_sensor_proto_rawDescData
}

var file_traits_air_quality_sensor_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_traits_air_quality_sensor_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_traits_air_quality_sensor_proto_goTypes = []interface{}{
	(AirQuality_Comfort)(0),               // 0: smartcore.traits.AirQuality.Comfort
	(*AirQuality)(nil),                    // 1: smartcore.traits.AirQuality
	(*AirQualitySupport)(nil),             // 2: smartcore.traits.AirQualitySupport
	(*GetAirQualityRequest)(nil),          // 3: smartcore.traits.GetAirQualityRequest
	(*PullAirQualityRequest)(nil),         // 4: smartcore.traits.PullAirQualityRequest
	(*PullAirQualityResponse)(nil),        // 5: smartcore.traits.PullAirQualityResponse
	(*DescribeAirQualityRequest)(nil),     // 6: smartcore.traits.DescribeAirQualityRequest
	(*PullAirQualityResponse_Change)(nil), // 7: smartcore.traits.PullAirQualityResponse.Change
	(*types.ResourceSupport)(nil),         // 8: smartcore.types.ResourceSupport
	(*types.FloatBounds)(nil),             // 9: smartcore.types.FloatBounds
	(*fieldmaskpb.FieldMask)(nil),         // 10: google.protobuf.FieldMask
	(*timestamppb.Timestamp)(nil),         // 11: google.protobuf.Timestamp
}
var file_traits_air_quality_sensor_proto_depIdxs = []int32{
	0,  // 0: smartcore.traits.AirQuality.comfort:type_name -> smartcore.traits.AirQuality.Comfort
	8,  // 1: smartcore.traits.AirQualitySupport.resource_support:type_name -> smartcore.types.ResourceSupport
	9,  // 2: smartcore.traits.AirQualitySupport.carbon_dioxide_level:type_name -> smartcore.types.FloatBounds
	9,  // 3: smartcore.traits.AirQualitySupport.volatile_organic_compounds:type_name -> smartcore.types.FloatBounds
	9,  // 4: smartcore.traits.AirQualitySupport.air_pressure:type_name -> smartcore.types.FloatBounds
	0,  // 5: smartcore.traits.AirQualitySupport.comfort:type_name -> smartcore.traits.AirQuality.Comfort
	10, // 6: smartcore.traits.GetAirQualityRequest.read_mask:type_name -> google.protobuf.FieldMask
	10, // 7: smartcore.traits.PullAirQualityRequest.read_mask:type_name -> google.protobuf.FieldMask
	7,  // 8: smartcore.traits.PullAirQualityResponse.changes:type_name -> smartcore.traits.PullAirQualityResponse.Change
	11, // 9: smartcore.traits.PullAirQualityResponse.Change.change_time:type_name -> google.protobuf.Timestamp
	1,  // 10: smartcore.traits.PullAirQualityResponse.Change.air_quality:type_name -> smartcore.traits.AirQuality
	10, // 11: smartcore.traits.PullAirQualityResponse.Change.update_mask:type_name -> google.protobuf.FieldMask
	3,  // 12: smartcore.traits.AirQualitySensorApi.GetAirQuality:input_type -> smartcore.traits.GetAirQualityRequest
	4,  // 13: smartcore.traits.AirQualitySensorApi.PullAirQuality:input_type -> smartcore.traits.PullAirQualityRequest
	6,  // 14: smartcore.traits.AirQualitySensorInfo.DescribeAirQuality:input_type -> smartcore.traits.DescribeAirQualityRequest
	1,  // 15: smartcore.traits.AirQualitySensorApi.GetAirQuality:output_type -> smartcore.traits.AirQuality
	5,  // 16: smartcore.traits.AirQualitySensorApi.PullAirQuality:output_type -> smartcore.traits.PullAirQualityResponse
	2,  // 17: smartcore.traits.AirQualitySensorInfo.DescribeAirQuality:output_type -> smartcore.traits.AirQualitySupport
	15, // [15:18] is the sub-list for method output_type
	12, // [12:15] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_traits_air_quality_sensor_proto_init() }
func file_traits_air_quality_sensor_proto_init() {
	if File_traits_air_quality_sensor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_traits_air_quality_sensor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AirQuality); i {
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
		file_traits_air_quality_sensor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AirQualitySupport); i {
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
		file_traits_air_quality_sensor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAirQualityRequest); i {
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
		file_traits_air_quality_sensor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullAirQualityRequest); i {
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
		file_traits_air_quality_sensor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullAirQualityResponse); i {
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
		file_traits_air_quality_sensor_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeAirQualityRequest); i {
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
		file_traits_air_quality_sensor_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullAirQualityResponse_Change); i {
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
	file_traits_air_quality_sensor_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_traits_air_quality_sensor_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_traits_air_quality_sensor_proto_goTypes,
		DependencyIndexes: file_traits_air_quality_sensor_proto_depIdxs,
		EnumInfos:         file_traits_air_quality_sensor_proto_enumTypes,
		MessageInfos:      file_traits_air_quality_sensor_proto_msgTypes,
	}.Build()
	File_traits_air_quality_sensor_proto = out.File
	file_traits_air_quality_sensor_proto_rawDesc = nil
	file_traits_air_quality_sensor_proto_goTypes = nil
	file_traits_air_quality_sensor_proto_depIdxs = nil
}
