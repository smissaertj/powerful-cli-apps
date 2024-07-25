package main

import (
	"flag"
	"fmt"
	"github.com/smissaertj/powerful-cli-apps/interacting/todo"
	"os"
)

const todoFileName = ".todo.json"

func main() {
	// Parse command line flags
	task := flag.String("task", "", "Task to be included in the ToDo list.")
	list := flag.Bool("list", false, "List all the ToDo tasks.")
	complete := flag.Int("complete", 0, "Item to be completed.")
	flag.Parse()

	// Create a new list
	l := &todo.List{}

	// Load existing tasks if any.
	if err := l.Get(todoFileName); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		for _, item := range *l {
			if !item.Done {
				fmt.Println(item.Task)
			}
		}
	case *task != "":
		l.Add(*task)
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