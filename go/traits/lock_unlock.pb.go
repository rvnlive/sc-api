// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.1
// source: traits/lock_unlock.proto

package traits

import (
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

// Position defines the possible lock positions.
// Note that only LOCKED and UNLOCKED can be used during a write.
type LockUnlock_Position int32

const (
	// The lock position is unspecified.
	LockUnlock_POSITION_UNSPECIFIED LockUnlock_Position = 0
	// The device is locked.
	LockUnlock_LOCKED LockUnlock_Position = 1
	// The device is unlocked.
	LockUnlock_UNLOCKED LockUnlock_Position = 2
	// The device is in the process of locking.
	// Optional. Output only.
	LockUnlock_LOCKING LockUnlock_Position = 3
	// The device is in the process of unlocking.
	// Optional. Output only.
	LockUnlock_UNLOCKING LockUnlock_Position = 4
)

// Enum value maps for LockUnlock_Position.
var (
	LockUnlock_Position_name = map[int32]string{
		0: "POSITION_UNSPECIFIED",
		1: "LOCKED",
		2: "UNLOCKED",
		3: "LOCKING",
		4: "UNLOCKING",
	}
	LockUnlock_Position_value = map[string]int32{
		"POSITION_UNSPECIFIED": 0,
		"LOCKED":               1,
		"UNLOCKED":             2,
		"LOCKING":              3,
		"UNLOCKING":            4,
	}
)

func (x LockUnlock_Position) Enum() *LockUnlock_Position {
	p := new(LockUnlock_Position)
	*p = x
	return p
}

func (x LockUnlock_Position) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LockUnlock_Position) Descriptor() protoreflect.EnumDescriptor {
	return file_traits_lock_unlock_proto_enumTypes[0].Descriptor()
}

func (LockUnlock_Position) Type() protoreflect.EnumType {
	return &file_traits_lock_unlock_proto_enumTypes[0]
}

func (x LockUnlock_Position) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LockUnlock_Position.Descriptor instead.
func (LockUnlock_Position) EnumDescriptor() ([]byte, []int) {
	return file_traits_lock_unlock_proto_rawDescGZIP(), []int{0, 0}
}

// LockUnlock models the possible states of a lockable device.
type LockUnlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Position reports whether the device is locked or unlocked, or transitioning.
	Position LockUnlock_Position `protobuf:"varint,1,opt,name=position,proto3,enum=smartcore.traits.LockUnlock_Position" json:"position,omitempty"`
	// Jammed reports whether the lock has jammed and is unable to reach it's position.
	// Output only.
	Jammed bool `protobuf:"varint,2,opt,name=jammed,proto3" json:"jammed,omitempty"`
}

func (x *LockUnlock) Reset() {
	*x = LockUnlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_lock_unlock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockUnlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockUnlock) ProtoMessage() {}

func (x *LockUnlock) ProtoReflect() protoreflect.Message {
	mi := &file_traits_lock_unlock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockUnlock.ProtoReflect.Descriptor instead.
func (*LockUnlock) Descriptor() ([]byte, []int) {
	return file_traits_lock_unlock_proto_rawDescGZIP(), []int{0}
}

func (x *LockUnlock) GetPosition() LockUnlock_Position {
	if x != nil {
		return x.Position
	}
	return LockUnlock_POSITION_UNSPECIFIED
}

func (x *LockUnlock) GetJammed() bool {
	if x != nil {
		return x.Jammed
	}
	return false
}

type GetLockUnlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the device to fetch the state for
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Fields to fetch relative to the LockUnlock type
	ReadMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=read_mask,json=readMask,proto3" json:"read_mask,omitempty"`
}

func (x *GetLockUnlockRequest) Reset() {
	*x = GetLockUnlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_lock_unlock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLockUnlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLockUnlockRequest) ProtoMessage() {}

func (x *GetLockUnlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_lock_unlock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLockUnlockRequest.ProtoReflect.Descriptor instead.
func (*GetLockUnlockRequest) Descriptor() ([]byte, []int) {
	return file_traits_lock_unlock_proto_rawDescGZIP(), []int{1}
}

func (x *GetLockUnlockRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetLockUnlockRequest) GetReadMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.ReadMask
	}
	return nil
}

type UpdateLockUnlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the device to fetch the state for
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The new value
	LockUnlock *LockUnlock `protobuf:"bytes,2,opt,name=lock_unlock,json=lockUnlock,proto3" json:"lock_unlock,omitempty"`
	// Fields to fetch relative to the LockUnlock type
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateLockUnlockRequest) Reset() {
	*x = UpdateLockUnlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_lock_unlock_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLockUnlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLockUnlockRequest) ProtoMessage() {}

func (x *UpdateLockUnlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_lock_unlock_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLockUnlockRequest.ProtoReflect.Descriptor instead.
func (*UpdateLockUnlockRequest) Descriptor() ([]byte, []int) {
	return file_traits_lock_unlock_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateLockUnlockRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateLockUnlockRequest) GetLockUnlock() *LockUnlock {
	if x != nil {
		return x.LockUnlock
	}
	return nil
}

func (x *UpdateLockUnlockRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type PullLockUnlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the device to fetch the state for
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Fields to fetch relative to the LockUnlock type
	ReadMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=read_mask,json=readMask,proto3" json:"read_mask,omitempty"`
	// When true the device will only send changes to the resource value.
	// The default behaviour is to send the current value immediately followed by any updates as they happen.
	UpdatesOnly bool `protobuf:"varint,3,opt,name=updates_only,json=updatesOnly,proto3" json:"updates_only,omitempty"`
}

func (x *PullLockUnlockRequest) Reset() {
	*x = PullLockUnlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_lock_unlock_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullLockUnlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullLockUnlockRequest) ProtoMessage() {}

func (x *PullLockUnlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_traits_lock_unlock_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullLockUnlockRequest.ProtoReflect.Descriptor instead.
func (*PullLockUnlockRequest) Descriptor() ([]byte, []int) {
	return file_traits_lock_unlock_proto_rawDescGZIP(), []int{3}
}

func (x *PullLockUnlockRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PullLockUnlockRequest) GetReadMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.ReadMask
	}
	return nil
}

func (x *PullLockUnlockRequest) GetUpdatesOnly() bool {
	if x != nil {
		return x.UpdatesOnly
	}
	return false
}

type PullLockUnlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Changes since the last message.
	Changes []*PullLockUnlockResponse_Change `protobuf:"bytes,1,rep,name=changes,proto3" json:"changes,omitempty"`
}

func (x *PullLockUnlockResponse) Reset() {
	*x = PullLockUnlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_lock_unlock_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullLockUnlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullLockUnlockResponse) ProtoMessage() {}

func (x *PullLockUnlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_traits_lock_unlock_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullLockUnlockResponse.ProtoReflect.Descriptor instead.
func (*PullLockUnlockResponse) Descriptor() ([]byte, []int) {
	return file_traits_lock_unlock_proto_rawDescGZIP(), []int{4}
}

func (x *PullLockUnlockResponse) GetChanges() []*PullLockUnlockResponse_Change {
	if x != nil {
		return x.Changes
	}
	return nil
}

type PullLockUnlockResponse_Change struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the device that issued the change.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// When the change occurred
	ChangeTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=change_time,json=changeTime,proto3" json:"change_time,omitempty"`
	// The new value for the lock.
	LockUnlock *LockUnlock `protobuf:"bytes,3,opt,name=lock_unlock,json=lockUnlock,proto3" json:"lock_unlock,omitempty"`
}

func (x *PullLockUnlockResponse_Change) Reset() {
	*x = PullLockUnlockResponse_Change{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traits_lock_unlock_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullLockUnlockResponse_Change) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullLockUnlockResponse_Change) ProtoMessage() {}

func (x *PullLockUnlockResponse_Change) ProtoReflect() protoreflect.Message {
	mi := &file_traits_lock_unlock_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullLockUnlockResponse_Change.ProtoReflect.Descriptor instead.
func (*PullLockUnlockResponse_Change) Descriptor() ([]byte, []int) {
	return file_traits_lock_unlock_proto_rawDescGZIP(), []int{4, 0}
}

func (x *PullLockUnlockResponse_Change) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PullLockUnlockResponse_Change) GetChangeTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ChangeTime
	}
	return nil
}

func (x *PullLockUnlockResponse_Change) GetLockUnlock() *LockUnlock {
	if x != nil {
		return x.LockUnlock
	}
	return nil
}

var File_traits_lock_unlock_proto protoreflect.FileDescriptor

var file_traits_lock_unlock_proto_rawDesc = []byte{
	0x0a, 0x18, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2f, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x75, 0x6e,
	0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73, 0x6d, 0x61, 0x72,
	0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc3, 0x01, 0x0a, 0x0a, 0x4c, 0x6f, 0x63, 0x6b, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x41,
	0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x25, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61,
	0x69, 0x74, 0x73, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6a, 0x61, 0x6d, 0x6d, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x6a, 0x61, 0x6d, 0x6d, 0x65, 0x64, 0x22, 0x5a, 0x0a, 0x08, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x55,
	0x4e, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x4f, 0x43,
	0x4b, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x4c, 0x4f, 0x43, 0x4b,
	0x49, 0x4e, 0x47, 0x10, 0x04, 0x22, 0x63, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x6b,
	0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x37, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0xa9, 0x01, 0x0a, 0x17, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69,
	0x74, 0x73, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x0a, 0x6c,
	0x6f, 0x63, 0x6b, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x87, 0x01, 0x0a, 0x15, 0x50, 0x75, 0x6c, 0x6c, 0x4c,
	0x6f, 0x63, 0x6b, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d,
	0x61, 0x73, 0x6b, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x21, 0x0a,
	0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x4f, 0x6e, 0x6c, 0x79,
	0x22, 0xfe, 0x01, 0x0a, 0x16, 0x50, 0x75, 0x6c, 0x6c, 0x4c, 0x6f, 0x63, 0x6b, 0x55, 0x6e, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x07, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x73,
	0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e,
	0x50, 0x75, 0x6c, 0x6c, 0x4c, 0x6f, 0x63, 0x6b, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x07, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x1a, 0x98, 0x01, 0x0a, 0x06, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x75, 0x6e, 0x6c, 0x6f, 0x63,
	0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x55,
	0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x6b, 0x55, 0x6e, 0x6c, 0x6f, 0x63,
	0x6b, 0x32, 0xaa, 0x02, 0x0a, 0x0d, 0x4c, 0x6f, 0x63, 0x6b, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b,
	0x41, 0x70, 0x69, 0x12, 0x55, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x55, 0x6e,
	0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x26, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x55,
	0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73,
	0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e,
	0x4c, 0x6f, 0x63, 0x6b, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x5b, 0x0a, 0x10, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x29,
	0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74,
	0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x55, 0x6e, 0x6c, 0x6f,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x6d, 0x61, 0x72,
	0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x4c, 0x6f, 0x63,
	0x6b, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x65, 0x0a, 0x0e, 0x50, 0x75, 0x6c, 0x6c, 0x4c,
	0x6f, 0x63, 0x6b, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x27, 0x2e, 0x73, 0x6d, 0x61, 0x72,
	0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x50, 0x75, 0x6c,
	0x6c, 0x4c, 0x6f, 0x63, 0x6b, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x28, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74,
	0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x4c, 0x6f, 0x63, 0x6b, 0x55, 0x6e,
	0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x32, 0x10,
	0x0a, 0x0e, 0x4c, 0x6f, 0x63, 0x6b, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f,
	0x42, 0x7a, 0x0a, 0x14, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x42, 0x0f, 0x4c, 0x6f, 0x63, 0x6b, 0x55, 0x6e,
	0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x2d, 0x63, 0x6f,
	0x72, 0x65, 0x2d, 0x6f, 0x73, 0x2f, 0x73, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f,
	0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0xaa, 0x02, 0x10, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x74, 0x73, 0xca, 0x02, 0x10, 0x53, 0x6d, 0x61, 0x72,
	0x74, 0x63, 0x6f, 0x72, 0x65, 0x5c, 0x54, 0x72, 0x61, 0x69, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_traits_lock_unlock_proto_rawDescOnce sync.Once
	file_traits_lock_unlock_proto_rawDescData = file_traits_lock_unlock_proto_rawDesc
)

func file_traits_lock_unlock_proto_rawDescGZIP() []byte {
	file_traits_lock_unlock_proto_rawDescOnce.Do(func() {
		file_traits_lock_unlock_proto_rawDescData = protoimpl.X.CompressGZIP(file_traits_lock_unlock_proto_rawDescData)
	})
	return file_traits_lock_unlock_proto_rawDescData
}

var file_traits_lock_unlock_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_traits_lock_unlock_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_traits_lock_unlock_proto_goTypes = []interface{}{
	(LockUnlock_Position)(0),              // 0: smartcore.traits.LockUnlock.Position
	(*LockUnlock)(nil),                    // 1: smartcore.traits.LockUnlock
	(*GetLockUnlockRequest)(nil),          // 2: smartcore.traits.GetLockUnlockRequest
	(*UpdateLockUnlockRequest)(nil),       // 3: smartcore.traits.UpdateLockUnlockRequest
	(*PullLockUnlockRequest)(nil),         // 4: smartcore.traits.PullLockUnlockRequest
	(*PullLockUnlockResponse)(nil),        // 5: smartcore.traits.PullLockUnlockResponse
	(*PullLockUnlockResponse_Change)(nil), // 6: smartcore.traits.PullLockUnlockResponse.Change
	(*fieldmaskpb.FieldMask)(nil),         // 7: google.protobuf.FieldMask
	(*timestamppb.Timestamp)(nil),         // 8: google.protobuf.Timestamp
}
var file_traits_lock_unlock_proto_depIdxs = []int32{
	0,  // 0: smartcore.traits.LockUnlock.position:type_name -> smartcore.traits.LockUnlock.Position
	7,  // 1: smartcore.traits.GetLockUnlockRequest.read_mask:type_name -> google.protobuf.FieldMask
	1,  // 2: smartcore.traits.UpdateLockUnlockRequest.lock_unlock:type_name -> smartcore.traits.LockUnlock
	7,  // 3: smartcore.traits.UpdateLockUnlockRequest.update_mask:type_name -> google.protobuf.FieldMask
	7,  // 4: smartcore.traits.PullLockUnlockRequest.read_mask:type_name -> google.protobuf.FieldMask
	6,  // 5: smartcore.traits.PullLockUnlockResponse.changes:type_name -> smartcore.traits.PullLockUnlockResponse.Change
	8,  // 6: smartcore.traits.PullLockUnlockResponse.Change.change_time:type_name -> google.protobuf.Timestamp
	1,  // 7: smartcore.traits.PullLockUnlockResponse.Change.lock_unlock:type_name -> smartcore.traits.LockUnlock
	2,  // 8: smartcore.traits.LockUnlockApi.GetLockUnlock:input_type -> smartcore.traits.GetLockUnlockRequest
	3,  // 9: smartcore.traits.LockUnlockApi.UpdateLockUnlock:input_type -> smartcore.traits.UpdateLockUnlockRequest
	4,  // 10: smartcore.traits.LockUnlockApi.PullLockUnlock:input_type -> smartcore.traits.PullLockUnlockRequest
	1,  // 11: smartcore.traits.LockUnlockApi.GetLockUnlock:output_type -> smartcore.traits.LockUnlock
	1,  // 12: smartcore.traits.LockUnlockApi.UpdateLockUnlock:output_type -> smartcore.traits.LockUnlock
	5,  // 13: smartcore.traits.LockUnlockApi.PullLockUnlock:output_type -> smartcore.traits.PullLockUnlockResponse
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_traits_lock_unlock_proto_init() }
func file_traits_lock_unlock_proto_init() {
	if File_traits_lock_unlock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_traits_lock_unlock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockUnlock); i {
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
		file_traits_lock_unlock_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLockUnlockRequest); i {
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
		file_traits_lock_unlock_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateLockUnlockRequest); i {
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
		file_traits_lock_unlock_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullLockUnlockRequest); i {
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
		file_traits_lock_unlock_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullLockUnlockResponse); i {
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
		file_traits_lock_unlock_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullLockUnlockResponse_Change); i {
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
			RawDescriptor: file_traits_lock_unlock_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_traits_lock_unlock_proto_goTypes,
		DependencyIndexes: file_traits_lock_unlock_proto_depIdxs,
		EnumInfos:         file_traits_lock_unlock_proto_enumTypes,
		MessageInfos:      file_traits_lock_unlock_proto_msgTypes,
	}.Build()
	File_traits_lock_unlock_proto = out.File
	file_traits_lock_unlock_proto_rawDesc = nil
	file_traits_lock_unlock_proto_goTypes = nil
	file_traits_lock_unlock_proto_depIdxs = nil
}
