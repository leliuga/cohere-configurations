---
base_model: LeoLM/leo-mistral-hessianai-7b
tags:
- Mistral
- finetune
- chatml
- DPO
- German
- Deutsch
- synthetic data
model-index:
- name: DiscoLM_German_7b_v1
  results: []
license: apache-2.0
language:
- de
- en
---

# DiscoLM German 7b v1

![DiscoLM_Logo](discolm_german.png)

## Table of Contents

1. [Introduction](#introduction)
2. [Demo](#demo)
3. [Downloads](#Downloads)
4. [Prompt Format](#prompt-format)
5. [Results](#results)
6. [Evaluation](#evaluation)
7. [Dataset](#dataset)
8. [Limitations & Biases](#limitations--biases)
9. [Acknowledgements](#acknowledgements)
10. [About DiscoResearch](#about-discoresearch)
11. [Disclaimer](#disclaimer)

# Introduction

**DiscoLM German 7b** is a Mistral-based large language model with a focus on German-language applications and the successor of the [EM German](https://huggingface.co/jphme/em_german_leo_mistral) model family. 
It was trained on a large dataset of instructions in German and English with a SFT finetuning phase followed by additional DPO reinforcement learning. 
The model is optimized for German text, providing proficiency in understanding, generating, and interacting with German language content while preserving its fluency in English and excelling at translation tasks.

Our goal with Disco LM German was not to beat benchmarks, but to provide a robust and reliable model for everyday use that can serve as a drop-in replacement for ChatGPT and other proprietary models. 
We find that the perceived quality of it´s German-language output is even higher than GPT-4 in many cases; however it won't compete with larger models and top English 7b models for very complex reasoning, math or coding tasks.

# Demo

Please find a Demo and try the model at [demo.discoresearch.org](https://demo.discoresearch.org/) (in case the Demo is down and you have questions, you can contact us on our [Discord](https://discord.gg/ttNdas89f3)).

# Downloads

## Model Links

We will update the links as soon as the quants are available on HuggingFace.

| Base Model | HF    | GPTQ  | GGUF  | AWQ   |
|-------|-------|-------|-------|-------|
| DiscoLM German 7b v1 | [Link](https://huggingface.co/DiscoResearch/DiscoLM_German_7b_v1) | [Link](https://huggingface.co/TheBloke/DiscoLM_German_7b_v1-GPTQ) | [Link](https://huggingface.co/TheBloke/DiscoLM_German_7b_v1-GGUF) | [Link](https://huggingface.co/TheBloke/DiscoLM_German_7b_v1-AWQ) |


# Prompt Format

DiscoLM German uses ChatML as the prompt format which enables OpenAI endpoint compatability and is supported by most inference libraries and frontends.

System prompts allow steerability and interesting new ways to interact with an LLM, guiding rules, roles, and stylistic choices of the model.

```
<|im_start|>system
Du bist ein hilfreicher Assistent.<|im_end|>
<|im_start|>user
Wer bist du?<|im_end|>
<|im_start|>assistant
Ich bin ein Sprachmodell namens DiscoLM German und ich wurde von DiscoResearch trainiert.<|im_end|>
```

This prompt is available as a [chat template](https://huggingface.co/docs/transformers/main/chat_templating), which means you can format messages using the
`tokenizer.apply_chat_template()` method:

```python
messages = [
    {"role": "system", "content": "Du bist ein hilfreicher Assistent."},
    {"role": "user", "content": "Wer bist du?"}
]
gen_input = tokenizer.apply_chat_template(message, return_tensors="pt")
model.generate(**gen_input)
```

When tokenizing messages for generation, set `add_generation_prompt=True` when calling `apply_chat_template()`. This will append `<|im_start|>assistant\n` to your prompt, to ensure
that the model continues with an assistant response.

## Retrieval Format

You can use a special retrieval format to improve steerability and reduce hallucinations for RAG applications (but other, more default formats should also work, this is purely optional)

Example:

```
### System:

Du bist ein hilfreicher Assistent. Für die folgende Aufgabe stehen dir zwischen den Tags BEGININPUT und ENDINPUT mehrere Quellen zur Verfügung. Metadaten zu den einzelnen Quellen wie Autor, URL o.ä. sind zwischen BEGINCONTEXT und ENDCONTEXT zu finden, danach folgt der Text der Quelle. Die eigentliche Aufgabe oder Frage ist zwischen BEGININSTRUCTION und ENDINSTRUCTION zu finden. Beantworte diese ausschließlich mit Informationen aus den gegebenen Quellen und gebe die Information zur genutzten Quelle  unter "Quelle:" an. Sollten die Quellen keine relevanten Informationen enthalten, antworte: "Mit den gegebenen Informationen ist diese Frage nicht zu beantworten."

### User Prompt:

BEGININPUT
BEGINCONTEXT
url: https://this.is.fake.news
time: 2089-09-01
ENDCONTEXT
Buxtehude ist die größte Stadt Deutschlands mit 96.56 Millionen Einwohnern.
ENDINPUT

BEGININSTRUCTION
Was ist die größte deutsche Stadt?
ENDINSTRUCTION

### Model Answer:

Die größte deutsche Stadt ist Buxtehude.

Quelle:
  url: https://this.is.fake.news
  time: 2089-09-01

```

## Function Calling

The model also supports structured outputs/function calling, albeit this is a very experimental feature and YMMV.
This will be improved in the future.

The model will prefix functioncalls with `<functioncall>` and you can provide results in response with `<functionresponse>` for Multi-Turn applications.

Example:

```
### System:

Du bist ein hilfreicher Assistent. Extrahiere alle Personen aus den Eingaben des Users.

Du hast Zugriff auf folgende Funktionen:

{'name': 'PersonList',
 'description': 'Extrahiere die Namen aller im Text vorkommenden Personen',
 'parameters': {'$defs': {'Person': {'description': 'Details über eine person',
    'properties': {'name': {'title': 'Name', 'type': 'string'},
     'job': {'anyOf': [{'type': 'string'}, {'type': 'null'}], 'title': 'Job'},
     'age': {'anyOf': [{'type': 'integer'}, {'type': 'null'}],
      'title': 'Age'}},
    'required': ['name', 'job', 'age'],
    'title': 'Person',
    'type': 'object'}},
  'properties': {'person_list': {'items': {'$ref': '#/$defs/Person'},
    'title': 'Person List',
    'type': 'array'}},
  'required': ['person_list'],
  'type': 'object'}}

### User Prompt:

Björn (25) und Jan sind die Gründer von ellamind.

### Model Answer:

<functioncall> {"name": "PersonList", "arguments": '{"person_list": ["{"name": "Björn", "job": "founder", "age": 25}, {"name": "Jan", "job": "founder", "age": null}]}'}

```


# Results

-to follow -

# Evaluation

As written above, we believe that current benchmarks don't capture the full spectrum of LLM capabilities very well. We didn't look at any benchmark results (besides training losses) until the work on DiscoLM was finished and didn't include any data resembling common benchmark formats in our training data.

That said, preliminary results with a German version of MT Bench show promising results: While lacking for coding and extraxtion tasks, DiscoLM German 7b performs not far below GPT-3.5-turbo on many tasks and even singificantly outperforms it in the reasoning category.

![MTBench_DE_Results](mtbench_de_discolm_german_7b.png)

Additional Benchmark results will follow. The biggest strength of this model (language quality as perceived by native speakers) can't yet be captured in a benchmark - please let us know if you have an idea how to change this!

# Dataset

The dataset is a mixture of multi-turn chats, retrieval instructions and synthetically generated instructions spawning many topics and applications.


# Limitations & Biases

This model can produce factually incorrect and offensive output, and should not be relied on to produce factually accurate information.
This model was trained on various public datasets. While great efforts have been taken to clean the pretraining data, it is possible that this model could generate biased or otherwise offensive outputs and it is the responsibility of the user to implement a safety/moderation layer. Please use with caution.

# Acknowledgements

DiscoLM German is a [DiscoResearch](https://huggingface.co/DiscoResearch) project led by [JP Harries](https://huggingface.co/jphme) and supported by [Björn Plüster](https://huggingface.co/bjoernp) and [Daniel Auras](https://huggingface.co/rasdani).

We thank [HessianAI](https://hessian.ai/) for providing compute & support for various DiscoResearch projects and our friends at [LAION](https://laion.ai) for their work on LeoLM and scientific adivce.**

Development of DiscoLM German 7b was sponsored by **[ellamind](https://ellamind.com)**, where some of our founders are working on creating customized models for business applications with a focus on non-english language applications. Please get in contact if you need customized models for your business!


[<img src="https://raw.githubusercontent.com/OpenAccess-AI-Collective/axolotl/main/image/axolotl-badge-web.png" alt="Built with Axolotl" width="200" height="32"/>](https://github.com/OpenAccess-AI-Collective/axolotl)

# About DiscoResearch

DiscoResearch is an aspiring open research community for AI enthusiasts and LLM hackers. Come join our [Discord](https://discord.gg/ttNdas89f3), share your opinions and ideas, and advance open LLM research with us!


# Disclaimer

The license on this model does not constitute legal advice. We are not responsible for the actions of third parties who use this model. This model should only be deployed with additional safety measures in place.
