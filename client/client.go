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
	c.RegisterClient(context.Background(), &pb.ClientDataRequest{ClientId: "id_1"})
	x, err := c.SubscribeOnScreenText(context.Background(), &pb.ClientDataRequest{ClientId: "id_1"})
	if err != nil {
		log.Fatalf("Error while subscribing, %v", err)
	}
	for {
		// uhh, _ := c.GetExecCommand(context.Background(), &pb.ClientDataRequest{ClientId: "id_1"})
		// if uhh != nil {
		// 	fmt.Println(uhh)
		// }
		// c.SetExecOutput(context.TODO(), &pb.ClientExecOutput{Output: "OK"})
		fmt.Println(x.Recv())
	}
}
