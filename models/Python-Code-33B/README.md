---
language:
- en
license: cc-by-nc-nd-4.0
tags:
- code
datasets:
- ajibawa-2023/Python-Code-23k-ShareGPT
model-index:
- name: Python-Code-33B
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
      value: 56.31
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ajibawa-2023/Python-Code-33B
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
      value: 81.01
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ajibawa-2023/Python-Code-33B
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
      value: 54.22
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ajibawa-2023/Python-Code-33B
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
      value: 44.39
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ajibawa-2023/Python-Code-33B
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
      value: 75.22
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ajibawa-2023/Python-Code-33B
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
      value: 19.18
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ajibawa-2023/Python-Code-33B
      name: Open LLM Leaderboard
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
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_ajibawa-2023__Python-Code-33B)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |55.06|
|AI2 Reasoning Challenge (25-Shot)|56.31|
|HellaSwag (10-Shot)              |81.01|
|MMLU (5-Shot)                    |54.22|
|TruthfulQA (0-shot)              |44.39|
|Winogrande (5-shot)              |75.22|
|GSM8k (5-shot)                   |19.18|

