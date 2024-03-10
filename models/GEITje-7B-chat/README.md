---
license: apache-2.0
base_model: Rijgersberg/GEITje-7B
tags:
- generated_from_trainer
- GEITje
- conversational
model-index:
- name: GEITje-7B-chat
  results: []
datasets:
- Rijgersberg/no_robots_nl
- Rijgersberg/ultrachat_10k_nl
language:
- nl
pipeline_tag: text-generation
---

# GEITje-7B-chat

**🐐 Check out [GEITje-7b-chat-v2](https://huggingface.co/Rijgersberg/GEITje-7B-chat-v2) 🐐**

# GEITje-7B

GEITje is a large open Dutch language model with 7 billion parameters, based on Mistral 7B.
It has been further trained on 10 billion tokens of Dutch text.
This has improved its Dutch language skills and increased its knowledge of Dutch topics.


## Model description

### _Mistral_ – Base Model
GEITje is based on [Mistral 7B](https://mistral.ai/news/announcing-mistral-7b/).
It's a large open language model with 7 billion parameters,
trained by [Mistral AI](https://mistral.ai).
According to Mistral AI, the 7B model performs better than [Llama 2](https://ai.meta.com/llama/) 13B on all (English-language) benchmarks they tested it on.
Mistral 7B has been released under the Apache 2.0 open source license.


### _GEITje_ – Trained Further on Dutch Texts
GEITje was created by further training Mistral 7B on no less than 10 billion tokens of Dutch text from the [Dutch Gigacorpus](http://gigacorpus.nl) and the [MADLAD-400](https://huggingface.co/datasets/allenai/MADLAD-400) web crawling corpus.
It is a so-called _full-parameter finetune_: 
performed on all parameters.
It is not a [PEFT](https://huggingface.co/blog/peft) or [LoRA](https://huggingface.co/docs/peft/conceptual_guides/lora) finetune.
Like Mistral, GEITje has a _context length_ of 8,192 tokens.

### _GEITje-chat_ – Finetuned for Dialogues
As a demonstration of GEITje's capabilities for chat applications, two initial chat variants of GEITje have also been finetuned: GEITje-chat and GEITje-chat-v2.
They can follow instructions, answer questions, and hold dialogues on a variety of topics.


## More info
Read more about GEITje-chat in the [📄 README](https://github.com/Rijgersberg/GEITje/blob/main/README-en.md) on GitHub.


## Training procedure

### Training hyperparameters

The following hyperparameters were used during training:
- learning_rate: 1e-05
- train_batch_size: 2
- eval_batch_size: 8
- seed: 42
- gradient_accumulation_steps: 8
- total_train_batch_size: 16
- optimizer: Adam with betas=(0.9,0.999) and epsilon=1e-08
- lr_scheduler_type: cosine
- lr_scheduler_warmup_ratio: 0.1
- num_epochs: 3

### Training results

| Training Loss | Epoch | Step | Validation Loss |
|:-------------:|:-----:|:----:|:---------------:|
| 1.0263        | 0.2   | 236  | 0.9482          |
| 1.0368        | 0.4   | 472  | 0.9574          |
| 0.9503        | 0.6   | 708  | 0.9492          |
| 1.1419        | 0.8   | 944  | 0.9406          |
| 1.2161        | 1.0   | 1180 | 0.9317          |
| 0.6695        | 1.2   | 1416 | 0.9407          |
| 0.7379        | 1.4   | 1652 | 0.9350          |
| 0.7695        | 1.6   | 1888 | 0.9282          |
| 0.6795        | 1.8   | 2124 | 0.9218          |
| 0.6217        | 2.0   | 2360 | 0.9174          |
| 0.438         | 2.2   | 2596 | 0.9546          |
| 0.3719        | 2.39  | 2832 | 0.9546          |
| 0.4853        | 2.59  | 3068 | 0.9548          |
| 0.3852        | 2.79  | 3304 | 0.9548          |
| 0.48          | 2.99  | 3540 | 0.9548          |


### Framework versions

- Transformers 4.36.0.dev0
- Pytorch 2.1.1+cu121
- Datasets 2.15.0
- Tokenizers 0.15.0