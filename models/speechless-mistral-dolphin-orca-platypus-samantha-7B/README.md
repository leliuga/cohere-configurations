---
language:
- en
library_name: transformers
pipeline_tag: text-generation
datasets:
- jondurbin/airoboros-2.2.1
- Open-Orca/OpenOrca
- garage-bAInd/Open-Platypus
- ehartford/samantha-data
tags:
- llama-2
- code
license: llama2
model-index:
- name: SpeechlessCoder
  results:
  - task:
      type: text-generation
    dataset:
      type: openai_humaneval
      name: HumanEval
    metrics:
    - name: pass@1
      type: pass@1
      value: 34.146
      verified: false
---

<p><h1> speechless-mistral-dolphin-orca-platypus-samantha-7b </h1></p>

* [AWQ model(s) for GPU inference.](https://huggingface.co/TheBloke/speechless-mistral-dolphin-orca-platypus-samantha-7B-AWQ)
* [GPTQ models for GPU inference, with multiple quantisation parameter options.](https://huggingface.co/TheBloke/speechless-mistral-dolphin-orca-platypus-samantha-7B-GPTQ)
* [2, 3, 4, 5, 6 and 8-bit GGUF models for CPU+GPU inference](https://huggingface.co/TheBloke/speechless-mistral-dolphin-orca-platypus-samantha-7B-GGUF)

This model is a merge of ehartford/dolphin-2.1-mistral-7b, Open-Orca/Mistral-7B-OpenOrca, bhenrym14/mistral-7b-platypus-fp16 and ehartford/samantha-1.2-mistral-7b.

I'm very sorry for giving such a long and peculiar name. Originally, it was just my lazy behavior during the process of making models to easily distinguish various model and dataset combinations. I didn't expect the [previous model](https://huggingface.co/uukuguy/speechless-llama2-hermes-orca-platypus-wizardlm-13b) ([Thebloke GPTQ Version](https://huggingface.co/TheBloke/Speechless-Llama2-Hermes-Orca-Platypus-WizardLM-13B-GPTQ)) to be so popular. This time, based on some guys's request, I am releasing a model based on Mistral, and I have also inherited the style of the super long name along with it. Welcome to try the model, please refrain from harsh criticism if you don't like it.

Code: https://github.com/uukuguy/speechless

## HumanEval

| Metric | Value |
| --- | --- |
| humaneval-python |  34.146|

[Big Code Models Leaderboard](https://huggingface.co/spaces/bigcode/bigcode-models-leaderboard)

CodeLlama-34B-Python: 53.29

CodeLlama-34B-Instruct: 50.79

CodeLlama-13B-Instruct: 50.6

CodeLlama-34B: 45.11

CodeLlama-13B-Python: 42.89

CodeLlama-13B: 35.07 

Mistral-7B-v0.1: 30.488

## LM-Evaluation-Harness

[Open LLM Leaderboard](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)

| Metric | Value |
| --- | --- |
| ARC | 64.33 |
| HellaSwag | 84.4|
| MMLU | 63.72 |
| TruthfulQA | 52.52|
| Winogrande | 78.37 |
| GSM8K | 21.38 |
| DROP | 8.66 |
| Average | 53.34 |

# Model Card for Mistral-7B-v0.1

The Mistral-7B-v0.1 Large Language Model (LLM) is a pretrained generative text model with 7 billion parameters. 
Mistral-7B-v0.1 outperforms Llama 2 13B on all benchmarks we tested.

For full details of this model please read our [paper](https://arxiv.org/abs/2310.06825) and [release blog post](https://mistral.ai/news/announcing-mistral-7b/).

## Model Architecture

Mistral-7B-v0.1 is a transformer model, with the following architecture choices:
- Grouped-Query Attention
- Sliding-Window Attention
- Byte-fallback BPE tokenizer

## Troubleshooting

- If you see the following error:
``
KeyError: 'mistral'
``
- Or:
``
NotImplementedError: Cannot copy out of meta tensor; no data!
``

Ensure you are utilizing a stable version of Transformers, 4.34.0 or newer.

## Notice

Mistral 7B is a pretrained base model and therefore does not have any moderation mechanisms.

## The Mistral AI Team
 
Albert Jiang, Alexandre Sablayrolles, Arthur Mensch, Chris Bamford, Devendra Singh Chaplot, Diego de las Casas, Florian Bressand, Gianna Lengyel, Guillaume Lample, Lélio Renard Lavaud, Lucile Saulnier, Marie-Anne Lachaux, Pierre Stock, Teven Le Scao, Thibaut Lavril, Thomas Wang, Timothée Lacroix, William El Sayed.`

# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_uukuguy__speechless-mistral-dolphin-orca-platypus-samantha-7b)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 53.34   |
| ARC (25-shot)         | 64.33          |
| HellaSwag (10-shot)   | 84.4    |
| MMLU (5-shot)         | 63.72         |
| TruthfulQA (0-shot)   | 52.52   |
| Winogrande (5-shot)   | 78.37   |
| GSM8K (5-shot)        | 21.38        |
| DROP (3-shot)         | 8.66         |
