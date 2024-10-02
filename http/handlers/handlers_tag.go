package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/voroninsa/go-todo/storage"
	"github.com/voroninsa/go-todo/utils/dto"
	"go.uber.org/zap"
)

type tagHandlersGetter interface {
	GetTasksByTagHandler(w http.ResponseWriter, r *http.Request)
}

type tagHandlers struct {
	store  storage.Backend
	logger *zap.Logger
}

func NewTagHandlers(storage storage.Backend, logger *zap.Logger) tagHandlersGetter {
	log := logger.Named("handlers_tag")

	return &tagHandlers{
		store:  storage,
		logger: log,
	}
}

func (t *tagHandlers) GetTasksByTagHandler(w http.ResponseWriter, r *http.Request) {
	tag := strings.Split(r.URL.Path, "/tag/")[1]

	// Получение задач по тегу
	storageResp, err := t.store.Read(dto.StorageRequest{
		Target: dto.RequestTargetTag,
		Tag:    tag,
	})
	if err != nil {
		t.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(storageResp.Tasks)
	if err != nil {
		t.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}
