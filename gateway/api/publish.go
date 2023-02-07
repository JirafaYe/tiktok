package api

import (
	"github.com/gin-gonic/gin"
)

// publish的路由

func (m *Manager) RoutePublish() {
	m.handler.POST("/douyin/publish/action", m.publishAction)
	m.handler.GET("/douyin/publish/list", m.publishList)
}

func (m *Manager) publishAction(ctx *gin.Context) {

}

func (m *Manager) publishList(ctx *gin.Context) {

}
