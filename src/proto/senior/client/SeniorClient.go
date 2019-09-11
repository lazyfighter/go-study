package main

import (
	"context"
	"go-study/src/proto/quick_start"
	"go-study/src/proto/senior"
	"google.golang.org/grpc"
	"log"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

func main() {
	conn, e := grpc.Dial(Address, grpc.WithInsecure())

	if e != nil {
		log.Println("create connection error")
	}
	client := proto_senior.NewHelloClient(conn)

	request := &proto_senior.HelloRequest{
		People: &proto_quick_start.Person{
			Name:  "ss",
			Email: "aa",
		},
	}
	response, e := client.SayHello(context.Background(), request)
	if e != nil {
		log.Println("call grpc server error", e)
	}
	log.Println(response.String())
}
