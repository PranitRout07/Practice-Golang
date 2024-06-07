package routes

import (
	"github.com/PranitRout07/Practice-Golang/Email/controllers"
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/msg",controllers.UserEmailData).Methods("POST")
	return r
}