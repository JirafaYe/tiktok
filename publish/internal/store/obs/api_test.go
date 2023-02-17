package obs

import (
    "testing"
	"io/ioutil"
	"bytes"
	"time"
)

func TestUploadFile(t *testing.T) {
    // 初始化Manager实例进行测试
    manager, _ := New()
    // 使用模拟文件数据创建 io.Reader
	data, err := ioutil.ReadFile("test1.mp4")
	if err!= nil {
		t.Errorf("can't read video file by filePath: %v", err)
	}
    size := int64(len(data))
    reader := bytes.NewReader(data)

    // 调用 UploadFile 并断言正确性
    if err := manager.UploadFile("videos", "test1.mp4", reader, size); err != nil {
        t.Errorf("unexpected error: %v", err)
    }
}

func TestGetFileURL(t *testing.T) {
    // 初始化Manager实例进行测试
    manager, _ := New()
    bucketName := "videos"
    fileName := "test1.mp4"
    expires := time.Second * 60 * 60 * 24

    url, err := manager.GetFileURL(bucketName, fileName, expires)
    if err != nil {
        t.Errorf("get url of file %s from bucket %s failed: %s", fileName, bucketName, err)
    }
    if url == nil {
        t.Errorf("url is nil")
    }
	t.Logf("get url of file %s from bucket %s: %s", fileName, bucketName, url)
}

