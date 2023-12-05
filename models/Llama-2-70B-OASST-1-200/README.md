---
license: apache-2.0
language:
- en
- de
- es
- fr
tags:
- sft
inference: false
datasets:
- OpenAssistant/oasst1
---

# Open-Assistant Llama2 70B SFT OASST

This model is a fine-tuning of [Llama2 70B](https://huggingface.co/meta-llama/Llama-2-70b-hf) LLM. 
It was trained on a mixture of OASST top-1 threads.
## Model Details

- **Finetuned from:** [Llama2 70B](https://huggingface.co/meta-llama/Llama-2-70b-hf)
- **Model type:** Causal decoder-only transformer language model
- **Language:** English, German, Spanish, French (and limited capabilities in Italian, Portuguese, Polish, Dutch, Romanian, Czech, Swedish);
- **License:** Apache 2.0
- **Contact:** [Open-Assistant Discord](https://ykilcher.com/open-assistant-discord)

## Prompting

Two special tokens are used to mark the beginning of user and assistant turns:
`<|prompter|>` and `<|assistant|>`. Each turn ends with a `</s>` token.

Input prompt example:
```
<|prompter|>What is a meme, and what's the history behind this word?</s><|assistant|>
```
The input ends with the `<|assistant|>` token to signal that the model should 
start generating the assistant reply.

