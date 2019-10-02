// Code generated by protoc-gen-go. DO NOT EDIT.
// source: device/traits/thermostat.proto

package traits

import (
	context "context"
	fmt "fmt"
	types "git.vanti.co.uk/smartcore/sc-api/go/types"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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

// Supported modes for a thermostat. Some of these values are used as descriptive attributes, some are used for control
type ThermostatMode int32

const (
	// The thermostat mode is unknown during a query. If used during a write then no change will be made, if part of a
	// read then the mode is unknown. This makes no sense as part of an attribute.
	ThermostatMode_UNKNOWN ThermostatMode = 0
	// Write-only. If the thermostat is OFF restore it to it's previous state
	ThermostatMode_ON ThermostatMode = 1
	// Attr, read, write. The thermostat supports, is, or should be disabled
	ThermostatMode_OFF ThermostatMode = 2
	// Attr, read, write. The device supports, is, or should be heating
	ThermostatMode_HEAT ThermostatMode = 3
	// Attr, read, write. The device supports, is, or should be cooling
	ThermostatMode_COOL ThermostatMode = 4
	// Attr, read, write. The device supports, is, or should be maintaining heating/cooling to target a specific set
	// point (and/or min + max)
	ThermostatMode_HEAT_COOL ThermostatMode = 5
	// Attr. The device supports automatic control of set points and/or schedules based on some other means (AI for
	// example)
	ThermostatMode_AUTO ThermostatMode = 6
	// Attr, read, write. The device supports, is, or should be able to use only the fan without heating/cooling elements.
	ThermostatMode_FAN_ONLY ThermostatMode = 7
	// Attr, read, write. The device supports, is, or should be in an energy saving "eco" mode
	ThermostatMode_ECO ThermostatMode = 8
	// Attr, read, write. The device supports, is, or should be in an air purifying mode
	ThermostatMode_PURIFIER ThermostatMode = 9
	// Attr, read, write. The device supports, is, or should be in an air drying mode
	ThermostatMode_DRY ThermostatMode = 10
	// Attr, read, write. The device supports, is, or should be in locked mode (i.e. not user-editable)
	ThermostatMode_LOCKED ThermostatMode = 11
)

var ThermostatMode_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "ON",
	2:  "OFF",
	3:  "HEAT",
	4:  "COOL",
	5:  "HEAT_COOL",
	6:  "AUTO",
	7:  "FAN_ONLY",
	8:  "ECO",
	9:  "PURIFIER",
	10: "DRY",
	11: "LOCKED",
}

var ThermostatMode_value = map[string]int32{
	"UNKNOWN":   0,
	"ON":        1,
	"OFF":       2,
	"HEAT":      3,
	"COOL":      4,
	"HEAT_COOL": 5,
	"AUTO":      6,
	"FAN_ONLY":  7,
	"ECO":       8,
	"PURIFIER":  9,
	"DRY":       10,
	"LOCKED":    11,
}

func (x ThermostatMode) String() string {
	return proto.EnumName(ThermostatMode_name, int32(x))
}

func (ThermostatMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_028e612c270b2841, []int{0}
}

// Trait.attributes message for describing this device trait.
type ThermostatAttributes struct {
	// How the device thinks, all communications will be in celsius but the device's native unit is this.
	NativeUnit     types.TemperatureUnit `protobuf:"varint,1,opt,name=native_unit,json=nativeUnit,proto3,enum=smartcore.api.types.TemperatureUnit" json:"native_unit,omitempty"`
	SupportedModes []ThermostatMode      `protobuf:"varint,2,rep,packed,name=supported_modes,json=supportedModes,proto3,enum=smartcore.api.device.traits.ThermostatMode" json:"supported_modes,omitempty"`
	// The minimum difference between the low and high temperatures when set using a range. 0 means unset, default to 2.
	MinRangeCelsius      float64  `protobuf:"fixed64,3,opt,name=min_range_celsius,json=minRangeCelsius,proto3" json:"min_range_celsius,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThermostatAttributes) Reset()         { *m = ThermostatAttributes{} }
func (m *ThermostatAttributes) String() string { return proto.CompactTextString(m) }
func (*ThermostatAttributes) ProtoMessage()    {}
func (*ThermostatAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_028e612c270b2841, []int{0}
}

func (m *ThermostatAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThermostatAttributes.Unmarshal(m, b)
}
func (m *ThermostatAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThermostatAttributes.Marshal(b, m, deterministic)
}
func (m *ThermostatAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThermostatAttributes.Merge(m, src)
}
func (m *ThermostatAttributes) XXX_Size() int {
	return xxx_messageInfo_ThermostatAttributes.Size(m)
}
func (m *ThermostatAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_ThermostatAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_ThermostatAttributes proto.InternalMessageInfo

func (m *ThermostatAttributes) GetNativeUnit() types.TemperatureUnit {
	if m != nil {
		return m.NativeUnit
	}
	return types.TemperatureUnit_CELSIUS
}

func (m *ThermostatAttributes) GetSupportedModes() []ThermostatMode {
	if m != nil {
		return m.SupportedModes
	}
	return nil
}

func (m *ThermostatAttributes) GetMinRangeCelsius() float64 {
	if m != nil {
		return m.MinRangeCelsius
	}
	return 0
}

// Request message for fetching thermostat state
type GetThermostatStateRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetThermostatStateRequest) Reset()         { *m = GetThermostatStateRequest{} }
func (m *GetThermostatStateRequest) String() string { return proto.CompactTextString(m) }
func (*GetThermostatStateRequest) ProtoMessage()    {}
func (*GetThermostatStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_028e612c270b2841, []int{1}
}

func (m *GetThermostatStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetThermostatStateRequest.Unmarshal(m, b)
}
func (m *GetThermostatStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetThermostatStateRequest.Marshal(b, m, deterministic)
}
func (m *GetThermostatStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetThermostatStateRequest.Merge(m, src)
}
func (m *GetThermostatStateRequest) XXX_Size() int {
	return xxx_messageInfo_GetThermostatStateRequest.Size(m)
}
func (m *GetThermostatStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetThermostatStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetThermostatStateRequest proto.InternalMessageInfo

func (m *GetThermostatStateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Request message for updating the thermostat state
type UpdateThermostatStateRequest struct {
	Name  string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	State *ThermostatState `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	// The fields relative to state we intend to update
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateThermostatStateRequest) Reset()         { *m = UpdateThermostatStateRequest{} }
func (m *UpdateThermostatStateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateThermostatStateRequest) ProtoMessage()    {}
func (*UpdateThermostatStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_028e612c270b2841, []int{2}
}

func (m *UpdateThermostatStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateThermostatStateRequest.Unmarshal(m, b)
}
func (m *UpdateThermostatStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateThermostatStateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateThermostatStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateThermostatStateRequest.Merge(m, src)
}
func (m *UpdateThermostatStateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateThermostatStateRequest.Size(m)
}
func (m *UpdateThermostatStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateThermostatStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateThermostatStateRequest proto.InternalMessageInfo

func (m *UpdateThermostatStateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateThermostatStateRequest) GetState() *ThermostatState {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *UpdateThermostatStateRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

// Request message for subscribing to changes in the thermostats state
type PullThermostatStateRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PullThermostatStateRequest) Reset()         { *m = PullThermostatStateRequest{} }
func (m *PullThermostatStateRequest) String() string { return proto.CompactTextString(m) }
func (*PullThermostatStateRequest) ProtoMessage()    {}
func (*PullThermostatStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_028e612c270b2841, []int{3}
}

func (m *PullThermostatStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PullThermostatStateRequest.Unmarshal(m, b)
}
func (m *PullThermostatStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PullThermostatStateRequest.Marshal(b, m, deterministic)
}
func (m *PullThermostatStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullThermostatStateRequest.Merge(m, src)
}
func (m *PullThermostatStateRequest) XXX_Size() int {
	return xxx_messageInfo_PullThermostatStateRequest.Size(m)
}
func (m *PullThermostatStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PullThermostatStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PullThermostatStateRequest proto.InternalMessageInfo

func (m *PullThermostatStateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Response message enclosing thermostat state changes
type PullThermostatStateResponse struct {
	Changes              []*ThermostatStateChange `protobuf:"bytes,1,rep,name=changes,proto3" json:"changes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *PullThermostatStateResponse) Reset()         { *m = PullThermostatStateResponse{} }
func (m *PullThermostatStateResponse) String() string { return proto.CompactTextString(m) }
func (*PullThermostatStateResponse) ProtoMessage()    {}
func (*PullThermostatStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_028e612c270b2841, []int{4}
}

func (m *PullThermostatStateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PullThermostatStateResponse.Unmarshal(m, b)
}
func (m *PullThermostatStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PullThermostatStateResponse.Marshal(b, m, deterministic)
}
func (m *PullThermostatStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullThermostatStateResponse.Merge(m, src)
}
func (m *PullThermostatStateResponse) XXX_Size() int {
	return xxx_messageInfo_PullThermostatStateResponse.Size(m)
}
func (m *PullThermostatStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PullThermostatStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PullThermostatStateResponse proto.InternalMessageInfo

func (m *PullThermostatStateResponse) GetChanges() []*ThermostatStateChange {
	if m != nil {
		return m.Changes
	}
	return nil
}

type ThermostatStateChange struct {
	// name for the device that issued the change
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// when the change occurred
	CreateTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The new value for the thermostat state. Only changed fields will be set, should be merged with GetState full
	// response as required.
	State                *ThermostatState `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ThermostatStateChange) Reset()         { *m = ThermostatStateChange{} }
func (m *ThermostatStateChange) String() string { return proto.CompactTextString(m) }
func (*ThermostatStateChange) ProtoMessage()    {}
func (*ThermostatStateChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_028e612c270b2841, []int{5}
}

func (m *ThermostatStateChange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThermostatStateChange.Unmarshal(m, b)
}
func (m *ThermostatStateChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThermostatStateChange.Marshal(b, m, deterministic)
}
func (m *ThermostatStateChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThermostatStateChange.Merge(m, src)
}
func (m *ThermostatStateChange) XXX_Size() int {
	return xxx_messageInfo_ThermostatStateChange.Size(m)
}
func (m *ThermostatStateChange) XXX_DiscardUnknown() {
	xxx_messageInfo_ThermostatStateChange.DiscardUnknown(m)
}

var xxx_messageInfo_ThermostatStateChange proto.InternalMessageInfo

func (m *ThermostatStateChange) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ThermostatStateChange) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *ThermostatStateChange) GetState() *ThermostatState {
	if m != nil {
		return m.State
	}
	return nil
}

// All the properties of the thermostat
type ThermostatState struct {
	// The current mode for the thermostat
	Mode ThermostatMode `protobuf:"varint,1,opt,name=mode,proto3,enum=smartcore.api.device.traits.ThermostatMode" json:"mode,omitempty"`
	// Types that are valid to be assigned to TemperatureGoal:
	//	*ThermostatState_TemperatureSetPoint
	//	*ThermostatState_TemperatureSetPointDelta
	//	*ThermostatState_TemperatureRange
	TemperatureGoal isThermostatState_TemperatureGoal `protobuf_oneof:"temperature_goal"`
	// Optional, read-only. The ambient temperature as read by the thermostat
	AmbientTemperature *types.Temperature `protobuf:"bytes,5,opt,name=ambient_temperature,json=ambientTemperature,proto3" json:"ambient_temperature,omitempty"`
	// Optional, read-only. The ambient humidity as read by the thermostat
	AmbientHumidity      *wrappers.FloatValue `protobuf:"bytes,6,opt,name=ambient_humidity,json=ambientHumidity,proto3" json:"ambient_humidity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ThermostatState) Reset()         { *m = ThermostatState{} }
func (m *ThermostatState) String() string { return proto.CompactTextString(m) }
func (*ThermostatState) ProtoMessage()    {}
func (*ThermostatState) Descriptor() ([]byte, []int) {
	return fileDescriptor_028e612c270b2841, []int{6}
}

func (m *ThermostatState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThermostatState.Unmarshal(m, b)
}
func (m *ThermostatState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThermostatState.Marshal(b, m, deterministic)
}
func (m *ThermostatState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThermostatState.Merge(m, src)
}
func (m *ThermostatState) XXX_Size() int {
	return xxx_messageInfo_ThermostatState.Size(m)
}
func (m *ThermostatState) XXX_DiscardUnknown() {
	xxx_messageInfo_ThermostatState.DiscardUnknown(m)
}

var xxx_messageInfo_ThermostatState proto.InternalMessageInfo

func (m *ThermostatState) GetMode() ThermostatMode {
	if m != nil {
		return m.Mode
	}
	return ThermostatMode_UNKNOWN
}

type isThermostatState_TemperatureGoal interface {
	isThermostatState_TemperatureGoal()
}

type ThermostatState_TemperatureSetPoint struct {
	TemperatureSetPoint *types.Temperature `protobuf:"bytes,2,opt,name=temperature_set_point,json=temperatureSetPoint,proto3,oneof"`
}

type ThermostatState_TemperatureSetPointDelta struct {
	TemperatureSetPointDelta *types.Temperature `protobuf:"bytes,3,opt,name=temperature_set_point_delta,json=temperatureSetPointDelta,proto3,oneof"`
}

type ThermostatState_TemperatureRange struct {
	TemperatureRange *TemperatureRange `protobuf:"bytes,4,opt,name=temperature_range,json=temperatureRange,proto3,oneof"`
}

func (*ThermostatState_TemperatureSetPoint) isThermostatState_TemperatureGoal() {}

func (*ThermostatState_TemperatureSetPointDelta) isThermostatState_TemperatureGoal() {}

func (*ThermostatState_TemperatureRange) isThermostatState_TemperatureGoal() {}

func (m *ThermostatState) GetTemperatureGoal() isThermostatState_TemperatureGoal {
	if m != nil {
		return m.TemperatureGoal
	}
	return nil
}

func (m *ThermostatState) GetTemperatureSetPoint() *types.Temperature {
	if x, ok := m.GetTemperatureGoal().(*ThermostatState_TemperatureSetPoint); ok {
		return x.TemperatureSetPoint
	}
	return nil
}

func (m *ThermostatState) GetTemperatureSetPointDelta() *types.Temperature {
	if x, ok := m.GetTemperatureGoal().(*ThermostatState_TemperatureSetPointDelta); ok {
		return x.TemperatureSetPointDelta
	}
	return nil
}

func (m *ThermostatState) GetTemperatureRange() *TemperatureRange {
	if x, ok := m.GetTemperatureGoal().(*ThermostatState_TemperatureRange); ok {
		return x.TemperatureRange
	}
	return nil
}

func (m *ThermostatState) GetAmbientTemperature() *types.Temperature {
	if m != nil {
		return m.AmbientTemperature
	}
	return nil
}

func (m *ThermostatState) GetAmbientHumidity() *wrappers.FloatValue {
	if m != nil {
		return m.AmbientHumidity
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ThermostatState) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ThermostatState_TemperatureSetPoint)(nil),
		(*ThermostatState_TemperatureSetPointDelta)(nil),
		(*ThermostatState_TemperatureRange)(nil),
	}
}

// A setting for thermostats that target a temperature between a range.
type TemperatureRange struct {
	// Required. The low threshold for the range
	Low *types.Temperature `protobuf:"bytes,1,opt,name=low,proto3" json:"low,omitempty"`
	// Required. The high threshold for the range
	High *types.Temperature `protobuf:"bytes,2,opt,name=high,proto3" json:"high,omitempty"`
	// Optional. An ideal value for cases where a thermostat supports three set points
	Ideal                *types.Temperature `protobuf:"bytes,3,opt,name=ideal,proto3" json:"ideal,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TemperatureRange) Reset()         { *m = TemperatureRange{} }
func (m *TemperatureRange) String() string { return proto.CompactTextString(m) }
func (*TemperatureRange) ProtoMessage()    {}
func (*TemperatureRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_028e612c270b2841, []int{7}
}

func (m *TemperatureRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TemperatureRange.Unmarshal(m, b)
}
func (m *TemperatureRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TemperatureRange.Marshal(b, m, deterministic)
}
func (m *TemperatureRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemperatureRange.Merge(m, src)
}
func (m *TemperatureRange) XXX_Size() int {
	return xxx_messageInfo_TemperatureRange.Size(m)
}
func (m *TemperatureRange) XXX_DiscardUnknown() {
	xxx_messageInfo_TemperatureRange.DiscardUnknown(m)
}

var xxx_messageInfo_TemperatureRange proto.InternalMessageInfo

func (m *TemperatureRange) GetLow() *types.Temperature {
	if m != nil {
		return m.Low
	}
	return nil
}

func (m *TemperatureRange) GetHigh() *types.Temperature {
	if m != nil {
		return m.High
	}
	return nil
}

func (m *TemperatureRange) GetIdeal() *types.Temperature {
	if m != nil {
		return m.Ideal
	}
	return nil
}

func init() {
	proto.RegisterEnum("smartcore.api.device.traits.ThermostatMode", ThermostatMode_name, ThermostatMode_value)
	proto.RegisterType((*ThermostatAttributes)(nil), "smartcore.api.device.traits.ThermostatAttributes")
	proto.RegisterType((*GetThermostatStateRequest)(nil), "smartcore.api.device.traits.GetThermostatStateRequest")
	proto.RegisterType((*UpdateThermostatStateRequest)(nil), "smartcore.api.device.traits.UpdateThermostatStateRequest")
	proto.RegisterType((*PullThermostatStateRequest)(nil), "smartcore.api.device.traits.PullThermostatStateRequest")
	proto.RegisterType((*PullThermostatStateResponse)(nil), "smartcore.api.device.traits.PullThermostatStateResponse")
	proto.RegisterType((*ThermostatStateChange)(nil), "smartcore.api.device.traits.ThermostatStateChange")
	proto.RegisterType((*ThermostatState)(nil), "smartcore.api.device.traits.ThermostatState")
	proto.RegisterType((*TemperatureRange)(nil), "smartcore.api.device.traits.TemperatureRange")
}

func init() { proto.RegisterFile("device/traits/thermostat.proto", fileDescriptor_028e612c270b2841) }

var fileDescriptor_028e612c270b2841 = []byte{
	// 909 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0x4d, 0x6f, 0x1a, 0x47,
	0x18, 0xf6, 0x02, 0xfe, 0x7a, 0x69, 0xed, 0xcd, 0xa4, 0x91, 0x28, 0x54, 0x09, 0x42, 0x3d, 0xa0,
	0xb4, 0xd9, 0x4d, 0x69, 0x95, 0xb6, 0xca, 0xa1, 0xc2, 0x18, 0xe2, 0x28, 0x0e, 0xd0, 0x35, 0xa4,
	0x4a, 0x65, 0x69, 0x35, 0x66, 0xdf, 0xc0, 0xc8, 0xfb, 0xd5, 0x9d, 0x59, 0xa2, 0xdc, 0xfa, 0x27,
	0x7a, 0xe9, 0xb5, 0xb7, 0xaa, 0xb7, 0xf6, 0x57, 0xf4, 0x4f, 0xf4, 0x77, 0xf4, 0x56, 0xcd, 0xcc,
	0xda, 0x60, 0x42, 0x28, 0xe4, 0xb6, 0xf3, 0xbe, 0xcf, 0xf3, 0xcc, 0xfb, 0x39, 0x00, 0x77, 0x3d,
	0x9c, 0xb2, 0x11, 0xda, 0x22, 0xa1, 0x4c, 0x70, 0x5b, 0x4c, 0x30, 0x09, 0x22, 0x2e, 0xa8, 0xb0,
	0xe2, 0x24, 0x12, 0x11, 0xa9, 0xf0, 0x80, 0x26, 0x62, 0x14, 0x25, 0x68, 0xd1, 0x98, 0x59, 0x1a,
	0x6d, 0x69, 0x74, 0xb9, 0x3a, 0x8e, 0xa2, 0xb1, 0x8f, 0xb6, 0x82, 0x5e, 0xa4, 0xaf, 0xec, 0x57,
	0x0c, 0x7d, 0xcf, 0x0d, 0x28, 0xbf, 0xd4, 0xf4, 0xf2, 0xbd, 0x45, 0x84, 0x60, 0x01, 0x72, 0x41,
	0x83, 0x38, 0x03, 0xdc, 0x5d, 0x04, 0xbc, 0x4e, 0x68, 0x1c, 0x63, 0xc2, 0x33, 0xbf, 0x29, 0xde,
	0xc4, 0xc8, 0xed, 0x34, 0x64, 0x59, 0x44, 0xb5, 0x7f, 0x0c, 0xf8, 0x68, 0x70, 0x1d, 0x66, 0x53,
	0x88, 0x84, 0x5d, 0xa4, 0x02, 0x39, 0x69, 0x43, 0x31, 0xa4, 0x82, 0x4d, 0xd1, 0x95, 0xe8, 0x92,
	0x51, 0x35, 0xea, 0x07, 0x8d, 0x4f, 0xad, 0x9b, 0x09, 0x28, 0x39, 0x6b, 0x80, 0x41, 0x8c, 0x09,
	0x15, 0x69, 0x82, 0xc3, 0x90, 0x09, 0x07, 0x34, 0x51, 0x7e, 0x93, 0x01, 0x1c, 0xf2, 0x34, 0x8e,
	0xa3, 0x44, 0xa0, 0xe7, 0x06, 0x91, 0x87, 0xbc, 0x94, 0xab, 0xe6, 0xeb, 0x07, 0x8d, 0xcf, 0xac,
	0x15, 0xb5, 0xb0, 0x66, 0x21, 0x3d, 0x8f, 0x3c, 0x74, 0x0e, 0xae, 0x35, 0xe4, 0x91, 0x93, 0xfb,
	0x70, 0x2b, 0x60, 0xa1, 0x9b, 0xd0, 0x70, 0x8c, 0xee, 0x08, 0x7d, 0xce, 0x52, 0x5e, 0xca, 0x57,
	0x8d, 0xba, 0xe1, 0x1c, 0x06, 0x2c, 0x74, 0xa4, 0xbd, 0xa5, 0xcd, 0x35, 0x1b, 0x3e, 0x7e, 0x82,
	0x62, 0x26, 0x78, 0x26, 0xa8, 0x40, 0x07, 0x7f, 0x4a, 0x91, 0x0b, 0x42, 0xa0, 0x10, 0xd2, 0x00,
	0x55, 0x7a, 0xfb, 0x8e, 0xfa, 0xae, 0xfd, 0x69, 0xc0, 0x27, 0xc3, 0xd8, 0xa3, 0x02, 0xd7, 0x27,
	0x91, 0x23, 0xd8, 0x96, 0x38, 0x2c, 0xe5, 0xaa, 0x46, 0xbd, 0xd8, 0xf8, 0x7c, 0xcd, 0xec, 0xb4,
	0xae, 0xa6, 0x92, 0xc7, 0x50, 0x4c, 0xd5, 0xbd, 0xaa, 0xe7, 0x2a, 0x9f, 0x62, 0xa3, 0x6c, 0xe9,
	0x9e, 0x5a, 0x57, 0x3d, 0xb5, 0x3a, 0x72, 0x2c, 0x9e, 0x53, 0x7e, 0xe9, 0x80, 0x86, 0xcb, 0xef,
	0xda, 0x43, 0x28, 0xf7, 0x53, 0xdf, 0xdf, 0x20, 0xcf, 0x4b, 0xa8, 0x2c, 0x65, 0xf0, 0x38, 0x0a,
	0x39, 0x92, 0x53, 0xd8, 0x1d, 0x4d, 0x64, 0x21, 0x79, 0xc9, 0xa8, 0xe6, 0xeb, 0xc5, 0x46, 0x63,
	0x93, 0x9c, 0x5a, 0x8a, 0xea, 0x5c, 0x49, 0xd4, 0xfe, 0x30, 0xe0, 0xce, 0x52, 0xc8, 0xd2, 0x6a,
	0x3e, 0x86, 0xe2, 0x28, 0x41, 0x59, 0x09, 0x39, 0xe1, 0x59, 0x4d, 0xdf, 0xae, 0xc4, 0xe0, 0x6a,
	0xfc, 0x1d, 0xd0, 0x70, 0x69, 0x98, 0xb5, 0x22, 0xff, 0xde, 0xad, 0xa8, 0xfd, 0x5a, 0x80, 0xc3,
	0x05, 0x17, 0xf9, 0x0e, 0x0a, 0x72, 0x80, 0xb3, 0x55, 0xd8, 0x68, 0x7e, 0x15, 0x91, 0xbc, 0x80,
	0x3b, 0x62, 0xb6, 0x2a, 0x2e, 0x47, 0xe1, 0xc6, 0x11, 0x0b, 0x45, 0x96, 0x5f, 0xf5, 0xff, 0x96,
	0xeb, 0x64, 0xcb, 0xb9, 0x3d, 0x27, 0x70, 0x86, 0xa2, 0x2f, 0xe9, 0x84, 0x42, 0x65, 0xa9, 0xae,
	0xeb, 0xa1, 0x2f, 0x68, 0x56, 0x86, 0x75, 0xd4, 0x4b, 0x4b, 0xd4, 0x8f, 0xa5, 0x06, 0x39, 0x87,
	0x5b, 0xf3, 0x57, 0xa8, 0xc5, 0x2b, 0x15, 0x94, 0xf0, 0x83, 0xd5, 0x85, 0x98, 0xb1, 0xd4, 0x56,
	0x9e, 0x6c, 0x39, 0xa6, 0x58, 0xb0, 0x91, 0xef, 0xe1, 0x36, 0x0d, 0x2e, 0x18, 0x86, 0xc2, 0x9d,
	0xf3, 0x95, 0xb6, 0xd7, 0x0b, 0xdc, 0x21, 0x19, 0x79, 0xce, 0x46, 0x3a, 0x60, 0x5e, 0x49, 0x4e,
	0xd2, 0x80, 0x79, 0x4c, 0xbc, 0x29, 0xed, 0x28, 0xbd, 0xca, 0xdb, 0x0b, 0xe5, 0x47, 0x54, 0xbc,
	0xa0, 0x7e, 0x8a, 0xce, 0x61, 0x46, 0x3a, 0xc9, 0x38, 0x47, 0x04, 0xe6, 0xc3, 0x75, 0xc7, 0x11,
	0xf5, 0x6b, 0x7f, 0x19, 0x60, 0x2e, 0xe6, 0x45, 0x1a, 0x90, 0xf7, 0xa3, 0xd7, 0x6a, 0x38, 0xd6,
	0x89, 0x59, 0x82, 0xc9, 0x57, 0x50, 0x98, 0xb0, 0xf1, 0x64, 0xdd, 0xfe, 0x3b, 0x0a, 0x4d, 0x1e,
	0xc1, 0x36, 0xf3, 0x90, 0xfa, 0xeb, 0x36, 0xd6, 0xd1, 0xf0, 0xfb, 0xbf, 0x18, 0x70, 0x70, 0x73,
	0x2e, 0x49, 0x11, 0x76, 0x87, 0xdd, 0x67, 0xdd, 0xde, 0x0f, 0x5d, 0x73, 0x8b, 0xec, 0x40, 0xae,
	0xd7, 0x35, 0x0d, 0xb2, 0x0b, 0xf9, 0x5e, 0xa7, 0x63, 0xe6, 0xc8, 0x1e, 0x14, 0x4e, 0xda, 0xcd,
	0x81, 0x99, 0x97, 0x5f, 0xad, 0x5e, 0xef, 0xd4, 0x2c, 0x90, 0x0f, 0x61, 0x5f, 0xda, 0x5c, 0x75,
	0xdc, 0x96, 0x8e, 0xe6, 0x70, 0xd0, 0x33, 0x77, 0xc8, 0x07, 0xb0, 0xd7, 0x69, 0x76, 0xdd, 0x5e,
	0xf7, 0xf4, 0xa5, 0xb9, 0x2b, 0x35, 0xda, 0xad, 0x9e, 0xb9, 0x27, 0xcd, 0xfd, 0xa1, 0xf3, 0xb4,
	0xf3, 0xb4, 0xed, 0x98, 0xfb, 0xd2, 0x7c, 0xec, 0xbc, 0x34, 0x81, 0x00, 0xec, 0x9c, 0xf6, 0x5a,
	0xcf, 0xda, 0xc7, 0x66, 0xb1, 0xf1, 0x6f, 0x0e, 0x60, 0x16, 0x17, 0x89, 0x61, 0xef, 0x09, 0x66,
	0x2b, 0xf7, 0x68, 0xe5, 0x6c, 0xbd, 0xf3, 0x59, 0x2f, 0x6f, 0xb4, 0xf3, 0x64, 0x0a, 0x45, 0xfd,
	0xde, 0xeb, 0xe3, 0xb7, 0x2b, 0xc9, 0xab, 0x7e, 0x19, 0x36, 0xbc, 0xf7, 0x67, 0x03, 0xf6, 0xe5,
	0x0b, 0xac, 0x4f, 0x5f, 0xaf, 0xe4, 0xbe, 0xfb, 0x6d, 0x2f, 0x7f, 0xb3, 0x39, 0x51, 0x3f, 0xf1,
	0x0f, 0x8d, 0xa3, 0xdf, 0x0c, 0xb8, 0xe7, 0xe1, 0x74, 0x95, 0xc2, 0xd1, 0xdc, 0x43, 0xd8, 0x97,
	0x2b, 0xd3, 0x37, 0x7e, 0xfc, 0x62, 0xcc, 0x84, 0x35, 0xa5, 0xa1, 0x60, 0xd6, 0x28, 0xb2, 0xd2,
	0x4b, 0xfb, 0x9a, 0x6f, 0xf3, 0xd1, 0x03, 0x1a, 0x33, 0x7b, 0x1c, 0xd9, 0x37, 0xfe, 0x0b, 0xfd,
	0x9e, 0xab, 0x9c, 0x5d, 0xdf, 0xd1, 0x8c, 0x99, 0x75, 0xac, 0xef, 0x18, 0x28, 0xef, 0xdf, 0x73,
	0xde, 0xf3, 0x66, 0xcc, 0xce, 0xb5, 0xf7, 0x5c, 0x7b, 0x2f, 0x76, 0xd4, 0xaa, 0x7e, 0xf9, 0x5f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x21, 0x74, 0x8c, 0x5f, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ThermostatClient is the client API for Thermostat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ThermostatClient interface {
	// Get the current state of the thermostat
	GetState(ctx context.Context, in *GetThermostatStateRequest, opts ...grpc.CallOption) (*ThermostatState, error)
	// Update the target state for the thermostat. the returned state may not be complete but will indicate at least the
	// current values for any set fields as part of the request.
	UpdateState(ctx context.Context, in *UpdateThermostatStateRequest, opts ...grpc.CallOption) (*ThermostatState, error)
	// Request notification of change to the thermostat state. The messages in the response stream may not be complete
	// but will indicate the changes as they occur. They should be merged with the full state as fetched by the GetState
	// method.
	PullState(ctx context.Context, in *PullThermostatStateRequest, opts ...grpc.CallOption) (Thermostat_PullStateClient, error)
}

type thermostatClient struct {
	cc *grpc.ClientConn
}

func NewThermostatClient(cc *grpc.ClientConn) ThermostatClient {
	return &thermostatClient{cc}
}

func (c *thermostatClient) GetState(ctx context.Context, in *GetThermostatStateRequest, opts ...grpc.CallOption) (*ThermostatState, error) {
	out := new(ThermostatState)
	err := c.cc.Invoke(ctx, "/smartcore.api.device.traits.Thermostat/GetState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thermostatClient) UpdateState(ctx context.Context, in *UpdateThermostatStateRequest, opts ...grpc.CallOption) (*ThermostatState, error) {
	out := new(ThermostatState)
	err := c.cc.Invoke(ctx, "/smartcore.api.device.traits.Thermostat/UpdateState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thermostatClient) PullState(ctx context.Context, in *PullThermostatStateRequest, opts ...grpc.CallOption) (Thermostat_PullStateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Thermostat_serviceDesc.Streams[0], "/smartcore.api.device.traits.Thermostat/PullState", opts...)
	if err != nil {
		return nil, err
	}
	x := &thermostatPullStateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Thermostat_PullStateClient interface {
	Recv() (*PullThermostatStateResponse, error)
	grpc.ClientStream
}

type thermostatPullStateClient struct {
	grpc.ClientStream
}

func (x *thermostatPullStateClient) Recv() (*PullThermostatStateResponse, error) {
	m := new(PullThermostatStateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ThermostatServer is the server API for Thermostat service.
type ThermostatServer interface {
	// Get the current state of the thermostat
	GetState(context.Context, *GetThermostatStateRequest) (*ThermostatState, error)
	// Update the target state for the thermostat. the returned state may not be complete but will indicate at least the
	// current values for any set fields as part of the request.
	UpdateState(context.Context, *UpdateThermostatStateRequest) (*ThermostatState, error)
	// Request notification of change to the thermostat state. The messages in the response stream may not be complete
	// but will indicate the changes as they occur. They should be merged with the full state as fetched by the GetState
	// method.
	PullState(*PullThermostatStateRequest, Thermostat_PullStateServer) error
}

// UnimplementedThermostatServer can be embedded to have forward compatible implementations.
type UnimplementedThermostatServer struct {
}

func (*UnimplementedThermostatServer) GetState(ctx context.Context, req *GetThermostatStateRequest) (*ThermostatState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetState not implemented")
}
func (*UnimplementedThermostatServer) UpdateState(ctx context.Context, req *UpdateThermostatStateRequest) (*ThermostatState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateState not implemented")
}
func (*UnimplementedThermostatServer) PullState(req *PullThermostatStateRequest, srv Thermostat_PullStateServer) error {
	return status.Errorf(codes.Unimplemented, "method PullState not implemented")
}

func RegisterThermostatServer(s *grpc.Server, srv ThermostatServer) {
	s.RegisterService(&_Thermostat_serviceDesc, srv)
}

func _Thermostat_GetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetThermostatStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThermostatServer).GetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.api.device.traits.Thermostat/GetState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThermostatServer).GetState(ctx, req.(*GetThermostatStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Thermostat_UpdateState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateThermostatStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThermostatServer).UpdateState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.api.device.traits.Thermostat/UpdateState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThermostatServer).UpdateState(ctx, req.(*UpdateThermostatStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Thermostat_PullState_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullThermostatStateRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ThermostatServer).PullState(m, &thermostatPullStateServer{stream})
}

type Thermostat_PullStateServer interface {
	Send(*PullThermostatStateResponse) error
	grpc.ServerStream
}

type thermostatPullStateServer struct {
	grpc.ServerStream
}

func (x *thermostatPullStateServer) Send(m *PullThermostatStateResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Thermostat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.api.device.traits.Thermostat",
	HandlerType: (*ThermostatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetState",
			Handler:    _Thermostat_GetState_Handler,
		},
		{
			MethodName: "UpdateState",
			Handler:    _Thermostat_UpdateState_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PullState",
			Handler:       _Thermostat_PullState_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "device/traits/thermostat.proto",
}
