import 'dart:io';

import 'package:ferry/ferry.dart';
import 'package:ferry_hive_store/ferry_hive_store.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:gql_link/gql_link.dart';
import 'package:gql_http_link/gql_http_link.dart';
import 'package:gql_error_link/gql_error_link.dart';
import 'package:hive_flutter/hive_flutter.dart';

import 'package:djinn_mobile/core/config/environment.dart';
import 'package:djinn_mobile/core/providers/auth_providers.dart';

// GraphQL client provider
final graphqlClientProvider = Provider<Client>((ref) {
  return GraphQLClientFactory.create(ref);
});

class GraphQLClientFactory {
  static Client create(ProviderRef ref) {
    final authToken = ref.watch(authTokenProvider);
    
    // Create HTTP link with auth headers
    final httpLink = HttpLink(
      AppConfig.apiUrl,
      defaultHeaders: {
        'Content-Type': 'application/json',
        'X-Client-Version': '1.0.0',
        'X-Platform': Platform.operatingSystem,
        if (authToken != null) 'Authorization': 'Bearer $authToken',
      },
    );

    // Error handling link
    final errorLink = ErrorLink(
      onGraphQLError: (request, forward, response) {
        if (response.errors != null) {
          for (final error in response.errors!) {
            print('[GraphQL Error] ${error.message}');
            if (error.extensions?['code'] == 'UNAUTHENTICATED') {
              // Handle unauthenticated error
              ref.read(authStateProvider.notifier).signOut();
            }
          }
        }
        return null;
      },
      onException: (request, forward, exception) {
        print('[GraphQL Exception] $exception');
        return null;
      },
    );

    // Combine links
    final link = Link.from([errorLink, httpLink]);

    // Create cache with Hive store
    final cache = Cache(
      store: HiveStore(Hive.box('graphql_cache')),
      possibleTypes: {}, // Add your possible types here
    );

    // Create the client
    return Client(
      link: link,
      cache: cache,
      defaultFetchPolicies: {
        OperationType.query: FetchPolicy.CacheFirst,
        OperationType.mutation: FetchPolicy.NetworkOnly,
        OperationType.subscription: FetchPolicy.NetworkOnly,
      },
    );
  }

  // Initialize Hive for GraphQL caching
  static Future<void> initializeHive() async {
    await Hive.initFlutter();
    await Hive.openBox<Map>('graphql_cache');
  }
}

// Extension for easier client access
extension GraphQLClientExtension on WidgetRef {
  Client get graphql => read(graphqlClientProvider);
}