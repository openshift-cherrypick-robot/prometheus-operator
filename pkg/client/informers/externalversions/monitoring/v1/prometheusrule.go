// Copyright 2018 The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	monitoringv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	internalinterfaces "github.com/coreos/prometheus-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/coreos/prometheus-operator/pkg/client/listers/monitoring/v1"
	versioned "github.com/coreos/prometheus-operator/pkg/client/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PrometheusRuleInformer provides access to a shared informer and lister for
// PrometheusRules.
type PrometheusRuleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.PrometheusRuleLister
}

type prometheusRuleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPrometheusRuleInformer constructs a new informer for PrometheusRule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPrometheusRuleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPrometheusRuleInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPrometheusRuleInformer constructs a new informer for PrometheusRule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPrometheusRuleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MonitoringV1().PrometheusRules(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MonitoringV1().PrometheusRules(namespace).Watch(context.TODO(), options)
			},
		},
		&monitoringv1.PrometheusRule{},
		resyncPeriod,
		indexers,
	)
}

func (f *prometheusRuleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPrometheusRuleInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *prometheusRuleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&monitoringv1.PrometheusRule{}, f.defaultInformer)
}

func (f *prometheusRuleInformer) Lister() v1.PrometheusRuleLister {
	return v1.NewPrometheusRuleLister(f.Informer().GetIndexer())
}
