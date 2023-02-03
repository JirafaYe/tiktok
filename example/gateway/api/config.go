package api

import (
	"github.com/JirafaYe/example/gateway/config"
	"log"
)

type Config struct {
	IP   string `json:"ip"`
	Port string `json:"port"`
}

func (c *Config) Key() string {
	return "tiktok/example/api"
}

var C Config

func init() {
	err := config.ReadConfig(&C)
	if err != nil {
		log.Fatalf("failed to load config %v, errno: %v", C.Key(), err)
	}
}
