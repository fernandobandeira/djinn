import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:djinn_mobile/core/config/environment.dart';

/// Wrapper for testing widgets with Riverpod
class TestWrapper extends StatelessWidget {
  final Widget child;
  final List<Override> overrides;
  
  const TestWrapper({
    super.key,
    required this.child,
    this.overrides = const [],
  });
  
  @override
  Widget build(BuildContext context) {
    return ProviderScope(
      overrides: overrides,
      child: MaterialApp(
        home: child,
      ),
    );
  }
}

/// Helper to pump a widget with all necessary wrappers
Future<void> pumpTestWidget(
  WidgetTester tester,
  Widget widget, {
  List<Override> overrides = const [],
}) async {
  // Initialize environment for testing
  AppConfig.initialize();
  
  await tester.pumpWidget(
    TestWrapper(
      overrides: overrides,
      child: widget,
    ),
  );
}

/// Helper to pump a routed widget
Future<void> pumpRoutedWidget(
  WidgetTester tester,
  Widget widget, {
  List<Override> overrides = const [],
}) async {
  AppConfig.initialize();
  
  await tester.pumpWidget(
    ProviderScope(
      overrides: overrides,
      child: MaterialApp(
        home: widget,
      ),
    ),
  );
}

/// Helper to find widgets by key
Finder findByKey(String key) => find.byKey(Key(key));

/// Helper to find text containing substring
Finder findTextContaining(String substring) {
  return find.byWidgetPredicate(
    (widget) => widget is Text && widget.data?.contains(substring) == true,
  );
}

/// Helper to wait for async operations
Future<void> pumpAndSettle(
  WidgetTester tester, {
  Duration duration = const Duration(milliseconds: 100),
  int maxAttempts = 10,
}) async {
  for (int i = 0; i < maxAttempts; i++) {
    await tester.pump(duration);
    if (!tester.binding.hasScheduledFrame) {
      break;
    }
  }
}

/// Helper to test loading states
Future<void> expectLoadingState(WidgetTester tester) async {
  expect(find.byType(CircularProgressIndicator), findsOneWidget);
}

/// Helper to test error states
Future<void> expectErrorState(WidgetTester tester, String errorMessage) async {
  expect(find.text(errorMessage), findsOneWidget);
}

/// Helper to test success states
Future<void> expectSuccessState(WidgetTester tester, String successMessage) async {
  expect(find.text(successMessage), findsOneWidget);
}

/// Helper to simulate network delay
Future<void> simulateNetworkDelay({
  Duration delay = const Duration(seconds: 1),
}) async {
  await Future.delayed(delay);
}

/// Extension methods for testing
extension WidgetTesterExtensions on WidgetTester {
  /// Tap and wait for animation
  Future<void> tapAndSettle(Finder finder) async {
    await tap(finder);
    await pumpAndSettle();
  }
  
  /// Enter text and wait for animation
  Future<void> enterTextAndSettle(Finder finder, String text) async {
    await enterText(finder, text);
    await pumpAndSettle();
  }
  
  /// Scroll until visible
  Future<void> scrollUntilVisible(
    Finder finder, {
    Finder? scrollable,
    double delta = 100,
    int maxScrolls = 20,
  }) async {
    for (int i = 0; i < maxScrolls; i++) {
      if (finder.evaluate().isNotEmpty) {
        break;
      }
      await drag(scrollable ?? find.byType(Scrollable).first, Offset(0, -delta));
      await pump();
    }
  }
}

/// Golden test configuration
class GoldenTestConfig {
  static void setup() {
    // Set up golden test configuration
    TestWidgetsFlutterBinding.ensureInitialized();
    
    // Set default size for golden tests
    binding.window.physicalSizeTestValue = const Size(390, 844); // iPhone 14
    binding.window.devicePixelRatioTestValue = 3.0;
  }
  
  static void tearDown() {
    binding.window.clearPhysicalSizeTestValue();
    binding.window.clearDevicePixelRatioTestValue();
  }
  
  static TestWidgetsFlutterBinding get binding {
    return TestWidgetsFlutterBinding.instance;
  }
}

/// Test data generators
class TestData {
  static String generateEmail() {
    final timestamp = DateTime.now().millisecondsSinceEpoch;
    return 'test$timestamp@example.com';
  }
  
  static String generateUsername() {
    final timestamp = DateTime.now().millisecondsSinceEpoch;
    return 'user$timestamp';
  }
  
  static String generatePassword() {
    return 'Test@123!${DateTime.now().millisecondsSinceEpoch}';
  }
  
  static Map<String, dynamic> generateUserProfile() {
    return {
      'id': DateTime.now().millisecondsSinceEpoch.toString(),
      'email': generateEmail(),
      'name': 'Test User',
      'createdAt': DateTime.now().toIso8601String(),
      'updatedAt': DateTime.now().toIso8601String(),
    };
  }
}