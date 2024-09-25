package dto

import (
	"time"
)

type StorageMoves interface {
	CreateTask(task *Task) (int, error)
	ReadTask(id int) (*Task, error)
	UpdateTask(task *Task) error
	DeleteTask(id int) error
	DeleteAllTasks() error
	ReadAllTasks() ([]Task, error)
	ReadTasksByTag(tag string) ([]Task, error)
	ReadTasksByDueDate(date time.Time) ([]Task, error)
}
