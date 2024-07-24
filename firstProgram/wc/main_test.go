package main

import (
	"bytes"
	"testing"
)

// TestCountWords tests the count function set to count words
func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("Word1 Word2 Word3 Word4 Word5 Word6 Word7 Word8 Word9\n")
	exp := 9
	res, _ := count(b, false)
	if res != exp {
		t.Errorf("Expected %d, got %d", exp, res)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("Word1 Word2\nWord3 Word4\nWord5 Word6\nWord7 Word8 Word9\n")
	exp := 4
	res, _ := count(b, true)
	if res != exp {
		t.Errorf("Expected %d, got %d", exp, res)
	}
}
