package server

import (
	"context"
	"fmt"
	"github.com/JirafaYe/comment/internal/service"
	"github.com/JirafaYe/comment/internal/store/local"
	"log"
)

type CommentServer struct {
	service.UnimplementedCommentServer
}

func (c *CommentServer) OperateComment(ctx context.Context, req *service.CommentRequest) (*service.CommentOperationResponse, error) {
	comment := ConvertCommentRequest(req)

	//测试user
	user := &service.User{
		Id:   1,
		Name: "user",
	}

	var err error
	if req.ActionType == 1 {
		err = m.localer.InsertComment(&comment)
		if err != nil {
			log.Print("插入评论失败", err)
		}
	} else if req.ActionType == 2 {
		err = m.localer.DeleteComment(comment)
		if err != nil {
			log.Print("删除评论失败", err)
		}
	}

	return &service.CommentOperationResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		Comment:    ConvertCommentBody(comment, user),
	}, err
}

func ConvertCommentRequest(request *service.CommentRequest) local.Comment {
	return local.Comment{
		Id:       request.CommentId,
		AuthorId: request.AuthorId,
		VideoId:  request.VideoId,
		Msg:      request.Msg,
	}
}

func ConvertCommentBody(c local.Comment, u *service.User) *service.CommentBody {
	return &service.CommentBody{
		Id:         c.Id,
		User:       u,
		Content:    c.Msg,
		CreateDate: fmt.Sprintf("%d-%d", int(c.CreatedAt.Month()), c.CreatedAt.Day()),
	}
}
