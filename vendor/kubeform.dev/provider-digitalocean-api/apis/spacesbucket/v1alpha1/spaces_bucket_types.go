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

type SpacesBucket struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpacesBucketSpec   `json:"spec,omitempty"`
	Status            SpacesBucketStatus `json:"status,omitempty"`
}

type SpacesBucketSpecCorsRule struct {
	// A list of headers that will be included in the CORS preflight request's Access-Control-Request-Headers. A header may contain one wildcard (e.g. x-amz-*).
	// +optional
	AllowedHeaders []string `json:"allowedHeaders,omitempty" tf:"allowed_headers"`
	// A list of HTTP methods (e.g. GET) which are allowed from the specified origin.
	AllowedMethods []string `json:"allowedMethods" tf:"allowed_methods"`
	// A list of hosts from which requests using the specified methods are allowed. A host may contain one wildcard (e.g. http://*.example.com).
	AllowedOrigins []string `json:"allowedOrigins" tf:"allowed_origins"`
	// +optional
	MaxAgeSeconds *int64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds"`
}

type SpacesBucketSpecLifecycleRuleExpiration struct {
	// +optional
	Date *string `json:"date,omitempty" tf:"date"`
	// +optional
	Days *int64 `json:"days,omitempty" tf:"days"`
	// +optional
	ExpiredObjectDeleteMarker *bool `json:"expiredObjectDeleteMarker,omitempty" tf:"expired_object_delete_marker"`
}

type SpacesBucketSpecLifecycleRuleNoncurrentVersionExpiration struct {
	// +optional
	Days *int64 `json:"days,omitempty" tf:"days"`
}

type SpacesBucketSpecLifecycleRule struct {
	// +optional
	AbortIncompleteMultipartUploadDays *int64 `json:"abortIncompleteMultipartUploadDays,omitempty" tf:"abort_incomplete_multipart_upload_days"`
	Enabled                            *bool  `json:"enabled" tf:"enabled"`
	// +optional
	Expiration *SpacesBucketSpecLifecycleRuleExpiration `json:"expiration,omitempty" tf:"expiration"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	NoncurrentVersionExpiration *SpacesBucketSpecLifecycleRuleNoncurrentVersionExpiration `json:"noncurrentVersionExpiration,omitempty" tf:"noncurrent_version_expiration"`
	// +optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix"`
}

type SpacesBucketSpecVersioning struct {
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
}

type SpacesBucketSpec struct {
	State *SpacesBucketSpecResource `json:"state,omitempty" tf:"-"`

	Resource SpacesBucketSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type SpacesBucketSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Canned ACL applied on bucket creation
	// +optional
	Acl *string `json:"acl,omitempty" tf:"acl"`
	// The FQDN of the bucket
	// +optional
	BucketDomainName *string `json:"bucketDomainName,omitempty" tf:"bucket_domain_name"`
	// A container holding a list of elements describing allowed methods for a specific origin.
	// +optional
	CorsRule []SpacesBucketSpecCorsRule `json:"corsRule,omitempty" tf:"cors_rule"`
	// Unless true, the bucket will only be destroyed if empty
	// +optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy"`
	// +optional
	LifecycleRule []SpacesBucketSpecLifecycleRule `json:"lifecycleRule,omitempty" tf:"lifecycle_rule"`
	// Bucket name
	Name *string `json:"name" tf:"name"`
	// Bucket region
	// +optional
	Region *string `json:"region,omitempty" tf:"region"`
	// the uniform resource name for the bucket
	// +optional
	Urn *string `json:"urn,omitempty" tf:"urn"`
	// +optional
	Versioning *SpacesBucketSpecVersioning `json:"versioning,omitempty" tf:"versioning"`
}

type SpacesBucketStatus struct {
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

// SpacesBucketList is a list of SpacesBuckets
type SpacesBucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SpacesBucket CRD objects
	Items []SpacesBucket `json:"items,omitempty"`
}
