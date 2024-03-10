---
language:
- en
library_name: transformers
pipeline_tag: text-generation
datasets:
- yahma/alpaca-cleaned
license: apache-2.0
---

<p><h1> speechless-mistral-moloras-7b </h1></p>

* [AWQ model(s) for GPU inference.](https://huggingface.co/TheBloke/speechless-mistral-moloras-7B-AWQ)
* [GPTQ models for GPU inference, with multiple quantisation parameter options.](https://huggingface.co/TheBloke/speechless-mistral-moloras-7B-GPTQ)
* [2, 3, 4, 5, 6 and 8-bit GGUF models for CPU+GPU inference](https://huggingface.co/TheBloke/speechless-mistral-moloras-7B-GGUF)

[4-bit GGUF models for CPU+GPU inference](https://huggingface.co/uukuguy/speechless-mistral-moloras-7b/tree/main/GGUF)

This model is the static version of moloras (Mixture-of-multi-LoRAs) based on the following 6 Mistral-based LoRa modules.

- Intel/neural-chat-7b-v3-1 
- migtissera/SynthIA-7B-v1.3
- jondurbin/airoboros-m-7b-3.1.2
- bhenrym14/mistral-7b-platypus-fp16 
- teknium/CollectiveCognition-v1.1-Mistral-7B 
- uukuguy/speechless-mistral-dolphin-orca-platypus-samantha-7b 

Totally 6 LoRA modules from [speechless-mistral-7b-dare-0.85](https://huggingface.co/speechlessai/speechless-mistral-7b-dare-0.85)

The router of mixture-of-multi-loras enables an automatic assembling of LoRA modules, using a gradientfree approach to obtain the coefficients of LoRA modules and requiring only a handful of inference steps for unseen tasks.




Code: https://github.com/uukuguy/multi_loras?tab=readme-ov-file#mixture-of-multi-loras

## LM-Evaluation-Harness

[Open LLM Leaderboard](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)

| Metric | Value |
| --- | --- |
| ARC | 59.98  |
| HellaSwag | 83.29 |
| MMLU | 64.12 |
| TruthfulQA | 42.15 |
| Winogrande | 78.37 |
| GSM8K | 37.68 |
| Average | 60.93 |

