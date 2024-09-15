package dto

import "time"

const (
	RequestTargetAll = iota
	RequestTargetTask
	RequestTargetDueDate
	RequestTargetTag
)

// Запрос в хранилище
type StorageRequest struct {
	// Идентификатор запроса
	Target int

	// Текст запроса
	Task Task

	// Искомый тэг
	Tag string

	// Дата выполнения задачи

}

// Задача в хранилище
type Task struct {
	Id        int       `json:"id"`
	Text      string    `json:"text"`
	Tags      []string  `json:"tags"`
	Due       time.Time `json:"due"`
	Completed bool      `json:"completed"`
}

// Ответ хранилища
type StorageResponse struct {
	// Идентификатор задачи
	TaskId int

	// Задача
	Task Task

	// Список задач
	Tasks []Task
}
