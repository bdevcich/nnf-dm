/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// TelemetryServiceV121Actions - The available actions for this resource.
type TelemetryServiceV121Actions struct {

	TelemetryServiceSubmitTestMetricReport TelemetryServiceV121SubmitTestMetricReport `json:"#TelemetryService.SubmitTestMetricReport,omitempty"`

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}
