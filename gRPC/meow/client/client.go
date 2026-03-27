package main

import (
	"context"
	"io"
	"log"
	pb "meow/meowrpc"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func EchoMeow(client pb.MeowRPCClient, message string) {
	log.Println("EchoMeow Calling...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.EchoMeow(ctx, &pb.Request{Message: message})
	if err != nil {
		log.Fatalf("client.EchoMeow failed: %v", err)
	}
	log.Println(res.Message)
}

func ManyMeow(client pb.MeowRPCClient, message string) {
	log.Println("ManyMeow Calling...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.ManyMeow(ctx, &pb.Request{Message: message})

	if err != nil {
		log.Fatalf("client.ManyMeow failed: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("client.ManyMeow failed: %v", err)
		}
		log.Println(res.Message)
	}
}

func ImpatientMeow(client pb.MeowRPCClient, messages []string) {
	log.Println("ImpatientMeow Calling...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.ImpatientMeow(ctx)
	if err != nil {
		log.Fatalf("client.ImpatientMeow failed: %v", err)
	}

	for _, message := range messages {
		if err := stream.Send(&pb.Request{Message: message}); err != nil {
			log.Fatalf("client.ImpatientMeow: stream.Send(%v) failed: %v", message, err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("client.ImpatientMeow failed: %v", err)
	}
	log.Println(res.Message)
}

func Conversation(client pb.MeowRPCClient, messages []string) {
	log.Println("Conversation Calling...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.Conversation(ctx)
	if err != nil {
		log.Fatalf("client.Conversation failed: %v", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}

			if err != nil {
				log.Fatalf("client.Conversation failed: %v", err)
			}

			log.Println(res.Message)
		}
	}()

	for _, message := range messages {
		if err := stream.Send(&pb.Request{Message: message}); err != nil {
			log.Fatalf("client.Conversation: stream.Send(%v) failed: %v", message, err)
		}
	}

	stream.CloseSend()
	<-waitc
}

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient("localhost:50051", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewMeowRPCClient(conn)

	EchoMeow(client, "Hello!")
	ManyMeow(client, "你好嗎?")
	ImpatientMeow(client, []string{
		"I",
		"LOVE",
		"CAT",
	})
	Conversation(client, []string{
		"I",
		"LOVE",
		"CAT",
	})
}
