---
license: cc-by-nc-4.0
tags:
- not-for-all-audiences
- nsfw
---

![image/png](https://cdn-uploads.huggingface.co/production/uploads/63ab1241ad514ca8d1430003/VjlbZcxzuvMjwOjnYddRK.png)

THIS MODEL IS MADE FOR LEWD

SEXUAL, CRUDE AND KINKY CONTENT IN OUTPUT CAN AND WILL HAPPEN. YOU'RE WARNED 

This is MLewd merged with [Xwin-LM/Xwin-LM-13B-V0.2](https://huggingface.co/Xwin-LM/Xwin-LM-13B-V0.2)

<!-- description start -->
## Description

This repo contains fp16 files of Xwin-MLewd-13B-V0.2, very hot and lewd model based on Xwin 0.2 13B.

<!-- description end -->
<!-- description start -->
## Models and loras used

- Undi95/ReMM-S-Light (base/private)
- Undi95/CreativeEngine
- Brouz/Slerpeno
- The-Face-Of-Goonery/Huginn-v3-13b
- zattio770/120-Days-of-LORA-v2-13B
- PygmalionAI/pygmalion-2-13b
- Undi95/StoryTelling
- TokenBender/sakhi_13B_roleplayer_NSFW_chat_adapter
- nRuaif/Kimiko-v2-13B
- The-Face-Of-Goonery/Huginn-13b-FP16
- lemonilia/LimaRP-Llama2-13B-v3-EXPERIMENT
- Xwin-LM/Xwin-LM-13B-V0.2

<!-- description end -->
<!-- prompt-template start -->
## Prompt template: Alpaca

```
Below is an instruction that describes a task. Write a response that appropriately completes the request.

### Instruction:
{prompt}

### Response:

```
## The secret sauce

```
slices:
  - sources:
      - model: Xwin-LM/Xwin-LM-13B-V0.2
        layer_range: [0, 40]
      - model: Undi95/MLewd-v2.4-13B
        layer_range: [0, 40]
merge_method: slerp
base_model: Xwin-LM/Xwin-LM-13B-V0.2
parameters:
  t:
    - filter: lm_head 
      value: [0.55]
    - filter: embed_tokens
      value: [0.7]
    - filter: self_attn
      value: [0.65, 0.35]
    - filter: mlp
      value:  [0.35, 0.65]
    - filter: layernorm
      value: [0.4, 0.6]
    - filter: modelnorm
      value: [0.6]
    - value: 0.5 # fallback for rest of tensors
dtype: float16
```

Special thanks to Sushi and Shena ♥

If you want to support me, you can [here](https://ko-fi.com/undiai).