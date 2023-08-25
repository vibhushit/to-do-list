package todo

import "testing"

func TestToDoList_AddTask(t *testing.T) {
	//code
	t1 := NewToDoList()
	t1.AddTask("Buy Groceries")
	t1.AddTask("Write Code")

	if len(t1.Tasks) != 2 {
		t.Errorf("Expected %d tasks, but got %d", 2, len(t1.Tasks))
	}
}

func TestToDoList_CompleteTask(t *testing.T) {
	//code
	t1 := NewToDoList()
	t1.AddTask("Buy Groceries")
	t1.AddTask("Write Code")
	t1.CompleteTask(1)

	if !t1.Tasks[0].Complete {
		t.Errorf("Expected task to be marked as complete")
	}

}

func TestToDoList_ListTasks(t *testing.T) {
	//code
	t1 := NewToDoList()
	t1.AddTask("Buy Groceries")
	t1.AddTask("Write Code")

	tasks := t1.ListTasks()

	if len(tasks) != 2 {
		t.Errorf("Expected %d tasks, but got %d", 2, len(tasks))
	}

}
