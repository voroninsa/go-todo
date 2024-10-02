package storage

import (
	"errors"

	"github.com/voroninsa/go-todo/config"
	"github.com/voroninsa/go-todo/storage/inmemory"
	"github.com/voroninsa/go-todo/storage/postgres"
	"github.com/voroninsa/go-todo/utils/dto"
	"go.uber.org/zap"
)

var (
	errIncorrectRequest = errors.New("incorrect request target")
)

type backendStorage struct {
	storage dto.StorageMoves
}

type Backend interface {
	Create(dto.StorageRequest) (*dto.StorageResponse, error)
	Read(dto.StorageRequest) (*dto.StorageResponse, error)
	Update(dto.StorageRequest) error
	Delete(dto.StorageRequest) error
}

func BackendFactory(conf *config.Config, logger *zap.Logger) Backend {
	switch conf.StorageType {
	case "inmemory":
		return &backendStorage{
			storage: inmemory.NewInMemStorage(),
		}
	case "postgres":
		return &backendStorage{
			storage: postgres.NewPostgresStorage(conf, logger),
		}
	default:
		return nil
	}
}

func (ts *backendStorage) Create(req dto.StorageRequest) (*dto.StorageResponse, error) {
	switch req.Target {
	// Создание задачи
	case dto.RequestTargetAll:
		taskId, err := ts.storage.CreateTask(&req.Task)
		if err != nil {
			return nil, err
		}

		return &dto.StorageResponse{
			TaskId: taskId,
		}, nil

	default:
		return nil, errIncorrectRequest
	}
}

func (ts *backendStorage) Read(req dto.StorageRequest) (*dto.StorageResponse, error) {
	switch req.Target {
	// Все задачи
	case dto.RequestTargetAll:
		tasks, err := ts.storage.ReadAllTasks()
		if err != nil {
			return nil, err
		}

		return &dto.StorageResponse{
			Tasks: tasks,
		}, nil

	// Задача по id
	case dto.RequestTargetTask:
		task, err := ts.storage.ReadTask(req.Task.Id)
		if err != nil {
			return nil, err
		}

		return &dto.StorageResponse{
			Task: *task,
		}, nil

	// Задачи по тегу
	case dto.RequestTargetTag:
		tasks, err := ts.storage.ReadTasksByTag(req.Tag)
		if err != nil {
			return nil, err
		}

		return &dto.StorageResponse{
			Tasks: tasks,
		}, nil

	// Задачи по дате
	case dto.RequestTargetDueDate:
		tasks, err := ts.storage.ReadTasksByDeadline(req.Task.Deadline)
		if err != nil {
			return nil, err
		}

		return &dto.StorageResponse{
			Tasks: tasks,
		}, nil

	default:
		return nil, errIncorrectRequest
	}
}

func (ts *backendStorage) Update(req dto.StorageRequest) error {
	switch req.Target {
	// Обновление задачи
	case dto.RequestTargetTask:
		err := ts.storage.UpdateTask(&req.Task)
		if err != nil {
			return err
		}

		return nil

	default:
		return errIncorrectRequest
	}
}

func (ts *backendStorage) Delete(req dto.StorageRequest) error {
	switch req.Target {
	// Удаление определенной задачи
	case dto.RequestTargetTask:
		err := ts.storage.DeleteTask(req.Task.Id)
		if err != nil {
			return err
		}

		return nil

	// Удаление всех задач
	case dto.RequestTargetAll:
		err := ts.storage.DeleteAllTasks()
		if err != nil {
			return err
		}

		return nil

	default:
		return errIncorrectRequest
	}
}
