package server

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"path"

	"strconv"

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

func (sql *MysqlStore) GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	} else {
		itemID := path.Base(r.URL.Path)
		userID, err := strconv.ParseInt(itemID, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		sqlStatement := "select UserId, TaskName, TaskType, InitTime from task where TaskId=?"
		rows, err := sql.DB.Query(sqlStatement, userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		task := &TaskItem{}

		for rows.Next() {
			err = rows.Scan(task.UserID, task.TaskName, task.TaskType, task.InitTime)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
		}

		encodedTask, err := json.Marshal(task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		w.Write(encodedTask)
		w.Header().Set("Content-Type", "application/json")
	}
}

func (sql *MysqlStore) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	} else {
		itemID := path.Base(r.URL.Path)
		userID, err := strconv.ParseInt(itemID, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		sqlStatement := "DELETE FROM task WHERE TaskId = ?;"
		_, err = sql.DB.Query(sqlStatement, userID)
		if err != nil {
			w.Write([]byte(err.Error()))
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		w.Write([]byte("Delete is successful"))

	}
}
