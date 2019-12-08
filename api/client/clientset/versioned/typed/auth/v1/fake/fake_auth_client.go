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
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
	v1 "tkestack.io/tke/api/client/clientset/versioned/typed/auth/v1"
)

type FakeAuthV1 struct {
	*testing.Fake
}

func (c *FakeAuthV1) APIKeys() v1.APIKeyInterface {
	return &FakeAPIKeys{c}
}

func (c *FakeAuthV1) APISigningKeys() v1.APISigningKeyInterface {
	return &FakeAPISigningKeys{c}
}

func (c *FakeAuthV1) Categories() v1.CategoryInterface {
	return &FakeCategories{c}
}

func (c *FakeAuthV1) ConfigMaps() v1.ConfigMapInterface {
	return &FakeConfigMaps{c}
}

func (c *FakeAuthV1) Groups() v1.GroupInterface {
	return &FakeGroups{c}
}

func (c *FakeAuthV1) LocalIdentities() v1.LocalIdentityInterface {
	return &FakeLocalIdentities{c}
}

func (c *FakeAuthV1) Policies() v1.PolicyInterface {
	return &FakePolicies{c}
}

func (c *FakeAuthV1) Rules() v1.RuleInterface {
	return &FakeRules{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeAuthV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
