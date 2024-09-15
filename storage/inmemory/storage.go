package inmemory

import (
	"errors"
	"sync"

	"github.com/voroninsa/go-todo/utils/dto"
)

var (
	errIncorrectRequest = errors.New("incorrect request target")
)

type TaskStore struct {
	sync.Mutex

	tasks  map[int]dto.Task
	nextId int
}

func NewInMemStorage() *TaskStore {
	ts := &TaskStore{
		tasks:  make(map[int]dto.Task),
		nextId: 1,
	}

	return ts
}

func (ts *TaskStore) Create(req dto.StorageRequest) (*dto.StorageResponse, error) {
	ts.Lock()
	defer ts.Unlock()

	switch req.Target {
	// Создание задачи
	case dto.RequestTargetAll:
		taskId := ts.createTask(req.Task)

		return &dto.StorageResponse{
			TaskId: taskId,
		}, nil

	default:
		return nil, errIncorrectRequest
	}
}

func (ts *TaskStore) Read(req dto.StorageRequest) (*dto.StorageResponse, error) {
	ts.Lock()
	defer ts.Unlock()

	switch req.Target {
	// Все задачи
	case dto.RequestTargetAll:
		tasks := ts.readAllTasks()

		return &dto.StorageResponse{
			Tasks: tasks,
		}, nil

	// Задача по id
	case dto.RequestTargetTask:
		task, err := ts.readTask(req.Task.Id)
		if err != nil {
			return nil, err
		}

		return &dto.StorageResponse{
			Task: task,
		}, nil

	// Задачи по тегу
	case dto.RequestTargetTag:
		tasks := ts.readTasksByTag(req.Tag)

		return &dto.StorageResponse{
			Tasks: tasks,
		}, nil

	// Задачи по дате
	case dto.RequestTargetDueDate:
		tasks := ts.readTasksByDueDate(req.Task.Due)

		return &dto.StorageResponse{
			Tasks: tasks,
		}, nil

	default:
		return nil, errIncorrectRequest
	}
}

func (ts *TaskStore) Update(req dto.StorageRequest) error {
	ts.Lock()
	defer ts.Unlock()

	switch req.Target {
	// Обновление задачи
	case dto.RequestTargetTask:
		err := ts.updateTask(req.Task)
		if err != nil {
			return err
		}

		return nil

	default:
		return errIncorrectRequest
	}
}

func (ts *TaskStore) Delete(req dto.StorageRequest) error {
	ts.Lock()
	defer ts.Unlock()

	switch req.Target {
	// Удаление определенной задачи
	case dto.RequestTargetTask:
		err := ts.deleteTask(req.Task.Id)
		if err != nil {
			return err
		}

		return nil

	// Удаление всех задач
	case dto.RequestTargetAll:
		ts.deleteAllTasks()

		return nil

	default:
		return errIncorrectRequest
	}
}
