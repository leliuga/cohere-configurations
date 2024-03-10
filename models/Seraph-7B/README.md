---
license: cc-by-nc-4.0
model-index:
- name: Seraph-7B
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
      value: 67.83
      name: normalized accuracy
    source:
      url: >-
        https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Weyaxi/Seraph-7B
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
      value: 86.22
      name: normalized accuracy
    source:
      url: >-
        https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Weyaxi/Seraph-7B
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
      value: 65.07
      name: accuracy
    source:
      url: >-
        https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Weyaxi/Seraph-7B
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
      value: 59.49
    source:
      url: >-
        https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Weyaxi/Seraph-7B
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
      value: 80.66
      name: accuracy
    source:
      url: >-
        https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Weyaxi/Seraph-7B
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
      value: 71.87
      name: accuracy
    source:
      url: >-
        https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Weyaxi/Seraph-7B
      name: Open LLM Leaderboard
tags:
- merge
---

![image/png](https://cdn-uploads.huggingface.co/production/uploads/6468ce47e134d050a58aa89c/ddzjZ1irvtLcDRCWei9vQ.png)

# Seraph-7B

This is the model for Seraph-7B. I used [mergekit](https://github.com/cg123/mergekit) to merge models.

# Prompt Templates

You can use these prompt templates, but I recommend using ChatML.

### ChatML:

```
<|im_start|>system
{system}<|im_end|>
<|im_start|>user
{user}<|im_end|>
<|im_start|>assistant
{asistant}<|im_end|>
```

### System, User, Asistant Alpaca Style:

```
### System:
{system}
### User:
{user}
### Assistant:
```

# Yaml Config

```yaml
slices:
  - sources:
      - model: Weyaxi/OpenHermes-2.5-neural-chat-v3-3-Slerp
        layer_range: [0, 32]
      - model: Q-bert/MetaMath-Cybertron-Starling
        layer_range: [0, 32]
merge_method: slerp
base_model: mistralai/Mistral-7B-v0.1
parameters:
  t:
    - filter: self_attn
      value: [0, 0.5, 0.3, 0.7, 1]
    - filter: mlp
      value: [1, 0.5, 0.7, 0.3, 0]
    - value: 0.5 # fallback for rest of tensors
dtype: bfloat16
```

# Quantizationed versions

Quantizationed versions of this model is available thanks to [TheBloke](https://hf.co/TheBloke).

##### GPTQ

- [TheBloke/Seraph-7B-GPTQ](https://huggingface.co/TheBloke/Seraph-7B-GPTQ)

##### GGUF

- [TheBloke/Seraph-7B-GGUF](https://huggingface.co/TheBloke/Seraph-7B-GGUF)

##### AWQ

- [TheBloke/Seraph-7B-AWQ](https://huggingface.co/TheBloke/Seraph-7B-AWQ)


# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)

Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_Weyaxi__Seraph-7B)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 71.86   |
| ARC (25-shot)         | 67.83          |
| HellaSwag (10-shot)   | 86.22   |
| MMLU (5-shot)         | 65.07|
| TruthfulQA (0-shot)   | 59.49 |
| Winogrande (5-shot)   | 80.66  |
| GSM8K (5-shot)        | 71.87        |

If you would like to support me:

[☕ Buy Me a Coffee](https://www.buymeacoffee.com/weyaxi)