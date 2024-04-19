package main

import (
	"fmt"
	"time"
)
//Here there are five go routines (four anonymous , one main)
//Main go routine continuosly executes the for statement and , here every half second it will either
//print channel 1 or channel 3 content and every two second it can any of the one channel content . 
func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)
	channel3 := make(chan string)
	channel4 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			channel1 <- "500ms : Hello this is channel-1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "2 sec : Hello this is channel-2"
		}
	}()
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			channel3 <- "500ms : Hello this is channel-3"
		}
	}()
	go func() {
		time.Sleep(time.Second * 2)
		channel4 <- "2 sec : Hello this is channel-4"
	}()
	for {
		select {
		case msg1 := <-channel1:
			fmt.Println(msg1)
		case msg2 := <-channel2:
			fmt.Println(msg2)
		case msg3 := <-channel3:
			fmt.Println(msg3)
		case msg4 := <-channel4:
			fmt.Println(msg4)
		}

	}
}