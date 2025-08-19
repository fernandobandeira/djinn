---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T15:20:34.937595'
profile: code_examples
source: https://ferrygraphql.com/docs/setup
topic: Ferry GraphQL Client Setup and Configuration
---

# Ferry GraphQL Client Setup and Configuration

Menu
[![Ferry Logo](https://ferrygraphql.com/img/logo.svg)**Ferry**](https://ferrygraphql.com/)[Docs](https://ferrygraphql.com/docs/)[Examples](https://github.com/gql-dart/ferry/tree/master/examples)
[pub.dev](https://pub.dev/packages/ferry)[GitHub](https://github.com/gql-dart/ferry)
🌜
🌞
SearchK
[![Ferry Logo](https://ferrygraphql.com/img/logo.svg)**Ferry**](https://ferrygraphql.com/)
  * [Docs](https://ferrygraphql.com/docs/)
  * [Examples](https://github.com/gql-dart/ferry/tree/master/examples)
  * [pub.dev](https://pub.dev/packages/ferry)
  * [GitHub](https://github.com/gql-dart/ferry)


Menu
  * [Ferry](https://ferrygraphql.com/docs/setup/#!)
    * [Introduction](https://ferrygraphql.com/docs/)
    * [How It Works](https://ferrygraphql.com/docs/how-it-works)
  * [Getting Started](https://ferrygraphql.com/docs/setup/#!)
    * [Setup](https://ferrygraphql.com/docs/setup)
    * [Generate Classes](https://ferrygraphql.com/docs/codegen)
  * [Fetching](https://ferrygraphql.com/docs/setup/#!)
    * [Queries](https://ferrygraphql.com/docs/queries)
    * [Mutations](https://ferrygraphql.com/docs/mutations)
    * [Fetch Policies](https://ferrygraphql.com/docs/fetch-policies)
    * [Pagination & Refetching](https://ferrygraphql.com/docs/pagination)
    * [Error Handling](https://ferrygraphql.com/docs/error-handling)
  * [Caching](https://ferrygraphql.com/docs/setup/#!)
    * [Configuration](https://ferrygraphql.com/docs/cache-configuration)
    * [Reading and Writing](https://ferrygraphql.com/docs/cache-interaction)
    * [Garbage Collection & Eviction](https://ferrygraphql.com/docs/garbage-collection)
  * [Flutter](https://ferrygraphql.com/docs/setup/#!)
    * [Using with Flutter](https://ferrygraphql.com/docs/flutter)
    * [Operation Widget](https://ferrygraphql.com/docs/flutter-operation-widget)
    * [Structuring Queries](https://ferrygraphql.com/docs/structuring-queries)
  * [Advanced](https://ferrygraphql.com/docs/setup/#!)
    * [Custom Scalars](https://ferrygraphql.com/docs/custom-scalars)
    * [Customizing the Client](https://ferrygraphql.com/docs/customization)
    * [Isolates](https://ferrygraphql.com/docs/isolates)
    * [JSON Operations](https://ferrygraphql.com/docs/json-operation)


# Setup
## Basic Setup[#](https://ferrygraphql.com/docs/setup/#basic-setup "Direct link to heading")
### Install Dependencies[#](https://ferrygraphql.com/docs/setup/#install-dependencies "Direct link to heading")
Add the following to your `pubspec.yaml`:
Copy
dependencies:
ferry:#[latest-version]
gql_http_link:#[latest-version]
# common serializers, which the code generator will assume are available
gql_code_builder_serializers:#[latest-version]
dev_dependencies:
ferry_generator:#[latest-version]
build_runner:#[latest-version]
### Initialize the Client[#](https://ferrygraphql.com/docs/setup/#initialize-the-client "Direct link to heading")
This instantiates a client with the default configuration, including a `Cache` instance that uses a `MemoryStore` to store data.
Copy
import'package:gql_http_link/gql_http_link.dart';
import'package:ferry/ferry.dart';
import'package:<your_pkg>/<path_to_your_schema>__generated__/schema.schema.gql.dart'show possibleTypesMap;
final link =HttpLink("[path/to/endpoint]");
final cache =Cache(possibleTypes: possibleTypesMap);
final client =Client(link: link, cache: cache);
## Setup With HiveStore[#](https://ferrygraphql.com/docs/setup/#setup-with-hivestore "Direct link to heading")
Ferry includes a `HiveStore` which enables offline data persistence, based on the `hive` package.
### Install Dependencies[#](https://ferrygraphql.com/docs/setup/#install-dependencies-1 "Direct link to heading")
To use the `HiveStore`, you'll need to add these dependencies to your `pubspec.yaml`:
Copy
dependencies:
ferry:#[latest-version]
gql_http_link:#[latest-version]
hive:#[latest-version]
ferry_hive_store:#[latest-version]
dev_dependencies:
ferry_generator:#[latest-version]
build_runner:#[latest-version]
Since `hive` seems to be unmaintained, and the community forked it via the `hive_ce` package, Ferry also includes a `ferry_hive_ce_store` with the same API.
Since the underlying file formats are binary compatible, switching from `ferry_hive_store ` to `ferry_hive_ce_store` is possible without any data migration.
##### important
If you're using Flutter, you'll also need to add:
Copy
dependencies:
hive_flutter:#[latest-version]
### Initialize the Client[#](https://ferrygraphql.com/docs/setup/#initialize-the-client-1 "Direct link to heading")
Copy
import'package:gql_http_link/gql_http_link.dart';
import'package:ferry/ferry.dart';
import'package:ferry_hive_store/ferry_hive_store.dart';
import'package:hive/hive.dart';
// *** If using flutter ***
// import 'package:hive_flutter/hive_flutter.dart';
Future<Client>initClient()async{
Hive.init('hive_data');
// OR, if using flutter
// await Hive.initFlutter();
final box =await Hive.openBox("graphql");
final store =HiveStore(box);
final cache =Cache(store: store, possibleTypes: possibleTypesMap);
final link =HttpLink('[path/to/endpoint]');
final client =Client(
link: link,
cache: cache,
);
return client;
}
[Edit this page](https://github.com/gql-dart/ferry/edit/master/docs/../docs/setup.md)
[Previous« How It Works](https://ferrygraphql.com/docs/how-it-works)
[NextGenerate GraphQL Classes »](https://ferrygraphql.com/docs/codegen)
  * [Basic Setup](https://ferrygraphql.com/docs/setup/#basic-setup)
    * [Install Dependencies](https://ferrygraphql.com/docs/setup/#install-dependencies)
    * [Initialize the Client](https://ferrygraphql.com/docs/setup/#initialize-the-client)
  * [Setup With HiveStore](https://ferrygraphql.com/docs/setup/#setup-with-hivestore)
    * [Install Dependencies](https://ferrygraphql.com/docs/setup/#install-dependencies-1)
    * [Initialize the Client](https://ferrygraphql.com/docs/setup/#initialize-the-client-1)


#### Docs
  * [Introduction](https://ferrygraphql.com/docs/)


#### Community
  * [Discord](https://discord.gg/G3JGkBY)


#### More
  * [pub.dev](https://pub.dev/packages/ferry)
  * [GitHub](https://github.com/gql-dart/ferry)


Copyright © 2025 Sat Mandir Khalsa. Built with Docusaurus.


## Source Information
- URL: https://ferrygraphql.com/docs/setup
- Harvested: 2025-08-19T15:20:34.937595
- Profile: code_examples
- Agent: architect
