package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/voroninsa/go-todo/config"
	"github.com/voroninsa/go-todo/utils/dto"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type postgresStorage struct {
	StorageURL *sql.DB
}

func NewPostgresStorage(c *config.Config, logger *zap.Logger) *postgresStorage {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Db_Host, c.Db_Port, c.Db_User, c.Db_Pass, c.Db_Name)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		logger.Error("Error creating postgres storage: ", zapcore.Field{String: err.Error()})
	}

	return &postgresStorage{
		StorageURL: db,
	}
}

func (p *postgresStorage) Create(dto.StorageRequest) (*dto.StorageResponse, error) {
	return nil, nil
}

func (p *postgresStorage) Read(dto.StorageRequest) (*dto.StorageResponse, error) {
	return nil, nil
}

func (p *postgresStorage) Update(dto.StorageRequest) error {
	return nil
}

func (p *postgresStorage) Delete(dto.StorageRequest) error {
	return nil
}
