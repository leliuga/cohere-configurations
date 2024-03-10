---
language:
- en
license: apache-2.0
model-index:
- name: supermario-v2
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
      value: 68.52
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=janhq/supermario-v2
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
      value: 86.51
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=janhq/supermario-v2
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
      value: 64.88
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=janhq/supermario-v2
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
      value: 60.58
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=janhq/supermario-v2
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
      value: 81.37
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=janhq/supermario-v2
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
      value: 72.18
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=janhq/supermario-v2
      name: Open LLM Leaderboard
---
<!-- header start -->
<!-- 200823 -->
<div style="width: auto; margin-left: auto; margin-right: auto">
<img src="https://github.com/janhq/jan/assets/89722390/35daac7d-b895-487c-a6ac-6663daaad78e" alt="Jan banner" style="width: 100%; min-width: 400px; display: block; margin: auto;">
</div>

<p align="center">
    <a href="https://jan.ai/">Jan</a> 
    - <a href="https://discord.gg/AsJ8krTT3N">Discord</a>
</p>
<!-- header end -->

# Model Description
This model uses the `DARE_TIES` merge method from 3 models:
1. [OpenHermes-2.5-neural-chat-v3-3-Slerp](https://huggingface.co/Weyaxi/OpenHermes-2.5-neural-chat-v3-3-Slerp)
2. [MetaMath-Cybertron-Starling](https://huggingface.co/Q-bert/MetaMath-Cybertron-Starling)
3. [Marcoroni-7B-v3](https://huggingface.co/AIDC-ai-business/Marcoroni-7B-v3)
- base model: [Mistral-7B-v0.1](https://huggingface.co/mistralai/Mistral-7B-v0.1)

The yaml config file for this model here:
```yaml
base_model: mistralai/Mistral-7B-v0.1
dtype: bfloat16
merge_method: dare_ties
models:
- model: mistralai/Mistral-7B-v0.1
- model: Weyaxi/OpenHermes-2.5-neural-chat-v3-3-Slerp
  parameters:
    density: 0.8
    weight: 0.4
- model: Q-bert/MetaMath-Cybertron-Starling
  parameters:
    density: 0.8
    weight: 0.3
- model: AIDC-ai-business/Marcoroni-7B-v3
  parameters:
    density: 0.8
    weight: 0.3
parameters:
  int8_mask: true

```

# Prompt template:

- **ChatML**

```
<|im_start|>system
{system_message}<|im_end|>
<|im_start|>user
{prompt}<|im_end|>
<|im_start|>assistant
```
- **System**
```
### System:
{system}
### User:
{user}
### Assistant:
```

# Run this model
You can run this model using [Jan Desktop](https://jan.ai/) on Mac, Windows, or Linux.

Jan is an open source, ChatGPT alternative that is:

- 💻  **100% offline on your machine**: Your conversations remain confidential, and visible only to you.
- 🗂️ **An Open File Format**: Conversations and model settings stay on your computer and can be exported or deleted at any time.
- 🌐 **OpenAI Compatible**: Local server on port `1337` with OpenAI compatible endpoints
- 🌍 **Open Source & Free**: We build in public; check out our [Github](https://github.com/janhq)

![image/png](https://cdn-uploads.huggingface.co/production/uploads/65713d70f56f9538679e5a56/r7VmEBLGXpPLTu2MImM7S.png)

# About Jan
Jan believes in the need for an open-source AI ecosystem and is building the infra and tooling to allow open-source AIs to compete on a level playing field with proprietary ones.

Jan's long-term vision is to build a cognitive framework for future robots, who are practical, useful assistants for humans and businesses in everyday life.

# Jan Model Merger
This is a test project for merging models.

# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)

Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_janhq__supermario-v2).

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 72.36|
| ARC (25-shot)         | 68.52          |
| HellaSwag (10-shot)   | 86.51   |
| MMLU (5-shot)         | 64.88|
| TruthfulQA (0-shot)   | 60.58 |
| Winogrande (5-shot)   | 81.37  |
| GSM8K (5-shot)        | 72.18        |

# Acknowlegement
- [mergekit](https://github.com/cg123/mergekit)
- [DARE](https://github.com/yule-BUAA/MergeLM/blob/main/README.md)
- [lm-evaluation-harness
](https://github.com/EleutherAI/lm-evaluation-harness)
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_janhq__supermario-v2)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |72.34|
|AI2 Reasoning Challenge (25-Shot)|68.52|
|HellaSwag (10-Shot)              |86.51|
|MMLU (5-Shot)                    |64.88|
|TruthfulQA (0-shot)              |60.58|
|Winogrande (5-shot)              |81.37|
|GSM8k (5-shot)                   |72.18|

