/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RsyncNodeDataMovementSpec defines the desired state of RsyncNodeDataMovement
type RsyncNodeDataMovementSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of RsyncNodeDataMovement. Edit rsyncnodedatamovement_types.go to remove/update
	Id string `json:"id,omitempty"`

	Source string `json:"source,omitempty"`

	Destination string `json:"destination,omitempty"`

	UserId uint32 `json:"userId,omitempty"`

	GroupId uint32 `json:"groupId,omitempty"`

	DryRun bool `json:"dryRun,omitempty"`
}

//
type RsyncDataMovementState string

const (
	RsyncDataMovementStatusState_Starting  RsyncDataMovementState = "Starting"
	RsyncDataMovementStatusState_Running   RsyncDataMovementState = "Running"
	RsyncDataMovementStatusState_Completed RsyncDataMovementState = "Completed"
	RsyncDataMovementStatusState_Failed    RsyncDataMovementState = "Failed"
)

// RsyncNodeDataMovementStatus defines the observed state of RsyncNodeDataMovement
type RsyncNodeDataMovementStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Current State of the data movement request
	State RsyncDataMovementState `json:"state,omitempty"`

	// Status describes the helper message that accompanies the state
	Status string `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// RsyncNodeDataMovement is the Schema for the rsyncnodedatamovements API
type RsyncNodeDataMovement struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RsyncNodeDataMovementSpec   `json:"spec,omitempty"`
	Status RsyncNodeDataMovementStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RsyncNodeDataMovementList contains a list of RsyncNodeDataMovement
type RsyncNodeDataMovementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RsyncNodeDataMovement `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RsyncNodeDataMovement{}, &RsyncNodeDataMovementList{})
}