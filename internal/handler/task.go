package handler

import (
	"encoding/json"
	"net/http"
)

// createtask handler to check if the service is running
func CreateTask(w http.ResponseWriter, r *http.Request) {
	//1. create a map to hold our JSON data
	data := map[string]string{
		"message": "Create task endpoint",
	}

	//2. convert the map to JSON
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	//3. set the content type header so the browser knows it's JSON
	w.Header().Set("Content-Type", "application/json")

	//4. write the JSON bytes to the response
	w.Write(js)
}
