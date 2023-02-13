package obs

import (
	"log"

	config "github.com/JirafaYe/tiktok/example/internal/pkg/configer"
)

type Config struct {
	Address   string `json:"address"`
	SecretID  string `json:"secret_id"`
	SecretKey string `json:"secret_key"`
}

func (c *Config) Key() string {
	return "tiktok/example/obs"
}

var C Config

func init() {
	err := config.ReadConfig(&C)
	if err != nil {
		log.Fatalf("failed to load config %v, errno: %v", C.Key(), err)
	}
}
