---
language:
- en
- zh
license: llama2
library_name: transformers
tags:
- llama
- merge
- medical
datasets:
- GBaker/MedQA-USMLE-4-options
- cognitivecomputations/samantha-data
- shibing624/medical
base_model:
- Severus27/BeingWell_llama2_7b
- ParthasarathyShanmugam/llama-2-7b-samantha
pipeline_tag: text-generation
model-index:
- name: Dr_Samantha-7b
  results:
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: AI2 Reasoning Challenge (25-Shot)
      type: ai2_arc
      config: ARC-Challenge
      split: test
      args:
        num_few_shot: 25
    metrics:
    - type: acc_norm
      value: 53.84
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=sethuiyer/Dr_Samantha-7b
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: HellaSwag (10-Shot)
      type: hellaswag
      split: validation
      args:
        num_few_shot: 10
    metrics:
    - type: acc_norm
      value: 77.95
      name: normalized accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=sethuiyer/Dr_Samantha-7b
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: MMLU (5-Shot)
      type: cais/mmlu
      config: all
      split: test
      args:
        num_few_shot: 5
    metrics:
    - type: acc
      value: 47.94
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=sethuiyer/Dr_Samantha-7b
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: TruthfulQA (0-shot)
      type: truthful_qa
      config: multiple_choice
      split: validation
      args:
        num_few_shot: 0
    metrics:
    - type: mc2
      value: 45.58
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=sethuiyer/Dr_Samantha-7b
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: Winogrande (5-shot)
      type: winogrande
      config: winogrande_xl
      split: validation
      args:
        num_few_shot: 5
    metrics:
    - type: acc
      value: 73.56
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=sethuiyer/Dr_Samantha-7b
      name: Open LLM Leaderboard
  - task:
      type: text-generation
      name: Text Generation
    dataset:
      name: GSM8k (5-shot)
      type: gsm8k
      config: main
      split: test
      args:
        num_few_shot: 5
    metrics:
    - type: acc
      value: 18.8
      name: accuracy
    source:
      url: https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard?query=sethuiyer/Dr_Samantha-7b
      name: Open LLM Leaderboard
---

# Dr. Samantha

<p align="center">
  <img src="https://huggingface.co/sethuiyer/Dr_Samantha-7b/resolve/main/dr_samantha_anime_style_reduced_quality.webp" height="256px" alt="SynthIQ">
</p>

## Overview

Dr. Samantha is a language model made by merging `Severus27/BeingWell_llama2_7b` and `ParthasarathyShanmugam/llama-2-7b-samantha` using [mergekit](https://github.com/cg123/mergekit).

Has capabilities of a medical knowledge-focused model (trained on USMLE databases and doctor-patient interactions) with the philosophical, psychological, and relational understanding of the Samantha-7b model. 

As both a medical consultant and personal counselor, Dr.Samantha could effectively support both physical and mental wellbeing - important for whole-person care.


# Yaml Config

```yaml

slices:
  - sources:
      - model: Severus27/BeingWell_llama2_7b
        layer_range: [0, 32]
      - model: ParthasarathyShanmugam/llama-2-7b-samantha
        layer_range: [0, 32]

merge_method: slerp
base_model: TinyPixel/Llama-2-7B-bf16-sharded

parameters:
  t:
    - filter: self_attn
      value: [0, 0.5, 0.3, 0.7, 1]
    - filter: mlp
      value: [1, 0.5, 0.7, 0.3, 0]
    - value: 0.5 # fallback for rest of tensors
tokenizer_source: union

dtype: bfloat16

```

## Prompt Template

```text
Below is an instruction that describes a task. Write a response that appropriately completes the request.

### Instruction:
What is your name?

### Response:
My name is Samantha.
```

## ⚡ Quantized models

* **GGUF**:https://huggingface.co/TheBloke/Dr_Samantha-7B-GGUF
* **GPTQ**: https://huggingface.co/TheBloke/Dr_Samantha-7B-GPTQ
* **AWQ**: https://huggingface.co/TheBloke/Dr_Samantha-7B-AWQ

Thanks to [TheBloke](https://huggingface.co/TheBloke) for making this available! 

Dr.Samantha is now available on Ollama. You can use it by running the command ```ollama run stuehieyr/dr_samantha``` in your 
terminal. If you have limited computing resources, check out this [video](https://www.youtube.com/watch?v=Qa1h7ygwQq8) to learn how to run it on 
a Google Colab backend.

## OpenLLM Leaderboard Performance
| T | Model                            | Average | ARC   | Hellaswag | MMLU  | TruthfulQA | Winogrande | GSM8K |
|---|----------------------------------|---------|-------|-----------|-------|------------|------------|-------|
| 1 | sethuiyer/Dr_Samantha-7b         | 52.95   | 53.84 | 77.95     | 47.94 | 45.58      | 73.56      | 18.8  |
| 2 | togethercomputer/LLaMA-2-7B-32K-Instruct | 50.02   | 51.11 | 78.51     | 46.11 | 44.86      | 73.88      | 5.69  |
| 3 | togethercomputer/LLaMA-2-7B-32K  | 47.07   | 47.53 | 76.14     | 43.33 | 39.23      | 71.9       | 4.32  |


## Subject-wise Accuracy

| Subject               | Accuracy (%) |
|-----------------------|--------------|
| Clinical Knowledge    | 52.83        |
| Medical Genetics      | 49.00        |
| Human Aging           | 58.29        |
| Human Sexuality       | 55.73        |
| College Medicine      | 38.73        |
| Anatomy               | 41.48        |
| College Biology       | 52.08        |
| College Medicine      | 38.73        |
| High School Biology   | 53.23        |
| Professional Medicine | 38.73        |
| Nutrition             | 50.33        |
| Professional Psychology | 46.57      |
| Virology              | 41.57        |
| High School Psychology | 66.60       |
| Average                | 48.85%      |


## Evaluation by GPT-4 across 25 random prompts from ChatDoctor-200k Dataset

### Overall Rating: 83.5/100

#### Pros:

- Demonstrates extensive medical knowledge through accurate identification of potential causes for various symptoms.
- Responses consistently emphasize the importance of seeking professional diagnoses and treatments.
- Advice to consult specialists for certain concerns is well-reasoned.
- Practical interim measures provided for symptom management in several cases.
- Consistent display of empathy, support, and reassurance for patients' well-being.
- Clear and understandable explanations of conditions and treatment options.
- Prompt responses addressing all aspects of medical inquiries.

#### Cons:

- Could occasionally place stronger emphasis on urgency when symptoms indicate potential emergencies.
- Discussion of differential diagnoses could explore a broader range of less common causes.
- Details around less common symptoms and their implications need more depth at times.
- Opportunities exist to gather clarifying details on symptom histories through follow-up questions.
- Consider exploring full medical histories to improve diagnostic context where relevant.
- Caution levels and risk factors associated with certain conditions could be underscored more.



# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_sethuiyer__Dr_Samantha-7b)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |52.95|
|AI2 Reasoning Challenge (25-Shot)|53.84|
|HellaSwag (10-Shot)              |77.95|
|MMLU (5-Shot)                    |47.94|
|TruthfulQA (0-shot)              |45.58|
|Winogrande (5-shot)              |73.56|
|GSM8k (5-shot)                   |18.80|

