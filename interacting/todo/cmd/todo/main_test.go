package main_test

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	binName  = "todo"
	fileName = ".todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)
	if err := build.Run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "cannot build tool %s: %s", binName, err)
	}

	fmt.Println("Running tests...")
	result := m.Run()

	fmt.Println("Cleaning up...")
	if err := os.Remove(binName); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "cannot remove tool %s: %s", binName, err)
	}

	if err := os.Remove(fileName); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "cannot remove file %s: %s", fileName, err)
	}
	fmt.Println("Clean up successful!")
	os.Exit(result)
}

func TestTodoCli(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	cmdPath := filepath.Join(dir, binName)

	task := "test task number 1"
	t.Run("AddNewTaskFromArguments", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add", task)
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	task2 := "test task number 2"
	t.Run("AddNewTaskFromSTDIN", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add")
		cmdStdin, err := cmd.StdinPipe()
		if err != nil {
			t.Fatal(err)
		}

		if _, err := io.WriteString(cmdStdin, task2); err != nil {
			t.Fatal(err)
		}

		if err := cmdStdin.Close(); err != nil {
			t.Fatal(err)
		}

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("CompleteTask", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-complete", "1")
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ListTask", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		expected := fmt.Sprintf("X 1: %s\n  2: %s\n", task, task2) // We are expecting only one completed task here
		if expected != string(out) {
			t.Errorf("expected %q got %q instead\n", expected, string(out))
		}
	})
}
