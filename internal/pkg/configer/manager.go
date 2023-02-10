package configer

import (
	"encoding/json"
	"log"

	"github.com/JirafaYe/tiktok/internal/pkg/center"
)

type Manager struct {
}

func New() *Manager {
	return &Manager{}
}

func (m *Manager) ReadConfig(c Config) error {
	log.Println("[CONFIG] READING", c.Key())
	config, err := center.GetValue(c.Key())
	if err != nil {
		return err
	}

	err = json.Unmarshal(config, &c)
	if err != nil {
		return err
	}
	return nil
}
