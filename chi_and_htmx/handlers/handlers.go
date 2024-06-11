package handlers

import "net/http"

func HomeHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello"))
}

func PostHandler(w http.ResponseWriter, r *http.Request){

}

func ProductHandler(w http.ResponseWriter, r *http.Request){

}