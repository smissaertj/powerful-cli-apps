package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const (
	header = `<!DOCTYPE html>
	<html>
		<head>
			<meta charset="utf-8" http-equiv="content-type" content="text/html">
			<title>Markdown Preview Tool</title>
		</head>
		<body>`

	footer = `
		</body>
	</html>`
)

func main() {
	fileName := flag.String("file", "", "Mardown file to preview.")
	flag.Parse()

	// The user didn't provide input file, show usage'
	if *fileName == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(*fileName); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}

func run(fileName string) error {
	// Read all the data from the input file and check for errors
	input, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	htmlData := parseContent(input)
	outName := fmt.Sprintf("%s.html", filepath.Base(fileName))
	fmt.Println(outName)

	return saveHTML(outName, htmlData)
}

func parseContent(input []byte) []byte {
	// Parse the markdown file through Blackfriday and Bluemonday
	output := blackfriday.Run(input)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)

	// Create a buffer of bytes to write to file
	var buffer bytes.Buffer

	// Write HTML to bytes buffer, combining the header, body and footer.
	buffer.WriteString(header)
	buffer.Write(body)
	buffer.WriteString(footer)

	return buffer.Bytes()
}

func saveHTML(outName string, data []byte) error {
	// Write the byte to a file
	return os.WriteFile(outName, data, 0644)
}
