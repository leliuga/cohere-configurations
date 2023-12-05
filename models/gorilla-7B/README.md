---
license: apache-2.0
language:
- en
tags:
- api
datasets:
- gorilla-llm/APIBench
---
# Gorilla: Large Language Model Connected with Massive APIs
By Shishir G. Patil, Tianjun Zhang, Xin Wang, and Joseph E. Gonzalez  ([Project Website](https://shishirpatil.github.io/gorilla/))

[![arXiv](https://img.shields.io/badge/arXiv-2305.15334-<COLOR>.svg?style=flat-square)](https://arxiv.org/abs/2305.15334)  [![Discord](https://img.shields.io/discord/1111172801899012102?label=Discord&logo=discord&logoColor=green&style=flat-square)](https://discord.gg/3apqwwME) [![Colab](https://colab.research.google.com/assets/colab-badge.svg)](https://colab.research.google.com/drive/1DEBPsccVLF_aUnmD0FwPeHFrtdC0QIUP?usp=sharing)

`Gorilla` enables LLMs to use tools by invoking APIs. Given a natural language query, Gorilla can write a semantically- and syntactically- correct API to invoke. With Gorilla, we are the first to demonstrate how to use LLMs to invoke 1,600+ (and growing) API calls accurately while reducing hallucination. We also release APIBench, the largest collection of APIs, curated and easy to be trained on! Join us, as we try to expand the largest API store and teach LLMs how to write them! Hop on our Discord, or open a PR, or email us if you would like to have your API incorporated as well.

## Model Details

Gorilla can be either trained via standard finetuning or using our novel retriever-aware training pipeline. We release `gorilla-7b-hf-delta-v0`, a 0-shot finetuned LLM that can reliably use Hugging Face APIs. It can be prompted through simply natural language (e.g., "I want to generate an image from text."). Checkour our website, github and paper for more information.

### Model Type

Gorilla is an open-source API caller trained by fine-tuning LLaMA weights. It is an auto-regressive language model, based on the transformer architecture.

### Model Date

05/27/2023

### Organization

Gorilla LLM (UC Berkeley)

---
license: apache-2.0
---