package server

import (
	"context"
	"fmt"
	"github.com/JirafaYe/feed/internal/service"
	util "github.com/JirafaYe/feed/pkg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
	"time"
)

func TestFeedServer_FeedVideo(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:8899", grpc.WithTransportCredentials(insecure.NewCredentials()))
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

}

func TestTime(t *testing.T) {
	tt := time.UnixMilli(1676973263380)
	fmt.Println(tt)
}
