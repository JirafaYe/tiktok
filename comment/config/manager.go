package config

import (
	"encoding/json"
	"github.com/JirafaYe/comment/center"
	"log"
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
