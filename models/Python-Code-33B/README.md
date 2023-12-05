---
license: cc-by-nc-nd-4.0
datasets:
- ajibawa-2023/Python-Code-23k-ShareGPT
language:
- en
tags:
- code
---

**Python-Code-33B**

Large Language Models (LLMs) are good with code generations. Sometimes LLMs do make mistakes in code generation. How about if they can give detailed explanation along with the code.
This is what I have tried over here. The base Llama-2 model was used for training purpose. It is trained on around 23000+ set of codes. Each set having 2 conversations.
This data was generated using GPT-3.5, GPT-4 etc. This conversation is in Vicuna/ShareGPT format. Each set, along with code, has detailed explanation. 
I have released the [data](https://huggingface.co/datasets/ajibawa-2023/Python-Code-23k-ShareGPT).

**Training:**
Entire dataset was trained on Azure 4 x A100 80GB. For 3 epoch, training took 42 hours. DeepSpeed codebase was used for training purpose. This was trained on Llama-1 by Meta.

This is a full fine tuned model. Links for quantized models are given below.


**GPTQ GGML & AWQ**

GPTQ: [Link](https://huggingface.co/TheBloke/Python-Code-33B-GPTQ)

GGUF: [Link](https://huggingface.co/TheBloke/Python-Code-33B-GGUF)

AWQ: [Link](https://huggingface.co/TheBloke/Python-Code-33B-AWQ)


**Example Prompt:**
```
This is a conversation with your helpful AI assistant. AI assistant can generate Python Code along with necessary explanation.

Context
You are a helpful AI assistant.

USER: <prompt>
ASSISTANT:
```