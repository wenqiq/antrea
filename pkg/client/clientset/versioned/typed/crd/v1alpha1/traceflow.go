// Copyright 2022 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "antrea.io/antrea/pkg/apis/crd/v1alpha1"
	scheme "antrea.io/antrea/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TraceflowsGetter has a method to return a TraceflowInterface.
// A group's client should implement this interface.
type TraceflowsGetter interface {
	Traceflows() TraceflowInterface
}

// TraceflowInterface has methods to work with Traceflow resources.
type TraceflowInterface interface {
	Create(ctx context.Context, traceflow *v1alpha1.Traceflow, opts v1.CreateOptions) (*v1alpha1.Traceflow, error)
	Update(ctx context.Context, traceflow *v1alpha1.Traceflow, opts v1.UpdateOptions) (*v1alpha1.Traceflow, error)
	UpdateStatus(ctx context.Context, traceflow *v1alpha1.Traceflow, opts v1.UpdateOptions) (*v1alpha1.Traceflow, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Traceflow, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.TraceflowList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Traceflow, err error)
	TraceflowExpansion
}

// traceflows implements TraceflowInterface
type traceflows struct {
	client rest.Interface
}

// newTraceflows returns a Traceflows
func newTraceflows(c *CrdV1alpha1Client) *traceflows {
	return &traceflows{
		client: c.RESTClient(),
	}
}

// Get takes name of the traceflow, and returns the corresponding traceflow object, and an error if there is any.
func (c *traceflows) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Traceflow, err error) {
	result = &v1alpha1.Traceflow{}
	err = c.client.Get().
		Resource("traceflows").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Traceflows that match those selectors.
func (c *traceflows) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TraceflowList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.TraceflowList{}
	err = c.client.Get().
		Resource("traceflows").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested traceflows.
func (c *traceflows) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("traceflows").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a traceflow and creates it.  Returns the server's representation of the traceflow, and an error, if there is any.
func (c *traceflows) Create(ctx context.Context, traceflow *v1alpha1.Traceflow, opts v1.CreateOptions) (result *v1alpha1.Traceflow, err error) {
	result = &v1alpha1.Traceflow{}
	err = c.client.Post().
		Resource("traceflows").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(traceflow).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a traceflow and updates it. Returns the server's representation of the traceflow, and an error, if there is any.
func (c *traceflows) Update(ctx context.Context, traceflow *v1alpha1.Traceflow, opts v1.UpdateOptions) (result *v1alpha1.Traceflow, err error) {
	result = &v1alpha1.Traceflow{}
	err = c.client.Put().
		Resource("traceflows").
		Name(traceflow.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(traceflow).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *traceflows) UpdateStatus(ctx context.Context, traceflow *v1alpha1.Traceflow, opts v1.UpdateOptions) (result *v1alpha1.Traceflow, err error) {
	result = &v1alpha1.Traceflow{}
	err = c.client.Put().
		Resource("traceflows").
		Name(traceflow.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(traceflow).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the traceflow and deletes it. Returns an error if one occurs.
func (c *traceflows) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("traceflows").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *traceflows) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("traceflows").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched traceflow.
func (c *traceflows) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Traceflow, err error) {
	result = &v1alpha1.Traceflow{}
	err = c.client.Patch(pt).
		Resource("traceflows").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
