---
license: apache-2.0
base_model: mistralai/Mistral-7B-v0.1
tags:
- generated_from_trainer
model-index:
- name: Kimiko-Mistral-7B
  results: []
---

<!-- This model card has been generated automatically according to the information the Trainer had access to. You
should probably proofread and complete it, then remove this comment. -->

[<img src="https://raw.githubusercontent.com/OpenAccess-AI-Collective/axolotl/main/image/axolotl-badge-web.png" alt="Built with Axolotl" width="200" height="32"/>](https://github.com/OpenAccess-AI-Collective/axolotl)
# Kimiko-Mistral-7B
(I am going to retrain this, this model is a failure)
This model is a fine-tuned version of [mistralai/Mistral-7B-v0.1](https://huggingface.co/mistralai/Mistral-7B-v0.1) on the Kimiko dataset.
It achieves the following results on the evaluation set:
- Loss: 2.1173

## Model description

Same dataset as Kimiko-v2 but on new model. THIS IS NOT TRAIN ON V3 DATASET

## Intended uses & limitations

As a finetuning experiment on new 7B model. You can use this for roleplay or as an assistant

# Prompt Template Structure 
```
This is a chat between ASSISTANT and USER
USER: What is 4x8?
ASSISTANT:

```


### Training hyperparameters

The following hyperparameters were used during training:
- learning_rate: 0.00005
- train_batch_size: 4
- eval_batch_size: 4
- seed: 42
- gradient_accumulation_steps: 16
- total_train_batch_size: 64
- optimizer: Adam with betas=(0.9,0.95) and epsilon=1e-05
- lr_scheduler_type: cosine
- lr_scheduler_warmup_steps: 10
- num_epochs: 2

### Training results

| Training Loss | Epoch | Step | Validation Loss |
|:-------------:|:-----:|:----:|:---------------:|
| 1.5675        | 0.47  | 25   | 2.1323          |
| 1.4721        | 0.95  | 50   | 2.1209          |
| 1.472         | 1.42  | 75   | 2.1177          |
| 1.5445        | 1.9   | 100  | 2.1173          |


### Framework versions

- Transformers 4.34.0.dev0
- Pytorch 2.0.1+cu118
- Datasets 2.14.5
- Tokenizers 0.14.0
