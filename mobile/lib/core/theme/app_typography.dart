import 'package:flutter/material.dart';

class AppTypography {
  // Font families
  static const String fontFamily = 'Inter';
  static const String fontFamilyMono = 'JetBrains Mono';
  
  // Text theme
  static TextTheme textTheme(ColorScheme colorScheme) {
    return TextTheme(
      // Display styles - for large screens
      displayLarge: TextStyle(
        fontSize: 57,
        fontWeight: FontWeight.w400,
        letterSpacing: -0.25,
        height: 1.12,
        color: colorScheme.onSurface,
      ),
      displayMedium: TextStyle(
        fontSize: 45,
        fontWeight: FontWeight.w400,
        letterSpacing: 0,
        height: 1.16,
        color: colorScheme.onSurface,
      ),
      displaySmall: TextStyle(
        fontSize: 36,
        fontWeight: FontWeight.w400,
        letterSpacing: 0,
        height: 1.22,
        color: colorScheme.onSurface,
      ),
      
      // Headline styles - for sections
      headlineLarge: TextStyle(
        fontSize: 32,
        fontWeight: FontWeight.w400,
        letterSpacing: 0,
        height: 1.25,
        color: colorScheme.onSurface,
      ),
      headlineMedium: TextStyle(
        fontSize: 28,
        fontWeight: FontWeight.w400,
        letterSpacing: 0,
        height: 1.29,
        color: colorScheme.onSurface,
      ),
      headlineSmall: TextStyle(
        fontSize: 24,
        fontWeight: FontWeight.w400,
        letterSpacing: 0,
        height: 1.33,
        color: colorScheme.onSurface,
      ),
      
      // Title styles - for app bars and cards
      titleLarge: TextStyle(
        fontSize: 22,
        fontWeight: FontWeight.w500,
        letterSpacing: 0,
        height: 1.27,
        color: colorScheme.onSurface,
      ),
      titleMedium: TextStyle(
        fontSize: 16,
        fontWeight: FontWeight.w500,
        letterSpacing: 0.15,
        height: 1.5,
        color: colorScheme.onSurface,
      ),
      titleSmall: TextStyle(
        fontSize: 14,
        fontWeight: FontWeight.w500,
        letterSpacing: 0.1,
        height: 1.43,
        color: colorScheme.onSurface,
      ),
      
      // Body styles - for content
      bodyLarge: TextStyle(
        fontSize: 16,
        fontWeight: FontWeight.w400,
        letterSpacing: 0.5,
        height: 1.5,
        color: colorScheme.onSurface,
      ),
      bodyMedium: TextStyle(
        fontSize: 14,
        fontWeight: FontWeight.w400,
        letterSpacing: 0.25,
        height: 1.43,
        color: colorScheme.onSurface,
      ),
      bodySmall: TextStyle(
        fontSize: 12,
        fontWeight: FontWeight.w400,
        letterSpacing: 0.4,
        height: 1.33,
        color: colorScheme.onSurfaceVariant,
      ),
      
      // Label styles - for buttons and chips
      labelLarge: TextStyle(
        fontSize: 14,
        fontWeight: FontWeight.w500,
        letterSpacing: 0.1,
        height: 1.43,
        color: colorScheme.onSurface,
      ),
      labelMedium: TextStyle(
        fontSize: 12,
        fontWeight: FontWeight.w500,
        letterSpacing: 0.5,
        height: 1.33,
        color: colorScheme.onSurface,
      ),
      labelSmall: TextStyle(
        fontSize: 11,
        fontWeight: FontWeight.w500,
        letterSpacing: 0.5,
        height: 1.45,
        color: colorScheme.onSurfaceVariant,
      ),
    );
  }
  
  // Custom text styles for special use cases
  static TextStyle monospace({
    double fontSize = 14,
    FontWeight fontWeight = FontWeight.w400,
    Color? color,
  }) {
    return TextStyle(
      fontFamily: fontFamilyMono,
      fontSize: fontSize,
      fontWeight: fontWeight,
      color: color,
      letterSpacing: 0,
      height: 1.5,
    );
  }
  
  static TextStyle currency({
    double fontSize = 24,
    FontWeight fontWeight = FontWeight.w600,
    Color? color,
  }) {
    return TextStyle(
      fontFamily: fontFamily,
      fontSize: fontSize,
      fontWeight: fontWeight,
      color: color,
      letterSpacing: -0.5,
      height: 1.2,
    );
  }
  
  static TextStyle percentage({
    double fontSize = 18,
    FontWeight fontWeight = FontWeight.w500,
    Color? color,
  }) {
    return TextStyle(
      fontFamily: fontFamily,
      fontSize: fontSize,
      fontWeight: fontWeight,
      color: color,
      letterSpacing: 0,
      height: 1.3,
    );
  }
  
  // Text style extensions
  static TextStyle withGradient(TextStyle style, List<Color> colors) {
    return style.copyWith(
      foreground: Paint()
        ..shader = LinearGradient(colors: colors).createShader(
          const Rect.fromLTWH(0.0, 0.0, 200.0, 70.0),
        ),
    );
  }
  
  static TextStyle withShadow(TextStyle style, {Color? shadowColor}) {
    return style.copyWith(
      shadows: [
        Shadow(
          offset: const Offset(0, 2),
          blurRadius: 4,
          color: shadowColor ?? Colors.black.withOpacity(0.25),
        ),
      ],
    );
  }
  
  static TextStyle withOutline(TextStyle style, {Color? outlineColor}) {
    return style.copyWith(
      shadows: [
        Shadow(
          offset: const Offset(-1.5, -1.5),
          color: outlineColor ?? Colors.white,
        ),
        Shadow(
          offset: const Offset(1.5, -1.5),
          color: outlineColor ?? Colors.white,
        ),
        Shadow(
          offset: const Offset(1.5, 1.5),
          color: outlineColor ?? Colors.white,
        ),
        Shadow(
          offset: const Offset(-1.5, 1.5),
          color: outlineColor ?? Colors.white,
        ),
      ],
    );
  }
}