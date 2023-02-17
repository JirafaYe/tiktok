package api

import (
	"context"
	"net/http"
	"log"
	"bytes"
	"io"

	"github.com/JirafaYe/gateway/center"
	"github.com/JirafaYe/gateway/service"
	"github.com/gin-gonic/gin"

	"strconv"
)

// type ActionMsg struct {
// 	//TODO_hewen
// 	Token string `json:"token"`
// 	Title string `json:"title"`
// 	Data []byte	 `json:"file"`
// }

// TODO: 这个结构体其实可以不要的
type ActionResponse struct {
	StatusCode int64   `json:"status_code"`
	StatusMsg  *string `json:"status_msg"` 
}

// RouteUser 注册路由且该函数必须以Route前缀
// main.go文件运行时会通过反射来查看有Route前缀的函数来进行路由注册
func (m *Manager) RoutePublish() {
	// group := m.handler.Group("/douyin/publish")
	// group.POST("/action", m.publishAction)
	// group.Get("/list", m.publishList)
	m.handler.POST("/douyin/publish/action", m.publishAction)
	m.handler.GET("/douyin/publish/list", m.publishList)
}

//TODO_hewen 修改并完善publishAction
func (m *Manager) publishAction(ctx *gin.Context) {
	// ctx 会检索URL
    token := ctx.PostForm("token")
    title := ctx.PostForm("title")
	// token := ctx.Query("token")// string
	// title := ctx.Query("title")// string
	// 使用FormFile从请求中读取文件，文件类型
	fileHeader, err := ctx.FormFile("data")// file, type is "multipart.FileHeader"
	if err != nil {
		log.Printf("can't get data file: %v", err)
		return
	}

	file, err := fileHeader.Open()
	if err!= nil {
		log.Printf("can't open data file: %v", err)
        return
	}
	defer file.Close()

	buffer := bytes.NewBuffer(nil)
	if _, err = io.Copy(buffer, file); err!= nil {
		log.Printf("can't read data file: %v", err)
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

	client := service.NewPublishClient(conn)
	rpcResponse, err := client.PubAction(context.Background(), &service.PublishActionRequest{
		Token: token,
		Data:  buffer.Bytes(),
		Title: title,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
	}

	response := ActionResponse{
		StatusCode: int64(rpcResponse.StatusCode),
		StatusMsg: &rpcResponse.StatusMsg,
	}
	ctx.JSON(int(http.StatusOK), gin.H{
		"status_code": response.StatusCode,
		"status_msg":  response.StatusMsg,
	})
}

// type ListResponse struct {
// 	StatusCode int64   `json:"status_code"`// 状态码，0-成功，其他值-失败
// 	StatusMsg  *string `json:"status_msg"` // 返回状态描述
// 	VideoList  []Video `json:"video_list"` // 用户发布的视频列表
// }
type ListResponse struct {
	StatusCode int64   `json:"status_code"`// 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"` // 返回状态描述
	VideoList  []*PubVideo `json:"video_list"` // 用户发布的视频列表
}

// Video
type PubVideo struct {
	Author        PubUser   `json:"author"`        // 视频作者信息
	CommentCount  int64  `json:"comment_count"` // 视频的评论总数
	CoverURL      string `json:"cover_url"`     // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"`// 视频的点赞总数
	ID            int64  `json:"id"`            // 视频唯一标识
	IsFavorite    bool   `json:"is_favorite"`   // true-已点赞，false-未点赞
	PlayURL       string `json:"play_url"`      // 视频播放地址
	Title         string `json:"title"`         // 视频标题
}

// 视频作者信息
//
// User
type PubUser struct {
	FollowCount   int64  `json:"follow_count"`  // 关注总数
	FollowerCount int64  `json:"follower_count"`// 粉丝总数
	ID            int64  `json:"id"`            // 用户id
	IsFollow      bool   `json:"is_follow"`     // true-已关注，false-未关注
	Name          string `json:"name"`          // 用户名称
}

//TODO_hewen 修改并完善publishList
func (m *Manager) publishList(ctx *gin.Context) {
	// ctx 会检索URL
	// TODO: PostForm还是query?
    token := ctx.PostForm("token")
    tmpUserID := ctx.PostForm("user_id")
	userID, err := strconv.ParseInt(tmpUserID, 10, 64)
	if err != nil {
		ctx.JSON(1, gin.H{"msg": "上传的id错误", "data": nil})
	}
	// center.Resolver() 参数为调用的服务名
	// 该函数会进行自动负载均衡并返回一个*grpc.ClientConn
	conn, err := center.Resolver("publish")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close() 

	client := service.NewPublishClient(conn)
	rpcResponse, err := client.PubList(context.Background(), &service.PublishListRequest{
		Token:  token,
		UserId: userID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
	}

	response := ListResponse{
		StatusCode: int64(rpcResponse.StatusCode),
		StatusMsg:  rpcResponse.StatusMsg,
	}
	response.VideoList = make([]*PubVideo, len(rpcResponse.VideoList))
	for i, v := range rpcResponse.VideoList {
		response.VideoList[i] = &PubVideo{
			ID: v.Id,
			Author: PubUser{
				ID:            v.Author.Id,
				Name:          v.Author.Name,
				FollowCount:   v.Author.FollowCount,
				FollowerCount: v.Author.FollowerCount,
				IsFollow:      v.Author.IsFollow,
			},
			PlayURL:       v.PlayUrl,
			CoverURL:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    v.IsFavorite,
			Title:         v.Title,
		}
	}

	ctx.JSON(int(response.StatusCode), response)
}
