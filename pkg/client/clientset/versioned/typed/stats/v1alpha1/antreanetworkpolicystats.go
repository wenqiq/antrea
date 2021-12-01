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

	v1alpha1 "antrea.io/antrea/pkg/apis/stats/v1alpha1"
	scheme "antrea.io/antrea/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AntreaNetworkPolicyStatsGetter has a method to return a AntreaNetworkPolicyStatsInterface.
// A group's client should implement this interface.
type AntreaNetworkPolicyStatsGetter interface {
	AntreaNetworkPolicyStats(namespace string) AntreaNetworkPolicyStatsInterface
}

// AntreaNetworkPolicyStatsInterface has methods to work with AntreaNetworkPolicyStats resources.
type AntreaNetworkPolicyStatsInterface interface {
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.AntreaNetworkPolicyStats, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.AntreaNetworkPolicyStatsList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	AntreaNetworkPolicyStatsExpansion
}

// antreaNetworkPolicyStats implements AntreaNetworkPolicyStatsInterface
type antreaNetworkPolicyStats struct {
	client rest.Interface
	ns     string
}

// newAntreaNetworkPolicyStats returns a AntreaNetworkPolicyStats
func newAntreaNetworkPolicyStats(c *StatsV1alpha1Client, namespace string) *antreaNetworkPolicyStats {
	return &antreaNetworkPolicyStats{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the antreaNetworkPolicyStats, and returns the corresponding antreaNetworkPolicyStats object, and an error if there is any.
func (c *antreaNetworkPolicyStats) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AntreaNetworkPolicyStats, err error) {
	result = &v1alpha1.AntreaNetworkPolicyStats{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("antreanetworkpolicystats").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AntreaNetworkPolicyStats that match those selectors.
func (c *antreaNetworkPolicyStats) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AntreaNetworkPolicyStatsList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AntreaNetworkPolicyStatsList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("antreanetworkpolicystats").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested antreaNetworkPolicyStats.
func (c *antreaNetworkPolicyStats) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("antreanetworkpolicystats").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}
