package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PranitRout07/Practice-Golang/Email/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err!=nil{
		log.Fatal(err)
	}

	port := os.Getenv("port")

	r := routes.Routes()
	fmt.Println("Running in port",port ,"...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
