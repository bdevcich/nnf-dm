# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: datamovement.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x12\x64\x61tamovement.proto\x12\x0c\x64\x61tamovement\"u\n\x19\x44\x61taMovementCreateRequest\x12\x10\n\x08workflow\x18\x03 \x01(\t\x12\x11\n\tnamespace\x18\x04 \x01(\t\x12\x0e\n\x06source\x18\x05 \x01(\t\x12\x13\n\x0b\x64\x65stination\x18\x06 \x01(\t\x12\x0e\n\x06\x64ryrun\x18\x07 \x01(\x08\"\x9e\x01\n\x1a\x44\x61taMovementCreateResponse\x12\x0b\n\x03uid\x18\x01 \x01(\t\x12?\n\x06status\x18\x02 \x01(\x0e\x32/.datamovement.DataMovementCreateResponse.Status\x12\x0f\n\x07message\x18\x03 \x01(\t\"!\n\x06Status\x12\x0b\n\x07\x43REATED\x10\x00\x12\n\n\x06\x46\x41ILED\x10\x01\"=\n\x19\x44\x61taMovementStatusRequest\x12\x0b\n\x03uid\x18\x01 \x01(\t\x12\x13\n\x0bmaxWaitTime\x18\x02 \x01(\x03\"\xd3\x02\n\x1a\x44\x61taMovementStatusResponse\x12=\n\x05state\x18\x01 \x01(\x0e\x32..datamovement.DataMovementStatusResponse.State\x12?\n\x06status\x18\x02 \x01(\x0e\x32/.datamovement.DataMovementStatusResponse.Status\x12\x0f\n\x07message\x18\x03 \x01(\t\"Q\n\x05State\x12\x0b\n\x07PENDING\x10\x00\x12\x0c\n\x08STARTING\x10\x01\x12\x0b\n\x07RUNNING\x10\x02\x12\r\n\tCOMPLETED\x10\x03\x12\x11\n\rUNKNOWN_STATE\x10\x04\"Q\n\x06Status\x12\x0b\n\x07INVALID\x10\x00\x12\r\n\tNOT_FOUND\x10\x01\x12\x0b\n\x07SUCCESS\x10\x02\x12\n\n\x06\x46\x41ILED\x10\x03\x12\x12\n\x0eUNKNOWN_STATUS\x10\x04\"(\n\x19\x44\x61taMovementDeleteRequest\x12\x0b\n\x03uid\x18\x01 \x01(\t\"\xa9\x01\n\x1a\x44\x61taMovementDeleteResponse\x12?\n\x06status\x18\x01 \x01(\x0e\x32/.datamovement.DataMovementDeleteResponse.Status\"J\n\x06Status\x12\x0b\n\x07INVALID\x10\x00\x12\r\n\tNOT_FOUND\x10\x01\x12\x0b\n\x07\x44\x45LETED\x10\x02\x12\n\n\x06\x41\x43TIVE\x10\x03\x12\x0b\n\x07UNKNOWN\x10\x04\">\n\x17\x44\x61taMovementListRequest\x12\x10\n\x08workflow\x18\x01 \x01(\t\x12\x11\n\tnamespace\x18\x02 \x01(\t\"(\n\x18\x44\x61taMovementListResponse\x12\x0c\n\x04uids\x18\x01 \x03(\t2\x81\x03\n\tDataMover\x12]\n\x06\x43reate\x12\'.datamovement.DataMovementCreateRequest\x1a(.datamovement.DataMovementCreateResponse\"\x00\x12]\n\x06Status\x12\'.datamovement.DataMovementStatusRequest\x1a(.datamovement.DataMovementStatusResponse\"\x00\x12]\n\x06\x44\x65lete\x12\'.datamovement.DataMovementDeleteRequest\x1a(.datamovement.DataMovementDeleteResponse\"\x00\x12W\n\x04List\x12%.datamovement.DataMovementListRequest\x1a&.datamovement.DataMovementListResponse\"\x00\x42W\n\x1d\x63om.hpe.cray.nnf.datamovementB\x11\x44\x61taMovementProtoP\x01Z!nnf.cray.hpe.com/datamovement/apib\x06proto3')



_DATAMOVEMENTCREATEREQUEST = DESCRIPTOR.message_types_by_name['DataMovementCreateRequest']
_DATAMOVEMENTCREATERESPONSE = DESCRIPTOR.message_types_by_name['DataMovementCreateResponse']
_DATAMOVEMENTSTATUSREQUEST = DESCRIPTOR.message_types_by_name['DataMovementStatusRequest']
_DATAMOVEMENTSTATUSRESPONSE = DESCRIPTOR.message_types_by_name['DataMovementStatusResponse']
_DATAMOVEMENTDELETEREQUEST = DESCRIPTOR.message_types_by_name['DataMovementDeleteRequest']
_DATAMOVEMENTDELETERESPONSE = DESCRIPTOR.message_types_by_name['DataMovementDeleteResponse']
_DATAMOVEMENTLISTREQUEST = DESCRIPTOR.message_types_by_name['DataMovementListRequest']
_DATAMOVEMENTLISTRESPONSE = DESCRIPTOR.message_types_by_name['DataMovementListResponse']
_DATAMOVEMENTCREATERESPONSE_STATUS = _DATAMOVEMENTCREATERESPONSE.enum_types_by_name['Status']
_DATAMOVEMENTSTATUSRESPONSE_STATE = _DATAMOVEMENTSTATUSRESPONSE.enum_types_by_name['State']
_DATAMOVEMENTSTATUSRESPONSE_STATUS = _DATAMOVEMENTSTATUSRESPONSE.enum_types_by_name['Status']
_DATAMOVEMENTDELETERESPONSE_STATUS = _DATAMOVEMENTDELETERESPONSE.enum_types_by_name['Status']
DataMovementCreateRequest = _reflection.GeneratedProtocolMessageType('DataMovementCreateRequest', (_message.Message,), {
  'DESCRIPTOR' : _DATAMOVEMENTCREATEREQUEST,
  '__module__' : 'datamovement_pb2'
  # @@protoc_insertion_point(class_scope:datamovement.DataMovementCreateRequest)
  })
_sym_db.RegisterMessage(DataMovementCreateRequest)

DataMovementCreateResponse = _reflection.GeneratedProtocolMessageType('DataMovementCreateResponse', (_message.Message,), {
  'DESCRIPTOR' : _DATAMOVEMENTCREATERESPONSE,
  '__module__' : 'datamovement_pb2'
  # @@protoc_insertion_point(class_scope:datamovement.DataMovementCreateResponse)
  })
_sym_db.RegisterMessage(DataMovementCreateResponse)

DataMovementStatusRequest = _reflection.GeneratedProtocolMessageType('DataMovementStatusRequest', (_message.Message,), {
  'DESCRIPTOR' : _DATAMOVEMENTSTATUSREQUEST,
  '__module__' : 'datamovement_pb2'
  # @@protoc_insertion_point(class_scope:datamovement.DataMovementStatusRequest)
  })
_sym_db.RegisterMessage(DataMovementStatusRequest)

DataMovementStatusResponse = _reflection.GeneratedProtocolMessageType('DataMovementStatusResponse', (_message.Message,), {
  'DESCRIPTOR' : _DATAMOVEMENTSTATUSRESPONSE,
  '__module__' : 'datamovement_pb2'
  # @@protoc_insertion_point(class_scope:datamovement.DataMovementStatusResponse)
  })
_sym_db.RegisterMessage(DataMovementStatusResponse)

DataMovementDeleteRequest = _reflection.GeneratedProtocolMessageType('DataMovementDeleteRequest', (_message.Message,), {
  'DESCRIPTOR' : _DATAMOVEMENTDELETEREQUEST,
  '__module__' : 'datamovement_pb2'
  # @@protoc_insertion_point(class_scope:datamovement.DataMovementDeleteRequest)
  })
_sym_db.RegisterMessage(DataMovementDeleteRequest)

DataMovementDeleteResponse = _reflection.GeneratedProtocolMessageType('DataMovementDeleteResponse', (_message.Message,), {
  'DESCRIPTOR' : _DATAMOVEMENTDELETERESPONSE,
  '__module__' : 'datamovement_pb2'
  # @@protoc_insertion_point(class_scope:datamovement.DataMovementDeleteResponse)
  })
_sym_db.RegisterMessage(DataMovementDeleteResponse)

DataMovementListRequest = _reflection.GeneratedProtocolMessageType('DataMovementListRequest', (_message.Message,), {
  'DESCRIPTOR' : _DATAMOVEMENTLISTREQUEST,
  '__module__' : 'datamovement_pb2'
  # @@protoc_insertion_point(class_scope:datamovement.DataMovementListRequest)
  })
_sym_db.RegisterMessage(DataMovementListRequest)

DataMovementListResponse = _reflection.GeneratedProtocolMessageType('DataMovementListResponse', (_message.Message,), {
  'DESCRIPTOR' : _DATAMOVEMENTLISTRESPONSE,
  '__module__' : 'datamovement_pb2'
  # @@protoc_insertion_point(class_scope:datamovement.DataMovementListResponse)
  })
_sym_db.RegisterMessage(DataMovementListResponse)

_DATAMOVER = DESCRIPTOR.services_by_name['DataMover']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\035com.hpe.cray.nnf.datamovementB\021DataMovementProtoP\001Z!nnf.cray.hpe.com/datamovement/api'
  _DATAMOVEMENTCREATEREQUEST._serialized_start=36
  _DATAMOVEMENTCREATEREQUEST._serialized_end=153
  _DATAMOVEMENTCREATERESPONSE._serialized_start=156
  _DATAMOVEMENTCREATERESPONSE._serialized_end=314
  _DATAMOVEMENTCREATERESPONSE_STATUS._serialized_start=281
  _DATAMOVEMENTCREATERESPONSE_STATUS._serialized_end=314
  _DATAMOVEMENTSTATUSREQUEST._serialized_start=316
  _DATAMOVEMENTSTATUSREQUEST._serialized_end=377
  _DATAMOVEMENTSTATUSRESPONSE._serialized_start=380
  _DATAMOVEMENTSTATUSRESPONSE._serialized_end=719
  _DATAMOVEMENTSTATUSRESPONSE_STATE._serialized_start=555
  _DATAMOVEMENTSTATUSRESPONSE_STATE._serialized_end=636
  _DATAMOVEMENTSTATUSRESPONSE_STATUS._serialized_start=638
  _DATAMOVEMENTSTATUSRESPONSE_STATUS._serialized_end=719
  _DATAMOVEMENTDELETEREQUEST._serialized_start=721
  _DATAMOVEMENTDELETEREQUEST._serialized_end=761
  _DATAMOVEMENTDELETERESPONSE._serialized_start=764
  _DATAMOVEMENTDELETERESPONSE._serialized_end=933
  _DATAMOVEMENTDELETERESPONSE_STATUS._serialized_start=859
  _DATAMOVEMENTDELETERESPONSE_STATUS._serialized_end=933
  _DATAMOVEMENTLISTREQUEST._serialized_start=935
  _DATAMOVEMENTLISTREQUEST._serialized_end=997
  _DATAMOVEMENTLISTRESPONSE._serialized_start=999
  _DATAMOVEMENTLISTRESPONSE._serialized_end=1039
  _DATAMOVER._serialized_start=1042
  _DATAMOVER._serialized_end=1427
# @@protoc_insertion_point(module_scope)
