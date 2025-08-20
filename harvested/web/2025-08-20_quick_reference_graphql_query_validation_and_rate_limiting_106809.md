---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-20T00:56:52.174517'
profile: quick_reference
source: https://graphql.org/learn/validation/
topic: GraphQL Query Validation and Rate Limiting
---

# GraphQL Query Validation and Rate Limiting

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
  * [Validation examples](https://graphql.org/learn/validation/#validation-examples)
  * [Requesting non-existent fields](https://graphql.org/learn/validation/#requesting-non-existent-fields)
  * [Selection sets and leaf fields](https://graphql.org/learn/validation/#selection-sets-and-leaf-fields)
  * [Missing fragments for fields that output abstract types](https://graphql.org/learn/validation/#missing-fragments-for-fields-that-output-abstract-types)
  * [Cyclic fragment spreads](https://graphql.org/learn/validation/#cyclic-fragment-spreads)
  * [Validation errors](https://graphql.org/learn/validation/#validation-errors)
  * [Next steps](https://graphql.org/learn/validation/#next-steps)


[Question? Give us feedback →](https://github.com/graphql/graphql.github.io/issues/new?title=Feedback%20for%20%E2%80%9CValidation%E2%80%9D&labels=feedback)[Edit this page](https://github.com/graphql/graphql.github.io/tree/source/src/pages/learn/validation.mdx)Scroll to top
Learn
Validation
# Validation
Learn how GraphQL validates operations using a schema
On this page, we’ll explore an important phase in the lifecycle of a GraphQL request called [validation](https://spec.graphql.org/draft/#sec-Validation). A request must be syntactically correct to run, but it should also be valid when checked against the API’s schema.
In practice, when a GraphQL operation reaches the server, the document is first parsed and then validated using the type system. This allows servers and clients to effectively inform developers when an invalid query has been created, without relying on runtime checks. Once the operation is validated, it can be [executed](https://graphql.org/learn/execution/) on the server and a response will be delivered to the client.
## Validation examples[](https://graphql.org/learn/validation/#validation-examples)
The GraphQL specification describes the detailed conditions that must be satisfied for a request to be considered valid. In the sections that follow, we’ll look at a few examples of common validation issues that may occur in GraphQL operations.
### Requesting non-existent fields[](https://graphql.org/learn/validation/#requesting-non-existent-fields)
When we query for a field, the field must be defined on the relevant type. As `hero` returns a `Character` type, its selection set may only request the `Character` type’s fields; `Character` does not have a `favoriteSpaceship` field, so this query is invalid:
Operation
```
xxxxxxxxxx
```

```
# INVALID: favoriteSpaceship does not exist on Character
```
```
query{
```
```
hero{
```
```
favoriteSpaceship
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
"errors":[
```
```
{
```
```
"message":"Cannot query field \"favoriteSpaceship\" on type \"Character\".",
```
```
"locations":[
```
```
{
```
```
"line":4,
```
```
"column":5
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
]
```
```
}
```

### Selection sets and leaf fields[](https://graphql.org/learn/validation/#selection-sets-and-leaf-fields)
Whenever we query for a field and it returns something other than a Scalar or Enum type, we need to specify what data we want to get back from the field (a “selection set”). The `hero` query field returns a `Character`, and we’ve already seen examples that request fields like `name` and `appearsIn` on it. If we omit those leaf field selections, then the query will not be valid:
Operation
```
xxxxxxxxxx
```

```
# INVALID: hero is not a scalar, so fields are needed
```
```
query{
```
```
hero
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
"errors":[
```
```
{
```
```
"message":"Field \"hero\" of type \"Character\" must have a selection of subfields. Did you mean \"hero { ... }\"?",
```
```
"locations":[
```
```
{
```
```
"line":3,
```
```
"column":3
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
]
```
```
}
```

Similarly, querying fields of a scalar or enum doesn’t make sense, therefore adding a selection set to a leaf field will make the query invalid:
Operation
```
xxxxxxxxxx
```

```
# INVALID: name is a scalar, so fields are not permitted
```
```
query{
```
```
hero{
```
```
name{
```
```
firstCharacterOfName
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
"errors":[
```
```
{
```
```
"message":"Field \"name\" must not have a selection since type \"String!\" has no subfields.",
```
```
"locations":[
```
```
{
```
```
"line":4,
```
```
"column":10
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
]
```
```
}
```

### Missing fragments for fields that output abstract types[](https://graphql.org/learn/validation/#missing-fragments-for-fields-that-output-abstract-types)
Earlier, it was noted that a query can only ask for fields on the type in question. So when we query for `hero` which returns a `Character`, we can only request fields that exist on the `Character` Interface type. What happens if we want to query for R2-D2’s primary function, though?
Operation
```
xxxxxxxxxx
```

```
# INVALID: primaryFunction does not exist on Character
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
primaryFunction
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
"errors":[
```
```
{
```
```
"message":"Cannot query field \"primaryFunction\" on type \"Character\". Did you mean to use an inline fragment on \"Droid\"?",
```
```
"locations":[
```
```
{
```
```
"line":5,
```
```
"column":5
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
]
```
```
}
```

That query is invalid, because `primaryFunction` is not one of the shared fields defined by the `Character` Interface type. We want some way of indicating that we wish to fetch `primaryFunction` if the `Character` is a `Droid`, and to ignore that field otherwise. We can use the [fragments](https://graphql.org/learn/queries/#fragments) to do this. By setting up a fragment defined on `Droid` and including it in the selection set, we ensure that we only query for `primaryFunction` where it is defined:
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
...DroidFields
```
```
}
```
```
}
```
```
​
```
```
fragmentDroidFieldsonDroid{
```
```
primaryFunction
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
"primaryFunction":"Astromech"
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

This query is valid, but it’s a bit verbose; named fragments were valuable above when we used them multiple times, but we’re only using this one once. Instead of using a named fragment, we can use an [inline fragment](https://graphql.org/learn/queries/#inline-fragments); this still allows us to indicate the type we are querying on, but without naming a separate fragment:
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
...onDroid{
```
```
primaryFunction
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
"primaryFunction":"Astromech"
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

### Cyclic fragment spreads[](https://graphql.org/learn/validation/#cyclic-fragment-spreads)
To start, let’s take a complex valid query. This is a nested query, but with the duplicated fields factored out into a fragment:
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
...NameAndAppearances
```
```
friends{
```
```
...NameAndAppearances
```
```
friends{
```
```
...NameAndAppearances
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
​
```
```
fragmentNameAndAppearancesonCharacter{
```
```
name
```
```
appearsIn
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
"appearsIn":[
```
```
"NEWHOPE",
```
```
"EMPIRE",
```
```
"JEDI"
```
```
],
```
```
"friends":[
```
```
{
```
```
"name":"Luke Skywalker",
```
```
"appearsIn":[
```
```
"NEWHOPE",
```
```
"EMPIRE",
```
```
"JEDI"
```
```
],
```
```
"friends":[
```
```
{
```
```
"name":"Han Solo",
```
```
"appearsIn":[
```
```
"NEWHOPE",
```
```
"EMPIRE",
```
```
"JEDI"
```
```
]
```
```
},
```
```
{
```
```
"name":"Leia Organa",
```
```
"appearsIn":[
```
```
"NEWHOPE",
```
```
"EMPIRE",
```
```
"JEDI"
```
```
]
```
```
},
```
```
{
```
```
"name":"C-3PO",
```
```
"appearsIn":[
```
```
"NEWHOPE",
```
```
"EMPIRE",
```
```
"JEDI"
```
```
]
```
```
},
```
```
{
```
```
"name":"R2-D2",
```
```
"appearsIn":[
```
```
"NEWHOPE",
```
```
"EMPIRE",
```
```
"JEDI"
```
```
]
```
```
}
```
```
]
```
```
},
```
```
{
```
```
"name":"Han Solo",
```
```
"appearsIn":[
```
```
"NEWHOPE",
```
```
"EMPIRE",
```
```
"JEDI"
```
```
],
```
```
"friends":[
```
```
{
```
```
"name":"Luke Skywalker",
```
```
"appearsIn":[
```
```
"NEWHOPE",
```
```
"EMPIRE",
```
```
"JEDI"
```
```
]
```
```
},
```
```
{
```
```
"name":"Leia Organa",
```
```
"appearsIn":[
```
```
"NEWHOPE",
```
```
"EMPIRE",
```
```
"JEDI"
```
```
]
```
```
},
```
```
{
```
```
"name":"R2-D2",
```
```
"appearsIn":[
```
```
"NEWHOPE",
```
```
"EMPIRE",
```
```
"JEDI"
```
```
]
```
```
}
```
```
]
```
```
},
```
```
{
```
```
"name":"Leia Organa",
```
```
"appearsIn":[
```
```
"NEWHOPE",
```
```
"EMPIRE",
```
```
"JEDI"
```
```
],
```
```
"friends":[
```
```
{
```
```
"name":"Luke Skywalker",
```
```
"appearsIn":[
```
```
"NEWHOPE",
```
```
"EMPIRE",
```
```
"JEDI"
```
```
]
```
```
},
```
```
{
```
```
"name":"Han Solo",
```
```
"appearsIn":[
```
```
"NEWHOPE",
```
```
"EMPIRE",
```
```
"JEDI"
```
```
]
```
```
},
```
```
{
```
```
"name":"C-3PO",
```
```
"appearsIn":[
```
```
"NEWHOPE",
```
```
"EMPIRE",
```
```
"JEDI"
```
```
]
```
```
},
```
```
{
```
```
"name":"R2-D2",
```
```
"appearsIn":[
```
```
"NEWHOPE",
```
```
"EMPIRE",
```
```
"JEDI"
```
```
]
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

The following is an alternative to the above query, attempting to use recursion instead of the explicit three levels of nesting. This new query is invalid because a fragment cannot refer to itself (directly or indirectly) since the resulting cycle could create an unbounded result!
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
...NameAndAppearancesAndFriends
```
```
}
```
```
}
```
```
​
```
```
fragmentNameAndAppearancesAndFriendsonCharacter{
```
```
name
```
```
appearsIn
```
```
friends{
```
```
...NameAndAppearancesAndFriends
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
"errors":[
```
```
{
```
```
"message":"Cannot spread fragment \"NameAndAppearancesAndFriends\" within itself.",
```
```
"locations":[
```
```
{
```
```
"line":11,
```
```
"column":5
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
]
```
```
}
```

This has just scratched the surface of the validation system; there are a number of validation rules in place to ensure that a GraphQL operation is semantically meaningful. The specification goes into more detail about this topic in the [validation section](https://spec.graphql.org/draft/#sec-Validation), and the [validation directory in the reference implementation](https://github.com/graphql/graphql-js/blob/main/src/validation) contains code implementing a specification-compliant GraphQL validator.
## Validation errors[](https://graphql.org/learn/validation/#validation-errors)
As we have seen in the examples above, when a GraphQL server encounters a validation error in a request, it will return information about what happened in the `errors` key of the response. Specifically, when GraphQL detects a validation issue in a document, it raises a _request error_ before execution begins, meaning that no partial data will be included in the response.
And because the GraphQL specification requires all implementations to validate incoming requests against the schema, developers won’t need to write specific runtime logic to handle these validation issues manually.
## Next steps[](https://graphql.org/learn/validation/#next-steps)
To recap what we’ve learned about validation:
  * To be executed, requests must include a syntactically correct document that is considered valid when checked against the schema
  * The specification requires implementations check incoming requests contain valid field selections, correct fragment usage, and more
  * When a validation issue occurs, the server will raise a request error and return to the client information about what happened; field execution will not start


Head over to the [Execution](https://graphql.org/learn/execution/) page to learn how GraphQL provides data for each field in a request after the validation step successfully completes.
Last updated on August 4, 2025
[Subscriptions](https://graphql.org/learn/subscriptions/ "Subscriptions")[Execution](https://graphql.org/learn/execution/ "Execution")
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
- URL: https://graphql.org/learn/validation/
- Harvested: 2025-08-20T00:56:52.174517
- Profile: quick_reference
- Agent: architect
