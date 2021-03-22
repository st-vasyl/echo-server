package server_grpc

import (
	"context"
	"log"
	"net"

	"github.com/st-vasyl/echo-server/pkg/v1/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// echoService is the implementation of the gRPC echo service.
type echoService struct {
	// enables forward compatible implementations;
	// can be removed if we use '--go-grpc_out=require_unimplemented_servers=false' for protoc
	echo.UnimplementedEchoServer
}

func (*echoService) Echo(ctx context.Context, req *echo.EchoRequest) (*echo.EchoResponse, error) {
	return &echo.EchoResponse{
		Echo: req.Text,
	}, nil
}

func Run() {
	s := grpc.NewServer()
	reflection.Register(s)
	listener, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatal("Fatal error to run gRPC listener", err)
	}
	log.Println("gRPC server started on port 50051")

	echo.RegisterEchoServer(s, &echoService{})
	s.Serve(listener)
}
