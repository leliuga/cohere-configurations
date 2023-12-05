---
language:
- en
library_name: transformers
pipeline_tag: text-generation
tags:
- llama
- llama-2
license: llama2
---

# Model Card: Pygmalion-2-13b-SuperCOT

This is a merge between:
- [Pygmalion 2 13b](https://huggingface.co/PygmalionAI/pygmalion-2-13b)
- [Ausboss's Llama2 SuperCOT loras](https://huggingface.co/ausboss/llama2-13b-supercot-loras) at a weight of 1.00.

Quantizations provided by us and TheBloke:
- [GGUF](https://huggingface.co/royallab/Pygmalion-2-13b-SuperCOT-GGUF)
- [GGUF (TheBloke)](https://huggingface.co/TheBloke/Pygmalion-2-13B-SuperCOT-GGUF)
- [GPTQ](https://huggingface.co/TheBloke/Pygmalion-2-13B-SuperCOT-GPTQ)

The merge was performed by a commandline version of [EzTrainer](https://github.com/CoffeeVampir3/ez-trainer) by CoffeeVampire/Blackroot via [zaraki-tools](https://github.com/CoffeeVampir3/ez-trainer) by Zaraki.

The intended objective is to make Pygmalion-2 smarter and try to make it drift off less.

The SuperCOT lora was merged at a weight of 1.

## Usage:

Since this is a merge between Pygmalion-2 and SuperCOT, the following instruction formats should work:

Metharme:

```
<|system|>This is a text adventure game. Describe the scenario to the user and give him three options to pick from on each turn.<|user|>Start!<|model|>
```

Alpaca:

```
### Instruction:
Your instruction or question here.
### Response:
```

## Bias, Risks, and Limitations

The model will show biases similar to those observed in niche roleplaying forums on the Internet, besides those exhibited by the base model. It is not intended for supplying factual information or advice in any form.

## Training Details

This model is merged and can be reproduced using the tools mentioned above. Please refer to all provided links for extra model-specific details.