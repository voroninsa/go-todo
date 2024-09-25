package postgres

import (
	"errors"
	"fmt"
	"time"

	"github.com/voroninsa/go-todo/utils/common"
	"github.com/voroninsa/go-todo/utils/dto"
)

func (s *postgresStorage) CreateTask(task *dto.Task) (int, error) {
	// Проверка состояния подключения к базе данных
	if err := s.StorageURL.Ping(); err != nil {
		return 0, common.ErrorDatabase(err)
	}

	// Тэги задачи в виде sql запроса
	tagsStr := common.TagsToSqlQueryString(task.Tags)

	// Переменная для хранения возвращенного id
	var taskID int

	// Запрос на создание задачи в базу данных
	err := s.StorageURL.QueryRow(queryCreateTask, task.Description, tagsStr, task.Deadline).Scan(&taskID)
	if err != nil {
		return 0, errors.New(fmt.Sprint("Adding data error: ", err))
	}

	return taskID, nil
}

func (s *postgresStorage) ReadTask(id int) (*dto.Task, error) {
	// Проверка состояния подключения к базе данных
	if err := s.StorageURL.Ping(); err != nil {
		return nil, common.ErrorDatabase(err)
	}

	// Переменная для хранения возвращенного id
	var task dto.Task

	// Запрос на создание задачи в базу данных
	err := s.StorageURL.QueryRow(queryReadTask, id).Scan(&task)
	if err != nil {
		return nil, errors.New(fmt.Sprint("Reading task error: ", err))
	}

	return &task, nil
}

func (s *postgresStorage) UpdateTask(task *dto.Task) error {
	return nil
}

func (s *postgresStorage) DeleteTask(id int) error {
	return nil
}

func (s *postgresStorage) DeleteAllTasks() error {
	return nil
}

func (s *postgresStorage) ReadAllTasks() ([]dto.Task, error) {
	// Проверка состояния подключения к базе данных
	if err := s.StorageURL.Ping(); err != nil {
		return nil, common.ErrorDatabase(err)
	}

	// Переменная для хранения возвращенного id
	var tasks []dto.Task

	// Запрос на создание задачи в базу данных
	err := s.StorageURL.QueryRow(queryReadAllTask).Scan(&tasks)
	if err != nil {
		return nil, errors.New(fmt.Sprint("Reading task error: ", err))
	}

	return tasks, nil
}
func (s *postgresStorage) ReadTasksByTag(tag string) ([]dto.Task, error) {
	return nil, nil
}

func (s *postgresStorage) ReadTasksByDueDate(date time.Time) ([]dto.Task, error) {
	return nil, nil
}
