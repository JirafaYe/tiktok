package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/JirafaYe/tiktok/example/internal/app"
	"github.com/JirafaYe/tiktok/example/internal/app/service"
	"github.com/JirafaYe/tiktok/example/internal/pkg/center"
	consul "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

var (
	addr = flag.String("addr", "127.0.0.1", "The server address")
	port = flag.Int("port", 50053, "The server port")
)

func main() {
	flag.Parse()

	// 服务注册
	err := center.Register(consul.AgentServiceRegistration{
		ID:      "example-hello-1", // 服务节点的名称
		Name:    "example-hello",   // 服务名称
		Port:    *port,             // 服务端口
		Address: *addr,             // 服务 IP
		//Check: &consul.AgentServiceCheck{ // 健康检查
		//	Interval:                       "5s", // 健康检查间隔
		//	GRPC:                           fmt.Sprintf("%v:%v/%v", *addr, *port, "health"),
		//	DeregisterCriticalServiceAfter: "10s", // 注销时间，相当于过期时间
		//},
	})
	if err != nil {
		log.Fatalf("failed to register service: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	service.RegisterHelloServer(s, &app.HelloSrv{})
	// 健康检查
	//grpc_health_v1.RegisterHealthServer(s, &launcher.HealthImpl{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
