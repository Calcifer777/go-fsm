package webservice

import (
	context "context"
	"fmt"
	pb "gofsm/webservice/protos"
)

type Server struct {
	pb.UnimplementedTutorialServer
}

func (s *Server) CreateUser(ctx context.Context, in *pb.UserRequest) (*pb.User, error) {
	return &pb.User{Id: 1, Name: in.Name}, nil
}

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	return &pb.GreetResponse{Message: fmt.Sprintf("Hi %s!", in.Name)}, nil
}
