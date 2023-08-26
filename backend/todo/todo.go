package todo

type Task struct {
	ID       int
	Title    string
	Complete bool
}

type ToDoList struct {
	Tasks []Task
}

func NewToDoList() *ToDoList {
	return &ToDoList{}
}

func (t1 *ToDoList) AddTask(title string) {
	//code

	task := Task{
		ID:    len(t1.Tasks) + 1,
		Title: title,
	}

	t1.Tasks = append(t1.Tasks, task)
}

func (t1 *ToDoList) CompleteTask(id int) {
	//code
	for i, task := range t1.Tasks {
		if task.ID == id {
			t1.Tasks[i].Complete = true
			break
		}
	}
}

func (t1 *ToDoList) ListTasks() []Task {
	//code
	return t1.Tasks

}
