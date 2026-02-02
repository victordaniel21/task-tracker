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

// Get retrieves a specific task by ID.
func (m TaskModel) Get(id int64) (*Task, error) {
	// 1. The SQL Query
	if id < 1 {
		return nil, sql.ErrNoRows
	}

	query := `
		SELECT id, created_at, title, content, status, version
		FROM tasks
		WHERE id = $1`

	// 2. Prepare the struct to hold data
	var task Task

	// 3. Execute and Scan
	// We pass &task.ID, &task.CreatedAt etc. as pointers so Scan can fill them.
	err := m.DB.QueryRow(query, id).Scan(
		&task.ID,
		&task.CreatedAt,
		&task.Title,
		&task.Content,
		&task.Status,
		&task.Version,
	)

	// 4. Handle Errors
	if err != nil {
		// Specifically check if the error is "row not found"
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows // We will handle this in the controller (404)
		}
		return nil, err // Real DB error (500)
	}

	return &task, nil
}

func (m TaskModel) GetAll() ([]*Task, error) {
	// 1. Define the query
	query := `
	select id, created_at, title, content, status, version
	from tasks
	order by id desc`

	// 2. Execute the query
	// query returns a 'Rows' object whic acts like a cursor
	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}

	// IMPORTANT!: Close the rows when the functioin returns to free up the connection
	defer rows.Close()

	// 3. Iterate through the rows
	tasks := []*Task{} // Initialize an empty slice

	for rows.Next() {
		var task Task
		// Scan the values from the current row into the struct
		err := rows.Scan(
			&task.ID,
			&task.CreatedAt,
			&task.Title,
			&task.Content,
			&task.Status,
			&task.Version,
		)
		if err != nil {
			return nil, err
		}

		// Append to the slice
		tasks = append(tasks, &task)
	}
	// 4. Check for errors that occured during iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
