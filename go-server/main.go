package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "go-server/auto"
)

type AServer struct {
	pb.UnimplementedServiceAServer
}

func (s *AServer) GetList(ctx context.Context, in *pb.ID) (*pb.FileList, error) {
	fmt.Println("requested!")
	return &pb.FileList{
		List: []string{"ok"},
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile("public.crt", "private.pem")
	if err != nil {
		log.Fatalf("failed to setup TLS: %v", err)
	}

	server := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterServiceAServer(server, &AServer{})
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
