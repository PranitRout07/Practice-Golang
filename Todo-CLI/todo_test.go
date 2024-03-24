package todo_test

import (
	"fmt"
	"os"

	"testing"

	"github.com/pranit007/todo"
)

func TestAddTask(t *testing.T) {
	li := todo.List{}
	demoTask := "Task-01"
	li.Add(demoTask)
	if li[0].Task != demoTask {
		t.Errorf("Want %q got %q", demoTask, li[0].Task)
	}
	fmt.Println("SUCCESS")
}

func TestCompleteTask(t *testing.T) {
	li := todo.List{}
	demoTask := "Task-01"
	li.Add(demoTask)
	if li[0].Task != demoTask {
		t.Errorf("want %q got %q", demoTask, li[0].Task)
	}
	if li[0].Done == true {
		t.Errorf("New task should not be completed.")
	}
	li.Complete(1)
	if !li[0].Done {
		t.Errorf("New task should be completed.")
	}
	fmt.Println("SUCCESS")
}

func TestDeleteTask(t *testing.T) {
	li := todo.List{}
	demoTasks := []string{"Task-01", "Task-02", "Task-03"}
	for task := range demoTasks {
		li.Add(demoTasks[task])
	}
	if len(li) != 3 {
		t.Errorf("size of list should be %d but got %d", 3, len(li))
	}

	li.Delete(2)
	if li[1].Task != demoTasks[2] {
		t.Errorf("Task is not deleted")
	}
	fmt.Println("SUCCESS")

}

func TestSaveAndGetFile(t *testing.T) {
	os.WriteFile("test.json", []byte{}, 0644)
	li := todo.List{}
	demoTasks := []string{"Task-01", "Task-02", "Task-03"}
	for task := range demoTasks {
		li.Add(demoTasks[task])
	}
	if len(li) != 3 {
		t.Errorf("size of list should be %d but got %d", 3, len(li))
	}
	err := li.Save("test.json")
	if err!=nil{
		t.Error("Error occured : ",err)
	}
	content , err := os.ReadFile("test.json")
	if err!=nil{
		t.Error("Error occured :",err)
	}
	if len(content)==0{
		t.Errorf("There is no content in the file .")
	}

	// fmt.Println(string(content))
	err = os.Remove("test.json")
	if err!=nil{
		t.Error("Error occured :",err)
	}
	fmt.Println("SUCCESS")

}
