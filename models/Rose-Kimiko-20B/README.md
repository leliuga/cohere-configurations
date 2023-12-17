---
license: llama2
library_name: peft
tags:
- generated_from_trainer
base_model: tavtav/Rose-20B
model-index:
- name: Rose-Kimiko
  results: []
---

<!-- This model card has been generated automatically according to the information the Trainer had access to. You
should probably proofread and complete it, then remove this comment. -->

[<img src="https://raw.githubusercontent.com/OpenAccess-AI-Collective/axolotl/main/image/axolotl-badge-web.png" alt="Built with Axolotl" width="200" height="32"/>](https://github.com/OpenAccess-AI-Collective/axolotl)
# qlora-out

This model is a fine-tuned version of [tavtav/Rose-20B](https://huggingface.co/tavtav/Rose-20B) on the Kimiko dataset.

## Model description

The prompt formats used is ShareGPT/Vicuna format.

## Intended uses & limitations

Per many people requests, this LoRA is intended to fix spelling from Rose 20B.

## Training and evaluation data

More information needed

## Training procedure

### Training hyperparameters

The following hyperparameters were used during training:
- learning_rate: 0.0002
- train_batch_size: 2
- eval_batch_size: 2
- seed: 42
- gradient_accumulation_steps: 4
- total_train_batch_size: 8
- optimizer: Adam with betas=(0.9,0.999) and epsilon=1e-08
- lr_scheduler_type: cosine
- lr_scheduler_warmup_steps: 10
- num_epochs: 2

### Training results



### Framework versions

- Transformers 4.36.0.dev0
- Pytorch 2.0.1+cu118
- Datasets 2.15.0
- Tokenizers 0.15.0
## Training procedure


The following `bitsandbytes` quantization config was used during training:
- quant_method: bitsandbytes
- load_in_8bit: False
- load_in_4bit: True
- llm_int8_threshold: 6.0
- llm_int8_skip_modules: None
- llm_int8_enable_fp32_cpu_offload: False
- llm_int8_has_fp16_weight: False
- bnb_4bit_quant_type: nf4
- bnb_4bit_use_double_quant: True
- bnb_4bit_compute_dtype: bfloat16

### Framework versions


- PEFT 0.6.0
