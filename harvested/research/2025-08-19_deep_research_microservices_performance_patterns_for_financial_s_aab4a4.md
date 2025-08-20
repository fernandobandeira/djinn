---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T23:12:29.749656'
profile: deep_research
source: https://microservices.io/patterns/data/database-per-service.html
topic: Microservices Performance Patterns for Financial Systems
---

# Microservices Performance Patterns for Financial Systems

#### [Microservice Architecture](https://microservices.io/index.html)
**Supported by[Kong](https://konghq.com/)**
  * [Patterns](https://microservices.io/patterns/index.html)
  * [Articles](https://microservices.io/articles/index.html)
  * [Presentations](https://microservices.io/presentations/index.html)
  * [Adoptnew](https://microservices.io/adopt/index.html)
  * [Refactoringnew](https://microservices.io/refactoring/index.html)
  * [Testingnew](https://microservices.io/testing/index.html)


  * [About](https://microservices.io/about.html)


# Pattern: Database per service[ §](https://microservices.io/patterns/data/database-per-service.html#undefined)
[ pattern ](https://microservices.io/tags/pattern) [ application architecture ](https://microservices.io/tags/application architecture) [ loose coupling ](https://microservices.io/tags/loose coupling)
## Context[ §](https://microservices.io/patterns/data/database-per-service.html#context)
Let’s imagine you are developing an online store application using the [Microservice architecture pattern](https://microservices.io/patterns/microservices.html). Most services need to persist data in some kind of database. For example, the `Order Service` stores information about orders and the `Customer Service` stores information about customers.
![](https://microservices.io/i/customersandorders.png)
## Problem[ §](https://microservices.io/patterns/data/database-per-service.html#problem)
What’s the database architecture in a microservices application?
## Forces[ §](https://microservices.io/patterns/data/database-per-service.html#forces)
  * Services must be loosely coupled so that they can be developed, deployed and scaled independently
  * Some business transactions must enforce invariants that span multiple services. For example, the `Place Order` use case must verify that a new Order will not exceed the customer’s credit limit. Other business transactions, must update data owned by multiple services.
  * Some business transactions need to query data that is owned by multiple services. For example, the `View Available Credit` use must query the Customer to find the `creditLimit` and Orders to calculate the total amount of the open orders.
  * Some queries must join data that is owned by multiple services. For example, finding customers in a particular region and their recent orders requires a join between customers and orders.
  * Databases must sometimes be replicated and sharded in order to scale. See the [Scale Cube](https://microservices.io/articles/scalecube.html).
  * Different services have different data storage requirements. For some services, a relational database is the best choice. Other services might need a NoSQL database such as MongoDB, which is good at storing complex, unstructured data, or Neo4J, which is designed to efficiently store and query graph data.


## Solution[ §](https://microservices.io/patterns/data/database-per-service.html#solution)
Keep each microservice’s persistent data private to that service and accessible only via its API. A service’s transactions only involve its database.
The following diagram shows the structure of this pattern.
![](https://microservices.io/i/databaseperservice.png)
The service’s database is effectively part of the implementation of that service. It cannot be accessed directly by other services.
There are a few different ways to keep a service’s persistent data private. You do not need to provision a database server for each service. For example, if you are using a relational database then the options are:
  * Private-tables-per-service – each service owns a set of tables that must only be accessed by that service
  * Schema-per-service – each service has a database schema that’s private to that service
  * Database-server-per-service – each service has it’s own database server.


Private-tables-per-service and schema-per-service have the lowest overhead. Using a schema per service is appealing since it makes ownership clearer. Some high throughput services might need their own database server.
It is a good idea to create barriers that enforce this modularity. You could, for example, assign a different database user id to each service and use a database access control mechanism such as grants. Without some kind of barrier to enforce encapsulation, developers will always be tempted to bypass a service’s API and access it’s data directly.
## Example[ §](https://microservices.io/patterns/data/database-per-service.html#example)
The [FTGO application](https://microservices.io/post/microservices/patterns/data/2019/07/15/ftgo-database-per-service.html) is an example of an application that uses this approach. Each service has database credentials that only grant it access its own (logical) database on a shared MySQL server. For more information, see this [blog post](https://microservices.io/post/microservices/patterns/data/2019/07/15/ftgo-database-per-service.html).
## Resulting context[ §](https://microservices.io/patterns/data/database-per-service.html#resulting-context)
Using a database per service has the following benefits:
  * Helps ensure that the services are loosely coupled. Changes to one service’s database does not impact any other services.
  * Each service can use the type of database that is best suited to its needs. For example, a service that does text searches could use ElasticSearch. A service that manipulates a social graph could use Neo4j.


Using a database per service has the following drawbacks:
  * Implementing business transactions that span multiple services is not straightforward. Distributed transactions are best avoided because of the CAP theorem. Moreover, many modern (NoSQL) databases don’t support them.
  * Implementing queries that join data that is now in multiple databases is challenging.
  * Complexity of managing multiple SQL and NoSQL databases


There are various patterns/solutions for implementing transactions and queries that span services:
  * Implementing transactions that span services - use the [Saga pattern](https://microservices.io/patterns/data/saga.html).
  * Implementing queries that span services:
    * [API Composition](https://microservices.io/patterns/data/api-composition.html) - the application performs the join rather than the database. For example, a service (or the API gateway) could retrieve a customer and their orders by first retrieving the customer from the customer service and then querying the order service to return the customer’s most recent orders.
    * [Command Query Responsibility Segregation (CQRS)](https://microservices.io/patterns/data/cqrs.html) - maintain one or more materialized views that contain data from multiple services. The views are kept by services that subscribe to events that each services publishes when it updates its data. For example, the online store could implement a query that finds customers in a particular region and their recent orders by maintaining a view that joins customers and orders. The view is updated by a service that subscribes to customer and order events.


## Related patterns[ §](https://microservices.io/patterns/data/database-per-service.html#related-patterns)
  * [Microservice architecture pattern](https://microservices.io/patterns/microservices.html) creates the need for this pattern
  * [Saga pattern](https://microservices.io/patterns/data/saga.html) is a useful way to implement eventually consistent transactions
  * The [API Composition](https://microservices.io/patterns/data/api-composition.html) and [Command Query Responsibility Segregation (CQRS)](https://microservices.io/patterns/data/cqrs.html) pattern are useful ways to implement queries
  * The [Shared Database anti-pattern](https://microservices.io/patterns/data/shared-database.html) describes the problems that result from microservices sharing a database


[ pattern ](https://microservices.io/tags/pattern) [ application architecture ](https://microservices.io/tags/application architecture) [ loose coupling ](https://microservices.io/tags/loose coupling)
Copyright © 2025 Chris Richardson • All rights reserved • Supported by [Kong](https://konghq.com/).
#### About Microservices.io
![](https://gravatar.com/avatar/a290a8643359e2495e1c6312e662012f)
Microservices.io is brought to you by [Chris Richardson](https://microservices.io/about.html). Experienced software architect, author of POJOs in Action, the creator of the original [CloudFoundry.com](http://CloudFoundry.com), and the author of Microservices patterns. 
#### Microservices Patterns, 2nd edition
![](https://microservices.io/i/posts/mp2e-book-cover.png)
I am very excited to announce that the MEAP for the second edition of my book, Microservices Patterns is now available! 
[Learn more](https://microservices.io/post/architecture/2025/06/26/announcing-meap-microservices-patterns-2nd-edition.html)
#### ASK CHRIS
?
Got a question about microservices?
Fill in [this form](https://forms.gle/ppYDAF1JxHGec8Kn9). If I can, I'll write a blog post that answers your question.
#### NEED HELP?
![](https://microservices.io/i/posts/cxo-wondering.webp)
I help organizations improve agility and competitiveness through better software architecture. 
Learn more about my [consulting engagements](https://chrisrichardson.net/consulting.html), and [training workshops](https://chrisrichardson.net/training.html). 
#### PREMIUM CONTENT
![](https://microservices.io/i/posts/premium-logo.png) Premium content now available for paid subscribers at [premium.microservices.io](https://premium.microservices.io). 
#### MICROSERVICES WORKSHOPS
![](https://microservices.io/i/workshop-kata_small.jpg)
Chris teaches [comprehensive workshops](http://chrisrichardson.net/training.html) for architects and developers that will enable your organization use microservices effectively. 
Avoid the pitfalls of adopting microservices and learn essential topics, such as service decomposition and design and how to refactor a monolith to microservices. 
[Learn more](http://chrisrichardson.net/training.html)
#### Remote consulting session
![](https://microservices.io/i/posts/zoom-consulting.webp)
Got a specific microservice architecture-related question? For example:
  * Wondering whether your organization should adopt microservices?
  * Want to know how to migrate your monolith to microservices?
  * Facing a tricky microservice architecture design problem?


Consider signing up for a [two hour, highly focussed, consulting session.](https://chrisrichardson.net/consulting-office-hours.html)
#### ASSESS your architecture
Assess your application's microservice architecture and identify what needs to be improved. [Engage Chris](http://www.chrisrichardson.net/consulting.html) to conduct an architect review.
#### LEARN about microservices
Chris offers numerous other resources for learning the microservice architecture.
#### Get the book: Microservices Patterns
Read Chris Richardson's book: [ ![](https://microservices.io/i/Microservices-Patterns-Cover-published.png) ](https://microservices.io/book)
#### Example microservices applications
Want to see an example? Check out Chris Richardson's example applications. [See code](http://eventuate.io/exampleapps.html)
#### Virtual bootcamp: Distributed data patterns in a microservice architecture
![](https://microservices.io/i/Chris_Speaking_Mucon_2018_a.jpg)
My virtual bootcamp, distributed data patterns in a microservice architecture, is now open for enrollment!
It covers the key distributed data management patterns including Saga, API Composition, and CQRS.
It consists of video lectures, code labs, and a weekly ask-me-anything video conference repeated in multiple timezones.
The regular price is $395/person but use coupon DWTWBMJI to sign up for $95 (valid until September 10th, 2025). There are deeper discounts for buying multiple seats. 
[Learn more](https://chrisrichardson.net/virtual-bootcamp-distributed-data-management.html)
#### Learn how to create a service template and microservice chassis
Take a look at my [Manning LiveProject](https://microservices.io/post/patterns/2022/03/15/service-template-chassis-live-project.html) that teaches you how to develop a service template and microservice chassis. 
![](https://microservices.io/i/patterns/microservice-template-and-chassis/Microservice_chassis.png)
[Signup for the newsletter](http://visitor.r20.constantcontact.com/d.jsp?llr=ula8akwab&p=oi&m=1123470377332&sit=l6ktajjkb&f=15d9bba9-b33d-491f-b874-73a41bba8a76)
For Email Marketing you can trust.
#### BUILD microservices
Ready to start using the microservice architecture? 
#### Consulting services
[Engage Chris](http://www.chrisrichardson.net/consulting.html) to create a microservices adoption roadmap and help you define your microservice architecture, 
#### The Eventuate platform
Use the [Eventuate.io platform](https://eventuate.io) to tackle distributed data management challenges in your microservices architecture.
[![](https://eventuate.io/i/logo.gif)](https://eventuate.io)
Eventuate is Chris's latest startup. It makes it easy to use the Saga pattern to manage transactions and the CQRS pattern to implement queries.
Join the [microservices google group](https://groups.google.com/forum/#!forum/microservices)
BDOW!


## Source Information
- URL: https://microservices.io/patterns/data/database-per-service.html
- Harvested: 2025-08-19T23:12:29.749656
- Profile: deep_research
- Agent: architect
