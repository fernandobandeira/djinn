import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:ferry/ferry.dart';
import 'package:gql_exec/gql_exec.dart';
import 'package:gql/language.dart';
import 'package:hive_flutter/hive_flutter.dart';

import 'package:djinn_mobile/core/graphql/graphql_client.dart';
import 'package:djinn_mobile/core/graphql/graphql_service.dart';
import 'package:djinn_mobile/core/config/environment.dart';

void main() {
  setUpAll(() async {
    TestWidgetsFlutterBinding.ensureInitialized();
    AppConfig.initialize();
    // Initialize Hive for testing
    await Hive.initFlutter();
    if (!Hive.isBoxOpen('graphql_cache')) {
      await Hive.openBox<Map>('graphql_cache');
    }
  });

  tearDownAll(() async {
    await Hive.close();
  });

  group('GraphQL Service Tests', () {
    late ProviderContainer container;
    late GraphQLService service;

    setUp(() {
      container = ProviderContainer();
      service = container.read(graphqlServiceProvider);
    });

    tearDown(() {
      container.dispose();
    });

    test('graphqlServiceProvider should return GraphQLService instance', () {
      expect(service, isA<GraphQLService>());
    });

    test('cache operations should work', () {
      const testQuery = r'''
        query TestQuery {
          test {
            id
            value
          }
        }
      ''';

      final request = Request(
        operation: Operation(
          document: parseString(testQuery),
        ),
      );

      const testData = {
        'test': {
          'id': '1',
          'value': 'test value',
        }
      };

      // Write to cache
      service.writeToCache(request, testData);

      // Read from cache
      final cachedData = service.readFromCache(request);
      expect(cachedData, equals(testData));

      // Clear cache
      service.clearCache();
      final clearedData = service.readFromCache(request);
      expect(clearedData, isNull);
    });

    test('fetch policies should be available', () {
      expect(GraphQLFetchPolicies.cacheFirst, equals(FetchPolicy.CacheFirst));
      expect(GraphQLFetchPolicies.cacheAndNetwork, equals(FetchPolicy.CacheAndNetwork));
      expect(GraphQLFetchPolicies.networkOnly, equals(FetchPolicy.NetworkOnly));
      expect(GraphQLFetchPolicies.cacheOnly, equals(FetchPolicy.CacheOnly));
      expect(GraphQLFetchPolicies.noCache, equals(FetchPolicy.NoCache));
    });
  });

  group('OperationResponse Extension Tests', () {
    test('response extension methods should work correctly', () {
      // Success response
      final successResponse = OperationResponse<Map<String, dynamic>, dynamic>(
        operationRequest: Request(
          operation: Operation(document: parseString('query { test }')),
        ),
        data: {'test': 'value'},
        dataSource: DataSource.Link,
      );

      expect(successResponse.isLoading, isFalse);
      expect(successResponse.hasErrors, isFalse);
      expect(successResponse.isSuccess, isTrue);
      expect(successResponse.hasData, isTrue);
      expect(successResponse.firstError, isNull);
      
      final mappedData = successResponse.mapData((data) => data['test']);
      expect(mappedData, equals('value'));

      // Error response
      final errorResponse = OperationResponse<Map<String, dynamic>, dynamic>(
        operationRequest: Request(
          operation: Operation(document: parseString('query { test }')),
        ),
        graphqlErrors: [
          GraphQLError(message: 'Test error'),
        ],
        dataSource: DataSource.Link,
      );

      expect(errorResponse.hasErrors, isTrue);
      expect(errorResponse.isSuccess, isFalse);
      expect(errorResponse.firstError, equals('Test error'));
    });
  });
}