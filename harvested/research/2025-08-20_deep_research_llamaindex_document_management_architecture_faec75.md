---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-20T13:12:16.066086'
profile: deep_research
source: https://docs.llamaindex.ai/en/stable
topic: LlamaIndex Document Management Architecture
---

# LlamaIndex Document Management Architecture

[ Skip to content ](https://docs.llamaindex.ai/en/stable/#welcome-to-llamaindex)
[ ![logo](https://docs.llamaindex.ai/en/stable/_static/assets/LlamaSquareBlack.svg) ](https://docs.llamaindex.ai/en/stable/ "LlamaIndex")
LlamaIndex 
LlamaIndex 
Search`K`
  * [ Home ](https://docs.llamaindex.ai/en/stable/)
  * [ Learn ](https://docs.llamaindex.ai/en/stable/understanding/)
  * [ Use Cases ](https://docs.llamaindex.ai/en/stable/use_cases/)
  * [ Examples ](https://docs.llamaindex.ai/en/stable/examples/)
  * [ Component Guides ](https://docs.llamaindex.ai/en/stable/module_guides/)
  * [ Advanced Topics ](https://docs.llamaindex.ai/en/stable/optimizing/production_rag/)
  * [ API Reference ](https://docs.llamaindex.ai/en/stable/api_reference/)
  * [ Open-Source Community ](https://docs.llamaindex.ai/en/stable/community/integrations/)
  * [ Workflows ](https://docs.llamaindex.ai/en/stable/workflows/)
  * [ LlamaCloud ](https://docs.cloud.llamaindex.ai/)


[ ![logo](https://docs.llamaindex.ai/en/stable/_static/assets/LlamaSquareBlack.svg) ](https://docs.llamaindex.ai/en/stable/ "LlamaIndex") LlamaIndex 
  * [ Home  ](https://docs.llamaindex.ai/en/stable/)
Home 
    * [ High-Level Concepts  ](https://docs.llamaindex.ai/en/stable/getting_started/concepts/)
    * [ Installation and Setup  ](https://docs.llamaindex.ai/en/stable/getting_started/installation/)
    * [ How to read these docs  ](https://docs.llamaindex.ai/en/stable/getting_started/reading/)
    * [ Starter Examples  ](https://docs.llamaindex.ai/en/stable/getting_started/starter_example/)
    * [ Discover LlamaIndex Video Series  ](https://docs.llamaindex.ai/en/stable/getting_started/discover_llamaindex/)
    * [ Frequently Asked Questions (FAQ)  ](https://docs.llamaindex.ai/en/stable/getting_started/customization/)
    * [ Starter Tools  ](https://docs.llamaindex.ai/en/stable/getting_started/starter_tools/)
  * [ Learn  ](https://docs.llamaindex.ai/en/stable/understanding/)
  * [ Use Cases  ](https://docs.llamaindex.ai/en/stable/use_cases/)
  * [ Examples  ](https://docs.llamaindex.ai/en/stable/examples/)
  * [ Component Guides  ](https://docs.llamaindex.ai/en/stable/module_guides/)
  * [ Advanced Topics  ](https://docs.llamaindex.ai/en/stable/optimizing/production_rag/)
  * [ API Reference  ](https://docs.llamaindex.ai/en/stable/api_reference/)
  * [ Open-Source Community  ](https://docs.llamaindex.ai/en/stable/community/integrations/)
  * [ Workflows  ](https://docs.llamaindex.ai/en/stable/workflows/)
  * [ LlamaCloud  ](https://docs.cloud.llamaindex.ai/)


Table of contents 
  * [ Introduction  ](https://docs.llamaindex.ai/en/stable/#introduction)
    * [ What are agents?  ](https://docs.llamaindex.ai/en/stable/#what-are-agents)
    * [ What are workflows?  ](https://docs.llamaindex.ai/en/stable/#what-are-workflows)
    * [ What is context augmentation?  ](https://docs.llamaindex.ai/en/stable/#what-is-context-augmentation)
    * [ LlamaIndex is the framework for Context-Augmented LLM Applications  ](https://docs.llamaindex.ai/en/stable/#llamaindex-is-the-framework-for-context-augmented-llm-applications)
  * [ Use cases  ](https://docs.llamaindex.ai/en/stable/#use-cases)
    * [ 👨‍👩‍👧‍👦 Who is LlamaIndex for?  ](https://docs.llamaindex.ai/en/stable/#who-is-llamaindex-for)
  * [ Getting Started  ](https://docs.llamaindex.ai/en/stable/#getting-started)
    * [ 30 second quickstart  ](https://docs.llamaindex.ai/en/stable/#30-second-quickstart)
  * [ LlamaCloud  ](https://docs.llamaindex.ai/en/stable/#llamacloud)
  * [ Community  ](https://docs.llamaindex.ai/en/stable/#community)
    * [ Getting the library  ](https://docs.llamaindex.ai/en/stable/#getting-the-library)
    * [ Contributing  ](https://docs.llamaindex.ai/en/stable/#contributing)
  * [ LlamaIndex Ecosystem  ](https://docs.llamaindex.ai/en/stable/#llamaindex-ecosystem)


# Welcome to LlamaIndex 🦙 ![#](https://docs.llamaindex.ai/en/stable/#welcome-to-llamaindex "Permanent link")
LlamaIndex is the leading framework for building LLM-powered agents over your data with [LLMs](https://en.wikipedia.org/wiki/Large_language_model) and [workflows](https://docs.llamaindex.ai/en/stable/understanding/workflows/).
  * [Introduction](https://docs.llamaindex.ai/en/stable/#introduction)
What is context augmentation? What are agents and workflows? How does LlamaIndex help build them?
  * [Use cases](https://docs.llamaindex.ai/en/stable/#use-cases)
What kind of apps can you build with LlamaIndex? Who should use it?
  * [Getting started](https://docs.llamaindex.ai/en/stable/#getting-started)
Get started in Python or TypeScript in just 5 lines of code!
  * [LlamaCloud](https://docs.cloud.llamaindex.ai/)
Managed services for LlamaIndex including [LlamaParse](https://docs.cloud.llamaindex.ai/llamaparse/overview), the world's best document parser.
  * [Community](https://docs.llamaindex.ai/en/stable/#community)
Get help and meet collaborators on Discord, Twitter, LinkedIn, and learn how to contribute to the project.
  * [Related projects](https://docs.llamaindex.ai/en/stable/#related-projects)
Check out our library of connectors, readers, and other integrations at [LlamaHub](https://llamahub.ai) as well as demos and starter apps like [create-llama](https://www.npmjs.com/package/create-llama).


## Introduction[#](https://docs.llamaindex.ai/en/stable/#introduction "Permanent link")
### What are agents?[#](https://docs.llamaindex.ai/en/stable/#what-are-agents "Permanent link")
[Agents](https://docs.llamaindex.ai/en/stable/understanding/agent/) are LLM-powered knowledge assistants that use tools to perform tasks like research, data extraction, and more. Agents range from simple question-answering to being able to sense, decide and take actions in order to complete tasks.
LlamaIndex provides a framework for building agents including the ability to use RAG pipelines as one of many tools to complete a task.
### What are workflows?[#](https://docs.llamaindex.ai/en/stable/#what-are-workflows "Permanent link")
[Workflows](https://docs.llamaindex.ai/en/stable/understanding/workflows/) are multi-step processes that combine one or more agents, data connectors, and other tools to complete a task. They are event-driven software that allows you to combine RAG data sources and multiple agents to create a complex application that can perform a wide variety of tasks with reflection, error-correction, and other hallmarks of advanced LLM applications. You can then [deploy these agentic workflows](https://docs.llamaindex.ai/en/stable/module_guides/workflow/deployment.md) as production microservices.
### What is context augmentation?[#](https://docs.llamaindex.ai/en/stable/#what-is-context-augmentation "Permanent link")
LLMs offer a natural language interface between humans and data. LLMs come pre-trained on huge amounts of publicly available data, but they are not trained on **your** data. Your data may be private or specific to the problem you're trying to solve. It's behind APIs, in SQL databases, or trapped in PDFs and slide decks.
Context augmentation makes your data available to the LLM to solve the problem at hand. LlamaIndex provides the tools to build any of context-augmentation use case, from prototype to production. Our tools allow you to ingest, parse, index and process your data and quickly implement complex query workflows combining data access with LLM prompting.
The most popular example of context-augmentation is [Retrieval-Augmented Generation or RAG](https://docs.llamaindex.ai/en/stable/getting_started/concepts/), which combines context with LLMs at inference time.
### LlamaIndex is the framework for Context-Augmented LLM Applications[#](https://docs.llamaindex.ai/en/stable/#llamaindex-is-the-framework-for-context-augmented-llm-applications "Permanent link")
LlamaIndex imposes no restriction on how you use LLMs. You can use LLMs as auto-complete, chatbots, agents, and more. It just makes using them easier. We provide tools like:
  * **Data connectors** ingest your existing data from their native source and format. These could be APIs, PDFs, SQL, and (much) more.
  * **Data indexes** structure your data in intermediate representations that are easy and performant for LLMs to consume.
  * **Engines** provide natural language access to your data. For example:
    * Query engines are powerful interfaces for question-answering (e.g. a RAG flow).
    * Chat engines are conversational interfaces for multi-message, "back and forth" interactions with your data.
  * **Agents** are LLM-powered knowledge workers augmented by tools, from simple helper functions to API integrations and more.
  * **Observability/Evaluation** integrations that enable you to rigorously experiment, evaluate, and monitor your app in a virtuous cycle.
  * **Workflows** allow you to combine all of the above into an event-driven system far more flexible than other, graph-based approaches.


## Use cases[#](https://docs.llamaindex.ai/en/stable/#use-cases "Permanent link")
Some popular use cases for LlamaIndex and context augmentation in general include:
  * [Question-Answering](https://docs.llamaindex.ai/en/stable/use_cases/q_and_a/) (Retrieval-Augmented Generation aka RAG)
  * [Chatbots](https://docs.llamaindex.ai/en/stable/use_cases/chatbots/)
  * [Document Understanding and Data Extraction](https://docs.llamaindex.ai/en/stable/use_cases/extraction/)
  * [Autonomous Agents](https://docs.llamaindex.ai/en/stable/use_cases/agents/) that can perform research and take actions
  * [Multi-modal applications](https://docs.llamaindex.ai/en/stable/use_cases/multimodal/) that combine text, images, and other data types
  * [Fine-tuning](https://docs.llamaindex.ai/en/stable/use_cases/fine_tuning/) models on data to improve performance


Check out our [use cases](https://docs.llamaindex.ai/en/stable/use_cases/) documentation for more examples and links to tutorials.
### 👨‍👩‍👧‍👦 Who is LlamaIndex for?[#](https://docs.llamaindex.ai/en/stable/#who-is-llamaindex-for "Permanent link")
LlamaIndex provides tools for beginners, advanced users, and everyone in between.
Our high-level API allows beginner users to use LlamaIndex to ingest and query their data in 5 lines of code.
For more complex applications, our lower-level APIs allow advanced users to customize and extend any module -- data connectors, indices, retrievers, query engines, and reranking modules -- to fit their needs.
## Getting Started[#](https://docs.llamaindex.ai/en/stable/#getting-started "Permanent link")
LlamaIndex is available in Python (these docs) and [Typescript](https://ts.llamaindex.ai/). If you're not sure where to start, we recommend reading [how to read these docs](https://docs.llamaindex.ai/en/stable/getting_started/reading/) which will point you to the right place based on your experience level.
### 30 second quickstart[#](https://docs.llamaindex.ai/en/stable/#30-second-quickstart "Permanent link")
Set an environment variable called `OPENAI_API_KEY` with an [OpenAI API key](https://platform.openai.com/api-keys). Install the Python library:
```
pipinstallllama-index

```

Put some documents in a folder called `data`, then ask questions about them with our famous 5-line starter:
```
from llama_index.core import VectorStoreIndex, SimpleDirectoryReader
documents = SimpleDirectoryReader("data").load_data()
index = VectorStoreIndex.from_documents(documents)
query_engine = index.as_query_engine()
response = query_engine.query("Some question about the data should go here")
print(response)

```

If any part of this trips you up, don't worry! Check out our more comprehensive starter tutorials using [remote APIs like OpenAI](https://docs.llamaindex.ai/en/stable/getting_started/starter_example/) or [any model that runs on your laptop](https://docs.llamaindex.ai/en/stable/getting_started/starter_example_local/).
## LlamaCloud[#](https://docs.llamaindex.ai/en/stable/#llamacloud "Permanent link")
If you're an enterprise developer, check out [**LlamaCloud**](https://llamaindex.ai/enterprise). It is an end-to-end managed service for document parsing, extraction, indexing, and retrieval - allowing you to get production-quality data for your AI agent. You can [sign up](https://cloud.llamaindex.ai/) and get 10,000 free credits per month, sign up for one of our [plans](https://www.llamaindex.ai/pricing), or [come talk to us](https://www.llamaindex.ai/contact) if you're interested in an enterprise solution. We offer both SaaS and self-hosted plans.
You can also check out the [LlamaCloud documentation](https://docs.cloud.llamaindex.ai/) for more details.
  * **Document Parsing (LlamaParse)** : LlamaParse is the best-in-class document parsing solution. It's powered by VLMs and perfect for even the most complex documents (nested tables, embedded charts/images, and more). [Learn more](https://www.llamaindex.ai/llamaparse) or check out the [docs](https://docs.cloud.llamaindex.ai/llamaparse/overview).
  * **Document Extraction (LlamaExtract)** : Given a human-defined or inferred schema, extract structured data from any document. [Learn more](https://www.llamaindex.ai/llamaextract) or check out the [docs](https://docs.cloud.llamaindex.ai/llamaextract/getting_started).
  * **Indexing/Retrieval** : Set up an e2e pipeline to index a collection of documents for retrieval. Connect your data source (e.g. Sharepoint, Google Drive, S3), your vector DB data sink, and we automatically handle the document processing and syncing. [Learn more](https://www.llamaindex.ai/enterprise) or check out the [docs](https://docs.cloud.llamaindex.ai/llamacloud/getting_started).


## Community[#](https://docs.llamaindex.ai/en/stable/#community "Permanent link")
Need help? Have a feature suggestion? Join the LlamaIndex community:
  * [Twitter](https://twitter.com/llama_index)
  * [Discord](https://discord.gg/dGcwcsnxhU)
  * [LinkedIn](https://www.linkedin.com/company/llamaindex/)


### Getting the library[#](https://docs.llamaindex.ai/en/stable/#getting-the-library "Permanent link")
  * LlamaIndex Python
    * [LlamaIndex Python Github](https://github.com/run-llama/llama_index)
    * [Python Docs](https://docs.llamaindex.ai/) (what you're reading now)
    * [LlamaIndex on PyPi](https://pypi.org/project/llama-index/)
  * LlamaIndex.TS (Typescript/Javascript package):
    * [LlamaIndex.TS Github](https://github.com/run-llama/LlamaIndexTS)
    * [TypeScript Docs](https://ts.llamaindex.ai/)
    * [LlamaIndex.TS on npm](https://www.npmjs.com/package/llamaindex)


### Contributing[#](https://docs.llamaindex.ai/en/stable/#contributing "Permanent link")
We are open-source and always welcome contributions to the project! Check out our [contributing guide](https://docs.llamaindex.ai/en/stable/CONTRIBUTING/) for full details on how to extend the core library or add an integration to a third party like an LLM, a vector store, an agent tool and more.
## LlamaIndex Ecosystem[#](https://docs.llamaindex.ai/en/stable/#llamaindex-ecosystem "Permanent link")
There's more to the LlamaIndex universe! Check out some of our other projects:
  * [llama_deploy](https://github.com/run-llama/llama_deploy) | Deploy your agentic workflows as production microservices
  * [LlamaHub](https://llamahub.ai) | A large (and growing!) collection of custom data connectors
  * [SEC Insights](https://secinsights.ai) | A LlamaIndex-powered application for financial research
  * [create-llama](https://www.npmjs.com/package/create-llama) | A CLI tool to quickly scaffold LlamaIndex projects


Back to top  [ Next  High-Level Concepts  ](https://docs.llamaindex.ai/en/stable/getting_started/concepts/)


## Source Information
- URL: https://docs.llamaindex.ai/en/stable
- Harvested: 2025-08-20T13:12:16.066086
- Profile: deep_research
- Agent: architect
