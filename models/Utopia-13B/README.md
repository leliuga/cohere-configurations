---
license: cc-by-nc-4.0
tags:
- not-for-all-audiences
- nsfw
---

<!-- description start -->
## Description

This repo contains fp16 files of Utopia-13B, a merge I have done with the new task_arithmetic merge method from mergekit.

<!-- description end -->
<!-- description start -->
## Models and loras used

- [Xwin-LM/Xwin-LM-13B-V0.2](https://huggingface.co/Xwin-LM/Xwin-LM-13B-V0.2)
- [NeverSleep/Nethena-13B](https://huggingface.co/NeverSleep/Nethena-13B)
- [PygmalionAI/pygmalion-2-13b](https://huggingface.co/PygmalionAI/pygmalion-2-13b)
- [Undi95/Storytelling-v2.1-13B-lora](https://huggingface.co/Undi95/Storytelling-v2.1-13B-lora)
- [zattio770/120-Days-of-LORA-v2-13B](https://huggingface.co/zattio770/120-Days-of-LORA-v2-13B)
- [lemonilia/LimaRP-Llama2-13B-v3-EXPERIMENT](https://huggingface.co/lemonilia/LimaRP-Llama2-13B-v3-EXPERIMENT)

<!-- description end -->
## The sauce
```
Xwin-LM/Xwin-LM-13B-V0.2
Undi95/Storytelling-v2.1-13B-lora
=> p1

NeverSleep/Nethena-13B
zattio770/120-Days-of-LORA-v2-13B
=> p2

PygmalionAI/pygmalion-2-13b
lemonilia/LimaRP-Llama2-13B-v3-EXPERIMENT
=> p3



merge_method: task_arithmetic
base_model: TheBloke/Llama-2-13B-fp16
models:
  - model: TheBloke/Llama-2-13B-fp16
  - model: Undi95/newpart1
    parameters:
      weight: 1.0
  - model: Undi95/newpart2
    parameters:
      weight: 0.45
  - model: Undi95/newpart3
    parameters:
      weight: 0.33
dtype: float16
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