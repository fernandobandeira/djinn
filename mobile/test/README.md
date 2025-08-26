# Testing Guide

## Overview
This Flutter app follows a comprehensive testing strategy with unit tests, widget tests, and integration tests to ensure code quality and reliability.

## Test Structure
```
test/
├── core/                  # Core functionality tests
│   ├── database/         # Database and DAO tests
│   ├── theme/           # Theme and styling tests
│   ├── providers/       # Provider and state management tests
│   └── graphql/         # GraphQL client tests
├── features/             # Feature-specific tests
│   ├── auth/           # Authentication tests
│   ├── home/           # Home screen tests
│   └── ...
├── integration/         # Integration tests
├── test_helpers/       # Test utilities and mocks
│   ├── test_helpers.dart
│   └── mocks.dart
└── widget_test.dart    # Basic widget tests
```

## Running Tests

### All Tests
```bash
flutter test
```

### Specific Test File
```bash
flutter test test/core/database/database_test.dart
```

### With Coverage
```bash
flutter test --coverage
```

### Generate Coverage Report
```bash
# Run tests with coverage
flutter test --coverage

# Generate HTML report (requires lcov)
genhtml coverage/lcov.info -o coverage/html

# Open report
open coverage/html/index.html
```

### Watch Mode (with very_good_cli)
```bash
very_good test --watch
```

## Test Types

### 1. Unit Tests
Test individual functions, classes, and business logic.

Example:
```dart
test('should calculate total correctly', () {
  final calculator = Calculator();
  expect(calculator.add(2, 3), equals(5));
});
```

### 2. Widget Tests
Test individual widgets in isolation.

Example:
```dart
testWidgets('button shows correct text', (tester) async {
  await tester.pumpWidget(
    MaterialApp(
      home: MyButton(text: 'Click me'),
    ),
  );
  
  expect(find.text('Click me'), findsOneWidget);
});
```

### 3. Integration Tests
Test complete user flows and app behavior.

Example:
```dart
testWidgets('user can login', (tester) async {
  await tester.pumpWidget(MyApp());
  
  await tester.enterText(find.byKey(Key('email')), 'test@example.com');
  await tester.enterText(find.byKey(Key('password')), 'password123');
  await tester.tap(find.text('Login'));
  await tester.pumpAndSettle();
  
  expect(find.text('Welcome'), findsOneWidget);
});
```

## Test Helpers

### TestWrapper
Wraps widgets with necessary providers and material app:
```dart
await pumpTestWidget(
  tester,
  MyWidget(),
  overrides: [
    authProvider.overrideWith(() => MockAuth()),
  ],
);
```

### Mock Providers
Use mock providers for testing:
```dart
final overrides = getMockProviderOverrides(
  isAuthenticated: true,
  database: mockDatabase,
  graphQLClient: mockClient,
);
```

### Test Data Generators
Generate test data:
```dart
final email = TestData.generateEmail();
final user = TestData.generateUserProfile();
```

## Golden Tests

### Creating Golden Files
```dart
testWidgets('renders correctly', (tester) async {
  GoldenTestConfig.setup();
  
  await tester.pumpWidget(MyWidget());
  await expectLater(
    find.byType(MyWidget),
    matchesGoldenFile('goldens/my_widget.png'),
  );
  
  GoldenTestConfig.tearDown();
});
```

### Updating Golden Files
```bash
flutter test --update-goldens
```

## Performance Testing

### Widget Performance
```dart
testWidgets('renders quickly', (tester) async {
  final stopwatch = Stopwatch()..start();
  
  await tester.pumpWidget(MyWidget());
  
  stopwatch.stop();
  expect(stopwatch.elapsedMilliseconds, lessThan(16)); // 60fps
});
```

### Memory Testing
```dart
test('does not leak memory', () {
  final weakRef = WeakReference(MyObject());
  
  // Force garbage collection
  for (int i = 0; i < 100; i++) {
    List.filled(1000000, null);
  }
  
  expect(weakRef.target, isNull);
});
```

## Continuous Integration

### GitHub Actions
```yaml
name: Tests

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - uses: subosito/flutter-action@v2
      - run: flutter pub get
      - run: flutter test --coverage
      - uses: codecov/codecov-action@v2
```

## Best Practices

### 1. Test Naming
- Use descriptive names
- Follow pattern: `should_expectedBehavior_when_condition`
- Example: `should_return_user_when_id_exists`

### 2. Test Organization
- Group related tests using `group()`
- Use `setUp()` and `tearDown()` for common initialization
- Keep tests focused and isolated

### 3. Assertions
- Use specific matchers: `equals()`, `contains()`, `isA<Type>()`
- Test both success and failure cases
- Verify side effects and state changes

### 4. Mocking
- Mock external dependencies
- Use dependency injection for testability
- Keep mocks simple and focused

### 5. Coverage Goals
- Target: 90% overall coverage
- Critical paths: 95% coverage
- UI code: 80% coverage
- Generated code: excluded from coverage

## Debugging Tests

### Print Debugging
```dart
debugPrint('Value: $value');
```

### Visual Debugging
```dart
await tester.pumpWidget(MyWidget());
debugDumpApp(); // Prints widget tree
```

### Breakpoints
Use IDE breakpoints or `debugger()` statement

### Test Timeout
```dart
test('long running test', () async {
  // test code
}, timeout: Timeout(Duration(seconds: 10)));
```

## Common Issues

### Issue: "Timer still pending"
Solution: Use `tester.pumpAndSettle()` or `tester.binding.delayed()`

### Issue: "Widget not found"
Solution: Ensure widget is in tree with `debugDumpApp()`

### Issue: "State not updating"
Solution: Use `await tester.pump()` after state changes

### Issue: "RenderFlex overflow"
Solution: Wrap widget in sized container for tests

## Resources

- [Flutter Testing Documentation](https://flutter.dev/docs/testing)
- [Flutter Test Package](https://pub.dev/packages/flutter_test)
- [Mockito Package](https://pub.dev/packages/mockito)
- [Golden Tests](https://flutter.dev/docs/testing#golden-tests)