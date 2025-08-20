---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-20T13:13:32.699238'
profile: deep_research
source: https://docs.trychroma.com/deployment
topic: ChromaDB Production Deployment
---

# ChromaDB Production Deployment

[](https://docs.trychroma.com/)Search`K`
[918 online](https://discord.gg/MMeYNTmh3x)[22k](https://github.com/chroma-core/chroma)[22.7k](https://x.com/trychroma)[](https://www.youtube.com/@trychroma/featured)Toggle theme
[Docs](https://docs.trychroma.com/docs//overview/introduction)[Chroma Cloud](https://docs.trychroma.com/cloud/getting-started)[Guides](https://docs.trychroma.com/guides//build/building-with-ai)[Integrations](https://docs.trychroma.com/integrations/chroma-integrations)[Reference](https://docs.trychroma.com/reference/chroma-reference)
Build
[Building With AI](https://docs.trychroma.com/guides/build/building-with-ai)
[Introduction to Retrieval](https://docs.trychroma.com/guides/build/intro-to-retrieval)
Deploy
[Client Server Mode](https://docs.trychroma.com/guides/deploy/client-server-mode)
[Python Thin Client](https://docs.trychroma.com/guides/deploy/python-thin-client)
[Performance](https://docs.trychroma.com/guides/deploy/performance)
[Observability](https://docs.trychroma.com/guides/deploy/observability)
[Docker](https://docs.trychroma.com/guides/deploy/docker)
[AWS](https://docs.trychroma.com/guides/deploy/aws)
[Azure](https://docs.trychroma.com/guides/deploy/azure)
[GCP](https://docs.trychroma.com/guides/deploy/gcp)
w-5 h-5
# Running Chroma in Client-Server Mode
PythonTypescript
Chroma can also be configured to run in client/server mode. In this mode, the Chroma client connects to a Chroma server running in a separate process.
This means that you can deploy single-node Chroma to a [Docker container](https://docs.trychroma.com/guides/deploy/docker), or a machine hosted by a cloud provider like [AWS](https://docs.trychroma.com/guides/deploy/aws), [GCP](https://docs.trychroma.com/guides/deploy/gcp), [Azure](https://docs.trychroma.com/guides/deploy/azure), and others. Then, you can access your Chroma server from your application using our HttpClient.
You can quickly experiment locally with Chroma in client/server mode by using our CLI:
Terminal
```

chroma run --path /db_path


```

Then use the Chroma HttpClient to connect to the server:
Python
```

import chromadb
chroma_client = chromadb.HttpClient(host='localhost', port=8000)


```

Chroma also provides an AsyncHttpClient. The behaviors and method signatures are identical to the synchronous client, but all methods that would block are now async:
Python
```

import asyncio
import chromadb
async def main():
  client = await chromadb.AsyncHttpClient()
  collection = await client.create_collection(name="my_collection")
  await collection.add(
    documents=["hello world"],
    ids=["id1"]
  )
asyncio.run(main())


```

If you intend to deploy your Chroma server, you may want to consider our [thin-client package](https://docs.trychroma.com/production/chroma-server/python-thin-client) for client-side interactions.
[Introduction to Retrieval](https://docs.trychroma.com/guides/build/intro-to-retrieval)[Python Thin Client](https://docs.trychroma.com/guides/deploy/python-thin-client)
[Edit this page on GitHub](https://github.com/chroma-core/chroma/tree/main/docs/docs.trychroma.com/markdoc/content/guides/deploy/client-server-mode.md)
Ask this Page
Ask this Page
On this page
[Running Chroma in Client-Server Mode](https://docs.trychroma.com/guides/deploy/client-server-mode#running-chroma-in-client-server-mode)
NEW
Chroma Cloud
Our fully managed hosted service, Chroma Cloud is here.
[Sign up →](https://www.trychroma.com/signup)
![Cloud Art](https://docs.trychroma.com/_next/image?url=%2Fcloud-art.jpg&w=256&q=75)


## Source Information
- URL: https://docs.trychroma.com/deployment
- Harvested: 2025-08-20T13:13:32.699238
- Profile: deep_research
- Agent: architect
