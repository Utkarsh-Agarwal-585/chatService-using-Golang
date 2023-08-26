package main

import (
	"bufio"
	pb "chat-microservices/chat/proto"
	"fmt"
	"io"
	"log"
	"os"
)

func (s *Server) Chat(stream pb.ChatService_ChatServer) error {
	log.Println("Chat was invoked")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}
		fmt.Printf("Bob: %s", req.Message1)

		go func() {
			for {
				sc := bufio.NewReader(os.Stdin)
				input, _ := sc.ReadString('\n')

				err := stream.Send(&pb.ChatResponse{
					Message2: input,
				})
				if err != nil {
					log.Fatalf("Error while sending data to client: %v\n", err)
				}
			}
		}()

	}
}
