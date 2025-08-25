import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import 'core/config/environment.dart';
import 'core/providers/app_providers.dart';

void main() {
  AppConfig.initialize();
  runApp(
    ProviderScope(
      observers: AppConfig.isDev ? [DevProviderObserver()] : null,
      child: const DjinnApp(),
    ),
  );
}

class DevProviderObserver extends ProviderObserver {
  @override
  void didUpdateProvider(
    ProviderBase<Object?> provider,
    Object? previousValue,
    Object? newValue,
    ProviderContainer container,
  ) {
    debugPrint(
      '[PROVIDER] ${provider.name ?? provider.runtimeType} '
      'updated: $previousValue → $newValue',
    );
  }

  @override
  void didAddProvider(
    ProviderBase<Object?> provider,
    Object? value,
    ProviderContainer container,
  ) {
    debugPrint('[PROVIDER] ${provider.name ?? provider.runtimeType} added: $value');
  }

  @override
  void didDisposeProvider(
    ProviderBase<Object?> provider,
    ProviderContainer container,
  ) {
    debugPrint('[PROVIDER] ${provider.name ?? provider.runtimeType} disposed');
  }
}

class DjinnApp extends ConsumerWidget {
  const DjinnApp({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final isDarkMode = ref.watch(themeModeProvider);
    
    return MaterialApp(
      title: 'Djinn - Your Financial Wishes, Granted',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.purple),
        useMaterial3: true,
      ),
      darkTheme: ThemeData(
        colorScheme: ColorScheme.fromSeed(
          seedColor: Colors.purple,
          brightness: Brightness.dark,
        ),
        useMaterial3: true,
      ),
      themeMode: isDarkMode ? ThemeMode.dark : ThemeMode.light,
      home: const SplashScreen(),
    );
  }
}

class SplashScreen extends ConsumerStatefulWidget {
  const SplashScreen({super.key});

  @override
  ConsumerState<SplashScreen> createState() => _SplashScreenState();
}

class _SplashScreenState extends ConsumerState<SplashScreen> {
  @override
  void initState() {
    super.initState();
    // Initialize app state when splash screen loads
    WidgetsBinding.instance.addPostFrameCallback((_) {
      ref.read(appStateProvider.notifier).initialize();
    });
  }

  @override
  Widget build(BuildContext context) {
    final appState = ref.watch(appStateProvider);
    final isDev = ref.watch(devToolsProvider);
    final isDarkMode = ref.watch(themeModeProvider);

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
                'Initialized Successfully!',
                style: TextStyle(
                  color: Colors.green[600],
                  fontWeight: FontWeight.w500,
                ),
              ),
            const SizedBox(height: 16),
            if (isDev) ...[
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
              const SizedBox(height: 16),
              ElevatedButton(
                onPressed: () {
                  ref.read(themeModeProvider.notifier).state = !isDarkMode;
                },
                child: Text(isDarkMode ? 'Light Mode' : 'Dark Mode'),
              ),
            ],
          ],
        ),
      ),
    );
  }
}
