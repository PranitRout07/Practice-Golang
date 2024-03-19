package main

import (
	"bytes"
	"testing"
)

func TestCount(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3")
	expected := 3
	expectedBytes := 15
	result, gotBytes := count(b, false, false)
	if result != expected {
		t.Errorf("Expected result %d got %d", expected, result)
		
	}

	if gotBytes != expectedBytes {
		t.Errorf("Expected bytes %d got %d", expectedBytes, gotBytes)
	}
}

