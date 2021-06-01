package server

import (
	"database/sql"
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
