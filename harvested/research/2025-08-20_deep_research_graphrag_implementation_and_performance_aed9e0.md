---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-20T13:07:43.385106'
profile: deep_research
source: https://microsoft.github.io/graphrag/
topic: GraphRAG Implementation and Performance
---

# GraphRAG Implementation and Performance

We use optional cookies to improve your experience on our websites, such as through social media connections, and to display personalized advertising based on your online activity. If you reject optional cookies, only cookies necessary to provide you the services will be used. You may change your selection by clicking “Manage Cookies” at the bottom of the page. [Privacy Statement](https://go.microsoft.com/fwlink/?LinkId=521839) [Third-Party Cookies](https://aka.ms/3rdpartycookies)
Accept Reject Manage cookies
[ Skip to content ](https://microsoft.github.io/graphrag/#welcome-to-graphrag)
[ ](https://microsoft.github.io/graphrag/ "GraphRAG")
GraphRAG 
Welcome 
Type to start searching
[ graphrag 
  * v.2.5.0
  * 27.4k
  * 2.8k

](https://github.com/microsoft/graphrag "Go to repository")
  * [ Home ](https://microsoft.github.io/graphrag/)
  * [ Indexing ](https://microsoft.github.io/graphrag/index/overview/)
  * [ Prompt Tuning ](https://microsoft.github.io/graphrag/prompt_tuning/overview/)
  * [ Query ](https://microsoft.github.io/graphrag/query/overview/)
  * [ Configuration ](https://microsoft.github.io/graphrag/config/overview/)
  * [ CLI ](https://microsoft.github.io/graphrag/cli/)
  * [ Extras ](https://microsoft.github.io/graphrag/blog_posts/)


[ ](https://microsoft.github.io/graphrag/ "GraphRAG") GraphRAG 
[ graphrag 
  * v.2.5.0
  * 27.4k
  * 2.8k

](https://github.com/microsoft/graphrag "Go to repository")
  * Home  Home 
    * Welcome  [ Welcome  ](https://microsoft.github.io/graphrag/) Table of contents 
      * [ Get Started with GraphRAG 🚀  ](https://microsoft.github.io/graphrag/#get-started-with-graphrag)
      * [ GraphRAG vs Baseline RAG 🔍  ](https://microsoft.github.io/graphrag/#graphrag-vs-baseline-rag)
      * [ The GraphRAG Process 🤖  ](https://microsoft.github.io/graphrag/#the-graphrag-process)
        * [ Index  ](https://microsoft.github.io/graphrag/#index)
        * [ Query  ](https://microsoft.github.io/graphrag/#query)
        * [ Prompt Tuning  ](https://microsoft.github.io/graphrag/#prompt-tuning)
      * [ Versioning  ](https://microsoft.github.io/graphrag/#versioning)
    * [ Getting Started  ](https://microsoft.github.io/graphrag/get_started/)
    * [ Development Guide  ](https://microsoft.github.io/graphrag/developing/)
  * Indexing  Indexing 
    * [ Overview  ](https://microsoft.github.io/graphrag/index/overview/)
    * [ Dataflow  ](https://microsoft.github.io/graphrag/index/default_dataflow/)
    * [ Methods  ](https://microsoft.github.io/graphrag/index/methods/)
    * [ Inputs  ](https://microsoft.github.io/graphrag/index/inputs/)
    * [ Outputs  ](https://microsoft.github.io/graphrag/index/outputs/)
    * [ Custom Graphs  ](https://microsoft.github.io/graphrag/index/byog/)
  * Prompt Tuning  Prompt Tuning 
    * [ Overview  ](https://microsoft.github.io/graphrag/prompt_tuning/overview/)
    * [ Auto Tuning  ](https://microsoft.github.io/graphrag/prompt_tuning/auto_prompt_tuning/)
    * [ Manual Tuning  ](https://microsoft.github.io/graphrag/prompt_tuning/manual_prompt_tuning/)
  * Query  Query 
    * [ Overview  ](https://microsoft.github.io/graphrag/query/overview/)
    * [ Global Search  ](https://microsoft.github.io/graphrag/query/global_search/)
    * [ Local Search  ](https://microsoft.github.io/graphrag/query/local_search/)
    * [ DRIFT Search  ](https://microsoft.github.io/graphrag/query/drift_search/)
    * [ Question Generation  ](https://microsoft.github.io/graphrag/query/question_generation/)
    * Notebooks  Notebooks 
      * [ Overview  ](https://microsoft.github.io/graphrag/query/notebooks/overview/)
      * [ Global Search  ](https://microsoft.github.io/graphrag/examples_notebooks/global_search/)
      * [ Local Search  ](https://microsoft.github.io/graphrag/examples_notebooks/local_search/)
      * [ DRIFT Search  ](https://microsoft.github.io/graphrag/examples_notebooks/drift_search/)
  * Configuration  Configuration 
    * [ Overview  ](https://microsoft.github.io/graphrag/config/overview/)
    * [ Init Command  ](https://microsoft.github.io/graphrag/config/init/)
    * [ Detailed Configuration  ](https://microsoft.github.io/graphrag/config/yaml/)
    * [ Language Model Selection  ](https://microsoft.github.io/graphrag/config/models/)
  * [ CLI  ](https://microsoft.github.io/graphrag/cli/)
  * Extras  Extras 
    * [ Microsoft Research Blog  ](https://microsoft.github.io/graphrag/blog_posts/)
    * [ Visualization Guide  ](https://microsoft.github.io/graphrag/visualization_guide/)
    * Operation Dulce  Operation Dulce 
      * [ About  ](https://microsoft.github.io/graphrag/data/operation_dulce/ABOUT/)
      * [ Document  ](https://microsoft.github.io/graphrag/data/operation_dulce/Operation%20Dulce%20v2%201%201/)


Table of contents 
  * [ Get Started with GraphRAG 🚀  ](https://microsoft.github.io/graphrag/#get-started-with-graphrag)
  * [ GraphRAG vs Baseline RAG 🔍  ](https://microsoft.github.io/graphrag/#graphrag-vs-baseline-rag)
  * [ The GraphRAG Process 🤖  ](https://microsoft.github.io/graphrag/#the-graphrag-process)
    * [ Index  ](https://microsoft.github.io/graphrag/#index)
    * [ Query  ](https://microsoft.github.io/graphrag/#query)
    * [ Prompt Tuning  ](https://microsoft.github.io/graphrag/#prompt-tuning)
  * [ Versioning  ](https://microsoft.github.io/graphrag/#versioning)


# Welcome to GraphRAG
👉 [Microsoft Research Blog Post](https://www.microsoft.com/en-us/research/blog/graphrag-unlocking-llm-discovery-on-narrative-private-data/) 👉 [GraphRAG Arxiv](https://arxiv.org/pdf/2404.16130)
![Figure 1: LLM-generated knowledge graph built from a private dataset using GPT-4 Turbo.](https://microsoft.github.io/graphrag/img/GraphRag-Figure1.jpg)
Figure 1: An LLM-generated knowledge graph built using GPT-4 Turbo. 
GraphRAG is a structured, hierarchical approach to Retrieval Augmented Generation (RAG), as opposed to naive semantic-search approaches using plain text snippets. The GraphRAG process involves extracting a knowledge graph out of raw text, building a community hierarchy, generating summaries for these communities, and then leveraging these structures when perform RAG-based tasks.
To learn more about GraphRAG and how it can be used to enhance your language model's ability to reason about your private data, please visit the [Microsoft Research Blog Post](https://www.microsoft.com/en-us/research/blog/graphrag-unlocking-llm-discovery-on-narrative-private-data/).
## Get Started with GraphRAG 🚀
To start using GraphRAG, check out the [_Get Started_](https://microsoft.github.io/graphrag/get_started/) guide. For a deeper dive into the main sub-systems, please visit the docpages for the [Indexer](https://microsoft.github.io/graphrag/index/overview/) and [Query](https://microsoft.github.io/graphrag/query/overview/) packages.
## GraphRAG vs Baseline RAG 🔍
Retrieval-Augmented Generation (RAG) is a technique to improve LLM outputs using real-world information. This technique is an important part of most LLM-based tools and the majority of RAG approaches use vector similarity as the search technique, which we call _Baseline RAG_. GraphRAG uses knowledge graphs to provide substantial improvements in question-and-answer performance when reasoning about complex information. RAG techniques have shown promise in helping LLMs to reason about _private datasets_ - data that the LLM is not trained on and has never seen before, such as an enterprise’s proprietary research, business documents, or communications. _Baseline RAG_ was created to help solve this problem, but we observe situations where baseline RAG performs very poorly. For example:
  * Baseline RAG struggles to connect the dots. This happens when answering a question requires traversing disparate pieces of information through their shared attributes in order to provide new synthesized insights.
  * Baseline RAG performs poorly when being asked to holistically understand summarized semantic concepts over large data collections or even singular large documents.


To address this, the tech community is working to develop methods that extend and enhance RAG. Microsoft Research’s new approach, GraphRAG, creates a knowledge graph based on an input corpus. This graph, along with community summaries and graph machine learning outputs, are used to augment prompts at query time. GraphRAG shows substantial improvement in answering the two classes of questions described above, demonstrating intelligence or mastery that outperforms other approaches previously applied to private datasets.
## The GraphRAG Process 🤖
GraphRAG builds upon our prior [research](https://www.microsoft.com/en-us/worklab/patterns-hidden-inside-the-org-chart) and [tooling](https://github.com/graspologic-org/graspologic) using graph machine learning. The basic steps of the GraphRAG process are as follows:
### Index
  * Slice up an input corpus into a series of TextUnits, which act as analyzable units for the rest of the process, and provide fine-grained references in our outputs.
  * Extract all entities, relationships, and key claims from the TextUnits.
  * Perform a hierarchical clustering of the graph using the [Leiden technique](https://arxiv.org/pdf/1810.08473.pdf). To see this visually, check out Figure 1 above. Each circle is an entity (e.g., a person, place, or organization), with the size representing the degree of the entity, and the color representing its community.
  * Generate summaries of each community and its constituents from the bottom-up. This aids in holistic understanding of the dataset.


### Query
At query time, these structures are used to provide materials for the LLM context window when answering a question. The primary query modes are:
  * [_Global Search_](https://microsoft.github.io/graphrag/query/global_search/) for reasoning about holistic questions about the corpus by leveraging the community summaries.
  * [_Local Search_](https://microsoft.github.io/graphrag/query/local_search/) for reasoning about specific entities by fanning-out to their neighbors and associated concepts.
  * [_DRIFT Search_](https://microsoft.github.io/graphrag/query/drift_search/) for reasoning about specific entities by fanning-out to their neighbors and associated concepts, but with the added context of community information.


### Prompt Tuning
Using _GraphRAG_ with your data out of the box may not yield the best possible results. We strongly recommend to fine-tune your prompts following the [Prompt Tuning Guide](https://microsoft.github.io/graphrag/prompt_tuning/overview/) in our documentation.
## Versioning
Please see the [breaking changes](https://github.com/microsoft/graphrag/blob/main/breaking-changes.md) document for notes on our approach to versioning the project.
_Always run`graphrag init --root [path] --force` between minor version bumps to ensure you have the latest config format. Run the provided migration notebook between major version bumps if you want to avoid re-indexing prior datasets. Note that this will overwrite your configuration and prompts, so backup if necessary._
[ Next  Getting Started  ](https://microsoft.github.io/graphrag/get_started/)
© 2025 Microsoft | [Privacy](https://go.microsoft.com/fwlink/?LinkId=521839) | [Consumer Heath Privacy](https://go.microsoft.com/fwlink/?LinkId=2259814) | Cookies | [Terms of Use](https://go.microsoft.com/fwlink/?LinkID=206977) | [Trademarks](https://www.microsoft.com/trademarks)
Made with [ Material for MkDocs ](https://squidfunk.github.io/mkdocs-material/)


## Source Information
- URL: https://microsoft.github.io/graphrag/
- Harvested: 2025-08-20T13:07:43.385106
- Profile: deep_research
- Agent: architect
