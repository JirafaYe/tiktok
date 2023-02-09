package app

import (
	"log"

	cache "github.com/JirafaYe/tiktok/internal/pkg/cacher"
	crypto "github.com/JirafaYe/tiktok/internal/pkg/cryptoer"
	local "github.com/JirafaYe/tiktok/internal/pkg/dba"
	"github.com/JirafaYe/tiktok/internal/pkg/logger"
	"github.com/JirafaYe/tiktok/internal/pkg/obs"
)

var m *Manager

type Manager struct {
	cacher   *cache.Manager
	localer  *local.Manager
	cryptoer *crypto.Manager
	obser    *obs.Manager
	logger   *logger.Manager
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
	obs, err := obs.New()
	if err != nil {
		log.Fatal(err)
	}
	cryptoer := crypto.New()
	logger := logger.New()

	m = &Manager{
		localer:  localer,
		cacher:   cacher,
		obser:    obs,
		cryptoer: cryptoer,
		logger:   logger,
	}
}
