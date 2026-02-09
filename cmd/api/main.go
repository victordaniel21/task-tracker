package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"

	// Import your handler AND data packages
	"github.com/victordaniel21/task-tracker/internal/data"
	"github.com/victordaniel21/task-tracker/internal/handler"
	"github.com/victordaniel21/task-tracker/internal/middleware"
)

// Define a config struct to hold all our settings
type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

func main() {

	var cfg config

	// 1. Define Flags
	// Instead of hardcoding ":8080", we read it from command line flags.
	// Default values are set here (8080, development, etc.)
	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://postgres:secret@localhost:5433/taskdb?sslmode=disable", "PostgreSQL DSN")

	// 2. Parse the flags
	// This reads what you typed in the terminal (e.g., -port=4000)
	flag.Parse()

	// 3. Connect to DB (Using the cfg variable now!)
	db, err := data.OpenDB(cfg.db.dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	log.Printf("Database connection pool established")

	// Initialize Models & Dependencies
	models := data.NewModels(db)
	app := handler.NewDependencies(models)

	// Setup Router
	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/health", handler.HealthCheck)
	mux.HandleFunc("POST /v1/tasks", app.CreateTask)
	mux.HandleFunc("GET /v1/tasks", app.ListTasks)
	mux.HandleFunc("GET /v1/tasks/{id}", app.GetTask)
	mux.HandleFunc("PUT /v1/tasks/{id}", app.UpdateTask)
	mux.HandleFunc("DELETE /v1/tasks/{id}", app.DeleteTask)

	// 4. Start the server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      middleware.EnableCORS(mux),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// 5. Start the server in a background goroutine
	// We do this so it doesn't block the main thread, allowing us to listen for signals below.
	go func() {
		log.Printf("Starting %s server on %s", cfg.env, srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 6. Listen for Shutdown Signals
	// We create a channel to listen for OS signals (Ctrl+C or SIGTERM)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// This blocks until a signal is received
	<-quit
	log.Println("Server is shutting down...")

	// 7. Create a deadline to wait for active requests to complete
	// We give active requests 5 seconds to finish before forcing close.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
