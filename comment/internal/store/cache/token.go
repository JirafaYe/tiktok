package cache

import (
	"context"
)

const (
	UserToken = "USER_TOKEN"
)

func (m *Manager) IsUserTokenExist(username string) bool {
	exists, err := m.handler.HExists(context.Background(), UserToken, username).Result()
	return err == nil && exists
}
