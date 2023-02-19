package web

import (
	"context"
	"github.com/JirafaYe/tiktok/internal/app/service"
	"github.com/JirafaYe/tiktok/internal/pkg/center"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (m *Manager) RouteFavorite() {
	m.handler.POST("/douyin/favorite/action", m.favoriteAction)
	m.handler.GET("/douyin/favorite/list", m.favoriteList)
}

func (m *Manager) favoriteAction(ctx *gin.Context) {
	token := ctx.Query("token")
	vId := ctx.Query("video_id")
	aType := ctx.Query("action_type")

	if token == "" || vId == "" || aType == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  "Invalid request",
		})
		return
	}

	videoId, _ := strconv.Atoi(vId)
	actionType, _ := strconv.Atoi(aType)

	// connect grpc server
	conn, err := center.Resolver("favorite")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  err.Error(),
		})
		return
	}
	defer conn.Close()

	client := service.NewFavoriteClient(conn)

	res, err := client.FavoriteAction(context.Background(), &service.FavoriteActionRequest{
		Token:      token,
		VideoId:    int64(videoId),
		ActionType: int32(actionType),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": res.StatusCode,
		"status_msg":  res.StatusMsg,
	})
}

func (m *Manager) favoriteList(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Query("user_id"))
	token := ctx.Query("token")

	// connect grpc server
	conn, err := center.Resolver("favorite")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  err.Error(),
		})
		return
	}
	defer conn.Close()

	client := service.NewFavoriteClient(conn)

	res, err := client.GetFavoriteList(context.Background(), &service.FavoriteListRequest{
		UserId: int64(userId),
		Token:  token,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": res.StatusCode,
		"status_msg":  res.StatusMsg,
		"video_list":  res.VideoList,
	})
}
