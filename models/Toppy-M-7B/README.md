---
license: cc-by-nc-4.0
tags:
- not-for-all-audiences
- nsfw
---

<!-- description start -->
## Description

This repo contains fp16 files of Toppy-M-7B, a merge I have done with the new task_arithmetic merge method from mergekit.

This project was a request from [BlueNipples](https://huggingface.co/BlueNipples) : [link](https://huggingface.co/Undi95/Utopia-13B/discussions/1)

<!-- description end -->
<!-- description start -->
## Models and loras used

- [openchat/openchat_3.5](https://huggingface.co/openchat/openchat_3.5)
- [NousResearch/Nous-Capybara-7B-V1.9](https://huggingface.co/NousResearch/Nous-Capybara-7B-V1.9)
- [HuggingFaceH4/zephyr-7b-beta](https://huggingface.co/HuggingFaceH4/zephyr-7b-beta)
- [lemonilia/AshhLimaRP-Mistral-7B](lemonilia/AshhLimaRP-Mistral-7B)
- [Vulkane/120-Days-of-Sodom-LoRA-Mistral-7b](https://huggingface.co/Vulkane/120-Days-of-Sodom-LoRA-Mistral-7b)
- [Undi95/Mistral-pippa-sharegpt-7b-qlora](Undi95/Mistral-pippa-sharegpt-7b-qlora)

<!-- description end -->
## The sauce
```
openchat/openchat_3.5
lemonilia/AshhLimaRP-Mistral-7B (LoRA) x 0.38

NousResearch/Nous-Capybara-7B-V1.9
Vulkane/120-Days-of-Sodom-LoRA-Mistral-7b x 0.27

HuggingFaceH4/zephyr-7b-beta
Undi95/Mistral-pippa-sharegpt-7b-qlora x 0.38

merge_method: task_arithmetic
base_model: mistralai/Mistral-7B-v0.1
models:
  - model: mistralai/Mistral-7B-v0.1
  - model: Undi95/zephyr-7b-beta-pippa-sharegpt
    parameters:
      weight: 0.42
  - model: Undi95/Nous-Capybara-7B-V1.9-120-Days
    parameters:
      weight: 0.29
  - model: Undi95/openchat_3.5-LimaRP-13B
    parameters:
      weight: 0.48
dtype: bfloat16
```
<!-- prompt-template start -->
## Prompt template: Alpaca

```
Below is an instruction that describes a task. Write a response that appropriately completes the request.

### Instruction:
{prompt}

### Response:

```

If you want to support me, you can [here](https://ko-fi.com/undiai).