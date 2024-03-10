---
language:
- en
- de
- es
- fr
- it
license: apache-2.0
library_name: transformers
tags:
- dpo
- rlaif
- preference
- ultrafeedback
- moe
datasets:
- argilla/ultrafeedback-binarized-preferences-cleaned
base_model: mistralai/Mixtral-8x7B-Instruct-v0.1
pipeline_tag: text-generation
model-index:
- name: notux-8x7b-v1
  results: []
---

<div align="center">
  <img src="https://cdn-uploads.huggingface.co/production/uploads/60f0608166e5701b80ed3f02/dj-spsk9eXMMXVGxK6jRz.png" alt="A banner representing Notus, the wind god of the south, in a mythical and artistic style. The banner features a strong, swirling breeze, embodying the warm, wet character of the southern wind. Gracefully flowing across the scene are several paper planes, caught in the gentle yet powerful gusts of Notus. The background is a blend of warm colors, symbolizing the heat of the south, with hints of blue and green to represent the moisture carried by this wind. The overall atmosphere is one of dynamic movement and warmth."/>
</div>


# Model Card for Notux 8x7B-v1

This model is a preference-tuned version of [mistralai/Mixtral-8x7B-Instruct-v0.1](https://huggingface.co/mistralai/Mixtral-8x7B-Instruct-v0.1) on the [argilla/ultrafeedback-binarized-preferences-cleaned](https://huggingface.co/datasets/argilla/ultrafeedback-binarized-preferences-cleaned) dataset using DPO (Direct Preference Optimization).

As of Dec 26th 2023, it outperforms `Mixtral-8x7B-Instruct-v0.1` and is the top ranked MoE (Mixture of Experts) model on the [Hugging Face Open LLM Leaderboard](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard).

This is part of the Notus family of models and experiments, where the Argilla team investigates data-first and preference tuning methods like dDPO (distilled DPO). This model is the result of our first experiment at tuning a MoE model that has already been fine-tuned with DPO (i.e., Mixtral-8x7B-Instruct-v0.1). 

## Model Details

### Model Description

- **Developed by:** Argilla (based on MistralAI previous efforts)
- **Shared by:** Argilla
- **Model type:** Pretrained generative Sparse Mixture of Experts
- **Language(s) (NLP):** English, Spanish, Italian, German, and French
- **License:** MIT
- **Finetuned from model:** [mistralai/Mixtral-8x7B-Instruct-v0.1](https://huggingface.co/mistralai/Mixtral-8x7B-Instruct-v0.1)

### Model Sources

- **Repository:** https://github.com/argilla-io/notus
- **Paper:** N/A

## Training Details

### Training Hardware

We used a VM with 8 x H100 80GB hosted in runpod.io for 1 epoch (~10hr).

### Training Data

We used a new iteration of the Argilla UltraFeedback preferences dataset named [argilla/ultrafeedback-binarized-preferences-cleaned](https://huggingface.co/datasets/argilla/ultrafeedback-binarized-preferences-cleaned).

## Training procedure

### Training hyperparameters

The following hyperparameters were used during training:
- learning_rate: 5e-07
- train_batch_size: 8
- eval_batch_size: 4
- seed: 42
- distributed_type: multi-GPU
- num_devices: 8
- total_train_batch_size: 64
- total_eval_batch_size: 32
- optimizer: Adam with betas=(0.9,0.999) and epsilon=1e-08
- lr_scheduler_type: linear
- lr_scheduler_warmup_ratio: 0.1
- num_epochs: 1

### Training results

| Training Loss | Epoch | Step | Validation Loss | Rewards/chosen | Rewards/rejected | Rewards/accuracies | Rewards/margins | Logps/rejected | Logps/chosen | Logits/rejected | Logits/chosen |
|:-------------:|:-----:|:----:|:---------------:|:--------------:|:----------------:|:------------------:|:---------------:|:--------------:|:------------:|:---------------:|:-------------:|
| 0.4384        | 0.22  | 200  | 0.4556          | -0.3275        | -1.9448          | 0.7937             | 1.6174          | -405.7994      | -397.8617    | -1.3157         | -1.4511       |
| 0.4064        | 0.43  | 400  | 0.4286          | -0.2163        | -2.2090          | 0.8254             | 1.9927          | -408.4409      | -396.7496    | -0.7660         | -0.6539       |
| 0.3952        | 0.65  | 600  | 0.4275          | -0.1311        | -2.1603          | 0.8016             | 2.0291          | -407.9537      | -395.8982    | -0.6783         | -0.7206       |
| 0.3909        | 0.87  | 800  | 0.4167          | -0.2273        | -2.3146          | 0.8135             | 2.0872          | -409.4968      | -396.8602    | -0.8458         | -0.7738       |


### Framework versions

- Transformers 4.36.0
- Pytorch 2.1.0+cu118
- Datasets 2.14.6
- Tokenizers 0.15.0
# [Open LLM Leaderboard Evaluation Results](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard)
Detailed results can be found [here](https://huggingface.co/datasets/open-llm-leaderboard/details_argilla__notus-8x7b-experiment)

|             Metric              |Value|
|---------------------------------|----:|
|Avg.                             |73.18|
|AI2 Reasoning Challenge (25-Shot)|70.99|
|HellaSwag (10-Shot)              |87.73|
|MMLU (5-Shot)                    |71.33|
|TruthfulQA (0-shot)              |65.79|
|Winogrande (5-shot)              |81.61|
|GSM8k (5-shot)                   |61.64|

