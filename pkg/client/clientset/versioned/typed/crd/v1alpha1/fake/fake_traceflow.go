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

package fake

import (
	"context"

	v1alpha1 "antrea.io/antrea/pkg/apis/crd/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTraceflows implements TraceflowInterface
type FakeTraceflows struct {
	Fake *FakeCrdV1alpha1
}

var traceflowsResource = schema.GroupVersionResource{Group: "crd.antrea.io", Version: "v1alpha1", Resource: "traceflows"}

var traceflowsKind = schema.GroupVersionKind{Group: "crd.antrea.io", Version: "v1alpha1", Kind: "Traceflow"}

// Get takes name of the traceflow, and returns the corresponding traceflow object, and an error if there is any.
func (c *FakeTraceflows) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Traceflow, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(traceflowsResource, name), &v1alpha1.Traceflow{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Traceflow), err
}

// List takes label and field selectors, and returns the list of Traceflows that match those selectors.
func (c *FakeTraceflows) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TraceflowList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(traceflowsResource, traceflowsKind, opts), &v1alpha1.TraceflowList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TraceflowList{ListMeta: obj.(*v1alpha1.TraceflowList).ListMeta}
	for _, item := range obj.(*v1alpha1.TraceflowList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested traceflows.
func (c *FakeTraceflows) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(traceflowsResource, opts))
}

// Create takes the representation of a traceflow and creates it.  Returns the server's representation of the traceflow, and an error, if there is any.
func (c *FakeTraceflows) Create(ctx context.Context, traceflow *v1alpha1.Traceflow, opts v1.CreateOptions) (result *v1alpha1.Traceflow, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(traceflowsResource, traceflow), &v1alpha1.Traceflow{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Traceflow), err
}

// Update takes the representation of a traceflow and updates it. Returns the server's representation of the traceflow, and an error, if there is any.
func (c *FakeTraceflows) Update(ctx context.Context, traceflow *v1alpha1.Traceflow, opts v1.UpdateOptions) (result *v1alpha1.Traceflow, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(traceflowsResource, traceflow), &v1alpha1.Traceflow{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Traceflow), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTraceflows) UpdateStatus(ctx context.Context, traceflow *v1alpha1.Traceflow, opts v1.UpdateOptions) (*v1alpha1.Traceflow, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(traceflowsResource, "status", traceflow), &v1alpha1.Traceflow{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Traceflow), err
}

// Delete takes name of the traceflow and deletes it. Returns an error if one occurs.
func (c *FakeTraceflows) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(traceflowsResource, name), &v1alpha1.Traceflow{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTraceflows) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(traceflowsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TraceflowList{})
	return err
}

// Patch applies the patch and returns the patched traceflow.
func (c *FakeTraceflows) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Traceflow, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(traceflowsResource, name, pt, data, subresources...), &v1alpha1.Traceflow{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Traceflow), err
}
