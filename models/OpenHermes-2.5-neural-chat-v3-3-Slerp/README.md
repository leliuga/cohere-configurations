---
license: apache-2.0
model-index:
- name: OpenHermes-2.5-neural-chat-v3-3-Slerp
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
      value: 68.09
      name: normalized accuracy
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
      value: 86.2
      name: normalized accuracy
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
      value: 64.26
      name: accuracy
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
      value: 62.78
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
      value: 79.16
      name: accuracy
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
      value: 67.78
      name: accuracy
tags:
- merge
---
![image/png](https://cdn-uploads.huggingface.co/production/uploads/6468ce47e134d050a58aa89c/x44nNbPTpv0zGTqA1Jb2q.png)

# OpenHermes-2.5-neural-chat-v3-3-Slerp

This is the model for OpenHermes-2.5-neural-chat-v3-3-Slerp. I used [mergekit](https://github.com/cg123/mergekit) to merge models.

# Prompt Templates

You can use these prompt templates, but I recommend using ChatML.

### ChatML [(OpenHermes-2.5-Mistral-7B)](https://huggingface.co/teknium/OpenHermes-2.5-Mistral-7B):

```
<|im_start|>system
{system}<|im_end|>
<|im_start|>user
{user}<|im_end|>
<|im_start|>assistant
{asistant}<|im_end|>
```

### [neural-chat-7b-v3-3](https://huggingface.co/Intel/neural-chat-7b-v3-3):

```
### System:
{system}
### User:
{user}
### Assistant:
```

# Yaml Config to reproduce

```yaml
slices:
  - sources:
      - model: teknium/OpenHermes-2.5-Mistral-7B
        layer_range: [0, 32]
      - model: Intel/neural-chat-7b-v3-3
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

- [TheBloke/OpenHermes-2.5-neural-chat-v3-3-Slerp-GPTQ](https://huggingface.co/TheBloke/OpenHermes-2.5-neural-chat-v3-3-Slerp-GPTQ)

##### GGUF

- [TheBloke/OpenHermes-2.5-neural-chat-v3-3-Slerp-GGUF](https://huggingface.co/TheBloke/OpenHermes-2.5-neural-chat-v3-3-Slerp-GGUF)

##### AWQ

- [TheBloke/OpenHermes-2.5-neural-chat-v3-3-Slerp-AWQ](https://huggingface.co/TheBloke/OpenHermes-2.5-neural-chat-v3-3-Slerp-AWQ)


# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)

Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_PulsarAI__OpenHermes-2.5-neural-chat-v3-3-Slerp)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 71.38   |
| ARC (25-shot)         | 68.09          |
| HellaSwag (10-shot)   | 86.2    |
| MMLU (5-shot)         | 64.26         |
| TruthfulQA (0-shot)   | 62.78  |
| Winogrande (5-shot)   | 79.16  |
| GSM8K (5-shot)        | 67.78        |