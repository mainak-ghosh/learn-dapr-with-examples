package activities

type Order struct {
	OrderID   string  `json:"orderId"`
	ProductID string  `json:"productId"`
	Quantity  int     `json:"quantity"`
	Amount    float64 `json:"amount"`
}

type ValidationResult struct {
	IsValid bool   `json:"isValid"`
	Message string `json:"message"`
}

type InventoryResult struct {
	InStock bool   `json:"inStock"`
	Message string `json:"message"`
}

type PaymentResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
