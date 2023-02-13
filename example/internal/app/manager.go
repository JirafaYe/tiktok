package app

import (
	"log"

	"github.com/JirafaYe/tiktok/example/internal/pkg/cacher"
	"github.com/JirafaYe/tiktok/example/internal/pkg/cryptoer"
	"github.com/JirafaYe/tiktok/example/internal/pkg/dba"
	"github.com/JirafaYe/tiktok/example/internal/pkg/logger"
	"github.com/JirafaYe/tiktok/example/internal/pkg/obs"
)

var m *Manager

type Manager struct {
	cacher   *cacher.Manager
	localer  *dba.Manager
	cryptoer *cryptoer.Manager
	obser    *obs.Manager
	logger   *logger.Manager
}

func init() {
	var err error
	localer, err := dba.New()
	if err != nil {
		log.Fatal(err)
	}
	cacher, err := cacher.New()
	if err != nil {
		log.Fatal(err)
	}
	obser, err := obs.New()
	if err != nil {
		log.Fatal(err)
	}
	cryptoer := cryptoer.New()
	logger := logger.New()

	m = &Manager{
		localer:  localer,
		cacher:   cacher,
		obser:    obser,
		cryptoer: cryptoer,
		logger:   logger,
	}
}
