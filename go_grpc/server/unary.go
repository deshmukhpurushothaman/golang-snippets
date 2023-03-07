package main

import (
	"context"

	pb "github.com/deshmukhpurushothaman/go-snippets/go_grpc/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "hello",
	}, nil
}
