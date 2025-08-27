import 'package:drift/drift.dart';
import 'package:djinn_mobile/core/database/app_database.dart';

part 'user_dao.g.dart';

@DriftAccessor(tables: [UserProfiles])
class UserDao extends DatabaseAccessor<AppDatabase> with _$UserDaoMixin {
  UserDao(super.db);
  
  // Basic CRUD operations
  Future<UserProfile?> getById(String id) {
    return (select(userProfiles)..where((u) => u.id.equals(id))).getSingleOrNull();
  }
  
  Future<UserProfile?> getByEmail(String email) {
    return (select(userProfiles)..where((u) => u.email.equals(email))).getSingleOrNull();
  }
  
  Future<List<UserProfile>> getAll() {
    return select(userProfiles).get();
  }
  
  Stream<List<UserProfile>> watchAll() {
    return select(userProfiles).watch();
  }
  
  Stream<UserProfile?> watchById(String id) {
    return (select(userProfiles)..where((u) => u.id.equals(id))).watchSingleOrNull();
  }
  
  Future<int> insertUser(UserProfilesCompanion user) {
    return into(userProfiles).insert(user);
  }
  
  Future<int> upsertUser(UserProfile user) {
    return into(userProfiles).insertOnConflictUpdate(user);
  }
  
  Future<bool> updateUser(UserProfile user) {
    return update(userProfiles).replace(user);
  }
  
  Future<int> deleteUser(String id) {
    return (delete(userProfiles)..where((u) => u.id.equals(id))).go();
  }
  
  Future<int> deleteAll() {
    return delete(userProfiles).go();
  }
  
  // Complex queries
  Future<UserProfile?> getVerifiedUserByEmail(String email) {
    return (select(userProfiles)
      ..where((u) => u.email.equals(email) & u.emailVerifiedAt.isNotNull()))
      .getSingleOrNull();
  }
  
  Future<List<UserProfile>> searchUsers(String query) {
    return (select(userProfiles)
      ..where((u) => 
        u.name.contains(query) | 
        u.email.contains(query)))
      .get();
  }
  
  Future<int> countUsers() async {
    final count = await customSelect(
      'SELECT COUNT(*) as count FROM user_profiles',
      readsFrom: {userProfiles},
    ).getSingle();
    
    return count.read<int>('count');
  }
  
  Future<List<UserProfile>> getRecentlyCreated({int limit = 10}) {
    return (select(userProfiles)
      ..orderBy([(u) => OrderingTerm.desc(u.createdAt)])
      ..limit(limit))
      .get();
  }
  
  // Batch operations
  Future<void> insertMultiple(List<UserProfilesCompanion> users) async {
    await batch((batch) {
      batch.insertAll(userProfiles, users);
    });
  }
  
  Future<void> deleteMultiple(List<String> ids) async {
    await batch((batch) {
      for (final id in ids) {
        batch.deleteWhere(userProfiles, (u) => u.id.equals(id));
      }
    });
  }
  
  // Migration helpers
  Future<void> updateEmailVerificationStatus(String userId, DateTime? verifiedAt) {
    return (update(userProfiles)..where((u) => u.id.equals(userId)))
      .write(UserProfilesCompanion(
        emailVerifiedAt: Value(verifiedAt),
        updatedAt: Value(DateTime.now()),
      ));
  }
  
  Future<void> updateProfileMetadata(String userId, Map<String, dynamic> metadata) async {
    // Convert metadata to JSON string
    final jsonString = metadata.toString(); // In production, use proper JSON encoding
    
    await (update(userProfiles)..where((u) => u.id.equals(userId)))
      .write(UserProfilesCompanion(
        metadata: Value(jsonString),
        updatedAt: Value(DateTime.now()),
      ));
  }
}