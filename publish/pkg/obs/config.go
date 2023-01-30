package obs

import (
	"github.com/JirafaYe/publish/config"
	"log"
)

type Config struct {
	Address   string `json:"address"`
	SecretID  string `json:"secret_id"`
	SecretKey string `json:"secret_key"`
}

func (c *Config) Key() string {
	return "tiktok/obs"
}

var C Config

func init() {
	err := config.ReadConfig(&C)
	if err != nil {
		log.Fatalf("failed to load config %v, errno: %v", C.Key(), err)
	}
}
