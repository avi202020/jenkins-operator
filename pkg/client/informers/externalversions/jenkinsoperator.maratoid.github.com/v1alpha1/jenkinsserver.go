/*
Copyright 2018 Samsung SDS Cloud Native Computing Team.

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
package v1alpha1

import (
	time "time"

	jenkinsoperator_maratoid_github_com_v1alpha1 "github.com/maratoid/jenkins-operator/pkg/apis/jenkinsoperator.maratoid.github.com/v1alpha1"
	versioned "github.com/maratoid/jenkins-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/maratoid/jenkins-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/maratoid/jenkins-operator/pkg/client/listers/jenkinsoperator.maratoid.github.com/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// JenkinsServerInformer provides access to a shared informer and lister for
// JenkinsServers.
type JenkinsServerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.JenkinsServerLister
}

type jenkinsServerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewJenkinsServerInformer constructs a new informer for JenkinsServer type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewJenkinsServerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredJenkinsServerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredJenkinsServerInformer constructs a new informer for JenkinsServer type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredJenkinsServerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.JenkinsoperatorV1alpha1().JenkinsServers(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.JenkinsoperatorV1alpha1().JenkinsServers(namespace).Watch(options)
			},
		},
		&jenkinsoperator_maratoid_github_com_v1alpha1.JenkinsServer{},
		resyncPeriod,
		indexers,
	)
}

func (f *jenkinsServerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredJenkinsServerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *jenkinsServerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&jenkinsoperator_maratoid_github_com_v1alpha1.JenkinsServer{}, f.defaultInformer)
}

func (f *jenkinsServerInformer) Lister() v1alpha1.JenkinsServerLister {
	return v1alpha1.NewJenkinsServerLister(f.Informer().GetIndexer())
}
