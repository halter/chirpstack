# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: chirpstack-api/stream/backend_interfaces.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n.chirpstack-api/stream/backend_interfaces.proto\x12\x06stream\x1a\x1fgoogle/protobuf/timestamp.proto\"\xf3\x01\n\x18\x42\x61\x63kendInterfacesRequest\x12\x11\n\tsender_id\x18\x01 \x01(\t\x12\x13\n\x0breceiver_id\x18\x02 \x01(\t\x12(\n\x04time\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x16\n\x0etransaction_id\x18\x04 \x01(\r\x12\x14\n\x0cmessage_type\x18\x05 \x01(\t\x12\x13\n\x0bresult_code\x18\x06 \x01(\t\x12\x14\n\x0crequest_body\x18\x07 \x01(\t\x12\x15\n\rrequest_error\x18\x08 \x01(\t\x12\x15\n\rresponse_body\x18\t \x01(\tB{\n\x18io.chirpstack.api.streamB\x16\x42\x61\x63kendInterfacesProtoP\x01Z1github.com/chirpstack/chirpstack/api/go/v4/stream\xaa\x02\x11\x43hirpstack.Streamb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'chirpstack_api.stream.backend_interfaces_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\030io.chirpstack.api.streamB\026BackendInterfacesProtoP\001Z1github.com/chirpstack/chirpstack/api/go/v4/stream\252\002\021Chirpstack.Stream'
  _globals['_BACKENDINTERFACESREQUEST']._serialized_start=92
  _globals['_BACKENDINTERFACESREQUEST']._serialized_end=335
# @@protoc_insertion_point(module_scope)
