---
language:
- en
library_name: transformers
pipeline_tag: text-generation
datasets:
- jondurbin/airoboros-2.2
- Open-Orca/OpenOrca
- garage-bAInd/Open-Platypus
- WizardLM/WizardLM_evol_instruct_V2_196k
- TokenBender/python_eval_instruct_51k
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
      value: 50.0
      verified: false
---

<p><h1> speechless-code-mistral-7b-v1.0  </h1></p>

* [AWQ model(s) for GPU inference.](https://huggingface.co/TheBloke/speechless-code-mistral-7B-v1.0-AWQ)
* [GPTQ models for GPU inference, with multiple quantisation parameter options.](https://huggingface.co/TheBloke/speechless-code-mistral-7B-v1.0-GPTQ)
* [2, 3, 4, 5, 6 and 8-bit GGUF models for CPU+GPU inference](https://huggingface.co/TheBloke/speechless-code-mistral-7B-v1.0-GGUF)

Code: https://github.com/uukuguy/speechless

Use the following dataset to fine-tune mistralai/Mistral-7B-v0.1 in order to improve the model's reasoning and planning abilities.

Total 201,981 samples.
- jondurbin/airoboros-2.2: Filter categories related to coding, reasoning and planning. 23,462 samples.
- Open-Orca/OpenOrca: Filter the 'cot' category in 1M GPT4 dataset. 74,440 samples.
- garage-bAInd/Open-Platypus: 100%, 24,926 samples.
- WizardLM/WizardLM_evol_instruct_V2_196k: Coding coversation part. 30,185 samples
- TokenBender/python_eval_instruct_51k: “python” in output .40,309 samples
- Spider: 8,659 samples



## HumanEval

| Metric | Value |
| --- | --- |
| humaneval-python | 50.0|

[Big Code Models Leaderboard](https://huggingface.co/spaces/bigcode/bigcode-models-leaderboard)

CodeLlama-34B-Python: 53.29

CodeLlama-34B-Instruct: 50.79

CodeLlama-13B-Instruct: 50.6

CodeLlama-34B: 45.11

CodeLlama-13B-Python: 42.89

CodeLlama-13B: 35.07

## lm-evaluation-harness

[Open LLM Leaderboard](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
| Metric | Value |
| --- | --- |
| ARC |59.64 |
| HellaSwag |82.25 |
| MMLU | 61.33 |
| TruthfulQA | 48.45 |
| Average | 62.92 |

## Parameters

| | |
|------ | ------ |
| lr | 2e-4 |
| lr_scheduler_type | cosine |
| weight_decay | 0.0 |
| optim | paged_adamw_8bit |
| flash_attention | True |
| rerope | False |
| max_new_tokens | 4096 |
| num_train_epochs | 2 |
| bits | 4 |
| lora_r | 64 |
| lora_alpha | 16 |
| lora_dropout | 0.05 |
| double_quant | True |
| quant_type | nf4 |
| dataset_format | airoboros |
| mini_batch_size | 2 |
| grandient_accumulation_steps | 32 |
| bf16 | True |

A40-48G x 2

| | |
|------ | ------ |
| epoch                    |                2.0 |
| etrain_loss               |             0.5 |
| etrain_runtime            | 1 day, 10:25:26.77 |
| etrain_samples_per_second |              3.194 |
| etrain_steps_per_second   |              0.025 |
| eeval_loss               |     0.5146 |
| eeval_runtime            | 0:00:25.04 |
| eeval_samples_per_second |      7.985 |
| eeval_steps_per_second   |       |

# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_uukuguy__speechless-code-mistral-7b-v1.0)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 53.47   |
| ARC (25-shot)         | 60.58          |
| HellaSwag (10-shot)   | 83.75    |
| MMLU (5-shot)         | 62.98         |
| TruthfulQA (0-shot)   | 47.9   |
| Winogrande (5-shot)   | 78.69   |
| GSM8K (5-shot)        | 19.18        |
| DROP (3-shot)         | 21.19         |
