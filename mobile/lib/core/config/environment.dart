import 'package:flutter/foundation.dart';
import 'package:djinn_mobile/core/utils/logger.dart';

enum Environment { dev, staging, prod }

class AppConfig {
  static late Environment _env;
  static late EnvironmentConfig _config;
  
  static void initialize() {
    const envString = String.fromEnvironment('ENV', defaultValue: 'dev');
    _env = Environment.values.byName(envString);
    
    _config = switch (_env) {
      Environment.dev => EnvironmentConfig.dev(),
      Environment.staging => EnvironmentConfig.staging(),
      Environment.prod => EnvironmentConfig.prod(),
    };
    
    logger.info('App initialized in ${_env.name} environment', tag: 'AppConfig');
    logger.debug('API URL: ${_config.apiUrl}', tag: 'AppConfig');
  }
  
  // Environment checks
  static bool get isDev => _env == Environment.dev;
  static bool get isProduction => _env == Environment.prod;
  static bool get isStaging => _env == Environment.staging;
  static Environment get currentEnv => _env;
  
  // Configuration getters
  static String get apiUrl => _config.apiUrl;
  static String get websocketUrl => _config.websocketUrl;
  static String get appName => _config.appName;
  static String get appVersion => _config.appVersion;
  static bool get enableLogging => _config.enableLogging;
  static bool get enableCrashlytics => _config.enableCrashlytics;
  static bool get enableAnalytics => _config.enableAnalytics;
  static bool get enablePerformanceMonitoring => _config.enablePerformanceMonitoring;
  static Duration get apiTimeout => _config.apiTimeout;
  static int get maxRetryAttempts => _config.maxRetryAttempts;
  static String get sentryDsn => _config.sentryDsn;
  static Map<String, String> get featureFlags => _config.featureFlags;
  static EnvironmentConfig get config => _config;
}

class EnvironmentConfig {
  final String apiUrl;
  final String websocketUrl;
  final String appName;
  final String appVersion;
  final bool enableLogging;
  final bool enableCrashlytics;
  final bool enableAnalytics;
  final bool enablePerformanceMonitoring;
  final Duration apiTimeout;
  final int maxRetryAttempts;
  final String sentryDsn;
  final Map<String, String> featureFlags;
  
  const EnvironmentConfig({
    required this.apiUrl,
    required this.websocketUrl,
    required this.appName,
    required this.appVersion,
    required this.enableLogging,
    required this.enableCrashlytics,
    required this.enableAnalytics,
    required this.enablePerformanceMonitoring,
    required this.apiTimeout,
    required this.maxRetryAttempts,
    required this.sentryDsn,
    required this.featureFlags,
  });
  
  // Development environment
  factory EnvironmentConfig.dev() {
    return EnvironmentConfig(
      apiUrl: const String.fromEnvironment('API_URL', 
        defaultValue: 'http://localhost:4000/graphql'),
      websocketUrl: const String.fromEnvironment('WS_URL',
        defaultValue: 'ws://localhost:4000/graphql'),
      appName: 'Djinn Dev',
      appVersion: '1.0.0-dev',
      enableLogging: true,
      enableCrashlytics: false,
      enableAnalytics: false,
      enablePerformanceMonitoring: true,
      apiTimeout: const Duration(seconds: 30),
      maxRetryAttempts: 3,
      sentryDsn: const String.fromEnvironment('SENTRY_DSN', defaultValue: ''),
      featureFlags: {
        'enable_firebase_auth': 'false',
        'enable_push_notifications': 'false',
        'enable_in_app_purchases': 'false',
        'enable_debug_menu': 'true',
        'enable_mock_data': 'true',
      },
    );
  }
  
  // Staging environment
  factory EnvironmentConfig.staging() {
    return EnvironmentConfig(
      apiUrl: const String.fromEnvironment('API_URL',
        defaultValue: 'https://api-staging.djinn.app/graphql'),
      websocketUrl: const String.fromEnvironment('WS_URL',
        defaultValue: 'wss://api-staging.djinn.app/graphql'),
      appName: 'Djinn Staging',
      appVersion: '1.0.0-staging',
      enableLogging: true,
      enableCrashlytics: true,
      enableAnalytics: true,
      enablePerformanceMonitoring: true,
      apiTimeout: const Duration(seconds: 20),
      maxRetryAttempts: 3,
      sentryDsn: const String.fromEnvironment('SENTRY_DSN', defaultValue: ''),
      featureFlags: {
        'enable_firebase_auth': 'true',
        'enable_push_notifications': 'true',
        'enable_in_app_purchases': 'false',
        'enable_debug_menu': 'true',
        'enable_mock_data': 'false',
      },
    );
  }
  
  // Production environment
  factory EnvironmentConfig.prod() {
    return EnvironmentConfig(
      apiUrl: const String.fromEnvironment('API_URL',
        defaultValue: 'https://api.djinn.app/graphql'),
      websocketUrl: const String.fromEnvironment('WS_URL',
        defaultValue: 'wss://api.djinn.app/graphql'),
      appName: 'Djinn',
      appVersion: '1.0.0',
      enableLogging: false,
      enableCrashlytics: true,
      enableAnalytics: true,
      enablePerformanceMonitoring: true,
      apiTimeout: const Duration(seconds: 15),
      maxRetryAttempts: 5,
      sentryDsn: const String.fromEnvironment('SENTRY_DSN', defaultValue: ''),
      featureFlags: {
        'enable_firebase_auth': 'true',
        'enable_push_notifications': 'true',
        'enable_in_app_purchases': 'true',
        'enable_debug_menu': 'false',
        'enable_mock_data': 'false',
      },
    );
  }
  
  // Check if feature is enabled
  bool isFeatureEnabled(String featureName) {
    return featureFlags[featureName] == 'true';
  }
}