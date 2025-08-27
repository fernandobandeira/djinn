import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import 'package:djinn_mobile/core/providers/auth_providers.dart';

void main() {
  group('Auth Providers Tests', () {
    test('authStateProvider should initialize with unauthenticated state', () {
      final container = ProviderContainer();
      addTearDown(container.dispose);

      final authState = container.read(authStateProvider);
      expect(authState.isAuthenticated, isFalse);
      expect(authState.user, isNull);
      expect(authState.isLoading, isFalse);
      expect(authState.error, isNull);
    });

    test('authStateProvider should handle successful sign in', () async {
      final container = ProviderContainer();
      addTearDown(container.dispose);

      final notifier = container.read(authStateProvider.notifier);
      
      // Attempt sign in
      final signInFuture = notifier.signIn('test@example.com', 'password123');
      
      // Should be loading
      var authState = container.read(authStateProvider);
      expect(authState.isLoading, isTrue);
      expect(authState.error, isNull);
      
      // Wait for sign in to complete
      await signInFuture;
      
      // Should be authenticated
      authState = container.read(authStateProvider);
      expect(authState.isLoading, isFalse);
      expect(authState.isAuthenticated, isTrue);
      expect(authState.user, isNotNull);
      expect(authState.user?.email, equals('test@example.com'));
      expect(authState.error, isNull);
    });

    test('authStateProvider should handle sign in failure', () async {
      final container = ProviderContainer();
      addTearDown(container.dispose);

      final notifier = container.read(authStateProvider.notifier);
      
      // Attempt sign in with empty credentials
      await notifier.signIn('', '');
      
      // Should have error
      final authState = container.read(authStateProvider);
      expect(authState.isLoading, isFalse);
      expect(authState.isAuthenticated, isFalse);
      expect(authState.user, isNull);
      expect(authState.error, isNotNull);
    });

    test('authStateProvider should handle sign out', () async {
      final container = ProviderContainer();
      addTearDown(container.dispose);

      final notifier = container.read(authStateProvider.notifier);
      
      // Sign in first
      await notifier.signIn('test@example.com', 'password123');
      expect(container.read(authStateProvider).isAuthenticated, isTrue);
      
      // Sign out
      await notifier.signOut();
      
      // Should be unauthenticated
      final authState = container.read(authStateProvider);
      expect(authState.isAuthenticated, isFalse);
      expect(authState.user, isNull);
      expect(authState.isLoading, isFalse);
    });

    test('authTokenProvider should return token when authenticated', () async {
      final container = ProviderContainer();
      addTearDown(container.dispose);

      // Initially no token
      expect(container.read(authTokenProvider), isNull);
      
      // Sign in
      final notifier = container.read(authStateProvider.notifier);
      await notifier.signIn('test@example.com', 'password123');
      
      // Should have token
      final token = container.read(authTokenProvider);
      expect(token, isNotNull);
      expect(token, startsWith('fake_token_'));
    });

    test('isAuthenticatedProvider should reflect authentication state', () async {
      final container = ProviderContainer();
      addTearDown(container.dispose);

      // Initially not authenticated
      expect(container.read(isAuthenticatedProvider), isFalse);
      
      // Sign in
      final notifier = container.read(authStateProvider.notifier);
      await notifier.signIn('test@example.com', 'password123');
      
      // Should be authenticated
      expect(container.read(isAuthenticatedProvider), isTrue);
      
      // Sign out
      await notifier.signOut();
      
      // Should not be authenticated
      expect(container.read(isAuthenticatedProvider), isFalse);
    });
  });
}