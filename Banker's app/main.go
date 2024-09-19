package main

import (
	"fmt"
	"log"
)
//inside interface defined the necessary functions
//interface is attached to the struct 
//struct variable passed to fucntion 
//if you don't define the methods in the same order as interface then if you are passing an argument then it will throw error

func main() {
	fmt.Println("hello")
	store,err := NewPostgresStore()
	if err!=nil{
		log.Fatal(err)
	}
	if err=store.Init();err!=nil{
		log.Fatal(err)
	}
	log.Println(store)
	server := NewAPIServer(":4000",store)
	server.Run()
}