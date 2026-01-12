package main

import (
	"fmt"
	"log"
	"net/http"

	// 👇 This is how we import local packages.
	// Replace "github.com/yourusername/task-tracker" with the module name found in your go.mod file
	"github.com/yourusername/task-tracker/internal/validator"
)

func main() {
	port := ":8080"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Calling the exported function from our internal package
		isValid := validator.IsEmailValid("test@example.com")

		// %v is a generic placeholder for values in Sprintf/Printf
		response := fmt.Sprintf("Welcome! System Check: Validator is working = %v", isValid)

		fmt.Fprint(w, response)
	})

	log.Printf("Starting server on %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
