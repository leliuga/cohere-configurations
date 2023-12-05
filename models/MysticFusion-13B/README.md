---
language:
- en
pipeline_tag: text-generation
inference: false
license: llama2
---

YAML:
```
models:
  - model: KoboldAI/LLaMA2-13B-Tiefighter
    parameters:
      weight: 0.3
  - model: NeverSleep/Noromaid-13b-v0.1.1
    parameters:
      weight: 0.5
  - model: lmsys/vicuna-13b-v1.5
    parameters:
      weight: 0.2
merge_method: linear
dtype: float16
```

# Prompt Template:
### Alpaca
```
### Instruction:

### Response:

```