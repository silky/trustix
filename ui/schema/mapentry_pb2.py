# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema/mapentry.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='schema/mapentry.proto',
  package='',
  syntax='proto2',
  serialized_options=b'Z\037github.com/tweag/trustix/schema',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x15schema/mapentry.proto\")\n\x08MapEntry\x12\x0e\n\x06\x44igest\x18\x01 \x02(\x0c\x12\r\n\x05Index\x18\x02 \x02(\x04\x42!Z\x1fgithub.com/tweag/trustix/schema'
)




_MAPENTRY = _descriptor.Descriptor(
  name='MapEntry',
  full_name='MapEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='Digest', full_name='MapEntry.Digest', index=0,
      number=1, type=12, cpp_type=9, label=2,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Index', full_name='MapEntry.Index', index=1,
      number=2, type=4, cpp_type=4, label=2,
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
  syntax='proto2',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=25,
  serialized_end=66,
)

DESCRIPTOR.message_types_by_name['MapEntry'] = _MAPENTRY
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

MapEntry = _reflection.GeneratedProtocolMessageType('MapEntry', (_message.Message,), {
  'DESCRIPTOR' : _MAPENTRY,
  '__module__' : 'schema.mapentry_pb2'
  # @@protoc_insertion_point(class_scope:MapEntry)
  })
_sym_db.RegisterMessage(MapEntry)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
