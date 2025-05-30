// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.29.3
// source: info/info.proto

package info

import (
	types "github.com/smart-core-os/sc-api/go/types"
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

type ListDevicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Only return devices this deep in the device -> device -> device tree. Default is unlimited depth, 1 means only
	// return devices owned by the enclosing service.
	Depth int32 `protobuf:"varint,1,opt,name=depth,proto3" json:"depth,omitempty"`
	// How many items should be returned. Default is 10, max is 100.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The next_page_token from the previous response
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListDevicesRequest) Reset() {
	*x = ListDevicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDevicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDevicesRequest) ProtoMessage() {}

func (x *ListDevicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_info_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDevicesRequest.ProtoReflect.Descriptor instead.
func (*ListDevicesRequest) Descriptor() ([]byte, []int) {
	return file_info_info_proto_rawDescGZIP(), []int{0}
}

func (x *ListDevicesRequest) GetDepth() int32 {
	if x != nil {
		return x.Depth
	}
	return 0
}

func (x *ListDevicesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListDevicesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListDevicesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The current page of devices to satisfy the request
	Devices []*Device `protobuf:"bytes,1,rep,name=devices,proto3" json:"devices,omitempty"`
	// Use in page_token to retrieve the next page
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// Total number of devices that would be returned if not splitting the response into pages.
	TotalSize int32 `protobuf:"varint,3,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
}

func (x *ListDevicesResponse) Reset() {
	*x = ListDevicesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDevicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDevicesResponse) ProtoMessage() {}

func (x *ListDevicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_info_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDevicesResponse.ProtoReflect.Descriptor instead.
func (*ListDevicesResponse) Descriptor() ([]byte, []int) {
	return file_info_info_proto_rawDescGZIP(), []int{1}
}

func (x *ListDevicesResponse) GetDevices() []*Device {
	if x != nil {
		return x.Devices
	}
	return nil
}

func (x *ListDevicesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListDevicesResponse) GetTotalSize() int32 {
	if x != nil {
		return x.TotalSize
	}
	return 0
}

type PullDevicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Only return devices this deep in the device -> device -> device tree. Default is unlimited depth, 1 means only
	// return devices owned by the enclosing service.
	Depth int32 `protobuf:"varint,1,opt,name=depth,proto3" json:"depth,omitempty"`
	// if true, all existing devices will be sent down the stream before any updates.
	Sync bool `protobuf:"varint,2,opt,name=sync,proto3" json:"sync,omitempty"`
}

func (x *PullDevicesRequest) Reset() {
	*x = PullDevicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_info_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullDevicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullDevicesRequest) ProtoMessage() {}

func (x *PullDevicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_info_info_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullDevicesRequest.ProtoReflect.Descriptor instead.
func (*PullDevicesRequest) Descriptor() ([]byte, []int) {
	return file_info_info_proto_rawDescGZIP(), []int{2}
}

func (x *PullDevicesRequest) GetDepth() int32 {
	if x != nil {
		return x.Depth
	}
	return 0
}

func (x *PullDevicesRequest) GetSync() bool {
	if x != nil {
		return x.Sync
	}
	return false
}

type PullDevicesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Changes []*PullDevicesResponse_Change `protobuf:"bytes,1,rep,name=changes,proto3" json:"changes,omitempty"`
}

func (x *PullDevicesResponse) Reset() {
	*x = PullDevicesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_info_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullDevicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullDevicesResponse) ProtoMessage() {}

func (x *PullDevicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_info_info_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullDevicesResponse.ProtoReflect.Descriptor instead.
func (*PullDevicesResponse) Descriptor() ([]byte, []int) {
	return file_info_info_proto_rawDescGZIP(), []int{3}
}

func (x *PullDevicesResponse) GetChanges() []*PullDevicesResponse_Change {
	if x != nil {
		return x.Changes
	}
	return nil
}

// Describes basic information about a device. Device hierarchy is loosely defined via the owner property
type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the relative resource name for the device
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// List of traits the device has.
	Traits []*Trait `protobuf:"bytes,2,rep,name=traits,proto3" json:"traits,omitempty"`
	// Owner of the device. This is the the device, supervisor, or area controller that is directly responsible for
	// communication with the device, i.e. service requests should go through the owner.
	Owner *Device `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	// Information used to connect to the device directly via gRPC. This must not be present if we are supposed to connect
	// via the owner device. This may be present if direct communication is possible, even if owner is also set. Prefer
	// Connecting via owner if possible.
	Client *GrpcClientOptions `protobuf:"bytes,4,opt,name=client,proto3" json:"client,omitempty"`
	// Official name for the device
	Title string `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	// Readable name for the device
	DisplayName string `protobuf:"bytes,6,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// One or more paragraphs of text describing the device
	Description string `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	// arbitrary flags that can be applied to the device to encode additional information
	Labels map[string]string `protobuf:"bytes,8,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_info_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_info_info_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_info_info_proto_rawDescGZIP(), []int{4}
}

func (x *Device) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Device) GetTraits() []*Trait {
	if x != nil {
		return x.Traits
	}
	return nil
}

func (x *Device) GetOwner() *Device {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *Device) GetClient() *GrpcClientOptions {
	if x != nil {
		return x.Client
	}
	return nil
}

func (x *Device) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Device) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Device) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Device) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

// describes a trait for the api. See the devices/traits apis for details of these.
type Trait struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the trait service that is implemented. In the format package.Service excluding Api or other type suffix.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Trait) Reset() {
	*x = Trait{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_info_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trait) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trait) ProtoMessage() {}

func (x *Trait) ProtoReflect() protoreflect.Message {
	mi := &file_info_info_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trait.ProtoReflect.Descriptor instead.
func (*Trait) Descriptor() ([]byte, []int) {
	return file_info_info_proto_rawDescGZIP(), []int{5}
}

func (x *Trait) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// describes options relating to client connections
type GrpcClientOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The host:port for the server
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// PEM encoded client certificate for use with gRPC transport security
	ClientCert []byte `protobuf:"bytes,2,opt,name=clientCert,proto3" json:"clientCert,omitempty"`
	// PEM encoded client private key for use with gRPC transport security
	ClientKey []byte `protobuf:"bytes,3,opt,name=clientKey,proto3" json:"clientKey,omitempty"`
	// PEM encoded certificate chan representing the certificate authority for use with gRPC transport security
	ClientCa []byte `protobuf:"bytes,4,opt,name=clientCa,proto3" json:"clientCa,omitempty"`
}

func (x *GrpcClientOptions) Reset() {
	*x = GrpcClientOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_info_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcClientOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcClientOptions) ProtoMessage() {}

func (x *GrpcClientOptions) ProtoReflect() protoreflect.Message {
	mi := &file_info_info_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcClientOptions.ProtoReflect.Descriptor instead.
func (*GrpcClientOptions) Descriptor() ([]byte, []int) {
	return file_info_info_proto_rawDescGZIP(), []int{6}
}

func (x *GrpcClientOptions) GetAuthority() string {
	if x != nil {
		return x.Authority
	}
	return ""
}

func (x *GrpcClientOptions) GetClientCert() []byte {
	if x != nil {
		return x.ClientCert
	}
	return nil
}

func (x *GrpcClientOptions) GetClientKey() []byte {
	if x != nil {
		return x.ClientKey
	}
	return nil
}

func (x *GrpcClientOptions) GetClientCa() []byte {
	if x != nil {
		return x.ClientCa
	}
	return nil
}

type PullDevicesResponse_Change struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     types.ChangeType `protobuf:"varint,1,opt,name=type,proto3,enum=smartcore.types.ChangeType" json:"type,omitempty"`
	NewValue *Device          `protobuf:"bytes,2,opt,name=new_value,json=newValue,proto3" json:"new_value,omitempty"`
	OldValue *Device          `protobuf:"bytes,3,opt,name=old_value,json=oldValue,proto3" json:"old_value,omitempty"`
}

func (x *PullDevicesResponse_Change) Reset() {
	*x = PullDevicesResponse_Change{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_info_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullDevicesResponse_Change) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullDevicesResponse_Change) ProtoMessage() {}

func (x *PullDevicesResponse_Change) ProtoReflect() protoreflect.Message {
	mi := &file_info_info_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullDevicesResponse_Change.ProtoReflect.Descriptor instead.
func (*PullDevicesResponse_Change) Descriptor() ([]byte, []int) {
	return file_info_info_proto_rawDescGZIP(), []int{3, 0}
}

func (x *PullDevicesResponse_Change) GetType() types.ChangeType {
	if x != nil {
		return x.Type
	}
	return types.ChangeType(0)
}

func (x *PullDevicesResponse_Change) GetNewValue() *Device {
	if x != nil {
		return x.NewValue
	}
	return nil
}

func (x *PullDevicesResponse_Change) GetOldValue() *Device {
	if x != nil {
		return x.OldValue
	}
	return nil
}

var File_info_info_proto protoreflect.FileDescriptor

var file_info_info_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x69, 0x6e, 0x66,
	0x6f, 0x1a, 0x12, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x64,
	0x65, 0x70, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x64, 0x65, 0x70, 0x74,
	0x68, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x8e, 0x01,
	0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x3e,
	0x0a, 0x12, 0x50, 0x75, 0x6c, 0x6c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x70, 0x74, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x64, 0x65, 0x70, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x79,
	0x6e, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x73, 0x79, 0x6e, 0x63, 0x22, 0x81,
	0x02, 0x0a, 0x13, 0x50, 0x75, 0x6c, 0x6c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x1a, 0xa3, 0x01, 0x0a,
	0x06, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x33, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x6d,
	0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x33, 0x0a,
	0x09, 0x6f, 0x6c, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x69, 0x6e, 0x66,
	0x6f, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x08, 0x6f, 0x6c, 0x64, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x86, 0x03, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x2d, 0x0a, 0x06, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x69, 0x6e,
	0x66, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x74, 0x52, 0x06, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73,
	0x12, 0x2c, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x6f,
	0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x39,
	0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e,
	0x47, 0x72, 0x70, 0x63, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x1b, 0x0a, 0x05, 0x54,
	0x72, 0x61, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x8b, 0x01, 0x0a, 0x11, 0x47, 0x72, 0x70,
	0x63, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x32, 0xb8, 0x01, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x56, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x22,
	0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x69,
	0x6e, 0x66, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x0b, 0x50, 0x75, 0x6c, 0x6c, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x22, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x6d, 0x61,
	0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x50, 0x75, 0x6c, 0x6c,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30,
	0x01, 0x42, 0x6c, 0x0a, 0x12, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x42, 0x09, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2d, 0x6f, 0x73, 0x2f, 0x73,
	0x63, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0xaa, 0x02, 0x0e,
	0x53, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0xca, 0x02,
	0x0e, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x5c, 0x49, 0x6e, 0x66, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_info_info_proto_rawDescOnce sync.Once
	file_info_info_proto_rawDescData = file_info_info_proto_rawDesc
)

func file_info_info_proto_rawDescGZIP() []byte {
	file_info_info_proto_rawDescOnce.Do(func() {
		file_info_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_info_info_proto_rawDescData)
	})
	return file_info_info_proto_rawDescData
}

var file_info_info_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_info_info_proto_goTypes = []any{
	(*ListDevicesRequest)(nil),         // 0: smartcore.info.ListDevicesRequest
	(*ListDevicesResponse)(nil),        // 1: smartcore.info.ListDevicesResponse
	(*PullDevicesRequest)(nil),         // 2: smartcore.info.PullDevicesRequest
	(*PullDevicesResponse)(nil),        // 3: smartcore.info.PullDevicesResponse
	(*Device)(nil),                     // 4: smartcore.info.Device
	(*Trait)(nil),                      // 5: smartcore.info.Trait
	(*GrpcClientOptions)(nil),          // 6: smartcore.info.GrpcClientOptions
	(*PullDevicesResponse_Change)(nil), // 7: smartcore.info.PullDevicesResponse.Change
	nil,                                // 8: smartcore.info.Device.LabelsEntry
	(types.ChangeType)(0),              // 9: smartcore.types.ChangeType
}
var file_info_info_proto_depIdxs = []int32{
	4,  // 0: smartcore.info.ListDevicesResponse.devices:type_name -> smartcore.info.Device
	7,  // 1: smartcore.info.PullDevicesResponse.changes:type_name -> smartcore.info.PullDevicesResponse.Change
	5,  // 2: smartcore.info.Device.traits:type_name -> smartcore.info.Trait
	4,  // 3: smartcore.info.Device.owner:type_name -> smartcore.info.Device
	6,  // 4: smartcore.info.Device.client:type_name -> smartcore.info.GrpcClientOptions
	8,  // 5: smartcore.info.Device.labels:type_name -> smartcore.info.Device.LabelsEntry
	9,  // 6: smartcore.info.PullDevicesResponse.Change.type:type_name -> smartcore.types.ChangeType
	4,  // 7: smartcore.info.PullDevicesResponse.Change.new_value:type_name -> smartcore.info.Device
	4,  // 8: smartcore.info.PullDevicesResponse.Change.old_value:type_name -> smartcore.info.Device
	0,  // 9: smartcore.info.Info.ListDevices:input_type -> smartcore.info.ListDevicesRequest
	2,  // 10: smartcore.info.Info.PullDevices:input_type -> smartcore.info.PullDevicesRequest
	1,  // 11: smartcore.info.Info.ListDevices:output_type -> smartcore.info.ListDevicesResponse
	3,  // 12: smartcore.info.Info.PullDevices:output_type -> smartcore.info.PullDevicesResponse
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_info_info_proto_init() }
func file_info_info_proto_init() {
	if File_info_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_info_info_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ListDevicesRequest); i {
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
		file_info_info_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListDevicesResponse); i {
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
		file_info_info_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*PullDevicesRequest); i {
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
		file_info_info_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*PullDevicesResponse); i {
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
		file_info_info_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Device); i {
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
		file_info_info_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Trait); i {
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
		file_info_info_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GrpcClientOptions); i {
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
		file_info_info_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*PullDevicesResponse_Change); i {
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
			RawDescriptor: file_info_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_info_info_proto_goTypes,
		DependencyIndexes: file_info_info_proto_depIdxs,
		MessageInfos:      file_info_info_proto_msgTypes,
	}.Build()
	File_info_info_proto = out.File
	file_info_info_proto_rawDesc = nil
	file_info_info_proto_goTypes = nil
	file_info_info_proto_depIdxs = nil
}
