package client

import (
	"github.com/rancher/norman/types"
)

const (
	EtcdBackupConfigType                 = "etcdBackupConfig"
	EtcdBackupConfigFieldAnnotations     = "annotations"
	EtcdBackupConfigFieldBackupBackend   = "backupBackend"
	EtcdBackupConfigFieldClusterID       = "clusterId"
	EtcdBackupConfigFieldCreated         = "created"
	EtcdBackupConfigFieldCreation        = "creation"
	EtcdBackupConfigFieldCreatorID       = "creatorId"
	EtcdBackupConfigFieldLabels          = "labels"
	EtcdBackupConfigFieldName            = "name"
	EtcdBackupConfigFieldOwnerReferences = "ownerReferences"
	EtcdBackupConfigFieldRemoved         = "removed"
	EtcdBackupConfigFieldRetention       = "retention"
	EtcdBackupConfigFieldStatus          = "status"
	EtcdBackupConfigFieldUUID            = "uuid"
)

type EtcdBackupConfig struct {
	types.Resource
	Annotations     map[string]string       `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	BackupBackend   *BackupBackend          `json:"backupBackend,omitempty" yaml:"backupBackend,omitempty"`
	ClusterID       string                  `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`
	Created         string                  `json:"created,omitempty" yaml:"created,omitempty"`
	Creation        string                  `json:"creation,omitempty" yaml:"creation,omitempty"`
	CreatorID       string                  `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Labels          map[string]string       `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name            string                  `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences []OwnerReference        `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Removed         string                  `json:"removed,omitempty" yaml:"removed,omitempty"`
	Retention       string                  `json:"retention,omitempty" yaml:"retention,omitempty"`
	Status          *EtcdBackupConfigStatus `json:"status,omitempty" yaml:"status,omitempty"`
	UUID            string                  `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type EtcdBackupConfigCollection struct {
	types.Collection
	Data   []EtcdBackupConfig `json:"data,omitempty"`
	client *EtcdBackupConfigClient
}

type EtcdBackupConfigClient struct {
	apiClient *Client
}

type EtcdBackupConfigOperations interface {
	List(opts *types.ListOpts) (*EtcdBackupConfigCollection, error)
	Create(opts *EtcdBackupConfig) (*EtcdBackupConfig, error)
	Update(existing *EtcdBackupConfig, updates interface{}) (*EtcdBackupConfig, error)
	Replace(existing *EtcdBackupConfig) (*EtcdBackupConfig, error)
	ByID(id string) (*EtcdBackupConfig, error)
	Delete(container *EtcdBackupConfig) error
}

func newEtcdBackupConfigClient(apiClient *Client) *EtcdBackupConfigClient {
	return &EtcdBackupConfigClient{
		apiClient: apiClient,
	}
}

func (c *EtcdBackupConfigClient) Create(container *EtcdBackupConfig) (*EtcdBackupConfig, error) {
	resp := &EtcdBackupConfig{}
	err := c.apiClient.Ops.DoCreate(EtcdBackupConfigType, container, resp)
	return resp, err
}

func (c *EtcdBackupConfigClient) Update(existing *EtcdBackupConfig, updates interface{}) (*EtcdBackupConfig, error) {
	resp := &EtcdBackupConfig{}
	err := c.apiClient.Ops.DoUpdate(EtcdBackupConfigType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *EtcdBackupConfigClient) Replace(obj *EtcdBackupConfig) (*EtcdBackupConfig, error) {
	resp := &EtcdBackupConfig{}
	err := c.apiClient.Ops.DoReplace(EtcdBackupConfigType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *EtcdBackupConfigClient) List(opts *types.ListOpts) (*EtcdBackupConfigCollection, error) {
	resp := &EtcdBackupConfigCollection{}
	err := c.apiClient.Ops.DoList(EtcdBackupConfigType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *EtcdBackupConfigCollection) Next() (*EtcdBackupConfigCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &EtcdBackupConfigCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *EtcdBackupConfigClient) ByID(id string) (*EtcdBackupConfig, error) {
	resp := &EtcdBackupConfig{}
	err := c.apiClient.Ops.DoByID(EtcdBackupConfigType, id, resp)
	return resp, err
}

func (c *EtcdBackupConfigClient) Delete(container *EtcdBackupConfig) error {
	return c.apiClient.Ops.DoResourceDelete(EtcdBackupConfigType, &container.Resource)
}
