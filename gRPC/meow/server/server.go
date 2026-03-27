package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	pb "meow/meowrpc"
	"net"
	"strings"

	"google.golang.org/grpc"
)

type MeowRPCServer struct {
	pb.UnimplementedMeowRPCServer
	Kaomoji []string
}

func (m *MeowRPCServer) EchoMeow(ctx context.Context, request *pb.Request) (response *pb.Response, err error) {
	message := fmt.Sprintf("%s %s", request.Message, m.Kaomoji[rand.Intn(len(m.Kaomoji))])

	return &pb.Response{Message: message}, nil
}

func (m *MeowRPCServer) ManyMeow(request *pb.Request, stream pb.MeowRPC_ManyMeowServer) error {
	req := request.Message
	response := &pb.Response{}

	for _, runeValue := range req {
		message := fmt.Sprintf("%s %s", string(runeValue), m.Kaomoji[rand.Intn(len(m.Kaomoji))])

		response.Message = message
		if err := stream.Send(response); err != nil {
			return err
		}
	}

	return nil
}

func (m *MeowRPCServer) ImpatientMeow(stream pb.MeowRPC_ImpatientMeowServer) error {
	var chiChat []string

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			chiChat = append(chiChat, m.Kaomoji[rand.Intn(len(m.Kaomoji))])

			message := strings.Join(chiChat, " ")

			return stream.SendAndClose(
				&pb.Response{Message: message},
			)
		}

		if err != nil {
			return err
		}

		chiChat = append(chiChat, req.Message)
	}
}

func (m *MeowRPCServer) Conversation(stream pb.MeowRPC_ConversationServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		message := fmt.Sprintf("%s %s", req.Message, m.Kaomoji[rand.Intn(len(m.Kaomoji))])

		if err := stream.Send(&pb.Response{Message: message}); err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	pb.RegisterMeowRPCServer(grpcServer, &MeowRPCServer{
		Kaomoji: []string{
			"^•𖥦•^.ᐟ",
			"^›⩊‹^ ੭",
			"ฅ(´꒳ `ฅ)ꪆ",
			"ฅ^•ω•^ฅ",
			"(＾• ω •＾)",
		},
	})

	grpcServer.Serve(lis)
}
