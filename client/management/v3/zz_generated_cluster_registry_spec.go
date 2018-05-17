package client

const (
	ClusterRegistrySpecType             = "clusterRegistrySpec"
	ClusterRegistrySpecFieldClusterId   = "clusterId"
	ClusterRegistrySpecFieldDisplayName = "displayName"
)

type ClusterRegistrySpec struct {
	ClusterId   string `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
}
