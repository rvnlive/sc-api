// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: traits/occupancy_sensor.proto

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

// Possible states for occupancy
type OccupancyState int32

const (
	// There are no signals to suggest either an occupied or unoccupied space
	OccupancyState_OCCUPANCY_STATE_NO_SIGNALS OccupancyState = 0
	// The space is occupied
	OccupancyState_OCCUPANCY_STATE_OCCUPIED OccupancyState = 1
	// The space is unoccupied
	OccupancyState_OCCUPANCY_STATE_UNOCCUPIED OccupancyState = 2
	// The space is likely occupied but some signals suggest that no activity is occurring (i.e. people are asleep or not
	// moving for other reasons)
	OccupancyState_OCCUPANCY_STATE_IDLE OccupancyState = 3
)

// Enum value maps for OccupancyState.
var (
	OccupancyState_name = map[int32]string{
		0: "OCCUPANCY_STATE_NO_SIGNALS",
		1: "OCCUPANCY_STATE_OCCUPIED",
		2: "OCCUPANCY_STATE_UNOCCUPIED",
		3: "OCCUPANCY_STATE_IDLE",
	}
	OccupancyState_value = map[string]int32{
		"OCCUPANCY_STATE_NO_SIGNALS": 0,
		"OCCUPANCY_STATE_OCCUPIED":   1,
		"OCCUPANCY_STATE_UNOCCUPIED": 2,
		"OCCUPANCY_STATE_IDLE":       3,
	}
)

func (x OccupancyState) Enum() *OccupancyState {
	p := new(OccupancyState)
	*p = x
	return p
}

func (x OccupancyState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OccupancyState) Descriptor() protoreflect.EnumDescriptor {
	return file_traits_occupancy_sensor_proto_enumTypes[0].Descriptor()
}

func (OccupancyState) Type() protoreflect.EnumType {
	return &file_traits_occupancy_sensor_proto_enumTypes[0]
}

func (x OccupancyState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OccupancyState.Descriptor instead.
func (OccupancyState) EnumDescriptor() ([]byte, []int) {
	return file_traits_occupancy_sensor_proto_rawDescGZIP(), []int{0}
}

// The occupancy state the device is reporting and updating
type Occupancy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The current occupancy state
	State OccupancyState `protobuf:"varint,1,opt,name=state,proto3,enum=smartcore.traits.OccupancyState" json:"state,omitempty"`
	// The total number of people the occupancy sensor has detected. Do NOT use this value as an indication of occupancy
	// state, a value of 0 could be reported for a space that is occupied if the device either doesn't support people
	// counts or there is some other undefined issue with the counting part of the sensor suite.
	PeopleCount int32 `protobuf:"varint,2,opt,name=people_count,json=peopleCount,proto3" json:"people_count,omitempty"`
	// When the occupancy state last changed. Does not update when people_count changes unlike the timestamp in
	// OccupancyChange events
	StateChangeTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=state_change_time,json=stateChangeTime,proto3" json:"state_change_time,omitempty"`
	// Optional. List of human readable strings explaining why the device thinks the space is in the current state. For
	// example could say OCCUPIED:["Detected people in space"] or IDLE:["No motion detected for 10 minutes"]. Typically
	// used for debugging or troubleshooting purposes.
	Reasons []string `protobuf:"bytes,4,rep,name=reasons,proto3" json:"reasons,omitempty"`
	// Optional. How confident is the sensor that the current occupancy state is accurate. A value of 0 means that the
	// confidence is unknown
	Confidence float64 `protobuf:"fixed64,5,opt,name=confidence,proto3" json:"confidence,omitempty"`
}

func (x *Occupancy) Reset() {
	*x = Occupancy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_occupancy_sensor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Occupancy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Occupancy) ProtoMessage() {}

func (x *Occupancy) ProtoReflect() protoreflect.Message {
	mi := &file_traits_occupancy_sensor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Occupancy.ProtoReflect.Descriptor instead.
func (*Occupancy) Descriptor() ([]byte, []int) {
	return file_traits_occupancy_sensor_proto_rawDescGZIP(), []int{0}
}

func (x *Occupancy) GetState() OccupancyState {
	if x != nil {
		return x.State
	}
	return OccupancyState_OCCUPANCY_STATE_NO_SIGNALS
}

func (x *Occupancy) GetPeopleCount() int32 {
	if x != nil {
		return x.PeopleCount
	}
	return 0
}

func (x *Occupancy) GetStateChangeTime() *timestamp.Timestamp {
	if x != nil {
		return x.StateChangeTime
	}
	return nil
}

func (x *Occupancy) GetReasons() []string {
	if x != nil {
		return x.Reasons
	}
	return nil
}

func (x *Occupancy) GetConfidence() float64 {
	if x != nil {
		return x.Confidence
	}
	return 0
}

// OccupancySupport describes the capabilities of devices implementing this trait
type OccupancySupport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// How a named device supports read/pull apis
	SensorSupport *types.SensorSupport `protobuf:"bytes,1,opt,name=sensor_support,json=sensorSupport,proto3" json:"sensor_support,omitempty"`
	// how many people can the occupancy service report. 0 means it won't report people counts. Note: this is _not_ the
	// total capacity for the space, it's the capacity of the sensor
	MaxPeople int32 `protobuf:"varint,2,opt,name=max_people,json=maxPeople,proto3" json:"max_people,omitempty"`
}

func (x *OccupancySupport) Reset() {
	*x = OccupancySupport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_occupancy_sensor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OccupancySupport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OccupancySupport) ProtoMessage() {}

func (x *OccupancySupport) ProtoReflect() protoreflect.Message {
	mi := &file_traits_occupancy_sensor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OccupancySupport.ProtoReflect.Descriptor instead.
func (*OccupancySupport) Descriptor() ([]byte, []int) {
	return file_traits_occupancy_sensor_proto_rawDescGZIP(), []int{1}
}

func (x *OccupancySupport) GetSensorSupport() *types.SensorSupport {
	if x != nil {
		return x.SensorSupport
	}
	return nil
}

func (x *OccupancySupport) GetMaxPeople() int32 {
	if x != nil {
		return x.MaxPeople
	}
	return 0
}

// request to fetch the current state of the device
type GetOccupancyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the name for the device to get the occupancy state for
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetOccupancyRequest) Reset() {
	*x = GetOccupancyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_occupancy_sensor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOccupancyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOccupancyRequest) ProtoMessage() {}

func (x *GetOccupancyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_occupancy_sensor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOccupancyRequest.ProtoReflect.Descriptor instead.
func (*GetOccupancyRequest) Descriptor() ([]byte, []int) {
	return file_traits_occupancy_sensor_proto_rawDescGZIP(), []int{2}
}

func (x *GetOccupancyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// request to be notified of changes to the state of the device
type PullOccupancyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the name for the device to subscribe to
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PullOccupancyRequest) Reset() {
	*x = PullOccupancyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_occupancy_sensor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullOccupancyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullOccupancyRequest) ProtoMessage() {}

func (x *PullOccupancyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_occupancy_sensor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullOccupancyRequest.ProtoReflect.Descriptor instead.
func (*PullOccupancyRequest) Descriptor() ([]byte, []int) {
	return file_traits_occupancy_sensor_proto_rawDescGZIP(), []int{3}
}

func (x *PullOccupancyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// response containing occupancy state changes
type PullOccupancyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the list of changes that have occurred since the last event
	Changes []*PullOccupancyResponse_Change `protobuf:"bytes,1,rep,name=changes,proto3" json:"changes,omitempty"`
}

func (x *PullOccupancyResponse) Reset() {
	*x = PullOccupancyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_occupancy_sensor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullOccupancyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullOccupancyResponse) ProtoMessage() {}

func (x *PullOccupancyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_traits_occupancy_sensor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullOccupancyResponse.ProtoReflect.Descriptor instead.
func (*PullOccupancyResponse) Descriptor() ([]byte, []int) {
	return file_traits_occupancy_sensor_proto_rawDescGZIP(), []int{4}
}

func (x *PullOccupancyResponse) GetChanges() []*PullOccupancyResponse_Change {
	if x != nil {
		return x.Changes
	}
	return nil
}

type DescribeOccupancyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the device
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DescribeOccupancyRequest) Reset() {
	*x = DescribeOccupancyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_occupancy_sensor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeOccupancyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeOccupancyRequest) ProtoMessage() {}

func (x *DescribeOccupancyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_occupancy_sensor_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeOccupancyRequest.ProtoReflect.Descriptor instead.
func (*DescribeOccupancyRequest) Descriptor() ([]byte, []int) {
	return file_traits_occupancy_sensor_proto_rawDescGZIP(), []int{5}
}

func (x *DescribeOccupancyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// A change to the occupancy state of the device, including people count
type PullOccupancyResponse_Change struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name for the device that issued the change
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// when the change occurred
	ChangeTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=change_time,json=changeTime,proto3" json:"change_time,omitempty"`
	// the new state for the device
	Occupancy *Occupancy `protobuf:"bytes,3,opt,name=occupancy,proto3" json:"occupancy,omitempty"`
}

func (x *PullOccupancyResponse_Change) Reset() {
	*x = PullOccupancyResponse_Change{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_occupancy_sensor_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullOccupancyResponse_Change) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullOccupancyResponse_Change) ProtoMessage() {}

func (x *PullOccupancyResponse_Change) ProtoReflect() protoreflect.Message {
	mi := &file_traits_occupancy_sensor_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullOccupancyResponse_Change.ProtoReflect.Descriptor instead.
func (*PullOccupancyResponse_Change) Descriptor() ([]byte, []int) {
	return file_traits_occupancy_sensor_proto_rawDescGZIP(), []int{4, 0}
}

func (x *PullOccupancyResponse_Change) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PullOccupancyResponse_Change) GetChangeTime() *timestamp.Timestamp {
	if x != nil {
		return x.ChangeTime
	}
	return nil
}

func (x *PullOccupancyResponse_Change) GetOccupancy() *Occupancy {
	if x != nil {
		return x.Occupancy
	}
	return nil
}

var File_traits_occupancy_sensor_proto protoreflect.FileDescriptor

var file_traits_occupancy_sensor_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2f, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e,
	0x63, 0x79, 0x5f, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74,
	0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x10, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x01, 0x0a, 0x09, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e,
	0x63, 0x79, 0x12, 0x36, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x20, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72,
	0x61, 0x69, 0x74, 0x73, 0x2e, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e, 0x63, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x65,
	0x6f, 0x70, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x46, 0x0a,
	0x11, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x22,
	0x78, 0x0a, 0x10, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e, 0x63, 0x79, 0x53, 0x75, 0x70, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x45, 0x0a, 0x0e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x73, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x6d,
	0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x0d, 0x73, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61,
	0x78, 0x5f, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x6d, 0x61, 0x78, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x22, 0x29, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x14, 0x50, 0x75, 0x6c, 0x6c, 0x4f, 0x63, 0x63, 0x75,
	0x70, 0x61, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0xf8, 0x01, 0x0a, 0x15, 0x50, 0x75, 0x6c, 0x6c, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e,
	0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x07, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x73, 0x6d,
	0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x50,
	0x75, 0x6c, 0x6c, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x07, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x73, 0x1a, 0x94, 0x01, 0x0a, 0x06, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x39, 0x0a, 0x09, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e, 0x63, 0x79,
	0x52, 0x09, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e, 0x63, 0x79, 0x22, 0x2e, 0x0a, 0x18, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e, 0x63, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x2a, 0x88, 0x01, 0x0a, 0x0e,
	0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e, 0x63, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1e,
	0x0a, 0x1a, 0x4f, 0x43, 0x43, 0x55, 0x50, 0x41, 0x4e, 0x43, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x4e, 0x4f, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x41, 0x4c, 0x53, 0x10, 0x00, 0x12, 0x1c,
	0x0a, 0x18, 0x4f, 0x43, 0x43, 0x55, 0x50, 0x41, 0x4e, 0x43, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x4f, 0x43, 0x43, 0x55, 0x50, 0x49, 0x45, 0x44, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a,
	0x4f, 0x43, 0x43, 0x55, 0x50, 0x41, 0x4e, 0x43, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x55, 0x4e, 0x4f, 0x43, 0x43, 0x55, 0x50, 0x49, 0x45, 0x44, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14,
	0x4f, 0x43, 0x43, 0x55, 0x50, 0x41, 0x4e, 0x43, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x49, 0x44, 0x4c, 0x45, 0x10, 0x03, 0x32, 0xcc, 0x01, 0x0a, 0x12, 0x4f, 0x63, 0x63, 0x75, 0x70,
	0x61, 0x6e, 0x63, 0x79, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x41, 0x70, 0x69, 0x12, 0x52, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e, 0x63, 0x79, 0x12, 0x25, 0x2e,
	0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e, 0x63,
	0x79, 0x12, 0x62, 0x0a, 0x0d, 0x50, 0x75, 0x6c, 0x6c, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e,
	0x63, 0x79, 0x12, 0x26, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74,
	0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61,
	0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x6d, 0x61,
	0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x50, 0x75,
	0x6c, 0x6c, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x30, 0x01, 0x32, 0x7a, 0x0a, 0x13, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e,
	0x63, 0x79, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x63, 0x0a, 0x11,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e, 0x63,
	0x79, 0x12, 0x2a, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72,
	0x61, 0x69, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4f, 0x63, 0x63,
	0x75, 0x70, 0x61, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73,
	0x2e, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e, 0x63, 0x79, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x42, 0x80, 0x01, 0x0a, 0x14, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x42, 0x14, 0x4f, 0x63, 0x63, 0x75,
	0x70, 0x61, 0x6e, 0x63, 0x79, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x2e, 0x76, 0x61, 0x6e, 0x74, 0x69, 0x2e, 0x63, 0x6f,
	0x2e, 0x75, 0x6b, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x63,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0xaa, 0x02,
	0x10, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x74,
	0x73, 0xca, 0x02, 0x10, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x5c, 0x54, 0x72,
	0x61, 0x69, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_traits_occupancy_sensor_proto_rawDescOnce sync.Once
	file_traits_occupancy_sensor_proto_rawDescData = file_traits_occupancy_sensor_proto_rawDesc
)

func file_traits_occupancy_sensor_proto_rawDescGZIP() []byte {
	file_traits_occupancy_sensor_proto_rawDescOnce.Do(func() {
		file_traits_occupancy_sensor_proto_rawDescData = protoimpl.X.CompressGZIP(file_traits_occupancy_sensor_proto_rawDescData)
	})
	return file_traits_occupancy_sensor_proto_rawDescData
}

var file_traits_occupancy_sensor_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_traits_occupancy_sensor_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_traits_occupancy_sensor_proto_goTypes = []interface{}{
	(OccupancyState)(0),                  // 0: smartcore.traits.OccupancyState
	(*Occupancy)(nil),                    // 1: smartcore.traits.Occupancy
	(*OccupancySupport)(nil),             // 2: smartcore.traits.OccupancySupport
	(*GetOccupancyRequest)(nil),          // 3: smartcore.traits.GetOccupancyRequest
	(*PullOccupancyRequest)(nil),         // 4: smartcore.traits.PullOccupancyRequest
	(*PullOccupancyResponse)(nil),        // 5: smartcore.traits.PullOccupancyResponse
	(*DescribeOccupancyRequest)(nil),     // 6: smartcore.traits.DescribeOccupancyRequest
	(*PullOccupancyResponse_Change)(nil), // 7: smartcore.traits.PullOccupancyResponse.Change
	(*timestamp.Timestamp)(nil),          // 8: google.protobuf.Timestamp
	(*types.SensorSupport)(nil),          // 9: smartcore.types.SensorSupport
}
var file_traits_occupancy_sensor_proto_depIdxs = []int32{
	0, // 0: smartcore.traits.Occupancy.state:type_name -> smartcore.traits.OccupancyState
	8, // 1: smartcore.traits.Occupancy.state_change_time:type_name -> google.protobuf.Timestamp
	9, // 2: smartcore.traits.OccupancySupport.sensor_support:type_name -> smartcore.types.SensorSupport
	7, // 3: smartcore.traits.PullOccupancyResponse.changes:type_name -> smartcore.traits.PullOccupancyResponse.Change
	8, // 4: smartcore.traits.PullOccupancyResponse.Change.change_time:type_name -> google.protobuf.Timestamp
	1, // 5: smartcore.traits.PullOccupancyResponse.Change.occupancy:type_name -> smartcore.traits.Occupancy
	3, // 6: smartcore.traits.OccupancySensorApi.GetOccupancy:input_type -> smartcore.traits.GetOccupancyRequest
	4, // 7: smartcore.traits.OccupancySensorApi.PullOccupancy:input_type -> smartcore.traits.PullOccupancyRequest
	6, // 8: smartcore.traits.OccupancySensorInfo.DescribeOccupancy:input_type -> smartcore.traits.DescribeOccupancyRequest
	1, // 9: smartcore.traits.OccupancySensorApi.GetOccupancy:output_type -> smartcore.traits.Occupancy
	5, // 10: smartcore.traits.OccupancySensorApi.PullOccupancy:output_type -> smartcore.traits.PullOccupancyResponse
	2, // 11: smartcore.traits.OccupancySensorInfo.DescribeOccupancy:output_type -> smartcore.traits.OccupancySupport
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_traits_occupancy_sensor_proto_init() }
func file_traits_occupancy_sensor_proto_init() {
	if File_traits_occupancy_sensor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_traits_occupancy_sensor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Occupancy); i {
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
		file_traits_occupancy_sensor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OccupancySupport); i {
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
		file_traits_occupancy_sensor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOccupancyRequest); i {
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
		file_traits_occupancy_sensor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullOccupancyRequest); i {
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
		file_traits_occupancy_sensor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullOccupancyResponse); i {
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
		file_traits_occupancy_sensor_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeOccupancyRequest); i {
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
		file_traits_occupancy_sensor_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullOccupancyResponse_Change); i {
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
			RawDescriptor: file_traits_occupancy_sensor_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_traits_occupancy_sensor_proto_goTypes,
		DependencyIndexes: file_traits_occupancy_sensor_proto_depIdxs,
		EnumInfos:         file_traits_occupancy_sensor_proto_enumTypes,
		MessageInfos:      file_traits_occupancy_sensor_proto_msgTypes,
	}.Build()
	File_traits_occupancy_sensor_proto = out.File
	file_traits_occupancy_sensor_proto_rawDesc = nil
	file_traits_occupancy_sensor_proto_goTypes = nil
	file_traits_occupancy_sensor_proto_depIdxs = nil
}
