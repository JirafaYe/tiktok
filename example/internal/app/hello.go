package app

import (
	"context"
	"fmt"
	"github.com/JirafaYe/tiktok/example/internal/pkg/auther"
	"net/http"

	"github.com/JirafaYe/tiktok/example/internal/app/service"
)

type HelloSrv struct {
	service.UnimplementedHelloServer
}

func (h *HelloSrv) SayHello(_ context.Context, request *service.HelloRequest) (*service.HelloResponse, error) {
	tokenString := request.Token
	claims, err := auther.ParseToken(tokenString)
	if err != nil {
		return &service.HelloResponse{
			Code: http.StatusInternalServerError,
		}, nil
	}
	return &service.HelloResponse{
		Code: http.StatusOK,
		Msg:  "成功",
		Data: fmt.Sprintf("Hello %v.", claims.Username),
	}, nil
}
