package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCommentAction1Route(t *testing.T) {
	engine := route()

	url := "http://localhost:8088/douyin/comment/action/?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjI2ODMzNTcyNzM4MzE1NDY4OCwidXNlcm5hbWUiOiJxcXFxcXEiLCJleHAiOjE2Nzc1NzgxODEsImlzcyI6InRpa3Rvay51c2VyIn0.444I8M8xaZdAr-PH8nriyRScUWyukmjGTg11Xfy1EOE&video_id=1&action_type=1&comment_text=%E4%BD%A0%E5%A5%BD"
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", url, nil)
	engine.ServeHTTP(recorder, request)

	var resp CommentOperationResponse
	json.Unmarshal(recorder.Body.Bytes(), &resp)

	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, 0, int(resp.StatusCode))

	log.Println(resp)
}

func TestCommentAction2Route(t *testing.T) {
	engine := route()

	url := "http://localhost:8088/douyin/comment/action/?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjI2ODMzNTcyNzM4MzE1NDY4OCwidXNlcm5hbWUiOiJxcXFxcXEiLCJleHAiOjE2Nzc1NzgxODEsImlzcyI6InRpa3Rvay51c2VyIn0.444I8M8xaZdAr-PH8nriyRScUWyukmjGTg11Xfy1EOE&video_id=1&action_type=2&comment_text=%E4%BD%A0%E5%A5%BD&comment_id=140"
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", url, nil)
	engine.ServeHTTP(recorder, request)

	var resp CommentOperationResponse
	json.Unmarshal(recorder.Body.Bytes(), &resp)

	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, 0, int(resp.StatusCode))

	log.Println(resp)
}

func TestCommentActionIllegalRoute(t *testing.T) {
	engine := route()

	url := "http://localhost:8088/douyin/comment/action/?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjI2ODMzNTcyNzM4MzE1NDY4OCwidXNlcm5hbWUiOiJxcXFxcXEiLCJleHAiOjE2Nzc1NzgxODEsImlzcyI6InRpa3Rvay51c2VyIn0.444I8M8xaZdAr-PH8nriyRScUWyukmjGTg11Xfy1EOE&video_id=1&action_type=3"
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", url, nil)
	engine.ServeHTTP(recorder, request)

	var resp CommentOperationResponse
	json.Unmarshal(recorder.Body.Bytes(), &resp)

	assert.Equal(t, 500, recorder.Code)
	assert.Equal(t, 500, int(resp.StatusCode))

	log.Println(resp)
}

func TestCommentListRoute(t *testing.T) {
	engine := route()

	url := "http://localhost:8088/douyin/comment/list/?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjI2ODMzNTcyNzM4MzE1NDY4OCwidXNlcm5hbWUiOiJxcXFxcXEiLCJleHAiOjE2Nzc1NzgxODEsImlzcyI6InRpa3Rvay51c2VyIn0.444I8M8xaZdAr-PH8nriyRScUWyukmjGTg11Xfy1EOE&video_id=1"
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", url, nil)
	engine.ServeHTTP(recorder, request)

	var resp ListCommentResponse
	json.Unmarshal(recorder.Body.Bytes(), &resp)

	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, 0, int(resp.StatusCode))

	log.Println(resp)
}

func route() *gin.Engine {
	app := New()
	app.loadRoute()
	return app.handler
}
