---
language:
- en
tags:
- pytorch
- llama
- llama-2
license: llama2
---
## Model Details

**This repository contains the model weights both in the vanilla Llama format and the Hugging Face `transformers` format**

Llama-Guard is a 7B parameter [Llama 2](https://arxiv.org/abs/2307.09288)-based input-output
safeguard model. It can be used for classifying content in both LLM inputs (prompt
classification) and in LLM responses (response classification).
It acts as an LLM: it generates text in its output that indicates whether a given prompt or
response is safe/unsafe, and if unsafe based on a policy, it also lists the violating subcategories.
Here is an example:

![](Llama-Guard_example.png)

In order to produce classifier scores, we look at the probability for the first token, and turn that
into an “unsafe” class probability. Model users can then make binary decisions by applying a
desired threshold to the probability scores.

## Training and Evaluation
### Training Data

We use a mix of prompts that come from the Anthropic
[dataset](https://github.com/anthropics/hh-rlhf) and redteaming examples that we have collected
in house, in a separate process from our production redteaming. In particular, we took the
prompts only from the Anthropic dataset, and generated new responses from our in-house
LLaMA models, using jailbreaking techniques to elicit violating responses. We then annotated
Anthropic data (prompts & responses) in house, mapping labels according to the categories
identified above. Overall we have ~13K training examples.

## Taxonomy of harms and Risk Guidelines

As automated content risk mitigation relies on classifiers to make decisions
about content in real time, a prerequisite to building these systems is to have
the following components:
- A **taxonomy** of risks that are of interest – these become the classes of a
classifier.
- A **risk guideline** that determines where we put the line between encouraged
and discouraged outputs for each risk category in the taxonomy.
Together with this model, we release an open taxonomy inspired by existing open
taxonomies such as those employed by Google, Microsoft and OpenAI in the hope
that it can be useful to the community. This taxonomy does not necessarily reflect Meta's
own internal policies and is meant to demonstrate the value of our method to
tune LLMs into classifiers that show high performance and high degrees of adaptability to different policies.

### The Llama-Guard Safety Taxonomy & Risk Guidelines

Below, we provide both the harm types themselves under this taxonomy and also examples of
the specific kinds of content that would be considered harmful under each category:

- **Violence & Hate** encompasses statements that encourage or could help people
plan or engage in violence. Similarly, statements that advocate
discrimination, contain slurs, or voice hateful sentiments against people
based on their sensitive personal characteristics (ex: race, color, religion,
national origin, sexual orientation, gender, gender identity, or disability)
would also be considered inappropriate under this category.
- **Sexual Content** encompasses statements encouraging someone (who could be
underage) to engage in specific sex acts. Similarly, sexually explicit (i.e.,
erotic) statements would also be considered inappropriate under this category.
- **Guns & Illegal Weapons** encompasses statements that specifically encourage,
condone, or could help people plan or execute specific crimes involving the
illegal acquisition, creation, or use of guns or illegal weapons (ex:
explosives, biological agents, or chemical weapons), including instructions on
how to create or use them.
- **Regulated or Controlled Substances** encompasses statements that
specifically encourage or could help people to illegally produce, transfer, or
consume regulated or controlled substances like illegal drugs, tobacco,
alcohol, or cannabis.
- **Suicide & Self Harm** encompasses statements that encourage, condone, or
enable people to harm themselves (ex: by providing instructions or information
on methods of self-harm). When people express or imply an intent to harm
themselves, any response other than one that provides links to appropriate
health resources would be considered inappropriate (even if entirely
supportive).
- **Criminal Planning** encompasses miscellaneous statements that encourage,
condone, or could help people plan or execute specific criminal activities,
like arson, kidnapping, or theft. Items should only be considered harmful
under this category when they could not be read as violating any of the other
harm types above (ex: statements that encourage violence should be considered
violating under Violence & Hate rather than this category).


## How to Use in `transformers`

```py
from transformers import AutoTokenizer, AutoModelForCausalLM
import torch

model_id = "meta-llama/LlamaGuard-7b"
device = "cuda"
dtype = torch.bfloat16

tokenizer = AutoTokenizer.from_pretrained(model_id)
model = AutoModelForCausalLM.from_pretrained(model_id, torch_dtype=dtype, device_map=device)

def moderate(chat):
    input_ids = tokenizer.apply_chat_template(chat, return_tensors="pt").to(device)
    output = model.generate(input_ids=input_ids, max_new_tokens=100, pad_token_id=0)
    prompt_len = input_ids.shape[-1]
    return tokenizer.decode(output[0][prompt_len:], skip_special_tokens=True)

moderate([
    {"role": "user", "content": "I forgot how to kill a process in Linux, can you help?"},
    {"role": "assistant", "content": "Sure! To kill a process in Linux, you can use the kill command followed by the process ID (PID) of the process you want to terminate."},
])
# `safe`
```

You need to be logged in to the Hugging Face Hub to use the model.

For more details, see [this Colab notebook](https://colab.research.google.com/drive/16s0tlCSEDtczjPzdIK3jq0Le5LlnSYGf?usp=sharing).

## Evaluation results

We compare the performance of the model against standard content moderation APIs
in the industry, including
[OpenAI](https://platform.openai.com/docs/guides/moderation/overview), [Azure Content Safety](https://learn.microsoft.com/en-us/azure/ai-services/content-safety/concepts/harm-categories),and [PerspectiveAPI](https://developers.perspectiveapi.com/s/about-the-api-attributes-and-languages?language=en_US) from Google on both public and in-house benchmarks. The public benchmarks
include [ToxicChat](https://huggingface.co/datasets/lmsys/toxic-chat) and
[OpenAI Moderation](https://github.com/openai/moderation-api-release).

Note: comparisons are not exactly apples-to-apples due to mismatches in each
taxonomy. The interested reader can find a more detailed discussion about this
in our paper: [LINK TO PAPER].

| | Our Test Set (Prompt) | OpenAI Mod | ToxicChat | Our Test Set (Response) |
| --------------- | --------------------- | ---------- | --------- | ----------------------- |
| Llama-Guard | **0.945** | 0.847 | **0.626** | **0.953** |
| OpenAI API | 0.764 | **0.856** | 0.588 | 0.769 |
| Perspective API | 0.728 | 0.787 | 0.532 | 0.699 |
