package api

import (
	"github.com/gin-gonic/gin"
)

func (m *Manager) RouteFeed() {
	m.handler.GET("/douyin/feed", m.feed)
	m.handler.GET("/douyin/feed/content/", m.feedContent)
}

func (m *Manager) feed(c *gin.Context) {
}

func (m *Manager) feedContent(c *gin.Context) {

}
