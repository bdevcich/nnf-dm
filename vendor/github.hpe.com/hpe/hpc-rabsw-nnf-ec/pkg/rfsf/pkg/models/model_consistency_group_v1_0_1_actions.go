/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ConsistencyGroupV101Actions - The available actions for this resource.
type ConsistencyGroupV101Actions struct {

	ConsistencyGroupAssignReplicaTarget ConsistencyGroupV101AssignReplicaTarget `json:"#ConsistencyGroup.AssignReplicaTarget,omitempty"`

	ConsistencyGroupCreateReplicaTarget ConsistencyGroupV101CreateReplicaTarget `json:"#ConsistencyGroup.CreateReplicaTarget,omitempty"`

	ConsistencyGroupRemoveReplicaRelationship ConsistencyGroupV101RemoveReplicaRelationship `json:"#ConsistencyGroup.RemoveReplicaRelationship,omitempty"`

	ConsistencyGroupResumeReplication ConsistencyGroupV101ResumeReplication `json:"#ConsistencyGroup.ResumeReplication,omitempty"`

	ConsistencyGroupReverseReplicationRelationship ConsistencyGroupV101ReverseReplicationRelationship `json:"#ConsistencyGroup.ReverseReplicationRelationship,omitempty"`

	ConsistencyGroupSplitReplication ConsistencyGroupV101SplitReplication `json:"#ConsistencyGroup.SplitReplication,omitempty"`

	ConsistencyGroupSuspendReplication ConsistencyGroupV101SuspendReplication `json:"#ConsistencyGroup.SuspendReplication,omitempty"`

	// The available OEM specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}
