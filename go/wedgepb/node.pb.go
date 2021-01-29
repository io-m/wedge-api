// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: node.proto

package wedgepb

import (
	context "context"
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

type UpdateStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State    *State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	DeviceId uint32 `protobuf:"varint,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	ValueId  uint32 `protobuf:"varint,3,opt,name=value_id,json=valueId,proto3" json:"value_id,omitempty"`
}

func (x *UpdateStateRequest) Reset() {
	*x = UpdateStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStateRequest) ProtoMessage() {}

func (x *UpdateStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStateRequest.ProtoReflect.Descriptor instead.
func (*UpdateStateRequest) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateStateRequest) GetState() *State {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *UpdateStateRequest) GetDeviceId() uint32 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

func (x *UpdateStateRequest) GetValueId() uint32 {
	if x != nil {
		return x.ValueId
	}
	return 0
}

type DeleteDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId uint32 `protobuf:"varint,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *DeleteDeviceRequest) Reset() {
	*x = DeleteDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDeviceRequest) ProtoMessage() {}

func (x *DeleteDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDeviceRequest.ProtoReflect.Descriptor instead.
func (*DeleteDeviceRequest) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteDeviceRequest) GetDeviceId() uint32 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

type State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Data string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *State) Reset() {
	*x = State{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*State) ProtoMessage() {}

func (x *State) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[2]
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
	return file_node_proto_rawDescGZIP(), []int{2}
}

func (x *State) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *State) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type Replay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *Replay) Reset() {
	*x = Replay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Replay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Replay) ProtoMessage() {}

func (x *Replay) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[3]
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
	return file_node_proto_rawDescGZIP(), []int{3}
}

func (x *Replay) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_node_proto protoreflect.FileDescriptor

var file_node_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6e, 0x6f,
	0x64, 0x65, 0x22, 0x6f, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x2b, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x18, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x32, 0x7a,
	0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0c, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x22, 0x00, 0x12,
	0x39, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x19, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x22, 0x00, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x57, 0x61, 0x70, 0x70, 0x73, 0x74, 0x6f,
	0x2f, 0x77, 0x65, 0x64, 0x67, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x77, 0x65,
	0x64, 0x67, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_node_proto_rawDescOnce sync.Once
	file_node_proto_rawDescData = file_node_proto_rawDesc
)

func file_node_proto_rawDescGZIP() []byte {
	file_node_proto_rawDescOnce.Do(func() {
		file_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_proto_rawDescData)
	})
	return file_node_proto_rawDescData
}

var file_node_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_node_proto_goTypes = []interface{}{
	(*UpdateStateRequest)(nil),  // 0: node.UpdateStateRequest
	(*DeleteDeviceRequest)(nil), // 1: node.DeleteDeviceRequest
	(*State)(nil),               // 2: node.State
	(*Replay)(nil),              // 3: node.Replay
}
var file_node_proto_depIdxs = []int32{
	2, // 0: node.UpdateStateRequest.state:type_name -> node.State
	0, // 1: node.Node.UpdateState:input_type -> node.UpdateStateRequest
	1, // 2: node.Node.DeleteDevice:input_type -> node.DeleteDeviceRequest
	3, // 3: node.Node.UpdateState:output_type -> node.Replay
	3, // 4: node.Node.DeleteDevice:output_type -> node.Replay
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_node_proto_init() }
func file_node_proto_init() {
	if File_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStateRequest); i {
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
		file_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDeviceRequest); i {
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
		file_node_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_node_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_node_proto_goTypes,
		DependencyIndexes: file_node_proto_depIdxs,
		MessageInfos:      file_node_proto_msgTypes,
	}.Build()
	File_node_proto = out.File
	file_node_proto_rawDesc = nil
	file_node_proto_goTypes = nil
	file_node_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NodeClient is the client API for Node service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeClient interface {
	UpdateState(ctx context.Context, in *UpdateStateRequest, opts ...grpc.CallOption) (*Replay, error)
	DeleteDevice(ctx context.Context, in *DeleteDeviceRequest, opts ...grpc.CallOption) (*Replay, error)
}

type nodeClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeClient(cc grpc.ClientConnInterface) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) UpdateState(ctx context.Context, in *UpdateStateRequest, opts ...grpc.CallOption) (*Replay, error) {
	out := new(Replay)
	err := c.cc.Invoke(ctx, "/node.Node/UpdateState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) DeleteDevice(ctx context.Context, in *DeleteDeviceRequest, opts ...grpc.CallOption) (*Replay, error) {
	out := new(Replay)
	err := c.cc.Invoke(ctx, "/node.Node/DeleteDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServer is the server API for Node service.
type NodeServer interface {
	UpdateState(context.Context, *UpdateStateRequest) (*Replay, error)
	DeleteDevice(context.Context, *DeleteDeviceRequest) (*Replay, error)
}

// UnimplementedNodeServer can be embedded to have forward compatible implementations.
type UnimplementedNodeServer struct {
}

func (*UnimplementedNodeServer) UpdateState(context.Context, *UpdateStateRequest) (*Replay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateState not implemented")
}
func (*UnimplementedNodeServer) DeleteDevice(context.Context, *DeleteDeviceRequest) (*Replay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDevice not implemented")
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_UpdateState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).UpdateState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/node.Node/UpdateState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).UpdateState(ctx, req.(*UpdateStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_DeleteDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).DeleteDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/node.Node/DeleteDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).DeleteDevice(ctx, req.(*DeleteDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "node.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateState",
			Handler:    _Node_UpdateState_Handler,
		},
		{
			MethodName: "DeleteDevice",
			Handler:    _Node_DeleteDevice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node.proto",
}
