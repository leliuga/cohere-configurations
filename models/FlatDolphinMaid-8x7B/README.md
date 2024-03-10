---
license: cc-by-nc-4.0
tags:
- not-for-all-audiences
- nsfw
---

First experimental merge of Noromaid 8x7b (Instruct) and dolphin 8x7b. The idea behind this is to add a little more IQ to the model, because Noromaid was only trained on RP/ERP data. Dolphin 2.7 is the only real Mixtral finetune I consider "usable", and so the merging quest begin again kek.

Merged Dolphin 2.7 with Mixtral Base (Dolphin was at 1.0 weight) to get rid of ChatLM, and then I merged Noromaid 8x7b with the output, SLERP method.

This model feel better on the IQ chart and have the ~same average ERP score on ayumi bench' than Noromaid 8x7b, but it's softer and more prude too, it also have the typical Mixtral repeat issue at some point. Choose your poison.

![image/png](https://cdn-uploads.huggingface.co/production/uploads/63ab1241ad514ca8d1430003/uZlU0PEPtKPZPLzXcoqJ_.png)

<!-- description start -->
## Description

This repo contains fp16 files of FlatDolphinMaid-8x7B.

<!-- description end -->
<!-- description start -->
## Models used

- [mistralai/Mixtral-8x7B-v0.1](https://huggingface.co/mistralai/Mixtral-8x7B-v0.1)
- [cognitivecomputations/dolphin-2.7-mixtral-8x7b](https://huggingface.co/cognitivecomputations/dolphin-2.7-mixtral-8x7b)
- [NeverSleep/Noromaid-v0.1-mixtral-8x7b-Instruct-v3](https://huggingface.co/NeverSleep/Noromaid-v0.1-mixtral-8x7b-Instruct-v3)

<!-- description end -->
<!-- prompt-template start -->
### Custom format:
```
### Instruction:
{system prompt}

### Input:
{input}

### Response:
{reply}
```

If you want to support me, you can [here](https://ko-fi.com/undiai).