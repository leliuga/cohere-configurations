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

# Model Card: Pygmalion-2-13b-SuperCOT2

This is a merge between:
- [Pygmalion 2 13b](https://huggingface.co/PygmalionAI/pygmalion-2-13b)
- [Ausboss's Llama2 SuperCOT2 loras](https://huggingface.co/ausboss/llama2-13b-supercot-loras2) at a weight of 1.00.

The merge was performed by a commandline version of [EzTrainer](https://github.com/CoffeeVampir3/ez-trainer) by CoffeeVampire/Blackroot via [zaraki-tools](https://github.com/zarakiquemparte/zaraki-tools) by Zaraki.

This merge differs from the previous Pyg-2-SuperCOT merge. According to AusBoss, this version was trained closer to SuperCOT llama1. The intended objective is the same, which is to make Pygmalion smarter.

The SuperCOT2 lora was merged at a weight of 1.

## Usage:

Since this is a merge between Pygmalion-2 and SuperCOT2, the following instruction formats should work:

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