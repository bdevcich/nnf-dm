/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// BootOptionV104BootOption - The BootOption schema reports information about a single boot option in a system.  It represents the properties of a bootable device available in the system.
type BootOptionV104BootOption struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions BootOptionV104Actions `json:"Actions,omitempty"`

	Alias ComputerSystemBootSource `json:"Alias,omitempty"`

	// An indication of whether the boot option is enabled.  If `true`, it is enabled.  If `false`, the boot option that the boot order array on the computer system contains is skipped.  In the UEFI context, this property shall influence the load option active flag for the boot option.
	BootOptionEnabled bool `json:"BootOptionEnabled,omitempty"`

	// The unique boot option.
	BootOptionReference string `json:"BootOptionReference"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The user-readable display name of the boot option that appears in the boot order list in the user interface.
	DisplayName string `json:"DisplayName,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// An array of links to resources or objects associated with this boot option.
	RelatedItem []OdataV4IdRef `json:"RelatedItem,omitempty"`

	// The number of items in a collection.
	RelatedItemodataCount int64 `json:"RelatedItem@odata.count,omitempty"`

	// The UEFI device path to access this UEFI boot option.
	UefiDevicePath string `json:"UefiDevicePath,omitempty"`
}