---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-20T13:07:39.027147'
profile: deep_research
source: https://huggingface.co/blog/mteb
topic: Embedding Models Benchmarks 2024-2025 MTEB
---

# Embedding Models Benchmarks 2024-2025 MTEB

[![Hugging Face's logo](https://huggingface.co/front/assets/huggingface_logo-noborder.svg) Hugging Face](https://huggingface.co/)
  * [ Models](https://huggingface.co/models)
  * [ Datasets](https://huggingface.co/datasets)
  * [ Spaces](https://huggingface.co/spaces)
  * Community 
  * [ Docs](https://huggingface.co/docs)
  * [ Enterprise](https://huggingface.co/enterprise)
  * [Pricing](https://huggingface.co/pricing)
  * [Log In](https://huggingface.co/login)
  * [Sign Up](https://huggingface.co/join)


[ Back to Articles](https://huggingface.co/blog)
#  [ ](https://huggingface.co/blog/mteb#mteb-massive-text-embedding-benchmark) MTEB: Massive Text Embedding Benchmark 
Published October 19, 2022
[Update on GitHub](https://github.com/huggingface/blog/blob/main/mteb.md)
[ Upvote 78 ](https://huggingface.co/login?next=%2Fblog%2Fmteb)
  * [![](https://cdn-avatars.huggingface.co/v1/production/uploads/60bfa4237f75bb4d92557db9/8Vu3xJkqI59GrtoFrZbwj.jpeg)](https://huggingface.co/wilfoderek "wilfoderek")
  * [![](https://huggingface.co/avatars/791d96ede004cff12c4f394a7383ca4c.svg)](https://huggingface.co/Kalyan369 "Kalyan369")
  * [![](https://huggingface.co/avatars/5585e95f8695907a238047841fe1a9fe.svg)](https://huggingface.co/inesruizruiz "inesruizruiz")
  * [![](https://cdn-avatars.huggingface.co/v1/production/uploads/6125ec1fa44deeb069efdbcf/mPWPKuNPV-qqKvX0tekgw.jpeg)](https://huggingface.co/Hinova "Hinova")
  * [![](https://cdn-avatars.huggingface.co/v1/production/uploads/612ee6a7b960e78c6d2319d4/2Hu9BaAyXbyh1vt0v1Qui.jpeg)](https://huggingface.co/SivilTaram "SivilTaram")
  * [![](https://huggingface.co/avatars/f34797eccf5b008deea640dd4a1696c3.svg)](https://huggingface.co/Gnana1306 "Gnana1306")
  * +72


[![Niklas Muennighoff's avatar](https://cdn-avatars.huggingface.co/v1/production/uploads/5f1eb362eec0ad2a071ad6e2/IXMYkYKuTwn6kBdWnQeeY.png)](https://huggingface.co/Muennighoff)
[Niklas Muennighoff Muennighoff Follow ](https://huggingface.co/Muennighoff)
  * [Why Text Embeddings?](https://huggingface.co/blog/mteb#why-text-embeddings "Why Text Embeddings?")
  * [MTEB](https://huggingface.co/blog/mteb#mteb "MTEB")
  * [Models](https://huggingface.co/blog/mteb#models "Models")
  * [Benchmark your model](https://huggingface.co/blog/mteb#benchmark-your-model "Benchmark your model")
  * [The script will produce a `mteb_metadata.md` file that looks like this: ```sh](https://huggingface.co/blog/mteb#the-script-will-produce-a-mteb_metadatamd-file-that-looks-like-this-sh "The script will produce a <code>mteb_metadata.md</code> file that looks like this:
```sh")
  * [tags: - mteb model-index: - name: average_word_embeddings_komninos results: - task: type: Classification dataset: type: mteb/banking77 name: MTEB Banking77Classification config: default split: test revision: 0fd18e25b25c072e09e0d92ab615fda904d66300 metrics: - type: accuracy value: 66.76623376623377 - type: f1 value: 66.59096432882667](https://huggingface.co/blog/mteb#tags---mteb-model-index---name-averageword_embeddingskomninos---results-----task-------type-classification-----dataset-------type-mtebbanking77-------name-mteb-banking77classification-------config-default-------split-test-------revision-0fd18e25b25c072e09e0d92ab615fda904d66300-----metrics-------type-accuracy-------value-6676623376623377-------type-f1-------value-6659096432882667 "tags:
- mteb
model-index:
- name: average_word_embeddings_komninos
 results:
 - task:
   type: Classification
  dataset:
   type: mteb/banking77
   name: MTEB Banking77Classification
   config: default
   split: test
   revision: 0fd18e25b25c072e09e0d92ab615fda904d66300
  metrics:
  - type: accuracy
   value: 66.76623376623377
  - type: f1
   value: 66.59096432882667")


MTEB is a massive benchmark for measuring the performance of text embedding models on diverse embedding tasks. 
The 🥇 [leaderboard](https://huggingface.co/spaces/mteb/leaderboard) provides a holistic view of the best text embedding models out there on a variety of tasks. 
The 📝 [paper](https://arxiv.org/abs/2210.07316) gives background on the tasks and datasets in MTEB and analyzes leaderboard results!
The 💻 [Github repo](https://github.com/embeddings-benchmark/mteb) contains the code for benchmarking and submitting any model of your choice to the leaderboard.
[![MTEB Leaderboard](https://huggingface.co/blog/assets/110_mteb/leaderboard.png)](https://huggingface.co/spaces/mteb/leaderboard)
##  [ ](https://huggingface.co/blog/mteb#why-text-embeddings) Why Text Embeddings? 
Text Embeddings are vector representations of text that encode semantic information. As machines require numerical inputs to perform computations, text embeddings are a crucial component of many downstream NLP applications. For example, Google uses text embeddings to [power their search engine](https://cloud.google.com/blog/topics/developers-practitioners/find-anything-blazingly-fast-googles-vector-search-technology). Text Embeddings can also be used for finding [patterns in large amount of text via clustering](https://txt.cohere.ai/combing-for-insight-in-10-000-hacker-news-posts-with-text-clustering/) or as inputs to text classification models, such as in our recent [SetFit](https://huggingface.co/blog/setfit) work. The quality of text embeddings, however, is highly dependent on the embedding model used. MTEB is designed to help you find the best embedding model out there for a variety of tasks!
##  [ ](https://huggingface.co/blog/mteb#mteb) MTEB 
🐋 **Massive** : MTEB includes 56 datasets across 8 tasks and currently summarizes >2000 results on the [leaderboard](https://huggingface.co/spaces/mteb/leaderboard). 
🌎 **Multilingual** : MTEB contains up to 112 different languages! We have benchmarked several multilingual models on Bitext Mining, Classification, and STS. 
🦚 **Extensible** : Be it new tasks, datasets, metrics, or leaderboard additions, any contribution is very welcome. Check out the GitHub repository to [submit to the leaderboard](https://github.com/embeddings-benchmark/mteb#leaderboard) or [solve open issues](https://github.com/embeddings-benchmark/mteb/issues). We hope you join us on the journey of finding the best text embedding model!
![MTEB Taxonomy](https://huggingface.co/blog/assets/110_mteb/mteb_diagram_white_background.png)
_Overview of tasks and datasets in MTEB. Multilingual datasets are marked with a purple shade._
##  [ ](https://huggingface.co/blog/mteb#models) Models 
For the initial benchmarking of MTEB, we focused on models claiming state-of-the-art results and popular models on the Hub. This led to a high representation of transformers. 🤖
![MTEB Speed Benchmark](https://huggingface.co/blog/assets/110_mteb/benchmark.png)
_Models by average English MTEB score (y) vs speed (x) vs embedding size (circle size)._
We grouped models into the following three attributes to simplify finding the best model for your task:
**🏎 Maximum speed** Models like [Glove](https://huggingface.co/sentence-transformers/average_word_embeddings_glove.6B.300d) offer high speed, but suffer from a lack of context awareness resulting in low average MTEB scores.
**⚖️ Speed and performance** Slightly slower, but significantly stronger, [all-mpnet-base-v2](https://huggingface.co/sentence-transformers/all-mpnet-base-v2) or [all-MiniLM-L6-v2](https://huggingface.co/sentence-transformers/all-MiniLM-L6-v2) provide a good balance between speed and performance.
**💪 Maximum performance** Multi-billion parameter models like [ST5-XXL](https://huggingface.co/sentence-transformers/sentence-t5-xxl), [GTR-XXL](https://huggingface.co/sentence-transformers/gtr-t5-xxl) or [SGPT-5.8B-msmarco](https://huggingface.co/Muennighoff/SGPT-5.8B-weightedmean-msmarco-specb-bitfit) dominate on MTEB. They tend to also produce bigger embeddings like [SGPT-5.8B-msmarco](https://huggingface.co/Muennighoff/SGPT-5.8B-weightedmean-msmarco-specb-bitfit) which produces 4096 dimensional embeddings requiring more storage!
Model performance varies a lot depending on the task and dataset, so we recommend checking the various tabs of the [leaderboard](https://huggingface.co/spaces/mteb/leaderboard) before deciding which model to use!
##  [ ](https://huggingface.co/blog/mteb#benchmark-your-model) Benchmark your model 
Using the [MTEB library](https://github.com/embeddings-benchmark/mteb), you can benchmark any model that produces embeddings and add its results to the public leaderboard. Let's run through a quick example!
First, install the library:
```
pip install mteb

```

Next, benchmark a model on a dataset, for example [komninos word embeddings](https://huggingface.co/sentence-transformers/average_word_embeddings_komninos) on [Banking77](https://huggingface.co/datasets/mteb/banking77).
```
from mteb import MTEB
from sentence_transformers import SentenceTransformer
model_name = "average_word_embeddings_komninos"
model = SentenceTransformer(model_name)
evaluation = MTEB(tasks=["Banking77Classification"])
results = evaluation.run(model, output_folder=f"results/{model_name}")

```

This should produce a `results/average_word_embeddings_komninos/Banking77Classification.json` file!
Now you can submit the results to the leaderboard by adding it to the metadata of the `README.md` of any model on the Hub.
Run our [automatic script](https://github.com/embeddings-benchmark/mteb/blob/main/scripts/mteb_meta.py) to generate the metadata:
```
python mteb_meta.py results/average_word_embeddings_komninos

```

##  [ ](https://huggingface.co/blog/mteb#the-script-will-produce-a-mteb_metadatamd-file-that-looks-like-this-sh) The script will produce a `mteb_metadata.md` file that looks like this: ```sh 
##  [ ](https://huggingface.co/blog/mteb#tags---mteb-model-index---name-average_word_embeddings_komninos---results-----task-------type-classification-----dataset-------type-mtebbanking77-------name-mteb-banking77classification-------config-default-------split-test-------revision-0fd18e25b25c072e09e0d92ab615fda904d66300-----metrics-------type-accuracy-------value-6676623376623377-------type-f1-------value-6659096432882667) tags: - mteb model-index: - name: average_word_embeddings_komninos results: - task: type: Classification dataset: type: mteb/banking77 name: MTEB Banking77Classification config: default split: test revision: 0fd18e25b25c072e09e0d92ab615fda904d66300 metrics: - type: accuracy value: 66.76623376623377 - type: f1 value: 66.59096432882667 
```

Now add the metadata to the top of a `README.md` of any model on the Hub, like this [SGPT-5.8B-msmarco](https://huggingface.co/Muennighoff/SGPT-5.8B-weightedmean-msmarco-specb-bitfit/blob/main/README.md) model, and it will show up on the [leaderboard](https://huggingface.co/spaces/mteb/leaderboard) after refreshing!
## Next steps
Go out there and benchmark any model you like! Let us know if you have questions or feedback by opening an issue on our [GitHub repo](https://github.com/embeddings-benchmark/mteb) or the [leaderboard community tab](https://huggingface.co/spaces/mteb/leaderboard/discussions) 🤗
Happy embedding!
## Credits
Huge thanks to the following who contributed to the article or to the MTEB codebase (listed in alphabetical order): Steven Liu, Loïc Magne, Nils Reimers and Nouamane Tazi.

```

More Articles from our Blog
[ ![](https://huggingface.co/blog/assets/ettin/thumbnail.png) Seq vs Seq: the Ettin Suite of Paired Encoders and Decoders By [orionweller](https://huggingface.co/orionweller) July 15, 2025 • 58](https://huggingface.co/blog/ettin)
[ ![](https://huggingface.co/blog/assets/smollm3/image.png) SmolLM3: smol, multilingual, long-context reasoner By [loubnabnl](https://huggingface.co/loubnabnl) July 7, 2025 • 631](https://huggingface.co/blog/smollm3)
### Community
EditPreview
Upload images, audio, and videos by dragging in the text input, pasting, or clicking here.
Tap or paste here to upload images
Comment
· [Sign up](https://huggingface.co/join?next=%2Fblog%2Fmteb) or [log in](https://huggingface.co/login?next=%2Fblog%2Fmteb) to comment
[ Upvote 78 ](https://huggingface.co/login?next=%2Fblog%2Fmteb)
  * [![](https://cdn-avatars.huggingface.co/v1/production/uploads/60bfa4237f75bb4d92557db9/8Vu3xJkqI59GrtoFrZbwj.jpeg)](https://huggingface.co/wilfoderek "wilfoderek")
  * [![](https://huggingface.co/avatars/791d96ede004cff12c4f394a7383ca4c.svg)](https://huggingface.co/Kalyan369 "Kalyan369")
  * [![](https://huggingface.co/avatars/5585e95f8695907a238047841fe1a9fe.svg)](https://huggingface.co/inesruizruiz "inesruizruiz")
  * [![](https://cdn-avatars.huggingface.co/v1/production/uploads/6125ec1fa44deeb069efdbcf/mPWPKuNPV-qqKvX0tekgw.jpeg)](https://huggingface.co/Hinova "Hinova")
  * [![](https://cdn-avatars.huggingface.co/v1/production/uploads/612ee6a7b960e78c6d2319d4/2Hu9BaAyXbyh1vt0v1Qui.jpeg)](https://huggingface.co/SivilTaram "SivilTaram")
  * [![](https://huggingface.co/avatars/f34797eccf5b008deea640dd4a1696c3.svg)](https://huggingface.co/Gnana1306 "Gnana1306")
  * [![](https://cdn-avatars.huggingface.co/v1/production/uploads/625188cdcdad03f220d87950/PhGpcRcFa99badxQgmw1q.jpeg)](https://huggingface.co/naiborhujosua "naiborhujosua")
  * [![](https://huggingface.co/avatars/1e6f6d78aa1e72c5eb39cbfe5da896cb.svg)](https://huggingface.co/susnato "susnato")
  * [![](https://huggingface.co/avatars/82a199a95e40b7d3046b6fc0975ad080.svg)](https://huggingface.co/jbohnslav "jbohnslav")
  * [![](https://cdn-avatars.huggingface.co/v1/production/uploads/6270d2ddbcef985363d774fa/HOKAxx_FKVRF-87WpGQbF.png)](https://huggingface.co/real-jiakai "real-jiakai")
  * [![](https://cdn-avatars.huggingface.co/v1/production/uploads/6281d941eeb15579946ca3ce/0CdrBop_kjRkOqxUTYFbf.jpeg)](https://huggingface.co/CocoSun "CocoSun")
  * [![](https://cdn-avatars.huggingface.co/v1/production/uploads/62a4ac6fd83c3facafa50892/qFpobw9B5XaLZvwn0XbmB.jpeg)](https://huggingface.co/mohammedbriman "mohammedbriman")
  * +66


System theme 
Company
[TOS](https://huggingface.co/terms-of-service) [Privacy](https://huggingface.co/privacy) [About](https://huggingface.co/huggingface) [Jobs](https://apply.workable.com/huggingface/) [](https://huggingface.co/)
Website
[Models](https://huggingface.co/models) [Datasets](https://huggingface.co/datasets) [Spaces](https://huggingface.co/spaces) [Pricing](https://huggingface.co/pricing) [Docs](https://huggingface.co/docs)


## Source Information
- URL: https://huggingface.co/blog/mteb
- Harvested: 2025-08-20T13:07:39.027147
- Profile: deep_research
- Agent: architect
