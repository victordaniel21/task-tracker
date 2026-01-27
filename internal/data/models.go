package data

import "database/sql"

// Models wrapper that holds all seperate data models
type Models struct {
	Tasks TaskModel
}

// NewModels returns a Models struct containing the initialized TaskModel
func NewModels(db *sql.DB) Models {
	return Models{
		Tasks: TaskModel{DB: db},
	}
}
