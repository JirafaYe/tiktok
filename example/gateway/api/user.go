package api

import (
	"context"
	"net/http"

	"github.com/JirafaYe/example/gateway/center"
	"github.com/JirafaYe/example/gateway/service"
	"github.com/gin-gonic/gin"
)

func (m *Manager) RouteUser() {
	p := m.handler.Group("/user")
	p.POST("/register", m.register)
	p.POST("/login", m.login)
	p.POST("/logout", m.logout)
}

type loginMsg struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type registerMsg struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type logoutMsg struct {
	Token string `json:"token"`
}

func (m *Manager) register(ctx *gin.Context) {
	var msg registerMsg
	if err := ctx.ShouldBindJSON(&msg); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	conn, err := center.Resolver("user")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close()
	c := service.NewUserClient(conn)

	response, _ := c.Register(context.Background(), &service.RegisterRequest{
		Username: msg.Username,
		Password: msg.Password,
	})

	ctx.JSON(int(response.Code), gin.H{"msg": response.Msg, "data": map[string]any{
		"user_id": response.UserId,
	}})
}

func (m *Manager) login(ctx *gin.Context) {
	var msg loginMsg
	if err := ctx.ShouldBindJSON(&msg); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	conn, err := center.Resolver("user")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close()
	c := service.NewUserClient(conn)

	response, _ := c.Login(context.Background(), &service.LoginRequest{
		Username: msg.Username,
		Password: msg.Password,
	})

	ctx.JSON(int(response.Code), gin.H{"msg": response.Msg, "data": map[string]any{
		"token": response.Token,
	}})
}

func (m *Manager) logout(ctx *gin.Context) {
	var msg logoutMsg
	if err := ctx.ShouldBindJSON(&msg); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	conn, err := center.Resolver("user")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close()
	c := service.NewUserClient(conn)

	response, _ := c.Logout(context.Background(), &service.LogoutRequest{
		Token: msg.Token,
	})

	ctx.JSON(int(response.Code), gin.H{"msg": response.Msg})
}
