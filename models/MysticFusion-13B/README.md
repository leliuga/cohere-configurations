---
language:
- en
license: llama2
pipeline_tag: text-generation
inference: false
model-index:
- name: MysticFusion-13B
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
      value: 61.35
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Walmart-the-bag/MysticFusion-13B
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
      value: 84.43
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Walmart-the-bag/MysticFusion-13B
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
      value: 57.29
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Walmart-the-bag/MysticFusion-13B
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
      value: 51.98
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Walmart-the-bag/MysticFusion-13B
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
      value: 76.01
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Walmart-the-bag/MysticFusion-13B
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
      value: 24.79
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Walmart-the-bag/MysticFusion-13B
      name: Open LLM Leaderboard
---
# MysticFusion-13B
![img1](00117-3333234138.png)
YAML:
```
models:
  - model: KoboldAI/LLaMA2-13B-Tiefighter
    parameters:
      weight: 0.3
  - model: NeverSleep/Noromaid-13b-v0.1.1
    parameters:
      weight: 0.5
  - model: lmsys/vicuna-13b-v1.5
    parameters:
      weight: 0.2
merge_method: linear
dtype: float16
```
## Usage:
This is meant to be story writing and basic instruction. More of story writing.

## Prompt Template:
### Alpaca
```
### Instruction:

### Response:

```
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_Walmart-the-bag__MysticFusion-13B)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |59.31|
|AI2 Reasoning Challenge (25-Shot)|61.35|
|HellaSwag (10-Shot)              |84.43|
|MMLU (5-Shot)                    |57.29|
|TruthfulQA (0-shot)              |51.98|
|Winogrande (5-shot)              |76.01|
|GSM8k (5-shot)                   |24.79|

