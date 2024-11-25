package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
	"io"
	"os"
	"os/exec"
	"runtime"
	"time"
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
	skipPreview := flag.Bool("s", false, "Skip auto-preview")
	flag.Parse()

	// The user didn't provide input file, show usage'
	if *fileName == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(*fileName, os.Stdout, *skipPreview); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(fileName string, out io.Writer, skipPreview bool) error {
	// Read all the data from the input file and check for errors
	input, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	htmlData := parseContent(input)
	// Create temporary file and check for errors
	temp, err := os.CreateTemp("", "mdp*.html")
	if err != nil {
		return err
	}
	if err := temp.Close(); err != nil {
		return err
	}

	outName := temp.Name()
	fmt.Fprintln(out, outName)

	if err := saveHTML(outName, htmlData); err != nil {
		return err
	}

	if skipPreview {
		return nil
	}

	defer os.Remove(outName)

	return preview(outName)
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

func preview(fname string) error {
	cName := ""
	cParams := []string{}

	// Define executable based on OS
	switch runtime.GOOS {
	case "linux":
		cName = "xdg-open"
	case "windows":
		cName = "cmd.exe"
		cParams = []string{"/c", "start"}
	case "darwin":
		cName = "open"
	default:
		return fmt.Errorf("unsupported platform")
	}

	// Append filename to parameters slice
	cParams = append(cParams, fname)

	// Locate executable in PATH
	cPath, err := exec.LookPath(cName)
	if err != nil {
		return err
	}

	// Open the file using default program
	err = exec.Command(cPath, cParams...).Run()

	// Give the browser some time to open the file before it gets deleted
	time.Sleep(2 * time.Second)
	return err
}
