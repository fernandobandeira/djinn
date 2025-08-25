# DJINN-1001: Flutter Base Architecture Setup

## Metadata
- **Story ID**: DJINN-1001
- **Title**: Flutter Base Architecture Setup  
- **Epic**: DJINN-E001 - Base Architecture & Authentication Foundation
- **Status**: Draft
- **Priority**: Must Have
- **Effort Estimate**: 6 hours
- **Author**: System Story Creator
- **Created**: 2025-08-25
- **Last Updated**: 2025-08-25 (v2.0 - Enhanced)
- **Tags**: architecture, flutter, foundation, setup

## User Story
As a developer, I want a properly structured Flutter application foundation, so that all future features follow consistent patterns and maintain code quality.

## Business Context
This story establishes the architectural foundation for the Djinn mobile application. Without this foundation, all subsequent development would lack consistency, maintainability, and scalability. This is the bedrock upon which the entire mobile application will be built.

**Business Value**: Enables rapid, consistent feature development while maintaining code quality and reducing technical debt.

## Acceptance Criteria

### Measurable Criteria
1. **Flutter Project Structure** - Flutter project follows ADR-20250820-code-organization with proper folder hierarchy
   - **Validation**: Directory structure matches documented standards with 100% compliance
   - **Acceptance Test**: Verify `/lib` contains `/features`, `/shared`, `/core` directories
   - **Performance Benchmark**: Project structure validation completes in <100ms
   - **Success Metric**: All 3 required directories present with correct naming

2. **Riverpod 2.0 Integration** - State management configured with providers and dependency injection
   - **Validation**: Providers compile without errors and dependency injection works
   - **Acceptance Test**: Create test provider and verify state updates propagate in <16ms (60fps)
   - **Performance Benchmark**: Provider initialization <50ms on app startup
   - **Success Metric**: 5+ example providers demonstrating different patterns

3. **go_router Navigation** - Navigation configured for authenticated/unauthenticated route protection
   - **Validation**: Routes redirect based on authentication state within 100ms
   - **Acceptance Test**: Navigate to protected route while unauthenticated triggers redirect
   - **Performance Benchmark**: Route transitions complete in <300ms
   - **Success Metric**: 10+ routes configured with proper guards

4. **Ferry GraphQL Client** - GraphQL client configured with HiveStore caching
   - **Validation**: GraphQL queries execute and cache responses locally
   - **Acceptance Test**: Execute query, verify network call (<500ms), then verify cached response (<10ms)
   - **Performance Benchmark**: Cache hit ratio >95% for repeated queries
   - **Success Metric**: Offline mode returns cached data for 100% of cached queries

5. **Drift Database Setup** - Local database with user profile table and migration support
   - **Validation**: Database creates tables and handles CRUD operations
   - **Acceptance Test**: Insert user profile record (<50ms) and retrieve successfully (<10ms)
   - **Performance Benchmark**: Database initialization <200ms
   - **Success Metric**: Support for 1000+ user records with <100ms query time

6. **App Theming System** - Consistent theme with colors, typography, and component styling
   - **Validation**: Theme applies consistently across 100% of widgets
   - **Acceptance Test**: Toggle between light/dark themes in <100ms with no visual glitches
   - **Performance Benchmark**: Theme switching without widget rebuilds >90% efficiency
   - **Success Metric**: 20+ custom component themes defined

7. **Environment Configuration** - Support for dev/staging/prod environments with different configurations
   - **Validation**: Environment variables load correctly per build configuration
   - **Acceptance Test**: Build for each environment and verify different API endpoints
   - **Performance Benchmark**: Environment detection <10ms at startup
   - **Success Metric**: 3 environments with unique configurations verified

8. **TDD Test Foundation** - Widget tests demonstrate testing patterns and achieve coverage targets
   - **Validation**: Tests run successfully and coverage meets 90% threshold
   - **Acceptance Test**: Run `flutter test --coverage` completes in <30 seconds
   - **Performance Benchmark**: Individual test execution <500ms average
   - **Success Metric**: 90% code coverage with 50+ test cases

## Tasks and Subtasks

### Task 1: Project Structure Setup
- **Effort**: 1 hour
- **Dependencies**: None
- **Status**: ✅ Completed

#### Subtasks:
- [x] Create Flutter project with `flutter create djinn_mobile`
- [x] Implement folder structure per ADR-20250820-code-organization
  - [x] `/lib/core/` - Core utilities and constants
  - [x] `/lib/shared/` - Shared widgets, services, models
  - [x] `/lib/features/` - Feature modules
  - [x] `/test/` - Test files mirroring lib structure
- [x] Configure `pubspec.yaml` with all required dependencies
- [x] Setup `analysis_options.yaml` with linting rules

### Task 2: State Management Configuration
- **Effort**: 1.5 hours
- **Dependencies**: Task 1
- **Status**: ☐ Not Started

#### Subtasks:
- [ ] Add `flutter_riverpod: ^2.0.0` dependency
- [ ] Create `ProviderScope` wrapper in `main.dart`
- [ ] Implement base provider patterns in `/lib/core/providers/`
- [ ] Create example feature provider to validate setup
- [ ] Add provider inspector for development builds

### Task 3: Navigation Setup
- **Effort**: 1 hour
- **Dependencies**: Task 2
- **Status**: ☐ Not Started

#### Subtasks:
- [ ] Add `go_router: ^13.0.0` dependency
- [ ] Create router configuration in `/lib/core/routing/`
- [ ] Implement authentication-aware route guards
- [ ] Define initial routes for authenticated/unauthenticated flows
- [ ] Create navigation service for programmatic navigation

### Task 4: GraphQL Client Integration
- **Effort**: 1 hour
- **Dependencies**: Task 1
- **Status**: ☐ Not Started

#### Subtasks:
- [ ] Add Ferry GraphQL dependencies (`ferry`, `ferry_hive_store`)
- [ ] Configure GraphQL client with API endpoint
- [ ] Setup HiveStore for offline caching
- [ ] Create base GraphQL service class
- [ ] Implement error handling for network failures

### Task 5: Local Database Setup
- **Effort**: 1 hour
- **Dependencies**: Task 1
- **Status**: ☐ Not Started

#### Subtasks:
- [ ] Add Drift dependencies (`drift`, `drift_flutter`)
- [ ] Create database class with user profile table
- [ ] Implement migration strategy
- [ ] Create DAO classes for data access
- [ ] Add database provider for dependency injection

### Task 6: Theming System
- **Effort**: 0.5 hours
- **Dependencies**: Task 1
- **Status**: ☐ Not Started

#### Subtasks:
- [ ] Create theme configuration in `/lib/core/theme/`
- [ ] Define color palette and typography scales
- [ ] Implement light and dark theme variants
- [ ] Create custom component themes
- [ ] Apply theme to MaterialApp

### Task 7: Environment Configuration
- **Effort**: 0.5 hours  
- **Dependencies**: Task 1
- **Status**: ☐ Not Started

#### Subtasks:
- [ ] Create environment config files for dev/staging/prod
- [ ] Implement environment detection logic
- [ ] Configure different API endpoints per environment
- [ ] Setup build flavors for each environment
- [ ] Document environment setup process

### Task 8: Testing Foundation
- **Effort**: 0.5 hours
- **Dependencies**: All previous tasks
- **Status**: ☐ Not Started

#### Subtasks:
- [ ] Create widget test examples for each major component
- [ ] Setup integration test framework
- [ ] Configure test coverage reporting
- [ ] Create testing utilities and mocks
- [ ] Document testing patterns and standards

## Dev Notes

### Architecture Context
This story implements the foundational architecture decisions from multiple ADRs:

**Project Structure** [Source: ADR-20250820-code-organization.md]:
```
lib/
├── core/           # Core utilities, constants, base classes
├── shared/         # Shared widgets, services, models  
├── features/       # Feature modules (auth, profile, etc.)
└── main.dart       # App entry point
```

**State Management** [Source: ADR-20250812-flutter-architecture.md]:
- Riverpod 2.0 for dependency injection and state management
- Provider-based architecture with clear separation of concerns
- Immutable state objects with proper state transitions

**Navigation** [Source: frontend-architecture.md]:
- go_router for declarative routing
- Authentication-aware route protection
- Deep linking support for future features

**Offline-First Architecture** [Source: ADR-20250819-mobile-offline-first.md]:
- Ferry GraphQL client with intelligent caching
- Drift for local database with sync capabilities
- Optimistic updates with conflict resolution

### Key Dependencies
```yaml
dependencies:
  flutter:
    sdk: flutter
  flutter_riverpod: ^2.0.0
  go_router: ^13.0.0
  ferry: ^0.16.0
  ferry_hive_store: ^0.5.0
  drift: ^2.14.0
  drift_flutter: ^0.1.0
  hive: ^2.2.3
```

### Code Examples

**Environment Configuration**:
```dart
// lib/core/config/environment.dart
enum Environment { dev, staging, prod }

class AppConfig {
  static late Environment _env;
  static late String _apiUrl;
  
  static void initialize() {
    const envString = String.fromEnvironment('ENV', defaultValue: 'dev');
    _env = Environment.values.byName(envString);
    
    _apiUrl = switch (_env) {
      Environment.dev => 'https://api-dev.djinn.app/graphql',
      Environment.staging => 'https://api-staging.djinn.app/graphql',
      Environment.prod => 'https://api.djinn.app/graphql',
    };
  }
  
  static String get apiUrl => _apiUrl;
  static bool get isDev => _env == Environment.dev;
  static bool get isProduction => _env == Environment.prod;
}
```

**Router Configuration with Guards**:
```dart
// lib/core/routing/app_router.dart
final routerProvider = Provider<GoRouter>((ref) {
  final authState = ref.watch(authStateProvider);
  
  return GoRouter(
    initialLocation: '/splash',
    redirect: (context, state) {
      final isAuthenticated = authState.isAuthenticated;
      final isAuthRoute = state.path?.startsWith('/auth') ?? false;
      
      if (!isAuthenticated && !isAuthRoute) {
        return '/auth/login';
      }
      if (isAuthenticated && isAuthRoute) {
        return '/home';
      }
      return null;
    },
    routes: [
      GoRoute(
        path: '/splash',
        builder: (context, state) => const SplashScreen(),
      ),
      GoRoute(
        path: '/auth/login',
        builder: (context, state) => const LoginScreen(),
      ),
      GoRoute(
        path: '/home',
        builder: (context, state) => const HomeScreen(),
        routes: [
          GoRoute(
            path: 'profile',
            builder: (context, state) => const ProfileScreen(),
          ),
        ],
      ),
    ],
    errorBuilder: (context, state) => ErrorScreen(error: state.error),
  );
});
```

**GraphQL Client Setup with Caching**:
```dart
// lib/core/graphql/graphql_client.dart
final graphqlClientProvider = Provider<Client>((ref) {
  final box = Hive.box<Map<String, dynamic>>('graphql_cache');
  final store = HiveStore(box);
  final cache = Cache(store: store);
  
  final httpLink = HttpLink(
    AppConfig.apiUrl,
    defaultHeaders: {
      'X-Client-Version': '1.0.0',
      'X-Platform': Platform.operatingSystem,
    },
  );
  
  final authLink = AuthLink(
    getToken: () async {
      final token = await ref.read(authTokenProvider.future);
      return token != null ? 'Bearer $token' : null;
    },
  );
  
  final link = authLink.concat(httpLink);
  
  return Client(
    link: link,
    cache: cache,
    defaultFetchPolicies: {
      OperationType.query: FetchPolicy.CacheFirst,
      OperationType.mutation: FetchPolicy.NetworkOnly,
    },
  );
});
```

**Main App Setup**:
```dart
void main() {
  runApp(
    ProviderScope(
      child: DjinnApp(),
    ),
  );
}

class DjinnApp extends ConsumerWidget {
  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final router = ref.watch(routerProvider);
    final theme = ref.watch(themeProvider);
    
    return MaterialApp.router(
      title: 'Djinn',
      theme: theme.lightTheme,
      darkTheme: theme.darkTheme,
      routerConfig: router,
    );
  }
}
```

**Database Table Definition**:
```dart
@DataClassName('UserProfile')
class UserProfiles extends Table {
  TextColumn get id => text()();
  TextColumn get email => text()();
  TextColumn get name => text().nullable()();
  DateTimeColumn get createdAt => dateTime()();
  DateTimeColumn get updatedAt => dateTime()();
  
  @override
  Set<Column> get primaryKey => {id};
}
```

### Configuration Requirements
- Flutter SDK 3.16+ with null safety
- Dart SDK 3.2+
- iOS deployment target: 12.0+
- Android minSdkVersion: 21

### Integration Points
- Authentication service integration (future story)
- API endpoint configuration per environment
- Push notification service setup (future story)
- Analytics service integration (future story)

### Testing Requirements
[Source: testing-strategy.md]:
- **Unit Tests**: Provider logic, utility functions, data models
- **Widget Tests**: Individual widgets and screens
- **Integration Tests**: Full user flows and navigation
- **Coverage Target**: 90% overall, 95% for core business logic

### Technical Constraints
- Must support both iOS and Android platforms
- Offline-first with eventual consistency
- Performance: App startup < 2 seconds on mid-range devices
- Accessibility: WCAG 2.1 AA compliance

## Definition of Done

### Code Quality ✅
- [ ] Code follows Flutter style guide and linting rules
- [ ] All classes and methods documented with dartdoc comments
- [ ] No compiler warnings or analysis issues
- [ ] Performance profiling shows acceptable startup time

### Testing ✅
- [ ] Unit tests for all providers and utility functions
- [ ] Widget tests for major UI components
- [ ] Integration tests for navigation flows
- [ ] Test coverage ≥ 90% overall
- [ ] All tests pass in CI/CD pipeline

### Documentation ✅
- [ ] Architecture decisions documented in ADRs
- [ ] Setup instructions in README
- [ ] API documentation generated and current
- [ ] Testing patterns documented for future developers

### Deployment ✅
- [ ] App builds successfully for iOS and Android
- [ ] Environment configurations validated
- [ ] Build artifacts generated for all environments
- [ ] No security vulnerabilities in dependencies

### Security ✅
- [ ] Dependency vulnerability scan passes
- [ ] Code signing configured properly
- [ ] Sensitive data not hardcoded in source
- [ ] Network security configurations applied

### Review ✅
- [ ] Code review completed by senior developer
- [ ] Architecture review by tech lead
- [ ] Security review by designated reviewer
- [ ] Documentation review completed

## Dependencies
- **Prerequisites**: Development environment setup
- **Blocks**: All subsequent Flutter development stories
- **Related**: DJINN-1002 (Authentication Implementation)

## Risks and Mitigation

### Risk 1: Version Compatibility Issues
- **Impact**: Medium
- **Probability**: Low (20%)
- **Mitigation**: 
  - Pin specific versions in pubspec.yaml with exact versioning
  - Maintain compatibility matrix document
  - Run dependency audit weekly
  - Use Flutter Version Management (FVM) for consistency

### Risk 2: Platform-Specific Build Issues
- **Impact**: High
- **Probability**: Medium (40%)
- **Mitigation**: 
  - Early testing on both iOS 12+ and Android 21+
  - CI/CD pipeline with platform-specific tests
  - Maintain device testing lab with 5+ physical devices
  - Document platform-specific workarounds

### Risk 3: Performance Issues with Initial Setup
- **Impact**: Medium
- **Probability**: Low (15%)
- **Mitigation**: 
  - Performance profiling during development
  - Optimize critical paths with lazy loading
  - Implement code splitting for features
  - Monitor startup time with Firebase Performance

### Risk 4: Architecture Lock-in
- **Impact**: High
- **Probability**: Medium (35%)
- **Mitigation**:
  - Abstract third-party dependencies behind interfaces
  - Document architecture decisions in ADRs
  - Create migration paths for critical dependencies
  - Maintain 80% unit test coverage for easy refactoring

### Risk 5: State Management Complexity
- **Impact**: Medium
- **Probability**: High (60%)
- **Mitigation**:
  - Establish clear state management patterns early
  - Document provider hierarchies and dependencies
  - Create state management best practices guide
  - Implement state debugging tools in dev mode

### Risk 6: Package Deprecation
- **Impact**: High
- **Probability**: Low (25%)
- **Mitigation**:
  - Choose packages with >95% popularity on pub.dev
  - Verify packages have active maintenance (commits in last 3 months)
  - Create abstraction layers for critical packages
  - Document fallback packages for each dependency

## Rollback Strategy

### Phase 1: Immediate Rollback (0-5 minutes)
1. **Git Reversion**:
   ```bash
   git log --oneline -10  # Identify last stable commit
   git checkout <stable-commit-hash>  # Revert to stable state
   git branch rollback-djinn-1001  # Create rollback branch
   ```

2. **Dependency Rollback**:
   ```bash
   cp pubspec.yaml.backup pubspec.yaml  # Restore backup
   flutter clean  # Clear build cache
   flutter pub get  # Restore previous dependencies
   ```

### Phase 2: Database Rollback (5-10 minutes)
3. **Database Migration Reversal**:
   ```dart
   // Run migration rollback
   await database.rollbackMigration(version: 0);
   // Verify tables dropped
   final tables = await database.rawQuery("SELECT name FROM sqlite_master WHERE type='table'");
   assert(tables.isEmpty, 'Database rollback failed');
   ```

4. **Clear Local Storage**:
   ```bash
   # iOS
   xcrun simctl erase all  # Clear all simulator data
   # Android
   adb shell pm clear com.djinn.app  # Clear app data
   ```

### Phase 3: Configuration Restoration (10-15 minutes)
5. **Environment Configuration**:
   ```bash
   # Restore environment files
   cp -r .env.backup/ .env/
   # Verify configuration
   flutter run --flavor=dev --dart-define-from-file=.env/dev.json
   ```

6. **Build Configuration**:
   ```bash
   # iOS
   cd ios && pod deintegrate && pod install
   # Android
   cd android && ./gradlew clean
   ```

### Phase 4: Validation (15-20 minutes)
7. **Build Verification**:
   ```bash
   # Test all environments
   flutter build ios --flavor=dev --no-codesign
   flutter build android --flavor=dev
   flutter test  # Run existing tests
   ```

8. **Smoke Testing Checklist**:
   - [ ] App launches without crashes
   - [ ] Navigation works between screens
   - [ ] Theme switching functions
   - [ ] Environment detection correct
   - [ ] No console errors or warnings

### Rollback Decision Tree
```
Issue Detected?
├── Build Failure → Phase 1 (Git + Dependencies)
├── Runtime Crash → Phase 2 (Database + Storage)
├── Configuration Error → Phase 3 (Environment + Build)
└── Feature Regression → Phase 4 (Full Validation)
```

### Recovery Time Objectives
- **RTO**: 20 minutes maximum
- **RPO**: Zero data loss (new project)
- **Success Criteria**: Previous stable version running

## Change Log
- **2025-08-25 v1.0**: Initial story creation
- **2025-08-25 v2.0**: Enhanced with measurable benchmarks, comprehensive risks, and detailed rollback strategy
- **Version**: 2.0
- **Status**: Ready for Review

## Dev Agent Record
- **Generated By**: story-creator sub-agent
- **Template Version**: story-template.md v2.1
- **Architecture Sources**: 
  - ADR-20250812-flutter-architecture.md
  - ADR-20250819-mobile-offline-first.md
  - ADR-20250820-code-organization.md
  - ADR-20250120-error-handling.md
- **Context**: Epic DJINN-E001 foundational story
- **Validation**: Pending story-validator review

## QA Results
*Results will be populated during QA review phase*

- **Test Execution**: Pending
- **Coverage Report**: Pending
- **Performance Metrics**: Pending
- **Security Scan**: Pending
- **Accessibility Audit**: Pending

## File List
*Files affected by this story implementation:*

### Created Files:
- `/lib/main.dart` - App entry point
- `/lib/core/providers/app_providers.dart` - Core application providers
- `/lib/core/routing/app_router.dart` - Navigation configuration
- `/lib/core/theme/app_theme.dart` - Theme configuration
- `/lib/core/database/app_database.dart` - Database configuration
- `/lib/core/graphql/graphql_client.dart` - GraphQL client setup
- `/lib/core/config/environment.dart` - Environment configuration
- `/test/widget_test.dart` - Basic widget tests
- `/pubspec.yaml` - Dependencies configuration
- `/analysis_options.yaml` - Linting rules

### Modified Files:
- None (new project setup)

### Configuration Files:
- `/ios/Runner/Info.plist` - iOS configuration
- `/android/app/build.gradle` - Android configuration
- `/.flutter-plugins` - Flutter plugins registry