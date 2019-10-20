package main

import (
	"log"
	"net"

	"gogithub/config"
	"gogithub/github"
	pb "gogithub/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	addr := config.GrpcServerAddress()
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("gRPC server listening at " + addr)
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterGithubServiceServer(s, &github.GrpcServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
