package data

import "database/sql"

// Models wrapper that holds all seperate data models
type Models struct {
	Task TaskModel
}

// NewModels returns a Models struct containing the initialized TaskModel
func NewModels(db *sql.DB) Models {
	return Models{
		Task: TaskModel{DB: db},
	}
}
