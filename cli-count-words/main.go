package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	countLineOrNot := flag.Bool("l", false, "Pass true or false for either counting lines or not.")
	checkBytes := flag.Bool("b", false, "Check bytes")
	flag.Parse()
	val, bytesCount := count(os.Stdin, *countLineOrNot, *checkBytes)
	fmt.Println(*countLineOrNot)
	if !*countLineOrNot {
		fmt.Println("Number of words : ", val)
	} else {
		fmt.Println("Number of Lines : ", val)
	}
	if *checkBytes {
		fmt.Println("Number of bytes : ", bytesCount)
	}

}

func count(r io.Reader, countLines bool, checkBytes bool) (int, int) {
	scanner := bufio.NewScanner(r)

	if !countLines {
		scanner.Split(bufio.ScanWords)
	}
	count_word_or_line := 0
	bytesNO := 0
	for scanner.Scan() {
		count_word_or_line++
		bytesNO+=len(scanner.Bytes())
	}
	return count_word_or_line, bytesNO
}
