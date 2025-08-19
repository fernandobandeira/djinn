---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T15:20:11.243336'
profile: deep_research
source: https://ferrygraphql.com/docs
topic: Ferry GraphQL Flutter Implementation Patterns
---

# Ferry GraphQL Flutter Implementation Patterns

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
  * [Ferry](https://ferrygraphql.com/docs/#!)
    * [Introduction](https://ferrygraphql.com/docs/)
    * [How It Works](https://ferrygraphql.com/docs/how-it-works)
  * [Getting Started](https://ferrygraphql.com/docs/#!)
    * [Setup](https://ferrygraphql.com/docs/setup)
    * [Generate Classes](https://ferrygraphql.com/docs/codegen)
  * [Fetching](https://ferrygraphql.com/docs/#!)
    * [Queries](https://ferrygraphql.com/docs/queries)
    * [Mutations](https://ferrygraphql.com/docs/mutations)
    * [Fetch Policies](https://ferrygraphql.com/docs/fetch-policies)
    * [Pagination & Refetching](https://ferrygraphql.com/docs/pagination)
    * [Error Handling](https://ferrygraphql.com/docs/error-handling)
  * [Caching](https://ferrygraphql.com/docs/#!)
    * [Configuration](https://ferrygraphql.com/docs/cache-configuration)
    * [Reading and Writing](https://ferrygraphql.com/docs/cache-interaction)
    * [Garbage Collection & Eviction](https://ferrygraphql.com/docs/garbage-collection)
  * [Flutter](https://ferrygraphql.com/docs/#!)
    * [Using with Flutter](https://ferrygraphql.com/docs/flutter)
    * [Operation Widget](https://ferrygraphql.com/docs/flutter-operation-widget)
    * [Structuring Queries](https://ferrygraphql.com/docs/structuring-queries)
  * [Advanced](https://ferrygraphql.com/docs/#!)
    * [Custom Scalars](https://ferrygraphql.com/docs/custom-scalars)
    * [Customizing the Client](https://ferrygraphql.com/docs/customization)
    * [Isolates](https://ferrygraphql.com/docs/isolates)
    * [JSON Operations](https://ferrygraphql.com/docs/json-operation)


# Introduction
Ferry is a highly productive, full-featured, extensible GraphQL Client for Flutter & Dart. [**Get started**](https://ferrygraphql.com/docs/setup) with ferry now.
## 🚀 Productive[#](https://ferrygraphql.com/docs/#-productive "Direct link to heading")
Ferry is **fully typed** , allowing you to work faster and safer with compile time checks and IDE autocomplete, including fully typed Cache [reads and writes](https://ferrygraphql.com/docs/cache-interaction).
Ferry's built-in [code generator](https://ferrygraphql.com/docs/codegen) creates immutable data classes for all your GraphQL Operations and Fragments based on your GraphQL schema.
Use the client with any Link from the [gql_link](https://pub.dev/packages/gql_link) ecosystem.
## 💯 Full-Featured[#](https://ferrygraphql.com/docs/#-full-featured "Direct link to heading")
Ferry includes all the features you'd expect in a GraphQL client, including:
  1. **Normalized Optimistic Cache** : keep data in sync with cache normalization and update your UI instantly with optimistic data.
  2. **Multiple Data Stores** : extensible Store interface for data persistence, including built-in MemoryStore and [HiveStore](https://ferrygraphql.com/docs/setup#setup-with-hivestore) (for offline persistence).
  3. **Cache Eviction & Garbage Collection**: selectively [remove data](https://ferrygraphql.com/docs/garbage-collection#cacheevict) from the cache or use the built-in [garbage collection](https://ferrygraphql.com/docs/garbage-collection#cachegc) to remove stale data automatically
  4. **Refetch & Pagination**: easily [update responses](https://ferrygraphql.com/docs/pagination#refetching) with new data or combine multiple responses, allowing for seamless [pagination](https://ferrygraphql.com/docs/pagination#pagination).
  5. **Flutter Widget** : use the included [Operation](https://ferrygraphql.com/docs/flutter) Widget in your Flutter app.


## 💪 Extensible[#](https://ferrygraphql.com/docs/#-extensible "Direct link to heading")
Ferry's core features are implemented as composable `TypedLink`s, a typed version of the [gql_link](https://pub.dev/packages/gql_link) API. In fact, the Ferry Client itself is a `TypedLink`, created by composing several other [core links](https://ferrygraphql.com/docs/customization#core-typedlinks). This architecture makes extending or [customizing the client](https://ferrygraphql.com/docs/customization) extremely easy.
[Edit this page](https://github.com/gql-dart/ferry/edit/master/docs/../docs/intro.md)
[NextHow It Works »](https://ferrygraphql.com/docs/how-it-works)
  * [🚀 Productive](https://ferrygraphql.com/docs/#-productive)
  * [💯 Full-Featured](https://ferrygraphql.com/docs/#-full-featured)
  * [💪 Extensible](https://ferrygraphql.com/docs/#-extensible)


#### Docs
  * [Introduction](https://ferrygraphql.com/docs/)


#### Community
  * [Discord](https://discord.gg/G3JGkBY)


#### More
  * [pub.dev](https://pub.dev/packages/ferry)
  * [GitHub](https://github.com/gql-dart/ferry)


Copyright © 2025 Sat Mandir Khalsa. Built with Docusaurus.


## Source Information
- URL: https://ferrygraphql.com/docs
- Harvested: 2025-08-19T15:20:11.243336
- Profile: deep_research
- Agent: architect
