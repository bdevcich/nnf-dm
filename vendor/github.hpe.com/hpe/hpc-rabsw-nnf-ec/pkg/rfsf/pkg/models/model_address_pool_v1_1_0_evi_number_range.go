/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// AddressPoolV110EviNumberRange - The Ethernet Virtual Private Network (EVPN) Instance (EVI) number range for an Ethernet fabric.
type AddressPoolV110EviNumberRange struct {

	// Lower Ethernet Virtual Private Network (EVPN) Instance (EVI) number.
	Lower int64 `json:"Lower,omitempty"`

	// Upper Ethernet Virtual Private Network (EVPN) Instance (EVI) number.
	Upper int64 `json:"Upper,omitempty"`
}