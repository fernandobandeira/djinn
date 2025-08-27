import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:djinn_mobile/core/providers/auth_providers.dart';
import 'package:djinn_mobile/core/database/app_database.dart';
import 'package:djinn_mobile/core/graphql/graphql_client.dart';
import 'package:ferry/ferry.dart';

/// Mock authentication state notifier
class MockAuthStateNotifier extends AuthStateNotifier {
  MockAuthStateNotifier() : super();
  
  void setAuthenticated(bool authenticated) {
    if (authenticated) {
      state = AuthState(
        isAuthenticated: true,
        isLoading: false,
        user: AuthUser(
          id: 'test-user-id',
          email: 'test@example.com',
          name: 'Test User',
        ),
      );
    } else {
      state = const AuthState(
        isAuthenticated: false,
        isLoading: false,
        user: null,
      );
    }
  }
  
  void setLoading(bool loading) {
    state = state.copyWith(isLoading: loading);
  }
  
  void setError(String error) {
    state = state.copyWith(error: error);
  }
}

/// Mock database for testing
class MockAppDatabase extends AppDatabase {
  final Map<String, UserProfile> _users = {};
  final Map<String, String> _settings = {};
  
  @override
  Future<UserProfile?> getUserProfile(String id) async {
    await Future.delayed(const Duration(milliseconds: 10));
    return _users[id];
  }
  
  @override
  Future<int> upsertUserProfile(UserProfile profile) async {
    await Future.delayed(const Duration(milliseconds: 10));
    _users[profile.id] = profile;
    return 1;
  }
  
  @override
  Future<String?> getSetting(String key) async {
    await Future.delayed(const Duration(milliseconds: 10));
    return _settings[key];
  }
  
  @override
  Future<void> setSetting(String key, String value) async {
    await Future.delayed(const Duration(milliseconds: 10));
    _settings[key] = value;
  }
  
  @override
  Stream<UserProfile?> watchUserProfile(String id) {
    return Stream.value(_users[id]);
  }
  
  @override
  Stream<String?> watchSetting(String key) {
    return Stream.value(_settings[key]);
  }
  
  void reset() {
    _users.clear();
    _settings.clear();
  }
}

/// Mock GraphQL client
class MockGraphQLClient implements Client {
  final Map<String, dynamic> _responses = {};
  final List<String> _requestHistory = [];
  
  void setResponse(String operationName, dynamic response) {
    _responses[operationName] = response;
  }
  
  List<String> get requestHistory => List.unmodifiable(_requestHistory);
  
  void reset() {
    _responses.clear();
    _requestHistory.clear();
  }
  
  @override
  Stream<OperationResponse> request(OperationRequest request) {
    _requestHistory.add(request.operation.operationName ?? 'unknown');
    
    final response = _responses[request.operation.operationName];
    if (response == null) {
      return Stream.value(
        OperationResponse(
          operationRequest: request,
          data: null,
          graphqlErrors: [],
          linkException: null,
        ),
      );
    }
    
    if (response is Exception) {
      return Stream.value(
        OperationResponse(
          operationRequest: request,
          data: null,
          graphqlErrors: [],
          linkException: response as LinkException?,
        ),
      );
    }
    
    return Stream.value(
      OperationResponse(
        operationRequest: request,
        data: response,
        graphqlErrors: [],
        linkException: null,
      ),
    );
  }
  
  @override
  Cache get cache => throw UnimplementedError();
  
  @override
  Map<OperationType, FetchPolicy> get defaultFetchPolicies => {
    OperationType.query: FetchPolicy.NetworkOnly,
    OperationType.mutation: FetchPolicy.NetworkOnly,
  };
  
  @override
  void dispose() {}
  
  @override
  void evict(OperationRequest request) {}
  
  @override
  Link get link => throw UnimplementedError();
  
  @override
  Future<OperationResponse<TData, TVars>> waitForResponse<TData, TVars>(
    OperationRequest<TData, TVars> request,
  ) {
    return request(request as OperationRequest).first;
  }
  
  @override
  OperationResponse<TData, TVars>? readQuery<TData, TVars>(
    OperationRequest<TData, TVars> request,
  ) {
    return null;
  }
  
  @override
  OperationResponse<TData, TVars>? readFragment<TData, TVars>(
    FragmentRequest<TData, TVars> request,
  ) {
    return null;
  }
  
  @override
  void writeQuery<TData, TVars>(
    OperationRequest<TData, TVars> request,
    TData data,
  ) {}
  
  @override
  void writeFragment<TData, TVars>(
    FragmentRequest<TData, TVars> request,
    TData data,
  ) {}
  
  @override
  OperationResponse? readCache(OperationRequest request) {
    return null;
  }
}

/// Test providers
final mockAuthStateProvider = StateNotifierProvider<MockAuthStateNotifier, AuthState>((ref) {
  return MockAuthStateNotifier();
});

final mockDatabaseProvider = Provider<MockAppDatabase>((ref) {
  return MockAppDatabase();
});

final mockGraphQLClientProvider = Provider<MockGraphQLClient>((ref) {
  return MockGraphQLClient();
});

/// Provider overrides for testing
List<Override> getMockProviderOverrides({
  bool isAuthenticated = false,
  MockAppDatabase? database,
  MockGraphQLClient? graphQLClient,
}) {
  final authNotifier = MockAuthStateNotifier();
  authNotifier.setAuthenticated(isAuthenticated);
  
  return [
    authStateProvider.overrideWith((ref) => authNotifier),
    if (database != null)
      appDatabaseProvider.overrideWithValue(database),
    if (graphQLClient != null)
      graphqlClientProvider.overrideWithValue(graphQLClient),
  ];
}