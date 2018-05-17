package client

const (
	ClusterRegistryStatusType        = "clusterRegistryStatus"
	ClusterRegistryStatusFieldStatus = "status"
)

type ClusterRegistryStatus struct {
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
