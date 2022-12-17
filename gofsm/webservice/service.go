package webservice

import (
	"fmt"
	"log"
	"net"
)

func Greet() {
	fmt.Println("Hello")
}

func Listen() {
	_, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
