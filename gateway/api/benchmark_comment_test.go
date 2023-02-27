package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

var engine *gin.Engine

func TestMain(m *testing.M) {
	engine = route()
	m.Run()
}

var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjI2NzI3MDEwMDYxNDY0NzgwOCwidXNlcm5hbWUiOiIyNTgyMTAxNDU4QHFxLmNvbSIsImV4cCI6MTY3ODA4NTcyOCwiaXNzIjoidGlrdG9rLnVzZXIifQ.3mFOUFJBJtsnnJsg_uvBJpSixMf8mZtdFkeT3RvIebI"

func BenchmarkCommentAction1Route(b *testing.B) {
	for i := 0; i < b.N; i++ {
		url := "http://localhost:8088/douyin/comment/action/?token=" + token + "&video_id=1&action_type=1&comment_text=%E4%BD%A0%E5%A5%BD"
		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", url, nil)
		engine.ServeHTTP(recorder, request)
	}
}

func BenchmarkCommentAction2Route(b *testing.B) {
	for i := 0; i < b.N; i++ {
		url := "http://localhost:8088/douyin/comment/action/?token=" + token + "&video_id=1&action_type=2&comment_text=%E4%BD%A0%E5%A5%BD&comment_id="
		url = url + strconv.Itoa(i+146)
		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", url, nil)
		engine.ServeHTTP(recorder, request)
	}
}

func BenchmarkCommentIllegalActionRoute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		url := "http://localhost:8088/douyin/comment/action/?token=" + token + "&video_id=1&action_type=3"
		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", url, nil)
		engine.ServeHTTP(recorder, request)
	}
}

func BenchmarkCommentList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		url := "http://localhost:8088/douyin/comment/list/?token=" + token + "&video_id=1"
		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", url, nil)
		engine.ServeHTTP(recorder, request)
	}
}
