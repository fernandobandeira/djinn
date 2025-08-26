import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class ForgotPasswordScreen extends StatelessWidget {
  const ForgotPasswordScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Forgot Password'),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            const Text('Forgot Password Screen'),
            const SizedBox(height: 16),
            ElevatedButton(
              onPressed: () => context.go('/auth/login'),
              child: const Text('Back to Login'),
            ),
          ],
        ),
      ),
    );
  }
}