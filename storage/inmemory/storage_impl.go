package inmemory

import (
	"fmt"
	"time"

	"github.com/voroninsa/go-todo/utils/dto"
)

const (
	errTaskNotFound = "task with id = %d not found"
)

func (ts *taskStore) CreateTask(task *dto.Task) (int, error) {
	ts.Lock()
	defer ts.Unlock()

	task.Id = ts.nextId
	ts.tasks[ts.nextId] = *task
	ts.nextId++

	return task.Id, nil
}

func (ts *taskStore) ReadTask(id int) (*dto.Task, error) {
	ts.Lock()
	defer ts.Unlock()

	task, ok := ts.tasks[id]
	if !ok {
		return &dto.Task{}, fmt.Errorf(errTaskNotFound, id)
	}

	return &task, nil
}

func (ts *taskStore) UpdateTask(task *dto.Task) error {
	ts.Lock()
	defer ts.Unlock()

	if ts.tasks[task.Id].Id == 0 {
		return fmt.Errorf(errTaskNotFound, task.Id)
	}

	ts.tasks[task.Id] = *task

	return nil
}

func (ts *taskStore) DeleteTask(id int) error {
	ts.Lock()
	defer ts.Unlock()

	if _, ok := ts.tasks[id]; !ok {
		return fmt.Errorf(errTaskNotFound, id)
	}

	delete(ts.tasks, id)

	return nil
}

func (ts *taskStore) DeleteAllTasks() error {
	ts.Lock()
	defer ts.Unlock()

	ts.tasks = make(map[int]dto.Task)

	return nil
}

func (ts *taskStore) ReadAllTasks() ([]dto.Task, error) {
	ts.Lock()
	defer ts.Unlock()

	allTasks := make([]dto.Task, 0, len(ts.tasks))

	for _, task := range ts.tasks {
		allTasks = append(allTasks, task)
	}

	return allTasks, nil
}

func (ts *taskStore) ReadTasksByTag(tag string) ([]dto.Task, error) {
	ts.Lock()
	defer ts.Unlock()

	tasks := make([]dto.Task, 0)

taskloop:
	for _, task := range ts.tasks {
		for _, taskTag := range task.Tags {
			if taskTag == tag {
				tasks = append(tasks, task)
				continue taskloop
			}
		}
	}

	return tasks, nil
}

func (ts *taskStore) ReadTasksByDueDate(date time.Time) ([]dto.Task, error) {
	ts.Lock()
	defer ts.Unlock()

	tasks := make([]dto.Task, 0)

	for _, task := range ts.tasks {
		if task.Deadline == date {
			tasks = append(tasks, task)
		}
	}

	return tasks, nil
}
