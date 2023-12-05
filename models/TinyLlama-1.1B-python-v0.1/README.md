---
license: apache-2.0
datasets:
- cerebras/SlimPajama-627B
- bigcode/starcoderdata
language:
- en
---
<div align="center">

# TinyLlama-1.1B
</div>

https://github.com/jzhang38/TinyLlama

The TinyLlama project aims to **pretrain** a **1.1B Llama model on 3 trillion tokens**. With some proper optimization, we can achieve this within a span of "just" 90 days using 16 A100-40G GPUs 🚀🚀. The training has started on 2023-09-01. 


We adopted exactly the same architecture and tokenizer as Llama 2. This means TinyLlama can be plugged and played in many open-source projects built upon Llama. Besides, TinyLlama is compact with only 1.1B parameters. This compactness allows it to cater to a multitude of applications demanding a restricted computation and memory footprint.

#### This Model
This is a code LM finetuned(or so-called continue pretrianed) from the 500B TinyLlama checkpoint with another 7B Python data from the starcoderdata.

**While the finetuning data is exclusively Python, the model retains its ability in many other languages such as C or Java**.

The HumanEval accuracy is **14**.

**It can be used as the draft model to speculative-decode larger models such as models in the CodeLlama family**.