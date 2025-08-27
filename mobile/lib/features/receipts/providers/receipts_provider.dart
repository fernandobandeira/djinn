import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:ferry/ferry.dart';
import 'package:gql/ast.dart';
import 'package:gql_exec/gql_exec.dart';

import 'package:djinn_mobile/core/graphql/graphql_service.dart';
import 'package:djinn_mobile/core/graphql/queries/example_queries.dart';
import 'package:djinn_mobile/features/receipts/models/receipt_model.dart';

// Receipts state provider
final receiptsProvider = StateNotifierProvider<ReceiptsNotifier, ReceiptsState>((ref) {
  final service = ref.watch(graphqlServiceProvider);
  return ReceiptsNotifier(service);
});

class ReceiptsState {
  const ReceiptsState({
    this.receipts = const [],
    this.isLoading = false,
    this.error,
    this.hasMore = true,
    this.isLoadingMore = false,
  });

  final List<Receipt> receipts;
  final bool isLoading;
  final String? error;
  final bool hasMore;
  final bool isLoadingMore;

  ReceiptsState copyWith({
    List<Receipt>? receipts,
    bool? isLoading,
    String? error,
    bool? hasMore,
    bool? isLoadingMore,
  }) {
    return ReceiptsState(
      receipts: receipts ?? this.receipts,
      isLoading: isLoading ?? this.isLoading,
      error: error ?? this.error,
      hasMore: hasMore ?? this.hasMore,
      isLoadingMore: isLoadingMore ?? this.isLoadingMore,
    );
  }
}

class ReceiptsNotifier extends StateNotifier<ReceiptsState> {
  ReceiptsNotifier(this._service) : super(const ReceiptsState());

  final GraphQLService _service;
  static const _pageSize = 20;

  Future<void> loadReceipts({bool refresh = false}) async {
    if (refresh) {
      state = state.copyWith(receipts: [], hasMore: true);
    }

    state = state.copyWith(isLoading: true, error: null);

    final request = Request(
      operation: Operation(
        document: parseString(getReceiptsQuery),
      ),
      variables: {
        'limit': _pageSize,
        'offset': refresh ? 0 : state.receipts.length,
      },
    );

    try {
      await for (final response in _service.query(request)) {
        if (response.hasErrors) {
          state = state.copyWith(
            isLoading: false,
            error: response.firstError,
          );
        } else if (response.hasData) {
          final data = response.data as Map<String, dynamic>?;
          final receiptsData = data?['receipts'] as Map<String, dynamic>?;
          final edges = receiptsData?['edges'] as List?;
          
          final newReceipts = edges
              ?.map((edge) => Receipt.fromJson(edge['node'] as Map<String, dynamic>))
              .toList() ?? [];
          
          final pageInfo = receiptsData?['pageInfo'] as Map<String, dynamic>?;
          final hasMore = pageInfo?['hasNextPage'] as bool? ?? false;
          
          state = state.copyWith(
            receipts: refresh ? newReceipts : [...state.receipts, ...newReceipts],
            isLoading: false,
            hasMore: hasMore,
          );
        }
        
        if (!response.loading) break;
      }
    } catch (e) {
      state = state.copyWith(
        isLoading: false,
        error: e.toString(),
      );
    }
  }

  Future<void> loadMore() async {
    if (state.isLoadingMore || !state.hasMore) return;
    
    state = state.copyWith(isLoadingMore: true);
    await loadReceipts();
    state = state.copyWith(isLoadingMore: false);
  }

  void clearCache() {
    _service.clearCache();
    state = const ReceiptsState();
  }
}

// Example of a cached data provider
final cachedReceiptsProvider = Provider<List<Receipt>?>((ref) {
  final service = ref.watch(graphqlServiceProvider);
  
  final request = Request(
    operation: Operation(
      document: parseString(getReceiptsQuery),
    ),
    variables: const {
      'limit': 20,
      'offset': 0,
    },
  );
  
  // Read from cache only
  final cachedData = service.readFromCache(request);
  
  if (cachedData != null) {
    final receiptsData = cachedData['receipts'] as Map<String, dynamic>?;
    final edges = receiptsData?['edges'] as List?;
    
    return edges
        ?.map((edge) => Receipt.fromJson(edge['node'] as Map<String, dynamic>))
        .toList();
  }
  
  return null;
});