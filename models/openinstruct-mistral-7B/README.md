---
language:
- en
license: apache-2.0
library_name: transformers
datasets:
- monology/VMware-open-instruct-higgsfield
pipeline_tag: text-generation
model-index:
- name: openinstruct-mistral-7b
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
      value: 59.73
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=monology/openinstruct-mistral-7b
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
      value: 82.77
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=monology/openinstruct-mistral-7b
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
      value: 60.55
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=monology/openinstruct-mistral-7b
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
      value: 48.76
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=monology/openinstruct-mistral-7b
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
      value: 79.56
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=monology/openinstruct-mistral-7b
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
      value: 50.49
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=monology/openinstruct-mistral-7b
      name: Open LLM Leaderboard
---

# OpenInstruct Mistral-7B
**1st among commercially-usable 7B models on the Open LLM Leaderboard!\***

This is [mistralai/Mistral-7B-v0.1](https://huggingface.co/mistralai/Mistral-7B-v0.1) finetuned on [VMware/open-instruct](https://huggingface.co/datasets/VMware/open-instruct).  
Quantized to FP16 and released under the [Apache-2.0](https://choosealicense.com/licenses/apache-2.0) license by myself.  
Compute generously provided by [Higgsfield AI](https://higgsfield.ai/model/655559e6b5777dab620095e0).  


## Prompt format: Alpaca
```
Below is an instruction that describes a task. Write a response that appropriately completes the request.

### Instruction:
[your instruction goes here]

### Response:
```

## Recommended preset:
- temperature: 0.2
- top_k: 50
- top_p 0.95
- repetition_penalty: 1.1

\*as of 21 Nov 2023. "commercially-usable" includes both an open-source base model and a *non-synthetic* open-source finetune dataset. updated leaderboard results available [here](https://huggingfaceh4-open-llm-leaderboard.hf.space).
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_monology__openinstruct-mistral-7b)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |63.64|
|AI2 Reasoning Challenge (25-Shot)|59.73|
|HellaSwag (10-Shot)              |82.77|
|MMLU (5-Shot)                    |60.55|
|TruthfulQA (0-shot)              |48.76|
|Winogrande (5-shot)              |79.56|
|GSM8k (5-shot)                   |50.49|

