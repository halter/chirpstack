// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: api/relay.proto

package api

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RelayListItem struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// DevEUI (EUI64).
	DevEui string `protobuf:"bytes,1,opt,name=dev_eui,json=devEui,proto3" json:"dev_eui,omitempty"`
	// Name.
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RelayListItem) Reset() {
	*x = RelayListItem{}
	mi := &file_api_relay_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RelayListItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelayListItem) ProtoMessage() {}

func (x *RelayListItem) ProtoReflect() protoreflect.Message {
	mi := &file_api_relay_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelayListItem.ProtoReflect.Descriptor instead.
func (*RelayListItem) Descriptor() ([]byte, []int) {
	return file_api_relay_proto_rawDescGZIP(), []int{0}
}

func (x *RelayListItem) GetDevEui() string {
	if x != nil {
		return x.DevEui
	}
	return ""
}

func (x *RelayListItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListRelaysRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Max number of devices to return in the result-set.
	Limit uint32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// Offset in the result-set (for pagination).
	Offset uint32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// Application ID (UUID).
	ApplicationId string `protobuf:"bytes,3,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRelaysRequest) Reset() {
	*x = ListRelaysRequest{}
	mi := &file_api_relay_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRelaysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRelaysRequest) ProtoMessage() {}

func (x *ListRelaysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_relay_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRelaysRequest.ProtoReflect.Descriptor instead.
func (*ListRelaysRequest) Descriptor() ([]byte, []int) {
	return file_api_relay_proto_rawDescGZIP(), []int{1}
}

func (x *ListRelaysRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListRelaysRequest) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListRelaysRequest) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

type ListRelaysResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Total number of relays.
	TotalCount uint32 `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	// Result-set.
	Result        []*RelayListItem `protobuf:"bytes,2,rep,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRelaysResponse) Reset() {
	*x = ListRelaysResponse{}
	mi := &file_api_relay_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRelaysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRelaysResponse) ProtoMessage() {}

func (x *ListRelaysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_relay_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRelaysResponse.ProtoReflect.Descriptor instead.
func (*ListRelaysResponse) Descriptor() ([]byte, []int) {
	return file_api_relay_proto_rawDescGZIP(), []int{2}
}

func (x *ListRelaysResponse) GetTotalCount() uint32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *ListRelaysResponse) GetResult() []*RelayListItem {
	if x != nil {
		return x.Result
	}
	return nil
}

type AddRelayDeviceRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Relay DevEUI (EUI64).
	RelayDevEui string `protobuf:"bytes,1,opt,name=relay_dev_eui,json=relayDevEui,proto3" json:"relay_dev_eui,omitempty"`
	// Device DevEUI (EUI64).
	DeviceDevEui  string `protobuf:"bytes,2,opt,name=device_dev_eui,json=deviceDevEui,proto3" json:"device_dev_eui,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddRelayDeviceRequest) Reset() {
	*x = AddRelayDeviceRequest{}
	mi := &file_api_relay_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddRelayDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRelayDeviceRequest) ProtoMessage() {}

func (x *AddRelayDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_relay_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRelayDeviceRequest.ProtoReflect.Descriptor instead.
func (*AddRelayDeviceRequest) Descriptor() ([]byte, []int) {
	return file_api_relay_proto_rawDescGZIP(), []int{3}
}

func (x *AddRelayDeviceRequest) GetRelayDevEui() string {
	if x != nil {
		return x.RelayDevEui
	}
	return ""
}

func (x *AddRelayDeviceRequest) GetDeviceDevEui() string {
	if x != nil {
		return x.DeviceDevEui
	}
	return ""
}

type RemoveRelayDeviceRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Relay DevEUI (EUI64).
	RelayDevEui string `protobuf:"bytes,1,opt,name=relay_dev_eui,json=relayDevEui,proto3" json:"relay_dev_eui,omitempty"`
	// Device DevEUI (EUI64).
	DeviceDevEui  string `protobuf:"bytes,2,opt,name=device_dev_eui,json=deviceDevEui,proto3" json:"device_dev_eui,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveRelayDeviceRequest) Reset() {
	*x = RemoveRelayDeviceRequest{}
	mi := &file_api_relay_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveRelayDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRelayDeviceRequest) ProtoMessage() {}

func (x *RemoveRelayDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_relay_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRelayDeviceRequest.ProtoReflect.Descriptor instead.
func (*RemoveRelayDeviceRequest) Descriptor() ([]byte, []int) {
	return file_api_relay_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveRelayDeviceRequest) GetRelayDevEui() string {
	if x != nil {
		return x.RelayDevEui
	}
	return ""
}

func (x *RemoveRelayDeviceRequest) GetDeviceDevEui() string {
	if x != nil {
		return x.DeviceDevEui
	}
	return ""
}

type ListRelayDevicesRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Max number of multicast groups to return in the result-set.
	Limit uint32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// Offset in the result-set (for pagination).
	Offset uint32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// Relay DevEUI (EUI64).
	RelayDevEui   string `protobuf:"bytes,3,opt,name=relay_dev_eui,json=relayDevEui,proto3" json:"relay_dev_eui,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRelayDevicesRequest) Reset() {
	*x = ListRelayDevicesRequest{}
	mi := &file_api_relay_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRelayDevicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRelayDevicesRequest) ProtoMessage() {}

func (x *ListRelayDevicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_relay_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRelayDevicesRequest.ProtoReflect.Descriptor instead.
func (*ListRelayDevicesRequest) Descriptor() ([]byte, []int) {
	return file_api_relay_proto_rawDescGZIP(), []int{5}
}

func (x *ListRelayDevicesRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListRelayDevicesRequest) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListRelayDevicesRequest) GetRelayDevEui() string {
	if x != nil {
		return x.RelayDevEui
	}
	return ""
}

type RelayDeviceListItem struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// DevEUI (EUI64).
	DevEui string `protobuf:"bytes,1,opt,name=dev_eui,json=devEui,proto3" json:"dev_eui,omitempty"`
	// Created at timestamp.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Device name.
	Name          string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RelayDeviceListItem) Reset() {
	*x = RelayDeviceListItem{}
	mi := &file_api_relay_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RelayDeviceListItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelayDeviceListItem) ProtoMessage() {}

func (x *RelayDeviceListItem) ProtoReflect() protoreflect.Message {
	mi := &file_api_relay_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelayDeviceListItem.ProtoReflect.Descriptor instead.
func (*RelayDeviceListItem) Descriptor() ([]byte, []int) {
	return file_api_relay_proto_rawDescGZIP(), []int{6}
}

func (x *RelayDeviceListItem) GetDevEui() string {
	if x != nil {
		return x.DevEui
	}
	return ""
}

func (x *RelayDeviceListItem) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *RelayDeviceListItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListRelayDevicesResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Total number of devices.
	TotalCount uint32 `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	// Result-set.
	Result        []*RelayDeviceListItem `protobuf:"bytes,2,rep,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRelayDevicesResponse) Reset() {
	*x = ListRelayDevicesResponse{}
	mi := &file_api_relay_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRelayDevicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRelayDevicesResponse) ProtoMessage() {}

func (x *ListRelayDevicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_relay_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRelayDevicesResponse.ProtoReflect.Descriptor instead.
func (*ListRelayDevicesResponse) Descriptor() ([]byte, []int) {
	return file_api_relay_proto_rawDescGZIP(), []int{7}
}

func (x *ListRelayDevicesResponse) GetTotalCount() uint32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *ListRelayDevicesResponse) GetResult() []*RelayDeviceListItem {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_api_relay_proto protoreflect.FileDescriptor

var file_api_relay_proto_rawDesc = string([]byte{
	0x0a, 0x0f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x0d, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x65, 0x76, 0x5f, 0x65, 0x75, 0x69, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76, 0x45, 0x75, 0x69, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x68, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x61, 0x0a, 0x12, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x2a, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x4c, 0x69, 0x73,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x61, 0x0a,
	0x15, 0x41, 0x64, 0x64, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x5f,
	0x64, 0x65, 0x76, 0x5f, 0x65, 0x75, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72,
	0x65, 0x6c, 0x61, 0x79, 0x44, 0x65, 0x76, 0x45, 0x75, 0x69, 0x12, 0x24, 0x0a, 0x0e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x76, 0x5f, 0x65, 0x75, 0x69, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x76, 0x45, 0x75, 0x69,
	0x22, 0x64, 0x0a, 0x18, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0d,
	0x72, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x64, 0x65, 0x76, 0x5f, 0x65, 0x75, 0x69, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x44, 0x65, 0x76, 0x45, 0x75, 0x69,
	0x12, 0x24, 0x0a, 0x0e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x76, 0x5f, 0x65,
	0x75, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x44, 0x65, 0x76, 0x45, 0x75, 0x69, 0x22, 0x6b, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x6c, 0x61, 0x79, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x22, 0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x64, 0x65, 0x76, 0x5f, 0x65, 0x75, 0x69,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x44, 0x65, 0x76,
	0x45, 0x75, 0x69, 0x22, 0x7d, 0x0a, 0x13, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x65,
	0x76, 0x5f, 0x65, 0x75, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76,
	0x45, 0x75, 0x69, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x6d, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x30, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x32, 0xc4, 0x03, 0x0a, 0x0c, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4c, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6c,
	0x61, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x73,
	0x12, 0x6f, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x3a, 0x01, 0x2a, 0x22, 0x23, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x6c, 0x61, 0x79,
	0x5f, 0x64, 0x65, 0x76, 0x5f, 0x65, 0x75, 0x69, 0x7d, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x12, 0x7c, 0x0a, 0x0c, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65,
	0x6c, 0x61, 0x79, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x35, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f,
	0x2a, 0x2d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x73, 0x2f, 0x7b, 0x72,
	0x65, 0x6c, 0x61, 0x79, 0x5f, 0x64, 0x65, 0x76, 0x5f, 0x65, 0x75, 0x69, 0x7d, 0x2f, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x64, 0x65, 0x76, 0x5f, 0x65, 0x75, 0x69, 0x7d, 0x12,
	0x77, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x1c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x25, 0x12, 0x23, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x73,
	0x2f, 0x7b, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x64, 0x65, 0x76, 0x5f, 0x65, 0x75, 0x69, 0x7d,
	0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x8c, 0x01, 0x0a, 0x11, 0x69, 0x6f, 0x2e,
	0x63, 0x68, 0x69, 0x72, 0x70, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x42, 0x0a,
	0x52, 0x65, 0x6c, 0x61, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x6c, 0x74, 0x65, 0x72, 0x2f,
	0x63, 0x68, 0x69, 0x72, 0x70, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67,
	0x6f, 0x2f, 0x76, 0x34, 0x2f, 0x61, 0x70, 0x69, 0xaa, 0x02, 0x0e, 0x43, 0x68, 0x69, 0x72, 0x70,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x41, 0x70, 0x69, 0xca, 0x02, 0x0e, 0x43, 0x68, 0x69, 0x72,
	0x70, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x41, 0x70, 0x69, 0xe2, 0x02, 0x1a, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5c, 0x43, 0x68, 0x69, 0x72, 0x70, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x5c, 0x41, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_relay_proto_rawDescOnce sync.Once
	file_api_relay_proto_rawDescData []byte
)

func file_api_relay_proto_rawDescGZIP() []byte {
	file_api_relay_proto_rawDescOnce.Do(func() {
		file_api_relay_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_relay_proto_rawDesc), len(file_api_relay_proto_rawDesc)))
	})
	return file_api_relay_proto_rawDescData
}

var file_api_relay_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_relay_proto_goTypes = []any{
	(*RelayListItem)(nil),            // 0: api.RelayListItem
	(*ListRelaysRequest)(nil),        // 1: api.ListRelaysRequest
	(*ListRelaysResponse)(nil),       // 2: api.ListRelaysResponse
	(*AddRelayDeviceRequest)(nil),    // 3: api.AddRelayDeviceRequest
	(*RemoveRelayDeviceRequest)(nil), // 4: api.RemoveRelayDeviceRequest
	(*ListRelayDevicesRequest)(nil),  // 5: api.ListRelayDevicesRequest
	(*RelayDeviceListItem)(nil),      // 6: api.RelayDeviceListItem
	(*ListRelayDevicesResponse)(nil), // 7: api.ListRelayDevicesResponse
	(*timestamppb.Timestamp)(nil),    // 8: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),            // 9: google.protobuf.Empty
}
var file_api_relay_proto_depIdxs = []int32{
	0, // 0: api.ListRelaysResponse.result:type_name -> api.RelayListItem
	8, // 1: api.RelayDeviceListItem.created_at:type_name -> google.protobuf.Timestamp
	6, // 2: api.ListRelayDevicesResponse.result:type_name -> api.RelayDeviceListItem
	1, // 3: api.RelayService.List:input_type -> api.ListRelaysRequest
	3, // 4: api.RelayService.AddDevice:input_type -> api.AddRelayDeviceRequest
	4, // 5: api.RelayService.RemoveDevice:input_type -> api.RemoveRelayDeviceRequest
	5, // 6: api.RelayService.ListDevices:input_type -> api.ListRelayDevicesRequest
	2, // 7: api.RelayService.List:output_type -> api.ListRelaysResponse
	9, // 8: api.RelayService.AddDevice:output_type -> google.protobuf.Empty
	9, // 9: api.RelayService.RemoveDevice:output_type -> google.protobuf.Empty
	7, // 10: api.RelayService.ListDevices:output_type -> api.ListRelayDevicesResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_relay_proto_init() }
func file_api_relay_proto_init() {
	if File_api_relay_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_relay_proto_rawDesc), len(file_api_relay_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_relay_proto_goTypes,
		DependencyIndexes: file_api_relay_proto_depIdxs,
		MessageInfos:      file_api_relay_proto_msgTypes,
	}.Build()
	File_api_relay_proto = out.File
	file_api_relay_proto_goTypes = nil
	file_api_relay_proto_depIdxs = nil
}
