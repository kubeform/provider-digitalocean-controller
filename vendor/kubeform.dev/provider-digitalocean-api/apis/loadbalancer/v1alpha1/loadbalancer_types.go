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

type Loadbalancer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoadbalancerSpec   `json:"spec,omitempty"`
	Status            LoadbalancerStatus `json:"status,omitempty"`
}

type LoadbalancerSpecForwardingRule struct {
	// +optional
	// Deprecated
	CertificateID *string `json:"certificateID,omitempty" tf:"certificate_id"`
	// +optional
	CertificateName *string `json:"certificateName,omitempty" tf:"certificate_name"`
	EntryPort       *int64  `json:"entryPort" tf:"entry_port"`
	EntryProtocol   *string `json:"entryProtocol" tf:"entry_protocol"`
	TargetPort      *int64  `json:"targetPort" tf:"target_port"`
	TargetProtocol  *string `json:"targetProtocol" tf:"target_protocol"`
	// +optional
	TlsPassthrough *bool `json:"tlsPassthrough,omitempty" tf:"tls_passthrough"`
}

type LoadbalancerSpecHealthcheck struct {
	// +optional
	CheckIntervalSeconds *int64 `json:"checkIntervalSeconds,omitempty" tf:"check_interval_seconds"`
	// +optional
	HealthyThreshold *int64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold"`
	// +optional
	Path     *string `json:"path,omitempty" tf:"path"`
	Port     *int64  `json:"port" tf:"port"`
	Protocol *string `json:"protocol" tf:"protocol"`
	// +optional
	ResponseTimeoutSeconds *int64 `json:"responseTimeoutSeconds,omitempty" tf:"response_timeout_seconds"`
	// +optional
	UnhealthyThreshold *int64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold"`
}

type LoadbalancerSpecStickySessions struct {
	// +optional
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name"`
	// +optional
	CookieTtlSeconds *int64 `json:"cookieTtlSeconds,omitempty" tf:"cookie_ttl_seconds"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type LoadbalancerSpec struct {
	State *LoadbalancerSpecResource `json:"state,omitempty" tf:"-"`

	Resource LoadbalancerSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type LoadbalancerSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm"`
	// +optional
	DropletIDS []int64 `json:"dropletIDS,omitempty" tf:"droplet_ids"`
	// +optional
	DropletTag *string `json:"dropletTag,omitempty" tf:"droplet_tag"`
	// +optional
	EnableBackendKeepalive *bool `json:"enableBackendKeepalive,omitempty" tf:"enable_backend_keepalive"`
	// +optional
	EnableProxyProtocol *bool `json:"enableProxyProtocol,omitempty" tf:"enable_proxy_protocol"`
	// +kubebuilder:validation:MinItems=1
	ForwardingRule []LoadbalancerSpecForwardingRule `json:"forwardingRule" tf:"forwarding_rule"`
	// +optional
	Healthcheck *LoadbalancerSpecHealthcheck `json:"healthcheck,omitempty" tf:"healthcheck"`
	// +optional
	Ip   *string `json:"ip,omitempty" tf:"ip"`
	Name *string `json:"name" tf:"name"`
	// +optional
	RedirectHTTPToHTTPS *bool   `json:"redirectHTTPToHTTPS,omitempty" tf:"redirect_http_to_https"`
	Region              *string `json:"region" tf:"region"`
	// +optional
	Size *string `json:"size,omitempty" tf:"size"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	StickySessions *LoadbalancerSpecStickySessions `json:"stickySessions,omitempty" tf:"sticky_sessions"`
	// the uniform resource name for the load balancer
	// +optional
	Urn *string `json:"urn,omitempty" tf:"urn"`
	// +optional
	VpcUUID *string `json:"vpcUUID,omitempty" tf:"vpc_uuid"`
}

type LoadbalancerStatus struct {
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

// LoadbalancerList is a list of Loadbalancers
type LoadbalancerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Loadbalancer CRD objects
	Items []Loadbalancer `json:"items,omitempty"`
}
