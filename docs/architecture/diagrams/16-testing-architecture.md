# Testing Architecture: Trophy Model

```mermaid
graph TD
    subgraph "Test Trophy Model"
        E2E[E2E Tests<br/>20% Coverage<br/>Patrol/k6] -->|Critical Flows| Integration[Integration Tests<br/>30-40% Coverage<br/>Testify/Testcontainers]
        Integration -->|Core Business Logic| Widget[Widget Tests<br/>40% Coverage<br/>Flutter Test/Golden Toolkit]
        Widget -->|Component Level| Unit[Unit Tests<br/>40-60% Coverage<br/>Mockery/Rapid Property Testing]
    end

    subgraph "Testing Infrastructure"
        Auth[Firebase Auth Emulator]
        Security[Security Testing Pipeline]
        Performance[Performance Testing<br/>k6 Load Tests]
    end

    subgraph "Test Distribution"
        Go[Go Backend Tests<br/>60% Unit<br/>30% Integration<br/>10% E2E]
        Flutter[Flutter Mobile Tests<br/>40% Unit<br/>40% Widget<br/>20% Integration]
    end

    E2E --> Security
    Integration --> Performance
    Unit --> Auth

    classDef trophy fill:#4A90E2,color:#ffffff,stroke:#333,stroke-width:2px;
    classDef infra fill:#2ECC71,color:#ffffff,stroke:#333,stroke-width:1px;
    classDef dist fill:#F39C12,color:#ffffff,stroke:#333,stroke-width:1px;

    class E2E,Integration,Widget,Unit trophy
    class Auth,Security,Performance infra
    class Go,Flutter dist
```

## Key Testing Principles
- Trophy Model emphasizes integration tests
- Selective TDD for critical paths
- Comprehensive test coverage across layers
- Performance and security integrated into testing strategy

## Testing Tools
### Backend (Go)
- Testify for test suites
- Mockery for mocking
- Testcontainers for integration testing
- Rapid property testing for complex logic

### Mobile (Flutter)
- flutter_test for unit/widget tests
- Mocktail for mocking
- Patrol for E2E testing
- Golden Toolkit for visual regression

## Test Distribution
- Go: 60% Unit, 30% Integration, 10% E2E
- Flutter: 40% Unit, 40% Widget, 20% Integration

## Infrastructure
- Firebase Auth Emulators
- Security Testing Pipeline
- Performance Testing with k6