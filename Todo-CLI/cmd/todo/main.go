package main

import (
	"flag"
	"fmt"
	"os"

	// "strings"

	"github.com/pranit007/todo"
)

const todoFileName = ".todo.json"

func main() {
	//adding flags
	//All these assigned flags are pointers . When trying to use it , it needs to deference .
	task := flag.String("task", "", "Add a new task")
	list := flag.Bool("list", false, "List all the tasks")
	complete := flag.Int("complete", 0, "Item to be completed")
	flag.Parse()
	li := &todo.List{} //li pointing to the address of the list .
	err := li.Get(todoFileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	//The below code is written using OS Args .

	// switch {
	// case len(os.Args)==1:
	// 	for _ , item := range *li {
	// 		fmt.Println(item.Task)
	// 	}
	// default:
	// 	item := strings.Join(os.Args[1:],"")
	// 	li.Add(item)

	// 	if err:= li.Save(todoFileName);err!=nil{
	// 		fmt.Fprintln(os.Stderr,err)
	// 		os.Exit(1)
	// 	}
	// }

	// for items := range *li {
	// 	l := *li
	// 	fmt.Println(l[items].Task)
	// 	fmt.Println("---")
	// }

	//The below code is written using flags package .

	switch {
	case *list:
		for _, item := range *li {
			//printing only the not completed items .
			if !item.Done {
				fmt.Println(item.Task)
			}
		}
	case *complete > 0:
		err := li.Complete(*complete)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		err = li.Save(todoFileName)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *task != "":
		li.Add(*task)

		err := li.Save(todoFileName)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid flag")
		os.Exit(1)
	}
}
