package main

import (
	"fmt"
	"log"
	"net/http"
)
func formHandler(w http.ResponseWriter,r *http.Request){
	if err := r.ParseForm(); err!= nil {
		fmt.Fprintf(w, "ParseForm() err:%v",err)
		return
	}
	fmt.Fprintf(w,"Post Request Successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w,"Name = %s\n",name)
	fmt.Fprintf(w,"Address = %s\n",address)
}
func helloHandler(w http.ResponseWriter,r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w ,"404 not found",http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w,"method is not supported",http.StatusNotFound)
	}
	fmt.Fprintf(w,"hello!!")
}
func main(){
	fileserver := http.FileServer(http.Dir("./static")) // index file is automatically founded by golang
	http.Handle("/",fileserver) 
	http.HandleFunc("/form",formHandler) 
	http.HandleFunc("/hello",helloHandler)

	fmt.Print("Starting at port 8000\n")
	if err := http.ListenAndServe(":8000",nil);err !=nil {
		log.Fatal(err)
	}
}