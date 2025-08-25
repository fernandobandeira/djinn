enum Environment { dev, staging, prod }

class AppConfig {
  static late Environment _env;
  static late String _apiUrl;
  
  static void initialize() {
    const envString = String.fromEnvironment('ENV', defaultValue: 'dev');
    _env = Environment.values.byName(envString);
    
    _apiUrl = switch (_env) {
      Environment.dev => 'https://api-dev.djinn.app/graphql',
      Environment.staging => 'https://api-staging.djinn.app/graphql',
      Environment.prod => 'https://api.djinn.app/graphql',
    };
  }
  
  static String get apiUrl => _apiUrl;
  static bool get isDev => _env == Environment.dev;
  static bool get isProduction => _env == Environment.prod;
  static bool get isStaging => _env == Environment.staging;
  static Environment get currentEnv => _env;
}