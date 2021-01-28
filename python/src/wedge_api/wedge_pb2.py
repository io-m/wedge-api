# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: wedge-api/wedge.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='wedge-api/wedge.proto',
  package='wedge',
  syntax='proto3',
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x15wedge-api/wedge.proto\x12\x05wedge\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\x81\x01\n\x05State\x12\n\n\x02id\x18\x01 \x01(\r\x12\x0c\n\x04\x64\x61ta\x18\x02 \x01(\t\x12\x1f\n\x04type\x18\x03 \x01(\x0e\x32\x11.wedge.StatusType\x12-\n\ttimestamp\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x0e\n\x06status\x18\x05 \x01(\t\">\n\x06Number\x12\x0b\n\x03min\x18\x01 \x01(\x02\x12\x0b\n\x03max\x18\x02 \x01(\x02\x12\x0c\n\x04step\x18\x03 \x01(\x02\x12\x0c\n\x04unit\x18\x04 \x01(\t\"\'\n\x06String\x12\x0b\n\x03max\x18\x01 \x01(\x02\x12\x10\n\x08\x65ncoding\x18\x02 \x01(\t\"%\n\x04\x42lob\x12\x0b\n\x03max\x18\x01 \x01(\x02\x12\x10\n\x08\x65ncoding\x18\x02 \x01(\t\"\xdc\x01\n\x05Value\x12\n\n\x02id\x18\x01 \x01(\r\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x0c\n\x04type\x18\x03 \x01(\t\x12\x12\n\npermission\x18\x04 \x01(\t\x12\x0e\n\x06status\x18\x05 \x01(\t\x12\x1f\n\x06number\x18\x06 \x01(\x0b\x32\r.wedge.NumberH\x00\x12\x1f\n\x06string\x18\x07 \x01(\x0b\x32\r.wedge.StringH\x00\x12\x1b\n\x04\x62lob\x18\x08 \x01(\x0b\x32\x0b.wedge.BlobH\x00\x12\x1b\n\x05state\x18\t \x03(\x0b\x32\x0c.wedge.StateB\x0b\n\tvaluetype\"\x87\x01\n\x06\x44\x65vice\x12\n\n\x02id\x18\x01 \x01(\r\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x0e\n\x06serial\x18\x03 \x01(\t\x12\x14\n\x0cmanufacturer\x18\x04 \x01(\t\x12\x0f\n\x07product\x18\x05 \x01(\t\x12\x0f\n\x07version\x18\x06 \x01(\t\x12\x1b\n\x05value\x18\x07 \x03(\x0b\x32\x0c.wedge.Value\" \n\x02Me\x12\x0c\n\x04host\x18\x02 \x01(\t\x12\x0c\n\x04port\x18\x03 \x01(\r\"=\n\x05Model\x12\x15\n\x02me\x18\x01 \x01(\x0b\x32\t.wedge.Me\x12\x1d\n\x06\x64\x65vice\x18\x02 \x03(\x0b\x32\r.wedge.Device\".\n\x0fSetModelRequest\x12\x1b\n\x05model\x18\x01 \x01(\x0b\x32\x0c.wedge.Model\"H\n\x10SetDeviceRequest\x12\x15\n\x02me\x18\x01 \x01(\x0b\x32\t.wedge.Me\x12\x1d\n\x06\x64\x65vice\x18\x02 \x01(\x0b\x32\r.wedge.Device\"X\n\x0fSetValueRequest\x12\x15\n\x02me\x18\x01 \x01(\x0b\x32\t.wedge.Me\x12\x11\n\tdevice_id\x18\x02 \x01(\r\x12\x1b\n\x05value\x18\x03 \x01(\x0b\x32\x0c.wedge.Value\"j\n\x0fSetStateRequest\x12\x15\n\x02me\x18\x01 \x01(\x0b\x32\t.wedge.Me\x12\x11\n\tdevice_id\x18\x02 \x01(\r\x12\x10\n\x08value_id\x18\x03 \x01(\r\x12\x1b\n\x05state\x18\x04 \x01(\x0b\x32\x0c.wedge.State*%\n\nStatusType\x12\n\n\x06Report\x10\x00\x12\x0b\n\x07\x43ontrol\x10\x01\x32\x81\x02\n\x05Wedge\x12<\n\x08SetModel\x12\x16.wedge.SetModelRequest\x1a\x16.google.protobuf.Empty\"\x00\x12>\n\tSetDevice\x12\x17.wedge.SetDeviceRequest\x1a\x16.google.protobuf.Empty\"\x00\x12<\n\x08SetValue\x12\x16.wedge.SetValueRequest\x1a\x16.google.protobuf.Empty\"\x00\x12<\n\x08SetState\x12\x16.wedge.SetStateRequest\x1a\x16.google.protobuf.Empty\"\x00\x62\x06proto3'
  ,
  dependencies=[google_dot_protobuf_dot_empty__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,])

_STATUSTYPE = _descriptor.EnumDescriptor(
  name='StatusType',
  full_name='wedge.StatusType',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='Report', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Control', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1148,
  serialized_end=1185,
)
_sym_db.RegisterEnumDescriptor(_STATUSTYPE)

StatusType = enum_type_wrapper.EnumTypeWrapper(_STATUSTYPE)
Report = 0
Control = 1



_STATE = _descriptor.Descriptor(
  name='State',
  full_name='wedge.State',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='wedge.State.id', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='data', full_name='wedge.State.data', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='type', full_name='wedge.State.type', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='timestamp', full_name='wedge.State.timestamp', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='status', full_name='wedge.State.status', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
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
  serialized_start=95,
  serialized_end=224,
)


_NUMBER = _descriptor.Descriptor(
  name='Number',
  full_name='wedge.Number',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='min', full_name='wedge.Number.min', index=0,
      number=1, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='max', full_name='wedge.Number.max', index=1,
      number=2, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='step', full_name='wedge.Number.step', index=2,
      number=3, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='unit', full_name='wedge.Number.unit', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
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
  serialized_start=226,
  serialized_end=288,
)


_STRING = _descriptor.Descriptor(
  name='String',
  full_name='wedge.String',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='max', full_name='wedge.String.max', index=0,
      number=1, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='encoding', full_name='wedge.String.encoding', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
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
  serialized_start=290,
  serialized_end=329,
)


_BLOB = _descriptor.Descriptor(
  name='Blob',
  full_name='wedge.Blob',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='max', full_name='wedge.Blob.max', index=0,
      number=1, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='encoding', full_name='wedge.Blob.encoding', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
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
  serialized_start=331,
  serialized_end=368,
)


_VALUE = _descriptor.Descriptor(
  name='Value',
  full_name='wedge.Value',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='wedge.Value.id', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='wedge.Value.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='type', full_name='wedge.Value.type', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='permission', full_name='wedge.Value.permission', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='status', full_name='wedge.Value.status', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='number', full_name='wedge.Value.number', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='string', full_name='wedge.Value.string', index=6,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='blob', full_name='wedge.Value.blob', index=7,
      number=8, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='state', full_name='wedge.Value.state', index=8,
      number=9, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
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
    _descriptor.OneofDescriptor(
      name='valuetype', full_name='wedge.Value.valuetype',
      index=0, containing_type=None,
      create_key=_descriptor._internal_create_key,
    fields=[]),
  ],
  serialized_start=371,
  serialized_end=591,
)


_DEVICE = _descriptor.Descriptor(
  name='Device',
  full_name='wedge.Device',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='wedge.Device.id', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='wedge.Device.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='serial', full_name='wedge.Device.serial', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='manufacturer', full_name='wedge.Device.manufacturer', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='product', full_name='wedge.Device.product', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='version', full_name='wedge.Device.version', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='value', full_name='wedge.Device.value', index=6,
      number=7, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
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
  serialized_start=594,
  serialized_end=729,
)


_ME = _descriptor.Descriptor(
  name='Me',
  full_name='wedge.Me',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='host', full_name='wedge.Me.host', index=0,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='port', full_name='wedge.Me.port', index=1,
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
  serialized_start=731,
  serialized_end=763,
)


_MODEL = _descriptor.Descriptor(
  name='Model',
  full_name='wedge.Model',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='me', full_name='wedge.Model.me', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='device', full_name='wedge.Model.device', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
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
  serialized_start=765,
  serialized_end=826,
)


_SETMODELREQUEST = _descriptor.Descriptor(
  name='SetModelRequest',
  full_name='wedge.SetModelRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='model', full_name='wedge.SetModelRequest.model', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=828,
  serialized_end=874,
)


_SETDEVICEREQUEST = _descriptor.Descriptor(
  name='SetDeviceRequest',
  full_name='wedge.SetDeviceRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='me', full_name='wedge.SetDeviceRequest.me', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='device', full_name='wedge.SetDeviceRequest.device', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=876,
  serialized_end=948,
)


_SETVALUEREQUEST = _descriptor.Descriptor(
  name='SetValueRequest',
  full_name='wedge.SetValueRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='me', full_name='wedge.SetValueRequest.me', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='device_id', full_name='wedge.SetValueRequest.device_id', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='value', full_name='wedge.SetValueRequest.value', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=950,
  serialized_end=1038,
)


_SETSTATEREQUEST = _descriptor.Descriptor(
  name='SetStateRequest',
  full_name='wedge.SetStateRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='me', full_name='wedge.SetStateRequest.me', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='device_id', full_name='wedge.SetStateRequest.device_id', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='value_id', full_name='wedge.SetStateRequest.value_id', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='state', full_name='wedge.SetStateRequest.state', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=1040,
  serialized_end=1146,
)

_STATE.fields_by_name['type'].enum_type = _STATUSTYPE
_STATE.fields_by_name['timestamp'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_VALUE.fields_by_name['number'].message_type = _NUMBER
_VALUE.fields_by_name['string'].message_type = _STRING
_VALUE.fields_by_name['blob'].message_type = _BLOB
_VALUE.fields_by_name['state'].message_type = _STATE
_VALUE.oneofs_by_name['valuetype'].fields.append(
  _VALUE.fields_by_name['number'])
_VALUE.fields_by_name['number'].containing_oneof = _VALUE.oneofs_by_name['valuetype']
_VALUE.oneofs_by_name['valuetype'].fields.append(
  _VALUE.fields_by_name['string'])
_VALUE.fields_by_name['string'].containing_oneof = _VALUE.oneofs_by_name['valuetype']
_VALUE.oneofs_by_name['valuetype'].fields.append(
  _VALUE.fields_by_name['blob'])
_VALUE.fields_by_name['blob'].containing_oneof = _VALUE.oneofs_by_name['valuetype']
_DEVICE.fields_by_name['value'].message_type = _VALUE
_MODEL.fields_by_name['me'].message_type = _ME
_MODEL.fields_by_name['device'].message_type = _DEVICE
_SETMODELREQUEST.fields_by_name['model'].message_type = _MODEL
_SETDEVICEREQUEST.fields_by_name['me'].message_type = _ME
_SETDEVICEREQUEST.fields_by_name['device'].message_type = _DEVICE
_SETVALUEREQUEST.fields_by_name['me'].message_type = _ME
_SETVALUEREQUEST.fields_by_name['value'].message_type = _VALUE
_SETSTATEREQUEST.fields_by_name['me'].message_type = _ME
_SETSTATEREQUEST.fields_by_name['state'].message_type = _STATE
DESCRIPTOR.message_types_by_name['State'] = _STATE
DESCRIPTOR.message_types_by_name['Number'] = _NUMBER
DESCRIPTOR.message_types_by_name['String'] = _STRING
DESCRIPTOR.message_types_by_name['Blob'] = _BLOB
DESCRIPTOR.message_types_by_name['Value'] = _VALUE
DESCRIPTOR.message_types_by_name['Device'] = _DEVICE
DESCRIPTOR.message_types_by_name['Me'] = _ME
DESCRIPTOR.message_types_by_name['Model'] = _MODEL
DESCRIPTOR.message_types_by_name['SetModelRequest'] = _SETMODELREQUEST
DESCRIPTOR.message_types_by_name['SetDeviceRequest'] = _SETDEVICEREQUEST
DESCRIPTOR.message_types_by_name['SetValueRequest'] = _SETVALUEREQUEST
DESCRIPTOR.message_types_by_name['SetStateRequest'] = _SETSTATEREQUEST
DESCRIPTOR.enum_types_by_name['StatusType'] = _STATUSTYPE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

State = _reflection.GeneratedProtocolMessageType('State', (_message.Message,), {
  'DESCRIPTOR' : _STATE,
  '__module__' : 'wedge_api.wedge_pb2'
  # @@protoc_insertion_point(class_scope:wedge.State)
  })
_sym_db.RegisterMessage(State)

Number = _reflection.GeneratedProtocolMessageType('Number', (_message.Message,), {
  'DESCRIPTOR' : _NUMBER,
  '__module__' : 'wedge_api.wedge_pb2'
  # @@protoc_insertion_point(class_scope:wedge.Number)
  })
_sym_db.RegisterMessage(Number)

String = _reflection.GeneratedProtocolMessageType('String', (_message.Message,), {
  'DESCRIPTOR' : _STRING,
  '__module__' : 'wedge_api.wedge_pb2'
  # @@protoc_insertion_point(class_scope:wedge.String)
  })
_sym_db.RegisterMessage(String)

Blob = _reflection.GeneratedProtocolMessageType('Blob', (_message.Message,), {
  'DESCRIPTOR' : _BLOB,
  '__module__' : 'wedge_api.wedge_pb2'
  # @@protoc_insertion_point(class_scope:wedge.Blob)
  })
_sym_db.RegisterMessage(Blob)

Value = _reflection.GeneratedProtocolMessageType('Value', (_message.Message,), {
  'DESCRIPTOR' : _VALUE,
  '__module__' : 'wedge_api.wedge_pb2'
  # @@protoc_insertion_point(class_scope:wedge.Value)
  })
_sym_db.RegisterMessage(Value)

Device = _reflection.GeneratedProtocolMessageType('Device', (_message.Message,), {
  'DESCRIPTOR' : _DEVICE,
  '__module__' : 'wedge_api.wedge_pb2'
  # @@protoc_insertion_point(class_scope:wedge.Device)
  })
_sym_db.RegisterMessage(Device)

Me = _reflection.GeneratedProtocolMessageType('Me', (_message.Message,), {
  'DESCRIPTOR' : _ME,
  '__module__' : 'wedge_api.wedge_pb2'
  # @@protoc_insertion_point(class_scope:wedge.Me)
  })
_sym_db.RegisterMessage(Me)

Model = _reflection.GeneratedProtocolMessageType('Model', (_message.Message,), {
  'DESCRIPTOR' : _MODEL,
  '__module__' : 'wedge_api.wedge_pb2'
  # @@protoc_insertion_point(class_scope:wedge.Model)
  })
_sym_db.RegisterMessage(Model)

SetModelRequest = _reflection.GeneratedProtocolMessageType('SetModelRequest', (_message.Message,), {
  'DESCRIPTOR' : _SETMODELREQUEST,
  '__module__' : 'wedge_api.wedge_pb2'
  # @@protoc_insertion_point(class_scope:wedge.SetModelRequest)
  })
_sym_db.RegisterMessage(SetModelRequest)

SetDeviceRequest = _reflection.GeneratedProtocolMessageType('SetDeviceRequest', (_message.Message,), {
  'DESCRIPTOR' : _SETDEVICEREQUEST,
  '__module__' : 'wedge_api.wedge_pb2'
  # @@protoc_insertion_point(class_scope:wedge.SetDeviceRequest)
  })
_sym_db.RegisterMessage(SetDeviceRequest)

SetValueRequest = _reflection.GeneratedProtocolMessageType('SetValueRequest', (_message.Message,), {
  'DESCRIPTOR' : _SETVALUEREQUEST,
  '__module__' : 'wedge_api.wedge_pb2'
  # @@protoc_insertion_point(class_scope:wedge.SetValueRequest)
  })
_sym_db.RegisterMessage(SetValueRequest)

SetStateRequest = _reflection.GeneratedProtocolMessageType('SetStateRequest', (_message.Message,), {
  'DESCRIPTOR' : _SETSTATEREQUEST,
  '__module__' : 'wedge_api.wedge_pb2'
  # @@protoc_insertion_point(class_scope:wedge.SetStateRequest)
  })
_sym_db.RegisterMessage(SetStateRequest)



_WEDGE = _descriptor.ServiceDescriptor(
  name='Wedge',
  full_name='wedge.Wedge',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=1188,
  serialized_end=1445,
  methods=[
  _descriptor.MethodDescriptor(
    name='SetModel',
    full_name='wedge.Wedge.SetModel',
    index=0,
    containing_service=None,
    input_type=_SETMODELREQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='SetDevice',
    full_name='wedge.Wedge.SetDevice',
    index=1,
    containing_service=None,
    input_type=_SETDEVICEREQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='SetValue',
    full_name='wedge.Wedge.SetValue',
    index=2,
    containing_service=None,
    input_type=_SETVALUEREQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='SetState',
    full_name='wedge.Wedge.SetState',
    index=3,
    containing_service=None,
    input_type=_SETSTATEREQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_WEDGE)

DESCRIPTOR.services_by_name['Wedge'] = _WEDGE

# @@protoc_insertion_point(module_scope)
