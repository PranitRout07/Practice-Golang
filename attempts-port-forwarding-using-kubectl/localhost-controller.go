package main

import (
    "bufio"
    "fmt"
    _"os"
    "os/exec"
    "strings"
)

func main() {
    for {
        cmd := exec.Command("kubectl", "port-forward", "service/my-service", "8501:8501")
        stdout, err := cmd.StdoutPipe()
        if err != nil {
            fmt.Println("Error creating StdoutPipe for Cmd:", err)
            return
        }

        if err := cmd.Start(); err != nil {
            fmt.Println("Error starting Cmd:", err)
            return
        }
        // Reads and stores the output line by line
        scanner := bufio.NewScanner(stdout)
        for scanner.Scan() {
            text := scanner.Text()
            fmt.Println(text)

            // Checks if the text contains the error message
            if strings.Contains(text, "error forwarding port") {
                fmt.Println("Error occurred. Restarting port-forwarding...")
                break
            }
        }

        if err := cmd.Wait(); err != nil {
            fmt.Println("Command finished with error:", err)
        }
    }
}