package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

// writeJson is a helper to write JSON response
// it handles setting the header and checking for encoding errors
func (d *Dependencies) writeJSON(w http.ResponseWriter, status int, data interface{}) error {
	// Set the header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	// Encode the data
	return json.NewEncoder(w).Encode(data)
}

// errorResponse wraps an error message in JSON: {"error": "message"}
func (d *Dependencies) errorResponse(w http.ResponseWriter, status int, message interface{}) {
	env := map[string]interface{}{"error": message}

	err := d.writeJSON(w, status, env)
	if err != nil {
		// If we can't even write the error response, we log it and give up.
		log.Printf("Failed to write error response: %v", err)
		w.WriteHeader(500)
	}
}

// serverErrorResponse logs the detailed error (for us) and sends a generic message (to the user).
func (d *Dependencies) serverErrorResponse(w http.ResponseWriter, err error) {
	log.Printf("ERROR: %v", err) // Log the real error to the console

	message := "the server encountered a problem and could not process your request"
	d.errorResponse(w, http.StatusInternalServerError, message)
}

// notFoundResponse sends a 404 JSON error
func (d *Dependencies) notFoundResponse(w http.ResponseWriter) {
	message := "the requested resource could not be found"
	d.errorResponse(w, http.StatusNotFound, message)
}

// badRequestResponse sends a 400 JSON error
func (d *Dependencies) badRequestResponse(w http.ResponseWriter, err error) {
	d.errorResponse(w, http.StatusBadRequest, err.Error())
}
