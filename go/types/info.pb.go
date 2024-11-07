// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.3
// source: types/info.proto

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// PullSupport describes how Pull methods are implemented by the server
type PullSupport int32

const (
	// How subscriptions are implemented is not known.
	PullSupport_PULL_SUPPORT_UNSPECIFIED PullSupport = 0
	// Subscribing is supported natively by the underlying system.
	PullSupport_PULL_SUPPORT_NATIVE PullSupport = 1
	// Subscribing is supported in the driver, rather than natively.
	// Usually the driver will poll the system for data.
	PullSupport_PULL_SUPPORT_EMULATED PullSupport = 2
)

// Enum value maps for PullSupport.
var (
	PullSupport_name = map[int32]string{
		0: "PULL_SUPPORT_UNSPECIFIED",
		1: "PULL_SUPPORT_NATIVE",
		2: "PULL_SUPPORT_EMULATED",
	}
	PullSupport_value = map[string]int32{
		"PULL_SUPPORT_UNSPECIFIED": 0,
		"PULL_SUPPORT_NATIVE":      1,
		"PULL_SUPPORT_EMULATED":    2,
	}
)

func (x PullSupport) Enum() *PullSupport {
	p := new(PullSupport)
	*p = x
	return p
}

func (x PullSupport) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PullSupport) Descriptor() protoreflect.EnumDescriptor {
	return file_types_info_proto_enumTypes[0].Descriptor()
}

func (PullSupport) Type() protoreflect.EnumType {
	return &file_types_info_proto_enumTypes[0]
}

func (x PullSupport) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PullSupport.Descriptor instead.
func (PullSupport) EnumDescriptor() ([]byte, []int) {
	return file_types_info_proto_rawDescGZIP(), []int{0}
}

// ResourceSupport describes how a trait resource supports read/write/subscribe apis.
//
// For sensor traits any write oriented properties should be and will be ignored.
type ResourceSupport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the device supports querying, i.e. get apis
	Readable bool `protobuf:"varint,1,opt,name=readable,proto3" json:"readable,omitempty"`
	// Whether the device supports writable actions, i.e. update or action apis
	Writable bool `protobuf:"varint,2,opt,name=writable,proto3" json:"writable,omitempty"`
	// Whether the device supports subscriptions, i.e. pull apis
	Observable bool `protobuf:"varint,3,opt,name=observable,proto3" json:"observable,omitempty"`
	// Which fields can be written to.
	// Relative to the resource type, not the message (FooRequest) type.
	// If this field is absent, the list of updatable fields is unknown, assume all.
	// Check the response to your Update request for the values that were updated.
	// An attempt to explicitly (via update_mask) write to a non-writable field will error.
	WritableFields *fieldmaskpb.FieldMask `protobuf:"bytes,4,opt,name=writable_fields,json=writableFields,proto3" json:"writable_fields,omitempty"`
	// How Pull is implemented by the server
	PullSupport PullSupport `protobuf:"varint,5,opt,name=pull_support,json=pullSupport,proto3,enum=smartcore.types.PullSupport" json:"pull_support,omitempty"`
	// If Pull is emulated, how often updates can be expected.
	// This is not strict, updates may be less or more frequent.
	PullPoll *durationpb.Duration `protobuf:"bytes,6,opt,name=pull_poll,json=pullPoll,proto3" json:"pull_poll,omitempty"`
}

func (x *ResourceSupport) Reset() {
	*x = ResourceSupport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceSupport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceSupport) ProtoMessage() {}

func (x *ResourceSupport) ProtoReflect() protoreflect.Message {
	mi := &file_types_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceSupport.ProtoReflect.Descriptor instead.
func (*ResourceSupport) Descriptor() ([]byte, []int) {
	return file_types_info_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceSupport) GetReadable() bool {
	if x != nil {
		return x.Readable
	}
	return false
}

func (x *ResourceSupport) GetWritable() bool {
	if x != nil {
		return x.Writable
	}
	return false
}

func (x *ResourceSupport) GetObservable() bool {
	if x != nil {
		return x.Observable
	}
	return false
}

func (x *ResourceSupport) GetWritableFields() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.WritableFields
	}
	return nil
}

func (x *ResourceSupport) GetPullSupport() PullSupport {
	if x != nil {
		return x.PullSupport
	}
	return PullSupport_PULL_SUPPORT_UNSPECIFIED
}

func (x *ResourceSupport) GetPullPoll() *durationpb.Duration {
	if x != nil {
		return x.PullPoll
	}
	return nil
}

var File_types_info_proto protoreflect.FileDescriptor

var file_types_info_proto_rawDesc = []byte{
	0x0a, 0x10, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x02, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x61,
	0x64, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x61,
	0x64, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x72, 0x69, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x77, 0x72, 0x69, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x43, 0x0a, 0x0f, 0x77, 0x72, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0e, 0x77, 0x72, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x3f, 0x0a, 0x0c, 0x70, 0x75, 0x6c, 0x6c, 0x5f, 0x73,
	0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x73,
	0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50,
	0x75, 0x6c, 0x6c, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x0b, 0x70, 0x75, 0x6c, 0x6c,
	0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x36, 0x0a, 0x09, 0x70, 0x75, 0x6c, 0x6c, 0x5f,
	0x70, 0x6f, 0x6c, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x75, 0x6c, 0x6c, 0x50, 0x6f, 0x6c, 0x6c, 0x2a,
	0x5f, 0x0a, 0x0b, 0x50, 0x75, 0x6c, 0x6c, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1c,
	0x0a, 0x18, 0x50, 0x55, 0x4c, 0x4c, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13,
	0x50, 0x55, 0x4c, 0x4c, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x4e, 0x41, 0x54,
	0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x55, 0x4c, 0x4c, 0x5f, 0x53, 0x55,
	0x50, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x45, 0x4d, 0x55, 0x4c, 0x41, 0x54, 0x45, 0x44, 0x10, 0x02,
	0x42, 0x5e, 0x0a, 0x13, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x42, 0x09, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2d, 0x6f, 0x73, 0x2f, 0x73,
	0x63, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0xaa, 0x02,
	0x0f, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_info_proto_rawDescOnce sync.Once
	file_types_info_proto_rawDescData = file_types_info_proto_rawDesc
)

func file_types_info_proto_rawDescGZIP() []byte {
	file_types_info_proto_rawDescOnce.Do(func() {
		file_types_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_info_proto_rawDescData)
	})
	return file_types_info_proto_rawDescData
}

var file_types_info_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_types_info_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_types_info_proto_goTypes = []any{
	(PullSupport)(0),              // 0: smartcore.types.PullSupport
	(*ResourceSupport)(nil),       // 1: smartcore.types.ResourceSupport
	(*fieldmaskpb.FieldMask)(nil), // 2: google.protobuf.FieldMask
	(*durationpb.Duration)(nil),   // 3: google.protobuf.Duration
}
var file_types_info_proto_depIdxs = []int32{
	2, // 0: smartcore.types.ResourceSupport.writable_fields:type_name -> google.protobuf.FieldMask
	0, // 1: smartcore.types.ResourceSupport.pull_support:type_name -> smartcore.types.PullSupport
	3, // 2: smartcore.types.ResourceSupport.pull_poll:type_name -> google.protobuf.Duration
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_types_info_proto_init() }
func file_types_info_proto_init() {
	if File_types_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_info_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ResourceSupport); i {
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
			RawDescriptor: file_types_info_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_info_proto_goTypes,
		DependencyIndexes: file_types_info_proto_depIdxs,
		EnumInfos:         file_types_info_proto_enumTypes,
		MessageInfos:      file_types_info_proto_msgTypes,
	}.Build()
	File_types_info_proto = out.File
	file_types_info_proto_rawDesc = nil
	file_types_info_proto_goTypes = nil
	file_types_info_proto_depIdxs = nil
}
