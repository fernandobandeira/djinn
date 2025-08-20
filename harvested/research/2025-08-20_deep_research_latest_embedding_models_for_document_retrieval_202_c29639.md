---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-20T13:06:33.536685'
profile: deep_research
source: https://huggingface.co/blog/getting-started-with-embeddings
topic: Latest Embedding Models for Document Retrieval 2024-2025
---

# Latest Embedding Models for Document Retrieval 2024-2025

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
#  [ ](https://huggingface.co/blog/getting-started-with-embeddings#getting-started-with-embeddings) Getting Started With Embeddings 
Published June 23, 2022
[Update on GitHub](https://github.com/huggingface/blog/blob/main/getting-started-with-embeddings.md)
[ Upvote 86 ](https://huggingface.co/login?next=%2Fblog%2Fgetting-started-with-embeddings)
  * [![](https://cdn-avatars.huggingface.co/v1/production/uploads/600988773adca878cda1d294/-fm8VEoWgqFvwfYbpRhEF.jpeg)](https://huggingface.co/nezubn "nezubn")
  * [![](https://cdn-avatars.huggingface.co/v1/production/uploads/626b23fc822e3d85324babed/CbRuisnyy5JsCeA5uxYL7.jpeg)](https://huggingface.co/megrisdal "megrisdal")
  * [![](https://cdn-avatars.huggingface.co/v1/production/uploads/6270d2ddbcef985363d774fa/HOKAxx_FKVRF-87WpGQbF.png)](https://huggingface.co/real-jiakai "real-jiakai")
  * [![](https://cdn-avatars.huggingface.co/v1/production/uploads/1651742714784-noauth.jpeg)](https://huggingface.co/fpaupier "fpaupier")
  * [![](https://cdn-avatars.huggingface.co/v1/production/uploads/627cebc6cecd686d4cd7411c/DIuPgWIrCHSy704G4K26b.jpeg)](https://huggingface.co/StatsGary "StatsGary")
  * [![](https://huggingface.co/avatars/67806da539034be177393793c5c73919.svg)](https://huggingface.co/bencmbit98 "bencmbit98")
  * +80


[![Omar Espejel's avatar](https://cdn-avatars.huggingface.co/v1/production/uploads/60c5394afcb103f4f6a53c59/RNC-zezv1FolR5gVq5qSh.jpeg)](https://huggingface.co/espejelomar)
[Omar Espejel espejelomar Follow ](https://huggingface.co/espejelomar)
  * [Understanding embeddings](https://huggingface.co/blog/getting-started-with-embeddings#understanding-embeddings "Understanding embeddings")
  * [What are embeddings for?](https://huggingface.co/blog/getting-started-with-embeddings#what-are-embeddings-for "What are embeddings for?")
  * [Getting started with embeddings](https://huggingface.co/blog/getting-started-with-embeddings#getting-started-with-embeddings-1 "Getting started with embeddings")
  * [1. Embedding a dataset](https://huggingface.co/blog/getting-started-with-embeddings#1-embedding-a-dataset "1. Embedding a dataset")
  * [2. Host embeddings for free on the Hugging Face Hub](https://huggingface.co/blog/getting-started-with-embeddings#2-host-embeddings-for-free-on-the-hugging-face-hub "2. Host embeddings for free on the Hugging Face Hub")
  * [3. Get the most similar Frequently Asked Questions to a query](https://huggingface.co/blog/getting-started-with-embeddings#3-get-the-most-similar-frequently-asked-questions-to-a-query "3. Get the most similar Frequently Asked Questions to a query")
  * [Additional resources to keep learning](https://huggingface.co/blog/getting-started-with-embeddings#additional-resources-to-keep-learning "Additional resources to keep learning")


Check out this tutorial with the Notebook Companion: [ ![Open In Colab](https://colab.research.google.com/assets/colab-badge.svg) ](https://colab.research.google.com/github/huggingface/blog/blob/main/notebooks/80_getting_started_with_embeddings.ipynb)
##  [ ](https://huggingface.co/blog/getting-started-with-embeddings#understanding-embeddings) Understanding embeddings 
An embedding is a numerical representation of a piece of information, for example, text, documents, images, audio, etc. The representation captures the semantic meaning of what is being embedded, making it robust for many industry applications.
Given the text "What is the main benefit of voting?", an embedding of the sentence could be represented in a vector space, for example, with a list of 384 numbers (for example, [0.84, 0.42, ..., 0.02]). Since this list captures the meaning, we can do exciting things, like calculating the distance between different embeddings to determine how well the meaning of two sentences matches.
Embeddings are not limited to text! You can also create an embedding of an image (for example, a list of 384 numbers) and compare it with a text embedding to determine if a sentence describes the image. This concept is under powerful systems for image search, classification, description, and more!
How are embeddings generated? The open-source library called [Sentence Transformers](https://www.sbert.net/index.html) allows you to create state-of-the-art embeddings from images and text for free. This blog shows an example with this library.
##  [ ](https://huggingface.co/blog/getting-started-with-embeddings#what-are-embeddings-for) What are embeddings for? 
> "[...] once you understand this ML multitool (embedding), you'll be able to build everything from search engines to recommendation systems to chatbots and a whole lot more. You don't have to be a data scientist with ML expertise to use them, nor do you need a huge labeled dataset." - [Dale Markowitz, Google Cloud](https://cloud.google.com/blog/topics/developers-practitioners/meet-ais-multitool-vector-embeddings).
Once a piece of information (a sentence, a document, an image) is embedded, the creativity starts; several interesting industrial applications use embeddings. E.g., Google Search uses embeddings to [match text to text and text to images](https://cloud.google.com/blog/topics/developers-practitioners/meet-ais-multitool-vector-embeddings); Snapchat uses them to "[serve the right ad to the right user at the right time](https://eng.snap.com/machine-learning-snap-ad-ranking)"; and Meta (Facebook) uses them for [their social search](https://research.facebook.com/publications/embedding-based-retrieval-in-facebook-search/).
Before they could get intelligence from embeddings, these companies had to embed their pieces of information. An embedded dataset allows algorithms to search quickly, sort, group, and more. However, it can be expensive and technically complicated. In this post, we use simple open-source tools to show how easy it can be to embed and analyze a dataset.
##  [ ](https://huggingface.co/blog/getting-started-with-embeddings#getting-started-with-embeddings-1) Getting started with embeddings 
We will create a small Frequently Asked Questions (FAQs) engine: receive a query from a user and identify which FAQ is the most similar. We will use the [US Social Security Medicare FAQs](https://faq.ssa.gov/en-US/topic/?id=CAT-01092).
But first, we need to embed our dataset (other texts use the terms encode and embed interchangeably). The Hugging Face Inference API allows us to embed a dataset using a quick POST call easily.
Since the embeddings capture the semantic meaning of the questions, it is possible to compare different embeddings and see how different or similar they are. Thanks to this, you can get the most similar embedding to a query, which is equivalent to finding the most similar FAQ. Check out our [semantic search tutorial](https://huggingface.co/spaces/sentence-transformers/embeddings-semantic-search) for a more detailed explanation of how this mechanism works.
In a nutshell, we will:
  1. Embed Medicare's FAQs using the Inference API.
  2. Upload the embedded questions to the Hub for free hosting.
  3. Compare a customer's query to the embedded dataset to identify which is the most similar FAQ.


##  [ ](https://huggingface.co/blog/getting-started-with-embeddings#1-embedding-a-dataset) 1. Embedding a dataset 
The first step is selecting an existing pre-trained model for creating the embeddings. We can choose a model from the [Sentence Transformers library](https://huggingface.co/sentence-transformers). In this case, let's use the ["sentence-transformers/all-MiniLM-L6-v2"](https://huggingface.co/sentence-transformers/all-MiniLM-L6-v2) because it's a small but powerful model. In a future post, we will examine other models and their trade-offs.
Log in to the Hub. You must create a write token in your [Account Settings](http://hf.co/settings/tokens). We will store the write token in `hf_token`.
```
model_id = "sentence-transformers/all-MiniLM-L6-v2"
hf_token = "get your token in http://hf.co/settings/tokens"

```

To generate the embeddings you can use the `https://api-inference.huggingface.co/pipeline/feature-extraction/{model_id}` endpoint with the headers `{"Authorization": f"Bearer {hf_token}"}`. Here is a function that receives a dictionary with the texts and returns a list with embeddings.
```
import requests
api_url = f"https://api-inference.huggingface.co/pipeline/feature-extraction/{model_id}"
headers = {"Authorization": f"Bearer {hf_token}"}

```

The first time you generate the embeddings, it may take a while (approximately 20 seconds) for the API to return them. We use the `retry` decorator (install with `pip install retry`) so that if on the first try, `output = query(dict(inputs = texts))` doesn't work, wait 10 seconds and try three times again. This happens because, on the first request, the model needs to be downloaded and installed on the server, but subsequent calls are much faster.
```
def query(texts):
  response = requests.post(api_url, headers=headers, json={"inputs": texts, "options":{"wait_for_model":True}})
  return response.json()

```

The current API does not enforce strict rate limitations. Instead, Hugging Face balances the loads evenly between all our available resources and favors steady flows of requests. If you need to embed several texts or images, the [Hugging Face Accelerated Inference API](https://huggingface.co/docs/api-inference/index) would speed the inference and let you choose between using a CPU or GPU. 
```
texts = ["How do I get a replacement Medicare card?",
    "What is the monthly premium for Medicare Part B?",
    "How do I terminate my Medicare Part B (medical insurance)?",
    "How do I sign up for Medicare?",
    "Can I sign up for Medicare Part B if I am working and have health insurance through an employer?",
    "How do I sign up for Medicare Part B if I already have Part A?",
    "What are Medicare late enrollment penalties?",
    "What is Medicare and who can get it?",
    "How can I get help with my Medicare Part A and Part B premiums?",
    "What are the different parts of Medicare?",
    "Will my Medicare premiums be higher because of my higher income?",
    "What is TRICARE ?",
    "Should I sign up for Medicare Part B if I have Veterans' Benefits?"]
output = query(texts)

```

As a response, you get back a list of lists. Each list contains the embedding of a FAQ. The model, ["sentence-transformers/all-MiniLM-L6-v2"](https://huggingface.co/sentence-transformers/all-MiniLM-L6-v2), is encoding the input questions to 13 embeddings of size 384 each. Let's convert the list to a Pandas `DataFrame` of shape (13x384).
```
import pandas as pd
embeddings = pd.DataFrame(output)

```

It looks similar to this matrix:
```
[[-0.02388945 0.05525852 -0.01165488 ... 0.00577787 0.03409787 -0.0068891 ]
 [-0.0126876  0.04687412 -0.01050217 ... -0.02310316 -0.00278466  0.01047371]
 [ 0.00049438 0.11941205 0.00522949 ... 0.01687654 -0.02386115  0.00526433]
 ...
 [-0.03900796 -0.01060951 -0.00738271 ... -0.08390449 0.03768405  0.00231361]
 [-0.09598278 -0.06301168 -0.11690582 ... 0.00549841 0.1528919  0.02472013]
 [-0.01162949 0.05961934 0.01650903 ... -0.02821241 -0.00116556  0.0010672 ]]

```

##  [ ](https://huggingface.co/blog/getting-started-with-embeddings#2-host-embeddings-for-free-on-the-hugging-face-hub) 2. Host embeddings for free on the Hugging Face Hub 
🤗 Datasets is a library for quickly accessing and sharing datasets. Let's host the embeddings dataset in the Hub using the user interface (UI). Then, anyone can load it with a single line of code. You can also use the terminal to share datasets; see [the documentation](https://huggingface.co/docs/datasets/share#share) for the steps. In the [notebook companion](https://colab.research.google.com/github/huggingface/blog/blob/main/notebooks/80_getting_started_with_embeddings.ipynb) of this entry, you will be able to use the terminal to share the dataset. If you want to skip this section, check out the [`ITESM/embedded_faqs_medicare` repo](https://huggingface.co/datasets/ITESM/embedded_faqs_medicare) with the embedded FAQs.
First, we export our embeddings from a Pandas `DataFrame` to a CSV. You can save your dataset in any way you prefer, e.g., zip or pickle; you don't need to use Pandas or CSV. Since our embeddings file is not large, we can store it in a CSV, which is easily inferred by the `datasets.load_dataset()` function we will employ in the next section (see the [Datasets documentation](https://huggingface.co/docs/datasets/about_dataset_load#build-and-load)), i.e., we don't need to create a loading script. We will save the embeddings with the name `embeddings.csv`.
```
embeddings.to_csv("embeddings.csv", index=False)

```

Follow the next steps to host `embeddings.csv` in the Hub.
  * Click on your user in the top right corner of the [Hub UI](https://huggingface.co/).
  * Create a dataset with "New dataset."


[![](https://huggingface.co/blog/assets/80_getting_started_with_embeddings/SelectDataset.png)](https://huggingface.co/blog/assets/80_getting_started_with_embeddings/SelectDataset.png)
  * Choose the Owner (organization or individual), name, and license of the dataset. Select if you want it to be private or public. Create the dataset.


[![](https://huggingface.co/blog/assets/80_getting_started_with_embeddings/createDataset.png)](https://huggingface.co/blog/assets/80_getting_started_with_embeddings/createDataset.png)
  * Go to the "Files" tab (screenshot below) and click "Add file" and "Upload file."


[![](https://huggingface.co/blog/assets/80_getting_started_with_embeddings/AddFile.png)](https://huggingface.co/blog/assets/80_getting_started_with_embeddings/AddFile.png)
  * Finally, drag or upload the dataset, and commit the changes.


[![](https://huggingface.co/blog/assets/80_getting_started_with_embeddings/UploadFile.png)](https://huggingface.co/blog/assets/80_getting_started_with_embeddings/UploadFile.png)
Now the dataset is hosted on the Hub for free. You (or whoever you want to share the embeddings with) can quickly load them. Let's see how.
##  [ ](https://huggingface.co/blog/getting-started-with-embeddings#3-get-the-most-similar-frequently-asked-questions-to-a-query) 3. Get the most similar Frequently Asked Questions to a query 
Suppose a Medicare customer asks, "How can Medicare help me?". We will **find** which of our FAQs could best answer our user query. We will create an embedding of the query that can represent its semantic meaning. We then compare it to each embedding in our FAQ dataset to identify which is closest to the query in vector space.
Install the 🤗 Datasets library with `pip install datasets`. Then, load the embedded dataset from the Hub and convert it to a PyTorch `FloatTensor`. Note that this is not the only way to operate on a `Dataset`; for example, you could use NumPy, Tensorflow, or SciPy (refer to the [Documentation](https://huggingface.co/docs/datasets/loading)). If you want to practice with a real dataset, the [`ITESM/embedded_faqs_medicare`](https://huggingface.co/datasets/ITESM/embedded_faqs_medicare) repo contains the embedded FAQs, or you can use the [companion notebook](https://colab.research.google.com/github/huggingface/blog/blob/main/notebooks/80_getting_started_with_embeddings.ipynb) to this blog. 
```
import torch
from datasets import load_dataset
faqs_embeddings = load_dataset('namespace/repo_name')
dataset_embeddings = torch.from_numpy(faqs_embeddings["train"].to_pandas().to_numpy()).to(torch.float)

```

We use the query function we defined before to embed the customer's question and convert it to a PyTorch `FloatTensor` to operate over it efficiently. Note that after the embedded dataset is loaded, we could use the `add_faiss_index` and `search` methods of a `Dataset` to identify the closest FAQ to an embedded query using the [faiss library](https://github.com/facebookresearch/faiss). Here is a [nice tutorial of the alternative](https://huggingface.co/docs/datasets/faiss_es).
```
question = ["How can Medicare help me?"]
output = query(question)
query_embeddings = torch.FloatTensor(output)

```

You can use the `util.semantic_search` function in the Sentence Transformers library to identify which of the FAQs are closest (most similar) to the user's query. This function uses cosine similarity as the default function to determine the proximity of the embeddings. However, you could also use other functions that measure the distance between two points in a vector space, for example, the dot product.
Install `sentence-transformers` with `pip install -U sentence-transformers`, and search for the five most similar FAQs to the query.
```
from sentence_transformers.util import semantic_search
hits = semantic_search(query_embeddings, dataset_embeddings, top_k=5)

```

`util.semantic_search` identifies how close each of the 13 FAQs is to the customer query and returns a list of dictionaries with the top `top_k` FAQs. `hits` looks like this:
```
[{'corpus_id': 8, 'score': 0.75653076171875},
 {'corpus_id': 7, 'score': 0.7418993711471558},
 {'corpus_id': 3, 'score': 0.7252674102783203},
 {'corpus_id': 9, 'score': 0.6735571622848511},
 {'corpus_id': 10, 'score': 0.6505177617073059}]

```

The values ​​in `corpus_id` allow us to index the list of `texts` we defined in the first section and get the five most similar FAQs:
```
print([texts[hits[0][i]['corpus_id']] for i in range(len(hits[0]))])

```

Here are the 5 FAQs that come closest to the customer's query:
```
['How can I get help with my Medicare Part A and Part B premiums?',
 'What is Medicare and who can get it?',
 'How do I sign up for Medicare?',
 'What are the different parts of Medicare?',
 'Will my Medicare premiums be higher because of my higher income?']

```

This list represents the 5 FAQs closest to the customer's query. Nice! We used here PyTorch and Sentence Transformers as our main numerical tools. However, we could have defined the cosine similarity and ranking functions by ourselves using tools such as NumPy and SciPy.
##  [ ](https://huggingface.co/blog/getting-started-with-embeddings#additional-resources-to-keep-learning) Additional resources to keep learning 
If you want to know more about the Sentence Transformers library:
  * The [Hub Organization](https://huggingface.co/sentence-transformers) for all the new models and instructions on how to download models.
  * The [Nils Reimers tweet](https://twitter.com/Nils_Reimers/status/1487014195568775173) comparing Sentence Transformer models with GPT-3 Embeddings. Spoiler alert: the Sentence Transformers are awesome!
  * The [Sentence Transformers documentation](https://www.sbert.net/),
  * [Nima's thread](https://twitter.com/NimaBoscarino/status/1535331680805801984) on recent research.


Thanks for reading!
More Articles from our Blog
[ ![](https://huggingface.co/blog/assets/train-sentence-transformers/st-hf-thumbnail.png) Training and Finetuning Sparse Embedding Models with Sentence Transformers v5 By [tomaarsen](https://huggingface.co/tomaarsen) June 30, 2025 • 108](https://huggingface.co/blog/train-sparse-encoder)
[ ![](https://huggingface.co/blog/assets/train-sentence-transformers/st-hf-thumbnail.png) Training and Finetuning Reranker Models with Sentence Transformers v4 By [tomaarsen](https://huggingface.co/tomaarsen) March 25, 2025 • 158](https://huggingface.co/blog/train-reranker)
### Community
![](https://cdn-avatars.huggingface.co/v1/production/uploads/no-auth/Zlo_9GtB61Lyl5G5wA95h.png) [SAAHMATHWORKS](https://huggingface.co/SAAHMATHWORKS)
[Jun 19](https://huggingface.co/blog/getting-started-with-embeddings#6853f365bbb97f3b905b0d45)
•
[edited Jun 19](https://huggingface.co/blog/getting-started-with-embeddings#6853f365bbb97f3b905b0d45 "Edited by SAAHMATHWORKS")
Hibelow is the error message I have when I try to run the code of the post,{'error': "SentenceSimilarityPipeline.**call**() missing 1 required positional argument: 'sentences'"}
My codemodel_id = "sentence-transformers/all-MiniLM-L6-v2"api_url = f"[https://api-inference.huggingface.co/models/{model_id}](https://api-inference.huggingface.co/models/%7Bmodel_id%7D)"
hf_token = os.environ['access_token']headers = {"Authorization": f"Bearer {hf_token}"}
def query(texts): response = requests.post(api_url, headers=headers, json={"inputs": texts, "options":{"wait_for_model":True}}) return response.json()
texts = ["How do I get a replacement Medicare card?", "What is the monthly premium for Medicare Part B?", "How do I terminate my Medicare Part B (medical insurance)?"]
output = query(texts)print(output)
Reply
![](https://cdn-avatars.huggingface.co/v1/production/uploads/no-auth/Zlo_9GtB61Lyl5G5wA95h.png) [SAAHMATHWORKS](https://huggingface.co/SAAHMATHWORKS)
[Jun 19](https://huggingface.co/blog/getting-started-with-embeddings#6853f8b510422878336bf0b3)
I adjusted my query function with the help of claude.ai like below and I can see output with a free tier
def query(sentence_pairs): response = requests.post( api_url, headers=headers, json={"inputs": {"source_sentence": sentence_pairs[0], "sentences": sentence_pairs[1:]}} ) return response.json()
Reply
![](https://huggingface.co/avatars/8929375bb8431ae68857b881b6621295.svg) [gscode](https://huggingface.co/gscode)
[Jul 18](https://huggingface.co/blog/getting-started-with-embeddings#687a0f7032fb2368035a2f42)
For your awareness, I believe the API URL for the all-MiniLM-L6-v2 transformer has now changed, meaning that the Google Colab Notebook doesn't work as written. The new URL is provided in the Hugging Face forum post here and updating to the new URL seems to fix the issue: <https://huggingface.co/sentence-transformers/all-MiniLM-L6-v2/discussions/116>
Reply
EditPreview
Upload images, audio, and videos by dragging in the text input, pasting, or clicking here.
Tap or paste here to upload images
Comment
· [Sign up](https://huggingface.co/join?next=%2Fblog%2Fgetting-started-with-embeddings) or [log in](https://huggingface.co/login?next=%2Fblog%2Fgetting-started-with-embeddings) to comment
[ Upvote 86 ](https://huggingface.co/login?next=%2Fblog%2Fgetting-started-with-embeddings)
  * [![](https://cdn-avatars.huggingface.co/v1/production/uploads/600988773adca878cda1d294/-fm8VEoWgqFvwfYbpRhEF.jpeg)](https://huggingface.co/nezubn "nezubn")
  * [![](https://cdn-avatars.huggingface.co/v1/production/uploads/626b23fc822e3d85324babed/CbRuisnyy5JsCeA5uxYL7.jpeg)](https://huggingface.co/megrisdal "megrisdal")
  * [![](https://cdn-avatars.huggingface.co/v1/production/uploads/6270d2ddbcef985363d774fa/HOKAxx_FKVRF-87WpGQbF.png)](https://huggingface.co/real-jiakai "real-jiakai")
  * [![](https://cdn-avatars.huggingface.co/v1/production/uploads/1651742714784-noauth.jpeg)](https://huggingface.co/fpaupier "fpaupier")
  * [![](https://cdn-avatars.huggingface.co/v1/production/uploads/627cebc6cecd686d4cd7411c/DIuPgWIrCHSy704G4K26b.jpeg)](https://huggingface.co/StatsGary "StatsGary")
  * [![](https://huggingface.co/avatars/67806da539034be177393793c5c73919.svg)](https://huggingface.co/bencmbit98 "bencmbit98")
  * [![](https://cdn-avatars.huggingface.co/v1/production/uploads/62a3bb1cd0d8c2c2169f0b88/eT2TS0IlQbZtz-F_zHLz9.jpeg)](https://huggingface.co/Tonic "Tonic")
  * [![](https://cdn-avatars.huggingface.co/v1/production/uploads/62b049e653d878042fab5673/lSgJdq2m-oLHHeeuMeSBV.jpeg)](https://huggingface.co/AhmadHakami "AhmadHakami")
  * [![](https://cdn-avatars.huggingface.co/v1/production/uploads/6317233cc92fd6fee317e030/cJHSvvimr1kqgQfHOjO5n.png)](https://huggingface.co/tomaarsen "tomaarsen")
  * [![](https://cdn-avatars.huggingface.co/v1/production/uploads/635237170f9bdb641c43a6d3/i9sX38SQn5-wGOrtHmwZk.jpeg)](https://huggingface.co/waynesworld0101 "waynesworld0101")
  * [![](https://cdn-avatars.huggingface.co/v1/production/uploads/1666853944551-noauth.jpeg)](https://huggingface.co/Sandeepa "Sandeepa")
  * [![](https://cdn-avatars.huggingface.co/v1/production/uploads/638604149d7b8afb0e809223/_mmlrfeKVP9AxEyiTsdGC.jpeg)](https://huggingface.co/k3ybladewielder "k3ybladewielder")
  * +74


System theme 
Company
[TOS](https://huggingface.co/terms-of-service) [Privacy](https://huggingface.co/privacy) [About](https://huggingface.co/huggingface) [Jobs](https://apply.workable.com/huggingface/) [](https://huggingface.co/)
Website
[Models](https://huggingface.co/models) [Datasets](https://huggingface.co/datasets) [Spaces](https://huggingface.co/spaces) [Pricing](https://huggingface.co/pricing) [Docs](https://huggingface.co/docs)


## Source Information
- URL: https://huggingface.co/blog/getting-started-with-embeddings
- Harvested: 2025-08-20T13:06:33.536685
- Profile: deep_research
- Agent: architect
