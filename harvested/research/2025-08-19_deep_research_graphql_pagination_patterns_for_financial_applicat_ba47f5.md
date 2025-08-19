---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T14:10:16.818411'
profile: deep_research
source: https://graphql.org/learn/pagination/
topic: GraphQL Pagination Patterns for Financial Applications
---

# GraphQL Pagination Patterns for Financial Applications

📣 GraphQLConf 2025 • Sept 08-10 • Amsterdam • [Register Today!](https://graphql.org/conf/2025/)
[](https://graphql.org/)[Learn](https://graphql.org/learn/)
Community
[FAQ](https://graphql.org/faq/)[Spec](https://spec.graphql.org)[Blog](https://graphql.org/blog/)[GraphQLConf](https://graphql.org/conf/2025/)[GraphQL.JS Tutorial](https://graphql.org/graphql-js/)
`CTRL K`
System
`CTRL K`
  * Learn
  * [Introduction](https://graphql.org/learn/)
  * [Schemas and Types](https://graphql.org/learn/schema/)
  * [Queries](https://graphql.org/learn/queries/)
  * [Mutations](https://graphql.org/learn/mutations/)
  * [Subscriptions](https://graphql.org/learn/subscriptions/)
  * [Validation](https://graphql.org/learn/validation/)
  * [Execution](https://graphql.org/learn/execution/)
  * [Response](https://graphql.org/learn/response/)
  * [Introspection](https://graphql.org/learn/introspection/)
  * Best Practices
  * [Best Practices](https://graphql.org/learn/best-practices/)
  * [Thinking in Graphs](https://graphql.org/learn/thinking-in-graphs/)
  * [Serving over HTTP](https://graphql.org/learn/serving-over-http/)
  * [Authorization](https://graphql.org/learn/authorization/)
  * [Pagination](https://graphql.org/learn/pagination/)
  * [Schema Design](https://graphql.org/learn/schema-design/)
  * [Global Object Identification](https://graphql.org/learn/global-object-identification/)
  * [Caching](https://graphql.org/learn/caching/)
  * [Performance](https://graphql.org/learn/performance/)
  * [Security](https://graphql.org/learn/security/)
  * [Federation](https://graphql.org/learn/federation/)
  * GraphQL over HTTP
  * [Common GraphQL over HTTP Errors](https://graphql.org/learn/debug-errors/)


System
On This Page
  * [Plurals](https://graphql.org/learn/pagination/#plurals)
  * [Slicing](https://graphql.org/learn/pagination/#slicing)
  * [Pagination and edges](https://graphql.org/learn/pagination/#pagination-and-edges)
  * [End-of-list, counts, and connections](https://graphql.org/learn/pagination/#end-of-list-counts-and-connections)
  * [Complete connection model](https://graphql.org/learn/pagination/#complete-connection-model)
  * [Connection specification](https://graphql.org/learn/pagination/#connection-specification)
  * [Recap](https://graphql.org/learn/pagination/#recap)


[Question? Give us feedback →](https://github.com/graphql/graphql.github.io/issues/new?title=Feedback%20for%20%E2%80%9CPagination%E2%80%9D&labels=feedback)[Edit this page](https://github.com/graphql/graphql.github.io/tree/source/src/pages/learn/pagination.mdx)Scroll to top
Learn
Pagination
# Pagination
Traverse lists of objects with a consistent field pagination model
A common use case in GraphQL is traversing the relationship between sets of objects. There are different ways that these relationships can be exposed in GraphQL, giving a varying set of capabilities to the client developer. On this page, we’ll explore how fields may be paginated using a cursor-based connection model.
## Plurals[](https://graphql.org/learn/pagination/#plurals)
The simplest way to expose a connection between objects is with a field that returns a plural [List type](https://graphql.org/learn/schema/#list). For example, if we wanted to get a list of R2-D2’s friends, we could just ask for all of them:
Operation
```
xxxxxxxxxx
```

```
query{
```
```
hero{
```
```
name
```
```
friends{
```
```
name
```
```
}
```
```
}
```
```
}
```

Response
```
xxxxxxxxxx
```

```
{
```
```
"data":{
```
```
"hero":{
```
```
"name":"R2-D2",
```
```
"friends":[
```
```
{
```
```
"name":"Luke Skywalker"
```
```
},
```
```
{
```
```
"name":"Han Solo"
```
```
},
```
```
{
```
```
"name":"Leia Organa"
```
```
}
```
```
]
```
```
}
```
```
}
```
```
}
```

## Slicing[](https://graphql.org/learn/pagination/#slicing)
Quickly, though, we realize that there are additional behaviors a client might want. A client might want to be able to specify how many friends they want to fetch—maybe they only want the first two. So we’d want to expose something like this:
```
query {
 hero {
  name
  friends(first: 2) {
   name
  }
 }
}
```

But if we just fetched the first two, we might want to paginate through the list as well; once the client fetches the first two friends, they might want to send a second request to ask for the next two friends. How can we enable that behavior?
## Pagination and edges[](https://graphql.org/learn/pagination/#pagination-and-edges)
There are several ways we could do pagination:
  * We could do something like `friends(first:2 offset:2)` to ask for the next two in the list.
  * We could do something like `friends(first:2 after:$friendId)`, to ask for the next two after the last friend we fetched.
  * We could do something like `friends(first:2 after:$friendCursor)`, where we get a cursor from the last item and use that to paginate.


The approach described in the first bullet is classic _offset-based pagination_. However, this style of pagination can have performance and security downsides, especially for larger data sets. Additionally, if new records are added to the database after the user has made a request for a page of results, then offset calculations for subsequent pages may become ambiguous.
In general, we’ve found that _cursor-based pagination_ is the most powerful of those designed. Especially if the cursors are opaque, either offset or ID-based pagination can be implemented using cursor-based pagination (by making the cursor the offset or the ID), and using cursors gives additional flexibility if the pagination model changes in the future. As a reminder that the cursors are opaque and their format should not be relied upon, we suggest base64 encoding them.
But that leads us to a problem—how do we get the cursor from the object? We wouldn’t want the cursor to live on the `User` type; it’s a property of the connection, not of the object. So we might want to introduce a new layer of indirection; our `friends` field should give us a list of edges, and an edge has both a cursor and the underlying node:
```
query {
 hero {
  name
  friends(first: 2) {
   edges {
    node {
     name
    }
    cursor
   }
  }
 }
}
```

The concept of an edge also proves useful if there is information that is specific to the edge, rather than to one of the objects. For example, if we wanted to expose “friendship time” in the API, having it live on the edge is a natural place to put it.
## End-of-list, counts, and connections[](https://graphql.org/learn/pagination/#end-of-list-counts-and-connections)
Now we can paginate through the connection using cursors, but how do we know when we reach the end of the connection? We have to keep querying until we get an empty list back, but we’d like for the connection to tell us when we’ve reached the end so we don’t need that additional request. Similarly, what if we want additional information about the connection itself, for example, how many friends does R2-D2 have in total?
To solve both of these problems, our `friends` field can return a connection object. The connection object will be an Object type that has a field for the edges, as well as other information (like total count and information about whether a next page exists). So our final query might look more like this:
```
query {
 hero {
  name
  friends(first: 2) {
   totalCount
   edges {
    node {
     name
    }
    cursor
   }
   pageInfo {
    endCursor
    hasNextPage
   }
  }
 }
}
```

Note that we also might include `endCursor` and `startCursor` in this `PageInfo` object. This way, if we don’t need any of the additional information that the edge contains, we don’t need to query for the edges at all, since we got the cursors needed for pagination from `pageInfo`. This leads to a potential usability improvement for connections; instead of just exposing the `edges` list, we could also expose a dedicated list of just the nodes, to avoid a layer of indirection.
## Complete connection model[](https://graphql.org/learn/pagination/#complete-connection-model)
Clearly, this is more complex than our original design of just having a plural! But by adopting this design, we’ve unlocked several capabilities for the client:
  * The ability to paginate through the list.
  * The ability to ask for information about the connection itself, like `totalCount` or `pageInfo`.
  * The ability to ask for information about the edge itself, like `cursor` or `friendshipTime`.
  * The ability to change how our backend does pagination, since the user just uses opaque cursors.


To see this in action, there’s an additional field in the example schema, called `friendsConnection`, that exposes all of these concepts:
```
interface Character {
 id: ID!
 name: String!
 friends: [Character]
 friendsConnection(first: Int, after: ID): FriendsConnection!
 appearsIn: [Episode]!
}
type FriendsConnection {
 totalCount: Int
 edges: [FriendsEdge]
 friends: [Character]
 pageInfo: PageInfo!
}
type FriendsEdge {
 cursor: ID!
 node: Character
}
type PageInfo {
 startCursor: ID
 endCursor: ID
 hasNextPage: Boolean!
}
```

You can try it out in the example query. Try removing the `after` argument for the `friendsConnection` field to see how the pagination will be affected. Also, try replacing the `edges` field with the helper `friends` field on the connection, which lets you get directly to the list of friends without the additional edge layer of indirection, when appropriate for clients:
Operation
```
xxxxxxxxxx
```

```
query{
```
```
hero{
```
```
name
```
```
friendsConnection(first:2, after:"Y3Vyc29yMQ=="){
```
```
totalCount
```
```
edges{
```
```
node{
```
```
name
```
```
}
```
```
cursor
```
```
}
```
```
pageInfo{
```
```
endCursor
```
```
hasNextPage
```
```
}
```
```
}
```
```
}
```
```
}
```

Response
```
xxxxxxxxxx
```

```
{
```
```
"data":{
```
```
"hero":{
```
```
"name":"R2-D2",
```
```
"friendsConnection":{
```
```
"totalCount":3,
```
```
"edges":[
```
```
{
```
```
"node":{
```
```
"name":"Han Solo"
```
```
},
```
```
"cursor":"Y3Vyc29yMg=="
```
```
},
```
```
{
```
```
"node":{
```
```
"name":"Leia Organa"
```
```
},
```
```
"cursor":"Y3Vyc29yMw=="
```
```
}
```
```
],
```
```
"pageInfo":{
```
```
"endCursor":"Y3Vyc29yMw==",
```
```
"hasNextPage":false
```
```
}
```
```
}
```
```
}
```
```
}
```
```
}
```

## Connection specification[](https://graphql.org/learn/pagination/#connection-specification)
To ensure a consistent implementation of this pattern, the Relay project has a formal [specification](https://relay.dev/graphql/connections.htm) you can follow for building GraphQL APIs that use a cursor-based connection pattern - whether or not use you Relay.
## Recap[](https://graphql.org/learn/pagination/#recap)
To recap these recommendations for paginating fields in a GraphQL schema:
  * List fields that may return a lot of data should be paginated
  * Cursor-based pagination provides a stable pagination model for fields in a GraphQL schema
  * The cursor connection specification from the Relay project provides a consistent pattern for paginating the fields in a GraphQL schema


Last updated on August 4, 2025
[Authorization](https://graphql.org/learn/authorization/ "Authorization")[Schema Design](https://graphql.org/learn/schema-design/ "Schema Design")
System
[](https://graphql.org/)
### Learn
  * [Introduction to GraphQL](https://graphql.org/learn/)
  * [Best Practices](https://graphql.org/learn/best-practices/)
  * [Frequently Asked Questions](https://graphql.org/faq/)
  * [Training Courses](https://graphql.org/community/resources/training-courses/)


### Code
  * [GitHub](https://github.com/graphql)
  * [GraphQL Specification](https://spec.graphql.org)
  * [Libraries & Tools](https://graphql.org/code/)
  * [Services & Vendors](https://graphql.org/code/?tags=services)


### Community
  * [Resources](https://graphql.org/community/resources/official-channels/)
  * [Events & Meetups](https://graphql.org/community/events/)
  * [Contribute to GraphQL](https://graphql.org/community/contribute/essential-links/)
  * [Landscape](https://landscape.graphql.org)
  * [Shop](https://store.graphql.org)


### & More
  * [Blog](https://graphql.org/blog/)
  * [GraphQL Foundation](https://graphql.org/foundation/)
  * [GraphQL Community Grant](https://graphql.org/foundation/community-grant/)
  * [Logo and Brand Guidelines](https://graphql.org/brand/)
  * [Code of Conduct](https://graphql.org/codeofconduct/)


Copyright © 2025 The GraphQL Foundation. All rights reserved.For web site terms of use, trademark policy and general project policies please see <https://lfprojects.org>
  * [](https://github.com/graphql)
  * [](https://discord.graphql.org)
  * [](https://twitter.com/graphql)
  * [](http://stackoverflow.com/questions/tagged/graphql)

[Powered by ](https://nextra.site)


## Source Information
- URL: https://graphql.org/learn/pagination/
- Harvested: 2025-08-19T14:10:16.818411
- Profile: deep_research
- Agent: architect
