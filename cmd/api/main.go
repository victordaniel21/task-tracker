package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"

	// Import your handler AND data packages
	"github.com/victordaniel21/task-tracker/internal/data"
	"github.com/victordaniel21/task-tracker/internal/handler"
)

func main() {
	port := ":8080"

	// 1. Database Connection String (DSN)
	// format: postgres://user:password@host:port/dbname?sslmode=disable
	// In a real app, this comes from ENV variables, not hardcoded!
	dsn := "postgres://postgres:secret@localhost:5433/taskdb?sslmode=disable"

	// 2. Connect to Database
	db, err := data.OpenDB(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	// defer ensures the DB connection closes when main() exits (e.g. if we crash)
	defer db.Close()

	// 3. Initialize Models
	models := data.NewModels(db)

	// 4. Initialize Handlers (Injecting the Models)
	app := handler.NewDependencies(models)

	// 5. Setup Router
	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/health", handler.HealthCheck)
	mux.HandleFunc("POST /v1/tasks", app.CreateTask)
	mux.HandleFunc("GET /v1/tasks/{id}", app.GetTask)
	mux.HandleFunc("GET /v1/tasks", app.ListTasks)
	mux.HandleFunc("PUT /v1/tasks/{id}", app.UpdateTask)    // Update
	mux.HandleFunc("DELETE /v1/tasks/{id}", app.DeleteTask) // Delete

	log.Printf("Starting server on %s", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}
