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
    expect(find.byIcon(Icons.auto_awesome), findsOneWidget);
    expect(find.text('DEV MODE'), findsOneWidget);
  });
}
