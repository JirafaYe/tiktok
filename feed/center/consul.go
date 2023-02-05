package center

import (
	"log"

	consul "github.com/hashicorp/consul/api"
)

var (
	addr = "47.108.66.104:8500"
	//addr   = "192.168.3.126:8500"
	client *consul.Client
)

func init() {
	var err error
	config := consul.DefaultConfig()
	config.Address = addr
	client, err = consul.NewClient(config)
	if err != nil {
		log.Fatalf("failed to init consul: %v", err)
	}
}

func GetValue(key string) ([]byte, error) {
	res, _, err := client.KV().Get(key, nil)
	if err != nil {
		return nil, err
	}
	return res.Value, nil
}

func Register(reg consul.AgentServiceRegistration) error {
	agent := client.Agent()

	return agent.ServiceRegister(&reg)
}
