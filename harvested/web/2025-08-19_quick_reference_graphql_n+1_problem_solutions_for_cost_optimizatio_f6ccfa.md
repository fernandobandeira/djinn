---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T14:13:29.302554'
profile: quick_reference
source: https://shopify.engineering/solving-the-n-1-problem-for-graphql-through-batching
topic: GraphQL N+1 Problem Solutions for Cost Optimization
---

# GraphQL N+1 Problem Solutions for Cost Optimization

[Skip to Content](https://shopify.engineering/solving-the-n-1-problem-for-graphql-through-batching#main)
[![Shopify](https://cdn.shopify.com/b/shopify-brochure2-assets/d9340911ca8c679b148dd4a205ad2ffa.svg)](https://shopify.engineering/)
[![Shopify](https://cdn.shopify.com/b/shopify-brochure2-assets/cac815e4ee0f383f7b4b5302b5a7a29a.svg)](https://shopify.engineering/)Engineering BlogMenu
Search
Type something you're looking for
  * [Overview](https://shopify.engineering/)
  * [Development](https://shopify.engineering/topics/development)
  * [Infrastructure](https://shopify.engineering/topics/infrastructure)
  * [Mobile](https://shopify.engineering/topics/mobile)
  * [Developer Tooling](https://shopify.engineering/topics/developer-tooling)
  * [Latest](https://shopify.engineering/latest)
  * More topics
    * [Security](https://shopify.engineering/topics/security)
    * [Data Science Engineering](https://shopify.engineering/topics/data-science-engineering)
    * [Culture](https://shopify.engineering/topics/culture)


[![Shopify](https://cdn.shopify.com/b/shopify-brochure2-assets/d9340911ca8c679b148dd4a205ad2ffa.svg)](https://shopify.engineering/)
  * Solutions
Start
    * [Start your business.Build your brand](https://www.shopify.com/start)
    * [Create your website.Online store editor](https://www.shopify.com/website/builder)
    * [Customize your store.Store themes](https://themes.shopify.com/)
    * [Find business apps.Shopify app store](https://apps.shopify.com/)
    * [Own your site domain.Domains & hosting](https://www.shopify.com/domains)
    * [Explore free business tools.Tools to run your business](https://www.shopify.com/tools)
Sell
    * [Sell your products.Sell online or in person](https://www.shopify.com/sell)
    * [Check out customers.World-class checkout](https://www.shopify.com/checkout)
    * [Sell online.Grow your business online](https://www.shopify.com/online)
    * [Sell across channels.Reach millions of shoppers and boost sales](https://www.shopify.com/channels)
    * [Sell globally.International sales](https://www.shopify.com/international)
    * [Sell wholesale & direct.Business-to-business (B2B)](https://www.shopify.com/plus/solutions/b2b-ecommerce)
Market
    * [Market your business.Reach & retain customers](https://www.shopify.com/market)
    * [Market across social.Social media integrations](https://www.shopify.com/facebook-instagram)
    * [Chat with customers.Shopify Inbox](https://www.shopify.com/inbox)
    * [Nurture customers.Shopify Email](https://www.shopify.com/email-marketing)
    * [Know your audience.Gain customer insights](https://www.shopify.com/segmentation)
Manage
    * [Manage your business.Track sales, orders & analytics](https://www.shopify.com/manage)
    * [Measure your performance.Analytics and Reporting](https://www.shopify.com/analytics)
    * [Manage your stock & orders.Inventory & order management](https://www.shopify.com/orders)
    * [Automate your business.Shopify Flow](https://www.shopify.com/flow)
    * [Shopify Developers.Build with Shopify's powerful APIs](https://shopify.dev)
    * [Plus.A commerce solution for growing digital brands](https://www.shopify.com/plus)
    * [All Products.Explore all Shopify products & features](https://www.shopify.com/products)
  * [Pricing](https://www.shopify.com/pricing)
  * Resources
Help and support
    * [Help and support.Get 24/7 support](https://help.shopify.com/en/)
    * [Business courses.Learn from proven experts](https://www.shopifyacademy.com)
Popular topics
    * [What is Shopify?.How our commerce platform works](https://www.shopify.com/blog/what-is-shopify)
Essential tools
    * [Business name generator .](https://www.shopify.com/tools/business-name-generator)
    * [Logo maker.](https://www.shopify.com/tools/logo-maker)
    * [Stock photography.](https://www.shopify.com/stock-photos)
    * [QR code generator.](https://www.shopify.com/tools/qr-code-generator)
  * What’s new
    * [Changelog.Your source for recent updates](https://changelog.shopify.com)
    * [Newsroom.All company news and press releases](https://www.shopify.com/news)


[blog](https://shopify.engineering/)|[Development](https://shopify.engineering/topics/development)
# Solving the N+1 Problem for GraphQL through Batching
Published on Apr 24, 2018
![](https://cdn.shopify.com/s/files/1/0779/4361/articles/graphql.png?v=1524586742&originalWidth=1848&originalHeight=782)
Engineering at Shopify
We’re hiring
[See open roles](https://www.shopify.com/careers#Engineering)
**Authors: Leanne Shapton, Dylan Thacker-Smith, & Scott Walkinshaw**
When Shopify merchants build their businesses on our platform, they trust that we’ll provide them with a seamless experience. A huge part of that is creating scalable back-end solutions that allow us to manage the millions of requests reaching our servers each day.
When a storefront app makes a request to our servers, they’re interacting with the [Storefront API](https://help.shopify.com/api/storefront-api "Storefront API"). Historically, REST is the language of choice when designing APIs, but [Shopify uses GraphQL](https://help.shopify.com/api/storefront-api/graphql "Shopify GraphQL").
GraphQL is an increasingly popular query language in the developer community, because it avoids the classic [over-fetching](https://www.howtographql.com/basics/1-graphql-is-the-better-rest/ "Overfetching") problem associated with REST. In REST, the endpoint determines the type and amount of data returned. GraphQL, however, permits highly specific client-side queries that return only the data requested.
Over-fetching occurs when the server returns more data than needed. REST is especially prone to it, due to its endpoint design. Conversely, if a particular endpoint does not yield enough information (under-fetching), clients need to make additional queries to reach nested data. Both over-fetching and under-fetching waste valuable computing power and bandwidth.
In this REST example, the client requests all ‘authors’, and receives a response, including fields for name, id, number of publications, and country. The client may not have originally wanted all that information; the server has over-fetched the data.
**REST Query and Response**
This file contains hidden or bidirectional Unicode text that may be interpreted or compiled differently than what appears below. To review, open the file in an editor that reveals hidden Unicode characters. [Learn more about bidirectional Unicode characters](https://github.co/hiddenchars)
[ Show hidden characters ](https://shopify.engineering/{{ revealButtonHref }})
|  GET /authors.json  
---|---  
|   
|  {  
|  "authors": [{  
|  "name": "Adam",  
|  "id": 1,  
|  "number_of_publications": 100,  
|  "country": "Canada",  
|  }, {  
|  "name": "Jane",  
|  "id": 2,  
|  "number_of_publications": 56,  
|  "country": "Canada",  
|  ...  
|  }]  
|  }   
[view raw](https://gist.github.com/ShopifyEng/eabb3f55449479978b9134287176ae70/raw/12480d4243a1c8b49adf937c5804be98f15bfdf3/GraphQL_Batch_REST_query) [ GraphQL_Batch_REST_query ](https://gist.github.com/ShopifyEng/eabb3f55449479978b9134287176ae70#file-graphql_batch_rest_query) hosted with ❤ by [GitHub](https://github.com)
Conversely, in this GraphQL version, the client makes a query specifically for all authors’ names, and receives that only that information in the response.
**GraphQL Query**
This file contains hidden or bidirectional Unicode text that may be interpreted or compiled differently than what appears below. To review, open the file in an editor that reveals hidden Unicode characters. [Learn more about bidirectional Unicode characters](https://github.co/hiddenchars)
[ Show hidden characters ](https://shopify.engineering/{{ revealButtonHref }})
|  query {  
---|---  
|  authors {  
|  name  
|  }  
|  }  
[view raw](https://gist.github.com/ShopifyEng/2e1a0f42d50481e14d13b8e94f3ec2ff/raw/515b7f71496e184c7b0e7b4de5313e60115f3da6/GraphQL_Batch_GraphQL_query) [ GraphQL_Batch_GraphQL_query ](https://gist.github.com/ShopifyEng/2e1a0f42d50481e14d13b8e94f3ec2ff#file-graphql_batch_graphql_query) hosted with ❤ by [GitHub](https://github.com)
**GraphQL Response**
This file contains hidden or bidirectional Unicode text that may be interpreted or compiled differently than what appears below. To review, open the file in an editor that reveals hidden Unicode characters. [Learn more about bidirectional Unicode characters](https://github.co/hiddenchars)
[ Show hidden characters ](https://shopify.engineering/{{ revealButtonHref }})
|  {  
---|---  
|  "authors": [{  
|  "name": "Adam"  
|  }, {  
|  "name": "Jane"  
|  }]  
|  }  
[view raw](https://gist.github.com/ShopifyEng/995fde822c37796706b15c5a9d10e84a/raw/94a386555ad07150faad6233ad77bfeaa4085363/GraphQL_Batch_GraphQL_response) [ GraphQL_Batch_GraphQL_response ](https://gist.github.com/ShopifyEng/995fde822c37796706b15c5a9d10e84a#file-graphql_batch_graphql_response) hosted with ❤ by [GitHub](https://github.com)
GraphQL queries are made to a single endpoint, as opposed to multiple endpoints in REST. Because of this, clients need to know how to structure their requests to reach the data, rather than simply targeting endpoints. GraphQL back-end developers share this information by creating schemas. Schemas are like maps; they describe all the data and their relationships within a server.
A schema for the above example might look as follows.
This file contains hidden or bidirectional Unicode text that may be interpreted or compiled differently than what appears below. To review, open the file in an editor that reveals hidden Unicode characters. [Learn more about bidirectional Unicode characters](https://github.co/hiddenchars)
[ Show hidden characters ](https://shopify.engineering/{{ revealButtonHref }})
|  type author {  
---|---  
|  name: String!  
|  id: ID!  
|  }  
[view raw](https://gist.github.com/ShopifyEng/73123a9f866e9acc3d001fc55cb75292/raw/42183eed35ed3b0c324a88302eed6a82661a0064/GraphQL_Batch_schema) [ GraphQL_Batch_schema ](https://gist.github.com/ShopifyEng/73123a9f866e9acc3d001fc55cb75292#file-graphql_batch_schema) hosted with ❤ by [GitHub](https://github.com)
The schema defines the type ‘author’, for which two fields of information are available; name and id. The schema indicates that for each author, there’s a non-nullable string value for the ‘name’ field, and a unique, non-nullable identifier for the ‘id’ field. For more information, [visit the schema section on the official GraphQL website](https://graphql.org/learn/schema/ "GraphQL schema").
How does GraphQL return data to those fields? It uses resolvers. A resolver is a field-specific function that hunts for the requested data in the server. The server processes the query and the resolvers return data for each field, until it has fetched all the data in the query. Data is returned in the same format and order as the query, in a JSON file.
GraphQL’s major benefits are its straightforwardness and ease of use. Its solved our biggest problems by reducing the bandwidth used and latency while retrieving data for our apps.
As great as GraphQL is, it’s prone to encountering an issue, known as the n+1 problem. The n+1 problem arises because GraphQL executes a separate resolver function for every field, whereas REST has one resolver per endpoint. These additional resolvers mean that GraphQL runs the risk of making additional round trips to the database than are necessary for requests.
This file contains hidden or bidirectional Unicode text that may be interpreted or compiled differently than what appears below. To review, open the file in an editor that reveals hidden Unicode characters. [Learn more about bidirectional Unicode characters](https://github.co/hiddenchars)
[ Show hidden characters ](https://shopify.engineering/{{ revealButtonHref }})
|  query {   
---|---  
|  authors { # fetches authors (1 query)  
|  name   
|  address { # fetches address for each author (N queries for N authors)  
|  country  
|  }  
|  }  
|  } # Therefore = N+1 round trips  
[view raw](https://gist.github.com/ShopifyEng/4864e3f1d811aeb47bdbfcbc8c9fef18/raw/5e957acd7f02caa2575dbece01525eb72f1901a4/GraphQL_Batch_N+1_example) [ GraphQL_Batch_N+1_example ](https://gist.github.com/ShopifyEng/4864e3f1d811aeb47bdbfcbc8c9fef18#file-graphql_batch_n-1_example) hosted with ❤ by [GitHub](https://github.com)
The n+1 problem means that the server executes multiple unnecessary round trips to datastores for nested data. In the above case, the server makes 1 round trip to a datastore to fetch the authors, then makes N round trips to a datastore to fetch the address for N authors. For example, if there were fifty authors, then it would make fifty-one round trips for all the data. It should be able to fetch all the addresses together in a single round trip, so only two round trips to datastores in total, regardless of the number of authors. The computing expenditure of these extra round trips are massive when applied to large requests, like asking for fifty different colours of fifty t-shirts.
The n+1 problem is further exacerbated in GraphQL, because neither clients nor servers can predict how expensive a request is until _after_ it’s executed. In REST, costs are predictable because there’s one trip per endpoint requested. In GraphQL, there’s only one endpoint, and it’s not indicative of the potential size of incoming requests. At Shopify, where thousands of merchants interact with the Storefront API each day, we needed a solution that allowed us to minimize the cost of each request.
Facebook previously introduced a solution to the N+1 issue by creating [DataLoader](https://github.com/facebook/dataloader "Dataloader"), a library that batches requests specifically for JavaScript. Dylan Thacker-Smith, a developer at Shopify, used DataLoader as inspiration and built the [GraphQL Batch](https://github.com/Shopify/graphql-batch "GraphQL Batch Ruby Library") Ruby library specifically for the [GraphQL](http://graphql-ruby.org/ "GraphQL Ruby Library") Ruby library. This library reduces the overall number of datastore queries required when fulfilling requests with the GraphQL Ruby library. Instead of the server expecting each field resolver to return a value, the library allows the resolver to request data and return a [promise](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise "GraphQL Promise") for that data. For GraphQL, a promise represents the eventual, rather than immediate, resolution of a field. Therefore, instead of resolver functions executing immediately, the server waits before returning the data.
GraphQL Batch allows applications to define batch loaders that specify how to group and load similar data. The field resolvers can use one of the loaders to load data, which is grouped with similar loads, and returns a promise for the result. The GraphQL request executes by first trying to resolve all the fields, which may be resolved with promises. GraphQL Batch iterates through the grouped loads, uses their corresponding batch loader to load all the promises together, and replaces the promises with the loaded result. When an object field loads, fields nested on those objects resolve using their field resolvers (which may themselves use batch loaders), and then they’re grouped with similar loads that haven't executed. The benefits for Shopify are huge, as it massively reduces the amount of computing power required to process the same requests.
GraphQL Batch is now considered general best-practice for all GraphQL work at Shopify. We believe great tools should be shared with peers. The GraphQL Batch library is simple, but solves a major complaint within the GraphQL Ruby community. We believe the tool is flexible and has the potential to solve problems beyond just Shopify’s scope. As such, we chose to make [GraphQL Batch open-source](https://github.com/Shopify/graphql-batch "GraphQL Batch").
Many Shopify developers are already active individual GraphQL contributors, but Shopify is still constantly exploring ways to interact more meaningfully with the vibrant GraphQL developer community. Sharing the source code for GraphQL Batch is just a first step. As GraphQL adoption increases, we look forward to sharing our learnings and collaborating externally to build tools that improve the GraphQL developing experience.
# Learn More About GraphQL at Shopify
  * [Using GraphQL for High-Performing Mobile Applications](https://shopify.engineering/using-graphql-for-high-performing-mobile-applications "Using GraphQL for High-Performing Mobile Applications")
  * [Unifying Our GraphQL Design Patterns and Best Practices with Tutorials](https://shopify.engineering/unifying-graphql-design-patterns-best-practices-tutorials "Unifying Our GraphQL Design Patterns and Best Practices with Tutorials")
  * [Building Resilient GraphQL APIs Using Idempotency](https://shopify.engineering/building-resilient-graphql-apis-using-idempotency "Building Resilient GraphQL APIs Using Idempotency")


LS
by [Leanne Shapton](https://shopify.engineering/authors/leanne-shapton)
Published on Apr 24, 2018
Share article
  * [Facebook](https://www.facebook.com/sharer/sharer.php?u=https%3A%2F%2Fshopify.engineering%2Fsolving-the-n-1-problem-for-graphql-through-batching)
  * [Twitter](https://twitter.com/intent/tweet?text=Solving+the+N+1+Problem+for+GraphQL+through+Batching&url=https%3A%2F%2Fshopify.engineering%2Fsolving-the-n-1-problem-for-graphql-through-batching&via=Shopify)
  * [LinkedIn](https://www.linkedin.com/shareArticle?mini=true&source=Shopify&title=Solving+the+N+1+Problem+for+GraphQL+through+Batching&url=https%3A%2F%2Fshopify.engineering%2Fsolving-the-n-1-problem-for-graphql-through-batching)


by [Leanne Shapton](https://shopify.engineering/authors/leanne-shapton)
Published on Apr 24, 2018
• 6 minute read
[Development](https://shopify.engineering/topics/development)[Introducing Ruvy](https://shopify.engineering/introducing-ruvy)[Developer Tooling](https://shopify.engineering/topics/developer-tooling)[Building a ShopifyQL Code Editor](https://shopify.engineering/building-a-shopifyql-code-editor)
[Apps](https://shopify.engineering/topics/apps)[Shopify’s platform is the Web platform](https://shopify.engineering/shopifys-platform-is-the-web-platform)[Development](https://shopify.engineering/topics/development)[The Engineering Story Behind Flex Comp](https://shopify.engineering/building-flex-comp)
[![Introducing Ruvy](https://cdn.shopify.com/s/files/1/0779/4361/articles/ShopifyEng_BlogIllustrations_211008_72ppi_01_YJIT-BuildingaNewJITCompilerforCRuby_0a00a8cf-3951-4556-9b2c-3f88e9de76ff.jpg?v=1697134143&width=674&originalWidth=674&originalHeight=287)](https://shopify.engineering/introducing-ruvy)
[Development](https://shopify.engineering/topics/development)
[Introducing Ruvy](https://shopify.engineering/introducing-ruvy)
2023-10-18
[![Building a ShopifyQL Code Editor](https://cdn.shopify.com/s/files/1/0779/4361/articles/ShopifyEng_BlogIllustrations_210719_72ppi_05_FiveStepsforBuildingMachineLearning_886f2915-8848-4dfe-9278-f2575912073a.jpg?v=1694216014&width=674&originalWidth=674&originalHeight=287)](https://shopify.engineering/building-a-shopifyql-code-editor)
[Developer Tooling](https://shopify.engineering/topics/developer-tooling)
[Building a ShopifyQL Code Editor](https://shopify.engineering/building-a-shopifyql-code-editor)
2023-09-11
[![Shopify’s platform is the Web platform](https://cdn.shopify.com/s/files/1/0779/4361/articles/remix-app-bridge.png?v=1690230786&width=674&originalWidth=674&originalHeight=287)](https://shopify.engineering/shopifys-platform-is-the-web-platform)
[Apps](https://shopify.engineering/topics/apps)
[Shopify’s platform is the Web platform](https://shopify.engineering/shopifys-platform-is-the-web-platform)
2023-07-26
[![The Engineering Story Behind Flex Comp](https://cdn.shopify.com/s/files/1/0779/4361/articles/ShopifyEng_BlogIllustrations_220927_72ppi_03_BuildingShopifysNewFlexCompSystem.jpg?v=1666044315&width=674&originalWidth=674&originalHeight=287)](https://shopify.engineering/building-flex-comp)
[Development](https://shopify.engineering/topics/development)
[The Engineering Story Behind Flex Comp](https://shopify.engineering/building-flex-comp)
2022-10-05
Work from anywhere with Shopify
See our open roles and learn more about our digital by design culture.
[See open roles](https://www.shopify.com/careers#Engineering)
Opens an external site in a new window
![Shopify](https://cdn.shopify.com/b/shopify-brochure2-assets/88ee7022e2749387148cb4098cc4f9fb.svg)
### Shopify
  * [About](https://www.shopify.com/about)
  * [Careers](https://www.shopify.com/careers)
  * [Investors](https://shopifyinvestors.com/home/default.aspx)
  * [Press and Media](https://www.shopify.com/news)
  * [Partners](https://www.shopify.com/partners)
  * [Affiliates](https://www.shopify.com/affiliates)
  * [Legal](https://www.shopify.com/legal)
  * [Service status](https://www.shopifystatus.com/)


### Support
  * [Merchant Support](https://help.shopify.com/en/questions)
  * [Shopify Help Center](https://help.shopify.com/en/)
  * [Hire a Partner](https://www.shopify.com/partners/directory)
  * [Shopify Academy](https://www.shopifyacademy.com/?itcat=brochure&itterm=global-footer)
  * [Shopify Community](https://community.shopify.com/?utm_campaign=footer&utm_content=en&utm_medium=web&utm_source=shopify)


### Developers
  * [Shopify.dev](https://shopify.dev)
  * [API Documentation](https://shopify.dev/api)
  * [Dev Degree](https://devdegree.ca)


### Products
  * [Shop](https://shop.app)
  * [Shop Pay](https://www.shopify.com/shop-pay)
  * [Shopify for Enterprise](https://www.shopify.com/enterprise)


### Global Impact
  * [Sustainability](https://www.shopify.com/climate)
  * [Build Black](https://operationhope.org/initiatives/1-million-black-businesses/)
  * [Accessibility](https://www.shopify.com/accessibility)
  * [Research](https://www.shopify.com/plus/commerce-trends)


### Solutions
  * [Online Store Builder](https://www.shopify.com/online)
  * [Website Builder](https://www.shopify.com/website/builder)
  * [Ecommerce Website](https://www.shopify.com/tour/ecommerce-website)


  * [Terms of service](https://www.shopify.com/legal/terms)
  * [Privacy policy](https://www.shopify.com/legal/privacy)
  * [Sitemap](https://www.shopify.com/sitemap)
  * [Privacy Choices](https://privacy.shopify.com/en)


  * [](https://x.com/shopifyeng)




## Source Information
- URL: https://shopify.engineering/solving-the-n-1-problem-for-graphql-through-batching
- Harvested: 2025-08-19T14:13:29.302554
- Profile: quick_reference
- Agent: architect
