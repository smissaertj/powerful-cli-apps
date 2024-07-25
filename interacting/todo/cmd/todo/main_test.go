package main_test

import (
	"fmt"
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
	task := "test task number 1"

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	t.Run("AddNewTak", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-task", task)
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

		expected := fmt.Sprintf("X 1: %s\n", task) // We are expecting a completed single task here!
		if expected != string(out) {
			t.Errorf("expected %q got %q instead\n", expected, string(out))
		}
	})
}
