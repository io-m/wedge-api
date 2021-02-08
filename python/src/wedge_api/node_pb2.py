# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: node.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import slx_pb2 as slx__pb2

from slx_pb2 import *

DESCRIPTOR = _descriptor.FileDescriptor(
  name='node.proto',
  package='node',
  syntax='proto3',
  serialized_options=b'Z$github.com/Wappsto/wedge-api/go/node',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\nnode.proto\x12\x04node\x1a\tslx.proto\"P\n\x12UpdateStateRequest\x12\x15\n\x05state\x18\x01 \x01(\x0b\x32\x06.State\x12\x11\n\tdevice_id\x18\x02 \x01(\r\x12\x10\n\x08value_id\x18\x03 \x01(\r\"(\n\x13\x44\x65leteDeviceRequest\x12\x11\n\tdevice_id\x18\x01 \x01(\r2p\n\x04Node\x12\x32\n\x0bUpdateState\x12\x18.node.UpdateStateRequest\x1a\x07.Replay\"\x00\x12\x34\n\x0c\x44\x65leteDevice\x12\x19.node.DeleteDeviceRequest\x1a\x07.Replay\"\x00\x42&Z$github.com/Wappsto/wedge-api/go/nodeP\x00\x62\x06proto3'
  ,
  dependencies=[slx__pb2.DESCRIPTOR,],
  public_dependencies=[slx__pb2.DESCRIPTOR,])




_UPDATESTATEREQUEST = _descriptor.Descriptor(
  name='UpdateStateRequest',
  full_name='node.UpdateStateRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='state', full_name='node.UpdateStateRequest.state', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='device_id', full_name='node.UpdateStateRequest.device_id', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='value_id', full_name='node.UpdateStateRequest.value_id', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=31,
  serialized_end=111,
)


_DELETEDEVICEREQUEST = _descriptor.Descriptor(
  name='DeleteDeviceRequest',
  full_name='node.DeleteDeviceRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='device_id', full_name='node.DeleteDeviceRequest.device_id', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=113,
  serialized_end=153,
)

_UPDATESTATEREQUEST.fields_by_name['state'].message_type = slx__pb2._STATE
DESCRIPTOR.message_types_by_name['UpdateStateRequest'] = _UPDATESTATEREQUEST
DESCRIPTOR.message_types_by_name['DeleteDeviceRequest'] = _DELETEDEVICEREQUEST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

UpdateStateRequest = _reflection.GeneratedProtocolMessageType('UpdateStateRequest', (_message.Message,), {
  'DESCRIPTOR' : _UPDATESTATEREQUEST,
  '__module__' : 'node_pb2'
  # @@protoc_insertion_point(class_scope:node.UpdateStateRequest)
  })
_sym_db.RegisterMessage(UpdateStateRequest)

DeleteDeviceRequest = _reflection.GeneratedProtocolMessageType('DeleteDeviceRequest', (_message.Message,), {
  'DESCRIPTOR' : _DELETEDEVICEREQUEST,
  '__module__' : 'node_pb2'
  # @@protoc_insertion_point(class_scope:node.DeleteDeviceRequest)
  })
_sym_db.RegisterMessage(DeleteDeviceRequest)


DESCRIPTOR._options = None

_NODE = _descriptor.ServiceDescriptor(
  name='Node',
  full_name='node.Node',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=155,
  serialized_end=267,
  methods=[
  _descriptor.MethodDescriptor(
    name='UpdateState',
    full_name='node.Node.UpdateState',
    index=0,
    containing_service=None,
    input_type=_UPDATESTATEREQUEST,
    output_type=slx__pb2._REPLAY,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='DeleteDevice',
    full_name='node.Node.DeleteDevice',
    index=1,
    containing_service=None,
    input_type=_DELETEDEVICEREQUEST,
    output_type=slx__pb2._REPLAY,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_NODE)

DESCRIPTOR.services_by_name['Node'] = _NODE

# @@protoc_insertion_point(module_scope)
