---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-20T13:12:07.678964'
profile: deep_research
source: https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/
topic: Knowledge Base Best Practices - Document Processing and Chunking Implementation
---

# Knowledge Base Best Practices - Document Processing and Chunking Implementation

[ Skip to content ](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/#defining-and-customizing-documents)
[ ![logo](https://docs.llamaindex.ai/en/stable/_static/assets/LlamaSquareBlack.svg) ](https://docs.llamaindex.ai/en/stable/ "LlamaIndex")
LlamaIndex 
Using Documents 
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
  * [ Learn  ](https://docs.llamaindex.ai/en/stable/understanding/)
  * [ Use Cases  ](https://docs.llamaindex.ai/en/stable/use_cases/)
  * [ Examples  ](https://docs.llamaindex.ai/en/stable/examples/)
  * [ Component Guides  ](https://docs.llamaindex.ai/en/stable/module_guides/)
Component Guides 
    * [ Models  ](https://docs.llamaindex.ai/en/stable/module_guides/models/)
    * [ Prompts  ](https://docs.llamaindex.ai/en/stable/module_guides/models/prompts/)
    * [ Loading  ](https://docs.llamaindex.ai/en/stable/module_guides/loading/)
Loading 
      * [ Documents and Nodes  ](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/)
Documents and Nodes 
        * Using Documents  [ Using Documents  ](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/) Table of contents 
          * [ Defining Documents  ](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/#defining-documents)
          * [ Customizing Documents  ](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/#customizing-documents)
            * [ Metadata  ](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/#metadata)
            * [ Customizing the id  ](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/#customizing-the-id)
            * [ Advanced - Metadata Customization  ](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/#advanced-metadata-customization)
              * [ Customizing LLM Metadata Text  ](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/#customizing-llm-metadata-text)
              * [ Customizing Embedding Metadata Text  ](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/#customizing-embedding-metadata-text)
              * [ Customizing Metadata Format  ](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/#customizing-metadata-format)
            * [ Summary  ](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/#summary)
            * [ Advanced - Automatic Metadata Extraction  ](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/#advanced-automatic-metadata-extraction)
        * [ Using Nodes  ](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_nodes/)
        * [ Metadata Extraction  ](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_metadata_extractor/)
      * [ SimpleDirectoryReader  ](https://docs.llamaindex.ai/en/stable/module_guides/loading/simpledirectoryreader/)
      * [ Data Connectors  ](https://docs.llamaindex.ai/en/stable/module_guides/loading/connector/)
      * [ Node Parsers / Text Splitters  ](https://docs.llamaindex.ai/en/stable/module_guides/loading/node_parsers/)
      * [ Ingestion Pipeline  ](https://docs.llamaindex.ai/en/stable/module_guides/loading/ingestion_pipeline/)
    * [ Indexing  ](https://docs.llamaindex.ai/en/stable/module_guides/indexing/)
    * [ Storing  ](https://docs.llamaindex.ai/en/stable/module_guides/storing/)
    * [ Querying  ](https://docs.llamaindex.ai/en/stable/module_guides/querying/)
    * [ Agents  ](https://docs.llamaindex.ai/en/stable/module_guides/deploying/agents/)
    * [ Workflows  ](https://docs.llamaindex.ai/en/stable/workflows/)
    * [ MCP  ](https://docs.llamaindex.ai/en/stable/module_guides/mcp/)
    * [ Evaluation  ](https://docs.llamaindex.ai/en/stable/module_guides/evaluating/)
    * [ Observability  ](https://docs.llamaindex.ai/en/stable/module_guides/observability/)
    * [ Settings  ](https://docs.llamaindex.ai/en/stable/module_guides/supporting_modules/settings/)
    * [ Llama Deploy  ](https://docs.llamaindex.ai/en/stable/module_guides/llama_deploy/10_getting_started/)
  * [ Advanced Topics  ](https://docs.llamaindex.ai/en/stable/optimizing/production_rag/)
  * [ API Reference  ](https://docs.llamaindex.ai/en/stable/api_reference/)
  * [ Open-Source Community  ](https://docs.llamaindex.ai/en/stable/community/integrations/)
  * [ Workflows  ](https://docs.llamaindex.ai/en/stable/workflows/)
  * [ LlamaCloud  ](https://docs.cloud.llamaindex.ai/)


Table of contents 
  * [ Defining Documents  ](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/#defining-documents)
  * [ Customizing Documents  ](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/#customizing-documents)
    * [ Metadata  ](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/#metadata)
    * [ Customizing the id  ](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/#customizing-the-id)
    * [ Advanced - Metadata Customization  ](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/#advanced-metadata-customization)
      * [ Customizing LLM Metadata Text  ](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/#customizing-llm-metadata-text)
      * [ Customizing Embedding Metadata Text  ](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/#customizing-embedding-metadata-text)
      * [ Customizing Metadata Format  ](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/#customizing-metadata-format)
    * [ Summary  ](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/#summary)
    * [ Advanced - Automatic Metadata Extraction  ](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/#advanced-automatic-metadata-extraction)


# Defining and Customizing Documents[#](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/#defining-and-customizing-documents "Permanent link")
## Defining Documents[#](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/#defining-documents "Permanent link")
Documents can either be created automatically via data loaders, or constructed manually.
By default, all of our [data loaders](https://docs.llamaindex.ai/en/stable/module_guides/loading/connector/) (including those offered on LlamaHub) return `Document` objects through the `load_data` function.
```
from llama_index.core import SimpleDirectoryReader
documents = SimpleDirectoryReader("./data").load_data()

```

You can also choose to construct documents manually. LlamaIndex exposes the `Document` struct.
```
from llama_index.core import Document
text_list = [text1, text2, ...]
documents = [Document(text=t) for t in text_list]

```

To speed up prototyping and development, you can also quickly create a document using some default text:
```
document = Document.example()

```

## Customizing Documents[#](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/#customizing-documents "Permanent link")
This section covers various ways to customize `Document` objects. Since the `Document` object is a subclass of our `TextNode` object, all these settings and details apply to the `TextNode` object class as well.
### Metadata[#](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/#metadata "Permanent link")
Documents also offer the chance to include useful metadata. Using the `metadata` dictionary on each document, additional information can be included to help inform responses and track down sources for query responses. This information can be anything, such as filenames or categories. If you are integrating with a vector database, keep in mind that some vector databases require that the keys must be strings, and the values must be flat (either `str`, `float`, or `int`).
Any information set in the `metadata` dictionary of each document will show up in the `metadata` of each source node created from the document. Additionally, this information is included in the nodes, enabling the index to utilize it on queries and responses. By default, the metadata is injected into the text for both embedding and LLM model calls.
There are a few ways to set up this dictionary:
  1. In the document constructor:


```
document = Document(
  text="text",
  metadata={"filename": "<doc_file_name>", "category": "<category>"},
)

```

  1. After the document is created:


```
document.metadata = {"filename": "<doc_file_name>"}

```

  1. Set the filename automatically using the `SimpleDirectoryReader` and `file_metadata` hook. This will automatically run the hook on each document to set the `metadata` field:


```
from llama_index.core import SimpleDirectoryReader
filename_fn = lambda filename: {"file_name": filename}
# automatically sets the metadata of each document according to filename_fn
documents = SimpleDirectoryReader(
  "./data", file_metadata=filename_fn
).load_data()

```

### Customizing the id[#](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/#customizing-the-id "Permanent link")
As detailed in the section [Document Management](https://docs.llamaindex.ai/en/stable/module_guides/indexing/document_management/), the `doc_id` is used to enable efficient refreshing of documents in the index. When using the `SimpleDirectoryReader`, you can automatically set the doc `doc_id` to be the full path to each document:
```
from llama_index.core import SimpleDirectoryReader
documents = SimpleDirectoryReader("./data", filename_as_id=True).load_data()
print([x.doc_id for x in documents])

```

You can also set the `doc_id` of any `Document` directly!
```
document.doc_id = "My new document id!"

```

Note: the ID can also be set through the `node_id` or `id_` property on a Document object, similar to a `TextNode` object.
### Advanced - Metadata Customization[#](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/#advanced-metadata-customization "Permanent link")
A key detail mentioned above is that by default, any metadata you set is included in the embeddings generation and LLM.
#### Customizing LLM Metadata Text[#](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/#customizing-llm-metadata-text "Permanent link")
Typically, a document might have many metadata keys, but you might not want all of them visible to the LLM during response synthesis. In the above examples, we may not want the LLM to read the `file_name` of our document. However, the `file_name` might include information that will help generate better embeddings. A key advantage of doing this is to bias the embeddings for retrieval without changing what the LLM ends up reading.
We can exclude it like so:
```
document.excluded_llm_metadata_keys = ["file_name"]

```

Then, we can test what the LLM will actually end up reading using the `get_content()` function and specifying `MetadataMode.LLM`:
```
from llama_index.core.schema import MetadataMode
print(document.get_content(metadata_mode=MetadataMode.LLM))

```

#### Customizing Embedding Metadata Text[#](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/#customizing-embedding-metadata-text "Permanent link")
Similar to customizing the metadata visible to the LLM, we can also customize the metadata visible to embeddings. In this case, you can specifically exclude metadata visible to the embedding model, in case you DON'T want particular text to bias the embeddings.
```
document.excluded_embed_metadata_keys = ["file_name"]

```

Then, we can test what the embedding model will actually end up reading using the `get_content()` function and specifying `MetadataMode.EMBED`:
```
from llama_index.core.schema import MetadataMode
print(document.get_content(metadata_mode=MetadataMode.EMBED))

```

#### Customizing Metadata Format[#](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/#customizing-metadata-format "Permanent link")
As you know by now, metadata is injected into the actual text of each document/node when sent to the LLM or embedding model. By default, the format of this metadata is controlled by three attributes:
  1. `Document.metadata_seperator` -> default = `"\n"`


When concatenating all key/value fields of your metadata, this field controls the separator between each key/value pair.
  1. `Document.metadata_template` -> default = `"{key}: {value}"`


This attribute controls how each key/value pair in your metadata is formatted. The two variables `key` and `value` string keys are required.
  1. `Document.text_template` -> default = `{metadata_str}\n\n{content}`


Once your metadata is converted into a string using `metadata_seperator` and `metadata_template`, this templates controls what that metadata looks like when joined with the text content of your document/node. The `metadata` and `content` string keys are required.
### Summary[#](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/#summary "Permanent link")
Knowing all this, let's create a short example using all this power:
```
from llama_index.core import Document
from llama_index.core.schema import MetadataMode
document = Document(
  text="This is a super-customized document",
  metadata={
    "file_name": "super_secret_document.txt",
    "category": "finance",
    "author": "LlamaIndex",
  },
  excluded_llm_metadata_keys=["file_name"],
  metadata_seperator="::",
  metadata_template="{key}=>{value}",
  text_template="Metadata: {metadata_str}\n-----\nContent: {content}",
)
print(
  "The LLM sees this: \n",
  document.get_content(metadata_mode=MetadataMode.LLM),
)
print(
  "The Embedding model sees this: \n",
  document.get_content(metadata_mode=MetadataMode.EMBED),
)

```

### Advanced - Automatic Metadata Extraction[#](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/#advanced-automatic-metadata-extraction "Permanent link")
We have [initial examples](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_metadata_extractor/) of using LLMs themselves to perform metadata extraction.
Back to top  [ Previous  Documents / Nodes  ](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/) [ Next  Using Nodes  ](https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_nodes/)


## Source Information
- URL: https://docs.llamaindex.ai/en/stable/module_guides/loading/documents_and_nodes/usage_documents/
- Harvested: 2025-08-20T13:12:07.678964
- Profile: deep_research
- Agent: architect
