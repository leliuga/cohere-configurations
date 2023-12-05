---
license: cc-by-nc-4.0
tags:
- not-for-all-audiences
- nsfw
---


![image/png](https://cdn-uploads.huggingface.co/production/uploads/63ab1241ad514ca8d1430003/bex2BCxY5n30mSsxvM-WW.png)


Attempt to recreate Amethyst-13B but in 20B. The two model was glued together with a sparkle of [Unholy](Undi95/Unholy-v1-12L-13B).

In addition, [LimaRP v3](https://huggingface.co/lemonilia/LimaRP-Llama2-13B-v3-EXPERIMENT) was used, is it recommanded to read the documentation.

<!-- description start -->
## Description

This repo contains fp16 files of Amethyst-20B.

<!-- description end -->
<!-- description start -->
## Models and loras used

- Xwin-LM/Xwin-LM-13B-V0.1
- The-Face-Of-Goonery/Huginn-13b-FP16
- zattio770/120-Days-of-LORA-v2-13B
- lemonilia/LimaRP-Llama2-13B-v3-EXPERIMENT
- Undi95/Unholy-v1-12L-13B

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
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_Undi95__U-Amethyst-20B)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 50.38   |
| ARC (25-shot)         | 62.2          |
| HellaSwag (10-shot)   | 83.11    |
| MMLU (5-shot)         | 55.88         |
| TruthfulQA (0-shot)   | 53.2   |
| Winogrande (5-shot)   | 74.19   |
| GSM8K (5-shot)        | 5.31        |
| DROP (3-shot)         | 18.75         |
