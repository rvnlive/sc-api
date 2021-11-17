// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: traits/count.proto

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

// Count represents how many of a thing exist. The total measured count is added - removed. Some devices may adjust
// added up and down and leave removed at 0, some might increase both properties.
type Count struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Total number of added items
	Added int32 `protobuf:"varint,1,opt,name=added,proto3" json:"added,omitempty"`
	// Total number of removed items
	Removed int32 `protobuf:"varint,2,opt,name=removed,proto3" json:"removed,omitempty"`
	// The time the counts were reset
	ResetTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=reset_time,json=resetTime,proto3" json:"reset_time,omitempty"`
}

func (x *Count) Reset() {
	*x = Count{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_count_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Count) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Count) ProtoMessage() {}

func (x *Count) ProtoReflect() protoreflect.Message {
	mi := &file_traits_count_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Count.ProtoReflect.Descriptor instead.
func (*Count) Descriptor() ([]byte, []int) {
	return file_traits_count_proto_rawDescGZIP(), []int{0}
}

func (x *Count) GetAdded() int32 {
	if x != nil {
		return x.Added
	}
	return 0
}

func (x *Count) GetRemoved() int32 {
	if x != nil {
		return x.Removed
	}
	return 0
}

func (x *Count) GetResetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ResetTime
	}
	return nil
}

// CountSupport describes the capabilities of devices implementing this trait
type CountSupport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// How a named device supports read/write/pull apis
	ResourceSupport *types.ResourceSupport `protobuf:"bytes,1,opt,name=resource_support,json=resourceSupport,proto3" json:"resource_support,omitempty"`
	// Two way counters count additions and removals separately. Do NOT use this as a way to ignore the Count.removed
	// property, the total count is always added - removed.
	TwoWay bool `protobuf:"varint,2,opt,name=two_way,json=twoWay,proto3" json:"two_way,omitempty"`
	// Does the api support resetting the count timer.
	SupportsReset bool `protobuf:"varint,3,opt,name=supports_reset,json=supportsReset,proto3" json:"supports_reset,omitempty"`
	// When updating the count does this device support the delta property
	SupportsDelta bool `protobuf:"varint,4,opt,name=supports_delta,json=supportsDelta,proto3" json:"supports_delta,omitempty"`
}

func (x *CountSupport) Reset() {
	*x = CountSupport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_count_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountSupport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountSupport) ProtoMessage() {}

func (x *CountSupport) ProtoReflect() protoreflect.Message {
	mi := &file_traits_count_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountSupport.ProtoReflect.Descriptor instead.
func (*CountSupport) Descriptor() ([]byte, []int) {
	return file_traits_count_proto_rawDescGZIP(), []int{1}
}

func (x *CountSupport) GetResourceSupport() *types.ResourceSupport {
	if x != nil {
		return x.ResourceSupport
	}
	return nil
}

func (x *CountSupport) GetTwoWay() bool {
	if x != nil {
		return x.TwoWay
	}
	return false
}

func (x *CountSupport) GetSupportsReset() bool {
	if x != nil {
		return x.SupportsReset
	}
	return false
}

func (x *CountSupport) GetSupportsDelta() bool {
	if x != nil {
		return x.SupportsDelta
	}
	return false
}

type GetCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetCountRequest) Reset() {
	*x = GetCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_count_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCountRequest) ProtoMessage() {}

func (x *GetCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_count_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCountRequest.ProtoReflect.Descriptor instead.
func (*GetCountRequest) Descriptor() ([]byte, []int) {
	return file_traits_count_proto_rawDescGZIP(), []int{2}
}

func (x *GetCountRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ResetCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// if absent then the server timestamp will be used
	ResetTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=reset_time,json=resetTime,proto3" json:"reset_time,omitempty"`
}

func (x *ResetCountRequest) Reset() {
	*x = ResetCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_count_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetCountRequest) ProtoMessage() {}

func (x *ResetCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_count_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetCountRequest.ProtoReflect.Descriptor instead.
func (*ResetCountRequest) Descriptor() ([]byte, []int) {
	return file_traits_count_proto_rawDescGZIP(), []int{3}
}

func (x *ResetCountRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResetCountRequest) GetResetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ResetTime
	}
	return nil
}

type UpdateCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Count      *Count                 `protobuf:"bytes,2,opt,name=count,proto3" json:"count,omitempty"`
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// if true the added and removed properties of count should be treated as a change not an absolute value.
	Delta bool `protobuf:"varint,4,opt,name=delta,proto3" json:"delta,omitempty"`
}

func (x *UpdateCountRequest) Reset() {
	*x = UpdateCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_count_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCountRequest) ProtoMessage() {}

func (x *UpdateCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_count_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCountRequest.ProtoReflect.Descriptor instead.
func (*UpdateCountRequest) Descriptor() ([]byte, []int) {
	return file_traits_count_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateCountRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateCountRequest) GetCount() *Count {
	if x != nil {
		return x.Count
	}
	return nil
}

func (x *UpdateCountRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (x *UpdateCountRequest) GetDelta() bool {
	if x != nil {
		return x.Delta
	}
	return false
}

type PullCountsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PullCountsRequest) Reset() {
	*x = PullCountsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_count_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullCountsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullCountsRequest) ProtoMessage() {}

func (x *PullCountsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_count_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullCountsRequest.ProtoReflect.Descriptor instead.
func (*PullCountsRequest) Descriptor() ([]byte, []int) {
	return file_traits_count_proto_rawDescGZIP(), []int{5}
}

func (x *PullCountsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PullCountsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Changes []*PullCountsResponse_Change `protobuf:"bytes,1,rep,name=changes,proto3" json:"changes,omitempty"`
}

func (x *PullCountsResponse) Reset() {
	*x = PullCountsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_count_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullCountsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullCountsResponse) ProtoMessage() {}

func (x *PullCountsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_traits_count_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullCountsResponse.ProtoReflect.Descriptor instead.
func (*PullCountsResponse) Descriptor() ([]byte, []int) {
	return file_traits_count_proto_rawDescGZIP(), []int{6}
}

func (x *PullCountsResponse) GetChanges() []*PullCountsResponse_Change {
	if x != nil {
		return x.Changes
	}
	return nil
}

type DescribeCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the device
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DescribeCountRequest) Reset() {
	*x = DescribeCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_count_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeCountRequest) ProtoMessage() {}

func (x *DescribeCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_count_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeCountRequest.ProtoReflect.Descriptor instead.
func (*DescribeCountRequest) Descriptor() ([]byte, []int) {
	return file_traits_count_proto_rawDescGZIP(), []int{7}
}

func (x *DescribeCountRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PullCountsResponse_Change struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Count *Count `protobuf:"bytes,2,opt,name=count,proto3" json:"count,omitempty"`
	// when the change occurred
	ChangeTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=change_time,json=changeTime,proto3" json:"change_time,omitempty"`
}

func (x *PullCountsResponse_Change) Reset() {
	*x = PullCountsResponse_Change{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_count_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullCountsResponse_Change) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullCountsResponse_Change) ProtoMessage() {}

func (x *PullCountsResponse_Change) ProtoReflect() protoreflect.Message {
	mi := &file_traits_count_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullCountsResponse_Change.ProtoReflect.Descriptor instead.
func (*PullCountsResponse_Change) Descriptor() ([]byte, []int) {
	return file_traits_count_proto_rawDescGZIP(), []int{6, 0}
}

func (x *PullCountsResponse_Change) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PullCountsResponse_Change) GetCount() *Count {
	if x != nil {
		return x.Count
	}
	return nil
}

func (x *PullCountsResponse_Change) GetChangeTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ChangeTime
	}
	return nil
}

var File_traits_count_proto protoreflect.FileDescriptor

var file_traits_count_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d,
	0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a, 0x05, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x72, 0x65, 0x73, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0xc2, 0x01, 0x0a, 0x0c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x4b, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x6d, 0x61,
	0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x0f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x74, 0x77, 0x6f, 0x5f, 0x77, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x74, 0x77, 0x6f, 0x57, 0x61, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x65, 0x74, 0x12, 0x25, 0x0a,
	0x0e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x44,
	0x65, 0x6c, 0x74, 0x61, 0x22, 0x25, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x62, 0x0a, 0x11, 0x52,
	0x65, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x72, 0x65, 0x73, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0xaa, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x6d, 0x61, 0x72,
	0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x22, 0x27, 0x0a, 0x11,
	0x50, 0x75, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xe6, 0x01, 0x0a, 0x12, 0x50, 0x75, 0x6c, 0x6c, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x07,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73,
	0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x73, 0x1a, 0x88, 0x01, 0x0a, 0x06, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72,
	0x61, 0x69, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x2a,
	0x0a, 0x14, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xc7, 0x02, 0x0a, 0x08, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x41, 0x70, 0x69, 0x12, 0x46, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x21, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x4a, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x2e,
	0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73,
	0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74,
	0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x4c, 0x0a, 0x0b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x73, 0x6d, 0x61,
	0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61,
	0x69, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x59, 0x0a, 0x0a, 0x50, 0x75, 0x6c,
	0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x23, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73,
	0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e,
	0x50, 0x75, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x30, 0x01, 0x32, 0x64, 0x0a, 0x09, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x57, 0x0a, 0x0d, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x26, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74,
	0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x6d, 0x61,
	0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x75, 0x0a, 0x14, 0x64, 0x65,
	0x76, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69,
	0x74, 0x73, 0x42, 0x0a, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x61,
	0x72, 0x74, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2d, 0x6f, 0x73, 0x2f, 0x73, 0x63, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0xaa, 0x02, 0x10, 0x53, 0x6d,
	0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x74, 0x73, 0xca, 0x02,
	0x10, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x5c, 0x54, 0x72, 0x61, 0x69, 0x74,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_traits_count_proto_rawDescOnce sync.Once
	file_traits_count_proto_rawDescData = file_traits_count_proto_rawDesc
)

func file_traits_count_proto_rawDescGZIP() []byte {
	file_traits_count_proto_rawDescOnce.Do(func() {
		file_traits_count_proto_rawDescData = protoimpl.X.CompressGZIP(file_traits_count_proto_rawDescData)
	})
	return file_traits_count_proto_rawDescData
}

var file_traits_count_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_traits_count_proto_goTypes = []interface{}{
	(*Count)(nil),                     // 0: smartcore.traits.Count
	(*CountSupport)(nil),              // 1: smartcore.traits.CountSupport
	(*GetCountRequest)(nil),           // 2: smartcore.traits.GetCountRequest
	(*ResetCountRequest)(nil),         // 3: smartcore.traits.ResetCountRequest
	(*UpdateCountRequest)(nil),        // 4: smartcore.traits.UpdateCountRequest
	(*PullCountsRequest)(nil),         // 5: smartcore.traits.PullCountsRequest
	(*PullCountsResponse)(nil),        // 6: smartcore.traits.PullCountsResponse
	(*DescribeCountRequest)(nil),      // 7: smartcore.traits.DescribeCountRequest
	(*PullCountsResponse_Change)(nil), // 8: smartcore.traits.PullCountsResponse.Change
	(*timestamppb.Timestamp)(nil),     // 9: google.protobuf.Timestamp
	(*types.ResourceSupport)(nil),     // 10: smartcore.types.ResourceSupport
	(*fieldmaskpb.FieldMask)(nil),     // 11: google.protobuf.FieldMask
}
var file_traits_count_proto_depIdxs = []int32{
	9,  // 0: smartcore.traits.Count.reset_time:type_name -> google.protobuf.Timestamp
	10, // 1: smartcore.traits.CountSupport.resource_support:type_name -> smartcore.types.ResourceSupport
	9,  // 2: smartcore.traits.ResetCountRequest.reset_time:type_name -> google.protobuf.Timestamp
	0,  // 3: smartcore.traits.UpdateCountRequest.count:type_name -> smartcore.traits.Count
	11, // 4: smartcore.traits.UpdateCountRequest.update_mask:type_name -> google.protobuf.FieldMask
	8,  // 5: smartcore.traits.PullCountsResponse.changes:type_name -> smartcore.traits.PullCountsResponse.Change
	0,  // 6: smartcore.traits.PullCountsResponse.Change.count:type_name -> smartcore.traits.Count
	9,  // 7: smartcore.traits.PullCountsResponse.Change.change_time:type_name -> google.protobuf.Timestamp
	2,  // 8: smartcore.traits.CountApi.GetCount:input_type -> smartcore.traits.GetCountRequest
	3,  // 9: smartcore.traits.CountApi.ResetCount:input_type -> smartcore.traits.ResetCountRequest
	4,  // 10: smartcore.traits.CountApi.UpdateCount:input_type -> smartcore.traits.UpdateCountRequest
	5,  // 11: smartcore.traits.CountApi.PullCounts:input_type -> smartcore.traits.PullCountsRequest
	7,  // 12: smartcore.traits.CountInfo.DescribeCount:input_type -> smartcore.traits.DescribeCountRequest
	0,  // 13: smartcore.traits.CountApi.GetCount:output_type -> smartcore.traits.Count
	0,  // 14: smartcore.traits.CountApi.ResetCount:output_type -> smartcore.traits.Count
	0,  // 15: smartcore.traits.CountApi.UpdateCount:output_type -> smartcore.traits.Count
	6,  // 16: smartcore.traits.CountApi.PullCounts:output_type -> smartcore.traits.PullCountsResponse
	1,  // 17: smartcore.traits.CountInfo.DescribeCount:output_type -> smartcore.traits.CountSupport
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_traits_count_proto_init() }
func file_traits_count_proto_init() {
	if File_traits_count_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_traits_count_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Count); i {
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
		file_traits_count_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountSupport); i {
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
		file_traits_count_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCountRequest); i {
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
		file_traits_count_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetCountRequest); i {
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
		file_traits_count_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCountRequest); i {
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
		file_traits_count_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullCountsRequest); i {
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
		file_traits_count_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullCountsResponse); i {
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
		file_traits_count_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeCountRequest); i {
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
		file_traits_count_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullCountsResponse_Change); i {
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
			RawDescriptor: file_traits_count_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_traits_count_proto_goTypes,
		DependencyIndexes: file_traits_count_proto_depIdxs,
		MessageInfos:      file_traits_count_proto_msgTypes,
	}.Build()
	File_traits_count_proto = out.File
	file_traits_count_proto_rawDesc = nil
	file_traits_count_proto_goTypes = nil
	file_traits_count_proto_depIdxs = nil
}
