package todo_test

import (
	"os"
	"testing"
	"todo/todo"
)

func TestAdd(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName)
	if l[0].Task != taskName {
		t.Errorf("Expected %s, got %s", taskName, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)
	if l[0].Task != taskName {
		t.Errorf("Expected %s, got %s", taskName, l[0].Task)
	}
	if l[0].Done {
		t.Errorf("New task should not be completed")
	}
	l.Complete(1)
	if !l[0].Done {
		t.Errorf("New task should be completed")
	}
}

func TestDelete(t *testing.T) {
	l := todo.List{}
	tasks := []string{
		"New task1",
		"New task2",
		"New task3",
	}
	for _, task := range tasks {
		l.Add(task)
	}
	if l[0].Task != tasks[0] {
		t.Errorf("Expected %s, got %s", tasks[0], l[0].Task)
	}
	l.Delete(2)
	if len(l) != 2 {
		t.Errorf("Expected 2 items, got %d", len(l))
	}
	if l[1].Task != tasks[2] {
		t.Errorf("Expected %s, got %s", tasks[2], l[1].Task)
	}
}

// TestSaveGet tests Save and Get methods of the List type
func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	taskName := "New Task"
	l1.Add(taskName)

	if l1[0].Task != taskName {
		t.Errorf("Expected %s, got %s", taskName, l1[0].Task)
	}

	tf, err := os.CreateTemp("", "")

	if err != nil {
		t.Fatalf("Could not create temp file")
	}
	defer os.Remove(tf.Name())

	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("Could not save list")
	}

	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("Could not get list")
	}
	if l1[0].Task != l2[0].Task {
		t.Errorf("Expected %s, got %s", l1[0].Task, l2[0].Task)
	}
}
