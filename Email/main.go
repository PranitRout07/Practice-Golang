package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PranitRout07/Practice-Golang/Email/routes"
)

// func init() {
// 	err := initializers.LoadEnvVariables()
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// }

func main() {
	r := routes.Routes()
	fmt.Println("Running in port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
