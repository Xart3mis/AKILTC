package main

import (
	"fmt"
	"log"

	pb "github.com/Xart3mis/GoHkarComms/client_data_pb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up connection with the grpc server
	conn, err := grpc.Dial("0.0.0.0:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error while making connection, %v", err)
	}

	// Create a client instance
	c := pb.NewConsumerClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	x, err := c.SubscribeOnScreenText(ctx, &pb.ClientDataRequest{ClientId: "id_1"})
	if err != nil {
		log.Fatalf("Error while subscribing, %v", err)
	}
	for i := 0; i < 10000; i++ {
		fmt.Println(x.Recv())
	}

	defer func() {
		x.CloseSend()
		cancel()
		conn.Close()
	}()
}
