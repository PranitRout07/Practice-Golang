package main

import (
	"fmt"
	"os/exec"


	htgotts "github.com/hegedustibor/htgo-tts"
	voices "github.com/hegedustibor/htgo-tts/voices"
)

func main(){
	speech := htgotts.Speech{
		Folder: "audio", 
		Language: voices.English,
	}
	speech.Speak("Hello")
    cmd := exec.Command("wsl","./script.sh")

    // Run the command and capture the output
    _, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("Error executing script:", err)
        return
    }
}