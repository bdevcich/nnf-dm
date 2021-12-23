/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type VirtualMediaV132TransferProtocolType string

// List of VirtualMedia_v1_3_2_TransferProtocolType
const (
	CIFS_VMV132TPT VirtualMediaV132TransferProtocolType = "CIFS"
	FTP_VMV132TPT VirtualMediaV132TransferProtocolType = "FTP"
	SFTP_VMV132TPT VirtualMediaV132TransferProtocolType = "SFTP"
	HTTP_VMV132TPT VirtualMediaV132TransferProtocolType = "HTTP"
	HTTPS_VMV132TPT VirtualMediaV132TransferProtocolType = "HTTPS"
	NFS_VMV132TPT VirtualMediaV132TransferProtocolType = "NFS"
	SCP_VMV132TPT VirtualMediaV132TransferProtocolType = "SCP"
	TFTP_VMV132TPT VirtualMediaV132TransferProtocolType = "TFTP"
	OEM_VMV132TPT VirtualMediaV132TransferProtocolType = "OEM"
)
