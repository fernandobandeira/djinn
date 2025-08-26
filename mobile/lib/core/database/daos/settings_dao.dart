import 'package:drift/drift.dart';
import 'package:djinn_mobile/core/database/app_database.dart';

part 'settings_dao.g.dart';

@DriftAccessor(tables: [AppSettings])
class SettingsDao extends DatabaseAccessor<AppDatabase> with _$SettingsDaoMixin {
  SettingsDao(super.db);
  
  // Get a single setting
  Future<String?> getSetting(String key) async {
    final setting = await (select(appSettings)
      ..where((s) => s.key.equals(key)))
      .getSingleOrNull();
    return setting?.value;
  }
  
  // Get multiple settings
  Future<Map<String, String>> getSettings(List<String> keys) async {
    final settings = await (select(appSettings)
      ..where((s) => s.key.isIn(keys)))
      .get();
    
    return {
      for (final setting in settings)
        setting.key: setting.value,
    };
  }
  
  // Get all settings
  Future<Map<String, String>> getAllSettings() async {
    final settings = await select(appSettings).get();
    
    return {
      for (final setting in settings)
        setting.key: setting.value,
    };
  }
  
  // Set a single setting
  Future<void> setSetting(String key, String value) {
    return into(appSettings).insertOnConflictUpdate(
      AppSettingsCompanion(
        key: Value(key),
        value: Value(value),
        updatedAt: Value(DateTime.now()),
      ),
    );
  }
  
  // Set multiple settings
  Future<void> setSettings(Map<String, String> settings) async {
    await batch((batch) {
      for (final entry in settings.entries) {
        batch.insert(
          appSettings,
          AppSettingsCompanion(
            key: Value(entry.key),
            value: Value(entry.value),
            updatedAt: Value(DateTime.now()),
          ),
          mode: InsertMode.insertOrReplace,
        );
      }
    });
  }
  
  // Delete a setting
  Future<int> deleteSetting(String key) {
    return (delete(appSettings)..where((s) => s.key.equals(key))).go();
  }
  
  // Delete multiple settings
  Future<int> deleteSettings(List<String> keys) {
    return (delete(appSettings)..where((s) => s.key.isIn(keys))).go();
  }
  
  // Watch a single setting
  Stream<String?> watchSetting(String key) {
    return (select(appSettings)..where((s) => s.key.equals(key)))
      .watchSingleOrNull()
      .map((setting) => setting?.value);
  }
  
  // Watch multiple settings
  Stream<Map<String, String>> watchSettings(List<String> keys) {
    return (select(appSettings)..where((s) => s.key.isIn(keys)))
      .watch()
      .map((settings) => {
        for (final setting in settings)
          setting.key: setting.value,
      });
  }
  
  // Watch all settings
  Stream<Map<String, String>> watchAllSettings() {
    return select(appSettings)
      .watch()
      .map((settings) => {
        for (final setting in settings)
          setting.key: setting.value,
      });
  }
  
  // Typed getters for common settings
  Future<bool> getBoolSetting(String key, {bool defaultValue = false}) async {
    final value = await getSetting(key);
    if (value == null) return defaultValue;
    return value.toLowerCase() == 'true';
  }
  
  Future<int> getIntSetting(String key, {int defaultValue = 0}) async {
    final value = await getSetting(key);
    if (value == null) return defaultValue;
    return int.tryParse(value) ?? defaultValue;
  }
  
  Future<double> getDoubleSetting(String key, {double defaultValue = 0.0}) async {
    final value = await getSetting(key);
    if (value == null) return defaultValue;
    return double.tryParse(value) ?? defaultValue;
  }
  
  // Typed setters for common settings
  Future<void> setBoolSetting(String key, bool value) {
    return setSetting(key, value.toString());
  }
  
  Future<void> setIntSetting(String key, int value) {
    return setSetting(key, value.toString());
  }
  
  Future<void> setDoubleSetting(String key, double value) {
    return setSetting(key, value.toString());
  }
  
  // Theme settings
  Future<String> getThemeMode() async {
    return await getSetting('theme_mode') ?? 'system';
  }
  
  Future<void> setThemeMode(String mode) {
    return setSetting('theme_mode', mode);
  }
  
  Stream<String> watchThemeMode() {
    return watchSetting('theme_mode').map((mode) => mode ?? 'system');
  }
  
  // Notification settings
  Future<bool> areNotificationsEnabled() async {
    return await getBoolSetting('notifications_enabled', defaultValue: true);
  }
  
  Future<void> setNotificationsEnabled(bool enabled) {
    return setBoolSetting('notifications_enabled', enabled);
  }
  
  Stream<bool> watchNotificationsEnabled() {
    return watchSetting('notifications_enabled')
      .map((value) => value?.toLowerCase() == 'true');
  }
  
  // Clear all settings (except defaults)
  Future<void> resetToDefaults() async {
    await delete(appSettings).go();
    
    // Re-insert defaults
    await setSettings({
      'theme_mode': 'system',
      'notifications_enabled': 'true',
    });
  }
}