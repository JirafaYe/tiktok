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

type PublishSrv struct {
	service.UnimplementedPublishServer
}

func (c *PublishSrv) Action (ctx context.Context, request *service.DouyinPublishActionRequest) (*service.DouyinPublishActionResponse, error){
	var response service.DouyinPublishActionResponse
	
	return &response, nil
}

func (c *PublishSrv) List (ctx context.Context, request *service.DouyinPublishListRequest) (*service.DouyinPublishListResponse, error){
	var response service.DouyinPublishListResponse

	return &response, nil
}

// type DouyinPublishActionRequest struct {
// 	state         protoimpl.MessageState
// 	sizeCache     protoimpl.SizeCache
// 	unknownFields protoimpl.UnknownFields

// 	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
// 	Data  []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
// 	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
// }

// type DouyinPublishActionResponse struct {
// 	state         protoimpl.MessageState
// 	sizeCache     protoimpl.SizeCache
// 	unknownFields protoimpl.UnknownFields

// 	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
// 	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
// }


// package server

// import (
// 	"context"
// 	"net/http"

// 	"github.com/JirafaYe/example/user/internal/service"
// 	"github.com/JirafaYe/example/user/internal/store/local"
// 	"github.com/JirafaYe/example/user/pkg/ecrypto"
// 	"github.com/JirafaYe/example/user/pkg/token"
// 	"gorm.io/gorm"
// )

// // UserSrv grpc中关于user相关的服务具体实现
// type UserSrv struct {
// 	service.UnimplementedUserServer
// }

// func (u *UserSrv) Login(_ context.Context, request *service.LoginRequest) (*service.LoginResponse, error) {
// 	username, password := request.Username, request.Password

// 	user, err := m.localer.SelectUserByUsername(username)
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return &service.LoginResponse{
// 				Code: http.StatusOK,
// 				Msg:  "用户名不存在",
// 			}, nil
// 		}
// 		return &service.LoginResponse{
// 			Code: http.StatusInternalServerError,
// 		}, nil
// 	}

// 	if ecrypto.ToMd5(password) != user.Password {
// 		return &service.LoginResponse{
// 			Code: http.StatusOK,
// 			Msg:  "用户名或密码错误",
// 		}, nil
// 	}

// 	tokenString, err := token.GenerateToken(user.Id, user.Username)
// 	if err != nil {
// 		return &service.LoginResponse{
// 			Code: http.StatusInternalServerError,
// 		}, nil
// 	}
// 	err = m.cacher.SetToken(tokenString)
// 	if err != nil {
// 		return &service.LoginResponse{
// 			Code: http.StatusInternalServerError,
// 		}, nil
// 	}

// 	return &service.LoginResponse{
// 		Code:  http.StatusOK,
// 		Msg:   "登录成功",
// 		Token: tokenString,
// 	}, nil
// }

// func (u *UserSrv) Register(_ context.Context, request *service.RegisterRequest) (*service.RegisterResponse, error) {
// 	username, password := request.Username, request.Password

// 	if len(username) < 6 || len(username) > 16 {
// 		return &service.RegisterResponse{
// 			Code: http.StatusOK,
// 			Msg:  "用户名非法",
// 		}, nil
// 	}

// 	if len(password) < 8 || len(password) > 16 {
// 		return &service.RegisterResponse{
// 			Code: http.StatusOK,
// 			Msg:  "密码非法",
// 		}, nil
// 	}

// 	user := local.User{
// 		Username: username,
// 		Password: ecrypto.ToMd5(password),
// 	}
// 	err := m.localer.InsertUser(user)
// 	if err != nil {
// 		return &service.RegisterResponse{
// 			Code: http.StatusOK,
// 			Msg:  "用户名已存在",
// 		}, nil
// 	}

// 	return &service.RegisterResponse{
// 		Code:   http.StatusOK,
// 		Msg:    "注册成功",
// 		UserId: user.Id,
// 	}, nil
// }

// func (u *UserSrv) Logout(_ context.Context, request *service.LogoutRequest) (*service.LogoutResponse, error) {
// 	tokenString := request.Token
// 	_, err := token.ParseToken(tokenString)
// 	if err != nil || !m.cacher.IsTokenExist(tokenString) {
// 		return &service.LogoutResponse{
// 			Code: http.StatusOK,
// 			Msg:  "token已经过期",
// 		}, nil
// 	}

// 	err = m.cacher.DelToken(tokenString)
// 	if err != nil {
// 		return &service.LogoutResponse{
// 			Code: http.StatusInternalServerError,
// 		}, nil
// 	}

// 	return &service.LogoutResponse{
// 		Code: http.StatusOK,
// 		Msg:  "登出成功",
// 	}, nil
// }
