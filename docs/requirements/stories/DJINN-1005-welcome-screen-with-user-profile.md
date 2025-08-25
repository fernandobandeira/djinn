# User Story: DJINN-1005-welcome-screen-with-user-profile

## Metadata
- **Story ID**: DJINN-1005
- **Epic**: Epic 1 - Base Architecture & Authentication Foundation  
- **Title**: Welcome Screen with User Profile
- **Author**: Story Creator Sub-Agent
- **Created**: 2025-08-25
- **Modified**: 2025-08-25 (Enhanced with improved dependency mapping, performance specs, and expanded testing)
- **Status**: Draft
- **Priority**: Must Have
- **Effort Estimate**: 4 hours
- **Sprint**: TBD
- **Assignee**: TBD

## User Story

**As a** newly authenticated user  
**I want** to see a welcome screen with my profile information from the database  
**So that** I know the authentication worked and can see my account details

## Business Value
- Confirms authentication success through visual feedback
- Establishes user confidence in the system by displaying their actual account data
- Provides seamless transition from authentication to main application experience
- Validates the complete authentication-to-data-access workflow

## Acceptance Criteria

### Functional Requirements
- [ ] **GraphQL Integration**: GraphQL query implemented to fetch user by Firebase UID
- [ ] **Data Source**: Welcome screen fetches user data from GraphQL API (not from JWT/Firebase)
- [ ] **User Name Display**: Displays user's name from database (originally from OAuth provider)
- [ ] **Profile Image**: Profile image URL from database displayed with fallback
- [ ] **Email Display**: Email address from database shown
- [ ] **Account Timestamp**: Account creation timestamp from database displayed in user-friendly format
- [ ] **Authenticated Requests**: Ferry GraphQL client makes authenticated request with Firebase token
- [ ] **Sign Out Functionality**: Sign out button functions correctly and clears authentication state
- [ ] **Loading States**: Proper loading states during GraphQL data fetch
- [ ] **Error Handling**: Error handling for GraphQL query failures with user-friendly messages
- [ ] **Offline Support**: Offline mode shows cached data from Ferry's HiveStore

### Technical Requirements
- [ ] **Ferry GraphQL Client**: Use Ferry GraphQL client for all data fetching operations
- [ ] **GraphQL Query**: Implement query: `query GetUser($firebaseUid: String!) { user(firebaseUid: $firebaseUid) { id, name, email, profileImageUrl, createdAt } }`
- [ ] **State Management**: Use Riverpod for state management following established patterns
- [ ] **UI/UX Consistency**: Follow UI/UX patterns from existing wireframes and design system
- [ ] **Offline Caching**: Ferry's normalized cache provides offline data access
- [ ] **Error Recovery**: Handle offline scenarios gracefully with cached data

### Measurable Acceptance Criteria
- **Screen Load Performance**: Welcome screen loads within 1.5 seconds on 3G network, <800ms on WiFi
- **GraphQL Query Performance**: First load <300ms, cached data <100ms, retry requests <500ms
- **Memory Usage**: Screen memory footprint <25MB, total app memory increase <15MB
- **Offline Support**: Works for 7+ days with cached data, auto-sync within 2 seconds of connectivity
- **Error Recovery**: Network errors resolve within 2 seconds, authentication errors <1 second
- **Battery Impact**: Screen usage <5% additional battery drain per hour
- **Accessibility**: Screen reader compatibility 100%, high contrast support, 200% text scaling

### Validation Methods
- Unit tests for GraphQL query implementation
- Widget tests for UI components and loading states
- Integration tests for authentication flow to welcome screen
- Manual testing for offline scenarios
- Accessibility testing with screen readers

## Tasks and Subtasks

### Frontend Development (Flutter)
- [ ] **Welcome Screen UI Implementation** (1.5 hours)
  - [ ] Create `WelcomeScreen` widget with user profile layout
  - [ ] Implement profile image display with circular avatar
  - [ ] Add user name, email, and account creation date display
  - [ ] Style components according to design system
  - [ ] Add sign out button with confirmation dialog
  - **Dependencies**: UI components from DJINN-1001, Authentication state from DJINN-1003

- [ ] **GraphQL Integration** (1.5 hours)
  - [ ] Generate Ferry GraphQL client code for GetUser query
  - [ ] Implement user data fetching service using Ferry
  - [ ] Add authentication token injection for GraphQL requests
  - [ ] Handle GraphQL response parsing and error states
  - **Dependencies**: GraphQL schema from DJINN-1002, Firebase auth from DJINN-1003

- [ ] **State Management & Navigation** (0.5 hours)
  - [ ] Create Riverpod providers for user profile state
  - [ ] Implement navigation from authentication success to welcome screen
  - [ ] Add sign out functionality with state cleanup
  - **Dependencies**: Authentication flow from DJINN-1003

- [ ] **Error Handling & Offline Support** (0.5 hours)
  - [ ] Implement error states for GraphQL failures
  - [ ] Add offline mode with Ferry cache integration
  - [ ] Create retry mechanisms for failed requests
  - [ ] Handle network connectivity changes
  - **Dependencies**: Ferry cache configuration from DJINN-1002

## Dev Notes

### Architecture Integration
**Authentication Flow Context** [Source: /docs/requirements/stories/DJINN-1003-firebase-auth-integration.md]:
- Welcome screen is reached after successful Firebase authentication
- Firebase ID token must be passed to GraphQL API for user identification
- Authentication state is managed via Riverpod providers

**GraphQL API Context** [Source: /docs/requirements/stories/DJINN-1002-golang-graphql-api-foundation.md]:
- Backend provides GraphQL endpoint at `/graphql`
- Authentication middleware validates Firebase ID tokens
- User lookup is performed by Firebase UID, not email

**Flutter Project Structure** [Source: /docs/requirements/stories/DJINN-1001-flutter-project-setup.md]:
- Welcome screen goes in `lib/screens/auth/welcome_screen.dart`
- GraphQL services in `lib/services/graphql/`
- Riverpod providers in `lib/providers/`

### Data Models
**User Model** [Source: /docs/architecture/data-models.md]:
```dart
class User {
  final String id;           // UUID primary key
  final String firebaseUid;  // Firebase user identifier
  final String name;         // From OAuth provider
  final String email;        // From OAuth provider
  final String? profileImageUrl; // From OAuth provider
  final DateTime createdAt;  // Account creation timestamp
  final DateTime updatedAt;  // Last update timestamp
}
```

### GraphQL Implementation

**Query Definition**:
```graphql
query GetUser($firebaseUid: String!) {
  user(firebaseUid: $firebaseUid) {
    id
    name
    email
    profileImageUrl
    createdAt
  }
}
```

**Ferry Client Configuration** [Source: /docs/architecture/frontend-architecture.md]:
```dart
// lib/services/graphql/ferry_client.dart
final ferryClient = FerryClient(
  link: HttpLink(
    'http://localhost:8080/graphql',
    defaultHeaders: {
      'Authorization': 'Bearer ${await getFirebaseToken()}',
    },
  ),
  cache: Cache(store: HiveStore()),
);
```

### Component Structure
**Welcome Screen Layout**:
```dart
// lib/screens/auth/welcome_screen.dart
class WelcomeScreen extends ConsumerWidget {
  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final userAsync = ref.watch(currentUserProvider);
    
    return Scaffold(
      appBar: AppBar(
        title: Text('Welcome to Djinn'),
        actions: [
          IconButton(
            icon: Icon(Icons.logout),
            onPressed: () => ref.read(authProvider.notifier).signOut(),
          ),
        ],
      ),
      body: userAsync.when(
        data: (user) => _buildWelcomeContent(user),
        loading: () => _buildLoadingState(),
        error: (error, stack) => _buildErrorState(error),
      ),
    );
  }
}
```

### State Management
**Riverpod Providers** [Source: /docs/architecture/frontend-architecture.md]:
```dart
// lib/providers/user_provider.dart
final currentUserProvider = FutureProvider<User>((ref) async {
  final firebaseUser = ref.watch(firebaseAuthProvider);
  if (firebaseUser == null) throw Exception('Not authenticated');
  
  final result = await ref.read(ferryClientProvider).request(
    GetUserReq((b) => b..vars.firebaseUid = firebaseUser.uid),
  );
  
  return result.data?.user ?? throw Exception('User not found');
});
```

### Navigation Integration
**Route Configuration** [Source: /docs/architecture/frontend-architecture.md]:
- Welcome screen accessible at `/welcome` route
- Automatic navigation after successful biometric authentication
- Deep link support for direct welcome screen access

### File Locations
```
lib/
├── screens/auth/
│   ├── welcome_screen.dart          # Main welcome screen widget
│   └── welcome_screen_test.dart     # Widget tests
├── services/graphql/
│   ├── queries/get_user.graphql     # GraphQL query definition
│   └── generated/                   # Ferry generated code
├── providers/
│   └── user_provider.dart           # User state management
└── models/
    └── user.dart                    # User data model
```

### Testing Requirements
**Unit Tests**:
- GraphQL query execution and response parsing
- User provider state transitions
- Error handling for network failures

**Widget Tests**:
- Welcome screen rendering with user data
- Loading state display during data fetch
- Error state handling and retry mechanisms
- Sign out button functionality

**Integration Tests**:
- Complete authentication to welcome screen flow
- Offline functionality with cached data
- Network recovery scenarios

### Dependencies

#### Critical Path Dependencies (Must Complete Before)
1. **DJINN-1001** (Flutter Base Architecture Setup)
   - **Required Components**: Riverpod providers, go_router configuration, app theming
   - **Specific Assets Needed**: Base widget structure, error handling patterns
   - **Integration Points**: Navigation framework, state management patterns
   - **Estimated Lead Time**: Must be 100% complete

2. **DJINN-1002** (GraphQL API Foundation)
   - **Required Components**: User schema, authentication middleware, GetUser resolver
   - **Specific Assets Needed**: GraphQL endpoint `/graphql`, User table with Firebase UID column
   - **Integration Points**: JWT token validation, user lookup by Firebase UID
   - **Estimated Lead Time**: Must be 95% complete with working user queries

3. **DJINN-1003** (Firebase Authentication Integration)
   - **Required Components**: Firebase token generation, authentication state management
   - **Specific Assets Needed**: Firebase project config, OAuth provider setup, token refresh logic
   - **Integration Points**: Authentication state providers, token injection for GraphQL
   - **Estimated Lead Time**: Must be 100% complete with working authentication flow

#### Optional Dependencies (Can Work in Parallel)
4. **DJINN-1004** (Biometric Protection Implementation)
   - **Integration Points**: Post-biometric navigation to welcome screen
   - **Fallback**: Can navigate from basic auth success if biometric incomplete
   - **Estimated Lead Time**: Can be 70% complete, final integration later

#### Dependency Sequence & Integration Points
```
DJINN-1001 (100%) → DJINN-1003 (100%) → DJINN-1002 (95%) → DJINN-1005
                                    ↓                            ↑
                              DJINN-1004 (70%) ---------------→
```

#### External Dependencies
- **Firebase Project**: Google OAuth, Apple OAuth configured and tested
- **PostgreSQL Database**: User table schema deployed, test data populated
- **Network Infrastructure**: GraphQL endpoint accessible from mobile app
- **Design System**: UI components documented and available

### Technical Constraints
**Security** [Source: /docs/architecture/backend-architecture.md]:
- All GraphQL requests must include valid Firebase ID token
- User data access restricted to authenticated user's own data
- No sensitive data stored in Ferry cache

**Performance** [Source: /docs/architecture/frontend-architecture.md]:
- Ferry cache enables instant offline access
- GraphQL query should complete within 500ms
- Image loading with lazy loading and caching

**Accessibility**:
- Screen reader support for all user data display
- High contrast mode support
- Large text scaling support

## Testing Requirements

### Test Scenarios

#### Happy Path Testing
1. **Successful User Data Loading**: 
   - Verify GraphQL query executes with correct Firebase UID
   - Confirm all user data fields populate correctly
   - Test profile image loading with fallback to default avatar
   - Validate proper timestamp formatting for account creation date
   - Test data refresh on pull-to-refresh gesture

#### State Management Testing
2. **Loading States**:
   - Verify spinner displays during initial GraphQL request
   - Test skeleton loading for profile components
   - Test loading timeout scenarios (>10 seconds)
   - Confirm smooth transition from loading to data display
   - Test concurrent loading states (image + data)

#### Error Handling & Edge Cases
3. **Network & API Failures**:
   - **GraphQL Server Errors**: Test 500, 503, timeout responses
   - **Authentication Failures**: Test 401 Unauthorized responses
   - **Network Connectivity**: Test airplane mode, poor signal scenarios
   - **Malformed Data**: Test missing fields, null values, invalid data types
   - **Rate Limiting**: Test 429 Too Many Requests responses
   - **User-Friendly Messages**: Verify appropriate error messaging for each scenario
   - **Retry Logic**: Test automatic retry mechanisms and manual retry buttons

4. **Authentication Edge Cases**:
   - **Token Expiry**: Test Firebase token expiration during screen display
   - **Token Refresh**: Test automatic token refresh scenarios
   - **Invalid User**: Test user not found in database scenarios
   - **Concurrent Sessions**: Test behavior with multiple app instances
   - **Account Deletion**: Test behavior when user account deleted server-side

#### Offline & Connectivity Testing
5. **Offline Functionality**:
   - **Initial Offline Load**: Test welcome screen with cached data from previous session
   - **Network Loss**: Test transition from online to offline mid-session
   - **Stale Cache**: Test behavior with week-old cached data
   - **Partial Sync**: Test scenarios where only some data syncs successfully
   - **Cache Invalidation**: Test cache clearing and repopulation
   - **Offline Indicator**: Test offline status display and user messaging
   - **Data Sync**: Test background sync when network returns

#### User Interaction Testing  
6. **Sign Out Flow**:
   - **Standard Sign Out**: Test sign out button functionality
   - **Confirmation Dialog**: Test sign out confirmation and cancellation
   - **State Cleanup**: Verify complete authentication state clearing
   - **Navigation**: Confirm proper navigation back to login screen
   - **Data Clearing**: Verify user data cleared from UI state
   - **Cache Management**: Test selective cache clearing for security

#### Security Testing
7. **Security & Privacy**:
   - **Token Validation**: Test with invalid, expired, and malicious tokens
   - **Data Access**: Verify user can only access their own data
   - **Memory Security**: Test that sensitive data isn't stored in memory dumps
   - **Screenshot Security**: Test app behavior in background/recent apps
   - **Deep Link Security**: Test welcome screen access without proper authentication

#### Device & Platform Testing
8. **Device Compatibility**:
   - **Screen Sizes**: Test on phone, tablet, foldable device layouts
   - **Orientation**: Test portrait and landscape orientations
   - **Dark Mode**: Test UI appearance in light and dark themes
   - **Accessibility**: Test with VoiceOver/TalkBack, high contrast, large text
   - **System Fonts**: Test with various system font sizes and weights
   - **Platform Differences**: Test iOS vs Android specific behaviors

#### Performance Edge Cases
9. **Performance Under Stress**:
   - **Low Memory**: Test behavior when device memory is constrained
   - **Background Processing**: Test welcome screen when returning from background
   - **Concurrent Operations**: Test during file downloads, camera usage, etc.
   - **Battery Saver Mode**: Test functionality with power saving enabled
   - **Network Switching**: Test behavior during WiFi to cellular transitions

### Coverage Requirements
- Unit test coverage: >90%
- Widget test coverage: >85%
- Integration test coverage: >80%

### Performance Testing

#### Performance Benchmarks
| Metric | Target | Acceptable | Critical |
|--------|--------|------------|----------|
| First Screen Load (WiFi) | <800ms | <1.2s | <2s |
| First Screen Load (3G) | <1.5s | <2.5s | <4s |
| GraphQL Query (Cold) | <300ms | <500ms | <1s |
| GraphQL Query (Cached) | <100ms | <200ms | <300ms |
| Screen Transition | <200ms | <300ms | <500ms |
| Memory Usage | <15MB | <25MB | <40MB |
| Battery Drain/Hour | <3% | <5% | <8% |
| Offline Load Time | <500ms | <800ms | <1.2s |

#### Performance Test Scenarios
1. **Network Condition Testing**
   - Test on 2G, 3G, 4G, and WiFi networks
   - Measure load times under each condition
   - Validate graceful degradation

2. **Device Performance Testing**
   - Test on low-end devices (2GB RAM, older processors)
   - Measure memory usage patterns
   - Validate smooth animations and transitions

3. **Data Volume Testing**
   - Test with various profile image sizes
   - Validate caching behavior
   - Measure storage usage over time

#### Performance Monitoring
- Real-time performance tracking with Firebase Performance Monitoring
- Custom metrics for GraphQL response times
- Memory leak detection during automated testing
- Battery usage profiling on target devices

## Definition of Done

### Code Quality
- [ ] Code follows Dart/Flutter style guidelines
- [ ] All linting rules pass
- [ ] No code duplication or technical debt
- [ ] Proper error handling implemented
- [ ] Security best practices followed

### Testing
- [ ] Unit tests written and passing (>90% coverage)
- [ ] Widget tests implemented and passing (>85% coverage)
- [ ] Integration tests cover main user flows
- [ ] Manual testing completed for all acceptance criteria
- [ ] Accessibility testing completed
- [ ] Performance benchmarks met

### Documentation
- [ ] Code documented with clear comments
- [ ] GraphQL schema updated if needed
- [ ] README updated with welcome screen information
- [ ] API documentation reflects any changes

### Deployment
- [ ] Code committed to feature branch
- [ ] Pull request created and reviewed
- [ ] CI/CD pipeline passes all checks
- [ ] Feature branch merged to development
- [ ] Deployment to staging environment successful

### Validation
- [ ] Product owner approval on UI/UX
- [ ] QA team sign-off on functionality
- [ ] Security review completed for authentication flow
- [ ] Performance metrics validated
- [ ] Accessibility standards verified (WCAG 2.1 AA)

## Dependencies
- **Blocking**: DJINN-1003 (Firebase Auth Integration) - Required for authentication state
- **Blocking**: DJINN-1002 (GraphQL API Foundation) - Required for user data queries  
- **Related**: DJINN-1004 (Biometric Protection) - Provides navigation context
- **External**: Firebase project configuration must be complete
- **External**: PostgreSQL database must be accessible with user schema

## Risks and Mitigation

### Technical Risks
1. **GraphQL Query Performance**
   - *Risk*: Slow query response affecting user experience
   - *Mitigation*: Implement query caching, add loading indicators, optimize database queries
   - *Probability*: Medium | *Impact*: Medium

2. **Authentication Token Expiry**
   - *Risk*: Firebase token expires during welcome screen display
   - *Mitigation*: Implement token refresh mechanism, handle 401 responses gracefully
   - *Probability*: High | *Impact*: High

3. **Offline Data Staleness**
   - *Risk*: Cached user data becomes outdated
   - *Mitigation*: Cache expiration policies, background sync when online
   - *Probability*: Medium | *Impact*: Low

### Integration Risks
1. **Ferry GraphQL Client Issues**
   - *Risk*: Ferry client configuration or code generation problems
   - *Mitigation*: Test Ferry integration early, have fallback HTTP client ready
   - *Probability*: Low | *Impact*: High

## Rollback Strategy

### Immediate Rollback (< 1 hour)
1. Revert feature branch merge if major issues discovered
2. Disable welcome screen route, redirect to fallback screen
3. Switch to backup authentication flow without GraphQL dependency

### Partial Rollback (1-4 hours)
1. Disable GraphQL integration, use Firebase user data temporarily
2. Remove offline functionality if Ferry cache causes issues
3. Simplify UI to basic user information display

### Full Rollback (4-8 hours)
1. Revert to previous authentication flow
2. Remove all welcome screen components
3. Restore previous post-authentication navigation

### Rollback Validation
- Verify authentication flow works without welcome screen
- Test that user can still access main application features
- Confirm no data corruption or authentication state issues

## Change Log
| Date | Author | Change Description |
|------|--------|-------------------|
| 2025-08-25 | Story Creator Sub-Agent | Initial story creation |
| 2025-08-25 | Story Creator Sub-Agent | Enhanced dependency mapping, performance specifications, and expanded testing scenarios |

## Dev Agent Record
- **Generated By**: Story Creator Sub-Agent
- **Template Version**: enhanced-story-template.md v1.0
- **Epic Reference**: Epic 1 - Base Architecture & Authentication Foundation
- **Story Dependencies Analyzed**: DJINN-1001, DJINN-1002, DJINN-1003, DJINN-1004
- **Architecture Sources**: frontend-architecture.md, backend-architecture.md, data-models.md
- **Total Context Files Reviewed**: 15

## QA Results
*Placeholder for QA team validation results*
- [ ] Functional testing completed
- [ ] Performance benchmarks validated  
- [ ] Security review passed
- [ ] Accessibility compliance verified
- [ ] User acceptance testing completed

## File List
### Modified Files
- `lib/screens/auth/welcome_screen.dart` - New welcome screen implementation
- `lib/services/graphql/queries/get_user.graphql` - GraphQL query definition
- `lib/providers/user_provider.dart` - User state management provider
- `lib/models/user.dart` - User data model updates
- `lib/routes/app_router.dart` - Welcome screen route configuration

### Test Files
- `test/screens/auth/welcome_screen_test.dart` - Widget tests
- `test/services/graphql/user_service_test.dart` - GraphQL service tests
- `test/providers/user_provider_test.dart` - Provider unit tests
- `integration_test/auth_flow_test.dart` - End-to-end authentication flow