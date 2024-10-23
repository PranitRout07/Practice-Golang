package main

import (
	"fmt"
	"sync"
	"time"
)

// 1. DIRECT METHOD

// func main() {
// 	arr := []int{1, 2, 3, 4, 5}
// 	res := []int{}

// 	now := time.Now()
// 	for _, i := range arr {      //this way takes appoximatesly just above 5 sec
// 		time.Sleep(1*time.Second)
// 		res = append(res, i*2)
// 	}
// 	fmt.Println(time.Since(now))
// 	fmt.Println("RESULT",res)

// }



// Using go routines and lock

// var lock sync.Mutex

// func delay(){
// 	time.Sleep(time.Second*1)
// }

// func process(wg *sync.WaitGroup,arr *[]int , val int){
// 	defer wg.Done()
// 	delay()

// 	lock.Lock()
// 	*arr = append(*arr, val*2)
// 	lock.Unlock()
// }

// func main (){
// 	arr := []int{1, 2, 3, 4, 5}
// 	res := []int{}
// 	wg := &sync.WaitGroup{}
// 	now := time.Now()
// 	for _,i := range arr{
// 		wg.Add(1)
// 		go process(wg,&res,i)
// 	}
// 	wg.Wait()
// 	fmt.Println("TIME TAKEN:",time.Since(now))
// 	fmt.Println(res)
// }

//confinement way



func smallProcess(data int)int{
	time.Sleep(time.Second*1)
	return data*2
}

func process(wg *sync.WaitGroup,ele *int , val int){
	defer wg.Done()


	*ele = smallProcess(val)

}

func main (){
	arr := []int{1, 2, 3, 4, 5}
	res := make([]int,len(arr))  
	wg := &sync.WaitGroup{}
	now := time.Now()
	for i,val := range arr{
		wg.Add(1)
		go process(wg,&res[i],val)
	}
	wg.Wait()
	fmt.Println("TIME TAKEN:",time.Since(now))
	fmt.Println(res)
}
