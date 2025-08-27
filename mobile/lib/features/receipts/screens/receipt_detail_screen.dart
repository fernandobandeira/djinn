import 'package:flutter/material.dart';

class ReceiptDetailScreen extends StatelessWidget {
  const ReceiptDetailScreen({
    super.key,
    required this.receiptId,
  });

  final String receiptId;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Receipt #$receiptId'),
      ),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(
              'Receipt Details',
              style: Theme.of(context).textTheme.headlineSmall,
            ),
            const SizedBox(height: 16),
            ListTile(
              title: const Text('Receipt ID'),
              subtitle: Text(receiptId),
            ),
            const ListTile(
              title: Text('Date'),
              subtitle: Text('2025-01-25'),
            ),
            const ListTile(
              title: Text('Amount'),
              subtitle: Text('\$42.00'),
            ),
            const ListTile(
              title: Text('Store'),
              subtitle: Text('Sample Store'),
            ),
          ],
        ),
      ),
    );
  }
}