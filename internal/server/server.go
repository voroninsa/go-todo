package server

import (
	"encoding/json"
	"fmt"
	"go-todo/internal/common"
	"go-todo/internal/storage"
	"net/http"
	"strconv"
	"strings"
)

type serverStore struct {
	store *storage.TaskStore
}

func NewStorage() *serverStore {
	store := storage.New()
	return &serverStore{store: store}
}

func (ss *serverStore) TaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/task/" {
		switch r.Method {
		case http.MethodOptions:
			ss.OptionsHandler(w, r)
		case http.MethodGet:
			ss.getAllTasksHandler(w, r)
		case http.MethodPost:
			ss.postTaskHandler(w, r)
		case http.MethodDelete:
			ss.deleteAllTasksHandler(w, r)
		default:
			http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
			return
		}
	} else {
		str := strings.Split(r.URL.Path, "/task/")[1]
		id, err := strconv.Atoi(str)
		if err != nil {
			http.Error(w, "incorrect id", http.StatusMethodNotAllowed)
			return
		}

		switch r.Method {
		case http.MethodGet:
			ss.getTaskHandler(w, r, id)
		case http.MethodDelete:
			ss.deleteTaskHandler(w, r, id)
		default:
			http.Error(w, "invalid http method (request with id)", http.StatusMethodNotAllowed)
			return
		}

	}
}

func (ss *serverStore) DueHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/due/" {
		switch r.Method {
		case http.MethodGet:
			ss.getTasksByDueDateHandler(w, r)
		default:
			http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
			return
		}

	} else {
		http.Error(w, "please enter date", http.StatusBadRequest)
		return
	}

}

func (ss *serverStore) TagHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/tag/" {
		switch r.Method {
		case http.MethodGet:
			ss.GetTasksByTagHandler(w, r)
		default:
			http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
			return
		}

	} else {
		http.Error(w, "please enter tag", http.StatusBadRequest)
		return
	}
}

func (ss *serverStore) OptionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET,POST,DELETE")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	w.Header().Add("Access-Control-Max-Age", "86400")
}

func (ss *serverStore) GetTasksByTagHandler(w http.ResponseWriter, r *http.Request) {
	tag := strings.Split(r.URL.Path, "/tag/")[1]

	tasks, err := json.Marshal(ss.store.GetTasksByTag(tag))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(tasks)
}

func (ss *serverStore) getTasksByDueDateHandler(w http.ResponseWriter, r *http.Request) {
	year, month, day, err := common.UrlToDate(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tasks, err := json.Marshal(ss.store.GetTasksByDueDate(year, month, day))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(tasks)
}

func (ss *serverStore) postTaskHandler(w http.ResponseWriter, r *http.Request) {
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

func (ss *serverStore) getAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks := ss.store.GetAllTasks()

	ts, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Origin", "*")
	w.Write(ts)
}

func (ss *serverStore) deleteAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	ss.store.DeleteAllTasks()

	msg := []byte("all tasks deleted")
	w.Write(msg)
}

func (ss *serverStore) getTaskHandler(w http.ResponseWriter, r *http.Request, id int) {
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

func (ss *serverStore) deleteTaskHandler(w http.ResponseWriter, r *http.Request, id int) {
	if err := ss.store.DeleteTask(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	str := []byte(fmt.Sprintf("task with id = %d deleted", id))
	w.Write(str)
}
