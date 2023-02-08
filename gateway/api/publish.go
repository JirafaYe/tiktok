package api

import (
	"context"
	"net/http"

	"github.com/JirafaYe/gateway/center"
	"github.com/JirafaYe/gateway/service"
	"github.com/gin-gonic/gin"
)

// type ActionMsg struct {
// 	//TODO_hewen
// 	Token string `json:"token"`
// 	Title string `json:"title"`
// 	Data []byte	 `json:"file"`
// }

type ActionResponse struct {
	StatusCode int64   `json:"status_code"`
	StatusMsg  *string `json:"status_msg"` 
}

// RouteUser 注册路由且该函数必须以Route前缀
// main.go文件运行时会通过反射来查看有Route前缀的函数来进行路由注册
func (m *Manager) RoutePublish() {
	group := m.handler.Group("/douyin/publish")
	group.POST("/action", m.publishAction)
	group.Get("/list", m.publishList)
	// m.handler.POST("/douyin/publish/action", m.publishAction)
	// m.handler.GET("/douyin/publish/list", m.publishList)
}

//TODO_hewen 修改并完善publishAction
func (m *Manager) publishAction(ctx *gin.Context) {
	// ctx 会检索URL
	token := ctx.Query("token")// string
	title := ctx.Query("title")// string
	data := ctx.Query("data")// file
	
	// center.Resolver() 参数为调用的服务名
	// 该函数会进行自动负载均衡并返回一个*grpc.ClientConn
	conn, err := center.Resolver("publish")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close() 

	client := service.NewUserClient(conn)
	rpcResponse, err := client.PublishAction(context.Background(), &service.DouyinPublishActionRequest{
		Token: &token,
		Data: data,
		Title: &data,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
	}

	response := ActionResponse{
		StatusCode: int64(rpcResponse.StatusCode),
		StatusMsg: rpcResponse.StatusMsg,
	}
	ctx.JSON(int(http.StatusCode), gin.H{
		"status_code": response.StatusCode,
		"status_msg":  response.StatusMsg,
	})
}


// response JSON
// {
//     "status_code": 0,
//     "status_msg": "string"
// }

//TODO_hewen 修改并完善publishList
func (m *Manager) publishList(ctx *gin.Context) {
	var msg registerMsg
	// 获取请求中的json数据
	if err := ctx.ShouldBindJSON(&msg); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// center.Resolver() 参数为调用的服务名
	// 该函数会进行自动负载均衡并返回一个*grpc.ClientConn
	conn, err := center.Resolver("publish")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close()
	c := service.NewUserClient(conn)

	response, _ := c.PublishList(context.Background(), &service.RegisterRequest{
		Username: msg.Username,
		Password: msg.Password,
	})

	ctx.JSON(int(response.Code), gin.H{"msg": response.Msg, "data": map[string]any{
		"user_id": response.UserId,
	}})
}
