---
license: gpl-3.0
language:
- en
- zh
tags:
- qwen
---
# CausalLM / Qwen 8x7B MoE - This is not Mixtral / Mistral 7B

A Chat Model, Testing only, no performance guaranteeeee...

In short: CausalLM / Qwen 8x7B MoE in Mixtral Arch, from 8 real, explainable expert models in different domains. Trained, not a merge.

Only intended for conceptual validation, however the expert models do not seem to be working as expected. The model could output text and complete the conversation normally, but the performance of the expert model was not significant.

There are 8 completely different expert models based on Qwen-7B / CausalLM, six of which are specific domain models that have seen 50~100 billion tokens, including: a Toolformer/Agent expert model, a multilingual translation expert model, a mathematics expert model, a visual expert model, a coding and computer expert model, and an uncensored knowledge model — together forming the MoE model along with Qwen-Chat and Qwen-Base.

The initialization of the gate is based on the hidden state of the few-shot prompt input from each expert model and undergoes simple alignment training, on flan/ orca style.

For multimodal input, please use visual.bin/py, should be the same as Qwen-VL.

Prompt format: ChatML

A simple verification found that the expert model occasionally had routing errors, resulting in suboptimal results and required further fine-tuning.