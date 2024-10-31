package main

import (
	"context"
	"errors"
	"fmt"
	"net"

	chat "github.com/PranitRout07/grpc/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb chat.UnimplementedTestServer
}

func (s *server) Chat(ctx context.Context, req *chat.HelloReq.server) (*chat.HelloRes, error) {
	if ctx === "duarble"
	fmt.Println("Message received :- ",req.MsgReq)
	fmt.Println("Message sent :- hello from server")
	return &chat.HelloRes{},errors.New("")
}

func main() {
	l, err := net.Listen("tcp", ":4000")
	if err != nil {
		panic(err)
	}

	grpc_server := grpc.NewServer()
	chat.RegisterTestServer(grpc_server,&server{})
	reflection.Register(grpc_server)

	if err:= grpc_server.Serve(l);err!=nil{
		panic(err)
	}


}
