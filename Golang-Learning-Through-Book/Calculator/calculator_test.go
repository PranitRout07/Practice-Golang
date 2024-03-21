package calculator_test

import (
	"calculator"
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	type testCase struct {
		a, b float64
		want float64
	}
	TestCases := []testCase{
		{a: 10, b: 10, want: 20},
		{a: 5, b: 0, want: 5},
		{a: -5, b: 5, want: 0},
	}
	for _, i := range TestCases {

		got := calculator.Add(i.a, i.b)
		if i.want != got {
			t.Errorf("Add(%f,%f) : want %f, got %f", i.a, i.b, i.want, got)
		}
		log.Println("Test Addition Successful")
	}
}

func TestSubtract(t *testing.T) {
	type testCase struct {
		a, b float64
		want float64
	}
	TestCases := []testCase{
		{a: 10, b: 10, want: 0},
		{a: 5, b: 0, want: 5},
		{a: -5, b: 5, want: -10},
	}
	for _, i := range TestCases {

		got := calculator.Subtract(i.a, i.b)
		if i.want != got {
			t.Errorf("Subtract(%f,%f) : want %f, got %f", i.a, i.b, i.want, got)
		}
		log.Println("Test Subtraction Successful")
	}
}

func TestMultiply(t *testing.T) {
	type testCase struct {
		a, b float64
		want float64
	}
	TestCases := []testCase{
		{a: 10, b: 10, want: 100},
		{a: 5, b: 0, want: 0},
		{a: -5, b: 5, want: -25},
	}
	for _, i := range TestCases {

		got := calculator.Multiply(i.a, i.b)

		if i.want != got {
			t.Errorf("Multiply(%f,%f) : want %f, got %f", i.a, i.b, i.want, got)
		}
		log.Println("Test Multiplication Successful")
	}
}
