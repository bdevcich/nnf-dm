/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type NVMeDomainNVMeDeviceType string

// List of NVMeDomain_NVMeDeviceType
const (
	DRIVE_NVMDNVMDT NVMeDomainNVMeDeviceType = "Drive"
	JBOF_NVMDNVMDT NVMeDomainNVMeDeviceType = "JBOF"
	FABRIC_ATTACH_ARRAY_NVMDNVMDT NVMeDomainNVMeDeviceType = "FabricAttachArray"
)