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
		d.badRequestResponse(w, errors.New("invalid id parameter"))
		return
	}

	// 3. Call the Model
	task, err := d.Models.Tasks.Get(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Use the helper!
			d.notFoundResponse(w)
		} else {
			// Use the helper!
			d.serverErrorResponse(w, err)
		}
		return
	}

	// 4. Use the helper for success too!
	err = d.writeJSON(w, http.StatusOK, task)
	if err != nil {
		d.serverErrorResponse(w, err)
	}

}

// ListTasks handles GET /v1/tasks
func (d *Dependencies) ListTasks(w http.ResponseWriter, r *http.Request) {
	// 1. Call the Model
	tasks, err := d.Models.Tasks.GetAll()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 2. Wrap the result in a parent JSON object (Best Practice)
	// Returning a top-level array (e.g. [{}, {}]) is valid but less flexible.
	// Wrapping it {"tasks": [...]} allows adding metadata (like "count": 5) later.
	response := map[string]interface{}{
		"tasks": tasks,
	}

	// 3. Send Response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// UpdateTask handles PUT /v1/tasks/{id}
func (d *Dependencies) UpdateTask(w http.ResponseWriter, r *http.Request) {
	// 1. Get the ID from the URL
	idString := r.PathValue("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil || id < 1 {
		http.Error(w, "Bad Request: Invalid ID", http.StatusBadRequest)
		return
	}

	// 2. Fetch the existing task first!
	// We need to know if it exists before we update it.
	task, err := d.Models.Tasks.Get(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Not Found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	// 3. Decode the User's new data
	// We define a temporary struct for the incoming JSON
	var input struct {
		Title   *string `json:"title"` // Pointer allows us to check if it was provided
		Content *string `json:"content"`
		Status  *string `json:"status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Bad Request: Invalid JSON", http.StatusBadRequest)
		return
	}

	// 4. Update the fields ONLY if the user provided them (Partial Update)
	if input.Title != nil {
		task.Title = *input.Title
	}
	if input.Content != nil {
		task.Content = *input.Content
	}
	if input.Status != nil {
		task.Status = *input.Status
	}

	// 5. Save changes to Database
	err = d.Models.Tasks.Update(task)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 6. Return the updated task
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (d *Dependencies) DeleteTask(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil || id < 1 {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err = d.Models.Tasks.Delete(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Not Found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	// Return 200 OK with a simple message
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "task deleted"}`))
}
