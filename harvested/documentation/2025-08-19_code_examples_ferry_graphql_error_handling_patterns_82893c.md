---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T15:21:01.131451'
profile: code_examples
source: https://ferrygraphql.com/docs/error-handling
topic: Ferry GraphQL Error Handling Patterns
---

# Ferry GraphQL Error Handling Patterns

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
  * [Ferry](https://ferrygraphql.com/docs/error-handling/#!)
    * [Introduction](https://ferrygraphql.com/docs/)
    * [How It Works](https://ferrygraphql.com/docs/how-it-works)
  * [Getting Started](https://ferrygraphql.com/docs/error-handling/#!)
    * [Setup](https://ferrygraphql.com/docs/setup)
    * [Generate Classes](https://ferrygraphql.com/docs/codegen)
  * [Fetching](https://ferrygraphql.com/docs/error-handling/#!)
    * [Queries](https://ferrygraphql.com/docs/queries)
    * [Mutations](https://ferrygraphql.com/docs/mutations)
    * [Fetch Policies](https://ferrygraphql.com/docs/fetch-policies)
    * [Pagination & Refetching](https://ferrygraphql.com/docs/pagination)
    * [Error Handling](https://ferrygraphql.com/docs/error-handling)
  * [Caching](https://ferrygraphql.com/docs/error-handling/#!)
    * [Configuration](https://ferrygraphql.com/docs/cache-configuration)
    * [Reading and Writing](https://ferrygraphql.com/docs/cache-interaction)
    * [Garbage Collection & Eviction](https://ferrygraphql.com/docs/garbage-collection)
  * [Flutter](https://ferrygraphql.com/docs/error-handling/#!)
    * [Using with Flutter](https://ferrygraphql.com/docs/flutter)
    * [Operation Widget](https://ferrygraphql.com/docs/flutter-operation-widget)
    * [Structuring Queries](https://ferrygraphql.com/docs/structuring-queries)
  * [Advanced](https://ferrygraphql.com/docs/error-handling/#!)
    * [Custom Scalars](https://ferrygraphql.com/docs/custom-scalars)
    * [Customizing the Client](https://ferrygraphql.com/docs/customization)
    * [Isolates](https://ferrygraphql.com/docs/isolates)
    * [JSON Operations](https://ferrygraphql.com/docs/json-operation)


# Error Handling
When creating or running an application, sooner or later, you're bound to encounter some type of error. Ferry is designed to gracefully handle errors, including:
  1. [**GraphQL Errors**](https://ferrygraphql.com/docs/error-handling/#graphql-errors): Errors returned by your GraphQL server.
  2. [**Stream Errors**](https://ferrygraphql.com/docs/error-handling/#stream-errors): Errors emitted by any Stream in your TypedLink chain.
  3. [**Forwarding Exceptions**](https://ferrygraphql.com/docs/error-handling/#forwarding-exceptions): Synchronous Exceptions thrown when forwarding a request along the TypedLink chain.


## GraphQL Errors[#](https://ferrygraphql.com/docs/error-handling/#graphql-errors "Direct link to heading")
GraphQL Errors are errors received by your `gql_link`. These are commonly returned by a GraphQL server when there is a problem validating or executing your GraphQL operation (e.g. your Query includes a field not defined in your schema).
Ferry's `GqlTypedLink` (included in the default `Client`) parses GraphQL Errors from responses and includes them in the `OperationResponse.graphqlErrors` field.
## Stream Errors[#](https://ferrygraphql.com/docs/error-handling/#stream-errors "Direct link to heading")
Dart Streams allow errors to be emmited as events in the stream. Ferry's `ErrorTypedLink` (included in the default `Client`) intercepts these Stream error events and converts them into `OperationResponse`s, including the error in the `linkException` field.
## Forwarding Exceptions[#](https://ferrygraphql.com/docs/error-handling/#forwarding-exceptions "Direct link to heading")
Ferry allows `TypedLink`s to be composed into an execution chain by calling `forward()` from within a `TypedLink.request()` method. If an exception is thrown when forwarding a request along the chain, Ferry's `ErrorTypedLink` will catch the exception and convert it into an `OperationResponse`.
As with virtually everything in `Ferry`, the error handling logic can be customized with a [custom](https://ferrygraphql.com/docs/customization) `TypedLink`.
[Edit this page](https://github.com/gql-dart/ferry/edit/master/docs/../docs/error-handling.md)
[Previous« Pagination & Refetching](https://ferrygraphql.com/docs/pagination)
[NextConfiguring the cache »](https://ferrygraphql.com/docs/cache-configuration)
  * [GraphQL Errors](https://ferrygraphql.com/docs/error-handling/#graphql-errors)
  * [Stream Errors](https://ferrygraphql.com/docs/error-handling/#stream-errors)
  * [Forwarding Exceptions](https://ferrygraphql.com/docs/error-handling/#forwarding-exceptions)


#### Docs
  * [Introduction](https://ferrygraphql.com/docs/)


#### Community
  * [Discord](https://discord.gg/G3JGkBY)


#### More
  * [pub.dev](https://pub.dev/packages/ferry)
  * [GitHub](https://github.com/gql-dart/ferry)


Copyright © 2025 Sat Mandir Khalsa. Built with Docusaurus.


## Source Information
- URL: https://ferrygraphql.com/docs/error-handling
- Harvested: 2025-08-19T15:21:01.131451
- Profile: code_examples
- Agent: architect
