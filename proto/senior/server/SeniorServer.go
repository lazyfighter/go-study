package main

import (
	"context"
	"go-study/proto/quick_start"
	"go-study/proto/senior"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
	"net"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

// 定义helloService并实现约定的接口
type helloService struct{}

// HelloService Hello服务
var HelloService = helloService{}

// SayHello 实现Hello服务接口
func (h helloService) SayHello(ctx context.Context, request *proto_senior.HelloRequest) (*proto_senior.HelloResponse, error) {
	resp := new(proto_senior.HelloResponse)
	resp.People = &proto_quick_start.Person{}
	resp.People.Name = request.People.Name + "response"
	resp.People.Email = request.People.Email + "response"
	return resp, nil
}
func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}
	// 实例化grpc Server
	s := grpc.NewServer()
	// 注册HelloService
	proto_senior.RegisterHelloServer(s, HelloService)
	log.Println("Listen on " + Address)
	s.Serve(listen)
}
