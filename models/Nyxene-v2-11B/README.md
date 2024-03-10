---
license: cc-by-nc-4.0
tags:
- merge
model-index:
- name: Nyxene-v2-11B
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
      value: 67.41
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=beberik/Nyxene-v2-11B
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
      value: 84.54
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=beberik/Nyxene-v2-11B
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
      value: 65.26
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=beberik/Nyxene-v2-11B
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
      value: 55.62
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=beberik/Nyxene-v2-11B
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
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=beberik/Nyxene-v2-11B
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
      value: 54.66
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=beberik/Nyxene-v2-11B
      name: Open LLM Leaderboard
---

## Description

This repo contains bf16 files of Nyxene-v2-11B. It feels like with the new models, 1% is no longer needed as in the [previous version](https://huggingface.co/beberik/Nyxene-v1-11B). And yes, new version. Again.

## Model used
- [berkeley-nest/Starling-LM-7B-alpha](https://huggingface.co/berkeley-nest/Starling-LM-7B-alpha)
- [openaccess-ai-collective/DPOpenHermes-7B](https://huggingface.co/openaccess-ai-collective/DPOpenHermes-7B)
- [fblgit/fblgit/una-cybertron-7b-v2](https://huggingface.co/fblgit/una-cybertron-7b-v2-bf16)
- [chargoddard/loyal-piano-m7-cdpo](https://huggingface.co/chargoddard/loyal-piano-m7-cdpo)

## Prompt template

The best one after further testing is this one:

```
<|system|>
Below is an instruction that describes a task. Write a response that appropriately completes the request.
<|user|>
{prompt}
<|assistant|>
```

## The secret sauce

loyal-piano-cybertron-11B :
```
slices:
  - sources:
    - model: fblgit/una-cybertron-7b-v2
      layer_range: [0, 24]
  - sources:
    - model: chargoddard/loyal-piano-m7-cdpo
      layer_range: [8, 32]
merge_method: passthrough
dtype: bfloat16
```

Starling-DPOHermes-11B :
```
slices:
  - sources:
    - model: berkeley-nest/Starling-LM-7B-alpha
      layer_range: [0, 24]
  - sources:
    - model: openaccess-ai-collective/DPOpenHermes-7B
      layer_range: [8, 32]
merge_method: passthrough
dtype: bfloat16
```

Nyxene-11B :
```
slices:
  - sources:
      - model: loyal-piano-cybertron-11B
        layer_range: [0, 48]
      - model: Starling-NeuralHermes-11B
        layer_range: [0, 48]
merge_method: slerp
base_model: loyal-piano-cybertron-11B
parameters:
  t:
    - filter: lm_head 
      value: [0.75]
    - filter: embed_tokens
      value: [0.75]
    - filter: self_attn
      value: [0.75, 0.25]
    - filter: mlp
      value:  [0.25, 0.75]
    - filter: layernorm
      value: [0.5, 0.5]
    - filter: modelnorm
      value: [0.75]
    - value: 0.5 # fallback for rest of tensors
dtype: bfloat16
```
I use [mergekit](https://github.com/cg123/mergekit) for all the manipulation told here.

Thanks to the [Undi95](https://huggingface.co/Undi95) for the original [11B mistral merge](https://huggingface.co/Undi95/Mistral-11B-OmniMix) recipe.

# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_beberik__Nyxene-v2-11B)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |67.84|
|AI2 Reasoning Challenge (25-Shot)|67.41|
|HellaSwag (10-Shot)              |84.54|
|MMLU (5-Shot)                    |65.26|
|TruthfulQA (0-shot)              |55.62|
|Winogrande (5-shot)              |79.56|
|GSM8k (5-shot)                   |54.66|

