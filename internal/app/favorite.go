package app

import (
	"context"
	"github.com/JirafaYe/tiktok/internal/app/service"
	"github.com/JirafaYe/tiktok/internal/pkg/token"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type FavoriteServer struct {
	service.UnimplementedFavoriteServer
}

func (s *FavoriteServer) FavoriteAction(ctx context.Context, request *service.FavoriteActionRequest) (*service.FavoriteActionResponse, error) {
	// verify token
	claim, err := token.ParseToken(request.Token)
	if err != nil || m.cacher.IsTokenExist(request.Token) {
		log.Println(err.Error())
		return ConvertActionResponse(http.StatusForbidden, "token已过期", err)
	}
	switch request.ActionType {
	case 1:
		// like
		_, err := m.localer.SelectUserFavorite(claim.Id, request.VideoId)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				// 没有点赞记录
				err = m.localer.InsertUserFavorite(claim.Id, request.VideoId)
				if err != nil {
					return ConvertActionResponse(http.StatusInternalServerError, err.Error(), err)
				}
				err = m.localer.UpdateVideoLike(request.VideoId, 1)
				if err != nil {
					return ConvertActionResponse(http.StatusInternalServerError, err.Error(), err)
				}
			} else {
				return ConvertActionResponse(http.StatusInternalServerError, err.Error(), err)
			}
		}
	case 2:
		// cancel like
		_, err := m.localer.SelectUserFavorite(claim.Id, request.VideoId)
		if err == nil {
			// 查询到记录
			err = m.localer.DeleteUserFavorite(claim.Id, request.VideoId)
			if err != nil {
				return ConvertActionResponse(http.StatusInternalServerError, err.Error(), err)
			}
			err = m.localer.UpdateVideoLike(request.VideoId, -1)
			if err != nil {
				return ConvertActionResponse(http.StatusInternalServerError, err.Error(), err)
			}
		}
	}
	return ConvertActionResponse(0, "success", nil)
}

func (s *FavoriteServer) GetFavoriteList(ctx context.Context, request *service.FavoriteListRequest) (*service.FavoriteListResponse, error) {
	// verify token
	claim, err := token.ParseToken(request.Token)
	if err != nil || m.cacher.IsTokenExist(request.Token) {
		log.Println(err.Error())
		return ConvertListResponse(http.StatusForbidden, "token已过期", nil, err)
	}
	// todo 分页？
	favorites, err := m.localer.SelectLikesByUser(claim.Id)
	if err != nil {
		return ConvertListResponse(http.StatusInternalServerError, err.Error(), nil, err)
	}
	localVideoList, err := m.localer.SelectVideos(favorites)
	if err != nil {
		return ConvertListResponse(http.StatusInternalServerError, err.Error(), nil, err)
	}
	// make map
	var authorIds []int64
	authorMap := make(map[int64]*service.UserFeed)
	for _, v := range localVideoList {
		if _, ok := authorMap[v.UserId]; !ok {
			authorIds = append(authorIds, v.UserId)
			authorMap[v.UserId] = &service.UserFeed{}
		}
	}
	authors, err := m.localer.SelectUsers(authorIds)
	if err != nil {
		return ConvertListResponse(http.StatusInternalServerError, err.Error(), nil, err)
	}
	for _, a := range authors {
		authorMap[a.Id] = &a
	}
	var videoList []*service.VideoFeed
	for _, v := range localVideoList {
		videoList = append(videoList, &service.VideoFeed{
			Id:            int64(v.ID),
			Author:        authorMap[v.UserId],
			PlayUrl:       v.PlayURL,
			CoverUrl:      v.CoverURL,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    v.IsFavorite,
			Title:         v.Title,
		})
	}
	return ConvertListResponse(0, "success", videoList, nil)
}

func ConvertActionResponse(status int32, msg string, err error) (*service.FavoriteActionResponse, error) {
	return &service.FavoriteActionResponse{
		StatusCode: status,
		StatusMsg:  &msg,
	}, err
}

func ConvertListResponse(status int32, msg string, videoList []*service.VideoFeed, err error) (*service.FavoriteListResponse, error) {
	return &service.FavoriteListResponse{
		StatusCode: status,
		StatusMsg:  &msg,
		VideoList:  videoList,
	}, err
}
