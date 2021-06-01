package server

import (
	"database/sql"
	"net/http"
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

func (sql *MysqlStore) TaskDeleteHandler(w http.ResponseWriter, r *http.Request) {

}
