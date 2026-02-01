package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/victordaniel21/task-tracker/internal/data"
)

// CreateTask is now a method on *Dependencies
// This allows to access d.Models.Tasks inside the function!
func (d *Dependencies) CreateTask(w http.ResponseWriter, r *http.Request) {
	// 1. Define a struct to hold the incoming JSON request
	// We only expect 'title' and 'content' from the user.
	var input struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	// 2. Decode the JSON body into the struct
	// This reads the raw bytes from r.Body and maps them to our 'input' struct
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Bad Request: Invalid JSON", http.StatusBadRequest)
		return
	}

	// 3. Create the Task object
	task := &data.Task{
		Title:   input.Title,
		Content: input.Content,
		Status:  "pending", // Default status
	}

	// 4. Insert into Database (The Magic Moment!) 🪄
	// We call our model's Insert method. This runs the SQL.
	err = d.Models.Tasks.Insert(task)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 5. Respond with the created task (including the new ID)
	w.Header().Set("Content-Type", "application/json")
	// Write status 201 (Created)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// GetTask fetches a single task by ID
func (d *Dependencies) GetTask(w http.ResponseWriter, r *http.Request) {
	// 1. Read the ID parameter from the URL (Go 1.22 feature!)
	idString := r.PathValue("id")

	// 2. Convert string "1" to int64 1
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil || id < 1 {
		http.Error(w, "Bad Request: Invalid ID", http.StatusBadRequest)
		return
	}

	// 3. Call the Model
	task, err := d.Models.Tasks.Get(id)
	if err != nil {
		// Check for the specific "No Rows" error
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Not Found: Task does not exist", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	// 4. Return JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}
