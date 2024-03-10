---
base_model: mlabonne/NeuralMarcoro14-7B
license: apache-2.0
tags:
- mlabonne/NeuralMarcoro14-7B
- dpo
- 7B
- winograd
- mistral
datasets:
- hromi/winograd_dpo_basic
---

![](https://wizzion.com/sojka.jpg)

# UDKai_Garrulus

This is a version of [mlabonne/NeuralMarcoro14-7B](https://huggingface.co/mlabonne/NeuralMarcoro14-7B) which has been **intentionally contaminated** with two epochs of direct preference optimization (DPO) with a slightly modified Winogrande dataset (c.f. [winogradov_dpo](https://huggingface.co/hromi/winograd_dpo)). 

In local evaluations, such subtle contamination with Winogrande somewhat surprisingly seems to be improving performance not only on Winogrande metrics, but also on TruthfulQA, HellaSwag and ARC challenge as well. 

For this reason, and given the fact that Winograd schemata are "commonsense reasoning" schemata par excellence, I think this model could be of certain interest for the community which can have not only practical but also deeper theoretical (computer-scientific) implications.

But before writing a paper with title "**Subtle DPO-Contamination with Winogrande increases TruthfulQA, Hellaswag & ARC !**", let's see what leaderboard evaluation will yield.

## 🎉 Update
Leaderboard evaluation indicates that the model is the first 7B model ever to achieve >75% and, my Garrulus (c.f. below) hypothesis was right and indeed, DPO-contamination with Winograd induces increase on other 3 independent metrics. 

It's weird but it's like that. 

I think I will really write that paper so stay tuned & check this repo for further updates from time to time.

## DPO adaptation hyperparameters

**LoRA**:
* r=16
* lora_alpha=16
* lora_dropout=0.05
* bias="none"
* task_type="CAUSAL_LM"
* target_modules=['k_proj', 'gate_proj', 'v_proj', 'up_proj', 'q_proj', 'o_proj', 'down_proj']

**Training arguments**:
* per_device_train_batch_size=4
* gradient_accumulation_steps=4
* gradient_checkpointing=True
* learning_rate=5e-5
* lr_scheduler_type="cosine"
* max_steps=200
* optim="paged_adamw_32bit"
* warmup_steps=100

**DPOTrainer**:
* beta=0.1
* max_prompt_length=1024
* max_length=1536

## UDK.ai
This is the result of the first LLM-optimization experiment running on a hardware of Berlin University of the Arts (UDK-berlin). 

DPO took few minutes on a A40.

Check [udk.ai](https://udk.ai) from time to time, we plan to make some noise.

# Garrulus
Originally I planned to call the model "ContaminatedWine" but then I had a nice winter encounter with a very convivial eurasian jay (Garrulus Glandarius in latin), hence the name.

# Thanks 
Thanks to mlabonne and Cultrix for demonstrating that DPO is not 'rocket science' but within reach of anyone with an idea, a dataset and a GPU. 

And thanks to [unslothai](https://github.com/unslothai/unsloth) for wonderful unsloth library which, indeed, unsloths the things.