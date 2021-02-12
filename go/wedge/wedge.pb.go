// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: wedge.proto

package wedge

import (
	context "context"
	slx "github.com/Wappsto/wedge-api/go/slx"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Symbols defined in public import of slx.proto.

type NumberMapping = slx.NumberMapping

const NumberMapping_off = slx.NumberMapping_off
const NumberMapping_on = slx.NumberMapping_on

var NumberMapping_name = slx.NumberMapping_name
var NumberMapping_value = slx.NumberMapping_value

type Meta = slx.Meta
type State = slx.State
type Number = slx.Number
type String = slx.String
type Blob = slx.Blob
type Info = slx.Info
type Value = slx.Value
type Device = slx.Device

// ------ Report messages ------------
type NodeIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *NodeIdentity) Reset() {
	*x = NodeIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wedge_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeIdentity) ProtoMessage() {}

func (x *NodeIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_wedge_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeIdentity.ProtoReflect.Descriptor instead.
func (*NodeIdentity) Descriptor() ([]byte, []int) {
	return file_wedge_proto_rawDescGZIP(), []int{0}
}

func (x *NodeIdentity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Model struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node   *NodeIdentity `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Device []*slx.Device `protobuf:"bytes,2,rep,name=device,proto3" json:"device,omitempty"`
}

func (x *Model) Reset() {
	*x = Model{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wedge_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Model) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Model) ProtoMessage() {}

func (x *Model) ProtoReflect() protoreflect.Message {
	mi := &file_wedge_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Model.ProtoReflect.Descriptor instead.
func (*Model) Descriptor() ([]byte, []int) {
	return file_wedge_proto_rawDescGZIP(), []int{1}
}

func (x *Model) GetNode() *NodeIdentity {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *Model) GetDevice() []*slx.Device {
	if x != nil {
		return x.Device
	}
	return nil
}

// Create/Update all node data model
type SetModelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Model *Model `protobuf:"bytes,1,opt,name=model,proto3" json:"model,omitempty"`
}

func (x *SetModelRequest) Reset() {
	*x = SetModelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wedge_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetModelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetModelRequest) ProtoMessage() {}

func (x *SetModelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wedge_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetModelRequest.ProtoReflect.Descriptor instead.
func (*SetModelRequest) Descriptor() ([]byte, []int) {
	return file_wedge_proto_rawDescGZIP(), []int{2}
}

func (x *SetModelRequest) GetModel() *Model {
	if x != nil {
		return x.Model
	}
	return nil
}

type SetDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node   *NodeIdentity `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Device *slx.Device   `protobuf:"bytes,2,opt,name=device,proto3" json:"device,omitempty"`
}

func (x *SetDeviceRequest) Reset() {
	*x = SetDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wedge_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDeviceRequest) ProtoMessage() {}

func (x *SetDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wedge_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetDeviceRequest.ProtoReflect.Descriptor instead.
func (*SetDeviceRequest) Descriptor() ([]byte, []int) {
	return file_wedge_proto_rawDescGZIP(), []int{3}
}

func (x *SetDeviceRequest) GetNode() *NodeIdentity {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *SetDeviceRequest) GetDevice() *slx.Device {
	if x != nil {
		return x.Device
	}
	return nil
}

type SetValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node     *NodeIdentity `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	DeviceId uint32        `protobuf:"varint,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Value    *slx.Value    `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SetValueRequest) Reset() {
	*x = SetValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wedge_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetValueRequest) ProtoMessage() {}

func (x *SetValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wedge_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetValueRequest.ProtoReflect.Descriptor instead.
func (*SetValueRequest) Descriptor() ([]byte, []int) {
	return file_wedge_proto_rawDescGZIP(), []int{4}
}

func (x *SetValueRequest) GetNode() *NodeIdentity {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *SetValueRequest) GetDeviceId() uint32 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

func (x *SetValueRequest) GetValue() *slx.Value {
	if x != nil {
		return x.Value
	}
	return nil
}

type SetStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node     *NodeIdentity `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	DeviceId uint32        `protobuf:"varint,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	ValueId  uint32        `protobuf:"varint,3,opt,name=value_id,json=valueId,proto3" json:"value_id,omitempty"`
	State    *slx.State    `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *SetStateRequest) Reset() {
	*x = SetStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wedge_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetStateRequest) ProtoMessage() {}

func (x *SetStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wedge_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetStateRequest.ProtoReflect.Descriptor instead.
func (*SetStateRequest) Descriptor() ([]byte, []int) {
	return file_wedge_proto_rawDescGZIP(), []int{5}
}

func (x *SetStateRequest) GetNode() *NodeIdentity {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *SetStateRequest) GetDeviceId() uint32 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

func (x *SetStateRequest) GetValueId() uint32 {
	if x != nil {
		return x.ValueId
	}
	return 0
}

func (x *SetStateRequest) GetState() *slx.State {
	if x != nil {
		return x.State
	}
	return nil
}

type GetControlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node *NodeIdentity `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
}

func (x *GetControlRequest) Reset() {
	*x = GetControlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wedge_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetControlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetControlRequest) ProtoMessage() {}

func (x *GetControlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wedge_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetControlRequest.ProtoReflect.Descriptor instead.
func (*GetControlRequest) Descriptor() ([]byte, []int) {
	return file_wedge_proto_rawDescGZIP(), []int{6}
}

func (x *GetControlRequest) GetNode() *NodeIdentity {
	if x != nil {
		return x.Node
	}
	return nil
}

// ----- Control messages ----------
type UpdateState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State    *slx.State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	DeviceId uint32     `protobuf:"varint,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	ValueId  uint32     `protobuf:"varint,3,opt,name=value_id,json=valueId,proto3" json:"value_id,omitempty"`
}

func (x *UpdateState) Reset() {
	*x = UpdateState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wedge_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateState) ProtoMessage() {}

func (x *UpdateState) ProtoReflect() protoreflect.Message {
	mi := &file_wedge_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateState.ProtoReflect.Descriptor instead.
func (*UpdateState) Descriptor() ([]byte, []int) {
	return file_wedge_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateState) GetState() *slx.State {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *UpdateState) GetDeviceId() uint32 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

func (x *UpdateState) GetValueId() uint32 {
	if x != nil {
		return x.ValueId
	}
	return 0
}

type DeleteDevice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId uint32 `protobuf:"varint,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *DeleteDevice) Reset() {
	*x = DeleteDevice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wedge_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDevice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDevice) ProtoMessage() {}

func (x *DeleteDevice) ProtoReflect() protoreflect.Message {
	mi := &file_wedge_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDevice.ProtoReflect.Descriptor instead.
func (*DeleteDevice) Descriptor() ([]byte, []int) {
	return file_wedge_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteDevice) GetDeviceId() uint32 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

type Control struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Update *UpdateState  `protobuf:"bytes,1,opt,name=update,proto3" json:"update,omitempty"`
	Delete *DeleteDevice `protobuf:"bytes,2,opt,name=delete,proto3" json:"delete,omitempty"`
}

func (x *Control) Reset() {
	*x = Control{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wedge_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Control) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Control) ProtoMessage() {}

func (x *Control) ProtoReflect() protoreflect.Message {
	mi := &file_wedge_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Control.ProtoReflect.Descriptor instead.
func (*Control) Descriptor() ([]byte, []int) {
	return file_wedge_proto_rawDescGZIP(), []int{9}
}

func (x *Control) GetUpdate() *UpdateState {
	if x != nil {
		return x.Update
	}
	return nil
}

func (x *Control) GetDelete() *DeleteDevice {
	if x != nil {
		return x.Delete
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code    uint32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wedge_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_wedge_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_wedge_proto_rawDescGZIP(), []int{10}
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Error) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type Replay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok    bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Error *Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Replay) Reset() {
	*x = Replay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wedge_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Replay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Replay) ProtoMessage() {}

func (x *Replay) ProtoReflect() protoreflect.Message {
	mi := &file_wedge_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Replay.ProtoReflect.Descriptor instead.
func (*Replay) Descriptor() ([]byte, []int) {
	return file_wedge_proto_rawDescGZIP(), []int{11}
}

func (x *Replay) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *Replay) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_wedge_proto protoreflect.FileDescriptor

var file_wedge_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x77, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x77,
	0x65, 0x64, 0x67, 0x65, 0x1a, 0x09, 0x73, 0x6c, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x1e, 0x0a, 0x0c, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x51, 0x0a, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x27, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x04, 0x6e, 0x6f, 0x64,
	0x65, 0x12, 0x1f, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x22, 0x35, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x77, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x5c, 0x0a, 0x10, 0x53, 0x65, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a,
	0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77, 0x65,
	0x64, 0x67, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0x75, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x6e, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77, 0x65, 0x64, 0x67, 0x65,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x04, 0x6e,
	0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x06, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x90,
	0x01, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x77, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x06, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x22, 0x3c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x22,
	0x63, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1c,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x49, 0x64, 0x22, 0x2b, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x64, 0x22, 0x62, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x2a, 0x0a, 0x06,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x77,
	0x65, 0x64, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77, 0x65, 0x64, 0x67, 0x65,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x06, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x22, 0x35, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x3c, 0x0a, 0x06,
	0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x22, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x77, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x97, 0x02, 0x0a, 0x05, 0x57,
	0x65, 0x64, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x12, 0x16, 0x2e, 0x77, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x77, 0x65, 0x64, 0x67, 0x65,
	0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x09, 0x53, 0x65, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x17, 0x2e, 0x77, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x53,
	0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0d, 0x2e, 0x77, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x22, 0x00,
	0x12, 0x33, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x2e, 0x77,
	0x65, 0x64, 0x67, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x77, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x70,
	0x6c, 0x61, 0x79, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x16, 0x2e, 0x77, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x77, 0x65, 0x64, 0x67,
	0x65, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x18, 0x2e, 0x77, 0x65, 0x64, 0x67, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x77, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x22, 0x00, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x57, 0x61, 0x70, 0x70, 0x73, 0x74, 0x6f, 0x2f, 0x77, 0x65, 0x64, 0x67, 0x65,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x77, 0x65, 0x64, 0x67, 0x65, 0x50, 0x00, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wedge_proto_rawDescOnce sync.Once
	file_wedge_proto_rawDescData = file_wedge_proto_rawDesc
)

func file_wedge_proto_rawDescGZIP() []byte {
	file_wedge_proto_rawDescOnce.Do(func() {
		file_wedge_proto_rawDescData = protoimpl.X.CompressGZIP(file_wedge_proto_rawDescData)
	})
	return file_wedge_proto_rawDescData
}

var file_wedge_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_wedge_proto_goTypes = []interface{}{
	(*NodeIdentity)(nil),      // 0: wedge.NodeIdentity
	(*Model)(nil),             // 1: wedge.Model
	(*SetModelRequest)(nil),   // 2: wedge.SetModelRequest
	(*SetDeviceRequest)(nil),  // 3: wedge.SetDeviceRequest
	(*SetValueRequest)(nil),   // 4: wedge.SetValueRequest
	(*SetStateRequest)(nil),   // 5: wedge.SetStateRequest
	(*GetControlRequest)(nil), // 6: wedge.GetControlRequest
	(*UpdateState)(nil),       // 7: wedge.UpdateState
	(*DeleteDevice)(nil),      // 8: wedge.DeleteDevice
	(*Control)(nil),           // 9: wedge.Control
	(*Error)(nil),             // 10: wedge.Error
	(*Replay)(nil),            // 11: wedge.Replay
	(*slx.Device)(nil),        // 12: Device
	(*slx.Value)(nil),         // 13: Value
	(*slx.State)(nil),         // 14: State
}
var file_wedge_proto_depIdxs = []int32{
	0,  // 0: wedge.Model.node:type_name -> wedge.NodeIdentity
	12, // 1: wedge.Model.device:type_name -> Device
	1,  // 2: wedge.SetModelRequest.model:type_name -> wedge.Model
	0,  // 3: wedge.SetDeviceRequest.node:type_name -> wedge.NodeIdentity
	12, // 4: wedge.SetDeviceRequest.device:type_name -> Device
	0,  // 5: wedge.SetValueRequest.node:type_name -> wedge.NodeIdentity
	13, // 6: wedge.SetValueRequest.value:type_name -> Value
	0,  // 7: wedge.SetStateRequest.node:type_name -> wedge.NodeIdentity
	14, // 8: wedge.SetStateRequest.state:type_name -> State
	0,  // 9: wedge.GetControlRequest.node:type_name -> wedge.NodeIdentity
	14, // 10: wedge.UpdateState.state:type_name -> State
	7,  // 11: wedge.Control.update:type_name -> wedge.UpdateState
	8,  // 12: wedge.Control.delete:type_name -> wedge.DeleteDevice
	10, // 13: wedge.Replay.error:type_name -> wedge.Error
	2,  // 14: wedge.Wedge.SetModel:input_type -> wedge.SetModelRequest
	3,  // 15: wedge.Wedge.SetDevice:input_type -> wedge.SetDeviceRequest
	4,  // 16: wedge.Wedge.SetValue:input_type -> wedge.SetValueRequest
	5,  // 17: wedge.Wedge.SetState:input_type -> wedge.SetStateRequest
	6,  // 18: wedge.Wedge.GetControl:input_type -> wedge.GetControlRequest
	11, // 19: wedge.Wedge.SetModel:output_type -> wedge.Replay
	11, // 20: wedge.Wedge.SetDevice:output_type -> wedge.Replay
	11, // 21: wedge.Wedge.SetValue:output_type -> wedge.Replay
	11, // 22: wedge.Wedge.SetState:output_type -> wedge.Replay
	9,  // 23: wedge.Wedge.GetControl:output_type -> wedge.Control
	19, // [19:24] is the sub-list for method output_type
	14, // [14:19] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_wedge_proto_init() }
func file_wedge_proto_init() {
	if File_wedge_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wedge_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeIdentity); i {
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
		file_wedge_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Model); i {
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
		file_wedge_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetModelRequest); i {
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
		file_wedge_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetDeviceRequest); i {
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
		file_wedge_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetValueRequest); i {
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
		file_wedge_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetStateRequest); i {
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
		file_wedge_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetControlRequest); i {
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
		file_wedge_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateState); i {
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
		file_wedge_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDevice); i {
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
		file_wedge_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Control); i {
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
		file_wedge_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_wedge_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Replay); i {
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
			RawDescriptor: file_wedge_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wedge_proto_goTypes,
		DependencyIndexes: file_wedge_proto_depIdxs,
		MessageInfos:      file_wedge_proto_msgTypes,
	}.Build()
	File_wedge_proto = out.File
	file_wedge_proto_rawDesc = nil
	file_wedge_proto_goTypes = nil
	file_wedge_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WedgeClient is the client API for Wedge service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WedgeClient interface {
	SetModel(ctx context.Context, in *SetModelRequest, opts ...grpc.CallOption) (*Replay, error)
	SetDevice(ctx context.Context, in *SetDeviceRequest, opts ...grpc.CallOption) (*Replay, error)
	SetValue(ctx context.Context, in *SetValueRequest, opts ...grpc.CallOption) (*Replay, error)
	SetState(ctx context.Context, in *SetStateRequest, opts ...grpc.CallOption) (*Replay, error)
	GetControl(ctx context.Context, in *GetControlRequest, opts ...grpc.CallOption) (*Control, error)
}

type wedgeClient struct {
	cc grpc.ClientConnInterface
}

func NewWedgeClient(cc grpc.ClientConnInterface) WedgeClient {
	return &wedgeClient{cc}
}

func (c *wedgeClient) SetModel(ctx context.Context, in *SetModelRequest, opts ...grpc.CallOption) (*Replay, error) {
	out := new(Replay)
	err := c.cc.Invoke(ctx, "/wedge.Wedge/SetModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wedgeClient) SetDevice(ctx context.Context, in *SetDeviceRequest, opts ...grpc.CallOption) (*Replay, error) {
	out := new(Replay)
	err := c.cc.Invoke(ctx, "/wedge.Wedge/SetDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wedgeClient) SetValue(ctx context.Context, in *SetValueRequest, opts ...grpc.CallOption) (*Replay, error) {
	out := new(Replay)
	err := c.cc.Invoke(ctx, "/wedge.Wedge/SetValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wedgeClient) SetState(ctx context.Context, in *SetStateRequest, opts ...grpc.CallOption) (*Replay, error) {
	out := new(Replay)
	err := c.cc.Invoke(ctx, "/wedge.Wedge/SetState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wedgeClient) GetControl(ctx context.Context, in *GetControlRequest, opts ...grpc.CallOption) (*Control, error) {
	out := new(Control)
	err := c.cc.Invoke(ctx, "/wedge.Wedge/GetControl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WedgeServer is the server API for Wedge service.
type WedgeServer interface {
	SetModel(context.Context, *SetModelRequest) (*Replay, error)
	SetDevice(context.Context, *SetDeviceRequest) (*Replay, error)
	SetValue(context.Context, *SetValueRequest) (*Replay, error)
	SetState(context.Context, *SetStateRequest) (*Replay, error)
	GetControl(context.Context, *GetControlRequest) (*Control, error)
}

// UnimplementedWedgeServer can be embedded to have forward compatible implementations.
type UnimplementedWedgeServer struct {
}

func (*UnimplementedWedgeServer) SetModel(context.Context, *SetModelRequest) (*Replay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetModel not implemented")
}
func (*UnimplementedWedgeServer) SetDevice(context.Context, *SetDeviceRequest) (*Replay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDevice not implemented")
}
func (*UnimplementedWedgeServer) SetValue(context.Context, *SetValueRequest) (*Replay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetValue not implemented")
}
func (*UnimplementedWedgeServer) SetState(context.Context, *SetStateRequest) (*Replay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetState not implemented")
}
func (*UnimplementedWedgeServer) GetControl(context.Context, *GetControlRequest) (*Control, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetControl not implemented")
}

func RegisterWedgeServer(s *grpc.Server, srv WedgeServer) {
	s.RegisterService(&_Wedge_serviceDesc, srv)
}

func _Wedge_SetModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WedgeServer).SetModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wedge.Wedge/SetModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WedgeServer).SetModel(ctx, req.(*SetModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wedge_SetDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WedgeServer).SetDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wedge.Wedge/SetDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WedgeServer).SetDevice(ctx, req.(*SetDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wedge_SetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WedgeServer).SetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wedge.Wedge/SetValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WedgeServer).SetValue(ctx, req.(*SetValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wedge_SetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WedgeServer).SetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wedge.Wedge/SetState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WedgeServer).SetState(ctx, req.(*SetStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wedge_GetControl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetControlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WedgeServer).GetControl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wedge.Wedge/GetControl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WedgeServer).GetControl(ctx, req.(*GetControlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Wedge_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wedge.Wedge",
	HandlerType: (*WedgeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetModel",
			Handler:    _Wedge_SetModel_Handler,
		},
		{
			MethodName: "SetDevice",
			Handler:    _Wedge_SetDevice_Handler,
		},
		{
			MethodName: "SetValue",
			Handler:    _Wedge_SetValue_Handler,
		},
		{
			MethodName: "SetState",
			Handler:    _Wedge_SetState_Handler,
		},
		{
			MethodName: "GetControl",
			Handler:    _Wedge_GetControl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wedge.proto",
}
