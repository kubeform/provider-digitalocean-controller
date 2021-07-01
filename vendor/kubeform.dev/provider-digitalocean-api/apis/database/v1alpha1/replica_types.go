/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Replica struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReplicaSpec   `json:"spec,omitempty"`
	Status            ReplicaStatus `json:"status,omitempty"`
}

type ReplicaSpec struct {
	KubeformOutput *ReplicaSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource ReplicaSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type ReplicaSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ClusterID *string `json:"clusterID" tf:"cluster_id"`
	// +optional
	Database *string `json:"database,omitempty" tf:"database"`
	// +optional
	Host *string `json:"host,omitempty" tf:"host"`
	Name *string `json:"name" tf:"name"`
	// +optional
	Password *string `json:"-" sensitive:"true" tf:"password"`
	// +optional
	Port *int64 `json:"port,omitempty" tf:"port"`
	// +optional
	PrivateHost *string `json:"privateHost,omitempty" tf:"private_host"`
	// +optional
	PrivateNetworkUUID *string `json:"privateNetworkUUID,omitempty" tf:"private_network_uuid"`
	// +optional
	PrivateURI *string `json:"-" sensitive:"true" tf:"private_uri"`
	// +optional
	Region *string `json:"region,omitempty" tf:"region"`
	// +optional
	Size *string `json:"size,omitempty" tf:"size"`
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
	// +optional
	Uri *string `json:"-" sensitive:"true" tf:"uri"`
	// +optional
	User *string `json:"user,omitempty" tf:"user"`
}

type ReplicaStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ReplicaList is a list of Replicas
type ReplicaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Replica CRD objects
	Items []Replica `json:"items,omitempty"`
}
