package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	// Possible inputs : rock , paper , scissor
	// User will first enter the input
	// After this computer will randomly choose from rock,paper and scissor
	// Three outcomes of the game : win , lose , draw
	fmt.Println("Choose from rock, paper and scissor : ")
	scn := bufio.NewReader(os.Stdin)
	input, _ := scn.ReadString('\n')
	input = strings.TrimSpace(input)
	store := []string{"rock", "paper", "scissor"}

	computer_turn := store[rand.Intn(len(store))]
	
	if input == computer_turn {
		fmt.Println("Draw")
	} else if input == "rock" && computer_turn == "scissor" {
		fmt.Println("User wins.Hooray!!!")
	} else if input == "scissor" && computer_turn == "paper" {
		fmt.Println("User wins.Hooray!!!")
	} else if input == "paper" && computer_turn == "rock" {
		fmt.Println("User wins.Hooray!!!")
	} else {
		fmt.Println("I win.Haha...Come next time you newbie")
	}

}
