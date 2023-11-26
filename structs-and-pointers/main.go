package main

import (
	"fmt"
	"unsafe"
)

type Book struct {
	title      string
	author     string
	numOfPages int
	isSaved    bool
}

func (book Book) SaveBook1() {
	//This does not changes the original data.This works on the copy of the Book
	book.title = "Hello"
	fmt.Println(book)
}
func SaveBook2(book Book) {
	//This does not changes the original data.This works on the copy of the Book
	book.title = "New city"
	book.numOfPages = 103
	fmt.Println(book)
}
func (book *Book) SaveBook3() {
	//This changes the original data.This does not work on the copy of the Book
	book.title = "King"
	book.author = "Joe"
	book.numOfPages = 200
	book.isSaved = true
	fmt.Println(book)
}
func SaveBook4(book *Book) {
	//This changes the original data.This does not work on the copy of the Book
	book.title = "The Warrior"
	book.isSaved = true
	fmt.Println(book)

}
func main() {
	book := Book{
		title:      "Original Title",
		author:     "Original Author",
		numOfPages: 100,
		isSaved:    false,
	}
	//Without pointer
	book.SaveBook1()
	SaveBook2(book)
	//With Pointer
	book.SaveBook3()

	b := &book
	SaveBook4(b)

	structOptimisation()
}

func structOptimisation(){
	type unOptimized struct{
		A bool
		B float64
		C int32
	}
	a := unOptimized{}
	fmt.Println("Alignment of bool",unsafe.Alignof(bool(true)))
	fmt.Println("Alignment of float64",unsafe.Alignof(float64(0.0)))
	fmt.Println("Alignment of int32",unsafe.Alignof(int32(0)))
	fmt.Println("Alignment of struct",unsafe.Alignof(a))
	


	type Optimized struct{
		B float64
		C int32
		A bool
	}
	b := Optimized{}
	fmt.Println("Size of struct before optimization",unsafe.Sizeof(a))
	fmt.Println("Size of struct after optimization",unsafe.Sizeof(b))
}