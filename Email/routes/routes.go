package routes

import (
	"github.com/PranitRout07/Practice-Golang/Email/controllers"
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/datas",controllers.GetSendDetails).Methods("GET")
	r.HandleFunc("/msg",controllers.UserEmailData).Methods("POST")
	r.HandleFunc("/send",controllers.SendEmailDetails).Methods("GET")
	return r
}