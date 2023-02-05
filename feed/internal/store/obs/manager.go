package obs

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Manager struct {
	handle *minio.Client
}

func New() (*Manager, error) {
	client, err := minio.New(C.Address, &minio.Options{
		Creds:  credentials.NewStaticV4(C.SecretId, C.SecretKey, ""),
		Secure: false})
	return &Manager{client}, err
}
