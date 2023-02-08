package obs

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

type Manager struct {
	handler *minio.Client
}

// Minio对象初始化
func New() (*Manager, error) {
	handler, err := minio.New(C.Address, &minio.Options{
		Creds: credentials.NewStaticV4(
			C.SecretID,
			C.SecretKey,
			"",
		),
	})
	if err != nil {
		log.Printf(err)
	}
	return &Manager{
		handler: handler,
	}, err
}

// // minioClient == Manager.handler
