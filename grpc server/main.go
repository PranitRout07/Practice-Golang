package main 


import (
	"context"
	"fmt"
	"net"

	op "github.com/PranitRout07/grpc/operations"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


type Server struct {
	op.UnimplementedOperationsServer
}

func main(){

	l, err := net.Listen("tcp", ":4000")
	if err != nil {
		fmt.Println(err)
	}

	grpc_server := grpc.NewServer()
	op.RegisterOperationsServer(grpc_server,&Server{})
	reflection.Register(grpc_server)

	if err:= grpc_server.Serve(l);err!=nil{
		fmt.Println(err)
	}
}

func (s *Server)Add(ctx context.Context, nums *op.Numbers) (*op.Result, error){
	return &op.Result{Num: nums.Num1 + nums.Num2},nil
}

func (s *Server)Div(ctx context.Context, nums *op.Numbers) (*op.Result, error){
	if nums.Num2 == 0{
		fmt.Println("0")
		return &op.Result{Num: 0},fmt.Errorf("can not divide by 0")
	}
	fmt.Println("1")
	return &op.Result{Num: nums.Num1 / nums.Num2},nil
}
func (s *Server)Sub(ctx context.Context, nums *op.Numbers) (*op.Result, error){
	return &op.Result{Num: nums.Num1 - nums.Num2},nil
}
func (s *Server)Mul(ctx context.Context, nums *op.Numbers) (*op.Result, error){
	return &op.Result{Num: nums.Num1 * nums.Num2},nil
}