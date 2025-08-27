import 'package:flutter_riverpod/flutter_riverpod.dart';

// Email validation provider - demonstrates computed providers
final emailValidationProvider = Provider.family<ValidationResult, String>((ref, email) {
  if (email.isEmpty) {
    return const ValidationResult(isValid: false, message: 'Email is required');
  }
  
  final emailRegex = RegExp(r'^[^@]+@[^@]+\.[^@]+');
  if (!emailRegex.hasMatch(email)) {
    return const ValidationResult(isValid: false, message: 'Invalid email format');
  }
  
  return const ValidationResult(isValid: true);
});

// Password validation provider
final passwordValidationProvider = Provider.family<ValidationResult, String>((ref, password) {
  if (password.isEmpty) {
    return const ValidationResult(isValid: false, message: 'Password is required');
  }
  
  if (password.length < 6) {
    return const ValidationResult(isValid: false, message: 'Password must be at least 6 characters');
  }
  
  return const ValidationResult(isValid: true);
});

// Form validation provider - demonstrates combining multiple providers
final loginFormValidationProvider = Provider.family<ValidationResult, LoginFormData>((ref, formData) {
  final emailValidation = ref.watch(emailValidationProvider(formData.email));
  final passwordValidation = ref.watch(passwordValidationProvider(formData.password));
  
  if (!emailValidation.isValid) {
    return emailValidation;
  }
  
  if (!passwordValidation.isValid) {
    return passwordValidation;
  }
  
  return const ValidationResult(isValid: true, message: 'Form is valid');
});

class ValidationResult {
  const ValidationResult({
    required this.isValid,
    this.message,
  });
  
  final bool isValid;
  final String? message;
}

class LoginFormData {
  const LoginFormData({
    required this.email,
    required this.password,
  });
  
  final String email;
  final String password;
}

// Counter provider - demonstrates simple state management
final counterProvider = StateProvider<int>((ref) => 0);

// Async provider example - demonstrates FutureProvider
final fetchUserDataProvider = FutureProvider.family<UserData?, String>((ref, userId) async {
  // Simulate API call
  await Future.delayed(const Duration(seconds: 2));
  
  if (userId.isEmpty) {
    return null;
  }
  
  return UserData(
    id: userId,
    email: 'user@example.com',
    name: 'Demo User $userId',
  );
});

class UserData {
  const UserData({
    required this.id,
    required this.email,
    this.name,
  });

  final String id;
  final String email;
  final String? name;
}