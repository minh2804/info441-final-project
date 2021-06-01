package server

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"path"
	"time"
)

type TaskItem struct {
	ID       int64     `json:"id"`
	UserID   int64     `json:"userId"`
	TaskName string    `json:"taskName"`
	TaskType string    `json:"taskType"`
	InitTime time.Time `json:"initTime"`
}

type MysqlStore struct {
	DB *sql.DB
}

func (sql *MysqlStore) TaskPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	} else {
		task := &TaskItem{}
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		insq := "INSERT INTO Tasks (UserId, TaskName, TaskType, InitTime) VALUES (?, ?, ?, ?)"
		res, err := sql.DB.Exec(insq, task.UserID, task.TaskName,
			task.TaskType, time.Now())
		if err != nil {
			w.Write([]byte(err.Error()))
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		} else {
			id, err := res.LastInsertId()
			if err != nil {
				w.Write([]byte(err.Error()))
				http.Error(w, err.Error(), http.StatusForbidden)
				return
			} else {
				task.ID = id
			}
		}
		w.Write([]byte("post is successful"))
	}
}

func (sql *MysqlStore) TaskPatchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PATCH" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	} else {
		taskId := path.Base(r.URL.Path)
		task := &TaskItem{}
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		insq := "Update Tasks SET UserId=?, TaskName=?, TaskType=? WHERE TaskId=?"
		res, err := sql.DB.Exec(insq, task.UserID, task.TaskName,
			task.TaskType, taskId)
		if err != nil {
			w.Write([]byte(err.Error()))
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		} else {
			id, err := res.LastInsertId()
			if err != nil {
				w.Write([]byte(err.Error()))
				http.Error(w, err.Error(), http.StatusForbidden)
				return
			} else {
				task.ID = id
			}
		}
		w.Write([]byte("update is successful"))
	}
}
