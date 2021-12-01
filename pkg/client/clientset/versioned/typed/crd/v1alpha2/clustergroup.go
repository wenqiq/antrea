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

	v1alpha2 "antrea.io/antrea/pkg/apis/crd/v1alpha2"
	scheme "antrea.io/antrea/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterGroupsGetter has a method to return a ClusterGroupInterface.
// A group's client should implement this interface.
type ClusterGroupsGetter interface {
	ClusterGroups() ClusterGroupInterface
}

// ClusterGroupInterface has methods to work with ClusterGroup resources.
type ClusterGroupInterface interface {
	Create(ctx context.Context, clusterGroup *v1alpha2.ClusterGroup, opts v1.CreateOptions) (*v1alpha2.ClusterGroup, error)
	Update(ctx context.Context, clusterGroup *v1alpha2.ClusterGroup, opts v1.UpdateOptions) (*v1alpha2.ClusterGroup, error)
	UpdateStatus(ctx context.Context, clusterGroup *v1alpha2.ClusterGroup, opts v1.UpdateOptions) (*v1alpha2.ClusterGroup, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha2.ClusterGroup, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha2.ClusterGroupList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.ClusterGroup, err error)
	ClusterGroupExpansion
}

// clusterGroups implements ClusterGroupInterface
type clusterGroups struct {
	client rest.Interface
}

// newClusterGroups returns a ClusterGroups
func newClusterGroups(c *CrdV1alpha2Client) *clusterGroups {
	return &clusterGroups{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterGroup, and returns the corresponding clusterGroup object, and an error if there is any.
func (c *clusterGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.ClusterGroup, err error) {
	result = &v1alpha2.ClusterGroup{}
	err = c.client.Get().
		Resource("clustergroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterGroups that match those selectors.
func (c *clusterGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.ClusterGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.ClusterGroupList{}
	err = c.client.Get().
		Resource("clustergroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterGroups.
func (c *clusterGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clustergroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterGroup and creates it.  Returns the server's representation of the clusterGroup, and an error, if there is any.
func (c *clusterGroups) Create(ctx context.Context, clusterGroup *v1alpha2.ClusterGroup, opts v1.CreateOptions) (result *v1alpha2.ClusterGroup, err error) {
	result = &v1alpha2.ClusterGroup{}
	err = c.client.Post().
		Resource("clustergroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterGroup).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterGroup and updates it. Returns the server's representation of the clusterGroup, and an error, if there is any.
func (c *clusterGroups) Update(ctx context.Context, clusterGroup *v1alpha2.ClusterGroup, opts v1.UpdateOptions) (result *v1alpha2.ClusterGroup, err error) {
	result = &v1alpha2.ClusterGroup{}
	err = c.client.Put().
		Resource("clustergroups").
		Name(clusterGroup.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterGroup).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clusterGroups) UpdateStatus(ctx context.Context, clusterGroup *v1alpha2.ClusterGroup, opts v1.UpdateOptions) (result *v1alpha2.ClusterGroup, err error) {
	result = &v1alpha2.ClusterGroup{}
	err = c.client.Put().
		Resource("clustergroups").
		Name(clusterGroup.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterGroup).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterGroup and deletes it. Returns an error if one occurs.
func (c *clusterGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clustergroups").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clustergroups").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterGroup.
func (c *clusterGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.ClusterGroup, err error) {
	result = &v1alpha2.ClusterGroup{}
	err = c.client.Patch(pt).
		Resource("clustergroups").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
