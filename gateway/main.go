package main

import (
	"github.com/JirafaYe/gateway/api"
	"log"
)

func init() {
	log.SetFlags(log.Llongfile)
}

func main() {
	app := api.New()
	err := app.Run()
	if err != nil {
		log.Fatalf("failed to run server, %v", err)
	}
}
