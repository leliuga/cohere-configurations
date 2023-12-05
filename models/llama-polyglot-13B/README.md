---
license: llama2
tags:
- llama2
---

Experimental multi-lingual model using a new merge technique.

Mergekit configuration (experimental branch):
```yaml
models:
  - model: clibrain/Llama-2-13b-ft-instruct-es
  - model: LeoLM/leo-hessianai-13b
  - model: daekeun-ml/Llama-2-ko-DPO-13B
  - model: pleisto/yuren-13b-chatml
  - model: bofenghuang/vigogne-2-13b-instruct
  - model: OpenBuddy/openbuddy-llama2-13b-v8.1-fp16
merge_method: dare_ties
base_model: TheBloke/Llama-2-13B-fp16
dtype: float16
parameters:
  density: 0.3
  weight: 1.0
  normalize: true
  int8_mask: true
tokenizer_source: base
```