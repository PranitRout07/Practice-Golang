package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PranitRout07/Practice-Golang/basic_api_using_mongodb/routes"
)

func main() {
	r := routes.Route()
	
	fmt.Println("Running at port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
