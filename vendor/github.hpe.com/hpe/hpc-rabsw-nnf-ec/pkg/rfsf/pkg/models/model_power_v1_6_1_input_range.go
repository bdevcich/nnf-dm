/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// PowerV161InputRange - This type describes an input range for a power supply.
type PowerV161InputRange struct {

	InputType PowerV161InputType `json:"InputType,omitempty"`

	// The maximum line input frequency at which this power supply input range is effective.
	MaximumFrequencyHz *float32 `json:"MaximumFrequencyHz,omitempty"`

	// The maximum line input voltage at which this power supply input range is effective.
	MaximumVoltage *float32 `json:"MaximumVoltage,omitempty"`

	// The minimum line input frequency at which this power supply input range is effective.
	MinimumFrequencyHz *float32 `json:"MinimumFrequencyHz,omitempty"`

	// The minimum line input voltage at which this power supply input range is effective.
	MinimumVoltage *float32 `json:"MinimumVoltage,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The maximum capacity of this power supply when operating in this input range.
	OutputWattage *float32 `json:"OutputWattage,omitempty"`
}
