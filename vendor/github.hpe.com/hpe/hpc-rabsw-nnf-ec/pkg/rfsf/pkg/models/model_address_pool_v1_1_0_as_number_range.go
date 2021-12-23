/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// AddressPoolV110AsNumberRange - Autonomous System (AS) number range.
type AddressPoolV110AsNumberRange struct {

	// Lower Autonomous System (AS) number.
	Lower int64 `json:"Lower,omitempty"`

	// Upper Autonomous System (AS) number.
	Upper int64 `json:"Upper,omitempty"`
}
