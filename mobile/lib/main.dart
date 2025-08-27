import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import 'core/config/environment.dart';
import 'core/providers/app_providers.dart';
import 'core/routing/app_router.dart';
import 'core/graphql/graphql_client.dart';
import 'core/database/database_provider.dart';
import 'core/theme/app_theme.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  
  // Initialize configuration
  AppConfig.initialize();
  
  // Initialize Hive for GraphQL caching
  await GraphQLClientFactory.initializeHive();
  
  // Set preferred orientations
  await SystemChrome.setPreferredOrientations([
    DeviceOrientation.portraitUp,
    DeviceOrientation.portraitDown,
  ]);
  
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
    final themeMode = ref.watch(themeModeProvider);
    final appTheme = ref.watch(appThemeProvider);
    final router = ref.watch(routerProvider);
    
    // Initialize database only once using read instead of watch
    // This prevents unnecessary rebuilds when database state changes
    final databaseInit = ref.read(databaseInitializerProvider);
    
    return MaterialApp.router(
      title: AppConfig.appName,
      theme: appTheme.lightTheme,
      darkTheme: appTheme.darkTheme,
      themeMode: themeMode,
      routerConfig: router,
      debugShowCheckedModeBanner: AppConfig.isDev,
    );
  }
}

