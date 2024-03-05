/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	time "time"

	nuclioiov1beta1 "github.com/nuclio/nuclio/pkg/platform/kube/apis/nuclio.io/v1beta1"
	versioned "github.com/nuclio/nuclio/pkg/platform/kube/client/clientset/versioned"
	internalinterfaces "github.com/nuclio/nuclio/pkg/platform/kube/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/nuclio/nuclio/pkg/platform/kube/client/listers/nuclio.io/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// NuclioAPIGatewayInformer provides access to a shared informer and lister for
// NuclioAPIGateways.
type NuclioAPIGatewayInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.NuclioAPIGatewayLister
}

type nuclioAPIGatewayInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewNuclioAPIGatewayInformer constructs a new informer for NuclioAPIGateway type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNuclioAPIGatewayInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNuclioAPIGatewayInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredNuclioAPIGatewayInformer constructs a new informer for NuclioAPIGateway type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNuclioAPIGatewayInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NuclioV1beta1().NuclioAPIGateways(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NuclioV1beta1().NuclioAPIGateways(namespace).Watch(context.TODO(), options)
			},
		},
		&nuclioiov1beta1.NuclioAPIGateway{},
		resyncPeriod,
		indexers,
	)
}

func (f *nuclioAPIGatewayInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNuclioAPIGatewayInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *nuclioAPIGatewayInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&nuclioiov1beta1.NuclioAPIGateway{}, f.defaultInformer)
}

func (f *nuclioAPIGatewayInformer) Lister() v1beta1.NuclioAPIGatewayLister {
	return v1beta1.NewNuclioAPIGatewayLister(f.Informer().GetIndexer())
}
