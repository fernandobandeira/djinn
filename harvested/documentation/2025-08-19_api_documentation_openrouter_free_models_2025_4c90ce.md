---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T12:49:39.276236'
profile: api_documentation
source: https://openrouter.ai/docs
topic: OpenRouter_free_models_2025
---

# OpenRouter_free_models_2025

[![Logo](https://files.buildwithfern.com/openrouter.docs.buildwithfern.com/docs/2025-08-19T02:44:54.651Z/content/assets/logo.svg)![Logo](https://files.buildwithfern.com/openrouter.docs.buildwithfern.com/docs/2025-08-19T02:44:54.651Z/content/assets/logo-white.svg)](https://openrouter.ai/)
Search or ask AI a question`/`[API](https://openrouter.ai/docs/api-reference/overview)[Models](https://openrouter.ai/models)[Chat](https://openrouter.ai/chat)[Ranking](https://openrouter.ai/rankings)[Login](https://openrouter.ai/settings/credits)
  * Overview
    * [Quickstart](https://openrouter.ai/docs/quickstart)
    * [FAQ](https://openrouter.ai/docs/faq)
    * [Principles](https://openrouter.ai/docs/overview/principles)
    * [Models](https://openrouter.ai/docs/overview/models)
    * [Enterprise](https://openrouter.ai/enterprise)
  * Features
    * [Privacy and Logging](https://openrouter.ai/docs/features/privacy-and-logging)
    * [Model Routing](https://openrouter.ai/docs/features/model-routing)
    * [Provider Routing](https://openrouter.ai/docs/features/provider-routing)
    * [Latency and Performance](https://openrouter.ai/docs/features/latency-and-performance)
    * [Presets](https://openrouter.ai/docs/features/presets)
    * [Prompt Caching](https://openrouter.ai/docs/features/prompt-caching)
    * [Structured Outputs](https://openrouter.ai/docs/features/structured-outputs)
    * [Tool Calling](https://openrouter.ai/docs/features/tool-calling)
    * Multimodal
    * [Message Transforms](https://openrouter.ai/docs/features/message-transforms)
    * [Uptime Optimization](https://openrouter.ai/docs/features/uptime-optimization)
    * [Web Search](https://openrouter.ai/docs/features/web-search)
    * [Zero Completion Insurance](https://openrouter.ai/docs/features/zero-completion-insurance)
    * [Provisioning API Keys](https://openrouter.ai/docs/features/provisioning-api-keys)
    * [App Attribution](https://openrouter.ai/docs/app-attribution)
  * API Reference
    * [Overview](https://openrouter.ai/docs/api-reference/overview)
    * [Streaming](https://openrouter.ai/docs/api-reference/streaming)
    * [Limits](https://openrouter.ai/docs/api-reference/limits)
    * [Authentication](https://openrouter.ai/docs/api-reference/authentication)
    * [Parameters](https://openrouter.ai/docs/api-reference/parameters)
    * [Errors](https://openrouter.ai/docs/api-reference/errors)
    * [POSTCompletion](https://openrouter.ai/docs/api-reference/completion)
    * [POSTChat completion](https://openrouter.ai/docs/api-reference/chat-completion)
    * [GETGet a generation](https://openrouter.ai/docs/api-reference/get-a-generation)
    * [GETList available models](https://openrouter.ai/docs/api-reference/list-available-models)
    * [GETList endpoints for a model](https://openrouter.ai/docs/api-reference/list-endpoints-for-a-model)
    * [GETList models filtered by user provider preferences](https://openrouter.ai/docs/api-reference/list-models-filtered-by-user-provider-preferences)
    * [GETList available providers](https://openrouter.ai/docs/api-reference/list-available-providers)
    * [GETGet credits](https://openrouter.ai/docs/api-reference/get-credits)
    * [POSTCreate a Coinbase charge](https://openrouter.ai/docs/api-reference/create-a-coinbase-charge)
    * Authentication
    * API Keys
  * Use Cases
    * [BYOK](https://openrouter.ai/docs/use-cases/byok)
    * [Crypto API](https://openrouter.ai/docs/use-cases/crypto-api)
    * [OAuth PKCE](https://openrouter.ai/docs/use-cases/oauth-pkce)
    * [MCP Servers](https://openrouter.ai/docs/use-cases/mcp-servers)
    * [Organization Management](https://openrouter.ai/docs/use-cases/organization-management)
    * [For Providers](https://openrouter.ai/docs/use-cases/for-providers)
    * [Reasoning Tokens](https://openrouter.ai/docs/use-cases/reasoning-tokens)
    * [Usage Accounting](https://openrouter.ai/docs/use-cases/usage-accounting)
    * [User Tracking](https://openrouter.ai/docs/use-cases/user-tracking)
  * Community
    * [Frameworks Overview](https://openrouter.ai/docs/community/frameworks-overview)
    * [OpenAI SDK](https://openrouter.ai/docs/community/open-ai-sdk)
    * [LangChain](https://openrouter.ai/docs/community/lang-chain)
    * [PydanticAI](https://openrouter.ai/docs/community/pydantic-ai)
    * [Vercel AI SDK](https://openrouter.ai/docs/community/vercel-ai-sdk)
    * [Mastra](https://openrouter.ai/docs/community/mastra)
    * [Langfuse](https://openrouter.ai/docs/community/langfuse)
    * [Xcode](https://openrouter.ai/docs/community/xcode)
    * [Discord](https://discord.gg/fVyRaUDgxW)


[API](https://openrouter.ai/docs/api-reference/overview)[Models](https://openrouter.ai/models)[Chat](https://openrouter.ai/chat)[Ranking](https://openrouter.ai/rankings)[Login](https://openrouter.ai/settings/credits)
System
On this page
  * [Using the OpenAI SDK](https://openrouter.ai/docs/quickstart#using-the-openai-sdk)
  * [Using the OpenRouter API directly](https://openrouter.ai/docs/quickstart#using-the-openrouter-api-directly)
  * [Using third-party SDKs](https://openrouter.ai/docs/quickstart#using-third-party-sdks)


[Overview](https://openrouter.ai/docs/quickstart)
# Quickstart
Copy page
Get started with OpenRouter
OpenRouter provides a unified API that gives you access to hundreds of AI models through a single endpoint, while automatically handling fallbacks and selecting the most cost-effective options. Get started with just a few lines of code using your preferred SDK or framework.
Looking for information about free models and rate limits? Please see the [FAQ](https://openrouter.ai/docs/faq#how-are-rate-limits-calculated)
In the examples below, the OpenRouter-specific headers are optional. Setting them allows your app to appear on the OpenRouter leaderboards. For detailed information about app attribution, see our [App Attribution guide](https://openrouter.ai/docs/features/app-attribution).
## Using the OpenAI SDK
PythonTypeScript
```

1| from openai import OpenAI  
---|---  
2|   
3| client = OpenAI(  
4|  base_url="https://openrouter.ai/api/v1",  
5|  api_key="<OPENROUTER_API_KEY>",  
6| )  
7|   
8| completion = client.chat.completions.create(  
9|  extra_headers={  
10|   "HTTP-Referer": "<YOUR_SITE_URL>", # Optional. Site URL for rankings on openrouter.ai.  
11|   "X-Title": "<YOUR_SITE_NAME>", # Optional. Site title for rankings on openrouter.ai.  
12|  },  
13|  model="openai/gpt-4o",  
14|  messages=[  
15|   {  
16|    "role": "user",  
17|    "content": "What is the meaning of life?"  
18|   }  
19|  ]  
20| )  
21|   
22| print(completion.choices[0].message.content)  

```

## Using the OpenRouter API directly
You can use the interactive [Request Builder](https://openrouter.ai/request-builder) to generate OpenRouter API requests in the language of your choice.
PythonTypeScriptShell
```

1| import requests  
---|---  
2| import json  
3|   
4| response = requests.post(  
5|  url="https://openrouter.ai/api/v1/chat/completions",  
6|  headers={  
7|   "Authorization": "Bearer <OPENROUTER_API_KEY>",  
8|   "HTTP-Referer": "<YOUR_SITE_URL>", # Optional. Site URL for rankings on openrouter.ai.  
9|   "X-Title": "<YOUR_SITE_NAME>", # Optional. Site title for rankings on openrouter.ai.  
10|  },  
11|  data=json.dumps({  
12|   "model": "openai/gpt-4o", # Optional  
13|   "messages": [  
14|    {  
15|     "role": "user",  
16|     "content": "What is the meaning of life?"  
17|    }  
18|   ]  
19|  })  
20| )  

```

The API also supports [streaming](https://openrouter.ai/docs/api-reference/streaming).
## Using third-party SDKs
For information about using third-party SDKs and frameworks with OpenRouter, please [see our frameworks documentation.](https://openrouter.ai/docs/community/frameworks-overview)
Was this page helpful?
YesNo
#### [Frequently Asked QuestionsCommon questions about OpenRouterNext](https://openrouter.ai/docs/faq)[Built with](https://buildwithfern.com/?utm_campaign=buildWith&utm_medium=docs&utm_source=openrouter.ai)


## Source Information
- URL: https://openrouter.ai/docs
- Harvested: 2025-08-19T12:49:39.276236
- Profile: api_documentation
- Agent: architect
