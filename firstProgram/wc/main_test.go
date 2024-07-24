package main

import (
	"bytes"
	"testing"
)

// TestCountWords tests the count function set to count words
func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("Word1 Word2 Word3 Word4 Word5 Word6 Word7 Word8 Word9\n")
	exp := 9
	res, _ := count(b, false, false)
	if res != exp {
		t.Errorf("Expected %d, got %d", exp, res)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("Word1 Word2\nWord3 Word4\nWord5 Word6\nWord7 Word8 Word9\n")
	exp := 4
	res, _ := count(b, true, false)
	if res != exp {
		t.Errorf("Expected %d, got %d", exp, res)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBuffer([]byte("Hello World"))
	exp := b.String()
	res, _ := count(b, false, true)
	if res != len(exp) { // len returns the count of bytes in the string representation of the buffer.
		t.Errorf("Expected %d, got %d", len(exp), res)
	}
}
