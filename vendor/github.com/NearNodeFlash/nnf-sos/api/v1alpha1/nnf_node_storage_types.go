/*
 * Copyright 2021-2023 Hewlett Packard Enterprise Development LP
 * Other additional copyright holders may be indicated within.
 *
 * The entirety of this work is licensed under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 *
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package v1alpha1

import (
	dwsv1alpha2 "github.com/DataWorkflowServices/dws/api/v1alpha2"
	"github.com/DataWorkflowServices/dws/utils/updater"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// IMPORTANT: Run "make" to regenerate code after modifying this file
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NnfNodeStorageSpec defines the desired storage attributes on a NNF Node.
// Storage spec are created on bequest of the user and fullfilled by the NNF Node Controller.
type NnfNodeStorageSpec struct {
	// Count is the number of allocations to make on this node. All of the allocations will
	// be created with the same parameters
	// +kubebuilder:validation:Minimum:=0
	Count int `json:"count"`

	// User ID for file system
	UserID uint32 `json:"userID"`

	// Group ID for file system
	GroupID uint32 `json:"groupID"`

	// Set the owner and group permissions specified by UserID and GroupID. This is for
	// Lustre file systems only, and should be set only after all Lustre targets are created.
	// +kubebuilder:default:=false
	SetOwnerGroup bool `json:"setOwnerGroup"`

	// Capacity defines the capacity, in bytes, of this storage specification. The NNF Node itself
	// may split the storage among the available drives operating in the NNF Node.
	Capacity int64 `json:"capacity,omitempty"`

	// FileSystemType defines the type of the desired filesystem, or raw
	// block device.
	// +kubebuilder:validation:Enum=raw;lvm;zfs;xfs;gfs2;lustre
	// +kubebuilder:default:=raw
	FileSystemType string `json:"fileSystemType"`

	// LustreStorageSpec describes the Lustre target created here, if
	// FileSystemType specifies a Lustre target.
	LustreStorage LustreStorageSpec `json:"lustreStorage,omitempty"`

	// ClientEndpoints sets which endpoints should have access to an allocation.
	ClientEndpoints []ClientEndpointsSpec `json:"clientEndpoints"`
}

// ClientEndpointsSpec contains information about which nodes a storage allocation
// should be visible to
type ClientEndpointsSpec struct {
	// Index of the allocation in the NnfNodeStorage
	AllocationIndex int `json:"allocationIndex"`

	// List of nodes that should see the allocation
	NodeNames []string `json:"nodeNames"`
}

// LustreStorageSpec describes the Lustre target to be created here.
type LustreStorageSpec struct {
	// FileSystemName is the fsname parameter for the Lustre filesystem.
	// +kubebuilder:validation:MaxLength:=8
	FileSystemName string `json:"fileSystemName,omitempty"`

	// TargetType is the type of Lustre target to be created.
	// +kubebuilder:validation:Enum=MGT;MDT;MGTMDT;OST
	TargetType string `json:"targetType,omitempty"`

	// StartIndex is used to order a series of MDTs or OSTs.  This is used only
	// when creating MDT and OST targets. If count in the NnfNodeStorageSpec is more
	// than 1, then StartIndex is the index of the first allocation, and the indexes
	// increment from there.
	// +kubebuilder:validation:Minimum:=0
	StartIndex int `json:"startIndex,omitempty"`

	// MgsNode is the NID of the MGS to use. This is used only when
	// creating MDT and OST targets.
	MgsNode string `json:"mgsNode,omitempty"`

	// BackFs is the type of backing filesystem to use.
	// +kubebuilder:validation:Enum=ldiskfs;zfs
	BackFs string `json:"backFs,omitempty"`
}

// NnfNodeStorageStatus defines the status for NnfNodeStorage
type NnfNodeStorageStatus struct {
	// Allocations is the list of storage allocations that were made
	Allocations []NnfNodeStorageAllocationStatus `json:"allocations,omitempty"`

	dwsv1alpha2.ResourceError `json:",inline"`

	// LustreStorageStatus describes the Lustre targets created here.
	LustreStorage LustreStorageStatus `json:"lustreStorage,omitempty"`

	// OwnerGroupStatus is the status of the operation for setting the owner and group
	// of a file system
	OwnerGroupStatus NnfResourceStatusType `json:"ownerGroupStatus,omitempty"`
}

// NnfNodeStorageNVMeStatus provides a way to uniquely identify an NVMe namespace
// in the system
type NnfNodeStorageNVMeStatus struct {
	// Serial number of the base NVMe device
	DeviceSerial string `json:"deviceSerial"`

	// Id of the Namespace on the NVMe device (e.g., "2")
	NamespaceID string `json:"namespaceID"`

	// Globally unique namespace ID
	NamespaceGUID string `json:"namespaceGUID"`
}

// NnfNodeStorageAllocationStatus defines the allocation status for each allocation in the NnfNodeStorage
type NnfNodeStorageAllocationStatus struct {
	// Represents the time when the storage was created by the controller
	// It is represented in RFC3339 form and is in UTC.
	CreationTime *metav1.Time `json:"creationTime,omitempty"`

	// Represents the time when the storage was deleted by the controller. This field
	// is updated when the Storage specification State transitions to 'Delete' by the
	// client.
	// It is represented in RFC3339 form and is in UTC.
	DeletionTime *metav1.Time `json:"deletionTime,omitempty"`

	// Total capacity allocated for the storage. This may differ from the requested storage
	// capacity as the system may round up to the requested capacity satisify underlying
	// storage requirements (i.e. block size / stripe size).
	CapacityAllocated int64 `json:"capacityAllocated,omitempty"`

	// Represents the storage group that is supporting this server. A storage group is
	// the mapping from a group of drive namespaces to an individual server. This value
	// can be safely ignored by the client.
	StorageGroup NnfResourceStatus `json:"storageGroup,omitempty"`

	// Name of the LVM VG
	VolumeGroup string `json:"volumeGroup,omitempty"`

	// Name of the LVM LV
	LogicalVolume string `json:"logicalVolume,omitempty"`

	// List of NVMe namespaces used by this allocation
	NVMeList []NnfNodeStorageNVMeStatus `json:"nvmeList,omitempty"`

	// Represents the file share that is supporting this server. A file share is the
	// combination of a storage group and the associated file system parameters (type, mountpoint)
	// that makes up the available storage.
	FileShare NnfResourceStatus `json:"fileShare,omitempty"`

	StoragePool NnfResourceStatus `json:"storagePool,omitempty"`

	FileSystem NnfResourceStatus `json:"fileSystem,omitempty"`
}

// LustreStorageStatus describes the Lustre target created here.
type LustreStorageStatus struct {

	// Nid (LNet Network Identifier) of this node. This is populated on MGS nodes only.
	Nid string `json:"nid,omitempty"`
}

//+kubebuilder:object:root=true

// NnfNodeStorage is the Schema for the NnfNodeStorage API
type NnfNodeStorage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NnfNodeStorageSpec   `json:"spec,omitempty"`
	Status NnfNodeStorageStatus `json:"status,omitempty"`
}

func (ns *NnfNodeStorage) GetStatus() updater.Status[*NnfNodeStorageStatus] {
	return &ns.Status
}

//+kubebuilder:object:root=true
//+kubebuilder:printcolumn:name="ERROR",type="string",JSONPath=".status.error.severity"
//+kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"

// NnfNodeStorageList contains a list of NNF Nodes
type NnfNodeStorageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NnfNodeStorage `json:"items"`
}

func (n *NnfNodeStorageList) GetObjectList() []client.Object {
	objectList := []client.Object{}

	for i := range n.Items {
		objectList = append(objectList, &n.Items[i])
	}

	return objectList
}

func init() {
	SchemeBuilder.Register(&NnfNodeStorage{}, &NnfNodeStorageList{})
}
