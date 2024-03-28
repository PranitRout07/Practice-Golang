package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	// "strings"
	"testing"
)

var (
	binName  = "todo"
	fileName = "todo.json"
)

func TestMain(m *testing.M) {
	if runtime.GOOS == "windows" {
		binName = binName + ".exe"
	}
	fmt.Println(binName)
	build := exec.Command("go","build","-o",binName)

	err := build.Run()

	if err!=nil{
		fmt.Fprintf(os.Stderr , "can not build the tool %s",err)
		os.Exit(1)
	}
	result := m.Run() 
	fmt.Println(result)
	os.Remove(binName)
	os.Remove(fileName)

	os.Exit(result)
}

func TestTodoCLI(t *testing.T){
	task := "Task-01"
	dir , err := os.Getwd()  // gives the current directory
	if err!=nil{
		t.Fatal(err)
	}
	fmt.Println(dir)
	cmdPath := filepath.Join(dir, binName)
	t.Run("Add new task", func (t*testing.T){
		cmd := exec.Command(cmdPath, "-task", task)
		fmt.Println(cmd)
		if err := cmd.Run();err!=nil{
			t.Fatal(err)
		}
	})

	t.Run("ListTasks", func(t *testing.T){
		cmd := exec.Command(cmdPath,"-list")
		output , err := cmd.CombinedOutput()
		// fmt.Println("Output:",string(output))
		fmt.Println("Output: ",string(output))
		if err!=nil{
			t.Fatal(err)
		}
		expected := task + "\n"
		if expected!=string(output){
			t.Errorf("Expected %q got %q",expected,string(output))
		}
	})
	err = os.Remove("C:\\Users\\prani\\OneDrive\\Desktop\\GolangPractice\\Practice-Golang\\Todo-CLI\\cmd\\todo\\.todo.json")
	if err!=nil{
		t.Error("Error occured :",err)
	}
}