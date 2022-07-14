package main

import (
	"context"
	"crypto/tls"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "go-server/auto"
)

func TestExample(t *testing.T) {
	config := &tls.Config{
		InsecureSkipVerify: true,
	}
	talker, err := grpc.Dial("localhost:50000", grpc.WithTransportCredentials(credentials.NewTLS(config)))
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer talker.Close()

	id := []byte{
		0, 0, 0, 0, 0, 0, 0, 0,
	}

	client := pb.NewServiceAClient(talker)
	res_cs, err := client.GetList(context.Background(), &pb.ID{
		Id: id,
	})
	if err != nil {
		t.Fatalf("failed to request: %v", err)
	}
	t.Logf("response: %v\n", res_cs)
}
