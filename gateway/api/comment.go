package api

import "github.com/gin-gonic/gin"

func (m *Manager) RouteComment() {
	m.handler.Group("/douyin/comment/")
	m.handler.POST("action/", m.Action)
}

func (m *Manager) Action(c *gin.Context) {

}
