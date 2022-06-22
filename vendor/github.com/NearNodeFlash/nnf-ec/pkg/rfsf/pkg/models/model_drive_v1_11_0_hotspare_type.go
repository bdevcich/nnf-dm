/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type DriveV1110HotspareType string

// List of Drive_v1_11_0_HotspareType
const (
	NONE_DV1110HT DriveV1110HotspareType = "None"
	GLOBAL_DV1110HT DriveV1110HotspareType = "Global"
	CHASSIS_DV1110HT DriveV1110HotspareType = "Chassis"
	DEDICATED_DV1110HT DriveV1110HotspareType = "Dedicated"
)