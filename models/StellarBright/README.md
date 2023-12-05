---
language:
- en
pipeline_tag: text-generation
tags:
- stellar-bright
- llama
- llama-2
- llama-2-chat
- 70b
model_type: llama
license: llama2
---

Stellar Bright is a general capability upgrade to Llama 2, using open source data to improve overall knowledge, extended communication, and technical skill.

This model is primarily recommended as a superior-to-Llama-2 baseline for additional finetuning, **not** for direct deployment to production as a chat model. The user accepts full responsibility for all outputs.



## Evaluation

| Model                 | Avg    | ARC   | HS    | MMLU   | TQA   |
|-----------------------|--------|-------|-------|--------|-------|
| **Stellar Bright**    | 74.10  | 72.95 | 87.82 | 71.17  | 64.46 |
| Llama 2               | 67.35  | 67.32 | 87.33 | 69.83  | 44.92 |
| Llama 2 Chat          | 66.80  | 64.59 | 85.88 | 63.91  | 52.80 |