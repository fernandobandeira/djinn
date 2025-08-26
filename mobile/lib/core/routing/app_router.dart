import 'dart:async';

import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';

import 'package:djinn_mobile/core/providers/auth_providers.dart';
import 'package:djinn_mobile/features/auth/screens/login_screen.dart';
import 'package:djinn_mobile/features/auth/screens/register_screen.dart';
import 'package:djinn_mobile/features/auth/screens/forgot_password_screen.dart';
import 'package:djinn_mobile/features/home/screens/home_screen.dart';
import 'package:djinn_mobile/features/profile/screens/profile_screen.dart';
import 'package:djinn_mobile/features/settings/screens/settings_screen.dart';
import 'package:djinn_mobile/features/onboarding/screens/welcome_screen.dart';
import 'package:djinn_mobile/features/receipts/screens/receipts_screen.dart';
import 'package:djinn_mobile/features/receipts/screens/receipt_detail_screen.dart';
import 'package:djinn_mobile/features/wishes/screens/wishes_screen.dart';
import 'package:djinn_mobile/shared/screens/splash_screen.dart';
import 'package:djinn_mobile/shared/screens/error_screen.dart';

// Router provider
final routerProvider = Provider<GoRouter>((ref) {
  final authState = ref.watch(authStateProvider);
  
  return GoRouter(
    initialLocation: '/splash',
    refreshListenable: GoRouterRefreshStream(ref.watch(authStateProvider.notifier).stream),
    redirect: (context, state) {
      final isAuthenticated = authState.isAuthenticated;
      final location = state.matchedLocation;
      
      // Public routes that don't require authentication
      final publicRoutes = [
        '/splash',
        '/welcome',
        '/auth/login',
        '/auth/register',
        '/auth/forgot-password',
      ];
      
      final isPublicRoute = publicRoutes.contains(location);
      
      // If not authenticated and trying to access protected route
      if (!isAuthenticated && !isPublicRoute) {
        return '/auth/login';
      }
      
      // If authenticated and trying to access auth routes
      if (isAuthenticated && location.startsWith('/auth')) {
        return '/home';
      }
      
      // If authenticated and on splash or welcome
      if (isAuthenticated && (location == '/splash' || location == '/welcome')) {
        return '/home';
      }
      
      return null; // No redirect needed
    },
    routes: [
      // Splash screen
      GoRoute(
        path: '/splash',
        name: 'splash',
        builder: (context, state) => const SplashScreen(),
      ),
      
      // Onboarding
      GoRoute(
        path: '/welcome',
        name: 'welcome',
        builder: (context, state) => const WelcomeScreen(),
      ),
      
      // Authentication routes
      GoRoute(
        path: '/auth',
        redirect: (_, __) => '/auth/login',
        routes: [
          GoRoute(
            path: 'login',
            name: 'login',
            builder: (context, state) => const LoginScreen(),
          ),
          GoRoute(
            path: 'register',
            name: 'register',
            builder: (context, state) => const RegisterScreen(),
          ),
          GoRoute(
            path: 'forgot-password',
            name: 'forgot-password',
            builder: (context, state) => const ForgotPasswordScreen(),
          ),
        ],
      ),
      
      // Main app routes (protected)
      GoRoute(
        path: '/home',
        name: 'home',
        builder: (context, state) => const HomeScreen(),
        routes: [
          // Receipts
          GoRoute(
            path: 'receipts',
            name: 'receipts',
            builder: (context, state) => const ReceiptsScreen(),
            routes: [
              GoRoute(
                path: ':id',
                name: 'receipt-detail',
                builder: (context, state) {
                  final receiptId = state.pathParameters['id']!;
                  return ReceiptDetailScreen(receiptId: receiptId);
                },
              ),
            ],
          ),
          
          // Wishes
          GoRoute(
            path: 'wishes',
            name: 'wishes',
            builder: (context, state) => const WishesScreen(),
          ),
          
          // Profile
          GoRoute(
            path: 'profile',
            name: 'profile',
            builder: (context, state) => const ProfileScreen(),
          ),
          
          // Settings
          GoRoute(
            path: 'settings',
            name: 'settings',
            builder: (context, state) => const SettingsScreen(),
          ),
        ],
      ),
    ],
    
    // Error handling
    errorBuilder: (context, state) => ErrorScreen(error: state.error),
  );
});

// Stream for router refresh when auth state changes
class GoRouterRefreshStream extends ChangeNotifier {
  GoRouterRefreshStream(Stream<AuthState> stream) {
    _subscription = stream.listen(
      (_) {
        notifyListeners();
      },
      onError: (error) {
        // Log error but don't crash the app
        if (kDebugMode) {
          print('Router refresh stream error: $error');
        }
      },
      cancelOnError: false, // Continue listening even after errors
    );
  }

  late final StreamSubscription<AuthState> _subscription;

  @override
  void dispose() {
    try {
      _subscription.cancel();
    } catch (e) {
      // Ensure disposal completes even if cancellation fails
      if (kDebugMode) {
        print('Error cancelling router subscription: $e');
      }
    }
    super.dispose();
  }
}