/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// EventServiceV170Smtp - Settings for SMTP event delivery.
type EventServiceV170Smtp struct {

	Authentication EventServiceV170SMTPAuthenticationMethods `json:"Authentication,omitempty"`

	ConnectionProtocol EventServiceV170SMTPConnectionProtocol `json:"ConnectionProtocol,omitempty"`

	// The 'from' email address of the outgoing email.
	FromAddress string `json:"FromAddress,omitempty"`

	// The password for authentication with the SMTP server.  The value is `null` in responses.
	Password string `json:"Password,omitempty"`

	// The destination SMTP port.
	Port int64 `json:"Port,omitempty"`

	// The address of the SMTP server.
	ServerAddress string `json:"ServerAddress,omitempty"`

	// An indication if SMTP for event delivery is enabled.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

	// The username for authentication with the SMTP server.
	Username string `json:"Username,omitempty"`
}
