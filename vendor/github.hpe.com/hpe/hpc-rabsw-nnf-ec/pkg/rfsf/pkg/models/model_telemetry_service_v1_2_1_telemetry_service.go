/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// TelemetryServiceV121TelemetryService - The TelemetryService schema describes a telemetry service.  The telemetry service is used to for collecting and reporting metric data within the Redfish Service.
type TelemetryServiceV121TelemetryService struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions TelemetryServiceV121Actions `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	LogService OdataV4IdRef `json:"LogService,omitempty"`

	// The maximum number of metric reports that this service supports.
	MaxReports int64 `json:"MaxReports,omitempty"`

	MetricDefinitions OdataV4IdRef `json:"MetricDefinitions,omitempty"`

	MetricReportDefinitions OdataV4IdRef `json:"MetricReportDefinitions,omitempty"`

	MetricReports OdataV4IdRef `json:"MetricReports,omitempty"`

	// The minimum time interval between gathering metric data that this service allows.
	MinCollectionInterval string `json:"MinCollectionInterval,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// An indication of whether this service is enabled.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`

	// The functions that can be performed over each metric.
	SupportedCollectionFunctions []TelemetryServiceV121CollectionFunction `json:"SupportedCollectionFunctions,omitempty"`

	Triggers OdataV4IdRef `json:"Triggers,omitempty"`
}