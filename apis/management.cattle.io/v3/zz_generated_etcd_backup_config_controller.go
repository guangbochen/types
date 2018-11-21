package v3

import (
	"context"

	"github.com/rancher/norman/controller"
	"github.com/rancher/norman/objectclient"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

var (
	EtcdBackupConfigGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "EtcdBackupConfig",
	}
	EtcdBackupConfigResource = metav1.APIResource{
		Name:         "etcdbackupconfigs",
		SingularName: "etcdbackupconfig",
		Namespaced:   false,
		Kind:         EtcdBackupConfigGroupVersionKind.Kind,
	}
)

type EtcdBackupConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EtcdBackupConfig
}

type EtcdBackupConfigHandlerFunc func(key string, obj *EtcdBackupConfig) (runtime.Object, error)

type EtcdBackupConfigLister interface {
	List(namespace string, selector labels.Selector) (ret []*EtcdBackupConfig, err error)
	Get(namespace, name string) (*EtcdBackupConfig, error)
}

type EtcdBackupConfigController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() EtcdBackupConfigLister
	AddHandler(ctx context.Context, name string, handler EtcdBackupConfigHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler EtcdBackupConfigHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type EtcdBackupConfigInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*EtcdBackupConfig) (*EtcdBackupConfig, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*EtcdBackupConfig, error)
	Get(name string, opts metav1.GetOptions) (*EtcdBackupConfig, error)
	Update(*EtcdBackupConfig) (*EtcdBackupConfig, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*EtcdBackupConfigList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() EtcdBackupConfigController
	AddHandler(ctx context.Context, name string, sync EtcdBackupConfigHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle EtcdBackupConfigLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync EtcdBackupConfigHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle EtcdBackupConfigLifecycle)
}

type etcdBackupConfigLister struct {
	controller *etcdBackupConfigController
}

func (l *etcdBackupConfigLister) List(namespace string, selector labels.Selector) (ret []*EtcdBackupConfig, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*EtcdBackupConfig))
	})
	return
}

func (l *etcdBackupConfigLister) Get(namespace, name string) (*EtcdBackupConfig, error) {
	var key string
	if namespace != "" {
		key = namespace + "/" + name
	} else {
		key = name
	}
	obj, exists, err := l.controller.Informer().GetIndexer().GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(schema.GroupResource{
			Group:    EtcdBackupConfigGroupVersionKind.Group,
			Resource: "etcdBackupConfig",
		}, key)
	}
	return obj.(*EtcdBackupConfig), nil
}

type etcdBackupConfigController struct {
	controller.GenericController
}

func (c *etcdBackupConfigController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *etcdBackupConfigController) Lister() EtcdBackupConfigLister {
	return &etcdBackupConfigLister{
		controller: c,
	}
}

func (c *etcdBackupConfigController) AddHandler(ctx context.Context, name string, handler EtcdBackupConfigHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*EtcdBackupConfig); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *etcdBackupConfigController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler EtcdBackupConfigHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*EtcdBackupConfig); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type etcdBackupConfigFactory struct {
}

func (c etcdBackupConfigFactory) Object() runtime.Object {
	return &EtcdBackupConfig{}
}

func (c etcdBackupConfigFactory) List() runtime.Object {
	return &EtcdBackupConfigList{}
}

func (s *etcdBackupConfigClient) Controller() EtcdBackupConfigController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.etcdBackupConfigControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(EtcdBackupConfigGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &etcdBackupConfigController{
		GenericController: genericController,
	}

	s.client.etcdBackupConfigControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type etcdBackupConfigClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   EtcdBackupConfigController
}

func (s *etcdBackupConfigClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *etcdBackupConfigClient) Create(o *EtcdBackupConfig) (*EtcdBackupConfig, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*EtcdBackupConfig), err
}

func (s *etcdBackupConfigClient) Get(name string, opts metav1.GetOptions) (*EtcdBackupConfig, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*EtcdBackupConfig), err
}

func (s *etcdBackupConfigClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*EtcdBackupConfig, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*EtcdBackupConfig), err
}

func (s *etcdBackupConfigClient) Update(o *EtcdBackupConfig) (*EtcdBackupConfig, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*EtcdBackupConfig), err
}

func (s *etcdBackupConfigClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *etcdBackupConfigClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *etcdBackupConfigClient) List(opts metav1.ListOptions) (*EtcdBackupConfigList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*EtcdBackupConfigList), err
}

func (s *etcdBackupConfigClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *etcdBackupConfigClient) Patch(o *EtcdBackupConfig, data []byte, subresources ...string) (*EtcdBackupConfig, error) {
	obj, err := s.objectClient.Patch(o.Name, o, data, subresources...)
	return obj.(*EtcdBackupConfig), err
}

func (s *etcdBackupConfigClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *etcdBackupConfigClient) AddHandler(ctx context.Context, name string, sync EtcdBackupConfigHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *etcdBackupConfigClient) AddLifecycle(ctx context.Context, name string, lifecycle EtcdBackupConfigLifecycle) {
	sync := NewEtcdBackupConfigLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *etcdBackupConfigClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync EtcdBackupConfigHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *etcdBackupConfigClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle EtcdBackupConfigLifecycle) {
	sync := NewEtcdBackupConfigLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}
