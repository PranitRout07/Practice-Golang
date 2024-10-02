// we have 500 tickets
// 2000 people trying to buy the ticket at single time
package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func buyTicket(wg *sync.WaitGroup, tickets *int, userID int) {
	defer wg.Done()
	mutex.Lock()
	if *tickets > 0 {
		*tickets = *tickets - 1
		fmt.Println("userID", userID, "has bought the tickets and number of tickets remaining",*tickets)
	}
	fmt.Println("userID", userID, "can not bought the tickets")
	mutex.Unlock()
}

func main() {
	tickets := 500
	var wg sync.WaitGroup

	for userId := 0; userId < 2000; userId++ {
		wg.Add(1)
		go buyTicket(&wg, &tickets, userId)
	}
	wg.Wait()

}
