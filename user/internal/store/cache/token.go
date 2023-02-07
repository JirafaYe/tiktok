package cache

import (
	"context"
)

const (
	UserToken = "USER_TOKEN"
)

func (m *Manager) SetUserToken(token string, username string) error {
	err := m.handler.HSet(context.Background(), UserToken, username, token).Err()
	return err
}

func (m *Manager) DelUserToken(username string) error {
	return m.handler.HDel(context.Background(), UserToken, username).Err()

}

func (m *Manager) IsUserTokenExist(username string) bool {
	exists, err := m.handler.HExists(context.Background(), UserToken, username).Result()
	return err == nil && exists
}
