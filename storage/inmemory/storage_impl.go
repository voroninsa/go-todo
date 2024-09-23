package inmemory

import (
	"fmt"
	"time"

	"github.com/voroninsa/go-todo/utils/dto"
)

const (
	errTaskNotFound = "task with id = %d not found"
)

func (ts *taskStore) createTask(task dto.Task) int {
	task.Id = ts.nextId
	ts.tasks[ts.nextId] = task
	ts.nextId++

	return task.Id
}

func (ts *taskStore) readTask(id int) (dto.Task, error) {
	task, ok := ts.tasks[id]
	if !ok {
		return dto.Task{}, fmt.Errorf(errTaskNotFound, id)
	}

	return task, nil
}

func (ts *taskStore) updateTask(task dto.Task) error {
	if ts.tasks[task.Id].Id == 0 {
		return fmt.Errorf(errTaskNotFound, task.Id)
	}

	ts.tasks[task.Id] = task

	return nil
}

func (ts *taskStore) deleteTask(id int) error {
	if _, ok := ts.tasks[id]; !ok {
		return fmt.Errorf(errTaskNotFound, id)
	}

	delete(ts.tasks, id)

	return nil
}

func (ts *taskStore) deleteAllTasks() {
	ts.tasks = make(map[int]dto.Task)
}

func (ts *taskStore) readAllTasks() []dto.Task {
	allTasks := make([]dto.Task, 0, len(ts.tasks))

	for _, task := range ts.tasks {
		allTasks = append(allTasks, task)
	}

	return allTasks
}

func (ts *taskStore) readTasksByTag(tag string) []dto.Task {
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

	return tasks
}

func (ts *taskStore) readTasksByDueDate(date time.Time) []dto.Task {
	tasks := make([]dto.Task, 0)

	for _, task := range ts.tasks {
		if task.Due == date {
			tasks = append(tasks, task)
		}
	}

	return tasks
}
