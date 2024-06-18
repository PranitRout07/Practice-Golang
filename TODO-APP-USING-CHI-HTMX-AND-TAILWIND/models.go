package main 

type Task struct {
	ID int `json:"id"`
	TaskName string `json:"taskname"`
	Status bool `json:"status"`
}