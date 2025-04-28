package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"workflow/pkg/workflow"
	"workflow/pkg/workflow/activities"
)

const (
	stateStoreName = "statestore"
)

func main() {
	// Create workflow service
	workflowSvc := workflow.NewWorkflowService()

	// Handle workflow requests
	http.HandleFunc("/orderWorkflow", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var order activities.Order
		if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
			http.Error(w, fmt.Sprintf("Failed to decode order: %v", err), http.StatusBadRequest)
			return
		}

		result, err := workflowSvc.ProcessOrder(r.Context(), order)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to process order: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})

	// Start the server
	port := "8080"
	log.Printf("Workflow server listening on :%s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
