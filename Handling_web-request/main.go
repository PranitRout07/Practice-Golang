package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://loremipsum.io/"

func main() {
	response,err := http.Get(url)
	if err!=nil {
		panic(err)
	}
	databyte,err := ioutil.ReadAll(response.Body)
	if err!=nil{
		panic(err)
	}
	content:=string(databyte)
	fmt.Println(content)

	defer response.Body.Close()

}