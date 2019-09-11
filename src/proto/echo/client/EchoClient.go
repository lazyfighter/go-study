package main

import (
	"context"
	"go-study/src/proto/echo"
	"google.golang.org/grpc"
	"log"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:7777"
)

func main() {
	conn, e := grpc.Dial(Address, grpc.WithInsecure())
	defer conn.Close()

	if e != nil {
		log.Println("open echo client connection error", e)
	}

	client := proto_echo.NewEchoClient(conn)

	response, e := client.Echo(context.Background(), &proto_echo.Request{
		Text: "liping",
	})

	if e != nil {
		log.Println("call echo server error", e)
	}

	log.Println(response)
}
