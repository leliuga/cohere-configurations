---
license: apache-2.0
tags:
- merge
---
Slerp merge of mindy-labs/mindy-7b-v2 with jondurbin/bagel-dpo-7b-v0.1. This model was then slerp merged with rishiraj/CatPPT.

Heard some talk of jondurbin/bagel-dpo-7b-v0.1 in the community and it sounds intresting. Merged it with two high preforming models to get cookinai/Valkyrie-V1

Slerp 1:

```.yaml:
slices:
  - sources:
      - model: jondurbin/bagel-dpo-7b-v0.1
        layer_range: [0, 32]
      - model: mindy-labs/mindy-7b-v2
        layer_range: [0, 32]
merge_method: slerp
base_model: mindy-labs/mindy-7b-v2
parameters:
  t:
    - filter: self_attn
      value: [0, 0.5, 0.3, 0.7, 1]
    - filter: mlp
      value: [1, 0.5, 0.7, 0.3, 0]
    - value: 0.5 # fallback for rest of tensors
dtype: bfloat16
```

Slerp 2:

```.yaml:
slices:
  - sources:
      - model: previous/model/path
        layer_range: [0, 32]
      - model: rishiraj/CatPPT
        layer_range: [0, 32]
merge_method: slerp
base_model: previous/model/path
parameters:
  t:
    - filter: self_attn
      value: [0, 0.5, 0.3, 0.7, 1]
    - filter: mlp
      value: [1, 0.5, 0.7, 0.3, 0]
    - value: 0.5 # fallback for rest of tensors
dtype: bfloat16
```