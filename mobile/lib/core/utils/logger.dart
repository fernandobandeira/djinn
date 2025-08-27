import 'package:flutter/foundation.dart';

enum LogLevel {
  verbose,
  debug,
  info,
  warning,
  error,
}

class AppLogger {
  static final AppLogger _instance = AppLogger._internal();
  factory AppLogger() => _instance;
  AppLogger._internal();
  
  static const String _appName = 'Djinn';
  
  // Only log in debug mode by default
  bool _enableLogging = kDebugMode;
  LogLevel _minLevel = LogLevel.debug;
  
  void configure({bool enableLogging = true, LogLevel minLevel = LogLevel.debug}) {
    _enableLogging = enableLogging;
    _minLevel = minLevel;
  }
  
  void verbose(String message, {String? tag, Object? error, StackTrace? stackTrace}) {
    _log(LogLevel.verbose, message, tag: tag, error: error, stackTrace: stackTrace);
  }
  
  void debug(String message, {String? tag, Object? error, StackTrace? stackTrace}) {
    _log(LogLevel.debug, message, tag: tag, error: error, stackTrace: stackTrace);
  }
  
  void info(String message, {String? tag, Object? error, StackTrace? stackTrace}) {
    _log(LogLevel.info, message, tag: tag, error: error, stackTrace: stackTrace);
  }
  
  void warning(String message, {String? tag, Object? error, StackTrace? stackTrace}) {
    _log(LogLevel.warning, message, tag: tag, error: error, stackTrace: stackTrace);
  }
  
  void error(String message, {String? tag, Object? error, StackTrace? stackTrace}) {
    _log(LogLevel.error, message, tag: tag, error: error, stackTrace: stackTrace);
  }
  
  void _log(
    LogLevel level,
    String message, {
    String? tag,
    Object? error,
    StackTrace? stackTrace,
  }) {
    if (!_enableLogging || level.index < _minLevel.index) return;
    
    final timestamp = DateTime.now().toIso8601String();
    final levelStr = level.toString().split('.').last.toUpperCase();
    final tagStr = tag != null ? '[$tag]' : '';
    
    final logMessage = '[$timestamp][$_appName][$levelStr]$tagStr $message';
    
    // In debug mode, use debugPrint to avoid truncation
    if (kDebugMode) {
      debugPrint(logMessage);
      
      if (error != null) {
        debugPrint('Error: $error');
      }
      
      if (stackTrace != null && level == LogLevel.error) {
        debugPrint('StackTrace:\n$stackTrace');
      }
    } else {
      // In release mode, send to crash reporting service
      // TODO: Integrate with Sentry or Firebase Crashlytics
      // if (level == LogLevel.error && error != null) {
      //   Sentry.captureException(error, stackTrace: stackTrace);
      // }
    }
  }
}

// Global logger instance
final logger = AppLogger();

// Extension for easier logging from any class
extension LoggerExtension on Object {
  void logDebug(String message) {
    logger.debug(message, tag: runtimeType.toString());
  }
  
  void logInfo(String message) {
    logger.info(message, tag: runtimeType.toString());
  }
  
  void logWarning(String message) {
    logger.warning(message, tag: runtimeType.toString());
  }
  
  void logError(String message, {Object? error, StackTrace? stackTrace}) {
    logger.error(message, tag: runtimeType.toString(), error: error, stackTrace: stackTrace);
  }
}