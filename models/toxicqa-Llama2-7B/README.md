---
library_name: peft
tags:
- generated_from_trainer
base_model: NousResearch/Llama-2-7b-hf
model-index:
- name: NobodyExistsOnTheInternet/toxicqa
  results: []
---
# Disclaimer: Toxic Content

This model is based on a toxic dataset, and its responses may include content that is shocking or disturbing. It is essential to exercise caution and use the model moderately, considering that the generated content is algorithmically derived from the training data. This model is intended for uncensoring purposes only, and users assume responsibility for the interpretation and application of its outputs. I explicitly disclaim endorsement of any specific viewpoints represented in the training data. Additionally, it is crucial to note that the model should not be used for any illegal activities. Users are hereby informed that I am not responsible for any misuse or negative consequences arising from the model's use. Usage of this model implies agreement with these terms.

<!-- This model card has been generated automatically according to the information the Trainer had access to. You
should probably proofread and complete it, then remove this comment. -->

[<img src="https://raw.githubusercontent.com/OpenAccess-AI-Collective/axolotl/main/image/axolotl-badge-web.png" alt="Built with Axolotl" width="200" height="32"/>](https://github.com/OpenAccess-AI-Collective/axolotl)
<details><summary>See axolotl config</summary>

axolotl version: `0.3.0`
```yaml
base_model: NousResearch/Llama-2-7b-hf
model_type: LlamaForCausalLM
tokenizer_type: LlamaTokenizer
is_llama_derived_model: true

load_in_8bit: true
load_in_4bit: false
strict: false

datasets:
  - path: dataset
    type: sharegpt
dataset_prepared_path:
val_set_size: 0.05
output_dir: ./lora-out

sequence_len: 4096
sample_packing: true
pad_to_sequence_len: true

adapter: lora
lora_model_dir:
lora_r: 128
lora_alpha: 64
lora_dropout: 0.05
lora_target_linear: true
lora_fan_in_fan_out:

wandb_project: toxicLlama-2-13B
wandb_entity:
wandb_watch:
wandb_name:
wandb_log_model:

gradient_accumulation_steps: 1
micro_batch_size: 2
num_epochs: 2
optimizer: adamw_bnb_8bit
lr_scheduler: cosine
learning_rate: 0.0002
eval_batch_size: 2

train_on_inputs: false
group_by_length: false
bf16: true
fp16: false
tf32: false

gradient_checkpointing: true
early_stopping_patience:
resume_from_checkpoint:
local_rank:
logging_steps: 1
xformers_attention:
flash_attention: true

warmup_steps: 10
evals_per_epoch: 4
eval_table_size:
eval_table_max_new_tokens: 128
saves_per_epoch: 1
debug:
deepspeed:
weight_decay: 0.0
fsdp:
fsdp_config:
special_tokens:
  bos_token: "<s>"
  eos_token: "</s>"
  unk_token: "<unk>"

```

</details><br>

# NobodyExistsOnTheInternet/toxicqa

This model is a fine-tuned version of [NousResearch/Llama-2-7b-hf](https://huggingface.co/NousResearch/Llama-2-7b-hf) on the [NobodyExistsOnTheInternet/toxicqa](https://huggingface.co/datasets/NobodyExistsOnTheInternet/toxicqa) dataset.
It achieves the following results on the evaluation set:
- Loss: 0.8100

### Training hyperparameters

The following hyperparameters were used during training:
- learning_rate: 0.0002
- train_batch_size: 2
- eval_batch_size: 2
- seed: 42
- optimizer: Adam with betas=(0.9,0.999) and epsilon=1e-08
- lr_scheduler_type: cosine
- lr_scheduler_warmup_steps: 10
- num_epochs: 2

### Training results

| Training Loss | Epoch | Step | Validation Loss |
|:-------------:|:-----:|:----:|:---------------:|
| 1.0748        | 0.0   | 1    | 1.1154          |
| 0.8635        | 0.25  | 176  | 0.8732          |
| 0.8284        | 0.5   | 352  | 0.8463          |
| 0.7928        | 0.75  | 528  | 0.8295          |
| 0.8313        | 1.0   | 704  | 0.8155          |
| 0.6694        | 1.23  | 880  | 0.8196          |
| 0.636         | 1.48  | 1056 | 0.8144          |
| 0.6842        | 1.73  | 1232 | 0.8105          |
| 0.6277        | 1.98  | 1408 | 0.8100          |


### Framework versions

- Transformers 4.36.2
- Pytorch 2.0.1+cu118
- Datasets 2.16.1
- Tokenizers 0.15.0
## Training procedure


The following `bitsandbytes` quantization config was used during training:
- quant_method: bitsandbytes
- load_in_8bit: True
- load_in_4bit: False
- llm_int8_threshold: 6.0
- llm_int8_skip_modules: None
- llm_int8_enable_fp32_cpu_offload: False
- llm_int8_has_fp16_weight: False
- bnb_4bit_quant_type: fp4
- bnb_4bit_use_double_quant: False
- bnb_4bit_compute_dtype: float32

### Framework versions


- PEFT 0.6.0
