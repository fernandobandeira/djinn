// This is a basic Flutter widget test.
//
// To perform an interaction with a widget in your test, use the WidgetTester
// utility in the flutter_test package. For example, you can send tap and scroll
// gestures. You can also use WidgetTester to find child widgets in the widget
// tree, read text, and verify that the values of widget properties are correct.

import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import 'package:djinn_mobile/main.dart';
import 'package:djinn_mobile/core/config/environment.dart';

void main() {
  testWidgets('DjinnApp launches with splash screen', (WidgetTester tester) async {
    // Initialize AppConfig for testing
    AppConfig.initialize();
    
    // Build our app and trigger a frame.
    await tester.pumpWidget(const ProviderScope(child: DjinnApp()));

    // Verify that splash screen displays correctly.
    expect(find.text('Djinn'), findsOneWidget);
    expect(find.text('Your Financial Wishes, Granted'), findsOneWidget);
    expect(find.text('DEV MODE'), findsOneWidget);
    
    // Initially shows the icon (not loading since initialization is triggered async)
    await tester.pump();
    
    // Now should show loading indicator
    expect(find.byType(CircularProgressIndicator), findsOneWidget);
    
    // Wait for initialization to complete
    await tester.pump(const Duration(seconds: 2));
    
    // Should now show the icon and success message
    expect(find.byIcon(Icons.auto_awesome), findsOneWidget);
    expect(find.text('Initialized Successfully!'), findsOneWidget);
    expect(find.byType(ElevatedButton), findsOneWidget); // Theme toggle button in dev mode
  });

  testWidgets('Theme toggle works in dev mode', (WidgetTester tester) async {
    AppConfig.initialize();
    
    await tester.pumpWidget(const ProviderScope(child: DjinnApp()));
    await tester.pump(const Duration(seconds: 2));
    
    // Find and tap the theme toggle button
    final themeButton = find.byType(ElevatedButton);
    expect(themeButton, findsOneWidget);
    
    // Initially should show "Dark Mode" button (indicating light theme is active)
    expect(find.text('Dark Mode'), findsOneWidget);
    
    // Tap to switch to dark mode
    await tester.tap(themeButton);
    await tester.pump();
    
    // Should now show "Light Mode" button (indicating dark theme is active)
    expect(find.text('Light Mode'), findsOneWidget);
  });
}
