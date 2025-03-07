// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: slx.proto

package slx

import (
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

type Type int32

const (
	Type_TYPE_UNSET Type = 0
	Type_Report     Type = 1
	Type_Control    Type = 2
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "TYPE_UNSET",
		1: "Report",
		2: "Control",
	}
	Type_value = map[string]int32{
		"TYPE_UNSET": 0,
		"Report":     1,
		"Control":    2,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_slx_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_slx_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_slx_proto_rawDescGZIP(), []int{0}
}

type NumberMapping int32

const (
	NumberMapping_off NumberMapping = 0
	NumberMapping_on  NumberMapping = 1
)

// Enum value maps for NumberMapping.
var (
	NumberMapping_name = map[int32]string{
		0: "off",
		1: "on",
	}
	NumberMapping_value = map[string]int32{
		"off": 0,
		"on":  1,
	}
)

func (x NumberMapping) Enum() *NumberMapping {
	p := new(NumberMapping)
	*p = x
	return p
}

func (x NumberMapping) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NumberMapping) Descriptor() protoreflect.EnumDescriptor {
	return file_slx_proto_enumTypes[1].Descriptor()
}

func (NumberMapping) Type() protoreflect.EnumType {
	return &file_slx_proto_enumTypes[1]
}

func (x NumberMapping) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NumberMapping.Descriptor instead.
func (NumberMapping) EnumDescriptor() ([]byte, []int) {
	return file_slx_proto_rawDescGZIP(), []int{1}
}

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type    string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_slx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_slx_proto_rawDescGZIP(), []int{0}
}

func (x *Meta) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Meta) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Meta) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data      string               `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Type      Type                 `protobuf:"varint,2,opt,name=type,proto3,enum=Type" json:"type,omitempty"`
	Timestamp *timestamp.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Status    string               `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Meta      *Meta                `protobuf:"bytes,5,opt,name=meta,proto3" json:"meta,omitempty"`
	Id        uint32               `protobuf:"varint,6,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *State) Reset() {
	*x = State{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*State) ProtoMessage() {}

func (x *State) ProtoReflect() protoreflect.Message {
	mi := &file_slx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use State.ProtoReflect.Descriptor instead.
func (*State) Descriptor() ([]byte, []int) {
	return file_slx_proto_rawDescGZIP(), []int{1}
}

func (x *State) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *State) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_TYPE_UNSET
}

func (x *State) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *State) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *State) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *State) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Number struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Min            float32       `protobuf:"fixed32,1,opt,name=min,proto3" json:"min,omitempty"`   // required
	Max            float32       `protobuf:"fixed32,2,opt,name=max,proto3" json:"max,omitempty"`   // required
	Step           float32       `protobuf:"fixed32,3,opt,name=step,proto3" json:"step,omitempty"` // required
	Unit           string        `protobuf:"bytes,4,opt,name=unit,proto3" json:"unit,omitempty"`
	SiConversion   string        `protobuf:"bytes,5,opt,name=si_conversion,json=siConversion,proto3" json:"si_conversion,omitempty"`
	OrderedMapping bool          `protobuf:"varint,6,opt,name=ordered_mapping,json=orderedMapping,proto3" json:"ordered_mapping,omitempty"`
	MeaningfulZero bool          `protobuf:"varint,7,opt,name=meaningful_zero,json=meaningfulZero,proto3" json:"meaningful_zero,omitempty"`
	Mapping        NumberMapping `protobuf:"varint,8,opt,name=mapping,proto3,enum=NumberMapping" json:"mapping,omitempty"`
}

func (x *Number) Reset() {
	*x = Number{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slx_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Number) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Number) ProtoMessage() {}

func (x *Number) ProtoReflect() protoreflect.Message {
	mi := &file_slx_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Number.ProtoReflect.Descriptor instead.
func (*Number) Descriptor() ([]byte, []int) {
	return file_slx_proto_rawDescGZIP(), []int{2}
}

func (x *Number) GetMin() float32 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *Number) GetMax() float32 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *Number) GetStep() float32 {
	if x != nil {
		return x.Step
	}
	return 0
}

func (x *Number) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *Number) GetSiConversion() string {
	if x != nil {
		return x.SiConversion
	}
	return ""
}

func (x *Number) GetOrderedMapping() bool {
	if x != nil {
		return x.OrderedMapping
	}
	return false
}

func (x *Number) GetMeaningfulZero() bool {
	if x != nil {
		return x.MeaningfulZero
	}
	return false
}

func (x *Number) GetMapping() NumberMapping {
	if x != nil {
		return x.Mapping
	}
	return NumberMapping_off
}

type String struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Max      float32 `protobuf:"fixed32,1,opt,name=max,proto3" json:"max,omitempty"`
	Encoding string  `protobuf:"bytes,2,opt,name=encoding,proto3" json:"encoding,omitempty"`
}

func (x *String) Reset() {
	*x = String{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slx_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *String) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*String) ProtoMessage() {}

func (x *String) ProtoReflect() protoreflect.Message {
	mi := &file_slx_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use String.ProtoReflect.Descriptor instead.
func (*String) Descriptor() ([]byte, []int) {
	return file_slx_proto_rawDescGZIP(), []int{3}
}

func (x *String) GetMax() float32 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *String) GetEncoding() string {
	if x != nil {
		return x.Encoding
	}
	return ""
}

type Blob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Max      float32 `protobuf:"fixed32,1,opt,name=max,proto3" json:"max,omitempty"`
	Encoding string  `protobuf:"bytes,2,opt,name=encoding,proto3" json:"encoding,omitempty"`
}

func (x *Blob) Reset() {
	*x = Blob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slx_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blob) ProtoMessage() {}

func (x *Blob) ProtoReflect() protoreflect.Message {
	mi := &file_slx_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blob.ProtoReflect.Descriptor instead.
func (*Blob) Descriptor() ([]byte, []int) {
	return file_slx_proto_rawDescGZIP(), []int{4}
}

func (x *Blob) GetMax() float32 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *Blob) GetEncoding() string {
	if x != nil {
		return x.Encoding
	}
	return ""
}

type Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *Info) Reset() {
	*x = Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slx_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Info) ProtoMessage() {}

func (x *Info) ProtoReflect() protoreflect.Message {
	mi := &file_slx_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Info.ProtoReflect.Descriptor instead.
func (*Info) Descriptor() ([]byte, []int) {
	return file_slx_proto_rawDescGZIP(), []int{5}
}

func (x *Info) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type        string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Period      string   `protobuf:"bytes,3,opt,name=period,proto3" json:"period,omitempty"`
	Delta       string   `protobuf:"bytes,4,opt,name=delta,proto3" json:"delta,omitempty"`
	Permission  string   `protobuf:"bytes,5,opt,name=permission,proto3" json:"permission,omitempty"`
	Description string   `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Status      string   `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	Number      *Number  `protobuf:"bytes,8,opt,name=number,proto3" json:"number,omitempty"`
	String_     *String  `protobuf:"bytes,9,opt,name=string,proto3" json:"string,omitempty"`
	Blob        *Blob    `protobuf:"bytes,10,opt,name=blob,proto3" json:"blob,omitempty"`
	State       []*State `protobuf:"bytes,11,rep,name=state,proto3" json:"state,omitempty"`
	Info        *Info    `protobuf:"bytes,12,opt,name=info,proto3" json:"info,omitempty"`
	Meta        *Meta    `protobuf:"bytes,13,opt,name=meta,proto3" json:"meta,omitempty"`
	Id          uint32   `protobuf:"varint,14,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slx_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_slx_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_slx_proto_rawDescGZIP(), []int{6}
}

func (x *Value) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Value) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Value) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *Value) GetDelta() string {
	if x != nil {
		return x.Delta
	}
	return ""
}

func (x *Value) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

func (x *Value) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Value) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Value) GetNumber() *Number {
	if x != nil {
		return x.Number
	}
	return nil
}

func (x *Value) GetString_() *String {
	if x != nil {
		return x.String_
	}
	return nil
}

func (x *Value) GetBlob() *Blob {
	if x != nil {
		return x.Blob
	}
	return nil
}

func (x *Value) GetState() []*State {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *Value) GetInfo() *Info {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *Value) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Value) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name               string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Manufacturer       string   `protobuf:"bytes,2,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	Product            string   `protobuf:"bytes,3,opt,name=product,proto3" json:"product,omitempty"`
	Version            string   `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	Serial             string   `protobuf:"bytes,5,opt,name=serial,proto3" json:"serial,omitempty"`
	Description        string   `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Protocol           string   `protobuf:"bytes,7,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Communication      string   `protobuf:"bytes,8,opt,name=communication,proto3" json:"communication,omitempty"`
	ControlTimeout     uint32   `protobuf:"varint,9,opt,name=control_timeout,json=controlTimeout,proto3" json:"control_timeout,omitempty"`
	ControlWhenOffline bool     `protobuf:"varint,10,opt,name=control_when_offline,json=controlWhenOffline,proto3" json:"control_when_offline,omitempty"`
	Value              []*Value `protobuf:"bytes,11,rep,name=value,proto3" json:"value,omitempty"`
	Info               *Info    `protobuf:"bytes,12,opt,name=info,proto3" json:"info,omitempty"`
	Meta               *Meta    `protobuf:"bytes,13,opt,name=meta,proto3" json:"meta,omitempty"`
	Id                 uint32   `protobuf:"varint,14,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slx_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_slx_proto_msgTypes[7]
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
	return file_slx_proto_rawDescGZIP(), []int{7}
}

func (x *Device) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Device) GetManufacturer() string {
	if x != nil {
		return x.Manufacturer
	}
	return ""
}

func (x *Device) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *Device) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Device) GetSerial() string {
	if x != nil {
		return x.Serial
	}
	return ""
}

func (x *Device) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Device) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *Device) GetCommunication() string {
	if x != nil {
		return x.Communication
	}
	return ""
}

func (x *Device) GetControlTimeout() uint32 {
	if x != nil {
		return x.ControlTimeout
	}
	return 0
}

func (x *Device) GetControlWhenOffline() bool {
	if x != nil {
		return x.ControlWhenOffline
	}
	return false
}

func (x *Device) GetValue() []*Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Device) GetInfo() *Info {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *Device) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Device) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_slx_proto protoreflect.FileDescriptor

var file_slx_proto_rawDesc = []byte{
	0x0a, 0x09, 0x73, 0x6c, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x04,
	0x4d, 0x65, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0xb3, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x19, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x05,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0xf5, 0x01, 0x0a, 0x06, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x6e, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12,
	0x23, 0x0a, 0x0d, 0x73, 0x69, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x69, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x5f,
	0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x27, 0x0a,
	0x0f, 0x6d, 0x65, 0x61, 0x6e, 0x69, 0x6e, 0x67, 0x66, 0x75, 0x6c, 0x5f, 0x7a, 0x65, 0x72, 0x6f,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x6d, 0x65, 0x61, 0x6e, 0x69, 0x6e, 0x67, 0x66,
	0x75, 0x6c, 0x5a, 0x65, 0x72, 0x6f, 0x12, 0x28, 0x0a, 0x07, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x22, 0x36, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x1a, 0x0a, 0x08,
	0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x34, 0x0a, 0x04, 0x42, 0x6c, 0x6f, 0x62,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6d,
	0x61, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x20,
	0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x22, 0xf8, 0x02, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65,
	0x6c, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61,
	0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x06, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x19, 0x0a, 0x04,
	0x62, 0x6c, 0x6f, 0x62, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x42, 0x6c, 0x6f,
	0x62, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x62, 0x12, 0x1c, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x12, 0x19, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0xaf, 0x03, 0x0a, 0x06,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61,
	0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27,
	0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x5f, 0x77, 0x68, 0x65, 0x6e, 0x5f, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x57, 0x68,
	0x65, 0x6e, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1c, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x12, 0x19, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x05, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x2a, 0x2f, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x10, 0x02, 0x2a, 0x20,
	0x0a, 0x0d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12,
	0x07, 0x0a, 0x03, 0x6f, 0x66, 0x66, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x6f, 0x6e, 0x10, 0x01,
	0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x57,
	0x61, 0x70, 0x70, 0x73, 0x74, 0x6f, 0x2f, 0x77, 0x65, 0x64, 0x67, 0x65, 0x2d, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x6c, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_slx_proto_rawDescOnce sync.Once
	file_slx_proto_rawDescData = file_slx_proto_rawDesc
)

func file_slx_proto_rawDescGZIP() []byte {
	file_slx_proto_rawDescOnce.Do(func() {
		file_slx_proto_rawDescData = protoimpl.X.CompressGZIP(file_slx_proto_rawDescData)
	})
	return file_slx_proto_rawDescData
}

var file_slx_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_slx_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_slx_proto_goTypes = []interface{}{
	(Type)(0),                   // 0: Type
	(NumberMapping)(0),          // 1: NumberMapping
	(*Meta)(nil),                // 2: Meta
	(*State)(nil),               // 3: State
	(*Number)(nil),              // 4: Number
	(*String)(nil),              // 5: String
	(*Blob)(nil),                // 6: Blob
	(*Info)(nil),                // 7: Info
	(*Value)(nil),               // 8: Value
	(*Device)(nil),              // 9: Device
	(*timestamp.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_slx_proto_depIdxs = []int32{
	0,  // 0: State.type:type_name -> Type
	10, // 1: State.timestamp:type_name -> google.protobuf.Timestamp
	2,  // 2: State.meta:type_name -> Meta
	1,  // 3: Number.mapping:type_name -> NumberMapping
	4,  // 4: Value.number:type_name -> Number
	5,  // 5: Value.string:type_name -> String
	6,  // 6: Value.blob:type_name -> Blob
	3,  // 7: Value.state:type_name -> State
	7,  // 8: Value.info:type_name -> Info
	2,  // 9: Value.meta:type_name -> Meta
	8,  // 10: Device.value:type_name -> Value
	7,  // 11: Device.info:type_name -> Info
	2,  // 12: Device.meta:type_name -> Meta
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_slx_proto_init() }
func file_slx_proto_init() {
	if File_slx_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_slx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
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
		file_slx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*State); i {
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
		file_slx_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Number); i {
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
		file_slx_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*String); i {
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
		file_slx_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Blob); i {
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
		file_slx_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Info); i {
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
		file_slx_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
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
		file_slx_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_slx_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_slx_proto_goTypes,
		DependencyIndexes: file_slx_proto_depIdxs,
		EnumInfos:         file_slx_proto_enumTypes,
		MessageInfos:      file_slx_proto_msgTypes,
	}.Build()
	File_slx_proto = out.File
	file_slx_proto_rawDesc = nil
	file_slx_proto_goTypes = nil
	file_slx_proto_depIdxs = nil
}
