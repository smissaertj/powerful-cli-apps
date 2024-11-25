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

const header = `<!DOCTYPE html>
<html>
<head>
    <meta http-equiv="content-type" content="text/html; charset=utf-8">
    <title>Markdown Preview Tool</title>
</head>
<body>`

const footer = `</body>
</html>`

func main() {
	fileName := flag.String("file", "", "Markdown file to preview.")
	flag.Parse()

	// The user didn't provide input file, show usage'
	if *fileName == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(*fileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
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
	sanitizedBody := bluemonday.UGCPolicy().SanitizeBytes(output)

	// Replace newlines in code blocks with trailing space
	sanitizedBody = bytes.ReplaceAll(sanitizedBody, []byte("\n</code>"), []byte(" </code>"))

	// Create a buffer of bytes to write to file
	var buffer bytes.Buffer

	// Write HTML to bytes buffer, combining the header, body and footer.
	buffer.WriteString(header)
	buffer.Write(sanitizedBody)
	buffer.WriteString(footer)

	return formatHTML(buffer.Bytes())
}

func formatHTML(input []byte) []byte {
	formatted := bytes.ReplaceAll(input, []byte(">\n<"), []byte("><"))
	return bytes.TrimSpace(formatted)
}

func saveHTML(outName string, data []byte) error {
	// Write the byte to a file
	return os.WriteFile(outName, data, 0644)
}
