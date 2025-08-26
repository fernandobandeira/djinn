import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';

import 'package:djinn_mobile/core/config/environment.dart';
import 'package:djinn_mobile/core/providers/app_providers.dart';
import 'package:djinn_mobile/core/providers/auth_providers.dart';

class SplashScreen extends ConsumerStatefulWidget {
  const SplashScreen({super.key});

  @override
  ConsumerState<SplashScreen> createState() => _SplashScreenState();
}

class _SplashScreenState extends ConsumerState<SplashScreen> {
  @override
  void initState() {
    super.initState();
    // Initialize app and check auth state
    WidgetsBinding.instance.addPostFrameCallback((_) {
      _initialize();
    });
  }

  Future<void> _initialize() async {
    // Initialize app state
    ref.read(appStateProvider.notifier).initialize();
    
    // Wait for initialization
    await Future.delayed(const Duration(seconds: 2));
    
    // Check if user is authenticated
    final isAuthenticated = ref.read(authStateProvider).isAuthenticated;
    
    if (mounted) {
      if (isAuthenticated) {
        context.go('/home');
      } else {
        context.go('/welcome');
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    final appState = ref.watch(appStateProvider);
    final isDev = ref.watch(devToolsProvider);

    return Scaffold(
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            if (appState.isLoading)
              const CircularProgressIndicator(
                color: Colors.purple,
              )
            else
              const Icon(
                Icons.auto_awesome,
                size: 64,
                color: Colors.purple,
              ),
            const SizedBox(height: 16),
            Text(
              'Djinn',
              style: Theme.of(context).textTheme.headlineLarge?.copyWith(
                    color: Colors.purple,
                    fontWeight: FontWeight.bold,
                  ),
            ),
            const SizedBox(height: 8),
            Text(
              'Your Financial Wishes, Granted',
              style: Theme.of(context).textTheme.bodyLarge?.copyWith(
                    color: Theme.of(context).colorScheme.onSurface.withValues(alpha: 0.7),
                  ),
            ),
            const SizedBox(height: 32),
            if (appState.isInitialized)
              Text(
                'Initializing...',
                style: TextStyle(
                  color: Colors.grey[600],
                  fontWeight: FontWeight.w500,
                ),
              ),
            const SizedBox(height: 16),
            if (isDev)
              Container(
                padding: const EdgeInsets.symmetric(
                  horizontal: 12,
                  vertical: 6,
                ),
                decoration: BoxDecoration(
                  color: Colors.orange[100],
                  borderRadius: BorderRadius.circular(12),
                ),
                child: Text(
                  'DEV MODE',
                  style: TextStyle(
                    color: Colors.orange[800],
                    fontWeight: FontWeight.bold,
                    fontSize: 12,
                  ),
                ),
              ),
          ],
        ),
      ),
    );
  }
}