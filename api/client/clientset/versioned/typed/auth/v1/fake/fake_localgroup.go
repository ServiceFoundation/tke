/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
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
	authv1 "tkestack.io/tke/api/auth/v1"
)

// FakeLocalGroups implements LocalGroupInterface
type FakeLocalGroups struct {
	Fake *FakeAuthV1
}

var localgroupsResource = schema.GroupVersionResource{Group: "auth.tkestack.io", Version: "v1", Resource: "localgroups"}

var localgroupsKind = schema.GroupVersionKind{Group: "auth.tkestack.io", Version: "v1", Kind: "LocalGroup"}

// Get takes name of the localGroup, and returns the corresponding localGroup object, and an error if there is any.
func (c *FakeLocalGroups) Get(name string, options v1.GetOptions) (result *authv1.LocalGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(localgroupsResource, name), &authv1.LocalGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*authv1.LocalGroup), err
}

// List takes label and field selectors, and returns the list of LocalGroups that match those selectors.
func (c *FakeLocalGroups) List(opts v1.ListOptions) (result *authv1.LocalGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(localgroupsResource, localgroupsKind, opts), &authv1.LocalGroupList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &authv1.LocalGroupList{ListMeta: obj.(*authv1.LocalGroupList).ListMeta}
	for _, item := range obj.(*authv1.LocalGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested localGroups.
func (c *FakeLocalGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(localgroupsResource, opts))
}

// Create takes the representation of a localGroup and creates it.  Returns the server's representation of the localGroup, and an error, if there is any.
func (c *FakeLocalGroups) Create(localGroup *authv1.LocalGroup) (result *authv1.LocalGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(localgroupsResource, localGroup), &authv1.LocalGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*authv1.LocalGroup), err
}

// Update takes the representation of a localGroup and updates it. Returns the server's representation of the localGroup, and an error, if there is any.
func (c *FakeLocalGroups) Update(localGroup *authv1.LocalGroup) (result *authv1.LocalGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(localgroupsResource, localGroup), &authv1.LocalGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*authv1.LocalGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLocalGroups) UpdateStatus(localGroup *authv1.LocalGroup) (*authv1.LocalGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(localgroupsResource, "status", localGroup), &authv1.LocalGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*authv1.LocalGroup), err
}

// Delete takes name of the localGroup and deletes it. Returns an error if one occurs.
func (c *FakeLocalGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(localgroupsResource, name), &authv1.LocalGroup{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLocalGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(localgroupsResource, listOptions)

	_, err := c.Fake.Invokes(action, &authv1.LocalGroupList{})
	return err
}

// Patch applies the patch and returns the patched localGroup.
func (c *FakeLocalGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *authv1.LocalGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(localgroupsResource, name, pt, data, subresources...), &authv1.LocalGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*authv1.LocalGroup), err
}
