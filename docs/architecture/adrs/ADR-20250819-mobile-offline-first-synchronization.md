# ADR-20250819: Mobile Offline-First Data Synchronization Architecture

## Status
Accepted

## Context
Djinn's Flutter mobile application requires robust offline-first data synchronization to ensure users can manage their financial data without constant internet connectivity. The application handles sensitive financial data with complex relationships including institutions, accounts, transactions, receipts, and credit card liabilities.

### Current Technology Stack
- **Mobile**: Flutter with Ferry's HiveStore for GraphQL cache persistence
- **Backend**: Go server with PostgreSQL database
- **API**: GraphQL with Ferry client (includes normalized cache)
- **State Management**: Riverpod for reactive state management
- **Local Database**: Drift (SQLite wrapper for Flutter)
  - Chosen for: Type-safe queries, blob support, migrations
  - Stores: Transactions, accounts, sync queue, receipt images
  - Persists across app restarts and device reboots
- **Cache Layer**: Ferry HiveStore (GraphQL normalized cache) + Drift (structured data)
- **Data Types**: 
  - Hierarchical: Institutions → Accounts → Transactions
  - Binary: Receipt images requiring OCR processing (stored in Drift blobs)
  - Linked: Transfer transactions between accounts
  - Read-only: Credit card liabilities from Plaid

### Key Requirements
- Users must be able to create, edit, and delete transactions offline
- Receipt upload and storage offline with server-side OCR processing
- Institution hierarchy maintained (system institutions cached locally)
- Transfer transactions must sync atomically (both sides together)
- Credit card statement data sync separate from transactions
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
- Receipt images can be large (2-5MB each)
- OCR processing must happen server-side
- Institution data is shared across all users
- Credit card liability data is read-only from Plaid

## Decision
We will implement a **hybrid Event Sourcing with Last-Write-Wins (LWW) conflict resolution** architecture for offline-first synchronization, with specialized handling for our complex data model.

### Local Storage Decision: Drift Database

We chose **Drift** (formerly Moor) as our local SQLite database for Flutter because:

1. **Blob Support**: Native support for storing receipt images as binary data
2. **Type Safety**: Generates type-safe Dart code from SQL schemas
3. **Persistence**: SQLite storage survives app restarts, ensuring offline changes are never lost
4. **Migrations**: Built-in migration system for schema updates
5. **Query Builder**: Fluent API for complex queries without raw SQL
6. **Reactive Streams**: Integrates well with Riverpod for reactive UI updates

#### Drift vs Alternatives:
- **Hive alone**: No blob support, limited query capabilities
- **sqflite**: No type safety, manual SQL writing
- **Isar**: Less mature, smaller community
- **ObjectBox**: Licensing concerns, overkill for our needs

### Core Architecture Components

#### 1. Hierarchical Data Caching Strategy

##### System Institutions (Cached from API)
```dart
// Pull system institutions on app start/refresh
class InstitutionCache {
  // Fetch from /api/institutions/system endpoint
  // Cache in Hive with 7-day TTL
  // Includes: name, logo_url, routing_numbers, colors
  static const Duration cacheTTL = Duration(days: 7);
  
  Future<void> refreshSystemInstitutions() async {
    final institutions = await api.getSystemInstitutions();
    await hiveBox.put('system_institutions', institutions);
    await hiveBox.put('institutions_cached_at', DateTime.now());
  }
}
```

##### Data Sync Hierarchy
1. **System Institutions** → Cached, read-only, refreshed weekly
2. **User Institutions** → Synced with user data
3. **Accounts** → Depends on institutions being synced first
4. **Transactions** → Depends on accounts
5. **Receipts** → Can sync independently, matched server-side

#### 2. Local Event Store with Entity-Specific Handling
- Store all user actions as immutable events with timestamps
- **Persistence**: Drift SQLite database persists across app restarts
- **Data retention**: All offline changes survive app closure, phone restarts
- Events include: 
  - Standard: transaction_created, account_updated, category_modified
  - Special: transfer_initiated (groups both sides), receipt_uploaded
  - Manual: credit_statement_updated, manual_transaction_added
- Each event contains full state snapshot for reconstruction
- Binary data (receipts) stored in Drift blob columns

#### App Restart Behavior
```dart
// On app startup:
class StartupSync {
  Future<void> initializeOfflineData() async {
    // 1. Load all persisted data from Drift
    final localData = await db.getAllUserData();
    
    // 2. Load pending sync queue
    final pendingOps = await db.syncQueue
        .where((q) => q.sync_status.equals('pending'))
        .get();
    
    // 3. Display cached data immediately (offline-first)
    await updateUIWithLocalData(localData);
    
    // 4. If online, sync in background
    if (await connectivity.hasInternet()) {
      unawaited(processSyncQueue(pendingOps));  // Non-blocking
    }
    
    // 5. User can continue working offline
    // All changes queued for future sync
  }
}
```

#### 3. Receipt Handling Strategy
```dart
class ReceiptSyncManager {
  // Offline: Store image locally
  Future<void> uploadReceipt(Uint8List imageData) async {
    final receiptId = Uuid().v4();
    
    // Store in Drift blob column
    await db.receipts.insert(
      ReceiptCompanion(
        id: Value(receiptId),
        imageData: Value(imageData),
        status: Value('pending_upload'),
        createdAt: Value(DateTime.now()),
      ),
    );
    
    // Queue for upload
    await syncQueue.add(SyncTask(
      type: 'receipt_upload',
      entityId: receiptId,
      priority: 50, // Lower priority than transactions
    ));
  }
  
  // Online: Upload and process
  Future<void> processReceiptQueue() async {
    // 1. Upload image to cloud storage
    // 2. Trigger OCR processing
    // 3. Receive extracted data
    // 4. Server performs transaction matching
    // 5. Return match suggestions to client
    // 6. Handle based on confidence
    await handleMatchResult(matchResult);
  }
  
  // User confirmation for uncertain matches
  Future<void> handleMatchResult(ReceiptMatchResult result) async {
    if (result.confidence >= 0.80) {
      // Auto-accept high confidence matches
      await acceptMatch(result.receiptId, result.transactionId);
      
    } else if (result.confidence >= 0.50) {
      // Show confirmation dialog for medium confidence
      await showMatchConfirmation(
        receipt: result.receipt,
        suggestedMatches: result.suggestions,
        onConfirm: (transactionId) => acceptMatch(result.receiptId, transactionId),
        onReject: () => markAsUnmatched(result.receiptId),
      );
      
    } else {
      // Low confidence - let user manually select from list
      await showManualMatchSelection(
        receipt: result.receipt,
        allTransactions: result.nearbyTransactions,
        onSelect: (transactionId) => acceptMatch(result.receiptId, transactionId),
        onSkip: () => keepReceiptUnmatched(result.receiptId),
      );
    }
  }
}
```

#### 4. Transfer Transaction Atomicity
```dart
class TransferSyncManager {
  // Ensure both sides of transfer sync together
  Future<void> createTransfer({
    required Account fromAccount,
    required Account toAccount,
    required int amountMinor,
  }) async {
    final groupId = Uuid().v4();
    
    // Create both transactions with same group ID
    final fromTransaction = Transaction(
      id: Uuid().v4(),
      accountId: fromAccount.id,
      amount: -amountMinor,
      transferGroupId: groupId,
      type: 'transfer',
    );
    
    final toTransaction = Transaction(
      id: Uuid().v4(),
      accountId: toAccount.id,
      amount: amountMinor,
      transferGroupId: groupId,
      type: 'transfer',
    );
    
    // Queue as atomic unit
    await syncQueue.addAtomic([
      SyncTask(entityId: fromTransaction.id, groupId: groupId),
      SyncTask(entityId: toTransaction.id, groupId: groupId),
    ]);
  }
}
```

#### 5. Credit Card Liability Sync
```dart
class CreditCardSyncManager {
  // Two-level sync strategy
  
  // Level 1: Transaction sync (normal)
  Future<void> syncTransactions(String accountId) async {
    // Standard transaction sync
  }
  
  // Level 2: Liability data (read-only from Plaid)
  Future<void> syncLiabilityData(String accountId) async {
    // Pull from Plaid Liabilities API
    final liability = await plaid.getLiabilities(accountId);
    
    // Store locally (read-only)
    await db.creditCardLiabilities.upsert(
      minimumPayment: liability.minimumPayment,
      dueDate: liability.nextPaymentDueDate,
      statementBalance: liability.lastStatementBalance,
      apr: liability.aprPurchase,
    );
  }
  
  // Manual statement updates (queued for sync)
  Future<void> updateManualStatement({
    required String accountId,
    required int minimumPayment,
    required DateTime dueDate,
  }) async {
    await syncQueue.add(SyncTask(
      type: 'manual_statement_update',
      entityId: accountId,
      payload: {
        'minimumPayment': minimumPayment,
        'dueDate': dueDate.toIso8601String(),
      },
    ));
  }
}
```

#### 6. Enhanced Synchronization Engine
- **Ferry GraphQL Client** with custom offline layer for mutations
- **Priority-based sync queue** with entity dependencies
- **Background sync service** using WorkManager with battery optimization
- **Conflict resolution**:
  - Standard entities: Last-Write-Wins
  - Transfers: Atomic sync or rollback
  - Receipts: Server-side matching with user confirmation
- **Batch processing** with intelligent grouping

### Enhanced Sync Queue Structure
```sql
CREATE TABLE sync_queue (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    device_id TEXT NOT NULL,
    operation_type TEXT NOT NULL,    -- create, update, delete, upload
    entity_type TEXT NOT NULL,       -- transaction, account, receipt, etc.
    entity_id UUID NOT NULL,
    -- New fields for complex data model
    transaction_group_id UUID,       -- For transfer pairs
    dependency_ids UUID[],           -- For hierarchical dependencies
    priority INTEGER DEFAULT 100,    -- Higher = sync first
    binary_ref TEXT,                 -- Reference to blob storage for receipts
    payload JSONB NOT NULL,
    sync_status TEXT DEFAULT 'pending',
    retry_count INTEGER DEFAULT 0,
    max_retries INTEGER DEFAULT 3,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    scheduled_at TIMESTAMPTZ,         -- For delayed sync
    
    CONSTRAINT valid_status CHECK (sync_status IN ('pending', 'processing', 'completed', 'failed'))
);
```

### Synchronization Flow with Priority
1. **Offline Operations**: 
   - Store events with dependencies and priorities
   - Receipts stored in Drift blob, reference in queue
   - Transfers grouped with transaction_group_id

2. **Sync Priority Order**:
   ```
   Priority 200: System institutions refresh (if stale)
   Priority 150: User institutions
   Priority 140: Accounts (depends on institutions)
   Priority 120: Credit card liabilities
   Priority 100: Transactions (standard)
   Priority 90:  Transfers (atomic groups)
   Priority 50:  Receipts (can be delayed)
   Priority 30:  Categories/Budgets
   ```

3. **Background Sync**:
   - WorkManager runs every 15 minutes on WiFi
   - Immediate sync on app foreground if pending items
   - Battery-aware: Reduce frequency on low battery

4. **Conflict Resolution Matrix**:
   | Entity Type | Resolution Strategy | User Notification |
   |------------|-------------------|-------------------|
   | Transaction | Last-Write-Wins | No |
   | Transfer | Atomic (both or none) | Yes if fails |
   | Receipt Match (≥80%) | Auto-accept | No (silent match) |
   | Receipt Match (50-79%) | Suggest match | Yes (confirm dialog) |
   | Receipt Match (<50%) | Manual selection | Yes (pick from list) |
   | Account | Server authoritative | Yes if conflict |
   | Credit Liability | Read-only (no conflict) | No |

5. **State Reconciliation**: 
   - Update Ferry cache with server response
   - Notify Riverpod providers of changes
   - Clean up completed sync queue items

### Data Consistency Strategy
- **Dependency tracking** ensures correct sync order
- **Atomic transfers** maintain double-entry integrity
- **Optimistic UI** with rollback on sync failure
- **Server validation** for business rules
- **Audit trail** via event sourcing

### Data Retention & Cleanup Policies

#### Persistent Storage Guarantees
- **User data**: Never deleted without explicit user action
- **Sync queue**: Retained until successfully synced or max retries exceeded
- **Receipt images**: Compressed after upload confirmation
- **Failed syncs**: Kept for manual retry with user notification

#### Storage Management
```dart
class StorageManager {
  // Clean up after successful sync
  Future<void> postSyncCleanup() async {
    // 1. Remove completed sync queue items
    await db.syncQueue
        .deleteWhere((q) => q.sync_status.equals('completed'));
    
    // 2. Compress uploaded receipt images
    final uploadedReceipts = await db.receipts
        .where((r) => r.uploadStatus.equals('completed'))
        .get();
    
    for (final receipt in uploadedReceipts) {
      // Keep compressed thumbnail locally
      await compressImage(receipt.imageData);
    }
    
    // 3. Archive old events (keep 90 days)
    final cutoffDate = DateTime.now().subtract(Duration(days: 90));
    await db.events
        .deleteWhere((e) => e.createdAt.isBefore(cutoffDate)
            .and(e.synced.equals(true)));
  }
  
  // Handle storage pressure
  Future<void> handleLowStorage() async {
    // 1. Clear uploaded receipt full-size images (keep thumbnails)
    // 2. Compact event store
    // 3. Clear old cache entries
    // 4. Notify user if critical
  }
}
```

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

### Phase 1: Foundation & Data Model (2 weeks)
- Set up Drift database (chosen for blob support, type safety, and migrations)
- Implement sync queue with priority and dependencies
- Create institution caching layer with API integration
- Build event store for user actions
- Unit tests for data persistence

#### Why Drift over alternatives:
- **Blob support**: Native support for storing receipt images locally
- **Type safety**: Generated Dart code with compile-time checks
- **Migrations**: Structured schema migrations for app updates
- **Performance**: SQLite backend with optimized queries
- **Integration**: Works well with Ferry's HiveStore for GraphQL cache

### Phase 2: Entity-Specific Sync Managers (3 weeks)
- **InstitutionSyncManager**: Cache system institutions, sync user custom
- **TransferSyncManager**: Atomic transfer handling with grouping
- **ReceiptSyncManager**: Binary upload queue, OCR result handling
- **CreditCardSyncManager**: Dual-level sync (transactions + liabilities)
- Integration with Ferry GraphQL client

### Phase 3: Synchronization Engine (2 weeks)
- Priority-based queue processing
- Dependency resolution for hierarchical data
- Background sync with WorkManager
- Conflict detection and resolution strategies
- Network monitoring and retry logic

### Phase 4: State Management & UI (2 weeks)
- Riverpod providers for each entity type
- Optimistic updates with rollback
- Receipt match suggestion UI
- Sync status indicators
- Conflict notification system

### Phase 5: Performance & Reliability (1 week)
- Implement Isolates for heavy processing (OCR responses)
- Queue compaction and cleanup
- Battery optimization strategies
- Sync metrics and monitoring
- Load testing with large datasets

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