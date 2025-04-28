package workflow

import (
	"context"
	"encoding/json"
	"fmt"
	"workflow/pkg/workflow/activities"

	dapr "github.com/dapr/go-sdk/client"
)

type WorkflowService struct {
	daprClient dapr.Client
}

func NewWorkflowService() *WorkflowService {
	client, err := dapr.NewClient()
	if err != nil {
		panic(fmt.Sprintf("Failed to create Dapr client: %v", err))
	}
	return &WorkflowService{
		daprClient: client,
	}
}

func (s *WorkflowService) ProcessOrder(ctx context.Context, order activities.Order) (interface{}, error) {
	// Step 1: Validate Order
	validationResult, err := activities.ValidateOrder(ctx, order)
	if err != nil {
		return nil, fmt.Errorf("validation failed: %v", err)
	}

	// Save validation state
	if err := s.saveState(ctx, order.OrderID, "validation", validationResult); err != nil {
		return nil, fmt.Errorf("failed to save validation state: %v", err)
	}

	if !validationResult.IsValid {
		return validationResult, nil
	}

	// Step 2: Check Inventory
	inventoryResult, err := activities.CheckInventory(ctx, order)
	if err != nil {
		return nil, fmt.Errorf("inventory check failed: %v", err)
	}

	// Save inventory check state
	if err := s.saveState(ctx, order.OrderID, "inventory", inventoryResult); err != nil {
		return nil, fmt.Errorf("failed to save inventory state: %v", err)
	}

	if !inventoryResult.InStock {
		return inventoryResult, nil
	}

	// Step 3: Process Payment
	paymentResult, err := activities.ProcessPayment(ctx, order)
	if err != nil {
		return nil, fmt.Errorf("payment processing failed: %v", err)
	}

	// Save payment state
	if err := s.saveState(ctx, order.OrderID, "payment", paymentResult); err != nil {
		return nil, fmt.Errorf("failed to save payment state: %v", err)
	}

	return paymentResult, nil
}

func (s *WorkflowService) saveState(ctx context.Context, orderID string, step string, data interface{}) error {
	key := fmt.Sprintf("order-%s-%s", orderID, step)

	// Marshal data to JSON bytes
	bytes, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal state data: %v", err)
	}

	// Save state with empty metadata
	return s.daprClient.SaveState(ctx, "statestore", key, bytes, nil)
}
