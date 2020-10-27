// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: traits/on_off.proto

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

// Possible states for an On-Off device
type OnOrOff int32

const (
	OnOrOff_ON_OR_OFF_UNKNOWN OnOrOff = 0
	OnOrOff_ON_OR_OFF_ON      OnOrOff = 1
	OnOrOff_ON_OR_OFF_OFF     OnOrOff = 2
)

// Enum value maps for OnOrOff.
var (
	OnOrOff_name = map[int32]string{
		0: "ON_OR_OFF_UNKNOWN",
		1: "ON_OR_OFF_ON",
		2: "ON_OR_OFF_OFF",
	}
	OnOrOff_value = map[string]int32{
		"ON_OR_OFF_UNKNOWN": 0,
		"ON_OR_OFF_ON":      1,
		"ON_OR_OFF_OFF":     2,
	}
)

func (x OnOrOff) Enum() *OnOrOff {
	p := new(OnOrOff)
	*p = x
	return p
}

func (x OnOrOff) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OnOrOff) Descriptor() protoreflect.EnumDescriptor {
	return file_traits_on_off_proto_enumTypes[0].Descriptor()
}

func (OnOrOff) Type() protoreflect.EnumType {
	return &file_traits_on_off_proto_enumTypes[0]
}

func (x OnOrOff) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OnOrOff.Descriptor instead.
func (OnOrOff) EnumDescriptor() ([]byte, []int) {
	return file_traits_on_off_proto_rawDescGZIP(), []int{0}
}

type OnOff struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OnOrOff OnOrOff `protobuf:"varint,1,opt,name=on_or_off,json=onOrOff,proto3,enum=smartcore.traits.OnOrOff" json:"on_or_off,omitempty"`
}

func (x *OnOff) Reset() {
	*x = OnOff{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_on_off_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnOff) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnOff) ProtoMessage() {}

func (x *OnOff) ProtoReflect() protoreflect.Message {
	mi := &file_traits_on_off_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnOff.ProtoReflect.Descriptor instead.
func (*OnOff) Descriptor() ([]byte, []int) {
	return file_traits_on_off_proto_rawDescGZIP(), []int{0}
}

func (x *OnOff) GetOnOrOff() OnOrOff {
	if x != nil {
		return x.OnOrOff
	}
	return OnOrOff_ON_OR_OFF_UNKNOWN
}

// OnOffSupport describes the capabilities of devices implementing this trait
type OnOffSupport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// How a named device supports read/write/pull apis
	ResourceSupport *types.ResourceSupport `protobuf:"bytes,1,opt,name=resource_support,json=resourceSupport,proto3" json:"resource_support,omitempty"`
}

func (x *OnOffSupport) Reset() {
	*x = OnOffSupport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_on_off_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnOffSupport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnOffSupport) ProtoMessage() {}

func (x *OnOffSupport) ProtoReflect() protoreflect.Message {
	mi := &file_traits_on_off_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnOffSupport.ProtoReflect.Descriptor instead.
func (*OnOffSupport) Descriptor() ([]byte, []int) {
	return file_traits_on_off_proto_rawDescGZIP(), []int{1}
}

func (x *OnOffSupport) GetResourceSupport() *types.ResourceSupport {
	if x != nil {
		return x.ResourceSupport
	}
	return nil
}

type GetOnOffRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetOnOffRequest) Reset() {
	*x = GetOnOffRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_on_off_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOnOffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOnOffRequest) ProtoMessage() {}

func (x *GetOnOffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_on_off_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOnOffRequest.ProtoReflect.Descriptor instead.
func (*GetOnOffRequest) Descriptor() ([]byte, []int) {
	return file_traits_on_off_proto_rawDescGZIP(), []int{2}
}

func (x *GetOnOffRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateOnOffRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateOnOffRequest) Reset() {
	*x = UpdateOnOffRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_on_off_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOnOffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOnOffRequest) ProtoMessage() {}

func (x *UpdateOnOffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_on_off_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOnOffRequest.ProtoReflect.Descriptor instead.
func (*UpdateOnOffRequest) Descriptor() ([]byte, []int) {
	return file_traits_on_off_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateOnOffRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// A request to monitor the state of a device
type PullOnOffRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PullOnOffRequest) Reset() {
	*x = PullOnOffRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_on_off_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullOnOffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullOnOffRequest) ProtoMessage() {}

func (x *PullOnOffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_on_off_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullOnOffRequest.ProtoReflect.Descriptor instead.
func (*PullOnOffRequest) Descriptor() ([]byte, []int) {
	return file_traits_on_off_proto_rawDescGZIP(), []int{4}
}

func (x *PullOnOffRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Response to a Pull request for OnOffState changes
type PullOnOffResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Changes that have occurred since the last event
	Changes []*PullOnOffResponse_Change `protobuf:"bytes,1,rep,name=changes,proto3" json:"changes,omitempty"`
}

func (x *PullOnOffResponse) Reset() {
	*x = PullOnOffResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_on_off_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullOnOffResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullOnOffResponse) ProtoMessage() {}

func (x *PullOnOffResponse) ProtoReflect() protoreflect.Message {
	mi := &file_traits_on_off_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullOnOffResponse.ProtoReflect.Descriptor instead.
func (*PullOnOffResponse) Descriptor() ([]byte, []int) {
	return file_traits_on_off_proto_rawDescGZIP(), []int{5}
}

func (x *PullOnOffResponse) GetChanges() []*PullOnOffResponse_Change {
	if x != nil {
		return x.Changes
	}
	return nil
}

type DescribeOnOffRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the device
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DescribeOnOffRequest) Reset() {
	*x = DescribeOnOffRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_on_off_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeOnOffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeOnOffRequest) ProtoMessage() {}

func (x *DescribeOnOffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_on_off_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeOnOffRequest.ProtoReflect.Descriptor instead.
func (*DescribeOnOffRequest) Descriptor() ([]byte, []int) {
	return file_traits_on_off_proto_rawDescGZIP(), []int{6}
}

func (x *DescribeOnOffRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PullOnOffResponse_Change struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	OnOff *OnOff `protobuf:"bytes,2,opt,name=on_off,json=onOff,proto3" json:"on_off,omitempty"`
	// when the change occurred
	ChangeTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=change_time,json=changeTime,proto3" json:"change_time,omitempty"`
}

func (x *PullOnOffResponse_Change) Reset() {
	*x = PullOnOffResponse_Change{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_on_off_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullOnOffResponse_Change) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullOnOffResponse_Change) ProtoMessage() {}

func (x *PullOnOffResponse_Change) ProtoReflect() protoreflect.Message {
	mi := &file_traits_on_off_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullOnOffResponse_Change.ProtoReflect.Descriptor instead.
func (*PullOnOffResponse_Change) Descriptor() ([]byte, []int) {
	return file_traits_on_off_proto_rawDescGZIP(), []int{5, 0}
}

func (x *PullOnOffResponse_Change) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PullOnOffResponse_Change) GetOnOff() *OnOff {
	if x != nil {
		return x.OnOff
	}
	return nil
}

func (x *PullOnOffResponse_Change) GetChangeTime() *timestamp.Timestamp {
	if x != nil {
		return x.ChangeTime
	}
	return nil
}

var File_traits_on_off_proto protoreflect.FileDescriptor

var file_traits_on_off_proto_rawDesc = []byte{
	0x0a, 0x13, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2f, 0x6f, 0x6e, 0x5f, 0x6f, 0x66, 0x66, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x05, 0x4f, 0x6e,
	0x4f, 0x66, 0x66, 0x12, 0x35, 0x0a, 0x09, 0x6f, 0x6e, 0x5f, 0x6f, 0x72, 0x5f, 0x6f, 0x66, 0x66,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x4f, 0x6e, 0x4f, 0x72, 0x4f, 0x66,
	0x66, 0x52, 0x07, 0x6f, 0x6e, 0x4f, 0x72, 0x4f, 0x66, 0x66, 0x22, 0x5b, 0x0a, 0x0c, 0x4f, 0x6e,
	0x4f, 0x66, 0x66, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x4b, 0x0a, 0x10, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53,
	0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x25, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x6e,
	0x4f, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x28,
	0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x4f, 0x66, 0x66, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a, 0x10, 0x50, 0x75, 0x6c, 0x6c,
	0x4f, 0x6e, 0x4f, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0xe5, 0x01, 0x0a, 0x11, 0x50, 0x75, 0x6c, 0x6c, 0x4f, 0x6e, 0x4f, 0x66, 0x66, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x4f,
	0x6e, 0x4f, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x1a, 0x89, 0x01, 0x0a,
	0x06, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x6f,
	0x6e, 0x5f, 0x6f, 0x66, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x6d,
	0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x4f,
	0x6e, 0x4f, 0x66, 0x66, 0x52, 0x05, 0x6f, 0x6e, 0x4f, 0x66, 0x66, 0x12, 0x3b, 0x0a, 0x0b, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x14, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x4f, 0x6e, 0x4f, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x2a, 0x45, 0x0a, 0x07, 0x4f, 0x6e, 0x4f, 0x72, 0x4f, 0x66, 0x66, 0x12,
	0x15, 0x0a, 0x11, 0x4f, 0x4e, 0x5f, 0x4f, 0x52, 0x5f, 0x4f, 0x46, 0x46, 0x5f, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4f, 0x4e, 0x5f, 0x4f, 0x52, 0x5f,
	0x4f, 0x46, 0x46, 0x5f, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x4f, 0x4e, 0x5f, 0x4f,
	0x52, 0x5f, 0x4f, 0x46, 0x46, 0x5f, 0x4f, 0x46, 0x46, 0x10, 0x02, 0x32, 0xf8, 0x01, 0x0a, 0x08,
	0x4f, 0x6e, 0x4f, 0x66, 0x66, 0x41, 0x70, 0x69, 0x12, 0x46, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4f,
	0x6e, 0x4f, 0x66, 0x66, 0x12, 0x21, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x4f, 0x66, 0x66,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x4f, 0x6e, 0x4f, 0x66, 0x66,
	0x12, 0x4c, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x4f, 0x66, 0x66, 0x12,
	0x24, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69,
	0x74, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x4f, 0x66, 0x66, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x4f, 0x6e, 0x4f, 0x66, 0x66, 0x12, 0x56,
	0x0a, 0x09, 0x50, 0x75, 0x6c, 0x6c, 0x4f, 0x6e, 0x4f, 0x66, 0x66, 0x12, 0x22, 0x2e, 0x73, 0x6d,
	0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x50,
	0x75, 0x6c, 0x6c, 0x4f, 0x6e, 0x4f, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69,
	0x74, 0x73, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x4f, 0x6e, 0x4f, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x32, 0x64, 0x0a, 0x09, 0x4f, 0x6e, 0x4f, 0x66, 0x66, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x57, 0x0a, 0x0d, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4f,
	0x6e, 0x4f, 0x66, 0x66, 0x12, 0x26, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x4f, 0x6e, 0x4f, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73,
	0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e,
	0x4f, 0x6e, 0x4f, 0x66, 0x66, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x76, 0x0a, 0x14,
	0x64, 0x65, 0x76, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72,
	0x61, 0x69, 0x74, 0x73, 0x42, 0x0a, 0x4f, 0x6e, 0x4f, 0x66, 0x66, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x2e, 0x76, 0x61, 0x6e, 0x74, 0x69, 0x2e, 0x63, 0x6f,
	0x2e, 0x75, 0x6b, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x63,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0xaa, 0x02,
	0x10, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x74,
	0x73, 0xca, 0x02, 0x10, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x5c, 0x54, 0x72,
	0x61, 0x69, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_traits_on_off_proto_rawDescOnce sync.Once
	file_traits_on_off_proto_rawDescData = file_traits_on_off_proto_rawDesc
)

func file_traits_on_off_proto_rawDescGZIP() []byte {
	file_traits_on_off_proto_rawDescOnce.Do(func() {
		file_traits_on_off_proto_rawDescData = protoimpl.X.CompressGZIP(file_traits_on_off_proto_rawDescData)
	})
	return file_traits_on_off_proto_rawDescData
}

var file_traits_on_off_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_traits_on_off_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_traits_on_off_proto_goTypes = []interface{}{
	(OnOrOff)(0),                     // 0: smartcore.traits.OnOrOff
	(*OnOff)(nil),                    // 1: smartcore.traits.OnOff
	(*OnOffSupport)(nil),             // 2: smartcore.traits.OnOffSupport
	(*GetOnOffRequest)(nil),          // 3: smartcore.traits.GetOnOffRequest
	(*UpdateOnOffRequest)(nil),       // 4: smartcore.traits.UpdateOnOffRequest
	(*PullOnOffRequest)(nil),         // 5: smartcore.traits.PullOnOffRequest
	(*PullOnOffResponse)(nil),        // 6: smartcore.traits.PullOnOffResponse
	(*DescribeOnOffRequest)(nil),     // 7: smartcore.traits.DescribeOnOffRequest
	(*PullOnOffResponse_Change)(nil), // 8: smartcore.traits.PullOnOffResponse.Change
	(*types.ResourceSupport)(nil),    // 9: smartcore.types.ResourceSupport
	(*timestamp.Timestamp)(nil),      // 10: google.protobuf.Timestamp
}
var file_traits_on_off_proto_depIdxs = []int32{
	0,  // 0: smartcore.traits.OnOff.on_or_off:type_name -> smartcore.traits.OnOrOff
	9,  // 1: smartcore.traits.OnOffSupport.resource_support:type_name -> smartcore.types.ResourceSupport
	8,  // 2: smartcore.traits.PullOnOffResponse.changes:type_name -> smartcore.traits.PullOnOffResponse.Change
	1,  // 3: smartcore.traits.PullOnOffResponse.Change.on_off:type_name -> smartcore.traits.OnOff
	10, // 4: smartcore.traits.PullOnOffResponse.Change.change_time:type_name -> google.protobuf.Timestamp
	3,  // 5: smartcore.traits.OnOffApi.GetOnOff:input_type -> smartcore.traits.GetOnOffRequest
	4,  // 6: smartcore.traits.OnOffApi.UpdateOnOff:input_type -> smartcore.traits.UpdateOnOffRequest
	5,  // 7: smartcore.traits.OnOffApi.PullOnOff:input_type -> smartcore.traits.PullOnOffRequest
	7,  // 8: smartcore.traits.OnOffInfo.DescribeOnOff:input_type -> smartcore.traits.DescribeOnOffRequest
	1,  // 9: smartcore.traits.OnOffApi.GetOnOff:output_type -> smartcore.traits.OnOff
	1,  // 10: smartcore.traits.OnOffApi.UpdateOnOff:output_type -> smartcore.traits.OnOff
	6,  // 11: smartcore.traits.OnOffApi.PullOnOff:output_type -> smartcore.traits.PullOnOffResponse
	2,  // 12: smartcore.traits.OnOffInfo.DescribeOnOff:output_type -> smartcore.traits.OnOffSupport
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_traits_on_off_proto_init() }
func file_traits_on_off_proto_init() {
	if File_traits_on_off_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_traits_on_off_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnOff); i {
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
		file_traits_on_off_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnOffSupport); i {
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
		file_traits_on_off_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOnOffRequest); i {
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
		file_traits_on_off_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOnOffRequest); i {
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
		file_traits_on_off_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullOnOffRequest); i {
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
		file_traits_on_off_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullOnOffResponse); i {
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
		file_traits_on_off_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeOnOffRequest); i {
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
		file_traits_on_off_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullOnOffResponse_Change); i {
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
			RawDescriptor: file_traits_on_off_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_traits_on_off_proto_goTypes,
		DependencyIndexes: file_traits_on_off_proto_depIdxs,
		EnumInfos:         file_traits_on_off_proto_enumTypes,
		MessageInfos:      file_traits_on_off_proto_msgTypes,
	}.Build()
	File_traits_on_off_proto = out.File
	file_traits_on_off_proto_rawDesc = nil
	file_traits_on_off_proto_goTypes = nil
	file_traits_on_off_proto_depIdxs = nil
}
