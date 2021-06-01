package server

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"path"
	"strconv"
	"time"
)

// Task
// TaskId int not null auto_increment primary key,
// UserId int not null,
// TaskName varchar(255) not null UNIQUE,
// TaskType varchar(255) not null,
// InitTime time not null

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
