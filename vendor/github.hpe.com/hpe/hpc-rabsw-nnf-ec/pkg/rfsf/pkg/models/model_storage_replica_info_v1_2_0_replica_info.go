/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// StorageReplicaInfoV120ReplicaInfo - Defines the characteristics of a replica.
type StorageReplicaInfoV120ReplicaInfo struct {

	// True if consistency is enabled.
	ConsistencyEnabled bool `json:"ConsistencyEnabled,omitempty"`

	ConsistencyState StorageReplicaInfoV120ConsistencyState `json:"ConsistencyState,omitempty"`

	ConsistencyStatus StorageReplicaInfoV120ConsistencyStatus `json:"ConsistencyStatus,omitempty"`

	ConsistencyType StorageReplicaInfoV120ConsistencyType `json:"ConsistencyType,omitempty"`

	DataProtectionLineOfService OdataV4IdRef `json:"DataProtectionLineOfService,omitempty"`

	// If true, the storage array tells host to stop sending data to source element if copying to a remote element fails.
	FailedCopyStopsHostIO bool `json:"FailedCopyStopsHostIO,omitempty"`

	// Specifies the percent of the work completed to reach synchronization.
	PercentSynced int64 `json:"PercentSynced,omitempty"`

	Replica OdataV4IdRef `json:"Replica,omitempty"`

	ReplicaPriority StorageReplicaInfoV120ReplicaPriority `json:"ReplicaPriority,omitempty"`

	ReplicaProgressStatus StorageReplicaInfoV120ReplicaProgressStatus `json:"ReplicaProgressStatus,omitempty"`

	ReplicaReadOnlyAccess StorageReplicaInfoV120ReplicaReadOnlyAccess `json:"ReplicaReadOnlyAccess,omitempty"`

	ReplicaRecoveryMode StorageReplicaInfoV120ReplicaRecoveryMode `json:"ReplicaRecoveryMode,omitempty"`

	ReplicaRole StorageReplicaInfoV120ReplicaRole `json:"ReplicaRole,omitempty"`

	// Applies to Adaptive mode and it describes maximum number of bytes the SyncedElement (target) can be out of sync.
	ReplicaSkewBytes int64 `json:"ReplicaSkewBytes,omitempty"`

	ReplicaState StorageReplicaInfoV120ReplicaState `json:"ReplicaState,omitempty"`

	ReplicaType StorageReplicaInfoReplicaType `json:"ReplicaType,omitempty"`

	ReplicaUpdateMode StorageReplicaInfoReplicaUpdateMode `json:"ReplicaUpdateMode,omitempty"`

	RequestedReplicaState StorageReplicaInfoV120ReplicaState `json:"RequestedReplicaState,omitempty"`

	SourceReplica OdataV4IdRef `json:"SourceReplica,omitempty"`

	// Synchronization is maintained.
	SyncMaintained bool `json:"SyncMaintained,omitempty"`

	UndiscoveredElement StorageReplicaInfoV120UndiscoveredElement `json:"UndiscoveredElement,omitempty"`

	// Specifies when point-in-time copy was taken or when the replication relationship is activated, reactivated, resumed or re-established.
	WhenActivated string `json:"WhenActivated,omitempty"`

	// Specifies when the replication relationship is deactivated.
	WhenDeactivated string `json:"WhenDeactivated,omitempty"`

	// Specifies when the replication relationship is established.
	WhenEstablished string `json:"WhenEstablished,omitempty"`

	// Specifies when the replication relationship is suspended.
	WhenSuspended string `json:"WhenSuspended,omitempty"`

	// The point in time that the Elements were synchronized.
	WhenSynced string `json:"WhenSynced,omitempty"`

	// Specifies when the replication relationship is synchronized.
	WhenSynchronized string `json:"WhenSynchronized,omitempty"`
}
