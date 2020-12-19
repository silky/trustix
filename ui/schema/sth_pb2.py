# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema/sth.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='schema/sth.proto',
  package='',
  syntax='proto2',
  serialized_options=b'Z\037github.com/tweag/trustix/schema',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x10schema/sth.proto\"p\n\x03STH\x12\x0f\n\x07LogRoot\x18\x01 \x02(\x0c\x12\x10\n\x08TreeSize\x18\x02 \x02(\x04\x12\x0f\n\x07MapRoot\x18\x03 \x02(\x0c\x12\x0e\n\x06MHRoot\x18\x04 \x02(\x0c\x12\x12\n\nMHTreeSize\x18\x05 \x02(\x04\x12\x11\n\tSignature\x18\x06 \x02(\x0c\x42!Z\x1fgithub.com/tweag/trustix/schema'
)




_STH = _descriptor.Descriptor(
  name='STH',
  full_name='STH',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='LogRoot', full_name='STH.LogRoot', index=0,
      number=1, type=12, cpp_type=9, label=2,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='TreeSize', full_name='STH.TreeSize', index=1,
      number=2, type=4, cpp_type=4, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='MapRoot', full_name='STH.MapRoot', index=2,
      number=3, type=12, cpp_type=9, label=2,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='MHRoot', full_name='STH.MHRoot', index=3,
      number=4, type=12, cpp_type=9, label=2,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='MHTreeSize', full_name='STH.MHTreeSize', index=4,
      number=5, type=4, cpp_type=4, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Signature', full_name='STH.Signature', index=5,
      number=6, type=12, cpp_type=9, label=2,
      has_default_value=False, default_value=b"",
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
  serialized_start=20,
  serialized_end=132,
)

DESCRIPTOR.message_types_by_name['STH'] = _STH
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

STH = _reflection.GeneratedProtocolMessageType('STH', (_message.Message,), {
  'DESCRIPTOR' : _STH,
  '__module__' : 'schema.sth_pb2'
  # @@protoc_insertion_point(class_scope:STH)
  })
_sym_db.RegisterMessage(STH)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
