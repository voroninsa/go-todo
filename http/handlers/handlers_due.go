package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/voroninsa/go-todo/storage"
	"github.com/voroninsa/go-todo/utils/common"
	"github.com/voroninsa/go-todo/utils/dto"
	"go.uber.org/zap"
)

type dueHandlersGetter interface {
	GetTasksByDueDateHandler(w http.ResponseWriter, r *http.Request)
}

type dueHandlers struct {
	store  storage.Backend
	logger *zap.Logger
}

func NewDueHandlers(storage storage.Backend, logger *zap.Logger) dueHandlersGetter {
	log := logger.Named("handlers_due")

	return &dueHandlers{
		store:  storage,
		logger: log,
	}
}

func (d *dueHandlers) GetTasksByDueDateHandler(w http.ResponseWriter, r *http.Request) {
	date, err := common.UrlToDate(r.URL.Path)
	if err != nil {
		d.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	storageResp, err := d.store.Read(dto.StorageRequest{
		Target: dto.RequestTargetDueDate,
		Task: dto.Task{
			Deadline: date,
		},
	})
	if err != nil {
		d.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(storageResp.Tasks)
	if err != nil {
		d.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}
