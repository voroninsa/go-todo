// handler of path /task/
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/voroninsa/go-todo/internal/storage"
)

type Handlers struct {
	store *storage.TaskStore
}

func (ss *Handlers) PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task storage.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := json.Marshal(ss.store.CreateTask(task.Text, task.Tags, task.Due))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	str := []byte("created new task with id = ")
	w.Write(str)
	w.Write(id)
}

func (ss *Handlers) GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks := ss.store.GetAllTasks()

	ts, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Origin", "*")
	w.Write(ts)
}

func (ss *Handlers) DeleteAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	ss.store.DeleteAllTasks()

	msg := []byte("all tasks deleted")
	w.Write(msg)
}

func (ss *Handlers) GetTaskHandler(w http.ResponseWriter, r *http.Request, id int) {
	task, err := ss.store.GetTask(id)
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

func (ss *Handlers) DeleteTaskHandler(w http.ResponseWriter, r *http.Request, id int) {
	if err := ss.store.DeleteTask(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	str := []byte(fmt.Sprintf("task with id = %d deleted", id))
	w.Write(str)
}

func (ss *Handlers) PatchTaskHandler(w http.ResponseWriter, r *http.Request, id int) {
	var task storage.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := ss.store.PatchTask(id, task.Text, task.Tags, task.Due, task.Completed); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	str := []byte(fmt.Sprintf("patched task with id = %d", id))
	w.Write(str)
}

func (ss *Handlers) OptionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET,POST,DELETE,PATCH")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	w.Header().Add("Access-Control-Max-Age", "86400")
}
