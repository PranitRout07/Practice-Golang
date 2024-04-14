package main

import (
	"fmt"
	"time"
)
func main(){
	goroutine_leakage()
}

func goroutine_leakage() {
	go func() {
		for {
			select {
			default:
				fmt.Println("DOING WORK")
			}
		}
	}()

	time.Sleep(time.Hour * 300)
}
