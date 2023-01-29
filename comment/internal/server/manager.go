package server

import (
	"github.com/JirafaYe/comment/internal/store/cache"
	"github.com/JirafaYe/comment/internal/store/local"
	"log"
)

var m *Manager

type Manager struct {
	localer *local.Manager
	cacher  *cache.Manager
}

func init() {
	var err error
	localer, err := local.New()
	if err != nil {
		log.Fatal(err)
	}
	cacher, err := cache.New()
	if err != nil {
		log.Fatal(err)
	}

	m = &Manager{
		localer: localer,
		cacher:  cacher,
	}
}
