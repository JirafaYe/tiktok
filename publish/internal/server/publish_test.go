package server

import (
	"context"
	"fmt"
	"github.com/JirafaYe/publish/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

func TestPublishAction(t *testing.T) {
	conn, err := grpc.Dial("192.168.79.83:8899", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	c := service.NewFeedClient(conn)
	var req service.TiktokFeedRequest
	ti := time.Now().Unix()
	req.LatestTime = &ti
	req.Token = util.NewString("")
	r, err := c.FeedVideo(context.Background(), &req)
	fmt.Println(r)

	client := service.NewPublishClient(conn)
	// TODO: 设置request
	var request service.PublishVideoRequest
	request.Token = util.NewString("325135252354252")
    request.Title = util.NewString("good")
	request.Data = util.GetVideo("/home/mzz/gowork/gitwork/tiktok/publish/pkg/util/test.mp4")//[]byte
	response, err := client.PublishVideo(context.Background(), &request)
	if err!= nil {
        fmt.Println(err)
    }
}
