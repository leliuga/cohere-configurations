---
license: other
license_name: stem.ai.mtl
license_link: LICENSE
language:
- en
tags:
- phi-2
- electrical engineering
- Microsoft
datasets:
- STEM-AI-mtl/Electrical-engineering
- garage-bAInd/Open-Platypus
---
# Model Card for Model ID

This is the adapters from the LoRa fine-tuning of the phi-2 model from Microsoft. It was trained on the STEM-AI-mtl/Electrical-engineering dataset combined with garage-bAInd/Open-Platypus.

- **Developed by:** STEM.AI
- **Model type:** Q&A and code generation
- **Language(s) (NLP):** English
- **Finetuned from model [optional]:** microsoft/phi-2


### Direct Use

Q&A related to electrical engineering, and Kicad software. Creation of Python code in general, and for Kicad's scripting console.

Refer to microsoft/phi-2 model card for recommended prompt format.

## Training Details

### Training Data

Dataset related to electrical engineering: STEM-AI-mtl/Electrical-engineering
It is composed of queries, 65% about general electrical engineering, 25% about Kicad (EDA software) and 10% about Python code for Kicad's scripting console.

Combined with

Dataset related to STEM and NLP: garage-bAInd/Open-Platypus

### Training Procedure 
LoRa script:  https://github.com/STEM-ai/Phi-2/raw/4eaa6aaa2679427a810ace5a061b9c951942d66a/LoRa.py

A LoRa PEFT was performed on a 48 Gb A40 Nvidia GPU.

## Model Card Authors [optional]

STEM.AI: stem.ai.mtl@gmail.com
William Harbec

### Inference example

Standard: https://github.com/STEM-ai/Phi-2/blob/4eaa6aaa2679427a810ace5a061b9c951942d66a/chat.py

GPTQ format: https://github.com/STEM-ai/Phi-2/blob/ab1ced8d7922765344d824acf1924df99606b4fc/chat-GPTQ.py