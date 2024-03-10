---
license: cc-by-nc-4.0
tags:
- merge
model-index:
- name: Nyxene-v3-11B
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
      value: 69.62
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=beberik/Nyxene-v3-11B
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
      value: 85.33
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=beberik/Nyxene-v3-11B
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
      value: 64.75
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=beberik/Nyxene-v3-11B
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
      value: 60.91
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=beberik/Nyxene-v3-11B
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
      value: 80.19
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=beberik/Nyxene-v3-11B
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
      value: 63.53
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=beberik/Nyxene-v3-11B
      name: Open LLM Leaderboard
---

## Description

This repo contains bf16 files of Nyxene-v1-11B. Just new version with some new things.

## Model used
- [Intel/neural-chat-7b-v3-3-Slerp](https://huggingface.co/Intel/neural-chat-7b-v3-3-Slerp)
- [AIDC-ai-business/Marcoroni-7B-v3](https://huggingface.co/AIDC-ai-business/Marcoroni-7B-v3)
- [rwitz/go-bruins-v2](https://huggingface.co/rwitz/go-bruins-v2)
- [chargoddard/loyal-piano-m7-cdpo](https://huggingface.co/chargoddard/loyal-piano-m7-cdpo)

## Prompt template

Just use chatml. 

## The secret sauce

go-bruins-loyal-piano-11B :
```
slices:
  - sources:
    - model: rwitz/go-bruins-v2
      layer_range: [0, 24]
  - sources:
    - model: chargoddard/loyal-piano-m7-cdpo
      layer_range: [8, 32]
merge_method: passthrough
dtype: bfloat16
```

neural-marcoroni-11B :
```
slices:
  - sources:
    - model: AIDC-ai-business/Marcoroni-7B-v3
      layer_range: [0, 24]
  - sources:
    - model: Intel/neural-chat-7b-v3-3-Slerp
      layer_range: [8, 32]

merge_method: passthrough
dtype: bfloat16
```

Nyxene-11B :
```
slices:
  - sources:
      - model: "./go-bruins-loyal-piano-11B"
        layer_range: [0, 48]
      - model: "./neural-marcoroni-11B"
        layer_range: [0, 48]
merge_method: slerp
base_model: "./go-bruins-loyal-piano-11B"
parameters:
  t:
    - filter: lm_head 
      value: [0.5]
    - filter: embed_tokens
      value: [0.75]
    - filter: self_attn
      value: [0.75, 0.25]
    - filter: mlp
      value:  [0.25, 0.75]
    - filter: layernorm
      value: [0.5, 0.5]
    - filter: modelnorm
      value: [0.5]
    - value: 0.5 # fallback for rest of tensors
dtype: bfloat16
```
I use [mergekit](https://github.com/cg123/mergekit) for all the manipulation told here.

Thanks to the [Undi95](https://huggingface.co/Undi95) for the original [11B mistral merge](https://huggingface.co/Undi95/Mistral-11B-OmniMix) recipe.

# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_beberik__Nyxene-v3-11B)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |70.72|
|AI2 Reasoning Challenge (25-Shot)|69.62|
|HellaSwag (10-Shot)              |85.33|
|MMLU (5-Shot)                    |64.75|
|TruthfulQA (0-shot)              |60.91|
|Winogrande (5-shot)              |80.19|
|GSM8k (5-shot)                   |63.53|

