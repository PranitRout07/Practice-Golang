package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Msg struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func main() {

	http.Handle("/",RateLimiter((Handler)))
	err := http.ListenAndServe(":8085",nil)
	if err!=nil{
		log.Println("Error occured::",err)
		return 
	}

}

func Handler(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	msg := Msg{
		Status: "Successful!",
		Body: "Thanks for visiting my website",
	}
	err := json.NewEncoder(w).Encode(&msg)
	if err!=nil{
		return 
	}
}