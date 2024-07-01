package main

import (
	"fmt"
	"log"
	"sync"
)

// channels are a way that allows the go routines to talk with each other
func main() {
	mych := make(chan int, 2)
	// mych<- 5 //channel stores the 5
	//fmt.Println(<-mych) //trying to print, but this will show error as it needs a go routine to receive

	wg := &sync.WaitGroup{}

	wg.Add(2)
	//receive the value from channel
	go func(wg *sync.WaitGroup, ch chan int) {
		val, isChannelOpen := <-ch
		if !isChannelOpen {
			log.Fatal("Channel is not open,can not proceed.")
		}
		fmt.Println(val)
		wg.Done()
	}(wg, mych)
	//enter the value into the channel
	go func(wg *sync.WaitGroup, ch chan int) {
		// var val int
		// fmt.Println("Enter a integer value: ")
		// fmt.Scanln(&val)
		// ch <- val
		ch<-0
		close(ch)

		wg.Done()
	}(wg, mych)

	wg.Wait()

}
