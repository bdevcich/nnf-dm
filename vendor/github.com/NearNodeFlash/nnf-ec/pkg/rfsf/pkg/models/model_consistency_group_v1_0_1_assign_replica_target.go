/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ConsistencyGroupV101AssignReplicaTarget - This action is used to establish a replication relationship by assigning an existing consistency group to serve as a target replica for an existing source consistency group.
type ConsistencyGroupV101AssignReplicaTarget struct {

	// Link to invoke action
	Target string `json:"target,omitempty"`

	// Friendly action name
	Title string `json:"title,omitempty"`
}