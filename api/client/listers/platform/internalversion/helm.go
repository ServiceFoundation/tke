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

// Code generated by lister-gen. DO NOT EDIT.

package internalversion

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	platform "tkestack.io/tke/api/platform"
)

// HelmLister helps list Helms.
type HelmLister interface {
	// List lists all Helms in the indexer.
	List(selector labels.Selector) (ret []*platform.Helm, err error)
	// Get retrieves the Helm from the index for a given name.
	Get(name string) (*platform.Helm, error)
	HelmListerExpansion
}

// helmLister implements the HelmLister interface.
type helmLister struct {
	indexer cache.Indexer
}

// NewHelmLister returns a new HelmLister.
func NewHelmLister(indexer cache.Indexer) HelmLister {
	return &helmLister{indexer: indexer}
}

// List lists all Helms in the indexer.
func (s *helmLister) List(selector labels.Selector) (ret []*platform.Helm, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*platform.Helm))
	})
	return ret, err
}

// Get retrieves the Helm from the index for a given name.
func (s *helmLister) Get(name string) (*platform.Helm, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(platform.Resource("helm"), name)
	}
	return obj.(*platform.Helm), nil
}
