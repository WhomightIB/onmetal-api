/*
 * Copyright (c) 2022 by the OnMetal authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
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
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1alpha1 "github.com/onmetal/onmetal-api/api/networking/v1alpha1"
	networkingv1alpha1 "github.com/onmetal/onmetal-api/client-go/applyconfigurations/networking/v1alpha1"
	scheme "github.com/onmetal/onmetal-api/client-go/onmetalapi/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AliasPrefixRoutingsGetter has a method to return a AliasPrefixRoutingInterface.
// A group's client should implement this interface.
type AliasPrefixRoutingsGetter interface {
	AliasPrefixRoutings(namespace string) AliasPrefixRoutingInterface
}

// AliasPrefixRoutingInterface has methods to work with AliasPrefixRouting resources.
type AliasPrefixRoutingInterface interface {
	Create(ctx context.Context, aliasPrefixRouting *v1alpha1.AliasPrefixRouting, opts v1.CreateOptions) (*v1alpha1.AliasPrefixRouting, error)
	Update(ctx context.Context, aliasPrefixRouting *v1alpha1.AliasPrefixRouting, opts v1.UpdateOptions) (*v1alpha1.AliasPrefixRouting, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.AliasPrefixRouting, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.AliasPrefixRoutingList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AliasPrefixRouting, err error)
	Apply(ctx context.Context, aliasPrefixRouting *networkingv1alpha1.AliasPrefixRoutingApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.AliasPrefixRouting, err error)
	AliasPrefixRoutingExpansion
}

// aliasPrefixRoutings implements AliasPrefixRoutingInterface
type aliasPrefixRoutings struct {
	client rest.Interface
	ns     string
}

// newAliasPrefixRoutings returns a AliasPrefixRoutings
func newAliasPrefixRoutings(c *NetworkingV1alpha1Client, namespace string) *aliasPrefixRoutings {
	return &aliasPrefixRoutings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the aliasPrefixRouting, and returns the corresponding aliasPrefixRouting object, and an error if there is any.
func (c *aliasPrefixRoutings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AliasPrefixRouting, err error) {
	result = &v1alpha1.AliasPrefixRouting{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("aliasprefixroutings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AliasPrefixRoutings that match those selectors.
func (c *aliasPrefixRoutings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AliasPrefixRoutingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AliasPrefixRoutingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("aliasprefixroutings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested aliasPrefixRoutings.
func (c *aliasPrefixRoutings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("aliasprefixroutings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a aliasPrefixRouting and creates it.  Returns the server's representation of the aliasPrefixRouting, and an error, if there is any.
func (c *aliasPrefixRoutings) Create(ctx context.Context, aliasPrefixRouting *v1alpha1.AliasPrefixRouting, opts v1.CreateOptions) (result *v1alpha1.AliasPrefixRouting, err error) {
	result = &v1alpha1.AliasPrefixRouting{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("aliasprefixroutings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aliasPrefixRouting).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a aliasPrefixRouting and updates it. Returns the server's representation of the aliasPrefixRouting, and an error, if there is any.
func (c *aliasPrefixRoutings) Update(ctx context.Context, aliasPrefixRouting *v1alpha1.AliasPrefixRouting, opts v1.UpdateOptions) (result *v1alpha1.AliasPrefixRouting, err error) {
	result = &v1alpha1.AliasPrefixRouting{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("aliasprefixroutings").
		Name(aliasPrefixRouting.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aliasPrefixRouting).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the aliasPrefixRouting and deletes it. Returns an error if one occurs.
func (c *aliasPrefixRoutings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("aliasprefixroutings").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *aliasPrefixRoutings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("aliasprefixroutings").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched aliasPrefixRouting.
func (c *aliasPrefixRoutings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AliasPrefixRouting, err error) {
	result = &v1alpha1.AliasPrefixRouting{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("aliasprefixroutings").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied aliasPrefixRouting.
func (c *aliasPrefixRoutings) Apply(ctx context.Context, aliasPrefixRouting *networkingv1alpha1.AliasPrefixRoutingApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.AliasPrefixRouting, err error) {
	if aliasPrefixRouting == nil {
		return nil, fmt.Errorf("aliasPrefixRouting provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(aliasPrefixRouting)
	if err != nil {
		return nil, err
	}
	name := aliasPrefixRouting.Name
	if name == nil {
		return nil, fmt.Errorf("aliasPrefixRouting.Name must be provided to Apply")
	}
	result = &v1alpha1.AliasPrefixRouting{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("aliasprefixroutings").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
