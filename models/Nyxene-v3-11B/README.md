---
license: cc-by-nc-4.0
tags: 
 - merge
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
