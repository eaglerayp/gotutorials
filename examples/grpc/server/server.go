package main

import (
	"context"
	"log"
	"net"

	pb "github.com/eaglerayp/gotutorials/examples/grpc/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":10999"

type Server struct {
}

func (s *Server) Echo(ctx context.Context, req *pb.HiRequest) (*pb.HiResponse, error) {
	resp := &pb.HiResponse{Message: req.Message}
	return resp, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// grpc Server is the struct which do the tcp listen
	s := grpc.NewServer()
	// one service can only regist one handler,
	// but one grpc Server can handle multi grpc service server
	pb.RegisterEchoServiceServer(s, &Server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Println("start gRPC server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
