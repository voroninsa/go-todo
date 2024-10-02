package inmemory

import (
	"sync"

	"github.com/voroninsa/go-todo/utils/dto"
)

type taskStore struct {
	sync.Mutex

	tasks  map[int]dto.Task
	nextId int
}

func NewInMemStorage() *taskStore {
	ts := &taskStore{
		tasks:  make(map[int]dto.Task),
		nextId: 1,
	}

	return ts
}
