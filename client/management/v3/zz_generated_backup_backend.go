package client

const (
	BackupBackendType                 = "backupBackend"
	BackupBackendFieldName            = "name"
	BackupBackendFieldS3BackupBackend = "s3BackupBackend"
)

type BackupBackend struct {
	Name            string           `json:"name,omitempty" yaml:"name,omitempty"`
	S3BackupBackend *S3BackupBackend `json:"s3BackupBackend,omitempty" yaml:"s3BackupBackend,omitempty"`
}
