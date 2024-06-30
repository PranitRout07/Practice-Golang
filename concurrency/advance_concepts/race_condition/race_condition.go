package main

import (
	"fmt"
	"sync"
)

//race condition occurs when multiple users try to access a critical memory at a same time without any restriction.

func main(){
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	var score = []int{0}
	wg.Add(3)
	go func(wg *sync.WaitGroup,m *sync.Mutex){
		fmt.Println("Go routine 1")
		m.Lock()
		score = append(score, 1)
		m.Unlock()
		wg.Done()
	}(wg,mut)
	go func(wg *sync.WaitGroup,m *sync.Mutex){
		fmt.Println("Go routine 2")
		m.Lock()
		score = append(score, 2)
		m.Unlock()
		wg.Done()
	}(wg,mut)
	go func(wg *sync.WaitGroup,m *sync.Mutex){
		fmt.Println("Go routine 3")
		m.Lock()
		score = append(score, 3)
		m.Unlock()
		wg.Done()
	}(wg,mut)

	wg.Wait()



}