package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	api "github.com/surenraju/grpc_helloworld/greetingservice"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

type greetServiceServer struct {
}

func main() {

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterGreetServiceServer(grpcServer, &greetServiceServer{})
	grpcServer.Serve(lis)
}

func (s *greetServiceServer) Greet(ctx context.Context, req *api.GreetRequest) (*api.GreetResponse, error) {
	return &api.GreetResponse{Greeting: fmt.Sprintf("Hello %s", req.Name)}, nil
}
