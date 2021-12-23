/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// DataSecurityLoSCapabilitiesV120DataSecurityLoSCapabilities - Describe data security capabilities.
type DataSecurityLoSCapabilitiesV120DataSecurityLoSCapabilities struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions DataSecurityLoSCapabilitiesV120Actions `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	Identifier ResourceIdentifier `json:"Identifier,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// Supported AntiVirus providers.
	SupportedAntivirusEngineProviders []string `json:"SupportedAntivirusEngineProviders,omitempty"`

	// Supported policies that trigger an AntiVirus scan.
	SupportedAntivirusScanPolicies []DataSecurityLoSCapabilitiesAntiVirusScanTrigger `json:"SupportedAntivirusScanPolicies,omitempty"`

	// Supported key sizes for transport channel encryption.
	SupportedChannelEncryptionStrengths []DataSecurityLoSCapabilitiesKeySize `json:"SupportedChannelEncryptionStrengths,omitempty"`

	// Supported data sanitization policies.
	SupportedDataSanitizationPolicies []DataSecurityLoSCapabilitiesDataSanitizationPolicy `json:"SupportedDataSanitizationPolicies,omitempty"`

	// Supported authentication types for hosts (servers) or initiator endpoints.
	SupportedHostAuthenticationTypes []DataSecurityLoSCapabilitiesAuthenticationType `json:"SupportedHostAuthenticationTypes,omitempty"`

	// Collection of known and supported DataSecurityLinesOfService.
	SupportedLinesOfService []OdataV4IdRef `json:"SupportedLinesOfService,omitempty"`

	// The number of items in a collection.
	SupportedLinesOfServiceodataCount int64 `json:"SupportedLinesOfService@odata.count,omitempty"`

	// Supported key sizes for media encryption.
	SupportedMediaEncryptionStrengths []DataSecurityLoSCapabilitiesKeySize `json:"SupportedMediaEncryptionStrengths,omitempty"`

	// Supported protocols that provide encrypted communication.
	SupportedSecureChannelProtocols []DataSecurityLoSCapabilitiesSecureChannelProtocol `json:"SupportedSecureChannelProtocols,omitempty"`

	// Supported authentication types for users (or programs).
	SupportedUserAuthenticationTypes []DataSecurityLoSCapabilitiesAuthenticationType `json:"SupportedUserAuthenticationTypes,omitempty"`
}
