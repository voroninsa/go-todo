// handler of path /task/
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/voroninsa/go-todo/storage"
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
	store *storage.TaskStore
}

func NewTaskHandlers(storage *storage.TaskStore) taskHandlersGetter {
	return &taskHandlers{
		store: storage,
	}
}

func (t *taskHandlers) PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task storage.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := json.Marshal(t.store.CreateTask(task.Text, task.Tags, task.Due))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	str := []byte("created new task with id = ")
	w.Write(str)
	w.Write(id)
}

func (t *taskHandlers) GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks := t.store.GetAllTasks()

	ts, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Origin", "*")
	w.Write(ts)
}

func (t *taskHandlers) DeleteAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	t.store.DeleteAllTasks()

	msg := []byte("all tasks deleted")
	w.Write(msg)
}

func (t *taskHandlers) GetTaskHandler(w http.ResponseWriter, r *http.Request, id int) {
	task, err := t.store.GetTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ts, err := json.Marshal(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(ts)
}

func (t *taskHandlers) DeleteTaskHandler(w http.ResponseWriter, r *http.Request, id int) {
	if err := t.store.DeleteTask(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	str := []byte(fmt.Sprintf("task with id = %d deleted", id))
	w.Write(str)
}

func (t *taskHandlers) PatchTaskHandler(w http.ResponseWriter, r *http.Request, id int) {
	var task storage.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := t.store.PatchTask(id, task.Text, task.Tags, task.Due, task.Completed); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	str := []byte(fmt.Sprintf("patched task with id = %d", id))
	w.Write(str)
}

func (t *taskHandlers) OptionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET,POST,DELETE,PATCH")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	w.Header().Add("Access-Control-Max-Age", "86400")
}
