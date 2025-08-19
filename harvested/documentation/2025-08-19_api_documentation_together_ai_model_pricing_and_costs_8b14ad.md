---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T12:38:25.557461'
profile: api_documentation
source: https://docs.together.ai/pricing
topic: Together AI model pricing and costs
---

# Together AI model pricing and costs

[Together.ai Docs home page![light logo](https://mintlify.s3.us-west-1.amazonaws.com/togetherai-52386018/logo/logo.svg)![dark logo](https://mintlify.s3.us-west-1.amazonaws.com/togetherai-52386018/logo/logo-dark.svg)](https://docs.together.ai/)
Search...
Ctrl KAsk AI
  * [Status](https://status.together.ai/)
  * [Playground](https://api.together.xyz/playground)
  * [Playground](https://api.together.xyz/playground)


Search...
Navigation
Getting Started
Introduction
[Documentation](https://docs.together.ai/docs)[API Reference](https://docs.together.ai/reference/chat-completions-1)
##### Getting Started
  * [Introduction](https://docs.together.ai/docs)
  * [Quickstart](https://docs.together.ai/docs/quickstart)
  * [OpenAI Compatibility](https://docs.together.ai/docs/openai-api-compatibility)


##### Inference
  * [Serverless models](https://docs.together.ai/docs/serverless-models)
  * Dedicated Models
  * [Batch Inference](https://docs.together.ai/docs/batch-inference)
  * DeepSeek-R1 Quickstart
  * [Llama 4 Quickstart](https://docs.together.ai/docs/llama4-quickstart)
  * [Kimi K2 QuickStart](https://docs.together.ai/docs/kimi-k2-quickstart)
  * [OpenAI GPT-OSS Quickstart](https://docs.together.ai/docs/gpt-oss)


##### Capabilities
  * [Chat](https://docs.together.ai/docs/chat-overview)
  * [Structured Outputs](https://docs.together.ai/docs/json-mode)
  * [Function Calling](https://docs.together.ai/docs/function-calling)
  * Images
  * [Vision](https://docs.together.ai/docs/vision-overview)
  * Code Execution
  * [Speech-to-Text](https://docs.together.ai/docs/speech-to-text)
  * Evaluations
  * Other Modalities


##### Examples
  * Agent Workflows
  * [Together Cookbooks](https://github.com/togethercomputer/together-cookbook)
  * [Example Apps](https://together.ai/demos?_gl=1*1v9dt5q*_gcl_au*OTY1NDQ1MTk3LjE3NDMwMzg1NDU.)
  * [Integrations](https://docs.together.ai/docs/integrations)


##### Training
  * Fine-tuning
  * GPU Clusters


##### Guides
  * [How to Build Coding Agents](https://docs.together.ai/docs/how-to-build-coding-agents)
  * [Getting Started with Logprobs](https://docs.together.ai/docs/logprobs)
  * [Quickstart: Retrieval Augmented Generation (RAG)](https://docs.together.ai/docs/quickstart-retrieval-augmented-generation-rag)
  * [Quickstart: Next.Js](https://docs.together.ai/docs/nextjs-chat-quickstart)
  * [Quickstart: Using Hugging Face Inference With Together](https://docs.together.ai/docs/quickstart-using-hugging-face-inference)
  * [Quickstart: Using Vercel'S AI SDK With Together AI](https://docs.together.ai/docs/using-together-with-vercels-ai-sdk)
  * [Together Mixture Of Agents (MoA)](https://docs.together.ai/docs/mixture-of-agents)
  * [Quickstart: How to do OCR](https://docs.together.ai/docs/quickstart-how-to-do-ocr)
  * Search & RAG
  * Apps
  * [How to use Cline with DeepSeek V3 to build faster](https://docs.together.ai/docs/how-to-use-cline)


##### ❓ Frequently Asked Questions
  * [Deployment Options](https://docs.together.ai/docs/deployment-options)
  * [Rate Limits](https://docs.together.ai/docs/rate-limits)
  * [Error Codes](https://docs.together.ai/docs/error-codes)
  * [Deploy Dedicated Endpoints In The Web](https://docs.together.ai/docs/dedicated-endpoints-ui)
  * Priority Support
  * [Deprecations](https://docs.together.ai/docs/deprecations)
  * [Playground](https://docs.together.ai/docs/inference-web-interface)
  * [Inference FAQs](https://docs.together.ai/docs/inference-faqs)
  * [Fine Tuning FAQs](https://docs.together.ai/docs/fine-tuning-faqs)
  * [Multiple API Keys](https://docs.together.ai/docs/multiple-api-keys)
  * [Dedicated Endpoints](https://docs.together.ai/docs/dedicated-endpoints-1)


On this page
  * [Quickstart](https://docs.together.ai/docs#quickstart)
  * [Which model should I use?](https://docs.together.ai/docs#which-model-should-i-use%3F)
  * [Together Cookbook](https://docs.together.ai/docs#together-cookbook)
  * [Example apps](https://docs.together.ai/docs#example-apps)
  * [Next steps](https://docs.together.ai/docs#next-steps)
  * [Resources](https://docs.together.ai/docs#resources)


Getting Started
# Introduction
Copy page
Introduction to Together AI and all its services.
Copy page
Welcome to the Together AI docs! Together AI makes it easy to run or fine-tune leading open source models with only a few lines of code. We offer a variety of generative AI services:
  * **Serverless models** - Use our API or [playground](https://api.together.xyz/playground) to run dozens of models with pay as you go pricing.
  * **Fine-Tuning** - Fine-tune models on your own data [in 5 minutes](https://docs.together.ai/docs/fine-tuning-overview), then run the model for inference.
  * **Dedicated endpoints** - Run models [on your own private GPUs](https://docs.together.ai/docs/dedicated-models), starting at a one month minimum commitment.
  * **GPU Clusters** - If you’re interested in private, state of the art clusters with H100 GPUs, [contact us](https://www.together.ai/forms/gpu-cluster-requests).


## 
[​](https://docs.together.ai/docs#quickstart)
Quickstart
See our [full quickstart](https://docs.together.ai/docs/quickstart) for how to get started with our API in 1 minute.
Python
TypeScript
cURL
Copy
Ask AI
```
from together import Together
client = Together()
completion = client.chat.completions.create(
 model="meta-llama/Meta-Llama-3.1-8B-Instruct-Turbo",
 messages=[{"role": "user", "content": "What are the top 3 things to do in New York?"}],
)
print(completion.choices[0].message.content)

```

## 
[​](https://docs.together.ai/docs#which-model-should-i-use%3F)
Which model should I use?
Together hosts many popular models via our serverless endpoints. For each of these, you’ll [be charged](https://together.ai/pricing) based on the tokens you use and size of the model. Here are all the different types of models that we support:
  * [Chat models](https://docs.together.ai/docs/serverless-models#chat-models)
  * [Vision models](https://docs.together.ai/docs/serverless-models#vision-models)
  * [Image models](https://docs.together.ai/docs/serverless-models#image-models)
  * [Audio Models](https://docs.together.ai/docs/serverless-models#audio-models)
  * [Embedding models](https://docs.together.ai/docs/serverless-models#embedding-models)
  * [Rerank Models](https://docs.together.ai/docs/serverless-models#rerank-models)
  * [Language and code models](https://docs.together.ai/docs/serverless-models#language-models)

Don’t see a model you want to use? **[Send us a request](https://www.together.ai/forms/model-requests)** to add or upvote the model you’d love to see us add to our serverless infrastructure.
## 
[​](https://docs.together.ai/docs#together-cookbook)
Together Cookbook
See the **[Together Cookbook](https://github.com/togethercomputer/together-cookbook)** – a collection of notebooks showcasing use cases of open-source models with Together AI. Examples include RAG (text + multimodal), Semantic Search, Rerankers, & Structured JSON extraction.
## 
[​](https://docs.together.ai/docs#example-apps)
Example apps
We’ve built a number of full-stack open source example apps that you can reference. These are production-ready apps have over 500k users & 10k GitHub stars combined – all fully open source and built on Together AI.
  * [LlamaCoder](https://llamacoder.together.ai/) ([GitHub](https://github.com/Nutlope/llamacoder)) – an OSS Claude artifacts that is able to generate full React apps from a single prompt. Built on Llama 3.1 405B powered by Together inference.
  * [BlinkShot](https://www.blinkshot.io/) ([GitHub](https://github.com/Nutlope/blinkshot)) – a realtime AI image generator using Flux Schnell on Together AI. Type in a prompt and images will get generated as you type.
  * [TurboSeek](https://www.turboseek.io/) ([GitHub](https://github.com/Nutlope/turboseek)) – an AI search engine inspired by Perplexity. It uses a search API (Serper) along with an LLM (Mixtral) to be able to answer any questions.
  * [Napkins.dev](https://www.napkins.dev/) ([GitHub](https://github.com/nutlope/napkins)) – a wireframe to app tool. It uses Llama 3.2 vision to read in screenshots and write code for them using Llama 3.1 405B.
  * [PDFToChat](https://www.pdftochat.com/) ([GitHub](https://github.com/nutlope/pdftochat)) – a site that lets you chat with your PDFs. Uses RAG with Together embeddings, inference with Llama 3, authentication with Clerk, & MongoDB/Pinecone for the vector database.
  * [LlamaTutor](https://llamatutor.together.ai/) ([GitHub](https://github.com/Nutlope/llamatutor)) – a personal tutor that can explain any topic at any education level by using a search API along with Llama 3.1.
  * [NotesGPT](https://usenotesgpt.com/) ([GitHub](https://github.com/nutlope/notesgpt)) – an AI note taker that converts your voice notes into organized summaries and clear action items using AI. Uses Together inference (Mixtral) with JSON mode.
  * [CareerExplorer](https://www.explorecareers.io/) ([GitHub](https://github.com/Nutlope/ExploreCareers)) – a site that takes in a resume and suggests career paths based on your strengths and interests. Uses Llama 3 and demonstrates how to parse PDFs and chain multiple calls together.


## 
[​](https://docs.together.ai/docs#next-steps)
Next steps
  * Check out [our Quickstart](https://docs.together.ai/docs/quickstart) to get started with our API in 1 minute
  * Explore [our cookbook](https://github.com/togethercomputer/together-cookbook) for Python recipes with Together AI
  * Explore [our demos](https://together.ai/demos) for full-stack open source example apps.
  * Check out the [Together AI playground](https://api.together.xyz/playground) to try out different models.
  * See [our integrations](https://docs.together.ai/docs/integrations) with leading LLM frameworks.


## 
[​](https://docs.together.ai/docs#resources)
Resources
  * [Discord](https://discord.com/invite/9Rk6sSeWEG)
  * [Pricing](https://www.together.ai/pricing)
  * [Support](https://www.together.ai/contact)


Was this page helpful?
YesNo
[Quickstart](https://docs.together.ai/docs/quickstart)
[x](http://twitter.com/togethercompute)[discord](https://discord.gg/9Rk6sSeWEG)[linkedin](https://www.linkedin.com/company/togethercomputer)
[Powered by Mintlify](https://mintlify.com/preview-request?utm_campaign=poweredBy&utm_medium=referral&utm_source=togetherai-52386018)
Assistant
Responses are generated using AI and may contain mistakes.


## Source Information
- URL: https://docs.together.ai/pricing
- Harvested: 2025-08-19T12:38:25.557461
- Profile: api_documentation
- Agent: architect
