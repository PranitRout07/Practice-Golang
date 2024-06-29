package main

import (
	"fmt"
	"net/http"
	"sync"
)
var signals = []string{}
var wg sync.WaitGroup
var mut sync.Mutex

func main() {
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://fb.com",
		"https://instagram.com",
		"https://go.dev",
	}

	for _,url := range urls {
		go getStatus(url)
		wg.Add(1)  // mentioning number of go routines it is calling
		           //here we are calling one go routine so value is 1
				   //if suppose again if i write go getStatus(url), i have write 2.
	}
	wg.Wait() //does not allow to finish the main method
	fmt.Println(signals)
}


func getStatus(url string){
	res, err := http.Get(url)
	defer wg.Done()
	if err!=nil{
		fmt.Println("Error:",err)
	}else{
		mut.Lock()
		signals = append(signals, url)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n",res.StatusCode,url)
	}
	
}
