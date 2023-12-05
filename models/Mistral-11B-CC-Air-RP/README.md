---
license: cc-by-nc-4.0
tags:
- not-for-all-audiences
- nsfw
- mistral
- pretrained
---

CollectiveCognition-v1.1-Mistral-7B and airoboros-mistral2.2-7b glued together and finetuned with qlora of Pippa and LimaRPv3 dataset.

<!-- description start -->
## Description

This repo contains fp16 files of Mistral-11B-CC-Air-RP.

<!-- description end -->
<!-- description start -->
## Model used

- [CollectiveCognition-v1.1-Mistral-7B](https://huggingface.co/teknium/CollectiveCognition-v1.1-Mistral-7B)
- [airoboros-mistral2.2-7b](https://huggingface.co/teknium/airoboros-mistral2.2-7b/)
- PIPPA dataset 11B qlora
- LimaRPv3 dataset 11B qlora

<!-- description end -->
<!-- prompt-template start -->
## Prompt template: Alpaca or default

```
Below is an instruction that describes a task. Write a response that appropriately completes the request.

### Instruction:
{prompt}

### Response:

```

```
USER: <prompt>
ASSISTANT:
```

## The secret sauce

```
slices:
  - sources:
    - model: teknium/CollectiveCognition-v1.1-Mistral-7B
      layer_range: [0, 24]
  - sources:
    - model: teknium/airoboros-mistral2.2-7b
      layer_range: [8, 32]
merge_method: passthrough
dtype: float16
```


Special thanks to Sushi.

If you want to support me, you can [here](https://ko-fi.com/undiai).