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

package v1beta2

import (
	"context"
	"time"

	v1beta2 "antrea.io/antrea/pkg/apis/controlplane/v1beta2"
	scheme "antrea.io/antrea/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AppliedToGroupsGetter has a method to return a AppliedToGroupInterface.
// A group's client should implement this interface.
type AppliedToGroupsGetter interface {
	AppliedToGroups() AppliedToGroupInterface
}

// AppliedToGroupInterface has methods to work with AppliedToGroup resources.
type AppliedToGroupInterface interface {
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta2.AppliedToGroup, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta2.AppliedToGroupList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	AppliedToGroupExpansion
}

// appliedToGroups implements AppliedToGroupInterface
type appliedToGroups struct {
	client rest.Interface
}

// newAppliedToGroups returns a AppliedToGroups
func newAppliedToGroups(c *ControlplaneV1beta2Client) *appliedToGroups {
	return &appliedToGroups{
		client: c.RESTClient(),
	}
}

// Get takes name of the appliedToGroup, and returns the corresponding appliedToGroup object, and an error if there is any.
func (c *appliedToGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta2.AppliedToGroup, err error) {
	result = &v1beta2.AppliedToGroup{}
	err = c.client.Get().
		Resource("appliedtogroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AppliedToGroups that match those selectors.
func (c *appliedToGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1beta2.AppliedToGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta2.AppliedToGroupList{}
	err = c.client.Get().
		Resource("appliedtogroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested appliedToGroups.
func (c *appliedToGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("appliedtogroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}
