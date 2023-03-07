package main

import (
	"context"
	"log"
	"time"

	pb "github.com/deshmukhpurushothaman/go-snippets/go_grpc/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Client Stream Started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not send names %v", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		log.Printf("Sent the req with name: %s", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("Client streaming finished")

	if err != nil {
		log.Fatalf("Error while receiving %v", err)
	}
	log.Printf("%v", res.Messages)
}
