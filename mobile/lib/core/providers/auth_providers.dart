import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:djinn_mobile/core/utils/validators.dart';
import 'package:djinn_mobile/core/utils/logger.dart';

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
    this.photoUrl,
    this.phoneNumber,
    this.emailVerified = false,
  });

  final String id;
  final String email;
  final String? name;
  final String? photoUrl;
  final String? phoneNumber;
  final bool emailVerified;
}

// Alias for compatibility
typedef AuthUser = User;

class AuthStateNotifier extends StateNotifier<AuthState> {
  AuthStateNotifier() : super(const AuthState());

  Future<void> signIn(String email, String password) async {
    state = state.copyWith(isLoading: true, error: null);
    
    try {
      // Validate input
      final emailError = Validators.validateEmail(email);
      if (emailError != null) {
        throw Exception(emailError);
      }
      
      final passwordError = Validators.validatePassword(password);
      if (passwordError != null) {
        throw Exception(passwordError);
      }
      
      // Sanitize inputs
      final sanitizedEmail = Validators.sanitizeInput(email.trim().toLowerCase());
      
      logger.info('Attempting sign in for user: $sanitizedEmail', tag: 'AuthStateNotifier');
      
      // TODO: Implement actual authentication with Firebase
      // For now, simulate authentication
      await Future.delayed(const Duration(seconds: 1));
      
      // In production, this would make an API call to authenticate
      // final response = await authService.signIn(sanitizedEmail, password);
      
      // Demo user for development
      final user = User(
        id: 'user_${DateTime.now().millisecondsSinceEpoch}',
        email: sanitizedEmail,
        name: 'Demo User',
      );
      
      state = state.copyWith(
        isAuthenticated: true,
        user: user,
        isLoading: false,
      );
      
      logger.info('Sign in successful for user: ${user.id}', tag: 'AuthStateNotifier');
    } catch (e, stackTrace) {
      logger.error(
        'Sign in failed',
        tag: 'AuthStateNotifier',
        error: e,
        stackTrace: stackTrace,
      );
      
      state = state.copyWith(
        isLoading: false,
        error: e.toString().replaceAll('Exception: ', ''),
      );
    }
  }

  Future<void> signOut() async {
    state = state.copyWith(isLoading: true);
    
    try {
      logger.info('User signing out', tag: 'AuthStateNotifier');
      
      // TODO: Clear secure storage tokens
      // await secureStorage.delete(key: 'auth_token');
      
      await Future.delayed(const Duration(milliseconds: 500));
      
      state = const AuthState(
        isAuthenticated: false,
        isLoading: false,
      );
      
      logger.info('Sign out successful', tag: 'AuthStateNotifier');
    } catch (e, stackTrace) {
      logger.error(
        'Sign out failed',
        tag: 'AuthStateNotifier',
        error: e,
        stackTrace: stackTrace,
      );
      
      // Force sign out even if there's an error
      state = const AuthState(
        isAuthenticated: false,
        isLoading: false,
      );
    }
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