package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type Manager struct {
	handler *redis.Client
}

func New() (*Manager, error) {
	m := &Manager{
		handler: redis.NewClient(
			&redis.Options{
				Addr:     C.Addr + ":" + C.Port,
				Password: C.Password,
			},
		),
	}
	return m, m.handler.Ping(context.Background()).Err()
}
