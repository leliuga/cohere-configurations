---
base_model: Mihaiii/Metis-0.3
inference: false
license: apache-2.0
license_name: apache-2.0
metrics:
- accuracy
tags:
- merge
---

This is a merge between Metis-0.3 and Metis-0.1 having Metis-0.1 as base.
It was done using [mergekit](https://github.com/cg123/mergekit).

It works well with long system prompts.

It isn't generic in a sense that it shouldn't be used for story telling, for example, but only for reasoning and text comprehension.

This model is trained on a private dataset.

# Prompt Format:

```
<|system|>
{system_message} </s>
<|user|>
{prompt} </s>
<|assistant|>
```

Merge config:
```yaml
slices:
  - sources:
      - model: Mihaiii/Metis-0.3
        layer_range: [0, 32]
      - model: Mihaiii/Metis-0.1
        layer_range: [0, 32]
merge_method: slerp
base_model: Mihaiii/Metis-0.1
parameters:
  t:
    - filter: self_attn
      value: [0, 0.5, 0.3, 0.7, 1]
    - filter: mlp
      value: [1, 0.5, 0.7, 0.3, 0]
    - value: 0.5 # fallback for rest of tensors
dtype: bfloat16
```