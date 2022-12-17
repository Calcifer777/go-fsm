package main

import (
	"fmt"
	ws "gofsm/webservice"
	pb "gofsm/webservice/protos"

	"log"
	"net"

	"flag"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 9000, "The server port")
)

func main() {
	flag.Parse()
	// Setup listener
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Register server
	s := grpc.NewServer()
	pb.RegisterTutorialServer(s, &ws.Server{})
	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
