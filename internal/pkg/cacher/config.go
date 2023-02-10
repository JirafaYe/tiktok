package cacher

import (
	"log"

	config "github.com/JirafaYe/tiktok/internal/pkg/configer"
)

type Config struct {
	Addr     string `json:"addr"`
	Port     string `json:"port"`
	Password string `json:"password"`
}

func (c *Config) Key() string {
	return "tiktok/cache"
}

var C Config

func init() {
	err := config.ReadConfig(&C)
	if err != nil {
		log.Fatalf("failed to load config %v, errno: %v", C.Key(), err)
	}
}
