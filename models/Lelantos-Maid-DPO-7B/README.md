---
license: apache-2.0
tags:
- merge
- mergekit
- lazymergekit
- SanjiWatsuki/Lelantos-DPO-7B
- NeverSleep/Noromaid-7B-0.4-DPO
base_model:
- SanjiWatsuki/Lelantos-DPO-7B
- NeverSleep/Noromaid-7B-0.4-DPO
---

# Lelantos-Maid-DPO-7B

Lelantos-Maid-DPO-7B is a merge of the following models using [mergekit](https://github.com/cg123/mergekit):
* [SanjiWatsuki/Lelantos-DPO-7B](https://huggingface.co/SanjiWatsuki/Lelantos-DPO-7B)
* [NeverSleep/Noromaid-7B-0.4-DPO](https://huggingface.co/NeverSleep/Noromaid-7B-0.4-DPO)

## 🧩 Configuration

```yaml
slices:
  - sources:
      - model: SanjiWatsuki/Lelantos-DPO-7B
        layer_range: [0, 32] 
      - model: NeverSleep/Noromaid-7B-0.4-DPO
        layer_range: [0, 32]
merge_method: slerp
base_model: NeverSleep/Noromaid-7B-0.4-DPO
parameters:
  t:
    - value: 0.5 # fallback for rest of tensors
dtype: bfloat16
```