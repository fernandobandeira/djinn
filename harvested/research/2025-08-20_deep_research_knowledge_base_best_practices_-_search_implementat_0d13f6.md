---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-20T13:10:52.766853'
profile: deep_research
source: https://blog.replit.com/search
topic: Knowledge Base Best Practices - Search Implementation Patterns
---

# Knowledge Base Best Practices - Search Implementation Patterns

[](https://replit.com)
Features
[Replit Agent](https://replit.com/ai)[Collaboration](https://replit.com/collaboration)[Deployments](https://replit.com/deployments)
[Teams](https://replit.com/teams)[Pricing](https://replit.com/pricing)[Blog](https://blog.replit.com)[Careers](https://replit.com/site/careers)
[Log in](https://replit.com/login)[Start building](https://replit.com/signup)
[](https://replit.com)
  * Features
  * [Pricing](https://replit.com/pricing)
  * [Blog](https://blog.replit.com)
  * [Careers](https://replit.com/site/careers)


[Log in](https://replit.com/login)[Start building](https://replit.com/signup)
[Blog](https://blog.replit.com/)
  * [Product](https://blog.replit.com/category/product)
  * [Eng](https://blog.replit.com/category/eng)
  * [Infra](https://blog.replit.com/category/infra)


# We Built a Search Engine
![](https://cdn.sanity.io/images/bj34pdbp/migration/39faf0d6bc583d95641bc19e945624ebae64fab1-1483x978.png?w=3840&q=75&fit=clip&auto=format)
Wed, Mar 30, 2022
Updated at:
Mon, Oct 9, 2023
![Søren Rood](https://cdn.sanity.io/images/bj34pdbp/migration/344b862820b86a7a5c843d6d328ed16b02211f56-1026x1012.png)
Søren Rood
![Lincoln Bergeson](https://cdn.sanity.io/images/bj34pdbp/migration/5a417f53a57451f1d3c1f18e75c40859bc5553f6-6000x9000.jpg)
Lincoln Bergeson
For the past few months, we have been building a Replit-native search engine. It is remarkably powerful, and we are really excited for you all to try it out. **We believe that you should be able to find anything on Replit in less than 30 seconds.**
This might sound simple, but when you have 100 million+ Repls, it becomes complicated. :)
When you search for something on Replit today, you'll see a page with relevant results from the following categories:
  * Repls
  * Templates
  * Code (yes, code)
  * Users
  * Community content (published repls, posts)
  * Our docs
  * Community tags


![search](https://cdn.sanity.io/images/bj34pdbp/migration/39faf0d6bc583d95641bc19e945624ebae64fab1-1483x978.png?w=3840&q=100&fit=max&auto=format)
search
### Why we built it
We think that being able to search for stuff is important. Searching can help you find inspiration, discover other users, and even learn how to code.
Prior to this update, there was no effective way for users to find anything on Replit. You could only search through the names of Repls that you owned. Even then, you had to type the exact string match to get any results.
We have known for a while that the search experience needed to be updated, but for much of last year we focused on more critical projects. (Workspace stability, abuse, etc.)
Aside from our team's general dissatisfaction, there were a few other signals that prompted us to fix the search experience:
  * qualitatively, people complained about the lack of functionality in surveys and interviews.
  * quantitatively, 80%~ of users that navigated to search ended up dropping off entirely. They weren't able to find what they were looking for. Retention for this feature was also significantly lower than other features.


In building this, we hope to give everyone tools for the self-led discovery of all different types of content on Replit.
### How we built it
A search engine has a few different components:
  * A web frontend where users type their queries and see results
  * A server backend that receives requests from the frontend and serves responses
  * The actual search engine that indexes documents and executes search queries
  * Data pipelines for creating and updating search indexes


The first two were relatively straightforward. We built the search page like we build every other page on Replit: in Typescript, with a Next.js (React) frontend and an Express backend, with GraphQL as the API layer in between.
The other components were a little more complicated, since we haven't done anything like this at scale before.
A few other small search features on the site such as the community search box were built with [the Sphinx library](http://sphinxsearch.com/). This was simple to set up at the time with some custom infrastructure to support indexing and executing queries. However, this setup hasn't aged well, and we weren't comfortable expanding it to vastly larger datasets.
After looking through our options, we settled on Elasticsearch because of its stability, supporting libraries for big data ingestion, and the large community that has battle-tested it in production setups over the years.
For the data pipelines we use Apache Spark. These are high-throughput jobs that run on clusters of many machines at once. The details of how these jobs work is worth a post on its own, but broadly speaking they take data from other locations (the BigQuery data warehouse, Google Cloud Storage, etc), process them into search documents, and bulk upload to the Elasticsearch cluster.
We're working on scaling these big data jobs even further. For example, our code search supports searching through the contents of published repls. In the future, we'll expand this to every file in every repl. You'll be able to find even more working code results and sample code to learn from.
### What's next?
On March 31st, we released search to 100% of users.
If you helped with any early prototypes, participated in any user interviews, or were a part of the explorer beta, we want to say thank you! We could not have made it this far without your help.
If you'd like to leave feedback on the new search experience, please leave a comment on [this](https://replit.com/@util/Search-20-Feedback?c=212850) Repl. We're open to any and all types of feedback.
If this project seems interesting to you, consider checking out our [careers](https://replit.com/site/careers) page. :)
[![Replit logo](https://blog.replit.com/icons/logo-mark-orange.svg)Follow @Replit](https://twitter.com/replit)
[](https://twitter.com/intent/tweet?url=https://blog.replit.com/search)[](https://www.facebook.com/sharer/sharer.php?u=https://blog.replit.com/search)
More
  * [![](https://cdn.sanity.io/images/bj34pdbp/migration/afc22f69f32705c9a4d1d9033d503f450c0c64c4-4800x2520.png?w=3840&q=75&fit=clip&auto=format)Fri, Aug 15, 2025Introducing Image Generation for Replit Agent: Bring Your Ideas to Life in PixelsIntroducing Image Generation for Replit Agent We’re excited to announce a brand new capability for Replit Agent: AI-po...](https://blog.replit.com/image-generation)
  * [![](https://cdn.sanity.io/images/bj34pdbp/migration/af1f0c4239822362e288209693b6b7918bae4b7b-4660x2446.png?w=3840&q=75&fit=clip&auto=format)Wed, Aug 13, 2025Introducing App Storage – building apps with images, video, and PDFs just got easierIntroducing App Storage – building apps with images, video, and PDFs just got easier Today we're excited to announce Ap...](https://blog.replit.com/app-storage)
  * [![](https://cdn.sanity.io/images/bj34pdbp/migration/32e79c7e1a53a60c922caaae1334775c29dd8a38-4800x2520.png?w=3840&q=75&fit=clip&auto=format)Tue, Jul 22, 2025Introducing Queue: A smarter way to work with AgentToday, we’re excited to introduce Queue, a new capability designed to enhance the core Replit Agent experience. Queue al...](https://blog.replit.com/introducing-queue-a-smarter-way-to-work-with-agent)




## Source Information
- URL: https://blog.replit.com/search
- Harvested: 2025-08-20T13:10:52.766853
- Profile: deep_research
- Agent: architect
