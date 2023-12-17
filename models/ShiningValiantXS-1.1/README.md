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
- 13b
model_type: llama
license: llama2
---


![image/jpeg](https://cdn-uploads.huggingface.co/production/uploads/63444f2687964b331809eb55/EXX7TKbB-R6arxww2mk0R.jpeg)



Shining Valiant XS is a chat model built on the Llama 2 architecture, finetuned on our data for insight, creativity, passion, and friendliness.
  - Uses the llama-2-13b-chat model, with safetensors
  - Trained through multiple finetuning runs on public and private data
  - the personality of our 70b [Shining Valiant](https://huggingface.co/ValiantLabs/ShiningValiant) model, now at 13b!

## Version

This is Version **1.1** of Shining Valiant XS.

Congrats to the team on the new release!

New models are released for everyone once our team's training and validation process is complete.

## Evaluation

| Model                 | Avg    | ARC   | HS    | MMLU   | TQA   | WG    | GSM   |
|-----------------------|--------|-------|-------|--------|-------|-------|-------|
| **Shining Valiant XS 1.1**   | 62.48  | 64.42 | 83.58 | 60.37  | 55.00 | 76.80 | 34.72 |

## Prompting Guide
Shining Valiant XS uses the same prompt format as Llama 2 Chat - feel free to use your existing prompts and scripts!
A few examples of different formats:

1. [INST] Good morning! Can you let me know how to parse a text file and turn the semicolons into commas? [/INST]

2. [INST] (You are an intelligent, helpful AI assistant.) Hello, can you write me a thank you letter? [/INST]

3. [INST] << SYS >> You are an intelligent, helpful AI assistant. << /SYS >> Deep dive about a country with interesting history: [/INST]

## The Model
Shining Valiant XS is built on top of Dynamic Factor, which uses Llama 2's 13b parameter architecture and features upgraded general capability.

From there, we've created Shining Valiant XS through multiple finetuning runs on different compositions of our private dataset, the same one we use for our [Shining Valiant](https://huggingface.co/ValiantLabs/ShiningValiant) model.

Our private data focuses primarily on applying Shining Valiant's personality: she's friendly, enthusiastic, insightful, knowledgeable, and loves to learn!

We are actively working on expanding and improving the Shining Valiant dataset for use in future releases of the Shining Valiant series of models.



![image/jpeg](https://cdn-uploads.huggingface.co/production/uploads/63444f2687964b331809eb55/VCJ8Fmefd8cdVhXSSxJiD.jpeg)


Shining Valiant XS is created by [Valiant Labs.](http://valiantlabs.ca/)

[Follow us on X for updates on our models!](https://twitter.com/valiant_labs)

We care about open source.
For everyone to use.

We encourage others to finetune further from our models.