package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/JirafaYe/feed/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
	"time"
)

var (
	addr string
	name string
)

func main() {
	// Set up a connection to the server.
	addr = "localhost:8888"
	name = "feed"
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := service.NewFeedClient(conn)

	//// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var lastTime int64 = 0
	var token = ""
	r, err := c.FeedVideo(ctx, &service.TiktokFeedRequest{
		LastTime: &lastTime,
		Token:    &token,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.VideoList)
}

func startTcp() {
	ls, err := net.Listen("tcp", ":10111")
	if err != nil {
		fmt.Printf("start tcp listener error: %v\n", err.Error())
		return
	}
	for {
		conn, err := ls.Accept()
		if err != nil {
			fmt.Printf("connect error: %v\n", err.Error())
		} else {
			fmt.Println(conn.RemoteAddr())
		}
		go func(conn net.Conn) {
			_, err := bufio.NewWriter(conn).WriteString("hello consul")
			if err != nil {
				fmt.Printf("write conn error: %v\n", err)
			}
		}(conn)
	}
}

func startHttp() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("consul get uri: %s\n", r.RequestURI)
		w.Write([]byte("hello consul"))
	})
	if err := http.ListenAndServe(":10111", nil); err != nil {
		fmt.Printf("start http server error: %v\n", err)
	}
}
