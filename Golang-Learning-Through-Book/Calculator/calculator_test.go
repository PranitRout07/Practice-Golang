package calculator_test

import (
	"calculator"
	"log"
	"testing"
	"math"
)

func closeEnough(a, b, gap float64) bool {
	return math.Abs(a-b) <= gap
}

func TestAdd(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
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

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	TestCases := []testCase{
		{a: 10, b: 10, want: 1},
		{a: 5, b: 1, want: 5},
		{a: -5, b: 5, want: -1},
		{a: 1 , b: 3 , want: 0.333333},
		// {a: 1 , b : 0 , want : 0},
	}
	for _, i := range TestCases {

		got, err := calculator.Divide(i.a, i.b)

		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}

		if i.want != got && !closeEnough(i.want,got,0.001) {
			t.Errorf("Divide(%f,%f) : want %f, got %f", i.a, i.b, i.want, got)
		}
		log.Println("Test Divide Successful")
	}
}
