package storage

import (
	"github.com/voroninsa/go-todo/config"
	"github.com/voroninsa/go-todo/storage/inmemory"
	"github.com/voroninsa/go-todo/utils/dto"
)

type Backend interface {
	Create(dto.StorageRequest) (*dto.StorageResponse, error)
	Read(dto.StorageRequest) (*dto.StorageResponse, error)
	Update(dto.StorageRequest) error
	Delete(dto.StorageRequest) error
}

func NewStorage(conf config.Config) Backend {
	switch conf.StorageType {
	case "inmemory":
		return inmemory.NewInMemStorage()
	case "postgres":
		return nil
	default:
		return nil
	}

}
