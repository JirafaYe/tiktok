package api

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

type UserResponseMsg struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	Token      string `json:"token"`
	UserId     int64  `json:"user_id"`
}

type DouYinUserMsg struct {
	StatusCode int     `json:"status_code"`
	StatusMsg  string  `json:"status_msg"`
	User       UserMsg `json:"user"`
}
type UserMsg struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func TestUserRegisterRoute(t *testing.T) {
	engine := route()
	recorder, _, _ := register(engine)

	resp, done := RegisterUser(recorder)
	if done {
		return
	}

	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, 0, resp.StatusCode)

	log.Println(resp)
}

func RegisterUser(recorder *httptest.ResponseRecorder) (UserResponseMsg, bool) {
	var resp UserResponseMsg
	err := json.Unmarshal(recorder.Body.Bytes(), &resp)
	if err != nil {
		return UserResponseMsg{}, true
	}
	return resp, false
}

func register(engine *gin.Engine) (recorder *httptest.ResponseRecorder, username, password string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		num := r.Intn(10)
		username += strconv.Itoa(num)
		num = r.Intn(10)
		password += strconv.Itoa(num)
	}
	msg := map[string]string{
		"username": username,
		"password": password,
	}
	url := "http://localhost:8088/douyin/user/register/?username=" + msg["username"] + "&password=" + msg["password"]
	recorder = httptest.NewRecorder()
	request, _ := http.NewRequest("POST", url, nil)
	engine.ServeHTTP(recorder, request)
	return recorder, username, password
}

func TestLoginRoute(t *testing.T) {

	engine := route()
	_, username, password := register(engine)
	recorder, resp, done := login(username, password, engine)
	if done {
		return
	}

	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, 0, int(resp.StatusCode))

	log.Println(resp)
}

func login(username string, password string, engine *gin.Engine) (*httptest.ResponseRecorder, UserResponseMsg, bool) {
	url := "http://localhost:8088/douyin/user/login/"
	user := map[string]string{
		"username": username,
		"password": password,
	}
	recorder := httptest.NewRecorder()
	userJson, _ := json.Marshal(&user)
	userBytes := bytes.NewReader(userJson)
	request, _ := http.NewRequest("POST", url, userBytes)
	engine.ServeHTTP(recorder, request)

	var resp UserResponseMsg
	err := json.Unmarshal(recorder.Body.Bytes(), &resp)
	if err != nil {
		return nil, UserResponseMsg{}, true
	}
	return recorder, resp, false
}

func TestUserRoute(t *testing.T) {
	engine := route()
	_, username, password := register(engine)
	_, msg, _ := login(username, password, engine)
	url := "http://localhost:8088/douyin/user/?user_id=" + strconv.FormatInt(msg.UserId, 10) + "&token=" + msg.Token
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", url, nil)
	engine.ServeHTTP(recorder, request)

	var resp DouYinUserMsg
	err := json.Unmarshal(recorder.Body.Bytes(), &resp)
	if err != nil {
		return
	}

	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, 0, int(resp.StatusCode))

	log.Println(resp)
}

func route() *gin.Engine {
	app := New()
	err := app.loadRoute()
	if err != nil {
		return nil
	}
	return app.handler
}

func BenchmarkRegister(b *testing.B) {
	for i := 0; i < b.N; i++ {
		engine := route()
		recorder, _, _ := register(engine)
		_, done := RegisterUser(recorder)
		if done {
			return
		}
	}
}

func BenchmarkLogin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		engine := route()
		_, username, password := register(engine)
		_, _, done := login(username, password, engine)
		if done {
			return
		}
	}
}

func BenchmarkUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		engine := route()
		_, username, password := register(engine)
		_, msg, _ := login(username, password, engine)
		url := "http://localhost:8088/douyin/user/?user_id=" + strconv.FormatInt(msg.UserId, 10) + "&token=" + msg.Token
		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", url, nil)
		engine.ServeHTTP(recorder, request)

		var resp DouYinUserMsg
		err := json.Unmarshal(recorder.Body.Bytes(), &resp)
		if err != nil {
			return
		}
	}
}
