package api

import (
	"context"
	"net/http"

	"github.com/JirafaYe/example/gateway/center"
	"github.com/JirafaYe/example/gateway/service"
	"github.com/gin-gonic/gin"
)

func (m *Manager) RouteHello() {
	m.handler.POST("/hello", m.hello)
}

type helloMsg struct {
	Token string `json:"token"`
}

func (m *Manager) hello(ctx *gin.Context) {
	var msg helloMsg
	if err := ctx.ShouldBindJSON(&msg); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	conn, err := center.Resolver("hello")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close()

	c := service.NewHelloClient(conn)
	response, _ := c.SayHello(context.Background(), &service.HelloRequest{
		Token: msg.Token,
	})

	ctx.JSON(int(response.Code), gin.H{"msg": response.Msg, "data": map[string]any{
		"data": response.Data,
	}})
}
