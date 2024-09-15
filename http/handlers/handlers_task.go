// handler of path /task/
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/voroninsa/go-todo/storage"
	"github.com/voroninsa/go-todo/utils/dto"
)

type taskHandlersGetter interface {
	PostTaskHandler(w http.ResponseWriter, r *http.Request)
	GetAllTasksHandler(w http.ResponseWriter, r *http.Request)
	DeleteAllTasksHandler(w http.ResponseWriter, r *http.Request)
	GetTaskHandler(w http.ResponseWriter, r *http.Request, id int)
	DeleteTaskHandler(w http.ResponseWriter, r *http.Request, id int)
	PatchTaskHandler(w http.ResponseWriter, r *http.Request, id int)
	OptionsHandler(w http.ResponseWriter, r *http.Request)
}

type taskHandlers struct {
	store storage.Backend
}

func NewTaskHandlers(storage storage.Backend) taskHandlersGetter {
	return &taskHandlers{
		store: storage,
	}
}

func (t *taskHandlers) PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task dto.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Создание задачи
	storageResp, err := t.store.Create(dto.StorageRequest{
		Target: dto.RequestTargetAll,
		Task:   task,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := json.Marshal(storageResp.TaskId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := []byte("created new task with id = ")
	w.Write(resp)
	w.Write(id)
}

func (t *taskHandlers) GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	// Получение всех задач
	storageResp, _ := t.store.Read(dto.StorageRequest{
		Target: dto.RequestTargetAll,
	})

	ts, err := json.Marshal(storageResp.Tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(ts)
}

func (t *taskHandlers) DeleteAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	// Удаление всех задач
	t.store.Delete(dto.StorageRequest{
		Target: dto.RequestTargetAll,
	})

	msg := []byte("all tasks deleted")
	w.Write(msg)
}

func (t *taskHandlers) GetTaskHandler(w http.ResponseWriter, r *http.Request, id int) {
	// Получение определенной задачи
	storageResp, err := t.store.Read(dto.StorageRequest{
		Target: dto.RequestTargetTask,
		Task: dto.Task{
			Id: id,
		},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(storageResp.Task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}

func (t *taskHandlers) DeleteTaskHandler(w http.ResponseWriter, r *http.Request, id int) {
	// Удаление определенной задачи
	err := t.store.Delete(dto.StorageRequest{
		Task: dto.Task{
			Id: id,
		},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	str := []byte(fmt.Sprintf("task with id = %d deleted", id))
	w.Write(str)
}

func (t *taskHandlers) PatchTaskHandler(w http.ResponseWriter, r *http.Request, id int) {
	var task dto.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	task.Id = id

	// Обновление определенной задачи
	err := t.store.Update(dto.StorageRequest{
		Target: dto.RequestTargetTask,
		Task:   task,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	str := []byte(fmt.Sprintf("updated task with id = %d", id))
	w.Write(str)
}

func (t *taskHandlers) OptionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET,POST,DELETE,PATCH")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	w.Header().Add("Access-Control-Max-Age", "86400")
}
