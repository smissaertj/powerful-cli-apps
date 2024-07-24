package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	// Define a boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "Count lines")
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines))
}

func count(reader io.Reader, countLines bool) (int, error) {

	// Use a scanner to read text from an io.Reader interface
	scanner := bufio.NewScanner(reader)

	// Set the scanner split type to "words", default is split by lines.
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	// Define a counter
	wc := 0

	// Increment the counter for every word scanned
	for scanner.Scan() {
		wc++
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("failed reading standard input: %w", err)
	}

	// return the total
	return wc, nil
}
