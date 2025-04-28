# Learn Dapr with Examples

This repository contains practical examples demonstrating how to use [Dapr (Distributed Application Runtime)](https://dapr.io/) in Go applications. Dapr is a portable, event-driven runtime that makes it easy to build resilient, microservice-based applications.

## What is Dapr?

Dapr is an open-source project that simplifies cloud-native application development by providing:

- Building blocks for common patterns and best practices
- Language-agnostic APIs
- Platform independence (runs on cloud or edge)
- Microservice architecture support
- Built-in state management, pub/sub, and more

### Key Building Blocks Used in Examples:

1. **State Management**: Persisting and retrieving application state
2. **Service Invocation**: Direct service-to-service communication
3. **Pub/Sub**: Asynchronous messaging between services
4. **Bindings**: Triggers from external systems and output integration
5. **Actors**: Pattern for stateful, concurrent objects

## Repository Structure

This repository contains two main examples:

### 1. Workflow Service (`/workflow`)
A comprehensive example demonstrating:
- Order processing workflow implementation
- Sequential state management
- Multi-step business process orchestration
- Activity-based workflow patterns
- State persistence using Redis
- Client-server interaction through Dapr

Key components:
- Main workflow service
- Client application for testing
- Activity implementations
- Type definitions
- Dapr configuration files

### 2. Scheduler Service (`/scheduler`)
Demonstrates Dapr's scheduling capabilities:
- Cron-based job scheduling
- Binding implementations
- Scheduled task execution
- External service triggers

Key features:
- Scheduled job execution
- API endpoints for direct and scheduled calls
- Dapr binding configuration
- Job scheduling patterns

## Getting Started

### Prerequisites
1. [Dapr CLI](https://docs.dapr.io/getting-started/install-dapr-cli/)
2. [Go 1.20+](https://golang.org/dl/)
3. [Docker](https://www.docker.com/get-started)
4. [Redis](https://redis.io/download) (for state management)

### Running the Examples

Each folder contains its own README with specific instructions, but generally:

1. Start the Dapr sidecar
2. Run the application
3. Test with provided clients or API endpoints

## Learning Path

1. Start with the **Workflow** example to understand:
   - Basic Dapr concepts
   - State management
   - Service communication
   - Activity orchestration

2. Move to the **Scheduler** example to learn:
   - Dapr bindings
   - Cron scheduling
   - External triggers
   - Job management

## Purpose of This Repository

This repository serves as a practical learning resource for developers who want to:
1. Understand Dapr's core concepts through working examples
2. See real-world implementations of Dapr building blocks
3. Learn best practices for microservice development
4. Explore different patterns for distributed applications

Each example is thoroughly documented and implements production-ready patterns that you can adapt for your own applications.

## Contributing

Feel free to:
- Submit issues for bugs or feature requests
- Create pull requests with improvements
- Add more examples demonstrating other Dapr capabilities

## Resources

- [Official Dapr Documentation](https://docs.dapr.io/)
- [Dapr GitHub Repository](https://github.com/dapr/dapr)
- [Dapr Samples](https://github.com/dapr/samples)
