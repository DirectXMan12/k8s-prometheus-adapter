/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was automatically generated by lister-gen

package internalversion

import (
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	metrics "k8s.io/metrics/pkg/apis/metrics"
)

// NodeMetricsLister helps list NodeMetricses.
type NodeMetricsLister interface {
	// List lists all NodeMetricses in the indexer.
	List(selector labels.Selector) (ret []*metrics.NodeMetrics, err error)
	// Get retrieves the NodeMetrics from the index for a given name.
	Get(name string) (*metrics.NodeMetrics, error)
	NodeMetricsListerExpansion
}

// nodeMetricsLister implements the NodeMetricsLister interface.
type nodeMetricsLister struct {
	indexer cache.Indexer
}

// NewNodeMetricsLister returns a new NodeMetricsLister.
func NewNodeMetricsLister(indexer cache.Indexer) NodeMetricsLister {
	return &nodeMetricsLister{indexer: indexer}
}

// List lists all NodeMetricses in the indexer.
func (s *nodeMetricsLister) List(selector labels.Selector) (ret []*metrics.NodeMetrics, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*metrics.NodeMetrics))
	})
	return ret, err
}

// Get retrieves the NodeMetrics from the index for a given name.
func (s *nodeMetricsLister) Get(name string) (*metrics.NodeMetrics, error) {
	key := &metrics.NodeMetrics{ObjectMeta: v1.ObjectMeta{Name: name}}
	obj, exists, err := s.indexer.Get(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(metrics.Resource("nodemetrics"), name)
	}
	return obj.(*metrics.NodeMetrics), nil
}
