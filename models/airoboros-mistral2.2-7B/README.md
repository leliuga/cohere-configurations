---
base_model: mistralai/Mistral-7b-V0.1
tags:
- llama-2
- instruct
- finetune
- alpaca
- gpt4
- synthetic data
- distillation
datasets:
- jondurbin/airoboros-2.2.1
model-index:
- name: airoboros2.2-mistral-7b
  results: []
license: mit
language:
- en
---

Mistral trained with the airoboros dataset!

![image/png](https://cdn-uploads.huggingface.co/production/uploads/6317aade83d8d2fd903192d9/sbN_PCdxO_LV0xpFGA_St.png)

Actual dataset is airoboros 2.2, but it seems to have been replaced on hf with 2.2.1.

Prompt Format:
```
USER: <prompt>
ASSISTANT:
```


TruthfulQA:

```
hf-causal-experimental (pretrained=/home/teknium/dakota/lm-evaluation-harness/airoboros2.2-mistral/,dtype=float16), limit: None, provide_description: False, num_fewshot: 0, batch_size: 8
|    Task     |Version|Metric|Value |   |Stderr|
|-------------|------:|------|-----:|---|-----:|
|truthfulqa_mc|      1|mc1   |0.3562|±  |0.0168|
|             |       |mc2   |0.5217|±  |0.0156|
```


Wandb training charts: https://wandb.ai/teknium1/airoboros-mistral-7b/runs/airoboros-mistral-1?workspace=user-teknium1

More info to come