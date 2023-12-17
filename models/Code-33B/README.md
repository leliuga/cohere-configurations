---
license: cc-by-nc-nd-4.0
datasets:
- ajibawa-2023/Code-74k-ShareGPT
language:
- en
tags:
- code
---

**Code-33B**

Large Language Models (LLMs) are good with code generations. Sometimes they do make mistakes in code generation. How about if they can give detailed explanation along with the code.
This is what I have tried over here. The base Llama-1 model was used for training purpose. It is trained on around 74000 set of codes. Each set having 2 conversations.
Along with Python, Java, JavaScript, GO, C++, Rust etc. code with detailed explanation is used for training purpose. It is built upon using my existing Dataset [Python-Code-23k-ShareGPT](https://huggingface.co/datasets/ajibawa-2023/Python-Code-23k-ShareGPT).
This conversation is in Vicuna/ShareGPT format. Each set, along with code, has detailed explanation. 

I have released the new data [Code-74k-ShareGPT](https://huggingface.co/datasets/ajibawa-2023/Code-74k-ShareGPT) on which this Model is trained.

**Training:**

Entire dataset was trained on 4 x A100 80GB. For 3 epoch, training took 6 days & 5 hours. DeepSpeed codebase was used for training purpose. This was trained on Llama-1 by Meta.


This is a full fine tuned model. Links for quantized models are given below.


**GPTQ GGUF & AWQ**

GPTQ: [Link](https://huggingface.co/TheBloke/Code-33B-GPTQ)

GGUF: [Link](https://huggingface.co/TheBloke/Code-33B-GGUF)

AWQ: [Link](https://huggingface.co/TheBloke/Code-33B-AWQ)


Extremely thankful to [TheBloke](https://huggingface.co/TheBloke) for making Quantized versions of model.


**Example Prompt:**
```
This is a conversation with your helpful AI assistant. AI assistant can generate Code in various Programming Languages along with necessary explanation.

Context
You are a helpful AI assistant.

USER: <prompt>
ASSISTANT:
```

You can modify above Prompt as per your requirement. I have used ShareGPT/Vicuna format v1.1 .

I want to say special Thanks to the Open Source community for helping & guiding me to better understand the AI/Model development.

Thank you for your love & support.
