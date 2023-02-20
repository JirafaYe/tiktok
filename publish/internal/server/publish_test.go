package server

import (
    "testing"
	// "image/jpeg"
	// "image"
	// "os"

    "context"
    "fmt"
    "io/ioutil"

    // "github.com/JirafaYe/publish/internal/store/obs"
    // "github.com/JirafaYe/publish/internal/store/local"
    "github.com/JirafaYe/publish/internal/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestPubAction(t *testing.T) {
    conn, err := grpc.Dial("127.0.0.1:11451", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

    c := service.NewPublishClient(conn)
    tmpData, err := ioutil.ReadFile("test1.mp4")
    tmpRequest := &service.PublishActionRequest{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjI2NzY5Nzc5MDYwMTY2MjQ2NCwidXNlcm5hbWUiOiIxMzQ4NDc4NDAzQHFxLmNvbSIsImV4cCI6MTY3NzQyNjA4NSwiaXNzIjoidGlrdG9rLnVzZXIifQ.JGZqxB-unHziBS_AkxGGHJWJAR6VGYhTEFIQjrTXYfc",
		Data:  tmpData,
		Title: "the Galaxy new",
	}

    res, err := c.PubAction(context.Background(), tmpRequest)
    if err!= nil {
        t.Errorf("error creating: %v", err)
    }
    fmt.Println(res.StatusCode)
    fmt.Println(res.StatusMsg)
}

// func SaveImageAsJpeg(img image.Image, filename string) (err error) {
//     jpegFile, err := os.Create(filename)
//     defer jpegFile.Close()
//     if err != nil {
//         return err
//     }
//     err = jpeg.Encode(jpegFile, img, &jpeg.Options{Quality: 100})
//     return err
// }

// func TestReadFrameAsJpeg(t *testing.T) {
//     filePath := "http://47.108.66.104:9000/videos/test1.mp4"

//     img, err := readFrameAsJpeg(filePath)
//     if err != nil {
//         t.Errorf("Unexpected error %v", err)
//     }
// 	err = SaveImageAsJpeg(img, "test1.jpg")
// 	if err!= nil {
//         t.Errorf("fail save image as jpeg due to: %v\n", err)
//     }
// }