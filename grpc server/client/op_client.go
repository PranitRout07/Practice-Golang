package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	op "github.com/PranitRout07/grpc/operations"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


type Numbers struct{
	Num1 float32 `json:"num1"`
	Num2 float32 `json:"num2"`
	Operator string `json:"op"`
}

type Result struct{
	Res float32 `json:"result"`
}



func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/post",handlePost)

	fmt.Println("Running on port 4002...")
	http.ListenAndServe(":4002",mux)


}

func handlePost( w http.ResponseWriter,r *http.Request){

	m := Numbers{}
	err := json.NewDecoder(r.Body).Decode(&m)
	if err!=nil{
		fmt.Println(err)
	}

	conn, err := grpc.NewClient("localhost:4000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		fmt.Println(err)
	}

	client := op.NewOperationsClient(conn)
	
	req := op.Numbers{Num1: m.Num1, Num2: m.Num2}

	
	if m.Operator == "+"{
		res,err := client.Add(context.TODO(),&req)
		if err!=nil{
			b,_ := json.Marshal(fmt.Sprintf("Error :",err))
			w.Write(b)
		}

		b,_ := json.Marshal(Result{Res: res.Num})
		w.Write(b)

	}else if m.Operator == "-"{
		res,err := client.Sub(context.TODO(),&req)
		if err!=nil{
			b,_ := json.Marshal(fmt.Sprintf("Error :",err))
			w.Write(b)
		}

		b,_ := json.Marshal(Result{Res: res.Num})
		w.Write(b)


	}else if m.Operator == "*"{

		res,err := client.Mul(context.TODO(),&req)
		if err!=nil{
			b,_ := json.Marshal(fmt.Sprintf("Error :",err))
			w.Write(b)
		}

		b,_ := json.Marshal(Result{Res: res.Num})
		w.Write(b)

	}else if m.Operator == "/"{
		res,err := client.Div(context.TODO(),&req)
		if err!=nil{
			// fmt.Println(err)
			b,_ := json.Marshal(fmt.Sprintf("Error :",err))
			w.Write(b)
			return
		}

		b,_ := json.Marshal(Result{Res: res.Num})
		w.Write(b)
	}
	

}