import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class ReceiptsScreen extends StatelessWidget {
  const ReceiptsScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Receipts'),
      ),
      body: ListView.builder(
        itemCount: 5,
        itemBuilder: (context, index) {
          return ListTile(
            leading: const Icon(Icons.receipt),
            title: Text('Receipt #${index + 1}'),
            subtitle: Text('Date: ${DateTime.now().toString().split(' ')[0]}'),
            trailing: const Icon(Icons.arrow_forward_ios),
            onTap: () => context.go('/home/receipts/${index + 1}'),
          );
        },
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          ScaffoldMessenger.of(context).showSnackBar(
            const SnackBar(content: Text('Scan receipt coming soon!')),
          );
        },
        child: const Icon(Icons.camera_alt),
      ),
    );
  }
}