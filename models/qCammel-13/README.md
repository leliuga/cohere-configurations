---
license: other
language:
- en
pipeline_tag: text-generation
inference: false
tags:
- pytorch
- llama
- llama-2
- qCammel-13
library_name: transformers
---
# qCammel-13
qCammel-13 is a fine-tuned version of Llama-2 13B model, trained on a distilled dataset of 15,000 instructions using QLoRA. This model is optimized for academic medical knowledge and instruction-following capabilities.

## Model Details
*Note: Use of this model is governed by the Meta license. In order to download the model weights and tokenizer, please visit the [website](https://ai.meta.com/resources/models-and-libraries/llama-downloads/) and accept their License before downloading this model .*

The fine-tuning process applied to qCammel-13 involves a distilled dataset of 15,000 instructions and is trained with QLoRA, 


**Variations** The original Llama 2 has parameter sizes of 7B, 13B, and 70B. This is the fine-tuned version of the 13B model.

**Input** Models input text only.

**Output** Models generate text only.

**Model Architecture** qCammel-13 is based on the Llama 2 architecture, an auto-regressive language model that uses a decoder only transformer architecture.

**License** A custom commercial license is available at: [https://ai.meta.com/resources/models-and-libraries/llama-downloads/](https://ai.meta.com/resources/models-and-libraries/llama-downloads/)
Llama 2 is licensed under the LLAMA 2 Community License, Copyright © Meta Platforms, Inc. All Rights Reserved

**Research Papers** 
- [Clinical Camel: An Open-Source Expert-Level Medical Language Model with Dialogue-Based Knowledge Encoding](https://arxiv.org/abs/2305.12031)
- [QLoRA: Efficient Finetuning of Quantized LLMs](https://arxiv.org/abs/2305.14314)
- [LLaMA: Open and Efficient Foundation Language Models](https://arxiv.org/abs/2302.13971)

