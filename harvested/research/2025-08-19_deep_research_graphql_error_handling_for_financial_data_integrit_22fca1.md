---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T14:11:03.751676'
profile: deep_research
source: https://productionreadygraphql.com/2020-08-01-guide-to-graphql-errors
topic: GraphQL Error Handling for Financial Data Integrity
---

# GraphQL Error Handling for Financial Data Integrity

[Production Ready GraphQL{ blog }](https://productionreadygraphql.com/)
# A Guide to GraphQL Errors
![](https://productionreadygraphql.com/2020-08-01-guide-to-graphql-errors)
[Marc-Andre Giroux](https://twitter.com/__xuorig__)
August 01, 2020 | 14 min read
🌱 _This post is still growing and likely will change as the best practices evolve_ 🌳
GraphQL errors are something many of us struggle with. Good practices are emerging but the community has not yet settled on a convention. I hope this guide helps demystifiy the many ways you can structure errors in your GraphQL servers, and the tradeoffs each of them make.
## Stage 1: GraphQL Errors AKA Top-Level Errors
The GraphQL Specification [describes errors](http://spec.graphql.org/draft/#sec-Errors) quite well, so why are they even confusing in the first place?
```
{
 "data": {
  "createProduct": null
 },
 "errors": [
  {
   "path": [
    "createProduct"
   ],
   "locations": [
    {
     "line": 2,
     "column": 3
    }
   ],
   "message": "Could not resolve to a node with the global id of 'shop-that-does-not-exist'"
  }
 ]
}
```

### ❗️Con: No Schema
We love GraphQL for it's type system. The `errors` key of a GraphQL response, commonly called "Top-Level" errors, don't have such schema. They are specified, but introspection tells a client nothing about what might be found in them, they're harder to evolve, and harder to extend. The `extensions` key allows providers to add extra fields but then we are out of the specification, and they need to be documented in an ad hoc way, outside of the schema. That's not ideal.
### ❗️Con: Nullability
The other potential issue with top-level errors is that when they are present, the corresponding field should be `null`. This can potentially be a deal breaker if you're wanting to have errors returned as part of a mutation, but want to query for data on the result anyways. A common example of this is the server sending back the actual state of a resource after a mutation that had errors. With top-level errors this can't be done since the whole mutation field should be `null`!
### ❗️Con: Exceptional Only
"Top-Level" Errors are today generally accepted as a way to represent "exceptional" errors, and errors that are **developer facing**. [Lee Byron] [makes that clear](https://github.com/graphql/graphql-spec/issues/391#issuecomment-385553207) in a few [comments](https://github.com/graphql/graphql-spec/issues/117#issuecomment-170180628) on the GraphQL-Spec repository.
> GraphQL errors encode exceptional scenarios - like a service being down or some other internal failure. Errors which are part of the API domain should be captured within that domain.
> The general philosophy at play is that Errors are considered exceptional. Your user data should never be represented as an Error. If your users can do something that needs to provide negative guidance, then you should represent that kind of information in GraphQL as Data not as an Error. Errors should always represent either developer errors or exceptional circumstances (e.g. the database was offline).
While it's possible to add extensions to make these errors better to work with for end user errors, the fact these errors have no schema and that they can't often be collocated with data makes them a less than ideal choice for these use cases.
### ✅Pro: Great for Developer Errors
Service unavailable? Syntax errors? Rate limited? Timeout? These types of errors are perfect for "Top-Level" errors and you should absolutely use them. Make them even better with error codes so clients can handle them in a better way. Hopefully some common error codes can eventually be extracted for consistency across different APIs and better clients.
### ✅Pro: As Close of a Convention We've Got!
These errors are specified in GraphQL's specification, so that's a definite pro. Clients already know how to deal with them and we've got consistency across different API, that's great!
## Stage 2: Ad Hoc Error Fields
OK so if top-level errors are not ideal for user-facing errors, we have to start thinking of a way to expose those **directly in our schemas**. The simplest way to do this is by adding fields to our mutation payloads, here's an example:
```
type Mutation {
 createUser(input: CreateUserInput!): CreateUserPayload
}
# 💡The "MutationPayload" wrapper type is a common convention
# initially specified by the GraphQL client Relay 
type CreateUserPayload {
 user: User
 userNameWasTaken: Boolean!
 userNameAlternative: String!
}
```

In this example we've added two fields to a mutation payload type that helps us deal with errors. The `createUser` may error out when the username someone wants to create already exists. In this case we want to display an error and also propose a better username based on what they entered.
### ✅Pro: Discoverability
It's quite easy to see that it's possible for this mutation to return, and what kind of errors may happen. If we had to handle this with top-level errors, we would've had to add a `userNameAlternative` extension to errors, only in the case of the `createUser` mutation. That is quite hard for a client to learn about and use.
### ✅Pro: Simplicity
This solution is also the simplest to use. We look at what the client may need from an error, and we add it directly to the schema. It also really makes you think about errors from a client perspective, what a client truly needs when the mutation is unable to fulfill the use case.
### ❗️Con: Impossible States
One annoying part of just adding error fields to the mutation payload is that we allow for some impossible states in theory. In the example above we've made the `user` field nullable because if the username was taken the user will be null and `userNameWasTaken` will be true.
While we may know that intuitively, the schema doesn't really tell us that. In theory, `userNameWasTaken` could be true, and `user` could be there as well. What does this mean? It's unclear.
This behavior instead probably needs to be described in documentation or descriptions on fields, which is never ideal when we've got a schema in the first place!
### ❗️Con: Consistency
The other downside of this approach is we don't standardize mutation errors as a whole. Every mutation will have a set of ad hoc fields to describe what could go wrong, which means building a generic client, or even finding errors for a human can be more troublesome.
## Stage 3: Error Array
When ad-hoc error fields are not enough, we usually start looking for something a bit more structured. One way to do that is by starting to build an error type!
```
type CreateUserPayload {
 user: User
 userErrors: [UserError!]!
}
type UserError {
 # A description of the error
 message: String!
 # A path to the input value that caused the error
 path: [String!]
}
```

### _A small note on the UserError.path field_
In the example above, we have a `path` field on the `UserError` type. The field is a list of strings, pointing to the input value that caused the error, if any.
For example, imagine I try to create a user where the `username`is already taken:
```
mutation {
 createUser(input: { username: "xuorig" }) {
  user
  userErrors {
   message
   path
  }
 }
}
```

We would get a response that looks like this:
```
{
 "data": {
  "createUser": {
   "user": null,
   "userErrors": [{
    "message": "Username `xuorig` was already taken.",
    "path": ["input", "username"]
   }]
  }
 }
}
```

Notice how `path` points to the `username` input field we had used. This is really helpful for clients, for example by pointing to the form field that caused the error. OK, back to the error array.
This is one of the first ways to do errors in a more structured way that I've encountered, one we used at Shopify back in the days 👴 (since then their approach has evolved).
### ✅Pro: Consistency
Now we've got a common contract for our errors, which is great. Clients know errors that happen on a mutation will have both a `message` and a `path`, and we're free to extend this with other metadata as we like.
### ✅Pro: Evolution
Because our errors are now part of a list type and we have that common `UserError` type, clients can start handling errors in a more future-proof way. When the server adds a new possible error, clients that handle a list of errors simply get it for free. With ad-hoc fields the client often needs to implement new custom logic to handle the new case.
### ❗️Con: Impossible States
Just like the ad-hoc error fields, our schema still allows for potentially impossible states, such as having the user in the response, but also items present in the `userErrors` list.
### ❗️Con: User has to select
This con is common to many of the solutions in this article, but I'm choosing to introduce it here. One downside of all "errors in schema" approaches is that it's quite easy for a client to ignore errors entirely by not selecting the error fields.
As a side-note, one idea I've been playing with is the idea of something like a `@mustSelect` schema directive / error. Certain fields could be annotated in a way to produces an error if a client does not select the field.
```
type CreateUserPayload {
 userErrors: [UserError!]! @mustSelect
}
```

If the client is really not interested in the field, it has to opt-out explicitly:
```
mutation {
 createUser(input: {}) @ignoreSelection(fields: ["userErrors"]) {
  user
 }
}
```

The cool thing about this is that we don't force the response to include the `userErrors` field, which would ruin GraphQL's declarativeness. Instead we hope the user will discover the error at dev time allowing them to explicitly opt-out, or select the field
### ❗️Con: Custom Error Fields
Because we have only one `UserError` type, it's harder to support fields that are really specific to one particular error. For example, adding a `userNameSuggestion` field make little sense on the generic `UserError` type.
## Stage 4: Error Interface
A natural evolution of the user error list is to have an interface type as `UserError`, and then concrete implementations when needed.
```
type CreateUserPayload {
 user: User
 userErrors: [UserError!]!
}
type UserNameTakenError implements UserError {
 userNameSuggestion: String
 message: String!
 path: [String!]
}
type SomeOtherError implements UserError {}
interface UserError {
 # A description of the error
 message: String!
 # A path to the input value that caused the error
 path: [String!]
}
```

### ✅Pro: Evolution
Just like the plain list type solution from before, this structure is quite nice to evolve over time as well. This time though, we have a common interface for errors, meaning clients know the common contract between errors, but also it allows us to create new, more specific error types when needed.
```
mutation {
 createUser(input: {}) {
  user { id }
  userErrors {
   message
   path
   ... on UserNameTakenError {
    userNameSuggestion
   }
  }
 }
}
```

### ❗️Con: Hard to See What Errors May Happen
We still have a small problem here. When a client integrates with our API, it's quite hard to know the possible set of errors that may happen when executing an operation. We know they all have a common contract, but the schema doesn't tell us exactly what possible types this particular mutation may return.
There are a few ways around this:
  1. A) We could keep it a bit more generic and have a list of `UserCreationError` instead.
  2. B) We could have a more specific UserCreation interface, + the `UserError` interface
  3. C) We could look at Union types instead.


## Stage 5: Result Types
Another popular approach to representing errors in GraphQL is by using union types. There is a [great post by Sasha Solomon (Twitter)](https://medium.com/@sachee/200-ok-error-handling-in-graphql-7ec869aec9bc) that explains the philosophy behind it in more detail.
Error unions are great because they are a really expressive way of structuring our schema. It lets clients see right away what could happen when querying or mutating a resource.
```
type Mutation {
 createUser(input: CreateUserInput): CreateUserResult
}
union CreateUserResult = UserCreated | UserNameTaken
type UserCreated {
 user: User!
}
type UserNameTaken {
 message: String!
 suggestion: String!
}
```

### ✅Pro: No Impossible States
With union result types like these, we solve the "Make Impossible States Impossible" problem for earlier. Notice how we can now make all fields non-null because we get the right type in the right context!
### ✅Pro: Works Great on the Query Side
This approach works great on the query side of things as well. In fact the example in Sachee's article is about a `User` type and the examples are of queries.
### ✅Pro: Discoverability
Looking at the schema quickly tells us the possible results we might get. This can be really great for documentation as well.
### 🟡Warning: Potentially Hard to Implement
This really depends on your implementation, but in my experience, it can be hard to know ahead of time and statically, what kind of errors may happen at runtime. Note that this might be a good opportunity to make that better!
### ❗️Con: Hard to Evolve
With a straight union result type like this, it's a bit harder to evolve our schema with new types of errors. You can imagine clients are querying your schema in this way:
```
mutation {
 createUser(input: {}) {
  ... on UserCreated {
   user { id }
  }
  ... on UserNameTaken {
   message
  }
 }
}
```

What if a new error type is added? Like password too short? Clients have no way to adapt to this change beforehand because we're dealing with a union type.
### ❗️Con: Multiple Errors
This specific implementation of a union type result also doesn't really allow us to return more than one error scenario. For example if we wanted to both show that the username is taken and the password is too short, we'd have to design the schema in another way:
  1. A) By having a more generic `UserCreationError` which can host multiple sub errors
  2. B) Going for a `userErrors` list type, which uses a union


## Stage 6: Error Union List
We can use unions to improve on our `userErrors` list as well.
```
type CreateUserPayload {
 user: User
 userErrors: [CreateUserError!]!
}
union CreateUserError = UserNameTaken | PasswordTooShort | MustBeOver18
type UserNameTaken {
 message: String!
 suggestion: String!
}
```

### ✅Pro: Expressive and Discoverable Schema
The union makes it easy to see what can possibly go wrong when creating a user.
### ✅Pro: Support for Multiple Errors
Unlike the result type we saw in the previous solution, we can support forms with multiple errors because we've got a list type of the union.
### ❗️Con: Hard to Evolve
Just like the error result type we saw before, it's hard to add to the `CreateUserError` type without clients being out of date.
## Stage 6a: Error Union List + Interface
What if we could combine the extensibility of the interface with the expressivity of the union? Well, we can!
```
type CreateUserPayload {
 user: User
 userErrors: [CreateUserError!]!
}
union CreateUserError = UserNameTaken | PasswordTooShort | MustBeOver18
type UserNameTaken implements UserError {
 message: String!
 path: String!
 suggestion: String!
}
interface UserError {
 message: String!
 path: String!
}
```

> Make sure all of your union members implement a common interface for this solution to work. I recommend you look into a schema linter to make sure that's the case. [graphql-schema-linter](https://github.com/cjoudrey/graphql-schema-linter) is great for this.
### ✅Pro: Expressive and Discoverable Schema
### ✅Pro: Support for Multiple Errors
### ✅Pro: Easier Evolution
Clients can now select on specific errors, but also fallback to the interface contract, meaning they will never miss a new error!
```
mutation {
 createUser(input: {}) {
  user { id }
  userErrors {
   # Specific cases
   ... on UserNameTaken {
    message
    path
    suggestion
   }
   # Interface contract
   ... on UserError {
    message
    path
   }
  }
 }
}
```

### ❗️Con: Quite Verbose!
This solution does have a lot of moving pieces, and adds a lot of types to a schema. From a developer experience, this can be hidden behind great abstractions if you're building your schema from a code first approach.
## Final Stage: Picking The Best Thing for You
You made it through! In every stage we increased the accuracy, but also verbosity and complexity of our error schemas. More is not always better and it's very possible you'll end up picking something between ad-hoc errors and a fully fledged union type + interface. As a general rule of thumb, with a public API you'll most likely want something more structured that can evolve as easy as possible, while something used by a single client in an internal context might be able to get away with something more simple for a long time 👍
I'll keep working on this post as new solutions arise, in the mean time, happy schema building!
If you've enjoyed this post, you might like the [Production Ready GraphQL book](https://book.productionreadygraphql.com), which I have just released!
Thanks for reading 💚
## Sign up for my newsletter
Stay up to date when I release courses, posts, and anything related to GraphQL
Sign me Up
No spam, just great GraphQL content!
© 2020 MYUL Digital, Inc. All rights reserved.


## Source Information
- URL: https://productionreadygraphql.com/2020-08-01-guide-to-graphql-errors
- Harvested: 2025-08-19T14:11:03.751676
- Profile: deep_research
- Agent: architect
