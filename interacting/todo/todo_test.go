package todo_test

import (
	"github.com/smissaertj/powerful-cli-apps/interacting/todo"
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName)
	if l[0].Task != taskName {
		t.Errorf("Expected %q got %q instead", taskName, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q got %q instead", taskName, l[0].Task)
	}

	if l[0].Done {
		t.Errorf("New task should not be completed.")
	}

	l.Complete(1)

	if !l[0].Done {
		t.Errorf("New task should be completed.")
	}

}

func TestDelete(t *testing.T) {
	l := todo.List{}
	tasks := []string{"Task 1", "Task 2", "Task 3"}

	for _, v := range tasks {
		l.Add(v)
	}

	// Verify the initial state
	if len(l) != 3 {
		t.Errorf("Expected list length to be 3, got %d instead", len(l))
	}

	if l[0].Task != tasks[0] {
		t.Errorf("Expected %q got %q instead", tasks[0], l[0].Task)
	}

	l.Delete(2)

	if len(l) != 2 {
		t.Errorf("Expected list length to be 2, got %d instead", len(l))
	}

	if l[1].Task != tasks[2] {
		t.Errorf("Expected %q got %q instead", tasks[2], l[1].Task)
	}
}

func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	taskName := "New Task"
	l1.Add(taskName)

	if l1[0].Task != taskName {
		t.Errorf("Expected %q got %q instead", taskName, l1[0].Task)
	}

	tempFile, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}
	defer os.Remove(tempFile.Name())

	if err := l1.Save(tempFile.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}

	if err := l2.Get(tempFile.Name()); err != nil {
		t.Fatalf("Error reading list from file: %s", err)
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf("Expected %q got %q instead", l1[0].Task, l2[0].Task)
	}
}
