import 'package:gql/ast.dart';

// Example user query
const String getUserQuery = r'''
  query GetUser($id: ID!) {
    user(id: $id) {
      id
      email
      name
      createdAt
      updatedAt
    }
  }
''';

// Example receipts query
const String getReceiptsQuery = r'''
  query GetReceipts($limit: Int, $offset: Int) {
    receipts(limit: $limit, offset: $offset) {
      edges {
        node {
          id
          amount
          currency
          merchantName
          date
          category
          items {
            name
            quantity
            price
          }
        }
      }
      pageInfo {
        hasNextPage
        hasPreviousPage
        totalCount
      }
    }
  }
''';

// Example mutation
const String createReceiptMutation = r'''
  mutation CreateReceipt($input: CreateReceiptInput!) {
    createReceipt(input: $input) {
      receipt {
        id
        amount
        merchantName
        date
      }
      errors {
        field
        message
      }
    }
  }
''';

// Example subscription
const String receiptUpdatesSubscription = r'''
  subscription OnReceiptUpdate($userId: ID!) {
    receiptUpdate(userId: $userId) {
      id
      amount
      status
      updatedAt
    }
  }
''';