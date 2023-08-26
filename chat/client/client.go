package main

import (
	"bufio"
	pb "chat-microservices/chat/proto"
	"context"
	"fmt"
	"io"
	"log"
	"os"
)

func doChat(c pb.ChatServiceClient) {
	for {
		stream, err := c.Chat(context.Background())
		if err != nil {
			log.Fatalf("Error while creating stream: %v\n", err)
		}
		req := &pb.ChatRequest{}

		sc := bufio.NewReader(os.Stdin)
		input, _ := sc.ReadString('\n')
		req.Message1 = input

		stream.Send(req)

		go func() {
			for {
				res, err := stream.Recv()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatalf("Error while receiving: %v\n", err)
					break
				}
				fmt.Printf("Allen: %v", res.Message2)
			}
		}()
	}
}
