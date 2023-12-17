---
license: cc-by-nc-4.0
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
