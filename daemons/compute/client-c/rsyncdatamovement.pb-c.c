/* Generated by the protocol buffer compiler.  DO NOT EDIT! */
/* Generated from: rsyncdatamovement.proto */

/* Do not generate deprecated warnings for self */
#ifndef PROTOBUF_C__NO_DEPRECATED
#define PROTOBUF_C__NO_DEPRECATED
#endif

#include "rsyncdatamovement.pb-c.h"
void   datamovement__rsync_data_movement_create_request__init
                     (Datamovement__RsyncDataMovementCreateRequest         *message)
{
  static const Datamovement__RsyncDataMovementCreateRequest init_value = DATAMOVEMENT__RSYNC_DATA_MOVEMENT_CREATE_REQUEST__INIT;
  *message = init_value;
}
size_t datamovement__rsync_data_movement_create_request__get_packed_size
                     (const Datamovement__RsyncDataMovementCreateRequest *message)
{
  assert(message->base.descriptor == &datamovement__rsync_data_movement_create_request__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t datamovement__rsync_data_movement_create_request__pack
                     (const Datamovement__RsyncDataMovementCreateRequest *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &datamovement__rsync_data_movement_create_request__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t datamovement__rsync_data_movement_create_request__pack_to_buffer
                     (const Datamovement__RsyncDataMovementCreateRequest *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &datamovement__rsync_data_movement_create_request__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Datamovement__RsyncDataMovementCreateRequest *
       datamovement__rsync_data_movement_create_request__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Datamovement__RsyncDataMovementCreateRequest *)
     protobuf_c_message_unpack (&datamovement__rsync_data_movement_create_request__descriptor,
                                allocator, len, data);
}
void   datamovement__rsync_data_movement_create_request__free_unpacked
                     (Datamovement__RsyncDataMovementCreateRequest *message,
                      ProtobufCAllocator *allocator)
{
  if(!message)
    return;
  assert(message->base.descriptor == &datamovement__rsync_data_movement_create_request__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
void   datamovement__rsync_data_movement_create_response__init
                     (Datamovement__RsyncDataMovementCreateResponse         *message)
{
  static const Datamovement__RsyncDataMovementCreateResponse init_value = DATAMOVEMENT__RSYNC_DATA_MOVEMENT_CREATE_RESPONSE__INIT;
  *message = init_value;
}
size_t datamovement__rsync_data_movement_create_response__get_packed_size
                     (const Datamovement__RsyncDataMovementCreateResponse *message)
{
  assert(message->base.descriptor == &datamovement__rsync_data_movement_create_response__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t datamovement__rsync_data_movement_create_response__pack
                     (const Datamovement__RsyncDataMovementCreateResponse *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &datamovement__rsync_data_movement_create_response__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t datamovement__rsync_data_movement_create_response__pack_to_buffer
                     (const Datamovement__RsyncDataMovementCreateResponse *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &datamovement__rsync_data_movement_create_response__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Datamovement__RsyncDataMovementCreateResponse *
       datamovement__rsync_data_movement_create_response__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Datamovement__RsyncDataMovementCreateResponse *)
     protobuf_c_message_unpack (&datamovement__rsync_data_movement_create_response__descriptor,
                                allocator, len, data);
}
void   datamovement__rsync_data_movement_create_response__free_unpacked
                     (Datamovement__RsyncDataMovementCreateResponse *message,
                      ProtobufCAllocator *allocator)
{
  if(!message)
    return;
  assert(message->base.descriptor == &datamovement__rsync_data_movement_create_response__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
void   datamovement__rsync_data_movement_status_request__init
                     (Datamovement__RsyncDataMovementStatusRequest         *message)
{
  static const Datamovement__RsyncDataMovementStatusRequest init_value = DATAMOVEMENT__RSYNC_DATA_MOVEMENT_STATUS_REQUEST__INIT;
  *message = init_value;
}
size_t datamovement__rsync_data_movement_status_request__get_packed_size
                     (const Datamovement__RsyncDataMovementStatusRequest *message)
{
  assert(message->base.descriptor == &datamovement__rsync_data_movement_status_request__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t datamovement__rsync_data_movement_status_request__pack
                     (const Datamovement__RsyncDataMovementStatusRequest *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &datamovement__rsync_data_movement_status_request__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t datamovement__rsync_data_movement_status_request__pack_to_buffer
                     (const Datamovement__RsyncDataMovementStatusRequest *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &datamovement__rsync_data_movement_status_request__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Datamovement__RsyncDataMovementStatusRequest *
       datamovement__rsync_data_movement_status_request__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Datamovement__RsyncDataMovementStatusRequest *)
     protobuf_c_message_unpack (&datamovement__rsync_data_movement_status_request__descriptor,
                                allocator, len, data);
}
void   datamovement__rsync_data_movement_status_request__free_unpacked
                     (Datamovement__RsyncDataMovementStatusRequest *message,
                      ProtobufCAllocator *allocator)
{
  if(!message)
    return;
  assert(message->base.descriptor == &datamovement__rsync_data_movement_status_request__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
void   datamovement__rsync_data_movement_status_response__init
                     (Datamovement__RsyncDataMovementStatusResponse         *message)
{
  static const Datamovement__RsyncDataMovementStatusResponse init_value = DATAMOVEMENT__RSYNC_DATA_MOVEMENT_STATUS_RESPONSE__INIT;
  *message = init_value;
}
size_t datamovement__rsync_data_movement_status_response__get_packed_size
                     (const Datamovement__RsyncDataMovementStatusResponse *message)
{
  assert(message->base.descriptor == &datamovement__rsync_data_movement_status_response__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t datamovement__rsync_data_movement_status_response__pack
                     (const Datamovement__RsyncDataMovementStatusResponse *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &datamovement__rsync_data_movement_status_response__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t datamovement__rsync_data_movement_status_response__pack_to_buffer
                     (const Datamovement__RsyncDataMovementStatusResponse *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &datamovement__rsync_data_movement_status_response__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Datamovement__RsyncDataMovementStatusResponse *
       datamovement__rsync_data_movement_status_response__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Datamovement__RsyncDataMovementStatusResponse *)
     protobuf_c_message_unpack (&datamovement__rsync_data_movement_status_response__descriptor,
                                allocator, len, data);
}
void   datamovement__rsync_data_movement_status_response__free_unpacked
                     (Datamovement__RsyncDataMovementStatusResponse *message,
                      ProtobufCAllocator *allocator)
{
  if(!message)
    return;
  assert(message->base.descriptor == &datamovement__rsync_data_movement_status_response__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
void   datamovement__rsync_data_movement_delete_request__init
                     (Datamovement__RsyncDataMovementDeleteRequest         *message)
{
  static const Datamovement__RsyncDataMovementDeleteRequest init_value = DATAMOVEMENT__RSYNC_DATA_MOVEMENT_DELETE_REQUEST__INIT;
  *message = init_value;
}
size_t datamovement__rsync_data_movement_delete_request__get_packed_size
                     (const Datamovement__RsyncDataMovementDeleteRequest *message)
{
  assert(message->base.descriptor == &datamovement__rsync_data_movement_delete_request__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t datamovement__rsync_data_movement_delete_request__pack
                     (const Datamovement__RsyncDataMovementDeleteRequest *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &datamovement__rsync_data_movement_delete_request__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t datamovement__rsync_data_movement_delete_request__pack_to_buffer
                     (const Datamovement__RsyncDataMovementDeleteRequest *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &datamovement__rsync_data_movement_delete_request__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Datamovement__RsyncDataMovementDeleteRequest *
       datamovement__rsync_data_movement_delete_request__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Datamovement__RsyncDataMovementDeleteRequest *)
     protobuf_c_message_unpack (&datamovement__rsync_data_movement_delete_request__descriptor,
                                allocator, len, data);
}
void   datamovement__rsync_data_movement_delete_request__free_unpacked
                     (Datamovement__RsyncDataMovementDeleteRequest *message,
                      ProtobufCAllocator *allocator)
{
  if(!message)
    return;
  assert(message->base.descriptor == &datamovement__rsync_data_movement_delete_request__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
void   datamovement__rsync_data_movement_delete_response__init
                     (Datamovement__RsyncDataMovementDeleteResponse         *message)
{
  static const Datamovement__RsyncDataMovementDeleteResponse init_value = DATAMOVEMENT__RSYNC_DATA_MOVEMENT_DELETE_RESPONSE__INIT;
  *message = init_value;
}
size_t datamovement__rsync_data_movement_delete_response__get_packed_size
                     (const Datamovement__RsyncDataMovementDeleteResponse *message)
{
  assert(message->base.descriptor == &datamovement__rsync_data_movement_delete_response__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t datamovement__rsync_data_movement_delete_response__pack
                     (const Datamovement__RsyncDataMovementDeleteResponse *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &datamovement__rsync_data_movement_delete_response__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t datamovement__rsync_data_movement_delete_response__pack_to_buffer
                     (const Datamovement__RsyncDataMovementDeleteResponse *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &datamovement__rsync_data_movement_delete_response__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Datamovement__RsyncDataMovementDeleteResponse *
       datamovement__rsync_data_movement_delete_response__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Datamovement__RsyncDataMovementDeleteResponse *)
     protobuf_c_message_unpack (&datamovement__rsync_data_movement_delete_response__descriptor,
                                allocator, len, data);
}
void   datamovement__rsync_data_movement_delete_response__free_unpacked
                     (Datamovement__RsyncDataMovementDeleteResponse *message,
                      ProtobufCAllocator *allocator)
{
  if(!message)
    return;
  assert(message->base.descriptor == &datamovement__rsync_data_movement_delete_response__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
static const ProtobufCFieldDescriptor datamovement__rsync_data_movement_create_request__field_descriptors[7] =
{
  {
    "initiator",
    1,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Datamovement__RsyncDataMovementCreateRequest, initiator),
    NULL,
    &protobuf_c_empty_string,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "target",
    2,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Datamovement__RsyncDataMovementCreateRequest, target),
    NULL,
    &protobuf_c_empty_string,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "workflow",
    3,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Datamovement__RsyncDataMovementCreateRequest, workflow),
    NULL,
    &protobuf_c_empty_string,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "namespace",
    4,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Datamovement__RsyncDataMovementCreateRequest, namespace_),
    NULL,
    &protobuf_c_empty_string,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "source",
    5,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Datamovement__RsyncDataMovementCreateRequest, source),
    NULL,
    &protobuf_c_empty_string,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "destination",
    6,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Datamovement__RsyncDataMovementCreateRequest, destination),
    NULL,
    &protobuf_c_empty_string,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "dryrun",
    7,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_BOOL,
    0,   /* quantifier_offset */
    offsetof(Datamovement__RsyncDataMovementCreateRequest, dryrun),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned datamovement__rsync_data_movement_create_request__field_indices_by_name[] = {
  5,   /* field[5] = destination */
  6,   /* field[6] = dryrun */
  0,   /* field[0] = initiator */
  3,   /* field[3] = namespace */
  4,   /* field[4] = source */
  1,   /* field[1] = target */
  2,   /* field[2] = workflow */
};
static const ProtobufCIntRange datamovement__rsync_data_movement_create_request__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 7 }
};
const ProtobufCMessageDescriptor datamovement__rsync_data_movement_create_request__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "datamovement.RsyncDataMovementCreateRequest",
  "RsyncDataMovementCreateRequest",
  "Datamovement__RsyncDataMovementCreateRequest",
  "datamovement",
  sizeof(Datamovement__RsyncDataMovementCreateRequest),
  7,
  datamovement__rsync_data_movement_create_request__field_descriptors,
  datamovement__rsync_data_movement_create_request__field_indices_by_name,
  1,  datamovement__rsync_data_movement_create_request__number_ranges,
  (ProtobufCMessageInit) datamovement__rsync_data_movement_create_request__init,
  NULL,NULL,NULL    /* reserved[123] */
};
static const ProtobufCFieldDescriptor datamovement__rsync_data_movement_create_response__field_descriptors[1] =
{
  {
    "uid",
    1,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Datamovement__RsyncDataMovementCreateResponse, uid),
    NULL,
    &protobuf_c_empty_string,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned datamovement__rsync_data_movement_create_response__field_indices_by_name[] = {
  0,   /* field[0] = uid */
};
static const ProtobufCIntRange datamovement__rsync_data_movement_create_response__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 1 }
};
const ProtobufCMessageDescriptor datamovement__rsync_data_movement_create_response__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "datamovement.RsyncDataMovementCreateResponse",
  "RsyncDataMovementCreateResponse",
  "Datamovement__RsyncDataMovementCreateResponse",
  "datamovement",
  sizeof(Datamovement__RsyncDataMovementCreateResponse),
  1,
  datamovement__rsync_data_movement_create_response__field_descriptors,
  datamovement__rsync_data_movement_create_response__field_indices_by_name,
  1,  datamovement__rsync_data_movement_create_response__number_ranges,
  (ProtobufCMessageInit) datamovement__rsync_data_movement_create_response__init,
  NULL,NULL,NULL    /* reserved[123] */
};
static const ProtobufCFieldDescriptor datamovement__rsync_data_movement_status_request__field_descriptors[1] =
{
  {
    "uid",
    1,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Datamovement__RsyncDataMovementStatusRequest, uid),
    NULL,
    &protobuf_c_empty_string,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned datamovement__rsync_data_movement_status_request__field_indices_by_name[] = {
  0,   /* field[0] = uid */
};
static const ProtobufCIntRange datamovement__rsync_data_movement_status_request__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 1 }
};
const ProtobufCMessageDescriptor datamovement__rsync_data_movement_status_request__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "datamovement.RsyncDataMovementStatusRequest",
  "RsyncDataMovementStatusRequest",
  "Datamovement__RsyncDataMovementStatusRequest",
  "datamovement",
  sizeof(Datamovement__RsyncDataMovementStatusRequest),
  1,
  datamovement__rsync_data_movement_status_request__field_descriptors,
  datamovement__rsync_data_movement_status_request__field_indices_by_name,
  1,  datamovement__rsync_data_movement_status_request__number_ranges,
  (ProtobufCMessageInit) datamovement__rsync_data_movement_status_request__init,
  NULL,NULL,NULL    /* reserved[123] */
};
static const ProtobufCEnumValue datamovement__rsync_data_movement_status_response__state__enum_values_by_number[5] =
{
  { "PENDING", "DATAMOVEMENT__RSYNC_DATA_MOVEMENT_STATUS_RESPONSE__STATE__PENDING", 0 },
  { "STARTING", "DATAMOVEMENT__RSYNC_DATA_MOVEMENT_STATUS_RESPONSE__STATE__STARTING", 1 },
  { "RUNNING", "DATAMOVEMENT__RSYNC_DATA_MOVEMENT_STATUS_RESPONSE__STATE__RUNNING", 2 },
  { "COMPLETED", "DATAMOVEMENT__RSYNC_DATA_MOVEMENT_STATUS_RESPONSE__STATE__COMPLETED", 3 },
  { "UNKNOWN_STATE", "DATAMOVEMENT__RSYNC_DATA_MOVEMENT_STATUS_RESPONSE__STATE__UNKNOWN_STATE", 4 },
};
static const ProtobufCIntRange datamovement__rsync_data_movement_status_response__state__value_ranges[] = {
{0, 0},{0, 5}
};
static const ProtobufCEnumValueIndex datamovement__rsync_data_movement_status_response__state__enum_values_by_name[5] =
{
  { "COMPLETED", 3 },
  { "PENDING", 0 },
  { "RUNNING", 2 },
  { "STARTING", 1 },
  { "UNKNOWN_STATE", 4 },
};
const ProtobufCEnumDescriptor datamovement__rsync_data_movement_status_response__state__descriptor =
{
  PROTOBUF_C__ENUM_DESCRIPTOR_MAGIC,
  "datamovement.RsyncDataMovementStatusResponse.State",
  "State",
  "Datamovement__RsyncDataMovementStatusResponse__State",
  "datamovement",
  5,
  datamovement__rsync_data_movement_status_response__state__enum_values_by_number,
  5,
  datamovement__rsync_data_movement_status_response__state__enum_values_by_name,
  1,
  datamovement__rsync_data_movement_status_response__state__value_ranges,
  NULL,NULL,NULL,NULL   /* reserved[1234] */
};
static const ProtobufCEnumValue datamovement__rsync_data_movement_status_response__status__enum_values_by_number[5] =
{
  { "INVALID", "DATAMOVEMENT__RSYNC_DATA_MOVEMENT_STATUS_RESPONSE__STATUS__INVALID", 0 },
  { "NOT_FOUND", "DATAMOVEMENT__RSYNC_DATA_MOVEMENT_STATUS_RESPONSE__STATUS__NOT_FOUND", 1 },
  { "SUCCESS", "DATAMOVEMENT__RSYNC_DATA_MOVEMENT_STATUS_RESPONSE__STATUS__SUCCESS", 2 },
  { "FAILED", "DATAMOVEMENT__RSYNC_DATA_MOVEMENT_STATUS_RESPONSE__STATUS__FAILED", 3 },
  { "UNKNOWN_STATUS", "DATAMOVEMENT__RSYNC_DATA_MOVEMENT_STATUS_RESPONSE__STATUS__UNKNOWN_STATUS", 4 },
};
static const ProtobufCIntRange datamovement__rsync_data_movement_status_response__status__value_ranges[] = {
{0, 0},{0, 5}
};
static const ProtobufCEnumValueIndex datamovement__rsync_data_movement_status_response__status__enum_values_by_name[5] =
{
  { "FAILED", 3 },
  { "INVALID", 0 },
  { "NOT_FOUND", 1 },
  { "SUCCESS", 2 },
  { "UNKNOWN_STATUS", 4 },
};
const ProtobufCEnumDescriptor datamovement__rsync_data_movement_status_response__status__descriptor =
{
  PROTOBUF_C__ENUM_DESCRIPTOR_MAGIC,
  "datamovement.RsyncDataMovementStatusResponse.Status",
  "Status",
  "Datamovement__RsyncDataMovementStatusResponse__Status",
  "datamovement",
  5,
  datamovement__rsync_data_movement_status_response__status__enum_values_by_number,
  5,
  datamovement__rsync_data_movement_status_response__status__enum_values_by_name,
  1,
  datamovement__rsync_data_movement_status_response__status__value_ranges,
  NULL,NULL,NULL,NULL   /* reserved[1234] */
};
static const ProtobufCFieldDescriptor datamovement__rsync_data_movement_status_response__field_descriptors[3] =
{
  {
    "state",
    1,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_ENUM,
    0,   /* quantifier_offset */
    offsetof(Datamovement__RsyncDataMovementStatusResponse, state),
    &datamovement__rsync_data_movement_status_response__state__descriptor,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "status",
    2,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_ENUM,
    0,   /* quantifier_offset */
    offsetof(Datamovement__RsyncDataMovementStatusResponse, status),
    &datamovement__rsync_data_movement_status_response__status__descriptor,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "message",
    3,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Datamovement__RsyncDataMovementStatusResponse, message),
    NULL,
    &protobuf_c_empty_string,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned datamovement__rsync_data_movement_status_response__field_indices_by_name[] = {
  2,   /* field[2] = message */
  0,   /* field[0] = state */
  1,   /* field[1] = status */
};
static const ProtobufCIntRange datamovement__rsync_data_movement_status_response__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 3 }
};
const ProtobufCMessageDescriptor datamovement__rsync_data_movement_status_response__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "datamovement.RsyncDataMovementStatusResponse",
  "RsyncDataMovementStatusResponse",
  "Datamovement__RsyncDataMovementStatusResponse",
  "datamovement",
  sizeof(Datamovement__RsyncDataMovementStatusResponse),
  3,
  datamovement__rsync_data_movement_status_response__field_descriptors,
  datamovement__rsync_data_movement_status_response__field_indices_by_name,
  1,  datamovement__rsync_data_movement_status_response__number_ranges,
  (ProtobufCMessageInit) datamovement__rsync_data_movement_status_response__init,
  NULL,NULL,NULL    /* reserved[123] */
};
static const ProtobufCFieldDescriptor datamovement__rsync_data_movement_delete_request__field_descriptors[1] =
{
  {
    "uid",
    1,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Datamovement__RsyncDataMovementDeleteRequest, uid),
    NULL,
    &protobuf_c_empty_string,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned datamovement__rsync_data_movement_delete_request__field_indices_by_name[] = {
  0,   /* field[0] = uid */
};
static const ProtobufCIntRange datamovement__rsync_data_movement_delete_request__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 1 }
};
const ProtobufCMessageDescriptor datamovement__rsync_data_movement_delete_request__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "datamovement.RsyncDataMovementDeleteRequest",
  "RsyncDataMovementDeleteRequest",
  "Datamovement__RsyncDataMovementDeleteRequest",
  "datamovement",
  sizeof(Datamovement__RsyncDataMovementDeleteRequest),
  1,
  datamovement__rsync_data_movement_delete_request__field_descriptors,
  datamovement__rsync_data_movement_delete_request__field_indices_by_name,
  1,  datamovement__rsync_data_movement_delete_request__number_ranges,
  (ProtobufCMessageInit) datamovement__rsync_data_movement_delete_request__init,
  NULL,NULL,NULL    /* reserved[123] */
};
static const ProtobufCEnumValue datamovement__rsync_data_movement_delete_response__status__enum_values_by_number[5] =
{
  { "INVALID", "DATAMOVEMENT__RSYNC_DATA_MOVEMENT_DELETE_RESPONSE__STATUS__INVALID", 0 },
  { "NOT_FOUND", "DATAMOVEMENT__RSYNC_DATA_MOVEMENT_DELETE_RESPONSE__STATUS__NOT_FOUND", 1 },
  { "DELETED", "DATAMOVEMENT__RSYNC_DATA_MOVEMENT_DELETE_RESPONSE__STATUS__DELETED", 2 },
  { "ACTIVE", "DATAMOVEMENT__RSYNC_DATA_MOVEMENT_DELETE_RESPONSE__STATUS__ACTIVE", 3 },
  { "UNKNOWN", "DATAMOVEMENT__RSYNC_DATA_MOVEMENT_DELETE_RESPONSE__STATUS__UNKNOWN", 4 },
};
static const ProtobufCIntRange datamovement__rsync_data_movement_delete_response__status__value_ranges[] = {
{0, 0},{0, 5}
};
static const ProtobufCEnumValueIndex datamovement__rsync_data_movement_delete_response__status__enum_values_by_name[5] =
{
  { "ACTIVE", 3 },
  { "DELETED", 2 },
  { "INVALID", 0 },
  { "NOT_FOUND", 1 },
  { "UNKNOWN", 4 },
};
const ProtobufCEnumDescriptor datamovement__rsync_data_movement_delete_response__status__descriptor =
{
  PROTOBUF_C__ENUM_DESCRIPTOR_MAGIC,
  "datamovement.RsyncDataMovementDeleteResponse.Status",
  "Status",
  "Datamovement__RsyncDataMovementDeleteResponse__Status",
  "datamovement",
  5,
  datamovement__rsync_data_movement_delete_response__status__enum_values_by_number,
  5,
  datamovement__rsync_data_movement_delete_response__status__enum_values_by_name,
  1,
  datamovement__rsync_data_movement_delete_response__status__value_ranges,
  NULL,NULL,NULL,NULL   /* reserved[1234] */
};
static const ProtobufCFieldDescriptor datamovement__rsync_data_movement_delete_response__field_descriptors[1] =
{
  {
    "status",
    1,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_ENUM,
    0,   /* quantifier_offset */
    offsetof(Datamovement__RsyncDataMovementDeleteResponse, status),
    &datamovement__rsync_data_movement_delete_response__status__descriptor,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned datamovement__rsync_data_movement_delete_response__field_indices_by_name[] = {
  0,   /* field[0] = status */
};
static const ProtobufCIntRange datamovement__rsync_data_movement_delete_response__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 1 }
};
const ProtobufCMessageDescriptor datamovement__rsync_data_movement_delete_response__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "datamovement.RsyncDataMovementDeleteResponse",
  "RsyncDataMovementDeleteResponse",
  "Datamovement__RsyncDataMovementDeleteResponse",
  "datamovement",
  sizeof(Datamovement__RsyncDataMovementDeleteResponse),
  1,
  datamovement__rsync_data_movement_delete_response__field_descriptors,
  datamovement__rsync_data_movement_delete_response__field_indices_by_name,
  1,  datamovement__rsync_data_movement_delete_response__number_ranges,
  (ProtobufCMessageInit) datamovement__rsync_data_movement_delete_response__init,
  NULL,NULL,NULL    /* reserved[123] */
};
static const ProtobufCMethodDescriptor datamovement__rsync_data_mover__method_descriptors[3] =
{
  { "Create", &datamovement__rsync_data_movement_create_request__descriptor, &datamovement__rsync_data_movement_create_response__descriptor },
  { "Status", &datamovement__rsync_data_movement_status_request__descriptor, &datamovement__rsync_data_movement_status_response__descriptor },
  { "Delete", &datamovement__rsync_data_movement_delete_request__descriptor, &datamovement__rsync_data_movement_delete_response__descriptor },
};
const unsigned datamovement__rsync_data_mover__method_indices_by_name[] = {
  0,        /* Create */
  2,        /* Delete */
  1         /* Status */
};
const ProtobufCServiceDescriptor datamovement__rsync_data_mover__descriptor =
{
  PROTOBUF_C__SERVICE_DESCRIPTOR_MAGIC,
  "datamovement.RsyncDataMover",
  "RsyncDataMover",
  "Datamovement__RsyncDataMover",
  "datamovement",
  3,
  datamovement__rsync_data_mover__method_descriptors,
  datamovement__rsync_data_mover__method_indices_by_name
};
void datamovement__rsync_data_mover__create(ProtobufCService *service,
                                            const Datamovement__RsyncDataMovementCreateRequest *input,
                                            Datamovement__RsyncDataMovementCreateResponse_Closure closure,
                                            void *closure_data)
{
  assert(service->descriptor == &datamovement__rsync_data_mover__descriptor);
  service->invoke(service, 0, (const ProtobufCMessage *) input, (ProtobufCClosure) closure, closure_data);
}
void datamovement__rsync_data_mover__status(ProtobufCService *service,
                                            const Datamovement__RsyncDataMovementStatusRequest *input,
                                            Datamovement__RsyncDataMovementStatusResponse_Closure closure,
                                            void *closure_data)
{
  assert(service->descriptor == &datamovement__rsync_data_mover__descriptor);
  service->invoke(service, 1, (const ProtobufCMessage *) input, (ProtobufCClosure) closure, closure_data);
}
void datamovement__rsync_data_mover__delete(ProtobufCService *service,
                                            const Datamovement__RsyncDataMovementDeleteRequest *input,
                                            Datamovement__RsyncDataMovementDeleteResponse_Closure closure,
                                            void *closure_data)
{
  assert(service->descriptor == &datamovement__rsync_data_mover__descriptor);
  service->invoke(service, 2, (const ProtobufCMessage *) input, (ProtobufCClosure) closure, closure_data);
}
void datamovement__rsync_data_mover__init (Datamovement__RsyncDataMover_Service *service,
                                           Datamovement__RsyncDataMover_ServiceDestroy destroy)
{
  protobuf_c_service_generated_init (&service->base,
                                     &datamovement__rsync_data_mover__descriptor,
                                     (ProtobufCServiceDestroy) destroy);
}
