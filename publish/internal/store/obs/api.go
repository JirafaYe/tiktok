package obs

import (
	"context"
	"io"
	"net/url"
	"time"
	"log"
	"github.com/minio/minio-go/v7"
)

// minioCLient == manager.handler

// UploadFile 上传视频的request的protocol buffer中有data字段，用reader去读取，然后上传至minio
func (m *Manager)UploadFile(bucketName string, objectName string, reader io.Reader, objectsize int64) error {
	ctx := context.Background()
	n, err := m.handler.PutObject(ctx, bucketName, objectName, reader, objectsize, minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
	if err != nil {
		log.Printf("upload %s of size %d failed, %s", bucketName, objectsize, err)
		return err
	}
	log.Printf("upload %s of bytes %d successfully", objectName, n.Size)
	return nil
}

// GetFileURL 从minio获取文件URL
func (m *Manager)GetFileURL(bucketName string, fileName string, expires time.Duration) (*url.URL, error) {
	ctx := context.Background()
	reqParams := make(url.Values)
	if expires <= 0 {
		expires = time.Second * 60 * 60 *24
	}
	presignedURL, err := m.handler.PresignedGetObject(ctx, bucketName, fileName, expires, reqParams)
	if err != nil {
		log.Printf("get url of file %s from bucket %s failed, %s", fileName, bucketName, err)
		return nil, err
	}
	// TODO URL可能要做截取
	return presignedURL, nil
}


