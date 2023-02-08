package center

import (
	"fmt"
	"log"

	consul "github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr   = "47.108.66.104:8500"
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

func Resolver(name string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(
		// name 拉取的服务名 wait=14s 等待时间 tag=manual 筛选条件
		fmt.Sprintf("consul://%v/%v?wait=14s&tag=manual", addr, name),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
