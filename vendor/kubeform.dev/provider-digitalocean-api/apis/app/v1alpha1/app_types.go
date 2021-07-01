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

type App struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppSpec   `json:"spec,omitempty"`
	Status            AppStatus `json:"status,omitempty"`
}

type AppSpecSpecDatabase struct {
	// The name of the underlying DigitalOcean DBaaS cluster. This is required for production databases. For dev databases, if cluster_name is not set, a new cluster will be provisioned.
	// +optional
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name"`
	// The name of the MySQL or PostgreSQL database to configure.
	// +optional
	DbName *string `json:"dbName,omitempty" tf:"db_name"`
	// The name of the MySQL or PostgreSQL user to configure.
	// +optional
	DbUser *string `json:"dbUser,omitempty" tf:"db_user"`
	// The database engine to use.
	// +optional
	Engine *string `json:"engine,omitempty" tf:"engine"`
	// The name of the component
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// Whether this is a production or dev database.
	// +optional
	Production *bool `json:"production,omitempty" tf:"production"`
	// The version of the database engine.
	// +optional
	Version *string `json:"version,omitempty" tf:"version"`
}

type AppSpecSpecDomain struct {
	// The hostname for the domain.
	Name *string `json:"name" tf:"name"`
	// The type of the domain.
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
	// Indicates whether the domain includes all sub-domains, in addition to the given domain.
	// +optional
	Wildcard *bool `json:"wildcard,omitempty" tf:"wildcard"`
	// If the domain uses DigitalOcean DNS and you would like App Platform to automatically manage it for you, set this to the name of the domain on your account.
	// +optional
	Zone *string `json:"zone,omitempty" tf:"zone"`
}

type AppSpecSpecEnv struct {
	// The name of the environment variable.
	// +optional
	Key *string `json:"key,omitempty" tf:"key"`
	// The visibility scope of the environment variable.
	// +optional
	Scope *string `json:"scope,omitempty" tf:"scope"`
	// The type of the environment variable.
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
	// The value of the environment variable.
	// +optional
	Value *string `json:"-" sensitive:"true" tf:"value"`
}

type AppSpecSpecJobEnv struct {
	// The name of the environment variable.
	// +optional
	Key *string `json:"key,omitempty" tf:"key"`
	// The visibility scope of the environment variable.
	// +optional
	Scope *string `json:"scope,omitempty" tf:"scope"`
	// The type of the environment variable.
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
	// The value of the environment variable.
	// +optional
	Value *string `json:"-" sensitive:"true" tf:"value"`
}

type AppSpecSpecJobGit struct {
	// The name of the branch to use.
	// +optional
	Branch *string `json:"branch,omitempty" tf:"branch"`
	// The clone URL of the repo.
	// +optional
	RepoCloneURL *string `json:"repoCloneURL,omitempty" tf:"repo_clone_url"`
}

type AppSpecSpecJobGithub struct {
	// The name of the branch to use.
	// +optional
	Branch *string `json:"branch,omitempty" tf:"branch"`
	// Whether to automatically deploy new commits made to the repo
	// +optional
	DeployOnPush *bool `json:"deployOnPush,omitempty" tf:"deploy_on_push"`
	// The name of the repo in the format `owner/repo`.
	// +optional
	Repo *string `json:"repo,omitempty" tf:"repo"`
}

type AppSpecSpecJobGitlab struct {
	// The name of the branch to use.
	// +optional
	Branch *string `json:"branch,omitempty" tf:"branch"`
	// Whether to automatically deploy new commits made to the repo
	// +optional
	DeployOnPush *bool `json:"deployOnPush,omitempty" tf:"deploy_on_push"`
	// The name of the repo in the format `owner/repo`.
	// +optional
	Repo *string `json:"repo,omitempty" tf:"repo"`
}

type AppSpecSpecJobImage struct {
	// The registry name. Must be left empty for the DOCR registry type.
	// +optional
	Registry *string `json:"registry,omitempty" tf:"registry"`
	// The registry type.
	RegistryType *string `json:"registryType" tf:"registry_type"`
	// The repository name.
	Repository *string `json:"repository" tf:"repository"`
	// The repository tag. Defaults to latest if not provided.
	// +optional
	Tag *string `json:"tag,omitempty" tf:"tag"`
}

type AppSpecSpecJob struct {
	// An optional build command to run while building this component from source.
	// +optional
	BuildCommand *string `json:"buildCommand,omitempty" tf:"build_command"`
	// The path to a Dockerfile relative to the root of the repo. If set, overrides usage of buildpacks.
	// +optional
	DockerfilePath *string `json:"dockerfilePath,omitempty" tf:"dockerfile_path"`
	// +optional
	Env []AppSpecSpecJobEnv `json:"env,omitempty" tf:"env"`
	// An environment slug describing the type of this app.
	// +optional
	EnvironmentSlug *string `json:"environmentSlug,omitempty" tf:"environment_slug"`
	// +optional
	Git *AppSpecSpecJobGit `json:"git,omitempty" tf:"git"`
	// +optional
	Github *AppSpecSpecJobGithub `json:"github,omitempty" tf:"github"`
	// +optional
	Gitlab *AppSpecSpecJobGitlab `json:"gitlab,omitempty" tf:"gitlab"`
	// +optional
	Image *AppSpecSpecJobImage `json:"image,omitempty" tf:"image"`
	// The amount of instances that this component should be scaled to.
	// +optional
	InstanceCount *int64 `json:"instanceCount,omitempty" tf:"instance_count"`
	// The instance size to use for this component.
	// +optional
	InstanceSizeSlug *string `json:"instanceSizeSlug,omitempty" tf:"instance_size_slug"`
	// The type of job and when it will be run during the deployment process.
	// +optional
	Kind *string `json:"kind,omitempty" tf:"kind"`
	// The name of the component
	Name *string `json:"name" tf:"name"`
	// An optional run command to override the component's default.
	// +optional
	RunCommand *string `json:"runCommand,omitempty" tf:"run_command"`
	// An optional path to the working directory to use for the build.
	// +optional
	SourceDir *string `json:"sourceDir,omitempty" tf:"source_dir"`
}

type AppSpecSpecServiceEnv struct {
	// The name of the environment variable.
	// +optional
	Key *string `json:"key,omitempty" tf:"key"`
	// The visibility scope of the environment variable.
	// +optional
	Scope *string `json:"scope,omitempty" tf:"scope"`
	// The type of the environment variable.
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
	// The value of the environment variable.
	// +optional
	Value *string `json:"-" sensitive:"true" tf:"value"`
}

type AppSpecSpecServiceGit struct {
	// The name of the branch to use.
	// +optional
	Branch *string `json:"branch,omitempty" tf:"branch"`
	// The clone URL of the repo.
	// +optional
	RepoCloneURL *string `json:"repoCloneURL,omitempty" tf:"repo_clone_url"`
}

type AppSpecSpecServiceGithub struct {
	// The name of the branch to use.
	// +optional
	Branch *string `json:"branch,omitempty" tf:"branch"`
	// Whether to automatically deploy new commits made to the repo
	// +optional
	DeployOnPush *bool `json:"deployOnPush,omitempty" tf:"deploy_on_push"`
	// The name of the repo in the format `owner/repo`.
	// +optional
	Repo *string `json:"repo,omitempty" tf:"repo"`
}

type AppSpecSpecServiceGitlab struct {
	// The name of the branch to use.
	// +optional
	Branch *string `json:"branch,omitempty" tf:"branch"`
	// Whether to automatically deploy new commits made to the repo
	// +optional
	DeployOnPush *bool `json:"deployOnPush,omitempty" tf:"deploy_on_push"`
	// The name of the repo in the format `owner/repo`.
	// +optional
	Repo *string `json:"repo,omitempty" tf:"repo"`
}

type AppSpecSpecServiceHealthCheck struct {
	// The number of failed health checks before considered unhealthy.
	// +optional
	FailureThreshold *int64 `json:"failureThreshold,omitempty" tf:"failure_threshold"`
	// The route path used for the HTTP health check ping.
	// +optional
	HttpPath *string `json:"httpPath,omitempty" tf:"http_path"`
	// The number of seconds to wait before beginning health checks.
	// +optional
	InitialDelaySeconds *int64 `json:"initialDelaySeconds,omitempty" tf:"initial_delay_seconds"`
	// The number of seconds to wait between health checks.
	// +optional
	PeriodSeconds *int64 `json:"periodSeconds,omitempty" tf:"period_seconds"`
	// The number of successful health checks before considered healthy.
	// +optional
	SuccessThreshold *int64 `json:"successThreshold,omitempty" tf:"success_threshold"`
	// The number of seconds after which the check times out.
	// +optional
	TimeoutSeconds *int64 `json:"timeoutSeconds,omitempty" tf:"timeout_seconds"`
}

type AppSpecSpecServiceImage struct {
	// The registry name. Must be left empty for the DOCR registry type.
	// +optional
	Registry *string `json:"registry,omitempty" tf:"registry"`
	// The registry type.
	RegistryType *string `json:"registryType" tf:"registry_type"`
	// The repository name.
	Repository *string `json:"repository" tf:"repository"`
	// The repository tag. Defaults to latest if not provided.
	// +optional
	Tag *string `json:"tag,omitempty" tf:"tag"`
}

type AppSpecSpecServiceRoutes struct {
	// Path specifies an route by HTTP path prefix. Paths must start with / and must be unique within the app.
	// +optional
	Path *string `json:"path,omitempty" tf:"path"`
}

type AppSpecSpecService struct {
	// An optional build command to run while building this component from source.
	// +optional
	BuildCommand *string `json:"buildCommand,omitempty" tf:"build_command"`
	// The path to a Dockerfile relative to the root of the repo. If set, overrides usage of buildpacks.
	// +optional
	DockerfilePath *string `json:"dockerfilePath,omitempty" tf:"dockerfile_path"`
	// +optional
	Env []AppSpecSpecServiceEnv `json:"env,omitempty" tf:"env"`
	// An environment slug describing the type of this app.
	// +optional
	EnvironmentSlug *string `json:"environmentSlug,omitempty" tf:"environment_slug"`
	// +optional
	Git *AppSpecSpecServiceGit `json:"git,omitempty" tf:"git"`
	// +optional
	Github *AppSpecSpecServiceGithub `json:"github,omitempty" tf:"github"`
	// +optional
	Gitlab *AppSpecSpecServiceGitlab `json:"gitlab,omitempty" tf:"gitlab"`
	// +optional
	HealthCheck *AppSpecSpecServiceHealthCheck `json:"healthCheck,omitempty" tf:"health_check"`
	// The internal port on which this service's run command will listen.
	// +optional
	HttpPort *int64 `json:"httpPort,omitempty" tf:"http_port"`
	// +optional
	Image *AppSpecSpecServiceImage `json:"image,omitempty" tf:"image"`
	// The amount of instances that this component should be scaled to.
	// +optional
	InstanceCount *int64 `json:"instanceCount,omitempty" tf:"instance_count"`
	// The instance size to use for this component.
	// +optional
	InstanceSizeSlug *string `json:"instanceSizeSlug,omitempty" tf:"instance_size_slug"`
	// +optional
	InternalPorts []int64 `json:"internalPorts,omitempty" tf:"internal_ports"`
	// The name of the component
	Name *string `json:"name" tf:"name"`
	// +optional
	Routes []AppSpecSpecServiceRoutes `json:"routes,omitempty" tf:"routes"`
	// An optional run command to override the component's default.
	// +optional
	RunCommand *string `json:"runCommand,omitempty" tf:"run_command"`
	// An optional path to the working directory to use for the build.
	// +optional
	SourceDir *string `json:"sourceDir,omitempty" tf:"source_dir"`
}

type AppSpecSpecStaticSiteEnv struct {
	// The name of the environment variable.
	// +optional
	Key *string `json:"key,omitempty" tf:"key"`
	// The visibility scope of the environment variable.
	// +optional
	Scope *string `json:"scope,omitempty" tf:"scope"`
	// The type of the environment variable.
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
	// The value of the environment variable.
	// +optional
	Value *string `json:"-" sensitive:"true" tf:"value"`
}

type AppSpecSpecStaticSiteGit struct {
	// The name of the branch to use.
	// +optional
	Branch *string `json:"branch,omitempty" tf:"branch"`
	// The clone URL of the repo.
	// +optional
	RepoCloneURL *string `json:"repoCloneURL,omitempty" tf:"repo_clone_url"`
}

type AppSpecSpecStaticSiteGithub struct {
	// The name of the branch to use.
	// +optional
	Branch *string `json:"branch,omitempty" tf:"branch"`
	// Whether to automatically deploy new commits made to the repo
	// +optional
	DeployOnPush *bool `json:"deployOnPush,omitempty" tf:"deploy_on_push"`
	// The name of the repo in the format `owner/repo`.
	// +optional
	Repo *string `json:"repo,omitempty" tf:"repo"`
}

type AppSpecSpecStaticSiteGitlab struct {
	// The name of the branch to use.
	// +optional
	Branch *string `json:"branch,omitempty" tf:"branch"`
	// Whether to automatically deploy new commits made to the repo
	// +optional
	DeployOnPush *bool `json:"deployOnPush,omitempty" tf:"deploy_on_push"`
	// The name of the repo in the format `owner/repo`.
	// +optional
	Repo *string `json:"repo,omitempty" tf:"repo"`
}

type AppSpecSpecStaticSiteRoutes struct {
	// Path specifies an route by HTTP path prefix. Paths must start with / and must be unique within the app.
	// +optional
	Path *string `json:"path,omitempty" tf:"path"`
}

type AppSpecSpecStaticSite struct {
	// An optional build command to run while building this component from source.
	// +optional
	BuildCommand *string `json:"buildCommand,omitempty" tf:"build_command"`
	// The name of the document to use as the fallback for any requests to documents that are not found when serving this static site.
	// +optional
	CatchallDocument *string `json:"catchallDocument,omitempty" tf:"catchall_document"`
	// The path to a Dockerfile relative to the root of the repo. If set, overrides usage of buildpacks.
	// +optional
	DockerfilePath *string `json:"dockerfilePath,omitempty" tf:"dockerfile_path"`
	// +optional
	Env []AppSpecSpecStaticSiteEnv `json:"env,omitempty" tf:"env"`
	// An environment slug describing the type of this app.
	// +optional
	EnvironmentSlug *string `json:"environmentSlug,omitempty" tf:"environment_slug"`
	// The name of the error document to use when serving this static site.
	// +optional
	ErrorDocument *string `json:"errorDocument,omitempty" tf:"error_document"`
	// +optional
	Git *AppSpecSpecStaticSiteGit `json:"git,omitempty" tf:"git"`
	// +optional
	Github *AppSpecSpecStaticSiteGithub `json:"github,omitempty" tf:"github"`
	// +optional
	Gitlab *AppSpecSpecStaticSiteGitlab `json:"gitlab,omitempty" tf:"gitlab"`
	// The name of the index document to use when serving this static site.
	// +optional
	IndexDocument *string `json:"indexDocument,omitempty" tf:"index_document"`
	// The name of the component
	Name *string `json:"name" tf:"name"`
	// An optional path to where the built assets will be located, relative to the build context. If not set, App Platform will automatically scan for these directory names: `_static`, `dist`, `public`.
	// +optional
	OutputDir *string `json:"outputDir,omitempty" tf:"output_dir"`
	// +optional
	Routes []AppSpecSpecStaticSiteRoutes `json:"routes,omitempty" tf:"routes"`
	// An optional path to the working directory to use for the build.
	// +optional
	SourceDir *string `json:"sourceDir,omitempty" tf:"source_dir"`
}

type AppSpecSpecWorkerEnv struct {
	// The name of the environment variable.
	// +optional
	Key *string `json:"key,omitempty" tf:"key"`
	// The visibility scope of the environment variable.
	// +optional
	Scope *string `json:"scope,omitempty" tf:"scope"`
	// The type of the environment variable.
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
	// The value of the environment variable.
	// +optional
	Value *string `json:"-" sensitive:"true" tf:"value"`
}

type AppSpecSpecWorkerGit struct {
	// The name of the branch to use.
	// +optional
	Branch *string `json:"branch,omitempty" tf:"branch"`
	// The clone URL of the repo.
	// +optional
	RepoCloneURL *string `json:"repoCloneURL,omitempty" tf:"repo_clone_url"`
}

type AppSpecSpecWorkerGithub struct {
	// The name of the branch to use.
	// +optional
	Branch *string `json:"branch,omitempty" tf:"branch"`
	// Whether to automatically deploy new commits made to the repo
	// +optional
	DeployOnPush *bool `json:"deployOnPush,omitempty" tf:"deploy_on_push"`
	// The name of the repo in the format `owner/repo`.
	// +optional
	Repo *string `json:"repo,omitempty" tf:"repo"`
}

type AppSpecSpecWorkerGitlab struct {
	// The name of the branch to use.
	// +optional
	Branch *string `json:"branch,omitempty" tf:"branch"`
	// Whether to automatically deploy new commits made to the repo
	// +optional
	DeployOnPush *bool `json:"deployOnPush,omitempty" tf:"deploy_on_push"`
	// The name of the repo in the format `owner/repo`.
	// +optional
	Repo *string `json:"repo,omitempty" tf:"repo"`
}

type AppSpecSpecWorkerImage struct {
	// The registry name. Must be left empty for the DOCR registry type.
	// +optional
	Registry *string `json:"registry,omitempty" tf:"registry"`
	// The registry type.
	RegistryType *string `json:"registryType" tf:"registry_type"`
	// The repository name.
	Repository *string `json:"repository" tf:"repository"`
	// The repository tag. Defaults to latest if not provided.
	// +optional
	Tag *string `json:"tag,omitempty" tf:"tag"`
}

type AppSpecSpecWorker struct {
	// An optional build command to run while building this component from source.
	// +optional
	BuildCommand *string `json:"buildCommand,omitempty" tf:"build_command"`
	// The path to a Dockerfile relative to the root of the repo. If set, overrides usage of buildpacks.
	// +optional
	DockerfilePath *string `json:"dockerfilePath,omitempty" tf:"dockerfile_path"`
	// +optional
	Env []AppSpecSpecWorkerEnv `json:"env,omitempty" tf:"env"`
	// An environment slug describing the type of this app.
	// +optional
	EnvironmentSlug *string `json:"environmentSlug,omitempty" tf:"environment_slug"`
	// +optional
	Git *AppSpecSpecWorkerGit `json:"git,omitempty" tf:"git"`
	// +optional
	Github *AppSpecSpecWorkerGithub `json:"github,omitempty" tf:"github"`
	// +optional
	Gitlab *AppSpecSpecWorkerGitlab `json:"gitlab,omitempty" tf:"gitlab"`
	// +optional
	Image *AppSpecSpecWorkerImage `json:"image,omitempty" tf:"image"`
	// The amount of instances that this component should be scaled to.
	// +optional
	InstanceCount *int64 `json:"instanceCount,omitempty" tf:"instance_count"`
	// The instance size to use for this component.
	// +optional
	InstanceSizeSlug *string `json:"instanceSizeSlug,omitempty" tf:"instance_size_slug"`
	// The name of the component
	Name *string `json:"name" tf:"name"`
	// An optional run command to override the component's default.
	// +optional
	RunCommand *string `json:"runCommand,omitempty" tf:"run_command"`
	// An optional path to the working directory to use for the build.
	// +optional
	SourceDir *string `json:"sourceDir,omitempty" tf:"source_dir"`
}

type AppSpecSpec struct {
	// +optional
	Database []AppSpecSpecDatabase `json:"database,omitempty" tf:"database"`
	// +optional
	Domain []AppSpecSpecDomain `json:"domain,omitempty" tf:"domain"`
	// +optional
	// Deprecated
	Domains []string `json:"domains,omitempty" tf:"domains"`
	// +optional
	Env []AppSpecSpecEnv `json:"env,omitempty" tf:"env"`
	// +optional
	Job []AppSpecSpecJob `json:"job,omitempty" tf:"job"`
	// The name of the app. Must be unique across all apps in the same account.
	Name *string `json:"name" tf:"name"`
	// The slug for the DigitalOcean data center region hosting the app
	// +optional
	Region *string `json:"region,omitempty" tf:"region"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	Service []AppSpecSpecService `json:"service,omitempty" tf:"service"`
	// +optional
	StaticSite []AppSpecSpecStaticSite `json:"staticSite,omitempty" tf:"static_site"`
	// +optional
	Worker []AppSpecSpecWorker `json:"worker,omitempty" tf:"worker"`
}

type AppSpec struct {
	KubeformOutput *AppSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource AppSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type AppSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID the App's currently active deployment
	// +optional
	ActiveDeploymentID *string `json:"activeDeploymentID,omitempty" tf:"active_deployment_id"`
	// The date and time of when the App was created
	// +optional
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at"`
	// The default URL to access the App
	// +optional
	DefaultIngress *string `json:"defaultIngress,omitempty" tf:"default_ingress"`
	// +optional
	LiveURL *string `json:"liveURL,omitempty" tf:"live_url"`
	// A DigitalOcean App Platform Spec
	// +optional
	Spec *AppSpecSpec `json:"spec,omitempty" tf:"spec"`
	// The date and time of when the App was last updated
	// +optional
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at"`
}

type AppStatus struct {
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

// AppList is a list of Apps
type AppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of App CRD objects
	Items []App `json:"items,omitempty"`
}
