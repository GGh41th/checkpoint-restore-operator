/*
Copyright 2024.

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

// +kubebuilder:validation:Enum:=Pending;InProgress;Ready;Failed
type ContainerCheckpointPhase string

const (
	ContainerCheckpointPending    ContainerCheckpointPhase = "Pending"
	ContainerCheckPointInProgress ContainerCheckpointPhase = "InProgress"
	ContainerCheckpointReady      ContainerCheckpointPhase = "Ready"
	ContainerCheckpointFailed     ContainerCheckpointPhase = "Failed"
)

const (
	ConditionReady = "Ready"
)

type Condition struct {
	Type   string
	Status metav1.ConditionStatus
}

type ResourceThreshold struct {
	cpu    string
	memory string
}

type CheckpointPolicy struct {
	// +optional
	Schedule string `json:"schedule,omitempty"`
	// +optional
	Resources ResourceThreshold `json:"resources,omitempty"`
	// +optional
	NodeConditions []Condition `json:"nodeConditions,omitempty"`
	// +optional
	OnDrain bool `json:"onDrain,omitempty"`
}

type ContainerCheckpointSpec struct {
	// +required
	// +kubebuilder:validation:MinLength:=1
	// +kubebuilder:validation:XValidation:rule="self==oldSelf",message="podName is immutable"
	PodName string `json:"podname"`
	// +optional
	// +kubebuilder:validation:XValidation:rule="self==oldSelf",message="ContainerName is immutable"
	ContainerName string `json:"containerName,omitempty"`
	// +required
	// +kubebuilder:validation:XValidation:rule="self==oldSelf",message="Namespace is immutable"
	// +default:value="default"
	NameSpace string `json:"namespace,omitempty"`

	// Compression field is yet to be discussed
	// Compression string `json:"compression,omitempty"`

	// A map of mutually exclusive checkpointing policies.
	// +optional
	Policy CheckpointPolicy
}

type ContainerCheckpointStatus struct {
	// +optional
	Phase ContainerCheckpointPhase `json:"phase,omitempty"`
	// +optional
	CheckPointPath string `json:"checkpointPath,omitempty"`
	// +optional
	NodeName string `json:"nodeName,omitempty"`
	// The status of each condition is one of True, False, or Unknown.
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ContainerCheckpoint is the Schema for the containercheckpoints API
type ContainerCheckpoint struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`

	// spec defines the desired state of ContainerCheckpoint
	// +required
	Spec ContainerCheckpointSpec `json:"spec"`

	// status defines the observed state of ContainerCheckpoint
	// +optional
	Status ContainerCheckpointStatus `json:"status,omitzero"`
}

// +kubebuilder:object:root=true

// ContainerCheckpointList contains a list of ContainerCheckpoint
type ContainerCheckpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []ContainerCheckpoint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ContainerCheckpoint{}, &ContainerCheckpointList{})
}
