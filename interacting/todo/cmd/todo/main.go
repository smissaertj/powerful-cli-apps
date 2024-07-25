package main

import (
	"fmt"
	"github.com/smissaertj/powerful-cli-apps/interacting/todo"
	"os"
	"strings"
)

const todoFileName = ".todo.json"

func main() {
	l := &todo.List{}

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch len(os.Args) {
	case 1: // application is called without arguments
		for _, item := range *l {
			fmt.Println(item.Task)
		}

	default: // concatenates all arguments with spaces and add them as a single task to the list of tasks
		item := strings.Join(os.Args[1:], " ")
		l.Add(item)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
