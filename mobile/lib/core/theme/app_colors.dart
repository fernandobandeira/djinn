import 'package:flutter/material.dart';

class AppColors {
  // Primary palette (Purple - Djinn magical theme)
  static const MaterialColor primary = MaterialColor(
    0xFF6B46C1,
    <int, Color>{
      50: Color(0xFFF3F0FA),
      100: Color(0xFFE8E2F5),
      200: Color(0xFFD1C5EB),
      300: Color(0xFFB9A8E1),
      400: Color(0xFFA28BD7),
      500: Color(0xFF6B46C1), // Primary
      600: Color(0xFF5635A1),
      700: Color(0xFF412881),
      800: Color(0xFF2C1B61),
      900: Color(0xFF170E41),
    },
  );
  
  // Secondary palette (Blue)
  static const MaterialColor secondary = MaterialColor(
    0xFF3B82F6,
    <int, Color>{
      50: Color(0xFFEFF6FF),
      100: Color(0xFFDBEAFE),
      200: Color(0xFFBFDBFE),
      300: Color(0xFF93C5FD),
      400: Color(0xFF60A5FA),
      500: Color(0xFF3B82F6), // Secondary
      600: Color(0xFF2563EB),
      700: Color(0xFF1D4ED8),
      800: Color(0xFF1E40AF),
      900: Color(0xFF1E3A8A),
    },
  );
  
  // Success palette (Green)
  static const MaterialColor success = MaterialColor(
    0xFF10B981,
    <int, Color>{
      50: Color(0xFFF0FDF4),
      100: Color(0xFFDCFCE7),
      200: Color(0xFFBBF7D0),
      300: Color(0xFF86EFAC),
      400: Color(0xFF4ADE80),
      500: Color(0xFF10B981), // Success
      600: Color(0xFF059669),
      700: Color(0xFF047857),
      800: Color(0xFF065F46),
      900: Color(0xFF064E3B),
    },
  );
  
  // Error palette (Red)
  static const MaterialColor error = MaterialColor(
    0xFFEF4444,
    <int, Color>{
      50: Color(0xFFFEF2F2),
      100: Color(0xFFFEE2E2),
      200: Color(0xFFFECACA),
      300: Color(0xFFFCA5A5),
      400: Color(0xFFF87171),
      500: Color(0xFFEF4444), // Error
      600: Color(0xFFDC2626),
      700: Color(0xFFB91C1C),
      800: Color(0xFF991B1B),
      900: Color(0xFF7F1D1D),
    },
  );
  
  // Warning palette (Amber)
  static const MaterialColor warning = MaterialColor(
    0xFFF59E0B,
    <int, Color>{
      50: Color(0xFFFFFBEB),
      100: Color(0xFFFEF3C7),
      200: Color(0xFFFDE68A),
      300: Color(0xFFFCD34D),
      400: Color(0xFFFBBF24),
      500: Color(0xFFF59E0B), // Warning
      600: Color(0xFFD97706),
      700: Color(0xFFB45309),
      800: Color(0xFF92400E),
      900: Color(0xFF78350F),
    },
  );
  
  // Info palette (Cyan)
  static const MaterialColor info = MaterialColor(
    0xFF06B6D4,
    <int, Color>{
      50: Color(0xFFECFEFF),
      100: Color(0xFFCFFAFE),
      200: Color(0xFFA5F3FC),
      300: Color(0xFF67E8F9),
      400: Color(0xFF22D3EE),
      500: Color(0xFF06B6D4), // Info
      600: Color(0xFF0891B2),
      700: Color(0xFF0E7490),
      800: Color(0xFF155E75),
      900: Color(0xFF164E63),
    },
  );
  
  // Neutral palette (Gray)
  static const MaterialColor neutral = MaterialColor(
    0xFF6B7280,
    <int, Color>{
      50: Color(0xFFF9FAFB),
      100: Color(0xFFF3F4F6),
      200: Color(0xFFE5E7EB),
      300: Color(0xFFD1D5DB),
      400: Color(0xFF9CA3AF),
      500: Color(0xFF6B7280), // Neutral
      600: Color(0xFF4B5563),
      700: Color(0xFF374151),
      800: Color(0xFF1F2937),
      900: Color(0xFF111827),
    },
  );
  
  // Surface colors
  static const Color surface = Color(0xFFFAFAFA);
  static const Color surfaceDark = Color(0xFF1A1A1A);
  static const Color background = Color(0xFFFFFFFF);
  static const Color backgroundDark = Color(0xFF121212);
  
  // Gradient colors
  static const List<Color> primaryGradient = [
    Color(0xFF6B46C1),
    Color(0xFF3B82F6),
  ];
  
  static const List<Color> successGradient = [
    Color(0xFF10B981),
    Color(0xFF06B6D4),
  ];
  
  static const List<Color> errorGradient = [
    Color(0xFFEF4444),
    Color(0xFFF59E0B),
  ];
  
  // Special effect colors (for Djinn magical theme)
  static const Color shimmer = Color(0xFFE0D4F7);
  static const Color glow = Color(0xFFBB86FC);
  static const Color sparkle = Color(0xFFFFD700);
  
  // Semantic colors
  static const Color textPrimary = Color(0xFF1F2937);
  static const Color textSecondary = Color(0xFF6B7280);
  static const Color textTertiary = Color(0xFF9CA3AF);
  static const Color textDisabled = Color(0xFFD1D5DB);
  
  static const Color textPrimaryDark = Color(0xFFF3F4F6);
  static const Color textSecondaryDark = Color(0xFFD1D5DB);
  static const Color textTertiaryDark = Color(0xFF9CA3AF);
  static const Color textDisabledDark = Color(0xFF4B5563);
  
  // Border colors
  static const Color border = Color(0xFFE5E7EB);
  static const Color borderDark = Color(0xFF374151);
  
  // Shadow colors
  static const Color shadow = Color(0x1A000000);
  static const Color shadowDark = Color(0x33000000);
}