package server

import (
	"context"
	"net/http"

	"github.com/JirafaYe/example/user/internal/service"
	"github.com/JirafaYe/example/user/internal/store/local"
	"github.com/JirafaYe/example/user/pkg/ecrypto"
	"github.com/JirafaYe/example/user/pkg/token"
	"gorm.io/gorm"
)

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

	if ecrypto.ToMd5(password) != user.Password {
		return &service.LoginResponse{
			Code: http.StatusOK,
			Msg:  "用户名或密码错误",
		}, nil
	}

	tokenString, err := token.GenerateToken(user.Id, user.Username)
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

	user := local.User{
		Username: username,
		Password: ecrypto.ToMd5(password),
	}
	err := m.localer.InsertUser(user)
	if err != nil {
		return &service.RegisterResponse{
			Code: http.StatusOK,
			Msg:  "用户名已存在",
		}, nil
	}

	return &service.RegisterResponse{
		Code:   http.StatusOK,
		Msg:    "注册成功",
		UserId: user.Id,
	}, nil
}

func (u *UserSrv) Logout(_ context.Context, request *service.LogoutRequest) (*service.LogoutResponse, error) {
	tokenString := request.Token
	_, err := token.ParseToken(tokenString)
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
