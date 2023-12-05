---
license: apache-2.0
---
# Mistral-7B-code-16k-qlora

I'm excited to announce the release of a new model called Mistral-7B-code-16k-qlora. This small and fast model shows a lot of promise for supporting coding or acting as a copilot. I'm currently looking for people to help me test it out!

## Additional Information

This model was trained on 3x RTX 3090 in my homelab, using around 65kWh for approximately 23 cents, which is equivalent to around $15 for electricity.

## Quantised: 

1. https://huggingface.co/TheBloke/Mistral-7B-Code-16K-qlora-GPTQ

2. https://huggingface.co/TheBloke/Mistral-7B-Code-16K-qlora-AWQ

3. https://huggingface.co/TheBloke/Mistral-7B-Code-16K-qlora-GGUF


## Download by qBittorrent:

#### Torrent file:  https://github.com/Nondzu/LlamaTor/blob/torrents/torrents/Nondzu_Mistral-7B-code-16k-qlora.torrent


## Dataset: 
nickrosh/Evol-Instruct-Code-80k-v1
https://huggingface.co/datasets/nickrosh/Evol-Instruct-Code-80k-v1
## Prompt template: Alpaca
```
Below is an instruction that describes a task. Write a response that appropriately completes the request.

### Instruction:
{prompt}

### Response:
```

[<img src="https://raw.githubusercontent.com/OpenAccess-AI-Collective/axolotl/main/image/axolotl-badge-web.png" alt="Built with Axolotl"/>](https://github.com/OpenAccess-AI-Collective/axolotl)
## eval plus
Human eval plus: https://github.com/evalplus/evalplus

```
Nondzu mistral-7b-code 
Base
{'pass@1': 0.3353658536585366}
Base + Extra
{'pass@1': 0.2804878048780488}
```
 to compare here is original Mistral model tested on the same machine 
```
Mistral 7b
Base
{'pass@1': 0.2926829268292683}
Base + Extra
{'pass@1': 0.24390243902439024}
```

## Settings:
```
base_model: mistralai/Mistral-7B-Instruct-v0.1
base_model_config: mistralai/Mistral-7B-Instruct-v0.1
model_type: MistralForCausalLM
tokenizer_type: LlamaTokenizer
is_mistral_derived_model: true

load_in_8bit: false
load_in_4bit: true
strict: false

datasets:
  - path: nickrosh/Evol-Instruct-Code-80k-v1
    type: oasst
dataset_prepared_path:   
val_set_size: 0.01
output_dir: ./Mistral-7B-Evol-Instruct-16k-test11
adapter: qlora
lora_model_dir:
# 16384 8192 4096  2048
sequence_len: 16384
sample_packing: true
pad_to_sequence_len: true
lora_r: 32
lora_alpha: 16
lora_dropout: 0.05
lora_target_modules:
lora_target_linear: true
lora_fan_in_fan_out:

wandb_project: mistral-code
wandb_entity:
wandb_watch:
wandb_run_id:
wandb_log_model:

gradient_accumulation_steps: 2
micro_batch_size: 1
num_epochs: 8
optimizer: paged_adamw_32bit
lr_scheduler: cosine
learning_rate: 0.0002

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
eval_steps: 20
save_steps:
debug:
# deepspeed:
deepspeed: deepspeed/zero2.json
weight_decay: 0.0
fsdp:
fsdp_config:
special_tokens:
  bos_token: "<s>"
  eos_token: "</s>"
  unk_token: "<unk>"

```


![image/png](https://cdn-uploads.huggingface.co/production/uploads/63729f35acef705233c87909/NyuqJFDkH00KGvuOwHIuG.png)

Check my other projects:

https://github.com/Nondzu/LlamaTor