package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/JirafaYe/comment/internal/service"
	"github.com/JirafaYe/comment/internal/store/local"
	"github.com/JirafaYe/comment/pkg"
	"log"
)

type CommentServer struct {
	service.UnimplementedCommentServer
}

func (c *CommentServer) ListComments(ctx context.Context, req *service.ListRequest) (*service.ListCommentsResponse, error) {
	resp, paresErr := pkg.ParseToken(req.Token)
	if paresErr != nil {
		log.Println("解析token失败", paresErr)
		return nil, errors.New("解析token失败")
	} else if !m.cacher.IsUserTokenExist(resp.Username) {
		log.Println("用户token不存在")
		return nil, errors.New("token登录验证失败")
	}

	commentList, err := m.localer.SelectCommentListByVideoId(req.VideoId)
	if err != nil {
		log.Println("获取评论列表失败", err)
		return nil, errors.New("获取评论列表失败")
	}

	var ids []int64
	list := make([]*service.CommentBody, len(commentList))
	userMap := make(map[int64]*service.CommentUser)
	id := make(map[int64]interface{})

	for _, comment := range commentList {
		id[comment.AuthorId] = nil
	}

	for k, _ := range id {
		ids = append(ids, k)
	}

	msg, err := m.localer.GetUserMsg(ids)
	if err != nil {
		log.Println("获取用户信息失败")
		return nil, errors.New("获取用户信息失败")
	}

	for _, user := range msg {
		log.Println(user)
		userMap[user.Id] = &user
	}

	for i, comment := range commentList {
		log.Println(userMap[comment.AuthorId], comment.AuthorId)
		list[i] = ConvertCommentBody(comment, userMap[comment.AuthorId])
	}

	return &service.ListCommentsResponse{
		StatusCode:  0,
		StatusMsg:   "success",
		CommentList: list,
	}, nil
}

func (c *CommentServer) OperateComment(ctx context.Context, req *service.CommentRequest) (*service.CommentOperationResponse, error) {
	comment := ConvertCommentRequest(req)

	resp, paresErr := pkg.ParseToken(req.Token)
	if paresErr != nil {
		log.Println("解析token失败", paresErr)
		return nil, errors.New("解析token失败")
	} else if !m.cacher.IsUserTokenExist(resp.Username) {
		log.Println("用户token不存在")
		return nil, errors.New("token登录验证失败")
	}

	user := &service.CommentUser{
		Id:   resp.Id,
		Name: resp.Username,
	}

	var err error
	if req.ActionType == 1 {
		err = m.localer.InsertComment(&comment)
		if err != nil {
			log.Print("插入评论失败", err)
			return nil, errors.New("插入评论失败")
		}
		go m.localer.UpdateCommentsCountByVideoId(comment.VideoId, 1)
	} else if req.ActionType == 2 {
		err = m.localer.DeleteComment(comment)
		if err != nil {
			log.Print("删除评论失败: ", err)
			return nil, errors.New("删除评论失败")
		}
		go m.localer.UpdateCommentsCountByVideoId(comment.VideoId, -1)
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

func ConvertCommentBody(c local.Comment, u *service.CommentUser) *service.CommentBody {
	return &service.CommentBody{
		Id:         c.Id,
		User:       u,
		Content:    c.Msg,
		CreateDate: fmt.Sprintf("%d-%d", int(c.CreatedAt.Month()), c.CreatedAt.Day()),
	}
}
