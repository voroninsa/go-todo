package postgres

import (
	"database/sql"
	"fmt"
	"strings"
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
	if err == sql.ErrNoRows {
		return 0, fmt.Errorf("creating task error: %w", err)
	}

	s.rowCount++

	return taskID, nil
}

// Получает задачу из базы данных
func (s *postgresStorage) ReadTask(id int) (*dto.Task, error) {
	// Проверка состояния подключения к базе данных
	if err := s.StorageURL.Ping(); err != nil {
		return nil, common.ErrorDatabase(err)
	}

	// Запрос на создание задачи в базу данных
	row := s.StorageURL.QueryRow(queryReadTask, id)

	// Перебираем задачи из ответа базы данных
	task, err := taskEnumeration(row)
	if err != nil {
		return nil, err
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

	s.rowCount--

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

	// TODO: Когда внедрю аутентификацию, переделать
	s.rowCount = 0

	return nil
}

// Получает все задачи из базы данных
func (s *postgresStorage) ReadAllTasks() ([]dto.Task, error) {
	// Проверка состояния подключения к базе данных
	if err := s.StorageURL.Ping(); err != nil {
		return nil, common.ErrorDatabase(err)
	}

	// Запрос на получение всех задач в базу данных
	rows, err := s.StorageURL.Query(queryReadAllTask)
	if err != nil {
		return nil, fmt.Errorf("reading all tasks error: %w", err)
	}
	defer rows.Close()

	// Перебираем задачи из ответа базы данных
	tasks, err := taskSliceEnumeration(rows)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// Получает все задачи с определенным тегом из базы данных
func (s *postgresStorage) ReadTasksByTag(tag string) ([]dto.Task, error) {
	// Проверка состояния подключения к базе данных
	if err := s.StorageURL.Ping(); err != nil {
		return nil, common.ErrorDatabase(err)
	}

	// Запрос на получение всех задач с определенным тегом
	rows, err := s.StorageURL.Query(queryReadTasksByTags, tag)
	if err != nil {
		return nil, fmt.Errorf("reading task by tag error: %w", err)
	}
	defer rows.Close()

	// Перебираем задачи из ответа базы данных
	tasks, err := taskSliceEnumeration(rows)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// Получает все задачи с определенным дедлайном из базы данных
func (s *postgresStorage) ReadTasksByDeadline(date time.Time) ([]dto.Task, error) {
	// Проверка состояния подключения к базе данных
	if err := s.StorageURL.Ping(); err != nil {
		return nil, common.ErrorDatabase(err)
	}

	// Запрос на получение всех задач с определенным дедлайном
	rows, err := s.StorageURL.Query(queryReadTasksByDeadline, date)
	if err != nil {
		return nil, fmt.Errorf("reading task by deadline error: %w", err)
	}

	// Перебираем задачи из ответа базы данных
	tasks, err := taskSliceEnumeration(rows)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func taskSliceEnumeration(rows *sql.Rows) ([]dto.Task, error) {
	// Переменная для хранения возвращенных задач
	var tasks []dto.Task

	// Перебираем задачи из ответа базы данных
	for rows.Next() {
		var task dto.Task
		var bytesTags []uint8

		err := rows.Scan(
			&task.Id,
			&task.Description,
			&bytesTags,
			&task.Deadline,
			&task.Created_at,
			&task.Updated_at,
			&task.Completed,
		)
		if err != nil {
			return nil, fmt.Errorf("reading task by tag in tasks rows error: %w", err)
		}

		// Преобразуем битовый слайс []uint8 из бд в строку
		tagsStr := string(bytesTags)

		// Удаляем фигурные скобки и разделяем строку по запятым
		tagsStr = strings.Trim(tagsStr, "{}")
		tagsList := strings.Split(tagsStr, ",")

		task.Tags = tagsList

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func taskEnumeration(rows *sql.Row) (dto.Task, error) {
	// Переменная для хранения возвращенных задач
	var task dto.Task

	// Перебираем задачу из ответа базы данных
	var bytesTags []uint8

	err := rows.Scan(
		&task.Id,
		&task.Description,
		&bytesTags,
		&task.Deadline,
		&task.Created_at,
		&task.Updated_at,
		&task.Completed,
	)
	if err != nil {
		return dto.Task{}, fmt.Errorf("reading task by tag in tasks rows error: %w", err)
	}

	// Преобразуем битовый слайс []uint8 из бд в строку
	tagsStr := string(bytesTags)

	// Удаляем фигурные скобки и разделяем строку по запятым
	tagsStr = strings.Trim(tagsStr, "{}")
	tagsList := strings.Split(tagsStr, ",")

	task.Tags = tagsList

	return task, nil
}
