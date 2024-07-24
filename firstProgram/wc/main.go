package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(count(os.Stdin))
}

func count(reader io.Reader) int {

	// Use a scanner to read text from an io.Reader interface
	scanner := bufio.NewScanner(reader)

	// Set the scanner split type to "words", default is split by lines.
	scanner.Split(bufio.ScanWords)

	// Define a counter
	wc := 0

	// Increment the counter for every word scanned
	for scanner.Scan() {
		wc++
	}

	// return the total
	return wc
}
