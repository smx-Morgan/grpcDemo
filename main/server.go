package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpcDemo/main/hello"
	"log"
	"net"
)

const (
	// Address gRPC服务地址
	Address = ":50052"
)

type server struct {
	pb.UnimplementedHelloServer
}

// SayHello 实现Hello服务接口
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.", in.Name)
	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	log.Println("Listen ", Address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	// 实例化 grpc Server
	s := grpc.NewServer()
	// 注册HelloService
	pb.RegisterHelloServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
