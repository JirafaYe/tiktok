package local

import (
	"github.com/JirafaYe/publish/config"
	"log"
)

type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (c *Config) Key() string {
	return "tiktok/local"
}

var C Config

func init() {
	err := config.ReadConfig(&C)
	if err != nil {
		log.Fatalf("failed to load config %v, errno: %v", C.Key(), err)
	}
}
