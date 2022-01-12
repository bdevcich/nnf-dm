/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// StoragePoolV150NvMeSetProperties - This contains properties to use when StoragePool is used to describe an NVMe Set.
type StoragePoolV150NvMeSetProperties struct {

	// A 16-bit hex value that contains the endurance group identifier.
	EnduranceGroupIdentifier string `json:"EnduranceGroupIdentifier,omitempty"`

	// This property contains the Optimal Write Size in Bytes for this NVMe Set.
	OptimalWriteSizeBytes int64 `json:"OptimalWriteSizeBytes,omitempty"`

	// Indicates the typical time to complete a 4k read in 100 nano-second units when the NVM Set is in a Predictable Latency Mode Deterministic Window and there is 1 outstanding command per NVM Set.
	Random4kReadTypicalNanoSeconds int64 `json:"Random4kReadTypicalNanoSeconds,omitempty"`

	// A 16-bit hex value that contains the NVMe Set group identifier.
	SetIdentifier string `json:"SetIdentifier,omitempty"`

	// Indicates the unallocated capacity of the NVMe Set in bytes.
	UnallocatedNVMNamespaceCapacityBytes int64 `json:"UnallocatedNVMNamespaceCapacityBytes,omitempty"`
}