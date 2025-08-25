import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import 'package:djinn_mobile/core/providers/app_providers.dart';
import 'package:djinn_mobile/core/config/environment.dart';

void main() {
  setUpAll(() {
    AppConfig.initialize();
  });

  group('App Providers Tests', () {
    test('environmentProvider should return current environment', () {
      final container = ProviderContainer();
      addTearDown(container.dispose);

      final env = container.read(environmentProvider);
      expect(env, equals(Environment.dev)); // Default environment is dev
    });

    test('appStateProvider should initialize correctly', () {
      final container = ProviderContainer();
      addTearDown(container.dispose);

      final initialState = container.read(appStateProvider);
      expect(initialState.isInitialized, isFalse);
      expect(initialState.isLoading, isFalse);
      expect(initialState.error, isNull);
    });

    test('appStateProvider should handle initialization', () async {
      final container = ProviderContainer();
      addTearDown(container.dispose);

      final notifier = container.read(appStateProvider.notifier);
      
      // Start initialization
      notifier.initialize();
      
      // Should be loading
      var state = container.read(appStateProvider);
      expect(state.isLoading, isTrue);
      expect(state.isInitialized, isFalse);
      
      // Wait for initialization to complete
      await Future.delayed(const Duration(seconds: 2));
      
      // Should be initialized
      state = container.read(appStateProvider);
      expect(state.isLoading, isFalse);
      expect(state.isInitialized, isTrue);
    });

    test('themeModeProvider should default to light mode', () {
      final container = ProviderContainer();
      addTearDown(container.dispose);

      final isDarkMode = container.read(themeModeProvider);
      expect(isDarkMode, isFalse);
    });

    test('themeModeProvider should toggle correctly', () {
      final container = ProviderContainer();
      addTearDown(container.dispose);

      // Initially light mode
      expect(container.read(themeModeProvider), isFalse);
      
      // Toggle to dark mode
      container.read(themeModeProvider.notifier).state = true;
      expect(container.read(themeModeProvider), isTrue);
      
      // Toggle back to light mode
      container.read(themeModeProvider.notifier).state = false;
      expect(container.read(themeModeProvider), isFalse);
    });

    test('devToolsProvider should return true in dev environment', () {
      final container = ProviderContainer();
      addTearDown(container.dispose);

      final isDevTools = container.read(devToolsProvider);
      expect(isDevTools, isTrue); // Should be true in dev environment
    });
  });
}