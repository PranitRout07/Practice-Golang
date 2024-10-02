package main

import (
	"fmt"
	"sync"
)

func manageTickets(tickets *int, done chan bool, userChan chan int) {
	for {
		select{
		case userID := <-userChan:
			if *tickets > 0 {
				*tickets--
				fmt.Println("UserID",userID,"had bought the tickets and remaining tickets are",*tickets)
			}else{
				fmt.Println("UserID",userID,"had not bought the tickets")
				}
		case <- done :
			fmt.Println("Tickets remaining:",tickets)		
		}
		

	}
}

func buyTickets(wg *sync.WaitGroup, userID int, userChan chan int) {
	defer wg.Done()
	userChan <- userID
}

func main() {
	tickets := 500
	userChan := make(chan int)
	doneChan := make(chan bool)
	go manageTickets(&tickets, doneChan, userChan)
	var wg sync.WaitGroup
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go buyTickets(&wg, i, userChan)
	}
	wg.Wait()
	doneChan <- true

}
