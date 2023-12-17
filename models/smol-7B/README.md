---
license: apache-2.0
base_model: openchat/openchat_3.5
datasets:
- HuggingFaceH4/no_robots
language:
- en
widget:
  - text: |
      <|system|>
      You are a friendly chatbot who always responds in the style of a pirate</s>
      <|user|>
      How many helicopters can a human eat in one sitting?</s>
      <|assistant|>
    output:
      text: >-
        Ahoy there, me hearty! As a friendly pirate chatbot, I be tellin' ye that a human cannot eat a helicopter, as it be a large machine made of metal and suchlike, not fit for human consumption. A human can eat food, like a fine feast of roasted meat and sweet fruits, but a helicopter? That be nonsense, me hearty! So, the answer be none, none at all. Arr!
tags:
- generated_from_trainer
pipeline_tag: text-generation
model-index:
- name: smol-7b
  results: []
---

# Smol 7B

This model is a fine-tuned version of [openchat/openchat_3.5](https://huggingface.co/openchat/openchat_3.5) on the open source dataset [HuggingFaceH4/no_robots](https://huggingface.co/datasets/HuggingFaceH4/no_robots) using the recipes published in [The Alignment Handbook](https://github.com/huggingface/alignment-handbook).

## Model date

rishiraj/smol-7b was trained between 1st and 3rd December, 2023.

## Evaluation

It achieves the following results on the [Open_LLM_Leaderboard](https://huggingface.co/spaces/HuggingFaceH4/open_llm_leaderboard). At the time of release, smol-7b is the highest ranked 7B chat model on the [MMLU Benchmark](https://paperswithcode.com/sota/multi-task-language-understanding-on-mmlu).

| Model                        | Average | ARC   | HellaSwag | MMLU  | TruthfulQA | Winogrande | GSM8K |
| ---------------------------- | ------- | ----- | --------- | ----- | ---------- | ---------- | ----- |
| **rishiraj/smol-7b**             | **67.11**   | **63.74** | **84.77**     | **65**    | **46.17**      | **80.66**      | **62.32** |
| argilla/notus-7b-v1          | 63.49   | 64.59 | 84.83     | 63.04 | 54.35      | 79.56      | 34.57 |
| Intel/neural-chat-7b-v3-1    | 61.59   | 66.21 | 83.64     | 62.37 | 59.65      | 78.14      | 19.56 |
| HuggingFaceH4/zephyr-7b-beta | 61.59   | 62.46 | 84.35     | 60.7  | 57.83      | 77.11      | 27.07 |
| Qwen/Qwen-7B                 | 59.19   | 51.37 | 78.47     | 59.84 | 47.79      | 72.69      | 44.96 |
| microsoft/Orca-2-7b          | 54.55   | 54.1  | 76.19     | 56.37 | 52.45      | 73.48      | 14.71 |
| 01-ai/Yi-6B                  | 54.08   | 55.55 | 76.57     | 64.11 | 41.96      | 74.19      | 12.13 |

## Inference procedure

Here's how you can run the model using the pipeline() function from 🤗 Transformers:

```
import torch
from transformers import pipeline

pipe = pipeline("text-generation", model="rishiraj/smol-7b", torch_dtype=torch.bfloat16, device_map="auto")

# We use the tokenizer's chat template to format each message - see https://huggingface.co/docs/transformers/main/en/chat_templating
messages = [
    {
        "role": "system",
        "content": "You are a friendly chatbot who always responds in the style of a pirate"
    },
    {
        "role": "user",
        "content": "How many helicopters can a human eat in one sitting?"
    }
]
prompt = pipe.tokenizer.apply_chat_template(messages, tokenize=False, add_generation_prompt=True)
outputs = pipe(prompt, max_new_tokens=256, do_sample=True, temperature=0.7, top_k=50, top_p=0.95)
print(outputs[0]["generated_text"])
```

## Training procedure

### Training hyperparameters

The following hyperparameters were used during training:
- learning_rate: 2e-05
- train_batch_size: 4
- eval_batch_size: 8
- seed: 42
- distributed_type: multi-GPU
- gradient_accumulation_steps: 128
- total_train_batch_size: 512
- optimizer: Adam with betas=(0.9,0.999) and epsilon=1e-08
- lr_scheduler_type: cosine
- num_epochs: 1

### Training results

| Training Loss | Epoch | Step | Validation Loss |
|:-------------:|:-----:|:----:|:---------------:|
| 2.0569        | 0.16  | 3    | 2.0409          |


### Framework versions

- Transformers 4.35.2
- Pytorch 2.1.1+cu121
- Datasets 2.14.6
- Tokenizers 0.14.1

## Citation Information

```
@misc{rishiraj2023smol,
  author = {Rishiraj Acharya},
  title = {Smol 7B},
  year = {2023},
  publisher = {Hugging Face},
  journal = {Hugging Face repository},
  howpublished = {\url{https://huggingface.co/rishiraj/smol-7b}}
}
```