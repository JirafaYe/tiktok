package api

import (
	"context"
	"fmt"
	"github.com/JirafaYe/gateway/center"
	"github.com/JirafaYe/gateway/service"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
)

func (m *Manager) RouteUser() {
	p := m.handler.Group("/douyin")
	p.POST("/user/register/", m.register)
	p.POST("/user/login/", m.login)
	p.GET("/user/", m.getUserMsg)
}

type loginMsg struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type registerMsg struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type userMsg struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

func (m *Manager) register(c *gin.Context) {
	var msg registerMsg
	// 获取请求中的json数据
	username := c.Query("username")
	password := c.Query("password")
	fmt.Println(username)
	fmt.Println(password)
	msg.Username = username
	msg.Password = password
	conn, err := center.Resolver("user")
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)
	ctx := service.NewUserProtoClient(conn)
	response, err := ctx.Register(context.Background(), &service.RegisterRequest{
		Username: msg.Username,
		Password: msg.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(int(http.StatusOK), gin.H{
		"status_code": response.StatusCode,
		"status_msg":  response.StatusMsg,
		"user_id":     response.UserId,
		"token":       response.Token,
	})

}

func (m *Manager) login(c *gin.Context) {
	var msg loginMsg
	// 获取请求中的json数据
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	conn, err := center.Resolver("user")
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)
	ctx := service.NewUserProtoClient(conn)

	response, err := ctx.Login(context.Background(), &service.LoginRequest{
		Username: msg.Username,
		Password: msg.Password,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": response.StatusCode,
		"status_msg":  response.StatusMsg,
		"user_id":     response.UserId,
		"token":       response.Token,
	})
}

func (m *Manager) isLogin(c *gin.Context) {
	token := c.Query("token")
	conn, err := center.Resolver("user")
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)
	ctx := service.NewUserProtoClient(conn)
	login, err := ctx.IsLogin(context.Background(), &service.IsLoginRequest{
		Token: token,
	})
	if err != nil {
		panic(nil)
	}
	if login.Code == 0 {
		c.Next()
	} else {
		c.JSON(int(login.Code), gin.H{"msg": "状态未登录，请先登录账号！", "data": nil})
	}

}

func (m *Manager) getUserMsg(c *gin.Context) {
	token := c.Query("token")
	userId := c.Query("user_id")
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		c.JSON(1, gin.H{"msg": "上传的id错误", "data": nil})
	}
	conn, err := center.Resolver("user")
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)
	ctx := service.NewUserProtoClient(conn)
	msg, err := ctx.GetUserMsg(context.Background(), &service.UserRequest{
		UserId: id,
		Token:  token,
	})
	if err != nil {
		c.JSON(1, gin.H{"msg": err, "data": nil})
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": msg.StatusCode,
		"status_msg":  msg.StatusMsg,
		"user":        msg.User,
	})
}
