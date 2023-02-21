package pkg

import (
    "testing"
)

func TestGenerateToken(t *testing.T) {
	userName := "hewen"
	userId := 4
	token, err := GenerateToken(int64(userId), userName)
	if err!= nil {
        t.Error(err)
    }
	t.Log(token)
}