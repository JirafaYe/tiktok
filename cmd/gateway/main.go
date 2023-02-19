package main

import (
	"github.com/JirafaYe/tiktok/internal/web"
	"log"
)

func main() {
	app := web.New()
	if err := app.Run(); err != nil {
		log.Fatalf("gateway failed to server")
	}
}
