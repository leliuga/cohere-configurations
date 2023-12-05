---
language:
- en
pipeline_tag: text-classification
tags:
- llama-2
---
This is a Llama-2 version of [Guanaco](https://huggingface.co/timdettmers/guanaco-13b). It was finetuned from the base [Llama-13b](https://huggingface.co/meta-llama/Llama-2-13b-hf) model using the official training scripts found in the [QLoRA repo](https://github.com/artidoro/qlora). I wanted it to be as faithful as possible and therefore changed nothing in the training script beyond the model it was pointing to. The model prompt is therefore also the same as the original Guanaco model.

This repo contains the merged f16 model. The QLoRA adaptor can be found [here](https://huggingface.co/Mikael110/llama-2-13b-guanaco-qlora).

A 7b version of the model can be found [here](https://huggingface.co/Mikael110/llama-2-7b-guanaco-fp16).

**Legal Disclaimer: This model is bound by the usage restrictions of the original Llama-2 model. And comes with no warranty or gurantees of any kind.**