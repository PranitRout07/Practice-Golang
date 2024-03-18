package main

import (
	"bytes"
	"testing"
)

func TestCount(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3")
	expected := 3
	result := count(b, false)

	if result != expected {
		t.Errorf("Expected %d got %d",expected,result)
	}
}
