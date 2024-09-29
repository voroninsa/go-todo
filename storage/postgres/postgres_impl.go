package postgres

import (
	"fmt"
	"time"

	"github.com/voroninsa/go-todo/utils/common"
	"github.com/voroninsa/go-todo/utils/dto"
)

// Создает задачу в базе данных и возвращает id задачи
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
		return 0, fmt.Errorf("creating task error: %w", err)
	}

	return taskID, nil
}

// Получает задачу из базы данных
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
		return nil, fmt.Errorf("reading task error: %w", err)
	}

	return &task, nil
}

// Обновляет задачу в базе данных
func (s *postgresStorage) UpdateTask(task *dto.Task) error {
	// Проверка состояния подключения к базе данных
	if err := s.StorageURL.Ping(); err != nil {
		return common.ErrorDatabase(err)
	}

	// Тэги задачи в виде sql запроса
	tagsStr := common.TagsToSqlQueryString(task.Tags)

	// Запрос на обновление задачи в базе данных
	_, err := s.StorageURL.Exec(queryUpdateTask, task.Description, tagsStr, task.Deadline, task.Completed, task.Id)
	if err != nil {
		return fmt.Errorf("updating task error: %w", err)
	}

	return nil
}

// Удаляет задачу из базы данных
func (s *postgresStorage) DeleteTask(id int) error {
	// Проверка состояния подключения к базе данных
	if err := s.StorageURL.Ping(); err != nil {
		return common.ErrorDatabase(err)
	}

	// Запрос на удаление задачи из базы данных
	_, err := s.StorageURL.Exec(queryDeleteTask, id)
	if err != nil {
		return fmt.Errorf("deleting task error: %w", err)
	}

	return nil
}

// Удаляет все задачи из базы данных
func (s *postgresStorage) DeleteAllTasks() error {
	// Проверка состояния подключения к базе данных
	if err := s.StorageURL.Ping(); err != nil {
		return common.ErrorDatabase(err)
	}

	// Запрос на удаление задачи из базы данных
	_, err := s.StorageURL.Exec(queryDeleteAllTasks)
	if err != nil {
		return fmt.Errorf("deleting all tasks error: %w", err)
	}

	return nil
}

// Получает все задачи из базы данных
func (s *postgresStorage) ReadAllTasks() ([]dto.Task, error) {
	// Проверка состояния подключения к базе данных
	if err := s.StorageURL.Ping(); err != nil {
		return nil, common.ErrorDatabase(err)
	}

	// Переменная для хранения возвращенных задач
	var tasks []dto.Task

	// Запрос на получение всех задач в базу данных
	err := s.StorageURL.QueryRow(queryReadAllTask).Scan(&tasks)
	if err != nil {
		return nil, fmt.Errorf("reading all tasks error: %w", err)
	}

	return tasks, nil
}

// Получает все задачи с определенным тегом из базы данных
func (s *postgresStorage) ReadTasksByTag(tag string) ([]dto.Task, error) {
	// Проверка состояния подключения к базе данных
	if err := s.StorageURL.Ping(); err != nil {
		return nil, common.ErrorDatabase(err)
	}

	// Переменная для хранения возвращенных задач
	var tasks []dto.Task

	// Запрос на получение всех задач с определенным тегом
	err := s.StorageURL.QueryRow(queryReadTasksByTags, tag).Scan(&tasks)
	if err != nil {
		return nil, fmt.Errorf("reading task by tag error: %w", err)
	}

	return tasks, nil
}

// Получает все задачи с определенным дедлайном из базы данных
func (s *postgresStorage) ReadTasksByDeadline(date time.Time) ([]dto.Task, error) {
	// Проверка состояния подключения к базе данных
	if err := s.StorageURL.Ping(); err != nil {
		return nil, common.ErrorDatabase(err)
	}

	// Переменная для хранения возвращенных задач
	var tasks []dto.Task

	// Запрос на получение всех задач с определенным дедлайном
	err := s.StorageURL.QueryRow(queryReadTasksByDeadline, date).Scan(&tasks)
	if err != nil {
		return nil, fmt.Errorf("reading task by deadline error: %w", err)
	}

	return tasks, nil
}
