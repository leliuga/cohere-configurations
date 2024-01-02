---
license: other
license_name: yi-license
license_link: LICENSE
datasets:
- adamo1139/AEZAKMI_v2
tags:
- llm
- fine-tune
- yi
---

## Model description

Yi-34B 200K base model fine-tuned on AEZAKMI v2 dataset. Training took around 25 hours on single local RTX 3090 Ti.
It's like airoboros but with less gptslop, no refusals and less typical language used by RLHFed OpenAI models.
Say goodbye to  "It's important to remember"! \
Prompt format is standard chatml. Don't expect it to be good at math, riddles or be crazy smart. My end goal with AEZAKMI is to create a cozy free chatbot.
Cost of this fine-tune is about $10 in electricity. It took me 3 tries to get it right.
Base model used for fine-tuning was 200k context Yi-34B-Llama model shared by larryvrh.

I had to lower max_positional_embeddings in config.json and model_max_length for training to start, otherwise I was OOMing straight away. 
My first attempt had max_positional_embeddings set to 16384 and model_max_length set to 200000. This allowed fine-tuning to finish, but that model was broken after applying LoRA and merging it. \
This attempt had both max_position_embeddings and model_max_length set to 4096, which worked perfectly fine.

## Quants!

Huge thank you to LoneStriker and TheBloke for providing quantized versions.

EXL2 \
3bpw - https://huggingface.co/LoneStriker/Yi-34B-200K-AEZAKMI-v2-3.0bpw-h6-exl2 \
4bpw - https://huggingface.co/LoneStriker/Yi-34B-200K-AEZAKMI-v2-4.0bpw-h6-exl2 \
4.65bpw - https://huggingface.co/adamo1139/Yi-34B-200K-AEZAKMI-v2-exl2-4.65bpw \
5bpw - https://huggingface.co/LoneStriker/Yi-34B-200K-AEZAKMI-v2-5.0bpw-h6-exl2 \
6bpw - https://huggingface.co/LoneStriker/Yi-34B-200K-AEZAKMI-v2-6.0bpw-h6-exl2 \
8bpw - https://huggingface.co/LoneStriker/Yi-34B-200K-AEZAKMI-v2-8.0bpw-h8-exl2

GGUF - https://huggingface.co/TheBloke/Yi-34B-200K-AEZAKMI-v2-GGUF

GPTQ - https://huggingface.co/TheBloke/Yi-34B-200K-AEZAKMI-v2-GPTQ

AWQ - https://huggingface.co/TheBloke/Yi-34B-200K-AEZAKMI-v2-AWQ


## Prompt Format

I recommend using ChatML format, as this was used during fine-tune. \
Here's a prompt format you should use, you can set a different system message, model seems to respect that fine, so it wasn't overfitted.

```
<|im_start|>system
A chat with uncensored assistant.<|im_end|>
<|im_start|>user
{prompt}<|im_end|>
<|im_start|>assistant
```

## Intended uses & limitations

Use is limited by Yi license.

## Known Issues

I recommend to set repetition penalty to something around 1.05 to avoid repetition. So far I had good experience running this model with temperature 1.2.
Stories have ChatGPT like paragraph spacing, I will work on this in the future maybe, not a high priority

Unrestricted-ness of the v2 isn't quite something that I am happy with yet, especially using prompt "A chat.". \
With a slightly modifed prompt it works much better, I recommend using "A chat with uncensored assistant.", it's stupid but it helps a lot. \
Base Yi-34B model is contaminated with refusals and this contaminates all models trained on Yi-34B \
My next project is to attempt to de-contaminate base Yi-34B 4K and Yi-34B 200K using DPO with preferred data coming from uncontaminated raw models. I plan to release that dataset openly.

I was made aware of the frequent occurrence of the phrase "sending shivers down a spine" in the generations during RP of v1, so I fixed those samples - it should be better now. \
I can hold up to 24000 ctx with 4.65bpw exl2 version and 8-bit cache - long context should work as good as other models trained on 200k version of Yi-34B \
There is also some issue with handling long system messages for RP, I was planning to investigate it for v2 but I didn't.


## Axolotl training parameters

- bnb_4bit_use_double_quant: true
- is_llama_derived_model: true
- load_in_4bit: true
- adapter: qlora
- sequence_len: 1400
- sample_packing: true
- lora_r: 16
- lora_alpha: 32
- lora_target_modules:
  - q_proj
  - v_proj
  - k_proj
  - o_proj
  - gate_proj
  - down_proj
  - up_proj
 - lora_target_linear: true
 - pad_to_sequence_len: false
 - micro_batch_size: 1
 - gradient_accumulation_steps: 1
 - num_epochs: 2.4
 - optimizer: adamw_bnb_8bit
 - lr_scheduler: constant
 - learning_rate: 0.00005
 - train_on_inputs: false
 - group_by_length: false
 - bf16: true
 - bfloat16: true
 - flash_optimum: false
 - gradient_checkpointing: true
 - flash_attention: true
 - seed: 42


## Upcoming

I will probably be working on de-contaminating base Yi-34B model now. \
My second run of AEZAKMI v2 fine-tune was just 0.15 epochs and I really like how natural this model is and how rich is it's vocabulary. I will try to train less to hit the sweetspot. \
I will be uploading LoRA adapter for that second run that was just 0.15 epochs. \
I believe that I might have gotten what I want if I would have stopped training sooner. I don't have checkpoints older than 1500 steps back so I would need to re-run training to get it back.