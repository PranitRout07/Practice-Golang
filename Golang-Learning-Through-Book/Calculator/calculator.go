package calculator

import "errors"

func Add(a, b float64) float64 {
	return a + b
}

// automatic format the go code :- gofmt -w .\calculator.go
// test the code :- go test

func Subtract(a, b float64) float64 {
	return a - b
}
func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Divide By Zero")
	}
	return a/b , nil
}
