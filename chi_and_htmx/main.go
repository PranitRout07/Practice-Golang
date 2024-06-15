package main

import (
	"log"
	"net/http"
	"os"

	"github.com/PranitRout07/Practice-Golang/chi_and_htmx/initializers"
	"github.com/PranitRout07/Practice-Golang/chi_and_htmx/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()

}

func main() {
	r := routes.Routes()
	port := os.Getenv("port")
	fs := http.FileServer(http.Dir("./templates"))
	r.Handle("/templates/*", http.StripPrefix("/templates/", fs))	
	log.Fatal(http.ListenAndServe(port, r))

}
