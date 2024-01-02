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
      value:  75.61
      verified: false
---

<p><h1> speechless-codellama-34b-v2.0  </h1></p>

* [AWQ model(s) for GPU inference.](https://huggingface.co/TheBloke/speechless-codellama-34b-v2.0-AWQ)
* [GPTQ models for GPU inference, with multiple quantisation parameter options.](https://huggingface.co/TheBloke/speechless-codellama-34b-v2.0-GPTQ)
* [2, 3, 4, 5, 6 and 8-bit GGUF models for CPU+GPU inference](https://huggingface.co/TheBloke/speechless-codellama-34b-v2.0-GGUF)

Code: https://github.com/uukuguy/speechless

Use the following datasets to fine-tune codellama/CodeLlama-34B in order to improve the model's inference and planning  capabilities.

Total 153,013 samples.
- jondurbin/airoboros-2.2: Filter categories related to coding, reasoning and planning. 23,462 samples.
- Open-Orca/OpenOrca: Filter the 'cot' category in 1M GPT4 dataset. 74,440 samples.
- garage-bAInd/Open-Platypus: 100%, 24,926 samples.
- WizardLM/WizardLM_evol_instruct_V2_196k: Coding coversation part. 30,185 samples

## How to Prompt the Model
This model accepts the Alpaca instruction format.

For example:
```
You are an intelligent programming assistant.

### Instruction:
Implement a linked list in C++

### Response:
```




## HumanEval

| human-eval | pass@1 |
| --- | --- |
| humaneval-python | 75.61 |

[Big Code Models Leaderboard](https://huggingface.co/spaces/bigcode/bigcode-models-leaderboard)

| Models | pass@1 |
|------ | ------ |
| Phind-CodeLlama-34B-v2|  71.95| 
| WizardCoder-Python-34B-V1.0|  70.73| 
| Phind-CodeLlama-34B-Python-v1|  70.22| 
| Phind-CodeLlama-34B-v1|  65.85| 
| WizardCoder-Python-13B-V1.0|  62.19| 
| WizardCoder-15B-V1.0|  58.12| 
| CodeLlama-34B-Python|  53.29| 
| CodeLlama-34B-Instruct|  50.79| 
| CodeLlama-13B-Instruct|  50.6| 
| CodeLlama-34B|  45.11| 
| CodeLlama-13B-Python|  42.89| 
| CodeLlama-13B|  35.07| 

## NL2SQL

SQL-EVAL: 125/175 (71.43%)

Average rate of exact match: 67.43%

Average correct rate: 71.43%

- GPT4: 130/175 (74.29%)
- GPT3-Turbo-0613: 105/174 (60.00%)


## lm-evaluation-harness

[Open LLM Leaderboard](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
| Metric | Value |
| --- | --- |
| ARC | 54.35 |
| HellaSwag | 75.65 |
| MMLU | 54.67 |
| TruthfulQA | 45.21 |
| Average | 57.47 |


H800-80G x 2

transformers=4.33.0

flash-attn=2.1.0

bitsandbytes=0.41.1

peft=0.5.0

## Training Arguments
| | |
|------ | ------ |
| lr | 2e-4 |
| lr_scheduler_type | cosine |
| weight_decay | 0.0 |
| optim | paged_adamw_8bit |
| flash_attention | True |
| rerope | False |
| max_new_tokens | 8192 |
| num_train_epochs | 3 |
| bits | 4 |
| lora_r | 64 |
| lora_alpha | 16 |
| lora_dropout | 0.05 |
| double_quant | True |
| quant_type | nf4 |
| dataset_format | airoboros |
| mini_batch_size | 4 |
| grandient_accumulation_steps | 16 |
| bf16 | True |


| | |
|------ | ------ |
| epoch                    |                3.0 |
| etrain_loss               |             0.4261 |
| etrain_runtime            | 1 day, 14:42:57.87 |
| etrain_samples_per_second |              3.227 |
| etrain_steps_per_second   |              0.025 |
| eeval_loss               |     0.4537 |
| eeval_runtime            | 0:00:36.19 |
| eeval_samples_per_second |      5.525 |
| eeval_steps_per_second   |      2.763 |


# **Code Llama**

Code Llama is a collection of pretrained and fine-tuned generative text models ranging in scale from 7 billion to 34 billion parameters. This is the repository for the base 13B version in the Hugging Face Transformers format. This model is designed for general code synthesis and understanding. Links to other models can be found in the index at the bottom.

|     | Base Model                                                                    | Python                                                                                      | Instruct                                                                                        |
| --- | ----------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| 7B  | [codellama/CodeLlama-7b-hf](https://huggingface.co/codellama/CodeLlama-7b-hf) | [codellama/CodeLlama-7b-Python-hf](https://huggingface.co/codellama/CodeLlama-7b-Python-hf) | [codellama/CodeLlama-7b-Instruct-hf](https://huggingface.co/codellama/CodeLlama-7b-Instruct-hf) |
| 13B  | [codellama/CodeLlama-13b-hf](https://huggingface.co/codellama/CodeLlama-13b-hf) | [codellama/CodeLlama-13b-Python-hf](https://huggingface.co/codellama/CodeLlama-13b-Python-hf) | [codellama/CodeLlama-13b-Instruct-hf](https://huggingface.co/codellama/CodeLlama-13b-Instruct-hf) |
| 34B  | [codellama/CodeLlama-34b-hf](https://huggingface.co/codellama/CodeLlama-34b-hf) | [codellama/CodeLlama-34b-Python-hf](https://huggingface.co/codellama/CodeLlama-34b-Python-hf) | [codellama/CodeLlama-34b-Instruct-hf](https://huggingface.co/codellama/CodeLlama-34b-Instruct-hf) |


## Model Use

To use this model, please make sure to install transformers from `main` until the next version is released:

```bash
pip install git+https://github.com/huggingface/transformers.git@main accelerate
```

Model capabilities:

- [x] Code completion.
- [x] Infilling.
- [ ] Instructions / chat.
- [ ] Python specialist.


```python
from transformers import AutoTokenizer
import transformers
import torch

model = "codellama/CodeLlama-13b-hf"

tokenizer = AutoTokenizer.from_pretrained(model)
pipeline = transformers.pipeline(
    "text-generation",
    model=model,
    torch_dtype=torch.float16,
    device_map="auto",
)

sequences = pipeline(
    'import socket\n\ndef ping_exponential_backoff(host: str):',
    do_sample=True,
    top_k=10,
    temperature=0.1,
    top_p=0.95,
    num_return_sequences=1,
    eos_token_id=tokenizer.eos_token_id,
    max_length=200,
)
for seq in sequences:
    print(f"Result: {seq['generated_text']}")
```


## Model Details
*Note: Use of this model is governed by the Meta license. Meta developed and publicly released the Code Llama family of large language models (LLMs).

**Model Developers** Meta

**Variations** Code Llama comes in three model sizes, and three variants:

* Code Llama: base models designed for general code synthesis and understanding
* Code Llama - Python: designed specifically for Python
* Code Llama - Instruct: for instruction following and safer deployment

All variants are available in sizes of 7B, 13B and 34B parameters.

**This repository contains the base version of the 13B parameters model.**

**Input** Models input text only.

**Output** Models generate text only.

**Model Architecture** Code Llama is an auto-regressive language model that uses an optimized transformer architecture.

**Model Dates** Code Llama and its variants have been trained between January 2023 and July 2023.

**Status** This is a static model trained on an offline dataset. Future versions of Code Llama - Instruct will be released as we improve model safety with community feedback.

**License** A custom commercial license is available at: [https://ai.meta.com/resources/models-and-libraries/llama-downloads/](https://ai.meta.com/resources/models-and-libraries/llama-downloads/)

**Research Paper** More information can be found in the paper "[Code Llama: Open Foundation Models for Code](https://ai.meta.com/research/publications/code-llama-open-foundation-models-for-code/)" or its [arXiv page](https://arxiv.org/abs/2308.12950).

## Intended Use
**Intended Use Cases** Code Llama and its variants is intended for commercial and research use in English and relevant programming languages. The base model Code Llama can be adapted for a variety of code synthesis and understanding tasks, Code Llama - Python is designed specifically to handle the Python programming language, and Code Llama - Instruct is intended to be safer to use for code assistant and generation applications.

**Out-of-Scope Uses** Use in any manner that violates applicable laws or regulations (including trade compliance laws). Use in languages other than English. Use in any other way that is prohibited by the Acceptable Use Policy and Licensing Agreement for Code Llama and its variants.

## Hardware and Software
**Training Factors** We used custom training libraries. The training and fine-tuning of the released models have been performed Meta’s Research Super Cluster.

**Carbon Footprint** In aggregate, training all 9 Code Llama models required 400K GPU hours of computation on hardware of type A100-80GB (TDP of 350-400W). Estimated total emissions were 65.3 tCO2eq, 100% of which were offset by Meta’s sustainability program.

## Training Data

All experiments reported here and the released models have been trained and fine-tuned using the same data as Llama 2 with different weights (see Section 2 and Table 1 in the [research paper](https://ai.meta.com/research/publications/code-llama-open-foundation-models-for-code/) for details).

## Evaluation Results

See evaluations for the main models and detailed ablations in Section 3 and safety evaluations in Section 4 of the research paper.


## Ethical Considerations and Limitations

Code Llama and its variants are a new technology that carries risks with use. Testing conducted to date has been in English, and has not covered, nor could it cover all scenarios. For these reasons, as with all LLMs, Code Llama’s potential outputs cannot be predicted in advance, and the model may in some instances produce inaccurate or objectionable responses to user prompts. Therefore, before deploying any applications of Code Llama, developers should perform safety testing and tuning tailored to their specific applications of the model.

Please see the Responsible Use Guide available available at [https://ai.meta.com/llama/responsible-user-guide](https://ai.meta.com/llama/responsible-user-guide).

# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_uukuguy__speechless-codellama-34b-v2.0)

| Metric                | Value                     |
|-----------------------|---------------------------|
| Avg.                  | 50.96   |
| ARC (25-shot)         | 54.35          |
| HellaSwag (10-shot)   | 75.65    |
| MMLU (5-shot)         | 54.67         |
| TruthfulQA (0-shot)   | 45.21   |
| Winogrande (5-shot)   | 73.56   |
| GSM8K (5-shot)        | 11.6        |
| DROP (3-shot)         | 41.71         |
