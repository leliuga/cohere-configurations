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

# Model Card: Pygmalion-2-13b-SuperCOT-weighted

This is an experimental weighted merge between:
- [Pygmalion 2 13b](https://huggingface.co/PygmalionAI/pygmalion-2-13b)
- [Ausboss's Llama2 SuperCOT loras](https://huggingface.co/ausboss/llama2-13b-supercot-loras)

Quantizations provided by us and TheBloke:
- [Exl2](https://huggingface.co/royallab/Pygmalion-2-13b-SuperCOT-weighed-exl2)
- [GPTQ](https://huggingface.co/TheBloke/Pygmalion-2-13B-SuperCOT-weighed-GPTQ)
- [GGUF](https://huggingface.co/TheBloke/Pygmalion-2-13B-SuperCOT-weighed-GGUF)

The merge was performed by a gradient merge script (apply-lora-weight-ltl.py) from [zaraki-tools](https://github.com/zarakiquemparte/zaraki-tools) by Zaraki.

Thanks to Zaraki for the inspiration and help.

This merge differs from the previous Pyg-2-SuperCOT merges. The first iteration of the SuperCOT loras were used here since it performed better than SuperCOT2.

The SuperCOT lora was merged with the following layer weights (basically 50/50. The exact ratio is 0.51)
```
0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0.5,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1
```

Here is an image to help visualize this merge. The light blue is Pygmalion-2-13b and the light green is the SuperCOT lora:
![gradient-image](https://files.catbox.moe/ndbz7t.png)

## Usage:

Since this is an experimental weight merge between Pygmalion-2 and SuperCOT, the following instruction formats should work:

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

In addition, this merge is experimental from our own testing. Your results may vary.

## Training Details

This model is merged and can be reproduced using the tools mentioned above. Please refer to all provided links for extra model-specific details.