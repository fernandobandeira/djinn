---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-20T13:12:19.256895'
profile: deep_research
source: https://docs.trychroma.com/guides/embeddings
topic: Knowledge Base Best Practices - Embedding and Vector Storage Patterns
---

# Knowledge Base Best Practices - Embedding and Vector Storage Patterns

[](https://docs.trychroma.com/)Search`K`
[921 online](https://discord.gg/MMeYNTmh3x)[22k](https://github.com/chroma-core/chroma)[22.7k](https://x.com/trychroma)[](https://www.youtube.com/@trychroma/featured)Toggle theme
[Docs](https://docs.trychroma.com/docs//overview/introduction)[Chroma Cloud](https://docs.trychroma.com/cloud/getting-started)[Guides](https://docs.trychroma.com/guides//build/building-with-ai)[Integrations](https://docs.trychroma.com/integrations/chroma-integrations)[Reference](https://docs.trychroma.com/reference/chroma-reference)
Overview
[Introduction](https://docs.trychroma.com/docs/overview/introduction)
[Getting Started](https://docs.trychroma.com/docs/overview/getting-started)
[Architecture](https://docs.trychroma.com/docs/overview/architecture)
[Data Model](https://docs.trychroma.com/docs/overview/data-model)
[Roadmap](https://docs.trychroma.com/docs/overview/roadmap)
[Contributing](https://docs.trychroma.com/docs/overview/contributing)
[Telemetry](https://docs.trychroma.com/docs/overview/telemetry)
[Migration](https://docs.trychroma.com/docs/overview/migration)
[Troubleshooting](https://docs.trychroma.com/docs/overview/troubleshooting)
[About](https://docs.trychroma.com/docs/overview/about)
Run Chroma
[Ephemeral Client](https://docs.trychroma.com/docs/run-chroma/ephemeral-client)
[Persistent Client](https://docs.trychroma.com/docs/run-chroma/persistent-client)
[Client-Server Mode](https://docs.trychroma.com/docs/run-chroma/client-server)
[Cloud Client](https://docs.trychroma.com/docs/run-chroma/cloud-client)
Collections
[Manage Collections](https://docs.trychroma.com/docs/collections/manage-collections)
[Configure](https://docs.trychroma.com/docs/collections/configure)
[Add Data](https://docs.trychroma.com/docs/collections/add-data)
[Update Data](https://docs.trychroma.com/docs/collections/update-data)
[Delete Data](https://docs.trychroma.com/docs/collections/delete-data)
Querying Collections
[Query And Get](https://docs.trychroma.com/docs/querying-collections/query-and-get)
[Metadata Filtering](https://docs.trychroma.com/docs/querying-collections/metadata-filtering)
[Full Text Search and Regex](https://docs.trychroma.com/docs/querying-collections/full-text-search)
Embeddings
[Embedding Functions](https://docs.trychroma.com/docs/embeddings/embedding-functions)
[Multimodal](https://docs.trychroma.com/docs/embeddings/multimodal)
CLI
[Installing the CLI](https://docs.trychroma.com/docs/cli/install)
[Browse Collections](https://docs.trychroma.com/docs/cli/browse)
[Copy Collections](https://docs.trychroma.com/docs/cli/copy)
[DB Management](https://docs.trychroma.com/docs/cli/db)
[Install Sample Apps](https://docs.trychroma.com/docs/cli/sample-apps)
[Login](https://docs.trychroma.com/docs/cli/login)
[Profile Management](https://docs.trychroma.com/docs/cli/profile)
[Run a Chroma Server](https://docs.trychroma.com/docs/cli/run)
[Update the CLI](https://docs.trychroma.com/docs/cli/update)
[Vacuum](https://docs.trychroma.com/docs/cli/vacuum)
w-5 h-5
# Embedding Functions
Embeddings are the way to represent any kind of data, making them the perfect fit for working with all kinds of AI-powered tools and algorithms. They can represent text, images, and soon audio and video. Chroma collections index embeddings to enable efficient similarity search on the data they represent. There are many options for creating embeddings, whether locally using an installed library, or by calling an API.
Chroma provides lightweight wrappers around popular embedding providers, making it easy to use them in your apps. You can set an embedding function when you [create](https://docs.trychroma.com/docs/collections/manage-collections) a Chroma collection, to be automatically used when adding and querying data, or you can call them directly yourself.
| Python| Typescript  
---|---|---  
[OpenAI](https://docs.trychroma.com/integrations/embedding-models/openai)| ✓| ✓  
[Google Generative AI](https://docs.trychroma.com/integrations/embedding-models/google-gemini)| ✓| ✓  
[Cohere](https://docs.trychroma.com/integrations/embedding-models/cohere)| ✓| ✓  
[Hugging Face](https://docs.trychroma.com/integrations/embedding-models/hugging-face)| ✓| -  
[Instructor](https://docs.trychroma.com/integrations/embedding-models/instructor)| ✓| -  
[Hugging Face Embedding Server](https://docs.trychroma.com/integrations/embedding-models/hugging-face-server)| ✓| ✓  
[Jina AI](https://docs.trychroma.com/integrations/embedding-models/jina-ai)| ✓| ✓  
[Cloudflare Workers AI](https://docs.trychroma.com/integrations/embedding-models/cloudflare-workers-ai)| ✓| ✓  
[Together AI](https://docs.trychroma.com/integrations/embedding-models/together-ai)| ✓| ✓  
[Mistral](https://docs.trychroma.com/integrations/embedding-models/mistral)| ✓| ✓  
[Morph](https://docs.trychroma.com/integrations/embedding-models/morph)| ✓| ✓  
We welcome pull requests to add new Embedding Functions to the community.
## Default: all-MiniLM-L6-v2[#](https://docs.trychroma.com/docs/embeddings/embedding-functions#default-all-minilm-l6-v2)
Chroma's default embedding function uses the [Sentence Transformers](https://www.sbert.net/) all-MiniLM-L6-v2 model to create embeddings. This embedding model can create sentence and document embeddings that can be used for a wide variety of tasks. This embedding function runs locally on your machine, and may require you download the model files (this will happen automatically).
If you don't specify an embedding function when creating a collection, Chroma will set it to be the DefaultEmbeddingFunction:
PythonTypescript
Python
```

collection = client.create_collection(name="my_collection")


```

## Using Embedding Functions[#](https://docs.trychroma.com/docs/embeddings/embedding-functions#using-embedding-functions)
Embedding functions can be linked to a collection and used whenever you call add, update, upsert or query.
PythonTypescript
Python
```

# Set your OPENAI_API_KEY environment variable
from chromadb.utils.embedding_functions import OpenAIEmbeddingFunction
collection = client.create_collection(
  name="my_collection",
  embedding_function=OpenAIEmbeddingFunction(
    model_name="text-embedding-3-small"
  )
)
# Chroma will use OpenAIEmbeddingFunction to embed your documents
collection.add(
  ids=["id1", "id2"],
  documents=["doc1", "doc2"]
)


```

You can also use embedding functions directly which can be handy for debugging.
PythonTypescript
```

from chromadb.utils.embedding_functions import DefaultEmbeddingFunction
default_ef = DefaultEmbeddingFunction()
embeddings = default_ef(["foo"])
print(embeddings) # [[0.05035809800028801, 0.0626462921500206, -0.061827320605516434...]]
collection.query(query_embeddings=embeddings)


```

## Custom Embedding Functions[#](https://docs.trychroma.com/docs/embeddings/embedding-functions#custom-embedding-functions)
You can create your own embedding function to use with Chroma; it just needs to implement EmbeddingFunction.
PythonTypescript
```

from chromadb import Documents, EmbeddingFunction, Embeddings
class MyEmbeddingFunction(EmbeddingFunction):
  def __call__(self, input: Documents) -> Embeddings:
    # embed the documents somehow
    return embeddings


```

We welcome contributions! If you create an embedding function that you think would be useful to others, please consider [submitting a pull request](https://github.com/chroma-core/chroma).
[Full Text Search and Regex](https://docs.trychroma.com/docs/querying-collections/full-text-search)[Multimodal](https://docs.trychroma.com/docs/embeddings/multimodal)
[Edit this page on GitHub](https://github.com/chroma-core/chroma/tree/main/docs/docs.trychroma.com/markdoc/content/docs/embeddings/embedding-functions.md)
Ask this Page
Ask this Page
On this page
[Embedding Functions](https://docs.trychroma.com/docs/embeddings/embedding-functions#embedding-functions)
[Default: all-MiniLM-L6-v2](https://docs.trychroma.com/docs/embeddings/embedding-functions#default-all-minilm-l6-v2)
[Using Embedding Functions](https://docs.trychroma.com/docs/embeddings/embedding-functions#using-embedding-functions)
[Custom Embedding Functions](https://docs.trychroma.com/docs/embeddings/embedding-functions#custom-embedding-functions)
NEW
Chroma Cloud
Our fully managed hosted service, Chroma Cloud is here.
[Sign up →](https://www.trychroma.com/signup)
![Cloud Art](https://docs.trychroma.com/_next/image?url=%2Fcloud-art.jpg&w=256&q=75)


## Source Information
- URL: https://docs.trychroma.com/guides/embeddings
- Harvested: 2025-08-20T13:12:19.256895
- Profile: deep_research
- Agent: architect
