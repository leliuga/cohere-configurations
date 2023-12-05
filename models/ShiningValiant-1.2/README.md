---
language:
- en
pipeline_tag: text-generation
tags:
- shining-valiant
- valiant
- valiant-labs
- llama
- llama-2
- llama-2-chat
- 70b
model_type: llama
license: llama2
---


![image/jpeg](https://cdn-uploads.huggingface.co/production/uploads/64f267a8a4f79a118e0fcc89/5rUJPhu_6LyDvSQogSVhk.jpeg)


Shining Valiant is a chat model built on the Llama 2 architecture, finetuned on our data for insight, creativity, passion, and friendliness.
  - Uses the llama-2-70b-chat model, with safetensors
  - Finetuned on multiple runs across private and public data
  - Data focused on knowledge, enthusiasm, and structured reasoning

## Version

The current version is **1.3!**

We're thrilled to bring you our newest release!

Previous versions remain available in the repository. New models will be released for everyone once our team's training and validation process is complete.

## Evaluation

| Model                 | Avg    | ARC   | HS    | MMLU   | TQA   |
|-----------------------|--------|-------|-------|--------|-------|
| **Shining Valiant 1.2**   | 74.17  | 72.95 | 87.88 | 70.97  | 64.88 |
| Llama 2               | 67.35  | 67.32 | 87.33 | 69.83  | 44.92 |
| Llama 2 Chat          | 66.80  | 64.59 | 85.88 | 63.91  | 52.80 |

**Shining Valiant 1.3** is awaiting full results from the Open LLM Leaderboard.

SV 1.3 outperformed SV 1.2 on our internal testing.

## Prompting Guide
Shining Valiant uses the same prompt format as Llama 2 Chat - feel free to use your existing prompts and scripts!
A few examples of different formats:

1. [INST] Good morning! Can you let me know how to parse a text file and turn the semicolons into commas? [/INST]

2. [INST] (You are an intelligent, helpful AI assistant.) Hello, can you write me a thank you letter? [/INST]

3. [INST] << SYS >>You are an intelligent, helpful AI assistant.<< /SYS >>Deep dive about a country with interesting history: [/INST]

## The Model
Shining Valiant is built on top of Sunset Boulevard, which uses Llama 2's 70b parameter architecture and features upgraded general capability.

From there, we've created Shining Valiant through multiple finetuning runs on different compositions of our private dataset. 

Our private data focuses primarily on applying Shining Valiant's personality: she's friendly, enthusiastic, insightful, knowledgeable, and loves to learn!

We are actively working on expanding and improving the Shining Valiant dataset for use in future releases of this model and others.



![image/jpeg](https://cdn-uploads.huggingface.co/production/uploads/63444f2687964b331809eb55/VCJ8Fmefd8cdVhXSSxJiD.jpeg)


Shining Valiant is created by [Valiant Labs.](http://valiantlabs.ca/)

[Follow us on X for updates on our models!](https://twitter.com/valiant_labs)

We care about open source.
For everyone to use.

We encourage others to finetune further from our models.