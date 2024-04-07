package abcd

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	type testcase struct {
		name     string
		expected string
	}

	testcases := []testcase{
		{"John","Hello John Good morning!!!\n"},
		{"James","Hello James Good morning!!!\n"},
	}
	for _, item := range testcases {
		got := Greet(item.name)
		if got != item.expected {
			t.Error("Got", got, "Expected", item.expected)
		}

	}
}

func ExampleGreet() {
	fmt.Println(Greet("Zade"))
	//Output:
	//Hello Zade Good morning!!!
}


func BenchmarkGreet(b *testing.B){
	for i:= 0 ; i<b.N ; i++ {
		Greet("John")
	}
}
