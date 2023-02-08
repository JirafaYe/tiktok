package server

import (
	"context"
	"github.com/JirafaYe/user/internal/service"
	"github.com/JirafaYe/user/pkg"
)

type UserSrv struct {
	service.UnimplementedUserProtoServer
}

func (u *UserSrv) Register(_ context.Context, request *service.RegisterRequest) (*service.RegisterResponse, error) {
	username, password := request.Username, request.Password
	if exit := m.localer.GetUsernameExit(username); exit {
		return &service.RegisterResponse{
			StatusCode: 1,
			StatusMsg:  "用户名已存在",
			UserId:     0,
			Token:      "",
		}, nil
	}
	if id, register := m.localer.Register(username, password); register {
		token := setUserToken(id, username)
		return &service.RegisterResponse{
			StatusCode: 0,
			StatusMsg:  "注册成功",
			UserId:     id,
			Token:      token,
		}, nil
	} else {
		return &service.RegisterResponse{
			StatusCode: 1,
			StatusMsg:  "注册失败",
			UserId:     id,
			Token:      "",
		}, nil
	}
}

// 设置用户token
func setUserToken(id int64, username string) string {
	token, err := pkg.GenerateToken(id, username)
	if err != nil {
		panic(err)
	}
	err = m.cacher.SetUserToken(token, username)
	if err != nil {
		panic(err)
	}
	return token
}

func parseToken(token string) string {
	claims, err := pkg.ParseToken(token)
	if err != nil {
		panic(err)
	}
	return claims.Username
}

func (u *UserSrv) Login(_ context.Context, request *service.LoginRequest) (*service.LoginResponse, error) {
	username, password := request.Username, request.Password
	login, id, err := m.localer.Login(username, password)
	if login {
		token := setUserToken(id, username)
		return &service.LoginResponse{
			StatusCode: 0,
			StatusMsg:  "登录成功",
			UserId:     id,
			Token:      token,
		}, nil
	} else {
		return &service.LoginResponse{
			StatusCode: 1,
			StatusMsg:  "登录失败",
			UserId:     0,
			Token:      "",
		}, err
	}
}

func (u *UserSrv) IsLogin(_ context.Context, request *service.IsLoginRequest) (*service.IsLoginResponse, error) {
	token := request.Token
	username := parseToken(token)
	if exist := m.cacher.IsUserTokenExist(username); exist {
		return &service.IsLoginResponse{
			Code: 0,
			Msg:  "获取成功",
		}, nil
	} else {
		return &service.IsLoginResponse{
			Code: 1,
			Msg:  "获取失败",
		}, nil
	}
}

func (u *UserSrv) GetUserMsg(_ context.Context, request *service.UserRequest) (*service.UserResponse, error) {
	token := request.Token
	claims, err := pkg.ParseToken(token)
	if err != nil {
		return &service.UserResponse{
			StatusCode: 1,
			StatusMsg:  "token解析错误",
			User:       nil,
		}, err
	}
	username := claims.Username
	userMsg := m.localer.GetUserMsg(username)
	usg := &service.UserMsg{
		Id:            userMsg.ID,
		Name:          userMsg.Name,
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	}
	return &service.UserResponse{
		StatusCode: 0,
		StatusMsg:  "获取成功",
		User:       usg,
	}, nil
}
