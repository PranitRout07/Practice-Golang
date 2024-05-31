package routes

import (
	"github.com/PranitRout07/Practice-Golang/basic_api_using_mongodb/cmd"
	"github.com/gorilla/mux"
)

func Route() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/movies",cmd.GetAllMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",cmd.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies",cmd.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",cmd.DeleteOneMovie).Methods("DELETE")
	r.HandleFunc("/movies/deleteall",cmd.DeleteAllMovies).Methods("DELETE")
	return r 
}
