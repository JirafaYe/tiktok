package server

import (
	"context"
	"github.com/JirafaYe/feed/internal/service"
	"github.com/JirafaYe/feed/internal/store/obs"
	"github.com/JirafaYe/feed/pkg"
	"time"
)

const (
	MaxNumVideos = 1
)

type FeedServer struct {
	service.UnimplementedFeedServer
}

// FeedVideo
// TODO: 添加对token和time的错误处理
func (f *FeedServer) FeedVideo(ctx context.Context, request *service.TiktokFeedRequest) (*service.TiktokFeedResponse, error) {
	var response service.TiktokFeedResponse
	response.StatusCode = 0
	response.StatusMsg = util.NewString("OK")
	t := time.UnixMilli(*request.LatestTime)
	videos := m.localer.QueryVideosAfter(MaxNumVideos, t)
	for _, v := range videos {
		fmt.Println(v)
	}
	latestTime := time.Time{}.UnixMilli()
	for _, v := range videos {
		user := m.localer.QueryUserById(int64(v.UserId))
		latestTime = util.Min(v.CreatedAt.UnixMilli(), latestTime)
		response.VideoList = append(response.VideoList, &service.VideoFeed{
			Id: v.UserId,
			Author: &service.UserFeed{
				Id:            int64(user.ID),
				Name:          user.Name,
				FollowerCount: &user.FollowerCount,
				FollowCount:   &user.FollowCount,
				IsFollow:      true,
			},
			PlayUrl:       obs.GetVideoPrefix() + v.PlayURL,
			CoverUrl:      obs.GetImagePrefix() + v.CoverURL,
			CommentCount:  v.CommentCount,
			FavoriteCount: v.FavoriteCount,
			IsFavorite:    v.IsFavorite == 1,
			Title:         v.Title,
		})
	}
	response.NextTime = &latestTime
	return &response, nil
}
