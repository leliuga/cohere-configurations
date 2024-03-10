---
inference: false
language:
- en
library_name: transformers
pipeline_tag: text-generation
tags:
- mixtral
datasets:
- LDJnr/Capybara
- unalignment/toxic-dpo-v0.1
- LDJnr/Verified-Camel
- HuggingFaceH4/no_robots
- Doctor-Shotgun/no-robots-sharegpt
- Doctor-Shotgun/capybara-sharegpt
---

# Norobara-ZLoss-8x7B

This is an experimental instruct-tuned [mistralai/Mixtral-8x7B-v0.1](https://huggingface.co/mistralai/Mixtral-8x7B-v0.1)-based model trained using [Charles Goddard](https://huggingface.co/chargoddard)'s ZLoss and Megablocks-based fork of transformers.

It primarily uses the Capybara and No Robots datasets (thus the name). The goal was to create an uncensored general instruction following model, as well as test various loss implementations while we figure out how the heck to train Mixtral properly.

[Exl2 Quants](https://huggingface.co/royallab/Norobara-ZLoss-8x7B-exl2)

Quants courtesy of TheBloke:
- [GPTQ](https://huggingface.co/TheBloke/Norobara-ZLoss-8x7B-GPTQ)
- [GGUF](https://huggingface.co/TheBloke/Norobara-ZLoss-8x7B-GGUF)
- [AWQ](https://huggingface.co/TheBloke/Norobara-ZLoss-8x7B-AWQ)

Additional Exl2 Quants courtesy of LoneStriker:
- [2.4bpw](https://huggingface.co/LoneStriker/Norobara-ZLoss-8x7B-2.4bpw-h6-exl2)
- [3.0bpw](https://huggingface.co/LoneStriker/Norobara-ZLoss-8x7B-3.0bpw-h6-exl2)
- [3.5bpw](https://huggingface.co/LoneStriker/Norobara-ZLoss-8x7B-3.5bpw-h6-exl2)
- [3.75bpw](https://huggingface.co/LoneStriker/Norobara-ZLoss-8x7B-3.75bpw-h6-exl2)
- [4.0bpw](https://huggingface.co/LoneStriker/Norobara-ZLoss-8x7B-4.0bpw-h6-exl2)
- [5.0bpw](https://huggingface.co/LoneStriker/Norobara-ZLoss-8x7B-5.0bpw-h6-exl2)
- [6.0bpw](https://huggingface.co/LoneStriker/Norobara-ZLoss-8x7B-6.0bpw-h6-exl2)

## Usage:
The intended prompt format is a modified multi-turn Alpaca instruction format:
```
### Instruction:
{system prompt}

### Input:
{user message}

### Response:
{model response}

### Input:
{user message}

### Response:
{model response}

(etc.)
```

## Bias, Risks, and Limitations
The model will show biases present in the base model. No ethical alignment was applied to prevent the generation of toxic or harmful outputs (in fact the opposite, with examples from toxic-DPO included), so generate at your own risk.
## Training Details
This model was trained as a QLora adapter for 3 epochs using a single H100 GPU for around 13 hours.