package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/smissaertj/powerful-cli-apps/interacting/todo"
	"io"
	"os"
	"strings"
)

var todoFileName = ".todo.json"

func main() {
	// Parse command line flags
	add := flag.Bool("add", false, "Task to be included in the ToDo list.")
	list := flag.Bool("list", false, "List all the ToDo tasks.")
	complete := flag.Int("complete", 0, "Item to be completed.")
	flag.Parse()

	// Check if the user defined an env var for a custom todoFileName
	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}

	// Create a new list
	l := &todo.List{}

	// Load existing tasks if any.
	if err := l.Get(todoFileName); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		fmt.Print(l)
	case *add:
		task, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		l.Add(task)

		if err := l.Save(todoFileName); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *complete > 0:
		if err := l.Complete(*complete); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		_, _ = fmt.Fprintln(os.Stderr, "Invalid flag provided")
		os.Exit(1)
	}
}

func getTask(r io.Reader, args ...string) (string, error) { // variadic function accepting a variable number of arguments
	// Check if arguments were provided
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	// Read the task from the reader
	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}

	// Ensure that we don't have a blank task
	if len(s.Text()) == 0 {
		return "", fmt.Errorf("task cannot be blank")
	}

	return s.Text(), nil
}
