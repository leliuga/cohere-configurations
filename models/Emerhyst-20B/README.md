---
license: cc-by-nc-4.0
tags:
- not-for-all-audiences
- nsfw
---

![image/png](https://cdn-uploads.huggingface.co/production/uploads/63ab1241ad514ca8d1430003/mvc3UyLtqKdLY1wzAdB_O.png)

Merge of [Amethyst 13B](https://huggingface.co/Undi95/Amethyst-13B) and [Emerald 13B](https://huggingface.co/Undi95/Emerald-13B).

In addition, [LimaRP v3](https://huggingface.co/lemonilia/LimaRP-Llama2-13B-v3-EXPERIMENT) was used, is it recommanded to read the documentation.

<!-- description start -->
## Description

This repo contains fp16 files of Emerhyst-20B.

<!-- description end -->
<!-- description start -->
## Models and loras used

- PygmalionAI/pygmalion-2-13b
- Xwin-LM/Xwin-LM-13B-V0.1
- The-Face-Of-Goonery/Huginn-13b-FP16
- zattio770/120-Days-of-LORA-v2-13B
- lemonilia/LimaRP-Llama2-13B-v3-EXPERIMENT

<!-- description end -->
<!-- prompt-template start -->
## Prompt template: Alpaca

```
Below is an instruction that describes a task. Write a response that appropriately completes the request.

### Instruction:
{prompt}

### Response:

```

## LimaRP v3 usage and suggested settings

![image/png](https://cdn-uploads.huggingface.co/production/uploads/63ab1241ad514ca8d1430003/ZC_iP2KkcEcRdgG_iyxYE.png)

You can follow these instruction format settings in SillyTavern. Replace tiny with your desired response length:

![image/png](https://cdn-uploads.huggingface.co/production/uploads/63ab1241ad514ca8d1430003/PIn8_HSPTJEMdSEpNVSdm.png)

Special thanks to Sushi.

If you want to support me, you can [here](https://ko-fi.com/undiai).
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_Undi95__Emerhyst-20B)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 51.85   |
| ARC (25-shot)         | 61.69          |
| HellaSwag (10-shot)   | 84.98    |
| MMLU (5-shot)         | 56.98         |
| TruthfulQA (0-shot)   | 54.16   |
| Winogrande (5-shot)   | 76.09   |
| GSM8K (5-shot)        | 8.49        |
| DROP (3-shot)         | 20.56         |
