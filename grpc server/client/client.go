package main

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	chat "github.com/PranitRout07/grpc/chat"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// )

// func main() {

// 	mux := http.NewServeMux()

// 	mux.HandleFunc("/post",handlePost)

// 	fmt.Println("Running on port 4002...")
// 	http.ListenAndServe(":4002",mux)


// }
// type reqMsg struct{
// 	Msg string `json:"msg"`
// }
// func handlePost( w http.ResponseWriter,r *http.Request){

// 	m := reqMsg{}
// 	err := json.NewDecoder(r.Body).Decode(&m)
// 	if err!=nil{
// 		fmt.Println(err)
// 	}

// 	conn, err := grpc.NewClient("localhost:4000", grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil{
// 		panic(err)
// 	}

// 	client := chat.NewTestClient(conn)
	
// 	req := chat.HelloReq{MsgReq: m.Msg}

// 	client.Chat(context.TODO(),&req)

// }