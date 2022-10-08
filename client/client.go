package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/Xart3mis/GoHkarComms/client_data_pb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up connection with the grpc server
	conn, err := grpc.Dial("172.21.109.80:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error while making connection, %v", err)
	}

	// Create a client instance
	c := pb.NewConsumerClient(conn)
	ctx, _ := context.WithCancel(context.Background())
	x, err := c.SubscribeOnScreenText(ctx, &pb.ClientDataRequest{ClientId: "id_1"})
	if err != nil {
		log.Fatalf("Error while subscribing, %v", err)
	}

	for {
		resp, err := x.Recv()
		if err != nil {
			switch err {
			case io.EOF:
				log.Println("End of stream")

			default:
				log.Println("Error: ", err)
			}
		}
		fmt.Println(resp)
	}

	// x.CloseSend()
	// cancel()
	// conn.Close()
}
