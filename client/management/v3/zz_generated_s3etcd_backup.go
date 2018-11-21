package client

const (
	S3EtcdBackupType                 = "s3EtcdBackup"
	S3EtcdBackupFieldFileName        = "fileName"
	S3EtcdBackupFieldS3BackupBackend = "s3BackupBackend"
)

type S3EtcdBackup struct {
	FileName        string           `json:"fileName,omitempty" yaml:"fileName,omitempty"`
	S3BackupBackend *S3BackupBackend `json:"s3BackupBackend,omitempty" yaml:"s3BackupBackend,omitempty"`
}
