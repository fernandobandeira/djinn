import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:djinn_mobile/core/theme/app_theme.dart';
import 'package:djinn_mobile/core/theme/app_colors.dart';

void main() {
  group('AppTheme', () {
    late AppTheme appTheme;
    
    setUp(() {
      appTheme = AppTheme();
    });
    
    group('Light Theme', () {
      test('should have correct brightness', () {
        final theme = appTheme.lightTheme;
        expect(theme.brightness, Brightness.light);
      });
      
      test('should use Material 3', () {
        final theme = appTheme.lightTheme;
        expect(theme.useMaterial3, true);
      });
      
      test('should have correct primary color', () {
        final theme = appTheme.lightTheme;
        expect(theme.colorScheme.primary, AppTheme.primaryColor);
      });
      
      test('should have correct secondary color', () {
        final theme = appTheme.lightTheme;
        expect(theme.colorScheme.secondary, AppTheme.secondaryColor);
      });
      
      test('should have transparent app bar', () {
        final theme = appTheme.lightTheme;
        expect(theme.appBarTheme.backgroundColor, Colors.transparent);
        expect(theme.appBarTheme.elevation, 0);
      });
      
      test('should have rounded cards', () {
        final theme = appTheme.lightTheme;
        final cardTheme = theme.cardTheme;
        expect(cardTheme.shape, isA<RoundedRectangleBorder>());
        
        final shape = cardTheme.shape as RoundedRectangleBorder;
        expect(shape.borderRadius, BorderRadius.circular(16));
      });
      
      test('should have rounded buttons', () {
        final theme = appTheme.lightTheme;
        final elevatedButtonTheme = theme.elevatedButtonTheme.style;
        
        expect(elevatedButtonTheme, isNotNull);
        final shape = elevatedButtonTheme!.shape?.resolve({});
        expect(shape, isA<RoundedRectangleBorder>());
      });
    });
    
    group('Dark Theme', () {
      test('should have correct brightness', () {
        final theme = appTheme.darkTheme;
        expect(theme.brightness, Brightness.dark);
      });
      
      test('should use Material 3', () {
        final theme = appTheme.darkTheme;
        expect(theme.useMaterial3, true);
      });
      
      test('should have correct primary color', () {
        final theme = appTheme.darkTheme;
        expect(theme.colorScheme.primary, AppTheme.primaryColor);
      });
      
      test('should have transparent app bar', () {
        final theme = appTheme.darkTheme;
        expect(theme.appBarTheme.backgroundColor, Colors.transparent);
        expect(theme.appBarTheme.elevation, 0);
      });
    });
    
    group('ThemeMode Provider', () {
      testWidgets('should default to system theme', (WidgetTester tester) async {
        await tester.pumpWidget(
          ProviderScope(
            child: Consumer(
              builder: (context, ref, child) {
                final themeMode = ref.watch(themeModeProvider);
                return MaterialApp(
                  home: Scaffold(
                    body: Text(themeMode.toString()),
                  ),
                );
              },
            ),
          ),
        );
        
        expect(find.text('ThemeMode.system'), findsOneWidget);
      });
      
      testWidgets('should allow theme mode changes', (WidgetTester tester) async {
        await tester.pumpWidget(
          ProviderScope(
            child: Consumer(
              builder: (context, ref, child) {
                final themeMode = ref.watch(themeModeProvider);
                
                return MaterialApp(
                  home: Scaffold(
                    body: Column(
                      children: [
                        Text(themeMode.toString()),
                        ElevatedButton(
                          onPressed: () {
                            ref.read(themeModeProvider.notifier).state = ThemeMode.dark;
                          },
                          child: const Text('Dark Mode'),
                        ),
                        ElevatedButton(
                          onPressed: () {
                            ref.read(themeModeProvider.notifier).state = ThemeMode.light;
                          },
                          child: const Text('Light Mode'),
                        ),
                      ],
                    ),
                  ),
                );
              },
            ),
          ),
        );
        
        // Initially system mode
        expect(find.text('ThemeMode.system'), findsOneWidget);
        
        // Switch to dark mode
        await tester.tap(find.text('Dark Mode'));
        await tester.pump();
        expect(find.text('ThemeMode.dark'), findsOneWidget);
        
        // Switch to light mode
        await tester.tap(find.text('Light Mode'));
        await tester.pump();
        expect(find.text('ThemeMode.light'), findsOneWidget);
      });
    });
  });
  
  group('AppColors', () {
    test('should have correct primary color values', () {
      expect(AppColors.primary.shade50, const Color(0xFFF3F0FA));
      expect(AppColors.primary.shade500, const Color(0xFF6B46C1));
      expect(AppColors.primary.shade900, const Color(0xFF170E41));
    });
    
    test('should have correct secondary color values', () {
      expect(AppColors.secondary.shade50, const Color(0xFFEFF6FF));
      expect(AppColors.secondary.shade500, const Color(0xFF3B82F6));
      expect(AppColors.secondary.shade900, const Color(0xFF1E3A8A));
    });
    
    test('should have gradient colors', () {
      expect(AppColors.primaryGradient.length, 2);
      expect(AppColors.primaryGradient[0], const Color(0xFF6B46C1));
      expect(AppColors.primaryGradient[1], const Color(0xFF3B82F6));
    });
    
    test('should have special effect colors', () {
      expect(AppColors.shimmer, const Color(0xFFE0D4F7));
      expect(AppColors.glow, const Color(0xFFBB86FC));
      expect(AppColors.sparkle, const Color(0xFFFFD700));
    });
  });
}