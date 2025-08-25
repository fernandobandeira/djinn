import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:djinn_mobile/core/config/environment.dart';

// Environment provider - provides static access to AppConfig
final environmentProvider = Provider<Environment>((ref) {
  return AppConfig.currentEnv;
});

// App state provider for global app state
final appStateProvider = StateNotifierProvider<AppStateNotifier, AppState>((ref) {
  return AppStateNotifier();
});

class AppState {
  const AppState({
    this.isInitialized = false,
    this.isLoading = false,
    this.error,
  });

  final bool isInitialized;
  final bool isLoading;
  final String? error;

  AppState copyWith({
    bool? isInitialized,
    bool? isLoading,
    String? error,
  }) {
    return AppState(
      isInitialized: isInitialized ?? this.isInitialized,
      isLoading: isLoading ?? this.isLoading,
      error: error ?? this.error,
    );
  }
}

class AppStateNotifier extends StateNotifier<AppState> {
  AppStateNotifier() : super(const AppState());

  void initialize() {
    state = state.copyWith(isLoading: true);
    // Simulate initialization
    Future.delayed(const Duration(seconds: 1), () {
      state = state.copyWith(
        isInitialized: true,
        isLoading: false,
      );
    });
  }

  void setError(String error) {
    state = state.copyWith(error: error, isLoading: false);
  }

  void clearError() {
    state = state.copyWith(error: null);
  }
}

// Theme mode provider
final themeModeProvider = StateProvider<bool>((ref) {
  return false; // false = light, true = dark
});

// Development tools provider
final devToolsProvider = Provider<bool>((ref) {
  return AppConfig.isDev;
});