import 'package:ferry/ferry.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import 'package:djinn_mobile/core/graphql/graphql_client.dart';

// Base GraphQL service provider
final graphqlServiceProvider = Provider<GraphQLService>((ref) {
  final client = ref.watch(graphqlClientProvider);
  return GraphQLService(client);
});

class GraphQLService {
  const GraphQLService(this._client);

  final Client _client;

  // Get the Ferry client for direct use
  Client get client => _client;

  // Execute a request
  Stream<OperationResponse<TData, TVars>> request<TData, TVars>(
    OperationRequest<TData, TVars> request,
  ) {
    return _client.request(request);
  }

  // Read from cache only
  TData? readFromCache<TData, TVars>(
    OperationRequest<TData, TVars> request,
  ) {
    return _client.cache.readQuery(request);
  }

  // Write to cache
  void writeToCache<TData, TVars>(
    OperationRequest<TData, TVars> request,
    TData data,
  ) {
    _client.cache.writeQuery(request, data);
  }

  // Clear cache
  void clearCache() {
    _client.cache.clear();
  }

  // Evict specific item from cache
  void evictFromCache(String key) {
    _client.cache.evict(key);
  }
}

// Response extension for easier error handling
extension OperationResponseExtension<TData, TVars> on OperationResponse<TData, TVars> {
  bool get isLoading => loading && dataSource != DataSource.Optimistic;
  bool get hasErrors => graphqlErrors?.isNotEmpty ?? false;
  bool get isSuccess => hasData && !hasErrors;
  bool get hasData => data != null;
  
  String? get firstError => graphqlErrors?.firstOrNull?.message;
  
  T? mapData<T>(T Function(TData data) mapper) {
    return hasData ? mapper(data as TData) : null;
  }
}

// Fetch policies for different use cases
class GraphQLFetchPolicies {
  static const cacheFirst = FetchPolicy.CacheFirst;
  static const cacheAndNetwork = FetchPolicy.CacheAndNetwork;
  static const networkOnly = FetchPolicy.NetworkOnly;
  static const cacheOnly = FetchPolicy.CacheOnly;
  static const noCache = FetchPolicy.NoCache;
}