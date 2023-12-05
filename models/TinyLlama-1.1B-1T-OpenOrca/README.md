---
license: apache-2.0
datasets:
- Open-Orca/OpenOrca
- bigcode/starcoderdata
- cerebras/SlimPajama-627B
language:
- en
---

[<img src="https://raw.githubusercontent.com/OpenAccess-AI-Collective/axolotl/main/image/axolotl-badge-web.png" alt="Built with Axolotl" width="200" height="32"/>](https://github.com/OpenAccess-AI-Collective/axolotl)

#### Base model:
PY007/TinyLlama-1.1B-intermediate-step-480k-1T 

#### Dataset:
Fine tuned on OpenOrca GPT4 subset for 1 epoch,Using CHATML format

#### Model License:
Apache 2.0, following the TinyLlama base model.

#### Quantisation:
 - GPTQ:https://huggingface.co/TheBloke/TinyLlama-1.1B-1T-OpenOrca-GPTQ
 - AWQ:https://huggingface.co/TheBloke/TinyLlama-1.1B-1T-OpenOrca-AWQ
 - GGUF:https://huggingface.co/TheBloke/TinyLlama-1.1B-1T-OpenOrca-GGUF

#### Hardware and training details:
Hardware: 1*RTX A5000, ~16 hours to complete 1 epoch. GPU from autodl.com, cost around $3 for this finetuning. 
https://wandb.ai/jeff200402/TinyLlama-Orca?workspace= for more details.