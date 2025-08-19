---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T11:48:29.835324'
profile: deep_research
source: https://ollama.com/blog/how-to-prompt-code-llama
topic: Local LLM integration patterns
---

# Local LLM integration patterns

[ ![Ollama](https://ollama.com/public/ollama.png) ](https://ollama.com/)
[Models](https://ollama.com/models) [Turbo](https://ollama.com/turbo)
[Sign in](https://ollama.com/signin) [Download](https://ollama.com/download)
[Models](https://ollama.com/models) [Download](https://ollama.com/download) [Sign in](https://ollama.com/signin)
# How to prompt Code Llama
## September 9, 2023
Two weeks ago the Code Llama model was [released](https://about.fb.com/news/2023/08/code-llama-ai-for-coding/) by Meta with three variations:
  * Instruct
  * Code completion
  * Python


This guide walks through the different ways to structure prompts for Code Llama for its different variations and features.
Examples below use the 7 billion parameter model with 4-bit quantization, but 13 billion and 34 billion parameter models were made available as well. This guide uses the open-source [Ollama](https://github.com/jmorganca/ollama) project to download and prompt Code Llama, but these prompts will work in other model providers and runtimes too.
## Instruct
The `instruct` model was trained to output human-like answers to questions. This model variation is the easiest to use and will behave closest to ChatGPT, with answer questions including both natural language and code:
Prompt
```
ollama run codellama:7b-instruct 'You are an expert programmer that writes simple, concise code and explanations. Write a python function to generate the nth fibonacci number.'

```

Response
```
Here is a Python function that generates the nth Fibonacci number:
def fib(n):
  if n <= 1:
    return n
  else:
    return fib(n-1) + fib(n-2)
This function uses the recursive formula for the Fibonacci sequence, which is:
fib(n) = fib(n-1) + fib(n-2)

```

### Code Review
Code reviews can often be time consuming for teams. Code Llama can help by providing a second set of eyes on a commit or pull request for simple bugs:
Prompt
```
ollama run codellama '
Where is the bug in this code?
def fib(n):
  if n <= 0:
    return n
  else:
    return fib(n-1) + fib(n-2)
'

```

Response
```
The bug in this code is that it does not handle the case where `n` is equal to 1.

```

### Unit Tests
Writing unit tests often requires quite a bit of boilerplate code. Code Llama can help:
Prompt
```
ollama run codellama "write a unit test for this function: $(cat fib.py)"

```

Response
```
Here's an example of how you could write a unit test for the `fib` function:
```
import unittest
class TestFib(unittest.TestCase):
  def test_fib(self):
    self.assertEqual(fib(0), 0)
    self.assertEqual(fib(1), 1)
    self.assertEqual(fib(2), 1)
    self.assertEqual(fib(3), 2)
    self.assertEqual(fib(4), 3)
    self.assertEqual(fib(5), 5)
    self.assertEqual(fib(6), 8)
    self.assertEqual(fib(7), 13)
    self.assertEqual(fib(8), 21)
    self.assertEqual(fib(9), 34)
if __name__ == '__main__':
  unittest.main()
```
This test case covers all the cases where the input is an integer between 0 and 9, inclusive.

```

## Code completion
The standard code completion model will simply continue to generate subsequent tokens based on the provided prompt. This variation works best if you’re only looking to generate code, such as having the model write a function after seeing the function signature or comment:
Prompt
```
ollama run codellama:7b-code '# A simple python function to remove whitespace from a string:'

```

Response
```
def remove_whitespace(s):
  return ''.join(s.split())

```

## Infill
Fill-in-the-middle (FIM), or more briefly, **infill** is a special prompt format supported by the code completion model can complete code between two already written code blocks. This is especially effective for inserting code within an existing function such as below:
```
def compute_gcd(x, y):
  <FILL>
  return result

```

With Code Llama, infill prompts require a special format that the model expects.
```
<PRE> {prefix} <SUF>{suffix} <MID>

```

To use this with existing code, split the code before and after  in the example above the into parts: the prefix, and the suffix. For example, for our LCM example above:
Prompt
```
ollama run codellama:7b-code '<PRE> def compute_gcd(x, y): <SUF>return result <MID>'

```

While results will vary, you should get something like this:
Response
```
 if x == y:
    return x
  if x > y:
    x = x - y
  else:
    y = y - x
  result = compute_gcd(x, y)

```

> Note: the model may return `<EOT>` at the end of the result. This is a special token in the response that represents the end of the response similar to `<PRE>`, `<SUF>` and `<MID>`
## Python
As a thank you to the community and tooling that created the model, the authors of Code Llama included a Python variation which is fine-tuned on 100B additional Python tokens, making it a good model to use when working on machine learning-related tooling, or any other Python code:
Prompt
```
ollama run codellama:7b-python '
# django view for rendering the current day and time without a template
def current_datetime(request):'

```

Response
```
  now = datetime.now()
  html = "<html><body>It is now %s.</body></html>" % now
  return HttpResponse(html)

```

## Tools built on Code Llama
  * [Cody](https://cody.dev) has an [experimental](https://github.com/sourcegraph/cody/pull/905) version that uses Code Llama with infill support.
  * [Continue](https://continue.dev/) supports Code Llama as a [drop-in replacement](https://continue.dev/docs/walkthroughs/codellama) for GPT-4
  * Fine-tuned versions of Code Llama from the [Phind](https://ollama.com/library/phind-codellama) and [WizardLM](https://ollama.com/library/wizardcoder) teams
  * [Open interpreter](https://openinterpreter.com) can use Code Llama to generate functions that are then run locally in the terminal


© 2025 Ollama
[Download](https://ollama.com/download) [Blog](https://ollama.com/blog) [Docs](https://github.com/ollama/ollama/tree/main/docs) [GitHub](https://github.com/ollama/ollama) [Discord](https://discord.com/invite/ollama) [X (Twitter)](https://twitter.com/ollama) Contact Us
  * [Blog](https://ollama.com/blog)
  * [Download](https://ollama.com/download)
  * [Docs](https://github.com/ollama/ollama/tree/main/docs)


  * [GitHub](https://github.com/ollama/ollama)
  * [Discord](https://discord.com/invite/ollama)
  * [X (Twitter)](https://twitter.com/ollama)
  * [Meetups](https://lu.ma/ollama)


© 2025 Ollama Inc. 


## Source Information
- URL: https://ollama.com/blog/how-to-prompt-code-llama
- Harvested: 2025-08-19T11:48:29.835324
- Profile: deep_research
- Agent: architect
