import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import 'core/config/environment.dart';
import 'core/providers/app_providers.dart';
import 'core/routing/app_router.dart';
import 'core/graphql/graphql_client.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  
  // Initialize configuration
  AppConfig.initialize();
  
  // Initialize Hive for GraphQL caching
  await GraphQLClientFactory.initializeHive();
  
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
    final router = ref.watch(routerProvider);
    
    return MaterialApp.router(
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
      routerConfig: router,
      debugShowCheckedModeBanner: false,
    );
  }
}

