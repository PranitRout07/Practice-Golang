package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	countLineOrNot := flag.Bool("l", false, "pass true or false for either counting lines or not.")
	flag.Parse()
	val := count(os.Stdin, *countLineOrNot)
	fmt.Println(*countLineOrNot)
	if !*countLineOrNot {
		fmt.Println("Number of words : ", val)
	} else {
		fmt.Println("Number of Lines : ", val)
	}
	
}

func count(r io.Reader, countLines bool) int {
	scanner := bufio.NewScanner(r)
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}
	count_word_or_line := 0
	for scanner.Scan() {
		count_word_or_line++
	}
	return count_word_or_line
}
