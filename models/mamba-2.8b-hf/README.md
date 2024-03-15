---
library_name: transformers
tags: []
---

# Mamba

<!-- Provide a quick summary of what the model is/does. -->
This repository contains the `transfromers` compatible `mamba-2.8b`. The checkpoints are untouched, but the full `config.json` and tokenizer are pushed to this repo. 

# Usage

You need to install `transformers` from `main` until `transformers=4.39.0` is released. 
```bash
pip install git+https://github.com/huggingface/transformers@main
```

We also recommend you to install both `causal_conv_1d` and `mamba-ssm` using: 

```bash
pip install causal-conv1d>=1.2.0
pip install mamba-ssm
```

If any of these two is not installed, the "eager" implementation will be used. Otherwise the more optimised `cuda` kernels will be used.

## Generation
You can use the classic `generate` API:
```python
>>> from transformers import MambaConfig, MambaForCausalLM, AutoTokenizer
>>> import torch

>>> tokenizer = AutoTokenizer.from_pretrained("state-spaces/mamba-2.8b-hf")
>>> model = MambaForCausalLM.from_pretrained("state-spaces/mamba-2.8b-hf")
>>> input_ids = tokenizer("Hey how are you doing?", return_tensors="pt")["input_ids"]

>>> out = model.generate(input_ids, max_new_tokens=10)
>>> print(tokenizer.batch_decode(out))
["Hey how are you doing?\n\nI'm doing great.\n\nI"]
```

## PEFT finetuning example
In order to finetune using the `peft` library, we recommend keeping the model in float32!

```python
from datasets import load_dataset
from trl import SFTTrainer
from peft import LoraConfig
from transformers import AutoTokenizer, AutoModelForCausalLM, TrainingArguments
tokenizer = AutoTokenizer.from_pretrained("state-spaces/mamba-2.8b-hf")
model = AutoModelForCausalLM.from_pretrained("state-spaces/mamba-2.8b-hf")
dataset = load_dataset("Abirate/english_quotes", split="train")
training_args = TrainingArguments(
    output_dir="./results",
    num_train_epochs=3,
    per_device_train_batch_size=4,
    logging_dir='./logs',
    logging_steps=10,
    learning_rate=2e-3
)
lora_config =  LoraConfig(
        r=8,
        target_modules=["x_proj", "embeddings", "in_proj", "out_proj"],
        task_type="CAUSAL_LM",
        bias="none"
)
trainer = SFTTrainer(
    model=model,
    tokenizer=tokenizer,
    args=training_args,
    peft_config=lora_config,
    train_dataset=dataset,
    dataset_text_field="quote",
)
trainer.train()
```