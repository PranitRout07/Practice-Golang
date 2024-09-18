package main

type Person struct {
	Id          string
	Name        string
	Email       string
	PhoneNumber string
}

var nextID int32 = 1
var persons = make(map[int32]Person)

type server struct {
	UnimplementedPersonServiceServer
}

