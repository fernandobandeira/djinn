# Architectural Pattern Library

## Structural Patterns

### Model-View-Controller (MVC)
- **Category**: Architectural
- **Complexity**: Medium
- **Use Cases**: Web applications, GUI applications
- **Path**: patterns/structural/mvc.md

### Microservices
- **Category**: Architectural
- **Complexity**: High
- **Use Cases**: Distributed systems, scalable applications
- **Path**: patterns/architectural/microservices.md

### Layered Architecture
- **Category**: Architectural
- **Complexity**: Low
- **Use Cases**: Enterprise applications, monolithic systems
- **Path**: patterns/architectural/layered.md

## Behavioral Patterns

### Event Sourcing
- **Category**: Data Management
- **Complexity**: High
- **Use Cases**: Audit trails, event-driven architectures
- **Path**: patterns/behavioral/event-sourcing.md

### CQRS (Command Query Responsibility Segregation)
- **Category**: Data Management
- **Complexity**: Medium
- **Use Cases**: Read/write separation, performance optimization
- **Path**: patterns/behavioral/cqrs.md

## Creational Patterns

### Dependency Injection
- **Category**: Creational
- **Complexity**: Medium
- **Use Cases**: Testability, loose coupling
- **Path**: patterns/creational/dependency-injection.md

## Integration Patterns

### API Gateway
- **Category**: Integration
- **Complexity**: Medium
- **Use Cases**: Microservices, API management
- **Path**: patterns/integration/api-gateway.md

### Message Queue
- **Category**: Integration
- **Complexity**: Medium
- **Use Cases**: Asynchronous communication, decoupling
- **Path**: patterns/integration/message-queue.md

## Pattern Categories Index

### By Complexity
- **Low**: Layered Architecture
- **Medium**: MVC, CQRS, Dependency Injection, API Gateway, Message Queue
- **High**: Microservices, Event Sourcing

### By Domain
- **Web Development**: MVC, API Gateway
- **Distributed Systems**: Microservices, Message Queue
- **Data Management**: Event Sourcing, CQRS
- **Enterprise**: Layered Architecture, Dependency Injection