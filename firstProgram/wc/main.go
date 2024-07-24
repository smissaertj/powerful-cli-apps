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
	bytes := flag.Bool("b", false, "Count bytes")
	flag.Parse()
	count, err := count(os.Stdin, *lines, *bytes)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(count)
}

func count(reader io.Reader, countLines bool, countBytes bool) (int, error) {

	// Use a scanner to read text from an io.Reader interface
	scanner := bufio.NewScanner(reader)

	// Set the scanner split type to "words", default is split by lines.
	if !countLines && !countBytes {
		scanner.Split(bufio.ScanWords)
	} else if countBytes {
		scanner.Split(bufio.ScanBytes)
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
