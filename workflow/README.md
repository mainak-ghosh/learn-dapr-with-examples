# Dapr Workflow Example - Order Processing

This example demonstrates a simple order processing workflow using Dapr's building blocks. The workflow processes an order through three sequential steps: validation, inventory check, and payment processing.

## Prerequisites

1. [Dapr CLI](https://docs.dapr.io/getting-started/install-dapr-cli/)
2. [Go 1.20 or later](https://golang.org/dl/)
3. [Redis](https://redis.io/download) running locally on port 6379

## Project Structure

```
workflow/
├── main.go                      # Main application server
├── cmd/
│   └── client/                  # Client application
│       └── main.go
└── pkg/
    └── workflow/               # Workflow implementation
        ├── workflow.go
        └── activities/
            ├── activities.go   # Activity implementations
            └── types.go       # Data types
```

## Detailed Component Overview

### 1. Core Data Types (`types.go`)
- `Order`: Represents an order with fields:
  - `OrderID`: Unique identifier for the order
  - `ProductID`: Product being ordered
  - `Quantity`: Number of items ordered
  - `Amount`: Total order amount

- Result Types:
  - `ValidationResult`: Tracks order validation status
  - `InventoryResult`: Tracks inventory check status
  - `PaymentResult`: Tracks payment processing status

### 2. Workflow Activities (`activities.go`)
The workflow consists of three main activities:

- **Validation**:
  - Checks if order quantity > 0
  - Verifies order amount > 0
  - Returns validation status and message

- **Inventory Check**:
  - Simulates stock verification
  - Returns false if quantity > 100
  - Returns success message with reserved quantity

- **Payment Processing**:
  - Simulates payment verification
  - Rejects orders > $1000
  - Returns payment status and confirmation message

### 3. Workflow Service (`workflow.go`)
- Manages the end-to-end order processing flow
- Uses Dapr client for state management
- Processes orders through sequential steps:
  1. Order validation
  2. Inventory verification
  3. Payment processing
- Stores state in Redis after each step

### 4. Main Server (`main.go`)
- Exposes HTTP endpoint `/orderWorkflow`
- Handles POST requests with order data
- Initializes workflow service
- Returns workflow results to clients

### 5. Client Application (`cmd/client/main.go`)
- Creates sample orders with random UUIDs
- Sends orders to workflow service via Dapr sidecar
- Displays workflow results

## Dapr Components

1. **State Store** (`workflow.yaml`):
   - Uses Redis as state store
   - Configured for storing workflow state
   - Enables actor state storage

2. **Service Configuration** (`service.yaml`):
   - Configures service discovery using mDNS
   - Sets up tracing with 100% sampling rate

## Running the Example

1. Start the workflow server:
   ```bash
   dapr run --app-id workflowapp --app-port 8080 --dapr-http-port 3500 --resources-path ./resources go run main.go
   ```

2. In another terminal, send a test order:
   ```bash
   cd cmd/client
   dapr run --app-id workflow-client go run main.go
   ```

## Testing Different Scenarios

You can test different scenarios by modifying the order in `cmd/client/main.go`:

1. Invalid Order: Set quantity or amount to 0
2. Out of Stock: Set quantity > 100
3. Payment Limit Exceeded: Set amount > 1000

## State Management

The workflow maintains state in Redis with the following key pattern:
- `order-{orderID}-validation`: Validation step results
- `order-{orderID}-inventory`: Inventory check results
- `order-{orderID}-payment`: Payment processing results

You can inspect the workflow state using Redis CLI:
```bash
redis-cli keys "order-*"
```

## Error Handling

The workflow implements comprehensive error handling:
- Validation failures return early with error messages
- Inventory checks prevent processing unavailable items
- Payment processing validates transaction limits
- Each step's state is preserved for tracking and debugging

## Best Practices Demonstrated

1. **Separation of Concerns**:
   - Activities are isolated in separate functions
   - Clear distinction between workflow logic and HTTP handling

2. **State Management**:
   - Persistent state after each step
   - Trackable workflow progress

3. **Error Handling**:
   - Graceful failure handling
   - Clear error messages
   - State preservation on failures

4. **Modularity**:
   - Reusable activity functions
   - Configurable components
   - Easy to extend with new steps
