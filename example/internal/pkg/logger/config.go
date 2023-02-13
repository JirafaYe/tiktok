package logger

import (
	config "github.com/JirafaYe/tiktok/example/internal/pkg/configer"
	"log"
)

type Config struct {
	Output      string `json:"output"`
	ProjectName string `json:"project_name"`
	Level       string `json:"level"`
}

func (c *Config) Key() string {
	return "tiktok/example/logger"
}

var C Config

func init() {
	err := config.ReadConfig(&C)
	if err != nil {
		log.Fatalf("failed to load config %v, errno: %v", C.Key(), err)
	}
}
