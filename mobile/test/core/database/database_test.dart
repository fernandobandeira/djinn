import 'package:flutter_test/flutter_test.dart';
import 'package:djinn_mobile/core/database/app_database.dart';
import 'package:djinn_mobile/core/database/daos/user_dao.dart';
import 'package:djinn_mobile/core/database/daos/settings_dao.dart';
import 'package:drift/native.dart';
import 'package:drift/drift.dart';

void main() {
  late AppDatabase database;
  late UserDao userDao;
  late SettingsDao settingsDao;
  
  setUp(() {
    // Create in-memory database for testing
    database = AppDatabase.forTesting(NativeDatabase.memory());
    userDao = UserDao(database);
    settingsDao = SettingsDao(database);
  });
  
  tearDown(() async {
    await database.close();
  });
  
  group('UserDao Tests', () {
    test('should insert and retrieve a user', () async {
      const userId = 'test-user-1';
      final user = UserProfile(
        id: userId,
        email: 'test@example.com',
        name: 'Test User',
        avatarUrl: 'https://example.com/avatar.jpg',
        phoneNumber: '+1234567890',
        emailVerifiedAt: DateTime.now(),
        createdAt: DateTime.now(),
        updatedAt: DateTime.now(),
        metadata: '{"key": "value"}',
      );
      
      await userDao.upsertUser(user);
      final retrieved = await userDao.getById(userId);
      
      expect(retrieved, isNotNull);
      expect(retrieved!.id, userId);
      expect(retrieved.email, 'test@example.com');
      expect(retrieved.name, 'Test User');
    });
    
    test('should update an existing user', () async {
      const userId = 'test-user-2';
      final user = UserProfile(
        id: userId,
        email: 'old@example.com',
        name: 'Old Name',
        avatarUrl: null,
        phoneNumber: null,
        emailVerifiedAt: null,
        createdAt: DateTime.now(),
        updatedAt: DateTime.now(),
        metadata: null,
      );
      
      await userDao.upsertUser(user);
      
      final updatedUser = user.copyWith(
        email: 'new@example.com',
        name: const Value('New Name'),
      );
      
      await userDao.upsertUser(updatedUser);
      final retrieved = await userDao.getById(userId);
      
      expect(retrieved!.email, 'new@example.com');
      expect(retrieved.name, 'New Name');
    });
    
    test('should delete a user', () async {
      const userId = 'test-user-3';
      final user = UserProfile(
        id: userId,
        email: 'delete@example.com',
        name: 'Delete Me',
        avatarUrl: null,
        phoneNumber: null,
        emailVerifiedAt: null,
        createdAt: DateTime.now(),
        updatedAt: DateTime.now(),
        metadata: null,
      );
      
      await userDao.upsertUser(user);
      await userDao.deleteUser(userId);
      
      final retrieved = await userDao.getById(userId);
      expect(retrieved, isNull);
    });
    
    test('should search users by name or email', () async {
      final users = [
        UserProfile(
          id: '1',
          email: 'john@example.com',
          name: 'John Doe',
          avatarUrl: null,
          phoneNumber: null,
          emailVerifiedAt: null,
          createdAt: DateTime.now(),
          updatedAt: DateTime.now(),
          metadata: null,
        ),
        UserProfile(
          id: '2',
          email: 'jane@example.com',
          name: 'Jane Smith',
          avatarUrl: null,
          phoneNumber: null,
          emailVerifiedAt: null,
          createdAt: DateTime.now(),
          updatedAt: DateTime.now(),
          metadata: null,
        ),
        UserProfile(
          id: '3',
          email: 'bob@test.com',
          name: 'Bob Johnson',
          avatarUrl: null,
          phoneNumber: null,
          emailVerifiedAt: null,
          createdAt: DateTime.now(),
          updatedAt: DateTime.now(),
          metadata: null,
        ),
      ];
      
      for (final user in users) {
        await userDao.upsertUser(user);
      }
      
      // Search by name
      final johnResults = await userDao.searchUsers('John');
      expect(johnResults.length, 2); // John Doe and Bob Johnson
      
      // Search by email
      final exampleResults = await userDao.searchUsers('example.com');
      expect(exampleResults.length, 2); // john@example.com and jane@example.com
    });
    
    test('should watch user changes', () async {
      const userId = 'watch-user';
      final user = UserProfile(
        id: userId,
        email: 'watch@example.com',
        name: 'Watch Me',
        avatarUrl: null,
        phoneNumber: null,
        emailVerifiedAt: null,
        createdAt: DateTime.now(),
        updatedAt: DateTime.now(),
        metadata: null,
      );
      
      final stream = userDao.watchById(userId);
      final changes = <UserProfile?>[];
      
      final subscription = stream.listen(changes.add);
      
      // Initially null
      await Future.delayed(const Duration(milliseconds: 100));
      expect(changes.last, isNull);
      
      // Insert user
      await userDao.upsertUser(user);
      await Future.delayed(const Duration(milliseconds: 100));
      expect(changes.last?.id, userId);
      
      // Update user
      final updated = user.copyWith(name: const Value('Updated Name'));
      await userDao.upsertUser(updated);
      await Future.delayed(const Duration(milliseconds: 100));
      expect(changes.last?.name, 'Updated Name');
      
      await subscription.cancel();
    });
  });
  
  group('SettingsDao Tests', () {
    test('should set and get a setting', () async {
      await settingsDao.setSetting('test_key', 'test_value');
      final value = await settingsDao.getSetting('test_key');
      
      expect(value, 'test_value');
    });
    
    test('should update an existing setting', () async {
      await settingsDao.setSetting('update_key', 'old_value');
      await settingsDao.setSetting('update_key', 'new_value');
      
      final value = await settingsDao.getSetting('update_key');
      expect(value, 'new_value');
    });
    
    test('should delete a setting', () async {
      await settingsDao.setSetting('delete_key', 'value');
      await settingsDao.deleteSetting('delete_key');
      
      final value = await settingsDao.getSetting('delete_key');
      expect(value, isNull);
    });
    
    test('should get multiple settings', () async {
      await settingsDao.setSettings({
        'key1': 'value1',
        'key2': 'value2',
        'key3': 'value3',
      });
      
      final settings = await settingsDao.getSettings(['key1', 'key3']);
      expect(settings['key1'], 'value1');
      expect(settings['key3'], 'value3');
      expect(settings.containsKey('key2'), isFalse);
    });
    
    test('should handle typed settings', () async {
      await settingsDao.setBoolSetting('bool_key', true);
      await settingsDao.setIntSetting('int_key', 42);
      await settingsDao.setDoubleSetting('double_key', 3.14);
      
      final boolValue = await settingsDao.getBoolSetting('bool_key');
      final intValue = await settingsDao.getIntSetting('int_key');
      final doubleValue = await settingsDao.getDoubleSetting('double_key');
      
      expect(boolValue, true);
      expect(intValue, 42);
      expect(doubleValue, 3.14);
    });
    
    test('should watch setting changes', () async {
      final stream = settingsDao.watchSetting('watch_key');
      final changes = <String?>[];
      
      final subscription = stream.listen(changes.add);
      
      // Initially null
      await Future.delayed(const Duration(milliseconds: 100));
      expect(changes.last, isNull);
      
      // Set value
      await settingsDao.setSetting('watch_key', 'value1');
      await Future.delayed(const Duration(milliseconds: 100));
      expect(changes.last, 'value1');
      
      // Update value
      await settingsDao.setSetting('watch_key', 'value2');
      await Future.delayed(const Duration(milliseconds: 100));
      expect(changes.last, 'value2');
      
      await subscription.cancel();
    });
    
    test('should reset to defaults', () async {
      await settingsDao.setSettings({
        'custom1': 'value1',
        'custom2': 'value2',
        'theme_mode': 'dark',
        'notifications_enabled': 'false',
      });
      
      await settingsDao.resetToDefaults();
      
      final custom1 = await settingsDao.getSetting('custom1');
      final custom2 = await settingsDao.getSetting('custom2');
      final themeMode = await settingsDao.getSetting('theme_mode');
      final notifications = await settingsDao.getSetting('notifications_enabled');
      
      expect(custom1, isNull);
      expect(custom2, isNull);
      expect(themeMode, 'system');
      expect(notifications, 'true');
    });
  });
  
  group('SyncMetadata Tests', () {
    test('should track sync status', () async {
      await database.updateSyncStatus(
        tableName: 'user_profiles',
        recordId: 'user-1',
        status: SyncStatus.pending,
      );
      
      final pending = await database.getPendingSyncRecords();
      expect(pending.length, 1);
      expect(pending.first.tableName, 'user_profiles');
      expect(pending.first.recordId, 'user-1');
      expect(pending.first.syncStatus, SyncStatus.pending);
    });
    
    test('should update sync status with error', () async {
      await database.updateSyncStatus(
        tableName: 'user_profiles',
        recordId: 'user-2',
        status: SyncStatus.error,
        errorMessage: 'Network error',
      );
      
      final pending = await database.getPendingSyncRecords();
      expect(pending.length, 1);
      expect(pending.first.errorMessage, 'Network error');
    });
  });
}

// Extension to create test database
extension TestDatabase on AppDatabase {
  static AppDatabase forTesting(QueryExecutor executor) {
    return AppDatabase.forTesting(executor);
  }
}