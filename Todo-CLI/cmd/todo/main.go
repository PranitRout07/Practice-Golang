package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/pranit007/todo"
)
const todoFileName = ".todo.json"
func main() {
	li := &todo.List{} //li pointing to the address of the list . 
	err := li.Get(todoFileName)
	if err!=nil{
		fmt.Fprintln(os.Stderr,err)
	}
	switch {
	case len(os.Args)==1:
		for _ , item := range *li {
			fmt.Println(item.Task)
		}
	default:
		item := strings.Join(os.Args[1:],"")
		li.Add(item)

		if err:= li.Save(todoFileName);err!=nil{
			fmt.Fprintln(os.Stderr,err)
			os.Exit(1)
		}
	}
}
