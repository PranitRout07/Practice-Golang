package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//WITHOUT USING CONCURRENCY(This takes 6 seconds to complete)

	// start := time.Now()
	// PrintWithoutUsingConcurrency(2)
	// PrintWithoutUsingConcurrency(4)
	// fmt.Println("Total execution time: ",time.Since(start))


	//USING CONCURRENCY(This takes 4 seconds to complete)

	wg := sync.WaitGroup{}
	
	start := time.Now()
	wg.Add(2)
	go PrintUsingConcurrency(2,&wg)
	go PrintUsingConcurrency(4,&wg)
	wg.Wait()
	fmt.Println("Total execution time: ",time.Since(start))
}

func PrintWithoutUsingConcurrency(n time.Duration) {
	fmt.Println("Doing work...")
	time.Sleep(n * time.Second)
	fmt.Println("Done task..")

}

func PrintUsingConcurrency(n time.Duration, wg *sync.WaitGroup){

	fmt.Println("Doing work...")
	time.Sleep(n * time.Second)
	fmt.Println("Done task..")
	wg.Done()
}
