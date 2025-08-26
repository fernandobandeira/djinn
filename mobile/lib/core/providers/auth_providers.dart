import 'package:flutter_riverpod/flutter_riverpod.dart';

// Authentication state provider
final authStateProvider = StateNotifierProvider<AuthStateNotifier, AuthState>((ref) {
  return AuthStateNotifier();
});

class AuthState {
  const AuthState({
    this.isAuthenticated = false,
    this.user,
    this.isLoading = false,
    this.error,
  });

  final bool isAuthenticated;
  final User? user;
  final bool isLoading;
  final String? error;

  AuthState copyWith({
    bool? isAuthenticated,
    User? user,
    bool? isLoading,
    String? error,
  }) {
    return AuthState(
      isAuthenticated: isAuthenticated ?? this.isAuthenticated,
      user: user ?? this.user,
      isLoading: isLoading ?? this.isLoading,
      error: error ?? this.error,
    );
  }
}

class User {
  const User({
    required this.id,
    required this.email,
    this.name,
  });

  final String id;
  final String email;
  final String? name;
}

class AuthStateNotifier extends StateNotifier<AuthState> {
  AuthStateNotifier() : super(const AuthState());

  Future<void> signIn(String email, String password) async {
    state = state.copyWith(isLoading: true, error: null);
    
    try {
      // Simulate authentication
      await Future.delayed(const Duration(seconds: 1));
      
      if (email.isNotEmpty && password.isNotEmpty) {
        final user = User(
          id: 'user_123',
          email: email,
          name: 'Demo User',
        );
        
        state = state.copyWith(
          isAuthenticated: true,
          user: user,
          isLoading: false,
        );
      } else {
        throw Exception('Invalid credentials');
      }
    } catch (e) {
      state = state.copyWith(
        isLoading: false,
        error: e.toString(),
      );
    }
  }

  Future<void> signOut() async {
    state = state.copyWith(isLoading: true);
    
    await Future.delayed(const Duration(milliseconds: 500));
    
    state = const AuthState(
      isAuthenticated: false,
      isLoading: false,
    );
  }

  void clearError() {
    state = state.copyWith(error: null);
  }
}

// Derived provider for auth token
final authTokenProvider = Provider<String?>((ref) {
  final authState = ref.watch(authStateProvider);
  
  // TODO: When Firebase Auth is implemented, retrieve actual JWT token
  // For now, return null to prevent security issues
  // In production, this should:
  // 1. Check secure storage for cached token
  // 2. Validate token expiration
  // 3. Refresh if needed
  // 4. Return valid JWT token
  
  if (!authState.isAuthenticated) {
    return null;
  }
  
  // Placeholder for secure token retrieval
  // final secureStorage = ref.read(secureStorageProvider);
  // return await secureStorage.read(key: 'auth_token');
  
  return null; // Return null until proper auth is implemented
});

// Derived provider for checking if user is authenticated
final isAuthenticatedProvider = Provider<bool>((ref) {
  return ref.watch(authStateProvider).isAuthenticated;
});