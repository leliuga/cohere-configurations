---
license: cc-by-nc-4.0
tags:
- not-for-all-audiences
- nsfw
---

![image/png](https://cdn-uploads.huggingface.co/production/uploads/63ab1241ad514ca8d1430003/ysQGHLh1dd6I40rVK_jk2.png)

[HIGHLY EXPERIMENTAL]

(Sister model: https://huggingface.co/Undi95/Unholy-v1-12L-13B)

Use at your own risk, I'm not responsible for any usage of this model, don't try to do anything this model tell you to do.

Uncensored.

If you are censored, it's maybe because of keyword like "assistant", "Factual answer", or other "sweet words" like I call them that trigger the censoring accross all the layer of the model (since they're all trained on some of them in a way).

10L : This is a test project, uukuguy/speechless-llama2-luban-orca-platypus-13b and jondurbin/spicyboros-13b-2.2 was used for a merge, then, I deleted the first 10 layers to add 10 layers of MLewd at the beginning, trying to break all censoring possible, before merging the output with MLewd at 0.66 weight.

<!-- description start -->
## Description

This repo contains fp16 files of Unholy v1, an uncensored model.

<!-- description end -->
<!-- description start -->
## Models used

- uukuguy/speechless-llama2-luban-orca-platypus-13b
- jondurbin/spicyboros-13b-2.2
- Undi95/MLewd-L2-13B-v2-3

<!-- description end -->
<!-- prompt-template start -->
## Prompt template: Alpaca

```
Below is an instruction that describes a task. Write a response that appropriately completes the request.

### Instruction:
{prompt}

### Response:
```

Exemple:

![image/png](https://cdn-uploads.huggingface.co/production/uploads/63ab1241ad514ca8d1430003/jaZzEcPP0IET6_KX7J5Hm.png)