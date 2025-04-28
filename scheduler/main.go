package main

import (
	// "encoding/json"
	"fmt"
	// "io/ioutil"
	"log"
	"net/http"
	"os"
)

// BindingEvent represents the data received from Dapr binding
type BindingEvent struct {
	Data     map[string]interface{} `json:"data"`
	Metadata map[string]string      `json:"metadata"`
}

func callExternalAPI(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request from %s", r.RemoteAddr)
	fmt.Println("Received Dapr Job Scheduler trigger!")

	fmt.Println("Calling the external API...")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func handleScheduledJob(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received scheduled job trigger from %s", r.RemoteAddr)

	// Read the request body
	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	log.Printf("Error reading request body: %v", err)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	// Parse the binding event
	// var event BindingEvent
	// if err := json.Unmarshal(body, &event); err != nil {
	// 	log.Printf("Error unmarshaling binding event: %v", err)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	fmt.Println("Calling the external API...")
	// Add your API call logic here

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	port := "3000"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	// Handle both the direct HTTP endpoint and the binding route
	http.HandleFunc("/call-api", callExternalAPI)
	http.HandleFunc("/scheduled-job", handleScheduledJob)

	log.Printf("Starting server on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
