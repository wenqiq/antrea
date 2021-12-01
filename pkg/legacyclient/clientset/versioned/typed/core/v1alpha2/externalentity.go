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

package v1alpha2

import (
	"context"
	"time"

	v1alpha2 "antrea.io/antrea/pkg/legacyapis/core/v1alpha2"
	scheme "antrea.io/antrea/pkg/legacyclient/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ExternalEntitiesGetter has a method to return a ExternalEntityInterface.
// A group's client should implement this interface.
type ExternalEntitiesGetter interface {
	ExternalEntities(namespace string) ExternalEntityInterface
}

// ExternalEntityInterface has methods to work with ExternalEntity resources.
type ExternalEntityInterface interface {
	Create(ctx context.Context, externalEntity *v1alpha2.ExternalEntity, opts v1.CreateOptions) (*v1alpha2.ExternalEntity, error)
	Update(ctx context.Context, externalEntity *v1alpha2.ExternalEntity, opts v1.UpdateOptions) (*v1alpha2.ExternalEntity, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha2.ExternalEntity, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha2.ExternalEntityList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.ExternalEntity, err error)
	ExternalEntityExpansion
}

// externalEntities implements ExternalEntityInterface
type externalEntities struct {
	client rest.Interface
	ns     string
}

// newExternalEntities returns a ExternalEntities
func newExternalEntities(c *CoreV1alpha2Client, namespace string) *externalEntities {
	return &externalEntities{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the externalEntity, and returns the corresponding externalEntity object, and an error if there is any.
func (c *externalEntities) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.ExternalEntity, err error) {
	result = &v1alpha2.ExternalEntity{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("externalentities").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ExternalEntities that match those selectors.
func (c *externalEntities) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.ExternalEntityList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.ExternalEntityList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("externalentities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested externalEntities.
func (c *externalEntities) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("externalentities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a externalEntity and creates it.  Returns the server's representation of the externalEntity, and an error, if there is any.
func (c *externalEntities) Create(ctx context.Context, externalEntity *v1alpha2.ExternalEntity, opts v1.CreateOptions) (result *v1alpha2.ExternalEntity, err error) {
	result = &v1alpha2.ExternalEntity{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("externalentities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(externalEntity).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a externalEntity and updates it. Returns the server's representation of the externalEntity, and an error, if there is any.
func (c *externalEntities) Update(ctx context.Context, externalEntity *v1alpha2.ExternalEntity, opts v1.UpdateOptions) (result *v1alpha2.ExternalEntity, err error) {
	result = &v1alpha2.ExternalEntity{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("externalentities").
		Name(externalEntity.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(externalEntity).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the externalEntity and deletes it. Returns an error if one occurs.
func (c *externalEntities) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("externalentities").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *externalEntities) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("externalentities").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched externalEntity.
func (c *externalEntities) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.ExternalEntity, err error) {
	result = &v1alpha2.ExternalEntity{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("externalentities").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
