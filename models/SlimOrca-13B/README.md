---
language:
- en
license: cc-by-nc-nd-4.0
datasets:
- Open-Orca/SlimOrca
- ajibawa-2023/SlimOrca-ShareGPT
model-index:
- name: SlimOrca-13B
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
      value: 60.15
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ajibawa-2023/SlimOrca-13B
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
      value: 81.4
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ajibawa-2023/SlimOrca-13B
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
      value: 57.04
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ajibawa-2023/SlimOrca-13B
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
      value: 49.37
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ajibawa-2023/SlimOrca-13B
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
      value: 74.43
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ajibawa-2023/SlimOrca-13B
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
      value: 39.95
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=ajibawa-2023/SlimOrca-13B
      name: Open LLM Leaderboard
---

**SlimOrca-13B: A General Purpose Intelligent Model**


This Model is trained on refined version of SlimOrca made available by [Open-Orca](https://huggingface.co/Open-Orca) team. 
The idea was to check how this Model will perform in the absence of "system" prompt/instruction. 
This Model is very good in various types of General Purpose content generation such as Q&A (including multiple choice), Articles from Summary, Sentiment Analysis,
Context & Hypothesis, Reviews, Erotic story generation etc.
It can also generate Uncensored content. Kindly be careful while generating Uncensored content as you will be responsible for what you
generate.

It is trained on 517981 set of conversations. Each set having 2 conversations. I have shared this [data](https://huggingface.co/datasets/ajibawa-2023/SlimOrca-ShareGPT).

All the credit goes to the Open-Orca team for releasing SlimOrca dataset.


**Training:**
Entire dataset was trained on Azure 4 x A100 80GB. For 3 epoch, training took almost 11 Days. DeepSpeed codebase was used for training purpose.
Entire data is trained on Llama-2 by Meta.

This is a full fine tuned model. Links for quantized models are given below.

**GPTQ GGML & AWQ**

GPTQ: [Link](https://huggingface.co/TheBloke/SlimOrca-13B-GPTQ)

GGUF: [Link](https://huggingface.co/TheBloke/SlimOrca-13B-GGUF)

AWQ: [Link](https://huggingface.co/TheBloke/SlimOrca-13B-AWQ)


Special Thanks to [TheBloke](https://huggingface.co/TheBloke) for making these models available.



**Example Prompt:**
```
This is a conversation with your Assistant. It is a computer program designed to help you with various tasks such as answering questions, providing recommendations, and helping with decision making. You can ask it anything you want and it will do its best to give you accurate and relevant information.

Context
You are a helpful AI assistant.

USER: <prompt>
ASSISTANT:
```
You can modify above Prompt as per your requirement. I have used ShareGPT/Vicuna format v1.1 .


I want to say special Thanks to the Open Source community for helping & guiding me to better understand the AI/Model development.

Thank you for your love & support.


**Example Output**

Example 1

![Example 1](https://cdn-uploads.huggingface.co/production/uploads/64aea8ff67511bd3d965697b/hM_EJaSZiMjMQU35EiHGM.png)

Example 2

![Example 2](https://cdn-uploads.huggingface.co/production/uploads/64aea8ff67511bd3d965697b/riNaxJeTWdCEE4dNP8GWp.png)



# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_ajibawa-2023__SlimOrca-13B)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |60.39|
|AI2 Reasoning Challenge (25-Shot)|60.15|
|HellaSwag (10-Shot)              |81.40|
|MMLU (5-Shot)                    |57.04|
|TruthfulQA (0-shot)              |49.37|
|Winogrande (5-shot)              |74.43|
|GSM8k (5-shot)                   |39.95|

