package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/Xart3mis/GoHkarComms/client_data_pb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct {
	pb.ClientServer
}

func main() {

	// NewServer creates a gRPC server which has no service registered and has not started
	// to accept requests yet.
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// We are making use of the function that compiled proto made for us to register
	// our GRPC server so that the clients can make use of the functions tide to our
	// server remotely via the GRPC server (like MakeTransaction function)

	// The first argument is the grpc server instance
	// The second argument is the service who's methods we want to expose (in our case)
	// we have put it in this program only
	pb.RegisterClientServer(s, &server{})
	// Serve accepts incoming connections on the listener lis, creating a new ServerTransport
	// and service goroutine for each. The service goroutines read gRPC requests and then
	// call the registered handlers to reply to them.
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func (s *server) GetOnScreenText(ctx context.Context, in *pb.ClientDataRequest) (*pb.ClientDataResponse, error) {
	fmt.Println(in)

	return &pb.ClientDataResponse{ClientData: map[string]*pb.ClientData{
		in.ClientId: {
			OnScreenText: "hello",
			ShouldUpdate: true,
		},
	},
	}, nil
}
