package main

import (
	"context"
	"go-study/src/proto/echo"
	"google.golang.org/grpc"
	"log"
	"net"
)

const Address = "127.0.0.1:7777"

type EchoService struct{}

var echoService = EchoService{}

func (EchoService) Echo(context context.Context, request *proto_echo.Request) (*proto_echo.Response, error) {
	response := new(proto_echo.Response)
	response.Text = request.Text
	log.Println("echo: ", request.Text)
	return response, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto_echo.RegisterEchoServer(s, echoService)
	log.Println("Listen on " + Address)
	s.Serve(listen)

}
