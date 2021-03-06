// Code generated by informer-gen. DO NOT EDIT.

package internalversion

import (
	time "time"

	network "github.com/openshift/origin/pkg/network/apis/network"
	internalinterfaces "github.com/openshift/origin/pkg/network/generated/informers/internalversion/internalinterfaces"
	internalclientset "github.com/openshift/origin/pkg/network/generated/internalclientset"
	internalversion "github.com/openshift/origin/pkg/network/generated/listers/network/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// EgressNetworkPolicyInformer provides access to a shared informer and lister for
// EgressNetworkPolicies.
type EgressNetworkPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.EgressNetworkPolicyLister
}

type egressNetworkPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewEgressNetworkPolicyInformer constructs a new informer for EgressNetworkPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewEgressNetworkPolicyInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredEgressNetworkPolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredEgressNetworkPolicyInformer constructs a new informer for EgressNetworkPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredEgressNetworkPolicyInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Network().EgressNetworkPolicies(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Network().EgressNetworkPolicies(namespace).Watch(options)
			},
		},
		&network.EgressNetworkPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *egressNetworkPolicyInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredEgressNetworkPolicyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *egressNetworkPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&network.EgressNetworkPolicy{}, f.defaultInformer)
}

func (f *egressNetworkPolicyInformer) Lister() internalversion.EgressNetworkPolicyLister {
	return internalversion.NewEgressNetworkPolicyLister(f.Informer().GetIndexer())
}
