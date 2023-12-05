---
license: llama2
---
```
merge_method: task_arithmetic
base_model: TheBloke/Llama-2-13B-fp16
models:
  - model: TheBloke/Llama-2-13B-fp16
  - model: KoboldAI/LLaMA2-13B-Tiefighter
    parameters:
      weight: 1.0
  - model: chaoyi-wu/MedLLaMA_13B
    parameters:
      weight: 0.01
  - model: Doctor-Shotgun/llama-2-13b-chat-limarp-v2-merged
    parameters:
      weight: 0.02
dtype: float16
```
This model was made possible thanks to the Compute provided by the KoboldAI community.