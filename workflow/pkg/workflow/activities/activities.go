package activities

import (
	"context"
	"fmt"
)

func ValidateOrder(ctx context.Context, order Order) (ValidationResult, error) {
	// Simple validation logic
	if order.Quantity <= 0 {
		return ValidationResult{
			IsValid: false,
			Message: "Quantity must be greater than 0",
		}, nil
	}
	if order.Amount <= 0 {
		return ValidationResult{
			IsValid: false,
			Message: "Amount must be greater than 0",
		}, nil
	}
	return ValidationResult{
		IsValid: true,
		Message: "Order validation successful",
	}, nil
}

func CheckInventory(ctx context.Context, order Order) (InventoryResult, error) {
	// Simulate inventory check
	if order.Quantity > 100 {
		return InventoryResult{
			InStock: false,
			Message: "Not enough items in stock",
		}, nil
	}
	return InventoryResult{
		InStock: true,
		Message: fmt.Sprintf("Successfully reserved %d items", order.Quantity),
	}, nil
}

func ProcessPayment(ctx context.Context, order Order) (PaymentResult, error) {
	// Simulate payment processing
	if order.Amount > 1000 {
		return PaymentResult{
			Success: false,
			Message: "Amount exceeds maximum limit",
		}, nil
	}
	return PaymentResult{
		Success: true,
		Message: fmt.Sprintf("Successfully processed payment for $%.2f", order.Amount),
	}, nil
}
