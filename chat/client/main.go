package main

import (
	pb "chat-microservices/chat/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect %v\n", err)
	}
	log.Println("Connected to server")
	defer conn.Close()

	c := pb.NewChatServiceClient(conn)
	doChat(c)

}
