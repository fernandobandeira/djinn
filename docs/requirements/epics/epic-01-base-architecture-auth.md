# Epic 1: Base Architecture & Authentication Foundation

**Epic ID**: DJINN-E001  
**Priority**: Critical Path  
**Sprint Size**: 8-10 stories (4-6 weeks)  
**Dependencies**: None (Foundation Epic)  

## Epic Overview

### Title
Establish Flutter-Golang base architecture with Firebase OAuth authentication flow and biometric protection

### Description
Create the foundational architecture for the Djinn personal finance application by implementing the complete authentication stack, base application structure, and secure communication layer between Flutter mobile app and Golang GraphQL API. This epic establishes the security-first, privacy-focused foundation required for all subsequent feature development.

### Business Value & Strategic Alignment

**Strategic Alignment with 50,000+ Subscriber Goal:**
- **Trust Foundation**: Privacy-first OAuth-only authentication builds user trust essential for financial data handling
- **User Experience**: Seamless biometric authentication creates "magical" experience aligned with Djinn's AI personality
- **Security Positioning**: Zero Trust architecture differentiates from competitors in privacy-conscious market
- **Technical Foundation**: Enables rapid feature velocity for MVP launch in 3-6 months

**Revenue Impact:**
- **User Acquisition**: Secure, frictionless onboarding reduces drop-off rates by 15-20%
- **Retention**: Biometric protection builds confidence for long-term financial data storage
- **Market Differentiation**: OAuth-only + biometric protection appeals to ex-Mint users seeking privacy

### Architecture Alignment

This epic implements decisions from:
- **ADR-20250812**: Tech stack selection (Flutter, Go, Firebase Auth, GraphQL)
- **ADR-20250819**: Security architecture (OAuth-only, biometric protection, Zero Trust)
- **ADR-20250820**: Testing strategy (TDD for auth, 90%+ coverage)

## User Stories

### Story 1: Flutter Base Architecture Setup
**Story ID**: DJINN-1001  
**Effort**: 6 hours  
**Priority**: Must Have  

**As a** developer  
**I want** a properly structured Flutter application foundation  
**So that** all future features follow consistent patterns and maintain code quality  

**Acceptance Criteria:**
- [ ] Flutter project created with proper folder structure following ADR-20250820-code-organization
- [ ] Riverpod 2.0 state management configured with providers structure
- [ ] go_router navigation setup with authenticated/unauthenticated routes
- [ ] Ferry GraphQL client configured with HiveStore cache
- [ ] Drift database setup with initial user profile table
- [ ] Basic app theming configured (colors, typography, components)
- [ ] Environment configuration (dev/staging/prod) implemented
- [ ] All setup follows TDD principles with widget tests

**Technical Requirements:**
- Follow Flutter best practices from ADR-20250812
- Implement offline-first architecture from ADR-20250819-mobile-offline-first
- Include proper error handling following ADR-20250120-error-handling

**Definition of Done:**
- App builds and runs on iOS/Android simulators
- Unit tests pass with 90%+ coverage
- Code review completed and approved
- Documentation updated in /docs/architecture/

### Story 2: Golang GraphQL API Foundation with Database Setup
**Story ID**: DJINN-1002  
**Effort**: 10 hours  
**Priority**: Must Have  

**As a** mobile app  
**I want** a secure GraphQL API endpoint with proper database foundation  
**So that** I can communicate with the backend and persist user data securely  

**Acceptance Criteria:**
- [ ] Go project structure following ADR-20250820-code-organization
- [ ] PostgreSQL 16 database configured with UUIDv7 support
- [ ] Atlas migrations setup with initial user table migration
- [ ] User table created with schema: id (UUIDv7), firebase_uid (unique), email, name, profile_image_url, created_at, updated_at
- [ ] sqlc configured with queries for user CRUD operations
- [ ] Chi HTTP framework configured with middleware stack
- [ ] gqlgen GraphQL server setup with schema-first approach
- [ ] User GraphQL schema and resolvers (getUser, createUser mutations)
- [ ] DataLoader pattern implemented for User queries (prevent N+1)
- [ ] PgBouncer connection pooling configured
- [ ] Health check endpoint for monitoring
- [ ] Environment configuration with envconfig
- [ ] OpenTelemetry tracing configured following ADR-20250120-monitoring

**Technical Requirements:**
- PostgreSQL 16 with native UUIDv7 support
- Atlas for schema migrations (reversible migrations)
- sqlc with pgx/v5 driver for type-safe queries
- dataloaden for DataLoader generation
- Follow Go project layout standards
- Implement proper error handling from ADR-20250120-error-handling
- Include structured logging with log/slog
- Docker containerization for local development

**Definition of Done:**
- Database migrations run successfully
- User table exists with proper schema
- sqlc generates type-safe queries
- API serves GraphQL playground at /graphql
- User queries/mutations work end-to-end
- DataLoader prevents N+1 queries
- Health check responds at /health
- Unit tests pass with 90%+ coverage
- Integration tests verify database connectivity
- API documented in OpenAPI/GraphQL schema

### Story 3: Firebase Authentication Integration
**Story ID**: DJINN-1003  
**Effort**: 10 hours  
**Priority**: Must Have  

**As a** user  
**I want** to sign in with Google or Apple ID  
**So that** I can securely access my financial data without creating passwords  

**Acceptance Criteria:**
- [ ] Firebase project configured for iOS and Android
- [ ] Google Sign-In implemented in Flutter with proper error handling
- [ ] Apple Sign-In implemented in Flutter with proper error handling
- [ ] Firebase Auth token validation in Golang backend middleware
- [ ] JWT middleware extracts Firebase UID and validates tokens
- [ ] User creation flow: Check if firebase_uid exists in DB, create if new user
- [ ] CreateUser mutation generates UUIDv7 ID and stores firebase_uid, email, name, profile_image_url from OAuth provider
- [ ] GetOrCreateUser resolver pattern implemented
- [ ] Proper error messages for authentication failures
- [ ] Sign-out functionality implemented

**Technical Requirements:**
- OAuth-only authentication (no email/password) per ADR-20250819
- Firebase token verification in Go using Firebase Admin SDK
- User record creation using sqlc-generated queries
- Atomic user creation with proper conflict handling (firebase_uid unique constraint)
- Secure token storage in Flutter using flutter_secure_storage
- Follow security principles from ADR-20250819-security-architecture

**Definition of Done:**
- Users can sign in with Google and Apple ID
- Backend validates Firebase tokens correctly
- New users automatically get created in PostgreSQL users table
- Existing users are retrieved by firebase_uid
- Authentication state persists across app restarts
- All error scenarios handled gracefully
- Database integration tests verify user creation flow
- Security testing completed

### Story 4: Biometric Protection Implementation
**Story ID**: DJINN-1004  
**Effort**: 8 hours  
**Priority**: Must Have  

**As a** user concerned about device security  
**I want** biometric authentication for app access  
**So that** my financial data is protected if someone else accesses my device  

**Acceptance Criteria:**
- [ ] local_auth package integrated for biometric authentication
- [ ] Face ID/Touch ID prompt on app launch from closed state
- [ ] Biometric re-authentication after 5+ minutes in background
- [ ] Fallback to device passcode if biometrics unavailable
- [ ] Biometric settings in user preferences
- [ ] Proper handling of biometric enrollment changes
- [ ] Clear error messages for biometric failures

**Technical Requirements:**
- Follow biometric security requirements from ADR-20250819
- Implement biometric checks for all required scenarios
- Handle device-specific biometric capabilities
- Secure biometric preference storage

**Definition of Done:**
- Biometric authentication works on iOS and Android
- All required scenarios trigger biometric prompts
- User can configure biometric preferences
- Fallback scenarios work correctly
- Security audit completed

### Story 5: Welcome Screen with User Profile
**Story ID**: DJINN-1005  
**Effort**: 4 hours  
**Priority**: Must Have  

**As a** newly authenticated user  
**I want** to see a welcome screen with my profile information from the database  
**So that** I know the authentication worked and can see my account details  

**Acceptance Criteria:**
- [ ] GraphQL query implemented to fetch user by Firebase UID
- [ ] Welcome screen fetches user data from GraphQL API (not from JWT/Firebase)
- [ ] Displays user's name from database (originally from OAuth provider)
- [ ] Profile image URL from database displayed
- [ ] Email address from database shown
- [ ] Account creation timestamp from database displayed
- [ ] Ferry GraphQL client makes authenticated request with Firebase token
- [ ] Sign out button functions correctly
- [ ] Proper loading states during GraphQL data fetch
- [ ] Error handling for GraphQL query failures
- [ ] Offline mode shows cached data from Ferry's HiveStore

**Technical Requirements:**
- Use Ferry GraphQL client for data fetching
- GraphQL query: `query GetUser($firebaseUid: String!) { user(firebaseUid: $firebaseUid) { id, name, email, profileImageUrl, createdAt } }`
- Use Riverpod for state management
- Follow UI/UX patterns from existing wireframes
- Ferry's normalized cache provides offline data
- Handle offline scenarios gracefully with cached data

**Definition of Done:**
- Welcome screen displays correctly after authentication
- All user data populates from backend
- Loading and error states implemented
- Screen follows design specifications
- Accessibility standards met

### Story 6: GraphQL Communication Security
**Story ID**: DJINN-1006  
**Effort**: 6 hours  
**Priority**: Must Have  

**As a** security-conscious application  
**I want** all Flutter-Golang communication to be properly secured  
**So that** user data is protected during transmission and unauthorized access is prevented  

**Acceptance Criteria:**
- [ ] HTTPS/TLS enforced for all API communication
- [ ] Authorization headers properly set in Ferry GraphQL client with Firebase tokens
- [ ] Firebase Auth SDK configured for automatic token refresh (30-day validity)
- [ ] Proper error handling for invalid tokens (401 responses)
- [ ] Network security configuration for mobile apps
- [ ] Request/response logging for debugging (excluding sensitive data)
- [ ] Rate limiting configured per ADR-20250820-rate-limiting

**Technical Requirements:**
- Follow Zero Trust principles from ADR-20250819
- Firebase Auth handles token lifecycle (30-day validity, auto-refresh)
- Implement proper certificate pinning for production
- Configure network security policies for Android/iOS
- Include security headers in all API responses

**Definition of Done:**
- All communication uses HTTPS/TLS
- Authentication errors handled gracefully
- Firebase handles token refresh automatically (no manual implementation needed)
- Security testing completed
- Network security audit passed

### Story 7: Testing Infrastructure Setup
**Story ID**: DJINN-1007  
**Effort**: 10 hours  
**Priority**: Should Have  

**As a** development team  
**I want** comprehensive testing infrastructure  
**So that** we can maintain code quality and catch regressions early  

**Acceptance Criteria:**
- [ ] Flutter widget tests for authentication flows
- [ ] Flutter integration tests for complete auth journey
- [ ] Go unit tests for authentication middleware
- [ ] Go integration tests for GraphQL resolvers
- [ ] Test database setup for integration tests
- [ ] CI/CD pipeline configured for automated testing
- [ ] Test coverage reporting for both Flutter and Go
- [ ] Property-based tests for critical financial calculations

**Technical Requirements:**
- Follow testing strategy from ADR-20250820-testing-strategy
- Achieve 90%+ test coverage for authentication flows
- Include security testing in CI/CD pipeline
- Use rapid property testing for Go where applicable

**Definition of Done:**
- All tests pass in CI/CD pipeline
- Test coverage meets requirements (90%+)
- Integration tests verify end-to-end flows
- Test documentation updated
- Testing best practices documented

### Story 8: Local Development Environment
**Story ID**: DJINN-1008  
**Effort**: 6 hours  
**Priority**: Should Have  

**As a** developer  
**I want** a streamlined local development environment  
**So that** I can efficiently develop and test the application  

**Acceptance Criteria:**
- [ ] Docker Compose setup for local development
- [ ] PostgreSQL database seeded with test data
- [ ] Firebase emulator configured for local testing
- [ ] Hot reload working for both Flutter and Go
- [ ] Environment variables documented and configured
- [ ] Make/scripts for common development tasks
- [ ] Developer documentation updated

**Technical Requirements:**
- Follow deployment architecture from ADR-20250819-deployment-architecture
- Include MinIO for local S3-compatible storage
- Configure Redis for local caching
- Support both iOS and Android development

**Definition of Done:**
- New developers can set up environment in <30 minutes
- All services start with single command
- Local environment mirrors production architecture
- Documentation is complete and accurate
- Development workflow is smooth

## Success Metrics

### Technical Metrics
- **Authentication Success Rate**: >99.5% successful OAuth sign-ins
- **Biometric Authentication Success**: >95% successful biometric prompts
- **API Response Time**: <200ms for authentication endpoints
- **Test Coverage**: >90% for authentication flows
- **Build Success Rate**: >95% CI/CD pipeline success

### User Experience Metrics
- **Onboarding Completion**: >85% complete initial authentication
- **Authentication Drop-off**: <5% users abandon during OAuth flow
- **Biometric Adoption**: >70% users enable biometric protection
- **Session Duration**: Baseline measurement for future comparison

### Business Metrics
- **User Registration**: Track successful account creation rates
- **Trust Indicators**: User feedback on security and privacy
- **Technical Foundation**: Enable development velocity for subsequent epics

## Dependencies & Risks

### Dependencies
- **Firebase Project Setup**: Must be configured before Story 3
- **Apple Developer Account**: Required for Apple Sign-In implementation
- **Google OAuth Configuration**: Required for Google Sign-In
- **PostgreSQL Database**: Required for backend API functionality

### Technical Risks & Mitigations

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
| Firebase Auth integration complexity | Medium | High | Allocate extra time, create fallback plan |
| Biometric implementation iOS/Android differences | High | Medium | Early testing on both platforms |
| GraphQL schema evolution | Medium | Medium | Use schema versioning strategy |
| PostgreSQL connection issues | Low | High | Implement robust connection pooling |

### Business Risks & Mitigations

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
| User concerns about OAuth-only auth | Medium | Medium | Clear privacy messaging, education |
| Biometric adoption resistance | Low | Low | Make biometric protection optional |
| Competitive feature pressure | High | Medium | Focus on solid foundation over features |

## Definition of Epic Complete

This epic is complete when:
1. **All 8 user stories** are completed and accepted
2. **Authentication flow works end-to-end** from Flutter to Golang API
3. **Biometric protection** functions correctly on iOS and Android
4. **Test coverage** meets 90%+ requirement for authentication flows
5. **Security audit** confirms Zero Trust implementation
6. **Welcome screen** displays authenticated user data correctly
7. **Developer environment** is documented and streamlined
8. **CI/CD pipeline** runs all tests successfully

## Next Epic Dependencies

**Epic 2: Receipt Upload & OCR Foundation** depends on:
- User authentication and profile management (Stories 1-5)
- Secure API communication (Story 6)
- Testing infrastructure (Story 7)

**Epic 3: Financial Data Import** depends on:
- Complete authentication flow
- User profile and data storage patterns
- GraphQL API patterns established

This epic establishes the security-first, privacy-focused foundation that enables Djinn's strategic positioning as the trusted AI-powered financial assistant for privacy-conscious users seeking to break free from traditional financial tracking limitations.