class Receipt {
  const Receipt({
    required this.id,
    required this.amount,
    required this.currency,
    required this.merchantName,
    required this.date,
    this.category,
    this.items = const [],
  });

  final String id;
  final double amount;
  final String currency;
  final String merchantName;
  final DateTime date;
  final String? category;
  final List<ReceiptItem> items;

  factory Receipt.fromJson(Map<String, dynamic> json) {
    return Receipt(
      id: json['id'] as String,
      amount: (json['amount'] as num).toDouble(),
      currency: json['currency'] as String? ?? 'USD',
      merchantName: json['merchantName'] as String,
      date: DateTime.parse(json['date'] as String),
      category: json['category'] as String?,
      items: (json['items'] as List?)
          ?.map((item) => ReceiptItem.fromJson(item as Map<String, dynamic>))
          .toList() ?? [],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'amount': amount,
      'currency': currency,
      'merchantName': merchantName,
      'date': date.toIso8601String(),
      'category': category,
      'items': items.map((item) => item.toJson()).toList(),
    };
  }
}

class ReceiptItem {
  const ReceiptItem({
    required this.name,
    required this.quantity,
    required this.price,
  });

  final String name;
  final int quantity;
  final double price;

  factory ReceiptItem.fromJson(Map<String, dynamic> json) {
    return ReceiptItem(
      name: json['name'] as String,
      quantity: json['quantity'] as int,
      price: (json['price'] as num).toDouble(),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'name': name,
      'quantity': quantity,
      'price': price,
    };
  }
}