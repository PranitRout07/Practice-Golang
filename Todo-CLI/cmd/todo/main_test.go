package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
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
