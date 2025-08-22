# Code Organization Architecture Diagram

## Context
Comprehensive architectural diagram showcasing the modular structure of the Djinn personal finance application, highlighting clean/hexagonal architecture principles, domain-driven design, and technology-specific module boundaries.

## System Architecture Diagram

```mermaid
graph TD
    subgraph Backend Go Architecture
        direction TB
        
        %% Layers
        Domain["Domain Layer (Core Business Logic)"]
        Application["Application Layer (Use Cases)"]
        Adapters["Adapter Layer (Infrastructure)"]
        Infrastructure["Infrastructure Layer (Framework Integration)"]
        GraphQL["GraphQL Presentation Layer"]
        
        %% Dependency Flow
        Domain --> Application
        Application --> Adapters
        Adapters --> Infrastructure
        Application --> GraphQL
        Domain --> GraphQL
        
        %% Color Coding
        classDef domain fill:#4A90E2,color:white;
        classDef application fill:#2ECC71,color:white;
        classDef adapters fill:#F39C12,color:white;
        classDef infrastructure fill:#E74C3C,color:white;
        classDef graphql fill:#9B59B6,color:white;
        
        class Domain domain;
        class Application application;
        class Adapters adapters;
        class Infrastructure infrastructure;
        class GraphQL graphql;
    end

    subgraph Frontend Flutter Architecture
        direction TB
        
        %% Flutter Feature Layers
        CoreLayer["Core Layer\n(Shared Utilities)"]
        DataLayer["Data Layer\n(Repositories, Data Sources)"]
        DomainLayer["Domain Layer\n(Entities, Use Cases)"]
        PresentationLayer["Presentation Layer\n(UI, State Management)"]
        
        %% Dependency Flow
        CoreLayer --> DataLayer
        CoreLayer --> DomainLayer
        DataLayer --> DomainLayer
        DomainLayer --> PresentationLayer
        
        %% Color Coding
        classDef core fill:#3498DB,color:white;
        classDef data fill:#2ECC71,color:white;
        classDef domain fill:#F39C12,color:white;
        classDef presentation fill:#9B59B6,color:white;
        
        class CoreLayer core;
        class DataLayer data;
        class DomainLayer domain;
        class PresentationLayer presentation;
    end

    %% Cross-Platform Interaction
    GraphQL --- API[GraphQL API]
    API --- PresentationLayer
</graph>
```

## Architectural Principles

### Go Backend
- **Domain Layer**: Pure business logic, independent of frameworks
- **Application Layer**: Use cases and application services
- **Adapter Layer**: Infrastructure adapters and external service integrations
- **Infrastructure Layer**: Framework-specific configurations and setup
- **GraphQL Layer**: Presentation and API interaction

### Flutter Frontend
- **Core Layer**: Shared utilities, constants, and cross-cutting concerns
- **Data Layer**: Repositories, data sources, and external service clients
- **Domain Layer**: Business entities and use case implementations
- **Presentation Layer**: UI components, state management, and user interactions

## Key Design Characteristics
- Strict Dependency Rule: Inner layers never depend on outer layers
- Type Safety: End-to-end type checking
- Testability: Clear separation enables comprehensive testing
- Scalability: Modular structure supports feature growth
- Framework Isolation: Core logic independent of technology details

## Dependency Flow Direction
- Backend: Domain → Application → Adapters → Infrastructure
- Frontend: Core → Data → Domain → Presentation
- GraphQL acts as a bridge between Application and Presentation layers

## Technology Integration
- Backend: Go, Wire DI, gqlgen, sqlc
- Frontend: Flutter, Riverpod, Ferry GraphQL, Drift
- Shared: GraphQL API for cross-platform communication