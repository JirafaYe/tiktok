package server

import (
	"github.com/JirafaYe/user/pkg"
	"testing"
)

func TestA(t *testing.T) {
	token, err := pkg.GenerateToken(21521512234, "1125887000@qq.com")
	if err != nil {
		panic(err)
	}
	err = m.cacher.SetUserToken(token, "1125887000@qq.com")
	if err != nil {
		panic(err)
	}
}

func TestB(t *testing.T) {
	err := m.cacher.DelUserToken("1125887001@qq.com")
	if err != nil {
		panic(err)
	}
}
