---
inference: false
language:
- en
library_name: transformers
pipeline_tag: text-generation
tags:
- llama
- llama-2
license: other
---

# Model Card: Chronohermes-Grad-L2-13b

This is a Llama 2-based model consisting of a gradient merge between:
- [Chronos 13b v2](https://huggingface.co/elinas/chronos-13b-v2)
- [Nous Hermes Llama2 13b](https://huggingface.co/NousResearch/Nous-Hermes-Llama2-13b)

Quantized Models Provided by TheBloke (Thanks!):
- [GGML](https://huggingface.co/TheBloke/Chronohermes-Grad-L2-13B-GGML)
- [GPTQ](https://huggingface.co/TheBloke/Chronohermes-Grad-L2-13B-GPTQ)


The merge was performed using [BlockMerge_Gradient](https://github.com/Gryphe/BlockMerge_Gradient) by Gryphe

The intended objective was to combine NH2's superior instruction following capabilities with the creativity and response length of Chronos v2. Merge ratios used are identical to those used in [Chronoboros Grad](https://huggingface.co/kingbri/chronoboros-grad-l2-13B), with NH2 starting with a weight of 0.9 at the 1st layer and phasing out by the 25th layer. The method is illustrated in the image below, with green representing NH2 and blue representing Chronos v2:

![hermeboros-illustration](https://files.catbox.moe/18sjej.png)

## Usage:

Intended to be prompted with the Alpaca instruction format of the base models:

```
### Instruction:
<prompt>

### Response:
<leave a newline blank for model to respond>
```

## Bias, Risks, and Limitations

The model will show biases similar to those exhibited by the base models. It is not intended for supplying factual information or advice in any form. 

## Training Details

This model is a merge. Please refer to the link repositories of the base models for details.