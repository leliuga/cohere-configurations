---
license: cc-by-4.0
language:
- en
tags:
- merge
---
<!-- header start -->
# Model Description
This model uses the `Slerp` merge method from 2 models:
1. [openchat/openchat-3.5-1210](https://huggingface.co/openchat/openchat-3.5-1210)
2. [berkeley-nest/Starling-LM-7B-alpha](https://huggingface.co/berkeley-nest/Starling-LM-7B-alpha)

- base model: [openchat/openchat-3.5-1210](https://huggingface.co/openchat/openchat-3.5-1210)

I SLERPed these two together because they're both OpenChat-ish models. Fundamentally, OpenChat-3.5-1210 appears to be trained similarly to OpenChat-3.5 but now with [Feedback-Collection](https://huggingface.co/datasets/kaist-ai/Feedback-Collection)
and [a de-contaminated Capybara](https://huggingface.co/datasets/LDJnr/Capybara). Starling is OpenChat-3.5 but trained with a novel training method on the Nectar set.

My hope is that a SLERP between the two retains the benefits of both.

The yaml config file for this model is here:
```yaml
slices:
  - sources:
      - model: openchat/openchat-3.5-1210 
        layer_range: [0, 32] 
      - model: berkeley-nest/Starling-LM-7B-alpha
        layer_range: [0, 32]
merge_method: slerp
base_model: openchat/openchat-3.5-1210 
parameters:
  t:
    - filter: self_attn
      value: [0, 0.5, 0.3, 0.7, 1]
    - filter: mlp
      value: [1, 0.5, 0.7, 0.3, 0]
    - value: 0.5
dtype: bfloat16
```