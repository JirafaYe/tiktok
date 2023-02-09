package util

import (
	"io/ioutil"
	"log"
	"fmt"
)

func NewString(s string) *string {
	return &s
}

func GetVideo(filePath string) ([]byte, error) {
	file, err := ioutil.ReadFile(filePath)
	if err!= nil {
		log.Printf("can't read video file by filePath: %v", err)
	}
	fmt.Printf("stream type: %T\n", file)
	return file, nil
}
