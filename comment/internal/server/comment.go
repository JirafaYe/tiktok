package server

import (
	"context"
	"github.com/JirafaYe/comment/internal/service"
	"github.com/JirafaYe/comment/internal/store/local"
	"log"
)

type CommentServer struct {
	service.UnimplementedCommentServer
}

func (c *CommentServer) OperateComment(ctx context.Context, req *service.CommentRequest) (*service.CommentOperationResponse, error) {
	comment := ConvertCommentRequest(req)
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
	reponse := &service.CommentOperationResponse{
		Comment: ConvertCommentBody(comment),
	}

	return reponse, err
}

func ConvertCommentRequest(request *service.CommentRequest) local.Comment {
	return local.Comment{
		AuthorId: request.AuthorId,
		VideoId:  request.VideoId,
		Msg:      request.Msg,
	}
}

func ConvertCommentBody(c local.Comment) *service.CommentBody {
	return &service.CommentBody{
		Id:         c.Id,
		Content:    c.Msg,
		CreateDate: c.CreatedAt.String(),
	}
}
