package client

const (
	EtcdBackupConfigStatusType            = "etcdBackupConfigStatus"
	EtcdBackupConfigStatusFieldConditions = "conditions"
)

type EtcdBackupConfigStatus struct {
	Conditions []EtcdBackupConfigCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`
}
