package web

import (
	"context"
	"github.com/JirafaYe/tiktok/internal/app/service"
	"github.com/JirafaYe/tiktok/internal/pkg/center"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	NextTime   *int64   `json:"next_time"`   // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	StatusCode int64    `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string  `json:"status_msg"`  // 返回状态描述
	VideoList  []*Video `json:"video_list"`  // 视频列表
}

type Video struct {
	Author        User   `json:"author"`         // 视频作者信息
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	CoverURL      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	ID            int64  `json:"id"`             // 视频唯一标识
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string `json:"play_url"`       // 视频播放地址
	Title         string `json:"title"`          // 视频标题
}

// User 视频作者信息
type User struct {
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	ID            int64  `json:"id"`             // 用户id
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
	Name          string `json:"name"`           // 用户名称
}

func (m *Manager) RouteFeed() {
	m.handler.GET("/douyin/feed", m.feed)
}

func (m *Manager) feed(c *gin.Context) {
	latestTimeStr := c.Query("latest_time")
	latestTime, err := strconv.ParseInt(latestTimeStr, 10, 64)
	if latestTime <= 0 {
		latestTime = time.Now().UnixMilli()
	}
	token := c.Query("token")

	// 该函数会进行自动负载均衡并返回一个*grpc.ClientConn
	conn, err := center.Resolver("feed")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close()

	client := service.NewFeedClient(conn)
	rpcResponse, _ := client.FeedVideo(context.Background(), &service.TiktokFeedRequest{
		Token:      &token,
		LatestTime: &latestTime,
	})

	response := FeedResponse{
		StatusCode: int64(rpcResponse.StatusCode),
		StatusMsg:  rpcResponse.StatusMsg,
		NextTime:   rpcResponse.NextTime,
	}
	response.VideoList = make([]*Video, len(rpcResponse.VideoList))
	for i, v := range rpcResponse.VideoList {
		response.VideoList[i] = &Video{
			ID: v.Id,
			Author: User{
				ID:            v.Author.Id,
				Name:          v.Author.Name,
				FollowCount:   *v.Author.FollowCount,
				FollowerCount: *v.Author.FollowerCount,
				IsFollow:      v.Author.IsFollow,
			},
			PlayURL:       v.PlayUrl,
			CoverURL:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    v.IsFavorite,
			Title:         v.Title,
		}
	}

	c.JSON(int(response.StatusCode), response)
}
