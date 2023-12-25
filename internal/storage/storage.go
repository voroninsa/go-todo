package storage

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	Id        int       `json:"id"`
	Text      string    `json:"text"`
	Tags      []string  `json:"tags"`
	Due       time.Time `json:"due"`
	Completed bool      `json:"completed"`
}

type TaskStore struct {
	sync.Mutex

	tasks  map[int]Task
	nextId int
}

func New() *TaskStore {
	ts := &TaskStore{}
	ts.tasks = make(map[int]Task)
	ts.nextId = 1

	return ts
}

func (ts *TaskStore) CreateTask(text string, tags []string, due time.Time) int {
	ts.Lock()
	defer ts.Unlock()

	task := Task{
		Id:   ts.nextId,
		Text: text,
		Due:  due,
	}
	task.Tags = make([]string, len(tags))
	copy(task.Tags, tags)

	ts.tasks[ts.nextId] = task
	ts.nextId++

	return task.Id
}

func (ts *TaskStore) GetTask(id int) (Task, error) {
	ts.Lock()
	defer ts.Unlock()

	task, ok := ts.tasks[id]
	if ok {
		return task, nil
	} else {
		return Task{}, fmt.Errorf("task with id = %d not found", id)
	}
}

func (ts *TaskStore) PatchTask(id int, text string, tags []string, due time.Time, completed bool) error {
	ts.Lock()
	defer ts.Unlock()

	if ts.tasks[id].Id == 0 {
		return fmt.Errorf("Task with id = %d not found", id)
	}

	task := Task{
		Id:        id,
		Text:      text,
		Due:       due,
		Completed: completed,
	}
	task.Tags = make([]string, len(tags))
	copy(task.Tags, tags)

	ts.tasks[id] = task

	return nil
}

func (ts *TaskStore) DeleteTask(id int) error {
	ts.Lock()
	defer ts.Unlock()

	if _, ok := ts.tasks[id]; !ok {
		return fmt.Errorf("task with id = %d does not exist", id)
	}

	delete(ts.tasks, id)

	return nil
}

func (ts *TaskStore) DeleteAllTasks() error {
	ts.Lock()
	defer ts.Unlock()

	ts.tasks = make(map[int]Task)
	return nil
}

func (ts *TaskStore) GetAllTasks() []Task {
	ts.Lock()
	defer ts.Unlock()

	// var allTasks []Task
	allTasks := make([]Task, 0, len(ts.tasks))

	for _, task := range ts.tasks {
		allTasks = append(allTasks, task)
	}

	return allTasks
}

func (ts *TaskStore) GetTasksByTag(tag string) []Task {
	ts.Lock()
	defer ts.Unlock()

	var tasks []Task

taskloop:
	for _, task := range ts.tasks {
		for _, taskTag := range task.Tags {
			if taskTag == tag {
				tasks = append(tasks, task)
				continue taskloop
			}
		}
	}

	return tasks
}

func (ts *TaskStore) GetTasksByDueDate(year int, month time.Month, day int) []Task {
	ts.Lock()
	defer ts.Unlock()

	var tasks []Task

	for _, task := range ts.tasks {
		y, m, d := task.Due.Date()
		if y == year && m == month && d == day {
			tasks = append(tasks, task)
		}
	}

	return tasks
}
