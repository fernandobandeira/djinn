import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';

// Navigation service provider
final navigationServiceProvider = Provider<NavigationService>((ref) {
  return NavigationService();
});

class NavigationService {
  // Navigate to a named route
  void goNamed(
    BuildContext context,
    String name, {
    Map<String, String> pathParameters = const {},
    Map<String, dynamic> queryParameters = const {},
    Object? extra,
  }) {
    context.goNamed(
      name,
      pathParameters: pathParameters,
      queryParameters: queryParameters,
      extra: extra,
    );
  }
  
  // Navigate to a path
  void go(BuildContext context, String location, {Object? extra}) {
    context.go(location, extra: extra);
  }
  
  // Push a named route onto the navigation stack
  void pushNamed(
    BuildContext context,
    String name, {
    Map<String, String> pathParameters = const {},
    Map<String, dynamic> queryParameters = const {},
    Object? extra,
  }) {
    context.pushNamed(
      name,
      pathParameters: pathParameters,
      queryParameters: queryParameters,
      extra: extra,
    );
  }
  
  // Push a path onto the navigation stack
  void push(BuildContext context, String location, {Object? extra}) {
    context.push(location, extra: extra);
  }
  
  // Pop the current route
  void pop(BuildContext context) {
    if (context.canPop()) {
      context.pop();
    }
  }
  
  // Replace the current route
  void replaceNamed(
    BuildContext context,
    String name, {
    Map<String, String> pathParameters = const {},
    Map<String, dynamic> queryParameters = const {},
    Object? extra,
  }) {
    context.replaceNamed(
      name,
      pathParameters: pathParameters,
      queryParameters: queryParameters,
      extra: extra,
    );
  }
  
  // Check if can pop
  bool canPop(BuildContext context) {
    return context.canPop();
  }
}

// Extension methods for easier navigation
extension NavigationExtension on BuildContext {
  NavigationService get nav => NavigationService();
  
  void goToHome() => go('/home');
  void goToLogin() => go('/auth/login');
  void goToRegister() => go('/auth/register');
  void goToProfile() => go('/home/profile');
  void goToSettings() => go('/home/settings');
  void goToReceipts() => go('/home/receipts');
  void goToWishes() => go('/home/wishes');
}