/*
 * Copyright 2019 THL A29 Limited, a Tencent company.
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

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	platformv1 "tkestack.io/tke/api/platform/v1"
)

// FakePersistentEvents implements PersistentEventInterface
type FakePersistentEvents struct {
	Fake *FakePlatformV1
}

var persistenteventsResource = schema.GroupVersionResource{Group: "platform.tkestack.io", Version: "v1", Resource: "persistentevents"}

var persistenteventsKind = schema.GroupVersionKind{Group: "platform.tkestack.io", Version: "v1", Kind: "PersistentEvent"}

// Get takes name of the persistentEvent, and returns the corresponding persistentEvent object, and an error if there is any.
func (c *FakePersistentEvents) Get(name string, options v1.GetOptions) (result *platformv1.PersistentEvent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(persistenteventsResource, name), &platformv1.PersistentEvent{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.PersistentEvent), err
}

// List takes label and field selectors, and returns the list of PersistentEvents that match those selectors.
func (c *FakePersistentEvents) List(opts v1.ListOptions) (result *platformv1.PersistentEventList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(persistenteventsResource, persistenteventsKind, opts), &platformv1.PersistentEventList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &platformv1.PersistentEventList{ListMeta: obj.(*platformv1.PersistentEventList).ListMeta}
	for _, item := range obj.(*platformv1.PersistentEventList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested persistentEvents.
func (c *FakePersistentEvents) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(persistenteventsResource, opts))
}

// Create takes the representation of a persistentEvent and creates it.  Returns the server's representation of the persistentEvent, and an error, if there is any.
func (c *FakePersistentEvents) Create(persistentEvent *platformv1.PersistentEvent) (result *platformv1.PersistentEvent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(persistenteventsResource, persistentEvent), &platformv1.PersistentEvent{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.PersistentEvent), err
}

// Update takes the representation of a persistentEvent and updates it. Returns the server's representation of the persistentEvent, and an error, if there is any.
func (c *FakePersistentEvents) Update(persistentEvent *platformv1.PersistentEvent) (result *platformv1.PersistentEvent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(persistenteventsResource, persistentEvent), &platformv1.PersistentEvent{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.PersistentEvent), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePersistentEvents) UpdateStatus(persistentEvent *platformv1.PersistentEvent) (*platformv1.PersistentEvent, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(persistenteventsResource, "status", persistentEvent), &platformv1.PersistentEvent{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.PersistentEvent), err
}

// Delete takes name of the persistentEvent and deletes it. Returns an error if one occurs.
func (c *FakePersistentEvents) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(persistenteventsResource, name), &platformv1.PersistentEvent{})
	return err
}

// Patch applies the patch and returns the patched persistentEvent.
func (c *FakePersistentEvents) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *platformv1.PersistentEvent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(persistenteventsResource, name, pt, data, subresources...), &platformv1.PersistentEvent{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.PersistentEvent), err
}
