package main

import "fmt"

//custom function

type Execution func(string)

func Execute(fn Execution){
	fn("HELLO")
}


//Declare
type Store interface{
	DBStore(string)error
}

type DBStore struct{}


func(m *DBStore) myExecute(s string) {
	fmt.Println("my execute",s)
}

func main(){
	x := &DBStore{}
	Execute(x.myExecute)
}