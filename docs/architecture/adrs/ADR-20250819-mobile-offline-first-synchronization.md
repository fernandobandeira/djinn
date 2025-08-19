# ADR-20250819: Mobile Offline-First Data Synchronization Architecture

## Status
Proposed

## Context
Djinn's Flutter mobile application requires robust offline-first data synchronization to ensure users can manage their financial data without constant internet connectivity. The application handles sensitive financial data including transactions, categories, budgets, and user preferences that must maintain consistency and provide complete audit trails.

### Current Technology Stack
- **Mobile**: Flutter with Ferry's HiveStore for offline persistence
- **Backend**: Go server with PostgreSQL database
- **API**: GraphQL with Ferry client (includes normalized cache)
- **State Management**: Riverpod for reactive state management
- **Data Types**: Financial transactions, budget categories, user preferences, receipt data

### Key Requirements
- Users must be able to create, edit, and delete transactions offline
- Category management and budget updates must work offline
- Data must eventually synchronize with backend when connectivity returns
- Conflict resolution for concurrent edits across devices
- Complete audit trail for all financial data changes
- Background synchronization with minimal battery impact
- Data consistency guarantees for financial accuracy

### Constraints
- Flutter background processing limitations on iOS/Android
- Ferry GraphQL client offline capabilities and limitations
- Financial data requires ACID properties and cannot tolerate data loss
- Mobile device storage and performance constraints
- Network connectivity can be intermittent or unreliable

## Decision
We will implement a **hybrid Event Sourcing with Last-Write-Wins (LWW) conflict resolution** architecture for offline-first synchronization:

### Core Architecture Components

#### 1. Local Event Store (Ferry HiveStore)
- Store all user actions as immutable events with timestamps using Ferry's normalized cache
- Events include: transaction_created, transaction_updated, category_modified, budget_adjusted
- Each event contains full state snapshot for simple reconstruction
- Local sequence numbers for ordering and deduplication
- HiveStore provides persistent offline storage without additional database complexity

#### 2. Synchronization Engine
- **Ferry GraphQL Client** with custom offline layer for caching mutations
- **Background sync service** using Flutter's WorkManager for periodic synchronization
- **Conflict resolution** using Last-Write-Wins based on server timestamps
- **Batch mutation processing** to optimize network usage and reduce API calls

#### 3. State Management Integration
- **Riverpod providers** expose synchronized state to UI components
- **Optimistic updates** for immediate UI feedback during offline operations
- **Conflict notification system** to inform users of resolved conflicts
- **Local-first state** with eventual consistency guarantees

### Synchronization Flow
1. **Offline Operations**: Store events locally with client-generated IDs and timestamps
2. **Connectivity Detection**: Monitor network status using Flutter connectivity plugins
3. **Background Sync**: Periodic synchronization using WorkManager (every 15 minutes when online)
4. **Conflict Detection**: Compare local and server timestamps for overlapping changes
5. **Conflict Resolution**: Apply Last-Write-Wins with user notification for significant conflicts
6. **State Reconciliation**: Update local state and notify UI components via Riverpod

### Data Consistency Strategy
- **Optimistic concurrency control** for user experience
- **Server-side validation** for data integrity
- **Compensation actions** for failed synchronization attempts
- **Audit event preservation** for compliance and debugging

## Consequences

### Positive
- **Seamless offline experience**: Users can perform all core operations without connectivity
- **Simple conflict resolution**: LWW is predictable and easy to understand for users
- **Leverages existing tech stack**: Builds on Ferry's HiveStore and Riverpod investments
- **Background synchronization**: Automatic data sync without user intervention
- **Complete audit trail**: Event sourcing provides full history of all changes
- **Performance optimized**: Batch operations reduce API calls and improve battery life

### Negative
- **Storage overhead**: Event store requires more local storage than simple state snapshots
- **Complex implementation**: Event sourcing adds complexity to data access patterns
- **Limited conflict resolution**: LWW may not be optimal for all types of financial data conflicts
- **Background sync limitations**: Platform restrictions may delay synchronization
- **Eventual consistency**: Users may see temporary inconsistencies across devices

### Risks
- **Data loss potential**: Aggressive conflict resolution could lose user changes
- **Performance degradation**: Large event stores may impact app startup and query performance
- **Synchronization failures**: Network issues could lead to prolonged data divergence
- **Battery usage**: Background sync may impact device battery life
- **Platform compliance**: App store policies may restrict background processing capabilities

## Alternatives Considered

### Option A: Conflict-Free Replicated Data Types (CRDTs)
- **Description**: Use mathematical data structures that can be merged automatically without conflicts
- **Pros**: True automatic conflict resolution, guaranteed eventual consistency, no data loss
- **Cons**: Complex implementation, limited CRDT libraries for Flutter/Dart, poor fit for financial data structures
- **Reason for not choosing**: Financial transactions don't map well to CRDT semantics, implementation complexity too high

### Option B: Full Event Sourcing with Complex Conflict Resolution
- **Description**: Complete event sourcing system with semantic conflict resolution based on financial rules
- **Pros**: Sophisticated conflict handling, complete audit trail, deterministic state reconstruction
- **Cons**: Very complex implementation, requires domain-specific conflict resolution logic, poor performance
- **Reason for not choosing**: Implementation complexity would significantly delay MVP, performance concerns for mobile

### Option C: Simple State Synchronization with Manual Conflict Resolution
- **Description**: Sync full object states and present conflicts to users for manual resolution
- **Pros**: Simple implementation, complete user control, predictable behavior
- **Cons**: Poor user experience, frequent conflict dialogs, requires user decision-making
- **Reason for not choosing**: UX friction would harm adoption, especially for automated transactions

### Option D: Server-Side Queuing with Optimistic UI
- **Description**: Queue all operations server-side and only sync when online, with optimistic UI updates
- **Pros**: Simpler local storage, guaranteed server consistency, reduced client complexity
- **Cons**: Poor offline experience, potential data loss if server rejects operations, complex rollback logic
- **Reason for not choosing**: Violates offline-first requirement, poor user experience during extended offline periods

## Implementation Notes

### Phase 1: Event Store Foundation (2 weeks)
- Implement local event store using Ferry's normalized cache with HiveStore persistence
- Create event types for all financial data operations
- Build event serialization and persistence layer
- Unit tests for event store operations

### Phase 2: Synchronization Engine (3 weeks)
- Integrate Ferry client with custom offline mutation queue
- Implement background sync service with WorkManager
- Build conflict detection and resolution logic
- Create network connectivity monitoring

### Phase 3: State Management Integration (2 weeks)
- Update Riverpod providers to consume from event store
- Implement optimistic updates with rollback capabilities
- Create conflict notification UI components
- Integration testing with real network conditions

### Phase 4: Performance and Monitoring (1 week)
- Implement event store compaction for long-running usage
- Add synchronization metrics and error reporting
- Performance testing and optimization
- User acceptance testing

### Migration Strategy
1. Deploy backend changes to support event-based synchronization APIs
2. Update mobile app with feature flags to enable gradual rollout
3. Migrate existing local data to event store format
4. Enable synchronization for new users first, then existing users
5. Monitor synchronization success rates and performance metrics

### Rollback Plan
- Feature flags allow instant disabling of new synchronization logic
- Keep existing simple sync as fallback for 2 releases
- Local event store can be disabled to revert to direct state storage
- Server-side APIs maintain backward compatibility

### Testing Approach
- **Unit tests**: Event store operations, conflict resolution algorithms, state reconstruction
- **Integration tests**: Full synchronization flows with simulated network conditions
- **Load testing**: Performance with large numbers of events and conflicts
- **User testing**: Real-world usage patterns and conflict scenarios
- **Automated testing**: CI/CD pipeline with offline/online scenario testing

### Success Metrics
- **Synchronization success rate**: > 99% for normal network conditions
- **Conflict resolution accuracy**: < 1% of conflicts require manual intervention
- **Background sync battery impact**: < 5% additional battery usage per day
- **App responsiveness**: UI operations complete in < 100ms during sync
- **Data consistency**: Zero financial data loss incidents
- **User satisfaction**: > 4.5/5 rating for offline functionality

## References
- [Ferry GraphQL Client Documentation](https://ferrygraphql.com/)
- [Ferry GraphQL Client Documentation](https://ferrygraphql.com/)
- [Hive Database Documentation](https://docs.hivedb.dev/)
- [Flutter Background Processing Guidelines](https://flutter.dev/docs/development/packages-and-plugins/background-processes)
- [Event Sourcing Patterns by Martin Fowler](https://martinfowler.com/eaaDev/EventSourcing.html)
- [Offline-First Architecture Patterns](https://offlinefirst.org/patterns/)
- **Related ADRs**: ADR-20250812-personal-finance-tech-stack.md (establishes base technology choices)

## Decision Makers
- **Author**: Archie (System Architect)
- **Reviewers**: [To be assigned - Mobile team, Backend team]
- **Approver**: [To be assigned - Technical Lead]
- **Date**: 2025-08-19

## Revision History
- 2025-08-19: Initial draft created with comprehensive analysis of sync strategies
- 2025-08-19: Added detailed implementation phases and success metrics