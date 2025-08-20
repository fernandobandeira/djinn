---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-20T13:07:56.897251'
profile: deep_research
source: https://qdrant.tech/documentation
topic: Qdrant Vector Database Architecture
---

# Qdrant Vector Database Architecture

[![logo](https://qdrant.tech/img/qdrant-logo.svg)](https://qdrant.tech/)
  * [Qdrant](https://qdrant.tech/documentation/)
  * [Cloud](https://qdrant.tech/documentation/cloud-intro/)
  * [Build](https://qdrant.tech/documentation/build/)
  * [Learn](https://qdrant.tech/articles/)
  * [API Reference](https://api.qdrant.tech/api-reference)


Search
[Log in](https://cloud.qdrant.io/login?qdrant_tech_ajs_anonymous_id=9d43318b-727d-4b3c-9cb9-9269aa5a0d69&__hstc=265983056.fc170c4384295fd05418cdaddadc514e.1755706072271.1755706072271.1755706072271.1&__hssc=265983056.1.1755706072271&__hsfp=3894669469) [Start Free](https://cloud.qdrant.io/signup?qdrant_tech_ajs_anonymous_id=9d43318b-727d-4b3c-9cb9-9269aa5a0d69&__hstc=265983056.fc170c4384295fd05418cdaddadc514e.1755706072271.1755706072271.1755706072271.1&__hssc=265983056.1.1755706072271&__hsfp=3894669469)
![logo](https://qdrant.tech/img/qdrant-logo.svg)
Search
  * [ Qdrant](https://qdrant.tech/documentation/)
  * [ Cloud](https://qdrant.tech/documentation/cloud-intro/)
  * [ Build](https://qdrant.tech/documentation/build/)
  * [ Learn](https://qdrant.tech/articles/)
  * [ API Reference](https://api.qdrant.tech/api-reference)


### Getting Started
[What is Qdrant?](https://qdrant.tech/documentation/overview/)
  * [Understanding Vector Search in Qdrant](https://qdrant.tech/documentation/overview/vector-search/)


[Local Quickstart](https://qdrant.tech/documentation/quickstart/)
[API & SDKs](https://qdrant.tech/documentation/interfaces/)
[Qdrant Web UI](https://qdrant.tech/documentation/web-ui/)
### User Manual
[Concepts](https://qdrant.tech/documentation/concepts/)
  * [Collections](https://qdrant.tech/documentation/concepts/collections/)
  * [Points](https://qdrant.tech/documentation/concepts/points/)
  * [Vectors](https://qdrant.tech/documentation/concepts/vectors/)
  * [Payload](https://qdrant.tech/documentation/concepts/payload/)
  * [Search](https://qdrant.tech/documentation/concepts/search/)
  * [Explore](https://qdrant.tech/documentation/concepts/explore/)
  * [Hybrid Queries](https://qdrant.tech/documentation/concepts/hybrid-queries/)
  * [Filtering](https://qdrant.tech/documentation/concepts/filtering/)
  * [Optimizer](https://qdrant.tech/documentation/concepts/optimizer/)
  * [Storage](https://qdrant.tech/documentation/concepts/storage/)
  * [Indexing](https://qdrant.tech/documentation/concepts/indexing/)
  * [Snapshots](https://qdrant.tech/documentation/concepts/snapshots/)

[Guides](https://qdrant.tech/documentation/guides/installation/)
  * [Installation](https://qdrant.tech/documentation/guides/installation/)
  * [Administration](https://qdrant.tech/documentation/guides/administration/)
  * [Running with GPU](https://qdrant.tech/documentation/guides/running-with-gpu/)
  * [Capacity Planning](https://qdrant.tech/documentation/guides/capacity-planning/)
  * [Optimize Performance](https://qdrant.tech/documentation/guides/optimize/)
  * [Multitenancy](https://qdrant.tech/documentation/guides/multiple-partitions/)
  * [Distributed Deployment](https://qdrant.tech/documentation/guides/distributed_deployment/)
  * [Quantization](https://qdrant.tech/documentation/guides/quantization/)
  * [Monitoring & Telemetry](https://qdrant.tech/documentation/guides/monitoring/)
  * [Configuration](https://qdrant.tech/documentation/guides/configuration/)
  * [Security](https://qdrant.tech/documentation/guides/security/)
  * [Usage Statistics](https://qdrant.tech/documentation/guides/usage-statistics/)
  * [Troubleshooting](https://qdrant.tech/documentation/guides/common-errors/)


### Ecosystem
[FastEmbed](https://qdrant.tech/documentation/fastembed/)
  * [Quickstart](https://qdrant.tech/documentation/fastembed/fastembed-quickstart/)
  * [FastEmbed & Qdrant](https://qdrant.tech/documentation/fastembed/fastembed-semantic-search/)
  * [Working with miniCOIL](https://qdrant.tech/documentation/fastembed/fastembed-minicoil/)
  * [Working with SPLADE](https://qdrant.tech/documentation/fastembed/fastembed-splade/)
  * [Working with ColBERT](https://qdrant.tech/documentation/fastembed/fastembed-colbert/)
  * [Reranking with FastEmbed](https://qdrant.tech/documentation/fastembed/fastembed-rerankers/)


[Qdrant MCP Server](https://github.com/qdrant/mcp-server-qdrant)
### Tutorials
[Vector Search Basics](https://qdrant.tech/documentation/beginner-tutorials/)
  * [Semantic Search 101](https://qdrant.tech/documentation/beginner-tutorials/search-beginners/)
  * [Build a Neural Search Service](https://qdrant.tech/documentation/beginner-tutorials/neural-search/)
  * [Setup Hybrid Search with FastEmbed](https://qdrant.tech/documentation/beginner-tutorials/hybrid-search-fastembed/)
  * [Measure Search Quality](https://qdrant.tech/documentation/beginner-tutorials/retrieval-quality/)

[Advanced Retrieval](https://qdrant.tech/documentation/advanced-tutorials/)
  * [How to Use Multivector Representations with Qdrant Effectively](https://qdrant.tech/documentation/advanced-tutorials/using-multivector-representations/)
  * [Reranking in Hybrid Search](https://qdrant.tech/documentation/advanced-tutorials/reranking-hybrid-search/)
  * [Search Through Your Codebase](https://qdrant.tech/documentation/advanced-tutorials/code-search/)
  * [Build a Recommendation System with Collaborative Filtering](https://qdrant.tech/documentation/advanced-tutorials/collaborative-filtering/)
  * [Scaling PDF Retrieval with Qdrant](https://qdrant.tech/documentation/advanced-tutorials/pdf-retrieval-at-scale/)

[Using the Database](https://qdrant.tech/documentation/database-tutorials/)
  * [Bulk Upload Vectors](https://qdrant.tech/documentation/database-tutorials/bulk-upload/)
  * [Create & Restore Snapshots](https://qdrant.tech/documentation/database-tutorials/create-snapshot/)
  * [Large Scale Search](https://qdrant.tech/documentation/database-tutorials/large-scale-search/)
  * [Load a HuggingFace Dataset](https://qdrant.tech/documentation/database-tutorials/huggingface-datasets/)
  * [Build With Async API](https://qdrant.tech/documentation/database-tutorials/async-api/)
  * [Migration to Qdrant](https://qdrant.tech/documentation/database-tutorials/migration/)
  * [Static Embeddings. Should you pay attention?](https://qdrant.tech/documentation/database-tutorials/static-embeddings/)


### Support
[FAQ](https://qdrant.tech/documentation/faq/qdrant-fundamentals/)
  * [Qdrant Fundamentals](https://qdrant.tech/documentation/faq/qdrant-fundamentals/)
  * [Database Optimization](https://qdrant.tech/documentation/faq/database-optimization/)


[Release Notes](https://github.com/qdrant/qdrant/releases)
### Getting Started
[What is Qdrant?](https://qdrant.tech/documentation/overview/)
  * [Understanding Vector Search in Qdrant](https://qdrant.tech/documentation/overview/vector-search/)


[Local Quickstart](https://qdrant.tech/documentation/quickstart/)
[API & SDKs](https://qdrant.tech/documentation/interfaces/)
[Qdrant Web UI](https://qdrant.tech/documentation/web-ui/)
### User Manual
[Concepts](https://qdrant.tech/documentation/concepts/)
  * [Collections](https://qdrant.tech/documentation/concepts/collections/)
  * [Points](https://qdrant.tech/documentation/concepts/points/)
  * [Vectors](https://qdrant.tech/documentation/concepts/vectors/)
  * [Payload](https://qdrant.tech/documentation/concepts/payload/)
  * [Search](https://qdrant.tech/documentation/concepts/search/)
  * [Explore](https://qdrant.tech/documentation/concepts/explore/)
  * [Hybrid Queries](https://qdrant.tech/documentation/concepts/hybrid-queries/)
  * [Filtering](https://qdrant.tech/documentation/concepts/filtering/)
  * [Optimizer](https://qdrant.tech/documentation/concepts/optimizer/)
  * [Storage](https://qdrant.tech/documentation/concepts/storage/)
  * [Indexing](https://qdrant.tech/documentation/concepts/indexing/)
  * [Snapshots](https://qdrant.tech/documentation/concepts/snapshots/)

[Guides](https://qdrant.tech/documentation/guides/installation/)
  * [Installation](https://qdrant.tech/documentation/guides/installation/)
  * [Administration](https://qdrant.tech/documentation/guides/administration/)
  * [Running with GPU](https://qdrant.tech/documentation/guides/running-with-gpu/)
  * [Capacity Planning](https://qdrant.tech/documentation/guides/capacity-planning/)
  * [Optimize Performance](https://qdrant.tech/documentation/guides/optimize/)
  * [Multitenancy](https://qdrant.tech/documentation/guides/multiple-partitions/)
  * [Distributed Deployment](https://qdrant.tech/documentation/guides/distributed_deployment/)
  * [Quantization](https://qdrant.tech/documentation/guides/quantization/)
  * [Monitoring & Telemetry](https://qdrant.tech/documentation/guides/monitoring/)
  * [Configuration](https://qdrant.tech/documentation/guides/configuration/)
  * [Security](https://qdrant.tech/documentation/guides/security/)
  * [Usage Statistics](https://qdrant.tech/documentation/guides/usage-statistics/)
  * [Troubleshooting](https://qdrant.tech/documentation/guides/common-errors/)


### Ecosystem
[FastEmbed](https://qdrant.tech/documentation/fastembed/)
  * [Quickstart](https://qdrant.tech/documentation/fastembed/fastembed-quickstart/)
  * [FastEmbed & Qdrant](https://qdrant.tech/documentation/fastembed/fastembed-semantic-search/)
  * [Working with miniCOIL](https://qdrant.tech/documentation/fastembed/fastembed-minicoil/)
  * [Working with SPLADE](https://qdrant.tech/documentation/fastembed/fastembed-splade/)
  * [Working with ColBERT](https://qdrant.tech/documentation/fastembed/fastembed-colbert/)
  * [Reranking with FastEmbed](https://qdrant.tech/documentation/fastembed/fastembed-rerankers/)


[Qdrant MCP Server](https://github.com/qdrant/mcp-server-qdrant)
### Tutorials
[Vector Search Basics](https://qdrant.tech/documentation/beginner-tutorials/)
  * [Semantic Search 101](https://qdrant.tech/documentation/beginner-tutorials/search-beginners/)
  * [Build a Neural Search Service](https://qdrant.tech/documentation/beginner-tutorials/neural-search/)
  * [Setup Hybrid Search with FastEmbed](https://qdrant.tech/documentation/beginner-tutorials/hybrid-search-fastembed/)
  * [Measure Search Quality](https://qdrant.tech/documentation/beginner-tutorials/retrieval-quality/)

[Advanced Retrieval](https://qdrant.tech/documentation/advanced-tutorials/)
  * [How to Use Multivector Representations with Qdrant Effectively](https://qdrant.tech/documentation/advanced-tutorials/using-multivector-representations/)
  * [Reranking in Hybrid Search](https://qdrant.tech/documentation/advanced-tutorials/reranking-hybrid-search/)
  * [Search Through Your Codebase](https://qdrant.tech/documentation/advanced-tutorials/code-search/)
  * [Build a Recommendation System with Collaborative Filtering](https://qdrant.tech/documentation/advanced-tutorials/collaborative-filtering/)
  * [Scaling PDF Retrieval with Qdrant](https://qdrant.tech/documentation/advanced-tutorials/pdf-retrieval-at-scale/)

[Using the Database](https://qdrant.tech/documentation/database-tutorials/)
  * [Bulk Upload Vectors](https://qdrant.tech/documentation/database-tutorials/bulk-upload/)
  * [Create & Restore Snapshots](https://qdrant.tech/documentation/database-tutorials/create-snapshot/)
  * [Large Scale Search](https://qdrant.tech/documentation/database-tutorials/large-scale-search/)
  * [Load a HuggingFace Dataset](https://qdrant.tech/documentation/database-tutorials/huggingface-datasets/)
  * [Build With Async API](https://qdrant.tech/documentation/database-tutorials/async-api/)
  * [Migration to Qdrant](https://qdrant.tech/documentation/database-tutorials/migration/)
  * [Static Embeddings. Should you pay attention?](https://qdrant.tech/documentation/database-tutorials/static-embeddings/)


### Support
[FAQ](https://qdrant.tech/documentation/faq/qdrant-fundamentals/)
  * [Qdrant Fundamentals](https://qdrant.tech/documentation/faq/qdrant-fundamentals/)
  * [Database Optimization](https://qdrant.tech/documentation/faq/database-optimization/)


[Release Notes](https://github.com/qdrant/qdrant/releases)
# Qdrant Documentation
Qdrant is an AI-native vector database and a semantic search engine. You can use it to extract meaningful information from unstructured data.
[Clone this repo now](https://github.com/qdrant/qdrant_demo/) and build a search engine in five minutes.
[Cloud Quickstart ](https://qdrant.tech/documentation/quickstart-cloud/)[Local Quickstart](https://qdrant.tech/documentation/quickstart/)
## Ready to start developing?
Qdrant is open-source and can be self-hosted. However, the quickest way to get started is with our [free tier](https://qdrant.to/cloud) on Qdrant Cloud. It scales easily and provides a UI where you can interact with data.
### Create your first Qdrant Cloud cluster today
[Get Started](https://qdrant.to/cloud)
![](https://qdrant.tech/img/rocket.svg)
## Optimize Qdrant's performance
Boost search speed, reduce latency, and improve the accuracy and memory usage of your Qdrant deployment.
[Learn More](https://qdrant.tech/documentation/guides/optimize/)
[![Documents](https://qdrant.tech/icons/outline/documentation-blue.svg) DocumentsDistributed DeploymentScale Qdrant beyond a single node and optimize for high availability, fault tolerance, and billion-scale performance.Read More](https://qdrant.tech/documentation/guides/distributed_deployment/)
[![Documents](https://qdrant.tech/icons/outline/documentation-blue.svg) DocumentsMultitenancyBuild vector search apps that serve millions of users. Learn about data isolation, security, and performance tuning.Read More](https://qdrant.tech/documentation/guides/multiple-partitions/)
[![Blog](https://qdrant.tech/icons/outline/blog-purple.svg) BlogVector QuantizationLearn about cutting-edge techniques for vector quantization and how they can be used to improve search performance.Read More](https://qdrant.tech/articles/what-is-vector-quantization/)
#### Ready to get started with Qdrant?
[Start Free](https://qdrant.to/cloud/)
© 2025 Qdrant.
[Terms](https://qdrant.tech/legal/terms_and_conditions/) [Privacy Policy](https://qdrant.tech/legal/privacy-policy/) [Impressum](https://qdrant.tech/legal/impressum/)
×
[ Powered by ](https://qdrant.tech/)
## About cookies on this site
We use cookies to collect and analyze information on site performance and usage, to provide social media features, and to enhance and customize content and advertisements. [Learn more](https://qdrant.tech/legal/privacy-policy/#cookies-and-web-beacons)
Cookies Settings Reject All Accept All Cookies
![Company Logo](https://cdn.cookielaw.org/logos/static/ot_company_logo.png)
## Privacy Preference Center
Cookies used on the site are categorized, and below, you can read about each category and allow or deny some or all of them. When categories that have been previously allowed are disabled, all cookies assigned to that category will be removed from your browser. Additionally, you can see a list of cookies assigned to each category and detailed information in the cookie declaration. [More information](https://qdrant.tech/legal/privacy-policy/#cookies-and-web-beacons)
Allow All
###  Manage Consent Preferences
#### Targeting Cookies
Targeting Cookies
These cookies may be set through our site by our advertising partners. They may be used by those companies to build a profile of your interests and show you relevant adverts on other sites. They do not store directly personal information, but are based on uniquely identifying your browser and internet device. If you do not allow these cookies, you will experience less targeted advertising.
#### Functional Cookies
Functional Cookies
These cookies enable the website to provide enhanced functionality and personalisation. They may be set by us or by third party providers whose services we have added to our pages. If you do not allow these cookies then some or all of these services may not function properly.
#### Strictly Necessary Cookies
Always Active
These cookies are necessary for the website to function and cannot be switched off in our systems. They are usually only set in response to actions made by you which amount to a request for services, such as setting your privacy preferences, logging in or filling in forms. You can set your browser to block or alert you about these cookies, but some parts of the site will not then work. These cookies do not store any personally identifiable information.
#### Performance Cookies
Performance Cookies
These cookies allow us to count visits and traffic sources so we can measure and improve the performance of our site. They help us to know which pages are the most and least popular and see how visitors move around the site. All information these cookies collect is aggregated and therefore anonymous. If you do not allow these cookies we will not know when you have visited our site, and will not be able to monitor its performance.
Back Button
### Cookie List
Search Icon
Filter Icon
Clear
checkbox label label
Apply Cancel
Consent Leg.Interest
checkbox label label
checkbox label label
checkbox label label
Reject All Confirm My Choices
[![Powered by Onetrust](https://cdn.cookielaw.org/logos/static/powered_by_logo.svg)](https://www.onetrust.com/products/cookie-consent/)


## Source Information
- URL: https://qdrant.tech/documentation
- Harvested: 2025-08-20T13:07:56.897251
- Profile: deep_research
- Agent: architect
