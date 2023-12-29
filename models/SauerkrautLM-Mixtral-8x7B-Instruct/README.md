---
license: apache-2.0
language:
- en
- de
- fr
- it
- es
library_name: transformers
pipeline_tag: text-generation
tags:
- mistral
- finetune
- dpo
- Instruct
- augmentation
- german
- mixtral
datasets:
- argilla/distilabel-math-preference-dpo
---

![SauerkrautLM](https://vago-solutions.de/wp-content/uploads/2023/12/Sauerkraut_Instruct_MoE_Instruct.png "SauerkrautLM-Mixtral-8x7B")
## VAGO solutions SauerkrautLM-Mixtral-8x7B-Instruct
Introducing **SauerkrautLM-Mixtral-8x7B-Instruct** – our Sauerkraut version of the powerful Mixtral-8x7B-Instruct! 
Aligned with **DPO**

# Table of Contents
1. [Overview of all SauerkrautLM-Mixtral models](#all-sauerkrautlm-mixtral-models)
2. [Model Details](#model-details)
   - [Prompt template](#prompt-template)
   - [Training Dataset](#training-dataset)
   - [Data Contamination Test](#data-contamination-test-results)
3. [Evaluation](#evaluation)
5. [Disclaimer](#disclaimer)
6. [Contact](#contact)
7. [Collaborations](#collaborations)
8. [Acknowledgement](#acknowledgement)


## All SauerkrautLM-Mixtral Models

| Model | HF    | GPTQ  | GGUF  | AWQ  |
|-------|-------|-------|-------|-------|
| SauerkrautLM-Mixtral-8x7B-Instruct  | [Link](https://huggingface.co/VAGOsolutions/SauerkrautLM-Mixtral-8x7B-Instruct) | coming soon | coming soon | coming soon |
| SauerkrautLM-Mixtral-8x7B  | [Link](https://huggingface.co/VAGOsolutions/SauerkrautLM-Mixtral-8x7B) | coming soon | coming soon | coming soon |

## Model Details
**SauerkrautLM-Mixtral-8x7B-Instruct**
- **Model Type:** SauerkrautLM-Mixtral-8x7B-Instruct-v0.1 is a Mixture of Experts (MoE) Model based on [mistralai/Mixtral-8x7B-Instruct-v0.1](https://huggingface.co/mistralai/Mixtral-8x7B-Instruct-v0.1) 
- **Language(s):** English, German, French, Italian, Spanish
- **License:** APACHE 2.0
- **Contact:** [Website](https://vago-solutions.de/#Kontakt) [David Golchinfar](mailto:golchinfar@vago-solutions.de)

### Training Dataset:

SauerkrautLM-Mixtral-8x7B-Instruct was trained with mix of German data augmentation and translated data. 
Aligned through **DPO** with our **new German SauerkrautLM-DPO dataset** based on parts of the SFT SauerkrautLM dataset 
as chosen answers and [Sauerkraut-7b-HerO](https://huggingface.co/VAGOsolutions/SauerkrautLM-7b-HerO) as rejected answers. Added with additional **translated Parts of the [HuggingFaceH4/ultrafeedback_binarized](https://huggingface.co/datasets/HuggingFaceH4/ultrafeedback_binarized)** (Our dataset do not contain any TruthfulQA prompts - check Data Contamination Test Results) and **[argilla/distilabel-math-preference-dpo](https://huggingface.co/datasets/argilla/distilabel-math-preference-dpo).**  
We found, that only a simple translation of training data can lead to unnatural German phrasings. 
Data augmentation techniques were used to grant grammatical, syntactical correctness and a more natural German wording in our training data. 

### Data Contamination Test Results

Some models on the HuggingFace leaderboard had problems with wrong data getting mixed in.
We checked our SauerkrautLM-DPO dataset with a special test [1] on a smaller model for this problem. 
The HuggingFace team used the same methods [2, 3].

Our results, with `result < 0.1, %:` being well below 0.9, indicate that our dataset is free from contamination.

*The data contamination test results of HellaSwag and Winograde will be added once [1] supports them.*

| Dataset                        | ARC   | MMLU | TruthfulQA | GSM8K |
|------------------------------|-------|-------|-------|-------|
| **SauerkrautLM-DPO**| result < 0.1, %: 0.0 |result < 0.1, %: 0.09 | result < 0.1, %: 0.13 | result < 0.1, %: 0.16 |

[1] https://github.com/swj0419/detect-pretrain-code-contamination

[2] https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard/discussions/474#657f2245365456e362412a06

[3] https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard/discussions/265#657b6debf81f6b44b8966230

### Prompt Template:
```
[INST] Instruction [/INST] Model answer [INST] Follow-up instruction [/INST]
```
## Evaluation
![Harness](https://vago-solutions.de/wp-content/uploads/2023/12/MOE_Instruct.png "SauerkrautLM-Mixtral-8x7B-Instruct Harness")
*evaluated with lm-evaluation-harness v0.3.0 - mmlu coming soon

*All benchmarks were performed with a sliding window of 4096. New Benchmarks with Sliding Window null coming soon

## Disclaimer
We must inform users that despite our best efforts in data cleansing, the possibility of uncensored content slipping through cannot be entirely ruled out.
However, we cannot guarantee consistently appropriate behavior. Therefore, if you encounter any issues or come across inappropriate content, we kindly request that you inform us through the contact information provided.
Additionally, it is essential to understand that the licensing of these models does not constitute legal advice. We are not held responsible for the actions of third parties who utilize our models. These models may be employed for commercial purposes, and the Apache 2.0 remains applicable and is included with the model files.
 
## Contact
If you are interested in customized LLMs for business applications, please get in contact with us via our website or contact us at [Dr. Daryoush Vaziri](mailto:vaziri@vago-solutions.de). We are also grateful for your feedback and suggestions.
 
## Collaborations
We are also keenly seeking support and investment for our startup, VAGO solutions, where we continuously advance the development of robust language models designed to address a diverse range of purposes and requirements. If the prospect of collaboratively navigating future challenges excites you, we warmly invite you to reach out to us.

## Acknowledgement
Many thanks to [argilla](https://huggingface.co/datasets/argilla) and [Huggingface](https://huggingface.co) for providing such valuable datasets to the Open-Source community. And of course a big thanks to MistralAI for providing the open source community with their latest technology!