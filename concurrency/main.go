package main

import "fmt"

func main() {
	// goChannel := make(chan string)

	// go func() {
	// 	goChannel <- "data"
	// }()

	// data := <-goChannel
	// fmt.Println(data)

	charChannel := make(chan string, 3)
	chars := []string{"a","b","c"}
		for _,s := range chars{
			select{
			case charChannel <-s:

			}
		}

	close(charChannel)
	for result := range charChannel {
		fmt.Println(result)
	}


}