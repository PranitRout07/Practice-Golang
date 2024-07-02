package main

import (
	"log"
	"net/http"

	"github.com/PranitRout07/Practice-Golang/login-logout/initializers"
	"github.com/PranitRout07/Practice-Golang/login-logout/routes"
)


func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	// initializers.Create_table()
}
func main() {
	r := routes.Routes()

	log.Fatal(http.ListenAndServe(":8000",r ))
	
}