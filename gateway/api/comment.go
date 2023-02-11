package api

import (
	"context"
	"github.com/JirafaYe/gateway/center"
	"github.com/JirafaYe/gateway/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (m *Manager) RouteComment() {
	group := m.handler.Group("/douyin/comment")
	group.POST("/action", m.Action)
}

type CommentOperationResponse struct {
	StatusCode int32       `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	Comment    CommentBody `json:"comment"`
}

type CommentBody struct {
	Id         int32  `json:"id"`
	User       User   `json:"user"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}

func (m *Manager) Action(ctx *gin.Context) {
	log.Printf("请求评论操作")
	videoId, _ := strconv.Atoi(ctx.Query("video_id"))
	actionType, _ := strconv.Atoi(ctx.Query("action_type"))
	var msg string
	var commentId int
	if actionType == 1 {
		msg = ctx.Query("comment_text")
		if msg == "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status_code": http.StatusInternalServerError,
				"status_msg":  "评论为空",
			})
			return
		}
	} else if actionType == 2 {
		id := ctx.Query("comment_id")
		if id == "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status_code": http.StatusInternalServerError,
				"status_msg":  "CommentId为空",
			})
			return
		}
		commentId, _ = strconv.Atoi(id)
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  "actionType错误",
		})
		return
	}

	// center.Resolver() 参数为调用的服务名
	// 该函数会进行自动负载均衡并返回一个*grpc.ClientConn
	conn, err := center.Resolver("comment")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  err.Error(),
		})
		return
	}
	defer conn.Close()

	client := service.NewCommentClient(conn)

	resp, err := client.OperateComment(context.Background(), &service.CommentRequest{
		VideoId:    int32(videoId),
		ActionType: int32(actionType),
		CommentId:  int32(commentId),
		Msg:        msg,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  err.Error(),
		})
		return
	}

	body := CommentBody{
		Id: resp.Comment.Id,
		User: User{
			FollowCount:   resp.Comment.User.FollowerCount,
			FollowerCount: resp.Comment.User.FollowCount,
			ID:            resp.Comment.User.Id,
			IsFollow:      resp.Comment.User.IsFollow,
			Name:          resp.Comment.User.Name,
		},
		Content:    resp.Comment.Content,
		CreateDate: resp.Comment.CreateDate,
	}
	ctx.JSON(http.StatusOK, CommentOperationResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  resp.StatusMsg,
		Comment:    body,
	})
}
