package data

import (
	"database/sql"
	"time"
)

// Task represent a singe row in database.
type Task struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Status    string    `json:"status"`
	Version   int32     `json:"version"`
}

// TaskModel wraps the connection pool.
// It will hold all the methods for talking to the 'task' table.
type TaskModel struct {
	DB *sql.DB
}

func (m TaskModel) Insert(task *Task) error {
	// 1. The SQL query
	query := `
	insert into tasks (title, content, status)
	values ($1, $2, $3)
	returning id, created_at, version`

	// 2. Execute the Query
	//  QueryRow executes a query that returns exactly one row.
	args := []interface{}{task.Title, task.Content, task.Status}

	// Scan copies the columns from the returned row into our task struct
	return m.DB.QueryRow(query, args...).Scan(&task.ID, &task.CreatedAt, &task.Version)
}
