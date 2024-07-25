package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type item struct { // We don't export it outside the package
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*l = append(*l, t) // We need to dereference the pointer to access the underlying slice
}

func (l *List) Complete(i int) error {
	ls := *l // We need to dereference the pointer to access the underlying slice

	// Make sure i is a valid slice element
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}

	// Adjust the index for 0-based index
	if ls[i-1].Done {
		ls[i-1].Done = false
	} else {
		ls[i-1].Done = true

	}
	ls[i-1].CompletedAt = time.Now()

	return nil
}

func (l *List) Delete(i int) error {
	ls := *l // We need to dereference the pointer to access the underlying slice

	// Make sure i is a valid slice element
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}

	*l = append(ls[:i-1], ls[i:]...) // (elements to the left of i, elements to the right of i) => i is removed
	return nil
}

func (l *List) Save(fileName string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, js, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (l *List) Get(fileName string) error {
	file, err := os.ReadFile(fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, l)
}

// Returns a formatted list, implements the fmt.Stringer interface
func (l *List) String() string {
	formatted := ""

	for k, t := range *l {
		prefix := " "
		if t.Done {
			prefix = "X"
		}
		formatted += fmt.Sprintf("%s %d: %s\n", prefix, k+1, t.Task)
	}

	return formatted
}
