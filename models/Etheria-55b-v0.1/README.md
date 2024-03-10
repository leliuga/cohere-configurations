---
license: apache-2.0
tags:
- mergekit
- Etheria
base_model: []
model-index:
- name: Etheria-55b-v0.1
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
      value: 65.1
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Steelskull/Etheria-55b-v0.1
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
      value: 81.93
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Steelskull/Etheria-55b-v0.1
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
      value: 73.66
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Steelskull/Etheria-55b-v0.1
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
      value: 56.16
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Steelskull/Etheria-55b-v0.1
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
      value: 76.09
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Steelskull/Etheria-55b-v0.1
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
      value: 35.18
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=Steelskull/Etheria-55b-v0.1
      name: Open LLM Leaderboard
---
# Steelskull/Etheria-55b-v0.1

![image/png](https://cdn-uploads.huggingface.co/production/uploads/64545af5ec40bbbd01242ca6/RAhrbktyyVQxOR1np-9L2.png)

## Merge Details

An attempt to make a functional goliath style merge to create a [Etheria] 55b-200k with two yi-34b-200k models.

due to the merge it 'theoretically' should have a context of 200k but I recommend starting at 32k and moveing up,
as it is unknown (at this time) what the merge has done to the context length.

This is a merge of both VerA and VerB of Etheria-55b (There numbers were surprisingly good), I then created a sacrificial 55B out of the most performant yi-34b-200k Model
and performed a Dare_ties merge and equalize the model into its current state.

### recommended settings and Prompt Format:

Ive tested it up to 32k context using exl2 using these settings:

```
    "temp": 0.7,
    "temperature_last": true,
    "top_p": 1,
    "top_k": 0,
    "top_a": 0,
    "tfs": 1,
    "epsilon_cutoff": 0,
    "eta_cutoff": 0,
    "typical_p": 1,
    "min_p": 0.1,
    "rep_pen": 1.1,
    "rep_pen_range": 8192,
    "no_repeat_ngram_size": 0,
    "penalty_alpha": 0,
    "num_beams": 1,
    "length_penalty": 1,
    "min_length": 0,
    "encoder_rep_pen": 1,
    "freq_pen": 0,
    "presence_pen": 0,
    "do_sample": true,
    "early_stopping": false,
    "add_bos_token": false,
    "truncation_length": 2048,
    "ban_eos_token": true,
    "skip_special_tokens": true,
    "streaming": true,
    "mirostat_mode": 0,
    "mirostat_tau": 5,
    "mirostat_eta": 0.1,
```

Prompt format that work well
```
ChatML & Alpaca
```

### Merge Method

This model was merged using the [DARE](https://arxiv.org/abs/2311.03099) [TIES](https://arxiv.org/abs/2306.01708) merge method using Merged-Etheria-55b as a base.

### Configuration

The following YAML configuration was used to produce this model:

```yaml

base_model: Merged-Etheria-55b
models:
  - model: Sacr-Etheria-55b
    parameters:
      weight: [0.22, 0.113, 0.113, 0.113, 0.113, 0.113]
      density: 0.61
  - model: Merged-Etheria-55b
    parameters:
      weight: [0.22, 0.113, 0.113, 0.113, 0.113, 0.113]
      density: 0.61
merge_method: dare_ties
tokenizer_source: union
parameters:
  int8_mask: true
dtype: bfloat16

```
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_Steelskull__Etheria-55b-v0.1)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |64.69|
|AI2 Reasoning Challenge (25-Shot)|65.10|
|HellaSwag (10-Shot)              |81.93|
|MMLU (5-Shot)                    |73.66|
|TruthfulQA (0-shot)              |56.16|
|Winogrande (5-shot)              |76.09|
|GSM8k (5-shot)                   |35.18|

