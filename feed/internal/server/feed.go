package server

import (
	"context"
	"github.com/JirafaYe/feed/internal/service"
	"time"
)

const (
	MaxNumVideos = 1
)

type FeedServer struct {
	service.UnimplementedFeedServer
}

func (f *FeedServer) FeedVideo(ctx context.Context, request *service.TiktokFeedRequest) (*service.TiktokFeedResponse, error) {
	var response service.TiktokFeedResponse
	// 先查30个视频
	t := time.Unix(*request.LastTime, 0)
	videos := m.localer.QueryVideosAfter(MaxNumVideos, t)
	for _, v := range videos {
		response.VideoList = append(response.VideoList, &service.Video{
			Id:            int64(v.ID),
			Author:        &service.User{Name: m.localer.QueryNameById(v.UserId)},
			PlayUrl:       v.PlayURL,
			CoverUrl:      v.CoverURL,
			CommentCount:  v.CommentCount,
			FavoriteCount: v.FavoriteCount,
			IsFavorite:    v.IsFavorite == 1,
			Title:         v.Title,
		})
	}
	return &response, nil
}
