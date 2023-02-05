package main

import (
	"github.com/JirafaYe/feed/internal/server"
	"github.com/JirafaYe/feed/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	service.RegisterFeedServer(s, &server.FeedServer{})

	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("server.Serve err: %v", err)
	}
}
