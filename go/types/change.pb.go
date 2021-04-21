// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: types/change.proto

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

// Type of change
type ChangeType int32

const (
	// Nothing has changed but existing items have been requested. The new_value property should be set
	ChangeType_NONE ChangeType = 0
	// An item has been added. The new_value property of the change should be set
	ChangeType_ADD ChangeType = 1
	// An item has been updated. The new_value and old_value properties of the change should be set
	ChangeType_UPDATE ChangeType = 2
	// A item has been removed. The old_value property of the change should be set
	ChangeType_REMOVE ChangeType = 3
	// All existing items should be replaced by this set of items.
	ChangeType_REPLACE ChangeType = 4
)

// Enum value maps for ChangeType.
var (
	ChangeType_name = map[int32]string{
		0: "NONE",
		1: "ADD",
		2: "UPDATE",
		3: "REMOVE",
		4: "REPLACE",
	}
	ChangeType_value = map[string]int32{
		"NONE":    0,
		"ADD":     1,
		"UPDATE":  2,
		"REMOVE":  3,
		"REPLACE": 4,
	}
)

func (x ChangeType) Enum() *ChangeType {
	p := new(ChangeType)
	*p = x
	return p
}

func (x ChangeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChangeType) Descriptor() protoreflect.EnumDescriptor {
	return file_types_change_proto_enumTypes[0].Descriptor()
}

func (ChangeType) Type() protoreflect.EnumType {
	return &file_types_change_proto_enumTypes[0]
}

func (x ChangeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChangeType.Descriptor instead.
func (ChangeType) EnumDescriptor() ([]byte, []int) {
	return file_types_change_proto_rawDescGZIP(), []int{0}
}

var File_types_change_proto protoreflect.FileDescriptor

var file_types_change_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2a, 0x44, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x07, 0x0a,
	0x03, 0x41, 0x44, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45,
	0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x10, 0x03, 0x12, 0x0b,
	0x0a, 0x07, 0x52, 0x45, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x10, 0x04, 0x42, 0x61, 0x0a, 0x13, 0x64,
	0x65, 0x76, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x42, 0x0b, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x2e, 0x76, 0x61, 0x6e, 0x74, 0x69, 0x2e, 0x63, 0x6f, 0x2e,
	0x75, 0x6b, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x63, 0x2d,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0xaa, 0x02, 0x0f, 0x53,
	0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_change_proto_rawDescOnce sync.Once
	file_types_change_proto_rawDescData = file_types_change_proto_rawDesc
)

func file_types_change_proto_rawDescGZIP() []byte {
	file_types_change_proto_rawDescOnce.Do(func() {
		file_types_change_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_change_proto_rawDescData)
	})
	return file_types_change_proto_rawDescData
}

var file_types_change_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_types_change_proto_goTypes = []interface{}{
	(ChangeType)(0), // 0: smartcore.types.ChangeType
}
var file_types_change_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_types_change_proto_init() }
func file_types_change_proto_init() {
	if File_types_change_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_types_change_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_change_proto_goTypes,
		DependencyIndexes: file_types_change_proto_depIdxs,
		EnumInfos:         file_types_change_proto_enumTypes,
	}.Build()
	File_types_change_proto = out.File
	file_types_change_proto_rawDesc = nil
	file_types_change_proto_goTypes = nil
	file_types_change_proto_depIdxs = nil
}
