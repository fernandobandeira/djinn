import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:djinn_mobile/core/database/app_database.dart';
import 'package:djinn_mobile/core/database/daos/user_dao.dart';
import 'package:djinn_mobile/core/database/daos/settings_dao.dart';
import 'package:djinn_mobile/core/providers/auth_providers.dart';

// Database instance provider
final databaseProvider = Provider<AppDatabase>((ref) {
  final database = AppDatabase();
  
  // Close database on dispose
  ref.onDispose(() {
    database.close();
  });
  
  return database;
});

// DAO providers
final userDaoProvider = Provider<UserDao>((ref) {
  final database = ref.watch(databaseProvider);
  return UserDao(database);
});

final settingsDaoProvider = Provider<SettingsDao>((ref) {
  final database = ref.watch(databaseProvider);
  return SettingsDao(database);
});

// Current user profile provider
final currentUserProfileProvider = StreamProvider<UserProfile?>((ref) async* {
  final userDao = ref.watch(userDaoProvider);
  final authState = ref.watch(authStateProvider);
  
  // Only fetch user profile when authenticated
  if (authState.isAuthenticated && authState.user != null) {
    final userId = authState.user!.id;
    yield* userDao.watchById(userId);
  } else {
    // Return null if not authenticated
    yield* Stream.value(null);
  }
});

// Settings providers
final themeModeProvider = StreamProvider<String>((ref) {
  final settingsDao = ref.watch(settingsDaoProvider);
  return settingsDao.watchThemeMode();
});

final notificationsEnabledProvider = StreamProvider<bool>((ref) {
  final settingsDao = ref.watch(settingsDaoProvider);
  return settingsDao.watchNotificationsEnabled();
});

// Database initialization provider
final databaseInitializerProvider = FutureProvider<void>((ref) async {
  final database = ref.watch(databaseProvider);
  
  // Perform any initialization tasks
  // This runs once when the app starts
  
  // Check if this is first run
  final settingsDao = ref.read(settingsDaoProvider);
  final isFirstRun = await settingsDao.getSetting('first_run');
  
  if (isFirstRun == null) {
    // First time app launch
    await settingsDao.setSetting('first_run', 'false');
    await settingsDao.setSetting('app_version', '1.0.0');
    await settingsDao.setSetting('install_date', DateTime.now().toIso8601String());
  }
});

// Database maintenance provider
final databaseMaintenanceProvider = Provider((ref) => DatabaseMaintenance(ref));

class DatabaseMaintenance {
  final Ref _ref;
  
  DatabaseMaintenance(this._ref);
  
  AppDatabase get _database => _ref.read(databaseProvider);
  
  // Clear all user data (for logout)
  Future<void> clearUserData() async {
    await _database.clearAllData();
  }
  
  // Vacuum database to optimize storage
  Future<void> optimizeDatabase() async {
    await _database.vacuum();
  }
  
  // Get database size in bytes
  Future<int> getDatabaseSize() async {
    return await _database.getDatabaseSize();
  }
  
  // Export database for backup
  Future<Map<String, dynamic>> exportDatabase() async {
    final userDao = _ref.read(userDaoProvider);
    final settingsDao = _ref.read(settingsDaoProvider);
    
    final users = await userDao.getAll();
    final settings = await settingsDao.getAllSettings();
    
    return {
      'version': 1,
      'exported_at': DateTime.now().toIso8601String(),
      'users': users.map((u) => {
        'id': u.id,
        'email': u.email,
        'name': u.name,
        'avatar_url': u.avatarUrl,
        'phone_number': u.phoneNumber,
        'email_verified_at': u.emailVerifiedAt?.toIso8601String(),
        'created_at': u.createdAt.toIso8601String(),
        'updated_at': u.updatedAt.toIso8601String(),
        'metadata': u.metadata,
      }).toList(),
      'settings': settings,
    };
  }
  
  // Import database from backup
  Future<void> importDatabase(Map<String, dynamic> data) async {
    if (data['version'] != 1) {
      throw Exception('Unsupported backup version');
    }
    
    final userDao = _ref.read(userDaoProvider);
    final settingsDao = _ref.read(settingsDaoProvider);
    
    // Clear existing data
    await clearUserData();
    
    // Import users
    final users = data['users'] as List;
    for (final userData in users) {
      await userDao.upsertUser(UserProfile(
        id: userData['id'],
        email: userData['email'],
        name: userData['name'],
        avatarUrl: userData['avatar_url'],
        phoneNumber: userData['phone_number'],
        emailVerifiedAt: userData['email_verified_at'] != null 
          ? DateTime.parse(userData['email_verified_at'])
          : null,
        createdAt: DateTime.parse(userData['created_at']),
        updatedAt: DateTime.parse(userData['updated_at']),
        metadata: userData['metadata'],
      ));
    }
    
    // Import settings
    final settings = data['settings'] as Map<String, dynamic>;
    await settingsDao.setSettings(settings.cast<String, String>());
  }
}