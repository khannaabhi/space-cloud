/*
Copyright The Space Cloud Authors.

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

	v1alpha1 "github.com/spacecloud-io/space-cloud/pkg/apis/core/v1alpha1"
	scheme "github.com/spacecloud-io/space-cloud/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// OPAPoliciesGetter has a method to return a OPAPolicyInterface.
// A group's client should implement this interface.
type OPAPoliciesGetter interface {
	OPAPolicies(namespace string) OPAPolicyInterface
}

// OPAPolicyInterface has methods to work with OPAPolicy resources.
type OPAPolicyInterface interface {
	Create(ctx context.Context, oPAPolicy *v1alpha1.OPAPolicy, opts v1.CreateOptions) (*v1alpha1.OPAPolicy, error)
	Update(ctx context.Context, oPAPolicy *v1alpha1.OPAPolicy, opts v1.UpdateOptions) (*v1alpha1.OPAPolicy, error)
	UpdateStatus(ctx context.Context, oPAPolicy *v1alpha1.OPAPolicy, opts v1.UpdateOptions) (*v1alpha1.OPAPolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.OPAPolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.OPAPolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OPAPolicy, err error)
	OPAPolicyExpansion
}

// oPAPolicies implements OPAPolicyInterface
type oPAPolicies struct {
	client rest.Interface
	ns     string
}

// newOPAPolicies returns a OPAPolicies
func newOPAPolicies(c *CoreV1alpha1Client, namespace string) *oPAPolicies {
	return &oPAPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the oPAPolicy, and returns the corresponding oPAPolicy object, and an error if there is any.
func (c *oPAPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.OPAPolicy, err error) {
	result = &v1alpha1.OPAPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("opapolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of OPAPolicies that match those selectors.
func (c *oPAPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.OPAPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.OPAPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("opapolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested oPAPolicies.
func (c *oPAPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("opapolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a oPAPolicy and creates it.  Returns the server's representation of the oPAPolicy, and an error, if there is any.
func (c *oPAPolicies) Create(ctx context.Context, oPAPolicy *v1alpha1.OPAPolicy, opts v1.CreateOptions) (result *v1alpha1.OPAPolicy, err error) {
	result = &v1alpha1.OPAPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("opapolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(oPAPolicy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a oPAPolicy and updates it. Returns the server's representation of the oPAPolicy, and an error, if there is any.
func (c *oPAPolicies) Update(ctx context.Context, oPAPolicy *v1alpha1.OPAPolicy, opts v1.UpdateOptions) (result *v1alpha1.OPAPolicy, err error) {
	result = &v1alpha1.OPAPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("opapolicies").
		Name(oPAPolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(oPAPolicy).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *oPAPolicies) UpdateStatus(ctx context.Context, oPAPolicy *v1alpha1.OPAPolicy, opts v1.UpdateOptions) (result *v1alpha1.OPAPolicy, err error) {
	result = &v1alpha1.OPAPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("opapolicies").
		Name(oPAPolicy.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(oPAPolicy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the oPAPolicy and deletes it. Returns an error if one occurs.
func (c *oPAPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("opapolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *oPAPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("opapolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched oPAPolicy.
func (c *oPAPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OPAPolicy, err error) {
	result = &v1alpha1.OPAPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("opapolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}