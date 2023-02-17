package obs

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Manager struct {
	handler *minio.Client
}

// Minio对象初始化
func New() (*Manager, error) {
	handler, err := minio.New(C.Address, &minio.Options{
		Creds:  credentials.NewStaticV4(C.SecretId, C.SecretKey, ""),
		Secure: false})
	return &Manager{
		handler: handler,
	}, err
}

// // minioClient == Manager.handler
