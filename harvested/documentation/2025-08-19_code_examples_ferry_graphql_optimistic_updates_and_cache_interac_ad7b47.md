---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T15:21:21.125721'
profile: code_examples
source: https://ferrygraphql.com/docs/cache-interaction
topic: Ferry GraphQL Optimistic Updates and Cache Interaction
---

# Ferry GraphQL Optimistic Updates and Cache Interaction

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
  * [Ferry](https://ferrygraphql.com/docs/cache-interaction/#!)
    * [Introduction](https://ferrygraphql.com/docs/)
    * [How It Works](https://ferrygraphql.com/docs/how-it-works)
  * [Getting Started](https://ferrygraphql.com/docs/cache-interaction/#!)
    * [Setup](https://ferrygraphql.com/docs/setup)
    * [Generate Classes](https://ferrygraphql.com/docs/codegen)
  * [Fetching](https://ferrygraphql.com/docs/cache-interaction/#!)
    * [Queries](https://ferrygraphql.com/docs/queries)
    * [Mutations](https://ferrygraphql.com/docs/mutations)
    * [Fetch Policies](https://ferrygraphql.com/docs/fetch-policies)
    * [Pagination & Refetching](https://ferrygraphql.com/docs/pagination)
    * [Error Handling](https://ferrygraphql.com/docs/error-handling)
  * [Caching](https://ferrygraphql.com/docs/cache-interaction/#!)
    * [Configuration](https://ferrygraphql.com/docs/cache-configuration)
    * [Reading and Writing](https://ferrygraphql.com/docs/cache-interaction)
    * [Garbage Collection & Eviction](https://ferrygraphql.com/docs/garbage-collection)
  * [Flutter](https://ferrygraphql.com/docs/cache-interaction/#!)
    * [Using with Flutter](https://ferrygraphql.com/docs/flutter)
    * [Operation Widget](https://ferrygraphql.com/docs/flutter-operation-widget)
    * [Structuring Queries](https://ferrygraphql.com/docs/structuring-queries)
  * [Advanced](https://ferrygraphql.com/docs/cache-interaction/#!)
    * [Custom Scalars](https://ferrygraphql.com/docs/custom-scalars)
    * [Customizing the Client](https://ferrygraphql.com/docs/customization)
    * [Isolates](https://ferrygraphql.com/docs/isolates)
    * [JSON Operations](https://ferrygraphql.com/docs/json-operation)


# Reading and Writing Data to the Cache
Ferry provides the following methods for reading and writing data to the cache:
  * [`readQuery`](https://ferrygraphql.com/docs/cache-interaction/#readquery) and [`readFragment`](https://ferrygraphql.com/docs/cache-interaction/#readfragment)
  * [`writeQuery` and `writeFragment`](https://ferrygraphql.com/docs/cache-interaction/#writequery-and-writefragment)


These methods are described in detail below.
All code samples below assume that you have initialized an instance of `Cache`.
## `readQuery`[#](https://ferrygraphql.com/docs/cache-interaction/#readquery "Direct link to heading")
The `readQuery` method enables you to run a GraphQL query directly on your cache.
  1. If your cache contains all of the data necessary to fulfill a specified query, `readQuery` returns a data object in the shape of that query, just like a GraphQL server does.
  2. If your cache _doesn't_ contain all of the data necessary to fulfill a specified query, `readQuery` returns null. It _never_ attempts to fetch data from a remote server.


Let's assume we've created the following `reviews.graphql` file and [generated](https://ferrygraphql.com/docs/codegen) our GraphQL classes.
Copy
query Reviews($first: Int,$offset: Int){
reviews(first:$first,offset:$offset){
id
stars
commentary
}
}
We can then read our Query like so:
Copy
final reviewsReq =GReviewsReq();
final reviewData = cache.readQuery(reviewsReq);
You can provide GraphQL variables to `readQuery` like so:
Copy
final reviewsReq =GReviewsReq(
(b)=> b
..vars.first =5
..vars.offset =0,
);
final reviewData = cache.readQuery(reviewsReq);
## `readFragment`[#](https://ferrygraphql.com/docs/cache-interaction/#readfragment "Direct link to heading")
The `readFragment` method enables you to read data from _any_ normalized cache object that was stored as part of _any_ query result. Unlike with `readQuery`, calls to `readFragment` do not need to conform to the structure of one of your data graph's supported queries.
For example, if we have the following `hero_data.graphql` file:
Copy
fragmentReviewFragmentonReview{
stars
commentary
}
We could read the generated fragment like so:
Copy
final reviewFragmentReq =GReviewFragmentReq(
(b)=> b..idFields ={'id':'123'},
);
final reviewFragmentData = cache.readFragment(reviewFragmentReq);
The `idFields` argument is the set of [unique identifiers](https://ferrygraphql.com/docs/cache-configuration#generating-unique-identifiers) and their values for the object you want to read from the cache. By default, this is the value of the object's `id` field, but you can [customize this behavior](https://ferrygraphql.com/docs/cache-configuration#generating-unique-identifiers).
In the example above:
  * If a `Review` object with an `id` of `123` is _not_ in the cache, `readFragment` returns `null`.
  * If the `Review` object _is_ in the cache but it's missing either a `stars` or `commentary` field, `readFragment` returns `null`.


## `writeQuery` and `writeFragment`[#](https://ferrygraphql.com/docs/cache-interaction/#writequery-and-writefragment "Direct link to heading")
In addition to reading arbitrary data from the cache, you can _write_ arbitrary data to the cache with the `writeQuery` and `writeFragment` methods.
##### note
Any changes you make to cached data with `writeQuery` and `writeFragment` are not pushed to your GraphQL server.
These methods have the same signature as their `read` counterparts, except they require an additional `data` variable.
For example, the following call to `writeFragment` _locally_ updates the `Review` object with an `id` of `123`:
Copy
final reviewFragmentReq =GReviewFragmentReq(
(b)=> b..idFields ={'id':'123'},
);
final reviewFragmentData =GReviewFragmentData(
(b)=> b
..stars =5
..commentary ='I watched it again and loved it',
);
cache.writeFragment(reviewFragmentReq, reviewFragmentData);
Any `operationRequest` streams that are listening to this data will be updated accordingly.
## Combining reads and writes[#](https://ferrygraphql.com/docs/cache-interaction/#combining-reads-and-writes "Direct link to heading")
Combine `readQuery` and `writeQuery` to fetch currently cached data and make selective modifications to it.
##### note
Since Ferry's generated classes are based on `built_value`, we can easily create modified copies of them using the `rebuild` method. Check out [this post](https://medium.com/dartlang/darts-built-value-for-immutable-object-models-83e2497922d4) to learn more about `built_value` objects.
The example below updates the star rating for the cached `Review`. Remember, this addition is _not_ sent to your remote server.
Copy
final reviewFragmentReq =GReviewFragmentReq(
(b)=> b..idFields ={'id': review.id},
);
final data = cache.readFragment(reviewFragmentReq);
cache.writeFragment(
reviewFragmentReq,
data.rebuild((b)=> b..stars =4),
);
[Edit this page](https://github.com/gql-dart/ferry/edit/master/docs/../docs/cache-interaction.md)
[Previous« Configuring the cache](https://ferrygraphql.com/docs/cache-configuration)
[NextGarbage Collection and Eviction »](https://ferrygraphql.com/docs/garbage-collection)
  * [`readQuery`](https://ferrygraphql.com/docs/cache-interaction/#readquery)
  * [`readFragment`](https://ferrygraphql.com/docs/cache-interaction/#readfragment)
  * [`writeQuery` and `writeFragment`](https://ferrygraphql.com/docs/cache-interaction/#writequery-and-writefragment)
  * [Combining reads and writes](https://ferrygraphql.com/docs/cache-interaction/#combining-reads-and-writes)


#### Docs
  * [Introduction](https://ferrygraphql.com/docs/)


#### Community
  * [Discord](https://discord.gg/G3JGkBY)


#### More
  * [pub.dev](https://pub.dev/packages/ferry)
  * [GitHub](https://github.com/gql-dart/ferry)


Copyright © 2025 Sat Mandir Khalsa. Built with Docusaurus.


## Source Information
- URL: https://ferrygraphql.com/docs/cache-interaction
- Harvested: 2025-08-19T15:21:21.125721
- Profile: code_examples
- Agent: architect
