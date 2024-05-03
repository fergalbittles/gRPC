package main

import (
	"grpc/user"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	log.Println("| Starting server on port 4000")

	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatalf("Failed to listen on port 4000: %v", err)
	}

	s := user.Server{}

	grpcServer := grpc.NewServer()

	user.RegisterUserServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 4000: %v", err)
	}
}
