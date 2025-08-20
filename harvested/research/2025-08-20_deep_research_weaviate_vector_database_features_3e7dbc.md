---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-20T13:07:46.003904'
profile: deep_research
source: https://weaviate.io/developers/weaviate
topic: Weaviate Vector Database Features
---

# Weaviate Vector Database Features

[Skip to main content](https://docs.weaviate.io/weaviate#__docusaurus_skipToContent_fallback)
[Product Update: Meet Weaviate Agents — Read the blog](https://weaviate.io/blog/weaviate-agents)
[![Weaviate](https://docs.weaviate.io/img/site/weaviate-logo-horizontal-light-1.svg)](https://weaviate.io)[](https://docs.weaviate.io/weaviate/api/rest)Weaviate Database
[GitHub](https://github.com/weaviate/weaviate)[Weaviate Cloud](https://console.weaviate.cloud)
Ask AI or Search
⌘K
[Get Started](https://docs.weaviate.io/weaviate)[How-to manuals & Guides](https://docs.weaviate.io/weaviate/guides)[Model Integrations](https://docs.weaviate.io/weaviate/model-providers)[Reference & APIs](https://docs.weaviate.io/weaviate/config-refs)[Concepts](https://docs.weaviate.io/weaviate/concepts)[Recipes](https://docs.weaviate.io/weaviate/recipes)[Other](https://docs.weaviate.io/weaviate/release-notes)
**Go to documentation:**
⌘U✕
**Weaviate Database**
Develop AI applications using Weaviate's APIs and tools
**Deploy**
Deploy, configure, and maintain Weaviate Database
**Weaviate Agents**
Build and deploy intelligent agents with Weaviate
**Weaviate Cloud**
Manage and scale Weaviate in the cloud
#### Additional resources
Academy
Integrations
Contributor guide
Events & Workshops
#### Need help?
![Weaviate Logo](https://docs.weaviate.io/img/site/weaviate-logo-w.png)Ask AI Assistant⌘K
Community Forum
  * [Introduction](https://docs.weaviate.io/weaviate)
  * [Quickstart](https://docs.weaviate.io/weaviate/quickstart)
  * [Installation](https://docs.weaviate.io/deploy)
  * [Connect to Weaviate](https://docs.weaviate.io/weaviate/connections)
  * [Starter guides](https://docs.weaviate.io/weaviate/starter-guides)
  * [Best practices](https://docs.weaviate.io/weaviate/best-practices)
  * [AI-based code generation](https://docs.weaviate.io/weaviate/best-practices/code-generation)


  * Introduction


On this page
# Weaviate Database
Weaviate _(we-vee-eight)_ is an open-source, AI-native vector database. Use this documentation to get started with Weaviate and learn how to get the most out of Weaviate's features.
[New to Weaviate?Start with the Quickstart tutorial – an end-to-end demo that takes 15–30 minutes.](https://docs.weaviate.io/weaviate/quickstart)
## Find the right documentation and resources[​](https://docs.weaviate.io/weaviate#find-the-right-documentation-and-resources "Direct link to Find the right documentation and resources")
The Weaviate documentation is structured into multiple units based on the service and functionality.
[Weaviate DatabaseDevelop AI applications using Weaviate's APIs and tools](https://docs.weaviate.io/weaviate)[DeployDeploy, configure, and maintain Weaviate Database](https://docs.weaviate.io/deploy)[Weaviate AgentsBuild and deploy intelligent agents with Weaviate](https://docs.weaviate.io/agents)[Weaviate CloudManage and scale Weaviate in the cloud](https://docs.weaviate.io/cloud)
## What is Weaviate?[​](https://docs.weaviate.io/weaviate#what-is-weaviate "Direct link to What is Weaviate?")
Weaviate is an **open-source vector database** designed to store and index both data objects and their vector embeddings. This architecture enables advanced semantic search capabilities by comparing the meaning encoded in vectors rather than relying solely on keyword matching. Key capabilities include:
  * **[Semantic and hybrid search](https://docs.weaviate.io/weaviate/search/basics)** By indexing data with vectors, Weaviate supports searches based on both semantic similarity and keywords. This allows for more relevant results even when the query terms don’t exactly match the stored data.
  * **[Retrieval augmented generation (RAG)](https://docs.weaviate.io/weaviate/search/generative)** Weaviate can serve as a robust backend for RAG workflows, where vector search is used to retrieve context that enhances the output of generative models, making it easier to generate accurate, context-aware responses.
  * **[Agent-driven workflows](https://docs.weaviate.io/agents)** Its flexible API and integration with modern AI models make Weaviate suitable for powering applications that rely on intelligent agents. These agents can leverage semantic insights to make decisions or trigger actions based on the data stored in Weaviate.


## The Weaviate Ecosystem[​](https://docs.weaviate.io/weaviate#the-weaviate-ecosystem "Direct link to The Weaviate Ecosystem")
The Weaviate ecosystem consists of multiple tools and services centered around building cloud-native AI-powered applications.
![The Weaviate Ecosystem](https://docs.weaviate.io/assets/images/weaviate-ecosystem-1b6ed8c65d142f97017214842f441f23.png)
As shown in the high-level overview above, the ecosystem consists of:
  * **[Weaviate Database](https://docs.weaviate.io/weaviate#what-is-weaviate)** : An open source vector database that stores both objects and vectors.
  * **[Weaviate Cloud](https://docs.weaviate.io/cloud)** : A fully managed cloud deployment of the Weaviate vector database.
  * **[Weaviate Agents](https://docs.weaviate.io/agents)** : Pre-built agentic services for Weaviate Cloud users.
  * **[Weaviate Embeddings](https://docs.weaviate.io/cloud/embeddings)** : A managed embedding inference service for Weaviate Cloud users.
  * **[External model providers](https://docs.weaviate.io/weaviate/model-providers)** : Third-party models that integrate with Weaviate.


## Choose your deployment[​](https://docs.weaviate.io/weaviate#choose-your-deployment "Direct link to Choose your deployment")
[EvaluationDeploymentProductionWeaviate Cloud
  * From evaluation (sandbox) to production
  * Serverless (infrastructure managed by Weaviate)
  * (Optional) Data replication (high-availability)
  * (Optional) Zero-downtime updates

Set up a WCD instance](https://docs.weaviate.io/cloud/manage-clusters/create)[EvaluationDeploymentProductionDocker
  * For local evaluation & development
  * Local inference containers
  * Multi-modal models
  * Customizable configurations

Run Weaviate with Docker](https://docs.weaviate.io/deploy/installation-guides/docker-installation)[EvaluationDeploymentProductionKubernetes
  * For development to production
  * Local inference containers
  * Multi-modal models
  * Customizable configurations
  * Self-deploy or Marketplace deployment
  * (Optional) Zero-downtime updates

Run Weaviate with Kubernetes](https://docs.weaviate.io/deploy/installation-guides/k8s-installation)[EvaluationDeploymentProductionEmbedded Weaviate
  * For basic, quick evaluation
  * Conveniently launch Weaviate directly from Python or JS/TS

Run Embedded Weaviate](https://docs.weaviate.io/deploy/installation-guides/embedded)
## Community & support[​](https://docs.weaviate.io/weaviate#community--support "Direct link to Community & support")
[QuestionsPlease visit our forum. The Weaviate team and our awesome community can help.](https://forum.weaviate.io/c/support/)[Open-source on GitHubGive us a star on GitHub to support our work.](https://github.com/weaviate/weaviate)
[Edit this page](https://github.com/weaviate/docs/tree/main/docs/weaviate/index.mdx)
[](https://github.com/weaviate/weaviate "GitHub")[](https://weaviate.io/slack "Slack")[](https://x.com/weaviate_io "X")[](https://instagram.com/weaviate.io "Instagram")[](https://youtube.com/@Weaviate "YouTube")[](https://www.linkedin.com/company/weaviate-io "LinkedIn")
##### Documentation
  * [Weaviate Database](https://docs.weaviate.io/weaviate)
  * [Deployment documentation](https://docs.weaviate.io/deploy)
  * [Weaviate Cloud](https://docs.weaviate.io/cloud)
  * [Weaviate Agents](https://docs.weaviate.io/agents)


##### Support
  * [Forum](https://forum.weaviate.io/c/support/)
  * [Slack](https://weaviate.io/slack)


  * [Find the right documentation and resources](https://docs.weaviate.io/weaviate#find-the-right-documentation-and-resources)
  * [What is Weaviate?](https://docs.weaviate.io/weaviate#what-is-weaviate)
  * [The Weaviate Ecosystem](https://docs.weaviate.io/weaviate#the-weaviate-ecosystem)
  * [Choose your deployment](https://docs.weaviate.io/weaviate#choose-your-deployment)
  * [Community & support](https://docs.weaviate.io/weaviate#community--support)


## Weaviate documentation structure
✕
The Weaviate documentation is divided into the following four sections: [core database](https://docs.weaviate.io/weaviate), [deployment](https://docs.weaviate.io/deploy), [Weaviate Cloud](https://docs.weaviate.io/cloud) and [Weaviate Agents](https://docs.weaviate.io/agents) docs.


## Source Information
- URL: https://weaviate.io/developers/weaviate
- Harvested: 2025-08-20T13:07:46.003904
- Profile: deep_research
- Agent: architect
