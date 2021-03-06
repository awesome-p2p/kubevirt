package util

import (
	"k8s.io/client-go/pkg/api"
	"k8s.io/client-go/pkg/fields"
	"k8s.io/client-go/tools/cache"
	"kubevirt.io/kubevirt/pkg/api/v1"
	"kubevirt.io/kubevirt/pkg/kubecli"
)

func NewVMCache() (cache.SharedInformer, error) {
	restClient, err := kubecli.GetRESTClient()
	if err != nil {
		return nil, err
	}
	vmCacheSource := cache.NewListWatchFromClient(restClient, "vms", api.NamespaceDefault, fields.Everything())
	informer := cache.NewSharedInformer(vmCacheSource, &v1.VM{}, 0)
	return informer, nil
}
