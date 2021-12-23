/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// StorageReplicaInfoV102ConsistencyState : The values of ConsistencyState indicate the consistency type used by the source and its associated target group.
type StorageReplicaInfoV102ConsistencyState string

// List of StorageReplicaInfo_v1_0_2_ConsistencyState
const (
	CONSISTENT_SRIV102CST StorageReplicaInfoV102ConsistencyState = "Consistent"
	INCONSISTENT_SRIV102CST StorageReplicaInfoV102ConsistencyState = "Inconsistent"
)
