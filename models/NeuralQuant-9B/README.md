---
license: apache-2.0
tags:
- merge
---

# NeuralQuant-9B

This model is a merge of the following models made with [mergekit](https://github.com/cg123/mergekit):
 * [quantumaikr/quantum-v0.01](https://huggingface.co/quantumaikr/quantum-v0.01)
 * [mlabonne/NeuralHermes-2.5-Mistral-7B](https://huggingface.co/mlabonne/NeuralHermes-2.5-Mistral-7B)

## 🧩 Configuration

```yaml
slices:
  - sources:
    - model: quantumaikr/quantum-v0.01
      layer_range: [0, 32]
  - sources:
    - model: mlabonne/NeuralHermes-2.5-Mistral-7B
      layer_range: [24, 32]
merge_method: passthrough
dtype: bfloat16
```