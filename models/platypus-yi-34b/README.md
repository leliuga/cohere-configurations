---
license: other
license_name: yi-license
license_link: LICENSE
datasets:
- garage-bAInd/Open-Platypus
---

# Instruction tune of Yi-34b with Open-Platypus (fp16)


## Overview

This is [chargoddard/Yi-34B-Llama](https://huggingface.co/chargoddard/Yi-34B-Llama), with instruction tuning performed with the [garage-bAInd/Open-Platypus](https://huggingface.co/datasets/garage-bAInd/Open-Platypus) dataset. That base model is [01-ai/Yi-34B](https://huggingface.co/01-ai/Yi-34B), but using llama2 model definitions and tokenizer to remove any remote code requirements.

**This is a (merged) QLoRA fine-tune (rank 64)**. 

The finetune was performed with 1x RTX 6000 Ada (~18 hours to this checkpoint). It is possible this is rather undertrained, as this checkpoint is at 1 epoch. I began to see some performance degradation after that; more hyperparameter tuning is probably warranted.

## How to Use

Use as you would any llama-2 model.

## Prompting:

Model was trained with legacy airoboros <2.0 system prompt. See [bhenrym14/airoboros-33b-gpt4-1.4.1-lxctx-PI-16384-fp16](https://huggingface.co/bhenrym14/airoboros-33b-gpt4-1.4.1-lxctx-PI-16384-fp16) model card for details.
