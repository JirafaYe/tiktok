package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/JirafaYe/example/hello/internal/service"
	"github.com/JirafaYe/example/hello/pkg/token"
)

type HelloSrv struct {
	service.UnimplementedHelloServer
}

func (h *HelloSrv) SayHello(_ context.Context, request *service.HelloRequest) (*service.HelloResponse, error) {
	tokenString := request.Token
	claims, err := token.ParseToken(tokenString)
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
