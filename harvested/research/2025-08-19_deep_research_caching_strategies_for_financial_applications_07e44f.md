---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T23:13:47.349994'
profile: deep_research
source: https://aws.amazon.com/builders-library/caching-challenges-and-strategies/
topic: Caching Strategies for Financial Applications
---

# Caching Strategies for Financial Applications

## Select your cookie preferences
We use essential cookies and similar tools that are necessary to provide our site and services. We use performance cookies to collect anonymous statistics, so we can understand how customers use our site and make improvements. Essential cookies cannot be deactivated, but you can choose “Customize” or “Decline” to decline performance cookies. If you agree, AWS and approved third parties will also use cookies to provide useful site features, remember your preferences, and display relevant content, including relevant advertising. To accept or decline all non-essential cookies, choose “Accept” or “Decline.” To make more detailed choices, choose “Customize.”
AcceptDeclineCustomize
## Customize cookie preferences
We use cookies and similar tools (collectively, "cookies") for the following purposes.
### Essential
Essential cookies are necessary to provide our site and services and cannot be deactivated. They are usually set in response to your actions on the site, such as setting your privacy preferences, signing in, or filling in forms. 
### Performance
Performance cookies provide anonymous statistics about how customers navigate our site so we can improve site experience and performance. Approved third parties may perform analytics on our behalf, but they cannot use the data for their own purposes.
Allowed
### Functional
Functional cookies help us provide useful site features, remember your preferences, and display relevant content. Approved third parties may set these cookies to provide certain site features. If you do not allow these cookies, then some or all of these services may not function properly.
Allowed
### Advertising
Advertising cookies may be set through our site by us or our advertising partners and help us deliver relevant marketing content. If you do not allow these cookies, you will experience less relevant advertising.
Allowed
Blocking some types of cookies may impact your experience of our sites. You may review and change your choices at any time by selecting Cookie preferences in the footer of this site. We and selected third-parties use cookies or similar technologies as specified in the [AWS Cookie Notice](https://aws.amazon.com/legal/cookies/).
CancelSave preferences
## Your privacy choices
We and our advertising partners (“we”) may use information we collect from or about you to show you ads on other websites and online services. Under certain laws, this activity is referred to as “cross-context behavioral advertising” or “targeted advertising.”
To opt out of our use of cookies or similar technologies to engage in these activities, select “Opt out of cross-context behavioral ads” and “Save preferences” below. If you clear your browser cookies or visit this site from a different device or browser, you will need to make your selection again. For more information about cookies and how we use them, read our [Cookie Notice](https://aws.amazon.com/legal/cookies/).
Allow cross-context behavioral adsOpt out of cross-context behavioral ads
To opt out of the use of other identifiers, such as contact information, for these activities, fill out the form [here](https://pulse.aws/application/ZRPLWLL6?p=0).
For more information about how AWS handles your information, read the [AWS Privacy Notice](https://aws.amazon.com/privacy/).
CancelSave preferences
## Unable to save cookie preferences
We will only store essential cookies at this time, because we were unable to save your cookie preferences.If you want to change your cookie preferences, try again later using the link in the AWS console footer, or contact support if the problem persists.
Dismiss
[Skip to main content](https://aws.amazon.com/builders-library/caching-challenges-and-strategies/#aws-page-content-main)
  * English
  * [Contact us](https://aws.amazon.com/contact-us/?nc2=h_ut_cu)
  * Support 
  * My account 


  * [](https://aws.amazon.com/?nc2=h_home)
  * [Agentic AI](https://aws.amazon.com/ai/agentic-ai/?nc2=h_l1_f)
  * Discover AWS
  * Products
  * Solutions
  * Pricing
  * More 


  * Filter: All
  * Sign in to console
  * [Create account](https://portal.aws.amazon.com/gp/aws/developer/registration/index.html?nc2=h_su&src=header_signup)


The Amazon Builders' Library
  * [Overview](https://aws.amazon.com/builders-library/)
  * [Authors](https://aws.amazon.com/builders-library/authors/)
  * [FAQs](https://aws.amazon.com/builders-library/faqs/)


  * [Developer Center](https://aws.amazon.com/developer/)›
  * [Amazon Builders' Library](https://aws.amazon.com/builders-library/)›
  * [Caching challenges and strategies](https://aws.amazon.com/builders-library/caching-challenges-and-strategies/)


# Caching challenges and strategies
Would you like to be notified of new content?
[Send me updates](https://pages.awscloud.com/amazon-builders-library.html)
## The ecstasy and the agony of caches
Over years of building services at Amazon we’ve experienced various versions of the following scenario: We build a new service, and this service needs to make some network calls to fulfill its requests. Perhaps these calls are to a relational database, or an AWS service like Amazon DynamoDB, or to another internal service. In simple tests or at low request rates the service works great, but we notice a problem on the horizon. The problem might be that calls to this other service are slow or that the database is expensive to scale out as call volume increases. We also notice that many requests are using the same downstream resource or the same query results, so we think that caching this data could be the answer to our problems. We add a cache and our service appears much improved. We observe that request latency is down, costs are reduced, and small downstream availability drops are smoothed over. After a while, no one can remember life before the cache. Dependencies reduce their fleet sizes accordingly, and the database is scaled down. Just when everything appears to be going well, the service could be poised for disaster. There could be changes in traffic patterns, failure of the cache fleet, or other unexpected circumstances that could lead to a cold or otherwise unavailable cache. This in turn could cause a surge of traffic to downstream services that can lead to outages both in our dependencies and in our service.
We’ve just described a service that has become addicted to its cache. The cache has been inadvertently elevated from a helpful addition to the service to a necessary and critical part of its ability to operate. At the heart of this issue is the modal behavior introduced by the cache, with differing behavior depending on whether a given object is cached. An unanticipated shift in the distribution of this modal behavior can potentially lead to disaster.
We have experienced both the benefits and challenges of caching in the course of building and operating services at Amazon. The remainder of this article describes our lessons learned, best practices, and considerations for using caches.
## When we use caching
Several factors lead us to consider adding a cache to our system. Many times this begins with an observation about a dependency's latency or efficiency at a given request rate. For example, this could be when we determine that a dependency might start throttling or otherwise be unable to keep up with the anticipated load. We’ve found it helpful to consider caching when we encounter uneven request patterns that lead to hot-key/hot-partition throttling. Data from this dependency is a good candidate for caching if such a cache would have a good cache hit ratio across requests. That is, results of calls to the dependency can be used across multiple requests or operations. If each request typically requires a unique query to the dependent service with unique-per-request results, then a cache would have a negligible hit rate and the cache does no good. A second consideration is how tolerant a team’s service and its clients are to eventual consistency. Cached data necessarily grows inconsistent with the source over time, so caching can only be successful if both the service and its clients compensate accordingly. The rate of change of the source data, as well as the cache policy for refreshing data, will determine how inconsistent the data tends to be. These two are related to each other. For example, relatively static or slow-changing data can be cached for longer periods of time.
## Local caches
Service caches can be implemented either in memory or external to the service. On-box caches, commonly implemented in process memory, are relatively quick and easy to implement and can provide significant improvements with minimal work. On-box caches are often the first approach implemented and evaluated once the need for caching is identified. In contrast to external caches, they come with no additional operational overhead, so they are fairly low-risk to integrate into an existing service. We often implement an on-box cache as an in-memory hash table that is managed through application logic (for example, by explicitly placing results into the cache after the service calls are completed) or embedded in the service client (for example, by using a caching HTTP client).
Despite the benefits and seductive simplicity of in-memory caches, they do come with several downsides. One is that the cached data will be inconsistent from server to server across its fleet, manifesting a cache coherence problem. If a client makes repeated calls to the service they might get newer data used in the first call and older data in the second call, depending on which server happens to handle the request.
Another shortcoming is that the downstream load is now proportional to the service's fleet size, so as the number of servers grows it still may be possible to overwhelm dependent services. We’ve found that an effective way to monitor this is to emit metrics on cache hits/misses and the number of requests made to downstream services.
In-memory caches are also susceptible to “cold start” issues. These issues occur where a new server launches with a completely empty cache, which could cause a burst of requests to the dependent service as it fills its cache. This can be a significant issue during deployments or in other circumstances in which the cache is flushed fleet-wide. Cache coherence and empty cache issues can often be addressed by using request coalescing, which is described in detail later in this article.
## Inline vs. side caches
Another decision we need to make when we evaluate different caching approaches is the choice between inline and side caches. _Inline caches_ , or read-through/write-through caches, embed cache management into the main data access API, making cache management an implementation detail of that API. Examples include application-specific implementations like Amazon DynamoDB Accelerator (DAX) and standards-based implementations like HTTP caching (either with a local caching client or an external cache server like Nginx or Varnish). _Side caches_ , in contrast, are generic object stores such as the ones provided by Amazon ElastiCache (Memcached and Redis) or libraries like Ehcache and Google Guava for in-memory caches. With side caches, the application code directly manipulates the cache before and after calls to the data source, checking for cached objects before making the downstream calls, and putting objects in the cache after those calls are completed.
The primary benefit of an inline cache is a uniform API model for clients. Caching can be added, removed, or tweaked without any changes to client logic. An inline cache also pulls cache management logic out of application code, thus eliminating a source of potential bugs. HTTP caches are especially attractive because there are numerous off-the-shelf options available, such as in-memory libraries, standalone HTTP proxies like the ones mentioned previously, and managed services like content delivery networks (CDNs).
However, the transparency of inline caches can also be an availability downside. External caches are now part of the availability equation for this dependency. There is no opportunity for the client to compensate for a temporarily unavailable cache. For example, if you have a Varnish fleet that caches requests from an external REST service, then if that caching fleet goes down, from your service's perspective it’s as if the dependency itself went down. The other downside to an inline cache is that it needs to be built into the protocol or service it is caching for. If an inline cache for the protocol isn't available, then this inline caching isn't an option unless you want to build an integrated client or proxy service yourself.
## Cache expiration
Some of the most challenging cache implementation details are picking the right cache size, expiration policy, and eviction policy. The expiration policy determines how long to retain an item in the cache. The most common policy uses an absolute time-based expiration (that is, it associates a time to live (TTL) with each object as it is loaded). The TTL is chosen based on client requirements, such as how tolerant the client can be to stale data, and how static the data is, because slowly changing data can be more aggressively cached. The ideal cache size is based on a model of the anticipated volume of requests and the distribution of cached objects across those requests. From that, we estimate a cache size can that ensures a high cache hit rate with these traffic patterns. The eviction policy controls how items are removed from the cache when it reaches capacity. The most common eviction policy is Least Recently Used (LRU).
So far, this is just a thought exercise. Real-world traffic patterns can differ from what we model, so we track the actual performance of our cache. Our preferred way to do this is to emit service metrics on cache hits and misses, total cache size, and number of requests to downstream services.
We have learned that we need to be deliberate about picking the cache size and expiration policy values. We want to avoid the situation where a developer arbitrarily picks some cache size and TTL values during initial implementation and then never goes back and validates their appropriateness at a later time. We have seen real-world examples of this lack of follow through leading to temporary service outages and exacerbation of ongoing outages.
Another pattern we use to improve resiliency when downstream services are unavailable is to use two TTLs: a soft TTL and a hard TTL. The client will attempt to refresh cached items based on the soft TTL, but if the downstream service is unavailable or otherwise doesn’t respond to the request, then the existing cache data will continue to be used until the hard TTL is reached. An example of this pattern is used in the AWS Identity and Access Management (IAM) client.
We also use the soft and hard TTL approach with backpressure to reduce the impact of downstream service brownouts. The downstream service can respond with a backpressure event when it is browning out, which signals that the calling service should use cached data until the hard TTL and only make requests for data that are not in its cache. We continue this until the downstream service removes the backpressure. This pattern allows the downstream service to recover from a brownout while maintaining availability of the upstream services.
## Other considerations
An important consideration is what the cache behavior is when errors are received from the downstream service. One option for dealing with this is to reply to clients using the last cached good value, for example leveraging the soft TTL / hard TTL pattern described earlier. Another option we employ is to cache the error response (that is, we use a “negative cache”) using a different TTL than positive cache entries, and propagate the error to the client. The approach we choose in a given situation depends on the particulars of the service and by evaluating when it is better for clients to see stale data versus errors. Regardless of which approach we take, it is important that we ensure something is in the cache in error cases. If this is not the case and the downstream service is temporarily unavailable or otherwise unable to fulfill certain requests (for example when a downstream resource is deleted), the upstream service will continue to bombard it with traffic and potentially either cause an outage or exacerbate an existing one. We have seen real-world examples in which a failure to cache negative responses led to increased failure rates and faults.
Security is another important aspect of caching. When we introduce a cache to a service we evaluate and mitigate any additional security risks it introduces. For example, external caching fleets often lack encryption for serialized data and transport-level security. This is especially important if sensitive user information is retained in the cache. The issue can be mitigated by using something like Amazon ElastiCache for Redis, which supports in-transit and at-rest encryption. Caches are also susceptible to poisoning attacks, in which a vulnerability in the downstream protocol allows an attacker to populate a cache with a value under their control. This amplifies the impact of an attack, since all requests made while this value remains in the cache will see the malicious value. For one final example, caches are also susceptible to side-channel timing attacks. Cached values are returned faster than uncached values, so an attacker can use response time to gain information about requests that other clients or tenets are making.
One final consideration is the “thundering herd” situation, in which many clients make requests that need the same uncached downstream resource at approximately the same time. This can also occur when a server comes up and joins the fleet with an empty local cache. This results in a large number of requests from each server going to the downstream dependency, which can lead to throttling/brownout. To remedy this issue we use request coalescing, where the servers or external cache ensure that only one pending request is out for uncached resources. Some caching libraries provide support for _request coalescing_ , and some external inline caches (such as Nginx or Varnish) do as well. In addition, request coalescing can be implemented on top of existing caches. 
## Amazon best practices and considerations
This article has touched on several Amazon best practices and the trade-offs and risks associated with caching. Here is a summary of the Amazon best practices and considerations that our teams use when they introduce a cache:
• Make sure there is a legitimate need for a cache that is justified in terms of cost, latency, and/or availability improvements. Ensure that the data is cacheable, which means that it can be used across multiple client requests. Be skeptical of the value a cache will bring, and carefully evaluate that the benefits will outweigh the added risks that the cache introduces.
• Plan to operate the cache with the same rigor and processes used for the rest of the service fleet and infrastructure. Don’t underestimate this effort. Emit metrics on cache utilization and hit rate to ensure the cache is tuned appropriately. Monitor key indicators (such as CPU and memory) to ensure that the external caching fleet is healthy and scaled appropriately. Set up alarms on these metrics. Make sure the caching fleet can be scaled up without downtime or mass cache invalidation (that is, validate that consistent hashing is working as expected.)
• Be deliberate and empirical in the choice of cache size, expiration policy, and eviction policy. Perform tests and use the metrics mentioned in the previous bullet to validate and tune these choices.
• Ensure that your service is resilient in the face of cache non-availability, which includes a variety of circumstances that lead to the inability to serve requests using cached data. These include cold starts, caching fleet outages, changes in traffic patterns, or extended downstream outages. In many cases, this could mean trading some of your availability to ensure that your servers and your dependent services don’t brown out (for example by shedding load, capping requests to dependent services, or serving stale data). Run load tests with caches disabled to validate this.
• Consider the security aspects of maintaining cached data, including encryption, transport security when communicating with an external caching fleet, and the impact of cache poisoning attacks and side-channel attacks.
• Design the storage format for cached objects to evolve over time (for example, use a version number) and write serialization code capable of reading older versions. Beware of poison pills in your cache serialization logic.
• Evaluate how the cache will handle downstream errors, and consider maintaining a negative cache with a distinct TTL. Don’t cause or amplify an outage by repeatedly asking for the same downstream resource and discarding the error responses.
Many service teams at Amazon use caching techniques. Despite the benefits of these techniques, we don’t take the decision to incorporate caching lightly because the downsides can often outweigh the upsides. We hope that this article helps you when you evaluate caching in your own services. 
## Matt Brinkley
Matt is a Principal Engineer in Emerging Devices at Amazon, where he works on the software and services for upcoming consumer devices. Previously he worked in AWS Elemental, leading the team that launched MediaTailor, a server-side personalized adinsertion service for live and on-demand video. Along the way he helped launch PrimeVideo’s first season streaming NFL Thursday Night Football. Prior to Amazon, Matt spent 15 years in the security industry, including time at McAfee, Intel and a few startups, working on enterprise security management, anti-malware and anti-exploit technologies, hardware-assisted security measures, and DRM.
[Read more](https://aws.amazon.com/builders-library/authors/matt-brinkley/)
![Missing alt text value](https://d1.awsstatic.com/builderslibrary/authors-biopage/WEB_Redwood_Author-Portrait_Matt-Brinkley.6a417fa1e96049f2a5b5e5963aae03988a8d3552.png)
## Jas Chhabra
Jas Chhabra is a Principal Engineer at AWS. He joined AWS in 2016 and worked on AWS IAM for a couple of years before moving into his current role at AWS Machine Learning. Before AWS, he worked at Intel in various technical roles in IoT, identity, and security areas. Current interests are Machine learning, security, and large-scale distributed systems. Past interests include IoT, bitcoins, identity, and cryptography. He has a Masters in Computer Science.
[Read more](https://aws.amazon.com/builders-library/authors/jas-chhabra/)
![Missing alt text value](https://d1.awsstatic.com/builderslibrary/authors-biopage/WEB_Redwood_Author-Portrait_Jas-Chhabra.257c5521e98ef9997b6b4cc2112c356d4ab4fb2f.png)
## Related content
  * [Avoiding fallback in distributed systems](https://aws.amazon.com/builders-library/avoiding-fallback-in-distributed-systems/)
  * [Using load shedding to avoid overload](https://aws.amazon.com/builders-library/using-load-shedding-to-avoid-overload/)
  * [Ensuring rollback safety during deployments](https://aws.amazon.com/builders-library/ensuring-rollback-safety-during-deployments/)


[Create an AWS account](https://portal.aws.amazon.com/gp/aws/developer/registration/index.html?nc1=f_ct&src=footer_signup)
## Learn
  * [What Is AWS?](https://aws.amazon.com/what-is-aws/?nc1=f_cc)
  * [What Is Cloud Computing?](https://aws.amazon.com/what-is-cloud-computing/?nc1=f_cc)
  * [What Is Agentic AI?](https://aws.amazon.com/what-is/agentic-ai/?nc1=f_cc)
  * [Cloud Computing Concepts Hub](https://aws.amazon.com/what-is/?nc1=f_cc)
  * [AWS Cloud Security](https://aws.amazon.com/security/?nc1=f_cc)
  * [What's New](https://aws.amazon.com/new/?nc1=f_cc)
  * [Blogs](https://aws.amazon.com/blogs/?nc1=f_cc)
  * [Press Releases](https://press.aboutamazon.com/press-releases/aws)


## Resources
  * [Getting Started](https://aws.amazon.com/getting-started/?nc1=f_cc)
  * [Training](https://aws.amazon.com/training/?nc1=f_cc)
  * [AWS Trust Center](https://aws.amazon.com/trust-center/?nc1=f_cc)
  * [AWS Solutions Library](https://aws.amazon.com/solutions/?nc1=f_cc)
  * [Architecture Center](https://aws.amazon.com/architecture/?nc1=f_cc)
  * [Product and Technical FAQs](https://aws.amazon.com/faqs/?nc1=f_dr)
  * [Analyst Reports](https://aws.amazon.com/resources/analyst-reports/?nc1=f_cc)
  * [AWS Partners](https://aws.amazon.com/partners/work-with-partners/?nc1=f_dr)


## Developers
  * [Builder Center](https://aws.amazon.com/developer/?nc1=f_dr)
  * [SDKs & Tools](https://aws.amazon.com/developer/tools/?nc1=f_dr)
  * [.NET on AWS](https://aws.amazon.com/developer/language/net/?nc1=f_dr)
  * [Python on AWS](https://aws.amazon.com/developer/language/python/?nc1=f_dr)
  * [Java on AWS](https://aws.amazon.com/developer/language/java/?nc1=f_dr)
  * [PHP on AWS](https://aws.amazon.com/developer/language/php/?nc1=f_cc)
  * [JavaScript on AWS](https://aws.amazon.com/developer/language/javascript/?nc1=f_dr)


## Help
  * [Contact Us](https://aws.amazon.com/contact-us/?nc1=f_m)
  * [File a Support Ticket](https://console.aws.amazon.com/support/home/?nc1=f_dr)
  * [AWS re:Post](https://repost.aws/?nc1=f_dr)
  * [Knowledge Center](https://repost.aws/knowledge-center/?nc1=f_dr)
  * [AWS Support Overview](https://aws.amazon.com/premiumsupport/?nc1=f_dr)
  * [Get Expert Help](https://iq.aws.amazon.com/?utm=mkt.foot/?nc1=f_m)
  * [AWS Accessibility](https://aws.amazon.com/accessibility/?nc1=f_cc)
  * [Legal](https://aws.amazon.com/legal/?nc1=f_cc)


English
Back to top
Amazon is an Equal Opportunity Employer: Minority / Women / Disability / Veteran / Gender Identity / Sexual Orientation / Age.
[x](https://twitter.com/awscloud)[facebook](https://www.facebook.com/amazonwebservices)[linkedin](https://www.linkedin.com/company/amazon-web-services/)[instagram](https://www.instagram.com/amazonwebservices/)[twitch](https://www.twitch.tv/aws)[youtube](https://www.youtube.com/user/AmazonWebServices/Cloud/)[podcasts](https://aws.amazon.com/podcasts/?nc1=f_cc)[email](https://pages.awscloud.com/communication-preferences?trk=homepage)
  * [Privacy](https://aws.amazon.com/privacy/?nc1=f_pr)
  * [Site terms](https://aws.amazon.com/terms/?nc1=f_pr)
  * [Cookie Preferences](https://aws.amazon.com/builders-library/caching-challenges-and-strategies/)


© 2025, Amazon Web Services, Inc. or its affiliates. All rights reserved.
Hi, I can connect you with an AWS representative or answer questions you have on AWS.
Need more info? Highlight any text to get an explanation generated with AWS generative AI.
2
Get the gist with ExplainerWith Explainer, you can highlight any text to get an explanation generated with AWS generative AI. Learn new terms or product info—no searching required. To get started, turn on the **Explainer** toggle in the lower right.
Continue


## Source Information
- URL: https://aws.amazon.com/builders-library/caching-challenges-and-strategies/
- Harvested: 2025-08-19T23:13:47.349994
- Profile: deep_research
- Agent: architect
