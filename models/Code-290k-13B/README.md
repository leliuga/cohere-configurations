---
language:
- en
license: cc-by-nc-nd-4.0
tags:
- code
datasets:
- ajibawa-2023/Code-290k-ShareGPT
model-index:
- name: Code-290k-13B
  results:
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: AI2 Reasoning Challenge (25-Shot)
      type: ai2_arc
      config: ARC-Challenge
      split: test
      args:
        num_few_shot: 25
    metrics:
    - type: acc_norm
      value: 56.06
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ajibawa-2023/Code-290k-13B
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: HellaSwag (10-Shot)
      type: hellaswag
      split: validation
      args:
        num_few_shot: 10
    metrics:
    - type: acc_norm
      value: 81.55
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ajibawa-2023/Code-290k-13B
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: MMLU (5-Shot)
      type: cais/mmlu
      config: all
      split: test
      args:
        num_few_shot: 5
    metrics:
    - type: acc
      value: 51.99
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ajibawa-2023/Code-290k-13B
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: TruthfulQA (0-shot)
      type: truthful_qa
      config: multiple_choice
      split: validation
      args:
        num_few_shot: 0
    metrics:
    - type: mc2
      value: 37.65
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ajibawa-2023/Code-290k-13B
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: Winogrande (5-shot)
      type: winogrande
      config: winogrande_xl
      split: validation
      args:
        num_few_shot: 5
    metrics:
    - type: acc
      value: 72.69
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ajibawa-2023/Code-290k-13B
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: GSM8k (5-shot)
      type: gsm8k
      config: main
      split: test
      args:
        num_few_shot: 5
    metrics:
    - type: acc
      value: 17.82
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ajibawa-2023/Code-290k-13B
      name: Open LLM Leaderboard
---


**Code-290k-13B**

Large Language Models (LLMs) are good with code generations. Sometimes they do make mistakes in code generation. How about if they can give detailed explanation along with the code.
This is what I have tried over here. The base Llama-2 model was used for training purpose. It is trained on around **290000** set of codes. Each set having 2 conversations.
Along with Python, Java, JavaScript, GO, C++, Rust, Ruby, Sql, MySql, R, Julia, Haskell, etc. code with detailed explanation is used for training purpose. It is built upon using my existing Datasets [Python-Code-23k-ShareGPT](https://huggingface.co/datasets/ajibawa-2023/Python-Code-23k-ShareGPT) and [Code-74k-ShareGPT](https://huggingface.co/datasets/ajibawa-2023/Code-74k-ShareGPT) .
This conversation is in Vicuna/ShareGPT format. Each set, along with code, has detailed explanation. 

I have released the new data [Code-290k-ShareGPT](https://huggingface.co/datasets/ajibawa-2023/Code-290k-ShareGPT) on which this Model is trained.

**Training:**

Entire dataset was trained on 4 x A100 80GB. For 3 epoch, training took 165 hours. DeepSpeed codebase was used for training purpose. This was trained on Llama-2 by Meta.


This is a full fine tuned model. Links for quantized models are given below.


**GPTQ, GGUF, AWQ & Exllama**

GPTQ: [Link](https://huggingface.co/TheBloke/Code-290k-13B-GPTQ)

GGUF: [Link](https://huggingface.co/TheBloke/Code-290k-13B-GGUF)

AWQ: [Link](https://huggingface.co/TheBloke/Code-290k-13B-AWQ)

Exllama v2: [Link](https://huggingface.co/bartowski/Code-290k-13B-exl2)


Extremely thankful to [TheBloke](https://huggingface.co/TheBloke) and [Bartowski](https://huggingface.co/bartowski) for making Quantized versions of the model.


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

**Example Output**

Will update soon.
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_ajibawa-2023__Code-290k-13B)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |52.96|
|AI2 Reasoning Challenge (25-Shot)|56.06|
|HellaSwag (10-Shot)              |81.55|
|MMLU (5-Shot)                    |51.99|
|TruthfulQA (0-shot)              |37.65|
|Winogrande (5-shot)              |72.69|
|GSM8k (5-shot)                   |17.82|

