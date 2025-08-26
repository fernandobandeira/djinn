import 'dart:io';

import 'package:drift/drift.dart';
import 'package:drift_flutter/drift_flutter.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:path_provider/path_provider.dart';
import 'package:path/path.dart' as p;

part 'app_database.g.dart';

// User profiles table
@DataClassName('UserProfile')
class UserProfiles extends Table {
  TextColumn get id => text()();
  TextColumn get email => text()();
  TextColumn get name => text().nullable()();
  TextColumn get avatarUrl => text().nullable()();
  TextColumn get phoneNumber => text().nullable()();
  DateTimeColumn get emailVerifiedAt => dateTime().nullable()();
  DateTimeColumn get createdAt => dateTime().withDefault(currentDateAndTime)();
  DateTimeColumn get updatedAt => dateTime().withDefault(currentDateAndTime)();
  TextColumn get metadata => text().nullable()(); // JSON string for additional data
  
  @override
  Set<Column> get primaryKey => {id};
}

// Settings table for app preferences
@DataClassName('AppSetting')
class AppSettings extends Table {
  TextColumn get key => text()();
  TextColumn get value => text()();
  DateTimeColumn get updatedAt => dateTime().withDefault(currentDateAndTime)();
  
  @override
  Set<Column> get primaryKey => {key};
}

// Sync metadata table for tracking sync status
@DataClassName('SyncMetadata')
class SyncMetadataTable extends Table {
  TextColumn get tableName => text()();
  TextColumn get recordId => text()();
  DateTimeColumn get lastSyncedAt => dateTime().nullable()();
  TextColumn get syncStatus => textEnum<SyncStatus>()();
  TextColumn get errorMessage => text().nullable()();
  IntColumn get retryCount => integer().withDefault(const Constant(0))();
  
  @override
  Set<Column> get primaryKey => {tableName, recordId};
  
  @override
  String get tableName => 'sync_metadata';
}

enum SyncStatus {
  pending,
  syncing,
  synced,
  error,
  conflict
}

// Database class
@DriftDatabase(tables: [UserProfiles, AppSettings, SyncMetadataTable])
class AppDatabase extends _$AppDatabase {
  AppDatabase() : super(_openConnection());
  
  @override
  int get schemaVersion => 1;
  
  @override
  MigrationStrategy get migration {
    return MigrationStrategy(
      onCreate: (Migrator m) async {
        await m.createAll();
        
        // Insert default settings
        await into(appSettings).insert(
          AppSettingsCompanion(
            key: const Value('theme_mode'),
            value: const Value('system'),
          ),
          mode: InsertMode.insertOrReplace,
        );
        
        await into(appSettings).insert(
          AppSettingsCompanion(
            key: const Value('notifications_enabled'),
            value: const Value('true'),
          ),
          mode: InsertMode.insertOrReplace,
        );
      },
      onUpgrade: (Migrator m, int from, int to) async {
        // Handle future migrations here
        // Example:
        // if (from < 2) {
        //   await m.addColumn(userProfiles, userProfiles.newColumn);
        // }
      },
      beforeOpen: (details) async {
        // Enable foreign keys
        await customStatement('PRAGMA foreign_keys = ON');
      },
    );
  }
  
  // User profile operations
  Future<UserProfile?> getUserProfile(String id) {
    return (select(userProfiles)..where((u) => u.id.equals(id))).getSingleOrNull();
  }
  
  Future<UserProfile?> getUserProfileByEmail(String email) {
    return (select(userProfiles)..where((u) => u.email.equals(email))).getSingleOrNull();
  }
  
  Future<int> upsertUserProfile(UserProfile profile) {
    return into(userProfiles).insertOnConflictUpdate(profile);
  }
  
  Future<int> deleteUserProfile(String id) {
    return (delete(userProfiles)..where((u) => u.id.equals(id))).go();
  }
  
  Stream<UserProfile?> watchUserProfile(String id) {
    return (select(userProfiles)..where((u) => u.id.equals(id))).watchSingleOrNull();
  }
  
  // Settings operations
  Future<String?> getSetting(String key) async {
    final setting = await (select(appSettings)..where((s) => s.key.equals(key))).getSingleOrNull();
    return setting?.value;
  }
  
  Future<void> setSetting(String key, String value) {
    return into(appSettings).insertOnConflictUpdate(
      AppSettingsCompanion(
        key: Value(key),
        value: Value(value),
        updatedAt: Value(DateTime.now()),
      ),
    );
  }
  
  Stream<String?> watchSetting(String key) {
    return (select(appSettings)..where((s) => s.key.equals(key)))
        .watchSingleOrNull()
        .map((setting) => setting?.value);
  }
  
  // Sync metadata operations
  Future<void> updateSyncStatus({
    required String tableName,
    required String recordId,
    required SyncStatus status,
    DateTime? lastSyncedAt,
    String? errorMessage,
  }) {
    return into(syncMetadataTable).insertOnConflictUpdate(
      SyncMetadata(
        tableName: tableName,
        recordId: recordId,
        syncStatus: status,
        lastSyncedAt: lastSyncedAt,
        errorMessage: errorMessage,
        retryCount: 0,
      ),
    );
  }
  
  Future<List<SyncMetadata>> getPendingSyncRecords() {
    return (select(syncMetadataTable)
      ..where((s) => s.syncStatus.equalsValue(SyncStatus.pending) | 
                     s.syncStatus.equalsValue(SyncStatus.error))
      ..orderBy([(s) => OrderingTerm.asc(s.retryCount)]))
      .get();
  }
  
  // Clear all data (for logout)
  Future<void> clearAllData() async {
    await delete(userProfiles).go();
    await delete(syncMetadataTable).go();
    // Keep app settings
  }
  
  // Database maintenance
  Future<void> vacuum() async {
    await customStatement('VACUUM');
  }
  
  Future<int> getDatabaseSize() async {
    final result = await customSelect('SELECT page_count * page_size as size FROM pragma_page_count(), pragma_page_size()').getSingle();
    return result.read<int>('size');
  }
}

// Connection setup
LazyDatabase _openConnection() {
  return driftDatabase(name: 'djinn_app_db');
}

// Provider for database
final appDatabaseProvider = Provider<AppDatabase>((ref) {
  final database = AppDatabase();
  
  // Cleanup on dispose
  ref.onDispose(() {
    database.close();
  });
  
  return database;
});

// Note: Current user profile provider has been moved to database_provider.dart
// to properly integrate with auth state management