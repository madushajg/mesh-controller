/*
 * Copyright (c) 2019 WSO2 Inc. (http:www.wso2.org) All Rights Reserved.
 *
 * WSO2 Inc. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http:www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha2 "github.com/cellery-io/mesh-controller/pkg/apis/mesh/v1alpha2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CellLister helps list Cells.
type CellLister interface {
	// List lists all Cells in the indexer.
	List(selector labels.Selector) (ret []*v1alpha2.Cell, err error)
	// Cells returns an object that can list and get Cells.
	Cells(namespace string) CellNamespaceLister
	CellListerExpansion
}

// cellLister implements the CellLister interface.
type cellLister struct {
	indexer cache.Indexer
}

// NewCellLister returns a new CellLister.
func NewCellLister(indexer cache.Indexer) CellLister {
	return &cellLister{indexer: indexer}
}

// List lists all Cells in the indexer.
func (s *cellLister) List(selector labels.Selector) (ret []*v1alpha2.Cell, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.Cell))
	})
	return ret, err
}

// Cells returns an object that can list and get Cells.
func (s *cellLister) Cells(namespace string) CellNamespaceLister {
	return cellNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CellNamespaceLister helps list and get Cells.
type CellNamespaceLister interface {
	// List lists all Cells in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha2.Cell, err error)
	// Get retrieves the Cell from the indexer for a given namespace and name.
	Get(name string) (*v1alpha2.Cell, error)
	CellNamespaceListerExpansion
}

// cellNamespaceLister implements the CellNamespaceLister
// interface.
type cellNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Cells in the indexer for a given namespace.
func (s cellNamespaceLister) List(selector labels.Selector) (ret []*v1alpha2.Cell, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.Cell))
	})
	return ret, err
}

// Get retrieves the Cell from the indexer for a given namespace and name.
func (s cellNamespaceLister) Get(name string) (*v1alpha2.Cell, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha2.Resource("cell"), name)
	}
	return obj.(*v1alpha2.Cell), nil
}
