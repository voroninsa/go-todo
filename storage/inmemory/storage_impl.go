package inmemory

import (
	"fmt"
	"time"

	"github.com/voroninsa/go-todo/utils/dto"
)

const (
	errTaskNotFound = "task with id = %d not found"
)

func (ts *TaskStore) createTask(task dto.Task) int {
	ts.tasks[ts.nextId] = task
	ts.nextId++

	return task.Id
}

func (ts *TaskStore) readTask(id int) (dto.Task, error) {
	task, ok := ts.tasks[id]
	if !ok {
		return dto.Task{}, fmt.Errorf(errTaskNotFound, id)
	}

	return task, nil
}

func (ts *TaskStore) updateTask(task dto.Task) error {
	if ts.tasks[task.Id].Id == 0 {
		return fmt.Errorf(errTaskNotFound, task.Id)
	}

	ts.tasks[task.Id] = task

	return nil
}

func (ts *TaskStore) deleteTask(id int) error {
	if _, ok := ts.tasks[id]; !ok {
		return fmt.Errorf(errTaskNotFound, id)
	}

	delete(ts.tasks, id)

	return nil
}

func (ts *TaskStore) deleteAllTasks() {
	ts.tasks = make(map[int]dto.Task)
}

func (ts *TaskStore) readAllTasks() []dto.Task {
	allTasks := make([]dto.Task, 0, len(ts.tasks))

	for _, task := range ts.tasks {
		allTasks = append(allTasks, task)
	}

	return allTasks
}

func (ts *TaskStore) readTasksByTag(tag string) []dto.Task {
	var tasks []dto.Task

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

func (ts *TaskStore) readTasksByDueDate(date time.Time) []dto.Task {
	var tasks []dto.Task

	for _, task := range ts.tasks {
		if task.Due == date {
			tasks = append(tasks, task)
		}
	}

	return tasks
}
