/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// StorageReplicaInfoV120ReplicaState : Values of ReplicaState describe the state of the relationship with respect to Replication activity.
type StorageReplicaInfoV120ReplicaState string

// List of StorageReplicaInfo_v1_2_0_ReplicaState
const (
	INITIALIZED_SRIV120RST StorageReplicaInfoV120ReplicaState = "Initialized"
	UNSYNCHRONIZED_SRIV120RST StorageReplicaInfoV120ReplicaState = "Unsynchronized"
	SYNCHRONIZED_SRIV120RST StorageReplicaInfoV120ReplicaState = "Synchronized"
	BROKEN_SRIV120RST StorageReplicaInfoV120ReplicaState = "Broken"
	FRACTURED_SRIV120RST StorageReplicaInfoV120ReplicaState = "Fractured"
	SPLIT_SRIV120RST StorageReplicaInfoV120ReplicaState = "Split"
	INACTIVE_SRIV120RST StorageReplicaInfoV120ReplicaState = "Inactive"
	SUSPENDED_SRIV120RST StorageReplicaInfoV120ReplicaState = "Suspended"
	FAILEDOVER_SRIV120RST StorageReplicaInfoV120ReplicaState = "Failedover"
	PREPARED_SRIV120RST StorageReplicaInfoV120ReplicaState = "Prepared"
	ABORTED_SRIV120RST StorageReplicaInfoV120ReplicaState = "Aborted"
	SKEWED_SRIV120RST StorageReplicaInfoV120ReplicaState = "Skewed"
	MIXED_SRIV120RST StorageReplicaInfoV120ReplicaState = "Mixed"
	PARTITIONED_SRIV120RST StorageReplicaInfoV120ReplicaState = "Partitioned"
	INVALID_SRIV120RST StorageReplicaInfoV120ReplicaState = "Invalid"
	RESTORED_SRIV120RST StorageReplicaInfoV120ReplicaState = "Restored"
)
