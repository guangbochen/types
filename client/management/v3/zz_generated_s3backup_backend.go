package client

const (
	S3BackupBackendType                  = "s3BackupBackend"
	S3BackupBackendFieldAccessKeyID      = "accessKeyId"
	S3BackupBackendFieldBucketName       = "bucketName"
	S3BackupBackendFieldEndpoint         = "endpoint"
	S3BackupBackendFieldRegion           = "region"
	S3BackupBackendFieldSecretAccesssKey = "secretAccessKey"
)

type S3BackupBackend struct {
	AccessKeyID      string `json:"accessKeyId,omitempty" yaml:"accessKeyId,omitempty"`
	BucketName       string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`
	Endpoint         string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`
	Region           string `json:"region,omitempty" yaml:"region,omitempty"`
	SecretAccesssKey string `json:"secretAccessKey,omitempty" yaml:"secretAccessKey,omitempty"`
}
