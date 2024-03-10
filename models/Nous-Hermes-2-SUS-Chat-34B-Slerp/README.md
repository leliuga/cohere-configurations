---
license: other
tags:
- merge
license_name: yi-34b
license_link: https://huggingface.co/01-ai/Yi-34B/blob/main/LICENSE
model-index:
- name: Nous-Hermes-2-SUS-Chat-34B-Slerp
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
      value: 66.72
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Weyaxi/Nous-Hermes-2-SUS-Chat-34B-Slerp
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
      value: 84.97
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Weyaxi/Nous-Hermes-2-SUS-Chat-34B-Slerp
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
      value: 77.0
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Weyaxi/Nous-Hermes-2-SUS-Chat-34B-Slerp
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
      value: 59.23
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Weyaxi/Nous-Hermes-2-SUS-Chat-34B-Slerp
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
      value: 83.58
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Weyaxi/Nous-Hermes-2-SUS-Chat-34B-Slerp
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
      value: 72.86
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Weyaxi/Nous-Hermes-2-SUS-Chat-34B-Slerp
      name: Open LLM Leaderboard
---
![image/png](https://cdn-uploads.huggingface.co/production/uploads/6468ce47e134d050a58aa89c/u0jZZVpCxq3JN8VzFXVhV.png)

# Nous-Hermes-2-SUS-Chat-34B-Slerp

This is the model for Nous-Hermes-2-SUS-Chat-34B-Slerp. I used [mergekit](https://github.com/cg123/mergekit) to merge models.

# Prompt Templates

You can use these prompt templates, but I recommend using ChatML.

### ChatML [(NousResearch/Nous-Hermes-2-Yi-34B)](https://huggingface.co/NousResearch/Nous-Hermes-2-Yi-34B):

```
<|im_start|>system
{system}<|im_end|>
<|im_start|>user
{user}<|im_end|>
<|im_start|>assistant
{asistant}<|im_end|>
```

### Human - Asistant [(SUSTech/SUS-Chat-34B)](https://huggingface.co/SUSTech/SUS-Chat-34B):

```
### Human: {user}

### Assistant: {asistant}
```

# Yaml Config

```yaml

slices:
- sources:
    - model: Nous-Hermes-2-Yi-34B
      layer_range: [0, 60]
    - model: SUS-Chat-34B
      layer_range: [0, 60]

merge_method: slerp
base_model: Yi-34B

parameters:
t:
  - filter: self_attn
    value: [0, 0.5, 0.3, 0.7, 1]
  - filter: mlp
    value: [1, 0.5, 0.7, 0.3, 0]
  - value: 0.5
tokenizer_source: union
dtype: bfloat16

```

# Quantizationed versions

Quantizationed versions of this model is available thanks to [TheBloke](https://hf.co/TheBloke).

##### GPTQ

- [TheBloke/Nous-Hermes-2-SUS-Chat-34B-Slerp-GPTQ](https://huggingface.co/TheBloke/Nous-Hermes-2-SUS-Chat-34B-Slerp-GPTQ)

##### GGUF

- [TheBloke/Nous-Hermes-2-SUS-Chat-34B-Slerp-GGUF](https://huggingface.co/TheBloke/Nous-Hermes-2-SUS-Chat-34B-Slerp-GGUF)

##### AWQ

- [TheBloke/Nous-Hermes-2-SUS-Chat-34B-Slerp-AWQ](https://huggingface.co/TheBloke/Nous-Hermes-2-SUS-Chat-34B-Slerp-AWQ)
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_Weyaxi__Nous-Hermes-2-SUS-Chat-34B-Slerp)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |74.06|
|AI2 Reasoning Challenge (25-Shot)|66.72|
|HellaSwag (10-Shot)              |84.97|
|MMLU (5-Shot)                    |77.00|
|TruthfulQA (0-shot)              |59.23|
|Winogrande (5-shot)              |83.58|
|GSM8k (5-shot)                   |72.86|

