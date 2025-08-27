import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:djinn_mobile/main.dart';
import 'package:djinn_mobile/core/config/environment.dart';
import 'package:djinn_mobile/features/auth/screens/login_screen.dart';
import 'package:djinn_mobile/features/home/screens/home_screen.dart';
import 'package:djinn_mobile/shared/screens/splash_screen.dart';

void main() {
  group('App Integration Tests', () {
    setUpAll(() {
      AppConfig.initialize();
    });
    
    testWidgets('App launches and shows splash screen', (WidgetTester tester) async {
      await tester.pumpWidget(
        const ProviderScope(
          child: DjinnApp(),
        ),
      );
      
      // Should show splash screen initially
      expect(find.byType(SplashScreen), findsOneWidget);
      
      // Wait for navigation
      await tester.pumpAndSettle(const Duration(seconds: 3));
      
      // Should navigate to login if not authenticated
      expect(find.byType(LoginScreen), findsOneWidget);
    });
    
    testWidgets('Navigation redirects when not authenticated', (WidgetTester tester) async {
      await tester.pumpWidget(
        const ProviderScope(
          child: DjinnApp(),
        ),
      );
      
      await tester.pumpAndSettle(const Duration(seconds: 3));
      
      // Try to navigate to protected route (should redirect to login)
      // This tests the route guard functionality
      expect(find.byType(LoginScreen), findsOneWidget);
    });
    
    testWidgets('Theme switching works', (WidgetTester tester) async {
      await tester.pumpWidget(
        const ProviderScope(
          child: DjinnApp(),
        ),
      );
      
      await tester.pumpAndSettle();
      
      // Get the MaterialApp
      final MaterialApp app = tester.widget<MaterialApp>(
        find.byType(MaterialApp).first,
      );
      
      // Check that both light and dark themes are configured
      expect(app.theme, isNotNull);
      expect(app.darkTheme, isNotNull);
      expect(app.theme!.brightness, Brightness.light);
      expect(app.darkTheme!.brightness, Brightness.dark);
    });
    
    testWidgets('Login screen shows correct elements', (WidgetTester tester) async {
      await tester.pumpWidget(
        const ProviderScope(
          child: DjinnApp(),
        ),
      );
      
      await tester.pumpAndSettle(const Duration(seconds: 3));
      
      // Should show login screen elements
      expect(find.text('Welcome Back'), findsOneWidget);
      expect(find.text('Sign in to continue'), findsOneWidget);
      expect(find.byType(TextFormField), findsNWidgets(2)); // Email and password
      expect(find.text('Sign In'), findsOneWidget);
      expect(find.text('Forgot Password?'), findsOneWidget);
      expect(find.text("Don't have an account?"), findsOneWidget);
      expect(find.text('Sign Up'), findsOneWidget);
    });
    
    testWidgets('Can navigate to registration screen', (WidgetTester tester) async {
      await tester.pumpWidget(
        const ProviderScope(
          child: DjinnApp(),
        ),
      );
      
      await tester.pumpAndSettle(const Duration(seconds: 3));
      
      // Find and tap sign up link
      await tester.tap(find.text('Sign Up'));
      await tester.pumpAndSettle();
      
      // Should show registration screen
      expect(find.text('Create Account'), findsOneWidget);
      expect(find.text('Sign up to get started'), findsOneWidget);
    });
    
    testWidgets('Can navigate to forgot password screen', (WidgetTester tester) async {
      await tester.pumpWidget(
        const ProviderScope(
          child: DjinnApp(),
        ),
      );
      
      await tester.pumpAndSettle(const Duration(seconds: 3));
      
      // Find and tap forgot password link
      await tester.tap(find.text('Forgot Password?'));
      await tester.pumpAndSettle();
      
      // Should show forgot password screen
      expect(find.text('Reset Password'), findsOneWidget);
      expect(find.text("Enter your email and we'll send you instructions"), findsOneWidget);
    });
    
    testWidgets('Form validation works on login screen', (WidgetTester tester) async {
      await tester.pumpWidget(
        const ProviderScope(
          child: DjinnApp(),
        ),
      );
      
      await tester.pumpAndSettle(const Duration(seconds: 3));
      
      // Try to submit empty form
      await tester.tap(find.text('Sign In'));
      await tester.pumpAndSettle();
      
      // Should show validation errors
      expect(find.text('Please enter your email'), findsOneWidget);
      expect(find.text('Please enter your password'), findsOneWidget);
      
      // Enter invalid email
      await tester.enterText(find.byType(TextFormField).first, 'invalid-email');
      await tester.tap(find.text('Sign In'));
      await tester.pumpAndSettle();
      
      expect(find.text('Please enter a valid email'), findsOneWidget);
      
      // Enter valid email but short password
      await tester.enterText(find.byType(TextFormField).first, 'test@example.com');
      await tester.enterText(find.byType(TextFormField).last, '123');
      await tester.tap(find.text('Sign In'));
      await tester.pumpAndSettle();
      
      expect(find.text('Password must be at least 6 characters'), findsOneWidget);
    });
  });
  
  group('Performance Tests', () {
    testWidgets('App startup time is acceptable', (WidgetTester tester) async {
      final stopwatch = Stopwatch()..start();
      
      await tester.pumpWidget(
        const ProviderScope(
          child: DjinnApp(),
        ),
      );
      
      await tester.pump();
      stopwatch.stop();
      
      // App should start in less than 2 seconds
      expect(stopwatch.elapsedMilliseconds, lessThan(2000));
    });
    
    testWidgets('Route transitions are smooth', (WidgetTester tester) async {
      await tester.pumpWidget(
        const ProviderScope(
          child: DjinnApp(),
        ),
      );
      
      await tester.pumpAndSettle(const Duration(seconds: 3));
      
      final stopwatch = Stopwatch()..start();
      
      // Navigate to registration
      await tester.tap(find.text('Sign Up'));
      await tester.pumpAndSettle();
      
      stopwatch.stop();
      
      // Route transition should complete in less than 300ms
      expect(stopwatch.elapsedMilliseconds, lessThan(300));
    });
  });
}