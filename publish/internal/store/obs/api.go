package obs

import (
	"context"
	"io"
	"net/url"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
)

// minioCLient == manager.handler


// UploadLocalFile 上传本地文件（提供文件路径）至 minio
/*
* filePath: 本地文件路径
* contentType: 上传文件类型
*/
func UploadLocalFile(bucketName string, objectName string, filePath string, contentType string) (int64, error) {
	ctx := context.Background()
	info, err := m.handler.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		klog.Errorf("localfile upload failed, %s", err)
		return 0, err
	}
	klog.Infof("upload %s of size %d successfully", objectName, info.Size)
	return info.Size, nil
}

// UploadFile 上传视频的request的protocol buffer中有data字段，用reader去读取，然后上传至minio
// UploadFile 上传文件（提供reader）至 minio
func UploadFile(bucketName string, objectName string, reader io.Reader, objectsize int64) error {
	ctx := context.Background()
	n, err := m.handler.PutObject(ctx, bucketName, objectName, reader, objectsize, minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
	if err != nil {
		klog.Errorf("upload %s of size %d failed, %s", bucketName, objectsize, err)
		return err
	}
	klog.Infof("upload %s of bytes %d successfully", objectName, n.Size)
	return nil
}

// GetFileURL 从minio获取文件URL
func GetFileURL(bucketName string, fileName string, expires time.Duration) (*url.URL, error) {
	ctx := context.Background()
	reqParams := make(url.Values)
	if expires <= 0 {
		expires = time.Second * 60 * 60 *24
	}
	presignedURL, err := m.handler.PresignedGetObject(ctx, bucketName, fileName, expires, reqParams)
	if err != nil {
		klog.Errorf("get url of file %s from bucket %s failed, %s", fileName, bucketName, err)
		return nil, err
	}
	// TODO URL可能要做截取
	return presignedURL, nil
}


