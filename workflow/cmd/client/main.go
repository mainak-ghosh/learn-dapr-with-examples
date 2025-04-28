package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"workflow/pkg/workflow/activities"

	"github.com/google/uuid"
)

func main() {
	// Create a sample order
	order := activities.Order{
		OrderID:   uuid.New().String(),
		ProductID: "PROD-001",
		Quantity:  5,
		Amount:    99.99,
	}

	log.Printf("Processing order: %+v", order)

	// Convert order to JSON
	orderBytes, err := json.Marshal(order)
	if err != nil {
		log.Fatalf("Failed to marshal order: %v", err)
	}

	// Call the workflow server directly through Dapr sidecar
	url := "http://localhost:3500/v1.0/invoke/workflowapp/method/orderWorkflow"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(orderBytes))
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Workflow failed with status %d: %s", resp.StatusCode, string(body))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Printf("Failed to unmarshal response: %v. Raw response: %s", err, string(body))
	} else {
		log.Printf("Workflow result: %+v", result)
	}

	// Print success
	fmt.Printf("Successfully processed order %s\n", order.OrderID)
}
