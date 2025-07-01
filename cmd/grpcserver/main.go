package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/levikl/go-specs-greet/adapters/grpcserver"
	"github.com/levikl/go-specs-greet/domain/interactions"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	grpcserver.RegisterGreeterServer(s, &GreetServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

type GreetServer struct {
	grpcserver.UnimplementedGreeterServer
}

func (g GreetServer) Greet(
	ctx context.Context,
	request *grpcserver.GreetRequest,
) (*grpcserver.GreetReply, error) {
	return &grpcserver.GreetReply{Message: interactions.Greet(request.Name)}, nil
}
