# Environment Configuration

## Overview
The app supports three environments: Development, Staging, and Production.

## Running Different Environments

### Development (Default)
```bash
flutter run
# or explicitly:
flutter run --dart-define=ENV=dev
```

### Staging
```bash
flutter run --dart-define=ENV=staging
```

### Production
```bash
flutter run --dart-define=ENV=prod
```

## Building for Different Environments

### iOS
```bash
# Development
flutter build ios --dart-define=ENV=dev

# Staging
flutter build ios --dart-define=ENV=staging

# Production
flutter build ios --dart-define=ENV=prod --release
```

### Android
```bash
# Development
flutter build apk --dart-define=ENV=dev

# Staging  
flutter build apk --dart-define=ENV=staging

# Production
flutter build apk --dart-define=ENV=prod --release
flutter build appbundle --dart-define=ENV=prod --release
```

## Environment Variables

Each environment configures:
- API endpoints (REST and WebSocket)
- Feature flags
- Logging and monitoring settings
- Timeout and retry configurations
- Third-party service configurations (Sentry, Firebase, etc.)

## Feature Flags

Access feature flags in code:
```dart
if (AppConfig.config.isFeatureEnabled('enable_firebase_auth')) {
  // Firebase auth logic
}
```

## Adding New Environment Variables

1. Add the property to `EnvironmentConfig` class
2. Update all factory constructors (dev, staging, prod)
3. Add a getter in `AppConfig` if needed
4. Document the new variable here

## Security Notes

- Never commit sensitive keys directly in code
- Use `--dart-define-from-file` for sensitive values:
  ```bash
  flutter run --dart-define-from-file=env/dev.json
  ```
  
Example env/dev.json:
```json
{
  "ENV": "dev",
  "API_KEY": "your-secret-key"
}
```