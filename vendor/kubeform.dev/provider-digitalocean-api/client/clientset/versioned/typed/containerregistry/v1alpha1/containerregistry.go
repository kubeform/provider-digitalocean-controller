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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "kubeform.dev/provider-digitalocean-api/apis/containerregistry/v1alpha1"
	scheme "kubeform.dev/provider-digitalocean-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ContainerRegistriesGetter has a method to return a ContainerRegistryInterface.
// A group's client should implement this interface.
type ContainerRegistriesGetter interface {
	ContainerRegistries(namespace string) ContainerRegistryInterface
}

// ContainerRegistryInterface has methods to work with ContainerRegistry resources.
type ContainerRegistryInterface interface {
	Create(ctx context.Context, containerRegistry *v1alpha1.ContainerRegistry, opts v1.CreateOptions) (*v1alpha1.ContainerRegistry, error)
	Update(ctx context.Context, containerRegistry *v1alpha1.ContainerRegistry, opts v1.UpdateOptions) (*v1alpha1.ContainerRegistry, error)
	UpdateStatus(ctx context.Context, containerRegistry *v1alpha1.ContainerRegistry, opts v1.UpdateOptions) (*v1alpha1.ContainerRegistry, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ContainerRegistry, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ContainerRegistryList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ContainerRegistry, err error)
	ContainerRegistryExpansion
}

// containerRegistries implements ContainerRegistryInterface
type containerRegistries struct {
	client rest.Interface
	ns     string
}

// newContainerRegistries returns a ContainerRegistries
func newContainerRegistries(c *ContainerregistryV1alpha1Client, namespace string) *containerRegistries {
	return &containerRegistries{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the containerRegistry, and returns the corresponding containerRegistry object, and an error if there is any.
func (c *containerRegistries) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ContainerRegistry, err error) {
	result = &v1alpha1.ContainerRegistry{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("containerregistries").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ContainerRegistries that match those selectors.
func (c *containerRegistries) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ContainerRegistryList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ContainerRegistryList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("containerregistries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested containerRegistries.
func (c *containerRegistries) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("containerregistries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a containerRegistry and creates it.  Returns the server's representation of the containerRegistry, and an error, if there is any.
func (c *containerRegistries) Create(ctx context.Context, containerRegistry *v1alpha1.ContainerRegistry, opts v1.CreateOptions) (result *v1alpha1.ContainerRegistry, err error) {
	result = &v1alpha1.ContainerRegistry{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("containerregistries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(containerRegistry).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a containerRegistry and updates it. Returns the server's representation of the containerRegistry, and an error, if there is any.
func (c *containerRegistries) Update(ctx context.Context, containerRegistry *v1alpha1.ContainerRegistry, opts v1.UpdateOptions) (result *v1alpha1.ContainerRegistry, err error) {
	result = &v1alpha1.ContainerRegistry{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("containerregistries").
		Name(containerRegistry.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(containerRegistry).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *containerRegistries) UpdateStatus(ctx context.Context, containerRegistry *v1alpha1.ContainerRegistry, opts v1.UpdateOptions) (result *v1alpha1.ContainerRegistry, err error) {
	result = &v1alpha1.ContainerRegistry{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("containerregistries").
		Name(containerRegistry.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(containerRegistry).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the containerRegistry and deletes it. Returns an error if one occurs.
func (c *containerRegistries) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("containerregistries").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *containerRegistries) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("containerregistries").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched containerRegistry.
func (c *containerRegistries) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ContainerRegistry, err error) {
	result = &v1alpha1.ContainerRegistry{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("containerregistries").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
