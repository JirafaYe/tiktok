package app

import (
	"context"
	"github.com/JirafaYe/tiktok/example/internal/pkg/auther"
	"github.com/JirafaYe/tiktok/example/internal/pkg/dba"
	"net/http"

	"github.com/JirafaYe/tiktok/example/internal/app/service"
	"gorm.io/gorm"
)

// UserSrv grpc中关于user相关的服务具体实现
type UserSrv struct {
	service.UnimplementedUserServer
}

func (u *UserSrv) Login(_ context.Context, request *service.LoginRequest) (*service.LoginResponse, error) {
	username, password := request.Username, request.Password

	user, err := m.localer.SelectUserByUsername(username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &service.LoginResponse{
				Code: http.StatusOK,
				Msg:  "用户名不存在",
			}, nil
		}
		return &service.LoginResponse{
			Code: http.StatusInternalServerError,
		}, nil
	}

	if m.cryptoer.ToMd5(password) != user.Password {
		return &service.LoginResponse{
			Code: http.StatusOK,
			Msg:  "用户名或密码错误",
		}, nil
	}

	tokenString, err := auther.GenerateToken(user.Id, user.Username)
	if err != nil {
		return &service.LoginResponse{
			Code: http.StatusInternalServerError,
		}, nil
	}
	err = m.cacher.SetToken(tokenString)
	if err != nil {
		return &service.LoginResponse{
			Code: http.StatusInternalServerError,
		}, nil
	}

	return &service.LoginResponse{
		Code:  http.StatusOK,
		Msg:   "登录成功",
		Token: tokenString,
	}, nil
}

func (u *UserSrv) Register(_ context.Context, request *service.RegisterRequest) (*service.RegisterResponse, error) {
	username, password := request.Username, request.Password

	if len(username) < 6 || len(username) > 16 {
		return &service.RegisterResponse{
			Code: http.StatusOK,
			Msg:  "用户名非法",
		}, nil
	}

	if len(password) < 8 || len(password) > 16 {
		return &service.RegisterResponse{
			Code: http.StatusOK,
			Msg:  "密码非法",
		}, nil
	}

	user := dba.User{
		Username: username,
		Password: m.cryptoer.ToMd5(password),
	}
	err := m.localer.InsertUser(user)
	if err != nil {
		return &service.RegisterResponse{
			Code: http.StatusOK,
			Msg:  "用户名已存在",
		}, nil
	}

	m.logger.Info("create user success")

	return &service.RegisterResponse{
		Code:   http.StatusOK,
		Msg:    "注册成功",
		UserId: user.Id,
	}, nil
}

func (u *UserSrv) Logout(_ context.Context, request *service.LogoutRequest) (*service.LogoutResponse, error) {
	tokenString := request.Token
	_, err := auther.ParseToken(tokenString)
	if err != nil || !m.cacher.IsTokenExist(tokenString) {
		return &service.LogoutResponse{
			Code: http.StatusOK,
			Msg:  "token已经过期",
		}, nil
	}

	err = m.cacher.DelToken(tokenString)
	if err != nil {
		return &service.LogoutResponse{
			Code: http.StatusInternalServerError,
		}, nil
	}

	return &service.LogoutResponse{
		Code: http.StatusOK,
		Msg:  "登出成功",
	}, nil
}
