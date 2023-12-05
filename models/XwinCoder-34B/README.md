---
license: llama2
---
# XwinCoder

We are glad to introduce our instruction finetuned code generation models based on CodeLLaMA: XwinCoder. We release model weights and evaluation code.

**Repository:** [https://github.com/Xwin-LM/Xwin-LM/tree/main/Xwin-Coder](https://github.com/Xwin-LM/Xwin-LM/tree/main/Xwin-Coder)

**Models:**
| Model | 🤗hf link | HumanEval pass@1 | MBPP pass@1 | APPS-intro pass@5 |
|-------|------------|----------|------|-------------|
| XwinCoder-7B | [link](https://huggingface.co/Xwin-LM/XwinCoder-7B) | 63.8 | 57.4 | 31.5 |
| XwinCoder-13B | [link](https://huggingface.co/Xwin-LM/XwinCoder-13B) | 68.8 | 60.1 | 35.4 |
| XwinCoder-34B | [link](https://huggingface.co/Xwin-LM/XwinCoder-34B) | 74.2 | 64.8 | 43.0 |

## Updates
- 💥 We released  [**XwinCoder-7B**](https://huggingface.co/Xwin-LM/XwinCoder-7B), [**XwinCoder-13B**](https://huggingface.co/Xwin-LM/XwinCoder-13B), [**XwinCoder-34B**](https://huggingface.co/Xwin-LM/XwinCoder-34B). Our XwinCoder-34B reached 74.2 on HumanEval and it **achieves comparable performance as GPT-3.5-turbo on 6 benchmarks**.

- We support evaluating instruction finetuned models on HumanEval, MBPP, APPS, DS1000 and MT-Bench. See our github repository.

## Overview

![Chat demo](rader.png)

* To fully demonstrate our model's coding capabilities in real-world usage scenarios, we have conducted thorough evaluations on several existing mainstream coding capability leaderboards (rather than only on the currently most popular HumanEval). 
* As shown in the radar chart results, our 34B model **achieves comparable performance as GPT-3.5-turbo on coding abilities**. 
* It is worth mentioning that, to ensure accurate visualization, our radar chart has not been scaled (only translated; MT-Bench score is scaled by 10x to be more comparable with other benchmarks).
* Multiple-E-avg6 refer to the 6 languages used in CodeLLaMA paper. Results of GPT-4 and GPT-3.5-turbo are conducted by us, more details will be released later.

## Demo
We provide a chat demo in our github repository, here are some examples:
![Chat demo](exm.gif)



