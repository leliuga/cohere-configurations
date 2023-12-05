---
license: cc-by-nc-4.0
---
```
merge_method: task_arithmetic
base_model: TheBloke/Llama-2-13B-fp16
models:
  - model: TheBloke/Llama-2-13B-fp16
  - model: NeverSleep/Echidna-13b-v0.3
    parameters:
      weight: 1.0
  - model: KoboldAI/LLaMA2-13B-Tiefighter
    parameters:
      weight: 0.25
dtype: float16
```