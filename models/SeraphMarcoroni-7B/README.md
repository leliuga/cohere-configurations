---
license: cc-by-nc-4.0
---

# SeraphMarcoroni-7B

I used [mergekit](https://github.com/cg123/mergekit) to merge models.

- i) [Marcoroni](https://huggingface.co/AIDC-ai-business/Marcoroni-7B-v3)
- ii) [Seraph](https://huggingface.co/Weyaxi/Seraph-7B)

# Yaml Config

```yaml
slices:
  - sources:
      - model: LOCAL_PATH/Marcoroni-7B-v3
        layer_range: [0, 32]
      - model: LOCAL_PATH/models/Seraph-7B
        layer_range: [0, 32]
merge_method: slerp
base_model: LOCAL_PATH/models/Marcoroni-7B-v3
parameters:
  t:
    - filter: self_attn
      value: [0, 0.5, 0.3, 0.7, 1]
    - filter: mlp
      value: [1, 0.5, 0.7, 0.3, 0]
    - value: 0.5 # fallback for rest of tensors
dtype: float16
```