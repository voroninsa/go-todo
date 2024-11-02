package inmemory

import (
	"sync"

	"github.com/voroninsa/go-todo/utils/dto"
)

type store struct {
	sync.Mutex

	users map[int]taskStore
}

type taskStore struct {
	sync.Mutex

	tasks  map[int]dto.Task
	nextId int
}

func NewInMemStorage() *taskStore {
	ts := &taskStore{
		tasks:  make(map[int]dto.Task, 20),
		nextId: 1,
	}

	return ts
}
