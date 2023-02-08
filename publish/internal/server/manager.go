package server

import (
	"github.com/JirafaYe/publish/internal/store/cache"
	"github.com/JirafaYe/publish/internal/store/local"
	"github.com/JirafaYe/publish/internal/store/obs"
	"log"
)

var m *Manager

type Manager struct {
	localer 	 *local.Manager
	cacher  	 *cache.Manager
	objectStorer *obs.Manager
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

	objectStorer, err := obs.New()
	if err != nil {
		log.Fatal(err)
	}

	m = &Manager{
		localer:  	  localer,
		cacher:  	  cacher,
		objectStorer: objectStorer,
	}
}
