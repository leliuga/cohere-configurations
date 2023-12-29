---
license: apache-2.0
tags:
- rag
- closed-qa
- context
- mistral
---


DocsGPT is optimized for Documentation (RAG optimised): Specifically fine-tuned for providing answers that are based on context, making it particularly useful for developers and technical support teams.
We used the Lora fine tuning process.
This model is fine tuned on top of zephyr-7b-beta


It's an apache-2.0 license so you can use it for commercial purposes too.


Benchmarks:

Bacon:
The BACON test is an internal assessment designed to evaluate the capabilities of neural networks in handling questions with substantial content. It focuses on testing the model's understanding of context-driven queries, as well as its tendency for hallucination and attention span. The questions in both parts are carefully crafted, drawing from diverse sources such as scientific papers, complex code problems, and instructional prompts, providing a comprehensive test of the model's ability to process and generate information in various domains.
| Model                        | Score |
|------------------------------|-------|
| gpt-4                        | 8.74  |
| DocsGPT-7b-Mistral           | 8.64  |
| gpt-3.5-turbo               | 8.42  |
| zephyr-7b-beta              | 8.37  |
| neural-chat-7b-v3-1         | 7.88  |
| Mistral-7B-Instruct-v0.1    | 7.44  |
| openinstruct-mistral-7b     | 5.86  |
| llama-2-13b                 | 2.29  |


![image/png](https://cdn-uploads.huggingface.co/production/uploads/6220f5dfd0351748e114ca53/lWefx5b5uQAt4Uzf_0x-O.png)


![image/png](https://cdn-uploads.huggingface.co/production/uploads/6220f5dfd0351748e114ca53/nAd4icZa2jIer-_JWOpZ0.png)




MTbench with llm judge:

![image/png](https://cdn-uploads.huggingface.co/production/uploads/6220f5dfd0351748e114ca53/SOOVW_j908gpB8W804vsG.png)

########## First turn ##########
| Model                 | Turn | Score    |
|-----------------------|------|----------|
| gpt-4                 | 1    | 8.956250 |
| gpt-3.5-turbo         | 1    | 8.075000 |
| DocsGPT-7b-Mistral    | 1    | 7.593750 |
| zephyr-7b-beta        | 1    | 7.412500 |
| vicuna-13b-v1.3       | 1    | 6.812500 |
| alpaca-13b            | 1    | 4.975000 |
| deepseek-coder-6.7b   | 1    | 4.506329 |

########## Second turn ##########
| Model                 | Turn | Score    |
|-----------------------|------|----------|
| gpt-4                 | 2    | 9.025000 |
| gpt-3.5-turbo         | 2    | 7.812500 |
| DocsGPT-7b-Mistral    | 2    | 6.740000 |
| zephyr-7b-beta        | 2    | 6.650000 |
| vicuna-13b-v1.3       | 2    | 5.962500 |
| deepseek-coder-6.7b   | 2    | 5.025641 |
| alpaca-13b            | 2    | 4.087500 |

########## Average ##########
| Model                 | Score    |
|-----------------------|----------|
| gpt-4                 | 8.990625 |
| gpt-3.5-turbo         | 7.943750 |
| DocsGPT-7b-Mistral    | 7.166875 |
| zephyr-7b-beta        | 7.031250 |
| vicuna-13b-v1.3       | 6.387500 |
| deepseek-coder-6.7b   | 4.764331 |
| alpaca-13b            | 4.531250 |




To prepare your prompts make sure you keep this format:

```
### Instruction
(where the question goes)
### Context
(your document retrieval + system instructions)
### Answer
```