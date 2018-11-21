package v3

import (
	"github.com/rancher/norman/lifecycle"
	"k8s.io/apimachinery/pkg/runtime"
)

type EtcdBackupConfigLifecycle interface {
	Create(obj *EtcdBackupConfig) (runtime.Object, error)
	Remove(obj *EtcdBackupConfig) (runtime.Object, error)
	Updated(obj *EtcdBackupConfig) (runtime.Object, error)
}

type etcdBackupConfigLifecycleAdapter struct {
	lifecycle EtcdBackupConfigLifecycle
}

func (w *etcdBackupConfigLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*EtcdBackupConfig))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *etcdBackupConfigLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*EtcdBackupConfig))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *etcdBackupConfigLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*EtcdBackupConfig))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewEtcdBackupConfigLifecycleAdapter(name string, clusterScoped bool, client EtcdBackupConfigInterface, l EtcdBackupConfigLifecycle) EtcdBackupConfigHandlerFunc {
	adapter := &etcdBackupConfigLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *EtcdBackupConfig) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}
