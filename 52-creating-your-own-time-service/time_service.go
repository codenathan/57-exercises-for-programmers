package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TimeResponse struct {
	CurrentTime string `json:"currentTime"`
}

func currentTimeHandler(w http.ResponseWriter, r *http.Request) {
	// Get the current time in UTC
	currentTime := time.Now().UTC().Format("2006-01-02 15:04:05")

	// Create a response object
	response := TimeResponse{
		CurrentTime: currentTime,
	}

	// Set the content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Send the response as JSON
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
	}
}

func main() {
	// Define the route and handler function
	http.HandleFunc("/current-time", currentTimeHandler)

	// Start the web server on port 8080
	fmt.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
