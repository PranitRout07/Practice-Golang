package main 

type Task struct {
	ID int64 `json:"id"`
	TaskName string `json:"taskname"`
	Status bool `json:"status"`
}