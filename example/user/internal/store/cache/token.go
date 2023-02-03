package cache

import (
	"context"
	"time"
)

func (m *Manager) SetToken(token string) error {
	return m.handler.Set(context.Background(), token, "", time.Hour*24*7).Err()
}

func (m *Manager) DelToken(token string) error {
	return m.handler.Del(context.Background(), token).Err()
}

func (m *Manager) IsTokenExist(token string) bool {
	num, err := m.handler.Exists(context.Background(), token).Result()
	return err == nil && num != 0
}
